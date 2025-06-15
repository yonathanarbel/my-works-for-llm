package main

import (
    "fmt"
    "os"
    "time"
    "math/rand"
    "encoding/json"
    "net/http"
    "log"
    "strings"
)

const articleText = `
Arbel_Galley (Do Not Delete) 1/2/2018 4:17 PM
Adminization:
Gatekeeping Consumer Contracts
Yonathan A. Arbel*
Large companies and debt collectors frequently file
unmeritorious claims against consumers. Recent high-profile actions
brought by the Consumer Financial Protection Bureau against J.P.
Morgan, Citibank, and other large debt collectors illustrate the breadth
and importance of this phenomenon. Due to the limited financial power
of individuals, consumers often do not defend against such baseless
claims, which results in the entry of millions of default judgments every
year. To combat this problem, policymakers and scholars have explored
a variety of court-based solutions that would make it easier for
consumers to defend in court, but these prove ineffectual.
To solve the problem of unmeritorious claiming, this Article
proposes a budget-friendly solution called “Adminization.” This novel
approach uses an administrative agency as a gatekeeper to civil
litigation that is tasked with detecting and sanctioning the filing of
baseless claims. The agency samples cases, using statistical methods
and potentially deep-learning algorithms, and then investigates selected
cases using agency auditors. When the auditors find wrongdoing, they
are instructed to levy large fines against wrongdoers. Unlike the current
system, Adminization subjects every plaintiff to the risk of thorough
investigation and large fines, thus undercutting the financial incentive
to engage in wrongful behavior. The importance of Adminization lies in
its cost-effectiveness, practicality, and political feasibility relative to the
court-based approaches that dominate the discussion today.
* Assistant Professor, University of Alabama School of Law. For insight and comments, I
am indebted to Aharon Barak, Oren Bar-Gill, I. Glenn Cohen, Andrew Crespo, Jesse Fried, Janet
Freilich, John Goldberg, John Golden, Patrick Goold, Richard Hynes, Louis Kaplow, Kobi Kastiel,
Jody S. Kraus, Duncan Kennedy, Ronald Krotoszynski, Ethan J. Leib, Andrew Lund, Gideon
Parchomovsky, David Rosenberg, Ronald J. Scalise, Steven Shavell, Ted Sichelman, Henry Smith,
Matthew Stephenson, Cass Sunstein, Daphna Renan, Duane Rudolph, and W. Mark C.
Weidemaier. I am deeply grateful for the wise advice I received from participants in the
Southeastern Workshop at Washington University at St. Louis and Villanova Business
Colloquium, and to the dedicated work of the editors at the Vanderbilt Law Review. The Harvard
Foundations of Private Law Center provided generous research support.
121
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
122 VANDERBILT LAW REVIEW [Vol. 71:1:121
INTRODUCTION ............................................................................. 122
I. ABUSE AND FRAUD IN CONSUMER CREDIT CONTRACTS....... 130
A. Abuse and Procedural Violations ........................... 132
B. Justice, Inaccessible ............................................... 137
C. Lack of Judicial Oversight ..................................... 140
II. ADMINIZATION................................................................... 142
A. Adminization: High-Level Outline ......................... 142
B. Main Features of Adminization .............................. 144
1. Audits and Fines ......................................... 144
2. Sampling, Artificial Intelligence,
and Resource Management ......................... 146
3. Third-Party Communications ...................... 151
C. Adminization of Consumer Credit Litigation ......... 152
III. THE FAILURE OF PARTICIPATION-BASED SOLUTIONS .......... 157
A. Lawyering Up ........................................................ 158
B. Throwing Judges into the Fray .............................. 163
C. Modifying the Legal Process ................................... 164
D. Arbitration and Class-Defense ............................... 167
E. A Pyrrhic Victory ................................................... 169
IV. CHALLENGES ..................................................................... 171
A. Legal Authority ...................................................... 171
B. Feasibility of Adminization & Political Economy ... 173
C. Regulatory Capture ................................................ 174
D. Costs and Incidence ............................................... 175
CONCLUSION ................................................................................ 177
INTRODUCTION
When Margaret Donnelly, an eighty-five-year-old widow
suffering from a congestive heart disease, woke up that morning, she
did not realize she was hours away from facing a warrant for her
arrest.1 But that was the message the county sheriff had for her. He
explained to her that a lawsuit was filed against her for a debt of $1,471
in the local court.2 He also informed her that because she failed to
appear in court, the judge entered a default judgment against her,
which she now had to pay from her meager income. This news caught
Ms. Donnelly by surprise—she never heard about the lawsuit.
1. Beth Healy, Dignity Faces a Steamroller, BOS. GLOBE (July 31, 2006),
https://www.bostonglobe.com/metro/2006/07/31/dignity-faces-
steamroller/SoK0TBVHzOzjLEpNqNrVYN/story.html [https://perma.cc/7MW9-AWX4].
2. Id.
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
2018] ADMINIZATION 123
Admittedly, even if she had, there was little she could do, as hiring a
lawyer would overextend her budget. This is despite the fact that the
case had no merit whatsoever: the debt was paid in full many years ago
and, in any event, no evidence was brought to support the claim.3
Worryingly, the lawsuit was part of a pattern of abusive lawsuits filed
by a local law firm that targeted over one hundred thousand consumers,
a practice facilitated by the difficulty consumers like Ms. Donnelley face
in accessing the courts and challenging these unmeritorious lawsuits.4
A large body of evidence shows that millions of others in the courts also
face “a silent, shameful crisis that inflicts suffering and costs the nation
money, legitimacy, and decency.”5
Open doors, they say, may tempt the saints. Every year, about
eight million debt claims are filed by large companies and debt buyers
against consumers.6 Of those, over six million lawsuits turn into default
judgments, with little, if any, judicial oversight.7 One in three
consumers is estimated to be at risk of facing such a lawsuit.8 As with
Ms. Donnelly’s case, many of these debt claims lack merit and involve
debts that are resolved, expired, inflated, and in some cases, outright
fraudulent.9 A recent study found, for example, that debt buyers
knowingly purchase debts that are well beyond the statute of
limitations, with at least twelve percent of the debt portfolio of large
debt buyers consisting of stale debt.10 In 2016, Citibank and two of its
3. Id.
4. The case is known to have lacked merit because of its unusual circumstances: Ms.
Donnelley decided to represent herself in court against this lawsuit. This required her to litigate
the case for over a year and travel twice to the courthouse—not an easy task for a person in her
circumstances—but finally the judge was convinced that the case lacked merit. Id. The same law
firm that unsuccessfully sued her had successfully sued thousands of others “by demanding money
they had no right to collect and on the basis of debts they could not prove.” See Press Release,
Attorney Gen. Maura Healey, AG Healey Sues Major Debt Collection Law Firm Over Widespread
Consumer Abuses (Dec. 23, 2015), http://www.mass.gov/ago/news-and-updates/press-
releases/2015/2015-12-23-debt-collection-lawsuit.html [https://perma.cc/23KS-R82L].
5. Martha Minow, Opinion, We Must Ensure Everyone Has Access to Equal Justice, BOS.
GLOBE (Oct. 23, 2014), https://www.bostonglobe.com/opinion/2014/10/23/must-ensure-everyone-
has-access-equal-justice/pZxzjjHhR0GI89o0lZTnhP/story.html [https://perma.cc/U79G-VK34].
Dean Minow is the vice chair of the Legal Services Corporation (“LSC”), an independent nonprofit
established by Congress to provide financial support for civil legal aid. See LEGAL SERVS. CORP.,
https://www.lsc.gov/ (last visited Oct. 22, 2017) [https://perma.cc/AS3N-BXJD].
6. This estimate is based on data on the overall volume of civil litigation and estimates taken
from different studies regarding the number of consumer debt cases on the docket. See infra note
42.
7. See infra notes 92–95 and accompanying text.
8. See infra note 42.
9. See infra Section I.A.
10. See FED. TRADE COMM’N, THE STRUCTURE AND PRACTICES OF THE DEBT BUYING INDUSTRY
T-12 (2013), https://www.ftc.gov/sites/default/files/documents/reports/structure-and-practices-
debt-buying-industry/debtbuyingreport.pdf [https://perma.cc/S8YY-Z479] [hereinafter FTC DEBT
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
124 VANDERBILT LAW REVIEW [Vol. 71:1:121
affiliates were ordered to pay $11 million and forego the collection of
$34 million in consumer debt for the filing of false affidavits which
misstated both the size of the debt and its age.11 J.P. Morgan Chase
reached a $136 million settlement for its role in selling debts that were
legally uncollectable to debt buyers.12 The Consumer Financial
Protection Bureau (“CFPB”) also recently took action against a large
debt buyer who was ordered to pay over $2.5 million for its attempt to
knowingly collect on “fraudulent debts, debts that consumers had paid
or settled, and debts that were so old that they could no longer be legally
collected.”13 The regulator itself concluded that “[t]he system for
resolving disputes about consumer debts is broken.”14
To solve the problem of such unmeritorious claiming, this Article
proposes the “Adminization” of civil litigation. Adminization places a
gatekeeper administrative agency between consumers and debt
collectors, which is tasked with autonomously investigating and finding
bad cases before they reach court. After filing and before litigation, a
sample of cases will be audited by an administrative agency, and where
fraud is found, large fines can be issued against the offender. Both
economic analysis and psychology suggest that the mere prospect of
detection can deter wrongful behavior, and much more so when it is
coupled with severe fines.15 Using samples, audits, and fines,
Adminization will provide a fresh and cost-effective solution to
consumer credit contracts litigation—the most common form of all civil
litigation.
A few different institutional arrangements could support
Adminization, such as federal agency review through one of the existing
INDUSTRY REPORT) (reporting that 9.2 million cases, or 12.1 percent of the debts in portfolios
purchased by six large debt buyers, involved debts that are six years and older).
11. See Citibank, N.A., CFPB No. 2016-CFPB-0004 (Feb. 22, 2016).
12. See Press Release, Consumer Fin. Prot. Bureau, CFPB, 47 States and D.C. Take Action
Against JPMorgan Chase for Selling Bad Credit Card Debt and Robo-Signing Court Documents
(July 8, 2015), http://www.consumerfinance.gov/newsroom/cfpb-47-states-and-d-c-take-action-
against-jpmorgan-chase-for-selling-bad-credit-card-debt-and-robo-signing-court-documents/
[https://perma.cc/HL2U-EFHR].
13. Press Release, Consumer Fin. Prot. Bureau, CFPB Takes Action Against Debt Collector
for Pursuing Disputed and Unverified Cellphone Debts (Dec. 7, 2015),
https://www.consumerfinance.gov/about-us/newsroom/cfpb-takes-action-against-debt-collector-
for-pursuing-disputed-and-unverified-cellphone-debts/ [https://perma.cc/ER3N-N55G]; see CFPB
v. Collecto, Inc., No 1:15-cv-14024 (D. Mass. Dec. 8, 2015).
14. FED. TRADE COMM’N, REPAIRING A BROKEN SYSTEM: PROTECTING CONSUMERS IN DEBT
COLLECTION LITIGATION AND ARBITRATION, at i (2010),
https://www.ftc.gov/sites/default/files/documents/reports/federal-trade-commission-bureau-
consumer-protection-staff-report-repairing-broken-system-protecting/debtcollectionreport.pdf
[https://perma.cc/FP8S-T2GD] [hereinafter FTC PROTECTING CONSUMERS REPORT].
15. See Gary S. Becker, Crime and Punishment: An Economic Approach, 76 J. POL. ECON.
169, 170 (1968) (developing the foundations for the theory of optimal fines given limited
enforcement resources).
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
2018] ADMINIZATION 125
consumer protection agencies (the Federal Trade Commission (“FTC”)
or the CFPB); state attorney general offices and state level consumer
agencies; or some combination thereof.16 More important than the
institutional locus is the process itself. In a nutshell, the agency would
be notified of every incoming lawsuit.17 Using its administrative
powers, the agency will select claims to be audited by competent agency
investigators; where wrongdoing and abuse are found, the agency will
use its statutory powers to levy fines against wrongdoers.18 To manage
the millions of cases that are filed every year, the agency will select for
audit only a small fraction of the cases, similar to the Internal Revenue
Service (“IRS”).19 The selection of cases will be random—to expose every
creditor to a potential risk of detection—although the Article also
explains how machine learning and algorithmic analysis could be used
to improve the accuracy of the selection process by focusing the auditors’
attention on the cases that are statistically most likely to involve
16. For the CFPB’s powers, see Dodd-Frank Wall Street Reform and Consumer Protection
(Dodd-Frank) Act, Pub. L. No. 111-203, 124 Stat. 1376 (2010) (codified in scattered sections of title
12 of the United States Code), and in particular 12 U.S.C. § 5511 (2012) (authorizing the CFPB to
“seek to implement and, where applicable, enforce Federal consumer financial law,” and assigning
it the function of “collecting, investigating, and responding to consumer complaints.” The
provisions codified at 12 U.S.C. §§ 5562–65 provide the agency with the requisite regulatory
powers to “engage in investigations and request information from covered persons, issue
subpoenas or civil investigative demands, conduct hearings and adjudication proceedings, and
commence civil actions in federal court seeking any appropriate or equitable relief against any
person that violates a federal consumer financial law.”). For the FTC’s powers, see the Fair Debt
Collection Practices Act, 15 U.S.C. § 1692l (2012), which also provides federal protection from
unlawful debt collection activity in state courts, as expounded in numerous cases. See, e.g., Phillips
v. Asset Acceptance, LLC, 736 F.3d 1076, 1079 (7th Cir. 2013) (suing for stale debt in a state court
is an “unfair” debt collection practice); Kimber v. Fed. Fin. Corp., 668 F. Supp. 1480 (D. Ala. 1987)
(same); Fox v. Citicorp Credit Servs., Inc., 15 F.3d 1507 (9th Cir. 1994) (filing a writ of garnishment
can itself be unlawful); Morgan v. Credit Adjustment Bd., 999 F. Supp. 803 (E.D. Va. 1998)
(threatening to sue, without such an intention, is unlawful). It should be noted that at the time of
writing, there is a growing uncertainty over the future of the CFPB and the Dodd-Frank Act, but
nothing in what follows depends specifically on the CFPB itself. See Excerpts from the Times’s
Interview with Trump, N.Y. TIMES (July 19, 2017),
https://www.nytimes.com/2017/07/19/us/politics/trump-interview-transcript.html?mcubz=3
[https://perma.cc/A2W7-MZ24] (reporting President Trump’s remark that “Dodd-Frank is going to
be, you know, modified, and again, I want rules and regulations. But you don’t want to choke,
right?”).
17. For a chart illustrating the process, see infra Section II.A. A similar system exists in
employment discrimination cases, where plaintiffs (employees) file their claims first with an
administrative agency. See Age Discrimination in Employment Act of 1967, 29 U.S.C. § 626(d)
(2012); Civil Rights Act of 1964, 42 U.S.C. § 2000e-5 (2012); Americans with Disabilities Act of
1990, 42 U.S.C. § 12117(a) (2012).
18. See 12 U.S.C. § 5565 (outlining the CFPB’s statutory powers). On the magnitude of fines,
see infra note 116. Between 2012 and 2016, the CFPB issued over $5 billion in penalties. Doug
Johnson, Total CFPB Penalties Top $5B, INSIDEARM (Mar. 24, 2016, 8:18 AM),
https://www.insidearm.com/news/00041798-total-cfpb-penalties-top-5b/ [https://perma.cc/CH6F-
FWGP].
19. The IRS audits only 0.8 percent of all individual filings. See infra note 112.
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
126 VANDERBILT LAW REVIEW [Vol. 71:1:121
wrongdoing. Fraud detection analysis is used extensively and fruitfully
by credit card companies to detect real-time fraud, and these techniques
hold great promise for application to Adminization. Until such
algorithms are proven in practice, however, Adminization can fully
depend on random selection.
Adminization employs an administrative agency to gatekeep
civil litigation, thus departing from traditional approaches focused on
finding solutions within the court system. Ultimately, these approaches
fail because the adversarial model requires information sourced and
presented by the parties.20 However, getting consumers to participate
in the process is an elusive problem of immense proportions.21 This
Article criticizes the pretension of participation-based solutions to
meaningfully solve the problems at hand by relying on a pure litigation
model. To affect real change, participation-based solutions would
require prohibitively costly reforms,22 a dramatic expansion of the legal
system, and the creation of impossible delays to all other civil matters.23
To illustrate, recall that there are eight million cases filed every year,
with about 6.4 million resulting in a default judgment.24 Any change
that would lead to the screening of even half of these 6.4 million cases
would require state courts—which are already clogged—to handle an
additional 3.2 million cases every year.
20. See ROBERT A. KAGAN, ADVERSARIAL LEGALISM: THE AMERICAN WAY OF LAW 117 (2003)
(“[A]dversarial legalism often transforms the civil justice system into an engine of injustice . . . .”);
infra Part I.
21. See, e.g., ACCESS TO JUSTICE INITIATIVE, U.S. DEP’T OF JUSTICE, FOUR-YEAR ANNIVERSARY
ACCOMPLISHMENTS (2014) https://www.justice.gov/sites/default/files/atj/legacy/2014/03/14/
accomplishments.pdf [https://perma.cc/H49U-JXDK] (“[T]he current deficiencies in . . . legal
services for the poor and middle class constitute not just a problem, but a crisis. And this crisis
appears as difficult and intransigent as any now before us.” (internal quotation marks omitted)
(quoting Eric Holder, Attorney Gen., Remarks at the Shriver Center Awards Dinner (Oct. 14,
2010))); The Obama White House, The White House Forum on Increasing Access to Justice,
YOUTUBE (Apr. 19, 2016), https://www.youtube.com/watch?v=162foSVT2Uk [https://perma.cc/
T5W9-HN5M] (recording conference at the White House with leading politicians, jurists, and
businesspeople, aimed to explore avenues of increasing access to justice, emphasizing the need for
innovation); see also DIRECTV, Inc. v. Imburgia, 136 S. Ct. 463, 471 (2015) (Ginsburg, J.,
dissenting) (choosing the contractual interpretation that would promote access to justice);
Memorandum on Establishment of the White House Legal Aid Interagency Roundtable, 2015
DAILY COMP. PRES. DOC. 643 (Sept. 24, 2015) https://obamawhitehouse.archives.gov/the-press-
office/2015/09/24/presidential-memorandum-establishment-white-house-legal-aid-interagency
[https://perma.cc/G3PR-S36Y] (ordering the establishment of a large interagency work group
designed to “enhance access to justice in our communities”).
22. Deborah L. Rhode, Whatever Happened to Access to Justice?, 42 LOY. L.A. L. REV. 869,
869 (2009) (“[D]espite [the] efforts [of legal professionals], an estimated four-fifths of the individual
legal needs of the poor, and a majority of the needs of middle-income Americans, remain unmet.”).
23. For a discussion and critique of participation-oriented solutions and their limitations, see
infra Part III.
24. See infra note 42.
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
2018] ADMINIZATION 127
The Trump Administration is placing tremendous pressure on
the legal aid project as a whole, as most clearly demonstrated by a
recent budget proposal that seeks to eliminate funding to the Legal
Services Corporation—the nonprofit organization tasked with
supporting legal aid.25 In contrast to the traditional legal aid model,
Adminization avoids the need to subsidize consumer participation or to
review millions of cases, thus presenting the most economically and
politically appealing solution. The use of sampling techniques for case
audits—techniques used extensively by agencies but almost never by
courts—allows the managing of cases on a large scale with a limited
budget. With the potential for large fines, Adminization will deter the
initial filing of wrongful suits, thus reducing the volume of claims
overall and offsetting its operational costs. Synergistically, the
reduction in case volume will free up the courts to screen more closely
the cases that do come before them, further bolstering the effectiveness
of the process. The existence of long-standing federal agencies, such as
the FTC, the cost-effectiveness of Adminization, and the direct control
Congress can exert on the budget dedicated to Adminization (compared
with court budgets, which are harder to control), promises a real
potential for bipartisan appeal.26
This Article’s contribution may be understood on four levels of
abstraction. First, Adminization presents a normatively attractive and
politically feasible solution to the pressing problem of unmeritorious
claiming in the context of consumer litigation. Second, Adminization
provides a model for reducing abusive claiming in other areas of civil
litigation that suffer from systemic power asymmetries, such as
housing, social benefits, elder law, and employee rights. Third, the
Article explores a particularly promising implementation of artificial
intelligence that is well within our current technological abilities: case
selection. Machine learning algorithms have made significant strides
over the last decade, with the latest among many previously
unfathomable advances being the triumph of software in the intractable
25. Charles Toutant, Legal Services Worried That Trump Will Take Ax to Agency, LAW.COM
(Jan. 25, 2017), http://www.law.com/sites/almstaff/2017/01/25/legal-services-worried-that-trump-
will-take-ax-to-agency/?slreturn=20170109141618 [https://perma.cc/34YH-TUR7] (also available
at http://www.njlawjournal.com/id=1202777693184/Legal-Services-Worried-That-Trump-Will-
Take-Ax-to-Agency?slreturn=20170814123738 [https://perma.cc/XLA8-PKAW]); Debra Cassens
Weiss, Trump Budget Eliminates Legal Services Corp. Funding, A.B.A. J. (Mar. 16, 2017, 8:45
AM), http://www.abajournal.com/news/article/trump_budget_eliminates_funding_for_legal
_services_corp/ [https://perma.cc/47QS-K537].
26. Dave Boyer, Consumer Financial Protection Bureau in Jeopardy Under Donald Trump,
WASH. TIMES (Nov. 29, 2016), http://www.washingtontimes.com/news/2016/nov/29/consumer-
financial-protection-bureau-in-jeopardy-u/ [https://perma.cc/6C37-BCCF].
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
128 VANDERBILT LAW REVIEW [Vol. 71:1:121
game of Go.27 Concurrently, there is a growing strain on judicial
resources,28 and full civil trials are becoming nearly extinct.29 These two
trends suggest the need, and promise, of utilizing algorithmic
decisionmaking within the legal process.30 Fourth, drawing on David
Engstrom’s recent Litigation Gatekeeper theory,31 business and startup
theory,32 and institutional economics, Adminization creates a new
model of regulation and demonstrates the important, yet unexplored,
complementarities between courts and agencies. This is in contrast to a
pervasive view, prominent in the writings of Jerry Mashaw for
example,33 that administration and civil litigation are somehow
inconsistent with each other.34 By looking past this illusory dichotomy,
27. See Artificial Intelligence: Google’s AlphaGo Beats Go Master Lee Se-dol, BBC (Mar. 12,
2016), http://www.bbc.com/news/technology-35785875 [https://perma.cc/AS5C-ZF3Z] (describing a
computer program’s defeat of a master Go player). On recent applications of artificial intelligence
to the law, see JUDICIAL APPLICATIONS OF ARTIFICIAL INTELLIGENCE (Giovanni Sartor & Luther
Branting eds., 2013).
28. See I. Glenn Cohen, Rationing Legal Services, 5 J. LEGAL ANALYSIS 221, 221 (2013)
(comparing methods of distributing access to legal assistance); Matthew J.B. Lawrence, Procedural
Triage, 84 FORDHAM L. REV. 79, 82 (2015) (considering a method of distributing access to hearings).
29. See infra note 92.
30. See RICHARD SUSSKIND & DANIEL SUSSKIND, THE FUTURE OF THE PROFESSIONS: HOW
TECHNOLOGY WILL TRANSFORM THE WORK OF HUMAN EXPERTS 128 (2016) (“[T]here is a new
generation of machines in action now, and these are systems . . . that can replace parts of, and
sometimes all of, certain kinds of professional work.”).
31. See David Freeman Engstrom, Agencies as Litigation Gatekeepers, 123 YALE L.J. 616
(2013) (developing a theory of the functions agencies can play in husbanding litigation).
Importantly, Engstrom’s focus is not participation problems, but how agencies can address
inefficiencies in private enforcement that would tend to result in excessive litigation. In an
important sense, his perspective is plaintiff-centric and not, as in here, defendant-centric. Despite
these differences, his view of the symbiotic relationships between agencies and courts in private
litigation and his core typology provide a foundation for the proposal outlined here.
32. See ERIC RIES, THE LEAN STARTUP: HOW TODAY’S ENTREPRENEURS USE CONTINUOUS
INNOVATION TO CREATE RADICALLY SUCCESSFUL BUSINESSES 4 (2011) (presenting a new “lean
startup” approach to business). Recently, various government officials have started experimenting
with methods inspired by the lean startup model. See Justine Brown, Governments Take a Lean
Startup Approach, GOV’T TECH. (Aug. 23, 2012), http://www.govtech.com/pcio/Governments-Take-
a-Lean-Startup-Approach.html [https://perma.cc/PTD3-K33H] (describing Ries’s work with
United States Chief Technology Officer Aneesh Chopra).
33. See JERRY L. MASHAW, BUREAUCRATIC JUSTICE: MANAGING SOCIAL SECURITY DISABILITY
CLAIMS 222 (1983) (concluding that societal ideals represented by courts and legislatures
contradict bureaucratic rationality).
34. See, e.g., KAGAN, supra note 20, at 125 (arguing, critically, that the persistence of the
adversarial system is due to a tradition of suspicion toward “any alternative that smacks of
hierarchically organized bureaucratic legalism or expert judgment”). Furthermore, Lon Fuller
argued that the morality of adjudication itself depends on procedural passiveness and on not
considering any evidence not presented by the parties themselves. LON FULLER, THE PROBLEMS
OF JURISPRUDENCE 706–07 (1949). This is inconsistent with the active role of agencies, as shall be
elaborated. More recently, David Engstrom noted a tendency in recent scholarship to place
“exclusively administrative regulation on the one hand and unbridled private enforcement on the
other.” Engstrom, supra note 31, at 622.
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
2018] ADMINIZATION 129
it is possible to conceive of solutions that will better serve both our
individualistic and democratic ideals of justice.35
The design of Adminization is calculated to avoid some of the
major legal and economic hurdles that threaten alternative proposed
solutions. By using already existing regulatory agencies to support the
platforms, most legal and constitutional concerns are assuaged, as
these agencies already investigate fraud and have enforcement powers.
Admittedly, the audit process is costly, but the use of case selection
minimizes the costs, and, perhaps even more significantly, the expected
benefit of Adminization—deterring the filing of unmeritorious claims—
could completely offset those costs. That Adminization is budget
friendly is especially important in light of the costs involved in the
alternatives, which are almost prohibitively high. Regulatory capture
is always a concern with agencies, but here, the concern is minimized
because Adminization diversifies regulatory activity between the court
and the agency. Diversifying our modes of regulation has the benefit of
requiring lobbyists to spread their efforts thin, thus reducing the
effectiveness of their investments. Finally, certain aspects of
Adminization are designed to appeal to creditors, thus promising the
possibility of building a large supporting coalition. Creditors, at least
those with a long-term view of the market, would find value in a system
that garners greater consumer confidence and legitimacy in the credit
market, as such attitudes are linked to a higher propensity to borrow
and repay debts. Compared to both the participation-based solutions
and the status quo, Adminization prevails in terms of effectiveness,
cost, and most importantly, justice.
The Article unfolds in four main parts. Part I describes the
problem of abuse in consumer credit litigation. Part II lays out the
Adminization framework, outlines its general principles, and applies it
to consumer credit litigation. Part III explains why participation-based
solutions are unlikely to solve the problem at hand. Part IV examines
some of the main challenges to Adminization. The Article concludes by
reflecting on the contribution of Adminization to civil litigation and
considering some future applications.
35. See Engstrom, supra note 31, at 622 (proposing that civil litigation and administration
are not “either/or options, but rather the outer poles of a rich continuum”). For other leading
examples of the view that litigation is the exclusive domain for private disputes, see ERNEST J.
WEINRIB, THE IDEA OF PRIVATE LAW 10 (1st rev. ed. 2012) (arguing that the harms individuals
suffer must be resolved within private law institutions, and that deviating from that would be
“fundamentally at odds with the nature of the entire [private law] enterprise”); and Owen M. Fiss,
Against Settlement, 93 YALE L.J. 1073, 1078–82 (1984) (rejecting out-of-court settlement of
disputes on the grounds that they fail to respect procedural rights).
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
130 VANDERBILT LAW REVIEW [Vol. 71:1:121
I. ABUSE AND FRAUD IN CONSUMER CREDIT CONTRACTS
Consumer credit contracts are the result of the ubiquitous and
valuable agreements made between consumers and businesses that
allow consumers to pay in the future for services or products tendered
today, which typically involve credit cards, hospital bills, and utilities.36
If a consumer fails to pay (“defaults”), creditors will often engage in
some type of informal collection methods before they file a claim,
consisting of dunning letters,37 phone calls, and quite rarely, face-to-
face collection attempts.38 The main leverages used at this stage are
credit reporting,39 psychological pressure, and social and peer
pressure.40 While not all uncollected debts result in a lawsuit, many
do.41 In fact, most civil litigation consists of these lawsuits, with eight
36. The most common sources of consumer debt are motor vehicle loans and credit card,
student, and medical debts. See, e.g., Marina Vornovytskyy et al., Household Debt in the U.S.: 2000
to 2011, U.S. CENSUS BUREAU 2–3 (2013), https://www.census.gov/content/dam/Census/library/
working-papers/2011/demo/debt-highlights-2011.pdf [https://perma.cc/3U99-27U3] (reporting the
composition of secured debt, which includes motor vehicle loans, and unsecured debt, which
includes credit card, student, and medical debts). The average consumer owes mostly mortgage
debt and then student loans, auto loans, and credit card debts. See FED. RESERVE BANK OF N.Y.,
QUARTERLY REPORT ON HOUSEHOLD DEBT AND CREDIT 3 (May 2013),
https://www.newyorkfed.org/medialibrary/Interactives/householdcredit/data/pdf/DistrictReport_
Q12013 [https://perma.cc/YSG2-S89B].
37. To dun is “[t]o demand payment from (a delinquent debtor).” Dun, BLACK’S LAW
DICTIONARY (10th ed. 2014).
38. As early as 1968, it was observed: “Debt-collection involves the very minimum of face-to-
face contact.” P.E. Rock, Observations on Debt Collection, 19 BRIT. J. SOC. 176, 178 (1968).
39. The cost of a bad credit score can be substantial. To give an example, a thirty-year-old
consumer with a car loan of $18,000, $5,000 in credit card debt, and a $400,000 mortgage will pay
$250,000 more in interest if she has the worst credit score relative to a consumer with the best
credit score. See Kathy Kristof, An Easy Way to Figure the Cost of Bad Credit, CBS NEWS (Oct. 22,
2014, 5:20 AM), http://www.cbsnews.com/news/new-tool-calculates-the-cost-of-bad-credit/
[https://perma.cc/Y7JY-7TTY].
40. The collection of debt is fraught with difficulty, because even when consumers have
assets, they may choose to hide and shield them from collection. See generally Yonathan A. Arbel,
Shielding of Assets and Lending Contracts, 48 INT’L REV. L. & ECON. 26 (2016) (surveying asset-
shielding techniques and proposing a theory of consumer behavior). Creditors, on the other hand,
often attempt to collect debt using abusive and potentially illegal techniques, as evidenced by the
large volume of consumer complaints. See Consumer Complaint Database, CONSUMER FIN.
PROTECTION BUREAU, http://www.consumerfinance.gov/data-research/consumer-complaints/ (last
visited Oct. 22, 2017) [https://perma.cc/U9DJ-YTKJ] (recording thousands of consumer complaints
submitted weekly). There is some mixed evidence to suggest that the stigma associated with the
inability to pay one’s debt is on the decline. See David B. Gross & Nicholas S. Souleles, An
Empirical Analysis of Personal Bankruptcy and Delinquency, 15 REV. FIN. STUD. 319, 345 (2002)
(finding evidence suggestive of a decline in stigma). But see Kartik Athreya, Shame as It Ever Was:
Stigma and Personal Bankruptcy, FED. RES. BANK RICHMOND ECON. Q., Spring 2004, at 1, 3
(arguing that the decline in stigma is not supported by an expected rise in interest rates).
41. Payday lending is an industry that specializes in low-stakes, low-duration, high-risk
loans to consumers without access to more formal credit. In this industry, where consumers are
frequently underfunded and the stakes are low, about ten percent of the cases go to litigation. See
Amanda E. Dawsey et al., Non-Judicial Debt Collection and the Consumer’s Choice Among
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
2018] ADMINIZATION 131
million filings a year, and one in three Americans facing a potential
lawsuit.42
Consumer credit litigation is handled by an adversarial system
that espouses the “sporting theory of justice,”43 the idea that truthful
information will emerge from the clash of self-interested participation
by the parties. The adversarial system imagines a judge who “views the
case from a peak of Olympian ignorance,” and “[t]he ignorance and
Repayment, Bankruptcy and Informal Bankruptcy, 87 AM. BANKR. L.J. 1, 4 n.12 (2013) (finding
that, in Virginia in 2008, only 11,717 of 104,832 uncollectable checks resulted in consumers being
sued for nonpayment); see also Richard M. Hynes, Broke but Not Bankrupt: Consumer Debt
Collection in State Courts, 60 FLA. L. REV. 1, 21–24 (2008) (arguing that civil lawsuit filings have
been stable despite an increase in borrowing); Robert A. Kagan, The Routinization of Debt
Collection: An Essay on Social Change and Conflict in the Courts, 18 LAW & SOC’Y REV. 323, 325–
26 (1984) (showing a decline in debt collection litigation in various state supreme courts).
42. In 2010, about fifteen million lawsuits were filed in U.S. civil courts, with 1.8 million in
small claims courts alone. See Small Claims Fall Sharply in Last Two Years, CT. STAT. PROJECT,
http://www.courtstatistics.org/Civil/2012W5CIVIL.aspx (last visited Oct. 22, 2017)
[https://perma.cc/NG6X-N9TX]. The most common claims were for consumer credit, accounting for
forty to sixty percent of the docket. Hynes, supra note 41, at 49 (reporting rates of at least sixty
percent in Virginia); Mary Spector, Debts, Defaults and Details: Exploring the Impact of Debt
Collection Litigation on Consumers and Courts, 6 VA. L. & BUS. REV. 257, 273 (2011) (finding
similar rates in Texas); Healy, supra note 1 (finding sixty percent); see also URBAN JUSTICE CTR.,
DEBT WEIGHT: THE CONSUMER CREDIT CRISIS IN NEW YORK CITY AND ITS IMPACT ON THE WORKING
POOR 8 (2007), https://cdp.urbanjustice.org/sites/default/files/CDP_Debt_Weight.pdf
[https://perma.cc/25B4-CFBZ] [hereinafter DEBT WEIGHT] (reporting that over fifty percent of total
filings in New York City were for consumer credit transactions); Due Process and Consumer Debt:
Eliminating Barriers to Justice in Consumer Credit Cases, APPLESEED 1 (2010),
https://www.appleseednetwork.org/wp-content/uploads/2012/05/Due-Process-and-Consumer-
Debt.pdf [https://perma.cc/G5YS-JVDJ] [hereinafter Due Process and Consumer Debt] (“Hundreds
of thousands of consumer credit cases are filed and adjudicated each year in the five boroughs of
New York City alone . . . .”). This rate amounts, in New York City alone, to at least 300,000
lawsuits annually. DEBT WEIGHT, supra, at 1 (reporting 320,000 cases annually in five New York
City boroughs); Conor P. Duffy, A Sum Uncertain: Preserving Due Process and Preventing Default
Judgments in Consumer Debt Buyer Lawsuits in New York, 40 FORDHAM URB. L.J. 1147, 1148
(2013) (finding a range of 100,000 to 300,000 lawsuits by debt buyers); Michael Virtanen, New
Rules Established for NY Debt Collection Cases, WASH. TIMES (Sept. 16, 2014),
http://www.washingtontimes.com/news/2014/sep/16/new-rules-established-for-ny-debt-collection-
cases/ [https://perma.cc/9L2S-S857] (estimating 160,000 annual lawsuits in New York City in
2013). For context, the total amount of consumer debt in 2015 was $3.4 trillion, which suggests
the overall economic significance of this field of law. Federal Reserve Statistical Release: Consumer
Credit July 2017, FED. RES. 1 (Sept. 8, 2017), http://www.federalreserve.gov/releases/g19/
current/g19.pdf [https://perma.cc/GA89-NEYX]. Additionally, one in three Americans has an
account in collections and one in twenty has a credit obligation that is thirty days past due. See
FED. RESERVE BANK OF N.Y., QUARTERLY REPORT ON HOUSEHOLD DEBT AND CREDIT, at i (May
2016), https://www.newyorkfed.org/medialibrary/interactives/householdcredit/data/pdf/HHDC
_2016Q1.pdf [https://perma.cc/JD7F-TCHC] (reporting five percent of debts in delinquency,
accounting for $613 billion); Caroline Ratcliffe et al., Delinquent Debt in America, URB. INST. 4, 9
(July 30, 2014), https://www.urban.org/sites/default/files/publication/22811/413191-delinquent-
debt-in-america.pdf [https://perma.cc/WJR8-K9LJ] (using data from Transunion).
43. See Roscoe Pound, The Causes of Popular Dissatisfaction with the Administration of
Justice, 40 AM. L. REV. 729, 738 (1906).
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
132 VANDERBILT LAW REVIEW [Vol. 71:1:121
unpreparedness of the judge are intended axioms of the system,”44
because the parties are supposed to be the source of all information.45
However, when consumers systematically underparticipate—as is the
case with consumer credit litigation—this entire edifice crumbles,
inviting fraud, abuse, and overall nonmeritorious claiming.
A. Abuse and Procedural Violations
The general view among specialists and practitioners in the field
of consumer credit litigation is that abuse is pervasive, with multiple
pieces of evidence showing both procedural and substantive
violations.46 First, regulators consistently find banks and debt buyers
filing abusive and nonmeritorious lawsuits on a mass scale, forging
affidavits, inflating amounts owed, pursuing debt claims whose veracity
they themselves have reason to doubt, and filing questionable lawsuits
that have long passed the statute of limitations.47 These regulatory
findings often translate to fines in the tens of millions of dollars, and
they reveal broad practices. For example, the FTC found that twelve
percent of all the debts handled by large debt buyers lie beyond the
statute of limitations, which often also implies that neither the debt
buyer nor the consumer have a clear sense of whether the debt is real.48
Second, and highly revealing, is the number and nature of
consumer complaints about debt collection practices.49 Before debt
44. Marvin E. Frankel, The Search for Truth: An Umpireal View, 123 U. PA. L. REV. 1031,
1042 (1975).
45. See, e.g., William B. Rubenstein, The Concept of Equality in Civil Procedure, 23 CARDOZO
L. REV. 1865, 1874 (2002) (noting that in adversarial adjudication, equal participation is
“important . . . because it is thought to contribute to accurate and acceptable dispute resolution”).
46. See Judith Fox, Rush to Judgment: How the Fair Debt Collection Practices Act Fails to
Protect Consumers in Judicial Debt Collection, 13 FLA. ST. U. BUS. REV. 37, 37 (2014) (“The shoddy
evidence commonly presented to Indiana courts . . . is at the very least deceptive, and is often
abusive.”); Sam Glover, Has the Flood of Debt Collection Lawsuits Swept Away Minnesotans’ Due
Process Rights?, 35 WM. MITCHELL L. REV. 1115, 1117 (2009) (“Forty-one percent of the total
default judgments . . . were filed by debt buyers who probably could not prevail on the merits in
most, if not all, of those lawsuits.”); Lauren Goldberg, Dealing in Debt: The High-Stakes World of
Debt Collection After FDCPA, 79 S. CAL. L. REV. 711, 741–45 (2006) (arguing that lower standards
in small claims courts “offer collection lawyers a swift sword of judgment against debtors and give
lawyers leeway to file cases that would not survive in general civil court”); Peter A. Holland, Junk
Justice: A Statistical Analysis of 4,400 Lawsuits Filed by Debt Buyers, 26 LOY. CONSUMER L. REV.
179, 198–203 (2014) (reviewing the literature).
47. See supra notes 11–13 and accompanying text. Another recent example includes an action
against two law firms that turned a large volume of unverified lawsuits into default judgments.
See New Century Fin. Servs., Inc., CFPB No. 2016-CFPB-0010 (Apr. 25, 2016).
48. See FTC DEBT INDUSTRY REPORT, supra note 10, at iv–v, T-12 (finding that twelve percent
of debts are over six years old and that most states imposed three- to six-year statutes of
limitations).
49. Fred Williams, FIGHT BACK AGAINST UNFAIR DEBT COLLECTION PRACTICES 5 (Jeanne
Glasser et. al eds., 2011) (“The number-one complaint is that collectors are demanding money that
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
2018] ADMINIZATION 133
claims become lawsuits they undergo a debt collection process, and
consumers can file a complaint to the regulator if their debt is unfairly
handled. In 2014, hundreds of thousands of complaints were filed with
the FTC,50 and an additional eighty-five thousand with the CFPB.51 In
terms of substance, the bulk of these complaints concern allegedly
invalid or unverified debts, abusive communications, and illegal
threats.52 This conforms with the view of experts that the debt collection
process is rife with abuse, fraud, and unfair practices,53 and the results
of a financial survey where thirty-seven percent of respondents reported
being overcharged or deceived by a financial institution.54 Worryingly,
weak demographics, such as the elderly, are reported to be targeted
specifically.55
people do not even owe . . . .”). But see Letter from Donald S. Clark, Sec’y, Fin. Trade Comm’n, to
Richard Cordray, Dir., Consumer Fin. Prot. Bureau (Feb. 21, 2014),
https://www.ftc.gov/system/files/documents/reports/federal-trade-commission-enforcement-fair-
debt-collection-practices-act-report-consumer-financial/140305debtcollectionletter.pdf
[https://perma.cc/CD95-6E35] (reporting that in 2013, misrepresentation of debt is second most
common to repeated calls by debt collectors).
50. The FTC reports 280,998 complaints in 2014. Press Release, Fed. Trade Comm’n, Identity
Theft Tops FTC’s Consumer Complaint Categories Again in 2014 (Feb. 27, 2015),
https://www.ftc.gov/news-events/press-releases/2015/02/identity-theft-tops-ftcs-consumer-
complaint-categories-again-2014 [https://perma.cc/ZU7N-K8DS]. There have been some doubts
raised as to the accuracy of the FTC reports, because they are unverified and may include multiple
complaints made by the same debtor. Michael Klozotsky, The Facts Behind the Fantasies About
Debt Collection Complaints, FORBES (June 22, 2012, 11:26 AM),
http://www.forbes.com/sites/insidearm/2012/06/22/the-facts-behind-the-fantasies-about-debt-
collection-complaints/print/ [https://perma.cc/EJ3K-VBK5].
51. CONSUMER FIN. PROT. BUREAU, SEMI-ANNUAL REPORT OF THE CONSUMER FINANCIAL
PROTECTION BUREAU 26 (2015), http://files.consumerfinance.gov/f/201506_cfpb_semi-annual-
report-spring-2015.pdf [https://perma.cc/F42B-MDJS]. The CFPB database lists 39,148
complaints in 2014 regarding debt collection. Consumer Complaint Database, supra note 40 (click
“View complaint data” then filter the “Data received” to be between Jan. 1, 2014 and Dec. 31, 2014,
and filter the product column to “Debt collection”).
52. See FED. TRADE COMM’N, CONSUMER SENTINEL NETWORK DATA BOOK FOR JANUARY–
DECEMBER 2014, at 77 (2015), https://www.ftc.gov/system/files/documents/reports/consumer-
sentinel-network-data-book-january-december-2014/sentinel-cy2014-1.pdf
[https://perma.cc/9U86-T9GW].
53. See Goldberg, supra note 46, at 713 (“[C]orruption is running rampant in the collection
industry and federal collection law is ill-equipped to stop it.”); Justin P. Nichols, Dumping the Fair
Debt Collection Practices Act, 16 J. CONSUMER & COM. L. 26, 26 (2012) (noting the corruption in
the debt collection industry); Note, Improving Relief from Abusive Debt Collection Practices, 127
HARV. L. REV. 1447, 1447 (2014) [hereinafter Improving Relief] (arguing that millions of Americans
have been subject to predatory litigation techniques).
54. Telephone Survey of Likely Voters, AMS. FOR FIN. REFORM & CTR. FOR RESPONSIBLE
LENDING 27 (2013), http://www.responsiblelending.org/sites/default/files/uploads/2013-crl-afr-full-
poll-results-toplines-july-12.pdf [https://perma.cc/Z5UE-5SQ3].
55. See Matthew W. Ludwig, Abuse, Harassment, and Deception: How the FDCPA Is Failing
America’s Elderly Debtors, 16 ELDER L.J. 135, 151–56 (2008) (detailing the targeted abuse of the
elderly); see also Goldberg, supra note 46, at 736–39 (describing targeted debt collection based on
sex, age, and income).
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
134 VANDERBILT LAW REVIEW [Vol. 71:1:121
The third piece of evidence of the level of abuse comes, ironically,
from the lack of evidence in a large fraction of all lawsuits.56 One judge
estimated that plaintiffs lack necessary evidence in ninety percent of
the cases,57 and another judge mused that many claims “lack a nano of
a modicum of a scintilla of a prima facie case so as to be entitled to a
judgment whether it be by default or otherwise.”58 An empirical study
found no evidence at all in forty-six percent of cases,59 and a recent
study showed that many debt buyers do not bother to acquire evidence
in the first place, buying debts that they have never verified
themselves.60 When evidence is produced, its quality tends to be very
poor. One study found a breakdown of the claimed debt to its principal,
interest, and other charges in only five percent of the cases. Information
regarding payment history and the date of default were likewise
missing.61 Moreover, much (arguably most) of the evidence that is
brought is “facially invalid,” as a study of six hundred cases found.62
This is congruent with the (potentially illegal) practice of “robo-
signing,” namely the automated signing of mass volumes of documents
without actual review, which many view as a major concern.63 Of
56. The lack of evidence is part of a broad industry practice of not producing evidence to
support debts. See Duffy, supra note 42, at 1162 (“Portfolios often lack essential collection
information . . . .”).
57. See Holland, supra note 46, at 184 (citing Jessica Silver-Greenberg, Problems Riddle
Moves to Collect Credit Card Debt, N.Y. TIMES: DEALBOOK (Aug. 12, 2012, 9:09 PM),
https://dealbook.nytimes.com/2012/08/12/problems-riddle-moves-to-collect-credit-card-debt/
[https://perma.cc/LTX8-HX42], which quotes Noach Dear, a civil court judge in Brooklyn). This
should not be read as saying that ninety percent of cases are fraudulent, only that creditors do not
find it cost-effective to produce evidence in light of the low rates of defendants’ appearances.
58. Am. Express Bank, FSB v. Dalbis, No. 300082/10, 2011 WL 873512, at *12 (N.Y. Civ. Ct.
Mar. 14, 2011) (internal quotation marks omitted).
59. Fox, supra note 46, at 45–46. Fox further notes that in the remaining cases, evidence was
sometimes completely fabricated. Id. at 46.
60. CONSUMER FIN. PROT. BUREAU, STUDY OF THIRD-PARTY DEBT COLLECTION OPERATIONS
22 (2016), http://files.consumerfinance.gov/f/documents/20160727_cfpb_Third_Party_Debt
_Collection_Operations_Study.pdf [https://perma.cc/3UWJ-VVG8] (finding in a survey of debt
buyers that evidence beyond that required to identify the debtor is often not acquired); see also
Peter A. Holland, The One Hundred Billion Dollar Problem in Small Claims Court: Robo-Signing
and Lack of Proof in Debt Buyer Cases, 6 J. BUS. & TECH. L. 259, 262 (2011) (discussing lack of
evidence); Rachel Terp & Lauren Bowne, Past Due: Why Debt Collection Practices and the Debt
Buying Industry Need Reform Now, CONSUMERS UNION 4–5 (Jan. 2011),
http://consumersunion.org/pdf/Past_Due_Report_2011.pdf [https://perma.cc/C3CS-WMEF]
(same).
61. Spector, supra note 42, at 291. Even attorney fees were explicitly itemized in only thirty
percent of the cases. Id. at 292.
62. DEBT WEIGHT, supra note 42, at 7, 9 (reporting that in “99.0% of the cases where default
judgments were entered, the materials underlying those applications constituted inadmissible
hearsay”). The main fault in most cases was affidavits signed by people with no personal
knowledge of the underlying debt. Id. at 20.
63. See 1 ROBERT J. HOBBS ET AL., NAT’L CONSUMER LAW CTR., FAIR DEBT COLLECTION
§ 5.5.2.13.4 (7th ed. 2011 & Supp. 2013) (noting that courts are split on whether robo-signing
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
2018] ADMINIZATION 135
course, absence of evidence is not evidence of absence.64 After all,
evidence is costly to produce, and so it may not pay to produce it if cases
are not scrutinized.65 Nonetheless, the reality is that a great deal of debt
is owned by parties who did not take part in the original transaction
and have no knowledge that the debt is real but at the same time have
strong financial incentive to try to collect it. The lack of evidence is thus
strongly suggestive of nonmeritorious lawsuits.
Finally, there is strong evidence of abuse in the process of
notifying consumers of lawsuits. As a rule, plaintiffs are required to
notify consumers of the lawsuit by serving them with a court
summons.66 Unfortunately, this rule engenders perverse incentives: if
the consumer fails to attend the hearing, the plaintiff is almost assured
to win the case. The result of this badly designed system of incentives
is manifested in the phenomenon of “sewer service”: the practice among
debt collectors of figuratively dumping the summons in the sewer while
signing an affidavit that alleges actual service. While it is hard to
gather evidence on the scope of this phenomenon,67 the evidence that
does exist points at a broad problem. For example, the New York Bar
estimates that each year sewer service affects “tens of thousands” of
New Yorkers,68 and a New York judge said that, in his view, an
violates the FDCPA); Matthew J. Petrozziello, Note, Who Can Enforce? The Murky World of Robo-
signed Mortgages, 67 RUTGERS U. L. REV. 1061, 1082 (2015) (finding limited judicial acceptance of
robo-signing as a serious violation of FDCPA); see also Improving Relief, supra note 53, at 1450
(“Robosigning represents a particularly significant threat to consumers . . . .”); Press Release,
Consumer Fin. Prot. Bureau, supra note 12 (reporting on banks allegedly involved in robo-signing).
64. However, lack of evidence may be suggestive of lack of merit and is obviously consistent
with it. For similar reasoning, see, for example, DEBT WEIGHT, supra note 42, at 7 (“[T]he debt
buyers’ consistent failure to provide relevant evidence in support of their claims suggests that they
do not possess such evidence.”). But this conclusion is too strong; evidence is costly to produce and
if most consumers do not contest cases, it is not worthwhile to produce it, even for cases with merit.
65. From an economic standpoint, evidence is only valuable instrumentally as measured by
its ability to influence outcomes. Because evidence is costly to produce, when we require evidence
from the parties, we face a trade-off between greater accuracy and greater costs. See Louis Kaplow,
Information and the Aim of Adjudication: Truth or Consequences?, 67 STAN. L. REV. 1303 (2015)
(arguing that overall consequences of judicial decisions, not the pursuit of truth, should be the
primary goal of the legal system).
66. FED. R. CIV. P. 4(c)(1) (“The plaintiff is responsible for having the summons and complaint
served within the time allowed by Rule 4(m) and must furnish the necessary copies to the person
who makes service.”). Notice is an essential part of due process. See Administrative Procedure Act,
5 U.S.C. §§ 554, 556–557 (2012) (requiring notice to parties to an agency hearing of the time, place,
and nature of the hearing; legal authority under which the hearing is held; and matters of fact and
law asserted); Goldberg v. Kelly, 397 U.S. 254, 266–70 (1970) (holding that procedural due process
requires adequate notice before terminating public welfare program).
67. Spector, supra note 42, at 287 (“Little information regarding non-service exists . . . .”).
Consumers may have an incentive to exaggerate claims of service failures.
68. N.Y.C. BAR ASS’N, OUT OF SERVICE: A CALL TO FIX THE BROKEN PROCESS SERVICE
INDUSTRY 11–12 (2010), http://www.nycbar.org/pdf/report/uploads/ProcessServiceReport4-10.pdf
[https://perma.cc/ 76KJ-G3P4] [hereinafter OUT OF SERVICE].
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
136 VANDERBILT LAW REVIEW [Vol. 71:1:121
“astonishing” amount of default judgments are the result of faulty
service.69 Indeed, a recent class action alleging sewer service in New
York recently settled for $59 million.70 More systematic studies found
similar indications. In one study of 350 consumers, none were properly
served.71 Another found service in only twelve percent of the cases,72
and a larger one found that faulty service was a cause for dismissal in
about twenty-one percent of the cases studied.73 This problem is hardly
new; a report from 1968 made by the U.S. Attorney’s Office for the
Southern District of New York claims that at least half of all default
judgments entered in the Civil Court for the County of New York were
supported by false affidavits of service.74 Even when service takes place,
it is poorly done. One study finds in a sample of ninety-one cases that
almost no summonses were served in person. Instead, the vast majority
of summonses were served either by “nail and mail” (i.e., affixing the
summons to the defendant’s door) or by delivery to a different individual
in the household.75 These methods were designed as last resorts, but
apparently some servers practice them frequently. This study showed
that while two law firms did not serve any debtor in person, another—
which presumably tried harder—successfully served eighteen percent
of its sample cases personally.76
Taken together, this evidence suggests a serious problem.
Skeptics, however, may worry that some of the evidence is only
anecdotal, that some of the violations are only formal, and that some of
the research is subject to methodological problems. Primarily, the
absence of evidence, and even sewer service, is not definite proof that
the underlying claim is unmeritorious. These concerns are not without
merit individually, but a broader look may assuage them. The
69. Due Process and Consumer Debt, supra note 42, at 12.
70. Benjamin Mueller, Victims of Debt Collection Scheme in New York Win $59 Million in
Settlement, N.Y. TIMES (Nov. 13, 2015), https://www.nytimes.com/2015/11/14/nyregion/victims-of-
debt-collection-scheme-in-new-york-win-59-million-in-settlement.html [https://perma.cc/S8D4-
QKV2]. The relevant class action is Sykes v. Mel S. Harris & Associates LLC, 780 F.3d 70 (2d Cir.
2015).
71. Justice Disserved, MFY LEGAL SERVS. 2 (2008), http://mobilizationforjustice.org/wp-
content/uploads/reports/Justice_Disserved.pdf [https://perma.cc/Q4C3-T4AN]. Another study
found that four out of fifteen surveyed consumers were not served. Hillard M. Sterling & Philip G.
Schrag, Default Judgments Against Consumers: Has the System Failed?, 67 DENV. U. L. REV. 357,
370 (1990).
72. Spector, supra note 42, at 287 (studying a sample of 507 cases).
73. Holland, supra note 46, at 210 (finding dismissal for lack of service in 925 out of 4,400
sampled cases).
74. Frank M. Tuerkheimer, Service of Process in New York City: A Proposed End to
Unregulated Criminality, 72 COLUM. L. REV. 847, 849 (1972).
75. Justice Disserved, supra note 71, at 5.
76. Id. The study also indicates that creditors vary considerably in their service method,
whether in person or by “nail and mail.” Id.
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
2018] ADMINIZATION 137
consistency of the evidence across studies, cases, and even anecdotes,
coupled with the experience of industry insiders and regulators, all
point toward the conclusion that fraud and abuse in consumer credit
litigation is a serious problem. The absence of contrary studies is not
strong evidence, but it is also relevant. And perhaps strongest of all, on
simple theoretical grounds of moral hazard, we would expect the
existence of financial incentives combined with weak consumer and
judicial supervision to breed significant abuse. It is with this in mind
that I now turn my attention to the role of consumers and judges in
providing adequate monitoring of creditor behavior.
B. Justice, Inaccessible
Consumers often find the courts inaccessible, resulting in low
levels of response to claims, appearance in court, and legal
representation. Even the most basic step of responding to lawsuits is
rarely taken: consumers respond to only five to twenty-three percent of
lawsuits77 (compared to seventy-two percent in tort cases78). Similarly,
consumers appear in only seven to twenty percent of cases.79
Representation rates stand at a much lower rate of only two to 8.7
percent overall (but forty-three percent of cases where the defendant
77. The defendant normally has three weeks to file an answer. See FED. R. CIV. P. 12(a)(1)
(twenty-one days); Fed. Mar. Comm’n v. S.C. State Ports Auth., 535 U.S. 743, 757 (2002) (noting a
common twenty-day period in Federal Maritime Commission administrative proceedings). This is
viewed as an important right. See Nelson v. Adams USA, Inc., 529 U.S. 460, 466 (2000) (“[The]
opportunity to respond, fundamental to due process, is the echo of the opportunity to respond to
original pleadings secured by Rule 12.”). On answer rates, see Judith Fox, Do We Have a Debt
Collection Crisis? Some Cautionary Tales of Debt Collection in Indiana, 24 LOY. CONSUMER L. REV.
355, 377 (2012) (3.6 percent); Holland, supra note 46, at 186 (less than twenty percent); Spector,
supra note 42, at 288 (22.87 percent); Due Process and Consumer Debt, supra note 42, at 2 (0.8 to
7.2 percent). In arbitration, consumers answer in roughly seventy percent of the cases. See CAROL
J. DEFRANCES & STEVEN K. SMITH, U.S. DEP’T OF JUSTICE, CONTRACT CASES IN LARGE COUNTIES
6 (1995), https://www.bjs.gov/content/pub/pdf/ccilc.pdf [https://perma.cc/LT4U-RYS2].
78. STEVEN K. SMITH ET AL., U.S. DEP’T OF JUSTICE, TORT CASES IN LARGE COUNTIES 1 (1995),
https://www.bjs.gov/content/pub/pdf/TCILC.PDF [https://perma.cc/Z9TL-ZXZT].
79. See Holland, supra note 46, at 208 (ten percent); Spector, supra note 42, at 288 (twenty
percent); see also SUSAN SHIN & CLAUDIA WILNER, NEW ECON. PROJECT, THE DEBT COLLECTION
RACKET IN NEW YORK: HOW THE INDUSTRY VIOLATES DUE PROCESS AND PERPETUATES ECONOMIC
INEQUALITY 14 (2013), http://www.neweconomynyc.org/wp-content/uploads/2014/08/
DebtCollectionRacketUpdated.pdf [https://perma.cc/CBF4-KZZV] (noting that eighteen percent of
consumers sued in New York City appeared in court and only seven percent of consumers sued
outside of New York City appeared in court); Sterling & Schrag, supra note 71, at 361 (finding
twenty-two percent appearance rate). But see Mary Spector & Ann Baddour, Collection Texas-
Style: An Analysis of Consumer Collection Practices in and out of the Courts, 67 HASTINGS L.J.
1427, 1462 (2016) (finding in Texas courts appearance in fifty-two percent of the cases).
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
138 VANDERBILT LAW REVIEW [Vol. 71:1:121
chose to appear).80 For a sense of magnitude, in the state of New York
alone, 1.8 million litigants proceeded pro se in 2014.81 At the same time,
creditors are almost always represented, an advantage that carries over
to settlement agreements.82 For example, Jeff Cook, an unemployed
plumber, signed off $651 out of his (legally protected and uncollectable)
unemployment benefits due to ignorance of his legal rights and pressure
from the creditor.83 Moreover, pro se debtors also impose costs on the
system, as the court has to deal with motions and requests which often
deviate from the standard of filings common among lawyers.84 Even the
more informal small claims courts present access problems, as they
relax traditional procedural safeguards, such as the rule prohibiting
hearsay, while allowing the plaintiff legal representation.85
There are several complementary explanations for participation
gaps: the lack of resources, sophistication, and legal knowledge;86
80. DEBT WEIGHT, supra note 42, at 16 (two out of 600 cases); Holland, supra note 46, at 187
(less than two percent); Spector, supra note 42, at 289 (8.68 percent, but 43.14 percent of those
who appeared).
81. PERMANENT COMM’N ON ACCESS TO JUSTICE, STATE OF N.Y. UNIFIED COURT SYS., REPORT
TO THE CHIEF JUDGE OF THE STATE OF NEW YORK 3, 24 (2015), http://nylawyer.nylj.com/adgifs/
decisions15/122915report.pdf [https://perma.cc/JV5C-FVQQ] [hereinafter N.Y. ACCESS TO JUSTICE
REPORT].
82. See DEBT WEIGHT, supra note 42, at 16 (“[One hundred percent] of plaintiffs initiating
consumer credit transaction cases reviewed in our study were represented by counsel . . . .”); Duffy,
supra note 42, at 1175 (noting that in New York courts, “100% of debt collector plaintiffs are
represented by counsel”). On abuse in settlements, see Jessica Silver-Greenberg, In Debt
Collecting, Location Matters, WALL ST. J. (July 18, 2011), https://www.wsj.com/articles/
SB10001424052702303365804576433763597389214 [https://perma.cc/6FXR-295D] (describing a
debtor having an unsupervised meeting with the creditor’s attorney, leading to the debtor letting
the creditor tap into his unemployment benefits); Due Process and Consumer Debt, supra note 42,
at 18 (observing that “[p]laintiffs’ counsel may pressure unrepresented defendants into
unfavorable settlements”). See also Fiss, supra note 35, at 1078–82, for a criticism of settlements
in civil trials generally, partly on the ground of imbalance of powers between the parties.
83. See Silver-Greenberg, supra note 82.
84. See MASS. PROB. & FAMILY COURT DEP’T, PRO SE LITIGANTS: THE CHALLENGE OF THE
FUTURE 12–16 (1997), http://www.mass.gov/courts/docs/courts-and-judges/courts/probate-and-
family-court/prosefinalreport.pdf [https://perma.cc/BR7A-6K6M] (discussing problems with pro se
litigants).
85. See Holland, supra note 60, at 263. In most states, plaintiffs may be represented by a
lawyer even in a small claims court, and while defendants may also be represented, this is
infrequent. See NAT’L INST. OF JUSTICE, U.S. DEP’T OF JUSTICE, SMALL CLAIMS COURT REFORM 6
(1983), https://www.ncjrs.gov/pdffiles1/Digitization/93351NCJRS.pdf [https://perma.cc/8LUV-
BF9J] (noting the power asymmetry between plaintiffs and pro se defendants in small claims
courts).
86. See KAGAN, supra note 20, at 122–24 (analyzing asymmetries in knowledge, wealth, and
sophistication); Victoria J. Haneman, The Ethical Exploitation of the Unrepresented Consumer, 73
MO. L. REV. 707, 711 (2008) (arguing that the current adversarial system exacerbates power
imbalances between represented creditors and unrepresented debtors); David Rosenberg &
Kathryn E. Spier, Incentives to Invest in Litigation and the Superiority of the Class Action, 6 J.
LEGAL ANALYSIS 305 (2014) (analyzing stake asymmetry and its distortive effects). See generally
Marc Galanter, Why the “Haves” Come Out Ahead: Speculations on the Limits of Legal Change, 9
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
2018] ADMINIZATION 139
problems with service;87 psychological barriers and biases;88 and power
asymmetries. But beyond this, perhaps the deepest reason for consumer
apathy to the legal process is that such apathy is often rational. That
is, the costs of full participation often exceed the potential benefits. A
first obstacle for most Americans is time. Appearing in court involves
taking a day off work, which spells a potential loss of $136 to the median
American, assuming she can obtain her employer’s permission. Besides
this cost, individuals must pay for travel, preparation, and most
significantly, representation. The average hourly cost of a consumer law
attorney is $361.89 Even assuming one finds a cheaper attorney with a
rate of, say, $200 per hour, handling a standard case will often take four
to eight hours, thus leading to a total cost of $800–$1,600 for an average
case. This cost is very close to the value of the case itself—a typical case
involves a debt of $3,000 ($820 in a small claims court).90 And because
lawyers do not guarantee a win, but must be paid in advance, their
value to consumers is quite doubtful, especially when one takes into
account risk aversion and liquidity constraints. Overall, then,
participation is a very costly and doubtful endeavor for many.
Proceeding pro se may save costs—and indeed, many consumers choose
this option—but it is still an involved and stressful experience that
presents consumers with many potential pitfalls.
LAW & SOC’Y REV. 95 (1974) (arguing that litigation gives a systematic advantage to sophisticated
players).
87. See supra notes 66–75 and accompanying text.
88. Sterling and Schrag tell of a case where a default judgment was entered despite the debtor
being present in court: when her name was called, the debtor got too nervous and preferred to stay
quiet. See Sterling & Schrag, supra note 71, at 369. One dominant psychological bias which may
be of relevance here is the tendency to overly discount future outcomes. See David Laibson, Golden
Eggs and Hyperbolic Discounting, 112 Q.J. ECON. 443, 445–46 (1997).
89. RONALD L. BURDGE, UNITED STATES CONSUMER LAW ATTORNEY FEE SURVEY REPORT
2013–2014, at 11 (2015), http://burdgelaw.com/NACA/US-Consumer-Law-Attorney-Fee-Survey-
Report-2015.pdf [https://perma.cc/Z7UV-XRV7]. The typical fee charged by an attorney “can range
from $500 to negotiate a simple credit card debt to more than $5,000 for more complex
negotiations.” Baran Bulkat, How Much Will a Lawyer Charge to Negotiate with My Creditors?,
NOLO, http://www.nolo.com/legal-encyclopedia/how-much-will-lawyer-charge-negotiate-with-my-
creditors.html (last visited Oct. 22, 2017) [https://perma.cc/D85K-U44M].
90. See Suzanne E. Elwell & Christopher D. Carlson, The Iowa Small Claims Court: An
Empirical Analysis, 75 IOWA L. REV. 433, 510 (1990) ($820 in small claims, CPI adjusted); Holland,
supra note 46, at 206 ($2,993.17). The average value of debts in collection is $1,387. Ctr. for
Microeconomic Data, Data Bank, FED. RES. BANK N.Y., https://www.newyorkfed.org/
microeconomics/databank.html (last visited Oct. 22, 2017) [https://perma.cc/4UAB-YKYK] (under
the “Credit Cards” section, click on “Delinquencies,” then click on the first link, “Quarterly Report
on Household Debt and Credit,” which will open an Excel spreadsheet, then go to page 18 of the
spreadsheet).
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
140 VANDERBILT LAW REVIEW [Vol. 71:1:121
C. Lack of Judicial Oversight
In the current system, the main safeguard against the filing of
abusive claims is judicial screening, but most cases are reduced to
default judgments with little judicial oversight.91 Civil trials are on the
verge of extinction, with full trials taking place in less than two percent
of cases,92 and judges doubting the need for factual examinations in
cases of consumer credit.93 Even more rudimentary examinations are
rare, and while the rates of default judgment vary considerably, it is
common to find that eighty percent of cases result in default
judgments.94 If multiplied by all relevant cases, this implies that 6.4
million cases of consumer credit every year turn into default judgments
with little judicial scrutiny.95 The minority of cases that are heard do
not follow any clear pattern, and it is unclear whether those are the
most deserving ones or simply ones where the consumer had sufficient
resources, grit, or conviction to appear. This leads to highly limited
judicial oversight.
Three factors contribute to limited oversight: First, the
adversarial nature of the process limits judges’ investigative authority,
thus exacerbating the informational problems resulting from consumer
91. Robert G. Bone, Procedure, Participation, Rights, 90 B.U. L. REV. 1011, 1015–16 (2010)
(considering the role of accuracy and procedural participation rights under various theories of
procedure).
92. On the “vanishing trial” phenomenon in civil litigation, see generally Marc Galanter, The
Vanishing Trial: An Examination of Trials and Related Matters in Federal and State Courts, 1 J.
EMPIRICAL LEGAL STUD. 459 (2004). See also John H. Langbein, The Disappearance of Civil Trial
in the United States, 122 YALE L.J. 522, 551–53 (2012) (noting the trend and claiming pretrial
procedure has made trials obsolete). Since the publication of Galanter’s work, the rate of civil trials
has declined from 0.6 percent in state courts to around 0.27 percent. See Court Statistics Project,
Court Statistics Project Data Viewer, NAT’L CTR. ST. CTS., http://www.ncsc.org/Sitecore/Content/
Microsites/PopUp/Home/CSP/CSP_Intro (last visited Oct. 22, 2017) [https://perma.cc/SYJ6-
PFAU]. In the consumer credit context, see SHIN & WILNER, supra note 79 (0 out of 200,000 cases);
Fox, supra note 46, at 44 (0 out of 1,000 cases); Holland, supra note 46, at 213 (21 out of 2,947
cases); Spector, supra note 42, at 297 (1 out of 446 cases); Debt Deception: How Debt Buyers Abuse
the Legal System to Prey on Lower-Income New Yorkers, LEGAL AID SOC’Y ET AL. 8 (May 2010),
http://mobilizationforjustice.org/wp-content/uploads/reports/DEBT-DECEPTION.pdf
[https://perma.cc/LY2Z-HRE4] [hereinafter Debt Deception] (0 cases in a sample of 336 cases in
New York courts). Taken together, this amounts to 22 out of 204,729 cases where a trial was
conducted.
93. See FTC PROTECTING CONSUMERS REPORT, supra note 14, at 7 & n.18 (estimated default
judgment rate of sixty to ninety-five percent); DEBT WEIGHT, supra note 42, at 9 (eighty percent
default judgment rate); Spector & Baddour, supra note 79, at 1449 (31.6 percent default judgment
rate in Texas). Most remaining cases are dismissed (commonly without prejudice), transferred, or
settled.
94. See supra note 93.
95. See supra note 42.
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
2018] ADMINIZATION 141
inexperience, rational apathy, and psychological barriers.96 Second,
creditors are repeat players and can more effectively scale their
experience and engage in forum shopping.97 Third, the overload of civil
courts’ dockets makes it difficult for judges to spend sufficient time
scrutinizing cases.98 All of these structural issues contribute to a low
level of judicial scrutiny.
On the outskirts of the judicial process are private settlements
in the courthouse. Troublingly, these often produce worse results for
consumers than they could expect under the law.99 Plaintiffs’ attorneys
are reported to often play a negative role in such settlements,
misinforming debtors of their rights and applying pressure.100 Judges
rarely scrutinize the resulting agreements and often rubberstamp
them, turning them into enforceable agreements.101
Overall, the system of handling consumer debt is an incubator
of abuse. Consumers are largely apathetic to the process and do not
respond to lawsuits or show up to hearings. Creditors routinely bring
nonmeritorious lawsuits that are neither verified nor supported by
evidence, and judges do not try cases or provide judicial oversight of
cases. The few cases that do receive scrutiny are haphazardly chosen
with no rationale or logic. This provides companies and debt collectors
with incentive to inflate their claims and bring bogus charges, and the
evidence we have suggests that this happens on a large scale. The
system affects millions of consumers and yet is deeply and inexcusably
96. See Haneman, supra note 86, at 720–21 (exploring how the adversarial system harms
unrepresented consumers); see also Amalia D. Kessler, Our Inquisitorial Tradition: Equity
Procedure, Due Process, and the Search for an Alternative to the Adversarial, 90 CORNELL L. REV.
1181, 1183–87 (2005) (tracing the origins of the adversarial procedure in American law).
97. See, e.g., Glover, supra note 46, at 1125 (“In Hennepin County, 76% of the total filings
were by original creditors or debt buyers who filed twenty-five or more lawsuits as of August
2008.”). On repeat players, see Galanter, supra note 86, at 97–104 (explaining that repeat players
enjoy advantages in litigation and have a systematic advantage over one-shotters); Leslie G.
Kosmin, The Small Claims Court Dilemma, 13 HOUS. L. REV. 934, 942–43 (1976) (explaining that,
even in small claims courts, unsophisticated debtors face a disadvantage). But see Assaf Hamdani
& Alon Klement, The Class Defense, 93 CALIF. L. REV. 685, 689–90 (2005) (proposing consolidation
of defendants to increase the incentive to defend them).
98. See Suein Hwang, Once-Ignored Consumer Debts Are Focus of Booming Industry, WALL
ST. J. (Oct. 25, 2004, 11:59 PM), https://www.wsj.com/articles/SB109865776922954118
[https://perma.cc/Z6W3-P6K].
99. See supra note 42.
100. See Russell Engler, Out of Sight and Out of Line: The Need for Regulation of Lawyers’
Negotiations with Unrepresented Poor Persons, 85 CALIF. L. REV. 79, 82 (1997) (reviewing
attorneys’ roles in improper negotiations with poor debtors).
101. Russell Engler, And Justice for All—Including the Unrepresented Poor: Revisiting the
Roles of the Judges, Mediators, and Clerks, 67 FORDHAM L. REV. 1987, 2019–20 (1999).
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
142 VANDERBILT LAW REVIEW [Vol. 71:1:121
flawed. As the FTC recognized: “[N]either litigation nor arbitration
currently provides adequate protection for consumers.”102
II. ADMINIZATION
Adminization is a model of civil litigation that is designed to
cost-effectively add oversight to the system. Section A lays out the main
principles of Adminization, Section B explores its main features, and
Section C applies it to consumer credit litigation.
A. Adminization: High-Level Outline
Parallel to civil litigation, we have an administrative system
that does not depend on user participation for its operation and
acquisition of information. When the police, the IRS, the Securities and
Exchange Commission (“SEC”), or the United States Department of
Agriculture—to give but a few examples—engage in their regulatory
activities, they do so on their own initiative, harnessing their expertise
and investigative powers.103 They do not wait for the regulated entities
to “participate”; rather, they independently seek and gather relevant
information. These agencies do not even need a complaint to start their
process; it is the agency itself that chooses when to intervene. Because
administrative agencies do not depend on participation to identify and
screen bad cases, they offer great promise for a system that suffers from
a participation problem.
The core idea underlying Adminization is that by tapping into
the powers of agencies, it will be possible to provide a threshold level of
consumer protection that is independent of consumer participation.
Adminization consists of a gatekeeper agency that uses its
administrative powers—most notably sampling, audits, and fines—to
investigate cases and sanction plaintiffs who file baseless claims. This,
in a nutshell, protects consumers and reduces the volume of unwanted
litigation. The following figure illustrates the operation of the agency in
the context of consumer credit litigation, with each of the steps and
features explained in detail later in this Part.
102. Holland, supra note 46, at 188 (quoting FTC PROTECTING CONSUMERS REPORT, supra note
14).
103. See, e.g., How Criminal Investigations Are Initiated, INTERNAL REVENUE SERV.,
https://www.irs.gov/uac/how-criminal-investigations-are-initiated (last visited Oct. 22, 2017)
[https://perma.cc/5678-6SBN] (explaining the process for initiating criminal investigations).
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
2018] ADMINIZATION 143
FIGURE 1: THE ADMINIZATION WORKFLOW
Before moving to cover the details, it is worth considering
Adminization from a jurisprudential perspective. The idea of
Adminization challenges the traditional view that posits a tension
between the “individualized justice” of civil litigation and the generic
and less equitable “bureaucratic management” by agencies.104 Under
this view, administration and litigation are understood as
multidimensional polar opposites, each on the other side of ex ante vs.
ex post regulation, proactive vs. reactive, rule-driven vs. standard-
driven, specialized vs. generalist judgment, public vs. private
enforcement, and government vs. individual disputes.105 However, this
emphasis on tensions hides much that is complementary between the
two systems. Recently, David Engstrom developed a theory of agencies
as litigation gatekeepers, which is focused on the productive coexistence
of courts and agencies.106 There are many instances of such peaceful
104. See MASHAW, supra note 33, at 222; see also JERRY L. MASHAW ET AL., ADMINISTRATIVE
LAW, THE AMERICAN PUBLIC LAW SYSTEM: CASES AND MATERIALS 310 (7th ed. 2014).
105. Richard A. Posner, Regulation (Agencies) Versus Litigation (Courts): An Analytical
Framework, in REGULATION VS. LITIGATION: PERSPECTIVES FROM ECONOMICS AND LAW 11, 13
(Daniel P. Kessler ed., 2011); see also Steven Shavell, A Fundamental Enforcement Cost Advantage
of the Negligence Rule over Regulation, 42 J. LEGAL STUD. 275, 275–76 (2013) (“Under regulation,
compliance with standards tends to be assessed before, or independently of, the occurrence of
harm . . . . Under the negligence rule, in contrast, compliance with standards is examined only on
the condition that harm transpires . . . .”). Administrative law scholars do not generally focus on
adjudicative processes. Michael Asimow, Five Models of Administrative Adjudication, 63 AM. J.
COMP. L. 3, 5 (2015) (“[A]djudication is not the glamor area of contemporary administrative
law . . . . Adjudication is administrative law at the retail rather than the wholesale level.”). When
they do, they mostly focus on individuals protecting themselves from the wrongdoings of
government agencies, legitimacy, judicial independence from agency heads, separation of powers,
and congressional ability to implement policies. While Adminization touches on these issues, its
focus is on the optimal design of institutions that promote due process, efficiency, and justice.
106. Engstrom, supra note 31, at 622 (“A systematic accounting of agency gatekeeping helps
us to see [the choice between private enforcement and regulation] not as either/or options, but
rather the outer poles of a rich continuum of institutional designs that tap agencies’ unique
position and capacity to engage with and rationalize private litigation efforts.”). Notably, Engstrom
is largely critical of “retail” (i.e., case-by-case) administrative processes. Additionally, he generally
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
144 VANDERBILT LAW REVIEW [Vol. 71:1:121
cooperation, such as administrative adjudication (workers’
compensation, social security, and asbestos claims tribunals107) and
specialized courts (drug, mental health, and domestic violence courts).
These examples are useful, especially in assuaging constitutional
concerns, but it should be noted that they do not fully capture the goal
of Adminization—to enhance civil litigation by augmenting it with
agency functions.108
B. Main Features of Adminization
Adminization involves three central features run by a central
agency: audits and fines, sampling, and third-party communications.
1. Audits and Fines
To overcome the participation gap in civil litigation, a core
feature of Adminization is agency-run audits and fines. The agency
takes claims and, by its own initiative, investigates the case, collects
evidence, interviews witnesses, gathers documents, and locates
relevant industry standards. An auditor reaches out to the parties, asks
them about the case, asks for evidence such as receipts and credit card
charges, and presents them with questions. The goal of the agency’s
investigations is to assess the validity and reasonableness of the claim,
and the process is akin to that of audits run by other agencies. One close
analogy is the Equal Employment Opportunity Commission (“EEOC”).
When employees file charges of discrimination in the workplace, the
EEOC is empowered to conduct investigations on behalf of the
employee.109 Like consumer credit litigation, these cases also involve
abstracts from participation problems and grounds most of the critique on the assumption that the
adversarial process itself is functional. Id. at 667, 685.
107. See JOSEPH W. LITTLE ET AL., WORKERS’ COMPENSATION: CASES AND MATERIALS 544–45
(7th ed. 2014) (describing some of the benefits of Adminization of workers’ claims); Lester
Brickman, The Asbestos Claims Management Act of 1991: A Proposal to the United States Congress,
13 CARDOZO L. REV. 1891, 1892 (1992) (arguing that the processing of asbestos claims should be
rendered by an administrative agency rather than the tort system). There are also calls now to
create administrative health courts. See Nora Freeman Engstrom, A Dose of Reality for Specialized
Courts: Lessons from the VICP, 163 U. PA. L. REV. 1631, 1633–35 (2015). In the context of tort law,
some have proposed a move to a no-fault system. See Stephen D. Sugarman, Doing Away with Tort
Law, 73 CALIF. L. REV. 555, 558–59 (1985). Finally, consumer arbitration may suggest yet another
solution, an issue addressed separately infra Section III.D.
108. See, e.g., Arthur L. Shipe, Private Litigation Before the Commodity Futures Trading
Commission, 33 ADMIN. L. REV. 153 (1981) (considering the constitutionality and desirability of
administrative adjudication of private rights in “complex cases” such as futures trading). On the
constitutional challenges, see infra Part IV.A.
109. 42 U.S.C. § 2000e-8(a) (2012); see also EEOC v. Shell Oil Co., 466 U.S. 54, 68–70 (1984)
(discussing the limits of the EEOC’s investigative powers).
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
2018] ADMINIZATION 145
private information. But the EEOC, through its broad investigatory
powers, including the subpoena power, is still able to acquire
considerable information. The EEOC handles close to one hundred
thousand charges every year.110 And while the EEOC audits cases on
behalf of plaintiffs and not defendants, it shares the objective of
increasing participation.111 Similarly, the IRS conducts about 1.2
million audits annually,112 the Department of Justice often takes over
private qui tam lawsuits under the False Claims Act using its own
investigatory powers,113 and the CFPB has extensive experience in
investigating consumer complaints. 114
At first blush, it may seem wasteful to have the agency collect
information that the parties naturally possess. On closer inspection,
however, such an approach is highly attractive. First and foremost, we
have already seen that consumers are not always able to use their
information effectively, nor are they always aware of what information
is most relevant to their case. An agent collecting information would be
able to direct the parties to the most pertinent evidence. Especially for
the weaker party, it is an entirely different experience to produce
evidence for trial and to answer leading questions from an experienced
investigator, who can ask the consumers questions such as “Do you
have a bank statement from November, 2005, so that we can see if you
indeed paid off your debt?” Secondly, the agency, being part of the
government, can have access to information that may not be available
to other parties, such as agency records—a treasure trove of
information on past behavior and industry practices. Moreover, through
its investigatory powers, the agency can access information that is in
the hands of third parties. Overall, putting the agency at the front of
the process, in charge of initiating actions and using its expertise to
110. See Press Release, Equal Emp’t Opportunity Comm’n, EEOC Releases Fiscal Year 2015
Enforcement and Litigation Data (Nov. 2, 2016), https://www.eeoc.gov/eeoc/newsroom/release/2-
11-16.cfm [https://perma.cc/9Y33-KMF8] (reporting 92,000 claims in 2015).
111. See Michael Selmi, The Value of the EEOC: Reexamining the Agency’s Role in Employment
Discrimination Law, 57 OHIO ST. L.J. 1, 3 (1996) (analyzing critically whether the EEOC mitigates
participation problems).
112. INTERNAL REVENUE SERV., DATA BOOK 2015, at 9 (2016), https://www.irs.gov/pub/irs-
soi/15databk.pdf [https://perma.cc/KV9F-VYLV] [hereinafter IRS DATA BOOK] (reporting about
147 million individual income tax returns and audits of 0.8 percent of those).
113. See Marc S. Raspant & David M. Laigaie, Current Practice and Procedure Under the
Whistleblower Provisions of the Federal False Claims Act, 71 TEMP. L. REV. 23, 38–40 (1998)
(describing the government’s role in qui tam actions under the False Claims Act).
114. The CFPB recently proposed a program under which it would examine the practices of
covered entities, which comprise approximately sixty percent of the market. See CONSUMER FIN.
PROT. BUREAU, EXAMINATION PROCEDURES: DEBT COLLECTION 28,
https://s3.amazonaws.com/files.consumerfinance.gov/f/documents/201210_cfpb_debt-collection-
examination-procedures.pdf (last updated Oct. 24, 2012) [https://perma.cc/A726-VFB9].
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
146 VANDERBILT LAW REVIEW [Vol. 71:1:121
gather and analyze information, relieves critical pressure from the
consumer.
A complementary feature of audits is the use of fines against
baseless claims. Where a case is found to involve abuse or fraud, the
agency will issue a fine. The goal is not to conduct a “mini-trial,” but
rather to inspect the case for plausibility and signs of abuse or fraud—
the use of false evidence, the processing of unverified debts, or the
claiming of nonexistent charges, to give but a few examples.115 The size
of the fine may be influenced by various considerations, and economic
theory provides a guidepost: the magnitude of fines should reflect,
among other considerations, the probability of evading detection.116 The
agency should calibrate the level of fines according to the perceived
accuracy and frequency of its audits. Like audits, the use of fines is
commonplace among agencies, which use them as a means of
sanctioning noncompliant behavior. Fines give “teeth” to the audit
process and guarantee that fraudulent claims will be met with a
sanction even in cases of underparticipation by the defendant. The fines
are then paid to the public coffer and can be used for various social
purposes (including financing the agency, although this may raise
conflicts of interests).
Taken together, the use of audits and fines that are initiated by
the agency would provide a bulwark against abuse for those cases where
underparticipation is a problem. An outstanding issue is the costliness
of such audits, as it will clearly be prohibitively costly to audit all
incoming cases. We now move to consider another feature of
Adminization that accounts for this highly relevant concern.
2. Sampling, Artificial Intelligence, and Resource Management
Both the judicial process and audits are resource-intensive
processes. Marginalist economic theory teaches that, given budgetary
constraints, it is desirable to allocate resources such that they have the
115. On the plausibility standard, see FED. R. CIV. P. 12(b)(6); Ashcroft v. Iqbal, 556 U.S. 662,
678–89 (2009) (“Only a complaint that states a plausible claim for relief survives a motion to
dismiss.”); Bell Atl. Corp. v. Twombly, 550 U.S. 544 (2007) (setting a plausibility test for the filings
of civil actions). See also Raymond H. Brescia & Edward J. Ohanianm, The Politics of Procedure:
An Empirical Analysis of Motion Practice in Civil Rights Litigation Under the New Plausibility
Standard, 47 AKRON L. REV. 329, 334–51 (2014) (reviewing the evolution of the standard and its
critique). On Rule 11, see FED. R. CIV. P. 11.
116. See Becker, supra note 15 (developing the foundations for the theory of optimal fines in
law enforcement). There is rich literature that examines the constraints on the use of fines to
supplement imperfect enforcement. The key reasons developed there—risk aversion and wealth
constraints—apply only weakly in the context of consumer credit litigation. See generally A.
Mitchell Polinsky & Steven Shavell, The Theory of Public Enforcement of Law, in 1 HANDBOOK OF
LAW AND ECONOMICS 403, 405 (A. Mitchell Polinsky & Steven Shavell eds., 2007).
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
2018] ADMINIZATION 147
greatest marginal productivity. This would often imply that different
cases should receive varying amounts of attention.117 However, civil
litigation handles attention allocation relatively poorly. Judges are not
free to dismiss cases simply because they want to devote more time to
hear other cases which are more deserving of judicial attention.118 A
judge is expected to give some attention to all the cases that come before
her, and lack of public interest is not a general reason for refusing to
hear cases. In contrast, agencies frequently allocate and prioritize
attention and resources based on priorities, with a clear example being
the IRS, which chooses only about one percent of all cases for in-depth
review.
Sampling is the process by which agencies choose the cases they
would like to prioritize and examine. There are a few approaches to
sampling, and the most straightforward and well-known one is random
sampling. This is the approach used, partially, by the IRS and the
TSA.119 A random sampling implies that each case has an equal chance
of being chosen for audit, thus imposing an equal risk of examination
on all participants. This approach has many upsides, with simplicity
being a main one. This approach also has a very clear drawback, in that
meritorious cases have an equal chance of being chosen for audit, thus
wasting resources. Another approach is to choose cases based on criteria
that are suggestive of risk. For example, the police may monitor known
sex offenders more closely than other citizens, and an insurance
company may only investigate claims of high value. This has the
drawback that if the criteria used to sample cases are known in
advance, then the system may be gamed.120 Moreover, prescreening the
cases that would be sampled can itself be resource-intensive, thus
reducing the benefit of using samples.
117. There are many advantages to the focusing of attention and there are even potential
economic gains from focusing enforcement efforts on arbitrary subgroups, like auditing more
closely the tax returns of people whose last name begins with A than those whose last name begins
with B. See Henrik Lando & Steven Shavell, The Advantage of Focusing Law Enforcement Effort,
24 INT’L REV. L. & ECON. 209, 209–10 (2004).
118. Adam M. Samaha, Randomization in Adjudication, 51 WM. & MARY L. REV. 1, 70–81
(2009) (exploring the role of randomization in adjudication and defending the use of case
randomization). The literature considers, to some extent, the role of managerial judges in
managing resources. See, e.g., Judith Resnik, Managerial Judges, 96 HARV. L. REV. 374, 380
(1982).
119. The samples are not purely random, and profiling (racial and other) is common. See Robin
Shepard Engel & Jennifer M. Calnon, Examining the Influence of Drivers’ Characteristics During
Traffic Stops with Police: Results from a National Survey, 21 JUST. Q. 49, 69–77 (2006) (finding
strong evidence of racial profiling in traffic police stops).
120. See Lando & Shavell, supra note 117, at 215 (arguing that a known enforcement focus
may increase crime if offenders can freely move to offend in unenforced areas).
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
148 VANDERBILT LAW REVIEW [Vol. 71:1:121
A promising sampling approach that can be used fruitfully in
Adminization is smart sampling—the use of Big Data and artificial
intelligence (“AI”) to profile risky cases using complex models. To be
clear at the outset, although I believe smart sampling to be highly
feasible and relatively inexpensive to develop, nothing in Adminization
depends directly on such sophisticated methods, and the system could
work on the basis of random sampling until smart sampling algorithms
prove workable. With this caveat in mind, smart sampling consists of
using machine learning algorithms to identify cases that are
statistically most likely to involve fraud based on the past resolution of
similar cases. Poring over the vast history of past cases, AI software can
identify those characteristics of a case that are most likely to correlate
with its eventual dismissal. Each of these characteristics is assigned a
risk weight. Based on a complex risk model, the software can decide the
probability with which a given case will be sampled. Smart sampling
can be done with great speed, at almost zero marginal cost, and
potentially with great accuracy. Unlike traditional criteria-based
sampling, smart sampling is not open to gaming by market
participants. The complexity of AI algorithms—which, ironically, is a
frequent criticism levied against them—presents a black box to those
who would seek to game the system.121 It is not surprising that the
private market is replete with AI-assisted fraud detection
algorithms.122 In the same spirit, agencies are starting to realize the
potential for machine learning for complaint handling. Today, the SEC
is developing an automated system that flags cases for review. The
system is based on an automated anomaly detection model that would
flag submissions for human review on the basis of statistical deviations
from the common filings.123
It may seem ambitious to develop a fraud-detecting software,
given the great diversity of cases and the complexity involved. And
while there is nothing simple about this task, it should be evaluated in
light of AI’s proven capabilities, especially bearing in mind the recent
121. See VIKTOR MAYER-SCHÖNBERGER & KENNETH CUKIER, BIG DATA: A REVOLUTION THAT
WILL TRANSFORM HOW WE LIVE, WORK, AND THINK 178 (2014) (“The basis of an algorithm’s
predictions may often be far too intricate for most people to understand.”); see also David Sussillo
& Omri Barak, Opening the Black Box: Low-Dimensional Dynamics in High-Dimensional
Recurrent Neural Networks, 25 NEURAL COMPUTATION 626, 627–29 (2013) (noting how a recurrent
neural network is viewed as a black box in terms of its implementation of its target functions).
122. See, e.g., Clifton Phua et al., A Comprehensive Survey of Data Mining-based Fraud
Detection Research, https://arxiv.org/pdf/1009.6119.pdf (last visited Dec. 19, 2017)
[https://perma.cc/E28Z-BWK2].
123. The model is called the “Automatic Quality Model” and is based, at least in part, on a
Jones Model: measuring the difference between a company’s discretionary accruals and those of
peer companies in the industry. See Douglas M. Boyle et al., Insights into the SEC’s Accounting
Quality Model, CPA J., May 2015, at 16, 18.
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
2018] ADMINIZATION 149
victory of AI over grandmaster Lee Sedol at the game of Go—a game so
rich in possibilities that it was considered to be impossibly stacked
against machines and in favor of human intuition.124 The closest
example of a working AI technology in fraud detection comes from the
credit card industry.125 Despite a daily volume of millions of
transactions,126 credit card companies effectively flag fraudulent
transactions, alerting human investigators of potential fraud.127 These
algorithms run in real time and evaluate each transaction against a
model of the specific consumer, placing alerts in the case of any
significant deviation from model-predicted behavior. Sifting through
the large dataset of past purchases, the consumer model is able to detect
when purchases are made in unexpected locations, times, or amounts.
Importantly, these algorithms, which run on an almost
incomprehensible volume of data with little to no human intervention,
manage to detect suspicious transactions with a relatively low level of
either false negatives or false positives. Another telling example is that
of spam filters. Until very recently, it seemed nearly impossible for a
computer to overcome the problem of spam identification, as the range
of richness of human communication is so vast. In 2002, for example,
Slate ran an article that pessimistically stated, “It’s time to give up . . .
spam has won. Spam is killing e-mail.”128 Pew predicted in 2002—based
on a large consensus—a rate of spam growth that would imply today
hundreds if not thousands of spam messages every day.129 Yet email
124. See Adrian Cho, “Huge leap forward”: Computer That Mimics Human Brain Beats
Professional at Game of Go, SCI. (Jan. 27, 2016, 1:00 PM), http://www.sciencemag.org/news/
2016/01/huge-leap-forward-computer-mimics-human-brain-beats-professional-game-go
[https://perma.cc/N3M6-K2EU] (“[F]or many years people have tried to sell the notion of Go as a
game in which computers can never beat humans.”); Cade Metz, Google’s AI Takes Historic Match
Against Go Champ with Third Straight Win, WIRED (Mar. 12, 2016, 3:21 AM),
https://www.wired.com/2016/03/third-straight-win-googles-ai-claims-victory-historic-match-go-
champ/ [https://perma.cc/9EQQ-KXKU].
125. See generally Richard J. Bolton & David J. Hand, Statistical Fraud Detection: A Review,
17 STAT. SCI. 235 (2002).
126. See VISA, ANNUAL REPORT 2013 (2013), https://s1.q4cdn.com/050606653/files/
doc_downloads/annual%20meeting/Visa%20Annual%20Report%202013%20final%20website.pdf
[https://perma.cc/SC7Q-3BXY] (reporting a daily volume of 160 million transactions).
127. Some of the methods include genetic algorithms, Bayesian classifiers, a hidden Markov
model, and, more recently, neural networks. See generally Krishna Kumar Tripathi & Mahesh A.
Pavaskar, Survey on Credit Card Fraud Detection Methods, 2 INT’L J. EMERGING TECH. &
ADVANCED ENGINEERING 721 (2012).
128. Kevin Werbach, Death by Spam, SLATE (Nov. 18, 2002, 10:35 AM),
http://www.slate.com/articles/technology/webhead/2002/11/death_by_spam.html
[https://perma.cc/D38A-54LZ].
129. Deborah Fallows, Email at Work: Few Feel Overwhelmed and Most Are Pleased with the
Way Email Helps Them Do Their Jobs, PEW RES. CTR. 5 (Dec. 8, 2002),
http://www.pewinternet.org/2002/12/08/email-at-work/ [https://perma.cc/K62R-HTGP] (citing
sources predicting a doubling of spam load every six months and a rate of growth from 2001 to
2006 of approximately six hundred percent).
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
150 VANDERBILT LAW REVIEW [Vol. 71:1:121
survived. Google reports that its email service, Gmail, filters ninety-
nine percent of all spam while only having a one percent rate of false
positives.130 Stated differently, Google reports that less than 0.1 percent
of email in the average inbox is spam while less than 0.05 percent of
wanted messages are in the spam folder.131
Another example illustrates the power of statistical fraud-
detection algorithms. Benford’s law is a decision rule that meets a
seemingly impossible challenge: How can one detect fraud in accounting
books without actually analyzing them? The astronomer Simon
Newcomb postulated in 1881—and later the physicist Frank Benford
proved—that one could identify potential fraud by simply looking at the
numbers reported in these ledgers, and more specifically, at the digits
themselves.132 If we count the frequency with which each digit appears
in financial accounts, a pattern emerges with surprising regularity. In
thirty percent of cases, the first digit of any number is one, but there is
only a 4.5 percent chance of it being a nine. For a variety of reasons,
naturally occurring numbers have greater likelihood of starting with
certain digits than others. Knowing this rule, we can count all the digits
that appear in a given account book. If much more than 4.5 percent of
the numbers start with nine, or much less than thirty percent of the
numbers start with one, then we have good reason to suspect that the
book was tampered with.133 Cooking the books will often leave a
footprint in the form of unnatural distribution of digits, and simply
counting the frequency of digits—without any real understanding of the
business—will indicate cases with suspected wrongdoing. Rules like
Benford’s law were developed by humans. Software would probably use
much more nuanced and sophisticated rules, taking account of every
facet of the case—from the identity of the parties through the amounts
indicated, and perhaps even seemingly irrelevant features like the font
used or the time of filing. However, the core ideas remain the same.
130. Cade Metz, Google Says Its AI Catches 99.9 Percent of Gmail Spam, WIRED (July 9, 2015,
2:00 PM), http://www.wired.com/2015/07/google-says-ai-catches-99-9-percent-gmail-spam/
[https://perma.cc/QD4U-3DU5].
131. Emil Protalinski, Google Now Uses an Artificial Neural Network to Fight Spam, Debuts
Gmail Postmaster Tools to Cut False Positives, VENTUREBEAT (July 9, 2015, 11:35 AM),
https://venturebeat.com/2015/07/09/google-launches-gmail-postmaster-tools-to-help-companies-
ensure-their-emails-arent-marked-as-spam/ [https://perma.cc/LUJ3-7UYE].
132. See Simon Newcomb, Note on the Frequency and Use of the Different Digits in Natural
Numbers, 4 AM. J. MATHEMATICS 39 (1881); Frank Benford, The Law of Anomalous Numbers, 78
PROC. AM. PHILOS. SOC’Y 551 (1938).
133. Cindy Durtschi et al., The Effective Use of Benford’s Law to Assist in Detecting Fraud in
Accounting Data, 5 J. FORENSIC ACCT. 17, 18–19 (2004).
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
2018] ADMINIZATION 151
To develop such sophisticated rules, we would need a large body
of training data.134 Ideally, the data will be “labeled,” i.e., each case will
be identified as either being with merit or without merit. Without such
data, machine learning cannot produce accurate predictions. Luckily,
this type of “big data” is readily available.135 As Andrew Crespo recently
noted, a by-product of the judicial process is a large body of unutilized
“systemic facts,” which are records of cases, claims, and resolutions.136
These present an almost perfect type of training data—the software can
scan the filings and all relevant facts of the case and then see how it
was decided. Of course, some of the data will have to be filtered, as many
cases are decided not on the merits. Yet, there is such a wealth of data
on all the millions of claims that are filed every year that even after
filtering, there will be a very large body of data. Moreover,
Adminization constantly produces new data. As part of the process,
cases are chosen for audit and are then subject to review—an
information producing process. Importantly, not only flagged cases will
be chosen, but also a few nonflagged cases. The results of the audit will
then be fed into the machine learning algorithm. If a flagged case is
proved to involve fraud, this will reinforce the rules used by the
software. If there was no fraud in a flagged case, this will prompt the
software to modify its decision rules—and the converse applies to
nonflagged cases. Over time, the system will self-modify based on the
results of the audit process, thus promising continuous improvement
and adaptation to changing circumstances.
3. Third-Party Communications
As previously discussed, the expectation that plaintiffs, who
stand to gain from consumer underparticipation, will effectively serve
134. A general view among computer scientists is that having a large dataset is at least as
important as good machine learning models to the development of effective algorithms. Google’s
Research Director, Peter Norvig, famously stated on Google’s success in this area: “We don’t have
better algorithms. We just have more data.” See Xavier Amatriain, Mining Large Streams of User
Data for Personalized Recommendations, ACM SIGKDD EXPLORATIONS, Dec. 2012, at 37, 43
(quoting Norvig in discussion on power of data); see also Pedro Domingos, A Few Useful Things to
Know About Machine Learning, 55 COMM. ACM, Oct. 2012, at 78 (explaining that simple
algorithms with large amounts of data are superior to sophisticated algorithms with modest
amounts of data).
135. On big data and the law, see Daniel Martin Katz, Quantitative Legal Prediction—Or—
How I Learned to Stop Worrying and Start Preparing for the Data-Driven Future of the Legal
Services Industry, 62 EMORY L.J. 909, 913–22 (2013).
136. Andrew Manuel Crespo, Systemic Facts: Toward Institutional Awareness in Criminal
Courts, 129 HARV. L. REV. 2049, 2065–66 (2016). Crespo’s argument is couched in the context of
the criminal law system; however, the spirit of his argument applies to civil litigation.
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
152 VANDERBILT LAW REVIEW [Vol. 71:1:121
consumers with court documents is highly unrealistic and will result in
many instances of sewer service.137
Under Adminization, the agency serves process, as well as all
other communications, thereby informing defendants of their rights.
This simple design feature will directly solve this structural problem,
which is wholly an artifact of a design that is incompatible with private
incentives. Indeed, pilot programs with third-party service by the court
were successful, which suggests even greater potential effectiveness if
done at scale by an agency.138 And while it may be possible to adapt
courts to provide services, agencies are naturally better designed to
provide such “outgoing” services, which involve reaching out to
individuals, locating them, and handling the necessary administrative
aspects. It will also allow courts to develop a more independent
approach to evaluating the quality of service if they are not implicated
in the process. In terms of finance, the service may still be funded as it
is today—by the plaintiff through fees. Moreover, taking advantage of
its disinterested role, the agency can also provide defendants with
educational materials to inform them of their rights, a function agencies
rarely perform today.139 With its communications, the agency could
provide informative, plain-language explanations of defendants’ rights
and duties, using simple illustrations, flowcharts, frequently asked
questions, and visual guides. In contrast, entrusting plaintiffs with this
task would again engender a moral hazard problem.
C. Adminization of Consumer Credit Litigation
The application of Adminization to consumer credit litigation
starts with the agency. The administrative overlay in the context of
consumer credit can be the CFPB, with its broad regulatory powers
under the Dodd-Frank Act.140 Indeed, it is possible to implement
Adminization using state or even local agencies; nothing here depends
critically on the use of federal agencies. Yet the advantages of scale, as
well as the broad existing powers of the CFPB, are very appealing, and
137. See supra Section I.A.
138. N.Y. COMP. CODES R. & REGS. tit. 22, § 208.6(h) (2017); OUT OF SERVICE, supra note 68,
at 11–12 (showing that sending court summons in addition to plaintiff summons resulted in an
increase in consumer participation, and that consumers often reported receiving only the court’s
summons).
139. See Thomas v. Law Firm of Simpson & Cybak, 354 F.3d 696, 699 (7th Cir. 2004) (“Nothing
in the FDCPA suggests that Congress intended creditors’ unilateral actions to obligate debt
collectors to inform debtors of their rights . . . .”), vacated, 358 F.3d 446 (7th Cir. 2004), and opinion
substituted, 392 F.3d 914 (7th Cir. 2004). Consumer education is at least partially a problem with
lack of incentive to learn. Since Adminization makes it easier to contest claims, learning
information becomes more attractive.
140. See infra Section IV.A.
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
2018] ADMINIZATION 153
so I will focus on this agency. The CFPB’s powers include the power to
investigate claims related to debt collection, the power to summon
witnesses, and the ability to issue fines.141 The existence of the CFPB’s
platform, its broad legislative powers, and its subject-matter expertise,
promise a smooth implementation at a relatively low marginal cost.
The process starts by filing a claim with the agency. A claim
could be initiated by the original creditor, or if the state permits, a debt
buyer. The claimant would be required to furnish rudimentary
information regarding the claim: the identity of the debtor and her last
known address, an estimate of the breakdown of the debt to its principal
and other fees, the origin of the claim, and the name of the original
creditor. The standard by which the quality of information is judged is
whether it provides a sufficient basis for a reasonable but
unsophisticated consumer to decide if the debt is real and accurate.142
The claimant would acknowledge, on pain of financial sanctions, that it
holds supporting evidence, although the current rules requiring an
affidavit may be relaxed.143
The agency will check the claim via an automatic machine
learning system that would screen and flag cases. The algorithms will
check, for example, whether the debt is time barred, whether the
interest rate exceeds statutorily allowed levels, and whether another
identical claim against the same debtor was filed by a different creditor.
If violations of bright line rules are identified, the claim will be
automatically rejected without prejudice and a notice will be sent to the
creditor, explaining the flaw. This will be beneficial to consumers in
that it will filter out empty claims that are currently filed against them;
specifically, this will solve the problem of “zombie debts,” which are
141. See infra Section IV.A.
142. A similar requirement exists under German law. See Sigmund A. Cohn, A Streamlined
Debt Collection Procedure in the Federal Republic of Germany, 2 B.C. INT’L & COMP. L. REV. 69,
71–79 (1978); European Consumer Ctr. Ger., The German Judicial System, EUR. CONSUMER CTR.
NETWORK 12–15 (Dec. 2010), https://www.evz.de/fileadmin/user_upload/eu-
verbraucher/PDF_Englisch/Brochures/Legal_sytem_Germany.pdf [https://perma.cc/PK3R-XFQ6];
Grozdana Šijanski & Jimmy Barber, The German Order for Payment Procedure (Mahnverfahren),
GERMAN L. ARCHIVE (2006), http://germanlawarchive.iuscomp.org/?p=343 [https://perma.cc/U7SA-
J5A9].
143. The insistence on signed affidavits in the legal system resulted in a large industry of robo-
signing. See generally Holland, supra note 60. Courts treat robo-signing with disdain. See, e.g.,
Intervale Ave Assoc v. Donlad, No. L & T 60527/12, 2013 WL 540153, at *4 (N.Y. Civ. Ct. Feb. 7,
2013) (“The courts have consistently demonstrated an intolerance for ‘robo-signing.’ ”). But the
problem of robo-signing is artificial because what should matter is the existence of evidence, not
the form of signature, and requiring personal knowledge for hundreds of thousands of debt claims
is grossly inefficient. If creditors can reliably present claims (at the pain of large financial
sanctions), this could achieve the same goals but at a lower cost.
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
154 VANDERBILT LAW REVIEW [Vol. 71:1:121
time barred actions that attempt to exploit consumer ignorance and
judicial passiveness.144
Besides the automatic screening of cases, the system will also
employ “smart sampling” to identify the cases which are most likely to
involve fraud. The exact algorithms will depend on implementation,
but, as a general matter, the learning system will synthesize statistical
information regarding the identity of the original creditor, the identity
and demographics of the debtor, the sums involved, the type of debt
involved, time of filings, and other case characteristics. If certain
creditors are known to engage in wrongdoing, this will increase the
likelihood that the case will be chosen for audit. If certain demographics
are targeted more frequently for abusive lawsuits—e.g., the elderly,
minorities, or the uneducated—then their cases will be flagged for audit
more frequently than other cases.
Flagged cases will be transferred to the agency’s auditors, who
will use their investigative powers to demand proof of the evidence
claimed by the creditor. The investigators will check if the evidence is
consistent, whether the case presents a cause, and, most importantly,
if there are any indications of fraud or abuse. In some cases, there will
be a need to acquire information from consumers. In these cases, the
investigators will approach consumers and ask for information. The
consumer will not be under any obligation to cooperate, but it should be
explained that an investigation can only advance the consumer’s case.
A friendly conversation could greatly advance the consumer’s interests,
as the auditor could lead with simple questions that would avoid the
need to present a legal case—“Do you have a receipt?”; “Do you have a
document showing that you were elsewhere on the date the alleged
purchase was made?”; “Did you file a complaint against identity theft?”;
etc.
If the audit reveals wrongdoing, the plaintiff will be issued a
fine. The findings of the investigation will be evaluated by the
professional staff at the agency, and where they find indications of
fraud, abuse, or other illegal practices, they can use their legal powers
to levy fines.145 The magnitude of the fine should reflect both the
severity of the offense and the likelihood of evasion. In general, large
fines would be required to deter companies from bringing abusive
lawsuits, since only a sample of cases are audited. As an administrative
action, such fines will be subject to appeal. This fine will be paid to
144. See Young Walgenkim, Killing “Zombie Debt” Through Clarity and Consistency in the
Fair Debt Collection Practices Act, 24 LOY. CONSUMER L. REV. 65, 65 (2011) (discussing the problem
of debt collectors attempting to “revive stale, paid-off, otherwise uncollectable debt”).
145. 12 U.S.C. § 5565 (2012).
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
2018] ADMINIZATION 155
either the agency or the government by the creditor, and these funds
may be used to finance the agency, although it will be prudent to avoid
potential conflicts of interest by not creating a direct link between fines
and agency funds.
The use of audits and fines will provide consumers with a basic
level of consumer protection. It will do so not by increasing participation
but rather by eroding the harsh consequences of underparticipation. By
using audits and fines, there will be an effective sanction against the
filing of fraudulent or unsupported claims, thus making participation
less critical and saving considerable resources. The use of audits and
fines also conforms to the prevalent but misguided expectation among
consumers today that by filing an answer the court will handle the issue
sua sponte.146
All the cases will then proceed to a “Communication” stage.
Unlike the current system, it is not the plaintiff but the agency that
would be responsible for contacting the consumer. The agency will use
its own databases, as well as information provided by the creditor, to
locate the consumer and communicate with them by email, mail, or
phone. This will address the root cause of the “sewer service”
problem.147 Here and throughout, the quality of communications should
be emphasized. Freed from the chains of legal language and procedure,
the communications should be made simple, friendly, easy to follow, and
graphic.148
All consumers will be sent a simple form. It will clearly inform
them of the fact of a claim made against them and its potential
implications. It should ask the consumer if she recalls making the
purchase from the original creditor and whether the principal and
charges seem correct.149 On this basis, the form will provide three
options: admitting the claim, contesting it, or ignoring it.150 Admitting
146. Due Process and Consumer Debt, supra note 42, at 18 (“Many defendants believe that
once they answer, the court will review their allegations and defenses sua sponte.”).
147. See supra Section I.A.
148. See D. James Greiner & Andrea J. Matthews, The Problem of Default, Part I (June 24,
2015) (unpublished), http://papers.ssrn.com/sol3/papers.cfm?abstract_id=2622140
[https://perma.cc/8BW2-VAET] (studying interventions that improve consumer participation).
149. This addresses a common problem today of the so-called “alphabet soup” of creditors,
where debtors receive debt claims from organizations with a name like ABC, which bears little
resemblance to the consumer’s experience of the origination of the debt (e.g., Best Buy). See
Roundtable on Data Integrity in Debt Collection: Life of a Debt, FED. TRADE COMMISSION &
CONSUMER FIN. PROTECTION BUREAU (2013), https://www.ftc.gov/system/files/documents/
public_events/71120/life-debt-roundtable-transcript.pdf [https://perma.cc/9XM5-L6AY]. For a
similar (although more onerous) recommendation, see FTC PROTECTING CONSUMERS REPORT,
supra note 14, at 16–17.
150. How to most clearly encourage consumer response is a question best left to
communications experts, who are frequently and regrettably missing from the design of most
governmental communications.
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
156 VANDERBILT LAW REVIEW [Vol. 71:1:121
will invite the consumer to make payments and, perhaps, financial
incentives (such as interest reduction) may be offered to fast-paying
consumers. If the consumer pays off the debt, the agency will provide a
confirmation letter that immunizes the consumer from any future
action based on this debt. The agency will then process the payment
and transfer it to the creditor. Alternatively, the consumer could offer a
settlement by proposing an affordable installment plan, which the
creditor may accept or reject.151 Many creditors should be willing to
accept reasonable payment plans, which offer greater recovery than
enforcement.
If the consumer contests the claim, the form will contain a few
sample checkboxes, which can be used later in litigation instead of a
more formal consumer response. Five checkboxes should be provided: “I
do not recognize the person to whom the debt is owed,” “I already paid
off this debt,” “the amount is wrong,” “another person owes this debt,”
and “other.” An open comment field should be available where the
putative debtor could write why the debt is wrong. Listing supporting
evidence should also be made easy but not mandatory. Contested cases
will be transferred to litigation, and only for those cases will the creditor
be required to provide a full body of evidence. The chief benefit of only
asking for evidence in contested cases is that it saves creditors the
immense costs of providing full evidence in all cases. This feature will
greatly increase the political appeal of this system to creditors.
Ignoring the claim will trigger a reassessment of the consumer’s
address: the agency should invest reasonable effort into searching for
the debtor using both its own resources and information procured from
the creditor. If the agency concludes that reasonable effort has been
taken, the communication should be deemed ignored and moved to
litigation, alongside all other contested cases.
Contested and ignored cases will be litigated, and the outcomes
of the process will be “fed” to the machine learning algorithms for future
improvements. These outcomes include the agency’s findings,
consumer’s response, and the court’s ruling. On this basis, the agency
will also be able to manage an internal score of creditor reputation, with
every finding of fraud lowering the creditor’s score. Low score creditors
will be chosen for audit more often—as the algorithm will take account
of their identity—whereas high score creditors will be subject to fewer
investigations.152 Creditor reputation could also be made public, thus
151. The consumer’s choice to admit the debt has important legal ramifications, and these
ramifications should be clearly explained.
152. For obvious reasons, the odds of being selected for audit, even for a creditor with the
highest level of reputation, must never be zero.
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
2018] ADMINIZATION 157
informing future consumers before they engage with a specific provider
of credit. This reputation system will provide greater compliance
incentives, especially since most debt collection lawsuits are brought by
a limited number of creditors.153
Adminization does not supplant litigation; rather, it
complements it. The continued use of litigation may raise some concern,
given its imperfections discussed above, yet the process will carry
significant advantages over the current system. First, and perhaps
most importantly, Adminization will significantly curtail the filing of
unmeritorious claims. Once plaintiffs internalize the risk associated
with filing frivolous claims—due to the real potential for fines—they
will be less inclined to file them. Second, and as a result, Adminization
provides important cost savings for the judicial system. The reduced
volume of filing (due to lesser incentive to file unmeritorious claims),
will lead to fewer cases on the docket. This will save considerable
resources for the courts, freeing them up to scrutinize other debt cases
more closely, thus further deterring the filing of unmeritorious
lawsuits. Third, Adminization is also highly beneficial for creditors. By
increasing the reliability and legitimacy of consumer credit contracts,
and by simplifying the process of producing judgments for uncontested
cases, there will be significant savings in the cost of providing credit—
savings that would be expected to be partially passed on to consumers.
From the consumer side, this will make the use of credit a safer option,
thus increasing the utilization of safe credit. This has important
implications, especially for people in poverty, for whom access to credit
is a persistent obstacle.154 No doubt, Adminization also involves certain
costs, but as I endeavor to show below, the costs are unlikely to be
prohibitively high and will mostly be offset by a reduction in the volume
of litigation. Perhaps more importantly, these costs pale in comparison
to any of the other alternatives currently considered, a topic to which I
now turn.155
III. THE FAILURE OF PARTICIPATION-BASED SOLUTIONS
In evaluating the desirability and effectiveness of Adminization,
it is important to be cognizant of the alternatives. The various
153. See DEBT WEIGHT, supra note 42, at 14 (finding that over fifty-eight percent of the cases
in the sample were brought by three debt buyers).
154. See, e.g., Dean Karlan & Jonathan Morduch, Access to Finance, in 5 HANDBOOK OF
DEVELOPMENT ECONOMICS 4702, 4703 (Dani Rodrik & Mark Rosenzweig eds., 2009) (“Expanding
access to financial services holds the promise to help reduce poverty and spur economic
development.”).
155. See infra Section IV.D.
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
158 VANDERBILT LAW REVIEW [Vol. 71:1:121
alternative solutions currently considered can be effectively grouped
under the general heading of participation-based solutions. The
common core idea, explored below, is that incentivizing and subsidizing
consumer participation would allow judges to have the information they
need to scrutinize cases.156 A thorough analysis of these proposed
solutions reveals, as I will show in this Part, that participation-based
solutions involve immense costs but marginal benefits, and that the
costs of Adminization pale in comparison to the costs and risks of
participation-based approaches.157 With this in mind, Adminization and
participation are not mutually exclusive, and a well-functioning system
should employ some degree of both approaches. My main contention is
not that participation solutions are without merit in some absolute
sense, but rather that—on the margin—there is much greater need for,
and a much higher return on, investments in administrative review as
a screening mechanism than greater and greater investments in more
traditional court-based solutions.
A. Lawyering Up
The most prominent call to solve the problem of abuse in civil
litigation has been to expand legal access through public subsidies of
legal services. Under this view, if consumers received subsidized access
to legal representation, they would more often stand up against wrongs,
assert their rights in court, and contest fraudulent claims.158 On this
view, the resulting rise in consumer participation will provide judges
with the information they need to screen out bad cases and prevent
156. On the dominant role of participation-based approaches in state legislatures, see, for
example, N.Y. ACCESS TO JUSTICE REPORT, supra note 81, at 3 (requesting $30 million in public
funding for legal assistance to “close the justice gap”); STATE BAR OF CAL., CIVIL JUSTICE
STRATEGIES TASK FORCE, REPORT & RECOMMENDATIONS 19 (2015),
http://board.calbar.ca.gov/docs/agendaItem/Public/agendaitem1000013003.pdf [https://perma.cc/
9W3Q-XF66] (recommending “that the State Bar support efforts to secure universal
representation”); Mission & Goals, TEX. ACCESS TO JUST. COMMISSION,
http://www.texasatj.org/mission-goals (last visited Oct. 23, 2017) [https://perma.cc/QQ34-WPFV]
(reporting their central goal of “reduc[ing] barriers to our judicial system”).
157. Proposals that primarily affect the debt collection industry, such as licensing
requirements, are excluded. Perhaps this type of ex ante regulation of debt collection is helpful,
but the New York experience—where licensing is employed—casts doubt. See N.Y.C., N.Y.,
ADMINISTRATIVE CODE § 20-490 (2017).
158. There are many reasons why consumers underparticipate in legal proceedings, leading to
potentially significant divergence between the social interest in the existence of lawsuits for
wrongful behavior and private incentives not to sue. See, e.g., Yonathan A. Arbel & Yotam Kaplan,
Tort Reform Through the Back Door: A Critique of Law and Apologies, 90 S. CAL. L. REV. 1199
(2017) (showing evidence that the simple tender of apology can cause consumers to avoid filing
lawsuits for meritorious claims of malpractice).
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
2018] ADMINIZATION 159
plaintiffs from taking advantage of consumers.159 This type of proposal,
often called a “civil Gideon” right, mirrors the right of indigent
defendants in criminal proceedings to an attorney.160 While this is
proposed as a primary solution to the problem, it is unworkable,
prohibitively costly, and of marginal effectiveness.
First, the sheer number of people who would be eligible for this
subsidy is staggering. The former president of the American Bar
Association (“ABA”) claimed that “one in five Americans now qualifies
for legal assistance,”161 and even that, he thought, was an
understatement: “[I]t’s not just the poor [who need assistance] . . . Too
many low- and moderate-income people cannot access legal
representation.”162 Yet, even his more conservative estimate implies
that sixty-four million people nationwide will be eligible for this
subsidy. And while not all of these people have legal issues, a significant
majority do, and those that do often have more than one. A recent study
found that about half of all low-income New Yorkers have experienced
legal issues in the course of a year, with about a third of them facing
three or more legal issues.163 Based on these estimates, which are
admittedly rough, we would expect there to be about thirty-two million
people who are both eligible for a subsidy and have a legal issue, ten
million of whom would have three or more such issues.
The cost of providing subsidies on such a scale is immense. The
ABA, which may have a reason to downplay the costs of legal aid,164
estimates the costs of expanding legal access at about $1.7 billion every
year.165 This is unlikely, as this amount is not much larger than the
159. See, e.g., DEBT WEIGHT, supra note 42, at 21 (calling on the state of New York to “[f]und
legal services for low-income and working poor individuals sued on alleged debts” and “[f]und the
provision of assistance, information and resources for pro se defendants”); Fox, supra note 46, at
75 (“Consumers need to be provided the legal assistance necessary to defend themselves in civil
debt litigation.”).
160. Gideon v. Wainwright, 372 U.S. 335, 339 (1963).
161. PERMANENT COMM’N ON ACCESS TO JUSTICE, REPORT TO THE CHIEF JUDGE OF THE STATE
OF NEW YORK—APPENDICES app. 7 at 41 (2015), https://www.nycourts.gov/
accesstojusticecommission/PDF/2015_Access_to_Justice-Appendices.pdf [https://perma.cc/T2GG-
ER4D].
162. Id.
163. N.Y. ACCESS TO JUSTICE REPORT, supra note 81, at 9.
164. Lawyers’ incentives exert considerable pressure on the choice of legal procedure. See, e.g.,
Yonathan A. Arbel, Contract Remedies in Action: Specific Performance, 118 W. VA. L. REV. 370,
388–89 (2015) (finding that lawyers tend to steer clients to opt for remedies that would facilitate
the collection of attorney’s fees).
165. The ABA finds an even lower amount—$1.7 billion, but this is based on the very strong
assumption that $100 worth of legal services will suffice for the common consumer. This rate is
the equivalent of less than an hour of work per case, which seems highly ambitious. The ABA also
notes that the United Kingdom spends $1.36 billion on legal services for the poor, which would
imply a U.S. cost of $8.16 billion (population weighted). TASK FORCE ON ACCESS TO CIVIL JUSTICE
ET AL., AM. BAR ASS’N, REPORT TO THE HOUSE OF DELEGATES 14 (2006),
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
160 VANDERBILT LAW REVIEW [Vol. 71:1:121
current cost of legal aid, estimated at $1.3 billion annually.166 More
realistically, Jessica Steinberg estimates that the costs would be three
times the ABA’s estimate: around $5.4 billion every year.167 My analysis
suggests that if we discard projections and instead look at the actual
costs of running the institutions that are currently assisting those in
need, we will find costs that are higher by at least an order of
magnitude. In New York, the Interest on Lawyer Account (“IOLA”) fund
reports that in 2013, a large group of supported organizations closed
296,621 cases with an overall budget totaling $266.6 million.168 This
implies a per-case cost of $897, which is the equivalent of 7.7 hours of
paralegal work per case at the national average rate of $116, or 2.5
attorney hours at the average rate of $361.169 Now, if indeed around
thirty-two million Americans would be eligible for assistance,170 then
the annual cost would amount to $28.7 billion—about seventeen times
more than the already expensive $1.7 billion estimate, which—to
emphasize—is the annual cost of running this system, not its overall
cost.
Admittedly, it is possible to cut some of the costs of legal aid,
primarily through domain restriction or through the use of means or
merit testing.171 Most clearly, the numbers given here include all issues
https://www.americanbar.org/content/dam/aba/administrative/legal_aid_indigent_defendants/ls_
sclaid_resolution_06a112a.authcheckdam.pdf [https://perma.cc/Q7YZ-H8Z2].
166. ALAN W. HOUSEMAN, CTR. FOR LAW & SOC. POLICY, CIVIL LEGAL AID IN THE UNITED
STATES: AN UPDATE FOR 2013, at 5 (2013), http://www.clasp.org/resources-and-
publications/publication-1/CIVIL-LEGAL-AID-IN-THE-UNITED-STATES-3.pdf
[https://perma.cc/573D-S4SB]. According to the report, the LSC is the largest provider of such
services, providing legal aid in eleven percent of the cases it handles.
167. See Jessica K. Steinberg, Demand Side Reform in the Poor People’s Court, 47 CONN. L.
REV. 741, 771 n.167 (2015). Steinberg extrapolates from an analysis made in Maryland, finding
that the national costs would be $5.4 billion. Id. However, this estimate is also conservative. It
assumes low payments to lawyers ($80 per hour), only four hours of work per case, and no overhead
and administrative costs, and also that pro bono services will not contract (a phenomenon known
as “crowding out”), that the rate of litigation will not increase, and that all those currently
represented will continue to hire a lawyer despite free legal services. See MD. ACCESS TO JUSTICE
COMM’N, Implementing a Civil Right to Counsel in Maryland, in ANNUAL REPORT 2010 app. 6 at
10 (2010), http://mdcourts.gov/mdatjc/pdfs/implementingacivilrighttocounselinmd2011.pdf
[https://perma.cc/L8HD-JWSK] [hereinafter Right to Counsel in Maryland]. Accounting for these
considerations would dramatically increase the costs involved.
168. INTEREST ON LAWYER ACCOUNT FUND OF THE STATE OF N.Y., ANNUAL REPORT 2014, at 2
(2014), https://www.iola.org/board/Grantee%20Annual%20Report%202014-
15/Annual%20Report%202014(final).pdf [https://perma.cc/F8GZ-TYUH].
169. See BURDGE, supra note 89, at 12.
170. To confirm this from another perspective, in Maryland, an average state in terms of
economy and inequality, one-sixth of the population qualifies for legal representation. See Right to
Counsel in Maryland, supra note 167, at 9 (reporting that approximately one million Marylanders
qualify for legal assistance from organizations funded by the Maryland Legal Services
Corporation). This would suggest a potential pool of at least fifty million eligible Americans
nationally.
171. Id. at 4.
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
2018] ADMINIZATION 161
where people might seek legal assistance, not only consumer litigation.
Restricting legal assistance to only the cases involving consumer credit
would be expected to reduce the overall costs of legal aid and make this
reform appear somewhat more realistic. While the cost reduction is
indeed likely, it is unlikely to be dramatic enough. As a preliminary
matter, someone would need to classify incoming complaints, and this
classification is costly and open to mistakes. More importantly, it is
necessary to recall the volume of consumer credit litigation—with over
eight million filings every year. Many of the people involved in such
litigation are likely to be in need of legal assistance, so that even by
itself, this category of cases is substantial.
A more promising avenue for cost reduction is means or merit
testing. Assistance could be made conditional on the financial needs of
consumers or the strength of the underlying case. By denying
assistance to people with means above a certain threshold, or to people
with weak cases, the costs of providing legal aid can be substantially
reduced. Both means and merit testing are indeed capable of cutting
costs, but they present their own issues. If the means threshold is high
(i.e., only people with low means are eligible), very few people will be
able to benefit from legal aid, which makes such reform unlikely to be
transformative enough. But if the threshold is set sufficiently low, the
whole point of means testing would be missed. Merit testing is likewise
difficult in this setting, because the consumers that are capable of
demonstrating the merits of their case to an administrator are those
who are least likely to need legal assistance in the first place, as they
could presumably also present their case to a judge. With respect to
both types of testing, it is also important to remember the
administrative infrastructure that would be required to support the
administration of testing, as well as the costs of both types of
mistakes—denying aid to deserving applicants and permitting aid to
irrelevant claims or claimants.172 From the consumer standpoint, such
testing often involves applicant-side costs and stigma,173 thus deterring
172. See Amartya Sen, The Political Economy of Targeting, in PUBLIC SPENDING AND THE POOR
11, 12–13 (Dominique van de Walle & Kimberly Nead eds., 1995); Wim van Oorschot, Targeting
Welfare: On the Functions and Dysfunctions of Means Testing in Social Policy, in WORLD POVERTY:
NEW POLICIES TO DEFEAT AN OLD ENEMY 171, 176 (Peter Townsend & David Gordon eds., 2002).
173. See Sen, supra note 172, at 13 (“Any system of subsidy that requires people to be identified
as poor and that is seen as a special benefaction for those who cannot fend for themselves would
tend to have some effects on their self-respect as well as on the respect accorded them by others.”);
Jennifer Stuber & Mark Schlesinger, Sources of Stigma for Means-Tested Government Programs,
63 SOC. SCI. & MED. 933, 944–45 (2006) (conducting empirical examination of the sources of
stigma); see also Bo Rothstein, The Universal Welfare State as a Social Dilemma, 13 RATIONALITY
& SOC’Y 213, 222–23 (2001) (explaining that in order to achieve sufficient support for such
measures, citizens must regard it as valuable and believe their fellow citizens are contributing).
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
162 VANDERBILT LAW REVIEW [Vol. 71:1:121
the neediest from seeking it.174 Overall, the discussed cost reduction
methods must make a difficult compromise: either set a low bar that
reduces the effectiveness of testing, or set a high bar but risk limiting
aid in a way that would mostly retain the status quo.175 While legal aid
can complement Adminization, it does not appear to be an appealing
substitute.
Not only are the costs of legal aid in this context extremely high,
the benefits are quite limited. Most cases are not genuinely disputed,176
and getting more consumers to court could drown the signal (valid
consumer defenses) in the flood of noise. Put formally, free
representation reduces Type I errors (enforcing unmeritorious claims)
but increases Type II errors (failing to enforce legitimate debts).
Whether one effect will be greater than the other is an open empirical
question, but even if the net effect is positive, the overall benefits will
be significantly limited by these costs. Moreover, the benefits will likely
be further diluted by rational creditor responses to such reforms, which
will likely consist of investing more in legal services to retain some of
their original advantage.177 This is even without taking into account
creditors’ market power and ability to influence the consumer contract
in ways that would mitigate the effects of legal access. At best, then, the
benefits will be modest but the costs will be immense. Equally worrying,
the costs are likely to expand over time, with little ability to control
them, as more and more people may claim eligibility.
174. Sen, supra note 172, at 13; van Oorschot, supra note 172, at 176; Dimitri Gugushvili &
Donald Hirsch, Means-Tested and Universal Approaches to Poverty: International Evidence and
How the UK Compares 2–3 (Ctr. for Research in Soc. Policy, Working Paper No. 640, 2014)
(empirical evidence); Thandika Mkandawire, Targeting and Universalism in Poverty Reduction
15–16 (Soc. Policy and Dev., United Nations Research Inst. for Soc. Dev., Working Paper No. 23,
2005).
175. These are not the only cost-cutting mechanisms. It is possible to offer a menu of more
limited services (such as a hotline for pro se claimants), and it is even possible to co-opt some of
the mechanisms developed in this paper, such as algorithmic screening of applicants and audit
review. Such proposals require sustained development and evaluation before they can be compared
to the alternatives considered here.
176. See infra notes 185–189.
177. Economic theory predicts that increasing one party’s investment in litigation (which is
similar to the effect of representation) can lead to an arms race that will greatly increase spending
but will not necessarily increase overall judicial accuracy. See Avery Katz, Judicial
Decisionmaking and Litigation Expenditure, 8 INT’L REV. L. & ECON. 127, 138–39 (1988). The
overall effect will be a reduction in the volume of litigation (as it is costlier to litigate) but an
increase in the intensity of litigation (because both parties “fight” harder). The net result requires
a more robust empirical analysis.
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
2018] ADMINIZATION 163
B. Throwing Judges into the Fray
Another common proposal is to have judges play a more active
role in litigation to level the playing field between the parties.178 Under
this view, judges should be more forgiving of consumers’ procedural
mistakes, allow more flexible deadlines, and furnish opportunities to
amend or correct what may be either mistakes or suboptimal litigation
tactics. According to more expansive versions, the judge would even
conduct examinations and seek settlements where possible.
This proposal is equally problematic. First, we do not know
whether inquisitorial systems produce systematically better results,
with a lingering concern that judges who produce their own evidence
are more prone to confirmation bias.179 Second, from an institutional
perspective, training judges to conduct inquisitorial functions requires
fundamental changes to the way legal education and training is
provided. But perhaps most troubling are the costs of these proposals.180
The more we ask judges to perform the activities of lawyers, the closer
we are to the first type of proposals, with public subsidies for private
lawyers. Discounting overhead and judicial staff, the median annual
salary of a judge is $132,500,181 compared with the median salary of a
178. See Steinberg, supra note 167, at 800 (“[J]udges should be active, frame legal issues, and
question parties and witnesses in order to develop legal claims.”); see also Amalia D. Kessler, Our
Inquisitorial Tradition: Equity Procedure, Due Process, and the Search for an Alternative to the
Adversarial, 90 CORNELL L. REV. 1181, 1274 (2005).
179. See Lon L. Fuller, The Adversary System, in TALKS ON AMERICAN LAW 30, 40 (Harold J.
Berman ed., 1961) (“An adversary presentation seems the only effective means for combating this
natural human tendency to judge too swiftly in terms of the familiar that which is not yet fully
known.”); Kathryn E. Spier, Litigation, in 1 HANDBOOK OF LAW AND ECONOMICS, supra note 116,
at 313–16 (presenting mixed theoretical accounts of the implications of an inquisitorial system);
John Thibaut et al., Comment, Adversary Presentation and Bias in Legal Decisionmaking, 86
HARV. L. REV. 386, 389–90 (1972) (noting that interested parties may vet evidence more thoroughly
than a judge). But see E. Allan Lind et al., Comment, Discovery and Presentation of Evidence in
Adversary and Nonadversary Proceedings, 71 MICH. L. REV. 1129, 1143 (1973) (arguing that bias
in evidence production will bias outcomes). On the empirical side, see Michael K. Block et al., An
Experimental Comparison of Adversarial Versus Inquisitorial Procedural Regimes, 2 AM. L. ECON.
REV. 170, 177–78 (2000) (finding that inquisitorial investigations fared better than adversarial
ones, but only if parties have no access or knowledge of the other party’s information and evidence).
180. See Resnik, supra note 118, at 380 (doubting that managerial judging reduces litigation
costs).
181. Survey of Judicial Salaries, NAT’L CTR. FOR ST. CTS. (2012),
http://www.ncsc.org/~/media/Files/PDF/Information%20and%20Resources/Judicial%20Salary/jud
icialsalaries.ashx [https://perma.cc/9AHV-WZMV]. The median salary for federal district court
judges is even higher, at $174,000 (both values are current to 2012). Judicial Compensation, U.S.
CTS., http://www.uscourts.gov/judges-judgeships/judicial-compensation (last visited Oct. 23, 2017)
[https://perma.cc/GZ5Z-A5ME].
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
164 VANDERBILT LAW REVIEW [Vol. 71:1:121
public interest attorney of $45,000–$75,000.182 The natural question is:
Would it not be cheaper to simply subsidize lawyers outright?
C. Modifying the Legal Process
The third type of reforms involve changes to legal procedure. For
example, to mitigate evidentiary problems, many propose that plaintiffs
should assert detailed knowledge of the claim, its origin, and of all other
evidence.183 Consequently, some states have imposed heavier
evidentiary burdens on creditors.184 Setting high evidentiary burdens
for the sake of controlling litigation may seem appealing, but it is a poor
solution to the problem at hand. Perhaps the most obvious point is that
in the absence of meaningful scrutiny, the mere production of evidence
cannot improve outcomes—similar to the problem of sewer service, how
can judges authenticate and verify the veracity of evidence? Moreover,
this proposal is extremely wasteful. Evidence production involves some
complex operations, as even discerning the amount of charges,
principal, and interest is not straightforward.185 Admittedly, the costs
per case are not high, but given the large volume of cases, these costs
quickly become a significant burden. Additionally, evidentiary bars are
182. Press Release, Nat’l Ass’n for Law Placement, New Public Interest and Public Sector
Salary Figures from NALP Show Little Growth Since 2004 (Oct. 18, 2012),
http://www.nalp.org/uploads/PressReleases/2012PISALPressRelease_Rev.pdf
[https://perma.cc/LEN7-N7BN]. For comparison, a lawyer’s median salary in 2012 was $113,530,
which is eighty-six percent of the salary of a state judge and sixty-five percent of the salary of a
federal district court judge. Bureau of Labor Statistics, May 2014 National Occupational
Employment and Wage Estimates: United States, U.S. DEP’T LAB.,
https://www.bls.gov/oes/2012/may/oes_nat.htm#23-0000 (last modified Mar. 25, 2015)
[https://perma.cc/B56N-S7X7]; see also sources cited supra note 181.
183. See Eric Y. Wu, Note, Vigilante Justice: Ensuring that Consumer Credit Plaintiffs Are Not
Above the Law in Collins Financial Services v. Vigilante, 60 AM. U. L. REV. 1561, 1563 (2011)
(advocating increased documentation requirements for default judgment); see also Fox, supra note
46, at 40 (documenting regulatory action aimed at increasing evidentiary standards); Glover,
supra note 46, at 1133 (“[A]t a bare minimum, courts should require the plaintiff to produce a valid
contract between the original creditor and the debtor.”); Goldberg, supra note 46, at 748 (proposing
that lawyers in small claims court be required to plead specific information about the debtor and
debt agreement); Debt Deception, supra note 92, at 2 (arguing in favor of “[p]rohibit[ing] debt
buyers from filing lawsuits without evidence”).
184. N.C. GEN. STAT. § 58-70-150 (2017) (requiring evidence of the original contract); see also;
MD. R. CIV. P. 3-306(d) (establishing additional requirements for affidavits filed by a plaintiff who
is not the original creditor); MASS. UNIF. SMALL CLAIMS R. 7(d) (outlining the circumstances that
impact whether further inquiry is needed in the event that a defendant does not appear for trial);
MINN. STAT. § 548.101 (2017) (proscribing requirements for assigned consumer debt default
judgments); Administrative Directive of the Chief Judge of the Court of Common Pleas for the
State of Delaware, No. 2012-2 (Aug. 22, 2012),
http://courts.delaware.gov/CommonPleas/docs/AD2012-2.pdf [https://perma.cc/TJ8Y-X5N7]
(establishing pleading requirements in consumer debt collection actions).
185. Duffy, supra note 42, at 1195 (“The amount due, however, is typically the result of
complicated, and often dynamic, contract terms . . . .”).
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
2018] ADMINIZATION 165
a “blunt calibration device[ ] . . . [that] risks screening out meritorious
and unmeritorious claims alike.”186 The concern that evidentiary bars
will deter the filing of meritorious lawsuits is heightened by the fact
that most cases are not even disputed. Evidence shows that in only 3.2
percent of debt collection cases by debt buyers did debtors informally
bother to dispute the debt.187 A qualitative in-depth (but small sample)
study found that only twenty percent of the cases were contested,
although perhaps half of them had some good faith defense of which
they were unaware.188 Some debtors, presumably those with the best
cases, do decide to go to court, but even those debtors fail about fifty
percent of the time.189 Even if we suppose that the rate of disputes
stands at the inflated rate of thirty percent, these reforms would require
the redundant production of evidence in all remaining cases (seventy
percent).190
Clearly, having robust evidence is also beneficial. Allowing court
judgments in the absence of evidence is a recipe for disaster. However,
the benefit of evidence only accrues if a sufficient number of cases are
scrutinized, which is hardly the case today.191 Additionally, there is also
an evidence requirement today, so one should consider whether the
marginal benefits that could result justify the requirement of high
evidentiary bars in all cases. Recall that Adminization is not meant as
a substitute, and setting evidentiary bars is recognized to be important.
The main contention here is that on the margin it would be more
productive to invest in administrative audits than to categorically
186. Engstrom, supra note 31, at 643.
187. FTC DEBT INDUSTRY REPORT, supra note 10, at 38. Admittedly, there are certain
recording issues involved, especially regarding verbal disputes. However, of the cases recorded,
only about fifty percent could be later verified, making the scope of genuine disputes much smaller.
Id. at 40. The FTC acknowledges that verbal disputes may not be properly recorded. Id. at 38; see
also Todd J. Zywicki, The Law and Economics of Consumer Debt Collection and Its Regulation 14–
15 (George Mason Univ. Law & Econ. Research Paper Series, Paper No. 15-33, 2015),
https://www.law.gmu.edu/assets/files/publications/working_papers/LS1517.pdf [https://perma.cc/
6M8L-4HN2].
188. See Sterling & Schrag, supra note 71, at 366 (studying claims against fifteen debtors).
Note, however, that the authors believe, based on interviews, that eight of the fifteen interviewees
had good-faith defenses of which they were personally unaware due to legal ignorance. Id. at 384.
189. See Holland, supra note 46, at 210 (finding that pro se debtors had at least some success
in fifty-three percent of the cases, although they only won trials in about one percent of the cases.
The represented debtors had favorable outcomes in about eighty-five percent of cases, but this only
applied to eight cases out of a sample of 4,400 cases). Id. Of course, trial outcomes are only
suggestive.
190. Means and merit testing would increase the benefit of evidence requirements, but would
introduce other types of errors and problems (e.g., debtors would “attorney-shop” for lenient
attorneys) and involve administrative costs. Most importantly, however, letting private attorneys
screen cases amounts to a de facto privatization of the process and, as such, should be evaluated
independently of the current system.
191. See Healy, supra note 1.
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
166 VANDERBILT LAW REVIEW [Vol. 71:1:121
require a higher bar of evidence in all eight million consumer credit
cases.
Other types of proposals offer conflicting recommendations on
the choice of venue. While in the past small claims courts have been
proposed as a solution, today some call for the transfer of cases to the
general civil courts, where a higher standard of proof might deter
creditors from filing.192 But this is similar to requiring more evidence,
and, as just argued, more evidence is unlikely to be the solution to the
problem. Others suggest that federal courts will provide a better
solution, due to their fee-shifting rules.193 Yet others suggest simply
narrowing creditors’ access to any court. Because litigation tends to be
lopsided, they reason, it will be best to allow litigation only after
creditors have exhausted informal collection efforts.194 However, since
informal collection is tainted with widespread abuse, it is hard to see
how such a proposal could improve the consumer’s situation.195
The last type of procedural reform tries to directly regulate
plaintiff behavior. For example, these proposals would require plaintiffs
to sign affidavits that they have taken due effort to locate the debtor,196
prove the timeliness of the claim,197 document service by means of GPS
technology, or educate the debtor of her rights.198 The problem with
these proposals, even setting aside their cost, is that they critically
depend on creditors with misaligned incentives. Financial incentives
exert a strong power, and as long as creditors stand to gain from
192. See, e.g., Goldberg, supra note 46, at 747–48.
193. See, e.g., Improving Relief, supra note 53, at 1464 (proposing a doctrine of “equitable
remand,” allowing federal courts to issue a vacatur of a state court judgment). This involves
reforming the Rooker-Feldman doctrine that limits federal courts’ power to intervene in state
courts’ judgments. See D.C. Court of Appeals v. Feldman, 460 U.S. 462 (1983); Rooker v. Fid. Tr.
Co., 263 U.S. 413 (1923).
194. See Goldberg, supra note 46, at 748 (“[T]o ensure that small-claims courts are truly a last
resort . . . lawyers should be required to inform the court of all prior communications with the
debtor and any extrajudicial collection efforts.”). In September of 2013, Minnesota passed a law
requiring creditors to provide advance notice of at least fourteen days to debtors of their intention
to file for a default judgment. MINN. STAT. § 548.101(a)(7) (2017); see also Glover, supra note 46,
at 1132 (advocating filing fees).
195. The authors of such proposals also seem aware of this inherent difficulty: “One possible
weakness . . . is that it may result in more aggressive extrajudicial collection pursuits and
consequently more violations of FDCPA.” Goldberg, supra note 46, at 749.
196. In Massachusetts, creditors in small claims courts are required to file a “Verification of
Defendant’s Address form, certifying that he or she has verified the defendant’s mailing address
in the manner set forth therein.” MASS. UNIF. SMALL CLAIMS R. 2(b).
197. See, e.g., Haneman, supra note 86, at 735–37.
198. See, e.g., FTC PROTECTING CONSUMERS REPORT, supra note 14, at 10 (“[J]urisdictions
should also consider amending service of process rules to require greater verification.”); OUT OF
SERVICE, supra note 68, at 3.
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
2018] ADMINIZATION 167
debtors’ failure to appear, they are bound to find loopholes and
shortcuts.199
D. Arbitration and Class-Defense
Two very different types of solutions include arbitration and
class litigation. Consumer arbitration is a growing trend.200 In theory,
it has various appealing characteristics that are relevant to some of the
problems Adminization addresses, most notably, arbitration’s ability to
overcome byzantine procedures and cut costs.201 Despite these benefits
(which many find empirically contestable202), arbitration does not solve
the structural issues that Adminization does. This is especially clear in
the case of consumer credit litigation where the FTC itself concluded
that arbitration fails to adequately protect consumers.203 The primary
reason for this failure is that arbitration is ultimately a contractual
instrument. As such, it tends to replicate the same market dynamics
that often lead to abuse in litigation. For example, creditors draft most
consumer agreements and affect the choice of arbitrators; as a result,
those cherry-picked arbitrators are often structurally impeded from
199. “[L]ife finds a way.” MICHAEL CRICHTON, JURASSIC PARK 160 (1990).
200. See AT&T Mobility, LLC v. Concepcion, 563 U.S. 333 (2011) (finding that the Federal
Arbitration Act preempted a California unconscionability rule for arbitration clauses, thereby
expanding the scope of such clauses in contracts related to class actions); David Horton & Andrea
Cann Chandrasekher, After the Revolution: An Empirical Study of Consumer Arbitration, 104 GEO.
L.J. 57, 70–76 (2015) (detailing the rise of consumer arbitration).
201. Arbitration allows consumers to avoid litigation costs, as emphasized by the industry. See
Letter from Am. Bankers Assoc. et al., to the Honorable Richard Cordray, Chairman, Bureau of
Consumer Fin. Prot. (July 13, 2015), http://www.aba.com/Advocacy/commentletters/Documents/cl-
jointArbitration2015.pdf [https://perma.cc/A5FB-LGMC]. Reducing costs for creditors could also
reduce the costs of consumer goods. See Stephen J. Ware, Paying the Price of Process: Judicial
Regulation of Consumer Arbitration Agreements, 2001 J. DISP. RESOL. 89, 91–93 (noting that the
passing of cost savings is an overlooked advantage for consumers of arbitration).
202. See JOHN O’DONNELL, PUB. CITIZEN, THE ARBITRATION TRAP: HOW CREDIT CARD
COMPANIES ENSNARE CONSUMERS 43 (2007), https://www.citizen.org/sites/default/files/
arbitrationtrap.pdf [https://perma.cc/W5A5-74T6] (highlighting the costs of arbitration to
consumers as a barrier). But see PETER B. RUTLEDGE, U.S. CHAMBER INST. FOR LEGAL REFORM,
ARBITRATION–A GOOD DEAL FOR CONSUMERS: A RESPONSE TO PUBLIC CITIZEN 10–12 (2008),
http://stmedia.startribune.com/documents/docload.pdf [https://perma.cc/3UEC-ZEKE] (critiquing
O’Donnell’s analysis); SEARLE CIVIL JUSTICE INST., NORTHWESTERN UNIV. SCH. OF LAW,
CONSUMER ARBITRATION BEFORE THE AMERICAN ARBITRATION ASSOCIATION 67–68 (2009),
https://www.adr.org/sites/default/files/document_repository/Searle%20Civil%20Justice%20Instit
ute%20Report%20on%20Consumer%20Arbitration.pdf [https://perma.cc/C9AF-GTDW]
(reviewing the empirical literature and finding mostly positive effects of arbitration for
consumers). Additionally, there is empirical doubt as to what extent consumers understand and
consent to arbitration clauses. See Jeff Sovern et al., “Whimsy Little Contracts” with Unexpected
Consequences: An Empirical Analysis of Consumer Understanding of Arbitration Agreements, 75
MD. L. REV. 1, 62–63 (2015) (finding broad consumer misunderstanding of arbitration clauses in
contracts).
203. See FTC PROTECTING CONSUMERS REPORT, supra note 14, at i.
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
168 VANDERBILT LAW REVIEW [Vol. 71:1:121
deterring fraud.204 Moreover, even well-intentioned arbitrators cannot
meaningfully investigate and audit cases where consumers do not
appear, and access to arbitration is still considered “prohibitively
expensive for consumers with relatively small claims.”205 A large study
of thirty-four thousand arbitration cases revealed statistics that are
similar to those of litigation, with ninety-four percent of arbitrations
being resolved in favor of creditors.206 The study also details evidence of
arbitrator shopping where pro-plaintiff arbitrators are sought more
often than prodefendant arbitrators.207 As a result, many are
disillusioned today with arbitration as a means of improving consumer
protection and remedying market flaws.208
Another alternative is the idea of class defense.209 Developed by
Assaf Hamdani and Alon Klement, the class defense mechanism is a
mirror image of the class action, only that it consolidates dispersed
defendants (rather than plaintiffs). When multiple defendants are sued
by a single plaintiff, the class defense mechanism would allow them to
be sued as a class, binding them all to the outcomes of litigation. The
aggregation of claims makes it more worthwhile to defend them, as the
joint stakes are large enough to pay a lawyer.210 Class defense has much
greater potential than the other proposals surveyed, primarily because
of its cost-effectiveness. Nonetheless, class defense is unlikely to fully
resolve the problems identified here. By their nature, class actions
apply only to cases meeting narrow criteria, and many consumer credit
204. See Richard M. Alderman, Pre-dispute Mandatory Arbitration in Consumer Contracts: A
Call for Reform, 38 HOUS. L. REV. 1237, 1256 (2001) (discussing the significant bias that favors
repeat players in the arbitration process); David S. Schwartz, Enforcing Small Print to Protect Big
Business: Employee and Consumer Rights Claims in an Age of Compelled Arbitration, 1997 WIS.
L. REV. 33, 60–61 (noting that corporate defendants may prefer arbitration over litigation due to
a belief that they will receive either sympathy or outright favorable bias from the arbitrator);
Nancy A. Welsh, Mandatory Predispute Consumer Arbitration, Structural Bias, and Incentivizing
Procedural Safeguards, 42 SW. L. REV. 187, 197 (2013) (noting that predispute arbitration is a
“system that is beset by structural bias”).
205. Sarah Rudolph Cole, On Babies and Bathwater: The Arbitration Fairness Act and the
Supreme Court’s Recent Arbitration Jurisprudence, 48 HOUS. L. REV. 457, 466 (2011).
206. See O’DONNELL, supra note 202, at 2 (examining approximately nineteen thousand cases
in which the arbitrator was appointed by the National Arbitration Forum).
207. Id. at 16–17 (exploring the incentives that exist for arbitrators to overwhelmingly side
against consumers). On the other hand, the Searle Civil Justice Institute found only weak evidence
of repeat-player effects in its review of the literature. See SEARLE CIVIL JUSTICE INST., supra note
202, at xiii.
208. See Cole, supra note 205, at 458–59 (detailing consumer concerns and attempted policy
responses to the growth in consumer arbitration agreements). But see SEARLE CIVIL JUSTICE INST.,
supra note 202, at 109–11 (reviewing the empirical evidence and discussing costs, due process,
speed, outcomes, and fairness considerations, and finding mostly positive effects).
209. See Hamdani & Klement, supra note 97, at 709–10 (proposing the mechanism of class
defense).
210. Class defense depends on fee shifting, so that if the class prevails, the representing
attorney can recover her fees from the plaintiff. Id. at 715–17.
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
2018] ADMINIZATION 169
cases would not meet those criteria. In this context, the most notable
issue is the requirement of commonality among the class members. The
Supreme Court’s recent decision in Wal-Mart Stores, Inc. v. Dukes
means that class members must share a very high degree of
commonality, and such degree of commonality is rare in civil litigation
generally, and especially in consumer credit contracts. 211 Having said
that, if implemented, a class defense mechanism could complement the
institution of Adminization.
E. A Pyrrhic Victory
Suppose, contrary to all of the foregoing, that these reforms
could work, and that they would bring a large portion of all consumers
to court. These consumers would plead and argue their cases and fight
against unfair charges, lack of evidence, fraud and abuse, or even more
technical issues, such as proper venue or setoffs and fees. Emboldened
and empowered, consumers would also appeal wrong decisions, and all
would have their day(s) in court. Consequently, reformers hope, the
accuracy of legal determination will rise and the scope of fraud and
consumer abuse will fall.
This optimistic view requires that the legal system will be able
to support this significant increase in litigation. The volume of civil
litigation is about fifteen million cases annually, and, as noted, it is
estimated that about eight million are consumer credit related.212
Today, in the vast majority of cases, consumers either do not appear or
do not respond. For example, the civil courts of the City of New York
saw 9,295 defendants out of 122,166 cases filed by nine large creditors
in 2008—this is about seven percent.213 Getting even one-third of all
New York consumers to appear means that the number of cases that
would be heard will rise from 9,295 to thirty-one thousand cases, and
an additional 4,340 appeals would be filed.214 Hence, any moderately
successful reform would then encumber the courts with a few million
new cases each year. This would increase the caseload of the same
courts that are currently criticized for being clogged and “overwhelmed”
211. See Wal-Mart Stores, Inc. v. Dukes, 564 U.S. 338, 349–360 (2011) (declining to certify a
class action due to lack of commonality among plaintiffs).
212. See supra note 42.
213. See Justice Disserved, supra note 71, at 4 (accounting for the volume of debt claims filed
by nine large creditors).
214. See Court Statistics Project, Caseload Highlights, NAT’L CTR. FOR ST. CTS. (Mar. 2007),
http://www.courtstatistics.org/~/media/Microsites/Files/CSP/DATA%20PDF/Vol14Num1CivilTria
lsonAppeal1.ashx [https://perma.cc/S6L9-KNA6] (finding a fourteen percent rate of appeal in all
civil matters).
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
170 VANDERBILT LAW REVIEW [Vol. 71:1:121
by consumer credit cases.215 To accommodate these cases, the system
would have to quadruple its capacity or else introduce impossibly long
queues to resolve disputes.216 Sending more cases to arbitration, it was
noted, is unlikely to be effective. Can we sustainably double, triple, or
even quadruple or more the national expense of the civil legal
system?217
This cost is “that which is seen,” but what about the cost “which
is not seen”?218 The less salient and more removed costs would be those
created by the response of creditors and debtors to such a change.
Debtors would be much more inclined to defend cases which have less
merit, hoping to win them through luck or through the attrition of the
creditor. Creditors will have to spend more resources on litigation to
win cases. Therefore, some creditors will not find it worthwhile to
pursue small claims, either because the case lacks merit or because the
costs exceed the value of the expected judgment. This will lead many to
215. See OUT OF SERVICE, supra note 68, at 11.
216. Not all of civil litigation is debt claims, so doubling the number of cases would lead to less
than double the resources. Nonetheless, debt claims are the majority of civil litigation, so the
necessary increase in resources in response to doubling the debt claims will be large. See supra
note 42. Also, some of the costs of litigation are fixed, so that more cases would not entail
necessarily more courtrooms. However, as we consider here, a very large increase in the volume of
cases would create a need for new infrastructure. The alternative to infrastructure, queueing, has
important costs well beyond the direct costs of waiting longer for cases to resolve; if the filing of a
lawsuit delays enforcement by a few years, this will provide a much stronger incentive to borrow
irresponsibly in the first place.
217. We do not know what the full cost of the legal system is, mainly because funding is
fragmented between local, state, and federal governments, as well as between user fees, fines,
charges, etc. Telephone Interview with William Raftery, Senior Knowledge & Info. Servs. Analyst,
Nat’l Ctr. for State Courts (Sept. 14, 2015). On the structure of funding, see RON MALEGA &
THOMAS H. COHEN, DEP’T OF JUSTICE, STATE COURT ORGANIZATION, 2011, at 8 (2013),
https://www.bjs.gov/content/pub/pdf/sco11.pdf [https://perma.cc/JNV8-HFF7]. For comparison, the
cost of the New York court system is $2 billion a year, and the cost of the Pennsylvania system is
$0.5 billion. Budget: Fiscal Year 2015-2016, N.Y. ST. UNIFIED CT. SYS., at v (2014),
https://www.nycourts.gov/admin/financialops/BGT15-16/2015-16-UCS-BUDGET.pdf
[https://perma.cc/PR29-DSMW]; Proposed Budget of the Unified Judicial System, Fiscal Year
2015–2016, SUP. CT. PA. (2015), http://www.pacourts.us/assets/files/setting-4778/file-
4372.pdf?cb=7127c2 [https://perma.cc/P6PC-QQZM] (noting that overall state contributions
amount to $11 billion a year); see also NAT’L GOVERNORS ASS’N & NAT’L ASS’N OF STATE BUDGET
OFFICERS, THE FISCAL SURVEY OF STATES: SPRING 2012 (2012),
https://www.nga.org/files/live/sites/NGA/files/pdf/FSS1206.PDF [https://perma.cc/KT83-3QKL];
Preparation and Submission of the Judicial Branch Budget, NAT’L CTR. ST. CTS.,
http://data.ncsc.org/QvAJAXZfc/opendoc.htm?document=Public%20App/SCO.qvw&host=QVS@qli
kviewisa&anonymous=true&bookmark=Document\BM09 (last visited Oct. 23, 2017)
[https://perma.cc/8PSQ-ZFXW] (outlining the various budget proposal processes across multiple
states).
218. See 1 CLAUDE FRÉDÉRIC BASTIAT, That Which is Seen, and That Which is Not Seen, in
THE BASTIAT COLLECTION 1, 1 (Ludwig von Mises Inst. ed., 2d ed. 2007) (1850). (“In the economy,
an act, a habit, an institution, a law, gives birth not only to an effect, but to a series of effects. Of
these effects, the first only is immediate; it manifests itself simultaneously with its cause—it is
seen. The others unfold in succession—they are not seen.” (emphasis added)).
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
2018] ADMINIZATION 171
invest more in informal debt collection219—an area that is hard to police
and is rife with abuse220—and others will altogether abandon small
consumer loans, an outcome undesirable for both creditors and debtors.
Indeed, there may be some important benefits to all of these
reform proposals: more legal accuracy, less fraud, and greater due
process and consumer participation. Adminization does not mean that
these routes should be abandoned; rather, we should diversify our
approach by using multiple institutions to address a multicausal
problem, thus improving outcomes and reducing costs. Adminization,
recall, is a complement, not a substitute, to litigation. Working in
tandem, litigation and Adminization pack a punch.
IV. CHALLENGES
A. Legal Authority
Adminization requires an agency infrastructure, and the use of
agencies to review consumer cases in state courts may introduce some
legal and constitutional issues. But such concerns are surmountable,
since Adminization relies on preexisting and legally proven
infrastructure such as the FTC, the CFPB, and state attorney general
offices.
At the federal level, both the FTC and the CFPB are already
legally empowered to oversee Adminization, although with certain
limits. For the CFPB, the primary source of authority is the Consumer
Financial Protection Act,221 enacted as Title X of the Dodd-Frank Act,
which established the CFPB and tasked it with implementing and
enforcing federal consumer protection law to ensure that “markets for
consumer financial products and services are fair, transparent, and
competitive.”222 A central objective is to protect consumers “from unfair,
deceptive, or abusive acts and practices.”223 To achieve these goals, the
law provides the CFPB with broad powers to conduct investigations,
request information from covered entities, issue subpoenas and civil
219. See Oren Gazal-Ayal & Limor Riza, Plea-Bargaining and Prosecution, in CRIMINAL LAW
AND ECONOMICS 145, 149 (Nuno Garoupa ed., 2009) (explaining that trial complexities lead to
greater frequency of plea bargains); Sergio G. Lazzarini et al., Order with Some Law:
Complementarity Versus Substitution of Formal and Informal Arrangements, 20 J.L. ECON. & ORG.
261, 262 (2004) (exploring substitution and complementarity between formal and informal norms).
220. See supra Section I.A; see also supra note 195 (explaining that advocates preferring
informal debt collection practices over litigation recognize inherent widespread abuse of such
practices).
221. Dodd-Frank Wall Street Reform and Consumer Protection (Dodd-Frank) Act § 1011(a),
12 U.S.C. § 5491 (2012).
222. 12 U.S.C. § 5511(a).
223. 12 U.S.C. § 5511(b)(2).
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
172 VANDERBILT LAW REVIEW [Vol. 71:1:121
investigative demands, hold hearings, bring lawsuits, and, importantly,
levy fines.224 The second relevant legislative authority is the Fair Debt
Collection Practices Act (“FDCPA”), which provides enforcement
powers to both the FTC and the CFPB.225 These powers are intended to
curb abusive debt collection practices and encourage fair debt collection
practices.226 The jurisdiction of both the FTC and the CFPB is
sufficiently broad, and their enforcement powers encompass any
provider of consumer financial services or its affiliates engaging in
“unfair, deceptive, or abusive acts or practices.”227 Moreover, the CFPB
has supervisory powers over large banks and certain “non-banks,” for
example, providers of credit such as payday lenders, auto lenders,
mortgage originators, and more recently, debt collectors with annual
earnings over $10 million, which is not a very high bar.228 In addition,
the FDCPA also empowers the FTC and the CFPB to investigate and
pursue actions against debt collectors.229 Using these broad powers,
these agencies regularly engage in enforcement activity that covers a
broad array of regulated entities, including banks, law firms that file
debt collection lawsuits with insufficient evidence, debt collectors, and
even individuals.230 As a consequence of these powers, both the CFPB
and the FTC have the necessary authority to support Adminization of
consumer debt litigation, allowing them to investigate and take
enforcement actions against those engaged in unfair, deceptive, or
abusive acts. Beyond the federal level, state attorney general offices are
generally equally empowered to investigate and prosecute consumer
abuse, although they are naturally limited to the jurisdiction of their
own states.231
A second related issue concerns the power of the federal agencies
to regulate consumer activity at the state level. This concern is directly
addressed by the FDCPA, which explicitly states that “[e]ven where
abusive debt collection practices are purely intrastate in character, they
224. 12 U.S.C. §§ 5562–65.
225. 15 U.S.C. § 1692l(a) (2012).
226. 15 U.S.C. § 1692(e).
227. 12 U.S.C. § 5531. For the definition of covered persons, see 12 U.S.C. § 1002(6).
228. 12 U.S.C. § 5514(a). Large debt collectors are considered “larger participants” under 12
U.S.C. § 5514(a)(1)(B), according to 12 C.F.R. § 1090.105(b) (2017).
229. 15 U.S.C. § 1692a(6).
230. See, e.g., In re Pressler & Pressler, LLP, CFPB No. 2016-CFPB-0009 (2016) (issuing
consent order regarding debt collection practices of a law firm); David Eghbali, CFPB No. 2016-
CFPB-0011 (2016) (issuing consent order regarding illegal mortgage-loan transaction
manipulations by a bank employee).
231. For example, the Consumer Interest Division at the Office of the Attorney General of the
State of Alabama is described as one of the Attorney General’s “most important responsibilities.”
Consumer Interest Division, ST. ALA. OFF. ATT’Y GEN., http://www.ago.state.al.us/Page-Consumer-
Protection (last visited Oct. 23, 2017) [https://perma.cc/AYZ3-7U8J].
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
2018] ADMINIZATION 173
nevertheless directly affect interstate commerce.”232 The Supreme
Court shared Congress’s view, holding that “commercial lending . . .
[has broad impact] on the national economy.”233 A large body of case law
affirmed this view.234 It is also worth mentioning that the CFPB has
recently emerged largely intact from a constitutional challenge that
sought to dismantle it.235 Whether the CFPB will continue to exist in
the current political climate is an open question, but such an issue does
not arise with respect to the other potential institutional arrangements.
B. Feasibility of Adminization & Political Economy
The success of regulatory reform does not depend solely on its
merits, but also on its political appeal. Would Adminization receive
sufficient political support? While predicting any sort of political
trajectory is difficult, I will note a few reasons why Adminization may
appeal to a variety of diverse interests that wield political power. First,
on the consumer side, this system of Adminization provides a robust
and meaningful form of protection that addresses some of the key
concerns of consumer associations today. Implementing Adminization
on a broad scale can improve the lives of millions of consumers over a
relatively short period. Consumer advocates and politicians seeking to
enhance the welfare of the middle and lower classes can mark a quick
victory with relatively little investment. It is equally important that
Adminization would also be appealing to creditors. As noted, they stand
to gain from a streamlined process, greater consumer confidence in the
credit market, and greater legitimacy of the debt collection process.
After all, there is reason to believe that greater legitimacy will translate
to greater consumer propensity to repay debts.236 Indeed, some
232. 15 U.S.C. § 1692(d).
233. Citizens Bank v. Alafabco, Inc., 539 U.S. 52, 58 (2003); see also Lewis v. BT Inv.
Managers, Inc., 447 U.S. 27, 38 (1980) (“[B]anking and related financial activities are of profound
local concern. . . . Nonetheless, it does not follow that these same activities lack important
interstate attributes.”); Perez v. United States, 402 U.S. 146, 154 (1971) (“Extortionate credit
transactions, though purely intrastate, may in the judgment of Congress affect interstate
commerce.”).
234. See supra note 16.
235. See PHH Corp. v. Consumer Fin. Prot. Bureau, 839 F.3d 1, 33 (D.C. Cir. 2016), reh’g en
banc granted, order vacated (Feb. 16, 2017). Note that while the D.C. Circuit panel decision holding
the CFPB unconstitutional was vacated and rehearing en banc granted, the court has not yet
issued its en banc opinion at the time of this writing. Thus there is still some uncertainty regarding
the CFPB’s future.
236. See, e.g., Jonathan Jackson et al., Why Do People Comply with the Law?, 52 BRIT. J.
CRIMINOLOGY 1051, 1059 (2012) (finding higher compliance from those who believe in the
legitimacy of the law); Tom R. Tyler, Psychological Perspectives on Legitimacy and Legitimation,
57 ANN. REV. PSYCHOL. 375, 377 (2006) (explaining legitimacy as the belief that the law should be
obeyed).
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
174 VANDERBILT LAW REVIEW [Vol. 71:1:121
creditors would object to Adminization, precisely because of its
effectiveness in deterring creditor fraud. But the opposition by creditors
will likely be much stronger to any of the other alternatives discussed
above, which tend to impose costs on all creditors, independently of the
nature of their claims. If not in absolute terms, then at least in relative
terms, Adminization should garner greater creditor support.
These mutual advantages to debtors and creditors promise a real
political possibility of implementation. As a case study, in Israel, where
a reform that streamlined the collection of small judgments was
proposed, an unlikely coalition emerged.237 Both creditors and
consumers joined hands in support of the reform; the creditors were
drawn to the streamlined process and the debtors to the greater
transparency and simplicity of the process as well as the concomitant
reduction in interest and fees.238 The only opponent was the Israel Bar
Association, which expressed concerns that the reform would make
parties less likely to retain the services of lawyers.239 Ultimately, the
Bar lost.240 It would seem that overall, Adminization offers a great
promise to both plaintiffs and defendants, and should be highly feasible
from a political-economy standpoint.
C. Regulatory Capture
In administrative law, there is a general concern with regulatory
capture.241 Here, one might worry that creditors will be able to lobby
the relevant agency, causing the agency to capitulate to their interests.
Without denying the potential dangers of regulatory capture in some
237. The initiative was not identical to Adminization; Israel already relies heavily on
administrative agencies to enforce small claims, and the reform proposal sought to create an
administrative process where lawyers would not be needed. In contrast, Adminization is meant to
supplement litigation.
238. See Protocol No. 661, LAW & JUST. COMMITTEE (July 24, 2012),
http://www.knesset.gov.il/protocols/data/rtf/huka/2012-07-24-02.rtf [https://perma.cc/EF32-
V5VY]; see also Protocol No. 7, LAW & JUST. COMMITTEE (June 16, 2015),
http://www.nevo.co.il/law_html/law103/20_ptv_309355.htm [https://perma.cc/H9X8-H3TM]
(endorsement of the process by a consumer NGO).
239. These concerns were expressed (and criticized by the court) in HCJ 6804/12 Israel Bar v.
Minister of Justice, unpublished (2013) (Isr.), http://elyon1.court.gov.il/files/12/040/068/
o10/12068040.o10.pdf [https://perma.cc/XK37-ZL9Y].
240. In the interest of disclosure, the author was a paid advisor for one of the commercial
companies supporting the reform.
241. See, e.g., STEVEN P. CROLEY, REGULATION AND PUBLIC INTERESTS: THE POSSIBILITY OF
GOOD REGULATORY GOVERNMENT 26–52 (2008) (noting that public choice theory is “an
abbreviation for analysis of how or why narrow regulatory interests routinely prevail over
others . . . [and] accounts for much academic skepticism toward public-law regulation”); Posner,
supra note 105, at 19 (“Agencies are subject to far more intense interest-group pressures than
courts.”).
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
2018] ADMINIZATION 175
cases,242 the concern in the abstract is unconvincing.243 Many
government agencies already operate in the world, and they are not all
hopelessly captured, despite a great variety of lobbyist groups.244 This
seems especially true of the CFPB, which has taken a hard line against
creditors during its years of operation. Additionally, consumers also
mobilize politically, as evidenced by the support garnered by politicians
such as Senators Elizabeth Warren and Bernie Sanders and political
movements such as Occupy Wall Street.245 Also, consumers are already
at a disadvantage in the courts today, as creditors are repeat players,
with more resources and greater ability to forum shop. Hence, neither
forum is immune to special interests.246 Ultimately, this challenge does
not appear especially worrying. The agency does not replace court
proceedings; it only adds an additional layer. Hence, the benefit to a
plaintiff of “capturing” the regulator is much diminished. Given the
great benefits that Adminization could provide and the low concrete
threat of regulatory capture, it would be misguided to let abstract
regulatory concerns inhibit meaningful reform.
D. Costs and Incidence
A final challenge relates to the cost of running the agency. This
concern may relate to the costs themselves or their “incidence,” i.e., the
idea that the public should bear the cost of Adminization. On reflection,
however, this challenge does not prove critical. Using the existing
platform of the CFPB means that set-up costs will be low. The most
242. See Daniel Carpenter & David A. Moss, Introduction to PREVENTING REGULATORY
CAPTURE: SPECIAL INTEREST INFLUENCE AND HOW TO LIMIT IT 3 (Daniel Carpenter & David A.
Moss eds., 2014) (“[O]bservers are quick to see capture as the explanation for almost any
regulatory problem . . . . At the same time, there appears to be a great deal of fatalism . . . about
the impossibility of ameliorating or preventing capture.”).
243. For general critique, see Engstrom, supra note 31, at 674–78. Engstrom argues, however,
that there is greater concern with capture when agencies conduct case-by-case adjudication. Id. at
678. His reasoning in this context seems to rely on a different model, where the agency substitutes
legal supervision rather than complements it—as Adminization does.
244. For a list of the most concentrated lobbyist groups, see Top Interest Groups Giving to
Members of Congress, 2018 Cycle, OPENSECRETS.ORG, http://www.opensecrets.org/industries/
mems.php (last visited Oct. 23, 2017) [https://perma.cc/P7YQ-BKKA] (select “2016” from the drop-
down list).
245. See Stephen Collinson, Does Elizabeth Warren Regret Not Running for President?, CNN
(Aug. 26, 2015), http://www.cnn.com/2015/08/25/politics/elizabeth-warren-joe-biden-elections-
2016/ [https://perma.cc/5DXH-PX4G] (predicting high potential for both Elizabeth Warren and
Bernie Sanders); 2016 Democratic Popular Vote, REALCLEARPOLITICS,
https://www.realclearpolitics.com/epolls/2016/president/democratic_vote_count.html (last visited
Oct. 23, 2017) [https://perma.cc/7Y6A-3XZ7] (noting twelve million votes in favor of Bernie
Sanders).
246. Cf. Posner, supra note 105, at 20 (“Courts are relatively immune to interest-group
pressures.”).
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
176 VANDERBILT LAW REVIEW [Vol. 71:1:121
significant cost is that of the audits, but the agency (and Congress) has
control over the frequency of audits, thus guaranteeing budgetary
control. For comparison, the IRS handles about 1.2 million audits every
year, which amounts to an audit rate of 0.8 percent of all its cases.247 To
achieve similar rates, Adminization would only require the auditing of
sixty-four thousand cases.248
Estimating the cost of audits is difficult; luckily, in situations
like these, Fermi Estimate often provides useful approximations
(within an order of magnitude).249 Collecting evidence in a case,
analyzing it, and contacting all the relevant parties should probably
take no more than ten hours on average for a skilled auditor. The
median annual salary of an IRS auditor is about $70,000,250 which,
using the standard divisor of 2,087 working hours per year, implies a
per-hour-cost of $33.50. To this we should add overhead, inefficiencies,
and some margin, so it is probably within a reasonable range to assume
that for every hour of work, an hour of similar cost should be added.
This means that the per-hour cost of audit is (again, using a very rough
estimate) about $70, giving us $700 per audited case, or a cost of $44.8
million. To verify, this estimate is consistent with the IRS estimate that
a case audit costs about $600.251 Even doubling this estimate, we are
still two orders of magnitude less than the cost of the leading
alternative. In fact, this cost is so low that there is reason to believe that
if Adminization reduces filings, it will be cheaper than the status quo,
where courts have to handle many cases that should not have been filed.
Another source of cost comes from the development of
algorithms. However, this cost would be largely a one-off expenditure
on development. Perhaps even more importantly from a policy
247. IRS DATA BOOK, supra note 112, at 21.
248. This is done under the assumption of eight million filings every year, of which 0.8 percent
would be audited. See supra note 42.
249. See Fermi Problem, WIKIPEDIA, https://en.wikipedia.org/wiki/Fermi_problem (last
updated June 8, 2017) [https://perma.cc/U7EH-WD2R] (“Fermi estimates generally work because
the estimations of the individual terms are often close to correct, and overestimates and
underestimates help cancel each other out.”).
250. See Average Salary for Tax Examiner, Collector, or Revenue Agent at U.S. Internal
Revenue Service (IRS), PAYSCALE, http://www.payscale.com/research/US/Employer=U.S.
_Internal_Revenue_Service_(IRS)/Salary/Job/Tax-Examiner%2c-Collector%2c-or-Revenue-Agent
(last visited Oct. 23, 2017) [https://perma.cc/3K5P-9BLR].
251. To verify, the IRS estimates that a $55 million budget increase will allow it to deal with
five hundred thousand additional cases (including individual audits, employment tax exams, and
collection matters). This implies a per-case cost of $574, which—despite the differences between
audited cases and the costs of handling the other types of cases—is still suggestive that the
analysis here is in the right order of magnitude. See IRS OVERSIGHT BD., FY2015 IRS BUDGET
RECOMMENDATION SPECIAL REPORT 1, 6 (2014), https://www.treasury.gov/IRSOB/reports/
Documents/IRSOB%20FY2015%20Budget%20Report-FINAL.pdf [https://perma.cc/J4FM-QV55].
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
2018] ADMINIZATION 177
perspective, Adminization can begin without these algorithms and
simply choose cases at random, similar to the IRS’s process.
The incidence objection holds that it should not be the public
purse that pays the costs of setting up an agency that delivers services
relating to purely private financial matters. While such objections may
have some merit in certain contexts, it is also of little relevance to the
case made here. This is because there are already large government
subsidies in place, namely the costs of running the courts, paying
judges’ and staff salaries, administrative processes, provision of legal
aid, etc. Court fees account for only twenty to thirty percent of the
overall cost of running the court itself, with the remainder being a
subsidy.252 Moreover, all the current reform proposals will involve a
much greater degree of subsidies. The question at this point is how to
best allocate those existing subsidies.
Overall, whatever costs Adminization entails and whatever
their incidence is, it is important to measure them in relation to both
the (impossibly) high costs of the alternatives and the benefits of
improving the system.
CONCLUSION
An old joke tells of a customer who dines in a restaurant and,
after finishing his meal, asks for the check, which the waiter promptly
brings him. The customer then decides to review the check in detail,
and discovers that, among the various items on the list, there is an
unrecognized item called “success.” Having no recollection of ordering
such a dish, the customer asks the waiter about the meaning of this
charge. “It is actually quite simple,” responds the waiter, “if the
customer pays, then it is a success.”
The success method is the calculated, strategic filing of
unmeritorious claims in the presence of lax screening mechanisms. This
Article demonstrates that civil litigation is systematically lacking in its
ability to screen unmeritorious claims in consumer credit litigation. The
review of the evidence shows problems of predatory debt collection
practices, sewer service, consumer underparticipation, lack of legal
representation, faulty and sometimes fraudulent evidence, and a lack
of supervision by judges. This results in a large system that invites the
use of the “success method.”
252. See GEOFFREY MCGOVERN & MICHAEL D. GREENBERG, WHO PAYS FOR JUSTICE?:
PERSPECTIVES ON STATE COURT SYSTEM FINANCING AND GOVERNANCE 17–18 (2014) (reporting
findings from Massachusetts and Utah, while noting, however, that in Florida the court system
seems to be profitable on net).
Electronic copy available at: https://ssrn.com/abstract=3015569

<<PAGE_BREAK>>

Arbel_Galley(Do Not Delete) 1/2/2018 4:17 PM
178 VANDERBILT LAW REVIEW [Vol. 71:1:121
The proposed solution to this problem is the use of a new mode
of regulation, in between courts and agencies, called Adminization: the
use of a gatekeeper agency to provide oversight when participation is
systematically lacking. It also offers a robust protection of due process
rights as a matter of both procedure and substance. This results in a
lean, cost-effective institution that could garner broad political support,
much more so than most of the other reform proposals currently
advocated. Consumers would enjoy greater access to justice at lower
costs and much broader protection of their rights. Creditors would
benefit from having greater consumer confidence in the credit market.
Future work will explore other applications of Adminization;
some prominent examples include housing, insurance and social
benefits fraud, employment law (suits against employees), civil rights,
and civil forfeitures. Each unique context brings its own nuance and
sensibilities, and the framework presented here can be usefully adapted
to meet these considerations. With the advance of algorithmic
decisionmaking, the growing budgetary pressures on courts, and the
pressure to improve outcomes for consumers of the legal system,
Adminization offers a glimpse into the future of our systems of
regulation.
Electronic copy available at: https://ssrn.com/abstract=3015569`

func main() {
    fmt.Println(articleText)
}
