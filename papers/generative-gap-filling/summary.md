# Generative Gap Filling

Yonathan A. Arbel and David A. Hoffman (2026)

1. ## TL;DR

   Contract law assumes that a textual gap marks the point where interpretation ends and judicial supplementation begins. This paper tests that premise by removing negotiated terms from real contracts and asking lay readers, law students, practicing lawyers, and six frontier language models to reconstruct them. Lay readers were correct 55% of the time, lawyers nearly 60%, and the models 88.3%; on 119 additional unseen commercial contracts, the models recovered the masked clause 87% of the time. The results suggest that surrounding contract language often makes the hypothetical bargain statistically legible, while also identifying where model-based gap filling fails and how courts and parties might govern its use.

2. ## Research Question and Contribution

   The paper asks whether a contract that appears incomplete has really "run out of meaning." It turns the interpretation-construction distinction into a testable question: if readers can predict a term that the parties actually negotiated but the researchers later hid, then the remaining document contains recoverable information about the supposed gap. The masking design manufactures a ground truth for legal interpretation, answering the objection that interpretive methods cannot be graded because real disputes lack an answer key.

3. ## Experimental Design

   The main study masks material terms in three executed agreements involving an artist, a contingency fee, and a bottle-supply arrangement. Participants receive the surrounding contract, a triggering scenario, and four possible answers. The human sample includes 465 attentive lay respondents recruited through Prolific, 77 law students, and 48 experienced lawyers. The model panel includes six frontier systems, twenty runs per model, with browsing disabled and answer order varied. A stricter follow-up question tests whether respondents recovered the operative meaning rather than guessed the headline result.

4. ## Main Results

   Lay respondents identified the hidden term 55% of the time, compared with 25% chance accuracy. Law students performed only marginally better, while practicing lawyers reached nearly 60% overall and 82.6% in the familiar contingency-fee setting. The model panel achieved 88.3% scenario-balanced accuracy: Claude Opus 4.6 reached 100%, Gemini 3.1 Pro, Grok 4.2, and GPT-5.4 scored between 88% and 97%, Qwen 3.6 scored 82%, and GLM 5.1 scored 70%. The stricter two-question measure reduced every group’s performance, and exposed a shared model failure on a technical pricing follow-up in the bottle contract.

5. ## Robustness and Scaling

   The authors perturb the original agreements to separate document-specific signals from general expectations about contract structure. Flipping the contractual language reduced model accuracy from 88% to 60%; withholding the contract produced 68% accuracy. The paper estimates that roughly two-thirds of model success comes from general contractual expectations and one-third from the particular language of the agreement. In a separate sample of 119 mostly 2025-2026 SEC EDGAR agreements, six models recovered the masked clause 87% of the time, with individual models between 84% and 92%. Performance was strongest for standardized instruments such as promissory notes and credit facilities, and weakest for bespoke or default-reversing provisions.

6. ## Doctrinal Implications

   The findings blur the conventional line between interpretation and construction. If sophisticated readers can recover omitted terms from the visible agreement, then some apparent gap filling may be evidence-based reconstruction rather than free-ranging judicial policymaking. The paper does not argue that statistical prediction automatically determines the legal answer. Instead, it separates the positive question—what the parties likely would have written—from the normative question—whether a court should enforce that reconstructed bargain.

7. ## Institutional Design

   The authors propose adversarial use rather than autonomous adjudication. Parties should disclose the model, version, prompts, and relevant outputs; opponents should be able to contest model selection and sensitivity; and courts should write narrow, reviewable opinions. Contracting parties can adopt "Choice of Model" clauses that designate a model or protocol for later interpretation, much as agreements choose law or forum. Better prediction may also reduce disputes by making outcomes easier to anticipate.

8. ## Limits and Failure Modes

   Accuracy is not uniform. Models struggle with bespoke clauses, provisions that reverse market defaults, and some technical follow-up questions. Training-data contamination cannot be ruled out completely, and model outputs remain sensitive to prompt design, model version, and calibration choices. The method also assumes a human-drafted agreement: when AI agents draft contracts, a model may reconstruct machine-generated regularities without recovering any human intent. The paper therefore treats model predictions as contestable evidence, not binding answers.

9. ## Relationship to Prior Work

   *Generative Gap Filling* is the empirical companion to Arbel and Hoffman’s *Generative Interpretation*, 99 N.Y.U. Law Review 451 (2024), which argued that language models could assist contract interpretation but acknowledged the absence of ground truth in litigated disputes. The new paper supplies that ground truth through masking and extends the project from interpreting existing language to reconstructing contractual terms that appear missing.

10. ## Citation

   Yonathan A. Arbel & David A. Hoffman, *Generative Gap Filling* (working paper 2026).
