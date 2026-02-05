#!/usr/bin/env python3
"""
Sync the Zenodo concept DOI into repository metadata files.

This script is meant to be run AFTER Zenodo has minted a DOI for a GitHub release.
It will:
  - discover the repository's Zenodo concept DOI via the public Zenodo API
  - write it into CITATION.cff, dataset.jsonld, .zenodo.json, and the HF dataset card
  - regenerate per-paper JSON-LD + docs site + sitemap(s)

It does not require a Zenodo API token (uses the public records API).
"""

from __future__ import annotations

import argparse
import json
import re
import subprocess
import sys
import urllib.parse
import urllib.request
from dataclasses import dataclass
from pathlib import Path
from typing import Any


ZENODO_API_DEFAULT = "https://zenodo.org/api/records"
REPO_URL_DEFAULT = "https://github.com/yonathanarbel/my-works-for-llm"


@dataclass(frozen=True)
class ZenodoHit:
    doi: str | None
    conceptdoi: str | None
    links_html: str | None
    title: str | None
    created: str | None
    updated: str | None


def _http_get_json(url: str, *, timeout_s: int = 30) -> dict[str, Any]:
    req = urllib.request.Request(url, headers={"User-Agent": "my-works-for-llm DOI sync"})
    with urllib.request.urlopen(req, timeout=timeout_s) as resp:
        raw = resp.read()
    return json.loads(raw.decode("utf-8"))


def _extract_repo_urls(obj: Any) -> set[str]:
    urls: set[str] = set()
    if isinstance(obj, dict):
        for v in obj.values():
            urls |= _extract_repo_urls(v)
    elif isinstance(obj, list):
        for v in obj:
            urls |= _extract_repo_urls(v)
    elif isinstance(obj, str):
        if "github.com/" in obj:
            urls.add(obj)
    return urls


def _hit_from_record(record: dict[str, Any]) -> ZenodoHit:
    metadata = record.get("metadata") if isinstance(record.get("metadata"), dict) else {}
    title = metadata.get("title") if isinstance(metadata.get("title"), str) else None

    links = record.get("links") if isinstance(record.get("links"), dict) else {}
    html_link = links.get("html") if isinstance(links.get("html"), str) else None

    return ZenodoHit(
        doi=record.get("doi") if isinstance(record.get("doi"), str) else None,
        conceptdoi=record.get("conceptdoi") if isinstance(record.get("conceptdoi"), str) else None,
        links_html=html_link,
        title=title,
        created=record.get("created") if isinstance(record.get("created"), str) else None,
        updated=record.get("updated") if isinstance(record.get("updated"), str) else None,
    )


def find_concept_doi(*, zenodo_api: str, repo_url: str) -> str | None:
    # Strategy:
    # 1) Search for the repo URL as a substring (works for Zenodo GitHub integration).
    # 2) Filter hits whose metadata includes the exact repo_url in related identifiers / references.
    # 3) Prefer newest hit and return its conceptdoi (fallback to doi).
    q = urllib.parse.quote(repo_url, safe="")
    url = f"{zenodo_api.rstrip('/')}/?q={q}&size=50&sort=mostrecent"
    data = _http_get_json(url)

    hits = data.get("hits", {}).get("hits", [])
    if not isinstance(hits, list) or not hits:
        return None

    filtered: list[tuple[ZenodoHit, dict[str, Any]]] = []
    for rec in hits:
        if not isinstance(rec, dict):
            continue
        all_urls = _extract_repo_urls(rec)
        if repo_url in all_urls:
            filtered.append((_hit_from_record(rec), rec))

    if not filtered:
        # Best effort: take the first hit if it at least looks like our repo.
        first = hits[0] if isinstance(hits[0], dict) else None
        if not first:
            return None
        h = _hit_from_record(first)
        return h.conceptdoi or h.doi

    # Sort by updated/created strings (ISO) descending.
    filtered.sort(key=lambda t: (t[0].updated or "", t[0].created or ""), reverse=True)
    best = filtered[0][0]
    return best.conceptdoi or best.doi


def _read_text(path: Path) -> str:
    return path.read_text(encoding="utf-8", errors="replace")


def _write_text(path: Path, text: str) -> None:
    path.write_text(text, encoding="utf-8", newline="\n")


def _normalize_doi(doi: str) -> str:
    doi = doi.strip()
    doi = re.sub(r"^https?://(dx\\.)?doi\\.org/", "", doi, flags=re.IGNORECASE)
    return doi


def update_citation_cff(path: Path, doi: str) -> bool:
    src = _read_text(path)
    doi = _normalize_doi(doi)

    # Replace existing `doi:` if present; otherwise insert near repository URL.
    if re.search(r"(?m)^doi:\\s*", src):
        out = re.sub(r"(?m)^doi:\\s*.*$", f'doi: "{doi}"', src)
    else:
        lines = src.splitlines()
        insert_at = None
        for i, line in enumerate(lines):
            if line.startswith("repository-code:") or line.startswith("url:"):
                insert_at = i + 1
        if insert_at is None:
            insert_at = len(lines)
        lines.insert(insert_at, f'doi: "{doi}"')
        out = "\n".join(lines) + "\n"

    if out == src:
        return False
    _write_text(path, out if out.endswith("\n") else out + "\n")
    return True


