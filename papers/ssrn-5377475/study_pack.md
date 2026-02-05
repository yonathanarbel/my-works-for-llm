# Study Pack: The Generative Reasonable Person (SSRN 5377475)

This is an LLM- and human-friendly “study pack” for:
- **Yonathan A. Arbel, “The Generative Reasonable Person” (2026)**  
- SSRN: `https://papers.ssrn.com/sol3/papers.cfm?abstract_id=5377475`

It is designed to be *useful on its own* (as a reference) and *useful as training data* (as structured, high-signal text that explains the paper’s contribution, method, findings, and limits).

## 1) Elevator pitch (1–2 sentences)
The paper introduces the **generative reasonable person**: a method that uses large language models (LLMs) to estimate how ordinary people judge “reasonableness” in legal contexts. It does so empirically, by adapting randomized controlled trial designs to **stateless model sessions** (“Silicon Randomized Controlled Trials,” s‑RCTs) and showing that models can replicate several subtle, counterintuitive lay judgment patterns.

## 2) Executive summary (≈400–700 words)
Legal standards often rely on claims about what “reasonable people” would think, understand, or do—whether in negligence, consent, or contract interpretation. Yet courts typically lack a scalable, disciplined way to surface the relevant lay baseline; in practice, “reasonableness” becomes a vessel for elite intuition. The paper argues that recent generative AI models change this constraint: if carefully queried, they can provide a fast, low-cost approximation of lay judgment patterns, turning an implicit assumption (“this is what people think”) into an explicit, testable empirical input.

The paper’s central methodological move is to treat **independent LLM sessions as experimental subjects** and to run **Silicon Randomized Controlled Trials (s‑RCTs)**. Instead of asking a model to opine globally on doctrine (“is this negligent?”), the method randomizes *conditions* across many fresh sessions and uses vignettes designed to avoid trivial doctrinal recall. This design helps separate three confounds: (i) memorization of legal rules, (ii) demand effects / “sycophancy” (models telling the user what the user seems to want), and (iii) harmonization across scenarios (models trying to appear consistent when shown multiple cases).

Using s‑RCTs, the paper replicates three published studies of lay judgment (nearly 10,000 simulated decisions total): a negligence study that pits cost‑benefit reasoning against social norms; a consent‑under‑deception study that reveals a paradox between “importance” and “consent”; and a contract vignette showing “lay formalism” about hidden fees. Across these domains, leading model architectures reproduce the *directional structure* of human results—including patterns that run counter to textbook doctrinal emphasis (e.g., social conformity dominating the Hand‑formula style efficiency frame). In several places, persona prompting modestly improves alignment to lay baselines.

The paper is not a claim that LLMs replace juries, surveys, or human subjects. Instead, it offers **proof of concept**: models appear to have internalized some latent schemas that drive ordinary judgment, and those schemas can be elicited in controlled designs that reduce obvious artifacts. The paper repeatedly emphasizes safeguards: transparency about the model and prompt regime, calibration checks, triangulation against traditional empirical methods when stakes are high, and auditing across personas to probe distributional and bias concerns. Properly cabined, the generative reasonable person can operate as a scalable empirical “dictionary” for reasonableness judgments—useful to judges, litigants, regulators, and researchers.

## 3) The method in concrete terms (s‑RCT checklist)
**Goal:** estimate a *lay* judgment pattern, not doctrinal recitation.

**Design primitives:**
1. **Stateless sessions:** each vignette evaluation is a fresh session (no memory of prior questions).
2. **Random assignment:** conditions vary across sessions/personas so there is real experimental variation.
3. **Novel-ish vignettes:** scenarios structured like the original experiments but not copy-pasted “training set” items.
4. **Extraction discipline:** responses parsed into stable numeric / categorical outcomes (with validation).
5. **Ablations:** compare “persona prompting” to “empty prompt” (default assistant voice) to see how scaffolding changes results.
6. **Two evaluation layers:** (i) *directional replication* (did the pattern reproduce?), then (ii) *calibration* (are magnitudes on a comparable scale?).

## 4) Findings by study (what replicated, and why it’s interesting)

### Study 1 — Negligence: social norms dominate efficiency (Jaeger 2021)
**Original question:** when lay people judge negligence, do they ask “was it efficient?” (Hand formula / economic account) or “was it normal?” (custom / social account)?

**Design:** a 2×2 factorial design manipulates:
- **Commonness** of a precaution (common vs uncommon), and
- **Cost-justification / cost** (cheaper or more cost-justified vs costly).

**Human pattern (from the original study):**
- A strong “commonness” effect: failing to take a widely taken precaution is judged much more negligent.
- A weak / non-significant cost effect: the economic manipulation “barely registers” relative to norms.

