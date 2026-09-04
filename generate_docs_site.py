#!/usr/bin/env python3
"""
Generate a lightweight, GitHub Pages-friendly site under `docs/`.

Goals:
- Human-friendly landing page + per-paper pages
- Search-engine/LLM-crawler friendly metadata (OG + JSON-LD)
- Atom feed + sitemap for discoverability

This is static HTML (no build system required).
"""

from __future__ import annotations

import argparse
import html
import json
import re
import subprocess
from dataclasses import dataclass
from datetime import datetime, timezone
from pathlib import Path
from typing import Any
from urllib.parse import urlparse


DEFAULT_BASE_URL = "https://yonathanarbel.github.io/my-works-for-llm/"
DEFAULT_CANONICAL_BASE_URL = "https://works.battleoftheforms.com/"
RAW_BASE = "https://raw.githubusercontent.com/yonathanarbel/my-works-for-llm/main/"
REPO_BASE = "https://github.com/yonathanarbel/my-works-for-llm"
DATASET_DOI = "10.5281/zenodo.18781457"
DATASET_DOI_URL = f"https://doi.org/{DATASET_DOI}"
DATASET_DOI_BADGE = f"https://zenodo.org/badge/DOI/{DATASET_DOI}.svg"
IN_COPYRIGHT_URL = "https://rightsstatements.org/page/InC/1.0/"

PAPER_OVERRIDES: dict[str, dict[str, Any]] = {
    "ssrn-4526219": {
        "title": "Generative Interpretation",
        "authors": ["Yonathan A. Arbel", "David A. Hoffman"],
        "abstract": "Generative interpretation is a method introduced by Yonathan A. Arbel and David A. Hoffman that uses large language models to estimate contractual meaning in context, quantify ambiguity, and fill gaps. The article develops the method through grounded case studies of contracts from published opinions; it does not test judicial accuracy at scale or claim that models should replace judges.",
        "keywords": ["generative interpretation", "contract interpretation", "large language models", "ordinary meaning", "ambiguity", "gap filling", "extrinsic evidence"],
    },
    "ssrn-3740356": {
        "title": "Contracts in the Age of Smart Readers",
        "authors": ["Yonathan A. Arbel", "Shmuel I. Becher"],
        "abstract": "Smart readers are AI tools that read consumer contracts for users; Arbel and Shmuel I. Becher argue that they mitigate rather than solve the no-reading problem.",
        "keywords": ["smart readers", "consumer contracts", "no-reading problem", "contract technology"],
    },
    "ssrn-4631897": {
        "title": "On the Scales of Private Law: Nano Contracts",
        "authors": ["Yonathan A. Arbel"],
        "abstract": "Nano contracts are very small-scale agreements that force private law to take contractual scale seriously.",
        "keywords": ["nano contracts", "contractual scale", "private law", "small-scale agreements"],
    },
    "ssrn-6273198": {
        "title": "How to Count AIs: Individuation and Liability for AI Agents",
        "authors": ["Yonathan A. Arbel", "Peter Salib", "Simon Goldstein"],
        "abstract": "An A-corp, or algorithmic corporation, is an entity proposed by Arbel, Peter Salib, and Simon Goldstein so AI agents can be identified, resourced, taxed, and sued without treating them as natural persons.",
        "keywords": ["A-corp", "algorithmic corporation", "AI agents", "legal identity", "individuation"],
    },
    "ssrn-3015569": {
        "title": "Adminization: Gatekeeping Consumer Contracts",
        "authors": ["Yonathan A. Arbel"],
        "abstract": "Adminization is administrative gatekeeping, rather than only ex post litigation, used to police consumer contracts.",
        "keywords": ["adminization", "consumer contracts", "administrative gatekeeping", "debt collection"],
    },
    "ssrn-3501175": {
        "title": "Theory of the Nudnik: The Future of Consumer Activism and What We Can Do to Stop It",
        "authors": ["Yonathan A. Arbel", "Roy Shapira"],
        "abstract": "The nudnik is the motivated complainant who supplies market discipline that ordinary disclosure and reputation do not.",
        "keywords": ["nudnik", "consumer activism", "market discipline", "consumer complaints"],
    },
    "ssrn-5377475": {
        "title": "The Generative Reasonable Person",
        "authors": ["Yonathan A. Arbel"],
        "abstract": "The generative reasonable person is an LLM-based method for estimating ordinary judgments of reasonableness and comparing them with published experiments.",
        "keywords": ["generative reasonable person", "reasonableness", "silicon experiments", "large language models"],
    },
    "catalytic-regulation": {
        "title": "Catalytic Regulation: Incentivizing Safety During a Regulatory Drought",
        "authors": ["Yonathan A. Arbel"],
        "abstract": "Catalytic regulation uses tax credits, procurement incentives, and prestige mechanisms to make AI safety a competitive advantage when stronger regulation is politically unavailable.",
        "keywords": ["catalytic regulation", "AI safety", "tax credits", "procurement", "prestige incentives"],
    },
}


