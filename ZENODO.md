# Zenodo DOI (recommended)

This repository includes a `.zenodo.json` file so Zenodo can mint a DOI for each tagged release.

This repo treats the **Zenodo concept DOI** (the DOI that points to “all versions”) as the canonical dataset DOI to embed in metadata.

## One-time setup

1. Sign in to Zenodo and connect your GitHub account.
2. In Zenodo → GitHub, enable DOI archiving for `yonathanarbel/my-works-for-llm`.

## Create a DOI for a release

1. Create a new GitHub release (tag) on `main` (e.g., `v2026.02.05`).
2. Zenodo will archive that release and mint a DOI.

## After a DOI exists

- Run `python sync_zenodo_doi.py` (or trigger the **Sync Zenodo DOI** GitHub Action).
  - This updates `CITATION.cff`, `dataset.jsonld`, `.zenodo.json`, the Hugging Face dataset card, and regenerates JSON-LD + sitemaps + `docs/`.