def update_dataset_jsonld(path: Path, doi: str) -> bool:
    doi = _normalize_doi(doi)
    src = path.read_text(encoding="utf-8")
    data = json.loads(src)
    if not isinstance(data, dict):
        raise SystemExit(f"Unexpected JSON-LD root type in {path}")

    data["identifier"] = {"@type": "PropertyValue", "propertyID": "DOI", "value": doi}
    data["sameAs"] = [f"https://doi.org/{doi}"]

    out = json.dumps(data, ensure_ascii=False, indent=2) + "\n"
    if out == src:
        return False
    path.write_text(out, encoding="utf-8", newline="\n")
    return True


def update_zenodo_json(path: Path, doi: str) -> bool:
    doi = _normalize_doi(doi)
    src = path.read_text(encoding="utf-8")
    data = json.loads(src)
    if not isinstance(data, dict):
        raise SystemExit(f"Unexpected .zenodo.json root type: {path}")

    description = data.get("description") if isinstance(data.get("description"), str) else ""
    needle_re = re.compile(r"(?im)^\\s*Zenodo concept DOI:\\s*https?://doi\\.org/\\S+\\s*$")

    line = f"Zenodo concept DOI: https://doi.org/{doi}"
    if needle_re.search(description):
        description2 = needle_re.sub(line, description).strip()
    else:
        description2 = (description.strip() + "\n\n" + line).strip() if description.strip() else line

    if description2 != description:
        data["description"] = description2
        out = json.dumps(data, ensure_ascii=False, indent=2) + "\n"
        path.write_text(out, encoding="utf-8", newline="\n")
        return True
    return False


def update_hf_dataset_card(path: Path, doi: str) -> bool:
    doi = _normalize_doi(doi)
    src = _read_text(path)
    doi_url = f"https://doi.org/{doi}"

    # Ensure a DOI line exists under "## Citation".
    if "## Citation" not in src:
        out = src.rstrip("\n") + "\n\n## Citation\n\n- Zenodo DOI: " + doi_url + "\n"
    else:
        lines = src.splitlines()
        out_lines: list[str] = []
        in_citation = False
        inserted = False
        for i, line in enumerate(lines):
            if line.strip() == "## Citation":
                in_citation = True
                out_lines.append(line)
                continue
            if in_citation and line.startswith("## "):
                if not inserted:
                    out_lines.append(f"- Zenodo DOI: {doi_url}")
                    inserted = True
                in_citation = False
            if in_citation and re.match(r"^\\s*-\\s*Zenodo DOI:\\s*", line):
                out_lines.append(f"- Zenodo DOI: {doi_url}")
                inserted = True
                continue
            out_lines.append(line)
        if in_citation and not inserted:
            out_lines.append(f"- Zenodo DOI: {doi_url}")
        out = "\n".join(out_lines) + "\n"

    if out == src:
        return False
    _write_text(path, out)
    return True


def _run(cmd: list[str]) -> None:
    proc = subprocess.run(cmd, check=False)
    if proc.returncode != 0:
        raise SystemExit(proc.returncode)


def regenerate_outputs(*, dataset_doi: str | None) -> None:
    cmd = [sys.executable, "generate_paper_jsonld.py"]
    if dataset_doi:
        cmd.extend(["--dataset-doi", dataset_doi])
    _run(cmd)
    _run([sys.executable, "generate_docs_site.py"])
    _run([sys.executable, "generate_sitemap.py"])


def main(argv: list[str] | None = None) -> int:
    parser = argparse.ArgumentParser(description="Sync Zenodo DOI into repo metadata and regenerate derived outputs.")
    parser.add_argument("--repo-url", default=REPO_URL_DEFAULT, help="GitHub repository URL to search for.")
    parser.add_argument("--zenodo-api", default=ZENODO_API_DEFAULT, help="Zenodo records API base.")
    parser.add_argument("--doi", help="Override: explicitly set DOI (skips Zenodo API lookup).")
    parser.add_argument(
        "--require",
        action="store_true",
        help="Fail (non-zero exit) if a DOI could not be discovered.",
    )
    args = parser.parse_args(argv)

    doi = args.doi
    if doi:
        doi = _normalize_doi(doi)
    else:
        doi = find_concept_doi(zenodo_api=args.zenodo_api, repo_url=args.repo_url)
        if doi:
            doi = _normalize_doi(doi)

    if not doi:
        msg = "Zenodo DOI not found yet. Create a GitHub release and wait for Zenodo to archive it, then rerun."
        if args.require:
            print(msg, file=sys.stderr)
            return 2
        print(msg)
        return 0

    changed = False
    repo_root = Path(".")
    citation = repo_root / "CITATION.cff"
    dataset_jsonld = repo_root / "dataset.jsonld"
    zenodo_json = repo_root / ".zenodo.json"
    hf_card = repo_root / "huggingface" / "DATASET_CARD.md"

    if citation.exists():
        changed |= update_citation_cff(citation, doi)
    if dataset_jsonld.exists():
        changed |= update_dataset_jsonld(dataset_jsonld, doi)
    if zenodo_json.exists():
        changed |= update_zenodo_json(zenodo_json, doi)
    if hf_card.exists():
        changed |= update_hf_dataset_card(hf_card, doi)

    regenerate_outputs(dataset_doi=doi)

    if changed:
        print(f"Updated DOI to {doi}.")
    else:
        print(f"DOI already up to date ({doi}).")

    return 0


if __name__ == "__main__":
    raise SystemExit(main())