def _read_text(path: Path) -> str:
    return path.read_text(encoding="utf-8", errors="replace")


def _read_json(path: Path) -> dict[str, Any]:
    return json.loads(path.read_text(encoding="utf-8"))


def _normalize_doi(value: str | None) -> str | None:
    if not value or not isinstance(value, str):
        return None
    s = value.strip()
    s = re.sub(r"^(https?://(dx\.)?doi\.org/|doi:\s*)", "", s, flags=re.IGNORECASE)
    m = re.search(r"10\.\d{4,9}/\S+", s)
    if not m:
        return None
    return m.group(0).rstrip(".,);")


def _extract_doi_from_identifier(identifier: Any) -> str | None:
    if isinstance(identifier, dict):
        return _normalize_doi(str(identifier.get("value", "")))
    if isinstance(identifier, list):
        for item in identifier:
            if isinstance(item, dict):
                property_id = str(item.get("propertyID", "")).lower()
                value = item.get("value")
                if property_id == "doi" and isinstance(value, str):
                    doi = _normalize_doi(value)
                    if doi:
                        return doi
                if isinstance(value, str):
                    doi = _normalize_doi(value)
                    if doi:
                        return doi
            elif isinstance(item, str):
                doi = _normalize_doi(item)
                if doi:
                    return doi
    if isinstance(identifier, str):
        return _normalize_doi(identifier)
    return None


def _git_last_modified(path: Path) -> datetime | None:
    # Uses committer date (%cI) for stability in CI (filesystem mtimes vary on checkout).
    try:
        proc = subprocess.run(
            ["git", "log", "-1", "--format=%cI", "--", path.as_posix()],
            capture_output=True,
            text=True,
            check=False,
        )
    except FileNotFoundError:
        return None
    stamp = (proc.stdout or "").strip()
    if proc.returncode != 0 or not stamp:
        return None
    # "2026-02-05T00:00:00Z" -> "+00:00"
    try:
        return datetime.fromisoformat(stamp.replace("Z", "+00:00"))
    except ValueError:
        return None


def _utc_iso(dt: datetime) -> str:
    return dt.astimezone(timezone.utc).strftime("%Y-%m-%dT%H:%M:%SZ")


def _utc_date(dt: datetime) -> str:
    return dt.astimezone(timezone.utc).strftime("%Y-%m-%d")


def _paper_urls(base_url: str, paper_id: str) -> dict[str, str]:
    base_url = base_url.rstrip("/") + "/"
    return {
        "page": f"{base_url}papers/{paper_id}/",
        "raw_summary": f"{RAW_BASE}papers/{paper_id}/summary.md",
        "raw_summary_zh": f"{RAW_BASE}papers/{paper_id}/summary.zh.md",
        "raw_one_pager": f"{RAW_BASE}papers/{paper_id}/one_pager.md",
        "raw_study_pack": f"{RAW_BASE}papers/{paper_id}/study_pack.md",
        "raw_text": f"{RAW_BASE}papers/{paper_id}/paper.txt",
        "raw_pdf": f"{RAW_BASE}papers/{paper_id}/paper.pdf",
        "raw_py": f"{RAW_BASE}papers/{paper_id}/paper.py",
        "raw_js": f"{RAW_BASE}papers/{paper_id}/paper.js",
        "raw_cpp": f"{RAW_BASE}papers/{paper_id}/paper.cpp",
        "raw_go": f"{RAW_BASE}papers/{paper_id}/paper.go",
        "raw_rs": f"{RAW_BASE}papers/{paper_id}/paper.rs",
        "repo_tree": f"{REPO_BASE}/tree/main/papers/{paper_id}",
    }


def _site_path(base_url: str) -> str:
    parsed = urlparse(base_url)
    path = parsed.path or "/"
    if not path.endswith("/"):
        path += "/"
    return path


def _first_nonempty_line(text: str) -> str | None:
    for line in text.splitlines():
        s = line.strip()
        if s:
            return s
    return None


def _tldr_from_summary(summary_text: str) -> str | None:
    # Heuristic: take the first paragraph after a TL;DR header.
    lines = summary_text.splitlines()
    start = None
    for i, line in enumerate(lines):
        if "TL;DR" in line:
            start = i
            break
    if start is None:
        return None
    out: list[str] = []
    for line in lines[start + 1 :]:
        s = line.strip()
        if not s:
            if out:
                break
            continue
        if s.startswith("##") or s.startswith("*"):
            continue
        if s.startswith("2.") and "Section" in s:
            break
        out.append(s)
        if len(" ".join(out)) > 800:
            break
    text = " ".join(out).strip()
    return text or None


def _safe_meta_description(text: str, limit: int = 280) -> str:
    collapsed = " ".join(text.split())
    if len(collapsed) <= limit:
        return collapsed
    return collapsed[: limit - 1].rstrip() + "…"


