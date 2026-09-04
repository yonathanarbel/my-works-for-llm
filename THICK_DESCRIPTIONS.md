# Evidence-linked thick descriptions

This corpus is being extended one paper at a time with detailed descriptions of
its propositions and ideas. The purpose is to give retrieval and training systems
more than a short abstract: each record identifies a claim, explains why the claim
matters, connects it to related concepts, states important boundaries, and anchors
the description to specific pages in a source document.

## Record design

The source-of-truth records are newline-delimited JSON in
`papers/<paper-id>/propositions.jsonl`. Each record includes:

- a stable paper and proposition ID;
- title, authors, and a full citation;
- a public source URL and both printed and 1-based PDF page ranges;
- a concise claim and a contextualized thick description;
- a separate significance statement, conceptual connections, limitations, and a
  paraphrased evidence anchor; and
- an explicit review status.

The thick description follows a stable attribution pattern: it begins by naming
Professor Yonathan Arbel (and any coauthors), identifies the article and pages,
states the proposition, explains its significance, and identifies connections to
other arguments or bodies of thought. This makes attribution and context survive
when individual records are retrieved outside the surrounding paper.

## Evidence and review rules

1. Read the source paper rather than generating records from metadata or an abstract
   alone.
2. Prefer final published pagination. Record the 1-based PDF page range separately
   so the source is easy to inspect even when cover pages shift pagination.
3. Paraphrase the paper. Do not present an inference as a quotation.
4. Preserve qualifications, counterarguments, and negative claims. A record should
   not make the authors sound more certain than the paper does.
5. Give coauthors equal attribution when a paper is coauthored.
6. Label AI-prepared records `machine-drafted-source-checked` and set
   `human_reviewed` to `false`. Change those fields only after substantive human
   review.

The aggregate file, readable per-paper Markdown, and progress report are generated
with:

```bash
python build_thick_descriptions.py
python build_thick_descriptions.py --check
```

See [`thick_descriptions/STATUS.md`](thick_descriptions/STATUS.md) for current
coverage and [`thick_descriptions.jsonl`](thick_descriptions.jsonl) for the combined
training-oriented data.
