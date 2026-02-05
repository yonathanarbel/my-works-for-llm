# The Generative Reasonable Person — one‑page summary

**Paper:** “The Generative Reasonable Person” (SSRN 5377475, 2026)  
**Author:** Yonathan A. Arbel (University of Alabama School of Law)  
**In this repo:** `papers/ssrn-5377475/paper.pdf`, `papers/ssrn-5377475/paper.txt`, `papers/ssrn-5377475/summary.md`

## What the paper introduces
The paper proposes the **generative reasonable person**: using large language models (LLMs) as a scalable way to *estimate* how ordinary people (lay decision-makers) judge **reasonableness**—in tort, consent, and contract settings—where courts often speak in “no reasonable person would…” terms without a clean empirical baseline.

## The core method (Silicon RCTs / s‑RCTs)
The paper adapts **randomized controlled trials** to LLMs by treating **stateless, independent model sessions** (API calls) as experimental “subjects”:
- Randomize experimental conditions *across sessions* so the model can’t harmonize answers across conditions.
- Use vignettes unlikely to be verbatim training memorization targets.
- Optionally add **persona prompting** (demographically specified respondents) and compare against an “empty prompt” baseline.
- Focus on whether models replicate *directional* human patterns (structure) and then separately examine *magnitudes* (calibration).

## What the replications show (three studies)
The paper runs silicon replications of three published experiments (nearly **10,000** simulated judgments total):

1. **Negligence (Jaeger, “The Empirical Reasonable Person,” 2021).**  
   A 2×2 design varies (i) whether a precaution is *common* and (ii) whether it is *cost‑justified*. Human subjects’ negligence judgments strongly track **social norms** (commonness) and barely track cost. LLMs reproduce the same *hierarchy* (norms dominate), though many models show **stronger cost sensitivity** than humans.

2. **Consent under deception (Sommers, “Commonsense Consent”).**  
   Subjects evaluate consent when deception is an **essential lie** (about what the thing is) versus a **material lie** (about a term that matters a lot). Humans show a paradox: material lies are judged as *more important to the victim* yet perceived as *less corrosive of consent* than essential lies. LLMs reproduce this **counterintuitive structure** across architectures.

3. **Contract interpretation / hidden fees (Furth‑Matzkin & Sommers, 2020).**  
   Lay subjects (and legal professionals) rate consent, fairness, and enforceability in a “no fee” pitch + fine‑print fee vignette. Lay reasoning follows a stable hierarchy: **enforceability > consent > fairness** (a kind of “lay formalism”). LLMs reproduce the hierarchy and, under many prompts, land **closer to lay baselines than to elite legal baselines**; persona prompting tends to nudge models further toward lay calibration.

## Why it matters
If properly validated and transparently deployed, generative reasonable people can:
- Provide **empirical guardrails** for judicial intuition (making departures from lay baselines explicit rather than implicit).
- Give **resource‑constrained litigants** access to simulated jury feedback (as a pre‑test / screen).
- Help **regulators and compliance teams** pilot‑test consumer understanding faster than surveys.

## Guardrails / limitations (the paper’s emphasis)
- **Structure vs magnitude:** models may capture *what matters* to people better than *how much* it matters.
- **Prompt/model sensitivity:** outputs depend on model choice and prompting regime; report them.
- **Transparency & triangulation:** treat outputs as inputs to be audited (disclose model, prompts, personas; validate against surveys/focus groups when stakes are high).
- **Bias & representativeness:** persona testing helps explore heterogeneity but requires careful auditing.