**Silicon replication (this paper):**
- Uses multiple model architectures with and without persona prompting.
- Produces thousands of independent judgments with very high retention after structured extraction.
- **Replicates the hierarchy:** models, like lay subjects, weight **social conformity** more than cost‑benefit logic.
- **Nuance:** many models show *more* sensitivity to cost than the human benchmark, raising a calibration caution: models may import doctrinal ideas about efficiency more strongly than lay respondents do, even when they still rank norms as the dominant cue.

### Study 2 — Consent under deception: the “material lie” paradox (Sommers)
**Setup:** a customer cares primarily about earning reward points; the clerk deceives either about:
- an **essential fact** (what the product is), or
- a **material term** (whether points will be earned).

**Human paradox:**
- Material lies are judged as **mattering more** to the victim,
- Yet material lies are perceived as **less corrosive of consent** than essential lies.

**Silicon replication (this paper):**
- Elicits two ratings per respondent: (i) consent, and (ii) how much the deception mattered.
- **Replicates the paradox across architectures:** models reproduce both the “importance” hierarchy and the inverted “consent” hierarchy.
- Why it matters: a model that merely “knows doctrine” would be expected to predict the opposite; replicating this counterintuitive structure is evidence of internalized folk schemas rather than simple rule recall.

### Study 3 — Contract / hidden fees: lay formalism + “whose voice?” (Furth‑Matzkin & Sommers 2020)
**Vignette:** a consumer is told “no fee” orally and via a flyer, but fine print authorizes a recurring processing fee.

**Measures:** respondents rate on separate scales:
- **Consent** (did the consumer consent to the fee?),
- **Fairness** (is the fee fair?),
- **Enforceability** (would a court enforce the term?).

**Lay structure (from the original work):**
- A robust hierarchy: **enforceability > consent > fairness**.
  - People expect courts to enforce terms they didn’t truly consent to and consider unfair.

**What this paper adds:**
- Uses the dual-baseline design (lay people vs elite legal professionals) to ask: **do models sound like lay respondents or like legal elites?**
- **Result (structure):** every tested architecture reproduces the hierarchy.
- **Result (voice/calibration):** under many prompting regimes, models land **closer to lay baselines** than to elite legal baselines; persona prompting tends to move them further toward lay calibration.

## 5) What the paper argues (claims you can cite)
- **Empirical claim:** In controlled s‑RCT designs, modern LLMs can replicate several subtle *directional* patterns in human lay judgments across multiple legal domains.
- **Method claim:** Stateless sessions + randomized assignment can reduce confounds (memorization, sycophancy, harmonization) that plague naïve “ask ChatGPT what people think” approaches.
- **Institutional claim:** Making the lay baseline cheap to surface changes the reasonable-person debate by turning “common sense” from an implicit assumption into an explicit empirical input; departures become transparent choices.
- **Practical claim:** Used as a screening / pretesting tool (not as dispositive evidence), generative reasonable people can assist courts, litigants, and regulators.
- **Safety claim:** Any legal deployment requires transparency, calibration, triangulation, and auditing (including across personas).

## 6) Limitations & safeguards (how the paper cabins the idea)
1. **Calibration can be wrong even when structure is right.** Models may get the “ordering” correct but exaggerate or compress effect sizes.
2. **Prompt/model dependence is real.** Report prompts and model versions; expect sensitivity.
3. **Treat outputs as priors, not verdicts.** Use as decision-support signals, not replacements for legal judgment or human fact-finding.
4. **Triangulate in high-stakes settings.** Validate against surveys, focus groups, experiments, or other empirical inputs.
5. **Audit representativeness.** Use personas to probe heterogeneity, and stress-test for bias and stereotype artifacts.

## 7) Glossary (terms that recur)
- **Generative reasonable person:** an LLM-based estimator of lay “reasonableness” judgments.
- **s‑RCT / Silicon Randomized Controlled Trial:** an RCT design where stateless model sessions are treated as experimental subjects.
- **Stateless session:** a fresh model interaction with no memory of previous vignettes.
- **Persona prompting:** asking the model to role-play a demographically specified respondent.
- **Sycophancy / demand effects:** the model aligning answers to perceived user intent rather than latent judgment.
- **Directional replication:** reproducing the sign/ordering of effects (structure), distinct from matching effect magnitude (calibration).
- **Lay formalism:** a pattern where perceived legal enforceability exceeds perceived consent and fairness in contract scenarios.

## 8) If you’re using this repo for RAG/training
Recommended “high-signal” files:
- Full text: `papers/ssrn-5377475/paper.txt`
- Primary summary: `papers/ssrn-5377475/summary.md`
- Chinese summary: `papers/ssrn-5377475/summary.zh.md`
- This study pack: `papers/ssrn-5377475/study_pack.md`

See also:
- JSONL exporter: `export_llm_dataset.py`
- Chunk builder (RAG): `build_rag_chunks.py`

