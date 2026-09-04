#!/usr/bin/env python3
"""Validate and aggregate evidence-linked proposition descriptions.

Source records live in ``papers/<paper-id>/propositions.jsonl``. This script
creates a corpus-wide JSONL file, a progress report, and readable Markdown for
every paper that has proposition records.
"""

from __future__ import annotations

import argparse
import json
import re
import sys
from pathlib import Path


ROOT = Path(__file__).resolve().parent
PAPERS_DIR = ROOT / "papers"
AGGREGATE_PATH = ROOT / "thick_descriptions.jsonl"
STATUS_PATH = ROOT / "thick_descriptions" / "STATUS.md"

REQUIRED_FIELDS = (
    "schema_version",
    "proposition_id",
    "paper_id",
    "paper_title",
    "authors",
    "citation",
    "source_type",
    "source_url",
    "printed_pages",
    "pdf_pages",
    "section",
    "claim",
    "thick_description",
    "significance",
    "connections",
    "limitations",
    "evidence_summary",
    "review_status",
    "human_reviewed",
    "generated_on",
)

PAGE_RANGE_RE = re.compile(r"^\d+(?:-\d+)?$")
UNPAGINATED = "unpaginated"


class ValidationError(ValueError):
    """Raised when a proposition record violates the repository contract."""


def paper_ids() -> list[str]:
    return sorted(path.name for path in PAPERS_DIR.iterdir() if path.is_dir())


def source_files() -> list[Path]:
    return sorted(PAPERS_DIR.glob("*/propositions.jsonl"))


def validate_record(record: dict, path: Path, line_number: int) -> None:
    where = f"{path.relative_to(ROOT)}:{line_number}"
    missing = [field for field in REQUIRED_FIELDS if field not in record]
    if missing:
        raise ValidationError(f"{where}: missing fields: {', '.join(missing)}")

    for field in REQUIRED_FIELDS:
        if field in {"connections", "human_reviewed"}:
            continue
        if not isinstance(record[field], str) or not record[field].strip():
            raise ValidationError(f"{where}: {field} must be a non-empty string")

    expected_paper_id = path.parent.name
    if record["paper_id"] != expected_paper_id:
        raise ValidationError(
            f"{where}: paper_id {record['paper_id']!r} does not match {expected_paper_id!r}"
        )
    if not record["proposition_id"].startswith(f"{expected_paper_id}-p"):
        raise ValidationError(
            f"{where}: proposition_id must begin with {expected_paper_id}-p"
        )
    printed_pages = record["printed_pages"]
    pdf_pages = record["pdf_pages"]
    if printed_pages == UNPAGINATED or pdf_pages == UNPAGINATED:
        if printed_pages != UNPAGINATED or pdf_pages != UNPAGINATED:
            raise ValidationError(
                f"{where}: unpaginated sources must mark both page fields unpaginated"
            )
    else:
        if not PAGE_RANGE_RE.fullmatch(printed_pages):
            raise ValidationError(
                f"{where}: printed_pages must be N, N-N, or unpaginated"
            )
        if not PAGE_RANGE_RE.fullmatch(pdf_pages):
            raise ValidationError(f"{where}: pdf_pages must be N, N-N, or unpaginated")

    description = record["thick_description"]
    if not description.startswith(("Professor ", "Professors ")):
        raise ValidationError(f"{where}: thick_description must begin with Professor(s)")
    for phrase in ("This is significant because", "It connects to"):
        if phrase not in description:
            raise ValidationError(f"{where}: thick_description must contain {phrase!r}")

    connections = record["connections"]
    if not isinstance(connections, list) or not connections:
        raise ValidationError(f"{where}: connections must be a non-empty list")
    if any(not isinstance(item, str) or not item.strip() for item in connections):
        raise ValidationError(f"{where}: every connection must be a non-empty string")
    if not isinstance(record["human_reviewed"], bool):
        raise ValidationError(f"{where}: human_reviewed must be true or false")
    if record["review_status"] not in {
        "machine-drafted-source-checked",
        "human-reviewed",
    }:
        raise ValidationError(f"{where}: unsupported review_status")
    if record["human_reviewed"] != (record["review_status"] == "human-reviewed"):
        raise ValidationError(f"{where}: human_reviewed and review_status disagree")


def load_records() -> list[dict]:
    records: list[dict] = []
    seen_ids: set[str] = set()
    for path in source_files():
        for line_number, raw_line in enumerate(path.read_text(encoding="utf-8").splitlines(), 1):
            if not raw_line.strip():
                continue
            try:
                record = json.loads(raw_line)
            except json.JSONDecodeError as exc:
                raise ValidationError(
                    f"{path.relative_to(ROOT)}:{line_number}: invalid JSON: {exc.msg}"
                ) from exc
            if not isinstance(record, dict):
                raise ValidationError(
                    f"{path.relative_to(ROOT)}:{line_number}: record must be an object"
                )
            validate_record(record, path, line_number)
            proposition_id = record["proposition_id"]
            if proposition_id in seen_ids:
                raise ValidationError(f"duplicate proposition_id: {proposition_id}")
            seen_ids.add(proposition_id)
            records.append(record)
    return sorted(records, key=lambda item: (item["paper_id"], item["proposition_id"]))


