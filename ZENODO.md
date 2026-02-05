# Zenodo DOI (recommended)

This repository includes a `.zenodo.json` file so Zenodo can mint a DOI for each tagged release.

## One-time setup

1. Sign in to Zenodo and connect your GitHub account.
2. In Zenodo → GitHub, enable DOI archiving for `yonathanarbel/my-works-for-llm`.

## Create a DOI for a release

1. Create a new GitHub release (tag) on `main` (e.g., `v2026.02.05`).
2. Zenodo will archive that release and mint a DOI.

## After a DOI exists

- Update `CITATION.cff` with the Zenodo DOI (recommended field: `identifiers`).
- Optionally update `dataset.jsonld` to include the DOI in `identifier`.

