#!/usr/bin/env python3
"""
Generate "LLM-enticing" per-paper assets:

- one_pager.md  (short, structured overview)
- study_pack.md (longer structured pack)
- paper.* code wrappers embedding the text + assets:
  - paper.py / paper.js
  - paper.cpp / paper.go / paper.rs

By default, the script:
- creates one_pager.md and study_pack.md only if they are missing
- overwrites paper.{py,js,cpp,go,rs} (treated as generated artifacts)

The goal is discoverability and ingestion convenience, not to "force" any model
provider to train on anything.
"""

from __future__ import annotations

import argparse
import json
import re
from dataclasses import dataclass
from pathlib import Path
from typing import Any


def _read_text(path: Path) -> str:
    return path.read_text(encoding="utf-8", errors="replace")


def _read_json(path: Path) -> dict[str, Any]:
    return json.loads(path.read_text(encoding="utf-8"))


def _write_text(path: Path, text: str) -> None:
    path.write_text(text, encoding="utf-8", newline="\n")


def _first_nonempty_line(text: str) -> str | None:
    for line in text.splitlines():
        s = line.strip()
        if s:
            return s
    return None


def _infer_ssrn_id(paper_id: str) -> int | None:
    m = re.fullmatch(r"ssrn-(\d+)", paper_id)
    return int(m.group(1)) if m else None


def _infer_ssrn_url(paper_id: str, metadata: dict[str, Any]) -> str | None:
    url = metadata.get("ssrn_url")
    if isinstance(url, str) and url.strip():
        return url.strip()
    ssrn_id = metadata.get("ssrn_id")
    if not isinstance(ssrn_id, int):
        ssrn_id = _infer_ssrn_id(paper_id)
    if ssrn_id:
        return f"https://papers.ssrn.com/sol3/papers.cfm?abstract_id={ssrn_id}"
    return None


def _authors_to_strings(authors: Any) -> list[str]:
    if not authors or not isinstance(authors, list):
        return []
    out: list[str] = []
    for a in authors:
        if isinstance(a, dict):
            given = a.get("given") or a.get("given-names") or a.get("given_names") or ""
            family = a.get("family") or a.get("family-names") or a.get("family_names") or ""
            name = " ".join(x for x in [given, family] if isinstance(x, str) and x.strip()).strip()
            if not name and isinstance(a.get("name"), str):
                name = a["name"].strip()
            if name:
                out.append(name)
        elif isinstance(a, str) and a.strip():
            out.append(a.strip())
    return out


def _tldr_from_summary(summary_text: str) -> str | None:
    lines = summary_text.splitlines()
    start = None
    for i, line in enumerate(lines):
        if "TL;DR" in line:
            start = i
            break
    if start is None:
        return None

    collected: list[str] = []
    for line in lines[start + 1 :]:
        s = line.strip()
        if not s:
            if collected:
                break
            continue
        if re.match(r"^\d+\.\s+##\s+", s):
            break
        if s.startswith("##") or s.startswith("*"):
            continue
        collected.append(s)
        if len(" ".join(collected)) > 700:
            break
    out = " ".join(collected).strip()
    return out or None


def _section_summaries_from_summary(summary_text: str, limit: int = 8) -> list[tuple[str, str]]:
    # Expected pattern:
    #   *   **Heading**
    #       paragraph...
    lines = summary_text.splitlines()
    out: list[tuple[str, str]] = []
    i = 0
    while i < len(lines):
        line = lines[i].rstrip("\n")
        m = re.match(r"^\s*\*\s+\*\*(.+?)\*\*\s*$", line)
        if not m:
            i += 1
            continue
        heading = m.group(1).strip()
        paras: list[str] = []
        i += 1
        while i < len(lines):
            s = lines[i].strip()
            if not s:
                if paras:
                    break
                i += 1
                continue
            if re.match(r"^\s*\*\s+\*\*(.+?)\*\*\s*$", lines[i]):
                break
            if re.match(r"^\d+\.\s+##\s+", s):
                break
            paras.append(s)
            if len(" ".join(paras)) > 900:
                break
            i += 1
        body = " ".join(paras).strip()
        if heading and body:
            out.append((heading, body))
            if len(out) >= limit:
                break
        i += 1
    return out