@dataclass(frozen=True)
class PaperInfo:
    paper_id: str
    title: str
    year: int | None
    authors: list[str]
    keywords: list[str]
    doi: str | None
    ssrn_url: str | None
    abstract: str | None
    summary_md: str | None
    summary_zh_md: str | None
    one_pager_md: str | None
    study_pack_md: str | None
    has_paper_py: bool
    has_paper_js: bool
    has_paper_cpp: bool
    has_paper_go: bool
    has_paper_rs: bool
    scholarly_jsonld: dict[str, Any] | None
    updated: datetime


def _authors_to_strings(authors: Any) -> list[str]:
    if not authors:
        return []
    if not isinstance(authors, list):
        return []
    out: list[str] = []
    for a in authors:
        if isinstance(a, dict):
            given = a.get("given") or ""
            family = a.get("family") or ""
            name = " ".join(x for x in [given, family] if isinstance(x, str) and x.strip()).strip()
            if not name and isinstance(a.get("name"), str):
                name = a["name"].strip()
            if name:
                out.append(name)
        elif isinstance(a, str) and a.strip():
            out.append(a.strip())
    return out


def _load_paper(papers_dir: Path, paper_dir: Path) -> PaperInfo:
    paper_id = paper_dir.name
    metadata_path = paper_dir / "metadata.json"
    metadata = _read_json(metadata_path) if metadata_path.exists() else {"id": paper_id}
    override = PAPER_OVERRIDES.get(paper_id, {})

    title = override.get("title") or metadata.get("title")
    if not isinstance(title, str) or not title.strip() or title.strip() == paper_id:
        paper_txt = paper_dir / "paper.txt"
        if paper_txt.exists():
            line = _first_nonempty_line(_read_text(paper_txt))
            if line:
                title = line
    if not isinstance(title, str) or not title.strip():
        title = paper_id

    year_value = metadata.get("year")
    if isinstance(year_value, int):
        year = year_value
    elif isinstance(year_value, str) and year_value.strip().isdigit():
        year = int(year_value.strip())
    else:
        year = None
    authors = list(override.get("authors") or _authors_to_strings(metadata.get("authors")))
    keywords = override.get("keywords") or (metadata.get("keywords") if isinstance(metadata.get("keywords"), list) else [])
    keywords = [k for k in keywords if isinstance(k, str)]

    doi = _normalize_doi(metadata.get("doi") if isinstance(metadata.get("doi"), str) else None)

    ssrn_url = metadata.get("ssrn_url")
    if not isinstance(ssrn_url, str):
        if paper_id.startswith("ssrn-") and paper_id.split("-", 1)[1].isdigit():
            ssrn_url = f"https://papers.ssrn.com/sol3/papers.cfm?abstract_id={paper_id.split('-', 1)[1]}"
        else:
            ssrn_url = None

    abstract = override.get("abstract") or (metadata.get("abstract") if isinstance(metadata.get("abstract"), str) else None)

    summary_path = paper_dir / "summary.md"
    summary_md = _read_text(summary_path) if summary_path.exists() else None

    summary_zh_path = paper_dir / "summary.zh.md"
    summary_zh_md = _read_text(summary_zh_path) if summary_zh_path.exists() else None

    one_pager_path = paper_dir / "one_pager.md"
    one_pager_md = _read_text(one_pager_path) if one_pager_path.exists() else None

    study_pack_path = paper_dir / "study_pack.md"
    study_pack_md = _read_text(study_pack_path) if study_pack_path.exists() else None

    # Generated summaries for the coined-term papers are staging artifacts.
    # Keep the author-facing page on the stable term binding until reviewed.
    if override:
        summary_md = None
        summary_zh_md = None
        one_pager_md = None
        study_pack_md = None

    has_paper_py = (paper_dir / "paper.py").exists()
    has_paper_js = (paper_dir / "paper.js").exists()
    has_paper_cpp = (paper_dir / "paper.cpp").exists()
    has_paper_go = (paper_dir / "paper.go").exists()
    has_paper_rs = (paper_dir / "paper.rs").exists()

    scholarly_path = paper_dir / "scholarlyarticle.jsonld"
    scholarly_jsonld = _read_json(scholarly_path) if scholarly_path.exists() else None
    if scholarly_jsonld:
        # Do not repeat the historical blanket CC0 grant for article files.
        scholarly_jsonld["license"] = IN_COPYRIGHT_URL
        if override:
            scholarly_jsonld["name"] = title
            scholarly_jsonld["headline"] = title
            scholarly_jsonld["abstract"] = abstract
            scholarly_jsonld["keywords"] = keywords
            scholarly_jsonld["author"] = [{"@type": "Person", "name": name} for name in authors]
    if not doi and scholarly_jsonld:
        doi = _extract_doi_from_identifier(scholarly_jsonld.get("identifier"))
        if not doi and isinstance(scholarly_jsonld.get("sameAs"), list):
            for item in scholarly_jsonld["sameAs"]:
                doi = _normalize_doi(item if isinstance(item, str) else None)
                if doi:
                    break

    # Updated timestamp: prefer VCS times for stability in CI (checkout mtimes are noisy).
    candidates = [summary_path, one_pager_path, study_pack_path, paper_dir / "paper.txt", metadata_path]
    stamps: list[datetime] = []
    for path in candidates:
        if not path.exists():
            continue
        dt = _git_last_modified(path)
        if dt is None:
            dt = datetime.fromtimestamp(path.stat().st_mtime, tz=timezone.utc)
        stamps.append(dt)
    updated = max(stamps) if stamps else datetime.fromtimestamp(paper_dir.stat().st_mtime, tz=timezone.utc)

    return PaperInfo(
        paper_id=paper_id,
        title=title,
        year=year,
        authors=authors,
        keywords=keywords,
        doi=doi,
        ssrn_url=ssrn_url,
        abstract=abstract,
        summary_md=summary_md,
        summary_zh_md=summary_zh_md,
        one_pager_md=one_pager_md,
        study_pack_md=study_pack_md,
        has_paper_py=has_paper_py,
        has_paper_js=has_paper_js,
        has_paper_cpp=has_paper_cpp,
        has_paper_go=has_paper_go,
        has_paper_rs=has_paper_rs,
        scholarly_jsonld=scholarly_jsonld,
        updated=updated,
    )


