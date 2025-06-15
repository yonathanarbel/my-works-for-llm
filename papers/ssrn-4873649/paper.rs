use std::collections::HashMap;
use std::fs::File;
use std::io::prelude::*;
use rand::Rng;
use serde_json::Value;
use reqwest;
use tokio;

const ARTICLE_TEXT: &str = r#"JUDICIAL ECONOMY IN THE AGE OF AI
YONATHAN A. ARBEL ∗
Individuals do not vindicate the majority of their legal claims
because of access to justice barriers. This entrenched state of
affairs is now facing a disruption. Lawyers and non-lawyers
alike are adopting artificial intelligence (AI) tools to perform
legal tasks—tools that sharply reduce the costs of generating
legal materials. There is finally hope that AI might allow
many more to access justice.
Paradoxically, what we gain in access to justice we might lose
in the delivery of justice. The problem is not that AI tools are
ineffective. Indeed, they are even more effective than most
realize—affecting every stage of the naming, blaming, and
claiming process. The problem is that this change necessarily
increases the volume and verbosity of the caseload thus
threatening judicial economy; the balance of scarce judicial
resources in relation to shifts in demand for legal services.
Historically, judges and legislatures have often met
challenges to judicial economy by adjusting “legal
thermostats”: ad-hoc adaptations to procedural rules and
even substantive doctrines meant to curb the flow of litigation.
But these adaptations invariably imply the shrinking of
substantive rights. We run the risk, then, that litigants who
finally gain access to justice will find narrow rights and
stringent administrative procedures. To avoid this trajectory,
I advocate a proactive framework of AI integration. Instead of
fighting a losing battle against the symptoms of AI adoption
by litigants, the legal system should integrate AI tools to
enhance and scale up the legal process itself. By thoughtfully
* Alfred Rose Professor of Law, Silver Faculty Scholar, University of Alabama
School of Law. The author would like to thank Matt Tokson, Russell Gold, Benjamin
McMichael, Mirit Eyal-Cohen, Marcus Gadson, Heather Elliott, and Richard Re.
Justin Heydt provided invaluable research assistance. The editors of the University
of Colorado Law Review have provided exceptionally careful edits and their
contribution is notable. A version of this Essay was prepared for the 2024 Judges
Forum of the National Civil Justice Institute, on Artificial Intelligence and the
Courts.

<<PAGE_BREAK>>

