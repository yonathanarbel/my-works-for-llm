#include <iostream>
#include <string>
#include <vector>
#include <map>
#include <fstream>
#include <random>
#include <algorithm>
#include <ctime>
#include <thread>

const std::string ARTICLE_TEXT = R"EOF(VILLANOVA
Public Law and Legal Theory Working Paper Series
Tort Reform through the Backdoor:
A Critique of Law and Apologies
By
Yonathan A. Arbel
and
Yotam Kaplan
September 9, 2016
Villanova University Charles Widger School of Law
Public Law and Legal Theory
Working Paper No. 2016-1030
This paper can be downloaded without charge from the
Social Science Research Network Electronic Paper
Collection at http://ssrn.com/abstract=2835482
Electronic copy available at: https://ssrn.com/abstract=2835482
Tort Reform through the Backdoor:
A Critique of Law and Apologies
Yonathan A. Arbel & Yotam Kaplan*
In this Article we show how the biggest tort reform of the
last decade was passed through the backdoor with the
blessing of its staunchest opponents. We argue that the
widely-endorsed apology law reform-----a change in the
national legal landscape that privileged apologies-----is, in
fact, a mechanism of tort reform, used to limit victims’
recovery and shield injurers from liability. While legal
scholars overlooked this effect, commercial interests seized
the opportunity and are in the process of transforming state
and federal law with the unwitting support of the public.
* Postdoctoral Fellow, Harvard Law School and Visiting Assistant Professor,
Villanova University School of Law; Private Law Fellow, Harvard Law School. We wish to
thank Janet Freilich, Meirav Furth, Oren Bar-Gill, John C.P. Goldberg, Patrick Goold, Kobi
Kastiel, Louis Kaplow, Steven Shavell, Yahli Shereshevsky, Henry E. Smith, Kathryn Spier,
and Gabriel Teninbaum for insightful comments and discussions. We are also thankful for
the suggestions of the participants at the Harvard John M. Olin Center Fellows Colloquium.
Generous financial support was provided by the Project on the Foundations of Private law at
Harvard Law School. Comments are welcome at yarbel@mail.law.harvard.edu.
Electronic copy available at: https://ssrn.com/abstract=2835482
8-Sep-16] Arbel & Kaplan - Draft 1
INTRODUCTION ........................................................................................................ 2
I. STRANGE BEDFELLOWS: OF ETHICISTS AND TORT REFORMERS ..................... 6
A. Apologies in Legal Scholarship .................................................................................. 6
B. Tort Reform .............................................................................................................. 12
C. How Tort Reformers Fought and Won the Apology Battle in State Legislatures ..... 15
II. COMMERCIAL APOLOGIES: THEORY AND PRACTICE ..................................... 20
A. A Theory of Apologies ............................................................................................. 20
1. The Goals of Tort Law and Apologies ................................................................ 20
2. A Unified Theory of Apologies in Tort Law ....................................................... 25
B. Commercial Apologies in Practice ........................................................................... 30
1. Delegation & Specialization................................................................................ 30
2. Professionalization & Training ........................................................................... 32
3. Diffusion of Responsibility ................................................................................. 33
4. Corporate Culture ................................................................................................ 34
C. Effectiveness of Commercial Apologies .................................................................. 36
III. CRITICAL ANALYSIS AND POLICY IMPLICATIONS ......................................... 39
A. Better Sorry than Safe .............................................................................................. 40
B. The Paradox of Excessive Apologies ....................................................................... 42
C. Apology as Disclosure .............................................................................................. 43
D. The Deficit of Apology Deficit ................................................................................ 45
E. Policy Implications ................................................................................................... 46
IV. CONCLUSION ........................................................................................... 48
V. APPENDIX: A MODEL OF LIABILITY FOR ACCIDENTS WITH APOLOGIES ....... 50
Electronic copy available at: https://ssrn.com/abstract=2835482
2 Draft, [8-Sep-16
INTRODUCTION
‘‘Capping malpractice payments . . . would do nothing to prevent unsafe practices or
ensure the provision of fair compensation to patients . . .
Apology offered by a health care provider during negotiations shall be kept confidential
and could not be used in any subsequent legal proceedings’’
--- Hillary Clinton & Barack Obama1
Why do large commercial interests-----insurance companies,
manufacturers, hospitals-----pledge millions of dollars to lobby for laws that
encourage apologies? What may explain this very recent interest of
commercial firms in the virtue of apologies? Why did tort reformers come
to adopt the rhetoric of regret, consilience, and penance? And how did the
largest tort reform of the last decades passed with the blessing of its
staunchest opponents?
Tort reform is a highly contentious social agenda. It is based on a
belief that litigation is inherently biased in favor of plaintiffs and must
therefore be reined-in by measures such as damage caps and screening
panels.2 Opponents of tort reform dispute this basic premise; they worry
that limitations on liability would unduly deprive accident victims of much-
needed compensation and would encourage negligent and reckless behavior.
The political pendulum slowly swings between these two positions.
In recent years, tort reformers have found a new and powerful
platform to advance their position, one that allowed them to strike a major
victory in their war against what they perceive as excessive liability.
Apology laws; laws designed to privilege apologies made by injurers,
making them inadmissible at trial. By co-opting the rhetoric and discourse
on apologies and the law-----independently developed by ethicists, dispute
resolution specialists, and legal theorists-----they found a way into the hearts
of legislators and the public. This maneuver has been so effective that even
long-standing opponents of tort reform, such as President Barack Obama,
express support for these reforms.3 In only two decades, 36 states have
adopted apology laws and there is currently a strong push to expand apology
1 Hillary Rodham Clinton & Barack Obama, Making Patient Safety the
Centerpiece, 354 N. ENG. J. MED. 2205, 2205 (2006).
2 See infra section I.B.
3 See infra section I.C.
Electronic copy available at: https://ssrn.com/abstract=2835482
8-Sep-16] Arbel & Kaplan - Draft 3
law reform to the federal level and other area of law.4
This Article argues and demonstrates that despite appearances,
apology laws are de-facto tort reform. Looking beyond the virtuous
rhetoric, the effect of apology laws on commercial actors is similar to that
of damages caps.5 At the heart of our argument is the overlooked claim that
apology laws undercut the deterrent effect of tort liability.6 We base our
argument on tort theory as well as research in psychology, economics,
sociology, and marketing. We contend that apology laws encourage
strategic apologies by commercial actors who do not express a real
commitment to avoid future wrongdoing. Commercial apologies exploit the
human tendency to forgive, which has myriad psychological, social, and
evolutionary reasons. For any of these reasons, victims forgive and settle
for a fraction of the value of their claims, foregoing hundreds of thousands
of dollars in compensation. Because commercial actors can anticipate in
advance that they will pay victims low amounts, they have less of an
incentive to invest in precautions that would prevent accidents in the first
place. In other words, apologies dilute deterrence, making it better to be
sorry than safe. We further argue that new market solutions and new trends
‘‘professionalize’’ and facilitate the tender of apologies by commercial
actors, thus greatly amplifying their harmful effects.7
4 See e.g., Dan M. Kahan & Eric A. Posner, Shaming White-Collar Criminals: A
Proposal for Reform of the Federal Sentencing Guidelines, 42 J. L. ECON. 365 (1999)
(calling to use shaming and apologies as a substitute to criminal sanctions); Chandler Farmer,
Striking a Balance: A Proposed Amendment to the Federal Rules of Evidence Excluding
Partial Apologies, 2 BELMONT L. REV. 243 (2015) (calling to create federal apology laws);
Lauren Gailey, “I’m Sorry” as Evidence? Why the Federal Rules of Evidence Should Include
a New Specialized Relevance Rule to Protect Physicians, 82 DEF. COUNS. J. 172 (2015);
Michael B. Runnels, Apologies All Around: Advocating Federal Protection for the Full
Apology in Civil Cases, 46 SAN DIEGO L. REV. 137 (2009). See also infra note 77 and
accompanying text.
5 Indeed, to the economist, apologies are a puzzle as ‘‘they must be regarded as
cheap talk’’ as ‘‘the only thing that is relevant is the expected magnitude of penalties.’’
Murat C. Mungan, Don’t Say You're Sorry Unless You Mean It: Pricing Apologies to
Achieve Credibility, 32 INT. REV. L. ECON. 178, 178 (2012). For this reason, Mungan
proposes that a special penalty will be levied on those who apologize. Id., at 179. Our
approach here is broader and we acknowledge the cost of delivering an apology, although
we believe that it is much smaller for commercial actors.
6 See infra section II.A.
7 We also explain that even if apologies are not merely strategic, i.e., they have a
real cost for the injurer, they may be socially undesirable. The reason, which we explain in
Electronic copy available at: https://ssrn.com/abstract=2835482
4 Draft, [8-Sep-16
The policy implications of our argument flow from understanding
the democratic gap inherent in these laws, as well as their potential harmful
implications for victims’ safety and welfare. Because our points here were
overlooked and avoided public scrutiny or scholarly analysis, we believe
that as a first measure, all planned future expansions of these laws-----to more
states, the federal level, and other areas of law-----should be suspended. The
effect of these laws on safety must be carefully evaluated, especially in the
context of medical malpractice where apologies are becoming
institutionalized and streamlined. Public discourse should internalize the
homomorphism of apology laws and tort reform and judge them
accordingly. Finally, judges should be made aware of the side effects of
apologies and learn to approach them with greater caution in commercial
settings.
Our argument explains, among other things, why we suddenly
witness deep interest from commercial actors in the virtues of apologies in
the context of private law.8 These reformers realized that by using the
uncontroversial rhetoric of apologies and penance they can mobilize
legislators from both parties. Hence, the support of apology laws by
commercial interest should not be viewed as a commendable fusion of social
and moral norms with business practices, but rather a self-interested
decision with potentially harmful social effects.
To provide a sense of the magnitude of the effect commercial
apologies have on victims, it is useful to consider the results of studies done
on payments to victims in states that enacted apology laws.9 These studies,
concentrating on hospitals, show a reduction of as much as 60% in payments
to victims. This translates to a reduction of $32,000-$65,000 in legal
payouts per case,10 which for many victims marks the difference between
greater detail in Section II.A.2, is that damages payments are transfers between individuals
which are, largely, socially neutral. Apologies, per this assumption, have real costs. While it
would be undesirable to replace a costless transfer with a costly action, injurers may
nonetheless do so, especially if they are encouraged by law.
8 We do not address in this paper the topic of public apologies or those made by
states, which raises distinct issues, see MARTHA MINOW, BETWEEN VENGEANCE AND
FORGIVENESS (1998); and Michael R. Marrus, Official Apologies and the Quest for
Historical Justice (Munk Centre, Occasional Paper III 2006).
9 See infra section II.C.
10 See Benjamin Ho & Elaine Liu, What’s an Apology Worth? Decomposing the
Effect of Apologies on Medical Malpractice Payments Using State Apology Laws, 8 J. EMPIR.
LEG. STUD. 179, 192 (2011) ($32,665); Benjamin Ho & Elaine Liu, Does Sorry Work? The
Electronic copy available at: https://ssrn.com/abstract=2835482
8-Sep-16] Arbel & Kaplan - Draft 5
being able to afford proper treatment for their accidents or suffering from
disability and poverty. For firms, on the other hand, the costs of apologies
are relatively marginal, and there is a large consensus that apologies are
cost-saving devices that can cut down costs by millions of dollars in
regulatory fines, judgments, and public outrage.11
The Article has three Parts. In Part I we explore the unexpected
camaraderie of ethicists and tort reformers. We show how the legal apology
movement was co-opted by the tort reform lobby to successfully effect tort
reform across the nation. Part II grounds apologies in tort theory and
explains how apologies can undermine deterrence in commercial settings.
Our theoretical analysis suggests that the problem is most acute if apologies
are cheap to produce and have a strong effect on victims. We then survey
recent developments in commercial apologies that show that commercial
apologies have indeed become cheaper and are highly effective. Part III
examines the theoretical and policy implications of these developments. We
argue that the evidence in support of apology law reform is weak and while
much empirical evidence is needed, the existing evidence is consistent with
the concern that apology laws undermine liability. After a brief conclusion,
an Appendix details our analysis using a formal economic model.
Impact of Apology Laws on Medical Malpractice, 43 J. RISK UNCERTAIN. 141 (2011)
($58,000-$73,000 for severe cases, $16,989-$24,017 for less severe cases, but -$3,132-$431
for insignificant cases, suggesting a potential increase in payouts for those cases). See also
Benjamin J. McMichael et al., Sorry is Never Enough: The Effect of State Apology Laws on
Medical Malpractice Liability Risk, manuscript (2016)
https://www.owen.vanderbilt.edu/faculty-and-research/upload/Apology-Paper-032516.pdf
($65,000).
11 See e.g., Erin O’hara O'Connor, Organizational Apologies: BP as a Case Study,
64 VAND. L. REV. 1957, 1977-1979 (2011) (discussing the role and effect of corporate
apologies).
Electronic copy available at: https://ssrn.com/abstract=2835482
6 Draft, [8-Sep-16
I. STRANGE BEDFELLOWS: OF ETHICISTS AND TORT REFORMERS
In recent decades, legal scholars from distinct disciplines-----ethicists,
dispute resolution experts, and sociologists-----have formed a movement that
challenged the traditional approach of the law to apologies. This Part tracks
the rise of this movement and its internal discourse. It then shows how the
rhetoric developed by this movement was co-opted by commercial interests
who lobbied for apology laws in state legislatures. These attempts were
immensely successful, and this part concludes by documenting the change
in the legal landscape.
A. Apologies in Legal Scholarship
In the early 90s, a movement of loosely formed ‘‘Legal Apologists’’
started to gain traction.12 The Legal Apologists critiqued the resolution of
conflict by the legal system for being overly abrasive to the relationship of
the parties. Instead, they argued that apologies can provide an effective and
wholesome solution to disputes. Despite the perception that apologies are
private and informal acts, they argued that the law has an important
facilitative role.13 In their view, the law should encourage individuals to
apologize or, at the very least, not stand in the way of those who wish to
apologize.
The Legal Apologists claimed that apologies have a wide-array of
benefits. When an individual is wronged, an apology by the responsible
party may acknowledge the harm done to the victim and the victim’s
12 See Hiroshi Wagatsuma & Arthur Rosett, The Implications of Apology: Law and
Culture in Japan and the United States, 20 L. & SOC'Y REV. 461, 487-88 (1986) (arguing
that incorporation of apologies into the American legal culture would reduce litigation and
repair relationships); Aviva Orenstein, Apology Excepted: Incorporating A Feminist
Analysis Into Evidence Policy Where You Would Least Expect It, 28 SW. U. L. REV. 221, 247
(1999) (advocating legal protection of apologies); Jonathan R. Cohen, Advising Clients to
Apologize, 72 S. CAL. L. REV. 1009 (1999). (Explaining the benefits to clients from
apologies). On the trend, see Aaron Lazare, The Healing Force of Apology in Medical
Malpractice and Beyond, 57 DEPAUL L. REV. 251, 251 (2008) (“Beginning in the early
1990s, there was a surge of academic and public interest in apologies.”).
13 See, e.g., Cohen, supra note 12, at 1011 (“Although a physician may wish to tell
a patient when he has made a mistake, lawyers often order doctors to say nothing.”); See also
Farmer, supra note 4, at 249 (calling apologies “legally dangerous”).
Electronic copy available at: https://ssrn.com/abstract=2835482
8-Sep-16] Arbel & Kaplan - Draft 7
agency;14 reduce feelings of anger and aggression by the victim;15 control
the attribution of fault to the responsible party;16 and start the process of
healing.17 As a consequence, apologies are said to mend the social fabric
14 See AARON LAZARE, ON APOLOGY 107 (2004) (considering acknowledgment of
harm as the foundation of an apology); Michael C. Jones, Can I Say I’m Sorry?: Examining
the Potential of an Apology Privilege in Criminal Law, 7 ARIZ. SUMMIT. L. REV. 563, 567
(2014) (by apologizing “[t]he offender acknowledges the harm he caused”).
15 See Erin A. O'Hara & Douglas Yarn, On Apology and Consilience, 77 WASH. L.
REV. 1121, 1124 (2002) (“In the face of a heartfelt apology, victims, . . . report feeling a near
instantaneous erosion of anger and pain.); Ken’ichi Ohbuchi et al., Apology as Aggression
Control: Its Role in Mediating Appraisal of and Response to Harm, 56 J. PERS. SOC.
PSYCHOL. 219 (1989) (testing empirically the effects of apologies on victim’s aggression,
finding soothing effects).
16 Psychologists find that apologies have a paradoxical effect. On the one hand,
apologies imply guilt and responsibility but on the other hand, experiments consistently find
that apologies reduce the attribution of fault to the wrongdoer and increase the belief that the
wrong happened for reason outside the wrongdoer’s control, See Bruce W. Darby & Barry
R. Schlenker, Children’s Reactions to Apologies, 43 J. PERS. SOC. PSYCHOL. 742, 745, 749
(1982) (finding that children attribute less responsibility to apologizing transgressors); and
Bernard Weiner et al., Public Confession and Forgiveness, 59 J. Pers. 281, 308 (1991)
(“confession alter perceptions of the confessor's moral character and causal attributions for
the negative action.”). On the paradox, see Jennifer K. Robbennolt, Apologies and
Reasonableness: Some Implications of Psychology for Torts, 59 DEPAUL L. REV. 489, 492
(2010).
17 See LAZARE, supra note 14, at 263 (listing the healing properties of apologies);
Margareth Etienne & Jennifer K. Robbennolt, Apologies and Plea Bargaining, 91
MARQUETTE L. REV. 295, 297 (2007) (arguing that victims of crimes find “emotional
restoration” and a “re-established sense of security” when receiving apologies). See also
Stephanos Bibas & Richard A. Bierschbach, Integrating Remorse and Apology into Criminal
Procedure, 114 YALE L.J. 85, 90 (2004) (apologies “heal offenders, victims, and
communities. Remorse and apology would teach offenders lessons, vindicate victims, and
encourage communities to welcome wrongdoers back into the fold”); Brent T. White, Say
You’re Sorry: Court-Ordered Apologies as a Civil Rights Remedy, 91 CORNELL L. REV.
1261, 1273-74 (2006); Deborah L. Levi, Note, The Role of Apology in Mediation, 72 N.Y.U.
L. REV. 1165, 1176-77 (1997) (arguing that apologies can be viewed as a form of
compensation as they heal part of the harm).
Electronic copy available at: https://ssrn.com/abstract=2835482
8 Draft, [8-Sep-16
torn by the transgression,18 restore prior relationships;19 and facilitate
negotiation.20 Importantly, the apology expresses a reestablished obligation
to refrain from future transgressions.21
For the Legal Apologists, all these advantages link to one
18 NICHOLAS TAVUCHIS, MEA CULPA: A SOCIOLOGY OF APOLOGY 13 (1991)
(“An apology thus speaks to an act that cannot be undone but that cannot go unnoticed
without compromising the current and future relationship of the parties, the legitimacy of the
violated rule, and the wider social web in which the participants are enmeshed.” He also
argues that apologies serve to reaffirm the victim’s membership in the community.); Barry
R. Schlenker & Bruce W. Darby, The Use of Apologies in Social Predicaments, 44(3) SOCIAL
PSYCHOLOGY QUARTERLY 271, 354 (1981) (noting that by apologizing the offender
“reaffirms the values of the rules that have been broken”); Erving Goffman, On Face-Work,
18 PSYCHIATRY 213, 220 (1955) (the apology is intended to “correct for the offense and
reestablish the expressive order”). See also Samul Oliner, Altruism, Apology, Forgiveness,
and Reconciliation as Public Sociology, in HANDBOOK OF PUBLIC SOCIOLOGY 375, 380
(“Through genuine apology and forgiveness, harmony may be restored”).
19 See Orenstein, supra note 12, at 241 (“apologies can transform individuals and
regenerate relationships.”). According to equity theory, individuals strive to a sense of equity
in their relationship which is disturbed by wrongdoing. The sense of imbalance is reported
to create anxiety, see generally Brad R.C. Kelln & John H. Ellard, An Equity Theory Analysis
of the Impact of Forgiveness and Retribution on Transgressor Compliance, 25 PERSONALITY
& SOC. PSYCHOL. BULL. 864 (1999). Apologies are found to restore the sense of equity by
demonstrating that the offender suffers too, See Robbennolt, supra note 16, at 492 and the
sources cited there. See also Kish Vinayagamoorthy, Apologies in the Marketplace, 33 PACE
L. REV. 1081, 1105 (2013) (arguing that the apology “reminds the transgressor of the value
of the relationship”) (citations omitted).
20 See Cohen, supra note 12, at 1020 (Indignity can be a large barrier to compromise,
and in many cases, an apology is needed”); Robin E. Ebert, Attorneys, Tell Your Clients to
Say They’re Sorry: Apologies in the Health Care Industry, 5 IND. HEAL. L. REV. 337, 339
(2015) (advocating apologies as a settlement strategy); Nancy L. Zisk, A Physician’s
Apology: An Argument Against Statutory Protection, 18 RICH. J.L. PUB. INT. 369, 390 (2015)
(arguing that because of the “powerful empirical data suggesting that physicians can reduce
their chances of being sued by communicating openly and honestly with their patients, . . . .
the conclusion seems inescapable that physicians must disclose mistakes and admit
responsibility for those mistakes.”). For a general discussion of research in emotion in
negotiations, see Max H. Bazerman et al., Negotiation, 51 ANN. REV. PSYCHOL. 279, 285-86
(2000).
21 See Gregg J. Gold & Bernard Weiner, Remorse, Confession, Group Identity and
Expectancies About Repeating Transgression, 22 BASIC APPL. PSYCHOL. 291 (2000);
Runnels, supra note 4, at 143-44 (“The apologetic offender will therefore be perceived as
less likely to engage in similar offending behavior in the future.”); Mihaela Mihai, Apology,
INTERNET ENCYCLOPEDIA OF PHILOSOPHY (2015) http://www.iep.utm.edu/apology/ (noting
that to be considered valid, the apology must imply an intention to refrain from similar
actions in the future).
Electronic copy available at: https://ssrn.com/abstract=2835482
8-Sep-16] Arbel & Kaplan - Draft 9
overarching theme: apologies facilitate dispute resolution in an effective
manner.22 By defusing victims’ desire for vindication,23 apologies avoid
disputes and encourage settlements, thus saving protracted legal proceedings
with their emotional and pecuniary costs.24
To demonstrate that these benefits are not merely theoretical, the
Legal Apologists have set to prove them empirically, mostly in lab settings.
The resulting studies have shown that victims of wrongful conduct report a
strong desire to receive an apology, express satisfaction once this need is
met, and, as a result, manifest a high willingness to settle and forego
litigation.25 A leading example is Jennifer Robbennolt’s work. In a series of
22 See Cohen, supra note 12, at 1061 (“encouraging apologies to occur early on may
prevent many injuries from escalating into legal disputes”); Farmer, supra note 4, at 244 (“A
sincere apology can help promote judicial economy by unlocking stalled settlement
negotiations . . . [and] can help ensure that impasse is avoided altogether.”); Ebert, supra
note 20, at 339 (noting that apologies can reduce litigation Jeffrey S. Helmreich, Does
“Sorry” Incriminate? Evidence, Harm and the Protection of Apology, 21 CORNELL J.L. PUB.
POL’Y 567, 567 (2012) (“Apology has proven a dramatically effective means of resolving
conflict and preventing litigation”);Orenstein, supra note 12, at 242 (“apologies can
substitute for costly litigation”); Zisk, supra note 20, at 390 (“In light of the powerful
empirical data suggesting that physicians can reduce their chances of being sued by
communicating openly and honestly with their patients, . . . . the conclusion seems
inescapable that physicians must disclose mistakes and admit responsibility for those
mistakes.”).
23 See supra note 17.
24 Steven Shavell and Mitchell Polinsky estimate that the costs of the legal system
absorb almost 50% of payments made by plaintiffs to defendants. Steven Shavell & A.
Mitchell Polinsky, The Uneasy Case for Product Liability, 123 HARV. L. REV. 1437, 1470
(2010) (“for each dollar that an accident victim receives in a settlement or judgment, it is
reasonable to assume that a dollar of legal and administrative expenses is incurred”).
25 See Thomas H. Gallagher et al., Patients' and Physicians' Attitudes Regarding
the Disclosure of Medical Errors, 289 JAMA 1001 (2003) (finding that patients expressed a
desire to receive an apology following a medical error); Gerald B. Hickson et al., Factors
That Prompted Families to File Medical Malpractice Claims Following Prenatal Injuries,
267 JAMA 1359, 1361 (1992) (noting that 24% of patients filed claims "when they realized
that physicians had failed to be completely honest with them about what happened, allowed
them to believe things that were not true, or intentionally misled them"); Marlynn L. May &
Daniel B. Stengel, Who Sues Their Doctors? How Patients Handle Medical Grievances, 24
L. & SOC’Y REV. 105 (1990) (finding that absence of apology motivates patients to bring
suit); Charles Vincent et al., Why Do People Sue Doctors? A Study of Patients and Relatives
Taking Legal Action, 343 LANCET 1609, 1612 (1994) (finding that 37% of respondents said
that they would not have sued had there been a full explanation and an apology and 14%
indicated that they would not have sued had there been an admission of negligence); Amy
B. Witman et al., How Do Patients Want Physicians to Handle Mistakes? A Survey of
Electronic copy available at: https://ssrn.com/abstract=2835482
10 Draft, [8-Sep-16
experimental studies, Robbennolt found that apologies increase victims’
belief that they would win their lawsuits, but, paradoxically, that they had
more favorable view of the injurer, were more willing to settle, and were
more receptive to lower settlement offers.26 Robbennolt also found that
victims who received an apology believed that the injurer is more likely to
be careful in the future.27
Armed with theory and evidence, the Legal Apologists quickly swept
legal academia. As others recently noted: “In the last two decades, apology
legal scholarship has become increasingly robust.”28 We found in our
analysis of the literature hundreds of articles on the issue, starting mostly in
the 90s and peaking in popularity in the 2000s.29
The ideas inspired by the movement quickly spread to other areas of
law, with apologies becoming the main item on the agenda for advocates of
Internal Medicine Patients in an Academic Setting, 156 ARCHIVES OF INTERNAL MED. 2565,
2566 (1996) (finding that 98% of respondents "desired or expected the physician's active
acknowledgement of an error.” And that "patients were significantly more likely to either
report or sue the physician when he or she failed to acknowledge the mistake."). See also
Nathalie Des Rosiers et al., Legal Compensation for Sexual Violence: Therapeutic
Consequences and Consequences for the Judicial System, 4 PSYCHOL. PUB. POL'Y & L. 433,
442 (1998) (survey of victims of sexual abuse that finds a desire for apologies); Piper Fogg,
Minnesota System Agrees to Pay $ 500,000 to Settle Pay-Bias Dispute, CHRON. HIGHER
EDUC., Feb. 14, 2003, at A12 (describing class-action plaintiff's disappointed reaction to the
settlement: "I want an apology," she said, "and I am never going to get it") (internal quotes
omitted); Editorial, The Paula Jones Settlement, WASH. POST, Nov. 15, 1998, at C6.
26 See Jennifer K. Robbennolt, Apologies and Settlement Levers, 3 J. EMPIRICAL
LEGAL STUD. 333 (2006); Jennifer K. Robbennolt, Apologies and Legal Settlement: An
Empirical Examination, 102 MICH. L. REV. 460, 462 (2003). See also Russell Korobkin &
Chris Guthrie, Psychological Barriers to Litigation Settlement: An Experimental Approach,
93 MICH. L. REV. 107, 148-50 (1994) (finding, but with low statistical significance, that
apologies affect willingness to settle).
27 See Robbennolt, supra note 16, at 506. For the effect of apologies outside the lab,
see infra section II.C..
28 Xuan-Thao Nguyen, Apologies as Intellectual Property Remedies: Lessons from
China, 44 CONN. L. REV. 883, 891 (2012) (“In the last two decades, apology legal scholarship
has become increasingly robust”).
29 Data acquired from Lexis Advance Search, Search terms: title(apolog*) OR
summary(apolog*) in title or summary of all Law Review and journal articles between 1984-
2015. A total of 326 results were found.
Electronic copy available at: https://ssrn.com/abstract=2835482
8-Sep-16] Arbel & Kaplan - Draft 11
‘‘restorative justice’’,30 ‘‘therapeutic jurisprudence’’,31 and alternative
dispute resolution, with special emphasis on mediation.32 Apologies were
offered as a mean of reforming diverse areas of law, such as criminal law,33
medical malpractice,34 tort law,35 and intellectual property.36 It was even
suggested that part of the Federal Register (‘‘probably one of the driest
publications ever printed’’) would include a section for governmental
apologies.37
This account of the literature will not be complete without
mentioning the internal divisions within the Legal Apologists. The most
common objections are that providing legal protection to apologies would
negate their moral value,38 that people would fake apologies and courts
would be ill-positioned to verify their authenticity,39 or that frequent
30 See Bibas & Bierschbach, supra note 17, at 103 (“Restorativists consider apology
and remorse important as part of a holistic process”); Alana Saulnier & Diane
Sivasubramaniam, Effects of Victim Presence and Coercion in Restorative Justice : An
Experimental Paradigm, 39 L. HUM. BEHAV. 378, 379 (2015). (“apology is central to
restorative justice”).
31 See Jones, supra note 14, at 565-68 (2014) (surveying therapeutic justice and
apologies in criminal law); See also Susan Daicoff, Apology, Forgiveness, Reconciliation &
Therapeutic Jurisprudence, 13 PEPP. DISP. RESOL. L.J. 131, 153-57 (2013) (surveying the
field of therapeutic jurisprudence).
32 See Bibas & Bierschbach, supra note 17, at 130-35 (advocating greater role for
mediation in criminal settings because it encourages apologies and remorse); Angela M.
Eastman, The Power of Apology and Forgiveness, 36 VT. B.J. 55 (2014). (Discussing the
effectiveness of apologies in dispute resolution); Levi, supra note 17, at 1165.
33 See generally Bibas & Bierschbach, supra note 17 (calling for a fuller integration
of apologies and expressions of regret into criminal procedure).
34 See Gailey, supra note 4, at 177-78.
35 See e.g., Daniel W. Shuman, The Role of Apology in Tort Law, 83 JUDICATURE
180 (2000).
36 See Nguyen, supra note 28.
37 See Eugene R. Fidell, Sorry, 71 Fed. Reg. 1 (2006), 8 GREEN BAG 155, 156
(2005).
38 See, e.g., TAVUCHIS, supra note 18, at 34 (explaining that the potential for
negative repercussions is an essential part of apologies); Lee Taft, Apology Subverted: The
Commodification of Apology, 109 YALE L. J. 1135, 1142 (2000) (arguing that the morality
of apologies derive from the exposure of the apologizing party to the consequences of the
wrongful act). Interestingly, victims may also abuse apologies by refusing to accept
apologies in order to use them as a basis for a lawsuit, see O'Hara & Yarn, supra note 15.
For a critique stating that apologies are helpful even when they do not admit blame, see
Helmreich, supra note 22, at 609.
39 On strategic apologies, see Ebert, supra note 20, at 364 (“a wrongdoer might
Electronic copy available at: https://ssrn.com/abstract=2835482
12 Draft, [8-Sep-16
apologies would lead victims to accept settlements that do not compensate
them fully.40 Despite these challenges, the movement itself is still going
strong, seemingly in the belief that none of these challenges is
insurmountable-----which, as we will argue, is most understandable if the
literature is read as focusing on interpersonal apologies.
B. Tort Reform
Moving from the high-minded Legal Apologists and their concern
with nuances of ethics, we consider the seemingly unrelated world of tort
reform. Tort reformers, known mostly for their activism in medical
malpractice and product liability, fight to limit the costs imposed on
defendants as a result of litigation, which they believe is excessive and
biased. They argue that the specter of excessive liability affects the industry
and especially physicians who are pressured to engage in so-called
apologize for the wrong reasons”); O'Hara & Yarn, supra note 15, at 1186 (“[A]pology can
be used as a tool for organizations to strategically take advantage of individual victims'
instincts to forgive in the face of apology.”); and Daniel Eisenberg, When Doctors Say,
“We're Sorry,” TIME, Aug. 15, 2007, at 50 (observing that many believe that “[a]pology laws
. . . could just usher in an epidemic of playacting.”). In one case, for example, a defendant
who was ordered by the court to apologize published an ad in the newspaper—later on the
same day—saying he was not really sorry. Amanda Garrett, Apologize or Go to Jail, Judge
Orders Criminals to Say, ‘I'm Sorry,’ to Victims, PLAIN DEALER, Oct. 9, 1999, at 1B. But
see Cohen, supra note 12, at 1065-66 (assuaging the concern that lawyers will advise clients
to strategically apologize because of their ethical obligations). Others believe that even
strategic apologies serve a useful social function. See Kahan & Posner, supra note 4
(advocating apologies as a shaming sanction); Orenstein, supra note 12, at 223 (“Even
apologies that originate from self-protection, which are not entirely sincere or fully contrite,
serve a vital social purpose.”). On courts’ ability, See Bibas & Bierschbach, supra note 17,
Jeffrie G. Murphy, Well Excuse Me! -- Remorse, Apology, and Criminal Sentencing, 38 ARIZ.
STATE L. J. 371, 376 (2006). (“[expressions of remorse] are matters about which the state is
probably incompetent to judge--it cannot even deliver the mail very efficiently”); Michael
M. O’Hear, Remorse, Cooperation, and “Acceptance of Responsibility”: The Structure,
Implementation, and Reform of Section 3E1.1 of the Federal Sentencing Guidelines, 91 NW.
U. L. REV. 1507, 1564 (1997) (expressing skepticism of courts’ ability to detect dishonest
apologies).
40 See Levi, supra note 17, at 1171 (“For instance, critics might ask, if a plaintiff
settles because she's emotionally fulfilled by an apology, isn't she being duped out of her
legal entitlement --an entitlement that the apology itself makes concrete?”); Gabriel H.
Teninbaum, How Medical Apology Programs Harm Patients, 15 CHAPMAN L. REV. 307, 309
(2011) ("modern apology programs appear to cool their marks out as a means of preventing
them from speaking to a lawyer and becoming educated about their legal rights.")
Electronic copy available at: https://ssrn.com/abstract=2835482
8-Sep-16] Arbel & Kaplan - Draft 13
‘‘defensive medicine’’, i.e., prescribing tests and procedures for the sole
purpose of reducing liability risk.41 Both the costs of liability and those of
defensive medicine are then passed on to the public in the form of higher
health costs (or, in other fields, in the form of higher costs of products and
services). To contain these costs, tort reformers suggest a series of methods
that would curb the threat of excessive liability, such as damages caps.
Opponents challenge these ideas, arguing that there is no evidence that
liability is excessive, that defensive medicine is prevalent, or that tort
reforms have any positive effects on the costs or quality of healthcare.42
To be clear-----and clarity is often lacking in this debate-----tort reform
is not about making the tort system more efficient.43 Both reformers and
their opponents are open to making the system work better at a lower cost.44
The focal point of contention is tort reform’s objective to reduce the
deterrent effect of tort liability. Tort reformers believe that damages in
litigation are too high and so overly-deter potential injurers, such as
physicians, which is the cause of ‘defensive medicine’ practices. Therefore,
their call is to cap money damages as means of curbing the over-deterrent
effect of litigation.45
41 David M. Studdert et al., Defensive Medicine Among High-Risk Specialist
Physicians in a Volatile Malpractice Environment, 293 JAMA 2609 (Finding in a survey of
824 physicians that 93% practice defensive medicine).
42 See, e.g., Myungho Paik et al., Will Tort Reform Bend the Cost Curve? Evidence
from Texas, 9 J. EMPIR. LEG. STUD. 173, 176-81, 209-11 (2012) (reviewing the literature and
conducting an empirical analysis of the effect on costs).
43 See generally Carl T Bogus, Syposium: Introduction: Genuine Tort Reform, 13
ROGER WILLIAMS U. L. REV. 1 (2008) (tracking the history of the tort reform movement and
noting the specific political meaning of the term). See also Rachel M. Janutis, The Struggle
Over Tort Reform and the Overlooked Legacy of the Progressives, 39 AKRON L. REV. 943
(2006).
44 For example, the leading Democratic legislation of the past decade, § 6801 to the
Patient Protection and Affordable Care Act (“Obamacare”), explicitly endorses efficiency
oriented reforms to tort law (“develop and test alternatives to the existing civil litigation
system as a way of improving patient safety, reducing medical errors, encouraging the
efficient resolution of disputes, increasing the availability
of prompt and fair resolution of disputes, and improving access
to liability insurance, while preserving an individual’s right to
seek redress in court.”).
45 See, e.g., Michael P Allen, A Survey and Some Commentary on Federal “Tort
Reform”, 39 AKRON L. REV. 909, 910 (2006) (“arguments about tort reform are really
arguments about restricting tort recoveries in one form or another.” Incidentally, the author
nonetheless uses a more expansive definition); Rachel M Janutis, The Struggle Over Tort
Electronic copy available at: https://ssrn.com/abstract=2835482
14 Draft, [8-Sep-16
In terms of political economy, the tort-reform debate pits consumers
and trial attorneys against professional, commercial, and business
interests.46 These opposing camps have mapped into political parties, with
Republicans being strong proponents of tort reform against the opposition
of Democrats, a somewhat ironic division in light of the history of tort law.47
Most notably, President George W. Bush has strongly favored tort reform
at the Federal level, calling to cap all money damages at $250,000,48 while
President Barack Obama has been largely opposed to damage caps.49
Tort reform has marked significant success. According to data
collected by Ronen Avraham in 2012, 21 states have a cap on non-economic
Reform and the Overlooked Legacy of the Progressives, 39 AKRON L. REV. 943, 944 (2006)
(explaining that tort reformers seek to “limit[] the availability of relief and the amount of
relief in personal injury actions”); Geoff Boehm, Debunking Medical Malpractice Myths :
Unraveling the False Premises Behind “Tort Reform”, 5 YALE J. HEAL. POL’Y, L. ETHICS
357, 358 (2005) (explaining tort reform as an attempt to limit victims’ rights through caps
on damages). Another pillar of tort reform is the screening of frivolous lawsuits, which tort
reformers believe are common. This would seem to be an attempt to make the system more
efficient, but opponents view this measure as an attempt to curb all litigation, regardless of
merit. See David A. Hyman & Charles Silver, Medical Malpractice Litigation and Tort
Reform: It’s the Incentives, Stupid, 59 VAND. L. REV. 1085, 1086-87 (2006) (arguing that
the true intent of tort reformers in this area is to make “the system less remunerative”).
46 See Todd J. Zywicki, Public Choice and Tort Reform. LAW AND ECONOMICS
WORKING PAPER (2000) (arguing that lawyers are pushing for expansion of tort liability),
Paul H. Rubin, Public Choice and Tort Reform, 124 PUBLIC CHOICE 223, 230 (2005)
(describing the tension between the different groups). See also Rachel M Janutis, The
Struggle Over Tort Reform and the Overlooked Legacy of the Progressives, 39 AKRON L.
REV. 943, 945-46 (2006).
47 See Stephen D. Sugarman, Ideological Flip-Flop : American Liberals Are Now
the Primary Supporters of Tort Law, in ESSAYS ON TORT, INSURANCE, LAW AND SOCIETY IN
HONOUR OF BILL W. DUFWA 1105 (Jure Forlag ed., 2006) (Identifying tort law with
conservative values and suggesting that the Democratic support of the tort system is a recent
one), Paul H. Rubin, Public Choice and Tort Reform, 124 PUBLIC CHOICE 223, 230-31
(2005) (explaining the mapping of these interests in partisan terms).
48 George W. Bush, Remarks at the University of Scranton in Scranton
Pennsylvania, January 16, 2003 (“for the sake of affordable and accessible health care in
America, we must have a limit on what they call non-economic damages. And I propose a
cap of $250,000”).
49 CBS, 60 Minutes, September, 11, 2009 (“"What I would be willing to do is to
consider any ideas out there that would actually work . . . [damages] caps will not do that.”).
In this interview, President Obama clarified a statement he gave to the congress,
acknowledging the potential importance of defensive medicine, see White House, Remarks
by the President to a Joint Session of Congress on Health Care, September, 9th, 2009.
Electronic copy available at: https://ssrn.com/abstract=2835482
8-Sep-16] Arbel & Kaplan - Draft 15
damages, 18 on punitive damages, and 22 on total compensation.50 Matter
and Stutzer recently found that Republican leadership in a state leads to a
large jump in the probability that tort reform will be undertaken.51
According to our own analysis, states that voted Republican in the 2012
election were far more likely to have some damages caps than those that
voted Democratic. Specifically, out of 24 Republican states, 19 had caps,
whereas out of 26 Democratic States, only 16 had caps.
Tort reform has made considerable inroads, but it also faces strong
opposition. First, politically, as we noted Democratic States are traditionally
averse to tort reform. Second, consumer and attorney lobby mounts a strong
opposition. And third, various courts have held damages caps
unconstitutional, mostly due to concerns of their limiting effect on the right
to a trial by jury.52 These challenges limit the ability of tort reformers to
push forward. The difficulty of advancing their agenda through ‘‘the front
door’’ has put pressure on reformers to find alternative venues for progress,
ones that could sidestep the political, interest-group, and legal obstacles.
Realizing this, reformers formed a new alliance with unlikely partners --- the
Legal Apologists.
C. How Tort Reformers Fought and Won the Apology Battle in State
Legislatures
Much to the envy of legal scholars everywhere, the Legal Apologists
have had a tremendous impact on policy. These ethicists and dispute
resolution specialists found surprising support from the pragmatic and well-
funded tort reform advocates.53 With the rhetoric of the legal apologists and
50 Ronen Avarham, Database of State Tort Law Reform, Version 5.1 (Clever),
https://law.utexas.edu/faculty/ravraham/dstlr.php.
51 See Ulrich Matter & Alois Stutzer, The Role of Party Politics in Medical
Malpractice Tort Reforms, 42 EUR. J. POLIT. ECON. 17 (2016),
52 See Bryan J Chase et al., Are Non-Economic Caps Constitutional?, 1 DEF.
COUNS. J. 154 (2015) (reviewing the judicial battle of the constitutionality of non-economic
damages caps).
53 NICK SMITH, JUSTICE THROUGH APOLOGIES: REMORSE, REFORM, AND
PUNISHMENT 283 (2015) (Arguing that “Tort reformers often bundle apology legislation
within” other tort reform measures”); Cohen, supra note 12, at 856 (2002) (suggesting that
apology laws depend on support by “insurance companies, medical associations and Fortune
500 companies”); Teninbaum, supra note 40, at 311 (“industry lobbyists exerted influence
on lawmakers to create special medical apology shield laws”); See also Runnels, supra note
4, at 484-85 (2011) (noting the lobby efforts of Sorry Works!, a coalition of physicians,
Electronic copy available at: https://ssrn.com/abstract=2835482
16 Draft, [8-Sep-16
the lobby efforts of tort reformers, the movement struck a chord with
legislators and judges across the country, prompting them to reform the law
to accommodate the use of apologies.
The same supporters of tort reform back apology laws: insurance
companies, medical associations, and large companies in diverse
industries.54 However, they do so using a new rhetoric, clearly
differentiating between apology laws and tort reforms.55 In Madison,
Wisconsin, for example, ‘‘[t]he medical lobby, supported by powerful
business groups, outmaneuvered trial lawyers . . . and won passage of the
‘I'm sorry" bill’.’’56 This lobby adopted a new rhetoric, arguing that apology
laws are useful not because they curb liability but because ‘‘at these difficult
times, people want, need and deserve compassion.’’57 Similarly, in
Massachusetts, various healthcare organizations lobbied for apology laws
explaining this as a move towards ‘‘a very proactive system where
physicians can advocate for patients who are injured rather than being told
they can't even talk to them.’’58
Tort reformers borrowed from Legal Apologists both the means and
the rhetoric to advance their goals. The most important item on the agenda
of reformers was the creation of ‘‘safe harbor’’ for apologies.59 Apologies
insurers, hospital administrators and patients).
54 Id.
55 Doug Wojcieszak et al., The Sorry Works! Coalition: Making the Case for Full
Disclosure., 32 JT. COMM’N J. QUAL. PATIENT SAF. 344, 344 (2006) (portraying apology
laws as a “middle-ground” approach to the “medical malpractice crisis”).
56 Cary Spivak & Kevin Crowe, ‘I’m Sorry’ Bill Latest Example of Doctors’ Clout,
June 28, 2014, JOURNAL SENTINEL.
57 Patrick Marley & Jason Stein, “Senate Passes Chemotherapy, Cannabis Oil Bills,
(Apr. 1, 2014), Journal Sentinel (quoting Sen. Leah Vukmir (R-Wauwatosa)).
58 Massachusetts Medical Society, Disclosure, Apology and Offer: A New
Approach to Medical Liability (June, 2012) http://www.massmed.org/News-and-
Publications/Vital-Signs/Back-Issues/Disclosure,-Apology-and-Offer--A-New-Approach-
to-Medical-Liability/.
59 See Peter H. Rehm & Denise R. Beatty, Legal Consequences of Apologizing, 1996
J. DISP. RESOL. 115, 128-29 (1996) (providing early support to apology safe harbor laws),
Orenstein, supra note 12 (calling for safe harbor laws). The nation’s apology laws originate
in Massachusetts. MASS. GEN. LAWS ANN. CH. 233, § 23D (West Supp. 1998). A retired
legislator’s daughter was hit by a car but the driver refused to apologize because of fear of
legal liability. This led to the adoption of the first apology law. See Taft, supra note 38, 1051-
52 (2000).
Electronic copy available at: https://ssrn.com/abstract=2835482
8-Sep-16] Arbel & Kaplan - Draft 17
often convey evidence of fault and are therefore admissible at trial.60
Reformers adopted the discourse and rhetoric of the Legal Apologists, who
argued that it would be wrong to punish people who ‘‘did the right thing’’
and apologized.61 The Legal Apologists further argued that existing
evidentiary rules make defendants fear apologies are ‘‘legal suicide’’62 and
provide an undue and unfair barrier to injurers from apologizing.63
The second item on the reformers’ agenda was the promotion of
apologies in less formal settings. Both reformers and legal apologists sought
to promote-----for very different reasons, of course-----the role of apologies in
mediation,64 settlement procedures,65 and the early stages of trial.66 Finally,
the third item was more institutional-----providing judges with the power to
mandate apologies as an additional or substitute aspect of sanctions.67
60 FED. R. EVID. 801(d)(2). Federal law only protects apologies if they are made
during settlement negotiations. See FED R. EVID. 408; Cohen, supra note 12, at 1032-36
(1999). An apology might also be inadmissible if it is implied from an offer to cover medical
expenses, FED R. EVID. 409. The rationale for this rule is that “such payment or offer [to pay
the victim's medical expenses] is usually made from humane impulses and not from an
admission of liability, and that to hold otherwise would tend to discourage assistance to the
injured person.” FED. R. EVID. 409 advisory committee's note.
61 See Cohen, supra note 12, at 864 ("The law should not punish people for taking
a moral step"); Orenstein, supra note 12, at 235-36 (“[A] justification for [these rules] arises
from a desire to reward goodness . . . . We do not want to punish the ‘blessed peacemakers[.]’
We certainly do not want to disadvantage individuals who do the right thing.”).
62 Eisenberg, supra note 39.
63 See Robbennolt, supra note 26, at 465 (“The conventional wisdom among legal
actors has been that an apology will be viewed as an admission of responsibility and will
lead to increased legal liability”, although she also notes that there is no empirical research
to supports this perception). Cohen, supra note 12, at 1010 (1999) (“If a lawyer contemplates
an apology, it may well be with a skeptical eye: Don't risk apology, it will just create
liability”).
64 See Levi, supra note 17.
65 See Elizabeth Latif, Apologetic Justice: Evaluating Apologies Tailored toward
Legal Solutions, 81 BOST. UNIV. L. REV. 289, 292 (2001).
66 See Etienne &. Robbennolt, supra note 17, at 299 (“encouraging apologies in
earlier stages of the criminal law process may be a laudable goal”); Bibas & Bierschbach,
supra note 17, at 128-29 (advocating that the tender of an apology would lead to lenient
charges, forego arrests, and deferral of prosecutions).
67 Cfr. White, supra note 17, at 1297 (“Requiring unrepentant officials to endure a
small amount of psychological discomfort [by coerced apologies] is a small price to pay to
help injured individuals”), and Latif, supra note 65, at 311 (forced apologies “can mitigate
anger, shame or educate the offender, or improve prospects for settlements”); Sharon
Elizabeth Rush, The Heart of Equal Protection: Education and Race, 23 N.Y.U. REV. L. &
Electronic copy available at: https://ssrn.com/abstract=2835482
18 Draft, [8-Sep-16
Reformers have been extremely successful, conquering 36 state
legislatures in only a decade.68 Additionally, courts have seemed to
internalize apology norms.69 Some courts are said to apply these norms
‘‘with gusto’’,70 leading them to treat apologies as valid grounds for
mitigating money damages,71 lowering sentencing,72 and exempting legal
SOC. CHANGE 1, 50-57 (1997) (advocating an equitable remedy of apologies in civil
litigation) with Pennsylvania Human Relations Comm'n v. Alto-Reste Park Cemetery Ass'n,
306 A.2d 881, 891 (1973) (Justice Pomeroy concurring) (“An apology is a communication
of the emotion of remorse for one's past acts. To order up that particular emotion, or any
other emotion, is beyond the reach of any government”. ); Levi, supra note 17, at 1178 (1997)
(arguing that involuntary apology is “just talk”), Nick Smith, Against Court-Ordered
Apologies, 16 NEW CRIM. L. REV 1 (2013) (arguing that court-ordered apologies serve little
function).
68 See Latif, supra note 65, 301 (reporting on California, Massachusetts, and Texas
in 2001). Compared with 36 states today, see EBS CONSULTING, Apology Protection Laws
in 36 States Letting Physicians be Human Again, Aug. 18, 2015, http://blog.ebs-
consulting.com/apology-protection-laws-in-36-states-letting-physicians-be-human-again.
See also Zisk, supra note 20, at 390 at 375 and n. 43; Ebert, supra note 20, at 366. The most
prevalent form of apology laws is safe-harbor to expressions of sympathy and empathy (e.g.,
“I am sorry you were hurt”), see, e.g., MONT. CODE. ANN. §26-1-814 (providing safe harbor
for a statement “expressing apology, sympathy, commiseration, condolence, compassion, or
a general sense of benevolence relating to the pain, suffering, or death of a person.”),
although several states provide a more robust protection and make inadmissible even
liability-assuming apologies (e.g., “I am sorry I hurt you through my negligence“).
69 Judges show reluctance to allow an apologetic admission of guilt to be the sole
basis for establishing the breach of a duty of care. In the medical context, see Ebert, supra
note 20, at 349 (“the use of apologies and other extrajudicial statements made by the
physician following a medical error are not alone sufficient to prove negligence.”). See also
Lashley v. Koeber, 156 P.2d 441, 445 (Cal. 1945) (physician’s admission that the mistake is
“all my own” is “insufficient to establish negligence”), Phinney v. Vinson, 605 A.2d 849,
850 (Vt. 1992) (Doctor’s apology is insufficient to establish a breach of standard of care).
But see Senesac v. Assocs. in Obstetrics & Gynecology, 449 A.2d 900, 901 (Vt. 1982) (“It
is conceivable that in some circumstances the extrajudicial admission of a defendant
physician could establish a prima facie case of negligence”). For more examples see Dan M.
Kahan, What Do Alternative Sanctions Mean?, 63 U. CHI. L. REV. 591, 634 & nn.171-72
(1996) (citing various examples of court-ordered apologies).
70 White, supra note 17, at 1268-69. See also Latif, supra note 65, at 296-98.
71 See, e.g., Groppi v. Leslie, 404 U.S. 496, 506 n.11 (1972) (providing mitigated
penalties for contempt due to an apology); Johnson v. Smith, 890 F. Supp. 726, 729 n.6 (N.D.
Ill. 1995) (apology mitigated punitive damages). See also Peter H. Rehm & Denise R.
Beatty, Legal Consequences of Apologizing, 1996 J. DISP. RESOL. 115 (1996) (reviewing
the legal effects of an apology).
72 See U.S. Sentencing Guidelines Manual § 3E1.1, cmt. n.3 (2003) (providing a
sentence reduction of two to three levels for clear demonstrations of acceptance of
Electronic copy available at: https://ssrn.com/abstract=2835482
8-Sep-16] Arbel & Kaplan - Draft 19
liability for crimes.73
This fast adoption amazed many: ‘‘Shortly after the idea of excluding
apologies from admissibility into evidence was raised in academic circles .
. . it rapidly spread to the policy arena’’.74 Yet, this success has not satiated
reformers’ appetite; they now seek to expand the scope of apology laws,75
apply them to other areas of civil and criminal law,76 enact them at the
federal level,77 and make them more uniform.78 Additionally, some
advocate that judges be able to compel the government to apologize in civil
rights cases.79
Tort reformers managed an impressive feat. On the one hand, they
draw on the resources and financial support of business interests that invest
hundreds of millions of dollars each year to advance tort reform.80 On the
other hand, they garnered large, bipartisan support. They even swayed
consumer advocates and lawyers which were willing to withdraw their
traditional opposition to tort reform in this context.81 Perhaps more
responsibility); Bibas & Bierschbach, supra note 17, at 92-95 (showing how criminal law
positively accounts for apologies in sentencing).
73 See, e.g., Kahan & Posner, supra note 4 (reporting of a judge substituting a 10
year punishment for embezzlement with an apology).
74 Cohen, supra note 12, at 819. See also Gailey, supra note 4, at 178-81 (surveying
the development of state apology laws).
75 See, e.g., Matthew Pillsbury, Say Sorry and Save: A Practical Argument for a
Greater Role for Apologies in Medical Malpractice Law, 1 S. NEW ENG. ROUNDTABLE
SYMP. L.J. 171, 200 (2006) (“As for situations where apologies are admissible, courts and
lawmakers across the country can learn from the strides made by their counterparts in other
states [where apologies are protected]”).
76 See, e.g., Jones, supra note 14, at 580-81(advocating an ‘apology privilege’ that
would create a safe harbor for apologies in criminal proceedings).
77 See supra note 4. See also Cohen, supra note 12, at 1061; Helmreich, supra note
22.
78 See Zisk, supra note 20, at 377-78 (noting that in Iowa, chiropractors are
protected, but chefs are not); See also Iowa Code § 622.31 (protecting only licensed
professionals).
79 See White, supra note 17 (advocating court-coerced apologies as a civil right
remedy).
80 Ronen Avraham, An Empirical Study of the Impact of Tort Reforms on Medical
Malpractice Settlement Payments, 36 J. LEGAL STUD. 183, 184 (2007).
81 PUBLIC CITIZEN, MEDICAL MALPRACTICE BRIEFING BOOK 85 (2004) (suggesting
apology laws as an alternative to tort reform); Pennsylvania Governor Signs Benevolent
Gesture Medical Professional Liability Act (Oct. 25, 2013), CLAIMS JOURNAL (reporting that
after a decade of back and forth battle between doctors and lawyers “the two professions
recently changed lobbying tactics by mutually agreeing on a new reform that both sides say
Electronic copy available at: https://ssrn.com/abstract=2835482
20 Draft, [8-Sep-16
remarkable is that despite his known opposition to tort reform, Barack
Obama co-signed a bill he authored with Hillary Clinton that sought to
establish federal apology safe harbors.82 Democratic lawmakers seem as
keen to adopt apology laws as Republican lawmakers, as evident by wide
adoption in both blue and red states.83
In sum, apology laws are promoted using the rhetoric of virtue,
improved communications, and ethics developed by legal intellectuals.
What is never explicitly noted, let alone considered, is the broader effects
of apology laws on incentives, harms, and other social costs. These issues
are simply suppressed and instead apology laws are framed as a neutral
measure that improves dispute resolution without sacrificing victims’ rights.
The acceptance of these laws by those who traditionally oppose tort reform
thus presents somewhat of a paradox. It is our task now to show why
apology laws undercut deterrence and are thus, in effect, comparable to
other measures of tort reform.
II. COMMERCIAL APOLOGIES: THEORY AND PRACTICE
We have seen that tort-reformers have joined hands with legal
scholars and have managed to change the law in most states. This Part first
provides a theoretical framework that allows evaluating the effect of
apologies on behavior. This theory highlights the importance of the cost of
apologies and their effectiveness; given that, this Part considers separately
both issues, showing how the costs of commercial apologies are declining
while their effectiveness remains considerable.
A. A Theory of Apologies
1. The Goals of Tort Law and Apologies
The two primary goals of tort law are compensation of victims and
will help.”)
82 See S.1784: National MEDiC Act, 109th Congress, 2005-2006
https://www.govtrack.us/congress/bills/109/s1784; Runnels, supra note 4, at 156 (discussing
the bill), Clinton & Obama, supra note 1, at 2206 (discussing the bill).
83 See Benjamin Ho & Elaine Liu, Does Sorry Work? The Impact of Apology Laws
on Medical Malpractice, 43 J. RISK UNCERTAIN. 141, 144 note 5 (2011) (noting that
regression analysis shows that “political composition in the State Senate and State House has
no significant explanatory power on the passage of apology laws.” They also find that
apology laws are not correlated with other tort reforms.).
Electronic copy available at: https://ssrn.com/abstract=2835482
8-Sep-16] Arbel & Kaplan - Draft 21
deterrence of wrongdoers; an important secondary goal is the reduction of
litigation costs.84
Controlling the costs of disputes, and litigation in particular, has
been the prominent theme in the writing of the Legal Apologists.85 They
argued that apologies help reduce litigation through the dissipation of
victim’s anger or need for vengeance. For the most part, the evidence seems
to support this assertion, although recent empirical work casts some doubt.86
From a cost-perspective, then, apologies seem to have a positive effect. If
one approaches a case after a dispute had already arisen, it may be appealing
to focus on controlling its costs. And indeed, most of the Legal Apologists
are conflict resolution experts who meet disputes after they arise.87 But tort
law adopts a broader perspective, and in this view, controlling litigation
costs is generally a secondary consideration.
What has been missing from the Legal Apologists’ analysis is the
effect of apologies on deterrence. In a fundamental oversight, the Legal
Apologists have failed to account for this central goal of tort law. Thus,
they have never accounted for the ex-ante effects of apologies on primary
behavior, namely, how does the possibility of apologizing after the fact
affect injurers’ decisions to engage in harmful activities in the first place?
How do apologies change the level of behavior? Would a more favorable
treatment of apologies by the legal system induce or suppress accidents?
Once considered, reflection reveals a tension between apologies and
deterrence. To the extent that apologies reduce the cost of an accident for
the injurer-----which is the point just discussed-----they provide the injurer with
less of a reason to avoid the accident. Put differently, if apologies allow the
injurer to limit exposure to liability, then the injurer has-----all other things
84 See generally STEVEN SHAVELL, FOUNDATIONS OF ECONOMIC ANALYSIS OF
LAW 192-93 (2004); Mark A Geistfeld, Compensation as a Tort Norm, in PHILOSOPHICAL
FOUNDATIONS OF THE LAW OF TORTS 65 (2013) (advancing compensation as the central
goal of tort law); Guido Calabresi, Some Thoughts on Risk Distribution and the Law of
Torts, 70 YALE L.J. 499 (1961) (loss spreading). See also Steven Shavell & A. Mitchell
Polinsky, The Uneasy Case for Product Liability, 123 HARV. L. REV. 1437 (2010)
(critically assessing the effectiveness of tort law in meeting these goals in the context of
product liability)
85 See supra note 22.
86 See infra II.C.
87 For example, Deborah L. Levi is a practicing mediator; Jennifer K. Robbennolt
is an expert on dispute resolution; and Robyn Carroll is an expert on dispute resolution and
mediation.
Electronic copy available at: https://ssrn.com/abstract=2835482
22 Draft, [8-Sep-16
being equal-----much lesser incentive to avoid the activity or to invest in
precautions. This does not mean that the injurer will not care at all or that
the effect of apologies is necessarily negative, but it does imply that injurers
will have less incentive to take care than they would otherwise.88
The other primary goal, compensation, fails to provide clear
guidance. The goal of compensating a victim is to restore her to her status
prior to the accident, by providing her with value that is as close as possible
to her loss.89 On first blush, apologies seem to undercut this goal, because-----
as demonstrated by the apologists themselves-----victims are willing to accept
lower payments in settlements when an apology is tendered.90 However, the
Legal Apologists emphasized that victims care for much more than financial
compensation and the positive emotional and expressive effect of an apology
may well be more important than the payment of money. The contention
here is that apologies are healing and valuable to the victim more than a
monetary payment. Let us call this view the therapeutic value theory of
apologies.
There are strong reasons to be skeptical of the therapeutic value
theory, particularly in the context of commercial apologies. While the
adherents of the therapeutic value theory argue that victims’ acceptance of
apologies is evidence of their value, there are several alternative, less
benign, reasons why victims might accept them and forgo sometimes
hundreds of thousands in compensation.91 We cover five reasons here.
The first two reasons why a victim may be accepting an apology
unwillingly has to do with pressure and manipulation. Gabriel Teninbaum
recently argued, for example, that apologies are used by sophisticated
commercial firms as means of beguiling victims.92 Teninbaum highlighted
certain practices of apologies used by firms that are meant to create
emotional pressure on victims to accept them, a decision that the victims
88 Of course, it may well be that there is too much deterrence in the baseline, so
the change will be favorable. However, there are many reasons to believe that the tort
system is generally under-deterrent, especially given injurers ability to shield assets after
an accident. See generally Yonathan A. Arbel, Shielding of Assets and Lending Contracts,
48 INT’L REV. L. & ECON. 26 (2016)
89 Livingstone v Rawyards Coal Co (1880) 5 App Cas 25,39 (‘‘Tort seeks to put
the victim in the position he was in before the tort.’’)
90 See e.g., supra note 26.
91 See infra section II.C.
92 See generally Teninbaum, supra note 40.
Electronic copy available at: https://ssrn.com/abstract=2835482
8-Sep-16] Arbel & Kaplan - Draft 23
will later come to regret.93 The strategic, deliberate use of apologies by
commercial firms is designed to maximize this effect and victims are
employing only limited agency in their decision to accept the apology. A
second reason concerns pressure that comes from sources besides the injurer
itself. Pursuant to an apology, victims may still wish to sue; however, they
are often subject to social or internal pressure to avoid doing so, lest they
be perceived as vengeful, unrelenting, or ungrateful. Research in
psychology shows that failure to accept an apology is associated with a
negative perception of the victim.94 Similarly, victims may experience an
internal or social pressure (perceived or real) not to sue, due to the social
norm of accepting apologies.95
The two other reasons are more epistemological. There is a real
question as to whether people understand the meaning of commercial
apologies and how they are different from interpersonal ones. When a firm
apologizes through one of its proxies, is that an expression of guilt? Of
whom? Given how dispersed the decisions and actions in a commercial firm
are, even an apology by the CEO reflects a sliver of the actual responsibility
(aside from the very general sense in which the CEO is the personification
of the firm, a loaded idea by itself).96 What does the apology say about the
future? Would a commercial firm be less likely to recidivate after an
93 Id., at 332 (“On its own, convincing an individual not to sue is no different than
any other “bad” settlement. What makes this different is the appearance of a system of
methods designed to dissuade patients from actually considering their rights before settling
for short money.”)
94 See Mark Bennett & Christopher Dewberry, ‘‘I’ve said I'm Sorry, Haven't I?’’
A Study of the Identity Implications and Constraints That Apologies Create For Their
Recipients, 13 CURR. PSYCHOL. 10 (1994). See also Joost M. Leunissen et al., The apology
mismatch: Asymmetries between victim's need for apologies and perpetrator's willingness
to apologize, 49 J. Exp. Soci. Psy. 315, 315 (2013) (‘‘Victims of transgressions are, in
turn, socialized into graciously accepting such apologies’’). This is in line with the view of
some economists that apologies create a “psychic cost” to suing. See Benjamin Ho & Elaine
Liu, Does Sorry Work? The Impact of Apology Laws on Medical Malpractice, 43 J. RISK
UNCERTAIN. 141, 148 (2011).
95 Some moral philosophers believe that there exists a duty to forgive, see
CHARLES GRISWOLD, FORGIVENESS: A PHILOSOPHICAL EXPLORATION, 67 (2007) (‘‘under
certain conditions it would be blameworthy not to forgive’’); and Espen Gamlund, The
Duty to Forgive Repentant Wrongdoers, 18 INT. J. PHILOS. STUD. 651 (2010) (arguing that
a limited duty to forgive exists), and so it is possible that some people have a mistaken
sense of duty to accept apologies, even when they are not genuine.
96 On diffusion of responsibility in firms, see infra section II.B.1 and II.B.3.
Electronic copy available at: https://ssrn.com/abstract=2835482
24 Draft, [8-Sep-16
apology? The meaning of this apology is an open question, and this leads
us to the fourth reason which has to do with firm anthropomorphism. It is
well known that people do not maintain a clear distinction between
individuals and firms, tending to endow brands and firms with personality.97
Humans have a strong tendency-----potentially related to evolutionary
reasons-----to accept apologies from other humans.98 The concern is that this
instinctive reaction is carried over to brands and firms without proper
reflection. People may intuitively interpret an apology by a firm in a similar
way to how they interpret an apology by a person, much like how
individuals feel that certain brands and companies are ‘warm’ or ‘evil’-----a
phenomenon known as brand personification.99
The final flaw, and perhaps the most fundamental one, is the
unrealistic magnitude of the hypothesized therapeutic effect. Even if
apologies have some healing effect, there must be some limit to the size of
this effect. If victims are willing to forgo small or perhaps even moderate
value claims in exchange for an apology, then it may be that they engage in
a conscious trade-off of pecuniary and non-pecuniary benefits, preferring
the latter to the former. However, the greater the amount the victim forgoes,
the less persuasive is the idea that there is a real trade-off of benefits. It is
less persuasive, we think, to argue that a disabled victim of an accident
prefers an apology to amounts large enough to considerably alleviate her
suffering. As we shall show, the effect of commercial apologies can be
measured sometimes in the hundreds of thousands of dollars, a fact that puts
considerable pressure on the therapeutic value theory.100
Overall, we think that there is good reason to suspect the therapeutic
theory of apologies, at least in commercial settings. We cannot completely
overrule the possibility that victims are sophisticated and fully understand
the difference between commercial and personal apologies and find them
97 See Martin Eisend & Nicola E. Stokburger-Sauer, Brand Personality: A Meta-
Analytic Review of Antecedents and Consequences, 24 MARK LETT 205, 205 (2013) (‘‘In
their pursuit of fulfilling self-definitional needs, individuals tend to increasingly perceive
brands as relationship partners.’’). Brand personality is understood as ‘‘the set of human
characteristics associated with a brand’’, Jennifer Aaker, Dimensions of Brand Personality,
34 J. MARK. RES. 347, 347 (1997).
98 See Yohsuke Ohtsubo & Esuka Watanabe, Do Sincere Apologies Need to be
Costly? Test of a Costly Signaling Model of Apology, 30 EVOL. HUM. BEHAV. 114 (2009)
(considering apologies as an evolutionary adaptation).
99 See generally Ronal Jay Cohen, Brand Personification: Introduction and
Overview, 31 PSYCHOL. MARK. 461 (2010).
100 See infra section II.C.
Electronic copy available at: https://ssrn.com/abstract=2835482
8-Sep-16] Arbel & Kaplan - Draft 25
nonetheless satisfying-----so satisfying that they are willing to forgo very large
amounts of money in exchange for an apology. Yet, we find this possibility
less probable than the other explanations proposed here. The following
thought experiment might elucidate our skepticism. Consider the common
victim of medical malpractice, who suffered a great harm from negligent
treatment. Suppose that after the accident, she receives an apology from the
hospital staff or the physician, and as a consequence, she is dropping the
lawsuit. Going back in time but knowing what the patient knows now, would
the patient undergo the same procedure again? If the answer is negative,
then it is unlikely that the apology really mended the harm, that it genuinely
compensated the victim for her loss.
Overall, we see tension between the goals of tort law. Cost-reduction
is seemingly favorable to apologies whereas deterrence argues against
apologies. Compensation fails to point in any clear direction, so it does not
provide guidance on how to resolve the tension. To account for this
complexity, we need a theory that accounts for the combined effects of cost-
reduction and deterrence.
2. A Unified Theory of Apologies in Tort Law
To evaluate the combined effect of apologies on behavior, we extend
the traditional model of accidents in tort law to account for apologies. An
informal presentation follows here and the interested reader would find the
formal explication in the Appendix.
In the basic model of tort liability, a potential injurer chooses
whether to engage in a risky activity. The activity has some benefit to the
injurer but may cause harm to the victim. The prototypical example of this
model is driving and the risk of an accident to a pedestrian. An important
aspect of the model is that litigation over the accident is costly. To win the
case, each party has to expend resources on retaining lawyers, hiring expert
witnesses, producing evidence, etc. In addition to these litigation costs,
there are also liability costs, which reflect the payments the injurer would
have to pay the victim if found liable (or if the parties settle). The social
goal is to find rules that minimize costs.101 On this point, it is worth
emphasizing that the economic analysis does not consider the payment of
liability costs to have any direct effect on social welfare --- when a person
pays an amount to another person, then the second person becomes richer
(a social benefit), but this benefit is completely offset by the loss of the first
person.
101 See supra note 84.
Electronic copy available at: https://ssrn.com/abstract=2835482
26 Draft, [8-Sep-16
To account for apologies, we add to the model the possibility that if
an accident occurs, the injurer may choose to apologize. In order to do that,
it is important to clearly identify the benefits and costs of apologies. On the
side of benefits, the literature points out to two potentially distinct effects:102
the victim is willing to settle more often and is demanding a lower amount
in settlement negotiations.103 The costs are those of tendering an apology,
which may involve loss of face, social stature, or reputation.104 Tendering
an apology is a private cost that is borne by the injurer.
Are apologies after an accident socially desirable? Based on the
extended model, we will now argue that the injurer will tend to apologize
either more or less often than is socially optimal, a point that was not fully
recognized in the literature.105 To see that, consider first the private
incentive to apologize, the way the injurer sees it. From this perspective,
the tender of the apology will involve a cost that the injurer bears, but the
apology will also have a double benefit----- the savings on the injurer’s
litigation and liability costs. If the benefits exceed the cost of apologizing,
the injurer would have an incentive to apologize.
From a social perspective, the calculus is different. From this point
of view, we would again count the cost of tendering the apology. However,
the benefits will be very different. First, the savings on liability costs will
not be counted. As noted, from the social perspective, the fact that the
injurer will save money by not paying the victim the full amount is not a
social benefit, as the injurer’s gain is the victim’s loss. At the same time,
the benefit of reducing litigation costs involves a saving to both the victim
and the injurer, but the injurer will only count her savings, whereas society
will care about the joint savings. We see there is reason for the private
incentive to apologize to diverge from the social optimum, leading the
102 See infra section II.C.
103 Alternatively, apologies reduce payments because juries and judges are more
lenient towards repentant injurers.
104 For example, one public official preferred being sent to prison than a halfway
house because he did not want to apologize, White, supra note 17, at 1269. See also Ebert,
supra note 20, at 334-35 (discussing ego and the difficulty physicians face in admitting their
professional shortcomings).
105 For a similar argument in the broader context of litigation, see Steven Shavell,
The Fundamental Divergence Between the Private and the Social Motive to Use the Legal
System, 26 J. LEGAL STUD. 575 (1997) (explaining that there would generally be too little or
too much litigation because parties’ private incentives to bring suit will often be too weak or
too strong relative to the social optimum).
Electronic copy available at: https://ssrn.com/abstract=2835482
8-Sep-16] Arbel & Kaplan - Draft 27
injurer to apologize too much or too little.106
Example 1. Suppose that an accidental poisonous leak from a nearby
factory caused the victim a harm of $5,000. Further suppose that tendering
an apology would cost $500, but that through this apology, the parties settle
the case-----thus each avoiding $200 in litigation costs. Finally, suppose that
because of the apology, the victim is willing to accept a payment of $2,500,
rather than the $5,000 the victim would have received in litigation. In this
example, an apology will not be socially desirable, as it costs $500, but only
saves $400 in litigation costs. (Recall that the $2,500 saving to the offender
is equal to the victim’s loss). On the other hand, by apologizing the injurer
could save $2,700 (2500+ 200) at a cost of only $400, thus giving her an
incentive to apologize. Since the private incentive to apologize exceeds what
is socially desirable, there will be too much of an incentive to apologize.
Example 1a. Suppose now that the apology costs only $300 to tender
but that it does not reduce the amount in settlement. In this case, the apology
will be socially valuable, as by investing $300, a total of $400 in litigation
expenses could be saved. The injurer, however, will not have an incentive
to invest $300 as this will only help her save her own litigation costs of
$200.
We see that the social and private incentive to apologize may
diverge. We would expect there to be too many apologies under a
combination of the following circumstances: apologies have a strong effect
on victims’ willingness to forgo parts of their claims, injurer’s own litigation
costs are high, and apologies are cheap. Indeed, there may also be cases
where injurers will have too little incentive to apologize, in which case,
apology laws would be desirable. Which of these two options is more
probable has to do with one’s assessment of the magnitude of the cost of
tendering an apology relative to the effect of apology on the victim. The
stronger the effect, or the lower the cost of apologies, the more we will be
concerned with having too many apologies.
The analysis should not stop here. How would the ability to
apologize affect the decision to undertake the risky activity in the first place?
Tort theory recognizes that injurer’s decision will be affected by how much
the injurer anticipates they would have to pay should an accident occur.
106 Note that at this stage, we do not take into account the possibility that making
the injurer pay will reduce her incentive to harm in the future. The analysis so far is made
‘‘ex-post’’, that is, under the assumption that an accident has already happened.
Electronic copy available at: https://ssrn.com/abstract=2835482
28 Draft, [8-Sep-16
Under the standard analysis, it is suggested that if the expected payment will
be equal to the harm, the injurer would have optimal incentives.107 With a
sanction equal to the harm, a factory will not produce goods with a value of
$5,000 if the expected harm from a pollution-related accident exceeds
$5,000. Making the factory owner pay $5,000 in the event of an accident
would make sure she would only have an incentive to produce when the
value of the goods exceeds $5,000.
This result changes when we consider apologies. When
contemplating the possibility of an accident, the injurer would take into
account several costs. If no apology is tendered, these costs include the
expected costs of litigation and the costs of liability (e.g., $5,000).108 And
if the injurer decides to tender an apology, then as just analyzed, the injurer
will save some of the costs of litigation and liability, but will have to pay
for the apology itself. In this sense, the cost of delivering the apology can
be thought of as a self-inflicted punishment. Nonetheless, the injurer does
not have to apologize, and she will only do so if the apology is, on net,
privately beneficial. It follows that the injurer will only apologize if she
expects that to reduce her costs. This point emphasizes that the only
potential effect of apologies is to reduce liability.
Part of this reduction in payments is benign, as apologies encourage
settlement of cases that would otherwise litigate. The savings on litigation
due to greater propensity to settle is thus a positive feature of apologies. But
apologies do more than encourage settlements: they also reduce payments
the injurer would have to make to victims. Because the injurer cares about
her own private costs in the event of an accident, this reduction means that
the injurer has less to worry about an accident and less interest to take
precautions against such an accident. Overall, then, apologies dilute
deterrence.
Example 2. Suppose now that a factory owner thinks about using a
production technique that would save $4,000 in production costs-----but will
cause a harm of $5,000 from pollution to one of the neighbors. Suppose, as
before, that apology costs $400 to tender and that it leads to a settlement for
107 See Richard A. Posner, A Theory of Negligence, 1 J. LEGAL STUD. 29 (1972);
William M. Landes & Richard A. Posner, The Positive Economic Theory of Tort Law, 15
GA. L. REV. 851 (1981). See also Allan E. Farnsworth, Your Loss or My Gain? The Dilemma
of the Disgorgement Principle in Breach of Contract, 94 YALE L.J. 1339 (1985) (discussing
optimal remedies in the context of contract law).
108 We are assuming, as is conventional, that liability is set to equal the harm, but
not the harm plus litigation costs, See generally A. Mitchell Polinsky & Steven Shavell,
Costly Litigation and Optimal Damages, 37 INT. REV. L. ECON. 86 (2014).
Electronic copy available at: https://ssrn.com/abstract=2835482
8-Sep-16] Arbel & Kaplan - Draft 29
$2,500, thus saving $200 in litigation costs for both the factory owner and
the neighbor. We have already noted that the factory owner will have an
incentive to apologize in this case. Given that, the factory owner knows that
if she decides to use this production technique, she will gain $4,000 in
savings and her costs from an accident would be $2,900 (apology cost and
the settlement payment). Hence, the factory owner will have an incentive to
undertake the activity, pocketing the $1,100 difference. From a social
perspective, however, the activity causes a harm of at least $5,000 and only
has a benefit of $4,000,109 thus making it undesirable.
Example 2a. Suppose, as in 1a, that the apology costs $300 to tender
and that it does not reduce the amount in settlement. In this case, we have
seen, the injurer will not apologize, hence apologies will not have any
effects on behavior. More generally, if apologies are very costly to make,
they will not influence behavior.
Tying the analysis together, apologies may lead to unwanted
behavior when they are cheap and effective. After an accident, there may
be an excessive incentive for the injurer to apologize. This concern will be
most pressing when, among other things, apologies are cheap and effective
in terms of their effect on victims’ demands in settlement negotiations.
Before an accident occurs, apologies would tend to reduce the injurer’s
incentive to take care, a problem that is again most pressing when apologies
are cheap and effective. It should be emphasized that this does not mean
that apologies are always undesirable; if the apology reduces the cost of an
accident to the injurer by less than the savings it entails in litigation costs to
both parties, it is desirable. However, once the effect of an apology exceeds
that amount, apologies are no longer socially desirable, as the
encouragement of risky behavior exceeds the value of saving on litigation
costs. The main conclusion here is worth repeating: if apologies are cheap
and effective (in terms of reducing the amounts victims ask for), they are
undesirable.
The analysis also carries a strong normative message. The law
influences the ‘‘cost’’ of apologies, because making them privileged reduces
their downside, thus making them cheaper. The literature showed no
appreciation to the notion that there is a benefit to the cost of apologies and
that there may be excessive apologies. Because of that, the central theme in
the literature is that apologies should unconditionally be made cheaper --- an
idea that should be rejected on grounds of public safety. The analysis further
109 To be precise, the total harm given an apology is $5,800, which includes the
litigation costs of both parties and the cost of tendering the apology.
Electronic copy available at: https://ssrn.com/abstract=2835482
30 Draft, [8-Sep-16
suggests that there is an optimal level of cost of apologies: to the extent that
legislators can influence apology costs, they should set apology costs cheap
enough to encourage apologies to reflect the savings from litigation costs,
but no more than that.
B. Commercial Apologies in Practice
For individuals, sorry may be the hardest word. But when
commercial players enter the arena and the stakes are high, the balance of
costs and benefits of apologies changes.110 As the theoretical framework
highlights the importance of the costs of apologies, we move now to
illustrate how these costs tend to be (relatively) low or are on the decline,
through four different mechanisms.
1. Delegation & Specialization
When an individual tries to render an apology, she is limited by her
own abilities. If she is a bad communicator, seems insincere, or is
uncharismatic, then she may easily botch the apology. She only has herself
to work with; it will normally not do for her to send someone else to
apologize on her behalf.111 With commercial apologies, the situation is
markedly different. Corporations, by necessity, always delegate their tasks
to individuals. The ability to delegate confers on corporations an advantage
in apologizing, as it allows them some leeway in the choice of the individual
to tender the apology.112 By selecting the best apologizers, a firm’s apology
can be made as good as its best employee. This can be crucial, as different
individuals have remarkably different abilities when it comes to apologies.
Here, the BP oil spill case is especially illustrative. After having recognized
that the CEO apology did not go over well,113 the company realized that its
apology is ineffective because the CEO was not an American and thus was
110 See, e.g., Yonathan A. Arbel, Contract Remedies in Action: Specific
Performance, 118 W. VA. L. REV. 369, 398-99 (Finding that animosity plays a lesser role
between commercial parties).
111 See, e.g., Holley S. Hodgins and Elizabeth Liebeskind, Apology versus Defense:
Antecedents and Consequences, 39 J. EXPERIMENTAL SOC. PSYCHOL. 297, 310 (2003).
112 In some cases, it may be expected that the CEO or a specific employee will make
the apology. But it seems that in practice most corporate apologies are delivered various
other employees, including customer representatives.
113 See O’hara supra note 11, 1985.
Electronic copy available at: https://ssrn.com/abstract=2835482
8-Sep-16] Arbel & Kaplan - Draft 31
not viewed as part of the affected group.114 The company pivoted and
delegated the task of apologizing to local, ethnically diverse employees,
who were members of communities affected by the spill.115 BP ran
television ads featuring these employees representing the company, who
clearly identified themselves as locals to the Gulf Coast area and
communicated their personal grief as a result of the accident.116
Certain social expectations may be seen as constraining corporate
leeway in delegate apology tasks, such as the expectation that the
apologizing party will be the wrongdoer (say the attending physician) or that
the CEO will assume residual responsibility in the spirit of President
Truman’s famous plaque ‘‘the buck stops here’’.117 On reflection, however,
this constraint leaves much slack. In many corporate settings, each action
of the corporate is a composite of many different actions and decisions taken
by a diffused mass. There is no natural way to assign blame to a single
employee for a defective automobile coming out of the assembly line. This
entails that there will not be any natural person from whom to demand the
apology. Likewise, some medical procedures involve one physician, but
many involve more than one, which allows the hospital to make a choice
between the medical personnel. Even apologies by corporate leaders leave
room for discretion, as the company can hire managers that are especially
adept at apologizing and may create corporate positions that are mostly
symbolic to fulfill functions such as PR, social responsibility, and apologies.
Finally, Commercial entities are not even limited to their current staff. They
can, and routinely do, retain specialized experts for the management of
crises, such as mediators, actors, and celebrities. A company may choose
to install, for example, a personable CEO in times of crisis. Likable
employees have significant effect; as one medical malpractice practitioner
reported patients ‘‘never sue the nice, contrite doctors. Their patients never
call our offices.’’118
114 Id. at 1986.
115 Id. at 1989.
116 Id. at 1989; BP, BP Gulf Coast Update: Our Ongoing Commitment, YOUTUBE,
https://www.youtube.com/watch?v=hoOfIR4Vk1o (“I was born here, I’m still here, and so
is BP. We’re committed to the gulf. For everyone who loves it and everyone who calls it
home”, apology presented by Iris Cross, BP Community Outreach).
117 Harry S. Truman Library & Museum, ‘The Buck Stops Here’ Desk Sign,
https://www.trumanlibrary.org/buckstop.htm (last visited Aug. 8, 2016).
118 Wojcieszak et al., supra note 55, at 347. See also Bruce W. Neckers, The Art of
Electronic copy available at: https://ssrn.com/abstract=2835482
32 Draft, [8-Sep-16
2. Professionalization & Training
To be effective, an apology needs to be, or at least appear to be,
sincere. However, sincerity is never observed, only inferred --- the victims
cannot look into the psyche of the injurer and must rely on signals and
heuristics. By studying these heuristics using modern scientific methods, a
body of scholarship developed that focuses on identifying and exploiting
their weaknesses. Experts have shown, for example, how injurers can
structure apologies for maximal effect by leveraging in-group bias,119 using
effective language,120 choosing the right employees for the task, 121 and
timing apologies correctly.122
These lessons are transferred to commercial actors by specialized
firms through seminars and workshops. These firms help organizations
implement apologies as part of their workflow, suggest ways to streamline
the process of apologies, and offer best practices.123 One such example is
Sorry Works!—an advocacy organization and a training company—claiming
to have “trained thousands of healthcare, insurance, and legal professionals
from coast to coast and around the world” on how to use disclosure and
the Apology, 81 MICH. B.J. 11 (2002) (recounting the story of a client who said that an
apology would substitute a lawsuit).
119 For example, Erin O'Hara suggests that corporate wrongdoers may use local
spokespeople in their apologies, to maximize effect. O’Hara, supra note 11, at 1986 (noting
that the corporate apology was ineffective because the CEO has “thick British accent” which
“probably exacerbated the negative connotations of his resentful statements because it
pegged him and the company as foreign”)
120 See, e.g., Ameeta Patel & Lamar Reinsch, Companies Can Apologize: Corporate
Apologies and Legal Liability, 66 BUS. COMMUN. Q. 9, 21-22 (arguing that corporations can
reap the benefits of apologies with diminished legal exposure by switching from active
language—“I am sorry for hurting you”—to passive language—“I am sorry you were hurt”).
121 See generally Leanne Ten Brinke & Gabriella S. Adams Saving Face? When
Emotion Displays During Public Apologies Mitigate Damage to Organizational
Performance, 130 ORG. BEHAV. & HUM. DEC. PROCESSES 1 (2015) (studying the market
effects of facial cues given by corporate wrongdoers).
122 See generally Jochen Witrz & Anna S. Mattila, Consumer Responses to
Compensation, Speed of Recovery and Apology after a Service Failure, 15 INT. J. SERV.
IND. MANAG. 150 (2004) (studying the effects of timing on apologies).
123 For example, many corporations have strict guidelines on complaint handling
that include guidelines on apologies. See Christian Homburg & Andreas Fürst, How
Organizational Complaint Handling Drives Customer Loyalty: An Analysis of the
Mechanistic and the Organic Approach, 69 J. MARK. 95 (2005).
Electronic copy available at: https://ssrn.com/abstract=2835482
8-Sep-16] Arbel & Kaplan - Draft 33
apology to combat medical malpractice suits.124 The experience of the “3Rs
Program” instituted by the physician-trust COPIC is another telling example:
As part of the program, physicians are coached on effective apologies,
training them on timing, structure, and content.125
The professionalization and training in the area of apologies give
commercial actors a unique advantage. They allow these commercial actors
to apologize more effectively and at a lower cost, benefitting from the
accumulated knowledge and experience.
3. Diffusion of Responsibility
Commercial entities enjoy a psychological advantage, as the
psychological cost for the employee to deliver the apology tends to be lower
than that of delivering a personal one. Psychologists argue that an effective
apology requires a person to create the impression of separate parts of her
personality-----a past offender, who committed a wrong and is thus worthy of
scorn, and a present repentant apologizer, who deserves forgiveness.126 This
is a challenging task because the more one accepts responsibility, the more
she might inspire indignation, whereas assuming too little responsibility
may be taken as a failure to take ownership of the wrongdoing. For a
diffused commercial entity, this difficulty may be less severe, because the
party apologizing and the party at fault are not necessarily the same person.
We have noted above how corporate actions are a composite of many
different decisions of various individuals, which dilutes the responsibility of
every single actor. To the extent that the party apologizing and the victim
are not the same, the dissociation makes it much easier to apologize. First,
because it is always easier to admit that someone else was wrong rather than
oneself,127 and second, because the offender may be cast in a bad light
124 SORRY WORKS!, http://www.sorryworks.net/.
125 See Teninbaum, supra note 40, at 317.
126 See generally Peter H. Kim et al., Removing the Shadow of Suspicion: The
Effects of Apology Versus Denial For Repairing Competence- Versus Integrity-Based Trust
Violations, 89 J. APPL PSYCHOL. 104 (2004). A famous articulation of this idea is by
sociologist Erving Goffman, writing: “an apology is a gesture through which an individual
splits himself into two parts, the part that is guilty of an offense and the part that dissociates
itself from the delict and affirms a belief in the offended rule", ERVING GOFFMAN,
RELATIONS IN PUBLIC 113 (1971). On the relationship between apology and guilt, see Bruce
N. Waller, Sincere Apology Without Moral Responsibility, 33 SOC. THEORY & PRAC. 441
(2007).
127 Apologies are sometimes coupled with some remedial action. Here, again,
Electronic copy available at: https://ssrn.com/abstract=2835482
34 Draft, [8-Sep-16
without negative implications for the image of the apologizing party.128
For example, when Mary Barra, GM’s CEO, took office she
immediately had to start apologizing for the company’s faulty ignition
switches incident---a horrible accident that claimed the lives of 124
individuals.129 Barra had no personal role in the incident, and therefore she
was able to apologize profusely without admitting any personal fault (or
harming her reputation); indeed, she apologized so effectively that she was
heaped with praise at her congressional hearing: ‘‘God bless you, and you’re
doing a good job’’ replied Senator Baxter to Barra’s apology.130 Even in a
closer case, such as BP’s oil spill, CEO Tony Hayward was not personally
responsible for the explosion; the company claimed that it was mostly its
subcontractors who were to blame, and even though the court found the
company was grossly negligent, the blame is not rested solely with the
CEO.131
4. Corporate Culture
Scholars studying corporate culture and crisis management argue
commercial actors have more options than individuals. As William Benoit noted: “It may be
possible to limit damage by firing one or more employees, but Hugh Grant cannot fire
himself”, WILLIAM L. BENIOT, ACCOUNTS, EXCUSES, AND APOLOGIES: A THEORY OF IMAGE
RESTORATION STRATEGIES 48 (2015).
128 An unexpected advantage commercial entities have is related to the
standardization of apologies. It may seem that spontaneous apologies are more powerful than
scripted ones. If this were the case, corporations might have been limited in their ability to
control the provision of apologies. However, research shows that strict guidelines actually
result in more effective apologies. One study finds that apologies by the call center for
reservation or billing mistakes have strong and significant effect on consumer satisfaction.
See Anna S. Mattila & Daniel J. Mount, The Role of Call Centers in Mollifying Disgruntled
Guests, 44 CORNELL HOTEL RESTAUR. ADM. Q. 75 (2003). In another large qualitative study,
researchers in the area of marketing found that corporations with stricter guidelines and rules
on apologies and complaint management result in greater consumer satisfaction and sense of
justice. See Christian Homburg & Andreas Fürst, How Organizational Complaint Handling
Drives Customer Loyalty: An Analysis of the Mechanistic and the Organic Approach, 69 J.
MARK. 95 (2005).
129 See Danielle Ivory and Bill Vlasic, $900 Million Penalty for G.M.’s Deadly
Defect Leaves Many Cold, N.Y. TIMES, Sept. 17, 2015.
130 Ben Geier, Why do Some People Love GM’s CEO Mary Barra, FORTUNE, Aug.
9th, 2014, (quoting Senator Barbara Boxer (D, CA)).
131 Campbell Robertson & Clifford Krauss, BP May Be Fined Up to $18 Billion for
Spill in Gulf, N.Y. TIMES, Sept. 4, 2014.
Electronic copy available at: https://ssrn.com/abstract=2835482
8-Sep-16] Arbel & Kaplan - Draft 35
that before the 1990s, commercial apologies were seen as stigmatizing.132
The 1990s saw a deep change in the stigma and reputational effects of
commercial apologies. The reasons are complex and many explanations are
offered:133 the creation of a broader ‘‘new culture of apology’’,134 the rise
of the internet, and the introduction of relationship management strategies
in the 1990s.135 Another potential driver of these changes is the discovery
in the marketing literature of the ‘‘recovery paradox’’, whereby apologizing
may actually improve consumer relations relative to their level prior to the
adverse incident.136 Whatever the true explanation is, experts see a strong
change in the way apologies are treated today relative to the 1990s.137 Today
the ‘‘[c]onventional wisdom’’ among scholars in business administration and
branding ‘‘holds that public apology in response to accusations of corporate
misconduct is one of the most important ways to restore a company's
reputation’’.138 Today the default has reversed, and it is expected that
companies would apologize: if in the past only the guilty apologized, today
not apologizing is a violation of consumers’ expectations.139 Moreover,
132 See LAZARE, supra note 14, at 7.
133 For other explanations, see Zohar Kampf, The Age of Apology: Evidence from
the Israeli Public Discourse, 19 SOC. SEMIOT. 257 (2009).
134 See Nicolaus Mills, The New Culture of Apology, 48 DISSENT 113, 114 (2001);
Mihai, supra note 21 (“A gesture formerly considered a sign of weakness has grown to
represent moral strength and a crucial step towards potential reconciliation”). See also Jeffrie
G. Murphy, Well Excuse Me!—Remorse, Apology, and Criminal Sentencing, 38 ARIZ. STATE
L.J. 371 (2006) (noting, and criticizing, the proliferation of apologies).
135 See Jan Breitsohl et al., Online Complaint Communication Strategy: An
Integrated Management Framework for E-Businesses, HANDB. E-BUS. STRATEG. MANAG.
907, 908 (2014); Michael Volkov, Successful Relationship Marketing: Understanding the
Importance of Complaints in a Consumer-Oriented Paradigm, 2 PROBL. PERSPECT. MANAG.
113 (2004).
136 See, e.g., James G. Maxham & Richard G. Netemeyer, A Longitudinal Study of
Complaining Customers’ Evaluations of Multiple Service Failures and Recovery Efforts, 66
J. MARK. 57 (2002) (showing in a longitudinal study the existence of a recovery paradox, but
also noting that it disappears if there are multiple adverse events).
137 See, e.g., Patel & Reinsch, supra note 120, 14-15 (noting that hard data is hard
to find but the impression is that commercial apologies are frequently used).
138 John G. Knight, Damien Mather, and Brianne Mathieson, The Key Role of
Sincerity In Restoring Trust In a Brand With a Corporate Apology, in MARKETING
DYNAMISM & SUSTAINABILITY: THINGS CHANGE, THINGS STAY THE SAME 192 (2015)
139 See Sean Tucker et al, Apologies and Transformational Leadership, 63 J. BUS.
ETHICS 195 (2006).
Electronic copy available at: https://ssrn.com/abstract=2835482
36 Draft, [8-Sep-16
apologies are taken to be a sign of strength and leadership.140 An employee
would thus find the personal costs of apologizing much lower than in the
past; institutions, like hospitals and insurance companies, often provide a
support system, assuring the injurer an apology is the right and honorable
thing to do. The increased popularity of apologies makes their social cost
lower, as the reputational effect is diminished (and per the recovery
paradox, actually becomes positive).
C. Effectiveness of Commercial Apologies
Commercial actors, we just argued, enjoy important advantages with
respect to tendering apologies. It is, therefore, natural to doubt whether
these apologies have an effect on victims. Would not individuals reject
apologies in commercial settings, seeing them as strategic, profit-
maximizing decisions? Would not the making of repeated apologies by the
same institution adulterate their effect?
In fact, commercial apologies are highly effective. Researchers
studying commercial entities in online settings puzzlingly noted after finding
strong effects that it seems ‘‘as if customers do not realize that they are
interacting with an employee who is paid to send apology emails and not
with an individual who experiences shame when apologizing.’’141 The
researchers concluded their field test by noting that ‘‘[we] find that a cheap-
talk apology yields significantly better outcomes for the firm than offering
a monetary compensation.’’142 The effectiveness of commercial apologies
can be learned from their prevalence,143 but it would be useful to look at
more direct evidence, which also gives a sense of the magnitude of the
effect.
The best evidence comes from the healthcare industry, which is the
best-studied area of commercial apologies, due to the large stakes involved
and the tragic frequency of accidents.144 Starting in the 1990s, hospitals
140 Id. at 195 (Finding that “ethical leaders who attempt to do the right thing with
their words and actions will be perceived as better leaders by followers … ethical leaders
apologize.”
141 See Johannes Abeler et al., The Power of Apology, 107 ECON. LETT. 233, 235
(2010).
142 Id., at 107.
143 See, e.g., BENIOT, supra note 127, at 61 (noting the pervasiveness of corporate
apologies).
144 THE NATIONAL PRACTITIONER DATA BANK,
Electronic copy available at: https://ssrn.com/abstract=2835482
8-Sep-16] Arbel & Kaplan - Draft 37
became aware that many patients sue for emotional reasons, as they resent
the lack of apology.145 This realization led to a series of successful
experiments with institutionalizing apologies.146 An example is the pioneer
program of The University of Michigan Health System. The university
adopted a policy of disclosure and apology that required hospital personnel
and physicians to disclose mistakes and apologize for them. A detailed
before-after analysis of this program reveals significant effects. First, the
monthly rate of claims (defined as requests for monetary compensation) has
fallen by 36%.147 This means that about one-third of the victims gave up
their claims in their entirety. Second, the number of lawsuits has fallen by
65%.148 Third, the cost per lawsuit has fallen from $405,921 to $228,308,
a saving of $177,603 (44%). Fourth, the costs of lawsuits have not only
fallen due to savings on legal costs; the hospital saved about 59% of the
compensation costs it would have had to pay patients.149
Another example is COPIC, an insurance trust founded by
physicians that designed the ‘‘3Rs Program’’: Recognition of the patient’s
harm, Response to the issue in a timely manner, and Resolution-----through
apology and a small offer of compensation. Looking at the data, the offers
of compensation are indeed small: in most cases, no payment is made at all
and in the rest, the payment is for only $5,300.150 The program led to
striking results-----a reduction of 50% in the number of malpractice claims
http://www.npdb.hrsa.gov/analysistool/ (reporting about 50,000 medical malpractice
payments and adverse events in 2014).
145 See supra note 25.
146 See, e.g., Steve S. Kraman & Ginny Hamm, Risk Management: Extreme Honesty
May be the Best Policy, 131 ANNALS OF INTERNAL MEDICINE 963 (finding financial savings
in hospitals that implemented a disclosure and compensation policy). See also ROBERT D.
TRUOG ET AL., TALKING WITH PATIENTS AND FAMILIES ABOUT MEDICAL ERROR 52–56
(2011).
147 See Allen Kachalia et al., Liability Claims and Costs Before and After
Implementation of a Medical Error Disclosure Program, 153 ANN. INTERN. MED. 213, 215
(2010). See also Michelle M Mello, David M Studdert & Allen Kachalia, The Medical
Liability Climate and Prospects for Reform, 312 JAMA 2146, 2149 (2014).
148 Kachalia et al., supra note 147, at 215.
149 Id.
150 See Richard C. Boothman et al., A Better Approach to Medical Malpractice
Claims? The University of Michigan Experience, 2 J. HEALTH LIFE SCI. L. 125, 147-48
(2009). Compare with average case costs of medical malpractice lawsuits of about $300,000,
see Seth Seabury et al., Defense Costs of Medical Malpractice Claims, 366 N. ENG. J.
MED. 1354, 1354 (2012).
Electronic copy available at: https://ssrn.com/abstract=2835482
38 Draft, [8-Sep-16
against COPIC physicians and a reduction in the costs of payments in
settlement of 23%.151 In one of the case records, a 66-year-old patient
suffered from an error that led to the removal of part of her ureter, which
required a painful invasive procedure for its treatment. The program settled
the entire case by paying her $3,898 to account for her out-of-pocket
expenses and, ‘generously’, also for her ‘‘gardening/lawn bills’’.152 These
two apology programs reduced significantly the number of compensation
requests, the number of lawsuits, and, most importantly for our purposes,
the amounts paid to patients.
Looking more broadly, economists Benjamin Ho and Elaine Liu find
that commercial apologies are highly effective. The two have investigated
how apology safe-harbor laws affect malpractice lawsuits. Their studies are
based on the fairly innocuous assumption that apology laws increase the
frequency of apologies. Because of that, if we see a change in outcomes
following the legislation of an apology law, that change would be
attributable to the effect of apologies. Based on this methodology, they find
that a state that adopts an apology law sees a reduction of about 17% in
payments for severe medical injuries,153 which is equivalent to a reduction in
payments of $58,00-73,000.154 This is remarkable, as the averages come
from all hospitals—not necessarily those who instituted an apology policy—
which suggests that the real effect can be much larger. Consistent with that,
a recent working paper found that apology laws lead to a reduction of $65,000
in payments to victims across all injury levels.155
Apologies in a commercial setting are effective beyond the medical
context. In a vignette study, researchers found that consumers express
greater willingness to purchase from companies which apologized in a way
that was perceived as sincere.156 In the commercial context of housing,
Russell Korobkin and Chris Guthrie find that participants playing the role
151 See Boothman et al., supra note 150, at 147-48; Wojcieszak et al., supra note
55, at 346.
152 See Richert E. Quinn & Mary C. Eichler, The 3Rs Program: The Colorado
Experience, 51 CLIN. OBSTET. GYNECOL. 709, 715 (2008).
153 See Benjamin Ho & Elaine Liu, Does Sorry Work? The Impact of Apology Laws
on Medical Malpractice, 43 J. RISK UNCERTAIN. 141, 143 (2011).
154 Id.
155 See McMichael et al., supra note 10.
156 See Denghua Yuan et al., Sorry Seems to be the Hardest Word : The Effect of
Self-Attribution when Apologizing for a Brand Crisis, (HKIBS Working Paper Series 073-
1314 (2014).
Electronic copy available at: https://ssrn.com/abstract=2835482
8-Sep-16] Arbel & Kaplan - Draft 39
of tenants were more likely to accept a settlement offer for an infraction of
landlord’s duties if they were told that the landlord apologized.157 These
results seem to carry over to the market: in market settings, e-apologies led
disappointed consumers to retract unfavorable reviews, at a rate much
greater than when they were offered monetary settlements.158 Moreover,
firms are said to perform better in the stock market after taking
responsibility for past failures.159
III. CRITICAL ANALYSIS AND POLICY IMPLICATIONS
The normative framework we provide in Part Error! Reference
ource not found. demonstrates that apologies can have detrimental social
157 See Korobkin & Guthrie, supra note 26, at 148 (reporting a 12% increase, but
note that this effect failed to reach statistical significance. Nonetheless, the size and sign of
the effect are consistent with our argument).
158 On eBay, customers can leave negative responses which can later be withdrawn
if seller’s feedback satisfies the consumer. A group of researchers collaborated with a very
large seller and randomly modified its response to a negative review left by a customer on
transactions with average value of 23.5 Euros: small monetary compensation (2.5 Euro);
large monetary compensation (5 euros), and an apology, electronically delivered by one of
the employees, without admitting to any legal liability and without any monetary
compensation. They found that small monetary compensation yields forgiveness (i.e.,
retraction of the negative review) in 19.3% of the cases; doubling the amount of
compensation only slightly increases forgiveness to 22.9%. The tender of apology outdid
both measures, with a forgiveness rate of 44.8%. See Abeler, supra note 141, 234.
159 Consider, for example, the Domino’s 2009 crisis, when a disgruntled employee
publicized a video of himself committing what we can euphemistically call “health code
violations” of customers’ pizzas. Soon after, Twitter was flooded with tweets deriding the
company and its products. Patrick Doyle, the company’s President, uploaded a response
video to YouTube, in which he said that he is sickened by the act, apologized and reported
corrective action. See Domino's President Responds to Prank Video, YOUTUBE (June 3,
2010) https://www.youtube.com/watch?v=dem6eA7-A2I. An empirical analysis of 20,773
tweets discovered that this was highly effective and the corporate brand, as reflected by
tweets, was restored to its original levels, See Hoh Kim et al., The Effect of Bad News and
CEO Apology of Corporate on User Responses in Social Media., 10 PLOS ONE e0126358
(2015). Others in the field reflected similar appreciation of the effectiveness of this apology,
and although far from being necessarily causally related, the brand is thriving. Domino’s
stock price is about ten times its value in 2009 . On firms’ performance, see Don Chance,
James Cicon, and Stephen P. Ferris, Poor Performance and the Value of Corporate Honesty,
33 J. CORP. FIN 1 (2015). Indeed, the return on investment in apology mechanisms was
estimated by researchers as being greater than 100% in some cases. See Christian Homburg
& Andreas Fürst, How Organizational Complaint Handling Drives Customer Loyalty: An
Analysis of the Mechanistic and the Organic Approach, 69 J. MARK. 95, 95 (2005).
Electronic copy available at: https://ssrn.com/abstract=2835482
40 Draft, [8-Sep-16
implications unless certain conditions are met. We have also shown that
commercial apologies are both cheap to tender and highly effective. In light
of this, we move to critically analyze the movement that transformed the
law and to outline necessary policy changes in response to this reform.
A. Better Sorry than Safe
Our theoretical analysis demonstrates that apologies are socially
undesirable if they are relatively cheap to tender and if they have strong
effects on the amounts victims seek. When these conditions obtain, the
problem is that sophisticated commercial actors would be able to anticipate,
before they engage in dangerous activities, that an apology would reduce
their exposure to liability for any ensuing accidents. Because of that, they
would have less incentive to be careful, which may increase the level of
accidents. Hence, they would find it preferable to be sorry rather than safe.
Indeed, if apologies are costly to tender or only mildly effective, this
concern does not arise. However, we believe our analysis above strongly
suggests the possibility of a problem, as commercial apologies are likely to
be both effective and cheaper to deliver in commercial settings. To
illustrate, in one case, a patient was willing to settle after the apology simply
because she felt the hospital took her case seriously.160 The hospital, on the
other hand, saved an approximate $3 million in liability payments in a
lawsuit that, according to the hospital’s estimation, was highly likely to
win.161 Of course, the apology itself had some cost, but nothing in the
evidence indicates this cost was large; indeed, this case is touted for its cost-
saving effect.162
A clear prediction that follows from our analysis is that apology laws
will increase the level or severity of accidents in states that adopt them
relative to non-apology laws states. This is in contrast to the hypothesis of
the Legal Apologists, which stipulates that apology laws will reduce levels
of litigation without a corresponding increase in the level of accidents. The
implications of our prediction are disconcerting. If commercial injurers can
easily escape liability, they would not have a real incentive to be safe. A
food company may employ less quality assurance procedures, a hospital
may order less expensive tests, and a large polluter may install fewer filters
and smoke scrapers than otherwise. To assess the validity of each
hypothesis, we need empirical data; unfortunately, the empirical data we
160 Boothman et al., supra note 150, at 157.
161 Id.
162 Id.
Electronic copy available at: https://ssrn.com/abstract=2835482
8-Sep-16] Arbel & Kaplan - Draft 41
have from the two studies on the topic is inconclusive, although it is largely
consistent with our prediction.
The most rigorous analysis to date was conducted by economists Ho
and Liu, who looked at the effect of apology laws on the level of disposed
medical malpractice claims.163 They find that apology laws increase the
number of disposed claims involving severe injuries by 21-27% and that
payments for severe injuries increase by 20-28%. This would seem to
suggest a rise in accidents and their severity, but the problem is that the data
consists only of disposed cases and the definition of disposed cases makes
it hard to draw any conclusions. In fact, Ho and Liu argue that the rise is
mostly attributable to the greater speed of processing claims and that over
time, there are fewer claims.164 But this conclusion is constrained by the
meaning and interpretation of disposed claim, a problematic category that
only includes complaints with positive money payments and so it does not
include all, if not most, of the accidents or all the cases where no payment
was made.165 Another limitation is that it is possible to make unreported
payments, and some hospitals seem to be doing so.166
Another problem is the tension between their findings and those of
another, more recent working paper.167 In this study, researchers obtained
data from an insurer that accounts not only for disposed claims with positive
payments, but for all claims that were filed with the insurer. Indeed, this
does not account for accidents that do not result in a formal claim, but it
163 Benjamin Ho & Elaine Liu, Does Sorry Work? The Impact of Apology Laws on
Medical Malpractice, 43 J. RISK UNCERTAIN. 141 (2011).
164 They indeed find that over time apology law states see a significant reduction in
disposed claims for non-severe injuries, but they also find an increase in the level of severe
injuries, which they interpret as resulting from a staggered effect of the apology law.
However, these findings are also consistent with the theory that apology laws increase the
severity of accidents. Id. at 162.
165 This is especially a problem, since most cases are disposed without payment, see
McMichael, supra note 10, at 16 (“Analysis of our data indicates that excluding claims that
involved no payment to a claimant results in excluding over half of all malpractice claims.”)
166 See Teninbaum, supra 40, at 316-17 (discussing rules in an apology program
that are designed to circumvent reporting requirements). See also Amitabh Chandra,
Shantanu Nundy, and Seth A. Seabury, The Growth Of Physician Medical Malpractice
Payments: Evidence From The National Practitioner Data Bank, HEALTH TRACKING, May
31, 2005 (estimating underreporting in data of about 20% of malpractice payments).
167 See McMichael, supra note 10.
Electronic copy available at: https://ssrn.com/abstract=2835482
42 Draft, [8-Sep-16
does provide a broader approach to the issue.168 The authors of the study
find that apology laws reduce payments to patients by 82%, which is
equivalent to a reduction of $65,000 in the average payment.169 They
explain this result as driven mostly by the increase of claims that do not
result in payment. In other words, they find that apologies mainly increase
the level of claims where no payment is made, but do not affect the level of
payments in other cases. They also find, however, that claims are more
likely to turn into a lawsuit under apology laws-----which is clearly
inconsistent with the goals of apology laws. Both these studies provide much
needed insight, but they do not clearly illuminate the key variable of
interest: the level of accidents. The lack of more focused research is
potentially attributable to the misunderstanding of the potential negative
effect of apologies on incentives, and we hope that this Article will spur
future research in this area.
B. The Paradox of Excessive Apologies
At the heart of the apology law reform is the argument that injurers
are wary of apologizing due to the legal ramifications of exposing
themselves to liability.170 To overcome this fear-----the argument goes-----
apologies should be privileged, shielding injurers from the evidentiary
implications of potential admission of fault. The Legal Apologists argue that
privileging apologies would encourage injurers to apologize, thus leading to
important benefits, most importantly, the control of litigation costs.171
This statement involves a potential paradox with no easy resolution.
The first argument-----that injurers do not apologize for fear of legal
liability-----assumes that unprivileged apologies encourage litigation. But at
the same time, the main reason that Legal Apologists argue that apologies
are desirable is that they encourage settlement and therefore discourage
litigation. It is seemingly paradoxical to argue that apologies both encourage
and discourage litigation. Resolving this paradox comes at a price.
For example, perhaps unprivileged apologies have disparate effect;
they reduce the incentive to bring suit but increase the probability that the
victim will prevail in a lawsuit by having better evidence. While coherent,
168 Id. at 10. On the other hand, their data is limited to only one specialty area,
which may introduce other kind of unanticipated bias.
169 Id. at 27.
170 See supra notes 62-63.
171 See supra note 22.
Electronic copy available at: https://ssrn.com/abstract=2835482
8-Sep-16] Arbel & Kaplan - Draft 43
this resolution also raises problems. It is unclear why the evidentiary
advantage of apologies does not entice more victims to file lawsuits. More
importantly, if privileging apologies will not reduce the level of litigation
but will only reduce the likelihood that the victim will prevail, then
apologies lose much of their luster.
Another possibility is to argue that apologies have a heterogeneous
impact on victims. Some victims will sue unless they receive an apology,
so apologies would reduce litigation in their case and are thus desirable.
Other victims would only sue if they receive an apology (as the apology will
provide them with sufficient evidence) and for this class of victims,
privileging apologies will reduce litigation costs. While coherent, this
resolution is also problematic, as it omits the class of victims who would
sue even in the presence of an apology. Privileging apologies will reduce
the likelihood that this class of victims will prevail in litigation, and thus
involves a cost.172 Whether this cost exceeds the benefit of controlling
litigation from the other group is an empirical question, which admits the
possibility that apologies will be undesirable.
C. Apology as Disclosure
A recurrent narrative, especially among medical professionals, is
that apologies help because they facilitate the communication of mistakes,
as put by Clinton and Obama:
Under our proposal, physicians would be given certain
protections from liability within the context of the
program, in order to promote a safe environment for
disclosure. By promoting better communication, this
legislation would provide doctors and patients with an
opportunity to find solutions outside the courtroom.173
On this account, privileging apologies would mean that injurers
would be more willing to admit their mistakes. The reason why admitting
172 To be clear, injurers would save a corresponding amount, as they would be
more likely to prevail in litigation. However, if we make the (natural) assumption that the
likelihood of prevailing at trial corresponds to the culpability of the injurer, then privileging
apologies would benefit mostly with culpable injurers, thus undermining deterrence.
173 See Clinton and Obama, supra note 1, at 2207.
Electronic copy available at: https://ssrn.com/abstract=2835482
44 Draft, [8-Sep-16
mistakes is important is an instrumental one; by recognizing mistakes, the
parties can learn and do better in the future.174
This logic may be applicable in many interpersonal settings, but it
transfers poorly to a commercial environment. Before touching on this
point, it should be noted that the basic assumption here --- that mistakes are
not divulged due to liability-----is doubted by many who believe the main
causes for hiding mistakes are factors such as culture and social norms,175
and indeed, studies comparing the rate of disclosure of errors in the United
States and countries with lower levels of liability for medical malpractice
find no difference in error reporting in hospitals.176 Adeeper problem is the
assumption that once identified, mistakes will be corrected. In many
commercial settings, learning from one’s mistakes is not simple. Taking
precautions will often involve investment in machinery, staff, and strict
regulation. These costs can be very high-----consider the cost of purchasing
an MRI machine or even of standard bloodwork procedures if done on a
large scale-----and it will certainly be contradictory to our approach in most
other areas of law to believe that actors will have sufficient incentive to
internalize the costs of their actions without the threat of any legal action.177
This inconsistency was noted by David Hyman and Richard Silver:
[I]t is naïve to think that error reporting and health care quality
would improve automatically by removing the threat of
liability. . . . No statistical study shows an inverse correlation
between malpractice exposure and the frequency of error
reporting, or indicates that malpractice liability discourages
providers from reporting mistakes.178
174 Id. at 2205.
175 TOM BAKER, THE MEDICAL MALPRACTICE MYTH 97 (2005) (arguing that “you
first have to prove that mistakes would be out in the open if there were no medical
malpractice lawsuits. That is clearly not the case.”).
176 Amy Widman, Liability and the Health Care Bill: An “Alternative” Perspective,
1 CAL. L. REV. CIRCUIT 57, 59 (2010); George J. Annas, The Patient’s Right to Safety —
Improving the Quality of Care Through Litigation Against Hospitals, 354 NEW ENG. J.
MED. 2063, 2065-66 (2006) (comparing with New Zealand)
177 Clinton and Obama proposed that savings from apology programs will be used
to reduce the premiums doctors pay-----but this would the equivalent of transferring money
from victims of accidents to physicians. They also proposed that some of the savings will
be used to ‘‘foster patient-safety initatives’’. Id. at 2207.
178 David A. Hyman & Charles Silver, The Poor State of Health Care Quality in the
U.S.: Is Malpractice Liability Part of the Problem or Part of the Solution?, 90 CORNELL L.
REV. 893, 898–99, 914 (2005)
Electronic copy available at: https://ssrn.com/abstract=2835482
8-Sep-16] Arbel & Kaplan - Draft 45
D. The Deficit of Apology Deficit
Motivating the entire movement of the Legal Apologists is the belief
in an apology deficit. The concern is that injurers have too little incentive
to apologize and therefore they need encouragement. It may seem odd in
retrospect, but besides anecdotal evidence, the point that there is a deficit
in apologies was never proven. Do we really have a deficit of apologies?
Are commercial actors shying away from apologizing?
The core problem is that even without any reform, commercial
injurers should have a strong incentive to apologize. As we have noted,
apologies create value to injurers by suppressing their litigiousness. If
apologies are value-creating, then just like any other goods, profit-
maximizing companies would seek to ‘‘produce’’ them. Indeed, given the
many benefits Legal Apologists ascribe to apologies, it would be odd if
companies would not provide them. The literature is in agreement that there
is a marked transition among companies from the age of ‘‘deny and defend’’
to ‘‘apologize and settle’’.179 Today, commentators agree, commercial
apologies have become commonplace.180 As early as 2002, well before most
states adopted apology laws, a survey of hospital risk managers revealed
that 68% would respond to a mistake with an apology, which suggests a
broad appreciation of the commercial benefits of apologies.181
Psychiatrist Aaron Lazare conducted a casual empirical analysis to
develop a basic intuition of the prevalence of commercial apologies, by
looking at the discourse on apologies in the media.182 To expand his
analysis, we reanalyzed the data using a larger database. Consistent with his
findings, Figure 1 illustrates the findings on the basis of a broad range of
media reports acquired from the EBSCO database, which includes 25
million media articles from the relevant time period.183 As can be seen, until
179 See Sandra Harris, Karen Grainger & Louise Mullany, The Pragmatics of
Political Apologies 17(6) DISCOURSE & SOCIETY 715 (2006).
180 Roy L. Brooks, The Age of Apology, in WHEN SORRY ISN’T ENOUGH: THE
CONTROVERSY OVER APOLOGIES AND REPARATIONS FOR HUMAN INJUSTICE, 8–11 (1999,
Roy L. Brooks, ed.).
181 See Rae M. Lamb, et al., Hospital Disclosure Practices: Results of a National
Survey, 22 HEALTH AFFAIRS 73 (2003)
182 LAZARE, supra note 14, at 6-7.
183 The Methodology consisted of search results for apology or sorry or related
Electronic copy available at: https://ssrn.com/abstract=2835482
46 Draft, [8-Sep-16
the 90s, apologies were hardly considered in the media. But starting in the
90s, there has been a growing interest that persists till today.
Figure 1: Apologies in Print: Mentions by Year
Data: EBSCO, 1971-2015
250
200
150
100
50
0
1 3 5 7 9 1 3 5 7 9 1 3 5 7 9 1 3 5 7 9 1 3 5
7 7 7 7 7 8 8 8 8 8 9 9 9 9 9 0 0 0 0 0 1 1 1
9 9 9 9 9 9 9 9 9 9 9 9 9 9 9 0 0 0 0 0 0 0 0
1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 2 2 2 2 2 2 2 2
Overall, the consensus in the literature on the ‘‘age of apologies’’ is
well reflected in this analysis. While this does not amount to a rigorous
analysis of the topic, it does suggest that the apology deficit may not exist.
E. Policy Implications
The on-going tort reform through apology laws is politically and
legally problematic. There are currently calls to further expand the ambit
of apology laws,184 and to encourage mediators and arbitrators, judges, and
juries to account for them.185 If past success and momentum are any
words in title or subject terms, restricted to magazines, newspapers, reviews and trade
publications, in the English language, between 1971 and 2015. A total of 4967 results were
located, which, after removing duplicates, was narrowed to 3747. Permalink to results:
http://search.ebscohost.com/login.aspx?direct=true&db=aph&bquery=(TI+apology)+OR+(
SU+apology)+OR+(TI+sorry)+OR+(SU+sorry)&cli0=LA99&clv0=Eng&type=1&site=eh
ost-live&scope=site (gated). To account for a potential bias due to the fact that more media
is produced today than in the past, we validated our findings by limiting search to the New
York Times, the Economist, New York Times Magazine, and the New Yorker – all existing
prior to 1971.
184 See Runnels, supra note 4, at 148 (2009); Cohen, supra note 12, at 1061. See
generally Gailey, supra note 4; Jones, supra note 14, at 580-81.
185 See, e.g., Robyn Carroll, Apologies as a Legal Remedy, 35 SYD. L. REV. 317
(2013).
Electronic copy available at: https://ssrn.com/abstract=2835482
8-Sep-16] Arbel & Kaplan - Draft 47
indication, these calls are likely to be translated into legislation in the near
future. Our analysis suggests that the case presented by reformists is lacking
in theoretical and empirical support. There is a question whether there is an
apology deficit and there is a real concern that apologies will be used to
circumvent legal liability for accidents by strategic actors. Politically, we
expressed the concern that apology laws have been used as a covert tort
reform, avoiding public scrutiny. These issues raise a few policy
implications.
First order of business is transparency. Apologies laws should be
understood-----and debated-----in terms of tort reform. The public, advocates,
and legislators, should be made aware of the social effects of apology laws.
This does not mean that apology laws should never be enacted-----the debate
on tort reform is an active one. However, the debate should be conducted
transparently, not in terms of virtue or penance, but in the more real terms
of reducing compensation to victims which may or may not be excessive.186
Second, a moratorium should be placed on all future expansions of
apology safe harbor laws. Besides the political concern, there are the social
concerns. Apology laws make the tender of apologies ‘‘cheaper’’ from the
viewpoint of the injurer, and the analysis demonstrates that reducing the
costs of apologies can lead to socially harmful outcomes, in the form of
risky behavior. The evidence we gathered suggests that this risk is real,
given the effectiveness of commercial apologies and their low cost.
Third, there is a push to encourage judges and juries to show
leniency in their judgments towards remorseful injurers.187 In a sense, these
186 Supporters of tort reform would also benefit from a better recognition of the
effect of apologies. There are many tools in the tort reformers’ toolkit, such as damages
caps, procedural adjustments, and panel screening of cases. Each of these tools has its own
advantages and shortcomings. Compared with damages caps, for example, apology laws
have the disadvantage of being impossible to calibrate. If one thinks that the true harm
from a medical accident is $250,000, then a damages cap at this level could rein in courts.
But the effect of apologies on victims is highly idiosyncratic and it does not allow for easy
corrections. On the other hand, apology laws encourage informal settlements, and this may
have merit of its own. Either way, a candid evaluation of alternatives would be prudent.
187 See, e.g., Bibas & Bierschbach, supra note 17, at 128-29 (advocating lenient
treatment of remorseful offenders). Interestingly, a new study provides preliminary evidence
suggesting that apologies have little effect on judges. See Jeffrey J. Rachlinski, Chris Guthrie
& Andrew J. Wistrich, Contrition in the Courtroom: Do Apologies Affect Adjudication?, 98
CORNELL L. REV. 1189 (2013) (finding in a vignette study that “a defendant's apology in
court is generally ineffective, sometimes counterproductive, and only occasionally
Electronic copy available at: https://ssrn.com/abstract=2835482
48 Draft, [8-Sep-16
initiatives are even more troublesome than the safe harbor laws, as safe
harbor laws protect apologies that can prevent litigation, but this policy
encourages apologies that do not even have this effect. Indeed, some have
argued that there is a case for treating apologizing defendants more
severely.188 We recognize that it may seem counter-intuitive to treat
remorseful and unremorseful injurers equally,189 but it is important to
remember that our discussion is limited only to commercial actors such as
companies, for whom the expression of remorse is at least suspect. In sum,
there should be a presumption against the preferential treatment of
commercial actors who apologize during trial.
Finally, the questions we raised here touch on important social
policies, but the data we currently have is limited. It will be important for
policymakers to devote funds and grants for studies in this area, and perhaps
there is room to use funding from Obamacare’s special allotment to this
end.190
IV. CONCLUSION
Over the last three decades, apology law reform has swept the
nation. Tort reformers and commercial interests provided funding to a
strong lobby that co-opted the rhetoric and discourse developed by a
movement of legal scholars we called the Legal Apologists. The work of
the Legal Apologists has contributed greatly to our philosophical, social,
and psychological understanding of the role of apologies in both the law and
in our daily lives. However, they have failed to articulate an account of
apologies in commercial settings and have not considered the potentially
socially harmful effects of apologies of this type. This oversight has not
been lost on tort reformers, who advocated apology law reformers to
effectively achieve tort reform through the backdoor.
We argued that making apologies cheaper may lead to socially
harmful outcomes. To support our claims, we developed a new model for
beneficial”).
188 Mungan argues that treating apologies more harshly helps differentiate between
sincere apologies (which are meant to relieve guilt) and non-sincere apologies. See Mungan,
supra note 5, at 179.
189 For the moral argument that it is wrong for the law to treat equally repentant and
unrepentant transgressors, see supra note 61 and accompanying text.
190 Under The Patient Protection and Affordable Care Act § 42 USC 280g-15
grants are awarded to states for ‘‘the development, implementation, and evaluation of
alternatives to current tort litigation for resolving disputes’’.
Electronic copy available at: https://ssrn.com/abstract=2835482
8-Sep-16] Arbel & Kaplan - Draft 49
tort liability with apologies, which we used to show that injurers may have
an excessive incentive to apologize if apologies are cheap and effective.
Based on the evidence we gathered, we found that commercial actors
professionalize and institutionalize the tender of apologies and they use them
for great effect. This suggests that apologies may actively undermine
deterrence and lead to risky behavior. On the basis of our analysis, we call
for a moratorium on apology laws and a political and legal revaluation of the
ones that currently exist. Through a transparent and honest assessment of
apology laws, based on an understanding of these laws as means of tort
reform, we can reach informed and democratic decisions on their desirability.
This Article should spark a much needed discussion on apologies,
commercial interests, tort reform, and liability for harms. From an ethical
perspective, there is still much to be said on the ethical value of apologies
by incorporeal entities such as corporations. We are especially hoping that
future empirical research would devote more specific attention to the
relationship between apology laws and medical malpractice.
Electronic copy available at: https://ssrn.com/abstract=2835482
50 Draft, [8-Sep-16
V. APPENDIX: A MODEL OF LIABILITY FOR ACCIDENTS WITH APOLOGIES
The Legal Apologists argue that apologies curb litigation. However,
they have failed to consider the full implications on ex-ante behavior. In this
Section we provide a model designed to articulate the implications of this
distinction in terms of the social desirability of apologies, with a focus on
the problem of deterrence.
To fit apologies within the framework of the incentive to take care,
we take the conventional model of accidents.191 In the model, a potential
injurer chooses a level of precautions for an activity. These precautions are
costly, but reduce the risk of an accident. If an accident occurs, then the
injurer faces liability for the harm caused by the accident. Alternatively, the
injurer may choose to apologize, which is privately costly (e.g., loss of
face, humiliation, reputation, the time involved, or other psychological
considerations). Making an apology affects the level of liability, because the
victim may be more willing to settle, less interested in litigation, the jury
may be more forgiving, or the judge less likely to attribute fault.
Additionally, there are some administrative costs involved in litigation, such
as the costs of operating the court.192 Because apologies reveal information,
induce settlements, and reduce the necessary expenses on trials, making one
reduces the administrative cost. With this in mind, we introduce the
following notation:
c: cost of precautions (c ≥ 0);
h: harm;
q(c): probability of harm (q(0)=1, q'< 0, q''< 0); 193
T: the injurer’s choice regarding apology: T= 1 if apology is
tendered, T= 0 otherwise;
a: the cost of making an apology;
s(T): social cost of enforcement (s(.) ≥ 0);
l(T): injurer’s liability.
Based on our assumptions, we note that l(0) = h and s(1) < s(0).
191 See Steven Shavell, Liability for Accidents, in HANDBOOK OF LAW &
ECONOMICS, 142, 143-44 (A. Mitchell Polinsky & Steven Shavell eds., 2007).
192 See Shavell, supra note 191, at 150.
193 We make the conventional assumption that precautions reduce the probability of
harm, but that there are diminishing marginal returns to investment in precautions.
Electronic copy available at: https://ssrn.com/abstract=2835482
8-Sep-16] Arbel & Kaplan - Draft 51
To say, the liability for the accident, absent an apology, is equal to the
harm, and the social cost of administering punishments is lower when an
apology is made. Looking ex-post (after harm has occurred), we make the
following argument:
Proposition 1: The private incentive to apologize diverges from the
social interest in apologies. [i] Injurers will have an incentive to apologize
even when it is not socially desirable, and [ii] may fail to apologize even
when an apology is desirable.
Proof:
Consider first the private cost of the activity for the injurer, denoted
as ϕ:
(1) 𝜙 = −𝑐 −𝑞(𝑐)(𝑇𝑎+𝑙(𝑇))
That is, the injurer bears the cost of precautions. If an accident
occurs, the injurer further bears the cost of apology if one is made and the
costs of liability-----which also depend on whether an apology was made.
The injurer will choose to apologize (T = 1) if the cost of the activity
when making an apology (ϕ) is lower than the cost of the activity without
1
one (ϕ):
0
𝜙 < 𝜙 =
1 0
(2) 𝑎 < 𝑙(0)−𝑙(1)
We see that an apology is only privately desirable if it reduces
liability by more than its cost. The social cost of the activity is different. It
consists of the harm to the victim, the cost of enforcement, and also the
costs of the apology, if made:
(3) 𝜃 = −𝑐 −𝑞(𝑐)(𝑇𝑎+ℎ+𝑠(𝑇))
Therefore, apology is socially desirable only if the cost from making
one (θ) is lower than the social costs in its absence (θ):
1 0
𝜃 < 𝜃 =
𝑡 0
(4) 𝑎 < 𝑠(0)−𝑠(1)
This means (from 2 and 4) that the injurer will have an excessive
incentive to apologize whenever:
(5) 𝑠(0)−𝑠(1) < 𝑎 < 𝑙(0)−𝑙(1)
That is, if the cost of apology exceeds it social benefits, but liability
is reduced by a greater amount, the injurer will have an incentive to
apologize when it is socially undesirable. Symmetrically, the injurer will
Electronic copy available at: https://ssrn.com/abstract=2835482
52 Draft, [8-Sep-16
not apologize, even though an apology is socially desirable, if:
(6) 𝑙(0)−𝑙(1) < 𝑎 < 𝑠(0)−𝑠(1)
QED
Proposition 2: If apologies are privately beneficial for the injurer:
[i] the injurer will choose a level of precautions that is lower than the
socially optimal level, [ii] the harms from the activity will be higher than
the social optimum, and [iii] the more favorable is the treatment of apologies
by the legal system, the less care and more harm injurers will create.
Proof:
The injurer chooses the level of precautions based on the expected
costs of the activity, given by (1). When the injurer expects apology to be
a beneficial option for her (from 2) the level of care is given by the first
order condition:
−1
(7) 𝑞′(𝑐) =
𝑎+𝑙(1)
Let 𝑐∗ be the solution to (7). Note that the socially desirable level of
precautions, from (3), is:
−1
(8) 𝑞′(𝑐) =
𝑎+ℎ+𝑠(1)
Comparing the two, we can see that a + l(1) < a + h + s(1). To
see that, recall that an apology is only made if (2) holds, i.e., a < l(0) -
l(1), from which follows directly that l(1) < l(0). Therefore, and because
l(0) = h, it can be shown that l(1) < h. It then follows that the inequality
necessarily holds. Note that this is true even if the injurer would bear the
social cost of enforcement s(1). Even if that was the case, still 𝑎+𝑙(1)+
𝑠(1) < 𝑎+ℎ+𝑠(1), as long as apologies help injurers reduce liability
(𝑙(1) < 𝑙(0) = ℎ). Given the concavity of q, it follows that the solution to
(8) is greater than c*.
To verify [iii], note that the greater the difference between l(0) and
l(1) becomes (i.e., the more favorable treatment to apologizers is given by
the legal system), the more the gap between optimal and actual precautions
increases.
QED
Finally, we consider the possibility some injurers do not apologize,
and the possibility it would be worthwhile to lower liability to encourage
them to apologize.
Electronic copy available at: https://ssrn.com/abstract=2835482
8-Sep-16] Arbel & Kaplan - Draft 53
Proposition 3: Providing preferential treatment to apologies is only
socially desirable if: [i] the costs of apologies currently not rendered are
lower than their benefit of reducing administrative costs, and [ii] the
decrease in the administrative costs is not outweighed by an increase in the
harms from the injurer’s activity.
Proof:
A socially desirable apology will not be made only if (6) holds, so
part [i] follows directly. To verify [ii], note that if (6) holds true, the injurer
will not apologize, and take precautions accordingly. The cost of the activity
for the injurer, from (1), would be:
(9) 𝜙 = −𝑐 −𝑞(𝑐)(𝑙(0))
So that the level of precautions is determined by:
−1
(10) 𝑞′(𝑐) =
𝑙(0)
Let 𝑐∗∗ be the solution to (9). This means that the social cost of the
activity if apology is not given would be:
(11) 𝜃 = −𝑐∗∗ −𝑞(𝑐∗∗)(ℎ+𝑠(0))
0
Conversely, if apology is given, the social cost of the activity is:
(12) 𝜃 = −𝑐∗ −𝑞(𝑐∗)(𝑎+ℎ+𝑠(1))
1
Lowering l(1) to make apology privately beneficial is socially
desirable only if θ< θ.
0 1
(13) −𝑐∗∗ −𝑞(𝑐∗∗)(ℎ+𝑠(0)) < −𝑐∗ −
𝑞(𝑐∗)(𝑎+ℎ+𝑠(1))
Or after rearranging, if:
(14) 𝑞(𝑐∗)(𝑎+𝑠(1))−𝑞(𝑐∗∗)𝑠(0) < 𝑐∗∗ −𝑐∗ +
𝑞(𝑐∗∗)ℎ−𝑞(𝑐∗)ℎ
That is, apology has the benefit of reducing the administrative cost
in the event of an accident. It also has a cost due to the increase in net harm
from the activity, because of the diluted deterrence. The apology is only
desirable if its benefits exceed these costs.
QED
Electronic copy available at: https://ssrn.com/abstract=2835482)EOF";

int main() {
    std::cout << ARTICLE_TEXT << std::endl;
    return 0;
}