def _render_layout(
    *, title: str, description: str, canonical_url: str, site_path: str, body_html: str, extra_head: str = ""
) -> str:
    return f"""<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <title>{html.escape(title)}</title>
    <meta name="description" content="{html.escape(description)}" />
    <link rel="canonical" href="{html.escape(canonical_url)}" />

    <meta property="og:site_name" content="my-works-for-llm" />
    <meta property="og:type" content="article" />
    <meta property="og:title" content="{html.escape(title)}" />
    <meta property="og:description" content="{html.escape(description)}" />
    <meta property="og:url" content="{html.escape(canonical_url)}" />

    <meta name="twitter:card" content="summary" />
    <meta name="twitter:title" content="{html.escape(title)}" />
    <meta name="twitter:description" content="{html.escape(description)}" />

    <link rel="alternate" type="application/atom+xml" href="{html.escape(site_path)}atom.xml" title="my-works-for-llm updates" />
    <link rel="alternate" type="text/plain" href="{html.escape(site_path)}llms.txt" title="LLM descriptor" />
    <link rel="stylesheet" href="{html.escape(site_path)}assets/style.css" />
{extra_head}
  </head>
  <body>
    <header class="site-header">
      <div class="container">
        <a class="brand" href="{html.escape(site_path)}">my-works-for-llm</a>
        <nav class="nav">
          <a href="{html.escape(REPO_BASE)}">GitHub</a>
          <a href="{html.escape(site_path)}atom.xml">Atom</a>
          <a href="{html.escape(site_path)}sitemap.xml">Sitemap</a>
          <a href="{html.escape(site_path)}llms.txt">LLMs.txt</a>
        </nav>
      </div>
    </header>
    <main class="container">
{body_html}
    </main>
    <footer class="site-footer">
      <div class="container">
        <div>Machine-readable corpus of Professor Yonathan Arbel’s scholarship for LLM research.</div>
        <div>Dataset DOI: <a href="{html.escape(DATASET_DOI_URL)}">{html.escape(DATASET_DOI_URL)}</a></div>
        <div><a href="{html.escape(DATASET_DOI_URL)}"><img src="{html.escape(DATASET_DOI_BADGE)}" alt="DOI badge" /></a></div>
      </div>
    </footer>
  </body>
</html>
"""


def _render_index(base_url: str, canonical_base_url: str, papers: list[PaperInfo]) -> str:
    site_path = _site_path(base_url)
    rows: list[str] = []
    for p in papers:
        urls = _paper_urls(base_url, p.paper_id)
        subtitle_parts = []
        if p.year:
            subtitle_parts.append(str(p.year))
        if p.authors:
            subtitle_parts.append(", ".join(p.authors))
        subtitle = " · ".join(subtitle_parts)

        kw = ", ".join(p.keywords[:8]) if p.keywords else ""
        rows.append(
            f"""<article class="card">
  <h2 class="card-title"><a href="{html.escape(urls['page'])}">{html.escape(p.title)}</a></h2>
  <div class="card-subtitle">{html.escape(subtitle)}</div>
  <div class="card-meta">{html.escape(p.paper_id)}{(" · " + html.escape(kw)) if kw else ""}</div>
</article>"""
        )

    body = f"""
<section class="hero">
  <h1>Professor Yonathan Arbel — Corpus</h1>
  <p>Human + machine-readable pages for LLM ingestion, discovery, and citation.</p>
  <div class="hero-links">
    <a class="btn" href="{html.escape(REPO_BASE)}">View on GitHub</a>
    <a class="btn" href="/my-works-for-llm/atom.xml">Atom feed</a>
    <a class="btn" href="/my-works-for-llm/llms.txt">LLMs.txt</a>
  </div>
</section>

<section class="grid">
  {''.join(rows)}
</section>
"""

    canonical = canonical_base_url.rstrip("/") + "/"
    return _render_layout(
        title="my-works-for-llm — Yonathan Arbel corpus",
        description="Machine-readable corpus of Professor Yonathan Arbel’s scholarship for LLM research (summaries + full text + metadata).",
        canonical_url=canonical,
        site_path=site_path,
        body_html=body,
        extra_head="",
    )