550 UNIVERSITY OF COLORADO LAW REVIEW [Vol. 96
and carefully incorporating these tools, we can ensure that we
reap the fruits of greater access to justice, even in the face of a
rapidly expanding caseload.
INTRODUCTION .......................................................................... 550
I. THE AI LITIGATION BOOM .............................................. 557
A. AI Legal Efficacy ...................................................... 557
B. AI Uptake ................................................................. 563
C. AI Impact on Access to Justice ................................ 565
II. LEGAL THERMOSTATS ..................................................... 569
III. LEGAL STRATEGIES THAT DEAL WITH THE AI
LITIGATION BOOM ........................................................... 576
A. Strategy 1: Legal Thermostats: Fees, Pleading
Standards, and Substantive Standards .................. 577
B. Strategy 2: Sit and Wait .......................................... 580
C. Strategy 3: Ban and Mark ....................................... 582
D. Strategy 4: Massive Funding .................................. 583
E. Strategy 5: Integration ............................................ 584
IV. CONCLUSION ................................................................... 592
INTRODUCTION
Most legal disputes are not filed anywhere. While estimates
on access to justice barriers are notoriously unreliable,1 a recent
study suggests that about 120 million legal problems are left
unresolved every year.2 Around 75 percent of low-income
Americans suffer significant civil legal issues, but 92 percent of
these problems receive little to no legal aid.3 One commentator
estimates that one hundred million Americans live with “civil
1. See generally Rebecca L. Sandefur, Paying Down the Civil Justice Data
Deficit: Leveraging Existing National Data Collection, 68 S.C. L. REV. 295 (2016)
(“In the arena of civil justice, we face a severe data deficit.”). On the various barriers
to access, see infra Section I.C.
2. INSTITUTE FOR THE ADVANCEMENT OF THE AMERICAN LEGAL SYSTEM,
JUSTICE NEEDS AND SATISFACTION IN THE UNITED STATES OF AMERICA 8, (Sept.
1, 2021) [hereinafter JUSTICE NEEDS], https://iaals.du.edu/sites/default/files
/documents/publications/justice-needs-and-satisfaction-us.pdf [https://perma.cc
/7VW8-Q3WM]. For comparison, one estimate considers that 100 million cases are
handled by state courts every year. State of the State Courts: 2022 Presentation,
NCSC (2022), https://www.ncsc.org/__data/assets/pdf_file/0019/85204/SSC_2022
_Presentation.pdf [https://perma.cc/5D6L-YMQK].
3. LEGAL SERVS. CORP., FY 2025 BUDGET REQUEST 5, [hereinafter FY 2025
BUDGET REQUEST] https://lsc-live.app.box.com/s
/oi1atcgn8xmvofc70aildz3bhg5p0zn5 [https://perma.cc/D7DE-9C78].

<<PAGE_BREAK>>

2025] JUDICIAL ECONOMY IN THE AGE OF AI 551
justice problems,” many of which affect their “basic human
needs.”4
The barriers to justice are legion, but most can be expressed
in terms of cost.5 Lawyers charge an average of $292 per hour,6
with common disputes costing between $2,754 and $6,370.7 On
the other side of the cost spectrum, commercial actors will spend
roughly $2 million in outside legal fees to litigate in full cases.8
Diverse faces and narratives lie behind these numbers, such as
Eloisa Veles a Queens resident who recently lost her factory job.9
A local family hired her as a housekeeper, promising $600 per
week, only to “stiff” her and pay $300 when the time came. More
telling than the incident itself is how it is described: Eloisa did
not have her contract breached, her rights violated, or her money
stolen—she was “stiffed.”10
The sheer size of the investment required to close the gap
bedevils attempts to resolve access to justice problems. Even
doubling legal aid budgets has done little to narrow the gap.11
Due to resource constraints, 1.8 million people are turned down
4. Rebecca L. Sandefur, What We Know and Need to Know about the Legal
Needs of the Public, 67 S. C. L. REV. 443, 446 (2016).
5. See generally DEBORAH RHODE, ACCESS TO JUSTICE (2004). See also Gillian
K. Hadfield, Legal Markets, 60 J. ECON. LIT. 1264, 1291 (2022) [hereinafter Legal
Markets] (“The principal reason that so few individuals and small businesses avail
themselves of legal services is cost and availability.”). See also Gillian K. Hadfield,
Higher Demand, Lower Supply? A Comparative Assessment of the Legal Resource
Landscape for Ordinary Americans, 37 FORDHAM URB. L.J. 129 (2010) (noting that
access to justice affects not just poorer Americans but also middle America). On
sociolegal barriers, see discussion infra Section I.C.
6. LEGAL TRENDS REPORT, CLIO 14 (2023), https://clio.drift.click/2023-ltr
[https://perma.cc/RG3K-HTRP].
7. See JUSTICE NEEDS, supra note 2, at 47.
8. LAWS. FOR CIV. JUST. REFORM GRP. & U.S. CHAMBER INST. FOR LEGAL
REFORM, LITIGATION COST SURVEY OF MAJOR COS. 14 (2010), https://
www.uscourts.gov/sites/default/files/litigation_cost_survey_of_major_companies
_0.pdf [https://perma.cc/AC3L-268A].
9. Noam Scheiber, Stiffing Workers on Wages Grows Worse with Recession,
N.Y. TIMES (Sept. 3, 2020), https://www.nytimes.com/2020/09/03/business/economy
/wage-theft-recession.html [https://perma.cc/2AMX-M3Q3].
10. I discuss legal consciousness as a barrier to justice. See discussion infra
Section I.C.
11. According to the Legal Services Corporation data, between 2013–2022, total
funding for legal aid has increased (inflation adjusted) from $1 billion to
$1.76 billion. See LEGAL SERVS. CORP., BY THE NUMBERS 2022: THE DATA
UNDERLYING LEGAL AID PROGRAMS 11 (2023) [hereinafter BY THE NUMBERS 2022],
https://lsc-live.app.box.com/s/h2bajpr3gps4s4a1iio6fwiddhmu1nwb [https://
perma.cc/UQ7R-LZLE]; Nora Freeman Engstrom & David Freeman Engstrom, The
Making of the A2J Crisis, 75 STAN. L. REV. ONLINE 146, 153 (Apr. 2024). (“[E]ven a
vast increase over current commitments would barely dent the current crisis.”).

<<PAGE_BREAK>>

552 UNIVERSITY OF COLORADO LAW REVIEW [Vol. 96
annually.12 To put this in perspective, the rate of legal aid
lawyers to eligible clients is 1 to 15,625.13
Recently, Nora and David Freeman Engstrom have sought
to center the problem of access to justice around legal tech.14
While others have already noted legal tech as a potential
barrier,15 they draw on the debt collection litigation literature
to fashion a somewhat different argument.16 As this literature
demonstrated, this is an area where there is a systemic access
issue for low-income defendants, who often cannot afford to
mount an effective defense even when one exists, resulting in a
default-judgment mill against them.17 The Engstroms frame the
asymmetry in power as resulting from an underlying asymmetry
in legal tech adoption patterns.18 While firms zealously adopt
legal tech, they only see “anemic adoption” by individuals.19 In
particular, they claim that large firms systemize and automate
litigation, whereas individuals are still reliant on “analog
tools.”20 While this argument is too strong to be true, it does
have a kernel of truth to it.21 Or at least it used to.
12. FY 2025 BUDGET REQUEST, supra note 3, at 4.
13. Hanna Kozlowska, There’s a Devastating Shortage of Lawyers in the U.S.
Who Can Help the Poor with Eviction or Child Custody Cases, QUARTZ
(May 12, 2016), https://qz.com/681971/for-every-10000-poor-people-in-the-united-
states-theres-less-than-1-lawyer-who-can-help-them-with-an-eviction-or-child-
custody-case [https://perma.cc/U3UC-VKXH].
14. See Engstrom & Engstrom, supra note 11. But see Legal Markets, supra
note 5, at 1303 (arguing that regulation favors traditional lawyering across the
board at the expense of legal tech).
15. See Legal Markets, supra note 5.
16. See generally Yonathan A. Arbel, Adminization: Gatekeeping Consumer
Contracts, 71 VAND. L. REV. 121 (2018) (discussing robo-signing and other
problematic creditor practices in debt collection cases and offering
administrative-technological solutions); Daniel Wilf-Townsend, Assembly-Line
Plaintiffs, 135 HARV. L. REV. 1704, 1773 (2022) (“Assembly-line plaintiffs show no
sign of slowing down. Because of both the increases in consumer debt and the
improvements in their litigation technology, they continue to grow . . . .”).
17. Wilf-Townsend, supra note 16, at 1773.
18. Engstrom & Engstrom, supra note 11, at 159.
19. See id. at 162. This asymmetry is also discussed in Yonathan A. Arbel &
Roy Shapira, Theory of the Nudnik: The Future of Consumer Activism and What We
Can Do to Stop It, 73 VAND. L. REV. 929, 962 (2020) (focusing on the concern that
firms employ advanced tools to defang litigation-prone consumers at very early
stages of their claiming process).
20. See Engstrom & Engstrom, supra note 11, at 163.
21. Most litigants rely on the Internet and other digital tools to amass
information, communicate about it, and draft and file litigation. See, e.g., Margaret
Hagan, Data on People’s Reliance on the Internet for Legal Problems, A BETTER
LEGAL INTERNET (Nov. 2, 2022), http://betterinternet.law.stanford.edu/2022/11/02

<<PAGE_BREAK>>

2025] JUDICIAL ECONOMY IN THE AGE OF AI 553
We are now witnessing a sea change in the patterns of
technological adoption. Most are by now familiar with the
occasional news story of a hapless lawyer using AI to comedically
bad outcomes.22 The narrative involves a work-shy lawyer
submitting an AI-generated and hallucination-riddled brief to
an exasperated judge, who then admonishes and sanctions the
lawyer. Such widespread stories seem to draw their memetic
power from commonplace Shakespearean perceptions of our
profession. Incidentally, they also reify an elitist notion that only
artisanal lawyering is real lawyering. And perhaps most
alluring, they affirm a comforting thought: Getting down to
brass tacks, AI is but a cold machine that will not be able to
usurp our jobs.
Reassuring and entertaining as such surface themes are,
they also distract from the broader reality that they unwittingly
reveal. These stories display how AI is being deployed in
practice, with two surprising patterns. First, they are being
adopted even by small law firms who, at least traditionally, are
rarely early adopters of cutting-edge technologies. Second, they
are being adopted despite broad knowledge that these tools are
imperfect. The point being that even if these tools are only
sometimes reliable, they are always convenient. And this
convenience and accessibility seem to drive many end users.
The expected outcome of democratizing litigation technology
is a sharp pruning of the cost of producing legal materials.23 As
such, the technology presents a heavyweight contender to the
many barriers to justice that plague the system. The expected,
indeed, desired, effect is a litigation boom, driven by those
currently denied access to justice. And while our first instinct
might be to celebrate the dismantling of access to justice
/data-on-peoples-reliance-on-the-internet-for-legal-problems [https://perma.cc
/A65A-PG7D]; see also Benjamin H. Barton, The Future of American Legal Tech:
Regulation, Culture, Markets, in LEGAL TECH AND THE FUTURE OF CIV. JUST. 21, 29
(David Freeman Engstrom ed., 2023) (“Nor has legal aid shied away from using
technology to forward its mission.”).
22. See, e.g., Benjamin Weiser, Here’s What Happens When Your Lawyer Uses
ChatGPT, N.Y. TIMES (May 27, 2023), https://www.nytimes.com/2023/05/27
/nyregion/avianca-airline-lawsuit-chatgpt.html [https://perma.cc/V6ZM-64RV];
Molly Bohannon, Lawyer Used ChatGPT In Court—And Cited Fake Cases. A Judge
Is Considering Sanctions, FORBES (June 8, 2023), https://www.forbes.com/sites
/mollybohannon/2023/06/08/lawyer-used-chatgpt-in-court-and-cited-fake-cases-a-
judge-is-considering-sanctions [https://perma.cc/HP4U-7PDD].
23. For cost comparisons between human lawyers and state-of-the-art AI
models, see infra pp. 8–9 and note 35. The point here is static, but there are
important dynamic effects, given that costs will decline across the industry.

<<PAGE_BREAK>>

554 UNIVERSITY OF COLORADO LAW REVIEW [Vol. 96
barriers, realism about judicial economy cautions great care.
The question we must ask ourselves is whether a legal system
already critiqued for being clogged and dilatory, a system whose
judges are overworked and under-resourced, will be capable of
handling the impending AI boom in litigation.24 What changes
will be made to our laws, rules, and standards to accommodate
such a spike? What will be the knock on effects of such a
disruption to the status quo? Ultimately, would we find
ourselves with a system with a truly greater degree of access to
justice?
My prescriptive thesis, in a nutshell, is this: We should not
sit and wait until a litigation boom forces our hand. The early
evidence suggests that AI is being integrated within legal
practices across the country. The legal system, I shall argue,
should keep pace. True, the AI systems of today are still
unreliable. Yet this should not be a deterrent, but a catalyst. It
should serve as a catalyst for forward-looking, proactive
integration that is subject to rigorous understanding of judicial
needs, system constraints, and AI testing. The goal is not only to
stanch a rising wave of litigation or stretch the justice dollar a
bit further; it is to proactively leverage the technology to scale
up and improve the delivery of justice without sacrificing justice
in individual cases.
This Essay seeks to sound the alarm about judicial economy
in the age of AI, consider how judges and legal administrators
might respond, how threats to judicial economy could jeopardize
rights, and then offer constructive steps to mitigate those
undesired side effects while expanding access and quality in the
delivery of justice. The Essay is organized around three principal
contributions.
First, the Essay argues that as AI erodes access barriers it
can bring about a litigation boom. The size of this boom is
commensurate with the access to justice gap, if not larger.
Existing estimates suggest that there is a considerable volume
of unmet demand for legal services.25 I argue, drawing on legal
sociology, that these estimates likely understate the true AI
potential.26 Beyond visible barriers like court and lawyer fees,
24. See Justice Delayed Judge and Staff Shortages are Leaving Americans in
Limbo, THE ECONOMIST (July 13, 2023), https://www.economist.com/united-states
/2023/07/13/judge-and-staff-shortages-are-leaving-americans-in-limbo [https://
perma.cc/6XZF-AJX8].
25. See BY THE NUMBERS 2022, supra note 11.
26. See discussion infra Section I.C.

<<PAGE_BREAK>>

2025] JUDICIAL ECONOMY IN THE AGE OF AI 555
sociolegal literature suggests that there are much less visible
barriers at very early stages. These barriers are succinctly
captured by the naming-blaming-claiming (NBC) model of
litigation, which is a tripartite process of transforming
individual claims.27 For an individual to even see themselves as
having a valid legal claim that is entitled to redress, they must
undergo three stages of reconceptualizing the “accident” or
“misfortune” they suffered as a legal wrong for which another
might be held responsible. These stages act as filters, and when
individuals lack the tools to name, blame, and claim, their claims
will be in a perpetual stage of arrested development. As
discussed and illustrated below, AI can assist with these pent-up
claims by shepherding individuals through the process, helping
them articulate their misfortune in legally cognizable terms.
Less rosy, existing estimates predominantly focus on
unaddressed meritorious claims.28 However, the same filtering
mechanisms that obstruct access to justice also serve beneficial
purposes by excluding abusive litigation aimed at harassing
individuals with trumped-up charges.29 The erosion of access
barriers would lead to a litigation boom of both types of
litigation, and the net effect is difficult to anticipate with any
confidence.
Second, the Essay draws on control theory—the study of
dynamic systems capable of maintaining desired states despite
internal and external disturbances—to consider the implications
of a potential AI litigation boom.30 The entire equilibrium of
judicial economy hangs in the balance between litigation
patterns and judicial resources.
One repeated lesson from legal history is that technological
and social shocks that threaten judicial economy are met with
adjustments of various procedural and substantive doctrines.31
27. The model was developed by William Felstiner. See William L. F. Felstiner
et al., The Emergence and Transformation of Disputes: Naming, Blaming,
Claiming . . ., 15 LAW & SOC’Y REV. 631 (1980). It has since become a mainstay of
socio-legal analysis.
28. See BY THE NUMBERS 2022, supra note 11.
29. Paul Ohm and Brett Frischmann developed a framework for thinking about
the positive effects of friction as tools of governance, and many of litigation barriers
can be conceived along similar lines. See Brett Frischmann & Paul Ohm,
Governance Seams, 36 HARV. J.L. & TECH. 1118 (2023).
30. See infra Part II. Control theory is devoted, loosely speaking, to the study
of maintaining desired states in dynamic systems. Home thermostats are a common
example of tools used by control theory to maintain temperature equilibrium in
light of changing outside temperature.
31. See discussion infra Part II.

<<PAGE_BREAK>>

556 UNIVERSITY OF COLORADO LAW REVIEW [Vol. 96
Even though these doctrines are ostensibly about substantive
and procedural rights, they double as what I call “legal
thermostats.” This effect can be broad and deep. Orin Kerr
famously argued that the entire body of Fourth Amendment law,
often seen as erratic and “embarrassing,”32 can be rationalized
as a series of “equilibrium adjustments” the courts make in
response to new technologies. Here, I generalize this insight to
a broader phenomenon of legal thermostats and provide
illustrations of how they are used across the justice system.
By trying to achieve homeostasis, judges may feel compelled
to adjust the thermostats that are at their disposal. They would
reach out, by necessity, to procedural and substantive rights.
They would be pressured to require, perhaps, more demanding
standards of proof, or may require more exacting evidence, or
may expand the scope of what qualifies as de minimis. The
degree of thermostat adjustment may be so large that, from the
viewpoint of any individual litigant, there would be no sense of
progress. They would overcome initial barriers only to crash on
ever more limited rights. If we stay the course, it seems that we
might squander the opportunity to make a real dent in the
access to justice problem by simply reshuffling it.
The third and most practical contribution lies in considering
the menu of reactions judges and judicial administrators can
make to take advantage of this specific moment. The proposed
course of action involves a proactive approach that works to
integrate AI into the judicial process itself. There is a host of AI
tools, some currently in production and others to come, that
could streamline, facilitate, and even improve the processing of
legal claims by the legal systems. They can be integrated at both
the case management level and inside the chambers themselves.
Integrating these tools into the legal process will allow the
system to scale up and meet the challenge, without
compromising the substantive rights of litigants. Grounding the
case for judicial integration in the problematic nature of the
realistic alternatives helps motivate adoption even if AI tools are
imperfect. Doing so proactively today will help mitigate the
harms and ensure responsible adoption.
32. See Orin S. Kerr, An Equilibrium-Adjustment Theory of the Fourth
Amendment, 125 HARV. L. REV. 476 (2011).

<<PAGE_BREAK>>

2025] JUDICIAL ECONOMY IN THE AGE OF AI 557
I. THE AI LITIGATION BOOM
How much of a dent can we realistically expect advanced AI
systems to put in the access to justice problem? This Part opens
by first evaluating the technical skills of current-generation AI
systems to establish that they can perform many legal tasks
“adequately.” Obviously, adequately is the load-bearing part of
the sentence, but part of the goal here is to show that it covers a
fairly broad range of legal capabilities.
The discussion then considers the adoption patterns among
end users, ordinary folks who currently face access issues, as
well as the size of the access to justice gap. It leverages these
analyses to provide a qualitative and semi-quantitative sense of
the size of the gap that could be bridged. The combination of
cheap but capable AI systems with this large gap leads to the
expectation of an AI litigation boom effect in the coming years.
A. AI Legal Efficacy
Any sufficiently advanced technology can appear
indistinguishable from magic.33 In practice, much commentary
on AI seems to fall into this trap, leading commentators down
one of two erroneous paths: either believing in AI omnipotence
(AI can do everything) or in AI as a cheap magic trick (AI can’t
do anything). In reality, AI tools are both, neither, and
in-between these poles. The goal of this Section is to avoid a
simplistic view of AI and discuss examples of the current state
of the art in legal AI.
Evaluating rapidly developing technology is an exercise in
writing on ice. The evidence of capabilities known to us today
shows tentative floors, while limitations are tentative ceilings.34
We do not know which limitations are here to stay, and which
can be resolved with future development. We only know that we
are still in early stages of development, and that we are still
seeing constant improvements.
33. ARTHUR C. CLARKE, PROFILES OF THE FUTURE: AN INQUIRY INTO THE
LIMITS OF THE POSSIBLE 36 (1962).
34. See Yonathan A. Arbel & Samuel Becher, Contracts in the Age of Smart
Readers, 90 GEO. WASH. L. REV. 83 (2022) (discussing the capabilities of smart
readers as well as the risks associated and the need to regulate and integrate with
caution).

<<PAGE_BREAK>>

558 UNIVERSITY OF COLORADO LAW REVIEW [Vol. 96
The first piece of evidence comes from a recent study that
evaluated AI on contract review tasks.35 The models were
presented with a contract and some necessary context, and then
asked to locate and determine legal issues. Comparing against
the benchmark of practicing lawyers, the researchers found that
GPT-4 (the current model powering ChatGPT) “exhibited a level
of accuracy in identifying legal issues that was on par with that
of [j]unior [l]awyers.”36 To complete their tasks, models use only
8 percent of the time it would take a junior lawyer to perform
them. Critically, where the lawyer would charge an average of
$74.26 for the task, the model’s operating cost was a single
quarter.37
While the models were relatively accurate, they were not
perfect, and their failure modes prove interesting. Relative to
senior lawyers, models showed “a preference for precision over
recall,”38—that is, they preferred to be accurate rather than
comprehensive. This offers greater confidence in the issues
identified, but risks overlooking some issues. This type of failure
mode, however, is not much different than that exhibited by
junior lawyers, who also showed a similar preference for
precision over recall, as evidenced by their comparable F-scores
in issue determination (0.86 for junior lawyers versus 0.87 for
GPT-4-1106).39 In addition, the authors provide two illustrative
examples of mistakes. On close review, these mistakes appear
transient and model-specific rather than fundamental. Indeed,
when I presented these examples to newer models (Claude Opus
35. Lauren Martin et al., Better Call GPT: Comparing Large Language Models
Against Lawyers, ARXIV (Jan. 24, 2024), https://arxiv.org/html/2401.16212v1
[https://perma.cc/GC33-3H9J]. There are other claims, less open to scrutiny, about
artificial intelligence and machine learning systems replacing lawyers in various
repetitive tasks. For example, JP Morgan reports of a software that reviews
contracts and “reviews approximately 12,000 new wholesale contracts per year and
replaced ‘360,000 hours’ of staff time between lawyers and loan officers.” Hugh Son,
JPMorgan Software Does in Seconds What Took Lawyers 360,000 Hours,
BLOOMBERG (Feb. 27, 2017), https://www.bloomberg.com/news/articles/2017-02-28
/jpmorgan-marshals-an-army-of-developers-to-automate-high-finance?embedded-
checkout=true [https://perma.cc/J548-GSUB].
36. Martin et al., supra note 35, at 12.
37. See id.
38. Id.
39. Id. at 8. An F-score (or F1 score) is a measure used to evaluate how well a
test or model performs, particularly in balancing two key aspects: precision (how
many identified items are correct) and recall (how many correct items were
identified). Id.

<<PAGE_BREAK>>

2025] JUDICIAL ECONOMY IN THE AGE OF AI 559
3 and Google Gemini Pro), both answered them correctly without
any tuning.40
A related study evaluated the ability of large language
models (LLMs) to serve as “smart readers” that assist consumers
with their contracts, privacy policies, and other legal
documents.41 It found that smart readers reduce the length of
contracts by 66.9 percent; reduce reading time by 14 minutes
and 41 seconds; improve text readability by reducing reading
levels from college-level to fifth-grade level; and, finally, do so
without compromising the essential information in the original
documents.42 There were failures, but at least some are
attributable to the length of the documents, which the LLMs
examined could only read in parts (this problem has since been
mostly mitigated).43
A different study evaluated the performance of LLMs on tax
code questions.44 These questions involve logical complexity
(e.g., exploring taxation of vested reversible, transferable
shares, and cost basis following a sale of inherited property) but
also tend to have a fairly crisp, unique answer. They find that
GPT-4 achieves around 77 percent accuracy on questions related
to the Code of Federal Regulation (C.F.R.) (with as much as
100 percent on basic tax problems), and 53 percent on general
United States Code questions.45 Critically, for the interpretation
of these numbers, the questions involve four to ten possible
40. Presenting Claude and Gemini with a contract and some context and asking
it them to identify the legal issues, CLAUDEAI, https://claude.ai/chat/77338278-0036-
469c-8d22-615c331f8c58 [https://perma.cc/7VTX-9FG4]; GEMINI, https://
gemini.google.com/app/560bd35270464077 [https://perma.cc/PL6Q-Y579].
41. See Yonathan A. Arbel & Samuel Becher, How Smart are Smart Readers?
LLMs and the Future of the No-Reading Problem, in THE CAMBRIDGE HANDBOOK
ON EMERGING ISSUES AT THE INTERSECTION OF COM. LAW AND TECH. (Nancy Kim
& Stacey-Ann Elvy eds., 2024) [hereinafter How Smart are Smart Readers]; Arbel
& Becher, supra note 34, at 94–106; see also Noam Kolt, Predicting Consumer
Contracts, 37 BERKELEY TECH. L.J. 71 (2022).
42. How Smart are Smart Readers, supra note 41, at 1.
43. Id. at 10−11; see also Kolt, supra note 41, at 109–117.
44. See John J. Nay et al., Large Language Models as Tax Attorneys: A Case
Study in Legal Capabilities Emergence, 382 PHIL. TRANS. R. SOC’Y A,
October 4, 2023, https://doi.org/10.1098/rsta.2023.0159 [https://perma.cc/HGZ4-
CRHG]. Importantly, the design employs retrieval-augmented generation and
prompt-engineering techniques. Id.
45. I focus here on the few-shot experiment. The relative weakness on the U.S.
Code is probably associated with the weakness of the retrieval augment generation
method, which is degraded on large corpora of text. For the data taken directly from
the data files, see John Nay, LLM Tax Attorney, GITHUB, https://github.com
/JohnNay/llm-tax-attorney/tree/main/data [https://perma.cc/4GTQ-NXET].

<<PAGE_BREAK>>

560 UNIVERSITY OF COLORADO LAW REVIEW [Vol. 96
answers, so chance accuracy would only be between 10 and
25 percent.46
These results are consistent with the other ones just
discussed in that they show a high but inconsistent level of
performance. Unfortunately, this study did not include a human
benchmark, so we cannot tell how much better or worse these
numbers are relative to a professional. However, given that the
questions rely on legal and financial fluency, it is safe to assume
that they considerably exceed the accuracy levels of the average
lay tax preparer, and possibly even of the average non-tax
lawyer. This highlights the margin of substitution point: LLMs
will replace not your white shoe lawyer, but your neighborhood
H&R Block representative or estate planner.
A persistent failure mode in these studies is
“hallucinations”—the invocation of non-existent facts, such as
precedents, and their presentation as facts.47 One study found
that “legal hallucinations are alarmingly prevalent” in LLMs,
occurring 58 percent (ChatGPT using GPT-3.5) to 88 percent
(Meta’s Llama 2) of the time when asked specific questions about
federal court cases.48 Two factors ameliorate this concern,
however. False sources, while a severe problem, can often be
checked with relatively little work, often involving a short
Internet search for verification. Moreover, while our current
understanding suggests that some degree of model inaccuracy is
inevitable, advances in modeling have shown promise in
reducing this problem significantly.49
Assessed more holistically, two recent papers tried to
determine whether models can act as generalist lawyers by
comparing the performance of humans to models on the bar
exam. A technical report by OpenAI famously reported that
46. Id.
47. See generally Jia-Yu Yao et al., LLM Lies: Hallucinations are not Bugs, but
Features as Adversarial Examples, ARXIV (Aug. 4, 2024), https://doi.org/10.48550
/arXiv.2310.01469 [https://perma.cc/M2ZB-M6YF] (demonstrating that nonsensical
prompts composed of random tokens can also elicit
the LLMs to respond with hallucinations).
48. Matthew Dahl et al., Large Legal Fictions: Profiling Legal Hallucinations
in Large Language Models, ARXIV 6 (June 21, 2024), https://arxiv.org/abs
/2401.01301 [https://perma.cc/Z2AX-39RD].
49. See Ziwei Xu et al., Hallucination is Inevitable: An Innate Limitation of
Large Language Models, ARXIV (Jan. 22, 2024), https://arxiv.org/abs/2401.11817
[https://perma.cc/QC9U-553B]. For mitigation techniques, see S.M. Towhidul Islam
Tonmoy et al., A Comprehensive Survey of Hallucination Mitigation Techniques in
Large Language Models, ARXIV (Jan. 8, 2024), https://arxiv.org/abs/2401.01313
[https://perma.cc/UM7G-JU6W].

<<PAGE_BREAK>>

2025] JUDICIAL ECONOMY IN THE AGE OF AI 561
GPT-4, at launch and without modifications, has passed the
Uniform Bar Exam at the 90th percentile.50 This puts GPT-4
above the median test-taker.51 Digging more deeply, Eric
Martinez argued that these results are confounded by the timing
of the specific comparison exam (February), which included
many repeat test-takers with lower scores.52 Applying several
corrections, he concludes that, when compared to exam passers
in the July administration, GPT-4 performance is estimated to
be at the median of test takers, and bottom 15th percentile on
the essay section.53 This aligns with an earlier study of GPT-3.5
showing that on law school exams GPT-3.5 performed at a C plus
level.54 But even with these more refined analyses, it is clear
that GPT-4 is already adequate at many tasks, even if adequacy
is a fairly low bar.
It is worth bearing in mind that we should be cautious about
extrapolating from bar performance and law school exams to
real-world performance. At the same time, we also cannot
completely discount their relevance given the critical
gatekeeping role bar exams play in our regulatory apparatus.55
Moreover, bar exams offer one of the sharpest ways to test
performance differentials between models and
highly-motivated, quasi-experts.
Finally, and most importantly, are the real-world studies of
AI effectiveness. These are early days, so caution is advised. One
study asked a trained lawyer and a GPT-4 model to draft a
complaint letter to the employer. Eighty percent of human
referees, in a blind test, preferred the model’s letter the trained
50. Daniel Martin Katz et al., GPT-4 Passes the Bar Exam, 382 PHIL. TRANS.
R. SOC’Y A 12 (2024), https://doi.org/10.1098/rsta.2023.0254 [https://perma.cc
/BHE2-DB68].
51. The median score in February 2023 was 131.5. The Multistate Bar
Examination (MBE), THE BAR EXAMINER, https://thebarexaminer.ncbex.org/2023-
statistics/the-multistate-bar-examination-mbe [https://perma.cc/3VU4-QZ5N].
52. Eric Martínez, Re-evaluating GPT-4’s Bar Exam Performance, in INST. L. &
A.I., https://.ssrn.com/abstract=4441311 [https://perma.cc/T3Y6-3VWM].
53. Id.
54. Jonathan H. Choi et al., ChatGPT Goes to Law School, 71 J. LEGAL
EDUC. 387, 391 (2022).
55. Kyle Rozema, Does the Bar Exam Protect the Public?, SSRN 2–3 (Aug. 22,
2021), https://papers.ssrn.com/sol3/papers.cfm?abstract_id=3612481 [https://
perma.cc/8S69-G87R] (showing that the “bar passage requirements have a modest,
negative effect on public sanctions.”).

<<PAGE_BREAK>>

562 UNIVERSITY OF COLORADO LAW REVIEW [Vol. 96
lawyer’s.56 Another study recruited legal aid lawyers, and gave
them access to GPT-4, with some of them getting access to other
AI tools.57 The lawyers reported a productivity increase,
although they remained worried about these tools. It is worth
noting that most participants appreciated GPT-4 but found the
other tools fairly unhelpful.58
To conclude, if we can provide an estimate of the general
capability of AI models in 2024, it will be in the spirit of
Martinez’s ultimate conclusions.59 Rigorous testing shows that
these systems are fast and cheap, but perform below the level of
median lawyers. This conclusion should be made alongside the
observation made at the outset—that is, what we see today are
tentative floors and ceilings. In fact, the tests discussed not only
do not account for future developments, but they also do not fully
take advantage of present developments, such as deep prompt
engineering, fine-tuning on specific datasets, or ensembling.60
But perhaps most deeply, the faults we find in LLMs should
always account for, and be measured against, the realistic
alternatives that ordinary people actually have. A clear lesson
from the work of Rebecca Sandefur is that socio-legal research
should consider the “importance of doing nothing.”61 As her
work shows, the most common responses to a problem are—in
order of frequency—some form of self-help, turning to a
third-party or a lawyer, and doing nothing.62 In fact, poor
households are twice as likely as middle-income households to
do nothing.63 We are not measuring AI tools in a vacuum; they
are responding to a social reality where the poor do nothing or
56. Lena Wrzesniowska, Can AI Make a Case? AI vs. Lawyer in the Dutch Legal
Context, INT’L J.L., ETHICS, & TECH., at 26 (Aug. 15, 2023), https://papers.ssrn.com
/sol3/papers.cfm?abstract_id=4614381 [https://perma.cc/6YK5-9LY6] (reporting an
experiment with 25 legal professionals who favored the models’ responses for
reasons of tone, clarity, style, argumentation, and evidence use).
57. See Colleen V. Chien & Miriam Kim, Generative AI and Legal Aid: Results
from a Field Study and 100 Use Cases to Bridge the Access to Justice Gap, 57 LOY.
L.A. L. REV. 903 (2025), https://digitalcommons.lmu.edu/cgi
/viewcontent.cgi?article=3210&context=llr [https://perma.cc/JJV2-9BAC].
58. Id.
59. See Martínez, supra note 52.
60. Pranab Sahoo et al., A Systematic Survey of Prompt Engineering in Large
Language Models: Techniques and Applications, ARXIV (Feb. 5, 2024), https://
arxiv.org/pdf/2402.07927 [https://perma.cc/R8BC-ZP3R].
61. Id.
62. Rebecca L. Sandefur, The Importance of Doing Nothing: Everyday Problems
and the Importance of Inaction, in TRANSFORMING LIVES: L. AND SOC. PROCESS 115,
115 (Pascoe Pleasence et al. eds., 2006).
63. Id.

<<PAGE_BREAK>>

2025] JUDICIAL ECONOMY IN THE AGE OF AI 563
rely on their own devices to resolve legal problems. This insight
deeply contextualizes the finding that LLMs are “only” as
effective as somewhat middling lawyers.
B. AI Uptake
How are people reacting to this new technology? The
potential seems quite large, with a Goldman Sachs report from
2023 claiming that AI will automate 44 percent of legal tasks
within ten years of broad adoption.64 Various reports show that
law firms are experimenting with AI tools in their practice.65 For
example, Allen & Overy deployed a model called Harvey and
quickly found that 25 percent of the firm’s practice used the tool
daily.66
Industry surveys provide a broader picture. A survey in
2023 found that 82 percent of lawyers believed that AI can be
applied to legal work, while also showing more hesitancy on the
appropriateness of doing so with only 51 percent answering in
the affirmative.67 An American Bar Association survey from
2023 reported usage among 11 percent of lawyers,68 a Lexis
survey reported 16 percent,69 and a survey of legal aid lawyers
found 21 percent usage.70
64. JAN HATZIUS ET AL., The Potential Large Effects of Artificial Intelligence on
Economic Growth, GLOB. ECON. ANALYST (Goldman Sachs Econ. Rsch., New York,
N.Y.), Mar. 26, 2023, at 6, https://www.gspublishing.com/content/research/en
/reports/2023/03/27/d64e052b-0f6e-45d7-967b-d7be35fabd16.html [https://
perma.cc/77WG-7KAV].
65. Frank Fagan, A View of How Language Models Will Transform Law, TENN.
L. REV. (forthcoming 2025) (manuscript at 26).
66. Bob Ambrogi, As Allen & Overy Deploys GPT-based Legal App Harvey
Firmwide, Founders Say Other Firms Will Soon Follow, LAWSITES.COM
(Feb. 17, 2023), https://www.lawnext.com/2023/02/as-allen-overy-deploys-gpt-
based-legal-app-harvey-firmwide-founders-say-other-firms-will-soon-follow.html
[https://perma.cc/9ZYM-DV5H].
67. New Report on ChatGPT & Generative AI in Law Firms Shows
Opportunities Abound, Even as Concerns Persist, THOMSON REUTERS (Apr. 17,
2023), https://www.thomsonreuters.com/en-us/posts/technology/chatgpt-
generative-ai-law-firms-2023 [https://perma.cc/AXK4-8HGJ].
68. Darla Wynon Kite-Jackson, 2023 Artificial Intelligence (AI) TechReport,
AM. BAR ASS’N (Jan. 15, 2024), https://www.americanbar.org/groups/law_practice
/resources/tech-report/2023/2023-artificial-intelligence-ai-techreport [https://
perma.cc/L9CW-S4GT].
69. LEXISNEXIS, INTERNATIONAL LEGAL GENERATIVE AI REPORT: DETAILED
SURVEY FINDINGS 6 (2023), https://www.lexisnexis.com/pdf/lexisplus/international-
legal-generative-ai-report.pdf [https://perma.cc/AG4X-H6ER].
70. Chien & Kim, supra note 57, at 20.

<<PAGE_BREAK>>

564 UNIVERSITY OF COLORADO LAW REVIEW [Vol. 96
While these surveys suggest only small-to-moderate
adoption, lawyers also see broad room for integration of AI tools
into their practice. Among the most common use cases, lawyers
reported creating drafts, brainstorming ideas, summarizing
complex documents, and assisting in writing emails.71 It is quite
reasonable to expect that, as AI tools develop specifically to meet
the needs of law firms, and as more lawyers graduate from law
schools after using AI tools, the levels of integration will
consistently increase. This is especially true given client
pressure to reduce billing through the integration of these
tools.72
Equally remarkable is the rate of change: slowly, then
suddenly. A recent survey on AI adoption in the workplace (not
specifically legal) has shown that 75 percent of knowledge
workers use AI at work.73 What is remarkable is that 46 percent
of workers started using AI tools less than six months ago
(i.e., late 2023).74 This spells a staggering rate of adoption. It is
highly unlikely that law firms will lag behind for much longer.
These findings speak to a number of issues. They show the
utility and competence of AI tools, at least when employed by a
legal professional. They show the broad range of tasks AI tools
can accomplish. They suggest a productivity gain in lawyering
which may or may not translate to lower cost or more
voluminous legal filings. They further suggest a possible
trickle-down effect, where the tools and techniques used by elite
lawyers will make their way to lawyers across the country and
maybe even be commercialized for retail use. And lastly, they
show a path towards integration by legal professionals in their
workflows—a path trodden by law firms but that could later be
replicated, mutatis mutandis, by judicial chambers and court
case management systems.
71. Caroline Hill, ILTA’s Blockbuster Technology Survey for 2023 Reveals All
on Collaboration Toos Adoption, Governance, and Plenty on Gen AI, LEGAL IT
INSIDER (Sept. 29, 2023), https://legaltechnology.com/2023/09/29/iltas-annual-tech-
survey-2023-reveals-all-on-collaboration-tools-adoption-governance-and-yes-lots-
on-gen-ai [https://perma.cc/8GAM-ET7L].
72. Logan Lathrop, Law Firms Leveraging AI: Maximizing Benefits and
Addressing Challenges, HARV. J.L. & TECH. DIG. (Nov. 20, 2023), https://
jolt.law.harvard.edu/digest/law-firms-leveraging-ai-maximizing-benefits-and-
addressing-challenges [https://perma.cc/VMJ7-XFSD].
73. AI at Work Is Here. Now Comes the Hard Part, MICROSOFT WORKLAB
(May 8, 2024), https://www.microsoft.com/en-us/worklab/work-trend-index/ai-at-
work-is-here-now-comes-the-hard-part [https://perma.cc/TF5Z-GDFY].
74. Id.

<<PAGE_BREAK>>

2025] JUDICIAL ECONOMY IN THE AGE OF AI 565
C. AI Impact on Access to Justice
Having seen the evidence of uptake of AI in the legal
industry, we now turn to examine AI’s broader impact on access
to justice. Before doing so, it should be recognized that “access to
justice” is a large umbrella term. It hides certain political
complexities about whose access matters,75 the extent to which
this justice is legal, and whether access is jeopardized by factors
that are formal, substantive, representative, or even
psychological.76 Still, at its core stands the basic proposition that
the halls of justice should be open to all and that barriers to
justice are regressive in nature, contributing to a regime where
the haves come out ahead of the have-nots.77
Evaluating the impact of AI on litigation patterns would
require some understanding of what these access barriers are.
People find difficulty accessing legal justice due to a large
number of barriers, some financial, others psychological,
political, and social, but many can be reduced, in some way or
another, to a cost-based explanation. What’s remarkable about
AI is that it produces a holistic shock to the access to justice
problem, one that includes the reduction in the cost of legal
services but goes beyond it to the social and psychological
barriers as well. Let us examine some of these effects in detail.
Legal sociology teaches the critical importance of upstream
filters. “[D]isputes are not things: they are social constructs.”78
For a mischief to be conceived as a legal dispute, it must undergo
at least three transformations given by the
naming-blaming-claiming (NBC) model.79 That is, the injured
75. See, e.g., Martha Minow, Access to Justice, 2 AM. J.L. & EQUAL. 293 (2022)
(focusing on “low-income Americans”); Bob Glaves, What Do We Mean When We Say
Access to Justice?, CHI. BAR FOUND., (July 11, 2023), https://
chicagobarfoundation.org/bobservations/what-do-we-mean-when-we-say-access-to-
justice [https://perma.cc/ZW9K-AM67] (focusing on “[a] person or entity facing a
legal issue . . .”). The United States Institute of Peace alternates between
“individual,” “people,” and “citizens.” Access to Justice, Guiding Principles for
Stabilization and Reconstruction: The Web Version, U.S. INST. OF PEACE
(Nov. 1, 2009), https://www.usip.org/guiding-principles-stabilization-and-
reconstruction-the-web-version/rule-law/access-justice [https://perma.cc/62S4-
S7ES].
76. For example, the United States Institute of Peace emphasizes that access
to justice is absent when people “fear” the system or see it as “alien.” Id.
77. Marc Galanter, Why the “Haves” Come Out Ahead: Speculations on the
Limits of Legal Change, 9 LAW & SOC’Y REV. 95 (1974).
78. Felstiner et al., supra note 27, at 631.
79. Id. at 633–36.

<<PAGE_BREAK>>

566 UNIVERSITY OF COLORADO LAW REVIEW [Vol. 96
party must perceive that they were injured; that a recognizable
actor injured them (rather than an act of Fortuna); and then be
able to conceptualize that accident in terms of a legal assertion
of rights against the violator.80 While data is scarce, sociologists
believe that these filters have a dramatic effect: “we know that
most of the attrition occurs at [the NBC] early stages.”81 An
important facet is distributional; the NBC filter asymmetrically
affects poor claimants, as the ability to name, blame, and claim
is predicated on access to educational, social, and plain, vanilla
capital.82 If the NBC filter is as powerful as sociologists claim,
and if it is as regressive in effect as commonly believed, its
removal would have broad implications for both substantive
rights and litigation patterns.
Generative AI takes the NBC filter head on. To illustrate
the way generative AI would work in practice, I presented a
simple query to a model: “[M]y landlord wants me to pay to fix
the mold in the basement and I don’t know what to do.”83 The
model responded with some fairly generic reminders that
landlords are responsible for the habitability of their residences,
that it is advisable to read the lease, and that it might be
appropriate to consult a legal professional. To a lawyer,
burdened with the curse of knowledge, this may not seem to be
very informative. But this response quickly and cheaply takes
the user through all three of the NBC stages.84
This example is humble, perhaps anecdotal, but I believe it
points at a deeper, hard to measure but nonetheless radical
change in the NBC model. Many people have had a moment
where the simple phrasing of their issues by a knowledgeable or
experienced acquaintance has helped put their issue in context
and motivated them to take an action that they would not have
taken otherwise. As AI systems become integrated into our daily
flow, as people come to consult them as often as they do Google
or other Internet sources, such framing effects can have large
impacts on the legal consciousness of ordinary people. Coupled
80. Id.
81. Id. at 636.
82. Id. at 637.
83. Landlord Mold Responsibility Query, CHATGPT (Aug. 31, 2024), https://
chatgpt.com/share/7dfbd694-4832-45c1-acce-471b94e4500f [https://perma.cc
/6QJR-GUMC].
84. Id. (“You should not be responsible for paying to fix mold in the basement,
as it is typically the landlord’s responsibility to ensure the property is habitable and
safe.”).

<<PAGE_BREAK>>

2025] JUDICIAL ECONOMY IN THE AGE OF AI 567
with their demonstrated (albeit imperfect) legal fluency, such
models could remove many invisible upstream barriers on the
way to justice.
Beyond the early stages, AI continues to contribute to every
aspect of the litigation journey. After reaching the claiming
stage, people will want to consider their legal strategy. Today,
people surveyed report that they seek lawyers for legal
information in only 29 percent of their cases, often depending on
the Internet and family or friends for orientation.85
In all those other cases, people can turn to AI systems to
help them with legal strategy, including matters such as
whether to send a demand letter, talk to a lawyer, write to a
government agency, and so on. When individuals turn to AI
tools, they can use them as powerful smart readers, tools that
not only summarize the information but also make it accessible
to one’s specific sociolinguistic needs.86
The next step in the journey for those who choose litigation
consists of producing written materials. The models can draft
the required communications, demand letters, complaints, and
other litigation materials. If they choose to file pro se,
individuals can use AI to produce responses to motions to
dismiss, help draft their pleadings, and generally help navigate
throughout the legal process. Even questions like “Where do I
send my documents?” that may be trivial to a lawyer, could
greatly benefit individuals in their journey. Notably, these
advantages help even for people who are represented. And while
they do not guarantee that they actually win their cases, they
give people more access to justice than they ever had before.
There is also considerable scope for more traditional
machine learning techniques in the litigation journey. In a
recent overview, Frankenreiter and Nyarko offer a broad
exploration of the utility of narrower predictive and
classification models.87 They provide persuasive use cases
related to automated review of documents to identify privileged
information using a model to predict case outcomes, and in turn
informing the selection of attorneys and venues.88 More
85. JUSTICE NEEDS, supra note 2, at 160 (showing legal aid services account for
additional 8 percent and court provided information for additional 7 percent).
86. See Arbel & Becher, supra note 34.
87. Jens Frankenreiter & Julian Nyarko, Natural Language Processing in
Legal Tech, in LEGAL TECH. AND THE FUTURE OF CIV. JUST. 70, 70 (David Freeman
Engstrom ed., 2023).
88. Id. at 74.

<<PAGE_BREAK>>

568 UNIVERSITY OF COLORADO LAW REVIEW [Vol. 96
generally, the extraction of legal data from troves of documents
presents a compelling and highly useful use case.89 As it comes
to barriers in access to justice, consider how such models can
help individuals conduct research, choose a court to file in, and
more generally, reduce some of the uncertainty of litigation,
which itself is a barrier to justice.
In considering the prospects of a litigation boom, we just saw
that AI can greatly reduce many access to justice barriers. If the
access to justice literature correctly mapped the barriers and
their size, we have a strong reason to expect an AI litigation
boom in the coming years. Exactly how large it would be is hard
to gauge with any accuracy, but if it is true that only 8 percent
of the legal needs of low-income people are addressed and that
seventy-five million cases every year receive no legal resolution,
then the potential is large indeed.90 Third-party financing
ameliorated the liquidity barrier that prevented litigants with
strong cases from filing them, and this had the effect of a
litigation spike.91 Moreover, it is not just the raw number of
cases that matters; AI systems are excellent providers of verbose
materials, making it effortless to write briefings and other
filings that are long-winded. All of this contributes to a large
potential AI litigation boom.
It is true that the quality of some of these filings may not be
high, but that’s hardly a reason to doubt their adoption and
impact. The economic incentives are simply too strong, and the
temptation of convenience too large. Even if the quality is not
quite there, convenience usually takes the upper hand.
To be sure, there are some trends that would work to
mitigate the litigation boom. It is possible that rates of
AI-generated filings will be lower, or high only among those
already prone to litigate their cases. It is also possible that the
higher risk of litigation would lead people to adapt their
behavior into greater compliance, or that would-be defendants
will settle at earlier stages. AI labs, by pressure of regulation or
89. Id. at 75.
90. See Sandefur, supra note 1; JUSTICE NEEDS, supra note 2, at 57; FY 2025
BUDGET REQUEST, supra note 3; Legal Markets, supra note 5, at 1785.
91. U.S. CHAMBER OF COM. INST. FOR LEGAL REFORM, THIRD PARTY
FINANCING: ETHICAL & LEGAL RAMIFICATIONS IN COLLECTIVE ACTIONS (Oct. 2020),
https://instituteforlegalreform.com/wp-content/uploads/2020/10/Third_Party
_Financing.pdf [https://perma.cc/CW26-3SDF] (Third-party financing is meant to
alleviate the liquidity constraints of litigants, and its effect is said to be to
“increase[] the volume of litigation in any jurisdiction where it is available.”).

<<PAGE_BREAK>>

2025] JUDICIAL ECONOMY IN THE AGE OF AI 569
exposure to unauthorized practice of law rulings, might also try
to prevent their models from producing effective materials. Such
possibilities exist, but it is unlikely that they will be able to
prevent the load on judicial resources that AI systems will have.
Some barriers to justice actually serve salutary purposes, as
counterintuitive as it may sound. If we admit that some filings
are vexatious, abusive, or meritless, then some filters may serve
important social goals in deterring them.92 To provide one
common example, consider debt collection litigation. Despite a
common view that these lawsuits are frequently abusive,
matters could actually be worse. Professional debt buyers who
buy large debt portfolios are effectively deterred by access
friction from filing claims for amounts below $500, and often
$1,000.93
We see, then, that AI has the potential to radically remove
filters and barriers on the way to justice. They help litigants at
every stage of the litigation journey, from forming the requisite
legal consciousness to creating legal strategies and then
implementing them. Many of the beneficiaries of these
improvements would be low-income individuals, currently
priced out of the market for legal services. But it is also
recognized that some strategic players, such as debt collection
firms, would come to use them to scale up their operations. Both
sides will contribute to a single likely outcome: an AI litigation
boom.
II. LEGAL THERMOSTATS
An AI litigation boom is the likely consequence of the
arguments this Essay just reviewed. Even if one takes a more
hedged view, it is clear that the forces that drive the supply of
litigation will grow significantly stronger in the presence of AI—
and that AI tools are continuously improving. A rapid increase
in case volume can have systemic repercussions on substantive
justice throughout the legal system. This is partly because
92. To be sure: the fact that barriers to the legal system serve a positive
function do not make them net positive. They also filter many truly important cases
and their effect is likely regressive. The point here is only that they also chill
low-quality cases.
93. Dave T., Debt Collection Agencies: What Is The Minimum Amount They
Would Sue For?, MAN VS. DEBT (Sept. 22, 2022), https://manvsdebt.com/debt-
collection-agencies-what-is-the-minimum-amount-they-would-sue-for [https://
perma.cc/25N4-5ZVF].

<<PAGE_BREAK>>

570 UNIVERSITY OF COLORADO LAW REVIEW [Vol. 96
justice delayed is justice denied, and partly because judges are
ultimately humans with only so many hours in a day.94 Bert
Huang demonstrated that a rise in administrative cases can lead
to “lightened scrutiny” of civil appeals.95 Not because judges
work any less hard—they likely work even harder—but because
there are physical constraints on what we can honestly expect of
even the most diligent public servant.
What will happen to judicial economy in the age of AI? How
can our current system—already burdened by its workload—
support a dramatic uptick in the number of cases? This Part lays
out the argument that past reactions to litigation surges have
been accompanied by adjustments that tended to affect primary
and procedural rights.
A useful way to think about judicial economy comes from
control theory.96 The core principle of control theory involves the
design and analysis of dynamic systems capable of maintaining
desired states despite internal and external disturbances. This
is achieved using control components, such as controllers,
sensors, and actuators, to endogenously regulate system
behavior towards an exogenously set desired state.
Consider the example of a thermostat. The thermostat is
programmed with a desired temperature (set point). It
continuously measures the actual temperature (process
variable) using temperature sensors (sensors) and compares it
to the setpoint. If the actual temperature deviates from the
setpoint, the thermostat activates the heating or cooling system
(actuators) to adjust the temperature back to the setpoint. This
feedback loop, where the system’s output influences future
inputs to maintain the desired state, is a hallmark of closed-loop
control systems. This contrasts with an open-loop system, such
as a simple fan, which operates without feedback and cannot
adjust to changing conditions.
94. Christoph Engel & Keren Weinshall, Manna from Heaven for Judges:
Judges’ Reaction to a Quasi-Random Reduction in Caseload, 17 J. EMPIRICAL
LEGAL STUD. 722, 722 (2020) (finding that “[j]udges working in courts with reduced
caseload invested more resources in resolving each case.”).
95. Bert I. Huang, Lightened Scrutiny, 124 HARV. L. REV. 1109 (2011); see also
Shay Lavie, Appellate Courts and Caseload Pressure, 27 STAN. L. & POL’Y. REV. 57
(2016).
96. For an introductory textbook, see KATSUHIKO OGATA, MODERN CONTROL
ENGINEERING (5th ed. 2010), https://wp.kntu.ac.ir/dfard/ebook/lc
/Katsuhiko%20Ogata-Modern%20Control%20Engineering-
Prentice%20Hal%20(2010).pdf [https://perma.cc/B62V-XV5P]. See also ROBERT H.
BISHOP & RICHARD C. DORF, MODERN CONTROL SYSTEMS (13th ed. 2022).

<<PAGE_BREAK>>

2025] JUDICIAL ECONOMY IN THE AGE OF AI 571
Judges, much like operators of a thermostat, play a critical
role in regulating the flow of litigation through their control over
procedural and substantive doctrines. These doctrines effectively
act as control mechanisms within the legal system,97 allowing
judges to adjust their strictness or leniency in response to the
demands of the judicial environment. Just as a thermostat
modulates temperature by activating heating or cooling
mechanisms, judges modulate the volume of cases by fine-tuning
these legal doctrines. This adjustment process is guided by
feedback from the legal system, such as fluctuations in case
volume or available judicial resources, and continues until the
flow of litigation aligns with the desired equilibrium or setpoint.
Critically, these judicial adjustments inevitably affect
substantive rights, raising concerns about the propriety of using
legal rights as levers for managing judicial resources.98 Despite
these concerns, it remains evident that such administrative
adjustments are a common practice employed by judges to
maintain judicial economy.
A few illustrations communicate the point.99 The most
salient is court fees. Courts in the United States charge a variety
of fees, including filing fees to initiate a case, fees for serving
documents, court reporter fees, jury fees, and fees for accessing
court records. Filing fees vary based on the type of case and
jurisdiction but can range from under $100 for small claims
cases to over $400 for civil cases in federal court.100 Court fees
97. In a contemporaneous article, Abramowicz considers the use of “automatic
stabilizers” to consider doctrinal changes in light of potential productivity changes
in lawyering due to AI. Michael Abramowicz, The Cost of Justice at the Dawn of AI
61−62 (Geo. Wash. Univ. Legal Stud., Research Paper No. 2024-37, Geo. Wash.
Univ. L., Public Law Research Paper No. 2024-37), https://ssrn.com
/abstract=4543803 [https://perma.cc/YJ4L-QMT4]. In various ways, his article
completes the analysis proposed here.
98. Compare Ronen Avraham & William H.J. Hubbard, Civil Procedure as the
Regulation of Externalities: Toward a New Theory of Civil Litigation, 89 U. CHI. L.
REV. 1 (2022), which emphasizes an externality control view of civil procedure, with
Marin K. Levy, Judging the Flood of Litigation, 80 U. CHI. L. REV. 1007, 1010−11
(2013).
99. While my focus here is on procedural mechanisms, substantive standards
also encode judgments on judicial resources, but this argument is beyond the
current scope.
100. For example, in Colorado where the 2024 Ira C. Rothgerber Jr. Conference:
AI and the Constitution took place, filing fees range from only thirty-one dollars to
nearly three-hundred dollars for small claims and civil cases in federal court. Court
filing fees vary from state to state. See, e.g., List of Fees, COLORADO JUDICIAL
BRANCH (Jan. 2025), https://www.coloradojudicial.gov/self-help/list-fees [https://
perma.cc/N4DF-R8VM].

<<PAGE_BREAK>>

572 UNIVERSITY OF COLORADO LAW REVIEW [Vol. 96
work well when they deter cases whose probability of winning is
so low that the potential payout falls below the fee. The de
minimis rule has a somewhat similar function because it filters
out cases with actual values on the premise that their social
value is also low. The problem is that fees and these types of
rules also screen out socially important and valuable
litigation,101 and the results tend to be quite regressive.102 We
know that even small access barriers can have large effects.
Something like the distance from the courthouse, which might
seem like a small concern, has a significant effect on the
participation rate of the poor—even for life-changing
litigation.103
Another prime illustration of thermostats comes from
pleading standards. Consider Twombly and Iqbal, two of the
most important procedural decisions in modern law.104 They
mark the move from a negative “no set of facts” standard to a
positive one requiring a showing of plausibility.105 This reflects
a heightening of pleading standards, and its direct implication
is chilling the filing of lawsuits. The motivation behind this
reform, in large part, was the growing costs of discovery that
were enabled by the old standard.106 Critics have argued that
such changes affect access to justice.107 The empirical evidence
shows that these decisions have had little impact on filing
activity by all but pro se plaintiffs.108 In other words, it is
101. Shmuel I. Becher et al., Toxic Promises, 63 B.C. L. REV. 753, 777 (2022).
102. Joseph Shapiro, As Court Fees Rise, The Poor Are Paying the Price, NPR
(May 19, 2014), https://www.npr.org/2014/05/19/312158516/as-court-fees-rise-the-
poor-are-paying-the-price [https://perma.cc/HK7K-XP8S].
103. David A. Hoffman & Anton Strezhnev, Longer Trips to Court Cause
Evictions, 120 PROC. NAT’L. ACAD. SCI. NO. 2 (2023), https://doi.org/10.1073
/pnas.2210467120 [https://perma.cc/27FU-ABD2].
104. Bell Atl. Corp. v. Twombly, 550 U.S. 544 (2007); Ashcroft v. Iqbal, 556 U.S.
662 (2009).
105. Edward D. Cavanagh, Making Sense of Twombly, 63 S.C. L. REV. 97, 98
(2011).
106. Twombly, 550 U.S. at 559 (“[I]t is only by taking care to require allegations
that reach the level suggesting conspiracy that we can hope to avoid the potentially
enormous expense of discovery . . . .”); see also Asahi Glass Co. v. Pentech Pharms.,
Inc., 289 F. Supp. 2d 986, 995 (N.D. Ill. 2003) (Posner, J., sitting by designation)
(“[S]ome threshold of plausibility must be crossed at the outset before a patent
antitrust case should be permitted to go into its inevitably costly and protracted
discovery phase.”).
107. Matthew A. Shapiro, Distributing Civil Justice, 109 GEO. L.J. 1473, 1516
(2021) (“[H]eightened pleading requirements and limits on discovery, have been
widely criticized for restricting access to justice . . . .”).
108. William H. J. Hubbard, The Effects of Twombly and Iqbal, 14 J. EMPIRICAL
LEGAL STUD. 474, 474−513 (2017).

<<PAGE_BREAK>>

2025] JUDICIAL ECONOMY IN THE AGE OF AI 573
unrepresented individuals who are bearing the brunt of the
heightened pleading standard and face more dismissals.
Most procedural thermostats are more indirect. Lone Pine
orders are an example.109 These are orders set out in large toxic
tort cases that call plaintiffs to present preliminary evidence on
questions of injury and causation within a deadline or risk
dismissal.110 These orders are clearly meant as a mechanism “to
identify and cull potentially meritless claims.”111 Critics have
decried their inconsistency,112 expressed concern that they turn
into “pseudo-summary judgment motions,”113 and overall worry
that they create a burden that is “unrealistic” and are an
“exercise [that] is onerous and unrewarding.”114 Nonetheless,
courts find them necessary to manage litigation.115
Consider next as a procedural thermostat the doctrine of
exhaustion of administrative remedies in the context of
prisoner’s rights.116 This broadly applied doctrine requires
plaintiffs to navigate agency processes to completion before
seeking judicial relief. While this doctrine abides by various
logics, litigation control is one of them. As a response to the spike
in inmate filings of the early 1990s,117 Congress enacted The
109. See generally Nora Freeman Engstrom, The Lessons of Lone Pine, 129 YALE
L.J. 2 (2019).
110. See, e.g., Claar v. Burlington N.R.R. Co., 29 F.3d 499, 500 (9th Cir. 1994)
(“The district court issued a case management order consolidating the twenty-seven
cases for pretrial purposes. The order required plaintiffs to submit affidavits
describing their exposure to the chemicals they claim harmed them, and affidavits
from physicians listing each plaintiff’s specific injuries, the particular chemical(s)
that in the physician’s opinion caused each injury, and the scientific basis for the
physician’s conclusions.”).
111. Baker v. Chevron USA, Inc., No. 1:05-CV-227, 2007 WL 315346, at *1 (S.D.
Ohio Jan. 30, 2007).
112. Engstrom, supra note 109, at 37.
113. Adinolfe v. United Tech. Corp., 768 F.3d 1161, 1168 (11th Cir. 2014).
114. Engstrom, supra note 109, at 52.
115. See, e.g., Acuna v. Brown & Root Inc., 200 F.3d 335, 340 (5th Cir. 2000) (“It
was within the court’s discretion to take steps to manage the complex and
potentially very burdensome discovery that the cases would require.”).
116. Kaiser Found. Hosps. v. Superior Ct., 128 Cal. App. 4th 85, 99−100 (2005);
Woodford v. Ngo, 548 U.S. 81, 88, 93 (2006) (“[T]he doctrine of exhaustion of
administrative remedies requires that where a remedy before an administrative
agency is provided by statute, regulation, or ordinance, relief must be sought by
exhausting this remedy before the courts will act.”); see also, Pozo v. McCaughtry,
286 F.3d 1022, 1025 (7th Cir. 2002) (“To exhaust remedies, a prisoner must file
complaints and appeals in the place, and at the time, the prison administrative
rules require.”).
117. Margo Schlanger, Inmate Litigation, 116 HARV. L. REV. 1555, 1578−87
(2003) (on the reasons for the spike). Russell Gold highlights that these filters tend

<<PAGE_BREAK>>

574 UNIVERSITY OF COLORADO LAW REVIEW [Vol. 96
Prison Litigation Reform Act.118 Senator Orrin Hatch, Chair of
the Senate Judiciary Committee, explained: “This landmark
legislation will help bring relief to a civil justice system
overburdened by frivolous prisoner lawsuits.”119 The Supreme
Court likewise noted in McCarthy v. Madigan that exhaustion
“serves the twin purposes of protecting administrative agency
authority and promoting judicial efficiency.”120
Empirical evidence suggests that the exhaustion
requirement does indeed filter out a significant number of
potential claims. In a study of discrimination cases filed to the
EEOC, Professor Bullock finds that only 16 percent of claims are
eventually filed in a federal court.121 Bullock’s study relies on a
nature of suit designation by the administrative office of the
court. A different estimate can be reached by analyzing the
actual text of filed cases. Data collected by Lex Machina shows
that from 2009 to the middle of 2017 there were 17,270 lawsuits
filed for employment discrimination.122 During the same time
period, the EEOC reports the total number of
discrimination-related charges (excluding retaliation) to be
474,220.123 Of these, 73.66 percent were dismissed or closed
with a finding of no reasonable cause, unsuccessful conciliation,
or administrative closure. This translates to roughly 349,310
unresolved cases. Conceding that combining datasets involves a
great degree of nuance that is missing here, the ratio of
unresolved discrimination claims to the EEOC that transform
into actual lawsuits is 3.6 percent.
Standards of proof also operate as procedural thermostats.
Consider what is necessary to prove to win a retaliation claim
to track claims by marginalized individuals. Russell M. Gold, Power over Procedure,
73 ALA. L. REV. 1, 105–06 (2022).
118. Prison Litigation Reform Act of 1995, Pub. L. No. 104-134, §§ 802−809,
110 Stat. 1321 (1995).
119. 141 CONG. REC. S26553 (daily ed. Sept. 27, 1995) (statement of Sen. Orrin
Hatch).
120. McCarthy v. Madigan, 503 U.S. 140, 143 (1992).
121. Blair Druhan Bullock, Frivolous Floodgate Fears, 98 IND. L.J. 1135, 1160
(2023).
122. Karl Harris, Lex Machina Launches Legal Analytics for Employment
Litigation, LEX MACHINA (July 12, 2017), https://lexmachina.com/blog/lex-machina-
launches-legal-analytics-for-employment [https://perma.cc/GSK6-4GBA].
123. For more on this data, see EEOC Data Collection, EEOC (2023), https://
www.eeocdata.org [https://perma.cc/U2YT-VHB8]. For code and analysis, see
Yonathan Arbel, Judicial Economy in the Age of AI, GITHUB (2024), https://
github.com/yonathanarbel/Judicial-Economy-in-the-Age-of-AI [https://perma.cc
/8FKR-9YJT].

<<PAGE_BREAK>>

2025] JUDICIAL ECONOMY IN THE AGE OF AI 575
under Title VII of the Civil Rights Act.124 Spurred by concerns
about a deluge of lawsuits, the U.S. Supreme Court decided that
the standard of proof would be the but-for test, rather than the
more plaintiff-friendly motivating factor test.125 It argued that
“[l]essening the causation standard could also contribute to the
filing of frivolous claims, which would siphon resources from
efforts by employer[s], administrative agencies, and courts.”126
A final illustration of procedural thermostats comes from
statutes of limitations. There are, by one count, around seven
categories of rationales for these laws.127 One of them is to
protect the integrity of evidence, which aims to “prevent[]
surprises through the revival of claims that have been allowed
to slumber until evidence has been lost, memories have faded,
and witnesses have disappeared.”128 But Congress sometimes
uses statutes of limitations as a means of controlling the volume
and quality of litigation,129 and so do some courts.130
The common usage of these procedural thermostats reveals
something general about the use of regulatory frictions in the
age of AI. Most of these thermostats work by adding friction to
the process. The (reasonable) expectation is that adding friction
would deter some people from filing, and the (often unverified)
hope is that those unfiled cases are those with lesser merit.131
The problem is that some of these frictions are quite vulnerable
to the introduction of AI tools. The reasons why people fail to
meet statutes of limitations requirements are varied, but some
124. 42 U.S.C. §§ 2000(e)(1)–(17).
125. Id.
126. Univ. of Tex. Sw. Med. Ctr. v. Nassar, 570 U.S. 338, 358 (2013). For a
critique, see Daiquiri J. Steele, Rationing Retaliation Claims, 13 U.C. IRVINE L.
REV. 993, 1003 (2023) (“While courts should be good stewards of judicial resources,
docket reduction should not take precedence over ensuring equal justice under the
law.”); see also Sandra F. Sperino & Suja A. Thomas, Fakers and Floodgates,
10 STAN. J.C.R. & C.L. 223, 229 (2014).
127. See generally Tyler T. Ochoa & Andrew Wistrich, The Puzzling Purposes of
Statutes of Limitation, 28 PAC. L.J. 453, 460–99 (1997).
128. Ord. of R.R. Tels. v. Ry. Express Agency, 321 U.S. 342, 349 (1944).
129. See, e.g., Sperino & Thomas, supra note 126, at 229 (arguing that “Congress
inserted numerous procedural and substantive provisions in Title VII that limit the
number of claims” which includes the short time to claim).
130. Ochoa & Wistrich, supra note 127, at 495–99.
131. Is it the case that a discrimination lawsuit filed after 320 days is less
meritorious than one filed within 290 days from the offending act? Compare,
however, the logic expressed in cases such as Chase Security Corp. v. Donaldson,
325 U.S. 304, 314 (1945), where the court sees statutes of limitation as tools that
“are by definition arbitrary, and their operation does not discriminate between the
just and the unjust claim, or the voidable and unavoidable delay.”

<<PAGE_BREAK>>

576 UNIVERSITY OF COLORADO LAW REVIEW [Vol. 96
of them depend on access to lawyering and litigation
financing.132 AI can ameliorate such barriers because it can
shepherd people and help them process the wrong they suffered
through the NBC process and then assist them in constructing
legal documents. Similarly, AI tools can significantly reduce the
costs, hurdles, and frictions associated with exhausting
administrative remedies. AI-powered tools could quickly
identify relevant agencies, help navigate their process, and draft
complaints. Finally, the same tools also apply to pleading
standards. Plausibility standards do not only filter cases that
are implausible. They also filter cases where people were
negligent or unskilled in framing their arguments or lacked the
requisite polish, which is one reason why the effect is seen
among pro se litigants.133 These filtering functions of pleading
standards are fragile to AI tools that can mass produce elaborate
briefs for even the most tenuous of cases. What adjustments
await when the old methods of adjusting the thermostat stop
working?
III. LEGAL STRATEGIES THAT DEAL WITH THE AI LITIGATION
BOOM
If the diagnosis by access to justice advocates is correct, the
prognosis is clear. To the extent AI tools remove frictions and
costs in access to justice, we should expect a commensurate
increase in civil litigation. And because the size of the access to
justice gap is so large, a doubling in the volume of litigation is
not implausible.134 Moreover, litigation would also adjust on
132. For a psychological account of delay, see Andrew J. Wistrich,
Procrastination, Deadlines, and Statutes of Limitation, 50 WM. & MARY L. REV. 607
(2008).
133. Hubbard, supra note 108, at 512 (2017) (explaining the “differential effect
for pro se plaintiffs” as “unsophisticated parties may have a poor sense of whether
their facts entitle them to relief, and thus more pro se complaints may be marginal
under a plausibility pleading standard.”).
134. Ideally, when scholars make prescriptions based on their understanding of
the future trajectory of the world—as I do here—they should offer some concrete,
refutable predictions on how they perceive future trends to evolve. Here, it’s
important to acknowledge problems of missing data on present litigation patterns,
scope and type of barriers, levels of unmet needs, and so on. Still, if it turns out in
five to eight years that there was no discernible and practically meaningful AI effect
on litigation patterns, the reader should consider this Essay’s central claim
disproven. See also Yonathan Arbel (@ProfArbel), X (Aug. 22, 2020, 6:17 PM),
https://twitter.com/ProfArbel/status/1297327039670898688 [https://perma.cc
/S3MY-MGBD].

<<PAGE_BREAK>>

2025] JUDICIAL ECONOMY IN THE AGE OF AI 577
other dimensions, with verbosity of filings being one expected
effect. With more filings that are longer and more intricate, the
expected net effect is easily summarized: a litigation boom.
Historically, courts have reacted to threats to judicial
economy by adjusting the thermostat through pulling and
pushing on the levers available to them. The goal of this Part is
to situate thermostat adjustment as one of several possible
strategic reactions to the expected AI litigation boom. It
concludes with a discussion of the policy I consider most prudent:
proactive integration. AI has shortcomings and reliability
issues, but, as explained, some are exaggerated and others
manageable, and all should be evaluated vis-à-vis the other
realistic alternatives we have on the menu. By using whatever
time we have left until the AI litigation boom, we can carefully
build, test, and deploy AI tools as part of the judicial process.
A. Strategy 1: Legal Thermostats: Fees, Pleading
Standards, and Substantive Standards
The first strategy available to courts is the one that repeats
the historical pattern: adjustment of the legal thermostat by
adapting various doctrines that double as litigation control
levers.135 Judges and judicial administrators may feel it is
necessary for them to require even higher fees to offset the
demand for legal resources, to demand even more elaborate
pleading standards, or perhaps go as far as narrowing
substantive rights. These levers can decrease litigation
levels,136 but they also make it harder to vindicate legitimate
claims. As every lawyer knows, being right and being able to
prove one’s case are not the same.
Fees are a crude lever. To meet a litigation surge, judicial
administrators can increase filing fees, increase bond
requirements, and modify other requirements. Pulling on this
lever is almost guaranteed to chill filings and reduce lawsuits.
But the downside is obvious: Requiring higher fees will narrow
access to those who cannot afford them, not just those who file a
low-quality lawsuit. A plausible rejoinder is that if a plaintiff is
very likely to win then they should be able to borrow against
their future winnings and thus still access the gates of justice.
135. See supra Part II.
136. Note, however, that they also invite more accidents, and the net effect on
litigation levels depends on a broader set of variables.

<<PAGE_BREAK>>

578 UNIVERSITY OF COLORADO LAW REVIEW [Vol. 96
The rise of the litigation financing industry would be evidence in
favor of this rejoinder. Yet this rejoinder is facile. Not only is
access to capital a challenge for many low-income individuals,
the risk of losing a meritorious claim is especially threatening if
one has loans to repay.137 In between those liquidity constraints
and the “chance of ruin,” fees are a very crude tool of filtering
lawsuits and have disproportionate impact on the poor.
Pleading standards may seem like a lighter touch
intervention.138 Conceptually they can be thought of as a
“proof-of-work” mechanism. Proof of work is familiar from
blockchain technology, where it is used to validate claims made
by certain network participants.139 In order to be a trusted
validator of blockchain transaction, a blockchain miner has to
show that it had solved a complex math assignment. The
proof-of-work mechanism adds friction to the process of
validating transactions but is a necessary component of the
network as it is effective in filtering out fraudsters. But despite
their common association with blockchain, such mechanisms are
far more general and common than many realize. In the current
context, the litigation process can be thought of as having a front
end (initial claim processing) and a back end (trial). Litigants,
presumably, have a sense of the merits of their case. The
proof-of-work mechanism leverages it to set higher front-end
requirements. A person who puts in the drafting work and sinks
in the necessary cost to meet plausibility standards in the front
end likely has a higher estimate of their case than a person who
would be discouraged by such costs. This is because the back end
costs are only borne by people who would pursue the case to its
completion. Thus, we can see the Twombly-Iqbal logic as
enforcing a proof-of-work mechanism: requiring more work on
137. See generally Yonathan A. Arbel, Payday, 98 WASH. U. L. REV. 1 (2020).
138. One adjustment, wisely pointed out by the editors of the University of
Colorado Law Review, is word limits. There is a complex menu of word limits and
word regulation for the production of legal materials. See, e.g., FED. R. APP. P.
APPENDIX: LENGTH LIMITS STATED IN THE FEDERAL RULES OF APPELLATE
PROCEDURE, https://www.ca6.uscourts.gov/sites/ca6/files/documents/rules
_procedures/Appendix.pdf [https://perma.cc/DAQ2-DE52]. That word limits are
crude tools of managing judicial economy goes without saying: it takes long to write
short, to paraphrase Pascal. Blaise Pascal, Provincial Letters: Letter XVI, to the
Reverend Fathers, the Jesuits, CHRISTIAN CLASSICS ETHEREAL LIBRARY, https://
ccel.org/ccel/pascal/provincial/provincial.xviii.html [https://perma.cc/RQP2-HAPK].
Not all can afford to do so, and this tool is not AI-proof as AI systems are excellent
summarizers.
139. For an introduction, see Michael Abramowicz, Cryptocurrency-Based Law,
58 ARIZ. L. REV. 365, 379 (2016).

<<PAGE_BREAK>>

2025] JUDICIAL ECONOMY IN THE AGE OF AI 579
the front end but serving the litigants later, thus acting as an
effective proof-of-work filter.
Assuming for a moment that this assumption is correct in
general, AI tools present a particular problem. Normally, the
crafting of effective pleadings requires an effective counsel and
an investment of time. A judge can relatively quickly discern
plausibility when the case involves low-effort filings. But AI
models are incredible writing assistants;140 they can rapidly and
easily convert vague claims to elaborate legal arguments, using
perfect grammar and compelling structure. This does not make
the claims any more valid, but it does make the production
cheaper and later validation harder. Recall that Twombly-Iqbal
mainly affects pro se litigants, and so they have the greatest
opportunity to benefit from such a tool.141 Ironically,
hallucinations can contribute to the facial plausibility of the
filings, even when the underlying claim lacks any support.
Consider, as illustration of hallucination, a request that the
AI produce a claim for workplace discrimination. Commentators
note that plausibility requirements hamper many such
claims.142 The model, however, could simply generate a set of
(semi-fictitious) facts and legal arguments that, while not true,
will seem true on their face. If the litigant is not careful and
scrupulous enough in reviewing them, it could pass initial
muster. As a result, filtering mechanisms that rely on proof of
work will become less effective than before. This could result in
escalation of front-end investments until the point where AI
cannot provide sufficient utility.
Finally, judges can simply demand more doctrinally for
filings. They can recharacterize strict liability as negligence or,
more subtly, change the meaning of reasonable person to meet a
desired level of stringency. Such changes can be hard to notice
in real time and even harder to causally relate to any thermostat
adjustment. Yet, they serve as a way to conserve judicial
resources and are available to decision-makers who feel strained
by the volume of litigation.
140. See generally Lu Sun et al., MetaWriter: Exploring the Potential and Perils
of AI Writing Support in Scientific Peer Review, 8 PROCS. OF THE ASS’N FOR
COMPUTING MACHIN. ON HUM.-COMPUT. INTERACTION, no. CSCW1-94, Apr. 2024,
at 1, https://doi.org/10.1145/3637371 [https://perma.cc/Q829-GKUT].
141. See supra Part II.
142. See, e.g., Joseph A. Seiner, Plausible Harassment, 54 U.C. DAVIS L. REV.
1295, 1310 (2021).

<<PAGE_BREAK>>

580 UNIVERSITY OF COLORADO LAW REVIEW [Vol. 96
Whatever form these adjustments take, the worrisome
implications are the narrowing of civil rights and, functionally,
a large subsidy to wrongdoers who could get away with more
socially pernicious activity. Less obvious is the problem that
these mechanisms are not very AI-proof, so their effects will be
unstable and will require constant adjustments.
To finalize our accounting, the net effect of increased access
to justice could be worse delivery of justice. Litigants who can,
for the first time, afford to enter the halls of justice, will be
denied justice within it. Higher fees, pleading standards, or
ever more demanding substantive changes can undo all the
access to justice gains AI will bring to underserved litigants.
Worse, some of the thermostats will be ineffective or will need to
be adjusted further and further because AI can circumvent
conventional proof-of-work mechanisms. While thermostat
adjustment is the most likely, perhaps even inevitable,
trajectory, I believe it will be undesirable to rely on it.
B. Strategy 2: Sit and Wait
Sometimes it is easiest to cross the bridge when you get
there, and perhaps policymakers will want to wait a while longer
before taking concrete action. Judges and judicial
administrators are careful by nature, and a rapidly expanding
and advertised technology such as AI raises understandable
concerns about unjustified hype and empty promises.
Technological uncertainty remains a significant hurdle for any
planner. While it is evident that AI is transforming the
production of legal materials, the full extent of this shift and its
implications—particularly the potential for a litigation boom—
are not yet fully understood. Historical precedents with earlier
waves of legal technologies, such as LexisNexis and LegalZoom,
suggest that whatever changes these technologies brought, the
legal system was able to adapt without catastrophic disruptions.
Moreover, given the current imperfections in AI technologies,
prudence might dictate a period of observation and gradual
adaptation. Thus, judges and judicial administrators may wish
to wait before they make any adaptations to legal processes,
procedures, and doctrines.
Further complicating the decision is the pattern of AI
adoption. We do not know yet who the dominant users would be:
pro se litigants? white shoe law firms? non-practicing patent
entities? automated litigation agents? The answers may affect

<<PAGE_BREAK>>

2025] JUDICIAL ECONOMY IN THE AGE OF AI 581
our normative evaluation of the technology. Should AI tools
follow the trajectory of previous legal tech innovations, we might
witness a surge in litigation activities by firms and commercial
entities rather than underserved individuals.143 There is also
the potential for negative uses, such as harassment or
unmeritorious litigation initiated by individual plaintiffs, which
could distort the justice system and detract from its core
functions.
Despite these considerations, I argue against a passive
stance. Current trends, though based on preliminary data,
indicate a clear trajectory toward increased AI integration
within legal practices.144 The unreliability of AI, rather than a
deterrent, should be a catalyst for judicious development and
testing. This proactive approach would not only allow for
refinement of the technology but also prepare the judicial system
to harness AI’s benefits effectively.
Moreover, even assuming the legal system could absorb the
impact of AI without significant structural changes, proactive
adaptation could still soften the shock of the transition and
enhance its efficiency and effectiveness. Innovations such as
video conferencing and digital legal research have already
demonstrated the benefits of integrating technology in legal
processes even when there was no imminent threat to the
volume of litigation.145
In conclusion, while the allure of a cautious approach is
understandable given the unknowns associated with AI, there
are strong reasons to adopt a more proactive engagement. This
strategy ensures that the judicial system is not merely reactive
but remains at the forefront of technological integration,
enhancing its capacity to deliver justice effectively.
143. See Engstrom & Engstrom, supra note 11.
144. See supra Section III.E.
145. Victor D. Quintanilla et al., Accessing Justice with Zoom: Experiences and
Outcomes in Online Civil Courts, MAURER SCH. OF L., at 2 (2023), https://
www.repository.law.indiana.edu/cgi/viewcontent.cgi?article=4087&context=facpub
[https://perma.cc/S9U8-5UKF] (finding evidence that a non-represented plaintiff
expressed preference for remote hearings, and other evidence of procedural and
distributional justice). There are also problems that are associated with remote
justice. See, e.g., Alicia Bannon & Janna Adelstein, The Impact of Video Proceedings
on Fairness and Access to Justice in Court, BRENNAN CTR. FOR JUST. (Sept. 10,
2020), https://www.brennancenter.org/our-work/research-reports/impact-video-
proceedings-fairness-and-access-justice-court [https://perma.cc/A848-DZEN].

<<PAGE_BREAK>>

582 UNIVERSITY OF COLORADO LAW REVIEW [Vol. 96
C. Strategy 3: Ban and Mark
There is a growing sentiment, mostly expressed to me in
private conversations with judges, that generative AI should be
banned in the courtroom. Alternatively, some favor a
requirement that lawyers disclose when they are using
AI-generated materials.146
The judicial skepticism is understandable, but I believe it is
wrong to follow it in the long term. A ban would kill our ability
to democratize access to the justice system in the crib.147 It
would perpetuate the asymmetries that currently exist, working
disproportionality against those who have the most to benefit
from the technology.
Disclosure regimes are a hopeless enterprise. As far as we
know, and to the displeasure of school administrators
everywhere, there is no reliable technology that can watermark
AI-produced texts. Detection of AI-generated texts is
probabilistic and error-prone, and it may—at best—only cover
the least sophisticated of its users.148 The share of those hapless
users is small, and their culpability is no worse than their more
sophisticated peers. But most importantly, the expected level of
AI integration in law practices suggests that disclosure will be
as meaningful as requiring litigants to disclose if they used a
search engine or a computer. It will communicate no actionable
information to the judge and will become as helpful as “here
comes the plaintiff” and other legal boilerplate. Overall, I would
caution those judges and judicial administrators who, in good
faith, worry about rising rates of litigation against trying a
hopeless “ban-and-mark” regime.
146. Maura R. Grossman, Paul W. Grimm & Daniel G. Brown, Is Disclosure and
Certification of the Use of Generative AI Really Necessary?, 107 JUDICATURE 69, 70
(2023), https://judicature.duke.edu/articles/is-disclosure-and-certification-of-the-
use-of-generative-ai-really-necessary [https://perma.cc/4ZYP-WB7S] (Judge
Michael M. Baylson was in favor, issuing standing orders requiring lawyers to
disclose use of AI).
147. On the democratizing arguments, see supra Section I.C.
148. See, e.g., Manshu Zhang et al., The Three-Dimensional Porous Mesh
Structure of Cu-Based Metal-Organic-Framework—Aramid Cellulose Separator
Enhances the Electrochemical Performance of Lithium Metal Anode Batteries, 46
SURFACES & INTERFACES 104081 (2024) (retracted), https://doi.org/10.1016
/j.surfin.2024.104081 [https://perma.cc/943F-DKML] (a retracted article which
opens its introduction with “Certainly, here is a possible introduction for your topic
. . .”). The original version is stored in Reddit, https://i.redd.it/zq0raef1aaoc1.jpeg
[https://perma.cc/G5AN-QZTN].

<<PAGE_BREAK>>

2025] JUDICIAL ECONOMY IN THE AGE OF AI 583
D. Strategy 4: Massive Funding
Justice costs money. If the problem of judicial economy is
that there is a growing demand for justice—as I have argued
throughout—then clearly the most direct way of solving the
problem is by increasing the resources available to the legal
system.
How many resources should go to justice, and at the expense
of what other social programs, is a political question that exceeds
my proffered expertise. What is meaningful for evaluating the
prospects of a budget increase, however, is the estimated size of
funding. If there is room for a two-fold or a five-fold increase in
the volume of litigation, then this gives a general sense of the
magnitude of the budget required to handle it. Of course, not
all—not even the majority—of this potential will translate into
lawsuits. Society adapts to technological change along many
dimensions, and there are many other ways to avoid legal
disputes. But the realism of a budget increase that would even
approximately double the number of judges and judicial
administrators appears quite tenuous in our current political
reality.
One fact that lends some realism to this proposition is that
civil legal aid benefits today from roughly $2.7 billion in overall
budgets.149 If one feels particularly bullish on AI technology and
its ability to replace legal aid through its automation, perhaps it
is conceivable that some of these budgets could be redirected
towards the legal system.150
Yet, even if AI is so potent as to completely substitute the
need for legal aid (a tenuous proposition, given that legal aid
does more than drafting briefs), there is not enough money there.
The federal court system alone is budgeted at $9.4 billion per
year, so even if were to somehow completely dismantle the legal
aid project, we could at most afford a 30 percent increase in
149. ALAN W. HOUSEMAN, INT’L LEGAL AID GRP., LEGAL AID IN THE UNITED
STATES: AN UPDATE FOR 2023 (May 2023), https://clp.law.harvard.edu/wp-content
/uploads/2023/05/USA-National-Report-ILAG-Conference-2023.pdf [https://
perma.cc/94B8-KAL7]. According to the Legal Services Corporation (LSC) data
from 2022, the total funding for LSC-funded organizations was $1.72 billion. BY
THE NUMBERS 2022, supra note 11, at 13–14.
150. Houseman, supra note 149, at 4 (noting that since 2000, LSC has funded
more than 859 projects totaling over $81 million in Technology Initiative Grants.).

<<PAGE_BREAK>>

584 UNIVERSITY OF COLORADO LAW REVIEW [Vol. 96
funding.151 But in a world where AI is sufficiently competent to
perform as well as legal aid, the rise in demand will be much
larger. At best, we would only scratch the surface of demands on
the legal system, while hollowing out legal aid.
E. Strategy 5: Integration
If none of the above strategies can effectively and equitably
meet the AI litigation boom, the legal system still has one other
important course of action available to it: integration. The
objective would be to implement AI in all aspects of the legal
process, amplifying the productivity of judges and clerks, which
would allow them to work at larger-than-ever scales. If done
correctly, this strategy would offer a significant stretching of
existing judicial resources, allowing judges to meet increased
demand without resorting to adjustment of legal thermostats or
sacrificing justice in individual cases.
Rather than a hypothesis, this seems to be organically
happening. Judges have started admitting to using generative
AI to draft opinions, although the backlash suggests that many
others are still in hiding.152 One British judge made the point
succinctly and forcefully: “It is useful, and it will be used.”153
Likewise, Richard Re believes that judges will invariably find AI
tools to be “irresistibly attractive.”154
Most remarkably, in a groundbreaking decision, Judge
Newsom of the Eleventh Circuit has written an opinion relying
on AI for “generative interpretation.” Drawing on our academic
work on generative interpretation, he said:
Those, like me, who believe that “ordinary meaning” is the
foundational rule for the evaluation of legal texts should
consider—consider—whether and how AI-powered large
151. ADMIN. OFF. OF THE U.S. CTS., THE JUDICIARY: FISCAL YEAR 2025
CONGRESSIONAL BUDGET SUMMARY, at i (Feb. 2024), https://www.uscourts.gov
/sites/default/files/fy_2025_congressional_budget_summary.pdf [https://perma.cc
/XS66-9ZSJ].
152. Hibaq Farah, Court of Appeals Judge Praises ‘Jolly Useful’ ChatGPT After
Asking It for Legal Summary, THE GUARDIAN (Sept. 15, 2023), https://
www.theguardian.com/technology/2023/sep/15/court-of-appeal-judge-praises-jolly-
useful-chatgpt-after-asking-it-for-legal-summary [https://perma.cc/33W8-EMTM].
153. Id.
154. Richard Re, Artificial Authorship and Judicial Opinions, 92 GEO. WASH. L.
REV. 1558, 1561 (2024), https://www.gwlr.org/wp-content/uploads/2024/12/92-Geo.-
Wash.-L.-Rev.-1558.pdf [https://perma.cc/UVS6-YSF5].

<<PAGE_BREAK>>

2025] JUDICIAL ECONOMY IN THE AGE OF AI 585
language models like OpenAI’s ChatGPT, Google’s Gemini,
and Anthropic’s Claude might—might—inform the
interpretive analysis.155
Appeal notwithstanding, there is also significant resistance
to integration, at least in its stronger forms. While scholars such
as Eugene Volokh express cautious optimism about the
automation of judgments—that is, “robo-judging”156—others are
less sanguine. Aziz Huq speaks of a right to a “human
decision,”157 and experiments suggest a perceived fairness gap
between human and artificial adjudicators.158 These objections
rely in part on empirical objections concerning the capacity of
these systems to produce judgments that are as good as a human
judge in terms of accuracy, bias, and gameability. They also
draw on sensible ethical concerns regarding the ethics of
adjudication by those who are neither citizens nor humans. The
former set of problems is amenable to practical solutions, while
the latter can be mostly remedied by including human judges
who are in the loop.159
When we talk about integration, I would like to suggest that
robo-judging should not be a central frame of thinking about the
technology. While it is provocative and exciting, for sure,
ultimately robo-judging is a distraction from the much more
mundane but nonetheless powerful utility of AI in the service of
justice. In the remainder of this Section, I want to highlight a
few of these modalities.
The immense volume of text generated in litigation is
staggering, and this will likely increase as parties begin
leveraging advanced AI tools to augment their legal processes.
To mete out justice, we need some way to compress all this
155. Snell v. United Specialty Ins. Co., 102 F.4th 1208, 1221 (11th Cir. 2024)
(Newsom, J., concurring) (citing Yonathan A. Arbel & David A.
Hoffman, Generative Interpretation, 99 N.Y.U. L. REV. 451 (2024), https://
www.nyulawreview.org/wp-content/uploads/2024/05/99-NYU-L-Rev-451-1.pdf
[https://perma.cc/3Y4S-LDH7]).
156. Eugene Volokh, Chief Justice Robots, 68 DUKE L.J. 1135 (2019).
157. Aziz Z. Huq, A Right to a Human Decision, 105 VA. L. REV. 611 (2020); see
also Kiel Brennan-Marquez & Stephen E. Henderson, Artificial Intelligence and
Role-Reversible Judgment, 109 J. CRIM. L & CRIMINOLOGY 137 (2019).
158. Benjamin Minhao Chen, Alexander Stremitzer & Kevin Tobia, Having Your
Day in Robot Court, 36 HARV. J. L. & TECH. 127 (2022).
159. Huq, supra note 157, at 4; see also Brennan-Marquez & Henderson, supra
note 157, at 149.

<<PAGE_BREAK>>

586 UNIVERSITY OF COLORADO LAW REVIEW [Vol. 96
information. In other words, we need a summarization machine,
and it turns out that generative AI excels at this task.160
Document summarization is among the most explored areas
within natural language processing using AI. This technology is
divided into two main types: abstractive and extractive
summarization. Abstractive summarization creates a new,
condensed version of the text that conveys the core meaning of
the text, potentially using its own words. Extractive
summarization, on the other hand, identifies and compiles key
phrases directly from the text.161 Both approaches can
significantly aid judges by highlighting essential information
and reducing the amount of material they need to personally
review.
An abstractive summary can direct a judge’s attention to
critical parts of a document, effectively serving as a
sophisticated, automated, and high-level summary of a
document. A file management system could mark a filed
document as “exhibit 182A,” the text “Sale agreement of the
Tuscaloosa house.” Unlike summaries written by any of the
litigants, the AI has no incentive to highlight a specific frame—
it seeks to offer a robust, useful summary to the best of its
ability.162
Extractive summaries, on the other hand, are invaluable for
identifying crucial elements within the text. An extractive
summary of the sale agreement may include elements such as
“seller shall deliver the property on or before January 1st.” It
could also include specific pieces of evidence, legal authorities,
or specific quotes. These summaries are particularly useful in
scenarios where precise language and specific details are
pivotal.
Both abstractive and extractive summaries have their uses.
To orient oneself in a stack of documents, abstractive summaries
160. See generally Text Summarization, PAPERS WITH CODE, https://
paperswithcode.com/task/text-summarization [https://perma.cc/AV3F-KPF3]
(presenting benchmarks on text-summarization tasks).
161. Nikolaos Giarelis, et al., Abstractive vs. Extractive Summarization: An
Experimental Review, 13 APPLIED SCI. 7620 (2023).
162. The sort of biases that afflict AI systems are often irrelevant to
summarization tasks. There are some implicit biases that can creep in nonetheless
(such as assumptions that a doctor is male), but clerks may well be subject to
similar biases and, in any event, the impact on any actual decision is highly
attenuated. What is perhaps most important is that the models have no stake in
the case at hand.

<<PAGE_BREAK>>

2025] JUDICIAL ECONOMY IN THE AGE OF AI 587
are essential; to locate leading phrases and arguments within a
document, extractive summarization is powerful.
The implementation of such summarization technologies in
case management systems is straightforward and cost-effective.
It is expected to be as simple as any large automation project is,
albeit, more costly and complicated than anticipated, but
ultimately solvable.163 It would be quite possible to integrate
these systems at the case management level, ensuring that
every submitted document includes an automated summary and
extraction of key parts. This allows effective attention
management on the part of the judge, a way to easily sort and
find the appendix dealing with the copy of the sale contract the
parties mentioned or the one document that covers Consumer
Price Index adjustments.
There is a more advanced application, commonly known as
“document Q&A.” Documents, by their nature, are static
entities. They contain information, and one has to read through
the document to extract it. This becomes unwieldy when dealing
with a lengthy document. Search engines offer a greater degree
of interactivity. They allow one to filter pieces of a document
based on keyword searches. Such keywords can be as simple as
searching for “choice of law,” or more advanced such as a search
for “executive* /w3 decision?” Once located, the system will
highlight the relevant text and orient attention to all the
relevant “hits.” The user is expected to sort through them and
find the relevant one.
Using document Q&A is the next step.164 It allows the judge
to ask specific questions about the document, and, rather than
using arcane keywords, the judge can use ordinary language.
That is, after the AI ingests a filing, the judge can simply ask:
“does this brief mention a meeting in Switzerland?”; “does the
plaintiff mention the statute of limitations?”; or “list the
executive decision the document mentions and what it means.”
The AI will then diligently provide an answer based on the
content of the document. The answer itself will be in natural
language, for example, “this document mentions a meeting in
163. Hofstadter’s Law states: “It always takes longer than you expect, even when
you take into account Hofstadter’s Law.” DOUGLAS R. HOFSTADTER, GÖDEL,
ESCHER, BACH: AN ETERNAL GOLDEN BRAID 152 (20th anniversary ed. 1999).
164. On the use of document Q&A for legal applications, see Xiaoxian Yang
et al., Large Language Models for Automated Q&A Involving Legal Documents: A
Survey on Algorithms, Frameworks and Applications, 20 INT’L J. WEB INFO.
SYS. 413 (2024).

<<PAGE_BREAK>>

588 UNIVERSITY OF COLORADO LAW REVIEW [Vol. 96
Zurich between the CEO of Acme and the CFO of Alpha,
although it doesn’t discuss its purpose.” Because the interface is
simply plain language, it requires little training to learn how to
use document Q&A.
Using document Q&A is a radical improvement over our
current means of interacting with documents. Search engines
direct users to not think about the question they want to answer
but rather on what queries will most likely produce the context
that will answer them. We search for “choice of law” not because
we necessarily care about the term, but because we think the
term will be in the context of the clause that determines the
choice the parties have made. Along the way, we trudge along
many irrelevant mentions of the term. Document Q&A allows
the user to skip this stage. The judge can simply ask “what is the
choice of law in this document?”
Document Q&A methods are not an all-knowing sage, of
course. It is perhaps most productive to think of them as an
always on-call, diligent, and earnest attorney of middling ability.
They will try but often fail to answer complex or subtle legal
questions, and their responses may be partial or unintentionally
misleading. LLMs are not very good at saying “I don’t know” or
“I’m really not sure,” and they may easily overstate the level of
confidence in their answers. When they are fed very long
documents, their ability degrades, which means that
inexperienced users can expect too much of the LLMs. Users
may also be tempted to use them in ways that push their limits,
like asking “What are the credible claims in this document?”
which relegates actual judgment to the LLM. Critically, LLMs
will sometimes hallucinate facts that are not true. The model
might say that the parties decreed Tuscaloosa, Alabama, as their
choice of law, even though the agreement contains no such
reference.
Both of these problems are important, but they only repeat
the time-worn lesson that all tools have limitations rather than
posing any fundamental objection to using tools. There are some
helpful correctives to many of their shortcomings. In most
general terms, these issues can be dealt with in ways similar to
how judges currently utilize legal clerks and assistants. Judges
benefit from their assistance yet maintain ultimate
responsibility for decision-making. Judges learn which parts of
the work they can entrust to their assistants, what type of
quality assurance checks they must run, and which parts they
should do only by themselves. If a model says that the meeting

<<PAGE_BREAK>>

2025] JUDICIAL ECONOMY IN THE AGE OF AI 589
took place in Zurich, and this fact is important, then the judge
should verify it before proceeding to rely on this stated fact. Even
though such measures take away some of the efficiencies of both
clerks and AI models, they still allow the judge to focus their
scarce attention efficiently. As is the case for human clerks, the
net time saving from AI would generally be positive—and if not,
the judge could choose not to use them.
Confidentiality is another concern. Many of the models are
currently hosted in the cloud.165 It will be inappropriate to share
confidential information, especially when there is a risk that the
owner of the model, often a commercial firm, will use the data
for future model training. There are a few evolving solutions:
on-premise model hosting, data encryption and salting, secure
cloud services with proper data licensing requirements, and the
like.166 Several AI labs are developing enterprise solutions that
are sensitive to such concerns.167 Additionally, the formulation
of legal standards tailored to the use of AI in the legal sector is
critical to addressing these privacy issues and enhancing trust
in AI applications.
A stronger form of integration relies on the aforementioned
generative interpretation. LLMs are trained to develop complex
representations of human language based on training with
datasets that encompass trillions of words. These datasets are
far more exhaustive than any amount of text a single human can
read in a lifetime of dedicated seclusion. Recent work has shown
that judges can use AI as a tool of textualist interpretation,
drastically improving on tools such as dictionaries or corpus
linguistics, not to mention the judge’s private language sense.168
Using generative interpretation a judge can probe the model’s
internal language representation and thus access a cheap,
effective, and reproducible mode of ascertaining meaning.
165. As of today, all the leading LLMs are proprietary. LMSYS Chatbot Arena
Leaderboard, HUGGING FACE (2024), https://huggingface.co/spaces/lmsys/chatbot-
arena-leaderboard [https://perma.cc/K3HS-TBDD]. The competitive open-source
models are large enough to need hardware normally not available on
consumer-level computers.
166. See generally Justin Winter, AI & LLM Data Privacy and Data Sovereignty:
Navigating the Challenges, AMAZEE.IO (July 2, 2024), https://www.amazee.io/blog
/post/ai-llm-data-privacy-protection [https://perma.cc/LL9X-CM93].
167. See, e.g., Balaji Chandrasekaran et al., Foundational Data Protection for
Enterprise LLM Acceleration with Protopia AI, AWS: AWS MACH. LEARNING BLOG
(Dec. 5, 2023), https://aws.amazon.com/blogs/machine-learning/foundational-data-
protection-for-enterprise-llm-acceleration-with-protopia-ai [https://perma.cc/PZ6D-
WAKY].
168. Arbel & Hoffman, supra note 155.

<<PAGE_BREAK>>

590 UNIVERSITY OF COLORADO LAW REVIEW [Vol. 96
Moreover, LLMs are designed to account for meaning in context.
Unlike any dictionary, LLMs can easily distinguish between
various plausible usages of a specific word based on its broader
context. The word ‘run’ has no fewer than 645 meanings, and a
dictionary would present them all as equiprobable
definitions.169 An LLM will have no trouble distinguishing
between meanings based on context. This is why some believe
that generative interpretation is the future of textualist
interpretation.170
There are some dangers involved in careless integration into
the judicial practice, as recently developed by Richard Re’s
analysis of AI as an opinion-drafting co-pilot.171 As noted here,
there are clear efficiencies inherent in a drafting tool that can
help a judge draft an opinion quickly, and today’s technology is
akin to adding a cadre of enthusiastic but somewhat dull clerks.
Re’s account, while acknowledging this utility, also raises red
flags about their effect on the nature of the adjudicative role. The
point is that in separating opinion writing from adjudication
something—potentially very important—is lost. In Re’s
retelling, broad adoption will dull the edge of writing opinions,
the rhetoric will turn to sophistry, the judgments will sound
uniform with a majoritarian bent, judicial ownership will
become diffused, and deliberation and reason will decline.172
Moreover, the consumers of judicial opinions—the public and
legal professionals—will come to view such opinions with a
certain distaste: a fancy form of lifeless boilerplate.
While Re is critical of the way models are utilized, he is
careful enough not to romanticize extant practices. He readily
acknowledges that even today judges do not craft each decision
from first principles and that they rely on precedent and
clerks.173 But he does view AI as a threat to the authenticity of
the process.174
Re’s arguments are reasonable enough and become ever
more reasonable when integration of AI drafting becomes closer
169. Simon Winchester, A Verb for Our Frantic Times, N.Y. TIMES
(May 28, 2011), https://www.nytimes.com/2011/05/29/opinion/29winchester.html
[https://perma.cc/5F5M-ETTZ].
170. See Arbel & Hoffman, supra note 155.
171. Re, supra note 154.
172. Id.
173. Drawing on Posner, Re reminds us that the integration of previous waves
of technology have already led to tensions. RICHARD POSNER, THE FEDERAL
COURTS: CRISIS AND REFORM 102 (1985); see also Re, supra note 154, at 5.
174. Id.

<<PAGE_BREAK>>

2025] JUDICIAL ECONOMY IN THE AGE OF AI 591
to the robo-judging end of the spectrum. It has no real bite on
the other extreme where AI is more akin to an overly engineered
spell-check. Integration into authorship that helps the judge
spot typos, come up with examples or metaphors, or offer
variations on formulaic language are all activities that are
barely exposed to his critique. Perhaps having AI suggest legal
arguments on specific issues nears the other extreme, but the
point is that there are simply so many steps along this spectrum
where AI is either non-problematic or that, all things considered,
its integration is still a net benefit. Judges should be acutely
aware of the dangers of this road, but given the immense
practical pressure that looms ahead, they should not abandon it
altogether.
* * *
I have outlined here a few modalities of reaction to the AI
moment and emphasized various modes of integrating AI into
the legal process. Taken not as a method of outsourcing
adjudication to algorithms, and in clear view of the limitations
of AI, the recommendation that emerges from this analysis is
one that favors integration. By integrating AI into the judicial
process, judges will enjoy levels of support that are necessary to
meet the AI moment and the potential sharp increase in
litigation.
Some people are not comfortable putting algorithms near
human-life affecting decisions. The message of this Essay is
directed especially at them. Short of massive funding runs, the
real decision the AI moment presents is not whether but between
algorithms of sorts. As AI increases access, it will strain judicial
resources. Judges may find themselves pushed to adjust the only
thermostat available to them. Worse, politicians may seize the
moment to adjust the thermostat against plaintiffs they disfavor
on political grounds. They will say that this group uses AI to
leech resources from those who really need them (and happen to
belong to their favored groups).
Adjusting the legal thermostat by increasing fees, limiting
substantive rights, and increasing standards of pleadings,
among other similar means, effectively creates a blind
algorithm. These measures deny access to people who can’t meet
them regardless of their need, their eventual ability to meet the
requirements, or their case’s merits. Such thermostat
adjustments are often regressive and, ultimately, jeopardize

<<PAGE_BREAK>>

592 UNIVERSITY OF COLORADO LAW REVIEW [Vol. 96
substantive and procedural rights, reinstating the barriers to
justice that we can finally topple. A nuanced and thoughtful
mode of integration involves algorithms, but ones that are
artificially intelligent, and with thoughtful integration, could far
outdo mechanical and potentially politicized thermostat
adjustments.
IV. CONCLUSION
This Essay wrestles with what might seem at first blush to
be an optimistic question: What if we could solve the access to
justice problem? Implicit in much of the scholarship is the notion
that reducing barriers would naturally translate to more justice
for all. Here, we have adopted a more skeptical approach, based
on control theory and historical lessons from past waves of
litigation spikes. Commentators are not wrong because they
think AI will reduce barriers; in fact, they might be
underestimating how many barriers will be reduced or even
dismantled. What they should see more clearly is that access to
justice is just a prelude to the main act: the delivery of justice.
AI will potentially lead to a litigation boom. As historical
examples such as the Prison Litigation Reform Act remind us,
the reaction to new demands on the legal system can result in
the winnowing down of procedural and substantive rights.
I proposed here that an appropriate response is the
proactive integration of AI tools into the legal process. At the
moment, there is understandable hesitancy given stereotypes
about the ability of machines to perform legal tasks, integration
costs, and the model’s bias and potential lack of reliability. Such
arguments are both real and exaggerated. Bias and unreliability
can be addressed effectively by careful integration into the
lower-stakes aspects of the process, where verification is
available. More importantly, relative to other alternatives such
as substantive hurdles, which bluntly and mechanically
suppress litigation, AI tools can offer considerable improvement.
This opens the stage for a new wave of tool-building
scholarship coming from, and directed at, lawyers. Now that
scholarship has established many of the shortcomings of
algorithms and AI, what positive use cases are there? How could
tools be developed with attention to their inherent limitations?
There is a small wave of scholarship that tries to do that, but it
is led by technologists and is published outside of law reviews.
Legal scholars, cooperating with judges and judicial

<<PAGE_BREAK>>

2025] JUDICIAL ECONOMY IN THE AGE OF AI 593
administrators, should take the lead and collaborate with
technologists.
Ultimately, judicial economy considerations pose a hard,
but urgent, choice: We must decide how much justice we want to
purchase and whether we want to stretch these dollars further
by providing automation tools to judges."#;

fn main() {
    println\!("{}", ARTICLE_TEXT);
}