# Usage Examples

This document provides practical examples for working with Professor Arbel's scholarship corpus.

## Command-Line Interface Examples

### Basic Operations

```bash
# List all papers
python corpus_cli.py list

# List papers with detailed file information
python corpus_cli.py list --verbose

# Show details about a specific paper
python corpus_cli.py show ssrn-3519630

# Show paper with full summary content
python corpus_cli.py show ssrn-3519630 --content
```

### Searching

```bash
# Search all fields for a term
python corpus_cli.py search "artificial intelligence"

# Search only in titles
python corpus_cli.py search "disclosure" --field title

# Search in abstracts
python corpus_cli.py search "contract law" --field abstract

# Search in summaries (full text search)
python corpus_cli.py search "consent" --field summary
```

### Statistics and Validation

```bash
# Display corpus statistics
python corpus_cli.py stats

# Validate corpus integrity
python corpus_cli.py validate

# List papers in different formats
python corpus_cli.py list --format table    # Default, nice display
python corpus_cli.py list --format simple   # Simple list of IDs
python corpus_cli.py list --format json     # JSON output for scripting
```

## Python API Examples

### Basic Usage

```python
from corpus_api import ArbelCorpus

# Initialize the corpus
corpus = ArbelCorpus()

# Get basic information
print(f"Total papers: {len(corpus)}")
papers = corpus.list_papers()

# Access a specific paper
paper = corpus.get_paper('ssrn-3519630')
print(f"Title: {paper.get_title()}")
print(f"Authors: {', '.join(paper.get_authors())}")
```

### Reading Paper Content

```python
# Get English summary
summary = paper.get_summary()
if summary:
    print(summary)

# Get Chinese summary
summary_zh = paper.get_summary('zh')
if summary_zh:
    print(summary_zh)

# Get full paper text
full_text = paper.get_full_text()
if full_text:
    print(f"Full text length: {len(full_text)} characters")

# Get abstract from metadata
abstract = paper.get_abstract()
```

### Working with Metadata

```python
# Access all metadata
metadata = paper.metadata
print(metadata)

# Get specific metadata fields
title = paper.get_title()
authors = paper.get_authors()
year = paper.get_year()
abstract = paper.get_abstract()

# Check what files are available
files = paper.get_available_files()
print(f"Available files: {files}")

# Get PDF path
pdf_path = paper.get_pdf_path()
if pdf_path:
    print(f"PDF available at: {pdf_path}")
```

### Searching Papers

```python
# Search all fields
results = corpus.search_papers("artificial intelligence")
print(f"Found {len(results)} papers")

# Search specific field
results = corpus.search_papers("contract", field="title")

# Iterate through results
for paper in results:
    print(f"{paper.paper_id}: {paper.get_title()}")
```

### Iterating Through Papers

```python
# Method 1: Get all papers at once
papers = corpus.list_papers()
for paper in papers:
    print(paper.get_title())

# Method 2: Use generator (memory efficient)
for paper in corpus.iterate_papers():
    if paper.has_summary():
        summary = paper.get_summary()
        # Process summary...
```

### Getting Corpus Statistics

```python
# Get comprehensive stats
stats = corpus.get_stats()
print(f"Total papers: {stats['total_papers']}")
print(f"With English summaries: {stats['with_english_summary']}")
print(f"With Chinese summaries: {stats['with_chinese_summary']}")
print(f"With PDF: {stats['with_pdf']}")
```

## LLM Training Pipeline Examples

### Example 1: Extract All Summaries for Training

```python
from corpus_api import ArbelCorpus

corpus = ArbelCorpus()
training_data = []

for paper in corpus.iterate_papers():
    summary = paper.get_summary()
    if summary:
        training_data.append({
            'id': paper.paper_id,
            'title': paper.get_title(),
            'text': summary,
            'metadata': paper.metadata
        })

print(f"Extracted {len(training_data)} summaries for training")
```

### Example 2: Create Multi-Granularity Dataset

```python
from corpus_api import ArbelCorpus

corpus = ArbelCorpus()
dataset = []

for paper in corpus.iterate_papers():
    entry = {
        'paper_id': paper.paper_id,
        'title': paper.get_title(),
        'authors': paper.get_authors(),
        'year': paper.get_year(),
    }
    
    # Add summary if available
    summary = paper.get_summary()
    if summary:
        entry['summary'] = summary
    
    # Add full text if available
    full_text = paper.get_full_text()
    if full_text:
        entry['full_text'] = full_text
    
    dataset.append(entry)

# Save to JSON
import json
with open('training_dataset.json', 'w') as f:
    json.dump(dataset, f, indent=2)
```