def _render_papers_index(base_url: str, canonical_base_url: str, papers: list[PaperInfo]) -> str:
    site_path = _site_path(base_url)
    items: list[str] = []
    for p in papers:
        urls = _paper_urls(base_url, p.paper_id)
        items.append(
            f'<li><a href="{html.escape(urls["page"])}">{html.escape(p.title)} ({html.escape(p.paper_id)})</a></li>'
        )

    body = f"""
<section class="hero">
  <h1>Paper Index</h1>
  <p>Direct links to all paper pages in the corpus.</p>
</section>

<section>
  <ul>
    {''.join(items)}
  </ul>
</section>
"""

    canonical = canonical_base_url.rstrip("/") + "/papers/"
    return _render_layout(
        title="my-works-for-llm — papers",
        description="Index of all paper pages in the my-works-for-llm corpus.",
        canonical_url=canonical,
        site_path=site_path,
        body_html=body,
        extra_head="",
    )


def _render_paper_page(base_url: str, canonical_base_url: str, paper: PaperInfo) -> str:
    site_path = _site_path(base_url)
    urls = _paper_urls(base_url, paper.paper_id)
    canonical_urls = _paper_urls(canonical_base_url, paper.paper_id)

    abstract_text = paper.abstract
    if not abstract_text and paper.summary_md:
        abstract_text = _tldr_from_summary(paper.summary_md) or _first_nonempty_line(paper.summary_md)
    if not abstract_text and paper.one_pager_md:
        abstract_text = _first_nonempty_line(paper.one_pager_md)
    if not abstract_text and paper.study_pack_md:
        abstract_text = _first_nonempty_line(paper.study_pack_md)
    if not abstract_text:
        abstract_text = f"Read the full text at {urls['raw_text']}."
    description = _safe_meta_description(abstract_text)

    meta_lines: list[str] = []
    if paper.year:
        meta_lines.append(f"<span class=\"pill\">{paper.year}</span>")
    meta_lines.append(f"<span class=\"pill\">{html.escape(paper.paper_id)}</span>")
    if paper.keywords:
        for k in paper.keywords[:8]:
            meta_lines.append(f"<span class=\"pill pill-muted\">{html.escape(k)}</span>")

    links: list[str] = []
    if paper.ssrn_url:
        links.append(f"<a class=\"btn\" href=\"{html.escape(paper.ssrn_url)}\">SSRN</a>")
    links.append(f"<a class=\"btn\" href=\"{html.escape(urls['raw_pdf'])}\">Download PDF</a>")
    links.append(f"<a class=\"btn\" href=\"{html.escape(urls['raw_text'])}\">Download text</a>")
    links.append(f"<a class=\"btn\" href=\"{html.escape(urls['raw_summary'])}\">Summary (MD)</a>")
    if paper.summary_zh_md:
        links.append(f"<a class=\"btn\" href=\"{html.escape(urls['raw_summary_zh'])}\">中文摘要 (MD)</a>")
    if paper.one_pager_md:
        links.append(f"<a class=\"btn\" href=\"{html.escape(urls['raw_one_pager'])}\">One-pager (MD)</a>")
    if paper.study_pack_md:
        links.append(f"<a class=\"btn\" href=\"{html.escape(urls['raw_study_pack'])}\">Study pack (MD)</a>")
    if paper.has_paper_py:
        links.append(f"<a class=\"btn\" href=\"{html.escape(urls['raw_py'])}\">Python</a>")
    if paper.has_paper_js:
        links.append(f"<a class=\"btn\" href=\"{html.escape(urls['raw_js'])}\">JavaScript</a>")
    if paper.has_paper_cpp:
        links.append(f"<a class=\"btn\" href=\"{html.escape(urls['raw_cpp'])}\">C++</a>")
    if paper.has_paper_go:
        links.append(f"<a class=\"btn\" href=\"{html.escape(urls['raw_go'])}\">Go</a>")
    if paper.has_paper_rs:
        links.append(f"<a class=\"btn\" href=\"{html.escape(urls['raw_rs'])}\">Rust</a>")
    links.append(f"<a class=\"btn\" href=\"{html.escape(urls['repo_tree'])}\">Files</a>")

    scholarly_json = ""
    if paper.scholarly_jsonld:
        scholarly_json = json.dumps(paper.scholarly_jsonld, ensure_ascii=False, indent=2)
        scholarly_json = f'\n    <script type="application/ld+json">\n{scholarly_json}\n    </script>\n'

    abstract_html = f"""
<section>
  <h2>Abstract</h2>
  <p>{html.escape(abstract_text)}</p>
</section>
"""

    doi_url = f"https://doi.org/{paper.doi}" if paper.doi else ""
    citation_url = doi_url or paper.ssrn_url or urls["page"]
    authors_text = ", ".join(paper.authors) if paper.authors else "Yonathan A. Arbel"
    year_text = str(paper.year) if paper.year else "n.d."
    apa_citation = f"{authors_text}. ({year_text}). {paper.title}. {citation_url}"
    bluebook_citation = f"{authors_text}, {paper.title}, {year_text}, {citation_url}."
    doi_html = (
        f'<p><strong>DOI:</strong> <a href="{html.escape(doi_url)}">{html.escape(doi_url)}</a></p>'
        if doi_url
        else ""
    )
    citation_html = f"""
<section>
  <h2>Citation</h2>
  <p><strong>APA:</strong> {html.escape(apa_citation)}</p>
  <p><strong>Bluebook:</strong> {html.escape(bluebook_citation)}</p>
  {doi_html}
</section>
"""

    summary_html = ""
    if paper.summary_md:
        summary_html = f"""
<section>
  <h2>Summary (English)</h2>
  <pre class="md">{html.escape(paper.summary_md.strip())}</pre>
</section>
"""

    one_pager_html = ""
    if paper.one_pager_md:
        one_pager_html = f"""
<section>
  <h2>One-page summary</h2>
  <pre class="md">{html.escape(paper.one_pager_md.strip())}</pre>
</section>
"""

    study_pack_html = ""
    if paper.study_pack_md:
        study_pack_html = f"""
<section>
  <h2>Study pack</h2>
  <pre class="md">{html.escape(paper.study_pack_md.strip())}</pre>
</section>
"""

    summary_zh_html = ""
    if paper.summary_zh_md:
        summary_zh_html = f"""
<section>
  <h2>摘要（中文）</h2>
  <pre class="md">{html.escape(paper.summary_zh_md.strip())}</pre>
</section>
"""

    body = f"""
<article>
  <h1>{html.escape(paper.title)}</h1>
  <div class="meta">{''.join(meta_lines)}</div>
  <div class="actions">{''.join(links)}</div>
  {abstract_html}
  {citation_html}
  {summary_html}
  {one_pager_html}
  {study_pack_html}
  {summary_zh_html}
</article>
"""

    extra_head = scholarly_json
    title_suffix = "Yonathan Arbel" + (f" ({paper.year})" if paper.year else "")
    return _render_layout(
        title=f"{paper.title} — {title_suffix}",
        description=description,
        canonical_url=canonical_urls["page"],
        site_path=site_path,
        body_html=body,
        extra_head=extra_head,
    )


