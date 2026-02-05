#!/usr/bin/env python3
"""
Publish a snapshot of this corpus to the Hugging Face Hub (dataset repo).

Requires:
  pip install huggingface_hub
and either:
  set HF_TOKEN env var
or:
  pass --token
"""

from __future__ import annotations

import argparse
import os
import shutil
import subprocess
from pathlib import Path


def main() -> int:
    parser = argparse.ArgumentParser(description="Publish the corpus to Hugging Face Hub (dataset).")
    parser.add_argument("--repo-id", required=True, help="HF dataset repo id, e.g. username/my-works-for-llm")
    parser.add_argument("--token", default=os.getenv("HF_TOKEN"), help="HF token (or set HF_TOKEN).")
    parser.add_argument("--private", action="store_true", help="Create/upload as a private dataset.")
    parser.add_argument(
        "--include",
        nargs="+",
        choices=["summary", "summary_zh", "one_pager", "study_pack", "paper"],
        default=["summary", "summary_zh", "one_pager", "study_pack", "paper"],
        help="Which content to export/upload.",
    )
    parser.add_argument("--workdir", default="dist/hf", help="Local staging directory (default: dist/hf).")
    args = parser.parse_args()

    if not args.token:
        raise SystemExit("Missing token. Pass --token or set HF_TOKEN.")

    try:
        from huggingface_hub import HfApi, upload_folder
    except ImportError as e:
        raise SystemExit("Missing dependency. Install with: pip install huggingface_hub") from e

    repo_root = Path(".")
    workdir = repo_root / args.workdir
    if workdir.exists():
        shutil.rmtree(workdir)
    workdir.mkdir(parents=True, exist_ok=True)

    # Build JSONL export
    jsonl_path = workdir / "arbel-corpus.jsonl"
    cmd = [
        "python",
        "export_llm_dataset.py",
        "--include",
        *args.include,
        "--out-jsonl",
        str(jsonl_path.relative_to(repo_root)).replace("\\", "/"),
    ]
    subprocess.run(cmd, check=True)

    # Dataset card (template)
    card_src = repo_root / "huggingface" / "DATASET_CARD.md"
    if card_src.exists():
        (workdir / "README.md").write_text(card_src.read_text(encoding="utf-8"), encoding="utf-8", newline="\n")

    api = HfApi(token=args.token)
    api.create_repo(repo_id=args.repo_id, repo_type="dataset", private=args.private, exist_ok=True)
    upload_folder(
        folder_path=str(workdir),
        repo_id=args.repo_id,
        repo_type="dataset",
        token=args.token,
        commit_message="Update dataset snapshot from GitHub",
    )

    print(f"Uploaded dataset snapshot to https://huggingface.co/datasets/{args.repo_id}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
