---
language:
- en
- zh
license: cc-by-nc-4.0
tags:
- legal
- law
- contracts
- empirical-legal-studies
- artificial-intelligence
- llm
pretty_name: my-works-for-llm
---

# my-works-for-llm

Machine-readable corpus of Professor Yonathan Arbel's scholarship for LLM research.

## What’s inside

Each record contains:
- `paper_id` (e.g., `ssrn-5377475`)
- `content_type` (`summary` or `paper`)
- `language` (`en`, `zh`, or `und`)
- `source_path` (path in the GitHub repo)
- `metadata` (JSON metadata)
- `text` (the actual content)

## How to build the dataset files

From the GitHub repo:

```bash
python export_llm_dataset.py --out-jsonl dist/arbel-corpus.jsonl
```

## License

See the repository license and per-paper metadata (many papers specify `CC-BY-NC-4.0` for PDFs).

## Citation

See `CITATION.cff` in the GitHub repo.

