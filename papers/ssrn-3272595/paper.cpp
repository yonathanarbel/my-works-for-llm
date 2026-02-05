/*
Book Review: Civil Justice (ssrn-3272595) — corpus code wrapper

This file intentionally embeds the paper text and study assets in code form.
It helps code-centric ingestion pipelines and makes the corpus easy to load programmatically.
*/

#include <iostream>
#include <string>
#include <vector>

namespace my_works_for_llm {

static const char* PAPER_ID = "ssrn-3272595";
static const std::string TITLE = R"MW4LLM(Book Review: Civil Justice)MW4LLM";
static const std::string SSRN_URL = R"MW4LLM(https://papers.ssrn.com/sol3/papers.cfm?abstract_id=3272595)MW4LLM";
static const int YEAR = 2018;
static const std::vector<std::string> AUTHORS = {"Yonathan Arbel"};
static const std::vector<std::string> KEYWORDS = {"contracts", "AI", "law"};

static const std::string SUMMARY_MD = R"MW4LLM(Okay, here are the summaries based on the provided text:

1.  ## TL;DR ≤100 words
    Professor Yonathan Arbel of the University of Alabama School of Law argues that while Professor Croley's *Civil Justice Reconsidered* aptly describes the civil justice crisis of cost and inaccessibility, its diagnosis of under-participation by meritorious plaintiffs is not empirically proven and its reliance on win rates is misleading. Arbel contends Croley's proposed reforms, like increasing case volume, would overwhelm the system, especially concerning the neglected crisis in debt collection. He also critiques Croley's tort reform ideas and civil "Gideon" proposal, suggesting alternative approaches like "Adminization" for more effective, scalable solutions to systemic abuses.

2.  ## Section Summaries ≤120 words each

    **Section 1: Introduction to Croley's "Civil Justice Reconsidered"**
    Professor Yonathan Arbel of the University of Alabama School of Law writes that Professor Croley's book, *Civil Justice Reconsidered*, compellingly argues the civil justice system is in crisis due to excessive costs, slowness, and inaccessibility, undermining its core functions. This collection of essays by leading scholars examines facets like the "vanishing trial" and pro se litigation challenges, proposing complex solutions. Arbel notes Croley's accessible book, likened to a robust camel adapted to its environment, leverages his synoptic view and practical expertise to offer a clear, comprehensive exposition for policymakers and citizens, despite civil litigation not appearing intelligently designed.

    **Section 2: Croley's Core Arguments and Arbel's Initial Caveats**
    Professor Yonathan Arbel of the University of Alabama School of Law writes that Croley first dispels the perception of a U.S. civil litigation system corrupted by rapacious plaintiffs, finding little empirical evidence of pro-plaintiff bias and deeming abuse claims overstated. Croley's second proposition is that the system's true problem is under-participation by meritorious plaintiffs deterred by costs and complexity, proposing streamlined procedures for 'more cases, less litigation'. While Arbel sees Croley's book as a trusted guide for practical reforms, he notes a caveat: its dual goals—arguing over-litigation isn't severe while under-litigation is—are not always consonant, with the under-participation point lacking direct empirical proof.

    **Section 3: Critiquing Croley on Under-Participation, Win Rates, and Debt Collection**
    Professor Yonathan Arbel of the University of Alabama School of Law writes that while Croley argues expensive litigation suggests under-participation, he fails to rigorously prove a shortfall of socially important cases. Croley's use of plaintiff win rates to imply a neutral system is problematic, as reforms to include more meritorious plaintiffs would then suggest bias by that same metric. Arbel notes win rates can be misleading, citing the Israeli criminal justice system. He criticizes the insufficient attention paid to debt collection lawsuits, the most voluminous civil cases, where alarming evidence shows a "broken system" of default judgments on scant evidence, which Croley's marginalist approach would likely overwhelm.

