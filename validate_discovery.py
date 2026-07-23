#!/usr/bin/env python3
"""Validate canonical and distribution metadata used by crawlers and dataset hubs."""

from __future__ import annotations

import json
import re
from pathlib import Path


ROOT = Path(__file__).resolve().parent
CANONICAL_BASE = "https://works.battleoftheforms.com/"


def require(condition: bool, message: str) -> None:
    if not condition:
        raise SystemExit(f"Discovery validation failed: {message}")


def main() -> int:
    dataset = json.loads((ROOT / "dataset.jsonld").read_text(encoding="utf-8"))
    require(dataset.get("url") == CANONICAL_BASE, "dataset.jsonld must use the canonical corpus URL")
    require(dataset.get("identifier", {}).get("value") == "10.5281/zenodo.18781457", "dataset DOI is missing")
    require(len(dataset.get("distribution", [])) >= 3, "dataset distributions are incomplete")

    card = (ROOT / "huggingface" / "DATASET_CARD.md").read_text(encoding="utf-8")
    for marker in ("size_categories:", "task_categories:", "configs:", "load_dataset("):
        require(marker in card, f"Hugging Face card is missing {marker}")

    robots = (ROOT / "docs" / "robots.txt").read_text(encoding="utf-8")
    require(f"Sitemap: {CANONICAL_BASE}sitemap.xml" in robots, "robots.txt sitemap is not canonical")

    html_files = [ROOT / "docs" / "index.html", *sorted((ROOT / "docs" / "papers").glob("*/index.html"))]
    require(len(html_files) > 1, "generated paper pages are missing")
    for html_file in html_files:
        page = html_file.read_text(encoding="utf-8", errors="replace")
        match = re.search(r'<link rel="canonical" href="([^"]+)"', page)
        require(bool(match), f"canonical link missing from {html_file.relative_to(ROOT)}")
        require(match.group(1).startswith(CANONICAL_BASE), f"non-canonical URL in {html_file.relative_to(ROOT)}")

    for schema_file in sorted((ROOT / "papers").glob("*/scholarlyarticle.jsonld")):
        schema = json.loads(schema_file.read_text(encoding="utf-8"))
        require(str(schema.get("url", "")).startswith(CANONICAL_BASE), f"non-canonical schema URL in {schema_file}")

    print(f"Discovery metadata valid: {len(html_files) - 1} paper pages checked")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