### Example 3: Bilingual Training Data

```python
from corpus_api import ArbelCorpus

corpus = ArbelCorpus()
bilingual_pairs = []

for paper in corpus.iterate_papers():
    en_summary = paper.get_summary('en')
    zh_summary = paper.get_summary('zh')
    
    if en_summary and zh_summary:
        bilingual_pairs.append({
            'paper_id': paper.paper_id,
            'english': en_summary,
            'chinese': zh_summary
        })

print(f"Found {len(bilingual_pairs)} bilingual paper pairs")
```

### Example 4: Filter by Topic and Export

```python
from corpus_api import ArbelCorpus

corpus = ArbelCorpus()

# Search for papers on a specific topic
ai_papers = corpus.search_papers("artificial intelligence")

# Export to text files for fine-tuning
for paper in ai_papers:
    summary = paper.get_summary()
    if summary:
        filename = f"ai_corpus/{paper.paper_id}.txt"
        with open(filename, 'w') as f:
            f.write(f"Title: {paper.get_title()}\n\n")
            f.write(summary)

print(f"Exported {len(ai_papers)} papers on AI")
```

### Example 5: Tokenization Pipeline

```python
from corpus_api import ArbelCorpus
# Assuming you have a tokenizer (e.g., from transformers)
# from transformers import AutoTokenizer

corpus = ArbelCorpus()
# tokenizer = AutoTokenizer.from_pretrained("your-model")

all_tokens = []
for paper in corpus.iterate_papers():
    text = paper.get_summary() or paper.get_full_text()
    if text:
        # tokens = tokenizer.encode(text)
        # all_tokens.extend(tokens)
        print(f"Processing {paper.paper_id}...")

print(f"Total papers processed: {len(corpus)}")
```

## Automation Examples

### Example 1: Daily Validation Script

```bash
#!/bin/bash
# validate_corpus.sh

echo "Validating corpus integrity..."
python corpus_cli.py validate

if [ $? -eq 0 ]; then
    echo "✓ Validation passed"
    python corpus_cli.py stats
else
    echo "✗ Validation failed"
    exit 1
fi
```

### Example 2: Export Statistics to JSON

```python
from corpus_api import ArbelCorpus
import json

corpus = ArbelCorpus()
stats = corpus.get_stats()

# Add paper list
stats['papers'] = [
    {
        'id': p.paper_id,
        'title': p.get_title(),
        'has_summary': p.has_summary(),
        'has_full_text': p.get_full_text() is not None
    }
    for p in corpus.list_papers()
]

with open('corpus_stats.json', 'w') as f:
    json.dump(stats, f, indent=2)

print("Statistics exported to corpus_stats.json")
```

### Example 3: Generate Paper Index HTML

```python
from corpus_api import ArbelCorpus

corpus = ArbelCorpus()

html = ['<html><body><h1>Professor Arbel\'s Papers</h1><ul>']

for paper in corpus.list_papers():
    title = paper.get_title()
    paper_id = paper.paper_id
    html.append(f'<li><a href="papers/{paper_id}/">{title}</a></li>')

html.append('</ul></body></html>')

with open('index.html', 'w') as f:
    f.write('\n'.join(html))

print("Generated index.html")
```

## Integration with Popular Tools

### Using with Pandas

```python
from corpus_api import ArbelCorpus
import pandas as pd

corpus = ArbelCorpus()

# Create a DataFrame of all papers
data = []
for paper in corpus.list_papers():
    data.append({
        'paper_id': paper.paper_id,
        'title': paper.get_title(),
        'year': paper.get_year(),
        'has_summary': paper.has_summary(),
        'has_pdf': paper.get_pdf_path() is not None,
    })

df = pd.DataFrame(data)
print(df.head())

# Analyze by year
print(df.groupby('year').size())
```

### Using with Hugging Face Datasets

```python
from corpus_api import ArbelCorpus
from datasets import Dataset

corpus = ArbelCorpus()

# Prepare data
data = {
    'paper_id': [],
    'title': [],
    'text': [],
}

for paper in corpus.iterate_papers():
    summary = paper.get_summary()
    if summary:
        data['paper_id'].append(paper.paper_id)
        data['title'].append(paper.get_title())
        data['text'].append(summary)

# Create HF dataset
dataset = Dataset.from_dict(data)
print(dataset)

# Push to hub (if desired)
# dataset.push_to_hub("username/arbel-papers")
```