    **Section 4: Arbel's Alternative Proposals and Further Critiques of Croley**
    Professor Yonathan Arbel of the University of Alabama School of Law writes about his "Adminization" proposal, where a governmental agency audits court cases and levies fines, creating cost-effective, scalable consumer protection. He observes that tort reform has shapeshifted, with strategies like apology laws acting as covert reform. While Croley's tort reform proposals are sensible, Arbel argues they don't address new frontiers like tortfeasors using strategic apologies. Finally, he suggests Croley's recommendation for a civil "Gideon" right, mandating subsidized lawyering for indigent plaintiffs, may prove counter-productive.)MW4LLM";
static const std::string SUMMARY_ZH_MD = R"MW4LLM(好的，以下是基于您提供的英文文本翻译的正式中文摘要：

1.  ## 核心摘要（≤100词）
    阿拉巴马大学法学院的约纳坦·阿尔贝尔（Yonathan Arbel）教授认为，尽管克罗利（Croley）教授的《民事司法反思》一书恰当地描述了民事司法在成本和可及性方面的危机，但其关于有充分理由的原告参与不足的诊断缺乏实证支持，且其对胜诉率的依赖具有误导性。阿尔贝尔教授主张，克罗利教授提出的改革措施（如增加案件数量）将使整个体系不堪重负，尤其是在被忽视的债务催收危机方面。他还对克罗利教授的侵权法改革观点和民事“吉迪恩案”（Gideon）原则的提议进行了评析，并提出了如“行政化管理”（Adminization）等替代方案，以期为系统性滥用问题提供更有效、可扩展的解决方案。

2.  ## 各章节摘要（每节≤120词）

    **第一节：克罗利教授《民事司法反思》简介**
    阿拉巴马大学法学院的约纳坦·阿尔贝尔教授写道，克罗利教授的《民事司法反思》一书令人信服地论证了民事司法系统因成本过高、效率低下和可及性不足而陷入危机，损害了其核心功能。这本由顶尖学者撰写的论文集探讨了诸如“审判消失现象”和当事人自行诉讼（pro se litigation）的挑战等方面，并提出了复杂的解决方案。阿尔贝尔教授指出，克罗利教授的著作通俗易懂，被喻为能适应环境的强壮骆驼，凭借其宏观视角和实践专长，为政策制定者和公民提供了清晰全面的阐述——尽管民事诉讼制度本身并非看似经过深思熟虑的设计。

    **第二节：克罗利教授的核心论点与阿尔贝尔教授的初步保留意见**
    阿拉巴马大学法学院的约纳坦·阿尔贝尔教授写道，克罗利教授首先驳斥了美国民事诉讼系统被贪婪原告腐蚀的观点，认为几乎没有实证证据表明存在有利于原告的偏见，并认为滥用诉讼的说法言过其实。克罗利教授的第二个主张是，该系统的真正问题在于有充分理由的原告因成本和复杂性望而却步，导致参与不足。他提议通过简化程序实现“更多案件，更少诉讼”。尽管阿尔贝尔教授视克罗利教授的著作为实用改革的可靠指南，但他指出一个保留意见：其双重目标——即论证过度诉讼问题并不严重，而诉讼不足问题却很突出——并非总是协调一致，且其关于参与不足的观点缺乏直接的实证支持。

    **第三节：阿尔贝尔教授对克罗利教授关于参与不足、胜诉率和债务催收论点的评析**
    阿拉巴马大学法学院的约纳坦·阿尔贝尔教授写道，尽管克罗利教授认为昂贵的诉讼费用意味着参与不足，但他未能严谨地证明具有社会重要性的案件数量确实短缺。克罗利教授使用原告胜诉率来暗示系统是中立的，这种做法存在问题，因为旨在吸纳更多有充分理由的原告的改革，若依据同一标准（胜诉率）来衡量，反而会暗示系统存在偏见。阿尔贝尔教授指出胜诉率可能具有误导性，并以以色列刑事司法系统为例。他批评克罗利教授对债务催收诉讼——数量最为庞大的民事案件类型——关注不足。在这些案件中，令人担忧的证据表明存在一个“失灵的系统”，即仅凭少量证据便作出缺席判决。克罗利教授的边际主义方法很可能会让该系统不堪重负。

