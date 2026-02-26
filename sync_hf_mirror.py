#!/usr/bin/env python3
"""
Mirror this GitHub repository's tracked files to a Hugging Face dataset repo.

One-command usage:
  python sync_hf_mirror.py

Requirements:
  pip install huggingface_hub
  HF_TOKEN or HUGGINGFACEHUB_API_TOKEN must be set.
"""

from __future__ import annotations

import argparse
import os
import shutil
import subprocess
from pathlib import Path


def _run_git(repo_root: Path, *args: str) -> str:
    result = subprocess.run(
        ["git", "-C", str(repo_root), *args],
        check=True,
        capture_output=True,
        text=True,
        encoding="utf-8",
        errors="replace",
    )
    return result.stdout.strip()


def _copy_tracked_files(repo_root: Path, mirror_dir: Path) -> int:
    tracked = _run_git(repo_root, "ls-files").splitlines()
    count = 0
    for rel in tracked:
        src = repo_root / rel
        if not src.is_file():
            continue
        dst = mirror_dir / rel
        dst.parent.mkdir(parents=True, exist_ok=True)
        shutil.copy2(src, dst)
        count += 1
    return count


def _strip_markdown_front_matter(text: str) -> str:
    if not text.startswith("---\n"):
        return text
    end = text.find("\n---\n", 4)
    if end == -1:
        return text
    return text[end + len("\n---\n") :]


def _write_hf_readme(mirror_dir: Path, repo_id: str, source_repo_url: str, head_sha: str) -> None:
    readme_path = mirror_dir / "README.md"
    existing = ""
    if readme_path.exists():
        existing = readme_path.read_text(encoding="utf-8", errors="replace")
    existing = _strip_markdown_front_matter(existing).lstrip()

    pretty_name = repo_id.split("/", 1)[-1]
    front_matter = "\n".join(
        [
            "---",
            "language:",
            "- en",
            "- zh",
            "license: cc0-1.0",
            "tags:",
            "- legal",
            "- law",
            "- llm",
            "- dataset",
            "- github-mirror",
            f"pretty_name: {pretty_name}",
            "---",
            "",
        ]
    )

    lead = "\n".join(
        [
            f"# {pretty_name}",
            "",
            f"Automated mirror of `{source_repo_url}`.",
            f"Source commit: `{head_sha}`.",
            "",
        ]
    )

    out = front_matter + lead
    if existing:
        out += "## Source README\n\n" + existing.rstrip() + "\n"
    readme_path.write_text(out, encoding="utf-8", newline="\n")


def main(argv: list[str] | None = None) -> int:
    parser = argparse.ArgumentParser(
        description="Mirror tracked Git files from this repo to a Hugging Face dataset repo."
    )
    parser.add_argument(
        "--repo-id",
        default=None,
        help=(
            "HF dataset repo id (e.g., username/my-works-for-llm-github-mirror). "
            "Default: inferred from your HF username and local folder name."
        ),
    )
    parser.add_argument(
        "--token",
        default=os.getenv("HF_TOKEN") or os.getenv("HUGGINGFACEHUB_API_TOKEN"),
        help="HF token (or set HF_TOKEN / HUGGINGFACEHUB_API_TOKEN).",
    )
    parser.add_argument(
        "--repo-root",
        default=str(Path(__file__).resolve().parent),
        help="Local repository root (default: this script's directory).",
    )
    parser.add_argument(
        "--workdir",
        default="dist/hf-github-mirror",
        help="Temporary local mirror directory inside repo-root (default: dist/hf-github-mirror).",
    )
    parser.add_argument(
        "--private",
        action="store_true",
        help="Create/upload mirror dataset as private.",
    )
    parser.add_argument(
        "--commit-message",
        default=None,
        help="Custom commit message for HF upload.",
    )
    args = parser.parse_args(argv)

    if not args.token:
        raise SystemExit("Missing token. Pass --token or set HF_TOKEN / HUGGINGFACEHUB_API_TOKEN.")

    try:
        from huggingface_hub import HfApi, upload_folder
    except ImportError as exc:
        raise SystemExit("Missing dependency. Install with: pip install huggingface_hub") from exc

    repo_root = Path(args.repo_root).resolve()
    mirror_dir = (repo_root / args.workdir).resolve()

    if mirror_dir.exists():
        shutil.rmtree(mirror_dir)
    mirror_dir.mkdir(parents=True, exist_ok=True)

    file_count = _copy_tracked_files(repo_root, mirror_dir)
    head_sha = _run_git(repo_root, "rev-parse", "--short", "HEAD")
    source_repo_url = _run_git(repo_root, "config", "--get", "remote.origin.url") or "unknown-source-repo"

    api = HfApi(token=args.token)
    repo_id = args.repo_id
    if not repo_id:
        who = api.whoami(token=args.token)
        username = who["name"] if isinstance(who, dict) and "name" in who else str(who)
        repo_id = f"{username}/{repo_root.name}-github-mirror"

    _write_hf_readme(mirror_dir, repo_id, source_repo_url, head_sha)

    api.create_repo(repo_id=repo_id, repo_type="dataset", private=args.private, exist_ok=True)
    upload_folder(
        folder_path=str(mirror_dir),
        repo_id=repo_id,
        repo_type="dataset",
        token=args.token,
        commit_message=args.commit_message or f"Mirror GitHub tracked files ({head_sha})",
    )

    print(f"Mirrored {file_count} files from commit {head_sha}")
    print(f"https://huggingface.co/datasets/{repo_id}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
