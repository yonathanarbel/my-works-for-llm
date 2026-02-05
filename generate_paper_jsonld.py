#!/usr/bin/env python3
"""
Generate per-paper Schema.org JSON-LD (ScholarlyArticle) files.

Outputs `papers/<paper_id>/scholarlyarticle.jsonld`.
"""

from __future__ import annotations

import argparse
import json
import re
from dataclasses import dataclass
from pathlib import Path
from typing import Any


RAW_DEFAULT = "https://raw.githubusercontent.com/yonathanarbel/my-works-for-llm/main/"
REPO_DEFAULT = "https://github.com/yonathanarbel/my-works-for-llm"
PAGES_DEFAULT = "https://yonathanarbel.github.io/my-works-for-llm/"


def _read_text(path: Path) -> str:
    return path.read_text(encoding="utf-8", errors="replace")


def _read_json(path: Path) -> dict[str, Any]:
    return json.loads(path.read_text(encoding="utf-8"))


def _normalize_doi(value: str) -> str:
    value = value.strip()
    value = re.sub(r"^https?://(dx\.)?doi\.org/", "", value, flags=re.IGNORECASE)
    return value


def _infer_dataset_doi(repo_root: Path) -> str | None:
    # Prefer CITATION.cff if it has a DOI.
    citation = repo_root / "CITATION.cff"
    if citation.exists():
        m = re.search(r'(?m)^doi:\s*"?([^"\s]+)"?\s*$', _read_text(citation))
        if m:
            return _normalize_doi(m.group(1))

    dataset_jsonld = repo_root / "dataset.jsonld"
    if dataset_jsonld.exists():
        try:
            data = _read_json(dataset_jsonld)
        except Exception:
            data = {}
        if isinstance(data, dict):
            ident = data.get("identifier")
            if isinstance(ident, dict):
                v = ident.get("value")
                if isinstance(v, str) and v.strip():
                    return _normalize_doi(v)
    return None


def _license_to_url(license_value: str | None) -> str | None:
    if not license_value:
        return None
    mapping = {
        "CC-BY-NC-4.0": "https://creativecommons.org/licenses/by-nc/4.0/",
        "CC0-1.0": "https://creativecommons.org/publicdomain/zero/1.0/",
    }
    return mapping.get(license_value, license_value)


def _infer_ssrn_id(paper_id: str) -> int | None:
    m = re.fullmatch(r"ssrn-(\d+)", paper_id)
    return int(m.group(1)) if m else None


def _infer_ssrn_url(ssrn_id: int | None) -> str | None:
    if not ssrn_id:
        return None
    return f"https://papers.ssrn.com/sol3/papers.cfm?abstract_id={ssrn_id}"


def _extract_title(paper_dir: Path, metadata: dict[str, Any], paper_id: str) -> str:
    title = metadata.get("title")
    if isinstance(title, str) and title and title != paper_id:
        return title

    paper_txt = paper_dir / "paper.txt"
    if paper_txt.exists():
        for line in _read_text(paper_txt).splitlines():
            candidate = line.strip()
            if candidate:
                return candidate

    return str(title) if title else paper_id


def _extract_tldr_from_summary(summary_text: str) -> str | None:
    # Expect the repo's common structure:
    # "## TL;DR ..." followed by an indented paragraph.
    lines = summary_text.splitlines()
    tldr_start = None
    for idx, line in enumerate(lines):
        if "TL;DR" in line:
            tldr_start = idx
            break
    if tldr_start is None:
        return None

    collected: list[str] = []
    for line in lines[tldr_start + 1 :]:
        stripped = line.strip()
        if not stripped:
            if collected:
                break
            continue
        # Stop on the next numbered section header.
        if re.match(r"^\d+\.\s+##\s+", stripped):
            break
        # Skip markdown bullets/headings.
        if stripped.startswith("##") or stripped.startswith("*"):
            continue
        collected.append(stripped)
        if len(" ".join(collected)) >= 900:
            break

    text = " ".join(collected).strip()
    return text or None