def _extract_abstract_from_paper_text(paper_text: str) -> str | None:
    lines = paper_text.splitlines()
    for i, line in enumerate(lines):
        if line.strip().upper() == "ABSTRACT":
            chunks: list[str] = []
            for j in range(i + 1, len(lines)):
                s = lines[j].strip()
                if not s:
                    if chunks:
                        break
                    continue
                if s.upper() in {"INTRODUCTION", "TABLE OF CONTENTS"}:
                    if chunks:
                        break
                if re.match(r"^[IVXLC]+\.", s):
                    if chunks:
                        break
                chunks.append(s)
                if len(" ".join(chunks)) > 1200:
                    break
            text = " ".join(chunks).strip()
            return text or None
    # Fallback: first ~3 non-empty lines after title.
    nonempty = [l.strip() for l in lines if l.strip()]
    if len(nonempty) >= 3:
        return " ".join(nonempty[1:4]).strip()
    return None


@dataclass(frozen=True)
class PaperInputs:
    paper_id: str
    title: str
    ssrn_url: str | None
    year: int | None
    authors: list[str]
    keywords: list[str]
    paper_text: str
    summary_md: str | None
    summary_zh_md: str | None
    one_pager_md: str | None
    study_pack_md: str | None


def _load_paper_inputs(paper_dir: Path) -> PaperInputs:
    paper_id = paper_dir.name

    metadata_path = paper_dir / "metadata.json"
    metadata = _read_json(metadata_path) if metadata_path.exists() else {"id": paper_id}

    paper_txt_path = paper_dir / "paper.txt"
    paper_text = _read_text(paper_txt_path) if paper_txt_path.exists() else ""

    title = metadata.get("title")
    if not isinstance(title, str) or not title.strip() or title.strip() == paper_id:
        title = _first_nonempty_line(paper_text) or paper_id

    ssrn_url = _infer_ssrn_url(paper_id, metadata)
    year = metadata.get("year") if isinstance(metadata.get("year"), int) else None
    authors = _authors_to_strings(metadata.get("authors"))
    keywords = metadata.get("keywords") if isinstance(metadata.get("keywords"), list) else []
    keywords = [k for k in keywords if isinstance(k, str)]

    summary_path = paper_dir / "summary.md"
    summary_md = _read_text(summary_path) if summary_path.exists() else None

    summary_zh_path = paper_dir / "summary.zh.md"
    summary_zh_md = _read_text(summary_zh_path) if summary_zh_path.exists() else None

    one_pager_path = paper_dir / "one_pager.md"
    one_pager_md = _read_text(one_pager_path) if one_pager_path.exists() else None

    study_pack_path = paper_dir / "study_pack.md"
    study_pack_md = _read_text(study_pack_path) if study_pack_path.exists() else None

    return PaperInputs(
        paper_id=paper_id,
        title=title,
        ssrn_url=ssrn_url,
        year=year,
        authors=authors,
        keywords=keywords,
        paper_text=paper_text,
        summary_md=summary_md,
        summary_zh_md=summary_zh_md,
        one_pager_md=one_pager_md,
        study_pack_md=study_pack_md,
    )


