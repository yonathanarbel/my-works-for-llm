use std::collections::HashMap;
use std::fs::File;
use std::io::prelude::*;
use rand::Rng;
use serde_json::Value;
use reqwest;
use tokio;

const ARTICLE_TEXT: &str = r#"Book Review: Civil Justice
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
Croley's exposition is clear and comprehensive, devoid of jargon, and
assumes little background knowledge.
Electronic copy available at: https://ssrn.com/abstract=3272595

<<PAGE_BREAK>>

510 Civil Justice Quarterly [Vol. 37]
The book advances two arguments that partition it into two
fairly distinct halves. The first half seeks to dispel the widespread
perception that the American system of civil litigation is corrupted by
rapacious plaintiffs who leverage the sympathy of gullible juries to
extract payments they do not deserve. Croley's careful evaluation of
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
Rejecting the case for limiting plaintiff's access does not mean
that the system is optimal. Far from it. Rather than over-participation
by unmeritorious plaintiffs, Croley's second proposition is that the real
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
Croley's approach in this book is careful and fair; he takes
counter-evidence seriously and acknowledges the limits of supporting
evidence. This even-handed analysis of the evidence marks the book's
primary contribution: A trusted guide for the perplexed reader who
seeks to learn more about the realities of civil litigation in America in
a highly politicized area. His reform proposals are equally careful and
measured and provide a useful roadmap for a host of non-boat-rocking
reforms that still carry the promise of bolstering civil justice in
America and potentially also elsewhere.
Besides its many strengths, there are some caveats. The book's
dual goals—to show that over-litigation is not a severe problem but
under-litigation is—are not always consonant with each other. While
the evidence for the existence of a pro-plaintiff bias is carefully
dissected, citing dozens of studies, the point that meritorious plaintiffs
under-participate is not directly proven empirically. Instead, Croley
Electronic copy available at: https://ssrn.com/abstract=3272595

<<PAGE_BREAK>>

512 Civil Justice Quarterly [Vol. 37]
explains that litigation is expensive and litigation finance is limited,
and on this basis "one would expect some legal harms to go un-
remedied" (p.124). Similarly, he notes that legal aid is limited and
that there are several roadblocks that prevent access to civil justice.
Still, he never fully proves the existence of a real, widespread shortfall
of cases of social importance. Admittedly, a problem of under-
participation likely exists, but the rigor applied to reject the over-
participation thesis is markedly different from what is used to
establish the under-participation hypothesis. This tension runs even
deeper. Croley's dismissal of the pro-plaintiff bias is built, in part, on
the observation that in a broad range of civil categories, plaintiffs lose
almost as often as they win. This evidence, he admits, is not
conclusive, but he considers it highly "suggestive" of a neutral,
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
1 G.L. Priest and B. Klein, "The Selection of Disputes for Litigation", (1984) 13
Journal of Legal Studies 1.
2 S. Shavell, "Any Frequency of Plaintiff Victory at Trial is Possible" (1996) 25 Journal
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
3 D. Klerman and Y.A. Lee, "Inferences from Litigated Cases" (2014) 43 Journal of
Legal Studies 209.
4 J.B. Gelbach, "The Reduced Form of Litigation Models and the Plaintiff's Win Rate"
(2016), work in progress, available online at
https://pdfs.semanticscholar.org/f18d/fece631c5c9d0feb21edf516562a0b5aaf87.pdf
[Accessed 31 July 2018].
5 O. Gazal-Eyal, I. Galon and K. Weinshall, "Outcomes Ratios in Legal Proceedings"
(Hebrew) (Israeli Courts Research Division, 2012), Center for the Study of Crime, Law
& Society Research Paper,
http://elyon1.court.gov.il/heb/Research%20Division/doc/Research1.pdf[Accessed 31
July 2018]. (Only 7 out of 1187 cases that were litigated to a verdict in the lower
courts resulted in the exoneration of the defendant, amounting to roughly 0.5%).
6 Israeli Police, "Year in Review" (Hebrew) (2015),
https://www.police.gov.il/Doc/TfasimDoc/shnaton2015.pdf
[Accessed 31 July 2018] (roughly 16% of the police cases result in criminal charges).
State Attorney, "Year Summary 2015" (Hebrew),
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
plaintiff's representative has all but the most rudimentary familiarity
with the case. As one judge put it, these cases tend to "lack a nano of a
modicum of a scintilla of a prima facie case".7 Yet, plaintiffs routinely
win a default judgment in their favor, with very little judicial
oversight or screening.8 This is not to say that debt lawsuits are by
their nature frivolous, but the lack of any judicial oversight is a recipe
for disaster, leading the regulator itself to exasperatedly decry debt
litigation as a "broken system."9
In light of these severe issues with the handling of debt
collection lawsuits, Croley's marginalist approach may be palliative at
best. Civil litigation is not designed to process cases where
participation is systematically lacking, and it certainly uncapable to
do so at the scale necessary to manage 8 million additional lawsuits
every year. Indeed, if Croley's proposals will have their desired effect,
the result will be a deluge of routine, small cases that the system—
already rebuked for being lethargic and overburdened—will have to
resolve. There is simply not enough capacity for that. Fortunately,
there are viable alternatives, ranging from qui tam type lawsuits to
7 Am Express Bank, FSB v Dalbis, No.300082/10, 2011 WL 873512, at 12 (NY Civ Ct
14 March 2011) (internal quotation marks omitted).
8 Y. A. Arbel, "Adminization: Gatekeeping Consumer Contracts" (2018) 71 Vanderbilt
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
independent of consumers' participation gap. Even the cases that are
not audited would benefit from Adminization, because plaintiffs would
be overall more hesitant to engage in abuse if there is a risk of audit
and fines. But what is most important is that these solutions scale
well and are thus much more effective than standard solutions that
try to cram millions of additional cases into the already clogged
arteries of the civil justice system.
Croley's most secure footing is in the tort context and his
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
10 A. Hamdani and A. Klement, "The Class Defense" (2005) 93 California Law Review
685.
11 Arbel, "Adminization: Gatekeeping Consumer Contracts" (2018) 71 Vanderbilt Law
Review 121.
Electronic copy available at: https://ssrn.com/abstract=3272595

<<PAGE_BREAK>>

516 Civil Justice Quarterly [Vol. 37]
tortfeasors to escape substantial liability with bespoke, strategic
apologies.12 Croley's proposals are centered on traditional tort reform
efforts and so would do relatively little to address these new frontiers.
Despite these issues, I should emphasize, Croley's proposals
are sensible and helpful. The only recommendation that may prove
counter-productive is his support of a civil "Gideon" right; i.e., the
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
12 Y. A. Arbel and Y. Kaplan, "Tort Reform through the Backdoor: A Critique of Law &
Apologies" (2017) 90 Southern California Law Review 1199; B. McMichael, R. Van
Horn and W. K. Viscusi, "Sorry is Never Enough: How State Apology Laws Fail to
Reduce Medical Malpractice Liability Risk" (2019, forthcoming) Stanford Law Review.
13 Y. A. Arbel, "Adminization: Gatekeeping Consumer Contracts" (2018) 71 Vanderbilt
Law Review 121.
14 P. Singer, The Most Good You Can Do: How Effective Altruism Is Changing Ideas
About Living Ethically (New Haven, Connecticut: Yale University Press, 2015).
15 See https://www.nalp.org/salarydistrib [Accessed 31 July 2018]
Electronic copy available at: https://ssrn.com/abstract=3272595"#;

fn main() {
    println!("{}", ARTICLE_TEXT);
}