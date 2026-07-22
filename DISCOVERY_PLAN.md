# Discovery and propagation plan

The goal is to maximize two distinct outcomes:

1. **Retrieval now:** make the scholarship easy for search engines and AI search
   products to find, understand, quote, and cite.
2. **Training later:** place stable, licensed copies in the public corpora and
   academic repositories that model builders are most likely to ingest.

No technical change can guarantee model training or citation. The practical
strategy is to publish one authoritative copy, expose high-signal machine
formats, strengthen the link graph, and measure actual crawl and citation data.

## P0: canonical discovery (days)

- [x] Use `https://works.battleoftheforms.com/` as the canonical corpus.
- [x] Serve paper pages, clean text, PDFs, JSON-LD, JSONL, claims, Q&A,
  `llms.txt`, feeds, and sitemaps from the canonical host.
- [x] Point the `works.yonathanarbel.com` and GitHub Pages mirrors at the
  canonical host through redirects or canonical metadata.
- [x] Allow search and major AI crawlers in `robots.txt`.
- [x] Disable Cloudflare AI-bot blocking and crawler labyrinth protection.
- [x] Publish an IndexNow ownership key and submit the canonical sitemap URLs.
- [ ] Enable Cloudflare Crawler Hints in the Cloudflare dashboard.
- [ ] Verify the canonical host in Google Search Console and Bing Webmaster
  Tools, then submit `sitemap.xml` and `sitemap-papers.xml`.

## P1: metadata and dataset distribution (days to weeks)

- [x] Publish a DOI-bearing Schema.org `Dataset` record with stable
  distributions and `sameAs` links.
- [x] Give every paper a canonical `ScholarlyArticle` URL and descriptive page
  title.
- [x] Make the Hugging Face snapshot directly loadable with a documented
  `load_dataset` example and explicit dataset-card metadata.
- [x] Add automated checks that prevent canonical and distribution metadata
  from drifting.
- [x] Confirm the Zenodo release archive contains resolved PDF binaries rather
  than Git LFS pointer files.
- [x] Request archival of the current main branch in Software Heritage.
- [ ] Request a Software Heritage refresh after this change is merged.

## P2: authoritative links and academic propagation (weeks)

- [ ] Add the canonical corpus URL to the Alabama Law faculty page, ORCID,
  Google Scholar homepage field, SSRN author profile, and GitHub profile.
- [ ] Set the corpus URL as the homepage for scholarship repositories and add
  consistent GitHub topics.
- [ ] Deposit eligible AI-related manuscripts in arXiv and ensure every paper
  has an open full-text record in Alabama Law Scholarly Commons.
- [ ] Reconcile DOI, SSRN, OpenAlex, Semantic Scholar, and institutional-repository
  identifiers in each paper's `sameAs` metadata.
- [ ] Add or improve sourced Wikidata records for the author and major papers;
  avoid autobiographical Wikipedia editing.

## Monthly measurement loop

Record the following in a dated issue or spreadsheet:

- Google Search Console and Bing indexed-page counts.
- Cloudflare AI Crawl Control requests by crawler, hostname, status, and path.
- Common Crawl index hits for the canonical domain.
- Hugging Face downloads and Zenodo views/downloads.
- Results of a fixed five-question probe set in ChatGPT, Claude, Gemini, and
  Perplexity with web search enabled: whether the answer cites the canonical
  site, a mirror, SSRN, or another repository.

Use the same questions each month. Change content only when the data identifies
a concrete retrieval gap—for example, a missing query formulation, claim,
citation, or authoritative inbound link.

## Operational rule

Generate and deploy from versioned sources, keep `works.battleoftheforms.com`
canonical everywhere, and treat mirrors as discovery paths rather than competing
canonical copies.
