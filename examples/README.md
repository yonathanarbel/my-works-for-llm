# Examples

This directory contains practical examples for working with Professor Arbel's scholarship corpus.

## Jupyter Notebook

### explore_corpus.ipynb

An interactive Jupyter notebook that demonstrates:
- Loading and exploring the corpus
- Getting statistics about papers
- Searching and filtering papers
- Reading paper content
- Analyzing distribution by year
- Creating training datasets
- Exporting data

**To run the notebook:**

```bash
# Install Jupyter if not already installed
pip install jupyter

# Navigate to the examples directory
cd examples

# Start Jupyter
jupyter notebook explore_corpus.ipynb
```

## Python Scripts

You can also run the examples as standalone Python scripts. See [USAGE_EXAMPLES.md](../USAGE_EXAMPLES.md) for more examples.

## Quick Examples

### List all papers

```python
from corpus_api import ArbelCorpus

corpus = ArbelCorpus(base_dir='..')
for paper in corpus.list_papers():
    print(f"{paper.paper_id}: {paper.get_title()}")
```

### Search papers

```python
from corpus_api import ArbelCorpus

corpus = ArbelCorpus(base_dir='..')
results = corpus.search_papers("artificial intelligence")
print(f"Found {len(results)} papers")
```

### Export summaries

```python
from corpus_api import ArbelCorpus
import json

corpus = ArbelCorpus(base_dir='..')
data = []

for paper in corpus.iterate_papers():
    summary = paper.get_summary()
    if summary:
        data.append({
            'id': paper.paper_id,
            'text': summary
        })

with open('summaries.json', 'w') as f:
    json.dump(data, f, indent=2)
```

## Contributing Examples

If you have interesting use cases or examples, please contribute them back to the repository!