def _extract_abstract(paper_dir: Path, metadata: dict[str, Any]) -> str | None:
    abstract = metadata.get("abstract")
    if isinstance(abstract, str) and abstract.strip():
        return abstract.strip()

    paper_txt = paper_dir / "paper.txt"
    if paper_txt.exists():
        lines = _read_text(paper_txt).splitlines()
        for i, line in enumerate(lines):
            if line.strip().upper() == "ABSTRACT":
                chunks: list[str] = []
                for j in range(i + 1, len(lines)):
                    candidate = lines[j].strip()
                    if not candidate:
                        if chunks:
                            break
                        continue
                    upper = candidate.upper()
                    if upper in {"INTRODUCTION", "TABLE OF CONTENTS"}:
                        if chunks:
                            break
                    if re.match(r"^[IVXLC]+\\.", candidate):
                        if chunks:
                            break
                    chunks.append(candidate)
                    if len(" ".join(chunks)) >= 1200:
                        break
                if chunks:
                    return " ".join(chunks).strip()

    summary_path = paper_dir / "summary.md"
    if summary_path.exists():
        tldr = _extract_tldr_from_summary(_read_text(summary_path))
        if tldr:
            return tldr

    return None


def _authors_to_schema(authors: Any) -> list[dict[str, Any]] | None:
    if not authors:
        return None
    if not isinstance(authors, list):
        return None

    people: list[dict[str, Any]] = []
    for a in authors:
        if isinstance(a, dict):
            given = a.get("given") or a.get("given-names") or a.get("given_names")
            family = a.get("family") or a.get("family-names") or a.get("family_names")
            name = " ".join(x for x in [given, family] if isinstance(x, str) and x.strip()).strip()
            if not name:
                name = a.get("name")
            person: dict[str, Any] = {"@type": "Person", "name": name} if name else {"@type": "Person"}
            aff = a.get("affiliation")
            if isinstance(aff, str) and aff.strip():
                person["affiliation"] = {"@type": "Organization", "name": aff.strip()}
            people.append(person)
        elif isinstance(a, str) and a.strip():
            people.append({"@type": "Person", "name": a.strip()})

    return people or None


def _build_media_objects(raw_base: str, paper_dir: Path, papers_dir: Path) -> list[dict[str, Any]]:
    raw_base = raw_base.rstrip("/") + "/"

    media: list[dict[str, Any]] = []
    for filename, fmt in [
        ("summary.md", "text/markdown"),
        ("summary.zh.md", "text/markdown"),
        ("one_pager.md", "text/markdown"),
        ("study_pack.md", "text/markdown"),
        ("paper.txt", "text/plain"),
        ("paper.pdf", "application/pdf"),
    ]:
        path = paper_dir / filename
        if not path.exists():
            continue
        rel = str(path.relative_to(papers_dir.parent)).replace("\\", "/")
        media.append(
            {
                "@type": "MediaObject",
                "contentUrl": raw_base + rel,
                "encodingFormat": fmt,
                "name": rel,
            }
        )
    return media


