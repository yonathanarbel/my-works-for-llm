#!/usr/bin/env python3
"""
Build a simple JSONL "chunk index" suitable for RAG pipelines.

This does not call any model provider. It just chunks `paper.txt` files into
uniform-sized records with useful metadata.
"""

from __future__ import annotations

import argparse
import json
import re
from pathlib import Path
from typing import Any, Iterable, Iterator


def _read_text(path: Path) -> str:
    return path.read_text(encoding="utf-8", errors="replace")


def _read_json(path: Path) -> dict[str, Any]:
    return json.loads(path.read_text(encoding="utf-8"))


def _normalize(text: str) -> str:
    text = text.replace("\r\n", "\n").replace("\r", "\n")
    text = text.replace("\f", "\n\n")
    return text


def _split_paragraphs(text: str) -> list[str]:
    text = _normalize(text)
    parts = [p.strip() for p in re.split(r"\n\s*\n+", text) if p.strip()]
    if len(parts) <= 1:
        parts = [l.strip() for l in text.splitlines() if l.strip()]
    return parts


def _chunk_paragraphs(paragraphs: list[str], max_chars: int, overlap: int) -> list[str]:
    chunks: list[str] = []
    current: list[str] = []
    current_len = 0

    def flush() -> None:
        nonlocal current, current_len
        if not current:
            return
        chunks.append("\n\n".join(current).strip())
        if overlap > 0:
            current = current[-overlap:]
            current_len = sum(len(p) + 2 for p in current)
        else:
            current = []
            current_len = 0

    for p in paragraphs:
        if not p:
            continue
        add_len = len(p) + 2
        if current and current_len + add_len > max_chars:
            flush()
        current.append(p)
        current_len += add_len

    if current:
        chunks.append("\n\n".join(current).strip())

    return [c for c in chunks if c]


def _iter_paper_dirs(papers_dir: Path, only: set[str] | None) -> Iterator[Path]:
    for d in sorted(papers_dir.iterdir()):
        if not d.is_dir():
            continue
        if only and d.name not in only:
            continue
        yield d


def _safe_str_list(value: Any) -> list[str]:
    if not isinstance(value, list):
        return []
    return [v for v in value if isinstance(v, str)]


def build_records(*, papers_dir: Path, max_chars: int, overlap: int, only: set[str] | None) -> Iterable[dict[str, Any]]:
    for paper_dir in _iter_paper_dirs(papers_dir, only):
        paper_id = paper_dir.name

        metadata_path = paper_dir / "metadata.json"
        metadata = _read_json(metadata_path) if metadata_path.exists() else {"id": paper_id}

        paper_path = paper_dir / "paper.txt"
        if not paper_path.exists():
            continue

        title = metadata.get("title")
        if not isinstance(title, str) or not title.strip() or title.strip() == paper_id:
            first_line = next((l.strip() for l in _read_text(paper_path).splitlines() if l.strip()), paper_id)
            title = first_line

        year = metadata.get("year") if isinstance(metadata.get("year"), int) else None
        keywords = _safe_str_list(metadata.get("keywords"))
        ssrn_url = metadata.get("ssrn_url")
        if not isinstance(ssrn_url, str) and paper_id.startswith("ssrn-") and paper_id.split("-", 1)[1].isdigit():
            ssrn_url = f"https://papers.ssrn.com/sol3/papers.cfm?abstract_id={paper_id.split('-', 1)[1]}"

        paragraphs = _split_paragraphs(_read_text(paper_path))
        chunks = _chunk_paragraphs(paragraphs, max_chars=max_chars, overlap=overlap)

        for idx, chunk in enumerate(chunks):
            yield {
                "id": f"{paper_id}#chunk={idx}",
                "paper_id": paper_id,
                "title": title,
                "year": year,
                "ssrn_url": ssrn_url,
                "keywords": keywords,
                "source_path": f"papers/{paper_id}/paper.txt",
                "chunk_index": idx,
                "chunk_count": len(chunks),
                "text": chunk,
            }


def main() -> int:
    parser = argparse.ArgumentParser(description="Build JSONL chunks for RAG pipelines.")
    parser.add_argument("--papers-dir", default="papers", help="Papers directory (default: papers).")
    parser.add_argument("--out", default="dist/rag_chunks.jsonl", help="Output JSONL path (default: dist/rag_chunks.jsonl).")
    parser.add_argument("--max-chars", type=int, default=8000, help="Max characters per chunk (default: 8000).")
    parser.add_argument("--overlap", type=int, default=1, help="Paragraph overlap between chunks (default: 1).")
    parser.add_argument("--paper", dest="paper_ids", nargs="*", help="Optional list of paper IDs to include.")
    args = parser.parse_args()

    papers_dir = Path(args.papers_dir)
    if not papers_dir.exists():
        raise SystemExit(f"Papers directory not found: {papers_dir}")

    out_path = Path(args.out)
    out_path.parent.mkdir(parents=True, exist_ok=True)

    only = set(args.paper_ids) if args.paper_ids else None

    records = build_records(
        papers_dir=papers_dir,
        max_chars=args.max_chars,
        overlap=max(0, args.overlap),
        only=only,
    )

    with out_path.open("w", encoding="utf-8", newline="\n") as f:
        for rec in records:
            f.write(json.dumps(rec, ensure_ascii=False))
            f.write("\n")

    print(f"Wrote {out_path}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())