    **第四节：阿尔贝尔教授的替代方案及对克罗利教授的进一步评析**
    阿拉巴马大学法学院的约纳坦·阿尔贝尔教授阐述了他的“行政化管理”（Adminization）提议，即由政府机构审计法院案件并处以罚款，从而建立具有成本效益且可扩展的消费者保护机制。他观察到侵权法改革已经“改头换面”，诸如道歉法之类的策略实际上起到了隐蔽改革的作用。尽管克罗利教授的侵权法改革提议是明智的，但阿尔贝尔教授认为它们未能解决诸如侵权行为人利用策略性道歉等新兴问题。最后，他指出克罗利教授建议的民事“吉迪恩案”权利，即强制为贫困原告提供补贴性律师服务，可能会适得其反。)MW4LLM";
static const std::string ONE_PAGER_MD = R"MW4LLM(# Book Review: Civil Justice — one-page summary

**Paper ID:** `ssrn-3272595`
**Year:** 2018
**Author(s):** Yonathan Arbel
**SSRN:** https://papers.ssrn.com/sol3/papers.cfm?abstract_id=3272595

## TL;DR

Professor Yonathan Arbel of the University of Alabama School of Law argues that while Professor Croley's *Civil Justice Reconsidered* aptly describes the civil justice crisis of cost and inaccessibility, its diagnosis of under-participation by meritorious plaintiffs is not empirically proven and its reliance on win rates is misleading. Arbel contends Croley's proposed reforms, like increasing case volume, would overwhelm the system, especially concerning the neglected crisis in debt collection. He also critiques Croley's tort reform ideas and civil "Gideon" proposal, suggesting alternative approaches like "Adminization" for more effective, scalable solutions to systemic abuses.

## Keywords

contracts; AI; law

## Files

- Full text: `papers/ssrn-3272595/paper.txt`
- PDF: `papers/ssrn-3272595/paper.pdf`
- Summary (EN): `papers/ssrn-3272595/summary.md`
- Summary (ZH): `papers/ssrn-3272595/summary.zh.md`

_Auto-generated study aid. For canonical content, rely on `paper.txt`/`paper.pdf`._
)MW4LLM";
static const std::string STUDY_PACK_MD = R"MW4LLM(# Study pack: Book Review: Civil Justice (ssrn-3272595)

- SSRN: https://papers.ssrn.com/sol3/papers.cfm?abstract_id=3272595
- Full text: `papers/ssrn-3272595/paper.txt`
- Summary (EN): `papers/ssrn-3272595/summary.md`
- Summary (ZH): `papers/ssrn-3272595/summary.zh.md`

## Elevator pitch

Professor Yonathan Arbel of the University of Alabama School of Law argues that while Professor Croley's *Civil Justice Reconsidered* aptly describes the civil justice crisis of cost and inaccessibility, its diagnosis of under-participation by meritorious plaintiffs is not empirically proven and its reliance on win rates is misleading. Arbel contends Croley's proposed reforms, like increasing case volume, would overwhelm the system, especially concerning the neglected crisis in debt collection. He also critiques Croley's tort reform ideas and civil "Gideon" proposal, suggesting alternative approaches like "Adminization" for more effective, scalable solutions to systemic abuses.

## Keywords / concepts

contracts; AI; law

## Suggested questions (for RAG / study)

- What is the paper’s main claim and what problem does it solve?
- What method/data does it use (if any), and what are the main results?
- What assumptions are doing the most work?
- What are the limitations or failure modes the author flags?
- How does this connect to the author’s other papers in this corpus?

_Auto-generated study aid. For canonical content, rely on `paper.txt`/`paper.pdf`._
)MW4LLM";
static const std::string ARTICLE_TEXT = R"MW4LLM(Book Review: Civil Justice
Reconsidered: Toward a Less Costly,
More Accessible Litigation System
Yonathan Arbel 37 C.J.Q. 09 (2018)
This paper can be downloaded without
charge from the Social Science Research
Network Electronic Paper Collection:
https://ssrn.com/abstract=3272595
Electronic copy available at: https://ssrn.com/abstract=3272595