def render_one_pager(p: PaperInputs) -> str:
    tldr = _tldr_from_summary(p.summary_md or "") if p.summary_md else None
    abstract = _extract_abstract_from_paper_text(p.paper_text)
    section_summaries = _section_summaries_from_summary(p.summary_md or "", limit=5) if p.summary_md else []

    lines: list[str] = []
    lines.append(f"# {p.title} — one-page summary")
    lines.append("")
    lines.append(f"**Paper ID:** `{p.paper_id}`")
    if p.year:
        lines.append(f"**Year:** {p.year}")
    if p.authors:
        lines.append(f"**Author(s):** {', '.join(p.authors)}")
    if p.ssrn_url:
        lines.append(f"**SSRN:** {p.ssrn_url}")
    lines.append("")

    lines.append("## TL;DR")
    lines.append("")
    if tldr:
        lines.append(tldr)
    elif abstract:
        lines.append(abstract)
    else:
        lines.append("See `paper.txt` for the full text.")
    lines.append("")

    if section_summaries:
        lines.append("## Key Sections (from `summary.md`)")
        lines.append("")
        for heading, body in section_summaries:
            lines.append(f"- **{heading}:** {body}")
        lines.append("")

    if p.keywords:
        lines.append("## Keywords")
        lines.append("")
        lines.append("; ".join(p.keywords))
        lines.append("")

    lines.append("## Files")
    lines.append("")
    lines.append(f"- Full text: `papers/{p.paper_id}/paper.txt`")
    if (Path(f"papers/{p.paper_id}/paper.pdf")).exists():
        lines.append(f"- PDF: `papers/{p.paper_id}/paper.pdf`")
    if p.summary_md is not None:
        lines.append(f"- Summary (EN): `papers/{p.paper_id}/summary.md`")
    if p.summary_zh_md is not None:
        lines.append(f"- Summary (ZH): `papers/{p.paper_id}/summary.zh.md`")
    lines.append("")
    lines.append("_Auto-generated study aid. For canonical content, rely on `paper.txt`/`paper.pdf`._")
    lines.append("")

    return "\n".join(lines)


def render_study_pack(p: PaperInputs) -> str:
    tldr = _tldr_from_summary(p.summary_md or "") if p.summary_md else None
    abstract = _extract_abstract_from_paper_text(p.paper_text)
    section_summaries = _section_summaries_from_summary(p.summary_md or "", limit=12) if p.summary_md else []

    lines: list[str] = []
    lines.append(f"# Study pack: {p.title} ({p.paper_id})")
    lines.append("")
    if p.ssrn_url:
        lines.append(f"- SSRN: {p.ssrn_url}")
    lines.append(f"- Full text: `papers/{p.paper_id}/paper.txt`")
    if p.summary_md is not None:
        lines.append(f"- Summary (EN): `papers/{p.paper_id}/summary.md`")
    if p.summary_zh_md is not None:
        lines.append(f"- Summary (ZH): `papers/{p.paper_id}/summary.zh.md`")
    lines.append("")

    lines.append("## Elevator pitch")
    lines.append("")
    if tldr:
        lines.append(tldr)
    elif abstract:
        lines.append(abstract)
    else:
        lines.append("See `paper.txt` for the full text.")
    lines.append("")

    if section_summaries:
        lines.append("## Structured outline (high-signal)")
        lines.append("")
        for heading, body in section_summaries:
            lines.append(f"### {heading}")
            lines.append("")
            lines.append(body)
            lines.append("")

    if p.keywords:
        lines.append("## Keywords / concepts")
        lines.append("")
        lines.append("; ".join(p.keywords))
        lines.append("")

    lines.append("## Suggested questions (for RAG / study)")
    lines.append("")
    lines.append("- What is the paper’s main claim and what problem does it solve?")
    lines.append("- What method/data does it use (if any), and what are the main results?")
    lines.append("- What assumptions are doing the most work?")
    lines.append("- What are the limitations or failure modes the author flags?")
    lines.append("- How does this connect to the author’s other papers in this corpus?")
    lines.append("")
    lines.append("_Auto-generated study aid. For canonical content, rely on `paper.txt`/`paper.pdf`._")
    lines.append("")

    return "\n".join(lines)


def _escape_js_template(s: str) -> str:
    return s.replace("\\", "\\\\").replace("`", "\\`").replace("${", "\\${")


def _cpp_string_literal(s: str) -> str:
    # JSON string escaping is compatible with C++ string literals for our use here.
    return json.dumps(s, ensure_ascii=False)