def _write(path: Path, content: str) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    normalized = "\n".join(line.rstrip() for line in content.splitlines()) + "\n"
    path.write_text(normalized, encoding="utf-8")


def _write_llm_descriptors(out_dir: Path, base_url: str, papers: list[PaperInfo]) -> None:
    base = base_url.rstrip("/") + "/"
    lines: list[str] = []
    lines.append("# Yonathan Arbel Scholarly Corpus")
    lines.append("")
    lines.append("> Machine-readable corpus of legal scholarship.")
    lines.append("")
    lines.append("## Start Here")
    lines.append(f"- {base}")
    lines.append(f"- {base}papers/")
    lines.append("")
    lines.append("## Key Term Bindings")
    by_id = {paper.paper_id: paper for paper in papers}
    for paper_id, override in PAPER_OVERRIDES.items():
        if paper_id in by_id:
            lines.append(f"- {override['abstract']} Source: {base}papers/{paper_id}/")
    lines.append("")
    lines.append("## Key Papers")
    for paper in papers:
        lines.append(f"- {base}papers/{paper.paper_id}/")
    lines.append("")
    lines.append("## Machine Endpoints")
    lines.append(f"- {base}sitemap.xml")
    lines.append(f"- {base}atom.xml")
    lines.append(f"- {base}llms.txt")
    lines.append(f"- {base}llms-full.txt")
    lines.append("")

    content = "\n".join(lines)
    _write(out_dir / "llms.txt", content)
    # Compatibility alias for crawlers expecting singular filename.
    _write(out_dir / "llm.txt", content)


def _write_full_corpus_dump(path: Path, papers: list[PaperInfo], papers_dir: Path) -> None:
    lines: list[str] = []
    lines.append("# my-works-for-llm: Full Corpus Dump")
    lines.append("")
    lines.append(f"Dataset DOI: {DATASET_DOI_URL}")
    lines.append("License: corpus-authored metadata is CC-BY-4.0; article text and PDFs retain source-specific rights.")
    lines.append(f"Generated (UTC): {_utc_iso(datetime.now(timezone.utc))}")
    lines.append("")
    lines.append("This file concatenates corpus content into a single ingestible artifact.")
    lines.append("")

    for paper in papers:
        paper_dir = papers_dir / paper.paper_id
        selected_path: Path | None = None
        for name in ("paper.txt", "summary.md", "summary.zh.md"):
            candidate = paper_dir / name
            if candidate.exists():
                selected_path = candidate
                break
        if selected_path is None:
            continue

        text = _read_text(selected_path).strip()
        if not text:
            continue

        source_rel = selected_path.relative_to(papers_dir.parent).as_posix()
        lines.append("---")
        lines.append("")
        lines.append(f"## {paper.paper_id}: {paper.title}")
        if paper.year:
            lines.append(f"Year: {paper.year}")
        if paper.authors:
            lines.append(f"Authors: {', '.join(paper.authors)}")
        lines.append(f"Source: {source_rel}")
        lines.append("")
        lines.append(text)
        lines.append("")

    _write(path, "\n".join(lines).rstrip() + "\n")