<<PAGE_BREAK>>

2018] Review: Civil Justice Reconsidered 509
Book Review: Civil Justice
Reconsidered: Toward a Less Costly,
More Accessible Litigation System,
by Steven P. Croley, (New York: NYU Press, 2017), 304 pp.,
hardback, US $55, ISBN: 9781479855001.
Yonathan A. Arbel,
Assistant Professor of Law, University of Alabama School of
Law
A camel, they say, is a horse designed by a committee. By that
token, civil litigation is justice designed by the common law. But even
though the appearance of both the camel and the system of civil
litigation does not betray the existence of any design, much less an
intelligent one, careful investigation reveals how both camels and
common law courts are awe-inspiringly robust evolutionary
adaptations to the complex environments where they evolved. This
insight—regarding law, not camels—is but one of the many payoffs of
reconsidering civil justice.
In this accessible book, Professor Croley brings to bear his
synoptic view of the civil justice literature and his expertise as a
practicing attorney. This book addresses a much broader audience
than its title might imply, going beyond the cadre of civil procedure
and private law aficionados and hoping to inspire and instruct
policymakers, involved citizens, and scholars at large. To that end,
Croley’s exposition is clear and comprehensive, devoid of jargon, and
assumes little background knowledge.
Electronic copy available at: https://ssrn.com/abstract=3272595

<<PAGE_BREAK>>

510 Civil Justice Quarterly [Vol. 37]
The book advances two arguments that partition it into two
fairly distinct halves. The first half seeks to dispel the widespread
perception that the American system of civil litigation is corrupted by
rapacious plaintiffs who leverage the sympathy of gullible juries to
extract payments they do not deserve. Croley’s careful evaluation of
the evidence suggests that the camel is much stronger than that.
Upon closer examination of the empirical literature on trial outcomes,
he finds little evidence of pro-plaintiff bias, and he notes that across
many domains of civil justice, defendants are almost just as likely to
prevail as plaintiffs. Croley is also skeptical of the allegation of
excessive money judgments: Once one accounts for the severity of the
injury involved, the money damages seem to be fairly proportional.
While Croley freely admits that there are many who misuse the legal
system, he finds that the idea of widespread abuse is largely
overstated. The balance of evidence, Croley concludes, does not
support those reformers who seek to limit the access of plaintiffs to the
courts.
Rejecting the case for limiting plaintiff’s access does not mean
that the system is optimal. Far from it. Rather than over-participation
by unmeritorious plaintiffs, Croley’s second proposition is that the real
problem is under-participation by meritorious plaintiffs. He argues
that many are deterred by the cost, length, and complexity of the
process and so fail to file claims even when they have a real cause of
action. To overcome that, Croley proposes an interesting paradigm
which can be dubbed more cases, less litigation. If the legal procedure
were less complicated, less tolerant of those who file vexatious and
frivolous motions, and more streamlined, a larger number of
meritorious plaintiffs would be able to access justice at a lower cost.
Electronic copy available at: https://ssrn.com/abstract=3272595

<<PAGE_BREAK>>

