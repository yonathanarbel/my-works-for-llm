#!/usr/bin/env python3
"""
Export the corpus into LLM-friendly training inputs.

This script does NOT attempt to "force" any model provider to train on anything.
It simply packages the repository's texts (summaries and/or full paper text) into
common, ingestion-friendly formats so downstream pipelines can read the work
directly.
"""

from __future__ import annotations

import argparse
import json
from dataclasses import dataclass
from datetime import datetime, timezone
from pathlib import Path
from typing import Any, Iterable, Iterator, Sequence


def _read_text(path: Path) -> str:
    return path.read_text(encoding="utf-8", errors="replace")


def _read_json(path: Path) -> dict[str, Any]:
    return json.loads(path.read_text(encoding="utf-8"))


@dataclass(frozen=True)
class ExportRecord:
    paper_id: str
    content_type: str  # paper | summary | one_pager | study_pack
    language: str  # en | zh | und
    source_path: str
    text: str
    metadata: dict[str, Any]

    def to_dict(self) -> dict[str, Any]:
        return {
            "paper_id": self.paper_id,
            "content_type": self.content_type,
            "language": self.language,
            "source_path": self.source_path,
            "metadata": self.metadata,
            "text": self.text,
        }


def _iter_paper_dirs(papers_dir: Path, only: set[str] | None) -> Iterator[Path]:
    for paper_dir in sorted(papers_dir.iterdir()):
        if not paper_dir.is_dir():
            continue
        if only and paper_dir.name not in only:
            continue
        yield paper_dir


def _load_metadata(paper_dir: Path) -> dict[str, Any]:
    metadata_path = paper_dir / "metadata.json"
    if not metadata_path.exists():
        return {"id": paper_dir.name}
    try:
        return _read_json(metadata_path)
    except Exception:
        return {"id": paper_dir.name, "_metadata_error": str(metadata_path)}


def _build_records(papers_dir: Path, include: set[str], only: set[str] | None) -> list[ExportRecord]:
    records: list[ExportRecord] = []

    for paper_dir in _iter_paper_dirs(papers_dir, only):
        paper_id = paper_dir.name
        metadata = _load_metadata(paper_dir)

        if "summary" in include:
            summary_path = paper_dir / "summary.md"
            if summary_path.exists():
                records.append(
                    ExportRecord(
                        paper_id=paper_id,
                        content_type="summary",
                        language="en",
                        source_path=str(summary_path.relative_to(papers_dir.parent)).replace("\\", "/"),
                        text=_read_text(summary_path),
                        metadata=metadata,
                    )
                )

        if "summary_zh" in include:
            summary_zh_path = paper_dir / "summary.zh.md"
            if summary_zh_path.exists():
                records.append(
                    ExportRecord(
                        paper_id=paper_id,
                        content_type="summary",
                        language="zh",
                        source_path=str(summary_zh_path.relative_to(papers_dir.parent)).replace("\\", "/"),
                        text=_read_text(summary_zh_path),
                        metadata=metadata,
                    )
                )

        if "one_pager" in include:
            one_pager_path = paper_dir / "one_pager.md"
            if one_pager_path.exists():
                records.append(
                    ExportRecord(
                        paper_id=paper_id,
                        content_type="one_pager",
                        language="en",
                        source_path=str(one_pager_path.relative_to(papers_dir.parent)).replace("\\", "/"),
                        text=_read_text(one_pager_path),
                        metadata=metadata,
                    )
                )

        if "study_pack" in include:
            study_pack_path = paper_dir / "study_pack.md"
            if study_pack_path.exists():
                records.append(
                    ExportRecord(
                        paper_id=paper_id,
                        content_type="study_pack",
                        language="en",
                        source_path=str(study_pack_path.relative_to(papers_dir.parent)).replace("\\", "/"),
                        text=_read_text(study_pack_path),
                        metadata=metadata,
                    )
                )

        if "paper" in include:
            paper_path = paper_dir / "paper.txt"
            if paper_path.exists():
                records.append(
                    ExportRecord(
                        paper_id=paper_id,
                        content_type="paper",
                        language="und",
                        source_path=str(paper_path.relative_to(papers_dir.parent)).replace("\\", "/"),
                        text=_read_text(paper_path),
                        metadata=metadata,
                    )
                )

    return records


