/*
Shielding of Assets and Lending Contracts (ssrn-2820650) — corpus code wrapper

This file intentionally embeds the paper text and study assets in code form.
It helps code-centric ingestion pipelines and makes the corpus easy to load programmatically.
*/

const PAPER_ID = "ssrn-2820650";
const TITLE = "Shielding of Assets and Lending Contracts";
const SSRN_URL = "https://papers.ssrn.com/sol3/papers.cfm?abstract_id=2820650";
const YEAR = 2016;
const AUTHORS = ["Yonathan Arbel"];
const KEYWORDS = ["contracts", "AI", "law"];

const SUMMARY_MD = `
Okay, here's the bullet list based on the provided text:

*   **Professor Yonathan Arbel of the University of Alabama School of Law writes that a debtor's wealth significantly influences their decision to shield assets, as wealthier individuals would need to shield a larger volume of assets at greater cost for it to be effective against creditors. Professor Yonathan Arbel of the University of Alabama School of Law writes that his paper develops a theory of asset shielding to explain this behavior, arguing that richer debtors often find it irrational to shield, while poorer debtors pose a higher shielding risk which impacts credit markets.**

1.  ## TL;DR ≤100 words
    Professor Yonathan Arbel of the University of Alabama School of Law argues that debtor wealth dictates asset shielding decisions. His theory posits that wealthier debtors often find shielding large asset volumes too costly and thus irrational. Conversely, poorer debtors present a higher shielding risk. This dynamic, where shielding is more rational for poorer debtors, significantly influences credit markets.

2.  ## Section Summaries ≤120 words each
    Professor Yonathan Arbel of the University of Alabama School of Law writes that a debtor's wealth is a key determinant in their choice to shield assets from creditors. He notes that for wealthier individuals, the sheer volume of assets requiring protection, and the associated costs, can make shielding an ineffective or irrational strategy. Professor Yonathan Arbel of the University of Alabama School of Law writes that his paper introduces a theory of asset shielding to elucidate these behaviors. This theory suggests that while richer debtors may forego shielding, poorer debtors are more likely to engage in it, thereby creating a higher shielding risk that has repercussions for credit markets.
`;

const SUMMARY_ZH_MD = `
好的，这是基于您提供的英文文本翻译的正式中文摘要：

*   **阿拉巴马大学法学院的约纳坦·阿尔伯（Yonathan Arbel）教授写道，债务人的财富状况显著影响其隐匿资产的决策，因为较富裕的个人若要有效对抗债权人，需以更高成本隐匿更大规模的资产。阿拉巴马大学法学院的约纳坦·阿尔伯教授在其论文中提出了一种资产隐匿理论来解释此行为，他认为，富裕债务人通常认为隐匿资产不尽合理，而贫困债务人则构成更高的资产隐匿风险，进而对信贷市场产生影响。**

1.  ## 内容摘要（不超过100字）
    阿拉巴马大学法学院的约纳坦·阿尔伯教授认为，债务人财富状况决定其资产隐匿决策。其理论阐明，富裕债务人常因隐匿大量资产成本过高而认为此举不理性；相反，贫困债务人则构成更高的资产隐匿风险。这种贫困债务人更倾向于选择资产隐匿的动态，对信贷市场具有显著影响。

2.  ## 各节摘要（每节不超过120字）
    阿拉巴马大学法学院的约纳坦·阿尔伯教授指出，债务人的财富是其决定是否向债权人隐匿资产的关键因素。他提到，对较富裕的个人而言，需要保护的资产规模庞大及相关高昂成本，可能使资产隐匿成为一种无效或不理性的策略。阿拉巴马大学法学院的约纳坦·阿尔伯教授在其论文中引入了一种资产隐匿理论以阐释这些行为模式。该理论认为，富裕债务人可能放弃隐匿资产，而贫困债务人则更倾向于这样做，由此带来更高的资产隐匿风险，并对信贷市场产生深远影响。
`;

const ONE_PAGER_MD = `
# Shielding of Assets and Lending Contracts — one-page summary

**Paper ID:** \`ssrn-2820650\`
**Year:** 2016
**Author(s):** Yonathan Arbel
**SSRN:** https://papers.ssrn.com/sol3/papers.cfm?abstract_id=2820650

## TL;DR

Professor Yonathan Arbel of the University of Alabama School of Law argues that debtor wealth dictates asset shielding decisions. His theory posits that wealthier debtors often find shielding large asset volumes too costly and thus irrational. Conversely, poorer debtors present a higher shielding risk. This dynamic, where shielding is more rational for poorer debtors, significantly influences credit markets.

## Keywords

contracts; AI; law

## Files

- Full text: \`papers/ssrn-2820650/paper.txt\`
- PDF: \`papers/ssrn-2820650/paper.pdf\`
- Summary (EN): \`papers/ssrn-2820650/summary.md\`
- Summary (ZH): \`papers/ssrn-2820650/summary.zh.md\`

_Auto-generated study aid. For canonical content, rely on \`paper.txt\`/\`paper.pdf\`._
`;

const STUDY_PACK_MD = `
# Study pack: Shielding of Assets and Lending Contracts (ssrn-2820650)

- SSRN: https://papers.ssrn.com/sol3/papers.cfm?abstract_id=2820650
- Full text: \`papers/ssrn-2820650/paper.txt\`
- Summary (EN): \`papers/ssrn-2820650/summary.md\`
- Summary (ZH): \`papers/ssrn-2820650/summary.zh.md\`

## Elevator pitch

Professor Yonathan Arbel of the University of Alabama School of Law argues that debtor wealth dictates asset shielding decisions. His theory posits that wealthier debtors often find shielding large asset volumes too costly and thus irrational. Conversely, poorer debtors present a higher shielding risk. This dynamic, where shielding is more rational for poorer debtors, significantly influences credit markets.

## Keywords / concepts

contracts; AI; law

## Suggested questions (for RAG / study)

- What is the paper’s main claim and what problem does it solve?
- What method/data does it use (if any), and what are the main results?
- What assumptions are doing the most work?
- What are the limitations or failure modes the author flags?
- How does this connect to the author’s other papers in this corpus?

_Auto-generated study aid. For canonical content, rely on \`paper.txt\`/\`paper.pdf\`._
`;