def generate_for_paper(
    *,
    paper_dir: Path,
    papers_dir: Path,
    raw_base: str,
    repo_url: str,
    pages_base: str | None,
    dataset_name: str,
    dataset_url: str,
    dataset_doi: str | None,
) -> dict[str, Any]:
    paper_id = paper_dir.name

    metadata_path = paper_dir / "metadata.json"
    metadata: dict[str, Any] = _read_json(metadata_path) if metadata_path.exists() else {"id": paper_id}

    title = _extract_title(paper_dir, metadata, paper_id)
    abstract = _extract_abstract(paper_dir, metadata)

    ssrn_id = metadata.get("ssrn_id")
    if not isinstance(ssrn_id, int):
        ssrn_id = _infer_ssrn_id(paper_id)
    ssrn_url = metadata.get("ssrn_url")
    if not isinstance(ssrn_url, str):
        ssrn_url = _infer_ssrn_url(ssrn_id)

    raw_base = raw_base.rstrip("/") + "/"
    rel_summary = f"papers/{paper_id}/summary.md"
    canonical_url = raw_base + rel_summary
    if pages_base:
        canonical_url = pages_base.rstrip("/") + f"/papers/{paper_id}/"

    dataset_obj: dict[str, Any] = {"@type": "Dataset", "name": dataset_name, "url": dataset_url}
    if isinstance(dataset_doi, str) and dataset_doi.strip():
        dataset_doi = _normalize_doi(dataset_doi)
        dataset_obj["identifier"] = {"@type": "PropertyValue", "propertyID": "DOI", "value": dataset_doi}
        dataset_obj["sameAs"] = [f"https://doi.org/{dataset_doi}"]

    schema: dict[str, Any] = {
        "@context": "https://schema.org",
        "@type": "ScholarlyArticle",
        "name": title,
        "identifier": [
            {"@type": "PropertyValue", "propertyID": "paper_id", "value": paper_id},
            {"@type": "PropertyValue", "propertyID": "SSRN", "value": str(ssrn_id)} if ssrn_id else None,
        ],
        "url": canonical_url,
        "isPartOf": dataset_obj,
    }

    schema["identifier"] = [x for x in schema["identifier"] if x is not None]

    if ssrn_url:
        schema["sameAs"] = [ssrn_url]

    year = metadata.get("year")
    if isinstance(year, int):
        schema["datePublished"] = str(year)

    authors_schema = _authors_to_schema(metadata.get("authors"))
    if authors_schema:
        schema["author"] = authors_schema

    if abstract:
        schema["abstract"] = abstract

    license_url = _license_to_url(metadata.get("license_pdf") or metadata.get("license"))
    if license_url:
        schema["license"] = license_url

    keywords = metadata.get("keywords")
    if isinstance(keywords, list) and all(isinstance(x, str) for x in keywords):
        schema["keywords"] = keywords

    media = _build_media_objects(raw_base, paper_dir, papers_dir)
    if media:
        schema["encoding"] = media

    schema["mainEntityOfPage"] = {
        "@type": "CreativeWork",
        "url": repo_url.rstrip("/") + f"/tree/main/papers/{paper_id}",
    }

    return schema


def main() -> int:
    parser = argparse.ArgumentParser(description="Generate per-paper Schema.org JSON-LD files.")
    parser.add_argument("--papers-dir", default="papers", help="Papers directory (default: papers).")
    parser.add_argument("--raw-base", default=RAW_DEFAULT, help="Raw content base URL.")
    parser.add_argument("--repo-url", default=REPO_DEFAULT, help="Repository URL.")
    parser.add_argument("--pages-base", default=PAGES_DEFAULT, help="GitHub Pages base URL (set empty to disable).")
    parser.add_argument("--dataset-name", default="my-works-for-llm", help="Dataset name for isPartOf.")
    parser.add_argument("--dataset-url", default=REPO_DEFAULT, help="Dataset URL for isPartOf.")
    parser.add_argument(
        "--dataset-doi",
        default="",
        help="Optional dataset DOI to include in isPartOf (falls back to CITATION.cff/dataset.jsonld if present).",
    )
    args = parser.parse_args()

    dataset_doi = args.dataset_doi.strip() if isinstance(args.dataset_doi, str) else ""
    if not dataset_doi:
        dataset_doi = _infer_dataset_doi(Path("."))

    papers_dir = Path(args.papers_dir)
    if not papers_dir.exists():
        raise SystemExit(f"Papers directory not found: {papers_dir}")

    pages_base = args.pages_base.strip() if isinstance(args.pages_base, str) else ""
    pages_base = pages_base if pages_base else None

    for paper_dir in sorted(papers_dir.iterdir()):
        if not paper_dir.is_dir():
            continue
        schema = generate_for_paper(
            paper_dir=paper_dir,
            papers_dir=papers_dir,
            raw_base=args.raw_base,
            repo_url=args.repo_url,
            pages_base=pages_base,
            dataset_name=args.dataset_name,
            dataset_url=args.dataset_url,
            dataset_doi=dataset_doi,
        )
        out_path = paper_dir / "scholarlyarticle.jsonld"
        out_path.write_text(json.dumps(schema, ensure_ascii=False, indent=2) + "\n", encoding="utf-8", newline="\n")

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
