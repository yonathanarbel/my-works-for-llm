use std::collections::HashMap;
use std::fs::File;
use std::io::prelude::*;
use rand::Rng;
use serde_json::Value;
use reqwest;
use tokio;

const ARTICLE_TEXT: &str = r#"

Electronic copy available at: https://ssrn.com/abstract=4809006

<<PAGE_BREAK>>

INTRODUCTION
Electronic copy available at: https://ssrn.com/abstract=4809006

<<PAGE_BREAK>>

𝑇𝑖𝑚𝑒 𝑎𝑛𝑑 𝐶𝑜𝑛𝑡𝑟𝑎𝑐𝑡 𝐼𝑛𝑡𝑒𝑟𝑝𝑟𝑒𝑡𝑎𝑡𝑖𝑜𝑛
Electronic copy available at: https://ssrn.com/abstract=4809006

<<PAGE_BREAK>>

I. TWO AND HALF THEORIES OF INTERPRETATION
Electronic copy available at: https://ssrn.com/abstract=4809006

<<PAGE_BREAK>>

𝑇𝑖𝑚𝑒 𝑎𝑛𝑑 𝐶𝑜𝑛𝑡𝑟𝑎𝑐𝑡 𝐼𝑛𝑡𝑒𝑟𝑝𝑟𝑒𝑡𝑎𝑡𝑖𝑜𝑛
Electronic copy available at: https://ssrn.com/abstract=4809006

<<PAGE_BREAK>>

*
Electronic copy available at: https://ssrn.com/abstract=4809006

<<PAGE_BREAK>>

𝑇𝑖𝑚𝑒 𝑎𝑛𝑑 𝐶𝑜𝑛𝑡𝑟𝑎𝑐𝑡 𝐼𝑛𝑡𝑒𝑟𝑝𝑟𝑒𝑡𝑎𝑡𝑖𝑜𝑛
II. CONTRACT INTERPRETATION: FINDING MEANING OR
PREDICTING INTENT?
Electronic copy available at: https://ssrn.com/abstract=4809006

<<PAGE_BREAK>>

“A sentence is never not in a context. We are never
not in a situation. . . . A set of interpretive assumptions is
always in force. A sentence that seems to need no
interpretation is already the product of one.” Fish
(1978)
[Textualists believe that t]he [judge can] . . . retir[e]
into that lawyer’s Paradise where all words have a fixed,
precisely ascertained meaning; where men may express
their purposes, not only with accuracy, but with fullness;
and where, if the writer has been careful, a lawyer,
having a document referred to him, may sit in his chair,
inspect the text, and answer all questions without raising
his eyes “(Thayer, 1898, at 428-29)
[I]t can hardly be insisted on too often or too
vigorously that language at its best is always a defective
Electronic copy available at: https://ssrn.com/abstract=4809006

<<PAGE_BREAK>>

𝑇𝑖𝑚𝑒 𝑎𝑛𝑑 𝐶𝑜𝑛𝑡𝑟𝑎𝑐𝑡 𝐼𝑛𝑡𝑒𝑟𝑝𝑟𝑒𝑡𝑎𝑡𝑖𝑜𝑛
and uncertain instrument, that words to not define
themselves, that terms and sentences in a contract, a
deed or a will do not apply themselves to external objects
and performances, that the meaning of such terms and
sentences consists of the ideas that they induce in the
mind of some individual person who uses or hears or
reads them, and that seldom in a litigated case do the
words of a contract convey one identical meaning to the
two contracting parties or to third persons. Therefore, it
is invariably necessary, before a court can give any
meaning to the words of a contract and can select one
meaning rather than other possible ones as the basis for
the determination of rights and other legal effects, that
extrinsic evidence shall be heard to make the court
aware of the “surrounding circumstances,” including
the other persons, objects, and events to which the words
can be applied and which caused the words to be used.
(emphasis added) (Corbin, 1960, at §535)
Electronic copy available at: https://ssrn.com/abstract=4809006

<<PAGE_BREAK>>

III. PRECISION AND ACCURACY
Electronic copy available at: https://ssrn.com/abstract=4809006

<<PAGE_BREAK>>