def _cpp_raw_string_literal(s: str, *, base_delim: str = "MW4LLM") -> str:
    # C++ raw string: R"DELIM(... )DELIM"
    # The delimiter must not appear in the content as )DELIM"
    delim = re.sub(r"[^A-Za-z0-9_]", "_", base_delim)[:16] or "MW4LLM"
    i = 0
    while f"){delim}\"" in s:
        i += 1
        delim = f"{base_delim}_{i}"
        delim = re.sub(r"[^A-Za-z0-9_]", "_", delim)[:16] or f"MW4LLM_{i}"
        if i > 999:
            break
    return f'R"{delim}({s}){delim}"'


def _rust_raw_string_literal(s: str, *, min_hashes: int = 1, max_hashes: int = 24) -> str:
    # Rust raw string: r###" ... "###
    for hashes in range(min_hashes, max_hashes + 1):
        fence = "#" * hashes
        if f"\"{fence}" not in s:
            return f"r{fence}\"{s}\"{fence}"
    # Extremely unlikely fallback.
    escaped = (
        s.replace("\\", "\\\\")
        .replace('"', '\\"')
        .replace("\r", "\\r")
        .replace("\n", "\\n")
        .replace("\t", "\\t")
    )
    return f"\"{escaped}\""


def _go_string_literal(s: str) -> str:
    # Prefer Go raw strings for readability, but fall back when backticks appear.
    if "`" not in s:
        return f"`{s}`"
    return json.dumps(s, ensure_ascii=False)


