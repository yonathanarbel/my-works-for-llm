/*
PAYDAY 6/2/2020 6:22 PM (ssrn-3547007) — corpus code wrapper

This file intentionally embeds the paper text and study assets in code form.
It helps code-centric ingestion pipelines and makes the corpus easy to load programmatically.
*/

const PAPER_ID = "ssrn-3547007";
const TITLE = "PAYDAY 6/2/2020 6:22 PM";
const SSRN_URL = "https://papers.ssrn.com/sol3/papers.cfm?abstract_id=3547007";
const YEAR = null;
const AUTHORS = [];
const KEYWORDS = [];

const SUMMARY_MD = `

`;

const SUMMARY_ZH_MD = `

`;

const ONE_PAGER_MD = `
# PAYDAY 6/2/2020 6:22 PM — one-page summary

**Paper ID:** \`ssrn-3547007\`
**SSRN:** https://papers.ssrn.com/sol3/papers.cfm?abstract_id=3547007

## TL;DR

PAYDAY FORTHCOMING: 98 WASH. U. L. REV. 1 (2020) Draft: Comments, Suggestions, and Critique Welcome!

## Files

- Full text: \`papers/ssrn-3547007/paper.txt\`
- PDF: \`papers/ssrn-3547007/paper.pdf\`

_Auto-generated study aid. For canonical content, rely on \`paper.txt\`/\`paper.pdf\`._
`;

const STUDY_PACK_MD = `
# Study pack: PAYDAY 6/2/2020 6:22 PM (ssrn-3547007)

- SSRN: https://papers.ssrn.com/sol3/papers.cfm?abstract_id=3547007
- Full text: \`papers/ssrn-3547007/paper.txt\`

## Elevator pitch

PAYDAY FORTHCOMING: 98 WASH. U. L. REV. 1 (2020) Draft: Comments, Suggestions, and Critique Welcome!

## Suggested questions (for RAG / study)

- What is the paper’s main claim and what problem does it solve?
- What method/data does it use (if any), and what are the main results?
- What assumptions are doing the most work?
- What are the limitations or failure modes the author flags?
- How does this connect to the author’s other papers in this corpus?

_Auto-generated study aid. For canonical content, rely on \`paper.txt\`/\`paper.pdf\`._
`;