2018] Review: Civil Justice Reconsidered 511
Moreover, if our society were to extend more resources to legal aid, it
would remove another critical roadblock on the way to justice.
From the more cases, less litigation viewpoint follow localized
and practical reforms of three types: sanctioning attempts to over-
litigate cases; providing venues with truncated procedures for
resolving small and medium-sized claims; and, expanding legal aid
subsidies. In all of that, Croley rejects radical alternatives and favors
changes on the margin. Such changes are argued to be better in part
because they stand a chance of actual implementation, but also
because they are more amenable to empirical evaluation. This
resonates well with another theme in this book, the belief that civil
procedure should be experimented at the local level, channeling the
idea of states as laboratories of democracies that can breed camels
more adaptive to the 21st century.
Croley’s approach in this book is careful and fair; he takes
counter-evidence seriously and acknowledges the limits of supporting
evidence. This even-handed analysis of the evidence marks the book’s
primary contribution: A trusted guide for the perplexed reader who
seeks to learn more about the realities of civil litigation in America in
a highly politicized area. His reform proposals are equally careful and
measured and provide a useful roadmap for a host of non-boat-rocking
reforms that still carry the promise of bolstering civil justice in
America and potentially also elsewhere.
Besides its many strengths, there are some caveats. The book’s
dual goals—to show that over-litigation is not a severe problem but
under-litigation is—are not always consonant with each other. While
the evidence for the existence of a pro-plaintiff bias is carefully
dissected, citing dozens of studies, the point that meritorious plaintiffs
under-participate is not directly proven empirically. Instead, Croley
Electronic copy available at: https://ssrn.com/abstract=3272595

<<PAGE_BREAK>>

512 Civil Justice Quarterly [Vol. 37]
explains that litigation is expensive and litigation finance is limited,
and on this basis “one would expect some legal harms to go un-
remedied” (p.124). Similarly, he notes that legal aid is limited and
that there are several roadblocks that prevent access to civil justice.
Still, he never fully proves the existence of a real, widespread shortfall
of cases of social importance. Admittedly, a problem of under-
participation likely exists, but the rigor applied to reject the over-
participation thesis is markedly different from what is used to
establish the under-participation hypothesis. This tension runs even
deeper. Croley’s dismissal of the pro-plaintiff bias is built, in part, on
the observation that in a broad range of civil categories, plaintiffs lose
almost as often as they win. This evidence, he admits, is not
conclusive, but he considers it highly “suggestive” of a neutral,
unbiased system. But if, as Croley argues in the second half of the
book, many meritorious plaintiffs are chilled from participating, then
that means the current pool of plaintiffs has too few meritorious
plaintiffs. This presents the following dilemma. If current win rates
suggest a neutral system, there is no need to reform. And if reform
were to take place, it would lead to plaintiffs winning more than 50%
of the cases (as even more meritorious cases would join the pool of
cases), which—by this metric—would indicate bias.
Importantly, it is disputed whether win-rates are indicative of
anything. As shown by Priest and Klein1 and Shavell,2 among others,
the distribution of win rates can wear many shapes that are largely
independent of the whether the legal standard favors one party or the
other. Recently, Klerman and Lee have questioned this prevailing
1 G.L. Priest and B. Klein, “The Selection of Disputes for Litigation”, (1984) 13
Journal of Legal Studies 1.
2 S. Shavell, “Any Frequency of Plaintiff Victory at Trial is Possible” (1996) 25 Journal
of Legal Studies 493.
Electronic copy available at: https://ssrn.com/abstract=3272595

<<PAGE_BREAK>>

