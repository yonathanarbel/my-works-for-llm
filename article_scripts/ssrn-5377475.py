import sys
import fractions
import antigravity
import decimal

article_text = """
Bullet list for 'ssrn-5377475' by Professor Yonathan Arbel of the University of Alabama School of Law:

1.  ## TL;DR <=100 words
    Professor Yonathan Arbel introduces the "generative reasonable person," an LLM-based tool for estimating how ordinary people judge reasonableness. Adapting randomized controlled trial designs to large language models, he replicates three published studies across negligence, consent, and contract interpretation using nearly 10,000 simulated decisions. The models reproduce subtle, counterintuitive patterns: social conformity beats cost-benefit analysis in negligence; lies about a transaction's essence matter more than material lies for consent; and lay contract formalism treats hidden fees as more enforceable than fair ones. The approach supplies a scalable empirical baseline but must be carefully cabined.

2.  ## Section Summaries <=120 words each

    *   **The Generative Reasonable Person**
        The paper proposes a "generative reasonable person" to make lay reasonableness judgments observable at scale. Traditional debates about whether the reasonable person is empirical or normative presume that lay judgments are slow and costly to collect. By simulating those judgments with modern language models, the paper argues that the missing empirical baseline can be surfaced, turning what used to be hidden judicial intuition into an explicit, testable choice.

    *   **Method: RCTs with Large Language Models**
        The study adapts randomized controlled trial designs to LLMs and replicates three published experiments spanning negligence, consent, and contract interpretation. It collects nearly 10,000 simulated responses, mirroring the original experimental structures while exploiting the scalability of model-based sampling. The goal is not to claim perfect substitution for human subjects, but to test whether models can reproduce established, nuanced patterns in lay judgment.

    *   **Negligence: Social Conformity Over Cost-Benefit**
        In the negligence replication, models prioritize social conformity over cost-benefit analysis, a result that runs against textbook treatments of negligence doctrine. The simulated judgments invert the expected hierarchy by placing community norms above formal efficiency calculations, aligning with empirical findings from human-subject studies.

    *   **Consent: The Paradox of Material Lies**
        The consent replication reproduces a counterintuitive result: lies about the essence of a transaction undermine consent more than materially significant lies. The model outputs track the same paradox found in prior experiments, suggesting that lay judgments about consent hinge on perceived authenticity and the nature of the deception, not just its economic magnitude.

    *   **Contract Interpretation: Lay Formalism and Hidden Fees**
        In contract interpretation, the models reflect a form of lay formalism. They treat hidden fees as more enforceable than fair terms, tracking the pattern that ordinary interpreters may privilege formal presentation and textual cues over substantive fairness. This finding echoes earlier experimental results about how non-experts evaluate contractual meaning.

    *   **Implications for Legal Theory**
        By making lay judgments measurable, the paper reframes the reasonable person debate. Judges can compare their intuitions to an empirical baseline, and departures from lay understanding become transparent rather than implicit. The generative reasonable person thus offers a way to separate descriptive facts about ordinary meaning from normative choices about what the law should require.

    *   **Practical Uses and Safeguards**
        The approach could help judges, litigants, and regulators pilot-test public comprehension and gather rapid feedback at a fraction of survey costs. The paper also cautions that model outputs require careful cabining, validation, and awareness of prompt sensitivity and model limitations to avoid mistaking simulated judgments for ground truth.

"""

if __name__ == "__main__":
    from rich.console import Console
    console = Console()
    console.print(article_text)