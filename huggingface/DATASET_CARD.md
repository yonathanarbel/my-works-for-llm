---
language:
- en
- zh
license: cc0-1.0
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

DOI: 10.5281/zenodo.18781458

Link: https://doi.org/10.5281/zenodo.18781458

## What’s inside

Each record contains:
- `paper_id` (e.g., `ssrn-5377475`)
- `content_type` (`summary`, `one_pager`, `study_pack`, or `paper`)
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

See the repository license and per-paper metadata (papers are dedicated under `CC0-1.0`).

## Citation

See `CITATION.cff` in the GitHub repo, or cite the Zenodo record directly:

```bibtex
@misc{https://doi.org/10.5281/zenodo.18781458,
  doi = {10.5281/ZENODO.18781458},
  url = {https://zenodo.org/doi/10.5281/zenodo.18781458},
  author = {Arbel, Yonathan A.},
  keywords = {contracts, artificial intelligence, legal scholarship, empirical legal studies, consumer protection},
  title = {my-works-for-llm: Machine-readable corpus of Yonathan Arbel's scholarship},
  publisher = {Zenodo},
  year = {2026},
  copyright = {Creative Commons Zero v1.0 Universal}
}
```