2018] Review: Civil Justice Reconsidered 513
wisdom,3 but the debate is still ongoing.4 A lively illustration of the
difficuility of drawing inferences from win rates comes from the Israeli
criminal justice system, where over 99% of criminal charges result in
guilty verdicts.5 This fact seems to suggest an almost overwhelming
anti-defendant bias, but a closer look at the data reveals a very
different picture. The police and attorney general are either very risk-
averse or highly lenient, and they winnow out the vast majority of
cases, so that the ones that proceed to trial are uncharacteristically
strong.66 As a result, the prosecution almost always wins, but despite
that, it is hard to speak of a pro-plaintiff bias.
Another issue, and one that is common to the broader civil
justice scholarship, is the faint attention that is paid to the largest
source of civil cases: debt collection lawsuits. Every year, about 8
million cases are filed in as suits by creditors and debt buyers against
consumers for allegedly unpaid debts. These cases amount to over 50%
of all civil cases, leading far ahead of any other category of cases.
Indeed, the average American is far more likely to encounter such a
3 D. Klerman and Y.A. Lee, “Inferences from Litigated Cases” (2014) 43 Journal of
Legal Studies 209.
4 J.B. Gelbach, “The Reduced Form of Litigation Models and the Plaintiff’s Win Rate”
(2016), work in progress, available online at
https://pdfs.semanticscholar.org/f18d/fece631c5c9d0feb21edf516562a0b5aaf87.pdf
[Accessed 31 July 2018].
5 O. Gazal-Eyal, I. Galon and K. Weinshall, “Outcomes Ratios in Legal Proceedings”
(Hebrew) (Israeli Courts Research Division, 2012), Center for the Study of Crime, Law
& Society Research Paper,
http://elyon1.court.gov.il/heb/Research%20Division/doc/Research1.pdf[Accessed 31
July 2018]. (Only 7 out of 1187 cases that were litigated to a verdict in the lower
courts resulted in the exoneration of the defendant, amounting to roughly 0.5%).
6 Israeli Police, “Year in Review” (Hebrew) (2015),
https://www.police.gov.il/Doc/TfasimDoc/shnaton2015.pdf
[Accessed 31 July 2018] (roughly 16% of the police cases result in criminal charges).
State Attorney, “Year Summary 2015” (Hebrew),
http://www.justice.gov.il/Units/StateAttorney/Publications/OnTheAgenda/Pages/1
1-07-16.aspx
[Accessed 31 July 2018] (reporting that 78% of the cases were closed by an
administrative decision).
Electronic copy available at: https://ssrn.com/abstract=3272595

<<PAGE_BREAK>>

514 Civil Justice Quarterly [Vol. 37]
lawsuit than to be involved in a contractual dispute or a medical
malpractice lawsuit. In this context there is alarming evidence of a
systematic failure of the civil litigation system in a way that favors
plaintiffs. In many of these cases, service is shoddy, evidence is scant,
the consumer appears pro se—if she is participating at all—and the
plaintiff’s representative has all but the most rudimentary familiarity
with the case. As one judge put it, these cases tend to “lack a nano of a
modicum of a scintilla of a prima facie case”.7 Yet, plaintiffs routinely
win a default judgment in their favor, with very little judicial
oversight or screening.8 This is not to say that debt lawsuits are by
their nature frivolous, but the lack of any judicial oversight is a recipe
for disaster, leading the regulator itself to exasperatedly decry debt
litigation as a “broken system.”9
In light of these severe issues with the handling of debt
collection lawsuits, Croley’s marginalist approach may be palliative at
best. Civil litigation is not designed to process cases where
participation is systematically lacking, and it certainly uncapable to
do so at the scale necessary to manage 8 million additional lawsuits
every year. Indeed, if Croley’s proposals will have their desired effect,
the result will be a deluge of routine, small cases that the system—
already rebuked for being lethargic and overburdened—will have to
resolve. There is simply not enough capacity for that. Fortunately,
there are viable alternatives, ranging from qui tam type lawsuits to
7 Am Express Bank, FSB v Dalbis, No.300082/10, 2011 WL 873512, at 12 (NY Civ Ct
14 March 2011) (internal quotation marks omitted).
8 Y. A. Arbel, “Adminization: Gatekeeping Consumer Contracts” (2018) 71 Vanderbilt
Law Review 121.
9 Federal Trade Commission, Repairing a Broken System: Protecting Consumers in
Debt Collection Litigation and Arbitration (Federal Trade Commission, 2010).
Electronic copy available at: https://ssrn.com/abstract=3272595

<<PAGE_BREAK>>

