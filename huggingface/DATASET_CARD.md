---
language:
- en
- zh
license: cc0-1.0
size_categories:
- n<1K
task_categories:
- question-answering
- text-generation
configs:
- config_name: default
  data_files: arbel-corpus.jsonl
  default: true
tags:
- legal
- law
- contracts
- empirical-legal-studies
- artificial-intelligence
- llm
- text
- datasets
pretty_name: my-works-for-llm
---

# my-works-for-llm

Machine-readable corpus of Professor Yonathan Arbel's scholarship for retrieval,
question answering, evaluation, and LLM research.

Canonical site: https://works.battleoftheforms.com/

DOI: https://doi.org/10.5281/zenodo.18781458

Source: https://github.com/yonathanarbel/my-works-for-llm

## Load the dataset

```python
from datasets import load_dataset

dataset = load_dataset("yonathanarbel/my-works-for-llm", split="train")
print(dataset.column_names)
print(dataset[0]["text"][:500])
```

For streaming:

```python
dataset = load_dataset(
    "yonathanarbel/my-works-for-llm",
    split="train",
    streaming=True,
)
```

## What’s inside

Each record contains:
- `paper_id` (e.g., `ssrn-5377475`)
- `content_type` (`summary`, `one_pager`, `study_pack`, or `paper`)
- `language` (`en`, `zh`, or `und`)
- `source_path` (path in the GitHub repo)
- `metadata` (JSON metadata)
- `text` (the actual content)

The single `train` split contains multiple representations of a paper. Use
`paper_id`, `content_type`, and `language` to group or filter records. The
`summary`, `one_pager`, and `study_pack` records are derived aids; use the
`paper` record when full source text is required.

## Intended uses and limitations

- Retrieval and citation experiments over legal scholarship.
- Question-answering and summarization prototypes.
- Corpus indexing and data-pipeline demonstrations.

The corpus is small and author-specific, so it is not representative of legal
scholarship as a whole. Verify quotations and legal propositions against the
paper text and cited publication. Publisher-hosted versions may have terms that
differ from the corpus-level CC0 dedication.

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