𝑇𝑖𝑚𝑒 𝑎𝑛𝑑 𝐶𝑜𝑛𝑡𝑟𝑎𝑐𝑡 𝐼𝑛𝑡𝑒𝑟𝑝𝑟𝑒𝑡𝑎𝑡𝑖𝑜𝑛
Electronic copy available at: https://ssrn.com/abstract=4809006

<<PAGE_BREAK>>

Electronic copy available at: https://ssrn.com/abstract=4809006

<<PAGE_BREAK>>

𝑇𝑖𝑚𝑒 𝑎𝑛𝑑 𝐶𝑜𝑛𝑡𝑟𝑎𝑐𝑡 𝐼𝑛𝑡𝑒𝑟𝑝𝑟𝑒𝑡𝑎𝑡𝑖𝑜𝑛
Figure 1 Contextualism vs. Textualism as Precision vs. Accuracy
Electronic copy available at: https://ssrn.com/abstract=4809006

<<PAGE_BREAK>>

Electronic copy available at: https://ssrn.com/abstract=4809006

<<PAGE_BREAK>>

𝑇𝑖𝑚𝑒 𝑎𝑛𝑑 𝐶𝑜𝑛𝑡𝑟𝑎𝑐𝑡 𝐼𝑛𝑡𝑒𝑟𝑝𝑟𝑒𝑡𝑎𝑡𝑖𝑜𝑛
IV. BIAS V. VARIANCE
2006
Electronic copy available at: https://ssrn.com/abstract=4809006

<<PAGE_BREAK>>

𝑀𝑆𝐸 = 𝐵𝑖𝑎𝑠2+𝑣𝑎𝑟𝑖𝑎𝑛𝑐𝑒+𝑖𝑟𝑟𝑒𝑑𝑢𝑐𝑖𝑏𝑙𝑒 𝑒𝑟𝑟𝑜𝑟𝑠
Electronic copy available at: https://ssrn.com/abstract=4809006

<<PAGE_BREAK>>

