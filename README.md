# Yonathan Arbel's Scholarship for LLM Research

This repository provides a machine-readable corpus of Professor Yonathan Arbel's scholarly articles, curated for use in Large Language Model (LLM) research.

## Project Goal

The primary goal of this project is to offer a structured and accessible collection of legal scholarship that can be readily used by researchers and developers working with LLMs. By providing full texts, summaries, and metadata in machine-readable formats, this corpus aims to facilitate various research avenues, including but not limited to text analysis, summarization, translation, and information retrieval within the legal domain.

## Repository Structure

The repository is organized as follows:

*   **`LICENSE`**: Contains the full text of the Creative Commons Attribution-NonCommercial 4.0 International License.
*   **`README.md`**: This file, providing an overview of the project.
*   **`dataset.jsonld`**: A JSON-LD file providing structured metadata about the dataset, adhering to schema.org standards.
*   **`llms.txt`**: An index file that lists direct links to the plain text summaries of each paper, useful for quick navigation or programmatic access.
*   **`papers/`**: This directory houses the individual scholarly articles. Each paper typically has its own subdirectory, named in a consistent format (e.g., `YYYY-Title-Of-Paper`). Inside each paper's subdirectory, you will find:
    *   `paper.pdf`: The original published version of the paper in PDF format.
    *   `paper.txt`: A plain text version of the paper, extracted for easier machine processing.
    *   `summary.md`: A concise summary of the paper in English, written in Markdown.
    *   `summary.zh.md`: A Chinese translation of the summary, also in Markdown.
    *   `metadata.json`: A JSON file containing metadata specific to the paper, such as publication date, journal, and keywords.

## What "Machine-Readable" Means

This corpus is designed to be "machine-readable" in several ways:

*   **Plain Text Versions**: Full texts of papers are provided in `.txt` format, removing complexities of PDF parsing for NLP tasks.
*   **Structured Summaries**: Summaries are available in Markdown (`.md`), a simple and structured format.
*   **Metadata**: Both dataset-level (`dataset.jsonld`) and paper-specific (`metadata.json`) metadata are provided in JSON format, allowing for easy parsing and integration into automated workflows.

## Usage for LLM Research

This corpus can be utilized in various LLM research projects, for example:

*   **Training and Fine-tuning Models**: The plain text versions of the papers and their summaries can be used to train or fine-tune models for tasks like:
    *   Legal text summarization.
    *   Abstract generation.
    *   Legal document Q&A.
*   **Cross-lingual Studies**: The English and Chinese summaries enable research into cross-lingual summarization, translation quality assessment, and comparative legal concept analysis.
*   **Legal Text Analysis**: Researchers can perform various NLP tasks such as topic modeling, keyword extraction, and sentiment analysis on the legal texts.
*   **Information Retrieval**: The structured metadata and text can be used to build and test information retrieval systems focused on legal documents.
*   **Citation Network Analysis**: While not directly included, the metadata can be a starting point for constructing citation networks (though this would require external data sources).

**Note**: This repository contains data (text, metadata) and does not include executable code or scripts. The intention is for this data to be used *with* external LLM tools, libraries, and research scripts.

## License

The contents of this repository, created by Professor Yonathan Arbel, are licensed under the **Creative Commons Attribution-NonCommercial 4.0 International License (CC BY-NC 4.0)**.

You can find a copy of the license in the [`LICENSE`](./LICENSE) file in this repository or read it online at [https://creativecommons.org/licenses/by-nc/4.0/](https://creativecommons.org/licenses/by-nc/4.0/).

Under this license, you are free to:
*   **Share** — copy and redistribute the material in any medium or format
*   **Adapt** — remix, transform, and build upon the material

Under the following terms:
*   **Attribution** — You must give appropriate credit, provide a link to the license, and indicate if changes were made. You may do so in any reasonable manner, but not in any way that suggests the licensor endorses you or your use.
*   **NonCommercial** — You may not use the material for commercial purposes.

## Navigation Aid

For a quick overview and direct links to the summaries of all papers included in this corpus, please refer to the `llms.txt` file. This can be particularly useful for programmatic access to the summaries.

## Citation

[To be added: A preferred citation format for this dataset if used in research.]
If you use this dataset in your research, please consider citing it. A recommended citation format will be provided here in the future.
For now, please acknowledge Professor Yonathan Arbel and this repository.