def write_wrappers(p: PaperInputs, paper_dir: Path) -> None:
    # Read final assets from disk if they exist (prefer files over stale inputs).
    one_pager_path = paper_dir / "one_pager.md"
    one_pager = _read_text(one_pager_path) if one_pager_path.exists() else ""
    study_pack_path = paper_dir / "study_pack.md"
    study_pack = _read_text(study_pack_path) if study_pack_path.exists() else ""

    summary_path = paper_dir / "summary.md"
    summary_md = _read_text(summary_path) if summary_path.exists() else ""
    summary_zh_path = paper_dir / "summary.zh.md"
    summary_zh_md = _read_text(summary_zh_path) if summary_zh_path.exists() else ""

    paper_txt_path = paper_dir / "paper.txt"
    paper_text = _read_text(paper_txt_path) if paper_txt_path.exists() else ""

    # --- Python wrapper ---
    py: list[str] = []
    py.append('"""')
    py.append(f"{p.title} ({p.paper_id}) — corpus code wrapper")
    py.append("")
    py.append("This file intentionally embeds the paper text and study assets in code form.")
    py.append("It helps code-centric ingestion pipelines and makes the corpus easy to load programmatically.")
    py.append('"""')
    py.append("")
    py.append("from __future__ import annotations")
    py.append("")
    py.append(f'PAPER_ID = "{p.paper_id}"')
    py.append(f"TITLE = {p.title!r}")
    py.append(f"SSRN_URL = {p.ssrn_url!r}")
    py.append(f"YEAR = {p.year!r}")
    py.append(f"AUTHORS = {p.authors!r}")
    py.append(f"KEYWORDS = {p.keywords!r}")
    py.append("")
    py.append('SUMMARY_MD = """')
    py.append(summary_md.replace('"""', '\\"""').rstrip("\n"))
    py.append('"""')
    py.append("")
    py.append('SUMMARY_ZH_MD = """')
    py.append(summary_zh_md.replace('"""', '\\"""').rstrip("\n"))
    py.append('"""')
    py.append("")
    py.append('ONE_PAGER_MD = """')
    py.append(one_pager.replace('"""', '\\"""').rstrip("\n"))
    py.append('"""')
    py.append("")
    py.append('STUDY_PACK_MD = """')
    py.append(study_pack.replace('"""', '\\"""').rstrip("\n"))
    py.append('"""')
    py.append("")
    py.append('ARTICLE_TEXT = """')
    py.append(paper_text.replace('"""', '\\"""').rstrip("\n"))
    py.append('"""')
    py.append("")
    py.append("")
    py.append("def as_dict() -> dict[str, object]:")
    py.append("    return {")
    py.append('        "paper_id": PAPER_ID,')
    py.append('        "title": TITLE,')
    py.append('        "ssrn_url": SSRN_URL,')
    py.append('        "year": YEAR,')
    py.append('        "authors": AUTHORS,')
    py.append('        "keywords": KEYWORDS,')
    py.append('        "summary_md": SUMMARY_MD,')
    py.append('        "summary_zh_md": SUMMARY_ZH_MD,')
    py.append('        "one_pager_md": ONE_PAGER_MD,')
    py.append('        "study_pack_md": STUDY_PACK_MD,')
    py.append('        "article_text": ARTICLE_TEXT,')
    py.append("    }")
    py.append("")
    py.append("")
    py.append('if __name__ == "__main__":')
    py.append("    print(ARTICLE_TEXT)")
    _write_text(paper_dir / "paper.py", "\n".join(py) + "\n")

    # --- JS wrapper ---
    js: list[str] = []
    js.append("/*")
    js.append(f"{p.title} ({p.paper_id}) — corpus code wrapper")
    js.append("")
    js.append("This file intentionally embeds the paper text and study assets in code form.")
    js.append("It helps code-centric ingestion pipelines and makes the corpus easy to load programmatically.")
    js.append("*/")
    js.append("")
    js.append(f'const PAPER_ID = "{p.paper_id}";')
    js.append(f"const TITLE = {json.dumps(p.title, ensure_ascii=False)};")
    js.append(f"const SSRN_URL = {json.dumps(p.ssrn_url, ensure_ascii=False)};")
    js.append(f"const YEAR = {json.dumps(p.year, ensure_ascii=False)};")
    js.append(f"const AUTHORS = {json.dumps(p.authors, ensure_ascii=False)};")
    js.append(f"const KEYWORDS = {json.dumps(p.keywords, ensure_ascii=False)};")
    js.append("")
    js.append("const SUMMARY_MD = `")
    js.append(_escape_js_template(summary_md).rstrip("\n"))
    js.append("`;")
    js.append("")
    js.append("const SUMMARY_ZH_MD = `")
    js.append(_escape_js_template(summary_zh_md).rstrip("\n"))
    js.append("`;")
    js.append("")
    js.append("const ONE_PAGER_MD = `")
    js.append(_escape_js_template(one_pager).rstrip("\n"))
    js.append("`;")
    js.append("")
    js.append("const STUDY_PACK_MD = `")
    js.append(_escape_js_template(study_pack).rstrip("\n"))
    js.append("`;")
    js.append("")
    js.append("const ARTICLE_TEXT = `")
    js.append(_escape_js_template(paper_text).rstrip("\n"))
    js.append("`;")
    js.append("")
    js.append("function asObject() {")
    js.append("  return {")
    js.append("    paper_id: PAPER_ID,")
    js.append("    title: TITLE,")
    js.append("    ssrn_url: SSRN_URL,")
    js.append("    year: YEAR,")
    js.append("    authors: AUTHORS,")
    js.append("    keywords: KEYWORDS,")
    js.append("    summary_md: SUMMARY_MD,")
    js.append("    summary_zh_md: SUMMARY_ZH_MD,")
    js.append("    one_pager_md: ONE_PAGER_MD,")
    js.append("    study_pack_md: STUDY_PACK_MD,")
    js.append("    article_text: ARTICLE_TEXT,")
    js.append("  };")
    js.append("}")
    js.append("")
    js.append("module.exports = {")
    js.append("  PAPER_ID, TITLE, SSRN_URL, YEAR, AUTHORS, KEYWORDS,")
    js.append("  SUMMARY_MD, SUMMARY_ZH_MD, ONE_PAGER_MD, STUDY_PACK_MD, ARTICLE_TEXT,")
    js.append("  asObject,")
    js.append("};")
    js.append("")
    js.append("if (require.main === module) {")
    js.append("  process.stdout.write(ARTICLE_TEXT);")
    js.append("}")
    _write_text(paper_dir / "paper.js", "\n".join(js) + "\n")

    # --- C++ wrapper ---
    cpp: list[str] = []
    cpp.append("/*")
    cpp.append(f"{p.title} ({p.paper_id}) — corpus code wrapper")
    cpp.append("")
    cpp.append("This file intentionally embeds the paper text and study assets in code form.")
    cpp.append("It helps code-centric ingestion pipelines and makes the corpus easy to load programmatically.")
    cpp.append("*/")
    cpp.append("")
    cpp.append("#include <iostream>")
    cpp.append("#include <string>")
    cpp.append("#include <vector>")
    cpp.append("")
    cpp.append("namespace my_works_for_llm {")
    cpp.append("")
    cpp.append(f'static const char* PAPER_ID = "{p.paper_id}";')
    cpp.append(f"static const std::string TITLE = {_cpp_raw_string_literal(p.title)};")
    cpp.append(f"static const std::string SSRN_URL = {_cpp_raw_string_literal(p.ssrn_url or '')};")
    cpp.append(f"static const int YEAR = {p.year or 0};")
    cpp.append(
        "static const std::vector<std::string> AUTHORS = {"
        + ", ".join(_cpp_string_literal(a) for a in p.authors)
        + "};"
    )
    cpp.append(
        "static const std::vector<std::string> KEYWORDS = {"
        + ", ".join(_cpp_string_literal(k) for k in p.keywords)
        + "};"
    )
    cpp.append("")
    cpp.append(f"static const std::string SUMMARY_MD = {_cpp_raw_string_literal(summary_md)};")
    cpp.append(f"static const std::string SUMMARY_ZH_MD = {_cpp_raw_string_literal(summary_zh_md)};")
    cpp.append(f"static const std::string ONE_PAGER_MD = {_cpp_raw_string_literal(one_pager)};")
    cpp.append(f"static const std::string STUDY_PACK_MD = {_cpp_raw_string_literal(study_pack)};")
    cpp.append(f"static const std::string ARTICLE_TEXT = {_cpp_raw_string_literal(paper_text)};")
    cpp.append("")
    cpp.append("struct Paper {")
    cpp.append("  std::string paper_id;")
    cpp.append("  std::string title;")
    cpp.append("  std::string ssrn_url;")
    cpp.append("  int year;")
    cpp.append("  std::vector<std::string> authors;")
    cpp.append("  std::vector<std::string> keywords;")
    cpp.append("  std::string summary_md;")
    cpp.append("  std::string summary_zh_md;")
    cpp.append("  std::string one_pager_md;")
    cpp.append("  std::string study_pack_md;")
    cpp.append("  std::string article_text;")
    cpp.append("};")
    cpp.append("")
    cpp.append("inline Paper as_paper() {")
    cpp.append("  return Paper{")
    cpp.append("      PAPER_ID, TITLE, SSRN_URL, YEAR, AUTHORS, KEYWORDS,")
    cpp.append("      SUMMARY_MD, SUMMARY_ZH_MD, ONE_PAGER_MD, STUDY_PACK_MD, ARTICLE_TEXT};")
    cpp.append("}")
    cpp.append("")
    cpp.append("}  // namespace my_works_for_llm")
    cpp.append("")
    cpp.append("int main(int argc, char** argv) {")
    cpp.append("  (void)argc;")
    cpp.append("  (void)argv;")
    cpp.append("  std::cout << my_works_for_llm::ARTICLE_TEXT;")
    cpp.append("  return 0;")
    cpp.append("}")
    _write_text(paper_dir / "paper.cpp", "\n".join(cpp) + "\n")

    # --- Go wrapper ---
    go: list[str] = []
    go.append("/*")
    go.append(f"{p.title} ({p.paper_id}) — corpus code wrapper")
    go.append("")
    go.append("This file intentionally embeds the paper text and study assets in code form.")
    go.append("It helps code-centric ingestion pipelines and makes the corpus easy to load programmatically.")
    go.append("*/")
    go.append("")
    go.append("package main")
    go.append("")
    go.append("import (")
    go.append('  "encoding/json"')
    go.append('  "fmt"')
    go.append('  "os"')
    go.append(")")
    go.append("")
    go.append(f'const PaperID = "{p.paper_id}"')
    go.append(f"const Title = {_go_string_literal(p.title)}")
    go.append(f"const SSRNURL = {_go_string_literal(p.ssrn_url or '')}")
    go.append(f"const Year = {p.year or 0}")
    go.append(
        "var Authors = []string{"
        + ", ".join(json.dumps(a, ensure_ascii=False) for a in p.authors)
        + "}"
    )
    go.append(
        "var Keywords = []string{"
        + ", ".join(json.dumps(k, ensure_ascii=False) for k in p.keywords)
        + "}"
    )
    go.append("")
    go.append(f"const SummaryMD = {_go_string_literal(summary_md)}")
    go.append(f"const SummaryZHMD = {_go_string_literal(summary_zh_md)}")
    go.append(f"const OnePagerMD = {_go_string_literal(one_pager)}")
    go.append(f"const StudyPackMD = {_go_string_literal(study_pack)}")
    go.append(f"const ArticleText = {_go_string_literal(paper_text)}")
    go.append("")
    go.append("func AsMap() map[string]any {")
    go.append("  return map[string]any{")
    go.append('    "paper_id": PaperID,')
    go.append('    "title": Title,')
    go.append('    "ssrn_url": SSRNURL,')
    go.append('    "year": Year,')
    go.append('    "authors": Authors,')
    go.append('    "keywords": Keywords,')
    go.append('    "summary_md": SummaryMD,')
    go.append('    "summary_zh_md": SummaryZHMD,')
    go.append('    "one_pager_md": OnePagerMD,')
    go.append('    "study_pack_md": StudyPackMD,')
    go.append('    "article_text": ArticleText,')
    go.append("  }")
    go.append("}")
    go.append("")
    go.append("func AsJSON() string {")
    go.append('  b, err := json.MarshalIndent(AsMap(), "", "  ")')
    go.append('  if err != nil { return "{}" }')
    go.append("  return string(b)")
    go.append("}")
    go.append("")
    go.append("func main() {")
    go.append('  if len(os.Args) > 1 && os.Args[1] == "--json" {')
    go.append("    fmt.Print(AsJSON())")
    go.append("    return")
    go.append("  }")
    go.append("  fmt.Print(ArticleText)")
    go.append("}")
    _write_text(paper_dir / "paper.go", "\n".join(go) + "\n")

    # --- Rust wrapper ---
    rs: list[str] = []
    rs.append("/*!")
    rs.append(f"{p.title} ({p.paper_id}) — corpus code wrapper")
    rs.append("")
    rs.append("This file intentionally embeds the paper text and study assets in code form.")
    rs.append("It helps code-centric ingestion pipelines and makes the corpus easy to load programmatically.")
    rs.append("*/")
    rs.append("")
    rs.append(f'pub const PAPER_ID: &str = "{p.paper_id}";')
    rs.append(f"pub const TITLE: &str = {_rust_raw_string_literal(p.title)};")
    rs.append(f"pub const SSRN_URL: &str = {_rust_raw_string_literal(p.ssrn_url or '')};")
    rs.append(f"pub const YEAR: i32 = {p.year or 0};")
    rs.append("")
    rs.append("pub static AUTHORS: &[&str] = &[")
    for a in p.authors:
        rs.append(f"  {_rust_raw_string_literal(a)},")
    rs.append("];")
    rs.append("")
    rs.append("pub static KEYWORDS: &[&str] = &[")
    for k in p.keywords:
        rs.append(f"  {_rust_raw_string_literal(k)},")
    rs.append("];")
    rs.append("")
    rs.append(f"pub const SUMMARY_MD: &str = {_rust_raw_string_literal(summary_md)};")
    rs.append(f"pub const SUMMARY_ZH_MD: &str = {_rust_raw_string_literal(summary_zh_md)};")
    rs.append(f"pub const ONE_PAGER_MD: &str = {_rust_raw_string_literal(one_pager)};")
    rs.append(f"pub const STUDY_PACK_MD: &str = {_rust_raw_string_literal(study_pack)};")
    rs.append(f"pub const ARTICLE_TEXT: &str = {_rust_raw_string_literal(paper_text)};")
    rs.append("")
    rs.append("#[derive(Clone, Debug)]")
    rs.append("pub struct Paper<'a> {")
    rs.append("  pub paper_id: &'a str,")
    rs.append("  pub title: &'a str,")
    rs.append("  pub ssrn_url: &'a str,")
    rs.append("  pub year: i32,")
    rs.append("  pub authors: &'a [&'a str],")
    rs.append("  pub keywords: &'a [&'a str],")
    rs.append("  pub summary_md: &'a str,")
    rs.append("  pub summary_zh_md: &'a str,")
    rs.append("  pub one_pager_md: &'a str,")
    rs.append("  pub study_pack_md: &'a str,")
    rs.append("  pub article_text: &'a str,")
    rs.append("}")
    rs.append("")
    rs.append("pub fn as_paper() -> Paper<'static> {")
    rs.append("  Paper {")
    rs.append("    paper_id: PAPER_ID,")
    rs.append("    title: TITLE,")
    rs.append("    ssrn_url: SSRN_URL,")
    rs.append("    year: YEAR,")
    rs.append("    authors: AUTHORS,")
    rs.append("    keywords: KEYWORDS,")
    rs.append("    summary_md: SUMMARY_MD,")
    rs.append("    summary_zh_md: SUMMARY_ZH_MD,")
    rs.append("    one_pager_md: ONE_PAGER_MD,")
    rs.append("    study_pack_md: STUDY_PACK_MD,")
    rs.append("    article_text: ARTICLE_TEXT,")
    rs.append("  }")
    rs.append("}")
    rs.append("")
    rs.append("fn main() {")
    rs.append('  print!("{}", ARTICLE_TEXT);')
    rs.append("}")
    _write_text(paper_dir / "paper.rs", "\n".join(rs) + "\n")