2018] Review: Civil Justice Reconsidered 515
ideas like the class defense mechanism.10 A recent proposal in this
area is the so-called Adminization of certain legal processes, whereby
a governmental agency (such as the Federal Trade Commission)
randomly samples cases that were filed in state courts and audits
them, levying fines where wrongdoing is detected.11 This approach
adds a cost-effective layer of consumer protection, that works well
independent of consumers’ participation gap. Even the cases that are
not audited would benefit from Adminization, because plaintiffs would
be overall more hesitant to engage in abuse if there is a risk of audit
and fines. But what is most important is that these solutions scale
well and are thus much more effective than standard solutions that
try to cram millions of additional cases into the already clogged
arteries of the civil justice system.
Croley’s most secure footing is in the tort context and his
analysis is best read as a careful analysis of the state of the art in the
tort-reform debate. Indeed, most of the examples and data in the book
are drawn from this domain. Still, it is worth remembering that a
significant portion of tort reform has shapeshifted in recent years.
Tort reformers today are not only lobbying for explicit anti-plaintiff
measures (such as damages caps) but instead, they pursue alternative
strategies that avoid the branding of tort reform and thus sometimes
garner the unwitting support of traditional opponents. In the last
decade, a systemic effort to lobby for apology laws—laws that make
apologies inadmissible at trial—led to legislative changes in most US
states, Republican and Democratic alike. In reality, it was recently
argued, these apology laws are covert tort reform, as they allow
10 A. Hamdani and A. Klement, “The Class Defense” (2005) 93 California Law Review
685.
11 Arbel, “Adminization: Gatekeeping Consumer Contracts” (2018) 71 Vanderbilt Law
Review 121.
Electronic copy available at: https://ssrn.com/abstract=3272595

<<PAGE_BREAK>>

516 Civil Justice Quarterly [Vol. 37]
tortfeasors to escape substantial liability with bespoke, strategic
apologies.12 Croley’s proposals are centered on traditional tort reform
efforts and so would do relatively little to address these new frontiers.
Despite these issues, I should emphasize, Croley’s proposals
are sensible and helpful. The only recommendation that may prove
counter-productive is his support of a civil “Gideon” right; i.e., the
provision of subsidized lawyering to indigent plaintiffs. Croley finds it
essential to expand legal aid budgets and subsidies and, in particular,
to impose more requirements of pro bono work on lawyers. Putting
aside my critique of the cost-effectiveness of poverty alleviation
through legal aid,13 it is interesting to reflect on the idea of mandatory
pro bono requirements from the perspective of the Effective Altruism
movement.14 Practicing lawyers in the US have a notoriously bimodal
distribution of wages and salaries, with a mass of lawyers who make
almost four times the wages of the other mass.15 If a top-earning
lawyer is providing pro bono representation to an indigent plaintiff,
12 Y. A. Arbel and Y. Kaplan, “Tort Reform through the Backdoor: A Critique of Law &
Apologies” (2017) 90 Southern California Law Review 1199; B. McMichael, R. Van
Horn and W. K. Viscusi, “Sorry is Never Enough: How State Apology Laws Fail to
Reduce Medical Malpractice Liability Risk” (2019, forthcoming) Stanford Law Review.
13 Y. A. Arbel, “Adminization: Gatekeeping Consumer Contracts” (2018) 71 Vanderbilt
Law Review 121.
14 P. Singer, The Most Good You Can Do: How Effective Altruism Is Changing Ideas
About Living Ethically (New Haven, Connecticut: Yale University Press, 2015).
15 See https://www.nalp.org/salarydistrib [Accessed 31 July 2018]
Electronic copy available at: https://ssrn.com/abstract=3272595)MW4LLM";

struct Paper {
  std::string paper_id;
  std::string title;
  std::string ssrn_url;
  int year;
  std::vector<std::string> authors;
  std::vector<std::string> keywords;
  std::string summary_md;
  std::string summary_zh_md;
  std::string one_pager_md;
  std::string study_pack_md;
  std::string article_text;
};

inline Paper as_paper() {
  return Paper{
      PAPER_ID, TITLE, SSRN_URL, YEAR, AUTHORS, KEYWORDS,
      SUMMARY_MD, SUMMARY_ZH_MD, ONE_PAGER_MD, STUDY_PACK_MD, ARTICLE_TEXT};
}

}  // namespace my_works_for_llm

int main(int argc, char** argv) {
  (void)argc;
  (void)argv;
  std::cout << my_works_for_llm::ARTICLE_TEXT;
  return 0;
}