const ARTICLE_TEXT = `
Shielding of Assets and Lending Contracts
(Forthcoming, International Review of Law & Economics)
Yonathan A. Arbel*
ABSTRACT
The primary means of enforcement of legal liabilities is through the seizure of debtors’ assets.
However, debtors can shield their assets in various ways and thereby reduce the power of en-
forcement. This paper studies the circumstances under which a debtor would choose to shield as-
sets and the value of assets that would be shielded.
A key idea is that borrower’s wealth mutes shielding incentives. Intuitively, avoiding debts
through shielding requires that enough assets will be shielded, for else the debts can be collected
from exposed assets. A wealthier debtor would thus need to shield more assets, and at a greater
cost, than a debtor with limited wealth. Using this basic understanding, I develop a theory of asset
shielding and explore its implications for incomplete lending contracts, explaining the role of eq-
uity agreements, equity cushions and collateral, and debt forgiveness, and explore the some of the
policy implications.
1
Electronic copy available at: https://ssrn.com/abstract=2820650
1. INTRODUCTION
The primary means of enforcement of civil legal liabilities, such as debt contracts, taxes, or tort
judgments, is through the seizure of debtors’ assets. However, as Section 2 discusses, debtors are
often in a position to circumvent asset seizure by the use of such methods as hiding cash, transfer-
ring ownership of property to family members, and using shell corporations. After shielding,
creditors will be dissuaded from collecting their debts and debtors may file for bankruptcy and
discharge their obligations.
Despite the importance of shielding decisions to the modern economy, the literature has not pro-
vided an account of when shielding will take place and the magnitude of assets that would be
shielded. This paper develops a theory of asset shielding that explains shielding behavior and its
impact on the credit market. It argues that richer debtors would often not find it in their self-
interest to shield assets. Conversely, poorer debtors, even if formally solvent, pose a shielding
risk. This risk would tend to harm debtors by limiting their access to credit and therefore they
would benefit from being able to commit not to shield. Such a commitment is difficult to secure
and the paper explores several public and private strategies of limiting shielding risk.
The model, described in Section 3, is a stylized lending model with incomplete contracts where
an entrepreneur borrows money from a lender for an investment and then faces an opportunity to
shield the investment returns at some per-unit cost. The focus on incomplete contracts is relevant
to high transaction costs environments where more sophisticated contracts are unavailable. An
important benefit of this assumption is that it allows us to generalize much of the analysis to in-
clude non-contractual settings, such as accidents, fines, and taxes. The assumption that there is
sufficient opportunity to shield assets ex-post (due to Hart and Moore 1989) is motivated by the
relative simplicity and rapidness of various shielding techniques and the costliness of monitoring
borrowers.
We start by examining the question how much value a borrower who decided to shield assets
would choose to shield, because this will inform all other decisions. In answering it, it is useful
to dispel first two common perceptions. Firstly, the borrower may be thought to attempt to shield
the amount he owes. On reflection, however, such a decision is revealed to be irrational. If the
borrower only shields what he owes, than he will often leave other assets exposed. The (oft-
overlooked) recourse principle (Scott 1913) holds that a debt-holder has a right to substitute his
debt with those other assets. The lender will then be able to collect despite shielding effort, ren-
dering shielding pointless. Secondly, the borrower may be thought to choose an amount to shield
such that the marginal cost of shielding does not exceed the marginal benefit, understood as the
amount shielded after deducting shielding expenses. But this too is erroneous: It assumes the
costs are borne by the borrower, but they are in fact taken from what is owed to the lender.
With this in mind, the first contribution of this paper is in explaining that the choice of value to
shield depends on debtor’s wealth and the size of the debt—but not on shielding costs. More spe-
cifically, the lowest value that a debtor will find rational to spend on shielding is given by the
2
Electronic copy available at: https://ssrn.com/abstract=2820650
difference between debtor’s wealth and debt.1 To illustrate the intuition underlying this result,
suppose the borrower owes $10,000 and has $25,000. Spending only $12,000 on shielding, for
example, would leave more than $10,000 exposed, so the lender would be able to collect in full
despite shielding effort. It follows that if the borrower spends on shielding any positive amount, it
must be greater than $15,000 for else shielding effort would be moot—and this minimal value
that must be shielded holds regardless of shielding costs.
Understanding how much would be shielded helps to determine whether shielding will take place
at all. When deciding whether to shield, the borrower compares the costs involved in shielding
sufficient value to the cost of debt repayment. If shielding costs exceed the debt, the borrower
would be better off repaying the debt than shielding. Now, since a wealthier borrower would
have to shield more, he would face higher shielding costs. From this follows the second contribu-
tion of this paper: the wealthier the borrower is, the less rational shielding becomes.2
Anticipating the ex-post shielding decisions has important effects on project finance. On the one
hand, shielding risk can lead to outright credit denial even for profitable investments. While nor-
mally the lender can be compensated for default risk by charging a higher interest, this is not nec-
essarily the case here. This is because charging higher interest would result in greater shielding
incentives may thus reduce lender’s payoff. This will result in credit rationing for certain profita-
ble investments even in the absence of the conventional reasons of adverse selection and moral
hazard in effort (Stiglitz and Weiss, 1981).
On the other hand, the analysis also implies a strong limit on the incentive to shield. If borrower’s
ex-post wealth is high, shielding will be too costly. And if the investment has high expected re-
turns, the lender need not worry about default due to shielding. As a result, high yield invest-
ments, even if risky in terms of variability of returns (e.g., start-ups), can enjoy easy access to
credit even without collateral. Stated differently, a lender might perversely prefer a risky invest-
ment to a safer one with the same expected value. This is because the upside in the risky invest-
ment is more likely to surpass the wealth threshold.
Given the problems asset shielding creates, both the lender and the borrower will have an incen-
tive to limit shielding opportunities and the analysis compares two alternatives. In the one, the
borrower “ties his hands” and commits not to shield. Such a commitment is of limited credibility,
because contractual sanctions are pecuniary and so will not be effective in exactly those circum-
stances when the borrower breaches the contract. Still, it may be possible for the borrower to limit
future shielding by permitting the creditor to repossess assets soon after default or by placing fu-
ture revenues in a hard-to-access account. In the other alternative, the parties opt for an equity
agreement instead of a debt contract. I show that the first approach mitigates the problem of
1 In the body of the analysis I develop a stronger version of this conclusion, showing that it would pay the borrower to
shield all assets conditional on deciding to shield.
2 The intuition underlying this result can be easily illustrated. If a borrower who owes $10,000 has $15,000 in assets,
then the minimum he would have to spend on shielding is $5,000. If he has $80,000, he would have to spend at least
$70,000. The greater spending implies a greater cost and lesser incentive to shield.
3
Electronic copy available at: https://ssrn.com/abstract=2820650
shielding but will not solve it, whereas the second approach may indeed solve the problem of as-
set shielding but may not always be practical.
Section 4 analyzes a few extensions of the basic model. I explain the role of collateral—
possessory and nonpossessory, explicit and implicit—in mitigating shielding risk and how equity
cushions may be useful. I suggest another motive for debt relief agreements—avoiding shielding
behavior—and suggest the structure of such agreements. I then consider how collection costs
could affect shielding behavior and show that there is both substitution and complementarity be-
tween collection costs and shielding behavior. Finally I discuss ex-ante shielding and consumer
lending.
I close with a brief analysis of a few regulatory implications. To ward off shielding, the legal sys-
tem could attack the problem directly—by making it more expensive to shield through sanctions
on shielding and closing shielding loopholes—or it can do so indirectly, by requiring minimal
asset requirements for actors engaged in potentially dangerous activities or by limiting the fines
and judgments imposed on asset-constrained individuals. Section 5 concludes.
There is rich work dating back to Shavell (1986) that studies the effect of insolvency on injurers’
incentives to take care (the ‘judgment proof problem’). For example, Ganuza and Gomez (2008)
advocate the use of ‘soft’ liability standards for insolvent injurers. For the most part, work in this
area takes wealth levels as exogenously set (Summers 1983, Dari-Mattiacci and De Geest 2002,
Ganuza and Gomez, 2008, Dari-Mattiacci and Mangan 2008, Wickelgren 2011) although some
important work considers the possibility that firm capitalization may be the result of strategic be-
havior (Shavell 2005, Che and Spier 2008, Veld and Hutchinson 2009, Ganuza and Gomez 2011).
However, even the latter strain in the literature only studies ex-ante moral hazard.3 This paper
complements this body of work by considering the ex-post moral hazard inherent in asset shield-
ing, which is of relevance regardless of ex-ante capitalization.
Other related work comes from the literature on credit. However, this scholarship too has largely
abstracted away from shielding decisions and focused instead on the costs and choice of default
(e.g., Leff 1970, Schwartz 1983, Gross and Souleles 2002, White 2007). In the recent contribu-
tion of Ellingsen and Kristiansen, they allow borrowers to “divert” assets, but this is different
from shielding as it is assumed to come at no cost besides the risk of enforcement (Ellingsen and
Kristiansen 2011). In practice, however, shielding is quite costly and generally goes unpunished.4
The part of the literature that relates most closely to asset shielding is the theory of costly state
falsification (owing to Lacker and Weinberg 1989). This work studies contracting in the face of
an ex-post moral hazard of asset hiding, similar to the inquiry at hand. However, this literature
focuses on sophisticated contracts and does not consider simpler contracts which are common in
practice (as noted by Lacker himself 1991). Moreover in many settings where shielding opportu-
nities are present—e.g., torts, civil judgments, fines, and taxes—contracting is impossible, so that
the study of shielding behavior with incomplete contracts becomes important.
3 Shavell (2005) adumbrates this possibility of ex-post moral hazard but does not analyze it.
4 See discussion in section 2.
4
Electronic copy available at: https://ssrn.com/abstract=2820650
The primary contributions of this paper relative to existing literature stem from the analysis of
optimal asset shielding decisions and the identification of the relationship between these decisions
and borrower’s wealth. The implications of this analysis on the credit market and on rich and
poor borrowers are likewise new. Similarly novel are the ideas that shielding is generally an all-
or-nothing proposition,5 that high wealth mutes the incentive to shield, and the interpretation of
the incentive to enter into debt relief contracts.
2. SHIELDING ASSETS FROM ENFORCEMENT
The term asset shielding (or equivalently, asset protection) is used here expansively, to account
for any action or omission by an individual that takes place after the creation of a specific debt
6
and that is intended to limit the lender’s ability to seize assets in case of default. Thus defined,
some examples are the case of Charles Kallestad, a borrower who was found to be shielding as-
sets worth hundreds of thousands of dollars through multiple transfers to an accomplice (U.S. v.
Kallestad 51 F.3d 1044 and trial documents); O.J. Simpson and Paul Bilzerian, who moved to
multi-million mansions in Florida after accruing millions of dollars in debts; and Jimmy Jen, a
debtor from California who was found to have shielded over $6 million using secret vaults, shell
corporations, and a sham divorce (Elinson 2010).7
The primary forms of shielding are consumption (e.g., purchasing perishable goods or expensive
dinners, giving to charity), concealment (e.g., hiding cash, mock transfers, failing to disclose in-
come), legal asset protection (e.g., investing in bankruptcy exempt property, creating special
trusts, and forming corporate structures), or obstruction of enforcement efforts (e.g., transferring
assets to foreign havens or changing addresses). The literature abounds with techniques and cas-
es (e.g., Che and Spier 2008, White 2007, Vandervort v. Vandervort, 134 P.3d 892, (Okla. Civ.
App. 2005)). Many of these techniques—although by no means all—are simple and rapid to af-
fect and do not require extensive advance planning, most notably asset concealment. This means
that a debtor would often be in a position to shield assets before the creditor seeks repayment.
This is especially true in loan agreements where the moment of realization of investment returns
is not precisely known in advance.
Shielding is costly. Some costs are direct—such as setting a shelter or retaining a lawyer—but
others are indirect—such as the opportunity cost of shielding and the cost involved in not being
able to freely use one’s assets. Whether the marginal cost of shielding increases or declines with
the amount shielded is hard to tell a-priori, but it will clearly be positive in most cases.
5 Ellingsen and Kristiansen (2011) reach a similar result but for a different reason.
6 Hence, I do not consider “preparatory” asset protection that takes place before the borrower has a specific debt.
7 O.J. Simpson escaped a $33.5 million judgment while residing in a Florida mansion and receiving a $25K monthly
pension, taking advantage of Florida’s generous homestead protection laws and the federal retirement laws (Alper
2007); Bilzerian defaulted on $140 million debts while residing in an 11 bedroom home in Florida taking advantage of
similar protections (Shenon, 2001); Jimmy Jen hid over $6 million dollars in vaults, shell corporations, and in his
wife’s possession following a sham divorce (Elinson 2010);
5
Electronic copy available at: https://ssrn.com/abstract=2820650
There are also two other types of costs: legality and reputation (Arbel, 2015). As a general matter,
shielding assets with the intention of avoiding a specific debt exposes the debtor to both civil and
criminal liability. Civil liability is seemingly somewhat effective,8 although the bringing of such a
lawsuit requires evidence of fraudulent transfer, which can be hard to obtain, and the successful
implementation of the judgment depends on the assets still being in jurisdiction.9 Moreover, even
preliminary steps, such as locating the debtor, prove difficult in practice (Stephen, Avril, and
Wrapson 2013). Criminal sanctions are seemingly much less effective (e.g., McCullough 1997).
In 2011 there were 1.3 million bankruptcy filings (U.S. Courts 2012), but the Executive Office of
10
the US Trustees (EOUST) referred only 1,968 cases to criminal proceedings (DOJ 2012) and in
only about 20 cases formal charges were filed.11 This is despite evidence noted below by the
EOUST itself that bankruptcy fraud is pervasive.
Another cost is reputation (e.g., loss of credit score or lending relationships). Reputation is like-
wise limited because it is often hard to observe whether the borrower defaulted strategically or
due to real financial inability. The multiplicity of credit channels further dilutes reputational ef-
fects, as information, even if obtained, may not fully propagate to all other lenders.
Overall, shielding is somewhat prevalent and seemingly highly effective. While we do not have
good data, one indication comes from reports of the EOUST, which is the agency in charge of
investigating bankruptcy fraud. In a random sample of 102 cases, they found that 17% of the
12
debtors “materially misstated” the true extent of their assets. This evidence joins other studies
that find strategic behavior in bankruptcy and in judgment enforcement (Fay, Hurst and White
2002 and Ning Zhu 2011, but see Sullivan, Westbrook and Warren 1989; Arbel, 2015), suggest-
13
ing that some debtors use bankruptcy to protect assets. The size of the debt collection industry
is also indicative, as their main service is collection from reluctant debtors. This industry collects
14
over 50 billion dollars annually (PWC 2008, EY 2012).
8 Creditor’s right to reverse fraudulent transfers stems, mainly, from either § 9 of the UFTA or §548 of the Bankruptcy
Code. Both Westlaw and LexisNexis record only about 450 cases in 2011 that cite to § 548.
9 In FTC v. Affordable Media, LLC 179 F. 3d 1228 (9th Cir. 1999), the debtors placed their assets in an offshore ac-
count. When ordered by a US judge to repatriate their assets, they faxed a request to the offshore trustee who refused to
implement it, because the trust mandate gave him the power to refuse requests made under duress.
10 The EOUST collects information from trustees who are required to inform the EOUST of all suspicions of criminal
activity (Handbook 2012)
11 Similarly, the database TracFed shows that over the last decade, there was a yearly average of about 60 cases prose-
cuted where the main charge was 8 USC §152 (the main section concerning asset fraud in bankruptcy). For the same
criteria, Westlaw records 56 open dockets in 2012 and the IRS reports opening investigations in only 44 cases in 2014
(http://www.irs.gov/uac/Statistical-Data-Bankruptcy-Fraud).
12 Of the cases referred, the causes for referral was: false oath or statement (33.2%), concealment of assets (24.8%), and
other bankruptcy fraud schemes (21.5%) (each case may have more than one allegation).
13 There are also reports of trillions of dollars hidden in offshore accounts, but it is unclear which part is motivated by
asset protection and which by the desire to avoid taxes and other laws.
14 Analysis of the Federal Reserve reports show that in the 10 years between 2003 to 2013, commercial banks have
reported average charge-offs of 5.5% (credit cards) and 0.91% (residential real estate) (Federal Reserve Website,
http://www.federalreserve.gov/releases/chargeoff/chgallnsa.htm) Similarly, in total, corporations write off an average
6
Electronic copy available at: https://ssrn.com/abstract=2820650
In sum, then, from the debtor’s perspective, the shielding of assets is a viable strategy with low
chance of criminal sanctions, but with some cost due to limited reputational effects and some risk
of reversal of the shielding if it is not executed carefully.
3. MODEL AND ANALYSIS
The asset shielding model spans four dates (see Figure 1) and involves two risk-neutral parties. At
Date 1, one of the parties (‘borrower’) has a positive-expected value investment that requires a
fixed investment of b. The borrower has no wealth,15 and seeks a loan from the other party (‘lend-
er’), in a competitive lending market with costs of capital normalized to zero. The parties negoti-
ate the interest on the loan r, so that the amount due at Date 4 is b + r. 16 If the lender agrees to
provide the loan at this price, the funds are invested.
If the investment is made, the earnings, e, are realized at Date 2 and are taken from the probabil-
ity distribution function f(e), and e is in [e̲, ē], with 0< e̲ <ē.
At date 3, the borrower decides whether to repay or shield. This is given by the amount t (t≥0)
the borrower spends on shielding, with t=0 being no-shielding. There is a shielding ‘technology’,
described by the function s(.), which gives the amount shielded by spending t. It is assumed that s'
is strictly increasing, continuous, and with s'<1.17
The borrower’s choice at Date 3 of the amount to shield implicitly defines the amount he leaves
exposed. At Date 4 the lender collects and by the recourse principle, the lender may collect from
all exposed assets up to the amount owed, i.e., Min(e-t, b+r). After that, the debt is discharged in
bankruptcy or the creditor gives up on collecting.
of 4.8% of their accounts receivable, which in 2009 totaled $253 billion (IRS Statistics, http://www.irs.gov/uac/Tax-
Stats-2). Note that this figure includes loss from debtors who are genuinely unable to repay, so it should only be under-
stood as a rough estimate on the upper limit of strategic debt avoidance.
15 This assumption bars the use of collateral and it is relaxed in the extensions of the model.
16 I restrict attention here to simple loan contracts, where the parties set a non-contingent payment schedule. I also ab-
stract from criminal sanctions, as Section 2 suggests they are rare; civil sanctions that reverse shielding are more com-
mon and they can be conceptualized as part of the expected cost of shielding.
17 This assumption captures the idea that there is always some cost to shielding, such as the opportunity cost involved
in losing interest on one’s savings, for otherwise the assets would be shielded in the baseline and the inquiry would be
trivial. I further assume here that there are no fixed shielding costs. I relax this assumption below.
7
Electronic copy available at: https://ssrn.com/abstract=2820650
Figure 1: Timing of the Model
Finally, the measure of social welfare adopted is the total utility of the parties, which in light of
the risk neutrality assumption, is simplified to their joint wealth. At the end of the Date 4, the so-
cial benefits from the investment, if made, are its earning e, and the social costs are the costs of
making the investment b as well as shielding costs c(s).
Timing, Information, and Order of Analysis: the game is solved by backwards induction. It is
assumed that all information is common knowledge, although shielding may not be verifiable to a
third party. The analysis starts with the shielding and collection stage (Dates 3 & 4), when the
amount of earning e and the interest r are fixed. The analysis then goes backwards to analyze the
parties’ decisions in the first stage (Date 1) when the investment contract, and more specifically r,
is negotiated. The results of the analysis will be compared to both the social optimum and to a
situation where shielding is limited and costly—either by law, technology, or borrower’s com-
mitment.
3.1. Stage 2 Analysis: Shielding
At stage 2, the borrower’s objective is to maximize profits by choosing t, the amount to spend on
shielding. This makes borrower’s payoff:
𝑠(𝑡)+max(0, 𝑒−𝑡−(𝑏+𝑟)) (1)
The first term here is the value to the borrower of shielded assets. By the recourse principle, all
exposed assets are potential substitutes for the debt. This means that after shielding, the lender
will collect the debt from exposed assets, leaving the borrower with what remains (with a lower
bound of zero). We then have the following result.
Proposition 1
1.1. If the borrower chooses to shield any positive amount of wealth, the optimal shielding sum,
t*, will be equal to his entire wealth; t*=e.
1.2. There exists a unique critical level of wealth e* that if e < e*, all of e is used in shielding, so
no debt is collected, and if e > e*, nothing is shielded. However, e* is an increasing function of r
and may exceed e̅.
Proof: See Appendix.
8
Electronic copy available at: https://ssrn.com/abstract=2820650
The intuition behind this proposition is as follows. The borrower has at this stage a debt of b+r
and a fixed wealth of e. We are assuming the borrower decided to shield assets and we are look-
ing to see how much value he would like to shield. This decision can be thought of as involving
two steps. First, the borrower contemplates whether spending on shielding a small amount is de-
sirable. If that amount is small indeed, the remaining assets would exceed the debt: i.e., e-t>b+r.
On reflection, shielding such a small amount is undesirable, for it leaves exposed enough value so
that despite shielding efforts, the debt will be collected in full. This is a direct implication of the
recourse principal. Understanding that, the borrower might seek to shield some larger amount, but
perhaps something that still falls short of his entire wealth (i.e, e-(b+r)<t<e). This would have
the benefit of avoiding some of the debt, but it will also—by definition—leave some assets ex-
posed. The lender will collect these assets which will satisfy only part of the debt. But with re-
spect to these exposed assets, the borrower reasons that leaving them exposed would mean that
they will be lost; if, instead, the borrower would shield them as well, he will retain some of their
value, and however small that is, it is still better than losing them to the lender. Hence, shielding
these assets too is in the borrower’s interests. We conclude that if shielding takes place, it would
be rational to shield all of the borrower’s assets.
Example 1. The borrower owes $6,000 and the project earned $10,000. For every dollar
spent on shielding, the borrower retains 20 cents. At first, the borrower contemplates spending
$3,000 on shielding (securing him a value of 0.20*3,000=$600). But since that would leave ex-
posed $7,000, the lender will be able to collect the debt in full. Then the borrower considers
spending $7,000 on shielding (securing him a value of 0.2*7,000=$1,400). This is preferable, but
still leaves $3,000 that will be collected. If the borrower shields these $3,000 too, he can further
increase his payoff by 0.2*3,000=$600, which is clearly superior to leaving them exposed. As a
result, if the borrower chooses to shield assets, he will spend the full $10,000 on shielding.
It is important to note that this conclusion is general. Even if shielding is very costly, it will still
be the case that if shielding takes place, the borrower’s entire wealth would be spent on shielding,
i.e., t=e. But of course, this does not mean that shielding will always take place. With this in
mind we turn to the other part of Proposition 1.
At the beginning of Stage 2 the borrower may either shield or repay, depending on his payoff
from each option. Repayment costs b+r while shielding will cost, based on the previous analysis,
e-s(e). There is some positive marginal cost to shielding due, for example, to the loss of interest
or asset mobility. This implies that when the amount spent on shielding is large enough, call it e*,
we will have e*-s*(e)=b+r; i.e., it will cost the same amount to shield one’s assets as it would
cost to repay the debt. In other words, e* is a wealth threshold; having more wealth than e* means
that it would cost more to shield than to repay, so that it will no longer be rational to shield assets.
We note three implications of the wealth threshold: shielding assets can be irrational if borrower’s
wealth is sufficiently high, for then shielding costs may exceed the cost of repayment.18 Addition-
18 Of course, the project may not produce sufficient earnings to pass this threshold (i.e., e*< ē), which can be very high
if it easy to shield. In this case, we shall see, the project will not be financed.
9
Electronic copy available at: https://ssrn.com/abstract=2820650
ally, e* is increasing in r –that is, a higher debt increases the incentive to shield assets. Lastly, if
shielding technology is inefficient and costly, due to legal or market conditions, that would have
the effect of increasing e-s(e), implying a lower wealth threshold. Taken together, the incentive to
shield increases in debt but decreases in wealth and enforcement.
Example 2. Consider the same borrower as before. If the project earns $7,000, spending
this entire amount on shielding would cost 0.8*7,000=$5,600, relative to a debt repayment cost of
$6,000. Hence, shielding would be preferable. But if the project earns $10,000, shielding would
cost $8,000, making it more costly than repaying $6,000. In fact, for every amount above $7,500,
it would cost the borrower more to shield than to repay—making it the wealth threshold. In other
words, if the return on the investment exceeds $7,500, the borrower would favor repayment to
shielding. The wealth threshold would change if the shielding technology were to become less
efficient, such that shielding would only leave 10 cents on the dollar (instead of 20). In that case,
the wealth threshold would fall to only $6,666.
It is of interest to note the relationship between the wealth threshold and the concept of solvency.
The borrower is technically solvent if e>b+r. However, this is not enough to surpass the wealth
threshold; that will happen only if e-s(e)>b+r, or e>b+r+s(e). Hence, borrower’s wealth must be
greater than the concept of solvency would imply by s(e) for repayment to take place.19 In other
words, solvency is a necessary but insufficient condition for debt repayment. The flip side of this
conclusion is that if the borrower chooses to repay, this will only happen when he can repay in
full – so that the lender can expect either no payment or full repayment.
The following figure illustrates these points:
19 This follows from Proposition 1.1: If borrower’s wealth falls below b+r, his entire wealth will be taken by the credi-
tor, so even if shielding is very wasteful, it is still better to shield and retain some value than lose all value.
10
Electronic copy available at: https://ssrn.com/abstract=2820650
Figure 2: Borrower’s Payoff under Shielding or Repayment as a Function of Earnings
As the figure shows, as the investment earnings rise, the borrower finds shielding more and more
costly. When earnings pass the e* threshold, the shielding costs exceed the costs of repaying the
debt and the borrower finds repayment optimal. As noted, having b+r in earning is insufficient to
warrant repayment, and the earnings must be greater than b+r by s(e*) for the borrower to repay.
Lastly, from the creditor’s perspective, repayment is binary due to the cliff’s edge nature of
shielding: For every e≤e*, the lender is not paid at all, and for every e>e* the creditor is repaid in
full.
Social Optimum. The social optimum is defined as the sum of the parties’ joint wealth:
𝑒−𝑏−(𝑡−𝑠(𝑡)) (2)
That is, the returns from the investment less the costs of making the investment and the costs in-
20
volved in shielding. Since (2) is strictly decreasing in t, it will be socially desirable that t=0, so
that no assets will be shielded; the simple reason is that asset protection absorbs resources but
creates no value aside from shifting wealth between the parties.
Private Decisions with Limits on Shielding. In some situations, shielding may be less effective
and more costly—either because the legal system makes shielding difficult or because the bor-
rower was able to “tie his hands” and make it more difficult to shield. To capture that, we will
consider a competing shielding technology s (), such that s '(t)<s(t) for all t, with full inability to
c c
shield being s '(t)=0 for all t. Under limited shielding, borrower’s payoff is given by:
c
𝑠 (𝑡)+max(0, 𝑒−𝑡−(𝑏+𝑟)) (3)
𝑐
And since s '(t)<s (t) for all t, it follows that the return on shielding will be lower, thus limiting
c c
the desirability of ex-post shielding. This is a reflection of the earlier point that if shielding tech-
20 The interest r does not factor in the parties’ joint wealth since it is a payment made between the parties.
11
Electronic copy available at: https://ssrn.com/abstract=2820650
nology is less effective, the wealth threshold decreases. In the extreme, when shielding is una-
vailable (i.e., s'(t)=0 for all t), the wealth threshold becomes zero, leading the borrower to never
shield, thus aligning his behavior with what the social optimum.
From the lender’s perspective, limits on shielding lower the wealth threshold implying a higher
probability of repayment. In the extreme, when the wealth threshold is zero, the lender receives a
payment of min(e, b+r) (compared to b+r only for e>e*).
Equity Agreement. We compare now the simple debt contract with a simple equity agreement
where the lender buys a stake in the borrower. Under this agreement, the borrower owes the lend-
er a fraction f of all earnings. To simplify the discussion of equity agreements, we will make the
fairly strong assumption that t-s(t)=kt for all t¸ that is, that the marginal cost of shielding is fixed
at k.
Proposition 2.
2. If the equity stake is small, i.e., if f<k, there will be no incentive to shield assets, but if it is
higher, the borrower would always rather shield than pay.
Proof: The proof is straightforward and so is its underlying intuition. Because of 1.1., if the bor-
rower decides to shield assets, he would attempt to shield all of them, which will come at a cost
of ke. Under equity, the borrower can meet his obligations by paying fe. It follows then that the
borrower would rather pay than shield if f<k. On the other side, if k<f, it will always be cheaper
to shield than to repay. The intuition being that because all assets are shielded, shielding imposes
a per-unit of wealth cost. Equity also imposes a per-unit of wealth cost. If the equity stake is
small, the cost per-unit of repaying the debt is lower than the cost of shielding assets. We con-
clude that a low equity interest can mitigate shielding incentives and, under the assumption of
fixed marginal cost, may solve them completely. QED.
The following figure illustrates this proposition. As can be seen, an equity agreement that requires
a fraction of costs that lies below the marginal shielding cost would provide the borrower an in-
centive to repay even in bad states of the world. Unlike a debt contract, the lender is thus assured
payment in all states.
12
Electronic copy available at: https://ssrn.com/abstract=2820650
Figure 3: Repayment and Shielding Costs as a Function of Earnings, Debt versus Equity
3.2. Stage 1 Analysis: Loan Negotiation
Let us examine now the parties’ decision regarding lending and borrowing in the first stage.
Based on the analysis just made, the parties negotiate with the expectation that the borrower
would repay only if e>e*. The parties negotiate over r and since the analysis implies that e* de-
pends on r, it will be useful to describe e* as e*(r).
The lender thus expects full repayment if, and only if, e>e*(r), making his expected payoff:
𝑒
∫ (𝑏+𝑟) 𝑓(𝑒)d𝑒 (4)
𝑒∗(𝑟)
The following proposition summarizes parties’ decisions in this model.
Proposition 3.
If assets can be shielded, then:
3.1. If the returns on the investment are expected to be sufficiently high –i.e., 𝐹(𝑒∗(𝑟)) < 1−
𝑏
, for some r, lending will take place despite the potential for asset shielding;
𝑏+𝑟
3.2. Asset shielding may result in denial of credit to net positive value investments through two
channels:
13
Electronic copy available at: https://ssrn.com/abstract=2820650
3.2.1. Value Reduction. Shielding reduces the expected value of the investment to a potentially
e∗(r)
negative value due to shielding costs(=∫ t(e) 𝑓(𝑒)d𝑒); or,
𝑒̲
3.2.2. Feedback Effect. Increasing interest may reduce expected payments, thus making interest
adjustments insufficient to compensate the lender for shielding risk, making it so that there will
not be an equilibrium interest rate.
Proof. See Appendix.
The first part of Proposition 2 reflects the notion that when wealth is high, shielding becomes ir-
rational. The lender will only lend if his participation constraint is met, that is, if the lender can
expect to recoup her investment b. This will happen when there is sufficient probability of the
wealth threshold being met, because then the debt plus interest are paid in full. As is familiar, the
potential for shielding in the other cases could be compensated through the higher interest rate.
And if the investment returns are expected to be high in all states of the world, it may be that the
interest charged would be at the risk-free rate, despite borrower’s technical ability to shield assets.
This result—that lending is rational even in the presence of effective shielding technology—could
partially explain credit markets in jurisdictions with poor enforcement and the practice of mostly
unsecured loans to high-payoff investments, such as start-ups. The following example illustrates:
Example 3: The borrower seeks a loan of $6,000 and the investment is certain to earn at
least $8,000. Since this earning amount is greater than the wealth threshold identified above of
$7,500, the lender knows, with certainty, that the borrower will have an incentive to repay.
Hence, even with an interest of 0, the lender will be able to recover the full $6,000.
Example 3a: Now the investment is expected to earn an amount between $0-30,000 with
a uniform distribution. If the interest is set at $6,000, the wealth threshold becomes $15,000 (for
the cost of shielding that amount is $12,000, the same as the debt plus interest). Since there is
50% chance of exceeding the wealth threshold, this secures the lender an expected payment of
0.5*12,000=$6,000. So the lender should be willing to lend even if shielding is expected to take
place in some states of the world. Borrower’s payoff below the wealth threshold has an expected
value of 0.2*7,500, and above the threshold of 22,500-12,000, for a total expected payoff of
0.5*0.2*7,500+0.5*(22,500-12,000)=$6,000.
The second part of the Proposition indicates two channels by which asset shielding destroys value
ex-ante. The first is the fact that shielding assets consumes resources that would otherwise be
available to the parties. If the expected surplus of the investment is $3,000, but the expected
shielding cost is $5,000, then the investment no longer has a positive expected value.
The second channel is subtly different. As was just noted, the risk of shielding means that a high-
er interest must be charged. But recall that the incentive to shield rises with the debt (e* is in-
creasing in r). Increasing r in (4) has the double effect of increasing payment conditional on re-
payment, but reducing the overall probability of repayment (by increasing the lower bound of the
integral). To compensate for the higher incentive to shield, an even higher interest may be
charged. But this can create a feedback effect, with higher interest increasing shielding incentives,
14
Electronic copy available at: https://ssrn.com/abstract=2820650
requiring higher interest, etc. And so, for various investments, there may not exist any equilibri-
um interest rate for which the lender will be willing to lend and the borrower willing to borrow.
An implication of this Proposition concerns the lender’s preference between safe and risky pro-
jects of the same expected value. Safer projects are characterized by lower variability of returns,
which also entails a lower probability of exceeding the wealth threshold. This would mean that a
lender my have a perverse preference to risky projects in the presence of asset shielding risk.
Note that denial of credit through these channels operates through a different mechanism than that
identified by Stiglitz and Weiss (1981). Instead of ex-ante moral hazard or adverse selection, here
the problem is one of ex-post moral hazard.
Social Optimum. At the beginning of the Date 1, the revenue from the investment is uncertain
and has an expected value of E(e). The net value of the investment is E(e) - b, which is assumed
to be positive. It was noted above that during the second stage, it is socially desirable for assets
not to be shielded. Therefore, assuming t=0, it is socially desirable for the positive expected value
investment to be undertaken. When assets may be shielded ex-post, the value of the investment
𝑒∗(𝑟)
becomes 𝐸(𝑒)−𝑏−∫ 𝑡(𝑒) 𝑓(𝑒)𝑑𝑒 , that is, the expected value falls by the costs of shield-
𝑒
ing. Hence, an otherwise socially desirable investment may become undesirable if shielding is
expected.
Private Decisions when Shielding is Limited. Lender’s payoff in this akin to (4):
𝑒
∫ (𝑏+𝑟) 𝑓(𝑒)d𝑒 (5)
𝑒∗(𝑟)
We have noted that e* decreases for lower s'(), i.e., if shielding technology is limited, the wealth
threshold will be lower. A lower threshold implies a higher probability of repayment, thus reduc-
ing interest rates and mitigating the problems just identified. Importantly, however, these prob-
lems may not be completely solved as long as s'()>0, for there could still be some incentive to
shield. The following example illustrates
Example 4. Suppose the same circumstances as in 3a. only that now shielding technolo-
gy is limited so that the borrower retains only 10 cents on each dollar shielded, instead of 20.
With this change, if the interest is set at $3,000, the wealth threshold is $10,000 (for
0.9*10,000=6000+3000). There is 2/3 chance of exceeding this threshold, thus securing the lend-
er an expected return of 2/3*9,000=6,000, so the lender would indeed be willing to lend at this
rate. This also means that the borrower pays $3,000 less in interest in this case, due to the limited
ability to shield. Borrower’s expected payoff under this example would be 1/3 * 0.1*5,000+ 2/3 *
(20,000-9000)=$7,500. This is an improvement of $1,500 over example 3a, all due to the limited
ability to shield assets.
The analysis changes when shielding is completely ineffective. In this case, the borrower is indif-
ferent between repaying and shielding for low earnings, e≤b+r, and will strictly prefer repaying
when earnings are high, e>b+r. This implies an expected return to the lender of min(e, b+r),
leading to a familiar prediction: while the interest rate may be positive, lending to a positive ex-
15
Electronic copy available at: https://ssrn.com/abstract=2820650
pected value investment will always take place. This is in contrast to the shielding model, where
lending to such projects was not guaranteed.
We see then that a credible commitment not to shield, and likewise effective enforcement tech-
nology, would lead to lower interest and can solve credit denial problems. This is clearly in the
borrower’s self-interest, ex-ante, to be able to commit to not shielding thus securing finance when
he otherwise would not and keeping a greater share of the investment surplus. The difficulty is
that standard contractual mechanisms involve only pecuniary sanctions, and thus provide no teeth
to a commitment not to shield
Equity Agreement. We turn now to examine the possibility of an equity agreement instead of a
debt contract. At Stage 1, the parties set f or r endogenously. Based on Proposition 2, it is an eq-
uity agreement will dominate ex-ante a debt contract from both the borrower’s and lender’s per-
spective.
Suppose first that there exists f<k that meets the lender’s participation constraint (i.e.,
𝑒̅
∫ 𝑓𝑒 𝑓(𝑒)𝑑𝑒 = 𝑏, recalling that for f<k there is no shielding). The lender would then find it ra-
𝑒
tional to lend, and the borrower would enjoy from improved access to credit at lower rates.
If it is impossible to set such f, then it will also be impossible to find r that would permit lending.
This is because for the debt contract to meet the lender’s participation constraint, there must be
some e'≥e*(r) for which ke'≥b+r. That is, there must be some of level of wealth for which the
borrower prefers repayment to shielding. But if this is the case, it is possible to set f=k, and offer
the lender a return that is at least as large as under debt. Hence, if the lender is willing to lend
with debt, he should be willing to lend with equity as well.
This conclusion, that equity stakes dominate debt contracts is strong and so it is important to note
its limits. The literature on costly-state falsification analyzes the optimality of such “shielding-
proof” contracts (Lacker and Weinberg 1989). As this literature highlights, the optimality of equi-
ty agreements is limited. While they eliminate waste due to shielding, they force allocations that
the parties may not find optimal ex-ante.
4. EXTENSIONS
I will now relax some of the assumptions made in the basic model, seeking to explore its rele-
vance to the use of collateral, debt relief, imperfect shielding techniques, and the shielding of in-
vestment proceeds.
4.1. Secured Credit, Collateral, and Equity Cushions
Parties sometimes often secured credit as means of reducing lending risk. Using this technique,
the lender receives a security interest in the borrower’s assets. How would that affect shielding
behavior?
Despite their name, security interests do not resolve the problem of asset shielding. While securi-
ty interests are generally effective in affecting the priorities among creditors in bankruptcy, they
16
Electronic copy available at: https://ssrn.com/abstract=2820650
only have a limited effect on shielding behavior. If the borrower owns a consumer good with a
lien on it, he may still sell it to a third party or hide it. The lien cannot prevent it and, in fact, a
third party may retain ownership of the assets despite creditor’s objection. This is not to say that a
lien is of no value: liens may make shielding more complicated and costly. But, as we saw, a
more expensive shielding technology does not necessarily solve the problem of asset shielding.
A much more effective mechanism is the use of possessory collateral. Suppose now that the bor-
rower starts with some positive wealth w. If the lender is in possession of the collateral, the bor-
rower will not be able to shield it. Hence, default will imply a cost of w on the borrower (loss of
collateral). This has two consequences. First, the wealth threshold is now implicitly given by e-
s(e)-w=b+r, which is lower than otherwise, meaning less shielding risk and, consequently, lower
interest rates and less credit denial. Second, because shielding costs w, it means that if w≥b+r,
there will be no incentive to shield when the collateral is large, as the cost will always be higher
than that of repayment. But possessory collateral is very costly to pledge and involves high and
sometime debilitating operational costs, making it of limited use.
Interestingly, even nonpossessory collateral is valuable in mitigating shielding risk – despite bor-
rower’s ability to shield it. The reason it is useful is that shielding the collateral itself is costly,
and this costs is added to the rest of the shielding costs. The wealth threshold is then implicitly
given by e+w-s(e+w)=b+r which is lower than otherwise (but higher than under possessory col-
lateral).21 This can be interpreted more broadly, as an advantage of equity cushions. Due to the
recourse principle, all of borrower’s equity cushion acts as implicit nonpossessory collateral, thus
reducing the incentive to shield. This means that collateral need not be pledged explicitly to be
effective.
Overall, this implies that high-wealth individuals pose a lesser lending risk, allowing them greater
access to credit markets. Additionally, this analysis highlights the function of nonpossessory col-
lateral, not only in the relationship between creditors but also between the debtor and the creditor.
4.2. Renegotiation and Debt Forgiveness
In situations where the borrower has a credible threat ex-post to shield assets, the parties may
want to renegotiate the debt. The lender would find such a renegotiation desirable, as with asset
shielding the lender is not repaid at all. This presents an opportunity for the lender to offer some
‘debt relief’ in exchange for partial payment.
From the borrower’s perspective, debt relief is beneficial if the lender is willing to accept an
amount that is equal or lower than borrower’s cost of shielding assets. That is, if the lender is
willing to forgo the debt claim in exchange for an amount no greater than e-s(e). if the lender is
21 An indirect effect is due to the somewhat lower marketability of secured assets, which will increase the marginal cost
of shielding those assets.
17
Electronic copy available at: https://ssrn.com/abstract=2820650
willing to accept that amount, the borrower would avoid the original debt at a lower cost than by
shielding assets. Hence, both parties can be made better off by ex-post debt relief.22
Is such an agreement realistic? If the borrower approaches the lender demanding debt relief,
would not the lender, who knows the debtor can pay, try to enforce the debt at that moment?
While plausible, we observe in practice many instances of debt relief where some debt is forgiven
or “settled” (NY City Bar 2012). Such agreements are generally based on the implicit under-
standing that the debtor is not giving literally everything he has, and instead, the parties negotiate
some lesser amount.23 This lends realism to the asset-shielding avoiding motive.
Ex-ante, the prospect of debt relief will mitigate the problems caused by asset shielding, but will
not solve them completely. First because of the transaction costs involved in ex-post renegotiation
and second because the lender will not be able to receive more than the cost of shielding in bad
states of the world (relative to project’s full returns in the case of no shielding), so the return may
still be insufficient to induce lending. Stated more abstractly, debt relief is akin to an equity
agreement in bad states of the world and a debt contract in good states of the world.
4.3. Costly Collection, Uncertain Collection, and Collection of Shielded Assets
In the main analysis, collection of exposed assets was certain. Relaxing this assumption may af-
fect the magnitude of assets that the borrower would seek to shield.24
Suppose first that collection has a fixed cost of c. 25 The lender will not find collection profitable
unless e-s(e)>c. This will allow the borrower to leave c in asset value exposed without fear of
collection. For the borrower, shielding costs are then reduced by c-s(c), thus further exacerbating
shielding incentives and their resultant negative effects.26
It may also be that collection costs depend on the amount collected. Let l(m) be the cost function
of collecting an amount m. Suppose that l(m) is either everywhere increasing or decreasing. Let
m* be the value of m for which l'(m*)=1. If collection costs are decreasing, the borrower could
leave up to m* in assets exposed and they will be essentially protected. This reduces the amount
the borrower needs to shield by m*, having a similar effect to that of a positive c.
If collection costs are increasing, collecting more than m* would be unprofitable. This implies for
the borrower that leaving assets exposed is akin to losing m* in value. Now, if m*>b+r, this will
not change the analysis. However, if m*<e-s(e), the borrower will find it better to leave some as-
22 The enforceability of debt relief agreement itself is not a problem, if debt relief is conditional on the payment being
made. If the borrower breaches, then he will have to incur the higher costs of shielding.
23 This is not intended to imply that the primary motivation for debt relief agreement is the avoidance of asset shielding.
There are other motivations for such agreements, such as saving collection costs. Whichever of these motivations plays
a more dominant role is an open empirical question.
24 An option that I do not address here directly is that shielded assets will be “de-shielded” by creditor’s efforts. This
option could be simply conceptualized as a higher cost of shielding.
25 Some fixed collection costs are court and sheriff fees, information requests, time spent on collection, and profession-
al assistance fees.
26 If e<b+r+c, the borrower may choose to repay only e-c.
18
Electronic copy available at: https://ssrn.com/abstract=2820650
sets exposed than to shield them. Note that in practice, lenders often pay a contingency fee to
debt-collectors, so that lenders may have an incentive to pursue all debts, regardless of costs, alt-
hough the debt collectors themselves may employ similar reasoning to this just noted.27
Let us examine now the situation where exposed assets are not always be collected due, for ex-
ample, to legal limitations or the difficulty of locating them. Suppose that assets exposed are col-
lected with a probability p. In this case, the borrower will compare his payoff conditional on
complete shielding s(e) with the payoff from leaving assets exposed, pe.28 If, for simplicity, mar-
ginal shielding cost is fixed at k, shielding will only be rational if k<p. From the lender’s perspec-
tive, not shielding is better than shielding, as expected recovery becomes (1-p)e instead of 0.
Finally, let us consider the case where assets can be ‘de-shielded’. If there is some probability that
assets that were shielded can be uncovered, then this risk can be conceptualized as part of the cost
of shielding, making shielding costs t-s'(t) instead of t-s(t). The implications of this were already
discussed above.
4.4. Ex-Ante Shielding
It is possible that the borrower could also shield the loan proceeds b. However, this possibility
will not change the analysis. A risk-neutral borrower would be better off investing the proceeds
and then shielding the greater expected returns than shielding the lower proceeds of the loan.
5. CONCLUSION AND POLICY IMPLICATIONS
This paper emphasizes the point that shielding is not worthwhile unless the borrower leaves ex-
posed fewer assets than his debt. A chief implication of this point is that shielding would not be
worthwhile if a borrower has high wealth, either due to high returns or to initial endowment. On
the other hand, a poorer borrower, even if solvent, may find shielding rational. Therefore, the so-
cial problems associated with asset shielding, wasted resources, and denial of credit, would be of
greater relevance for either poor borrowers or investments that do not succeed.
The conclusions of the analysis are quite general and tend to hold even when collection is costly
or uncertain and when the borrower has opportunity to shield prior to making the investment. Re-
negotiation mitigates some of the social problems of asset shielding, but still involves higher in-
terest and some degree of credit denial. Collateral can be used to further ameliorate the problem,
but it will not generally solve it. Equity agreements do provide a promising solution, although
they may not always be feasible in practice.
In closing, I would like to briefly highlight a few regulatory implications of this analysis. As the
analysis shows, asset shielding can lead to a social loss. The parties may not always be able to
solve it directly. One reason can be transaction costs, but a more interesting problem is that con-
27 The debt collection industry commission average is reported to be around 19% (Fedaseyeu 2013)
28 The actual decision may be more complicated, as it may be rational to shield only some assets but leave others ex-
posed. But this will not change the qualitative nature of the conclusion to follow.
19
Electronic copy available at: https://ssrn.com/abstract=2820650
tractual sanctions are mostly pecuniary. As a result, in exactly those circumstances when a party
shields assets, these sanctions would be irrelevant.
Given those limitations, the most direct means of controlling the problem is making shielding
more costly, by, for example, reversing suspicious transactions, extending clawback periods, us-
ing criminal sanctions more extensively, or limiting bankruptcy exemptions.29 More generally,
the recourse principle limits shielding and so exceptions to this rule (as in non-recourse loans)
should be carefully considered.
Another possible mechanism is insurance requirements and use of vicarious liability (as in Pitch-
ford 1995 and Mattiacci and Parisi 2003). Having a third-party involved, such as an insurance
company, would reduce lender’s risk, but will not necessarily solve the problem, as the borrower
may still be able to shield assets. It is only if third parties have an advantage over creditors in
monitoring borrowers—which they might have as in the case of multiple lenders—that third party
involvement may be a solution to this problem. Future work may develop the conditions under
which such a system can be useful.
The legal system may also control the problems of asset shielding more indirectly, by manipulat-
ing the wealth threshold. This can be done by requiring actors engaged in dangerous activities to
meet a minimal capital requirement. Such an equity cushion, we have noted, reduces shielding
incentives by reducing the wealth threshold. The legal system may also reduce shielding incen-
tives by imposing lower fines on under-capitalized debtors (in the spirit of Ganuza and Gomez
2008), or, similarly, order the payment of fines in installments over a long period of time. We
noted that the wealth threshold is lower if the debt is lower, so that such solutions may mitigate
the problem.
This outline of policy interventions suggest that there is much that can be done to control shield-
ing risk. Understanding the problem of ex-post asset shielding and the underlying incentives
could inform work in these areas.
* I thank Christopher Avery, Oren Bar-Gill, Osnat Jacobi, Louis Kaplow, Roy Shapira, Steven
Shavell, and Kathryn Spier for valuable comments. I would also like to thank the John M. Olin
Center for Law, Economics, and Business at Harvard Law School for financial support.
APPENDIX
Proof of Proposition 1.
1.1. I will now show that it is optimal to set t* at either 0 or e (i.e., there are no intermediate values of
t). To see the first case, suppose that a low t is considered such that t < e - (b + r), i.e., t is less
than borrower’s ex-post wealth less the debt obligation. This means that the second term in (1) is
positive (i.e., max(0, e-t-(b+r))>0). We can therefore simplify (1) to s(t) + e - t - b – r. And be-
cause s(t)-t<0,it would be best to set t=0. If, instead, a higher t is considered, such that t > e - (b
29 When shielding leads to suboptimal care, the legal system may seek to directly
20
Electronic copy available at: https://ssrn.com/abstract=2820650
+ r), that would make borrower’s payoff in (1) =s(t), which is clearly optimal to set at its highest
30
value, i.e., e. QED.
Note: the proof is relatively general and it holds regardless of the marginal cost of shielding (as
long as it is positive), so that it applies even if the marginal cost of shielding exceeds the marginal
amount shielded.
1.2. By 1.1., the optimal level of t, conditional on shielding, is e.. Hence, shielding costs can be
expressed as e-s(e). As s'(e)<1, Shielding costs are increasing in e. This implies that for some
level of e, we will have e-s(e)=b+r. let e* denote this level. The benefit of shielding is avoiding
up to b+r in costs, so when e>e*, the costs of shielding exceed the benefit. It is straightforward
to see from this formulation that for higher r the level of e* also increases.
1.3. Because the cost of shielding for any e>e* is greater than b+r, the borrower will be better off
paying b+r and retaining e-(b+r) than shielding. QED.
Proof of Proposition 2.
2.1. Let us look at the case where, for simplicity, r=0 and e>e*(0). In all states of the world bor-
rower’s wealth exceeds the threshold identified in 1.2., so shielding will not take place, and the
expected return in (5) becomes simply b+0. Hence, the loan is guaranteed to be repaid in full and
the interest would be indeed set at zero.
If, however, e<e*(0), it is still possible that the loan will be provided. This will be the case if
there exists r for which:
ē
∫ (𝑏+𝑟) 𝑓(𝑒)d𝑒 ≥ 𝑏
𝑒∗(𝑟)
But it must be that r >0 in this case, because there will not be any payment in the states of the
world when e<e*(0).31
Note. For lending to take place despite shielding risk, the expected investment earnings have to
𝑒∗(𝑟)
be higher than they would be under no asset shielding by ∫ e 𝑓(𝑒)𝑑𝑒, which reflects the re-
𝑒
duction in repayments due to asset shielding. This provides one measure of assessing the social
cost of shielding.
2.2.1. The value of the investment is reduced in all cases where e<e*(r) because the borrower
𝑒∗(𝑟)
spends e-s(e) on shielding. This reduces the (joint) value of the investment by ∫ (𝑒−
0
𝑠(𝑒)) 𝑓(𝑒)d𝑒; hence, if the net value of the investment, 𝐸(𝑒)−𝑏, is lower than this expected re-
duction, the investment will have a negative expected value and there will not exist a level of in-
terest for which the lender would be willing to lend.
30 The same result of full shielding applies, and for the same reason, when borrower’s wealth is low, i.e., e<b+r.
31 For example, suppose c(p) is a constant $0.5 per dollar, b = $10,000, and there is a 20% chance that e̲=$1,000 will
be realized and a 80% chance that e̅=$40,000 will be realized. If r=$2,500, then e*=$37,500 and there is 80% chance
that $12,500 will be repaid, meaning an expected payment of 0.8*12,500=$10,000 – enough to induce the lender to
lend.
21
Electronic copy available at: https://ssrn.com/abstract=2820650
2.2.2. Suppose first that the parties contemplate setting r=0 and that e<e*(0). The amount repaid
ē
is given by (3) as ∫ (𝑏) 𝑓(𝑒)d𝑒. Since this will be lower than b due to shielding, the interest
𝑒∗(0)
must be increased to some higher level, denoted as r1>0. By slightly increasing the interest, the
lender increases his expected payoff by the probability of repayment times the higher payment,
i.e., (1-F(e*(r1))(b+r1). At the same time, however, the higher debt would reduce the probability
of repayment, so that the expected repayment falls by the lost revenues, i.e., F(e*(r1)-F(e*(0)))b.
Now, if under the new schedule the net amount received (the gain from the higher interest less the
loss from the lower probability of repayment) falls below b, r1 would have to be set even higher.
It is plain to see now that further increasing the interest may repeat this dynamic, so that r would
need to be further and further increased. At the extreme, r could be set any arbitrarily high level,
r≥e̅, but then surely no payments will be made, because if the debt is greater than borrower’s
wealth it always pays to shield (by 1.1.). Hence, it may be that no interest rate would exist such
that the necessary amount will be repaid, and the loan will be denied (even if, after deducting the
costs of shielding, it has a positive value).32 QED.
* Terence M. Considine and Private Law Fellow, Harvard Law School. yar-
bel@sjd.law.harvard.edu. I would like to thank I thank Christopher Avery, Oren Bar-Gill, Louis
Kaplow, Koby Kastiel, Roy Shapira, Steven Shavell, Henry Smith, and Kathryn Spier for valua-
ble comments, as well as to the anonymous reviewer. I would also like to thank the John M. Olin
Center for Law, Economics, and Business at Harvard Law School for financial support.
REFERENCES
Alper, Elijah, M. 2007. Opportunistic Informal Bankruptcy: How BAPCPA may Fail to Make
Wealthy Debtors Pay Up. Columbia Law Review 107(8):1908-1943
NY City Bar. 2012. “Profiteering from Financial Distress: An Examination of the Debt Settle-
ment Industry”, available at
http://www2.nycbar.org/pdf/report/uploads/DebtSettlementWhitePaperCivilCtConsumerAffairsR
eportFINAL5.11.12.pdf
Che, Yeon Koo, and Kathryn E. Spier. 2008. “Strategic Judgment Proofing.” RAND Journal of
Economics 39 (4): 926–48.
Dari-Mattiacci, Giuseppe, and Gerrit De Geest. 2005. “Judgment Proofness Under Four Different
Precaution Technologies.” Journal of Institutional and Theoretical Economics 161 (1): 38–56.
32 For example, suppose shielding cost is a constant $0.5 per dollar, the loan is $10,000 and that there is equal probabil-
ity of returns being either $10,000 or $30,000—i.e., expected return is $20,000. If r=0, the critical level is $30,000, so
there is %50 chance of the borrower repaying, implying an expected return to the lender of $5,000. To adjust, the inter-
est must be set to some positive amount, but this will only increase the critical level to be above $30,000, implying 0
overall returns to the lender.
22
Electronic copy available at: https://ssrn.com/abstract=2820650
Dari-Mattiacci, Giuseppe, and Barbara Mangan. 2008. “Disappearing Defendants Versus Judg-
ment Proof Injurers.” Economica 75 (300): 749–65.
Dari-Mattiacci, Giuseppe, and Francesco Parisi. 2003. “The Cost of Delegated Control: Vicarious
Liability, Secondary Liability And Mandatory Insurance.” International Review of Law and Eco-
nomics 23 (4): 453–75.
Elinson, Zusha 2010. “City Finds Millions, but Reaps Little in Case,” New York Times, Decem-
ber, 11 2010.
Arbel, Yonathan A., 2015. “Contract Remedies in Action: Specific Performance”, West Virginia
Law Review 118:100-141.
Ellingsen, Tore and Eirik Gaard Kristiansen 2011. “Financial Contracting Under Imperfect En-
forcement,” The Quarterly Journal of Economics 126 (1): 323-371.
Ernest & Young (EY). 2012. “The Impact of Third Party Debt Collection on the National and
State Economies”.
Fay, Scott, Eric Hurst, and Michelle White. 2002. “The Household Bankruptcy Decision,” Ameri-
can Economics Review 92:706–718.
Ganuza, Juan-José, and Fernando Gómez. 2008. “Realistic Standards: Optimal Negligence with
Limited Liability.” The Journal of Legal Studies 37 (2): 577–94.
Ganuza, Juan José, and Fernando Gomez. 2011. “Soft Negligence Standards and the Strategic
Choice of Firm Size.” Journal of Legal Studies 40 (2): 438–66.
Gross, David B. and Nicholas S. Souleles. 2002. “Do Liquidity Constraints and Interest Rates
Matter for Consumer Behavior? Evidence from Credit Card Data,” Quarterly Journal of Econom-
ics 117 (1): 149-185.
Hart, Oliver and John Moore. 1989. Default and Renegotiation: A Dynamic Model of Debt, Quar-
terly Journal of Economics 113(1): 1-41.
Lacker, Jeffery M. 1991. “Why is there Debt? Economic Review,” July-Aug:3–19.
Lacker, Jeffery M. and John A. Weinberg. 1989. “Optimal Contracts under Costly State Falsifica-
tion,” Journal of Political Economy, 97(6):1345-1363.
Leff, Arthur A. 1970. “Injury, Ignorance and Spite — The Dynamics of Coercive Collection,”
Yale Law Journal 80 (1):1-46.
McCullough, Ralph C. 1997. “Bankruptcy Fraud: Crime without Punishment II,” Commercial
Law Journal 102:1-54.
Pitchford, Rohan. 1995. “How Liable Should a Lender Be? The Case of Judgment-Proof Firms
and Environmental Risk.” American Economic Review 85 (5): 1171–86.
Price Waterhouse Coopers (PWC) 2008, “Value of Third-party Debt Collection to the US. Econ-
omy in 2007: Survey and Analysis.”
Schwartz, Alan. 1983. “The Enforceability of Security Interests in Consumer Goods,” Journal of
Law & Economics 26(1): 117-162.
Scott, Austin W. 1913. “The Right to Follow Money Wrongfully Mingled with Other Money,”
Harvard Law Review 27(2): 125-138.
Stephen, Lea E. G, Avril J. Mewse, and Wendy Wrapson. 2013. “The Psychology of Debt in
Poor Households in Britain” in A Debtor World (Ralph Brubaker, Robert M. Lawless, and
Charles J. Tabb eds.).
23
Electronic copy available at: https://ssrn.com/abstract=2820650
Shavell, Steven. 1986. “The Judgment Proof Problem.” International Review of Law and Eco-
nomics 6. Springer: 45-68.
Shavell, Steven. 2005. “Minimum Asset Requirements and Compulsory Liability Insurance as
Solutions to the Judgment-Proof Problem.” The RAND Journal of Economics 36 (1): 63.
Shenon, Philip. 2001. “Home Exemptions Snag Bankruptcy Bill.” New York Times, April 6 2001.
Sullivan, Teresa A., Jay L. Westbrook, and Elizabeth Warren. 1989. As We Forgive Our Debtors:
Bankruptcy and Consumer Credit in America. New York, NY: Oxford University Press.
Summers, John. 1983. “The Case of the Disappearing Defendant: An Economic Analysis.” U.
Penn. L. Rev 1422 (1979): 145–72.
Stiglitz, Joseph E. and Andrew Weiss. 1981. “Credit Rationing in Markets with Imperfect Infor-
mation,” The American Economic Review 71 (3):393-410.
US Courts. 2012. Bankruptcy Statistics.
http://www.uscourts.gov/Statistics/BankruptcyStatistics/12-month-period-ending-march.aspx.
US DOJ. 2012. Public Report: Borrower Audits by the United States Trustee Program.
Veld, van’t Klaas, and Emma Hutchinson. 2009. “Excessive Spending by Firms to Avoid Acci-
dents: Is It a Concern in Practice?” International Review of Law and Economics 29 (4): 324–35.
White, Michelle J. 2007. “Abuse or Protection: The Economics of Bankruptcy Reform under
BAPCPA.” University of Illinois Law Review 2007 (1):275-304.
Wickelgren, Abraham L. “Settlement and the Strict Liability-Negligence Comparison.” U of Tex-
as Law, Law and Econ Research Paper No. 213.
Zhu, Ning. 2011. “Household Consumption and Personal Bankruptcy”, Journal of Legal Studies,
40 (1): 1-37
24
Electronic copy available at: https://ssrn.com/abstract=2820650
`;

function asObject() {
  return {
    paper_id: PAPER_ID,
    title: TITLE,
    ssrn_url: SSRN_URL,
    year: YEAR,
    authors: AUTHORS,
    keywords: KEYWORDS,
    summary_md: SUMMARY_MD,
    summary_zh_md: SUMMARY_ZH_MD,
    one_pager_md: ONE_PAGER_MD,
    study_pack_md: STUDY_PACK_MD,
    article_text: ARTICLE_TEXT,
  };
}

module.exports = {
  PAPER_ID, TITLE, SSRN_URL, YEAR, AUTHORS, KEYWORDS,
  SUMMARY_MD, SUMMARY_ZH_MD, ONE_PAGER_MD, STUDY_PACK_MD, ARTICLE_TEXT,
  asObject,
};

if (require.main === module) {
  process.stdout.write(ARTICLE_TEXT);
}