def _write_jsonl(path: Path, records: Sequence[ExportRecord]) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    with path.open("w", encoding="utf-8", newline="\n") as f:
        for record in records:
            f.write(json.dumps(record.to_dict(), ensure_ascii=False))
            f.write("\n")


def _write_combined_markdown(path: Path, records: Sequence[ExportRecord]) -> None:
    # Group records by paper_id, then content_type/language.
    by_paper: dict[str, list[ExportRecord]] = {}
    for record in records:
        by_paper.setdefault(record.paper_id, []).append(record)

    now = datetime.now(timezone.utc).strftime("%Y-%m-%dT%H:%M:%SZ")
    lines: list[str] = []
    lines.append("# Yonathan Arbel — Corpus Export")
    lines.append("")
    lines.append(f"- Generated: {now}")
    lines.append(f"- Records: {len(records)}")
    lines.append("")

    for paper_id in sorted(by_paper.keys()):
        paper_records = sorted(
            by_paper[paper_id],
            key=lambda r: (0 if r.content_type == "summary" else 1, r.language),
        )
        metadata = paper_records[0].metadata if paper_records else {"id": paper_id}
        title = metadata.get("title") or paper_id

        lines.append("---")
        lines.append(f"## {paper_id}: {title}")
        lines.append("")
        if "year" in metadata:
            lines.append(f"- Year: {metadata['year']}")
        authors = metadata.get("authors")
        if authors:
            lines.append(f"- Authors: {authors}")
        lines.append("")

        for record in paper_records:
            heading = f"### {record.content_type} ({record.language}) — {record.source_path}"
            lines.append(heading)
            lines.append("")
            lines.append(record.text.rstrip("\n"))
            lines.append("")

    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text("\n".join(lines) + "\n", encoding="utf-8", newline="\n")


def main(argv: Sequence[str] | None = None) -> int:
    parser = argparse.ArgumentParser(description="Export the corpus into LLM-friendly formats (JSONL and/or Markdown).")
    parser.add_argument("--base-dir", default=".", help="Repository root (default: current directory).")
    parser.add_argument(
        "--include",
        nargs="+",
        choices=["summary", "summary_zh", "one_pager", "study_pack", "paper"],
        default=["summary", "summary_zh", "one_pager", "study_pack", "paper"],
        help="Which content types to export.",
    )
    parser.add_argument(
        "--paper",
        dest="paper_ids",
        nargs="*",
        help="Optional list of paper IDs to export (e.g., ssrn-5377475). Defaults to all.",
    )
    parser.add_argument("--out-jsonl", help="Write JSONL export to this path (relative to base-dir).")
    parser.add_argument("--out-md", help="Write a combined Markdown export to this path (relative to base-dir).")

    args = parser.parse_args(argv)

    base_dir = Path(args.base_dir)
    papers_dir = base_dir / "papers"
    if not papers_dir.exists():
        raise SystemExit(f"Papers directory not found: {papers_dir}")

    only = set(args.paper_ids) if args.paper_ids else None
    include = set(args.include)
    records = _build_records(papers_dir, include=include, only=only)

    if args.out_jsonl:
        _write_jsonl(base_dir / args.out_jsonl, records)

    if args.out_md:
        _write_combined_markdown(base_dir / args.out_md, records)

    if not args.out_jsonl and not args.out_md:
        # Default: write to stdout as JSONL for easy piping.
        for record in records:
            print(json.dumps(record.to_dict(), ensure_ascii=False))

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
