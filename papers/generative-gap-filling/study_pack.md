# Study pack: Generative Gap Filling (generative-gap-filling)

- Full text: `papers/generative-gap-filling/paper.txt`
- Summary (EN): `papers/generative-gap-filling/summary.md`

## Elevator pitch

Contract law assumes that a textual gap marks the point where interpretation ends and judicial supplementation begins. This paper tests that premise by removing negotiated terms from real contracts and asking lay readers, law students, practicing lawyers, and six frontier language models to reconstruct them. Lay readers were correct 55% of the time, lawyers nearly 60%, and the models 88.3%; on 119 additional unseen commercial contracts, the models recovered the masked clause 87% of the time. The results suggest that surrounding contract language often makes the hypothetical bargain statistically legible, while also identifying where model-based gap filling fails and how courts and parties might govern its use.

## Structured outline

### The empirical gap in gap filling

Contract doctrine distinguishes interpretation from construction on the assumption that a gap means the document no longer contains evidence of the parties’ intent. The paper converts that assumption into a falsifiable prediction: readers should not be able to reconstruct a deliberately hidden negotiated term from the remaining agreement.

### Creating a ground truth

The authors mask terms from executed contracts, provide a realistic triggering scenario, and ask participants to choose among four reconstructions. Because the parties’ actual language is known, the method can grade interpretive accuracy rather than merely compare plausible answers.

### Human and model performance

Lay respondents reached 55% accuracy, law students performed marginally better, and experienced lawyers approached 60%. Six frontier language models reached 88.3% overall. Legal experience helped most where domain familiarity supplied useful priors; both humans and models still made systematic errors.

### Robustness and generalization

Direction-flipped contracts and no-contract baselines show that model predictions combine general market expectations with agreement-specific information. A further test on 119 unseen commercial agreements produced 87% model accuracy, close to the main study’s result. Errors clustered in bespoke provisions and clauses that reversed a conventional default.

### Doctrinal and institutional consequences

The results weaken the claim that construction necessarily begins where textual evidence ends. They do not collapse prediction into legal obligation: a court must still decide whether the reconstructed bargain should govern. The recommended procedure emphasizes disclosure, adversarial presentation, sensitivity testing, and narrow judicial reasoning.

### Choice of Model clauses

Parties can specify a model, version, or selection protocol in advance, analogous to choices of law or forum. Such clauses can discipline later interpretation while encouraging transactional lawyers to draft with model legibility and contestability in mind.

### Limits and future contracts

Model performance varies by contract type, prompt, and task; contamination cannot be eliminated completely; and a technically correct top-line answer can mask errors on finer implications. The method also depends on human-authored agreements. AI-drafted contracts may encode regularities without preserving any human mental state for a later model to recover.

## Keywords / concepts

contract law; contract interpretation; gap filling; large language models; legal AI; masked language modeling; empirical legal studies; hypothetical bargain; Choice of Model clauses; commercial contracts

## Suggested questions (for RAG / study)

- What is the paper’s main claim and what problem does it solve?
- What method/data does it use (if any), and what are the main results?
- What assumptions are doing the most work?
- What are the limitations or failure modes the author flags?
- How does this connect to the author’s other papers in this corpus?

_Auto-generated study aid. For canonical content, rely on `paper.txt`/`paper.pdf`._
