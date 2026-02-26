# Quick Start Guide

Get up and running with Professor Arbel's scholarship corpus in minutes.

## Installation

```bash
# Clone the repository
git clone https://github.com/yonathanarbel/my-works-for-llm.git
cd my-works-for-llm

# Install dependencies
pip install -r requirements.txt

# Optional: Install as a package for CLI shortcuts
pip install -e .
```

## 5-Minute Tour

### 1. List All Papers

```bash
python corpus_cli.py list
```

Or with the installed CLI:
```bash
arbel-corpus list
```

### 2. View Corpus Statistics

```bash
python corpus_cli.py stats
```

You'll see:
- Total number of papers (25+)
- Papers with summaries in English and Chinese
- File type distribution
- Total corpus size

### 3. Search for Papers

```bash
python corpus_cli.py search "artificial intelligence"
```

### 4. View a Specific Paper

```bash
python corpus_cli.py show ssrn-3519630 --content
```

### 5. Validate Corpus

```bash
python corpus_cli.py validate
```

## Using the Python API

Create a Python script or use in a Jupyter notebook:

```python
from corpus_api import ArbelCorpus

# Load the corpus
corpus = ArbelCorpus()

# Print basic stats
print(f"Total papers: {len(corpus)}")

# Get a specific paper
paper = corpus.get_paper('ssrn-3519630')
print(f"Title: {paper.get_title()}")

# Read the summary
summary = paper.get_summary()
print(summary[:200] + "...")  # First 200 characters

# Search papers
results = corpus.search_papers("contract")
print(f"Found {len(results)} papers about contracts")

# Iterate through all papers
for paper in corpus.iterate_papers():
    print(f"- {paper.paper_id}: {paper.get_title()}")
```

## Common Tasks

### Export All Summaries to Text Files

```python
from corpus_api import ArbelCorpus
from pathlib import Path

corpus = ArbelCorpus()
output_dir = Path('exported_summaries')
output_dir.mkdir(exist_ok=True)

for paper in corpus.iterate_papers():
    summary = paper.get_summary()
    if summary:
        output_file = output_dir / f"{paper.paper_id}.txt"
        output_file.write_text(summary)

print(f"Exported summaries to {output_dir}")
```

### Create a Training Dataset (JSON)

```python
from corpus_api import ArbelCorpus
import json

corpus = ArbelCorpus()
dataset = []

for paper in corpus.iterate_papers():
    summary = paper.get_summary()
    if summary:
        dataset.append({
            'id': paper.paper_id,
            'title': paper.get_title(),
            'text': summary,
            'year': paper.get_year()
        })

with open('training_data.json', 'w') as f:
    json.dump(dataset, f, indent=2)

print(f"Created training dataset with {len(dataset)} papers")
```

### Search and Filter Papers

```bash
# Search in all fields
arbel-corpus search "contract law"

# Search only titles
arbel-corpus search "disclosure" --field title

# Search in abstracts
arbel-corpus search "artificial intelligence" --field abstract
```

### Get Paper Information in JSON

```bash
# Export all papers metadata to JSON
arbel-corpus list --format json > papers.json

# Use in scripts
python -c "
from corpus_api import ArbelCorpus
import json

corpus = ArbelCorpus()
papers_data = [
    {'id': p.paper_id, 'title': p.get_title()}
    for p in corpus.list_papers()
]
print(json.dumps(papers_data, indent=2))
" > papers_simple.json
```

## Next Steps

- **Read the full documentation**: See [README.md](README.md) for detailed information
- **Explore examples**: Check [USAGE_EXAMPLES.md](USAGE_EXAMPLES.md) for more use cases
- **Integrate with your workflow**: Use the Python API in your data pipelines
- **Contribute**: Add new papers or improve the tools

## Troubleshooting

### "rich library not found"

```bash
pip install rich
```

### "Papers directory not found"

Make sure you're running commands from the repository root directory:

```bash
cd /path/to/my-works-for-llm
python corpus_cli.py list
```

### Permission errors on Linux/Mac

Make the CLI script executable:

```bash
chmod +x corpus_cli.py
./corpus_cli.py list
```

## Getting Help

```bash
# General help
arbel-corpus --help

# Command-specific help
arbel-corpus list --help
arbel-corpus search --help
arbel-corpus show --help
```

## Key Features at a Glance

- ✅ **25+ papers** from Professor Yonathan Arbel
- ✅ **Bilingual support** (English and Chinese summaries)
- ✅ **Multiple formats** (PDF, TXT, Markdown)
- ✅ **CLI tool** for quick exploration
- ✅ **Python API** for programmatic access
- ✅ **Search functionality** across papers
- ✅ **Validation tools** for data integrity
- ✅ **LLM-ready** formats and scripts
- ✅ **Rich metadata** in Schema.org format
- ✅ **Open source** under CC0 1.0 license

Happy exploring! 📚