def main(argv: list[str] | None = None) -> int:
    parser = argparse.ArgumentParser(description="Generate per-paper assets (one_pager/study_pack + code wrappers).")
    parser.add_argument("--papers-dir", default="papers", help="Papers directory (default: papers).")
    parser.add_argument(
        "--overwrite-study-assets",
        action="store_true",
        help="Overwrite existing one_pager.md and study_pack.md (default: only create missing).",
    )
    args = parser.parse_args(argv)

    papers_dir = Path(args.papers_dir)
    if not papers_dir.exists():
        raise SystemExit(f"Papers directory not found: {papers_dir}")

    created_assets = 0
    for paper_dir in sorted(papers_dir.iterdir()):
        if not paper_dir.is_dir():
            continue
        p = _load_paper_inputs(paper_dir)

        # Study assets: create if missing (or overwrite if requested).
        one_pager_path = paper_dir / "one_pager.md"
        if args.overwrite_study_assets or not one_pager_path.exists():
            if not one_pager_path.exists() and not args.overwrite_study_assets:
                created_assets += 1
            _write_text(one_pager_path, render_one_pager(p))

        study_pack_path = paper_dir / "study_pack.md"
        if args.overwrite_study_assets or not study_pack_path.exists():
            if not study_pack_path.exists() and not args.overwrite_study_assets:
                created_assets += 1
            _write_text(study_pack_path, render_study_pack(p))

        # Wrappers: always overwrite (generated files).
        write_wrappers(p, paper_dir)

    print(f"Done. Created/updated assets for {papers_dir}. (New study assets created: {created_assets})")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