𝑇𝑖𝑚𝑒 𝑎𝑛𝑑 𝐶𝑜𝑛𝑡𝑟𝑎𝑐𝑡 𝐼𝑛𝑡𝑒𝑟𝑝𝑟𝑒𝑡𝑎𝑡𝑖𝑜𝑛
2
𝐵𝑖𝑎𝑠2 = (𝐸 (𝑓̂(𝑥;𝐷)−𝑓(𝑥))
𝐷
𝑉𝑎𝑟𝑖𝑎𝑛𝑐𝑒 = 𝐸 [(𝑓̂(𝑥;𝐷)−𝐸 [𝑓̂(𝑥;𝐷)]) 2 ]
𝐷 𝐷
𝑓̂
1992
Electronic copy available at: https://ssrn.com/abstract=4809006

<<PAGE_BREAK>>

Electronic copy available at: https://ssrn.com/abstract=4809006

<<PAGE_BREAK>>

𝑇𝑖𝑚𝑒 𝑎𝑛𝑑 𝐶𝑜𝑛𝑡𝑟𝑎𝑐𝑡 𝐼𝑛𝑡𝑒𝑟𝑝𝑟𝑒𝑡𝑎𝑡𝑖𝑜𝑛
Electronic copy available at: https://ssrn.com/abstract=4809006

<<PAGE_BREAK>>

V. INTERPRETATION VERSUS SIMULATION
Electronic copy available at: https://ssrn.com/abstract=4809006

<<PAGE_BREAK>>

𝑇𝑖𝑚𝑒 𝑎𝑛𝑑 𝐶𝑜𝑛𝑡𝑟𝑎𝑐𝑡 𝐼𝑛𝑡𝑒𝑟𝑝𝑟𝑒𝑡𝑎𝑡𝑖𝑜𝑛
Electronic copy available at: https://ssrn.com/abstract=4809006

<<PAGE_BREAK>>

Electronic copy available at: https://ssrn.com/abstract=4809006

<<PAGE_BREAK>>

𝑇𝑖𝑚𝑒 𝑎𝑛𝑑 𝐶𝑜𝑛𝑡𝑟𝑎𝑐𝑡 𝐼𝑛𝑡𝑒𝑟𝑝𝑟𝑒𝑡𝑎𝑡𝑖𝑜𝑛
“Should [Alice] be absolutely confident that [Bob]
prefers the new arrangement, it would not be a breach of
[Alice’s] promise for her to leave a message for [Bob]
simply informing [Bob] of the new plan. [Bob’s]’s actual
consent is not important where there is no uncertainty
about [Bob’s]’s understanding of her interests.”
“Because intimates know more about each other,
they can more reliably assess and act on a richer account
of each other’s evolving interests; to the extent this holds
true, they can adopt and continually update an ex-post
view.”
Electronic copy available at: https://ssrn.com/abstract=4809006

<<PAGE_BREAK>>

Electronic copy available at: https://ssrn.com/abstract=4809006

<<PAGE_BREAK>>

𝑇𝑖𝑚𝑒 𝑎𝑛𝑑 𝐶𝑜𝑛𝑡𝑟𝑎𝑐𝑡 𝐼𝑛𝑡𝑒𝑟𝑝𝑟𝑒𝑡𝑎𝑡𝑖𝑜𝑛
Electronic copy available at: https://ssrn.com/abstract=4809006

<<PAGE_BREAK>>

Electronic copy available at: https://ssrn.com/abstract=4809006

<<PAGE_BREAK>>

𝑇𝑖𝑚𝑒 𝑎𝑛𝑑 𝐶𝑜𝑛𝑡𝑟𝑎𝑐𝑡 𝐼𝑛𝑡𝑒𝑟𝑝𝑟𝑒𝑡𝑎𝑡𝑖𝑜𝑛
Electronic copy available at: https://ssrn.com/abstract=4809006

<<PAGE_BREAK>>

CONCLUSION
Electronic copy available at: https://ssrn.com/abstract=4809006

<<PAGE_BREAK>>

𝑇𝑖𝑚𝑒 𝑎𝑛𝑑 𝐶𝑜𝑛𝑡𝑟𝑎𝑐𝑡 𝐼𝑛𝑡𝑒𝑟𝑝𝑟𝑒𝑡𝑎𝑡𝑖𝑜𝑛
1. Arbel, Yonathan and David Hoffman. “Generative
Interpretation.” New York University Law Review
(Forthcoming, 2024).
2. Bagchi, Aditi. "Separating Contract and Promise." Florida
State University Law Review 38 (2011): 709-758.
3. Barnett, Randy E. "The Sound of Silence: Default Rules and
Contractual Consent." Virginia Law Review 78, no. 4 (1992):
821-911.
4. Bridgeman, Curtis. "Default Rules, Penalty Default Rules,
and New Formalism." Florida State University Law Review
33, no. 3 (2006): 683-720.
5. Corbin, Arthur Linton. Corbin on Contracts. 1st ed. St. Paul,
MN: West Publishing Co., 1960.
6. DiMatteo, Larry A. "Reason and Context: A Dual Track
Theory of Interpretation." Penn State Law Review 109, no. 2
(2004): 397-486.
7. Eisenberg, Melvin A. Foundational Principles of Contract
Law. New York: Oxford University Press, 2018.
8. Fish, Stanley. "Normal Circumstances, Literal Language,
Direct Speech Acts, the Ordinary, the Everyday, the Obvious,
What Goes without Saying, and Other Special Cases."
Critical Inquiry 4, no. 4 (1978): 625-644.
9. Geman, Donald, Edir Bienenstock, and René Doursat.
"Neural Networks and the Bias/Variance Dilemma." Neural
Computation 4, no. 1 (1992): 1-58.
10. Gilson, Ronald J., Charles F. Sabel, and Robert E. Scott.
"Text and Context: Contract Interpretation as Contract
Design." Cornell Law Review 100, no. 1 (2014): 23-98.
11. Goldberg, Victor P. "Impossibility and Related Excuses."
Journal of Institutional and Theoretical Economics (JITE) /
Zeitschrift Für Die Gesamte Staatswissenschaft 144, no. 1
(1988): 100-116.
12. Greene, Marjorie. "Theories of Interpretation in the Law of
Contracts." The University of Chicago Law Review 6, no. 3
(1939): 374-394.
13. Klass, Gregory. "Interpretation and Construction in Contract
Law." Georgetown Law Faculty Publications and Other
Works, 2018.
https://papers.ssrn.com/sol3/papers.cfm?abstract_id=2913228
.
Electronic copy available at: https://ssrn.com/abstract=4809006

<<PAGE_BREAK>>

14. Leech, Geoffrey, and Jenny Thomas. "Language, Meaning
and Context: Pragmatics." In An Encyclopedia of Language,
edited by N.E. Collinge, 173-206. London: Routledge, 1989.
15. Linzer, Peter. "The Comfort of Certainty: Plain Meaning and
the Parol Evidence Rule." Fordham Law Review 71, no. 3
(2002): 799-836.
16. Listokin, Yair. "The Meaning of Contractual Silence: A Field
Experiment." Journal of Legal Analysis 2, no. 2 (2010): 397-
416.
17. Maggs, Gregory E. "Karl Llewellyn's Fading Imprint on the
Jurisprudence of the Uniform Commercial Code." University
of Colorado Law Review 71, no. 2 (2000): 541-588.
18. Martinez, Eric, and Kevin Tobia. "What Do Law Professors
Believe About Law and the Legal Academy?" Georgetown
Law Journal 112 (forthcoming 2023): 1120189.
19. McElroy, R. G., and Glanville Williams. "The Coronation
Cases—I." Modern Law Review 4, no. 4 (April 1941): 241-
60. https://doi.org/10.1111/j.1468-2230.1940.tb00777.x.
20. Mitchell, Catherine. Interpretation of Contracts. 2nd ed.
London: Routledge-Cavendish, 2019.
21. Posner, Richard A. "The Law and Economics of Contract
Interpretation." Law & Economics Working Paper No. 229,
John M. Olin Program in Law and Economics, University of
Chicago Law School, 2004.
22. Saussure, Ferdinand de. Course in General Linguistics.
Edited by Charles Bally and Albert Sechehaye. Translated by
Wade Baskin. New York: Philosophical Library, 1959.
23. Schwartz, Alan, and Robert E. Scott. "Contract Interpretation
Redux." The Yale Law Journal 119, no. 5 (2010): 926-964.
24. Schwartz, Alan, and Robert E. Scott. "The Limits of Contract
Law." Yale Law Journal 113 (2003): 541, 573.
25. Scott, Robert E. "The Rise and Fall of Article 2." Louisiana
Law Review 62, no. 4 (2002): 1009-1064.
26. Scott, Robert E., and George G. Triantis. "Anticipating
Litigation in Contract Design." The Yale Law Journal 115,
no. 4 (2006): 814-879.
27. Shackel, Nicholas. "The Vacuity of Postmodernist
Methodology." Metaphilosophy 36, no. 3 (2005): 295-320.
28. Shalev-Shwartz, Shai, and Shai Ben-David. Understanding
Machine Learning: From Theory to Algorithms. New York:
Cambridge University Press, 2014.
Electronic copy available at: https://ssrn.com/abstract=4809006

<<PAGE_BREAK>>

𝑇𝑖𝑚𝑒 𝑎𝑛𝑑 𝐶𝑜𝑛𝑡𝑟𝑎𝑐𝑡 𝐼𝑛𝑡𝑒𝑟𝑝𝑟𝑒𝑡𝑎𝑡𝑖𝑜𝑛
29. Silverstein, Joshua M. "The Contract Interpretation Policy
Debate: A Primer." Stanford Journal of Law, Business &
Finance 26, no. 2 (2021): 222-294.
30. Stempel, Jeffrey W., and Erik S. Knutsen. "Rejecting Word
Worship: An Integrative Approach to Judicial Construction
of Insurance Policies." University of Cincinnati Law Review
90, no. 2 (2021): 561-636.
31. Thayer, James Bradley. A Preliminary Treatise on Evidence
at the Common Law. Boston: Little, Brown, and Company,
1898.
32. Wellman, Henry M. "Developing a Theory of Mind." In The
Wiley-Blackwell Handbook of Childhood Cognitive
Development, edited by Usha Goswami, 258-284. 2nd ed.
Malden, MA: Wiley-Blackwell, 2011.
Electronic copy available at: https://ssrn.com/abstract=4809006"#;

fn main() {
    println\!("{}", ARTICLE_TEXT);
}