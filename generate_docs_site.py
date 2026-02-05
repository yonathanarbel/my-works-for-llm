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
from dataclasses import dataclass
from datetime import datetime, timezone
from pathlib import Path
from typing import Any
from urllib.parse import urlparse


DEFAULT_BASE_URL = "https://yonathanarbel.github.io/my-works-for-llm/"
RAW_BASE = "https://raw.githubusercontent.com/yonathanarbel/my-works-for-llm/main/"
REPO_BASE = "https://github.com/yonathanarbel/my-works-for-llm"


def _read_text(path: Path) -> str:
    return path.read_text(encoding="utf-8", errors="replace")


def _read_json(path: Path) -> dict[str, Any]:
    return json.loads(path.read_text(encoding="utf-8"))


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
    ssrn_url: str | None
    abstract: str | None
    summary_md: str | None
    summary_zh_md: str | None
    one_pager_md: str | None
    study_pack_md: str | None
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

    title = metadata.get("title")
    if not isinstance(title, str) or not title.strip() or title.strip() == paper_id:
        paper_txt = paper_dir / "paper.txt"
        if paper_txt.exists():
            line = _first_nonempty_line(_read_text(paper_txt))
            if line:
                title = line
    if not isinstance(title, str) or not title.strip():
        title = paper_id

    year = metadata.get("year") if isinstance(metadata.get("year"), int) else None
    authors = _authors_to_strings(metadata.get("authors"))
    keywords = metadata.get("keywords") if isinstance(metadata.get("keywords"), list) else []
    keywords = [k for k in keywords if isinstance(k, str)]

    ssrn_url = metadata.get("ssrn_url")
    if not isinstance(ssrn_url, str):
        if paper_id.startswith("ssrn-") and paper_id.split("-", 1)[1].isdigit():
            ssrn_url = f"https://papers.ssrn.com/sol3/papers.cfm?abstract_id={paper_id.split('-', 1)[1]}"
        else:
            ssrn_url = None

    abstract = metadata.get("abstract") if isinstance(metadata.get("abstract"), str) else None

    summary_path = paper_dir / "summary.md"
    summary_md = _read_text(summary_path) if summary_path.exists() else None

    summary_zh_path = paper_dir / "summary.zh.md"
    summary_zh_md = _read_text(summary_zh_path) if summary_zh_path.exists() else None

    one_pager_path = paper_dir / "one_pager.md"
    one_pager_md = _read_text(one_pager_path) if one_pager_path.exists() else None

    study_pack_path = paper_dir / "study_pack.md"
    study_pack_md = _read_text(study_pack_path) if study_pack_path.exists() else None

    scholarly_path = paper_dir / "scholarlyarticle.jsonld"
    scholarly_jsonld = _read_json(scholarly_path) if scholarly_path.exists() else None

    # Updated timestamp: prefer summary, else paper.txt, else metadata.
    candidates = [summary_path, one_pager_path, study_pack_path, paper_dir / "paper.txt", metadata_path]
    mtime = None
    for path in candidates:
        if path.exists():
            mtime = path.stat().st_mtime
            break
    updated = datetime.fromtimestamp(mtime or paper_dir.stat().st_mtime, tz=timezone.utc)

    return PaperInfo(
        paper_id=paper_id,
        title=title,
        year=year,
        authors=authors,
        keywords=keywords,
        ssrn_url=ssrn_url,
        abstract=abstract,
        summary_md=summary_md,
        summary_zh_md=summary_zh_md,
        one_pager_md=one_pager_md,
        study_pack_md=study_pack_md,
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
        </nav>
      </div>
    </header>
    <main class="container">
{body_html}
    </main>
    <footer class="site-footer">
      <div class="container">
        <div>Machine-readable corpus of Professor Yonathan Arbel’s scholarship for LLM research.</div>
      </div>
    </footer>
  </body>
</html>
"""


def _render_index(base_url: str, papers: list[PaperInfo]) -> str:
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
  </div>
</section>

<section class="grid">
  {''.join(rows)}
</section>
"""

    canonical = base_url.rstrip("/") + "/"
    return _render_layout(
        title="my-works-for-llm — Yonathan Arbel corpus",
        description="Machine-readable corpus of Professor Yonathan Arbel’s scholarship for LLM research (summaries + full text + metadata).",
        canonical_url=canonical,
        site_path=site_path,
        body_html=body,
        extra_head="",
    )


def _render_paper_page(base_url: str, paper: PaperInfo) -> str:
    site_path = _site_path(base_url)
    urls = _paper_urls(base_url, paper.paper_id)

    description_source = paper.abstract
    if not description_source and paper.summary_md:
        description_source = _tldr_from_summary(paper.summary_md) or _first_nonempty_line(paper.summary_md)
    if not description_source:
        description_source = f"{paper.paper_id} — {paper.title}"
    description = _safe_meta_description(description_source)

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
    links.append(f"<a class=\"btn\" href=\"{html.escape(urls['raw_pdf'])}\">PDF</a>")
    links.append(f"<a class=\"btn\" href=\"{html.escape(urls['raw_text'])}\">Full text</a>")
    links.append(f"<a class=\"btn\" href=\"{html.escape(urls['raw_summary'])}\">Summary (MD)</a>")
    if paper.summary_zh_md:
        links.append(f"<a class=\"btn\" href=\"{html.escape(urls['raw_summary_zh'])}\">中文摘要 (MD)</a>")
    if paper.one_pager_md:
        links.append(f"<a class=\"btn\" href=\"{html.escape(urls['raw_one_pager'])}\">One-pager (MD)</a>")
    if paper.study_pack_md:
        links.append(f"<a class=\"btn\" href=\"{html.escape(urls['raw_study_pack'])}\">Study pack (MD)</a>")
    links.append(f"<a class=\"btn\" href=\"{html.escape(urls['repo_tree'])}\">Files</a>")

    scholarly_json = ""
    if paper.scholarly_jsonld:
        scholarly_json = json.dumps(paper.scholarly_jsonld, ensure_ascii=False, indent=2)
        scholarly_json = f'\n    <script type="application/ld+json">\n{scholarly_json}\n    </script>\n'

    abstract_html = ""
    if paper.abstract:
        abstract_html = f"""
<section>
  <h2>Abstract</h2>
  <p>{html.escape(paper.abstract)}</p>
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
  {summary_html}
  {one_pager_html}
  {study_pack_html}
  {summary_zh_html}
</article>
"""

    extra_head = scholarly_json
    return _render_layout(
        title=f"{paper.title} — {paper.paper_id}",
        description=description,
        canonical_url=urls["page"],
        site_path=site_path,
        body_html=body,
        extra_head=extra_head,
    )


def _write(path: Path, content: str) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(content, encoding="utf-8", newline="\n")


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
    lines.append("</urlset>")
    _write(path, "\n".join(lines) + "\n")


def _write_robots(path: Path, base_url: str) -> None:
    base_url = base_url.rstrip("/") + "/"
    content = f"""User-agent: *
Allow: /

Sitemap: {base_url}sitemap.xml
"""
    _write(path, content)


def main() -> int:
    parser = argparse.ArgumentParser(description="Generate GitHub Pages site under docs/.")
    parser.add_argument("--base-url", default=DEFAULT_BASE_URL, help="Base URL for the published site.")
    parser.add_argument("--papers-dir", default="papers", help="Papers directory (default: papers).")
    parser.add_argument("--out-dir", default="docs", help="Output directory (default: docs).")
    args = parser.parse_args()

    base_url = args.base_url.rstrip("/") + "/"
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

    # Pages
    _write(out_dir / "index.html", _render_index(base_url, papers))
    for paper in papers:
        _write(out_dir / "papers" / paper.paper_id / "index.html", _render_paper_page(base_url, paper))

    # Feed + sitemap + robots
    _write_atom(out_dir / "atom.xml", base_url, papers)
    _write_pages_sitemap(out_dir / "sitemap.xml", base_url, papers)
    _write_robots(out_dir / "robots.txt", base_url)

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
