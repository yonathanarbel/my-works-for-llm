"""Readable corpus entry for Thinking the Unthinkable: AI in the Service of Justice."""

ARTICLE_TEXT = r'''# Thinking the Unthinkable: AI in the Service of Justice

Professor Yonathan Arbel argues that large language models can provide useful, context-sensitive evidence about ordinary meaning and thereby supplement dictionaries, precedent, interpretive canons, and judicial intuition. Judge Kevin Newsom's concurrence in the Eleventh Circuit's Snell trampoline-insurance dispute illustrates the approach's movement into judicial practice.

## The interpretive problem

Conventional sources promote stability but may be static, context-poor, or detached from current and vernacular usage. Interpretive error can affect insurance coverage, contractual obligation, statutory meaning, litigation cost, and consistency.

## Generative interpretation

Because LLMs learn statistical patterns from unusually broad text corpora, they may offer evidence about how language is used across domains and communities. The method draws on Arbel's joint work with David Hoffman and treats model output as linguistic evidence rather than legal authority.

## The Snell concurrence

Judge Newsom asked several models whether installing an in-ground trampoline counts as landscaping. Their answers were affirmative and useful to his ordinary-meaning inquiry, but other factors determined the legal result. The episode is an early judicial engagement with generative interpretation, not binding acceptance of it.

## Applications and safeguards

Possible applications include insurance terms, prenuptial agreements, and commercial jargon. Arbel recommends querying multiple models, recording prompts and outputs for replication, treating results as advisory rather than determinative, and maintaining human review for bias, hallucination, and error.

## Institutional impact and limits

Lawyers may need new prompting and evaluation skills; judges may gain broader language evidence; and clients may benefit from more predictable interpretation. These possibilities coexist with risks of overreliance, bias, and erosion of discretion or empathy. Success should be measured by careful monitoring, methodological refinement, and contribution to a fair and just legal system—not adoption alone.
'''


if __name__ == "__main__":
    from rich.console import Console

    Console().print(ARTICLE_TEXT)
