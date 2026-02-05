#!/usr/bin/env python3
"""
Generate sitemap.xml for the corpus.

    Includes summaries (English/Chinese when present), extended study assets, and
    full-text paper.txt files so crawlers and downstream tooling can discover
    canonical URLs.
"""

from __future__ import annotations

import argparse
from datetime import datetime, timezone
from pathlib import Path
import subprocess
from typing import Iterable


SITEMAP_NS = "http://www.sitemaps.org/schemas/sitemap/0.9"


def _git_lastmod_date(path: Path) -> str | None:
    # Use VCS timestamps for stability in CI (filesystem mtimes vary on checkout).
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
    try:
        dt = datetime.fromisoformat(stamp.replace("Z", "+00:00"))
    except ValueError:
        return None
    return dt.astimezone(timezone.utc).strftime("%Y-%m-%d")


def _iter_urls(papers_dir: Path) -> Iterable[tuple[str, str]]:
    """Yield (relative_path, lastmod_yyyy_mm_dd)."""
    for paper_dir in sorted(papers_dir.iterdir()):
        if not paper_dir.is_dir():
            continue

        candidates = [
            paper_dir / "summary.md",
            paper_dir / "summary.zh.md",
            paper_dir / "one_pager.md",
            paper_dir / "study_pack.md",
            paper_dir / "paper.txt",
            paper_dir / "scholarlyarticle.jsonld",
        ]

        for path in candidates:
            if not path.exists():
                continue
            lastmod = _git_lastmod_date(path) or datetime.fromtimestamp(path.stat().st_mtime, tz=timezone.utc).strftime(
                "%Y-%m-%d"
            )
            rel = str(path.relative_to(papers_dir.parent)).replace("\\", "/")
            yield rel, lastmod


def build_sitemap_xml(base_url: str, papers_dir: Path) -> str:
    base = base_url.rstrip("/") + "/"

    lines: list[str] = []
    lines.append('<?xml version="1.0" encoding="UTF-8"?>')
    lines.append(f'<urlset xmlns="{SITEMAP_NS}">')
    for rel, lastmod in _iter_urls(papers_dir):
        lines.append("  <url>")
        lines.append(f"    <loc>{base}{rel}</loc>")
        lines.append(f"    <lastmod>{lastmod}</lastmod>")
        lines.append("  </url>")
    lines.append("</urlset>")
    return "\n".join(lines) + "\n"


def main() -> int:
    parser = argparse.ArgumentParser(description="Generate sitemap.xml for the corpus.")
    parser.add_argument(
        "--base-url",
        default="https://raw.githubusercontent.com/yonathanarbel/my-works-for-llm/main/",
        help="Base URL prefix to use in <loc> entries.",
    )
    parser.add_argument("--out", default="sitemap.xml", help="Output path (default: sitemap.xml).")
    args = parser.parse_args()

    repo_root = Path(".")
    papers_dir = repo_root / "papers"
    if not papers_dir.exists():
        raise SystemExit(f"Papers directory not found: {papers_dir}")

    xml = build_sitemap_xml(args.base_url, papers_dir)
    Path(args.out).write_text(xml, encoding="utf-8", newline="\n")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
