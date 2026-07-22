# Generative Gap Filling — one-page summary

**Paper ID:** `generative-gap-filling`
**Year:** 2026
**Author(s):** Yonathan A. Arbel, David A. Hoffman

## TL;DR

Contract law assumes that a textual gap marks the point where interpretation ends and judicial supplementation begins. This paper tests that premise by removing negotiated terms from real contracts and asking lay readers, law students, practicing lawyers, and six frontier language models to reconstruct them. Lay readers were correct 55% of the time, lawyers nearly 60%, and the models 88.3%; on 119 additional unseen commercial contracts, the models recovered the masked clause 87% of the time. The results suggest that surrounding contract language often makes the hypothetical bargain statistically legible, while also identifying where model-based gap filling fails and how courts and parties might govern its use.

## Design

- Three executed contracts with negotiated terms masked and converted into four-choice reconstruction tasks.
- Human comparison groups: 465 attentive lay respondents, 77 law students, and 48 experienced lawyers.
- Model comparison group: six frontier systems, twenty runs per model, browsing disabled, answer order varied.
- Robustness checks: direction-flipped agreements, no-contract baselines, and 119 additional mostly 2025-2026 SEC EDGAR contracts.

## Headline Findings

- Lay readers: 55% correct; lawyers: nearly 60%; chance: 25%.
- Frontier models: 88.3% correct in the main study and 87% on 119 unseen contracts.
- Flipping the contract reduced model accuracy to 60%; withholding it produced 68% accuracy.
- Models were strongest on standardized instruments and weakest on bespoke or default-reversing terms.
- A stricter follow-up exposed a shared model failure on a technical pricing provision, underscoring that headline accuracy does not eliminate important errors.

## Legal Takeaway

Some apparent contractual gaps may be recoverable from the agreement itself rather than supplied through unconstrained judicial policy. The authors propose treating model predictions as contestable evidence, with disclosure, adversarial testing, and party-adopted "Choice of Model" clauses—not as autonomous or binding adjudication.

## Keywords

contract law; contract interpretation; gap filling; large language models; legal AI; masked language modeling; empirical legal studies; hypothetical bargain; Choice of Model clauses; commercial contracts

## Files

- Full text: `papers/generative-gap-filling/paper.txt`
- PDF: `papers/generative-gap-filling/paper.pdf`
- Summary (EN): `papers/generative-gap-filling/summary.md`

_Auto-generated study aid. For canonical content, rely on `paper.txt`/`paper.pdf`._