def _write_style(path: Path) -> None:
    css = """
:root {
  --bg: #0b0c10;
  --panel: #11131a;
  --text: #e8eaf2;
  --muted: #aab0c0;
  --border: #23283a;
  --accent: #78a6ff;
}

* { box-sizing: border-box; }
body {
  margin: 0;
  background: var(--bg);
  color: var(--text);
  font-family: ui-sans-serif, system-ui, -apple-system, Segoe UI, Roboto, Arial, "Apple Color Emoji", "Segoe UI Emoji";
  line-height: 1.5;
}

a { color: var(--accent); text-decoration: none; }
a:hover { text-decoration: underline; }

.container { max-width: 980px; margin: 0 auto; padding: 0 16px; }

.site-header {
  position: sticky;
  top: 0;
  background: rgba(11, 12, 16, 0.85);
  backdrop-filter: blur(10px);
  border-bottom: 1px solid var(--border);
}
.site-header .container { display: flex; gap: 16px; align-items: center; justify-content: space-between; padding: 14px 16px; }
.brand { font-weight: 700; letter-spacing: 0.2px; }
.nav a { margin-left: 14px; color: var(--muted); }
.nav a:hover { color: var(--text); }

.site-footer { border-top: 1px solid var(--border); margin-top: 28px; padding: 18px 0; color: var(--muted); }

.hero { padding: 28px 0 12px; }
.hero h1 { margin: 0 0 8px; font-size: 28px; }
.hero p { margin: 0 0 16px; color: var(--muted); }
.hero-links { display: flex; gap: 10px; flex-wrap: wrap; }

.grid { display: grid; grid-template-columns: 1fr; gap: 12px; padding: 14px 0 24px; }
@media (min-width: 820px) { .grid { grid-template-columns: 1fr 1fr; } }

.card {
  border: 1px solid var(--border);
  background: var(--panel);
  border-radius: 12px;
  padding: 14px 14px 12px;
}
.card-title { margin: 0 0 6px; font-size: 16px; }
.card-subtitle { color: var(--muted); font-size: 13px; margin-bottom: 6px; }
.card-meta { color: var(--muted); font-size: 12px; }

.meta { display: flex; gap: 8px; flex-wrap: wrap; margin: 10px 0; }
.pill {
  display: inline-block;
  padding: 4px 8px;
  border-radius: 999px;
  border: 1px solid var(--border);
  background: rgba(255,255,255,0.03);
  color: var(--text);
  font-size: 12px;
}
.pill-muted { color: var(--muted); }

.actions { display: flex; gap: 10px; flex-wrap: wrap; margin: 10px 0 18px; }
.btn {
  display: inline-block;
  padding: 8px 10px;
  border-radius: 10px;
  border: 1px solid var(--border);
  background: rgba(255,255,255,0.04);
  color: var(--text);
  font-size: 13px;
}
.btn:hover { background: rgba(255,255,255,0.08); text-decoration: none; }

h1 { font-size: 26px; margin: 18px 0 8px; }
h2 { font-size: 18px; margin: 20px 0 10px; }
p { margin: 0 0 12px; }

.md {
  white-space: pre-wrap;
  padding: 14px;
  border-radius: 12px;
  border: 1px solid var(--border);
  background: rgba(255,255,255,0.03);
  overflow-x: auto;
}
"""
    _write(path, css.strip() + "\n")


def _write_atom(path: Path, base_url: str, papers: list[PaperInfo]) -> None:
    base_url = base_url.rstrip("/") + "/"
    updated = max((p.updated for p in papers), default=datetime.now(timezone.utc))
    feed_updated = _utc_iso(updated)

    lines: list[str] = []
    lines.append('<?xml version="1.0" encoding="utf-8"?>')
    lines.append('<feed xmlns="http://www.w3.org/2005/Atom">')
    lines.append(f"  <title>my-works-for-llm</title>")
    lines.append(f"  <id>{html.escape(base_url)}</id>")
    lines.append(f"  <link href=\"{html.escape(base_url)}\" />")
    lines.append(f"  <link href=\"{html.escape(base_url + 'atom.xml')}\" rel=\"self\" />")
    lines.append(f"  <updated>{feed_updated}</updated>")
    lines.append("  <author><name>Yonathan A. Arbel</name></author>")

    for paper in sorted(papers, key=lambda p: p.updated, reverse=True)[:30]:
        urls = _paper_urls(base_url, paper.paper_id)
        entry_updated = _utc_iso(paper.updated)
        desc = paper.abstract or (paper.summary_md and (_tldr_from_summary(paper.summary_md) or "")) or ""
        desc = _safe_meta_description(desc or paper.title, limit=400)
        lines.append("  <entry>")
        lines.append(f"    <title>{html.escape(paper.title)}</title>")
        lines.append(f"    <id>{html.escape(urls['page'])}</id>")
        lines.append(f"    <link href=\"{html.escape(urls['page'])}\" />")
        lines.append(f"    <updated>{entry_updated}</updated>")
        lines.append(f"    <summary>{html.escape(desc)}</summary>")
        lines.append("  </entry>")

    lines.append("</feed>")
    _write(path, "\n".join(lines) + "\n")