const ARTICLE_TEXT = `
PAYDAY 6/2/2020 6:22 PM
PAYDAY
FORTHCOMING: 98 WASH. U. L. REV. 1 (2020)
Draft: Comments, Suggestions, and Critique Welcome!
Yonathan A. Arbel
Legislation lags behind technology all too often. While trillions of
dollars are exchanged in online transactions—safely, cheaply, and
instantaneously—workers still must wait two weeks to a month to
receive payments from their employers. In the modern economy,
workers are effectively lending money to their employers, as they wait
for earned wages to be paid.
The same worker who taps a credit card to pay for groceries in
semi-automated checkout lines depends on dated payroll systems
that only transfer payments on a “payday.” Workers, especially those
living paycheck-to-paycheck, are hard-pressed to meet their daily
needs and turn to expensive, short-term credit products—notably,
payday lenders. While the need for credit is a real one, credit
providers charge a steep price, often culminating in endless debt
spirals. So, why does the payday still exist?
This Article studies various explanations—economic, historical,
behavioral, and legal. A primary conclusion is that the payday owes
its existence to legacy legal architecture. That is, payday is a software
problem, not a hardware problem. The hardware—i.e., money and
payroll technology—is here. We can pay workers daily; in fact, gig
economy workers in developing countries will often be paid more
quickly than an American employee for the same work. What holds
us back is our legal software: Dated Eisenhower-era legislation that
failed to anticipate technological change. Surprisingly, even pro-
worker legislation, such as minimum wage laws, inadvertently
encourage the practice.
By revealing the overlooked and dated legal infrastructure that
sustains the payday, the Article suggests a path for legal reform.
Daily streams of payment to workers are feasible, practical, and far
more efficient than most people realize. A focused reform could
effectively bring an end to the puzzling and pernicious practice of
 Assistant Professor of Law, University of Alabama, School of Law. For suggestions
and comments I would like to thank John Acevedo, Oren Bar-Gill, Matt Bodie, Deepa Das
Acevedo, Stephanie Didwania, Shahar Dillbary, Don Dornberg, Brian Galle, Victor
Goldberg, Avery Katz, Jonathan Korn, Ron Krotoszynski, Ronald Mann, Ben McMichael,
Jonathan Nash, Mike Pardo, Michael Pressman, Robert Scott, Steve Shavell, Jonathan Korn,
Vincent Buccola, and Paige Skiba, and participants in the following conferences and
workshops—Law and Economic Organization (Columbia), Contracts Conference XIV,
Faculty Workshop (SMU), Midwestern Law & Economics, Law & Economics Workshop
(TAU), and Private & Commercial Law (Hebrew U). Tori Moffa, Hamilton Millwee, Tamara
Imam, and Brett Linley provided diligent and careful research assistance. A companion
playlist can be found here: http://bit.ly/payday_comp
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

2 Draft[Vol. __
having workers lend money to their employers while they wait for
their payday.
INTRODUCTION ........................................................................................ 3
I.THE PAYDAY PUZZLE ............................................................................ 8
A. The Two Employment Contracts ......................................... 8
B. The Puzzle of K2 ................................................................. 11
II.PAYDAY: HISTORICAL, LEGAL, SOCIAL, AND ECONOMIC
EXPLANATIONS ............................................................................ 15
A. Path-Dependence ................................................................ 16
B. The Synchronization of Bills and the Payday ................... 20
C. Employer Power and Lack of Sophistication .................... 22
D. Collateral ............................................................................. 26
E. Behavioral Biases ............................................................... 28
F. Legislation ........................................................................... 33
G. Check Cutting Costs ........................................................... 37
III.ABOLISHING THE PAYDAY ................................................................ 41
A. The Stakes of Abolishing the Payday ................................ 41
B. Alternatives to Abolition .................................................... 45
IV.A WORLD WITHOUT THE PAYDAY ..................................................... 48
A. Changing by Information ................................................... 51
B. Changing by Leading .......................................................... 52
C. Fixing Employment Law .................................................... 53
D. Improving Money Technology ............................................ 54
V.THE DAY AFTER PAYDAY: CONCLUDING THOUGHTS ........................ 57
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

2020] Payday 3
INTRODUCTION
Legislation often lags behind technology. As Guido Calabresi
observed, “laws are governing us that would not and could not be
enacted today.”1 This failure is resounding in the context of
employment contracts. Payment technology has made incredible
advances, and today trillions of dollars are traded in the online
economy, moving between parties almost instantaneously.2 At the
same time, workers still wait for weeks until a formal “payday” to
receive their hard-earned wages. While workers sell their labor today,
employers only pay them in the future, leveraging wages as another
line of credit.
We seem to take the payday’s existence for granted,3 but it
exacts a heavy price. Workers who wait for payment need to support
themselves; the vicissitudes of everyday life—a sudden toothache, a
flat tire, a stain on their only clean work shirt—demand money, now.4
With many workers living paycheck-to-paycheck,5 the current payday
system pushes them to payday lenders and other short-term credit
providers that dot the modern urban landscape.6 A payday loan is
1 GUIDO CALABRESI, COMMON LAW FOR THE AGE OF STATUTES 2 (1982).
2 16.9%+ Growth for Online Payment Gateway Market Size Raising to USD 4020
Million by 2024, MARKETWATCH (May 21, 2019), https://www.marketwatch.com/press-
release/169-growth-for-online-payment-gateway-market-size-raising-to-usd-4020-million-
by-2024-2019-05-21.
3 The modern literature has mostly neglected this question. This omission is perhaps
most glaring in law and economics analyses of employment contracts, but it is by no means
confined to these works. See, e.g., RICHARD POSNER, THE ECONOMIC ANALYSIS OF LAW (9th
ed. 2014) (reviewing major topics but neglecting pay frequency); MARK ROTHSTEIN & LANCE
LIEBMAN, EMPLOYMENT LAW, 420–21 (2011) (adumbrating pay frequency). But cf. JOHN R.
COMMONS & JOHN B. ANDREWS, PRINCIPLES OF LABOR LEGISLATION 50–52 (1916) (noting the
credit nature of the payday).
4 The three leading reasons why individuals borrow from alternative lenders (such as
payday lenders, pawn shops, and rent-to-own stores) are basic living expenses, making up for
lost income, and house or car repairs. Neil Bhutta et al., Consumer Borrowing After Payday
Loan Bans, 59 J.L. & ECON. 225, 240 (2016). See also Rob Levy & Joshua Sledge, A Complex
Portrait: An Examination of Small-Dollar Credit Consumers, CTR. FOR FIN. SERVS.
INNOVATION, 12 (Aug. 2012),
https://www.fdic.gov/news/conferences/consumersymposium/2012/A%20Complex%20Port
rait.pdf [https://perma.cc/D78A-RLT3] (reporting that approximately 37% of very short-term
borrowers borrowed because “[they] had a bill or payment due before [their] paycheck
arrived.” In addition, 30% of respondents borrowed to meet some unexpected expense). This
borrowing likely result from the payday. See also Nicholas Bianchi & Rob Levy, Know Your
Borrower: The Four Need Cases of Small-Dollar Credit Consumers, CTR. FOR FIN. SERVS.
INNOVATION, 12 (2013), https://s3.amazonaws.com/cfsi-innovation-files/wp-
content/uploads/2017/01/26054909/Know-Your-Borrower-The-Four-Need-Cases-of-Small-
Dollar-Credit-Consumers.pdf [https://perma.cc/G5AB-W4PG] (finding that 32% of
consumers borrow because of misaligned cash flow and 32% to meet an unexcepted expense).
Again, both reasons can be mitigated by regularized pay.
5 15% of households reported having spent more than they earned over the last year.
Bricker et al., Changes in U.S. Family Finances from 2013 to 2016, 103 FED. RES. BULL. 1, 8
(2017).
6 Paige Marta Skiba, Regulation of Payday Loans: Misguided?, 69 WASH. & LEE L. REV.
1023, 1031 n.22 (2012) (noting that “payday lenders outnumber both Starbucks and
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

4 Draft[Vol. __
meant to help the worker bridge the gap until payday, but it involves
interest rates are on average twenty times higher than those of credit
cards.7 A $300 loan can quickly balloon into thousands of dollars of
outstanding debt, leading many borrowers to a debt spirals that can
culminate in deep financial distress and even bankruptcy.8
This Article begins by framing the payday in the context of the
employment contract. The employment relationship is, at its core, an
exchange of money for labor.9 The payday also injects into this
relationship a credit transaction, one where the employee is lending
money to the employer. But this is a credit transaction that is
completely artificial from the viewpoint of financial theory. Put
simply, workers should not be in the business of lending money to
their employers.10 Not only do workers lack capital or comparative
specialization in lending, but they are also badly positioned to deal
with counterparty risk.11 A value-creating credit transaction moves
money from those who have it to those who need; not from the
Walmart employee to Walmart.
If the payday does not serve a clear financial purpose, what
might explain its dogged persistence? The Article evaluates a variety
of reasons: economic, sociological, historical, legislative, and even
psychological. The primary conclusion is that the payday is a software
problem, not a hardware problem. The hardware of the economy, both
McDonalds”). Roughly 64% of all adult Americans have at least one credit card and carry an
average balance of $4,800. In addition, roughly 50% of all credit card holders carry a
revolving balance on which they pay interest. CONSUMER FIN. PROT. BUREAU, THE CONSUMER
CREDIT CARD MARKET 46, 48 n.16, 56 (2017),
https://files.consumerfinance.gov/f/documents/cfpb_consumer-credit-card-market-
report_2017.pdf [https://perma.cc/A9UA-TS3E].
7 Levy & Sledge, supra note 4, at 12 (reporting that approximately 37% of borrowers
borrowed because “I had a bill or payment due before my paycheck arrived. In addition, 30%
of respondents borrowed to meet some unexpected expense”). This issue can also be
considered as resulting from payday. See also Bianchi & Levy, supra note 4, at 12 (finding
that 32% of consumers borrow because of misaligned cash flow and 32% to meet an excepted
expense). Both reasons can be mitigated by regularized pay
8 CONSUMER FIN. PROT. BUREAU, CFPB DATA POINT: PAYDAY LENDING 4 (2014),
https://files.consumerfinance.gov/f/201403_cfpb_report_payday-lending.pdf
[https://perma.cc/KD2Y-YXJ7] (finding that 80% of payday loans are rolled over or followed
by an additional loan and that 15% of loans are followed by a loan sequence of at least 10
loans).
To experience firsthand the process of obtaining a payday loan, I borrowed $200 from a
payday lender in Tuscaloosa, Alabama. I signed a postdated check to the benefit of the lender
for $235, representing a 638.75% APR. See @ProfArbel, TWITTER (Nov. 22, 2019, 5:05 PM),
https://twitter.com/ProfArbel/status/1198014702762283008. Sociologist Lisa Sevron worked
for a payday lender and reported her experiences in LISA SEVRON, THE UNBANKING OF
AMERICA (2018).
9 See, e.g., COMMONS & ANDREWS, supra note 3, at 2 (describing the employment
contract as a “relation between a propertyless [sic] seller of himself, on the one hand, and a
propertied buyer on the other”).
10 See infra Part I.B for a discussion of this point.
11 Counterparty risk is defined as the “the likelihood or probability that one of those
involved in a transaction might default on its contractual obligation.” Chris B. Murphy,
Counterparty Risk, INVESTOPEDIA (May 14, 2019),
https://www.investopedia.com/terms/c/counterpartyrisk.asp [https://perma.cc/K7EA-
AEUA].
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

2020] Payday 5
money and payroll technologies, has greatly advanced over the last
century, allowing us to quickly and cheaply pay for both goods and
services. To wit, a freelancer doing work in India for an American
employer as part of the gig economy, who performs the same work as
an American employee, will often be paid faster than the American
counterpart.12 What hinders progress is our legal software:13
Eisenhower-era legislation that failed to keep pace with modern
technology. In fact, as this Article reveals, the culprit is often pro-
worker legislation, which stands in the way of progress, sometimes
actively encouraging longer pay periods.
This Article’s central message is that abolishing the payday is
desirable, efficient, and surprisingly feasible. To move to a system of
daily pay, two challenges of legal origin must be overcome: compliance
costs and payment costs. To assure compliance with legal norms,
employers must verify payments—and doing so daily can be
expensive. Transferring money to employees is also costly, given the
sizable minority of workers who are unbanked and under-banked.14
How can we offer payments at scale without compromising
compliance costs or burdening workers with check-cashing costs?
To address these issues and others, the proposed framework
offers to decouple compliance from pay.15 Every day, workers are to
receive roughly 93% of their daily pay, leaving some slack until a
biweekly “accounting day.”16 On accounting day, the employer verifies
compliance and makes true-up adjustments as needed. To address
issues of money transfer, which are of particular concern for the
unbanked and the underbanked, I explore the increasing use of
digital money and payroll cards. The Article concludes that moving
to daily streams of payment is both feasible and desirable, although
it contemplates a transition period. By abolishing payday, we can
spare employees the indignities of the payday, increase consumer
liquidity, enhance worker autonomy, reduce the size of the payday
lending industry, and improve the American economy as a whole.17
12 In the online platform upwork.com, hourly workers receive weekly pay five day
afterwards (a 12 day cycle). See https://community.upwork.com/t5/Announcements/Faster-
payouts-for-hourly-contracts/m-p/739876. Freelancer.com allows some contractors to
withdraw payments within a single business day after verification.
https://www.freelancer.com/support/freelancer/payments/daily-withdrawals
13 Conceptualizing legislation as software is a productive metaphor and suggests a
different paradigm to that envisioned in CALABRESI, supra note 1. Both legislation and
software need to be updated to account for new circumstances and new information; both
need to combine efforts of different groups, sometimes with different agendas; both worry
about documentation of designer intent; and, both face complex inter-dependancies. Software
technology has created a number of interesting solutions to these problems that the legal
literature is yet to address, such as alpha and beta versions, periodic updates, branches, and
commits. See generally Git Theory, GITHUB, https://github.com/SCOREC/core/wiki/Git-
Theory [https://perma.cc/5GFU-2682] (last visited Feb. 13, 2020).
14 I discuss the phenomenon and problems of the unbanked and underbanked infra notes
199–212 and accompanying text.
15 See infra Part I.A.Changing by Information
16 For a discussion of the methodology behind this framework, see infra Part IV.
17 For conceptual clarity, daily streams of payments are no longer payday in the
conventional sense of a special day which aggregates pay for multiple days of work.
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

6 Draft[Vol. __
This Article highlights the importance of regularly updating our
legal software. Payday legislation started as a mode of progressive
reform towards the end of the nineteenth century. Overcoming initial
resistance from legislators and courts, payday laws were passed to
discourage predatory behavior of companies, which were lending to
their employees at usurious rates. Remarkably, despite the poor
money and payroll technologies that existed at the time, the
legislation was effective and for a short period of time, workers were
paid weekly. By an ironic twist of fate, it is possibly the rise of the
welfare state that led to the move from weekly to the much slower
biweekly pay.18 The birth of the welfare state was spurred by the
introduction of social security and social security taxes. The
administrative burden occasioned by various related laws, such as the
Federal Insurance Contributions Act (FICA), the Federal
Unemployment Tax Act (FUTA), the Fair Labor Standards Act
(FLSA), and tax withholdings made frequent pay more difficult. Thus,
the same laws that were meant to protect employees ended up
harming them in an unanticipated way: by depressing the frequency
of pay, they increased the need for expensive short-term credit
solutions.
This Article unfolds in four Parts. Part I sets the stage by
explaining the tenuous relationship between employment contracts
and the payday. Part II explores a variety of reasons for the existence
of the payday and evaluates whether any counsels in favor of keeping
of this practice. Part III explains why the payday should be abolished
and Part IV explains how this could be achieved in practice.
To understand why the payday exists, Part I covers the basic
theory of employment contracts. It explains why the payday is not a
natural part of employment contracts and why, from a finance
perspective, it is an artificial and inefficient credit transaction.
If financial logic doesn’t explain the existence of payday, what
does? Part II explores a variety of potential reasons and
justifications—historical, legal, economic, psychological, and
sociological. Special attention is given to a psychological attempt to
justify the payday: the idea that the payday helps employees
overcome some of the behavioral challenges of saving and budgeting
their own money.19 Refuting this idea is important because some
might worry that moving to daily streams of payment would lead to
profligacy among employees. To this end, I present empirical evidence
that frequent pay does not increase spending. In fact, there is some
reason to worry that infrequent pay may result in excessive spending,
because of the higher availability of cash on hand. Most important,
however, is the argument that employer-side savings are extremely
risky, as they expose employees to opportunistic behavior,
18 There are various terms of art used to describe pay frequency. For expositional
simplicity, this Article refers to payment modes that are more frequent than once a month and
less frequent than once a week as ‘biweekly.’ See infra Part I.A.
19 See, e.g. Christopher A. Parsons & Edward D. Van Wesep, The Timing of Pay, 109 J.
FIN. ECON. 373 (2013).
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

2020] Payday 7
counterparty risk, and employer bankruptcy.20 To the extent that
workers need help managing money, an insured, trusted financial
institution provides a much more robust solution than postponing
wages.
Part II highlights one especially worrisome reason for the
continued existence of the payday: ineffective legislation. For public
sector employees, legislation often mandates by fiat long pay periods.
The President of the United States is paid, by law, on a monthly
basis.21 In the private sector, badly drafted legislation also
encourages late payments; in particular, and not without irony, wage
and hour legislation unwittingly encourages long pay schedules.
These defects, overlooked by employment law scholars and
policymakers, have substantial consequences for the welfare of
employees.22
Understanding the sources of the payday allows the
development of solutions. Part III first explains the large stakes
involved in abolishing the payday. It then explains why abolishing
the payday is imperative and why seemingly more moderate
alternatives, such as advance payments, are insufficient and risky. It
closes by examining the legislative changes that would be needed to
abolish the payday.
The key proposal here, developed in Part IV, is to move from
biweekly pay to daily streams of payment of the good faith estimate
of the employee’s daily pay. Every two weeks, the employer will have
an “accounting day,” and will add to the day’s pay any shortfall in
payments. For the part of the workforce that is either unbanked or
underbanked, payments can be made using pay cards and similar
Fintech solutions. While there are some nuances and practical
considerations in implementing this proposal, it is important to
recognize at the outset that it does not derogate from the rights of
either employees or employers. By contrast, adopting this proposal
will greatly advance the welfare of all American employees and would
also take a bite out of the large payday lending industry, increase
worker autonomy, and correct some historical defects in legislation.
In fact, implementing this proposal only requires modest changes to
the legislative framework.23
20 See also Shlomo Benartzi et al., The Law and Economics of Company Stock in 401(k)
Plans, 50 J.L. & ECON. 45, 46 (2007) (arguing that employees over-invest in their employers’
stock and that “investing a dollar in company stock . . . is often worth only 50 cents.”).
21 3 U.S.C. § 102 (2004); U.S. CONST. art. II § 1(7) (“The President shall, at stated Times,
receive for his Services, a Compensation…”).
22 The Restatement of Employment Law defers to the employer’s choice regarding the
payday. RESTATEMENT OF EMP’T LAW § 3.01 cmt. a (AM. LAW INST. 2015) (“Employees also
have a right to be paid the compensation they have earned on a timely basis, usually in
conformity with the employer's normal payroll practices.”)
23 Pay frequency interacts in complex ways with a variety of workers’ rights and issues,
such as wage theft, wage discrimination, and minimum wage. For example, frequent pay
would expand workers’ ability to sue for equal-pay violations, as the Lilly Ledbetter Fair Pay
Act of 2009, Pub. L. No. 111-2, 123 Stat. 5 (codified at 42 U.S.C. § 2000e-5 (2009)) holds
that every payment resets the 180-day statute of limitations. In general, frequent pay will tend
to expand worker rights and, at the very least, will not derogate them.
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

8 Draft[Vol. __
Even if one disagrees with any of the specific policy
prescriptions, the key message of this Article is that the payday
should not be taken as a neutral or natural fact of the modern
economy. The existence of the payday has substantial consequences
in terms of efficiency, distribution, and autonomy. While we live in an
exceptional period with historically low-interest rates, the harms of
the payday will only be amplified as interest rates rise.24 The recent
outbreak of Covid-19 powerfully demonstrates the importance of
liquidity. Many of the recent developments in Fintech suggest that
the payday lives on borrowed time. It is perhaps time to call this loan.
I. THE PAYDAY PUZZLE
A. The Two Employment Contracts
What is the purpose of an employment contract? Roughly 130
million Americans are considered employees and are thus parties to
an employment contract.25 These contracts feature a great deal of
variability, as they each stipulate different norms the employee must
abide by—the employee’s various rights, benefits, and perquisites, as
well as the employee’s duties, obligations, and loyalties. Still, at its
core, the contract is premised on a very basic economic transaction: a
“bargained-for exchange of labor for consideration.”26 The
employment contract is an exchange relationship, which the parties
seek to optimize according to their own circumstances.27
This exchange transaction stands at the heart of the
employment contract, and I denote it here as K . In this K , the
1 1
employee is selling labor, broadly defined as time, skill, effort, and
any other aspect of his or her human or social capital. In
consideration, the employer gives the employee “money,” which could
include wages, tips, perquisites, in-kind transfers, and any other
value that redounds to the employee from the employer. When the
employment contract describes the employee’s duties, it outlines the
scope of labor that is exchanged. When the employment contract
stipulates the employee’s pay and benefits, it states the payment that
is exchanged for this labor. The concept of K is sufficiently capacious
1
and abstract to capture all employment contracts, despite the fact
that they differ in almost any other respect. In this high level of
24 Federal Funds Rate—62 Year Historical Chart, MACROTRENDS,
https://www.macrotrends.net/2015/fed-funds-rate-historical-chart [https://perma.cc/KLS7-
97UK] (last visited Feb. 13, 2020) [hereinafter MACROTRENDS].
25 BUREAU OF LAB. STAT., LABOR FORCE STATISTICS FROM THE CURRENT POPULATION
SURVEY (2019), https://www.bls.gov/cps/cpsaat08.htm [https://perma.cc/5XN7-HFLA].
26 Vanskike v. Peters, 974 F.2d 806, 809 (7th Cir. 1992).
27 PATRICK BOLTON & MATHIAS DEWATRIPONT, CONTRACT THEORY 4 (2005) (studying
the optimal design of the exchange relationship).
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

2020] Payday 9
abstraction, we can say that K is responsible for the annual exchange
1
of at least 6.4 trillion dollars.28
What both economists and lawyers will often miss is another
striking regularity in modern employment contracts. Besides the K
1
aspect of the transaction, most contracts also include a payday—a gap
in time between the moment work is rendered and payment is
transferred. Almost all payments by the employer are paid in
arrears—that is, after the employee “gave” their labor to the
employer. The following figure summarizes the frequency of the
payday and the typical lag involved in payments, based on data made
available by the Bureau of Labor Statistics for the years 2014 and
2019:
Figure 1 - Pay Frequency, US Private Businesses
Sources: Bureau of Labor Statistics (BLS); Private Payroll Company (ADP)
2014 2017 (ADP) 2019
66%
61%
56%
32% 34%
22%
11%11%
5%
WEEKLY BIWEEKLY OR MONTHLY
SEMIMONTHLY
This figure summarizes pay frequency data, based on a very
large sample of nonfarm employees.29 The chart shows that most
American employees are paid twice a month, on either a biweekly or
a semimonthly basis.30 The difference between biweekly and
semimonthly is fairly subtle; a biweekly payday takes place every
fourteen days, while a semimonthly payday takes place twice a
month, on two separate days (e.g., the 1st and the 20th). Given the
existence of fifty-two workweeks in a year, this means that a biweekly
28 Based on the product of 130 million full time employees, supra note 25, working 52
weeks per year and earning on average $956 per week,
https://www.bls.gov/news.release/empsit.t19.htm, the
29 BUREAU OF LAB. STAT., LENGTH OF PAY PERIODS IN THE CURRENT EMPLOYMENT
STATISTICS SURVEY (2019), https://www.bls.gov/ces/publications/length-pay-period.htm#fn5
[https://perma.cc/G39D-JJQQ] (farm workers were excluded from this study). The ADP data
comes from a private payroll company, ADP, as reported in Tomaz Cajner et al., Using
Payroll Processor Microdata to Measure Aggregate Labor Market Activity, Finance and
Economics Discussion Series 2018-005, BD. OF GOVERNORS OF THE FED. RES. SYS. 44 (2018),
https://doi.org/10.17016/FEDS.2018.005 [https://perma.cc/LN9E-Y2AA].
30 Biweekly also denotes twice a week; however, in the wage payment context, it is used
to denote payment frequency of once every 14 days.
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

10 Draft[Vol. __
payday translates to either twenty-six or twenty-seven paydays per
year, whereas a semimonthly payday entails a fixed number of
twenty-four paydays. Beyond the twice-monthly pay, a sizable
minority of employees are paid weekly and a small minority on a
monthly basis.31
The existence of a payday may seem obvious—indeed, many
take it for granted—but it hides significant complexity. The worker is
providing work today: stacking the shelves, cleaning the floor,
building a wall, attending to customers, etc. But for services rendered
today, the employee is only paid in the future, on payday. In other
words, payments in the economy are, by and large, in arrears.
Thus, the very idea of the payday implies a temporal distance
between the moment the employee is providing services, the quid of
K , and the moment she is paid, the quo of K . As noted by Commons
1 1
and Andrews in their 1908 treatise on labor law:32
When the laborer starts to work for [the employer], he
also becomes, for a time, a creditor. He contributes his
services in advance of compensation. He is a temporary
investor in the business. While he works he passes over
to the employer the title to his product, and retains a
claim for wages. When his wages are paid his investment
is liquidated.
The economic classification of this aspect of the transaction is
straightforward. When a person buys a car from the dealership, he or
she can pay on the spot for the exchange. But he or she can also agree
with the dealer to pay in the future, perhaps in monthly installments.
This is the financing part of the exchange. By the same logic, when
an employer receives services today but pays for them in the future,
on payday, this is a credit transaction. In addition to the exchange
relationship, what we called K , the employment contract thus
1
embodies a second credit transaction, what we might call K . This K
2 2
contains the agreement between the parties to defer payment for
money earned until payday. The parties will not always explicitly set
the payday in the contract, but of course, they agree to some kind of
payday—and this part of the agreement, explicit or implicit, makes
K . As in any credit agreement, we can identify three parts: an
2
employee-lender, an employer-borrower, and wages-principal.
A natural question is whether this is a true credit transaction,
as K does not seem to indicate any interest rates. This, however,
2
should not be too distracting: Credit transactions do not require
explicit quotes of interest or even any interest at all to count as credit
transactions. Consider how auto traders will sometimes offer ‘zero-
interest financing.’ The auto trader will not really offer a free loan,
but rather, will build the cost of the loan into the price of the car.
Some part of the price, then, can be seen as interest—the premium
31 The data collection methodology is not sufficiently clear to discern what share of
American employees are paid on shorter time spans than weekly.
32 COMMONS & ANDREWS, supra note 3, at 50.
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

2020] Payday 11
the dealer charges for offering ‘free’ finance. And even if the trader
charges no interest at all, it would still be a loan that would have to
be repaid on pain of default and collection. That is to say, a loan is a
loan even if it does not involve interest payments.33
It may be tempting to try and define the problem away. If we
were to define the unit of work as two weeks’ full of work, there
wouldn’t be K , because the payment is only due when the work-unit
2
is completed. On reflection, however, such definitional games are
unpersuasive. Defining work in two-weeks units is ad-hoc and does
not map any underlying transfer of value. Effort, skill, and time do
not come in two-weeks increments, rather, they are continuous. The
worker’s daily expenses, as well, do not come in such neat packages.
In fact, employers have attempted to redefine labor units; in one case,
they sought to define work as a year’s full of work.34 This way,
employers hoped, they did not have to pay until the end of the year
and if the employee quits—or is encouraged to quit—before the end
of the year, they could avoid the obligation to pay. For sound policy
reasons, courts and legislators rejected this view.35 More
theoretically, if the employee is understood to be selling time, then
time does not come at two-week increments.36
*
To quickly recap, so far, we have considered the existence of two
“contracts” implicit in the employment relationship: K and K . K is
1 2 1
the standard exchange of labor for money; K is the credit transaction
2
whereby payments for K will only be made on payday. The K loan
1 2
includes some “interest” payment in the form of higher than
otherwise wages. With this in mind, we can turn our attention to how
odd K appears from a finance perspective.
2
B. The Puzzle of K2
Finance theory teaches that, at the most fundamental level,
loans create value by moving money from those who have it to those
who need it.37 Banks lend money to cash-strapped businesses,
33 Loans also have a maturity date; here, the it is the payday. In a biweekly K2, the worker
lends 1/14 of the salary daily to the employer. The period until maturity shortens every day;
at first, the loan is for 13 days, but on the last day of work, the loan is only for that same day.
On average, the maturity date is 6.5 days in the future and the loan is remade every two weeks.
In a daily pay system, the loan mature on the same day it is paid, so it involves minimal
interest, and so I do not explore here the possibility of hourly pay.
34 Britton v. Turner, 6 N.H. 481, 481, 485–86 (1834) (holding that, despite the employee
quitting before the end of the stipulated year of work, the employer still had an obligation to
pay under restitution); Matthew T. Bodie, Employment As Fiduciary Relationship, 105 GEO.
L.J. 819, 840, n. 133 (2017) (“Modern wage payment schemes require that employees be
paid . . . for all time worked, regardless of the length of term”).
35 See infra notes 121–129 and accompanying text.
36 With independent contractors, it is sometimes the case that payment is made on an
project-completion basis (even though, even there, advances are common). Employment
contracts, however, normally separate work tasks and payment, and pay on the basis of time
worked.
37 See generally Dan Bernhardt, Money and Loans, 56 REV. ECON. STUD. 89 (1989).
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

12 Draft[Vol. __
venture capitalists to promising entrepreneurs, and bondholders to
growing companies. Such transactions create value because it they
are mutually advantageous. A loan enables the borrower to seize
profitable investment opportunities and smooth consumption over
time. At the same time, the loan also allows the lender to use its
money as a source of profit, through interest payments. As long as the
interest payment is between the value to the borrower and the cost of
lending to the lender, the parties would find a credit transaction
mutually advantageous.
This basic logic of finance is well recognized; however, applying
it to employment contracts presents a puzzle. As we just saw, K is a
2
ubiquitous part of the economy. It covers the Walmart employee
stocking the shelves, the grocery store teller working the register, and
the cook at McDonald’s flipping burgers; it covers employees from
store clerks to university professors to executives. In all of these
cases, K facilitates a loan from employees to employers—it is a loan
2
from those with little money to those with more money. Why, then, is
the Walmart employee lending money to Walmart? Why are service
technicians lending money to Comcast? And why are police officers
lending money to the government?
It may be tempting to answer these questions with the same
logic as any other financial transaction. The borrower (employer)
borrows because it benefits from having cash on hand and the lender
(employee) lends because it profits from the interest payment. On
reflection, however, the benefits to employers are vastly exceeded by
costs to employees. The intuition is straightforward: households are
in no position to lend money to firms.
The benefit to employers from K loans is relatively small. One
2
reason for that is that some employers do not even need cash.
Consider how the publicly listed firm Alphabet, despite holding $117
billion in cash, still uses K with its janitors, programmers, and
2
marketers.38 Apple holds $100 billion in cash, and Microsoft lags with
only $50 billion, yet both use the payday.39 The federal government is
also not particularly cash hungry, and yet it mandates the use of a
biweekly payday in all of its employment contracts.40 Even employees
of the Federal Reserve—which quite literally prints money—are paid
on a biweekly basis.41 This offers some evidence that the reason
behind K is not liquidity.
2
Still, many employers, especially small businesses, are not as
cash rich as these companies, and they do stand to benefit from liquid
funds. However, even for those employers, the benefit from K loans
2
is smaller than first appears. To estimate the size of the benefit,
consider the cost of borrowing from alternative lenders. After all,
employers borrow from banks, capital markets, specialized lenders,
38 Richard Waters, Google Parent Alphabet Overtakes Apple to Become New King of
Cash, FIN. TIMES (July 31, 2019), https://www.ft.com/content/332dd974-b349-11e9-8cb2-
799a3a8cf37b.
39 Id.
40 5 U.S.C. § 5504 (2018).
41 Telephone Interview with Payroll Department, Federal Reserve (Feb. 5, 2020).
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

2020] Payday 13
and a variety of other sources. In 2019, for example, the weighted
average interest rate on loans to small businesses ranged from 5.1%
to 5.66%.42 If we use this rate to measure the gain the business
receives from paying a typical employee in arrears, it becomes clear
that the gain is fairly small. For an employee earning $50,000
annually who is paid on a monthly basis, the annual payday credit
benefit to the employer is only $108.43
The benefit to employers is not large, but the cost to employees
of lending money is significant. Employees are not in a position, nor
do they have the skills, to lend money to their employers.44
Monitoring and secured credit, two common features of credit
transactions, are all but absent in the employee-employer
relationship. And, of course, to lend money one needs liquid cash. But
workers are often subject to severe liquidity constraints which make
it very costly for them to offer loans to their employers. 40% of
Americans with a credit card carry a credit card balance,45 and
roughly 16% of households reported essential expenses that are
unmet, with 11.5% percent reporting unpaid utilities.46 In a survey,
roughly 21% of households reported difficulty in accessing credit for
their own needs.47
Rather than being providers of cheap credit, households often
turn to expensive credit products to finance daily expenses—such as
42 See FED. RES. BANK OF KAN. CITY (Sep. 2019),
https://www.kansascityfed.org/~/media/files/publicat/research/indicatorsdata/smallbusiness/
2019/sbls_aggregatedatapdf_sept2019.pdf?la=en [https://perma.cc/N6TN-SEYB].
43 The calculation assumes daily compounding with 5% APR and the average salary in
2019 of $50,000. The calculation itself is not straightforward due to compounding, but it can
be approximated in the following manner.
$50,000 per annum implies a salary payment of $137 per day. At the beginning of the
month, the employee has to wait roughly 30 days to be paid. On the last day of the month,
however, the employee receives pay on the same day. On average, then, each payment is
delayed by 15 days . 5% APR implies a daily interest rate of 0.014% (0.05/365). This means
that the employee is lending every day of the year, on average, $137 for 15 days at a rate of
0.014%. Overall, the value of this transaction is 365∗137∗15∗0.00014=105.01. (The
difference between $105 and 108 is due to compounding).
44 Most lenders will not lend absent a credit check and, where a large part of their
portfolio is staked with a specific borrower, would require contractual controls. Neither of
these characterize household lending decisions to employers.
45 CONSUMER FIN. PROT. BUREAU, supra note 6, at 55. The bottom 20% of Americans
have a median $2,000 in financial assets. BD. OF GOVERNORS OF THE FED. RES. SYS., 2016
SURVEY OF CONSUMER FINANCES, https://www.federalreserve.gov/econres/scfindex.htm
[https://perma.cc/6UC5-AEQG] (last visited Feb. 13, 2020) (Table G.19); cf. Kathleen Elkins,
Here’s How Much Money Americans Have in Savings at Every Income Level, CNBC (Oct.
11, 2018, 12:02 AM), https://www.cnbc.com/2018/09/27/heres-how-much-money-
americans-have-in-savings-at-every-income-level.html [https://perma.cc/Y4QV-BXDZ]
(“29% of households have less than $1,000 in savings.”).
46 JULIE SIEBENS, U.S. CENSUS BUREAU, EXTENDED MEASURES OF WELL-BEING: LIVING
CONDITIONS IN THE UNITED STATES: 2011 11, Table 3 (2013),
https://www.census.gov/prod/2013pubs/p70-136.pdf [https://perma.cc/X87A-BK5C].
47 See Bricker et al., supra note 5, at 27 (2017). See also Matt Tatham, The Number of
Americans with Bank Accounts Rises, EXPERIAN (Mar. 25, 2019),
https://www.experian.com/blogs/ask-experian/research/the-decline-of-the-unbanked-and-
underbanked/ [https://perma.cc/C94N-8H6F] (In 2017, nearly 20% of respondents were
underbanked and 14.1 million adults had no bank account).
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

14 Draft[Vol. __
payday lenders, credit card companies, advance tax refunds, and
pawnshops—and the size of these industries illustrates the need felt
by households.48 The cost of such borrowing is considerable. Congress
estimated (quite crudely) that every late-paid dollar costs the
employee an additional dollar—i.e., 100% cost of borrowing.49
However, the real costs tend to be even higher. When households
borrow, they use a variety of sources, which include bank loans (with
a ~10% cost of borrowing on average),50 credit cards (a 16% cost of
borrowing),51 and payday lenders (typically 400%).52 For those
households that use payday lending regularly, the cost of finance can
amount to a large percentage of their annual earnings.
The costs to households are not strictly financial. The liquidity
crunch has broader effects on household welfare. Lack of access to
funds is not only a financial issue; concerns with liquidity create
financial stress, which is associated with higher mortality and worse
health outcomes.53
Judged in terms of the standard model of credit, K fails to
2
produce social value. The cost of lending by the household far exceed
the benefits that accrue to the employers. True, larger employers
would reap larger benefits, but the costs to employees would scale by
the same factor. And, to be sure, if the employer does not bear these
costs, the employer might not care about them and excessively engage
in K , even if it comes at a severe cost to the worker. I will return to
2
the private incentive of employers later,54 but for now, the key point
is that from a social perspective, K destroys value because the costs
2
of the loan exceed its benefits. We—society—want businesses to
borrow using capital markets and lenders that can, more accurately,
48 In 2016, the revenue of short-term lenders (i.e., fee and interest payments) was $57.9
billion dollars. See ERIC WILSON & EVA WOLKOWITZ, CTR. FOR FIN. SERVS. INNOVATION, 2017
FINANCIALLY UNDERSERVED MARKET SIZE STUDY(2017), https://s3.amazonaws.com/cfsi-
innovation-files-2018/wp-content/uploads/2017/04/27001546/2017-Market-Size-
Report_FINAL_4.pdf [https://perma.cc/DV99-99NA].
49 29 U.S.C. § 216(b) (2018) (“Any employer [in violation] . . . shall be liable to the
employee . . . in the amount of their unpaid minimum wages . . . and in an additional equal
amount as liquidated damages.”).
50 See BD. OF GOVERNORS OF THE FED. RES. SYS., CONSUMER CREDIT, DECEMBER 2019,
(Feb. 7, 2020), https://www.federalreserve.gov/releases/g19/current/default.htm
[https://perma.cc/M4W3-5P92].
51 See Kelly Dilworth, Average Credit Card Interest Rates: Week of May 27, 2020,
Creditcards.com (May, 27, 2020), https://www.creditcards.com/credit-card-news/rate-
report.php. Timely payment of credit-card balance would avoid these interest charges, but in
practice, 47% of Americans carry a balance on their credit cards and so they pay interest on
credit-card purchases. CONSUMER FIN. PROT. BUREAU, supra note 6, at 55–56.
52 CFPB, What is a Payday Loan?, (Jun 2, 2017)
https://www.consumerfinance.gov/ask-cfpb/what-is-a-payday-loan-en-1567/
53 Todd H. Baker, FinTech Alternatives to Short-Term Small- Dollar Credit: Helping
Low-Income Working Families Escape the High-Cost Lending Trap Trap 8 (Harv. Kennedy
Sch., M-RCBG Working Paper Series No. 75, 2017),
https://www.hks.harvard.edu/sites/default/files/centers/mrcbg/files/75_final.pdf
[https://perma.cc/YB4B-6FUA].
54 I return to this point infra Section I.C
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

2020] Payday 15
price and monitor risk. We do not want to create a line of credit which
consists of employees’ wages.
Importantly, K is not a one-off transaction, so value may be
2
destroyed multiple times. It is not just that households need to bridge
the first two weeks of employment; instead, K involves a continuous
2
cycle of borrowing and repayment. Consider a hypothetical low-pay
employee starting work on January 1, 2020, with only a small amount
of cash on hand. The employee is paid biweekly and so has to borrow
on January 1 against future earnings to support daily expenses. Come
payday on January 15, the employee is paid and finally has cash on
hand. But the employee also owes money. Now the employee has to
repay the loan, plus interest, and make do with whatever is left. If the
remainder is insufficient, the employee will have to borrow again.
And again. And again. In the worst case scenario, a debt spiral
emerges—the employee would need to borrow back-to-back on a
revolving basis.55
*
Employment contracts include two key components, K and K .
1 2
K is the basic exchange of labor for capital. K is a credit transaction
1 2
that is superimposed on the employment relationship. However,
unlike K , the credit transaction of K does not generate social value—
1 2
from a financial perspective at least. The absence of financial logic
presents the payday puzzle. Households are in no position to lend
money to their employers, at least in the general case. Businesses
have better access to liquidity, pay lower interest rates, and do not
face the same pressures as individuals do when funds are running
out. We shall now explore alternative, non-financial reasons for the
existence of the payday, the topic of the next Part.
II. PAYDAY: HISTORICAL, LEGAL, SOCIAL, AND ECONOMIC
EXPLANATIONS
The payday is a fixture of modern employment contracts. As
Part I just demonstrated, however, there is nothing natural—from a
financial perspective—in the modern matrimony between K and K .
1 2
If K serves any social function, it is not one that is rooted in financial
2
logic. So what reasons could there be for the continued existence of
the payday in today’s economy?
In trying to answer this question, a broad range of possible
reasons present themselves: historical, legal, social, and economic.
My goal here is to examine and evaluate the leading reasons on the
basis of two criteria: first as an explanation and then as a
justification. This tracks the difference between understanding why
a social practice exists and understanding whether it should persist.
55 A survey in England found that one in five payday borrowers were unable to repay the
debt on time, leading to a debt cycle. Jill Insley, Payday Loan Borrowers 'Trapped In Debt
Spiral,’ THE GUARDIAN (May 18, 2012),
https://www.theguardian.com/money/2012/may/18/payday-loan-borrowers-trapped-debt-
spiral [https://perma.cc/5DZF-QFZJ].
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

16 Draft[Vol. __
The explanation for why the train is late—the conductor is a late
riser—is causally satisfactory, but it does not present a justification.
Similarly, as we will see, many of the possible explanations for the
payday fail as justifications.
A. Path-Dependence
Justice Holmes once observed that the path of the law is not logic;
it is experience.56 From keyboard layouts to tax legislation, path-
dependence explains a variety of social arrangements.57 In these
cases, past choices, justified by historical contingencies, continue to
affect decisions far into the future. Once adopted, too many social
arrangements become dependent on past historical events, making
the transition to an alternative system (even if superior) too costly.58
Consider, for example, how obsolete area codes are in phone numbers
today; although they feel natural, logical, and perhaps inevitable for
participants in the system, they are hard to explain to outsiders. In a
similar manner, the payday may be yet another instance of inefficient
social equilibria that results from path-dependence. This conclusion
becomes clear within a historical analysis that considers how the
payday emerged in an environment with inferior money and payroll
technologies.
The first moral exhortation on the payday is in the Bible, where
it is admonished that one should not "take advantage of a hired
worker who is poor and needy. . . Pay them their wages each day
before sunset, because they are poor and are counting on it.”59
Whether daily pay was indeed broadly practiced in the old world with
any regularity, though, is historically unclear.
Moving to the modern era, under early English common law,
employers were initially only required to pay within the pre-agreed
pay period; in the absence of a specific agreement, the default was
payment at the end of the contract.60 The old default presumably
reflects the idea of piece-rate work, which was a common mode of
employment in the 18th century.61 By contrast, if one counts certain
agricultural workers, such as sharecroppers, as wage laborers, their
56 Oliver Wendell Holmes, The Path of the Law, 10 HARV. L. REV. 457, 457 (1897).
57 See generally Paul A. David, Path Dependence, Its Critics, and the Quest for
“Historical Economics,” in THE EVOLUTION OF ECONOMIC INSTITUTIONS (2007); Mirit Eyal-
Cohen, Path-Dependence in Temporary Legislation (unpublished manuscript) (on file with
author) (discussing the role of path dependence in tax legislation).
58 David, supra note 57, at 10–12.
59 Deuteronomy 24:14–15; see also Leviticus 19:13 (“Do not hold back the wages of a
hired worker overnight.”).
60 ROBERT GILDERSLEEVE PATERSON, WAGE PAYMENT LEGISLATION IN THE UNITED
STATES 68–70 (1918). The duration of the employment contract was imputed, in part, from
the pay period. Jay M. Feinman, The Development of the Employment at Will Rule, 20 AM. J.
LEGAL HIST. 118, 120–21 (1976). For development of similar ideas in early American law,
see id. at 125–26.
61 E. P. Thompson, Time, Work-Discipline, and Industrial Capitalism, 38 PAST &
PRESENT 56, 78–79 (1967).
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

2020] Payday 17
pay was only seasonal.62 Still, as early as the 17th century, we find
growing indications of weekly and even daily wages in England.63 By
the 19th century wage work became the dominant form of payment
and English workers were commonly paid on a weekly basis.64 Across
the pond, American workers in the 19th century were also paid
commonly on a wage basis, but it seems like factory workers and
many other employees were only paid on a monthly basis.65
As wage payment evolved in the 19th century, it faced a critical
challenge; both payroll and money technologies were nascent and
highly inefficient.66 Taken for granted today, the use of a standard
unit of currency—the federal US dollar—was not always common in
the early American republic and the Supreme Court labored to
encourage its use.67 Monitoring hours worked and computing pay also
proved challenging, especially if one has to compute withholdings,
garnishments, benefits, and deductions for a large workforce.68 And
then there is the difficulty of disbursing pay—consider the
illuminating complaint of a nineteenth-century business owner:69
If the larger mills should pay once a week it would entail
considerable more expense. The Pacific Company
employs between five and six thousand hands, and it
would be extremely difficult for the paymaster to visit all
these people once a week, carrying his trunk up and
62 See generally Joseph D. Reid Jr., Sharecropping as an Understandable Market
Response: The Post-Bellum South, 62 J. ECON. HIST. 106, 109-120 (Mar. 1973).
63 JAMES E. THOROLD ROGERS, SIX CENTURIES OF WORK AND WAGES: THE HISTORY OF
ENGLISH LABOUR, 430 (1884),
https://socialsciences.mcmaster.ca/econ/ugcm/3ll3/rogers/sixcenturies.pdf
[https://perma.cc/552G-2LDU]; Peter H. Lindert & Jeffrey G. Williamson, English Workers’
Living Standards During the Industrial Revolution: A New Look, 36 ECON. HIST. REV. 1, 13
n.38 (Feb. 1983); Jeremy Boulton, Wage Labour in Seventeenth-Century London, 49 ECON.
HIST. REV. 268 (1996) (noting daily pay).
64 Thompson, supra note 61.
65 PATERSON, supra note 60 at 77 (noting the “custom of monthly wage payments which
prevailed in most lines of industry prior to 1885”). See also FRANCES PERKINS & ISADOR
LUBIN, BUREAU OF LABOR STATISTICS, HISTORY OF WAGES IN THE UNITED STATES FROM
COLONIAL TIMES TO 1928, 93 (1934) (noting that in 1777 pay-per-product was abolished in
the glass industry in favor of monthly pay). However, this source does not find any regular
pay period across industries. See, e.g., id. at 90, 92. See also CHRISTOPHER L. TOMLINS, LAW,
LABOR, AND IDEOLOGY IN THE EARLY AMERICAN REPUBLIC 275 (1993) (citing M'Millan and
M'Millan v. Vanderlip, 12 Johnson 165 (N.Y. 1815)).
66 I turn to the technological issues infra Part I.F. In addition, wage work requires time
technology—a watch—as emphasized by Thompson, supra note 61, and more conceptually,
“the abstraction of a man's labour from both his person and the product of his work. . . . [and]
a method of measuring the labour one has purchased, for purposes of payment, commonly by
introducing a second abstraction, namely labour-time.” MOSES I. FINLEY, THE ANCIENT
ECONOMY 65 (1973).
67 SHARON ANN MURPHY, OTHER PEOPLE'S MONEY: HOW BANKING WORKED IN THE
EARLY AMERICAN REPUBLIC (2017).
68 See infra Part I.G.
69 Esther Redmount et al., The Effect of Wage Payment Reform on Workers’ Labor
Supply, Wages, and Welfare, 72 J. ECON. HISTORY 1064, 1069 (2012).
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

18 Draft[Vol. __
down stairs, and taking receipts from each one. He has
to go to the help so as not to stop the work.
These difficulties with cash and computation seem dated today,
but they were of utter importance in the time when wage pay evolved.
The evolution of the payday faced another formative moment
towards the end of the nineteenth century. As part of a large
movement of workers, wage and salary workers started organizing
and lobbying for legislation that would mandate more frequent pay.70
Their efforts were initially met with strong resistance. Many
legislators were unresponsive, and even when the legislature was
responsive, courts were reluctant to approve pay frequency
legislation.71 Such regulation was seen as an unwarranted imposition
on the parties’ freedom of contract and a due process violation.72
The first large win for workers was in Massachusetts.73 The
charismatic governor of Massachusetts, George D. Robinson, was a
champion of regular pay. In the legislative hearing, he urged that a
weekly payday be implemented for several reasons. The proposed law
would increase worker autonomy, limit the scope of debt collection
lawsuits, increase the use of cash (a major concern at the time),74 and
instill a better sense of money management among employees.75 He
also noted that the experience from voluntary weekly pay was
favorable and thus refuted many of the chief concerns. Workers still
saved and did not “waste their earnings in frequent debaucherie
[sic].”76 Indeed, even large employers found that the system was
practicable and added few costs.77
The weekly payday in Massachusetts signaled a national
change. Reports on the enforcement of this law seem positive.78 Other
70 PATERSON, supra note 60, at 70.
71 LINDLEY D. CLARK & STANLEY J. TRACY, BUREAU OF LAB. STATISTICS, LAWS
RELATING TO PAYMENTS OF WAGES 16–19 (1926),
https://fraser.stlouisfed.org/files/docs/publications/bls/bls_0408_1926.pdf
[https://perma.cc/MHS5-RU69].
72 See PATERSON, supra note 60, at 92–93 (documenting twelve cases where wage
regulation was deemed unconstitutional and fourteen where it was also constitutional).
73 Massachusetts Acts of 1879 Ch 128 p. 483; cited in PATERSON, supra note 60, at 70;
Redmount et al., supra note 69, at 1024. See also PATERSON, supra note 60, at 68 (noting that
wage period laws are “comparatively recent origin”). For the legislative history, see Am. Mut.
Liab. Ins. Co. v. Comm'r of Labor & Indus., 163 N.E.2d 19, 21 (1959).
74 See MURPHY, supra note 67Error! Bookmark not defined., at 17–20.
75 See George D. Robinson, Address of His Excellency George D. Robinson to the Two
Branches of the Legislature of Massachusetts 33 (Jan. 3, 1884) (“[T]he lesson of economy be
practically taught every day.”). See also id. at 36–38.
76 Id. at 33.
77 Id. at 34 (“It is, I submit, always wise and salutary to devise legislation of such a
character as will reach the humblest and the poorest citizen, who has no voice but his own to
present his needs, — no power in combination with others to emphasize his opinions.”).
78 KAN. DEP’T OF LAB. AND INDUS. INDU THIRD ANNUAL REPORT OF THE BUREAU OF
LABOR AND INDUSTRIAL STATISTICS, 324–25 (1888). See also See ST. OF N.Y., SEVENTH
ANNUAL REPORT OF THE FACTORY INSPECTORS OF THE STATE OF NEW YORK, 45–46 (1893);
Robinson, supra note 75, at 36–38.
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

2020] Payday 19
states followed suit and adopted weekly or biweekly pay periods.79
Courts, too, changed their attitude and grew increasingly accepting
of such provisions.80 One reason for this growing acceptance was the
concern that employers use their bargaining power to offer unfair
loans (advances) to employees.81 Another was the concern that
regular payment is “much more a matter of life and death to a
workingman . . . than to the employing corporation.”82 Even the
Supreme Court weighed in and held that states are well within their
powers to regulate pay frequency legislation.83 This ruling came only
nine years after Lochner,84 but it withstood Lochner era standards,
as it was seen more as a form of preventing fraud and abuse than
substantive regulation of the terms of the deal.85
The boom in payday regulation was followed by a quick bust. As
soon as 1908, most states had already moved to the modern system of
biweekly pay.86 Massachusetts was the last bastion of weekly pay,87
but even there the practice has changed drastically. In 1959, the
weekly pay law was still on the books, but many companies were
paying biweekly.88 In a high-profile case, the Supreme Judicial Court
of Massachusetts ruled that weekly pay was still the norm,89 but the
decision recognized that it was perhaps time for a change.90 Others
criticized the decision for creating “unnecessary paper work . . . and
add[ing] administrative burdens.”91 Soon thereafter, the legislature
changed the law to allow for biweekly pay.92
79 PATERSON, supra note 60, at 70–88. A few states adopted a monthly pay obligation.
Id. at 88–92. One example of weekly pay is 1891 R.I. Pub. Laws 38.
80 Claudio J. Katz, Protective Labor Legislation in the Courts: Substantive Due Process
and Fairness in the Progressive Era, 31 L. & HIST. REV. 275, 288 (2013).
81 Steven L. Willborn, Indirect Threats to the Wages of Low-Income Workers:
Garnishment and Payday Loans, 45 STETSON. L. REV. 35, 40 (2015); State v. Brown & Sharpe
Mfg. Co., A. 246, 252 (1892) (grounding payday legislation in a concern with “the greed of
corporate capital.”).
82 COMMONS & ANDREWS, supra note 3, at 51.
83 Erie R. Co. v. Williams, 233 U.S. 685 (1914).
84 Lochner v. New York, 198 U.S. 45 (1905).
85 See Adkins v. Children's Hosp. of D.C., 261 U.S. 525 (1923), overruled in part
by W. Coast Hotel Co. v. Parrish, 300 U.S. 379 (noting that “[i]n none of the statutes thus
sustained was the liberty of [the parties] interfered with. Their tendency and purpose was to
prevent unfair, and perhaps fraudulent, methods in the payment of wages.”). See also David
E. Bernstein, Lochner Era Revisionism, Revised: Lochner and the Origins of Fundamental
Rights Constitutionalism, 92 GEO. L.J. 1, 9 (2003); David N. Mayer, Substantive Due
Process Rediscovered: The Rise and Fall of Liberty of Contract, 60 MERCER L. REV. 563,
650 (2009).
86 COMMONS & ANDREWS, supra note 3, at 51.
87 Emilie Tavel, Companies Request Talk With State On Weekly Pay Ruling: 'Ample
Time' Indicated, CHRISTIAN SCI. MONITOR (1908).
88 Id.
89 Am. Mut. Liab. Ins. Co. v. Comm'r of Labor & Indus., 163 N.E.2d 19 (1959).
90 Id. at 22. (“[m]any good reasons may today exist for the payment of wages less often
than weekly, including the greater financial responsibility of most employers, the payment of
family obligations on a monthly basis or better family financial security than existed in years
gone by.” )
91 Tavel, supra note 87.
92 Mass. Gen. Laws ch. 149, § 148 (1990). For the reforming act, see Mass. Gen. Laws
ch. 133, § 502, approved July 20, 1992, and by § 599 made effective as of July 1, 1992.
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

20 Draft[Vol. __
Labor historian Nelson Lichtenstein proposed a more
provocative explanation for the decline of weekly pay. In the 30s, as
part of the New Deal, President Franklin D. Roosevelt introduced the
Federal Insurance Contributions Act (FICA) tax as part of the social
security reform. In 1938, Congress introduced the minimum wage
and the Fair Labor Standards Act (FLSA).93 Then, in 1943, Congress
also introduced the payroll tax, which required employers to withhold
federal income tax from employees’ pay.94 The result was an increased
administrative load on employers who had to compute pay without
computers.95 According to Lichtenstein, the effect of this legislation
was to make weekly pay too expensive, leading to a push to move to
biweekly pay.96 There is a bitter irony here, as legislation that is
ostensibly pro-worker might have had this unanticipated adverse
consequence on pay frequency. The same legislation that guarantees
minimum wage, unemployment insurance, and Medicare may be
inadvertently pushing employees into the hands of payday lenders
and other short-term credit providers.
*
Path-dependence may explain why we still have the payday
today: we are relying on a century-old body of legislation that was
optimized to deal with inferior money and payroll technology.
Defaults tend to become sticky and even the presence of financial
incentive to contract out of them may not be enough to overcome their
viscid pull.97 Being the first-mover to break a social equilibrium
carries risks and costs, and free-riding logic may result in inaction
(consider, again, our dated system of area codes). However compelling
as an explanation, path-dependence is only a weak justification for
the continuation of this practice. Fin-de-siècle labor wars, concerns
with scrip and truck, difficulties of computing wages by hand, and
heavy coin chests carried among work sites – are considerations that
carry little weight in the age of modern payroll and money
technology.98
B. The Synchronization of Bills and the Payday
Another potential reason for the continued existence of the
payday is the seeming alignment of the timing of bill payments and
the payday. Today, households pay most of their bills—utilities, rent,
93 Fair Labor Standards Act, 29 U.S.C. § 201 (2018).
94 Current Tax Payment Act of 1943, Pub. L. No. 68–120, 57 Stat. 126 (June 9, 1943).
95 For proportion, today, roughly 30% of the pay is made through “fringe benefits” which
are often paid to third parties and require a more complex set of computations. BUREAU OF
LAB. STATS., EMPLOYER COSTS FOR EMPLOYEE COMPENSATION HISTORICAL TABLES 2 (2019),
https://www.bls.gov/web/ecec/ececqrtn.pdf [https://perma.cc/9GDV-GVVT].
96 Chris Hayes, The Breakdown: Why Are We Paid Every Two Weeks?, THE NATION (Jan.
21, 2011), https://www.thenation.com/article/breakdown-why-are-we-paid-every-two-
weeks/ [https://perma.cc/3S2D-P5KU].
97 See generally Alan Schwartz & Robert E. Scott, The Common Law of Contract and
the Default Rule Project, 102 VA. L. REV. 1523, 1566–69 (2016).
98 See infra Part I.F.Legislation
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

2020] Payday 21
mortgage, internet, phone, insurance payments, and so on—on a
monthly basis.99 Monthly outlays place the payday into a larger social
equilibrium, with both ingoing and outgoing money streams being
closely tied together. Monthly bill payments, it is worth noting, are a
somewhat recent historical development—a fact that played a role in
the debates over longer pay periods.100
The synchronization of bills and the payday appears, at first
glance, harmonious; like clockwork, money comes in and goes out. But
this is deceptive. Households pay bills for goods and services that they
consume or use throughout the month. Whereas households consume
daily, they only pay monthly.101 This means that the service provider
is not only providing the service, but it is also providing credit: selling
electricity today but receiving payment only at the end of the month.
We see here K attaching again to a primary transaction, the sale of
2
electricity, only that this time around it is the household that borrows
rather than lends.
Economic logic dictates that utility providers charge for this
service and for the risk of default. Households, however, are not the
most reliable borrowers. Some households default on their utility
payments, and the cost borne by all other households is greater for
this reason.102 After all, the provider bears both the cost of not having
access to their earned payments and the risk of default by the
household. Hypothetically, out of every $150 in the electric bill,
perhaps $10 can be seen as interest. Exactly how much households
today pay for this loan is not clear, but the overall economic effect is
likely to be noticeable.103
Consider, then, the situation from the individual’s perspective.
Jane is working all month as a store clerk, but she is paid at the end
of the month. Throughout the month, she needs to consume groceries,
utilities, and other everyday expenses, but her employer will not pay
her until the end of the month. For groceries, she uses her credit
card—paying a few dozen dollars on her revolving balance. For
utilities, she doesn’t need to borrow per se, but she is paying a higher
price, perhaps a dozen more dollars. And while most of her daily
expenses are financed by someone else, she is lending money to her
employer. Somehow, on each transaction, she is on the losing end.
Being a risky borrower, Jane is paying a large amount to the utility
company in implicit interest; being an unsophisticated, under-
99 See e.g., Nevada Public Utilities Commission, http://puc.nv.gov/FAQ/Utility_Bills/
(“Generally, meters are read monthly for electric, natural gas and water services, and monthly
bills are generated for phone services.”)
100 Am. Mut. Liab. Ins. Co. v. Comm'r of Labor & Indus., 163 N.E.2d 19 (1959).
101 Technically, mortgage payments are in arrears, but rent is most often paid in advance.
102 See Residential Energy Consumption Survey (RECS), U.S. ENERGY INFO. ADMIN.
(2015), https://www.eia.gov/consumption/residential/data/2015/hc/php/hc11.1.php
[https://perma.cc/3CWE-VUDT] (roughly 10% of all households received disconnect
notices). The cost of default by some households is then spread to the bills of all other paying
households.
103 The savings from abolishing K2 will be split between the utility providers and the
end-consumer—but the exact split requires a more nuanced analysis of the market and tariff
regulation.
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

22 Draft[Vol. __
capitalized lender, Jane is receiving less in wage premium than her
cost of borrowing.104
Overall, households both borrow and lend, always on worse
terms. Borrowing and lending do not offset each other; instead, they
amplify each other, being two unnecessary and costly credit
transactions. Rather than clockwork, bills and the payday are more
like oarsmen—rowing in opposite directions, only to stay in place.105
As an explanation, the synchronization of bills and pay may make
some sense, syncing income and expense. As a justification, however,
it fails completely; there is no reason to preserve one for the other. If
anything, it would be socially desirable to abolish both. On reflection,
this synchronization seems to be contributing to the path-dependent
pull of historical considerations, making it all the more harder to
imagine breaking away from the biweekly pay convention—although
it makes the case for abolishing the payday more compelling.
C. Employer Power and Lack of Sophistication
Another potential reason for the persistence of K is rooted in
2
the unequal distribution of power and sophistication between
employers and employees. If employers enjoy strong bargaining
power, they may insist on K as a source of cheap credit. And if
2
employees are unsophisticated, they may yield to such demands with
negotiation, not realizing that K is a essentially a credit transaction.
2
In the standard economic model of wages, what determines
wages is marginal productivity—how much value the employee is
producing for the employer.106 A more productive worker would
receive higher wages. In this model, one consistent idea is that of a
wage premium or a “compensating wage differential.”107 If the
104 The loan from the utility company relives some of the liquidity pressure of the
household, but as explained, this is a form of (forced) credit that comes at a cost, albeit implicit
in the price of utilities. .
105 The reasons for K2 in this context are likely to be distinctive from the ones in the
employment context. It is possible that houesholds prefer lump sums outlays because they
allow for easier detection of overcharges or give them power in disputes vis-à-vis the
company. This is a fertile area for future research.
106 PIERRE CAHUC ET AL., LABOR ECONOMICS 82–83 (2d ed. 2014) (exploring, in the
simple model of labor demand, optimal wages).
107 The existence of wage premiums was consistently confirmed. See e.g., Don Fullerton
& Gilbert E. Metcalf, Tax Incidence, NAT’L BUREAU OF ECON. RES. 28 (Mar. 2002),
https://www.nber.org/papers/w8829.pdf [https://perma.cc/KQE7-N5XH] (noting that the
shared incidence of payroll taxes “has been tested and confirmed repeatedly.”); Johnathan
Deslauriers et al., Estimating the Impacts of Payroll Taxes: Evidence from Canadian
Employer-Employee Tax Data, INST. OF LAB. ECON. (2018), http://ftp.iza.org/dp11598.pdf
[https://perma.cc/8VAY-L6Z4] (“The consensus is that [payroll] taxes are partially to
completely shifted to workers, at least in the long run.”). But see Emmanuel Saez, et al.,
Payroll Taxes, Firm Behavior, and Rent Sharing: Evidence from a Young Workers' Tax Cut
in Sweden, 109 AM. ECON. R. 5, 1717 (2017) (finding firm-level differences in incidence and
employment effects of payroll tax cuts). In other contexts, see e.g., John M. Abowd & Orley
Ashenfelter, Anticipated Unemployment, Temporary Layoffs, and Compensating Wage
Differentials, STUDIES IN LABOR MARKETS (1981) (premium for risk of layoffs); John R.
Graham et al., Employee Costs of Corporate Bankruptcy, NAT’L BUREAU OF ECON. RES. 25922
(Jun. 2019) (Bankruptcy risk).
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

2020] Payday 23
employee produces some additional benefit to the employer beyond
his direct labor output, the wage would be adjusted upwards to
include a wage premium. That is, if the employee agrees to receive
payment infrequently, the employer would be willing to pay a wage
premium relative to an employee who is paid frequently.
The size of the wage premium for payday, as well as its very
existence, are empirical questions that were never investigated.
Some complicating factors are market failures, market organization,
and regulation. Now, on theoretical grounds alone, it is clear that
employers will not be willing to offer a wage premium that fully
compensates the worker. To do so, the employer would have to pay
them their costs of lending—but as we just saw, the costs to
employees exceed the benefits to employers.108
Still, if employers do not have to pay a full wage premium, they
may use employee wages as a line of credit. To be able to extract such
a benefit, employers must wield considerable bargaining power. And
while it is clear that many employers do, in fact, wield such power
(think of a single employer in a small town), this surely this does not
describe the entire economy. Outside of monopsonic employers, the
distribution of power is far more heterogeneous. Even middle-class
employees often find themselves in a position to negotiate portions of
their salary and benefits, and firms invest considerably in the
retention efforts of these employees. Yet, we do not find daily pay
common even among these employees.109
A deeper challenge to the asymmetric power explanation lies in
the idea of effective pay. Even supposing that the employer can avoid
paying a wage premium, the employer would find better and worse
ways to exercise its bargaining power. Both the employer and the
employee care about more than the per-hour wage; they care about
the entire package of pay, benefits, work conditions, and duties—that
is, they care about the effective wage. The more benefits the employer
provides, the more costly it becomes to employee workers, even if the
per-hour wage remains the same.
108 See supra Part I.B.
109 Parsons & Van Wesep, supra note 19, 374 (showing that frequency of pay falls with
income, so that middle and high-income workers are paid less frequently than low-income
workers).
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

24 Draft[Vol. __
Now, even for employers who wield enough power that they can
unilaterally dictate the terms of employment, the choice of effective
pay requires some balancing. Set too low, few workers would come to
work and those who do would work fewer hours and leave at the first
opportunity. The profit-maximizing employer would want to offer the
minimal package of pay and benefits that still attracts enough
workers. If the benefits are high, the employer can offer a somewhat
lower wage and still attract enough workers; if the employer cuts
benefits, it would likely have to offer more in the way of pay to attract
the same number of employees. The following figure illustrates this
basic tradeoff:
Figure 2 -- Effective Pay with Different Mixes of Per-Hour Pay and
Pay Frequency
As the Figure illustrates, paying more frequently allows the
employer to pay less per-hour while maintaining the same effective
wage. In designing the optimal mix, the employer would compare its
own costs in providing frequent pay against the savings in direct wage
payments.110 If it is indeed the case that the employee’s cost of
infrequent pay is higher than the employer benefit, even the
asymmetrical powered employer would tend to favor more frequent
pay because it would allow her to reduce paid wages while
maintaining the same effective pay that is needed to retain
employees. Thus, even selfish, dominant employer who is committed
to profit-maximization may find it better to pay less but more
frequently.
While asymmetric power fails to explain the payday
phenomenon in general, there is one area in which it provides a more
110 The value of infrequent pay also includes savings on check-cutting costs and stronger
leverage against the employee, issues that are analyzed infra Section II.D. and II.G.
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

2020] Payday 25
cogent explanation: minimum wage employees.111 Potentially a
design flaw, the minimum wage legislation does not consider effective
pay, only nominal wages. Consider, an employer who—before the
minimum wage—were paying $7 an hour and still attracting enough
workers. Now suppose the legislator requires a minimum wage
payment of $7.25 per hour. If the employer complies and pays more,
the employee’s effective wage is raised above the market clearing
equilibrium. The employer can offset that increase and reduce
effective pay by paying less frequently, thus still keeping compliance
with the letter (but not the spirit) of minimum wage law. Hence, there
is a theoretical possibility that, in the presence of minimum wages,
employers would seek longer payment periods.112 This possibility has
not been investigated in the voluminous literature on the effects of
minimum wages and should be analyzed in future research, because
it is very worrisome.113 Minimum wage employees are also most likely
to suffer low access to liquidity and improving their liquidity should
be an important policy consideration.114
As for lack of sophistication, it may have some explanatory
power, but it does raise some questions. It may be that many
employees lack the financial sophistication to properly classify K as
2
a credit transaction. But what they lack in academic sophistication of
this sort, they have in terms of skin-in-the-game. One does not need
a degree in finance to understand that getting paid every day will
make life easier than waiting a month to be paid. As workers
viscerally feel the consequences of the payday, we would expect them
to gravitate more towards employers who pay regularly. Indeed, one
consequence of the Covid-19 pandemic is the increased demand for
daily pay.115
Overall, employers’ market power and employees lack of
sophistication may explain some part of the practice of payday,
although it seems unpersuasive as a general explanation—especially
given the fact that we find prolonged payment periods even among
111 I emphasize that this is only a possible effect, as the literature on the effects of
minimum wages is complex, nuanced, and hotly-debated. Here I consider the classic wage
model, noting that its applicability in different markets may be limited. See generally David
Neumark, The Employment Effects of Minimum Wages: Some Questions We Need to Answer,
NAT’L BUREAU OF ECON. RES. 23584, at 1 (Oct. 2017), https://www.nber.org/papers/w23584
(“the debate among researchers about the employment effects of minimum wages remains
intense and unsettled.”).
112 There are limitations on the frequency of pay, as discussed infra Section I.F.
113 There is empirical evidence that employers sometimes cut fringe benefits in response
to higher minimum wages. Jeffrey Clemens et al., The Minimum Wage, Fringe Benefits, and
Worker Welfare, NAT’L BUREAU OF ECON. RES. 24635 (May 2018),
https://www.nber.org/papers/w24635.pdf [https://perma.cc/7HGY-ETMK].
114 See e.g., Jonathan Morduch, Poverty and Vulnerability, 84 AMERICAN ECONOMIC
REVIEW, 221, 221 (1994) (noting the “reasonably universal phenomenon by which the lack
of collateral limits borrowing by the poor in bad times”)
115 See Ellen Sheng, Companies Offer Cash-Strapped Employees Daily Pay Cards and
Other Flex-Pay Options as a Lifeline, CNBC (Mar. 30, 2020). As this Article was in edits in
the midst of the pandemic, it is too early to determine its long term labor market effects.
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

26 Draft[Vol. __
higher-paid employees.116 As a justification, however, both reasons
fail. Both information gaps and market monopolies, are types of
market failures—and there is little appeal to market outcomes that
result from market failures.
D. Collateral
A different reason for the existence of the payday grounds the
practice in the need of employers to retain their employees.
Employers worry that employees may decide to quit midstream,
leaving the employer stranded without the necessary personnel or
skill necessary to produce their products or serve their clientele.
Contract law can protect employers against this possibility—they can
require the employee to give notice. But such protection is quite weak,
as employees can be judgment-proof and the cost of litigation can be
prohibitive.117 Postponing pay thus creates collateral and, with it,
leverage; if the employee disappears, the employer may threaten to
expropriate this collateral.118
As an explanation, the idea of collateral faces a challenge in
explaining why the payday is used even when there is little flight risk
or when employees are not judgment proof.
Indeed, the average worker stays with his or her employer for at least
four years.119 It may still be true that employers are reluctant to sue
employees for reputational reasons, rendering the employment
contract unprotective of the employer’s interests. But the same logic,
the same concern with reputational effects, would also lead employers
to avoid sequestering the collateral.120 In any case, collateral offers a
plausible explanation for some of the practices of payday. As a
justification, however, things are more complicated.
As a society, we decided that employers should not be allowed
to sequester earned wages, even when the employee quits. Employers
are legally prohibited from taking earned wages in retaliation for the
116 As of 2013, the average hourly pay per period was $18.6/weekly; $24.8 biweekly;
$29.7 semimonthly; and $28.4 monthly. Matt Burgess, How Frequently Do Private
Businesses Pay Workers?, 3 PAY & BENEFITS 3 (2014).
117 See e.g, 80 Fed. Reg. 62,958 (Oct. 16, 2015) (Private employers
asking, in the context of new rules on commissions, “that DOL permit employers
to withhold a portion of wages as an incentive for the employee to complete the
contract period and to discourage workers from leaving to work in other
industries”). A deeper problem is strategic judgment proofing through asset
shielding. See generally Yonathan A. Arbel, Shielding of Assets and
Lending Contracts, 48 INT'L REV. L. & ECON. 26 (2016)
118 Redmount et al., supra note 69, at 1065.
119 BUREAU OF LAB. STATS., EMPLOYEE TENURE IN 2018 (2018),
https://www.bls.gov/news.release/pdf/tenure.pdf [https://perma.cc/SC4C-ESPD] (based on
the median).
120 Reputational concerns may indeed push the employer to sue and sequester the
collateral to develop a reputation for “toughness.”
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

2020] Payday 27
worker quitting.121 A large number of jurisdictions have enacted
“final pay statutes,” which compel the payment of all unpaid wages
upon termination, or soon thereafter.122 Final pay statutes are often
accompanied by penalties and fee-shifting provisions to further
compel employers to make timely payments.123 State courts have
likewise recognized the public policy imperative in favor of prompt
payment.124 The policy underlying these statutes is widely endorsed:
The Supreme Court held that legislation requiring prompt payment
upon discharge—i.e., payment without “abatement or deduction”—is
constitutional.125 The Department of Labor denounced any pay
practices that have the effect of payment deferral.126 In some
jurisdictions, courts adhere to the “faithless servant” doctrine, which
denies employees any pay (even in quantum meruit) if they are
disloyal to their employers.127 However, disloyalty is generally
understood to mean unlawful competition with the employer or
perhaps dissemination of trade secrets.128 This doctrine is of little
relevance, then, to employees who quit midstream.129
While collateral fails as a justification for withholding earned
wages, it does provide justification to a subtly different issue—not the
existence of payday but why payday is always in arrears. The reason
that employers do not prepay employees is clearly rooted in the
difficulty of recovering unearned wages from an employee who
absconds. Anticipating this difficulty in recovery, some workers may
want to assume positions just for the sake of prepayments, making
the hiring process difficult and costly. Hence, a “reverse” K , where
2
the employer lends money to the employee, is not a general solution—
a point worth remembering as we move to the normative discussion.
Overall, collateral may explain the practice of payday to some
extent and may justify the absence of “reverse K ,” but it fails to
2
121 See, e.g., Britton v. Turner, 6 N.H. 481, 489–92 (1834) (establishing the duty to pay
an employee for part performance); Pineda v. Bank of Am., 241 P.3d 870, 877 (Cal. 2010)
(“[t]he public policy in favor of full and prompt payment of an employee's earned wages is
fundamental and well established…”) (quoting Smith v. Superior Court, 137 P.3d 218, 220
(Cal. 2010)).
122 William C. Martucci & Jennifer K. Oldvader, Avoiding Another Wage-and-Hour
Pitfall:State Late-Paycheck Laws, EMP. REL. TODAY 71 (2011).
123 See e.g., Alaska Stat. § 23.05.140(d), Del. Code Ann. tit. 19, § 1103(b); Va. Code §
40.1-29 (2), (f).
124 See, e.g., Pineda, 241 P.3d at 877 (“[T]he public policy in favor of full and prompt
payment of an employee's earned wages is fundamental and well established…”) (quoting
Smith v. Superior Court, 137 P.3d 218, 220 (Cal. 2010)).
125 St. Louis, Iron Mountain & St. Paul Ry. Co. v. Paul, 173 U.S. 404, 405 (1899).
126 See, e.g., 80 Fed. Reg. 62,957, 62985-86 (Oct. 16, 2015) (to be codified at 20 C.F.R.
pt. 655).
127 See Charles A. Sullivan, Mastering the Faithless Servant?: Reconciling Employment
Law, Contract Law, and Fiduciary Duty, 2011 WIS. L. REV. 777, 779 (2011).
128 See Alan Hyde, What Should the Proposed Restatement of Employment Law Say
About Remedies?, 16 EMP. RTS. & EMP. POL'Y J. 497, 508 (2012).
129 See STEVEN L. WILLBORN ET AL., EMPLOYMENT LAW: CASES AND MATERIALS 609–12
(4th ed. 2007).
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

28 Draft[Vol. __
justify K as a social practice. Employers, we have decided as a
2
society, should not sequester earned wages.
E. Behavioral Biases
The reasons discussed so far were mostly concentrated on the
employer. Another potential reason for the existence of the payday
comes from the employees and their own well-being. While the
question of the payday was mostly neglected in the legal literature,130
a recent theory in financial economics argues that the payday caters
to psychological biases of employees and helps resolve them.131 This
theory was published in the leading Journal of Financial Economics,
, and it is based on a common intuition; still, for the following reasons
I believe it fails as both an explanation and as a justification for the
existence of the payday.
The idea goes as follows: people find it difficult to budget and
control their expenses. When employers pay frequently, individuals
are more likely to spend the money in their pockets due to behavioral
biases such as “present-bias” that prevent them from considering the
full, long-term implications of their behavior. The same way as some
of us would benefit from a pizzeria that would only sell us a few slices,
employees are said to benefit from having infrequent pay. Under this
account, employers are delaying payments as a service to employees,
sparing employees from their weak impulse control.132
This theory is not without evidence. The basic proposition—that
households need help budgeting money—is consistent with some
evidence showing that the timing of payments influences household
money management. One study showed that pension recipients
consume the fewest calories the week before the benefits are paid,
perhaps suggesting a difficulty in saving evenly across the entire pay
period.133 Similarly, another study showed that individuals make the
most of their food and necessity purchases right after receiving
benefit payments.134 The authors and economists Parsons and Van
Wesep further argue that their findings are consistent with the fact
that low-paid employees are paid more frequently than higher-paid
employees. To them, this is simply the result of low-paid employees
being more presently-biased than their wealthier counterparts and
130 See supra note 3.
131 Parsons & Van Wesep, supra note 19.
132 Id, at 374
133 Giovanni Mastrobuoni & Matthew Weinberg, Heterogeneity in Intra-Monthly
Consumption Patterns, Self-Control, and Savings at Retirement, 1 AM. ECON. J.: ECON. POL’Y
163, 164 (2009). See also Jani-Petri Laamanen et al., Once or Twice a Month? The Impact of
Payment Frequency on Consumption Patterns, (2019),
https://pdfs.semanticscholar.org/794a/c54611eeff7bd40efa93729cada5e0e03fa5.pdf
[https://perma.cc/24ND-DYMA].
134 Melvin Stephens Jr., "3rd of tha Month": Do Social Security Recipients Smooth
Consumption Between Checks?, 93 AM. ECON. REV. 406 (2003).
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

2020] Payday 29
lacking a financial buffer, making their need for money exceed their
desire to save.135
There is no doubt that saving money can be difficult, but this
point should not be taken to mean that workers need their employers
to help them save. If that were the case, we would expect to see at
least some workers asking their employers to delay payments—so
this dog doesn’t bark. More generally, this behavioral explanation
fails, both on theoretical and empirical grounds.
First, consider how behavioral biases may work in this context
in the exact opposite direction. By waiting until payday, employees
receive a larger paycheck than they would if they were to be paid on
an ongoing basis. This large payment can create a sense of windfall—
an illusion of plenty.136 This behavioral bias may lead individuals to
spend more on luxuries than when individuals operate under a sense
of scarcity. Indeed, the concern with the illusion of plenty was
precisely the reason some legislators enacted frequent pay legislation:
“[large payments could mean] dissipation on payday of a large part of
the accumulated sums by irresponsible employees with consequent
adverse effect on family and community.”137 Empirical evidence
suggests that this concern is not only theoretical. For example,
research shows that individuals consider tax refunds to be “extra”
money, leading them to spend it more easily than their “regular”
money.138 Similarly, when benefits are paid in a lump sum, one finds
a spike in drug use, hospitalization, and mortality—as some
135 Parsons & Van Wesep, supra note 19, at 389 (“Insofar as education and wealth
correlate negatively with time-inconsistency, more educated and more wealthy workers
should be, and are empirically, paid less frequently.”).
136 See e.g., Hal R. Arkes et al., The Psychology of Windfall Gains, 59 ORGANIZATIONAL
BEHAVIOR & HUMAN DECISION PROCESSES 331 (1994) (finding higher propensity to spend
money viewed as a windfall); C. Yiwei Zhang & Abigail B. Sussman, The Role of Mental
Accounting in Household Spending and Investing Decisions, p. 69, in CLIENT PSYCHOLOGY
(2018)(noting the evidence of a higher propensity to spend windfalls on luxury items). Such
debates are not new; in nineteenth-century Massachusetts, mill owners thought that moving
to weekly pay would lead to more employee intoxication, but “our treasurer determined to
give it a fair trial and the result exceeded our anticipations, for we found that instead of
increasing drunkenness, it has had a contrary effect, so far as we could ascertain by the
working days of our operatives.” Redmount et al., supra note 69, at 1069–70.
137 Am. Mut. Liab. Ins. Co. v. Comm'r of Labor & Indus., 163 N.E.2d 19, 21 (1959). See
also Rebekah D. Provost, Punishing and Deterring the Unknowing: Mandatory Treble
Damages Under the Massachusetts Wage Act, 18 SUFFOLK J. TRIAL & APP. ADVOC. 305, 311
(2013). Payday was a special occasion in turn of the century America, when mostly men
engaged in communal binge drinking, spending a large portion of their payday wages. See
also MADELON POWERS, FACES ALONG THE BAR: LORE AND ORDER IN THE WORKINGMAN'S
SALOON, 1870–1920 52–53 (1998); COMMONS & ANDREWS, supra note 3, at 52 (noting that
some states had special legislation mandating payment during pay hours, to avoid the payment
bar-rooms). In contrast, some legislators expressed concern that too-frequent pay would lead
to “frequent debaucheries . . . .” See Robinson, supra note 75, at 33.
138 See William Adams, Liran Einav & Jonathan Levin, Liquidity Constraints and
Imperfect Information in Subprime Lending, 99 AM. EC. REV. 49, pp. 49-50 (2009) (finding a
sharp increase in auto purchases in the subprime market during tax refund season); Brian
Baug et al., Disentangling Financial Constraints, Precautionary Savings, And Myopia:
Household Behavior Surronding Federal Tax Returns, NBER Working Paper 19783, 2 (2014)
(finding a large temporary increase in expenses following tax refunds, which is interpreted as
suggesting myopic behavior).
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

30 Draft[Vol. __
individuals purchase excess drugs and alcohol.139 One recent report
notes a spike in child abuse on payday, as adults engage in excessive
drinking.140 Another study compared the expenditure profile of
benefits recipients who receive payment twice a month with those
who receive a larger payment once a month. It found that the single
payment leads to high within-month variability, with most of the
money spent early, thus concluding that “two temporally separate
payments might lead to smoother spending than just one
payment.”141 It is also possible that it is easier to save pennies than
dollars, which is the business model of a few recent start-ups.142
Second, to explain why low-income workers are paid more
regularly than higher-income workers, Parsons and Van Wesep posit
that low-paid workers are more prone to present-bias.143 How likely
is this assumption? Are middle-income employees more money
conscious and less likely to overspend than their paycheck-to-
paycheck counterparts?144 And even if that were the case, low-paid
employees are hardly a homogenous or static group. A large body of
research documents earning mobility, suggesting that many (but of
course, not all) employees are on their path to higher earnings in the
future—think interns, students working a side job, or a manager-
track employee working the ranks.145 Of course, pay raises do not
come with an antidote to present bias.
Third, there is a subtle legal point that belies this explanation.
The entire utility of delayed pay is undermined if employees can ask
139 Laamanen et al., supra note 133, at 4.
140 Martin Selsoe Sorensen, Greenland Calls On Denmark to Help Fight Child Sexual
Abuse, N.Y. TIMES, Sept. 27, 2019, at A10,
https://www.nytimes.com/2019/09/27/world/europe/greenland-sexual-abuse-tasiilaq-
denmark.html [https://perma.cc/AFM8-XVUP] (“Pay days are the worst time for the children
of Tasiilaq, . . . With their salaries or social benefits in hand, many adults tend to drink and
parents become too inebriated to look after their children . . . . So on the last Friday of every
month, officials open a sports hall in the district as a shelter to keep children away from sexual
abuse.”).
141 Laamanen et al., supra note 133, at 20.
142 See, e.g., ACORNS, https://www.acorns.com/ [https://perma.cc/MR26-Y7GY] (last
visited Dec. 25, 2019) (a micro-investing platform with corresponding app that allows
customers to invest spare change into an aggregated portfolio managed by industry
professionals). To be clear, I do not consider the windfall bias as necessarily stronger than
myopia—but I note that both are equally plausible forces that operate in opposing directions.
143 Parsons & Van Wesep, supra note 19, at 389 (“Insofar as education and wealth
correlate negatively with time-inconsistency, more educated and more wealthy workers
should be, and are empirically, paid less frequently”)
At least in the aggregate data presented by the Bureau of Labor Statistics, the correlation
between pay and pay frequency breaks if one excludes weekly paid employees—in fact, semi-
monthly paid employees are paid somewhat more than monthly-paid employees. See BUREAU
OF LAB. STAT., supra note 28.
144 Economists are divided on these questions. See Leandro Carvalho, Poverty and Time
Preference, (RAND Lab. & Population, Working Paper No. WR-759, 2010), 2-3
https://www.rand.org/content/dam/rand/pubs/working_papers/2010/RAND_WR759.pdf
[https://perma.cc/CC65-V63F].
145 See, e.g., Katharine Bradbury, Levels and Trends in the Income Mobility of U.S.
Families, 1977–2012, 21 n. 21 (Federal Reserve Bank of Boston, Working Paper No. 16-8,
2016) (“[F]or those starting poor, an average of 58 percent moved out of the poorest group”).
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

2020] Payday 31
employers to advance their wages—and if employees are indeed
present-biased, they would be expected to do so. The authors
themselves admit that wage advances “will cause our results thus far
to unravel, implying a need for regulation.”146 They argue, however,
that the law prevents advances because “regulators in 45 U.S. states
require wages to be paid at a minimum frequency.”147 This is,
however, a mistake. Pay frequency legislation does not require
minimum pay frequency (but a maximum) and, more importantly,
does not bar wage advances.148 The legality of wage advances means
that if workers are indeed blinded by present-bias, they could and
would use wage advances to squander their pay.
The last problem with this explanation is that it overstates the
difficulty of saving money. A recent study concludes that “pay
frequency does not affect household’s savings” and that the amount
of money that households spend over the month has no relation to the
frequency of pay.149 The evidence also suggests that households do not
change their spending categories based on pay frequency. Perhaps it
is because of conflicting behavioral biases, or perhaps it is due to other
reasons—but in practice, the withholding of pay does not significantly
change either saving or consumption patterns. Moreover, while
almost all households are paid infrequently, most households
demonstrably manage liquid assets without squandering them
recklessly. Over 55% of households have liquid assets at their
disposal, thus demonstrating their ability to save and manage money
without a third-party.150 As these households are demonstrably
capable of not wasting the money sitting in their checking and savings
accounts, they do not suffer from such a degree of present-bias that
would make them dependent on the employer’s paternalistic
withholding of cash. Indeed, if this explanation were persuasive, we
might expect to see workers asking their employers to delay
payments, so they can save better—but of course, such behavior is
rarely observed.
Overall, then, while the inability to save may explain a portion of
the payday phenomenon, it fails as a general explanation. However, I
want to make a stronger claim; behavioral biases also fail as a
justification for the payday. To show this, I would like to take a step
back from the question of whether employees need help saving and
focus on the question of whether employers should be the ones who
help them save.
146 Parsons & Van Wesep, supra note 19, at 382.
147 Id.
148 See Jim Hawkins, Earned Wages Access and the End of Payday Lending, 32
(forthcoming Bos. Uni. L. Rev., 2020),
https://papers.ssrn.com/sol3/papers.cfm?abstract_id=3514856 (“Currently, no states
specifically regulate [earned wage advance] transactions”)
149 Inés Berniell, Pay Cycles: Individual and Aggregate Effects of Paycheck Frequency
(Apr. 2019), 19 (unpublished manuscript),
https://inesberniell.weebly.com/uploads/9/1/2/2/91228902/pay_cycles_berniell_in
es.pdf [https://perma.cc/4FH2-LGJ9].
150 See FED. RES. SYS., supra note 5, at 17. Conditional on having financial assets, the
median family held $23,500 in assets. Id. at 18.
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

32 Draft[Vol. __
The core of the problem is simple: employers are unreliable
agents for the management of employee savings.151 There is a reason
why pension funds, such as 401(k)s, are not owned by employers.152
Employers are not some neutral bank; in practice, wage theft—the
withholding of pay due—is “rampant in the low-wage workforce.”153
Employers (and the government is no exception) sometimes
unilaterally suspend pay.154 Moreover, keeping money with one’s
employer also gives the employer leverage, and the employer may
abuse it.155 Worse, unlike banks, employers are not insured against
bankruptcy.156 Thus, using employers as vaults not only exposes
employees to abuse but also to the risk of bankruptcy, a risk over
which the employee has little control.157 Bankruptcy risk also exposes
another problem with employer-side saving. Employers, after all, are
also humans and are inherently not immune to the same present bias
that would lead employees to squander money. The manager may be
tempted to spend the money on a new machine, a shiny business
opportunity, or a private car, not leaving enough slack to pay
wages.158 Given these problems, employees are likely better off facing
their own temptations than dealing with those of their employers. At
the very least, employees would benefit from having reliable, insured
third parties manage their savings (such as their 401(k) retirement
accounts), rather than having their bankruptcy-prone employer
manage them.159
151 A key component of prudent financial planning is diversification. Tying one’s money
with one’s place of employment is the opposite of diversification. See e.g., Sarah O’Brien,
Don’t overlook the risk that comes with your employee stock options, CNBC (Feb, 27, 2018),
https://www.cnbc.com/2018/02/27/employee-stock-options-can-come-with-expensive-
risks.html.
152 The Employee Retirement Income Security Act of 1974 (ERISA) was enacted to
minimize the “looting and mismanagement that had previously plagued private pensions” by
borrowing a trust law model. Natalya Shnitser, Trusts No More: Rethinking the Regulation of
Retirement Savings in the United States, 2016 BYU L. REV. 629, 642 (2016).
153 Llezlie L. Green, Wage Theft in Lawless Courts, 107 CAL. L. REV. 1303, 1309 (2019).
See also IMMIGRANTS’ RIGHTS/INTERNATIONAL HUMAN RIGHTS CLINIC, SETON HALL
UNIVERSITY, ALL WORK AND NO PAY: DAY LABORERS, WAGE THEFT, AND WORKPLACE
JUSTICE IN NEW JERSEY (2011), https://www.immigrationresearch.org/report/seton-hall-
university-school-law/all-work-and-no-pay-day-laborers-wage-theft-and-workplace-ju
[https://perma.cc/3MHZ-S5H7].
154 The concern with suspended government pay is longstanding. See, e.g., Payless
Payday, WASH. POST, Aug. 17, 1949, at 10 (noting that “[y]ear after year, Federal employees
[sic] face suspensions of income”).
155 See supra Part I.D.
156 The Federal Deposit Insurance Act of 1950, Pub. L. 81–797, 64 Stat. 873 (1950)
157 Aside from bankruptcy, letting the employer control more money provides it with
leverage which it can use against the employee in various ways, making quitting, for example,
more difficult. Employers are also less efficient than financial institutions at making
periodical payments.
158 Even without present-bias, large debts can exacerbate risk taking by managers. See
e.g., Zhiyao Chen and Ran Duchin, Do Nonfinancial Firms Use Financial Assets to Take
Risk? At 2 (May 1, 2019). Available at SSRN: https://ssrn.com/abstract=3284205 or
http://dx.doi.org/10.2139/ssrn.3284205 (“A vast body of theoretical work predicts that firms
will invest in riskier projects as they become distressed)
159 The authors foresee this objection but dismiss it: “[I]t is not particularly important
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

2020] Payday 33
Besides this core problem, one must also consider that
withholding pay from employees is a particularly severe form of
paternalism. Proposing to withhold property from individuals
because one thinks they are insufficiently responsible to handle it is
a very strong claim that would require very strong evidence. But as
noted, the evidence suggests otherwise.160 In fact, there are some
deeply disturbing stories of how the larger lump sum payments, due
to infrequent pay, result in substance abuse.161
*
This section tackled the argument that the payday serves
employees by helping them budget their own money. It showed why
this intuitive idea fails as an explanation—among other things, it
neglects to consider how larger paychecks can invite excess spending.
More critically, this section argued that this theory also fails as a
justification for the payday; this type of paternalism requires an
excessive degree of trust in employers. Thus, whatever limited
explanatory power this theory has, it is insufficient to justify this
practice.
F. Legislation
Employment law is highly regulated at both the federal and
state level, and the payday is no exception. As this section shows, the
payday is affected by both federal and state legislation in ways both
visible and invisible.
Legislation provides the most direct explanation of the payday
in the public sector. As a result of extensive pay regulation, most
public employees are paid on a biweekly or a semimonthly basis.
Federal legislation sets a biweekly pay period.162 State laws,
similarly, will often set a biweekly or a semimonthly pay schedule for
state employees.163 Local governments also pay twice a month. Of the
200 largest cities in the United States, 189 (94.5%) pay on a biweekly
who conducts the timing-welfare calculation, as long as someone does.” Parsons & Van
Wesep, supra note 19, at 383. The fact that, despite the considerable risks, employers are the
ones who supposedly save for employees is too important to be casually dismissed.
160 See supra Part I.E
161 Martin Selsoe Sorensen, Greenland Calls On Denmark to Help Fight Child Sexual
Abuse, N.Y. Times (Sep. 27, 2019) (“Pay days are the worst time for the children of Tasiilaq,
… With their salaries or social benefits in hand, many adults tend to drink and parents become
too inebriated to look after their children . . . That’s when an already high rate of sexual abuse
rises”)
162 5 U.S.C. § 5504(a) (2018) (“The pay period for an employee covers two
administrative workweeks.”). See also U.S. BUREAU OF THE CENSUS, STATISTICAL ABSTRACT
OF THE UNITED STATES 1986, 302 (1986),
https://www.google.com/books/edition/Statistical_Abstract_of_the_United_State/R7M_0H
XXZ48C?hl=en&gbpv=1&dq=biweekly+payroll&pg=PA302&printsec=frontcover
[hereinafter 1986 ABSTRACT] (noting that ““most Federal employees are paid on a biweekly
basis.”)
163 ALA. CODE § 36-6-1 (2020) (semi-monthly); N.Y. STATE FIN. LAW § 200(1)
(McKinney 2018) (biweekly); WASH. ADMIN. CODE § 82-50-021 (2020) (semi-monthly).
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

34 Draft[Vol. __
or semimonthly basis.164 With 22 million Americans employed as
government employees,165 we thus find legislation to be a direct
explanation for pay practices in this sector.
As for the private sector, the analysis is far more nuanced. In
the private sector, no law sets pay frequency directly. Instead, states
set payment frequency floors—the requirement that the employer
won’t delay pay for longer than, normally, two weeks. It is possible
that the private sector simply imitates pay practices in the public
sector, but this possibility seems weak, given the stakes involved. If
employers could set lower pay with more frequent pay, then the
analysis above suggests that it would be profitable for them to do so.
To the extent that debt spirals also affect worker productivity,
stability, and reliability, we would expect the private sector to be
somewhat responsive to such pressures.166
Indirectly, however, legislation inadvertently incentivizes
infrequent pay. Ironically, it is mostly pro-worker legislation that
promotes infrequent pay. Legal scholars, however, have failed to note
and grapple with this complexity.
Take minimum wage laws. We have already seen these laws fail
to regulate the interaction of minimum wage and pay frequency—
thus, allowing employers to cut back on pay frequency without
violating the law.167 Another problem emerges in the context of tipped
and commission-based employees—an important part of the
workforce, with approximately 4.3 million tipped workers in the
United States.168 For these employees, the Fair Labor Standards Act
(FLSA) permits employers to pay below minimum wage, so long as
the lower wage plus tips averages to the minimum wage over the pay
period.169 As a result, employers are induced to set a long pay period,
164 CITY OF CHESAPEAKE, VA AUDIT SERVICES DEP’T, CITY PAYROLL CYCLES SPECIAL
AUDIT App. B (2019),
http://www.cityofchesapeake.net/Assets/documents/departments/aud
it/department-audits/Payroll+Cycle+Full+Report.pdf
[https://perma.cc/NVH4-4XMR].
165 Employment, Hours, and Earnings from the Current Employment Statistics survey
(National), U.S. BUREAU OF LAB. STATS. (Apr. 18, 2020),
https://data.bls.gov/timeseries/CES9000000001 [https://perma.cc/8N8F-A2DS] (last visited
Feb. 13, 2020).
166 A set of economic arguments involve nominal rigidity of wages or “sticky wages”—
the failure of payments to adjust, mostly downward, to changing market conditions. This may
further explain the pattern of biweekly pay, although even sticky wages are thought to adjust
in the long run. See generally Alessandro Barattieri et al., Some Evidence on the Importance
of Sticky Wages, 6 AM. ECON. J.: MACROECON. 70 (2014).
167 See supra Part I.C.
168 SYLVIA A. ALLEGRETTO & DAVID COOPER, ECON. POL’Y INST., TWENTY-THREE
YEARS AND STILL WAITING FOR CHANGE: WHY IT’S TIME TO GIVE TIPPED WORKERS THE
REGULAR MINIMUM WAGE 23 (2014), https://www.epi.org/files/2014/EPI-CWED-BP379.pdf
[https://perma.cc/ZSH7-LFZ6]. This estimate does not cover commission-based employees.
169 Fair Labor Standards Act, 29 U.S.C. §203(m) (2018); 48B AM. JUR. 2D Labor and
Labor Relations § 3108 (2020). See also DEP’T OF LABOR, FIELD OPERATIONS HANDBOOK
CHAPTER 30, RECORDS, MINIMUM WAGE, AND PAYMENT OF WAGES 30b01(b)(1) (2016),
https://www.dol.gov/whd/FOH/FOH_Ch30.pdf (“[T]he salary is sufficient to meet the
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

2020] Payday 35
so as to average daily variations in pay, as the following example
demonstrates.
Suppose that an employee makes $1,300 in tips in one week and
$100 in the next. The average is $700—well beyond the biweekly
federal minimum of $580—so the employer need not pay the employee
any extra amount.170 But what happens if the pay period is shorter?
Suppose an employer instead paid on a weekly basis. The employee
makes $1,300 the first week, well beyond the minimum wage, so the
employer would again not need to compensate the employee. But in
the second week, the employee only makes $100, well below the
weekly minimum wage of $290. By making the payday shorter, the
employer now has to pay the employee an extra $190. As this example
demonstrates, for tipped and commissioned employees, minimum
wage legislation unwittingly incentivizes longer pay periods.
Admittedly, for many tipped employees this harm is mitigated by the
common practice of paying cash tips daily—but as many tips are paid
on credit, this problem remains important.171
Overtime legislation presents a similar averaging problem,
although to a lesser extent. If overtime obligations are tied to pay
frequency, employers would want to prolong pay periods to smooth
periods of high work. In a biweekly pay period, the employer could
avoid paying overtime in the first week if there is less work in the
second week. The Department of Labor and many courts have taken
the view that employers are not allowed to average over more than
one week and that overtime legislation is done on a weekly basis.172
Some courts, however, has taken a different approach, as explained
by Judge Easterbrook: “[I] is unlikely that Congress meant to require
employers to pay overtime in the lean weeks when the fat weeks more
than make up.”173 To the extent that employers are allowed to average
minimum wage requirements for all hours worked during the pay period . . . .”); id. at
30b05(c)(1) (“There is no requirement that wages be paid weekly, as long as some regular
pay period (such as biweekly or monthly) is established . . . . The only requirement is that
employees receive prompt payment of the minimum wage covering all hours worked during
the pay period.”). see CAL. CODE REGS. tit. 8 § 11010(4)(B) (2020) (“[M]inimum wage for all
hours worked in the payroll period”)
170 The employer still owes the employee the federal minimum wage per hour for a tipped
employee of $2.13/hr, which is not included in the example for the sake of clarity. Fair Labor
Standards Act, 29 U.S.C. §203(m) (2018)
171 See MONEYTIPS, Should You Use Your Credit Card to Tip?, CBS NEWS (Jun. 16,
2015, 1:34 PM), https://www.cbsnews.com/news/should-you-use-your-credit-card-to-tip/
[https://perma.cc/64EZ-ELUH] (noting the time lag associated with credit card tips).
172 29 C.F.R. § 778.104 (2020) (“[FLSA] takes a single workweek as its standard and
does not permit averaging of hours”); Overtime Frequently Asked Questions, N.Y. DEP’T. OF
LAB., https://www.labor.ny.gov/legal/counsel/pdf/overtime-frequently-
asked-questions.pdf [https://perma.cc/8YCB-QBPT] (last visited Dec. 25, 2019). See
also Fernandez v. Centerplate/NBSE, 441 F.3d 1006, 1007 (D.C. Cir. 2006) (“FLSA
requires employers to pay overtime compensation for time worked in excess of forty hours
per week, but not for time worked in excess of eight hours per day”); Freixa v. Prestige
Cruise Servs., LLC, 853 F.3d 1344, 1346 (11th Cir. 2017).
173 Walton v. United Consumers Club, Inc., 786 F.2d 303, 307 (7th Cir. 1986). See also
Triple "AAA" Co. v. Wirtz, 378 F.2d 884, 887 (10th Cir. 1967) (allowing for averaging over
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

36 Draft[Vol. __
pay over pay periods, they would have an incentive to prolong that
period.
A much deeper problem with overtime legislation concerns the
definition of salaried employees, who are a large minority of the
working population.174 A salary is a fixed payment that does not
depend on actual hours worked.175 The FLSA permits employers to
avoid paying overtime to salaried employees.176 Because employers
might abuse this system by designating employees as salaried
employees, the FLSA sets clear criteria as to which class of workers
are exempt from overtime obligations, the “exempt” employee.177 For
example, if a worker is docked pay for working fewer hours, then the
employer can no longer claim that the worker is exempt from
overtime pay.178
The problem is that the FLSA also imposes a formal test, the
“salary basis test”: an employee cannot be considered salaried “if the
employee regularly receives each pay period on a weekly, or less
frequent basis.”179 This test links pay frequency and pay status, and
it leads to the absurd result that an employer who pays employees
frequently will also have to pay overtime, whereas an employer who
chooses infrequent pay can also avoid overtime pay.180 This outcome
directly contradicts FLSA’s purpose to protect workers from “labor
conditions [that are] detrimental to the maintenance of the minimum
standard of living necessary for [the] health, efficiency, and general
well-being of workers.”181 By tying pay frequency to legal protections,
the law deters employers from paying employees daily, lest they be
considered unsalaried.
a year); Forster v. Smartstream, Inc., No. 3:13-CV-866-J-PDB, 2016 WL 70605, at *6 (M.D.
Fla. Jan. 6, 2016); Schwind v. EW & Assocs., Inc., 371 F. Supp. 2d 560, 568 (S.D.N.Y. 2005);
Gatto v. Mortg. Specialists of Ill., Inc., 442 F. Supp. 2d 529, 542 (N.D. Ill. 2006).
174 BUREAU OF LAB. STATS., CHARACTERISTICS OF MINIMUM WAGE WORKERS 2018 (Mar.
2019), https://www.bls.gov/opub/reports/minimum-wage/2018/home.htm
[https://perma.cc/3PBD-SQYE].
175 Garrett Reid Krueger, Straight-Time Overtime and Salary Basis: Reform of the Fair
Labor Standards Act, 70 WASH. L. REV. 1097, 1103 (1995) (“Typically, salaried employees
do not ‘punch a clock,’ are not paid by the hour, and are not docked pay if they do not work
forty hours in a given week.”).
176 Fair Labor Standards Act, 29 U.S.C. § 213(a)(1) (2018).
177 Robert L. Levin, Salaried or Hourly: Do Your Exempt Employees Meet the “Salary
Test” Under the FLSA?, 11 LAB. LAW. 25, 25 (1995). When employers pay employees who
work in a “fluctuating workweek” arrangement, they need to pay only one half of the regular
hourly rate, rather than 1.5 of that rate. The hourly rate, oddly, is lower the more overtime
hours the employee clocks, a practice approved by the courts. See generally C.W. Von
Bergen, Using the Fluctuating Workweek Compensation Method to Reduce Overtime
Expenses in Public Organizations, 40 PUB. PERSONNEL MGMT. 165 (2011).
178 Brock v. The Claridge Hotel & Casino, 846 F.2d 180, 184–85 (3d Cir. 1988).
179 29 C.F.R. § 541.602 (2018).
180 To the best of my knowledge, this topic was never litigated, so it is an open question
how the courts would rule. My conversations with practitioners suggest a general belief that
courts would be willing to divorce pay frequency from the actual definition of salaried
employees, although given the plain language of the text, it is unclear how they would reach
this outcome.
181 29 U.S.C. § 202(a) (2018).
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

2020] Payday 37
In these various ways, legislation explains pay practices. In the
public sector, the explanation is simple fiat; but why not pay public
sector employees more often? The low return on treasury bonds shows
that the government can easily borrow at low rates.182 Private
employers may conform to public sector standards and are, in any
case, incentivized to delay payments because of well-intentioned but
poorly-drafted legislation. And while fiat and bad legislative design
may explain the payday, it does not justify it.
G. Check Cutting Costs
Paying workers is expensive. This section explores the various
costs involved in paying workers, and highlights how these costs can
be an obstacle to regular pay. Schematically, paying involves four
different stages: (1) determining pay due; (2) calculating withholding
for compliance purposes; (3) transferring payments; and (4) receiving
payments. The first two stages involve costs that are affected by
payroll technology; the latter two involve costs due to money
technology.
The first cost is that of the determination of payment due. This
is mostly a technological problem, and it has largely been resolved.
Determining due pay for salaried employees is almost trivial in
modern times.183 For other types of employees, the determination
may be somewhat more complex—but not by a large margin. The
employee time clock was patented in 1891,184 and with the broad
integration of computers and mobile devices in the modern workplace,
most time-tracking today is automated.185 True, employers want to
verify every reported work hour—a task that does not scale up well,
however, this difficulty is inherent to the employment relationship for
reasons other than pay frequency, and, as we shall see, this concern
can be effectively resolved with careful design of pay obligations.186
A seemingly more serious cost is compliance. Even after
assessing the employees’ wages, the employer must still verify that it
is properly assessing compulsory and voluntary deductions, that
levies are effectively put aside, that child support and alimony
payments are correctly computed, and that any wage garnishments
182 Daily Treasury Yield Curve Rates, DEP’T OF TREASURY,
https://www.treasury.gov/resource-center/data-chart-center/interest-
rates/pages/textview.aspx?data=yield [https://perma.cc/RT6G-6CGB] (last visited Feb. 13,
2020).
183 The biweekly pay is simply given by dividing the annual salary by 24 for a semi-
monthly paid employee. Withholdings and deductions can complicate the calculus, but with
payroll software, these issues are generally easily resolved in practice.
184 U.S. Patent No. 452,894 (issued May 26, 1891).
185 See https://www.tsheets.com/resources/time-tracking-survey (finding in a survey of
954 employees that only 25% track time with paper or a timesheet).
Roughly 2.9% of US employees are reportedly working remotely at least half of the time,
requiring alternative arrangements (such as salary or software tracking). Brie Weiler
Reynolds, The State of the Remote Job Marketplace, FLEXJOBS (Mar. 27, 2018),
https://www.flexjobs.com/blog/post/state-of-the-remote-job-marketplace/
[https://perma.cc/GNQ8-M55S].
186 See infra Section IV.A.
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

38 Draft[Vol. __
are deducted. Then, the employer must verify compliance with all
minimum wage and overtime legislation. Finally, the employer must
keep a record of hours worked and communicate this information to
the employee. These challenges may have been enormous in the past,
as properly computing withholding manually is a long, arduous, and
error-prone process. But today, none of these challenges are especially
onerous with the advent of the modern computer and payroll
software. The per payroll cost of paying an employee in medium-sized
companies appears to be between $1 and $5, although companies
differ significantly in their pricing methodologies.187 Completing a
payroll “run” may involve a real cost, but this cost is no longer
prohibitive.188
Despite the availability of software, employers still want to verify
the accuracy of all payments, because failure to comply can result in
criminal, civil, and ethical sanctions. The FLSA, for example, imposes
criminal fines and even imprisonment for failures to comply.189 This
liability also extends to corporate officers.190 The consequences can
also be disciplinary for some professionals. One lawyer was put on
probation for eighteen months for failure to file and pay various
federal, state, and local payroll tax obligations on a timely basis.191
The FLSA also includes a civil sanction: failure to pay wages can
result in liquidated damages equal to all unpaid wages192 and
attorney fees.193 Given the costs of mistakes, the employer will want
to include safeguards—such as manual revision of at least some of the
paystubs. Under the current system, these safeguards should be
187 See Real-world Payroll Services Prices From BuyerZone Buyers, BUYERZONE,
https://www.buyerzone.com/hr-personnel/payroll-services/ar-prices-payroll-large/
[https://perma.cc/HX32-TFCW] (last visited Dec. 25, 2019); See also Easy,
modern payroll starting at just $45, $25/month, GUSTO, https://gusto.com/product/pricing
[https://perma.cc/CX2L-8XWL] (last visited Dec. 25, 2019); Our Prices,
CORPORATE PAYROLL SERVICES, https://www.corpay.com/pricing/
[https://perma.cc/7EKG-C8N3]
(last visited Dec. 25, 2019); Online Payroll Service Prices, PRICE IT THERE,
https://priceithere.com/online-payroll-cost/ [https://perma.cc/994N-6HPD]
(last visited Dec. 25, 2019).
188 As I discuss later, the compliance cost would remain largely the same under my
proposal, because the verification process will only take place once every two weeks. See
infra Part I.A.
189 29 U.S.C. § 216(a) (2018) (setting a fine of up to $10,000 and imprisonment of up to
six months for willful violations).
190 Erwin v. United States, 591 F.3d 313, 320 (4th Cir. 2010); Hina Shah, Broadening
Low-Wage Workers’ Access to Justice: Guaranteeing Unpaid Wages in Targeted Industries,
28 HOFSTRA LAB. & EMPL. L. J. 9, 30 (2010).
191 In re Finestrauss, 32 A.3d 978, 979 (Del. 2011).
192 See 29 U.S.C. § 216(b) (2018). The liquidated damages can be reduced if the
employer shows good-faith and reasonable grounds for underpayment, 29 U.S.C. § 260
(2018). The standard norm, however, is double damages.See Kinney v. D.C., 994 F.2d 6, 12
(D.C. Cir. 1993) (citing Walton v. United Consumers Club, Inc., 786 F.2d 303, 310 (7th Cir.
1986)).
193 29 U.S.C. § 216(b) (2018) (“The court in such action shall, in addition to any
judgment awarded to the plaintiff or plaintiffs, allow a reasonable attorney’s fee to be paid by
the defendant, and costs of the action.”).
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

2020] Payday 39
employed at every pay cycle, and because they do not scale well,
increasing the pay frequency can drastically increase costs.
Illustrative was the momentary expression of horror when, in an
interview with a payroll director for a large organization, I mentioned
the possibility of moving to a daily payday.194
Overall, payroll technology is sufficiently mature to resolve the
basic aspect of calculating pay; however, an outstanding issue is the
problem of verification and compliance. These processes do not scale
well and become increasingly costly with higher-frequency pay.
Moving to money technology, for most employers and employees,
transferring money is a largely invisible process. Roughly 87% of
households are paid using direct deposit,195 a money transfer
technology that involves the Automated Clearing House (ACH)
system. Normally, there are no charges on the employee side; but
employers are charged roughly $0.26–$0.50 per transfer.196
Employers also incur an additional administrative cost (in terms of
personnel and IT) of $0.11–$0.25, suggesting a total cost of $0.37–
$0.75 per single employee payment for one pay period.197 These costs
are not substantial by themselves, although moving from biweekly to
daily payments can increase costs by $5.18–$10.5 per two weeks.198
Even for a minimum wage employee, this is roughly the cost of
another hour of work—a real, but not prohibitive, cost.
The problem is the “Other America.”199 In 2017, 14.1 million
adults were unbanked, meaning they did not have either a checking
or a savings account.200 In addition, 48.9 million were “underbanked,”
i.e., they were using non-banks for financial products (such as check
cashing, payday lending, or money orders) despite having a bank
account.201 As a consequence, 27.6% of households receive some of
their payments in a paper check or a money order, and 7.9% receive
payments in cash.202 The under- and unbanked are also poorer on
average.203
194 Interview with anonymous payroll director, (Nov. 19, 2019) (details on file with
author).
195 FED. DEPOSIT INS. CORP., FDIC NATIONAL SURVEY OF UNBANKED AND
UNDERBANKED HOUSEHOLDS 12 (2017),
https://www.fdic.gov/householdsurvey/2017/2017report.pdf [https://perma.cc/239V-XYEF].
196 Payment Cost Benchmarking Survey, ASS’N FOR FIN. PROF’LS. 8-9 (2015),
https://www.bottomline.com/application/files/faster-cost-effective-afp-payments-cost-
benchmark-survey-gen-us-srr-1510.pdf [https://perma.cc/RXY7-26EF].
197 id.
198 The range of total costs per-transfer is $0.37-0.75. Moving from biweekly to daily
pay multiplies the number of transactions by (at most) 14, giving the estimate above.
In addition, ACH only recently (in 2016) started moving to a same-day process, a
transition that is still on going.
199 MICHAEL HARRINGTON, THE OTHER AMERICA (1962) (documenting the spread of
poverty in the U.S.).
200 FDIC National Survey of Unbanked and Underbanked Households, supra note 195,
at 1.
201 Id.
202 Id. at 12. Note that the percentages do not add up to 100% as households may be paid
in more than one method.
203 See https://www.fdic.gov/householdsurvey/2012_unbankedreport.pdf at 18.
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

40 Draft[Vol. __
Employees not paid via direct deposit are mostly paid by check or
money orders—two dated, lengthy, unreliable, and expensive money
technologies. For the employer, the simple cost of writing a check is
estimated at $4 per check.204 Checks are also physical objects, which
add friction and costs related to security and delivery. Even the
delivery of checks is unreliable; one employee described her
experience: “the checks . . . were delivered by oft-delayed trucks that,
living paycheck-to-paycheck, sometimes left her family in dire
financial straits.”205
Checks must be cashed somehow, and cash checking services
flourish around the nation.206 These services offer immediate money
for checks, but because checks are such a slow and unreliable
technology, these businesses assume a considerable risk for their
services.207 A check can be easily forged and, even if authentic, can
still bounce. Cash checking services provide a real service, but they
charge high rates. One study reports a range of 1.5%–3.3% of the
check’s face value.208 This means that, on average, there is a cost of
$40 per paycheck for typical households with full-time workers to even
access their earned money.209 If used regularly over one’s career, the
household will spend $41,600 in fees—money that could otherwise be
used to build wealth for retirement.210 Indeed, some of these fees are
avoidable, by cashing the check at the bank of issue (i.e., the
employer’s bank), but this involves the time and cost of travel to the
bank.211 Getting to the location, safely carrying the check, and
204 Vipal Monga, U.S. Companies Cling to Writing Paper Checks, Wall
Street J. (Mar. 10, 2014). Another estimate suggests a per-check cost of $1.22,
see ADAM J. LEVITIN, CONSUMER FINANCE LAW: MARKETS AND REGULATION,
at 348 (2018).
205 Michael M. Oswalt & César F. Rosado Marzán, Organizing the State: The "New
Labor Law" Seen from the Bottom-Up, 39 BERKELEY J. EMP. & LAB. L. 415, 453 (2018). In
1908, a similar report emerged: the “pay car” containing payments was delayed, suspending
all of the employees monthly pay. Missouri Pacific Postpones Payday, Gives No Reason,
ST. LOUIS. POST, Feb. 17, 1908.
206 Check cashing is not unique to the unbanked. See Michael S. Barr, Banking the Poor,
21 YALE J. ON REG. 121, 144 (2004).
207 Beware of Fake Checks, FED. DEPOSIT INS. CORP. CONSUMER NEWS (Aug., 26, 2019),
https://www.fdic.gov/consumers/consumer/news/august2019.pdf [https://perma.cc/EA3N-
ZUGE].
Anecdotally, online users report their experience running a check cashing service as
“Check Cashing Business is a Big NO. . . . it will ruin you.” V. Sheth, Answer to “Does check
cashing business bring good profit?, QUORA (Jun. 20, 2019), https://www.quora.com/Does-
check-cashing-business-bring-good-profit [https://perma.cc/H6E9-T8EJ].
208 Barr, supra note 206, at 146–47. See also Robin A. Prager, Determinants of the
Locations of Alternative Financial Service Providers, 45 REV. INDUS. ORG. 21, 24 (2014).
209 Matt Fellowes & Mia Mabanta, Banking on Wealth: America’s New Retail Banking
Infrastructure and its Wealth-Building Potential, BROOKINGS 3 (2008),
https://www.brookings.edu/wp-content/uploads/2016/06/01_banking_fellowes.pdf.
[https://perma.cc/49G6-F7JC].
210 Id. at 14.
211 Barr, supra note 206, at 145 (“[A] large portion of the unbanked manage to avoid
paying high costs for at least some of their financial services.”).
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

2020] Payday 41
waiting in line are non-trivial costs; especially since paydays tend to
be synchronized, leading to congestion.212
Finally, the use of cash presents its own difficulties. Roughly 8%
of households are paid in cash.213 Paying with cash requires carrying
large amounts, which involves administrative overhead. More
importantly, the perfect liquidity of cash invites theft risk, both for
the employer and for the employee. Carrying large amounts of cash
exposes one to risks, and there is little wonder why most people prefer
to carry small amounts of cash on their person.
In conclusion, while money technology has improved
dramatically over the last century, many employees are still being
paid using dated technologies—checks and cash. These dated
payment methods impose significant costs, making daily payment
prohibitively expensive. While digital money exists and offers
important efficiencies, it still has to overcome the under-banking gap
and other issues of implementation.
III. ABOLISHING THE PAYDAY
The payday is a common feature of employment contracts. The
payday implicates a credit transaction (K ), but this credit
2
transaction is not motivated by the logic of credit. Instead, the
investigation of this practice suggests that it owes a large part of its
vitality to outdated legislation and money technology. Even the most
sympathetic justifications of the payday—those which are rooted in
employee psychology—still leave the current arrangements in a poor
light.
The goal of this Part is to explain why abolishing the payday
and moving to daily streams of payment is critical, valuable, and more
effective than some intermediate solutions that are currently being
proposed. If abolition initially strikes one as radical, recall that in the
19th century, weekly payment systems were already in place214—
during a time in which one had to do all calculations by hand and
transport a chest with coins between work locations.215 Daily
payments are well within reach today.
A. The Stakes of Abolishing the Payday
Suspending for a moment the how, let us consider the
implications of abolishing the payday and moving to daily streams of
payment.
212 The concern with congestion is a longstanding one. See e.g., Congested Payday,
WASH. POST (May 13, 1941) (“Residents of Washington . . . always know when payday
arrives. For twice each month they are subjected to major and minor inconveniences”). At
one point, President Roosevelt ordered the spreading of payday to 20 days for this reason.
Federal Paydays To Be Increased: WASHINGTON POST (Oct. 17, 1942) (On file).
213 Beware of Fake Checks, supra note 207.
214 See supra note 65.
215 See Redmount et al, supra note 69, at 1096.. xx
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

42 Draft[Vol. __
In the first instance, the current biweekly payday harms
workers. True, paying employees more frequently will not make
households wealthier; but it will make them more capable of meeting
life challenges as they come. Over the last few years, interest rates
were at a historic low; but the stakes of abolishing payday will only
increase if interest rates revert to their historical rates.216 Lack of
liquidity is associated with a variety of negative health outcomes.217
Abolishing the payday would help ease some of this pressure. The
stress of thinking about how to pay for groceries the next day,
whether one should skip the next dentist appointment, or the
arguments with one’s partner can be alleviated with greater control
over one’s finances. Indeed, the records from the nineteenth-century
move to weekly payments suggest a marked increase in reported
employee well-being.218
Greater liquidity also allows one to seize opportunities as they
present themselves. Some of these opportunities are humdrum,
although consequential for one’s financial health, like buying
discounted items in bulk. Other opportunities can have even larger
effects, like buying a ticket to fly out to an interview with another
employer. It is perhaps natural for a well-off reader to discount the
difficulty insufficient liquidity imposes on life choices, but even the
cost of dry cleaning or a haircut can prevent some from attending a
job interview.219
One potential negative aspect of abolishing payday is that it will
restrict credit access to businesses. Firms today borrow at cheap rates
through the withholding of pay and abolishing the payday might limit
their access to credit, especially if the firm is a small business. This
issue should not be overstated. Worker wages should not be an open
line of credit. When the firm taps into this source of credit, it exposes
workers to the risk of its own bankruptcy and it imposes on them the
costs associated with low liquidity. While Walmart enjoys the float
from withholding pay, the costs endured by its employees far exceed
this benefit.
Small businesses are often under more severe credit pressures,
and for many of them, access to credit is even more essential.220 This
consideration, however, does not mean that workers wages should be
the solution. In fact, it may suggest more caution with exposing
workers’ wages to business risks. If the small business is over-
216 MACROTRENDS, supra note 24.
217 See Lorraine T. Dean & Lauren Hersch Nicholas, Using Credit Scores to Understand
Predictors and Consequences of Disease, 108 Am. J. Public Health 1503, 1504 (2018).
218 Redmount et al., supra note 69, at 1083.
219 A Good Impression Begins With the Way you Dress, JAILS TO JOBS
https://jailstojobs.org/free-interview-clothes/ [https://perma.cc/62LC-8C8F]
(last visited Dec. 24, 2019); Amelia Ward, Barber Gives Homeless Free Haircuts and
Trains Them to Work In His Shop, LAD BIBLE (Dec. 21, 2019),
https://www.ladbible.com/news/daily-ladness-barber-gives-homeless-people-free-haircuts-
and-trains-them-up-20191220 [https://perma.cc/5PR4-W82D].
220 Claire Kramer Mills et al., Growing Pains: Examining Small Business Access to
Affordable Credit in Low-Income Areas, 2019 CONSUMER & COMMUNITY CONTEXT 22, 23-24
(2019)
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

2020] Payday 43
extended, using unpaid wages to finance operations jeopardizes
workers. As a society, we face a basic choice as to who should be the
source of liquidity to small businesses—should it be workers with
their salaries or sophisticated credit markets, which are capable of
evaluating, monitoring, and pricing risk. Keeping the payday to
finance business operations is a policy choice that is available to us—
but it appears a bad one: workers should not be in the business of
lending money to their employers.
Another related negative consequence of abolishing the payday
is the elimination of the wage premium associated with it. As
discussed, one might expect a wage premium for longer paydays for
two reasons. First, the employer receives the benefit of holding (and
using) the money until the payday, and, second, the employer saves
the costs of making more regular payments. Now, for those who think
that employees do not receive a meaningful wage premium today for
K loans, this consequence is largely irrelevant. Even for those who
2
believe that there may be a wage effect, there is some reason to doubt
its magnitude, if not its existence. The single study that evaluated the
effect of moving from the monthly payday to the weekly payday—
while admittedly dated and incomplete—found that this move
actually led to an increase in the effective pay and well-being of
employees.221 This is, in part, because workers chose to work more
when pay was more frequent (what economists call a ‘income
effect’).222 This finding should not be overstated because of various
methodological and data issues, but it at least suggests that the
effects of abolishing the payday may be more nuanced than what
appears at first sight.
Whatever the case might be about the wage premium, daily pay
would also have strong positive effects. Most directly, more frequent
pay would remove workers from the unnatural position of lending
money to their employers. The employer’s benefit from retaining this
money is more than offset by the worker’s need for the money. In a
very early decision, the Supreme Court clearly recognized this point:
“[t]here [is] certainly . . . advantage to those who work for a living of
a ready purchasing power for their needs over the use of credit.”223
The lack of purchasing power manifests itself in many ways—most
painfully, in the cost of short-credit solutions. The average American
has $5,673 in revolving credit card debt,224 on which they pay 16%
APR ($580 per year, roughly).225 Credit cards appear cheap relative
to the burgeoning installment loans industry, which charges an
221 Redmount et al., supra note 69, at 1083.
222 Jim Chappelow, Income Effect, INVESTOPEDIA (May, 28, 2020).
223 Erie R.R. Co. v. Williams, 233 U.S. 685, 704 (1914).
224 Jeff Herman, Average Credit Card Debt Statistics, CREDITCARDS.COM (Jul. 16, 2019),
https://www.creditcards.com/credit-card-news/credit-card-debt-statistics-1276.php
[https://perma.cc/9VFQ-X97P].
225 See Dilworth, supra note 51. This is not an exact calculation, as the households do
not carry the same balance throughout the year, and it does not account for monthly
compounding.
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

44 Draft[Vol. __
effective APR of 40%–90%.226 The installment loans industry serves
10 million Americans annually and earns over $10 billion in finance
charges.227 And this industry is still cheaper than the payday lending
industry, which charges a typical 400% APR.228
I do not mean to argue that abolishing the payday would abolish
either the payday industry or the short-term credit industry.229
People borrow for many reasons—smoothing consumption, pursuing
opportunities, bracing shocks, etc.230 The demand for short-term
credit solutions is based on real need, and the lack of liquidity due to
the payday is but one of them. Still, there is little doubt that short-
term credit solutions are very expensive and can often lead to
inescapable debt spirals. Thus, achieving even a meaningful
reduction in the demand for these services is a worthy social goal. To
get a sense of the potential impact, consider the results of a study that
examined the effects of an unexpected $600 tax rebate on payday
borrowing. Using a variation in the timing of the rebate, the
researchers found a large and marked effect on the demand for
payday loans. In their analysis, payday borrowers were roughly 16%
less likely to borrow from payday lenders within two pay cycles of
receipt of the rebate.231 This effect, unfortunately, disappeared after
two pay cycles.232
Another important potential effect of abolishing the payday is
that it may also lead to the abolition of the wasteful monthly utility
payment practice. As noted, households consume daily but pay
monthly. In consuming now and paying later, households are
essentially borrowing from utility providers. And of course, this credit
transaction comes at a cost; utility providers charge for offering credit
services. This credit transaction is artificial; it may be an artifact of
the payday itself. With greater liquidity, perhaps service providers
can be made to charge households on a daily basis as well. By moving
to daily payments, the cost of utilities can decline by what is now the
cost of the interest payments that are implicit in the monthly bill. If
the technology is ripe—and to a large extent it already is—then the
costs of these additional transactions would be trivial. This means
that removing this unnecessary credit transaction may result in
dramatic savings—think about a household that borrows a few
hundred dollars every month and its annual cost of doing so,
226 State Laws Put Installment Loan Borrowers at Risk, PEW 3 (Oct. 17, 2018),
https://www.pewtrusts.org/en/research-and-analysis/reports/2018/10/17/state-laws-put-
installment-loan-borrowers-at-risk [https://perma.cc/7AEP-NXZD].
227 Id. at 2.
228 See CFPB, supra note 52.
229 Cf. Hawkins, supra note 148, at 7(“[E]arned wage advances have the potential to end
payday lending”).
230 See Robert B. Nielsen et al., Consumer Finances of Low-Income Families, in
HANDBOOK OF CONSUMER FINANCE RESEARCH 167, 169 (2016) (“Credit can help low-income
consumers smooth consumption, invest in human capital, and build assets, but the high cost
of credit can crowd out current consumption and saving”).
231 Marianne Bertrand & Adair Morse, What Do High-Interest Borrowers Do with Their
Tax Rebate?, 99 AMERICAN ECON. REV. 418, 422 (2009).
232 Id.
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

2020] Payday 45
multiplied by almost all households.233 How much of the savings will
actually be passed on to households is an important question, and
while there is no reason to assume that all of the savings will be
passed to households, there is also no general reason to assume that
none of the savings will pass. It is enough for now to note that even if
some of the savings will pass, the effect of abolishing the payday on
households can be significant.
Overall, paying workers more frequently would have an
important positive effect on their well-being and reduce the demand
for short-term credit solutions.
B. Alternatives to Abolition
At this point, I hope, the question is no longer whether the
payday is worth preserving, but rather what the viable alternatives
are. As I propose the abolition of the payday in favor of daily streams
of payment, I should explain why other more “moderate” solutions are
ill-advised.
What is perhaps the leading alternative to dealing with the
problem of the payday is the use of wage-advances. Today, there is a
flurry of activity in this space by Fintech companies that compete over
a variety of wage advance solutions.234 These products go by different
names—wages on demand, earned income access, advance wage
payment—but they all share a basic structure: the employee is paid
ahead of the payday as part of the anticipated pay.235 The advance is
paid by either the employer or a third party, which specializes in
making advances against the employee’s wages. In a strict sense,
these are not really advances, as they mostly apply to earned wages.
Hence, the employee is not paid early but is instead lending less. But
whatever the terminology, the effect is the same—bridging the gap
between earning one’s pay and the payday. Thus, the concentration
of activity in this sector is a good indication of the size of the problem
of K and vividly demonstrates K ’s inefficiency.
2 2
Such advances can offer a response to short-term liquidity
shocks, such as a car that suddenly needs a costly repair or an
emergency hospital visit. Nonetheless, advances are a flawed,
incomplete, and potentially harmful solution to the underlying
problem—justified only if deeper solutions are unavailable but
otherwise a band-aid for a lost limb.
The central objection is cost. Paying employees in advance
involves cost on the side of either the employer or the third-party
company. Someone has to hold sufficient capital, handle requests, and
233 As noted, millions of households default on utility payments and the costs of default
are spread, at least in part, among all other consumers. See supra note 102.
234 See Hawkins, supra note 148, at 5–6; Nakita Q. Cuttino, The Rise of “FringeTech”:
Regulatory Risks in Early Wage Access, 115 NW. U. L. REV. (forthcoming, 2020).
235 Stephen T. Middlebrook, What Business Lawyers Need to Know About Wage Advance
Products, A.B.A. 4 (Sep. 5, 2019), https://businesslawtoday.org/2019/09/business-lawyers-
need-know-wage-advance-products/ [https://perma.cc/JJQ9-EWTT].
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

46 Draft[Vol. __
create mechanisms to ensure proper deductions come payday. Few
would be willing to bear this cost for free.
While Fintech and terms such as “wage on-demand” sound novel,
the history of employer advances is longstanding, and it is not
wholesome. The first wage payment laws emerged as a response to
concerns with “employers that took improper deductions from worker
wages or forced them to borrow from employers.”236 The effect of these
issues is reflected in the memorable “Sixteen Tons,” written by Merle
Travis in 1946 and modeled after his father’s experiences working in
the coal mines:237
You load sixteen tons, what do you get?
Another day older and deeper in debt
Saint Peter don't you call me 'cause I can't go
I owe my soul to the company store238
To combat abuse, states passed legislation that regularized
paydays and limited employer’s ability to deduct fees and interest
from employees’ wages.239 It is not without irony, notes Professor
Willborn, that “the payday loan industry had arisen to do almost
exactly what employers were doing prior to the state wage- payment
laws.”240
Today, there is still great regulatory uncertainty regarding
advances.241 While some view these as services that provide the
consumer with much-needed credit, others see them as opportunities
to profit at employees’ expense.242 The relevant framework, even at
the Federal level, is complex—involving the interpretation of the
Truth in Lending Act (TILA), the Equal Credit Opportunity Act
(ECOA), the Fair Credit Reporting Act (FCRA), the Fair Debt
Collection Practices Act (FDCPA), and the Consumer Financial
Protection Act (CFPA).243 Article 9 of the Uniform Commercial Code
(UCC) also adds complexity, as it views the sale of accounts (i.e.,
future payments) as a secured transaction, thus subjecting it to its
burdensome framework.244 Some state laws also require licenses to
lend, limit wage assignments, and impose usury limits.245 This results
236 Willborn, supra note 81, at 40.
237 Sixteen Tons: The Story Behind the Legend, TENNESSEE ERNIE FORD,
https://www.ernieford.com/sixteen-tons [https://perma.cc/9HSZ-B3GF] (last visited Feb. 13,
2020). The song resonates other problems of the time – payment in script and company credit
and wage theft.
238 MERLE TRAVIS, SIXTEEN TONS (Capitol Records, 1947).
239 Willborn, supra note 81, at 40.
240 Id, at 41.
241 See Cuttino, supra note 234, 39–45.
242 For the debate, see Hawkins, supra note 148,, at 36–40.
243 Adam Levitin, What Is “Credit”? AfterPay, Earnin’, and ISAs, CREDIT SLIP (Jul. 16,
2019, 1:31 PM), https://www.creditslips.org/creditslips/2019/07/what-is-credit-afterpay-
earnin-and-isas.html [https://perma.cc/CG8U-B5T7] (arguing that, inasmuch as no finance
charges are levied, some advance products are exempt from TILA but subject to other forms
of credit legislation).
244 U.C.C. § 9-109(a)(3) (AM. LAW INST. & UNIF. LAW COMM’N 1977).
245 Hawkins, supra note 148,at 15–24.
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

2020] Payday 47
in a very complex regulatory landscape, and employers explain their
reluctance to offer advances in this complexity.246
Third-party advance companies are for-profit companies, and
they turn a profit by charging fees, commissions, and, oddly, tips. One
such product is called Earnin’, where users are encouraged to leave a
tip of $0-14 per $100 advanced; failure to leave a tip is believed to
restrict the user’s access to cash.247 A $14 charge per $100 is very close
to the cost of payday lending ($15). Another study of Fintech
companies finds that the average APR ranges from 20% to 145%.248
It is damning with faint praise to say that these products, “although
[expensive] in absolute terms, appear[] clearly superior to [short-term
loan] alternatives.”249
Some of the costs are less visible. Professor Jim Hawkins recently
reviewed the contracts used by market players. He found that despite
Fintech companies' self-attestation to being “concerned with their
social impact” and notwithstanding the intense regulatory scrutiny,
their contracts are “surprisingly unfriendly” to the consumer.250
Arbitration, disclaimers of warranties, unilateral contract
amendments, and high fees are some of the more common issues.251
It is highly likely that, even if permitted to operate, purveyors of
advances will be held under strict regulation.252
Reforming laws to facilitate advances would result in a
complicated and costly patchwork of legislation. It is inevitable that
some advance companies will go the way of many lenders in the past:
resorting to abusive terms, one-sided “mistakes,” and excessive rates.
The issue is not so much that companies seek to profit; it is that the
problem they seek to solve is an artifact of badly-designed legislation
and dated money technology. Treating this problem directly can
resolve the liquidity problem directly without requiring the
development of a newly-regulated industry. Although the focus
should be on eradicating the payday entirely, advance payments are
a step in the right direction. They highlight, quite clearly, the
unreasonable burden K imposes on workers. They also develop
2
technologies and solutions for regularizing payments. And, to the
246 Id. at 8.
247 Kevin Dugan, Cash-advance App Earnin Gets Subpoenaed by NY Regulator, N.Y.
POST (Mar. 28, 2019), https://nypost.com/2019/03/28/cash-advance-app-earnin-gets-
subpoenaed-by-ny-regulator-source/ [https://perma.cc/R6MR-N8MJ] (“Earnin encouraged
users to leave a tip of anywhere between zero and $14 on a $100 weekly loan. Users who
don’t leave a tip appear to have their credit restricted. Meanwhile, a $14 tip would equate to
a 730- percent APR—nearly thirty times higher than New York’s 25 percent cap.”). In
evaluating the costs, one should consider the regulatory uncertainty; if it would ever be
resolved, one might expect greater competition in this space.
248 Todd H. Baker, FinTech Alternatives to Short-Term Small-Dollar Credit: Helping
Low-Income Working Families Escape the High-Cost Lending Trap, HARV. KENNEDY SCH.
46 (2017), https://www.hks.harvard.edu/sites/default/files/centers/mrcbg/files/75_final.pdf
[https://perma.cc/Z8SF-FEDY].
249 Id.
250 Hawkins, supra note 148,at 23–24.
251 Id.
252 Id. at 43–48 (proposing regulation).
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

48 Draft[Vol. __
extent the solutions provided here would take time and political will
to implement, wage advances can serve as an interim solution.
IV. A WORLD WITHOUT THE PAYDAY
The abolition of the payday requires steps that are far more
conservative than the goal might seem to imply. Indeed, while the
problems caused by the payday are severe, the solutions are fairly
mild. This suggests a low-hanging policy fruit: large effects with small
changes—ones that do not risk complex, unanticipated systemic
effects. With sufficient goodwill, this policy can be implemented in a
very short time span, dramatically improving the welfare of millions
of Americans. In this Part I will present a workable alternative to
payday and then move to discuss several complementary ways to
implement it. Each method has somewhat different benefits and
costs, and much can be achieved even if only some of these methods
are implemented.
Let us first reflect on the two most important barriers to regular
pay: compliance costs and money technology for the underbanked.
Both of these issues create a scale problem: while payroll software can
fairly accurately estimate pay, the costs of inadvertent compliance
errors are high, thus requiring human supervision and authorization
for each payment. Whereas ACH money transfers provide a fairly
cheap solution, even with daily payments, it is inapplicable with
respect to the unbanked and the underbanked who must rely on
inefficient alternative money instruments, such as checks.
In considerations of these issues, I propose the following.253 At
the end of each day, employers will be required to pay employees at
least 93% of a good-faith estimate of their earned income.254 The
payday will be replaced by an “accounting day,” or a true-up, once
every two weeks when the employer must complete a final calculation
of the employee’s full earned income for the period. After making this
calculation, including all adjustments for unclaimed deductions,
bonuses, commissions, etc. , the employer will adjust the daily pay to
reflect outstanding amounts. If no adjustments are necessary, the
employer will pay the employee the daily 7% shortfall, which would
come to an extra day’s worth of pay, once every two weeks. As long
as the employer makes a good faith estimate of the daily pay, the
employer will not be held liable for regulatory compliance issues for
daily pay—such liability will only follow if, as is today, the employer
fails to pay in full on accounting day.255
253 The mode of reform can be legislative, but it is worth noting that the Restatement of
Employment Law also recognizes the possibility of changes to employment law through the
common law. See RESTATEMENT OF EMP’T LAW § 3.01 (AM. LAW INST. 2015) (“wage-
payment laws . . . do not generally preclude common-law development because they are
based on contract principles found in the common law”).
254 The choice of 93% is meant to create enough reserve to capture a full day’s wage. So
if the employee works 10 days in a 14 days period, and earns $1923, the employee’s daily pay
will be $137.6, and the biweekly adjustment will pay the employee an additional $134.6.
255 Given the predictability of pay for most professions, and the low profit from
underpaying every day, this duty is not expected to generate considerable friction or litigation.
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

2020] Payday 49
Under this proposal, employers will not pay their workers their
full daily pay but only an estimate of part of it. The reason why
employers will not be required to pay in full is grounded in a few
considerations. It is very difficult to know the total amounts due to
employees, given all the possible deductions, taxes, and levies. Hence,
some estimation may be unavoidable, and this means that there will
often be errors, either of over- or under-payment. If employers are not
afforded some margin of error, that would require them to carefully
review each payment—and the costs of doing so daily may be
prohibitive.256 Another important consideration is that it is arguably
harder for the employer to collect money owed from the employee than
vice versa, given the greater mobility of the employee and lack of
collateral. Leaving 7% of the income to the last day of the fourteen-
day period is calculated to create a buffer that, on the one hand,
allows the employee to keep most of the daily pay and, on the other
hand, accounts for potential errors in daily estimates. Subject to
further experimentation, this margin should be sufficient to allow
employers to make offsets against mistakes in overpaying
employees.257 It also means that the employee is receiving on the last
day of the biweekly period an extra day’s pay (which is deducted from
their on-going payments). This feature may appeal to those who think
employer-based budgeting is helpful.
The design of biweekly pay is meant to address two concerns:
wage monitoring and compliance-cost control. Wage theft is an
important concern, and monitoring daily payments may be harder
than monitoring the transfers of larger lump sums.258 Of course, once
the employee grows accustomed to daily pay, he or she could detect
deviations by comparing actual payments to normal payments. Still,
with possible daily fluctuations, deviations are harder to detect. To
deal with this problem, on accounting day, the employer would
produce a pay stub that accounts for all of the biweekly payments.
The employee can then compare this amount to amounts paid, just as
easily as can be done today.259 The second function is controlling
compliance costs. As noted, a large part of the cost of making
payments is due to the need to verify compliance with a variety of
different laws. Because the final accounting is only done once every
two weeks, the employer would not need to engage in more
compliance than they do today, besides the fairly trivial calculation of
93% of the expected daily pay. Note that the employer does not bear
liability for small or unintentional deviations, making it unnecessary
to verify daily payments with the same degree of attention as the
biweekly pay.
256 See supra Part I.F (discussing costs of payroll).
257 In most industries, a much smaller buffer would be needed—and perhaps no buffer is
even needed for salaried employees with fix wages. Still, it is prudent to start with a moderate
buffer in experimenting with the implementation of this proposal.
258 On wage theft, see supra note 153 and accompanying text.
259 It may be necessary to add in the bank’s user-interface support for easy comparisons
of employer-pay per wage period. Such technology is already implemented in the apps and
websites of most banks, which allow users to filter deposits by recipient per period.
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

50 Draft[Vol. __
One remaining issue is the control of money-transfer costs. As
noted, this is not an issue for the majority of workers, who are banked
and can benefit from ACH transfers, but it is still a pressing and
painful issue for the under- and unbanked. The solution here is
technological and I explore in greater length the use of pay cards as a
viable solution to this problem.260 In addition to pay cards, others
have proposed non-technological alternatives, such as postal and
public banking, which can also mitigate these issues.261
An optional addition to this proposal would be to allow
employees to elect a biweekly payday. That is, the daily pay would be
presented as an option alongside biweekly pay, and employees could
elect which payment option they prefer. In terms of preserving
employee choice, this would seem superior, as those employees who
find biweekly pay more manageable would elect it. Such a choice may
be preferred by some—if the worker has no need for liquidity or finds
it difficult to budget otherwise. But for the reasons I laid out earlier,
I believe employer-side savings is a bad idea due to the counterparty
risk.262 If employees need help budgeting, bank-side savings
programs are a superior alternative. And if employees want to lend
money, they can always do so in explicit capital markets, where there
is more robust competition for their money. Hence, there is legitimate
concern that presenting this option may be a trap for the unwary and
will serve little other function.263
The final part of this proposal is that it envisions a transition
and experimentation period. Wages and payments are systemic
issues; they affect every part of the economy. The urgent need for
reform should be tempered with patience and understanding that
immediate implementation may be harmful. Instead, an
announcement of a target date for daily pays in a few years, perhaps
coupled with a transition to weekly pay, is likely the most prudent
course of action.
Implementing this reform would require some legislative
changes. The key changes are focused on changing labor laws that
impede more frequent pay; changing our money infrastructure;
improving market education; and changing the market by leadership.
Each of these interventions, summarized in the Table below, is
developed in the following subsections:
Promoting Frequent Pay
Method Type of Change Notes
260 See infra Part I.D.
261 The U.S. postal banking system was abolished in 1966. On its history and for a
proposal to reinstate it, see MEHRSA BARADARAN, HOW THE OTHER HALF BANKS, 183–226.
See also Know the Facts, CAMPAIGN FOR POSTAL BANKING,
http://www.campaignforpostalbanking.org/know-the-facts/ [https://perma.cc/GW9R-XJKB]
(last visited May 12, 2020).
262 See supra Part III.A.
263 A more compelling reason to favor biweekly pay is if the check-cutting costs are high,
the employee could be paid more by being paid less frequently. However, this is a
transitionary issue until the money and payroll technology are sufficiently advanced.
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

2020] Payday 51
Information Demand-side pressure by Least intrusive
making implicit interest
transparent
Leading/fiat Changing legislation to Requires
encourage and mandate political will
frequent pay but can have
cascading
effects
Fixing Removing inadvertent Requires
employment incentives to reduce pay political will
law frequency but involves
relatively few
partisan issues
Money Making payments cheaper, Long term
technology especially to the underbanked investment
with positive
externalities
A.Changing by Information
One reason why the payday persists is related to the employer’s
power in employment negotiation. Perhaps employees are
insufficiently aware of the credit nature of K . If that is the case,
2
employees would also be unaware of the true cost of K and would not
2
demand an appropriate wage premium. This imbalance of
information or sophistication tilts the balance in favor of the employer
and leads to inefficiently infrequent pay periods.
This idea—that individuals misprice credit transactions—is a
central impetus for the enactment of TILA. Congress diagnosed that
consumers engage in “uninformed use of credit” and prescribed
“meaningful disclosure of credit terms.”264 By conspicuously
disclosing credit terms using a uniform standard, TILA hopes to
improve consumer finance decisions.265
The logic of TILA can be brought to bear on payday.266 If
employers want to borrow money from employees through the
payday, they might be required to disclose the fact that payday is a
credit transaction. This can be done in the written employment
contract or in a separate disclosure. More importantly, the employer
might be required to display the (implicit) interest rate in this
transaction. Using the same language as that used when consumers
borrow—the so-called Schumer’s Box—the employer will be required
to disclose how much the employee receives in exchange for the
extension of credit. This disclosure would allow workers, subject to
264 Truth in Lending Act § 102(a), 15 U.S.C. § 1601(a) (2018).
265 See generally Hosea H. Harvey, Opening Schumer's Box: The Empirical Foundations
of Modern Consumer Finance Disclosure Law, 48 U. MICH. J.L. REFORM 59 (2014)
266 For illustration, see Figure 2 in Part II.C., which illustrates how effective pay is
comprised of both per-hour wages and frequency of payments.
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

52 Draft[Vol. __
the general caveats about disclosure in general,267 to better
understand the meaning of the credit element of the payday and to
“shop” effectively—that is, to understand how pay frequency
compares to the cost of borrowing from other sources and choose, if
given the option, a shorter pay period. The following figure illustrates
using a typical employee who earns $1,923 biweekly.268
Figure 3 – A “Schumer’s Box” for
Interest Rates and Interest Charges
Annual Percentage Rate 5% This the amount of interest paid to you.
(APR) for Biweekly Pay Compare to your cost of borrowing.
Biweekly Interest $4 This is the amount of interest the
Charges Paid to employer pays you for delaying your
Employee payments by two weeks
Total Wage (Biweekly) $1923 $1919 base + $4 interest
The use of disclosure also has one substantive implication in the
context of the minimum wage. If an employer borrows money, it
should identify the portion of the pay that is the wage premium. The
remaining pay is the pay-for-work portion of the wage paid to the
worker. A prolonged pay period undercuts the minimum wage
obligations of the employer; paying $7.25 hourly with a daily payday
is not the same as paying it monthly. In the latter case, the effective
pay is much lower and the employer is arguably failing to meet the
minimum wage obligations, at least in spirit. That federal legislation
does not account for this difference suggests a serious blind spot, even
among legislators and judges. Once advertised, courts could start
paying better attention to determine the proper baseline envisioned
by the FLSA—is it daily pay, weekly pay, or something else?
B.Changing by Leading
Another potential explanation for the persistence of the payday
is that government employees are paid biweekly.269 Social norms can
have a significant effect on market outcomes, and if the government
declares a certain pay period to be the standard, then this
pronouncement might have downstream effects on private employers.
If this explanation carries any explanatory power, it opens the
road to straightforward intervention. Under Title 5 of the United
States Code, all federal employees are to be paid once every two
267 See Yonathan A. Arbel and Andrew Toler, ALL-CAPS, (U. of Ala. Legal Stud. Working
Paper No. 3519630, 2020) (providing evidence of the failure of the most common mode of
conspicuous disclosures, disclosure via all-caps)
268 Wage data based on Measures of Central Tendency for Wage Data, SOC. SECURITY
ADMIN., https://www.ssa.gov/oact/cola/central.html [https://perma.cc/QM3L-FENF] (last
visited Feb. 13, 2020).
269 See supra Part I.E.
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

2020] Payday 53
administrative workweeks.270 This period could be changed to a daily
payment of 93% of the daily pay, subject to a biweekly accounting.
Notably, the change will not infringe on any employee’s rights. Nor
will this reform require large substantive changes. Admittedly,
changing federal legislation is not easy, and I do not mean to discount
the political and procedural challenges, especially because state law
is so diverse and will also have to be amended. However, the
importance of the goal, and its non-partisan nature, promise some
optimism.
C.Fixing Employment Law
One impediment to abolishing the payday is, ironically,
minimum wage legislation. As I have noted, the FLSA makes
employers average the minimum wage payments over the entire pay
period. This incentivizes employers to extend the pay period as much
as possible so they can benefit from averaging. If a tipped employee
is making above minimum wage in week one and below minimum
wage in week two, the employer could avoid compensating the
employee for week two by setting a biweekly payday. We also saw that
overtime legislation, at least in theory, does not have this flaw.
The faulty legislative design opens the door to a number of
potential interventions. The key to all of these options is to divorce
the averaging period from the pay period. Hence, the option with the
least effect on the status quo would permit employers to choose their
accounting periods. The accounting period would substitute today’s
payday and would be the day on which the employer will average the
employee’s pay and see if any amounts are still due to meet the
minimum wage requirement for the accounting period. The length of
the accounting period could be regulated by the same limitations set
today by state legislation on pay periods. This way, the employer
would pay the employee each day of the week and then, come
accounting day, make sure that a minimum wage was paid. If there
was any shortfall in payments, the employer would add it to that day’s
pay. Over a two-week period (or however long the accounting period
is) the employee would be paid the exact same amount the employee
would have been paid under the payday—but at more frequent
intervals. This aspect of the proposal means that neither employee
nor employer rights are harmed by this transition, yet the indignities
of the payday are avoided.
Similarly, overtime legislation should divorce pay frequency
from the definition of who is a salaried employee; there is no reason
to tie the definition to the (in)frequency of pay.271 A daily-paid
employee can equally be salaried or unsalaried, and the frequency of
pay need not reflect on this determination.
270 5 U.S.C. § 5504(a) (2018).
271 If one believes that this definition tracks any meaningful practical distinction, it is
possible to use the accounting period instead of the current pay period as the measure of the
period.
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

54 Draft[Vol. __
Finally, employers’ compliance with wage and hour laws should
be evaluated at the accounting period. Thus, if an employer makes a
compliance error on a specific day, this should not be a cause for a
lawsuit. The goal is to reduce ongoing compliance costs, and allowing
lawsuits to proceed based on random errors would undermine this
goal. At the same time, employers are still under a duty to make a
good-faith estimate of the 93% pay the employee deserves. This
means that employers do not have a carte-blanche right to underpay
employees daily. While one-off or even occasional mistakes should not
be grounds for a lawsuit, the employee should be allowed to sue for
systematic mistakes if they are done in the employer’s favor. Hence,
the proposal does not derogate in any way from minimum wage laws
or overtime laws under the status-quo; it neither increases pay nor
reduces it. The only effect is on pay frequency.
D.Improving Money Technology
Transferring money is more difficult than would appear at first
glance. I have already noted the various costs associated with bank
transfers, the difficulty of storing and handling cash, and the many
costs of writing and liquidating checks.
Digital money is clearly the future, and, to a growing extent, it
is the present.272 In particular, employers are now increasingly using
payroll cards.273 A payroll card is akin to a debit card and is issued by
a bank or another financial institution. The account is not attached
to any depository account, and thus, the card owner is spared the cost
and difficulty of opening a bank account. Instead, the owner charges
the card against the available balance. In 2017, roughly 3.4% of
households reported receiving income with a payroll card,274 and the
CFPB estimates a 6% growth, amounting to $44.6 billion loaded onto
these cards.275 In 2015, nineteen state governments were already
using payroll cards,276 and one survey suggests that 7 million
workers were using them in 2014.277
272 An estimated 4% of Americans hold only a prepaid card. Analysis based on data
presented in FED. RES. SYS., supra note 5, at 18–19.
273 On the other hand, a survey by the FDIC found that usage of prepaid cards by
households ranged between 7.9% (2013), 9.8% (2015), and 9.2% (2017). FED. DEPOSIT INS.
CORP., supra note 195, at 7. 9.2% of households using a prepaid card reported receiving it as
a payroll card. Id.
274 Id. at 12. One estimate by a consulting firm from 2015 estimated that payroll
payments will exceed check payments by 2017, but FDIC data still shows that checks are far
more common. See Madeline K. Aufseeser, Checkmate: U.S. Payroll Card Programs Trump
Paper Checks, AITE GROUP (2015), http://www.aitegroup.com/report/checkmate-us-payroll-
card-programs-trump-paper-checks [https://perma.cc/3FE4-ZRUB].
275 See Prepaid Accounts Under the Electronic Fund Transfer Act (Regulation E), 12
C.F.R. § 1005 (2018); see also Truth in Lending Act (Regulation Z), 12 C.F.R. § 226
(2018).
276 Rating State Gov’t Payroll Cards, NAT’L CONSUMER L. CTR. 3 (2015).
277 Gregg Gelzinis et al., How Workers Get Paid Is Changing: Consumer Protections
Need to Catch Up, CTR. FOR AM. PROGRESS,
https://www.americanprogress.org/issues/economy/reports/2019/01/17/465223/workers-get-
paid-changing-consumer-protections-need-catch/ [https://perma.cc/DN9K-NJD5].
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

2020] Payday 55
Payroll cards are convenient, safe, and allow the immediate use
of the funds paid. Importantly, the employee does not have to have a
bank account to use a payroll card. This means that one’s
creditworthiness and legal status are not hurdles. Moreover, the
employee need not maintain a minimum balance in his or her bank
account or pay fees. The cost of depositing funds is also reportedly
low: $0.35 in deposit costs.278 It is not surprising, then, that many low-
paid employees view payroll cards positively.279
There are also various concerns with payroll cards, many of
which will be familiar to users of bank accounts.280 One concern is the
insurance of amounts deposited on these cards—what prevents a “run
on the card”? Then there is the issue of fees: ATM-use fees, point of
sale fees (i.e., a transaction fee), overdraft fees, and even balance
inquiry fees.281 By one estimate, the average per-employee fees were
$20 per month.282 To make things worse, the fees are badly
disclosed,283 among other concerns.284
Federal legislation partially covers payroll cards. Under the
Electronic Fund Transfer Act (EFTA) and Regulation E,285 financial
institutions that offer payroll card accounts must make account
information available to consumers by specific means, but they are
exempted from providing periodic statements.286 In addition, the
financial institution must allow consumers to report errors and limit
customers’ liability for unauthorized transfers.287 In April 2019, a new
CFPB rule came into effect, expanding the fraud, error, and
unauthorized charges protections to these cards; requiring simplified
disclosure; and providing for easy access to information.288 State
legislation in this area is developing. Roughly half of the states have
some laws that regulate payroll cards.289 The regulations usually
278 NEW YORK STATE ATT’Y GEN., PINCHED BY PLASTIC: THE IMPACT OF PAYROLL CARDS
ON LOW-WAGE WORKERS 3 (2014),
http://www.ag.ny.gov/pdfs/Pinched%20by%20Plastic.pdf [https://perma.cc/9EG4-9DVC].
279 Oswalt & Marzán, supra note 205, at 453.
280 Payroll cards do not bear interest, but given the typical rates in checking accounts,
this concern is of little practical consequence. See Liran Haim & Ronald Mann, Putting
Stored-Value Cards in Their Place, 18 LEWIS & CLARK L. REV. 989, 1008 (2014). .
281 See supra note 220, at 6.
282 Id. at 9.
283 Id. at 13.
284 Jessica Silver-Greenberg & Stephanie Clifford, Paid via Card, Workers Feel Sting of
Fees, N.Y. TIMES (June 30, 2013), http://goo.gl/VzyTz [https://perma.cc/CZ74-Z7T8].
285 Electronic Fund Transfer Act (EFTA), 15 U.S.C. § 1693 (2012).
286 Electronic Fund Transfer Act, FED. RES. 18,
https://www.federalreserve.gov/boarddocs/caletters/2008/0807/08-07_attachment.pdf
[https://perma.cc/XL2T-72P5] (last visited Dec. 28, 2019).
287 FED. RES., REGULATION E: ELECTRONIC FUND TRANSFER ACT 6,
https://www.federalreserve.gov/boarddocs/supmanual/cch/efta.pdf [https://perma.cc/48JL-
EHDY] (last visited Dec. 28, 2019).
288 See generally CONSUMER FIN. PROT. BUREAU, PREPAID ACCOUNTS UNDER THE
ELECTRONIC FUND TRANSFER ACT (REGULATION E) AND THE TRUTH IN LENDING ACT
https://files.consumerfinance.gov/f/documents/20161005_cfpb_Final_Rule_Prepaid_Accou
nts.pdf [https://perma.cc/KZL4-DKJL] (last visited Dec. 28, 2019).
289 Rachel Blakely-Gray, Pay Card Laws By State and Regulations, PATRIOT SOFTWARE
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

56 Draft[Vol. __
permit the use of these cards but impose some limits on fees and set
rules on proper fee disclosure.290 Finally, a series of class actions were
filed against employers who offered payroll card programs for failing
to obtain employee consent and for violating wage and hour laws.291
In one of these cases, a court in Pennsylvania ruled that the
mandatory use of pay cards that impose fees is illegal.292
Facilitating the use of payroll cards is an important step
towards the abolition of the payday. The recent CFPB regulation
offers an initial framework, safeguarding certain employee rights,
although more experimentation is needed.293 Still, the fragmented
nature of state legislation impedes much innovation.294 Admittedly,
it is difficult to design a fee structure that would make payroll cards
profitable to operate and yet not encumber poor households with
additional expenses. Still, others have made the case that increasing
access to banking through public subsidies can be justified both as a
matter of redistribution and efficiency.295
Against this regulatory backdrop, positive steps can be taken to
promote payroll cards, at least for an initial period of adoption, such
as offering certain tax subsidies or requiring all employers to offer
this option.
A less obvious hurdle in the way of payroll cards is pro-employee
regulation that mandates that employers offer the choice of payment
methods. The Electronic Fund Transfer Act and Regulation E prohibit
employers from forcing employees to receive wages via pay card.296
New York law requires employers to provide employees with at least
one fee-free method of payment every payday.297 This choice creates
unanticipated problems: if, when setting a daily payday, employers
(June 5, 2019), https://www.patriotsoftware.com/payroll/training/blog/pay-card-laws-by-
state/ [https://perma.cc/KMQ8-LHAG].
290 See generally id. (surveying state laws); see also Sarah Jane Hughes & Stephen T.
Middlebrook, Are These Game Changers? Developments in the Law Affecting Virtual
Currencies, Prepaid Payroll Cards, Online Tribal Lending, and Payday Lenders, 70 BUS. L.
261, 265 (2015).
291 See Joint Stipulation of Class Action Settlement and Release ¶¶ 42, 47, Chavez v.
PVH Corp., No. C 13-01797-LHK (N.D. Cal. Feb. 12, 2014); First Amended Compl. ¶¶ 1-4,
Lapan v. PVH Corp., No. C 13-05006-YGR (N.D. Cal. Dec. 9, 2013); Plaintiffs' Original
Petition ¶¶ 14, 16, 19, Branson v. Destiny Foods, Inc., No. D-1-GN-14-001131 (Travis Cty.,
Tex., May 13, 2014).
292 See Siciliano v. Mueller, 149 A.3d 863 (Pa. Super. 2016).
293 See Haim & Mann, supra note 280, at 1014–19.
294 See Benjamin Lo, Fatal Fragments: The Effect of Money Transmission Regulation
on Payments Innovation, 18 YALE J. L. & TECH. 111 (2016).
295 Barr, supra note 206, at 237 (“[N]etwork externalities in electronic payments
systems and distribution networks suggest that net social benefit could be obtained through
further expansion.”).
296 See Letter from Richard Cordray, Dir., CFPB, to Sen. Richard Blumenthal (Sept. 12,
2013), http://goo.gl/KOqFzB [https://perma.cc/3KCJ-2QSN]; CONSUMER FIN. PROT.
BUREAU, CFPB BULL. NO. 2013-10: PAYROLL CARD ACCOUNTS (REGULATION E) 1–2 (Sept.
12, 2013), http://goo.gl/98d8I6 [https://perma.cc/FM77-RZJS].
297 N.Y. LAB. LAW § 191 (2018); see also N.Y. STATE DEP'T. OF LAB., aOp. No. RO-08-
0001
https://labor.state.ny.us/legal/counsel/pdf/Direct%20Deposit%20of%20Wages/Payroll%20c
ards%20Letter%2010-29-2009.pdf [https://perma.cc/57T7-D4ZY] (Oct. 29, 2009).
Electronic copy available at: https://ssrn.com/abstract=3547007

<<PAGE_BREAK>>

2020] Payday 57
must pay some employees in cash or check, this cost could be
significant. Employee choice, then, can undermine the viability of
payment streams.
The solution, however, is straightforward. The daily pay option
can be made open only to employees who are willing to use pay cards
or bank transfers. Relative to today, where all employees are paid on
long intervals, employees who favor cash will not be harmed by
having this additional option. But for all other employees, this option
would greatly advance their wellbeing.
V. THE DAY AFTER PAYDAY: CONCLUDING THOUGHTS
A complicated dynamic of dated legislation, path-dependence,
and inefficient money technology has contributed to the economy-
wide practice of paying employees in arrears. This dynamic puts
employees in the absurd position of lending money to their employers.
This feature of the modern economy is clearly a software
problem, not a hardware problem. We can, and should, pay workers
in at least the same frequency we pay overseas vendors. Instead, our
antiquated system of payments creates significant financial stress,
leading households to borrow from payday lender and other providers
of short-term credit products.
Abolishing the payday might take time, as it will face
resistance. No change is easy. However, the case for paying people for
their work is too compelling to ignore. Paying employees late may
made sense when we had to compute wages by hand and carry cash
chests between worksites. But it makes little sense when sending
digital money has become so ubiquitous that our vocabulary includes
new verbs to describe instantaneous money transfer – e.g., “I will
Venmo you the money tomorrow,” and “I just Paypaled you.”
With our new hardware, it is time to update our legal software.
Electronic copy available at: https://ssrn.com/abstract=3547007
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
