# The Readability of Contracts: Big Data Analysis

## TL;DR

Professor Yonathan Arbel challenges the empirical foundations of the plain-language movement. Using nearly two million commercial, consumer-credit, privacy-policy, and franchise documents, plus a large comparison corpus of texts adults voluntarily read, he finds that consumer credit-card agreements have median readability scores close to ordinary news. Privacy policies are somewhat harder, while commercial contracts and franchise disclosure documents are substantially harder. The results therefore do not support a general claim that consumer contracts are linguistically beyond the reach of most American adults.

The paper also finds that the standard formulas used to establish and regulate readability are unreliable. Different formulas correlate weakly, and different software implementations of the same Flesch-Kincaid formula produce results that differ by an average of 4.6 grade levels on identical text. The formulas explain only a modest share of actual comprehension and behavior in the empirical literature. Arbel further traces the familiar claim that most Americans read only at a sixth- or eighth-grade level to unsupported or mischaracterized sources; direct literacy and educational data do not sustain it.

## Research design

- The source-draft reports a final corpus of 1,935,680 contracts: roughly 1.2 million SEC-filed commercial contracts and 735,680 consumer or hybrid agreements.
- A weakly supervised classifier uses lawyer-supplied contract titles as initial labels, removes those titles before training, and classifies contracts by substantive content. The selected TF-IDF and LightGBM approach achieved 96.67% balanced accuracy within the defined categories and 78% in a broader blind audit.
- The study improves robustness through strict preprocessing, multiple sentence tokenizers, averaging across software implementations of each formula, and a Composite Readability Measure that averages across formulas.
- A benchmark corpus of approximately 1.8 million news articles, reviews, debates, books, and reference texts anchors formula scores in materials adults actually choose to read.

## Principal findings

- Credit-card agreements fall near CNN and Fox News in median composite readability; privacy policies fall near financial news and *The New Republic*.
- Commercial contracts and franchise disclosure documents occupy the difficult end of the distribution.
- Consumer contracts are much longer than comparison texts even when their sentence-level readability is similar. The rejection of a generalized linguistic crisis is therefore not a rejection of length and attention problems.
- Commercial contracts became more complex from 2001 to 2022, while consumer-contract readability remained broadly stable. Most contract categories did not grow materially longer.
- Noncompete, confidentiality, and post-employment agreements showed an especially notable increase in difficulty.
- ALL-CAPS remains common in commercial, franchise, and credit documents even though it does not improve readability.

## Normative implications

Arbel argues that readability is neither necessary nor sufficient to cure consumer-market pathologies. Clear disclosure cannot make a substantively exploitative bargain fair, while a fair and standard agreement need not become harmful merely because its prose is imperfect. Regulators should pay more attention to market structure, consumer choice, substantive terms, enforcement, and the needs of genuinely vulnerable readers.

The paper recommends shifting from one-size-fits-all firm-side drafting mandates toward consumer-side tools, including AI-powered smart readers that can simplify, translate, visualize, summarize, and interpret the same contract differently for different users. Such tools require standards for accuracy, bias, privacy, security, and liability. The larger lesson is methodological: good intentions and repeated citations do not substitute for validation of the empirical premises on which legal reform rests.

## Scope and cautions

The consumer sample is extensive but not exhaustive and overrepresents regulated credit-card and privacy documents. Social-media terms, retail purchases, and consumer-to-consumer agreements are not fully represented. Readability formulas remain imperfect even after robust averaging. The article therefore makes comparative claims about the observed corpus, not a claim that every consumer can understand every contract or that clear drafting has no value.

**Canonical citation:** Yonathan A. Arbel, *The Readability of Contracts: Big Data Analysis*, 21 J. Empirical Legal Stud. 927 (2024).