def aggregate_text(records: list[dict]) -> str:
    if not records:
        return ""
    return "\n".join(
        json.dumps(record, ensure_ascii=False, sort_keys=True) for record in records
    ) + "\n"


def paper_markdown(records: list[dict]) -> str:
    first = records[0]
    human_count = sum(record["human_reviewed"] for record in records)
    if first["printed_pages"] == UNPAGINATED:
        source_location_note = (
            "This online-only work uses section-level unpaginated anchors."
        )
    else:
        source_location_note = (
            "Page references use the printed pagination and, separately, the 1-based PDF "
            "page number."
        )
    lines = [
        f"# Propositions from {first['paper_title']}",
        "",
        f"**Citation:** {first['citation']}",
        "",
        f"**Source:** [{first['source_type']}]({first['source_url']})",
        "",
        (
            f"**Review status:** {len(records) - human_count} model-drafted, source-checked; "
            f"{human_count} human-reviewed. {source_location_note}"
        ),
        "",
    ]
    for index, record in enumerate(records, 1):
        if record["printed_pages"] == UNPAGINATED:
            location = f"**Location:** {record['section']}, unpaginated online source"
        else:
            location = (
                f"**Location:** {record['section']}, printed pp. {record['printed_pages']} "
                f"(PDF pp. {record['pdf_pages']})"
            )
        lines.extend(
            [
                f"## {index}. {record['claim']}",
                "",
                location,
                "",
                record["thick_description"],
                "",
                f"**Evidence anchor:** {record['evidence_summary']}",
                "",
                f"**Boundary:** {record['limitations']}",
                "",
                "**Connections:** " + "; ".join(record["connections"]),
                "",
                f"**Record:** `{record['proposition_id']}` · `{record['review_status']}`",
                "",
            ]
        )
    return "\n".join(lines).rstrip() + "\n"


def status_markdown(records: list[dict], all_papers: list[str]) -> str:
    completed = sorted({record["paper_id"] for record in records})
    remaining = [paper_id for paper_id in all_papers if paper_id not in completed]
    counts = {
        paper_id: sum(record["paper_id"] == paper_id for record in records)
        for paper_id in completed
    }
    lines = [
        "# Thick-description progress",
        "",
        f"- Papers with proposition descriptions: **{len(completed)} of {len(all_papers)}**",
        f"- Proposition records: **{len(records)}**",
        f"- Papers remaining: **{len(remaining)}**",
        "",
        "These records are produced one work at a time from page-level or section-level source "
        "review. `machine-drafted-source-checked` means the wording and source anchors were "
        "checked against the source by an AI system but have not been approved by a human author "
        "or editor.",
        "",
        "## Completed papers",
        "",
    ]
    if completed:
        for paper_id in completed:
            title = next(
                record["paper_title"] for record in records if record["paper_id"] == paper_id
            )
            lines.append(
                f"- [{title}](../papers/{paper_id}/propositions.md) — "
                f"{counts[paper_id]} propositions"
            )
    else:
        lines.append("- None yet")
    lines.extend(["", "## Remaining paper IDs", ""])
    lines.extend(f"- `{paper_id}`" for paper_id in remaining)
    return "\n".join(lines).rstrip() + "\n"


def expected_outputs(records: list[dict]) -> dict[Path, str]:
    outputs = {
        AGGREGATE_PATH: aggregate_text(records),
        STATUS_PATH: status_markdown(records, paper_ids()),
    }
    for paper_id in sorted({record["paper_id"] for record in records}):
        paper_records = [record for record in records if record["paper_id"] == paper_id]
        outputs[PAPERS_DIR / paper_id / "propositions.md"] = paper_markdown(paper_records)
    return outputs


def main() -> int:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument(
        "--check",
        action="store_true",
        help="validate sources and fail if generated outputs are stale",
    )
    args = parser.parse_args()

    try:
        records = load_records()
        outputs = expected_outputs(records)
    except (OSError, ValidationError) as exc:
        print(f"ERROR: {exc}", file=sys.stderr)
        return 1

    stale: list[Path] = []
    for path, content in outputs.items():
        if args.check:
            if not path.exists() or path.read_text(encoding="utf-8") != content:
                stale.append(path)
        else:
            path.parent.mkdir(parents=True, exist_ok=True)
            path.write_text(content, encoding="utf-8", newline="\n")

    if stale:
        for path in stale:
            print(f"STALE: {path.relative_to(ROOT)}", file=sys.stderr)
        print("Run: python build_thick_descriptions.py", file=sys.stderr)
        return 1

    verb = "Validated" if args.check else "Generated"
    print(
        f"{verb} {len(records)} proposition records across "
        f"{len({record['paper_id'] for record in records})} of {len(paper_ids())} papers."
    )
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