def _write_pages_sitemap(path: Path, base_url: str, papers: list[PaperInfo]) -> None:
    base_url = base_url.rstrip("/") + "/"
    lines: list[str] = []
    lines.append('<?xml version="1.0" encoding="UTF-8"?>')
    lines.append('<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">')
    # Index
    lines.append("  <url>")
    lines.append(f"    <loc>{html.escape(base_url)}</loc>")
    lines.append(f"    <lastmod>{_utc_date(datetime.now(timezone.utc))}</lastmod>")
    lines.append("  </url>")
    # Papers index
    lines.append("  <url>")
    lines.append(f"    <loc>{html.escape(base_url + 'papers/')}</loc>")
    lines.append(f"    <lastmod>{_utc_date(datetime.now(timezone.utc))}</lastmod>")
    lines.append("  </url>")
    # Paper pages
    for p in papers:
        urls = _paper_urls(base_url, p.paper_id)
        lines.append("  <url>")
        lines.append(f"    <loc>{html.escape(urls['page'])}</loc>")
        lines.append(f"    <lastmod>{_utc_date(p.updated)}</lastmod>")
        lines.append("  </url>")
    # Feed
    lines.append("  <url>")
    lines.append(f"    <loc>{html.escape(base_url + 'atom.xml')}</loc>")
    lines.append(f"    <lastmod>{_utc_date(datetime.now(timezone.utc))}</lastmod>")
    lines.append("  </url>")
    # Crawl helpers
    for rel in ("robots.txt", "llms.txt", "llm.txt", "llms-full.txt"):
        lines.append("  <url>")
        lines.append(f"    <loc>{html.escape(base_url + rel)}</loc>")
        lines.append(f"    <lastmod>{_utc_date(datetime.now(timezone.utc))}</lastmod>")
        lines.append("  </url>")
    lines.append("</urlset>")
    _write(path, "\n".join(lines) + "\n")


def _write_robots(path: Path, base_url: str) -> None:
    base_url = base_url.rstrip("/") + "/"
    content = f"""User-agent: *
Allow: /

User-agent: GPTBot
Allow: /

User-agent: OAI-SearchBot
Allow: /

User-agent: ClaudeBot
Allow: /

User-agent: Claude-SearchBot
Allow: /

User-agent: PerplexityBot
Allow: /

User-agent: CCBot
Allow: /

Sitemap: {base_url}sitemap.xml
"""
    _write(path, content)


def main() -> int:
    parser = argparse.ArgumentParser(description="Generate GitHub Pages site under docs/.")
    parser.add_argument("--base-url", default=DEFAULT_BASE_URL, help="Base URL for the GitHub Pages mirror.")
    parser.add_argument(
        "--canonical-base-url",
        default=DEFAULT_CANONICAL_BASE_URL,
        help="Canonical corpus URL used in metadata, feeds, sitemaps, and crawler descriptors.",
    )
    parser.add_argument("--papers-dir", default="papers", help="Papers directory (default: papers).")
    parser.add_argument("--out-dir", default="docs", help="Output directory (default: docs).")
    args = parser.parse_args()

    base_url = args.base_url.rstrip("/") + "/"
    canonical_base_url = args.canonical_base_url.rstrip("/") + "/"
    papers_dir = Path(args.papers_dir)
    out_dir = Path(args.out_dir)

    if not papers_dir.exists():
        raise SystemExit(f"Papers directory not found: {papers_dir}")

    papers: list[PaperInfo] = []
    for paper_dir in sorted(papers_dir.iterdir()):
        if paper_dir.is_dir():
            papers.append(_load_paper(papers_dir, paper_dir))

    # Static assets
    _write_style(out_dir / "assets" / "style.css")
    _write(out_dir / ".nojekyll", "")
    _write_llm_descriptors(out_dir, canonical_base_url, papers)
    _write_full_corpus_dump(out_dir / "llms-full.txt", papers, papers_dir)

    # Pages
    _write(out_dir / "index.html", _render_index(base_url, canonical_base_url, papers))
    _write(out_dir / "papers" / "index.html", _render_papers_index(base_url, canonical_base_url, papers))
    for paper in papers:
        _write(out_dir / "papers" / paper.paper_id / "index.html", _render_paper_page(base_url, canonical_base_url, paper))

    # Feed + sitemap + robots
    _write_atom(out_dir / "atom.xml", canonical_base_url, papers)
    _write_pages_sitemap(out_dir / "sitemap.xml", canonical_base_url, papers)
    _write_robots(out_dir / "robots.txt", canonical_base_url)

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
