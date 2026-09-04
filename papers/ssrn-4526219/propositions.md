# Propositions from Generative Interpretation

**Citation:** Yonathan A. Arbel & David A. Hoffman, Generative Interpretation, 99 N.Y.U. L. Rev. 451 (2024)

**Source:** [final published PDF](https://www.nyulawreview.org/wp-content/uploads/2024/05/99-NYU-L-Rev-451-1.pdf)

**Review status:** 12 model-drafted, source-checked; 0 human-reviewed. Page references use the printed pagination and, separately, the 1-based PDF page number.

## 1. Generative interpretation uses language models as an aid for reconstructing contractual meaning

**Location:** Introduction, printed pp. 455-460 (PDF pp. 5-10)

Professors Yonathan A. Arbel and David A. Hoffman claim, in “Generative Interpretation” on pages 455–460, that large language models can examine an agreement together with relevant context and generate disciplined estimates of what the parties meant. They call this method generative interpretation and present it as a lower-cost, more replicable, and more transparent adjunct to judicial interpretation. This is significant because it reframes language models from generic legal chatbots into instruments for testing interpretive intuitions against linguistic patterns. It connects to the article’s later case studies of ordinary meaning, ambiguity, gap filling, and extrinsic evidence, while leaving the ultimate legal judgment with courts.

**Evidence anchor:** The introduction defines the method, describes its expected cost and consistency benefits, previews grounded contract case studies, and repeatedly characterizes the model as a judicial aid rather than a decisionmaker.

**Boundary:** The authors do not claim that a plausible model answer is ground truth, that their case studies measure judicial accuracy at scale, or that models should replace judges.

**Connections:** contract interpretation methodology; computational linguistics; judicial decision support; access to justice

**Record:** `ssrn-4526219-p01` · `machine-drafted-source-checked`

## 2. Contract interpretation is substantially a backward-looking prediction about meaning, but prediction cannot settle every legal question

**Location:** Part I.A, Interpretation as Prediction, printed pp. 461-464 (PDF pp. 11-14)

Professors Yonathan A. Arbel and David A. Hoffman claim, in “Generative Interpretation” on pages 461–464, that leading approaches to contract interpretation share a substantial predictive ambition: they seek to reconstruct what the parties, reasonable parties, or the relevant linguistic community would have understood at formation. The approaches diverge over whose meaning counts, what evidence should inform the prediction, and what legal consequences follow. This is significant because it identifies a common task that a language model can assist without pretending that interpretive theory has become value-free. It connects to debates over subjective intent, objective meaning, textualism, contextualism, and the distinction between an empirical prediction and a court’s normative choice.

**Evidence anchor:** Part I.A organizes interpretive approaches around backward-looking prediction and then identifies unresolved questions about the target, evidentiary basis, and legal significance of that prediction.

**Boundary:** The authors treat prediction as an important shared component of interpretation, not as a complete account of every interpretive theory or every judicial responsibility.

**Connections:** objective theory of contract; subjective intent; interpretation versus construction; empirical legal analysis

**Record:** `ssrn-4526219-p02` · `machine-drafted-source-checked`

## 3. Existing interpretive methods trade off evidentiary richness, cost, consistency, and bias

**Location:** Part I.B–D, Textualism, Contextualism, and Empirical Methods, printed pp. 464-473 (PDF pp. 14-23)

Professors Yonathan A. Arbel and David A. Hoffman claim, in “Generative Interpretation” on pages 464–473, that neither textualism nor contextualism escapes institutional tradeoffs. Textualism controls cost and can improve predictability, yet dictionaries, canons, and judges’ linguistic intuitions leave room for selection and hindsight; contextualism admits richer evidence, yet discovery and factfinding are costly and can expose decisionmakers to bias and strategic behavior. This is significant because the familiar doctrinal disagreement partly reflects the limitations of available interpretive technologies rather than an unavoidable choice between text and context. It connects to corpus linguistics and survey experiments, which discipline intuition in useful ways but remain constrained by context, sample design, expense, or limited judicial adoption.

**Evidence anchor:** The article compares textual and contextual regimes, critiques apparently objective aids such as dictionaries and canons, and evaluates corpus and survey methods as partial empirical responses.

**Boundary:** The discussion does not establish that any one conventional method is uniformly inferior; it emphasizes that each serves values and incurs costs that vary by dispute and party.

**Connections:** textualism; contextualism; corpus linguistics; survey evidence; litigation costs

**Record:** `ssrn-4526219-p03` · `machine-drafted-source-checked`

## 4. LLMs can produce context-sensitive linguistic predictions even though their internal reasoning remains opaque

**Location:** Part II.A, Grounding Generative Interpretation, printed pp. 473-483 (PDF pp. 23-33)

Professors Yonathan A. Arbel and David A. Hoffman claim, in “Generative Interpretation” on pages 473–483, that transformer-based language models use learned statistical relationships and attention to predict language in context across a vast body of training material. That architecture lets a model integrate more of a contract and its surroundings than a dictionary lookup or a narrow corpus query, but the output remains a prediction rather than a transparent causal explanation of how people actually spoke or thought. This is significant because interpretive usefulness can coexist with mechanistic opacity: a tool may test linguistic probabilities without supplying a human-style rationale for them. It connects to the interpretability problem in machine learning, the law’s demand for reason-giving, and the need to distinguish an evidentiary signal from a judicial explanation.

**Evidence anchor:** Part II.A describes training, embeddings, attention, context windows, and probabilistic output, while stressing that even model builders cannot fully explain particular predictions.

**Boundary:** The authors do not equate next-token prediction with human understanding, causal proof, or a self-justifying legal conclusion.

**Connections:** transformer attention; language-model prediction; explainable AI; judicial reason-giving

**Record:** `ssrn-4526219-p04` · `machine-drafted-source-checked`

## 5. A language model can check judicial confidence about ordinary meaning by exposing a competing probabilistic reading

**Location:** Part II.B, The Ordinary Meaning Problem, printed pp. 483-485 (PDF pp. 33-35)

Professors Yonathan A. Arbel and David A. Hoffman claim, in “Generative Interpretation” on pages 483–485, that the Famiglio prenuptial dispute shows how a language model can test a court’s asserted ordinary meaning. After receiving the agreement and the sequence of two divorce filings, the model treated the second filing as the more natural date for calculating years of marriage, contrary to the appellate court’s confident reliance on the indefinite article and a golf-course analogy. This is significant because the model’s contrary reading makes judicial certainty itself contestable even when it does not prove that the judge was wrong. It connects to ordinary-meaning doctrine, representativeness of judicial intuitions, probabilistic language, and the possible relevance of private meaning or trade context.

**Evidence anchor:** The article reconstructs the prenup dispute, contrasts the court’s first-filing interpretation with the model’s second-filing prediction, and explains why the divergence should prompt further reflection.

**Boundary:** One model response cannot establish the parties’ actual intent, and case-specific extrinsic evidence could justify a meaning different from the model’s public-language prediction.

**Connections:** ordinary meaning; judicial intuition; probabilistic semantics; private meaning

**Record:** `ssrn-4526219-p05` · `machine-drafted-source-checked`

## 6. Model outputs can represent ambiguity as a distribution of plausible readings rather than a binary intuition

**Location:** Part II.C, The Ambiguity Problem, printed pp. 485-492 (PDF pp. 35-42)

Professors Yonathan A. Arbel and David A. Hoffman claim, in “Generative Interpretation” on pages 485–492, that language models can help courts see ambiguity as a spectrum of plausible interpretations. In Trident, several models leaned against a borrower’s asserted prepayment right while still displaying a minority probability; in Ellington, repeated outputs across multiple prompt variations more often read “other affiliates” to include later-created affiliates than the state high court did. This is significant because a distribution can expose plausible minority meanings and check a court’s confidence without collapsing the legal ambiguity threshold into a model score. It connects to summary-judgment screening, linguistic communities and private meanings, robustness testing across prompts and models, and the separate judicial question of how much plausibility is legally enough.

**Evidence anchor:** The section compares multi-model results in Trident and repeated, varied prompts in Ellington, then treats the resulting distributions as checks on confidence rather than dispositive rulings.

**Boundary:** The models do not decide whether a reading is legally reasonable, numerical scores should not be treated as calibrated facts, and prompt framing is itself a consequential choice.

**Connections:** contractual ambiguity; probability distributions; minority linguistic meanings; prompt robustness; summary judgment

**Record:** `ssrn-4526219-p06` · `machine-drafted-source-checked`

## 7. LLMs can test proposed gap fillers against the whole agreement and reveal both convergence and unresolved disagreement

**Location:** Part II.D, Filling Gaps, printed pp. 492-495 (PDF pp. 42-45)

Professors Yonathan A. Arbel and David A. Hoffman claim, in “Generative Interpretation” on pages 492–495, that a language model can help a court ask what parties would likely have provided for an omitted contingency by evaluating candidate rules against the full agreement. In the Haines sewage-contract study, two models rejected termination at will and were open to several duration rules, yet they differed over whether the city’s obligations expanded with future community growth. This is significant because agreement between models can strengthen a textual inference while disagreement can direct attention to overlooked provisions and competing limiting principles. It connects to default rules, incomplete contracts, the boundary between interpretation and construction, and the common-law practice of implying terms.

**Evidence anchor:** The authors feed two models the lengthy 1924 agreement, compare responses on duration and scope, and highlight both shared rejection of termination at will and divergent readings of an expansion clause.

**Boundary:** The models disagreed on an important scope question, and their assessments cannot choose the legally proper gap-filling rule or establish the historical parties’ actual counterfactual agreement.

**Connections:** incomplete contracts; default rules; interpretation and construction; implied terms; model disagreement

**Record:** `ssrn-4526219-p07` · `machine-drafted-source-checked`

## 8. Adding extrinsic evidence sequentially can reveal its marginal effect on an interpretation

**Location:** Part II.E, From Text to Context, printed pp. 495-497 (PDF pp. 45-47)

Professors Yonathan A. Arbel and David A. Hoffman claim, in “Generative Interpretation” on pages 495–497, that contextual evidence can be introduced to a model in stages to test how each addition changes the predicted meaning of a contract. In Stewart, they begin with a sparse construction agreement and an assumed payment default, then add evidence of a phone conversation and an asserted industry custom to observe changes in the models’ assessment of monthly payment. This is significant because the direction of change can help a court estimate whether expensive discovery into a category of extrinsic evidence is likely to matter. It connects to contextualism, the marginal probative value of evidence, proportional discovery, and staged sensitivity analysis.

**Evidence anchor:** The case study describes a baseline prompt and two successive additions of context, then compares how model assessments move as the phone call and trade custom enter the record.

**Boundary:** The assumed legal default may be contestable, the historical record is sparse, and the authors caution that the models’ self-reported confidence levels should not be read literally; the direction of change is more informative.

**Connections:** extrinsic evidence; contextualism; sensitivity analysis; proportional discovery; industry custom

**Record:** `ssrn-4526219-p08` · `machine-drafted-source-checked`

## 9. The relevant institutional test is whether generative interpretation is good enough for ordinary, resource-constrained adjudication

**Location:** Part III.A, Applications and Pitfalls, printed pp. 499-503 (PDF pp. 49-53)

Professors Yonathan A. Arbel and David A. Hoffman claim, in “Generative Interpretation” on pages 499–503, that the practical benchmark for generative interpretation is not whether it always surpasses ideal, labor-intensive linguistic analysis. The more urgent comparison is with ordinary adjudication in resource-deprived courts, where inexpensive and accessible model assistance may improve consistency, settlement calibration, and the position of parties who lack repeat-player expertise. This is significant because it places access to justice and opportunity cost at the center of technology assessment instead of comparing automation only with the best imaginable human performance. It connects to unequal legal information, litigation budgets, predictive settlement, clearer ex ante contracting, and a more broadly accessible form of textual analysis.

**Evidence anchor:** The article links low-cost prediction to information equality and settlement while expressly asking whether the method is sufficiently competent for ordinary courts, rather than superior to careful artisanal analysis in all cases.

**Boundary:** The authors present this as a promise and a competency question, not as proof that current unspecialized models are reliable in every ordinary case; their case studies are curated rather than representative.

**Connections:** access to justice; resource-constrained courts; repeat-player advantage; predictive settlement; institutional comparison

**Record:** `ssrn-4526219-p09` · `machine-drafted-source-checked`

## 10. Reliable legal use requires cross-checking outputs and governing prompts, models, and disclosure

**Location:** Part III.A.1–2, Hallucinations and Manipulation, printed pp. 503-505 (PDF pp. 53-55)

Professors Yonathan A. Arbel and David A. Hoffman claim, in “Generative Interpretation” on pages 503–505, that hallucinations and strategic prompt design require procedural safeguards around generative interpretation. They propose comparing models and multiple inputs, scrutinizing party-supplied framing, disclosing the model and prompts, and allowing contracting parties to specify a model in advance. This is significant because reproducibility depends on governing the whole interpretive setup, not merely preserving a model’s final sentence. It connects to adversarial presentation, expert-method disclosure, model versioning, contractual choice of interpretive method, and the creation of a persistent record that later readers can audit.

**Evidence anchor:** The pitfalls discussion responds to fabricated outputs and framing effects with multiple-model checks, varied inputs, disclosure, and possible ex ante party choice over the model used.

**Boundary:** Cross-model agreement is not truth, safeguards can add cost, and party control over model selection cannot eliminate manipulation, hallucination, or changes between model versions.

**Connections:** AI governance; reproducibility; adversarial procedure; model choice clauses; methodological disclosure

**Record:** `ssrn-4526219-p10` · `machine-drafted-source-checked`

## 11. Majoritarian training data, adversarial inputs, opacity, and linguistic drift define the domain in which LLM interpretation is safe and useful

**Location:** Part III.A.3–7, Interpretability, Bias, Attacks, and Linguistic Drift, printed pp. 505-509 (PDF pp. 55-59)

Professors Yonathan A. Arbel and David A. Hoffman claim, in “Generative Interpretation” on pages 505–509, that courts must limit and qualify LLM use because model opacity, majoritarian training patterns, adversarial inputs, and temporal drift can distort contractual meaning. A model may suppress a local, private, or minority linguistic practice; hidden instructions in a document may manipulate it; and contemporary training data may misread an old agreement through later usage or later decisions. This is significant because the same scale that makes an LLM sensitive to public language can make it unreliable for historically bounded or nonmajoritarian meaning. It connects to algorithmic bias, cybersecurity, historical corpus methods, linguistic communities, and the authors’ insistence that models assist textual analysis rather than make human-critical legal decisions.

**Evidence anchor:** The section enumerates interpretability limits, majoritarian and discriminatory effects, adversarial attacks, and contamination by later language, then recommends caution, tailored models, and a preserved methodological record.

**Boundary:** The article offers risk-management directions rather than a demonstrated cure; specialized time-bounded models may not exist, and transparent disclosure does not make an opaque model fully explainable.

**Connections:** algorithmic bias; prompt injection; historical semantics; private language; human judicial responsibility

**Record:** `ssrn-4526219-p11` · `machine-drafted-source-checked`

## 12. Generative interpretation offers a contingent third path between textualism and contextualism while preserving party choice and judicial authority

**Location:** Part III.B and Conclusion, Beyond the Textualist/Contextualist Divide, printed pp. 510-514 (PDF pp. 60-64)

Professors Yonathan A. Arbel and David A. Hoffman claim, in “Generative Interpretation” on pages 510–514, that LLM-assisted interpretation can disrupt the inherited choice between predictable but narrow textualism and information-rich but expensive contextualism. If models can absorb broader evidence consistently and estimate the incremental value of context, courts may be able to relax categorical exclusions of extrinsic evidence while parties retain the ability to choose, constrain, or reject the method. This is significant because it treats interpretive doctrine as partly dependent on adjudicatory technology and gives the new method possible distributive consequences for uncounseled and poorer parties. It connects to party autonomy, interpretive defaults, the parol evidence rule, relational contracting, and the prospect of a distinct methodology that supplements rather than replaces judicial judgment.

**Evidence anchor:** The final section argues that LLMs can combine predictability with broader evidence, discusses party control and distributional effects, and closes by preserving a separate role for judicial legal judgment.

**Boundary:** The method is not appropriate for every contract, parties should be able to opt out or specify alternatives, and the authors expressly stop short of replacing judges or resolving the legal significance of model predictions.

**Connections:** textualism-contextualism divide; party autonomy; interpretive defaults; parol evidence rule; distributional effects

**Record:** `ssrn-4526219-p12` · `machine-drafted-source-checked`
