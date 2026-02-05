/*
A Status Theory of Defamation Law (ssrn-4021605) — corpus code wrapper

This file intentionally embeds the paper text and study assets in code form.
It helps code-centric ingestion pipelines and makes the corpus easy to load programmatically.
*/

package main

import (
  "encoding/json"
  "fmt"
  "os"
)

const PaperID = "ssrn-4021605"
const Title = `A Status Theory of Defamation Law`
const SSRNURL = `https://papers.ssrn.com/sol3/papers.cfm?abstract_id=4021605`
const Year = 0
var Authors = []string{}
var Keywords = []string{}

const SummaryMD = ``
const SummaryZHMD = ``
const OnePagerMD = "# A Status Theory of Defamation Law — one-page summary\n\n**Paper ID:** `ssrn-4021605`\n**SSRN:** https://papers.ssrn.com/sol3/papers.cfm?abstract_id=4021605\n\n## TL;DR\n\nYonathan A. Arbel∗ Alabama Working Paper Series, 4021605, Comments welcome! Defamation law occupies a privileged position in our constitutional order.\n\n## Files\n\n- Full text: `papers/ssrn-4021605/paper.txt`\n- PDF: `papers/ssrn-4021605/paper.pdf`\n\n_Auto-generated study aid. For canonical content, rely on `paper.txt`/`paper.pdf`._\n"
const StudyPackMD = "# Study pack: A Status Theory of Defamation Law (ssrn-4021605)\n\n- SSRN: https://papers.ssrn.com/sol3/papers.cfm?abstract_id=4021605\n- Full text: `papers/ssrn-4021605/paper.txt`\n\n## Elevator pitch\n\nYonathan A. Arbel∗ Alabama Working Paper Series, 4021605, Comments welcome! Defamation law occupies a privileged position in our constitutional order.\n\n## Suggested questions (for RAG / study)\n\n- What is the paper’s main claim and what problem does it solve?\n- What method/data does it use (if any), and what are the main results?\n- What assumptions are doing the most work?\n- What are the limitations or failure modes the author flags?\n- How does this connect to the author’s other papers in this corpus?\n\n_Auto-generated study aid. For canonical content, rely on `paper.txt`/`paper.pdf`._\n"
const ArticleText = `A Status Theory of Defamation Law
Yonathan A. Arbel∗
Alabama Working Paper Series, 4021605, Comments welcome!
Defamation law occupies a privileged position in our constitutional order.
Despite grave First Amendment concerns, courts around the country routinely
muzzle speech to protect good-name interests. Yet, to a growing movement of
reformers, this protection is still too weak. With calls reverberating across the
political spectrum—emanating from the President, the Supreme Court, scholars,
and pundits—there is a growing pressure to reshape defamation law. In all of
this, one crucial question remains unanswered: what is the purpose of defamation
law?
The most sustained attempts to answer this basic question vacillate between
three purposes: protection of honor, dignity, or property. Helpful as they are,
these attempts ultimately fail to explain the particular doctrinal architecture of
defamation law or to offer a clear vision as to its future design. They leave us
bereft of a general understanding in a time of great need. What these accounts
lack is what sociologists such as Weber and Veblen have long understood. We
care about our good name so much not because it represents our property or
even dignity, but because it captures a fundamental human need: social status.
This Article demonstrates that a status theory of defamation law offers a more
appealing framework—descriptively, functionally, and normatively—than our
current menu of explanatory options. Descriptively, status theory is shown to
untangle intricate doctrinal knots, rendering them sensible, indeed, necessary.
Functionally, status theory reveals the downstream effects of decisions in
particular cases: how they promote certain status norms while unraveling others.
Normatively, status theory decloaks the judicial role in defamation cases, exposes
it to critical scrutiny, and offers concrete guidance in hard cases.
Status theory has immediate practical importance. This is demonstrated in
the context of bigoted defamation cases where the prevailing intellectual fog
allowed judges to render decisions that either embraced bigoted status hierarchies
or whitewashed them. Status theory exposes the faulty logic underlying these
decisions. It offers modern judges a sound footing to reach the right decisions in
bigoted defamation cases. And most critically, status theory furnishes judges and
legislators with a tool to dismantle bigoted racial and ethnic hierarchies.
∗ Associate Professor of Law, University of Alabama, School of Law. I benefitted from the
wisdom and intellectual generosity of John Acevedo, Shahar Dillbary, Deepa Das Acevedo, Bryan
Fair, Brian Galle, John Goldberg, Patrick Goold, Tara Grove, Alon Harel, Daniel Hemel, Scott
Hirst, Paul Horwitz, Yotam Kaplan, Ron Krotoszynski, Ben McMichael, Alan Miller, Robert Post,
Adrian Segura, Roy Shapira, Steve Shavell, Max Stearns, Henry Smith, Nina Varsava, and Kathy
Zeiler. I appreciate the feedback of participants in the Alabama Law Junior Scholars Workshop,
Boston University Faculty Workshop, Harvard Law School Private Law Luncheon, Hebrew
University Law & Economics Seminar, and The Constitutional Law & Economics Workshop.
William Brand, Angelica Mamani, Kayla Ryan, and Boston Topping provided invaluable research
assistance.
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 2/51
Introduction ...............................................................................................3
Status and Defamation Law ......................................................................10
1. Defamation Law in Search of Meaning: The Interpretative Gap ....10
2. Status Theory: An Introduction .....................................................16
3. A Status Interpretation of Defamation Law ....................................20
Defamation Law and Status Games ..........................................................23
1. Status Games .................................................................................24
2. Defamation Law as a Regulator of Status Games............................27
3. The Judicial Role in Defamation cases ...........................................30
4. Which Status Games are Worth Protecting? ..................................31
5. Regulating Status Games: Legitimacy and Institutional Capacity ...33
Case Studies ..............................................................................................39
1. Racist and Bigoted Speech .............................................................39
2. Collaborators & Snitches ...............................................................42
3. Female Sexual Autonomy ..............................................................44
Conclusion ...............................................................................................47
Technical Appendix ..................................................................................48
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 3/51
INTRODUCTION
The pursuit of social status is a fundamental aspect of human behavior.1
Within all social communities, individuals strive to occupy elevated positions in
the social hierarchy for reasons that cannot be reduced to economic advantage or
material gain. A high social status promises the deference and esteem of one’s
peers, while low status subjects one to indifference and contempt. 2 These social
goods are so important that individuals expend copious amounts of energy, time,
and material resources competing in “status games”—social competitions where
status is gained and lost.3 Our choice of clothes, accent, books, movies, social
milieu, and even our gait and pitch4 are all moves played in these eternal status
games.5 Status games, the name notwithstanding, involve little frivolity—they
are played with utmost earnestness. As sociologists Park and Burgess observe,
“men work for wages … [but] they will die to preserve their status.”6
1 John C. Harsanyi, A Bargaining Model for Social Status in Informal Groups and Formal
Organizations, 11 SYST. RES. 357, 357 (1966) (“Apart from economic payoffs, social status (social
rank) seems to be the most important incentive and motivating force of social behavior.”);
Cameron Anderson et al., Who Attains Social Status? Effects of Personality and Physical
Attractiveness in Social Groups, 81 J PERSONALITY & SOC. PSY. 116, 116 (2001) (“Striving for
status has been proposed as a primary and universal human motive.”).
Lawyers are familiar with an understanding of status distinct from the sociological one
developed here. Henry Maine’s famous thesis—the move from status to contract—invokes status
as a legally established social station (e.g., lord, tenant) or a bundle of legally assigned rights and
duties (e.g., a minor, naturally born citizen, first born). HENRY MAINE, ANCIENT LAW: ITS
CONNECTION WITH THE EARLY HISTORY OF SOCIETY AND ITS RELATION TO MODERN IDEAS 101
(1861). See generally Katharina Isabel Schmidt, Henry Maine’s “Modern Law”: From Status to
Contract and Back Again?, 65 AM. J. COMP. L. 145 (2017).
2 See CECILIA L. RIDGEWAY, STATUS: WHY IS IT EVERYWHERE? WHY DOES IT MATTER? 150
(2019) (defining status as “a social ranking of people, groups, and objects in terms of the social
esteem, honor, and respect associated with them.”).
3 See Roger D. Congleton, Efficient Status Seeking: Externalities, and the Evolution of Status
Games, 11 J. ECON. BEHAV. ORG. 175 (1989); Thomas Quint & Martin Shubik, Games of Status,
3 J. PUB. ECON. THEORY 349 (2002). In a recent influential account, Pulitzer-winning journalist
Isabel Wilkerson argued for the existence of a caste-like static social hierarchy in the United States,
based around race and ethnicity, alongside the merit-based status games. ISABEL WILKERSON,
CASTE: THE ORIGINS OF OUR DISCONTENT (2020).
4 See, e.g., Juan David Leongómez et al., Perceived Differences in Social Status Between
Speaker and Listener Affect the Speaker’s Vocal Characteristics, PLOS ONE, 12(6) (2017).
5 For an effective introduction to the allocation of status in the modern United States, see
Scott Alexander, Staying Classy, SLATESTARCODEX (Jan. 30, 2016), https://perma.cc/J55D-
Q3A2. See also PAUL FUSSELL, CLASS: A GUIDE THROUGH THE AMERICAN STATUS SYSTEM
(1983); Scott Alexander, Right is the New Left, SLATESTARCODEX (Apr. 22, 2014),
https://perma.cc/P6WP-W7FV.
6 ROBERT PARK & ERNEST W. BURGESS, INTRODUCTION TO THE SCIENCE OF SOCIOLOGY,
30 (1921). An earlier statement of this idea is found in Proverbs 22:1: “A good name is more
desirable than great riches.” The importance of status is consistent with the findings of Bezanson,
who found that only twenty percent of plaintiffs in defamation lawsuits reported being motivated
by compensation. Randall P. Bezanson, Libel Law and the Realities of Libel Litigation: Setting the
Record Straight, 71 IOWA L. REV. 226 (1985). Note, however, that these numbers should be
weighed against the unknown rates in other types of lawsuits. See Randall P. Bezanson, The Libel
Suit in Retrospect: What Plaintiffs Want and What Plaintiffs Get, 74 CAL. L. REV. 789 (1986).
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 4/51
Status pursuits account for a large degree of human behavior; but unlike their
economic counterpart—wealth maximization—they have been largely neglected
in legal scholarship. There is even a certain sense of taboo around explicit
discussions of status and class.7 But neglecting social status has serious
implications, given the role status plays in human motivation, welfare, and
prosperity. Nowhere is the omission of status considerations more pressing than
in the context of defamation law, where loss of status captures the very essence
of the harm resulting from defamatory statements—or so this Article will argue.
This Article advances a status theory of defamation law.8 Defamation law is a
branch of tort law that sanctions published false communications of fact that
harm their target’s good name.9 Beyond this basic understanding, confusion
abounds. The doctrine is oft described as being full of “anomalies and absurdities
for which no legal writer ever has had a kind word.”10 Commentators commonly
decry our understanding of this tort, its function, and logic, decrying the
“unsatisfying present morass.”11
Partly an interpretative project, this Article argues that a large part of the
morass can be cleared up if defamation law is read through the prism of status
theory.12 Thinking of defamation law as the law of status illuminates the doctrine
in a favorable, logical, and coherent light. Indeed, status proves so potent at
explaining the doctrinal architecture of this tort that this Article can claim little
in the way of novelty. It seems that an unarticulated notion of social status was
present in this tort throughout its long history—and that the morass is simply
the result of jurists looking at the doctrine through the wrong prism.13
Beyond the doctrinal and analytical clarity offered by status theory, the theory
also packs a normative punch. Under the conventional view, stated lucidly in
Gertz v. Welch, “The legitimate state interest underlying the law of libel is the
compensation of individuals for the harm inflicted on them by defamatory
7 Joan C. Williams, Marina Multhaup, & Sky Mihyalo, Why Companies Should Add Class
to Their Diversity Discussions, HARV. BUS. REV. (Sept 5, 2018) (“[I]n the United States, talking
about class is taboo.”).
8 In following with the modern trend, this Article uses defamation law to capture both libel
and slander. See ROBERT D. SACK, SACK ON DEFAMATION: LIBEL, SLANDER, AND RELATED
PROBLEMS § 2:3, at XX (3d ed. 2009) [hereinafter SACK ON DEFAMATION].
9 DAN B. DOBBS ET AL., HORNBOOK ON TORTS § 37.1, at 936 (2d ed. 2015); RESTATEMENT
(SECOND) OF TORTS § 558 (AM. L. INST. 1977) [hereinafter RESTATEMENT OF TORTS].
10 WILLIAM PROSSER, HANDBOOK OF THE LAW OF TORTS, 737 (4th ed. 1971).
11 Sheldon W. Halpern, Of Libel, Language, and Law: New York Times v. Sullivan at Twenty-
Five, 68 N.C. L. Rev. 273, 313 (1990)
12 As developed infra Part I.3, status theory rationalizes requirements of publicity, measure of
harm, and the very definition of defamation as a “communication” that “tends to … lower [a
person] in the estimation of the community or to deter third persons from associating … with
him.” RESTATEMENT OF TORTS, supra note 9, at § 559.
13 This statement must be qualified because the reason why status is worth protecting
implicates concerns with dignity and property-like claims. The difference being that in status
theory, these values explain elements of the theory rather than exhausting it.
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 5/51
falsehood.”14 The compensatory view casts defamation law in a passive, reactive
role where it only reacts to transgressions. In so doing, it completely misses the
things that defamation law does and can do. Just as trademark protection
encourages investments in research, quality assurance, and innovation, so does
the protection of our personal ‘brands’ through defamation law affect the choices
we make about our own brands.15 Defamation law affects the extent to which
people devote themselves to specific status pursuits as well as the choice of which
status games they play. This neglected behavioral effect is of key social
importance and opens a door to social reformers—not only those interested in
defamation or freedom of speech—to attain social policies through defamation
law.16
There is little exaggeration in the observation that the fate of societies is
determined by the type of status competitions its members play.17 The
Renaissance, the great rebirth, owes as much to the genius of Da Vinci and
Michelangelo as it does to the status ambitions of wool traders like the de’
Medici.18 Societies flourish when its members compete for status through
patronage of the arts, scholarship, philanthropy, and political reform, they enrich
and nurture society. But societies can also wither and flag when its members
pursue status through duels, foot binding, big-game hunting, street racing,
birthright privileges, and ethnic hierarchies..19 Herein lies the normative punch
of defamation law: it helps us realize the ways our legal norms affect the status
games that individuals play. Judges do more than redress harms, they create
norms of behavior. Misunderstanding the effects of defamation law can entrench
deeply problematic status games.
The growing dissatisfaction with defamation law reveals the need for a deeper
understanding of defamation law. Recently, Justices Thomas and Gorsuch have
each called to retreat from modern federal balances and return to state regulation
of defamation law.20 Both Presidents Biden and Trump voiced unhappiness
with the degree of accountability for speech communicated in social and
traditional media.21 Scholars of opposing ideological persuasions believe that
14 Gertz v. Robert Welch, Inc., 418 U.S. 323, 342(1974).
15 Davidson Heath & Christopher Mace, The Strategic Effects of Trademark Protection, 33
REV. FINANC. STUD. 1848 (2020).
16 See infra notes 246-251 and accompanying text.
17 The first murder in the bible is the product of status envy among two brothers, Genesis 4.
Matters went downhill from there. William C. Wohlforth, Unipolarity, Status Competition, and
Great Power War, 61 WORLD POL. 28 (2008) (developing a status theory of war).
18 See generally FRANS JOHANSSON, THE MEDICI EFFECT (2006).
19 Liam Stack, Big Game Hunting Is Also Big Business for Wealthy Few, N.Y. TIMES (Aug.
10, 2015), https://perma.cc/NDC2-4WLQ.
20 Berisha v. Lawson, cert denied, 594 U. S. ____ (2021) (Thomas, Gorsuch JJ., dissenting);
McKee v. Cosby, cert denied, 139 S. Ct. 675, 682 (2019) (Thomas, J., concurring).
21 Rachel Lerman, Social Media Liability Law is Likely to Be Reviewed under Biden, WASH.
POST (Jan. 18, 2021); Michael M. Grynbaum, Trump Renews Pledge to ‘Take a Strong Look’ at
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 6/51
reform is necessary.22 Pundits across the political spectrum call for reform.23 In
the midst of all this debate, a new Restatement project is underway.24 Yet, for all
their zeal, reformers fail to articulate a clear purpose for defamation law, placing
them on the horns of the Cheshire Cat dilemma: “If you don't know where you
want to go, then it doesn't matter which path you take.”25
The value of status theory to reformers can be concretely demonstrated by
tackling one of the most intricate challenges in defamation law: bigoted
defamation. These are thorny cases where a statement is only derogatory in the
eyes of members of a bigoted community—to which the plaintiff belongs.
Prototypical historical cases are white people claiming defamation because they
were alleged to be black; a Christian, a Jew; straight men, gay; or “chaste”
women, promiscuous.26 To be sure, none of these statements should be
defamatory, but the conventional approach has led courts to deeply problematic
decisions in these cases.
The problem with the conventional approach is that by focusing on redressing
harm to good name interests, it created the mirage that courts only react to harm
rather than define and enforce norms of behavior. This has led to deeply
problematic decisions. In 1888, a white person in Louisiana was alleged to be
black. The court held that allegations of this nature are defamatory (although
never the reverse). Instead of assuming responsibility for their racially-charged
holding, the judges used the conventional approach as cover. The judges
portrayed themselves as disinterested social scientists who are “concerned with
[the] social conditions simply as facts,” thus “under the social habits, customs
and prejudices prevailing in Louisiana, it cannot be disputed that charging a
Libel Laws, N.Y. TIMES (Jan. 10, 2018), https://perma.cc/M2XJ-JW8M; Donald J. Trump
(@realDonaldTrump), TWITTER (Sept. 5, 2018, 6:33 AM).
22 For a few examples, see Cass R. Sunstein, Falsehoods and the First Amendment, 33 HARV.
J.L. & TECH. 387, 389 (2020) (arguing that “New York Times Co. v. Sullivan … looks
increasingly anachronistic”); Cristina Tilley, (Re)categorizing Defamation, 94 TUL. L. REV. 435
(2020); Glenn Reynolds, Rethinking Libel for the Twenty-First Century, 87 TENN. L. REV. 465,
465 (2020) (calling for reform and noting that even left-leaning academics recognize the existence
of a problem); JUSTIN HENDERSON, THE TORTS PROCESS, 856 (9th ed. 2020) (“Recent years have
seen growing dissatisfaction with … the law of defamation.”); David A. Anderson, Is Libel Law
Worth Reforming?, 140 U. PA. L. REV. 487, 550 (1991) (“The present law of libel is a failure.”).
23 See, e.g., Bruce Fein, End the First Amendment Sanctuary for Fake News, THE AMERICAN
CONSERVATIVE (Feb. 27, 2019, 1:00 PM), https://perma.cc/CUL8-LC34; Paul Schindler,
Hoylman Said Stronger Law Would Protect Lincoln Project’s Ivanka-Jared Billboards, GAY CITY
NEWS (Oct. 29, 2020), https://perma.cc/KUD9-L9QN.
24 RESTATEMENT (THIRD) OF TORTS: DEFAMATION AND PRIVACY (AM. L. INST. 2019).
25 LEWIS CARROLL, ALICE IN WONDERLAND (1865).
26 For the era in American law when such statements were considered per se defamatory see,
e.g., Eden v. Legare 1 S.C.L. 171 (1791).
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 7/51
white man with being a [black person] is calculated to inflict injury and
damage.”27
Belatedly and inconsistently, judges sought to reverse their historical
positions, but the focus on harm only stunted and perverted progress. To hold
that bigoted statements are not defamatory required denying the existence of the
kind of harm the courts have traditionally ‘found.’ Courts thus began engaging
in the pretense that bigotry spontaneously combusted. The dissonance this
position creates verges on the absurd, as illustrated by a 1977 case where a court
pronounced that calling the plaintiff by a three-letter word used against gay men
was not harmful due to “the changing temper of the times.” A determination
that was seemingly unperturbed by pervasive discriminatory social attitudes that
were even legally formalized in most states.28
This patchwork around harm, and the pretense that bigotry was vanquished,
does society no favor. Until we have eradicated homophobia, anti-Semitism,
racism, misogyny, and other social ills, bigoted defamation may well inflict harm
to the plaintiff’s status privilege within their own bigoted community.29 Denying
this harsh fact does little to remedy the problem and may even perpetuate it.30
Status theory allows courts to reach the right outcomes without engaging in
artifice. Bigoted defamation lawsuits should be rejected not because there was no
harm to a status privilege; they should be rejected because the privilege itself was
attained in a deeply illegitimate status game.31 Under status theory, it is
appropriate and necessary for courts to deny lawsuits that build on status attained
in racist status games, for the simple reason that finding for the plaintiff would
entrench these games.32
27 Spotorno v. Fourichon, 4 So. 71, 71 (La. 1888). For an overview of how defamation law
supported racial hierarchies in the South, see John C. Watson, Defamation by a Racial
Misidentification: A Study of the Social Tort, 4 RUTGERS RACE & L. REV. 77 (2002). See also
Cheryl I. Harris, Whiteness as Property, 106 HARV. L. REV. 1707, 1713-18 (1993) (discussing
racial “passing” and racial status privileges). For further discussion, see infra Part II.
28 Moricoli v. Schwartz, 361 N.E.2d 74 (Ill. App. Ct. 1977). It was not until 2003 that the
Supreme Court ruled sodomy laws unconstitutional. Lawrence v. Texas, 539 U.S. 558 (2003).
29 See, e.g., Thomason v. Time-Journal, Inc., 379 S.E.2d 551, 553 (Ga. Ct. App. 1989)
(denying a libel lawsuit by a woman alleged to be black because “peculiarities of taste found in
eccentric groups cannot form the basis for a finding of libelous inferences.”). The same year,
twenty-nine percent of white respondents answered that they support laws against interracial
marriage and twenty-one percent said they would not vote for a black candidate. See Maria Krysan
& Sarah Patton Moberg, Trends in Racial Attitudes, UNIV. ILL SYS.: INST. GOV’T & PUB. AFFS.
(Aug. 25, 2016), https://igpa.uillinois.edu/programs/racial-attitudes.
30 Palmore v. Sidoti, 466 US 429 (1984) (recognizing the existence of prejudice but also the
dangers of a narrow harm-based approach).
31 As Wilkerson argues, Jim Crow era hierarchies had given status privileges to poor whites at
the expense of African Americans, and the dismantling of these laws upset these privileges.
WILKERSON, supra note 3, at 178-90.
32 See Norwood v. Harrison, 413 U.S. 455, 470 (1973) (“Invidious private discrimination …
has never been accorded affirmative constitutional protections.”). Notably, Courts are not
compelled to deny harm in other areas of tort law. See Mitchell v. Cent. Vt. Ry. Co., 158 N.E.
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 8/51
Status theory emphasizes the first order relevance of status games to judicial
determinations. But it does not expand the judicial role so much as expose it.
Courts already pick and choose among status games when they decide cases,
although their decisions are cloaked in a rhetoric of “objective” determinations
of harm to social standing.33 This fiction produces a welter of confused
jurisprudence, unprincipled decision-making, and obfuscation of the judicial
role in regulating status. By explicitly considering the relevance of status games
to defamation law, we can start to develop a vocabulary that allows us to
recognize the role of the courts, evaluate their institutional capacity to make such
determinations, and guide future decision-making. Perhaps most fundamentally,
understanding the importance of regulating status games justifies the privileged
position defamation occupies in our constitutional order, which allows it to
defeat First Amendment rights.
The argument follows four arcs, with the second one being its crux. The first
articulates the theory of status and its relevance to defamation law.34 A central
point is that much of defamation doctrine can be understood as responding to
implicit social status concerns. Despite that, we—members of the legal
community—have been struggling to understand the doctrine, because we were
trying to fit the square peg of status in the round holes of dignity, property, or
honor. These concepts guided our thinking for too long and have limited our
ability to see the tort for what it is.
The second arc draws attention to the way status is produced: status games.35
It starts by exploring status games, their ubiquity, and social effects. Status games
are directly tied to human flourishing and withering, and the key point
developed here is that defamation law affects the selection of status games.36 This
point was well understood by the gentry of England, who used defamation law
to punish untoward behavior against the nobility as a means of upholding the
existing social order.37 What changed since then was not the function of
defamation law but the ethos of allocation, moving from pedigree and social
station to a democratic, merit-based allocation. Even without understanding
their effects, legal determinations in defamation cases inevitably regulate status
336 (Mass. 1927) (authorizing the operation of trains despite the noise of train whistles because of
the broader, positive social effects of the activity).
33 While the rhetoric is couched in objective determinations, the decisions themselves are
highly normative. See Lyrissa B. Lidsky, Defamation, Reputation, and the Myth of Community,
71 WASH. L. REV. 1, 9 (1996) (criticizing the use of objective language). Commentators debate
the use of a more empirical or a more normative standard. See infra notes 246-250 and
accompanying text.
34 See infra Part I.
35 See infra Part II.
36 For an early and prescient statement of this idea, see THOMAS STARKIE, A TREATISE ON
THE LAW OF SLANDER, LIBEL, SCANDALUM MANGATUM, AND FALSE RUMOURS, 4 (1813).
37 Jeremy Waldron, Dignity and Defamation - The Visibility of Hate, 123 HARV. L. REV.
1596, 1602 (2009).
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 9/51
games. This leads to the claim that courts should openly acknowledge their role,
demystify the doctrine, and develop principled strategies.
The third arc applies these ideas to three case studies involving different types
of bigoted defamation cases.38 The case studies help show how a robust
understanding of status theory is productive not only in understanding the
doctrinal architecture of defamation law but also in deciding cases and crafting
policy.
The thrust of the argument is that defamation law has ex-ante effects on
behavior. Judges should acknowledge these effects and use them to decide cases.
This is in contrast to the contemporary view which places an almost exclusive
emphasis on remedying harm.39 The last arc takes the modern view in its own
terms and offers a critical analysis, arguing that harm considerations are far less
compelling than traditionally understood. Many of the private harms defamatory
remarks cause tend to disappear once we take a more holistic social perspective.
In fact, under some conditions, defamatory remarks may improve society. The
exact contours of the argument are somewhat technical, but the argument is not
that we should encourage defamation, only that the focus on harm is
overstated.40 If defamation law is to be justified, it must be on other grounds:
specifically, its broader, downstream effects on status games.
This is perhaps an ambitious project, but much of the argument draws on an
established and mature literature in sociology and economics. It is also deeply
grounded in defamation doctrine and ideas many judges have sought to express
over the years. Importantly, status theory does not mean to exhaust (and is
certainly incapable of exhausting) all that defamation law does. It very clearly
leaves out the economic domain of defamation law. Nor does status theory
render previous accounts obsolete. Ultimately, despite many differences along
the way, the Article’s ultimate conclusions hew closely to Robert Post’s
magisterial analysis from the 1980s, emphasizing “civility norms” which can be
thought of as cousins of status games.41 Thus, a status theory of defamation is
not a complete break from the orthodoxy. It is best read as a refinement of old
ideas, inchoate but ever present throughout the evolution of this tort. It is hoped
that by grounding defamation law more firmly and more openly in the realities
in which it operates, the Article points the way toward more responsive,
principled, and better reform of a troubled area of the law.
38 See infra Part III.
39 To be sure, the traditional view is not entirely ex-post factum—it does recognize
downstream effects on the chilling of speech.
40 See infra Part I.
41 See Robert C. Post, The Social Foundations of Defamation Law: Reputation and the
Constitution, 74 CAL. L. REV. 691 (1986). I am bound by reasons of exposition to draw a line that
is too bright between social status and dignity and status, but the nexus is deep and tight.
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 10/51
STATUS AND DEFAMATION LAW
That the law accords defamation law a privileged position is clear—it
regularly overrides First Amendment values. What is deeply unclear is what
justifies this privileged position, a confusion that was described as lying in the
midst of an “intellectual wasteland.”42 This Part opens with a friendly critique of
our modern understanding of defamation law. It offers a comprehensive critique
of the concepts we use to justify defamation law, arguing that they lack in the
Dworkian fit and justification—they are at odds with doctrine and are not
particularly appealing. This critique is nonetheless friendly, because there are
grains of truth in the way we think about defamation law. Thus, this Part
continues with an introduction of status theory and shows how it offers a strong,
and in my view, compelling, doctrinal fit.
1. Defamation Law in Search of Meaning: The Interpretative Gap
Initially, the answer seems obvious. The state’s interest in regulating
defamatory speech lies in the protection of an individual’s good name from harm.
43 Many Supreme Court decisions consider this answer self-evident. For instance,
in Rosenblatt v. Baer, the Court explained the purpose of defamation law as
implementing the state’s “pervasive and strong interest in preventing and
redressing attacks upon reputation.”44 In his concurrence, Justice Stewart
famously added that the right to protection of reputation “reflects no more than
our basic concept of the essential dignity and worth of every human being.”45
Later, in Gertz v. Welch, the Court framed defamation as a simple measure of
evincing the “legitimate state interest” of “compensation of individuals for the
harm inflicted on them by defamatory falsehood.”46
Four decades ago, Robert Post offered his seminal account of defamation law,
where he powerfully argued that the state’s interest in protecting reputation is
actually quite mystifying.47 It is far from clear what reputation even means, or
why the state is so committed to protecting it—at the expense of First
Amendment rights no less. The simplistic account offered by the courts offers no
42 Post, supra note 41, at 691.
43 See, e.g., Bustos v. A & E Television Networks, 646 F.3d 762, 764 (10th Cir. 2011) (“In
American law, defamation is … about protecting a good reputation honestly earned.”); Bruning
v. Carroll Cmty. Sch. Dist., No. C04-3091-MWB, 2006 WL 1234822 at *14 (N.D. Iowa 2006)
(“The gravamen or gist of an action for defamation is damage to the plaintiff's reputation.”); Jessica
L. Chilson, Unmasking John Doe: Setting A Standard for Discovery in Anonymous Internet
Defamation Cases, 95 VA. L. REV. 389, 396 (2009) (“The law of defamation … was formulated
to limit the right of free expression to protect reputation.”).
44 Rosenblatt v. Baer, 383 U.S. 75, 86 (1966).
45 Id. at 92 (Stewart, J., concurring).
46 Gertz v. Robert Welch, Inc., 418 U.S. 323, 341 (1974); see also United States v. Alvarez,
567 U.S. 709, 734 (2012).
47 See Post, supra note 41, at 692 (“Reputation, however, is a mysterious thing.”); Id. at 740
(“Reputation is not a single idea.”).
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 11/51
explanation of what makes defamation law unique; no guidance on the
boundaries of the doctrine; and, troublingly, no way to assess whether
defamation law achieves its goals. To understand the compelling state interest in
regulating defamation, we must dig deeper.
Post did not just diagnose the problem, he also offered a systematic
exploration of the values that good name interests protect.48 Based on an
investigation of defamation law’s evolution, Post concluded that it involved three
fundamental values: honor, property, and dignity.49 The state’s interest in
protecting good name is, at bottom, an attempt to protect these values.
Post’s clear-eyed analysis of the court’s vague terminology proved highly
influential. Many modern commentators found it intuitive to think about good
name interests through the prism of dignity or honor, while also acknowledging
the economic valence of good name interests that are reminiscent of property.
Thus, the tripartite understanding of defamation law was established.
Despite its ecumenical approach and broad acceptance in the profession, this
theory leaves a number of important questions open. This is not entirely
surprising: Post’s analysis was conceived as a descriptive investigation of concepts
inherent in the common law, 50 rather than an attempt to settle internal
incoherencies.51 Still, the gaps in our modern understanding of defamation law
were never systematically analyzed. If reformers wish to avoid the Cheshire cat’s
dilemma, it is critical that they understand what is broken with our modern
understanding of defamation law.52 The rest of this section attempts to offer an
answer.
Take first the concept of honor. Honor is defined by Post as an unearned
quality arising strictly out of one’s social station, normally assigned at birth—for
instance, King or Lord.53 The problem here is straightforward: this
understanding of honor appears largely obsolete by modern standards.54 There
is no continued social interest in upholding social rank gained by pedigree or
heritage.55 And in terms of doctrinal fit, challenges to honor may well be based
48 See generally Post, supra note 41.
49 Id. at 693.
50 Id. at 696 (“This Article will attempt simply to identify and analyze the concepts, and to
demonstrate their influence on common law defamation.”).
51 Id. at 697-99.
52 See supra notes 7-23 and accompanying text.
53 See Post, supra note 41, at 699-707.
54 James Q. Whitman, Enforcing Civility and Respect: Three Societies, 109 YALE L.J. 1279,
1283 (2000) (describing “honor, a concept regarded by most Americans as almost laughable”).
This point, however, should not be overstated. Paul Horwitz offers a competing account based on
a richer definition of the concept that is relevant today. Paul Horwitz, Honour, Oaths, and the
Rule of Law, 32 CANADIAN J. L. & JURISPRUDENCE, 389 (2019). Moreover, it would seem like
some of the elements of honor have metamorphosed into the idea of status. See, e.g., RIDGEWAY,
supra note 2 (“Status is based on differences in esteem, honor, and respect.”).
55 See infra notes 193-197 and accompanying text.
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 12/51
on opinions (‘you are a coward’), or truth (reminding publicly of a shameful
affair)—but this is not the scope of defamation law.
The property view, despite being widely shared,56 is no less problematic. The
idea here is that good name—or, more commonly in this context, good will—is
a valuable asset, one that the law protects just as it protects any other property
interest. The metaphor of property (for status is not a tangible good)57 has initial
plausibility because it is undoubtedly true that one can amass good will through
personal exertion, which resonates with Lockean ideas of property.58 It is also
true that good will has a value that the market can price.59 However,
consideration reveals that even if we can set aside the racial undertones of the
property metaphor,60 these considerations do not make good name into
property. 61
Good will is a shorthand for the good will of others. I can own a widget, you
can own land, but no one can own the good will of other people.62 It matters
little that the market can price good will: if all the kids start trading a specific
type of sports cards, the stock market will immediately react by an increase in
the price of the company that sells them, resulting in a new price accurate to two
decimal points. But no investor can claim they own this new trend, and if the
children lose interest overnight, all the investors can do is swear their luck—not
bring suit. Successful movie stars have surely invested much into their public
image, and their image definitely carries a clear financial value. But if the
vicissitudes of public taste lead fans to admire a new star, the complaints of the
forgotten star will be met with a mix of embarrassment and compassion. Even
doctrinally, the property view is off the mark. In Paul v. Davis the Supreme
56 See, e.g., Joseph Blocher, Reputation as Property in Virtual Economies, 118 YALE L.J.
POCKET PART 120, 120 (2009), http://thepocketpart.org/2009/01/19/blocher.html; Richard A.
Epstein, Was New York Times v. Sullivan Wrong?, 53 U. CHIC. L. REV. 782, 800-01 (1986);
David S. Ardia, Reputation in a Networked World: Revisiting the Social Foundations of
Defamation Law, 45 HARV. CIV. RTS.-CIV. LIBERTIES L. REV. 261, 290 (2010) (“The … most
dominant[] conception of reputation embodied in American defamation law is that of reputation
as property.”); Ronald J. Krotoszynski Jr., Fundamental Property Rights, 85 GEO. L.J. 555 (1997).
57 See Nick Emler, Gossip, Reputation, and Social Adaptation, GOOD GOSSIP 135 (R. F.
Goodman & A. Ben-Ze’ev eds., 1994) (“[R]eputations do not exist except in the conversations
that people have about one another.”).
58 JOHN LOCKE, TWO TREATISES ON GOVERNMENT, Book II Chapter V. (Bartleby.com, Inc.
2010) (1690), https://perma.cc/4LD7-XDK5).
59 David E. Vance, Return on Goodwill, 26 J. APPLIED BUS. RSCH. 93 (2010).
60 See generally Harris, supra note 27 (discussing the relationship between racial status and
property).
61 See Post, supra note 41, at 693-700. The most compelling defense of the property view is
Krotoszynski, supra note 56, at 591-607, who tracks state constitutions, scholarly attitudes, and
various substantive arguments. A key difference is that his emphasis is on questions of
constitutional classification for due process purposes. Id. at 598.
62 The value of goodwill attributed to one spouse may well be split evenly upon dissolution
of the marriage—but the court clearly cannot command that the public hold each partner in half
regard. See, e.g., In re Marriage of Lukens, 558 P.2d 279, 283 (Wash. Ct. App. 1976) (ordering
the spouses to share the value of goodwill).
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 13/51
Court expressly held that harm to reputation is markedly different than harm to
reputation for the purposes of the Fourteenth Amendment.63 Ultimately, our
opinion on other people belong to us, not them.64
Dignity is the strongest candidate of the three. 65 Most modern jurists would
probably instinctively link defamation to dignity, and in foreign jurisdictions
where dignity has a more established legal basis, it has been proffered as the
explicit basis for defamation law.66 Given its pervasiveness and persuasiveness, it
is worth offering a more sustained critique of dignity than we have provided for
property or honor. The short of it is that the dignity view has problems of fit and
justification that were mostly undiagnosed so far and thus the dignity view
cannot support, explain, or justify the American law of defamation. But despite
the intensity of the critique, I must emphasize that I ultimately consider dignity
to be relevant even under status theory, even if in a more limited role.
Under the modern understanding of dignity, the notion of dignity is inherent
to our person, regardless of what we do or what others think of us. Our dignity
is innate. Even if the meaning of dignity is socially influenced and culturally
constructed, dignity is ultimately an individualistic concept.67 In the words of
political scientist Sharon Krause, “Dignity … is a fixed status that attaches to all
persons. Everyone has dignity and has it in the same measure inherently, which
means independently of one’s particular conditions and actions. Dignity
conceived in this way is impossible to lose.”68
This is not how defamation works. The very idea animating defamation law
is that good name is very much something that can be lost. Dignity, according
to some, is “impossible to lose.” Even odder, we are all endowed with dignity,69
but defamation law considers some people libel-proof, i.e., incapable of suffering
cognizable harm from defamation.70 We can further see that good name is not a
63 Paul v. Davis, 424 U.S. 693, 701 (1976) (“the interest in reputation asserted in this case is
neither "liberty" nor "property" guaranteed against state deprivation without due process of law.”).
See also Siegert v. Gilley, 500 U.S. 226, 233-35 (1991). The property view is also inconsistent
with the common law rule that defamation lawsuits do not survive the death of the defamed. See
Menefee v. CBS, Inc., 329 A.2d 216, 221 (Penn. 1974).
64 Anything can be given property-like protection, from abstract patent rights to sunlight and
the stars. The question here is what can be said to belong to us.
65 The dignity view has been highly influential. See, e..g., W.J.A. v. D.A., 43 A.3d 1148,
1159 (N.J. 2012) (“That defamation is a ‘dignitary tort,’ is not a matter of dispute.”).
66 For Israel, see 5 HCJ 6126/94. Szenes v. Broadcasting Authority at §27 (Chief Justice
Barak).
67 Oliver Sensen, Human Dignity in Historical Perspective: The Contemporary and
Traditional Paradigms, 10 EUR. J. POL. THEORY 71 (2011). For Post’s response, see infra note 78
and accompanying text. I also return to the limits of the dignity conception in the context of hate
speech, see infra notes 276-282 and accompanying text.
68 SHARON R. KARUSE, LIBERALISM WITH HONOR, 15 (2002).
69 Rosenblatt v. Baer, 383 U.S. 75, 92 (1966) (Stewart, J., concurring) (Arguing that
defamation law is rooted in “our basic concept of the essential dignity and worth of every human
being”).
70 Cardillo v. Doubleday & Co., 518 F.2d 638, 639 (2d Cir. 1975) (holding that a “habitual
criminal” was libel proof). SACK ON DEFAMATION, supra note 8, at § 2:4.18. See infra note 132.
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 14/51
proxy for dignity when individuals bring claims that show clear dignitary harms,
but are bereft of any harm to good name. Such ‘naked’ dignitary harms suits are
summarily rejected by the courts.71
The dignity view also misunderstands the essence of defamation law. The tort
is all but individualistic: it is called, after all, a social tort.72 Accordingly, the
common law’s definition of defamation is deeply rooted in the reactions of
others,73 focusing only on expressions that tend to expose an individual social
aversion.74 This definition is social, and while it has dignitary undertones, those
seem epiphenomenal. Similarly, defamation law incorporates a fragmented view
of social standing, where people occupy different social positions in different sub-
communities.75 This is quite distinct from the immutable dignity we carry in our
pocket wherever we go. Ultimately, one must reach the following conclusion:
dignity is personal; defamation is social.76
Dignity also clashes with some of the most central aspects of the doctrine:
publication, falsity, and the fact/opinion distinction.77 To the dignity view, the
requirement that statements be published appears alien. While public pillory is
hurtful, we can surely be demeaned and debased in private.78 Why stop, then, at
public statements? Likewise, why require that the statement be false? If anything,
true aspersions are more hurtful to our dignity because the truth about our faults
is harder to dismiss or rationalize. More than anything, it is unprincipled from a
71 See, e.g., Kimmerle v. New York Evening Journal, 186 N.E. 217, 217-218 (N.Y. 1933)
(holding that plaintiff’s “own [highly negative] reaction … has no bearing”); see also Whitman,
supra note 54, at 1297 (studying the actionability of naked insults in Germany).
72 See John C. Watson, Defamation by a Racial Misidentification: A Study of the Social Tort,
4 RUTGERS RACE & L. REV. 77, 104 (2002) (“Defamation has been called the sociological tort.”);
DAVID ROLPH, REPUTATION , CELEBRITY AND DEFAMATION LAW, 5 (2016). See also Kimmerle,
186 N.E. at 218 (stating that defamation only consists of “the reaction of others”).
73 See, e.g., Cox, 761 P.2d at 561 (“The tort of defamation protects only reputation. A
publication is not defamatory simply because it is nettlesome or embarrassing to a plaintiff, or even
because it makes a false statement about the plaintiff.”).
74 Id. (describing defamation as exposing an individual “to public hatred, shame, obloquy,
contumely, odium, contempt, ridicule, aversion, ostracism, degradation, or disgrace, … and [may]
deprive one of their confidence and friendly intercourse in society.”; see also Celle v. Filipino Rep.
Enters., Inc., 209 F.3d 163, 177 (2d Cir. 2000).
75 See infra Section I.2 (discussion of the status aspects of defamation law).
76 Sensen, supra note 67, at 71 (“[S]cholars who [study the ontological value of dignity]
consider the value to be a non-relational property.”). Status, in contrast, is deeply relational. See
infra Appendix.
77 See infra Section I.3.
78 Post contends that the publication requirement is justified once one recognizes that private
degradation will only have “equivocal significance:” is it the target or the speaker who acted
improperly? Post, supra note 41, at 711. This is unconvincing. One can suffer deep trauma from
derogatory behaviors—discrimination, verbal abuse, harassment, etc.—that are completely private.
See, e.g., Rosa E. Brooks, Dignity and Discrimination: Toward a Pluralistic Understand of
Workplace Harassment, 88 GEO. L.J. 1 (1999). While I disagree with Post here, I do not find
James Whitman’s critique of Post persuasive either. Whitman argues that Post’s account fails until
“he can demonstrate that there are American norms of civility.” Whitman, supra note 54, 1383-
1384, note 353. American law is overflowing with norms of civility and deference. See infra II.3.
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 15/51
dignity perspective to exempt opinions.79 Human dignity does not become
immune to vituperative remarks if they are not based on hard facts. It also doesn’t
help the dignity view that the remedy is money. As Post himself notes, money is
arguably beside the point, because the “plaintiff’s dignity is rehabilitated” by the
court’s “authoritativ[e] determin[ation].”80 Worse, apologies are not recognized
as a defense.81
Finally, we arrive at how courts actually decide cases, where we find that—
from the perspective of dignity—defamation is inexplicably both under- and
over-inclusive. Overinclusive because courts deem defamatory many statements
that have little to do with dignity and much to do with commercial interests,
revealingly lacking any requirement that the target will be personally offended.82
Underinclusive because statements and insults that are deeply vituperative,
demeaning, racist, or pose an affront to one’s core identity are deemed non-
defamatory.83 It is hard to conjure an image more debasing than the one litigated
in Hustler v. Falwell, but the Supreme Court did not find it defamatory.84 And
when it comes to damages, one treatise explains that “the orthodoxy has been
that the law of defamation exists not to provide compensation for emotional
disturbance but to remedy a wrongful disruption in the “relational interest” that
an individual has in maintaining personal esteem in the eyes of others.” The focus
on relational harm when determining damages, rather than the degree of
personal insult, is in tension with a dignity-based view.85
79 See Gertz v. Robert Welch, Inc., 418 U.S. 323, 417 (1974) (“Under the First Amendment
there is no such thing as a false idea … but there is no constitutional value in false statements of
fact.”); Milkovich v. Lorain J. Co., 497 U.S. 1, 12 (1990); RESTATEMENT OF TORTS, supra note
9, at § 566.
80 Post, supra note 41, 42, at 638. See also Pierre N. Leval, The No-Money, No-Fault Libel
Suit: Keeping “Sullivan” in Its Proper Place, 101 HARV. L. REV. 1287 (Proposing the award of
judgments without compensation).
81 See generally Jane E. Kirtley, Getting to the Truth: Fake News, Libel Laws, and “Enemies
of the American People,” 43 HUMAN RIGHTS Magazine 4, https://perma.cc/DJK8-W6XX.
82 See, e.g., Dun & Bradstreet, Inc. v. Greenmoss Builders, Inc., 472 U.S. 749 (1985) (Powell,
J., plurality opinion) (stating that a false report of corporate bankruptcy is defamatory); Blake v.
Ann-Marie Giustibelli, P.A., 182 So. 3d 881, 883–84 (Fla. Dist. Ct. App. 2016) (affirming
$350,000 in damages for online defamatory reviews of attorney services). Such interests are best
seen as reputational concerns (rather than social status).
83 Political affiliation is often seen as a core part of an individual’s identity. Yet, false
allegations of political affiliation are not defamatory. See, e.g. Cox v. Hatch, 761 P.2d 556, 562
(Utah 1988), Frinzi v. Hanson, 140 N.W.2d 259, 262 (Wis. 1966). Many would consider
allegations that they are dead to be an affront to their dignity; but these statements are not
defamatory, see, e.g., Cardiff v. Brooklyn Eagle, 75 N.Y.S.2d 222, 224 (Sup. Ct. 1947); Decker
v. Princeton Packet, Inc., 561 A.2d 1122 (N.J. 1989).
84 Hustler Magazine, Inc. v. Falwell, 485 U.S. 46 (1988) (holding that a cartoon of a minister
“engaged in a drunken incestuous rendezvous with his mother in an outhouse” is not defamatory
because it was not understood as describing facts). The false light doctrine allows recovery for
emotional injuries resulting from publications that are not necessarily false. See Braun v. Flynt,
726 F.2d 245, 247 (5th Cir. 1984).
85 See DOBBS, supra note 9, AT § 9:26. See also Smith v. Durden, 2012-NMSC-010, 276
P.3d 943 (N.M. 2012).; Richie v. Paramount Pictures Corp., 544 N.W.2d 21, 28 (Minn. 1996).
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 16/51
Perhaps some sophisticated refinement of the idea of dignity might answer
these challenges. Indeed, status theory does not make dignity irrelevant. But I
think it is fair to question whether dignity is a natural fit here, keeping our eyes
open to other alternatives. After all, the human preoccupation with good name
is older than our legal understanding of dignity and has been the subject of
intense research by sociologists and economists. There is voluminous literature
in these disciplines, utilizing a variety of methodological tools, which closely
study the meaning and importance of good name. The challenges faced by the
standard classification suggest that there is a potential for great profit in learning
from these schools of thought.
2. Status Theory: An Introduction
Given the uncertainty about the values that underlie defamation law, the only
sound premise, shared by all, is that defamation law protects good name interests.
The most promising way forward, then, is to understand what these interests are.
A core insight found in sociology and economics is that good name reflects
two distinct but interrelated human pursuits—reputation and status.86 The
overly simplistic way to introduce them is to think about reputation as an
economic concept, a measure of the desirability of transacting with its subject.
Status is a social concept, a measure of the desirability of affiliating, befriending,
or socializing with its subject. The nation’s leading surgeon has good reputation;
the President has high status. Both values are of deep, sometimes mortal,
importance to individuals, and together they explain a large part of the law of
defamation.
Thus, it should be clear at the outset that status theory is not meant as an
exclusive interpretation of defamation law. Defamation law, I believe, cares
about both social status and reputation, and while these two concepts are often
intertwined, they possess different core meanings. Having wrestled with the
concept of reputation elsewhere,87 I will only make here brief remarks about
reputation and instead focus on exploring the theory of status.88
The concept of status emerges from an old tradition in sociology, dating back
to at least Weber.89 Sociologists define status as “the prestige accorded to
86 The legal literature uses inconsistent terminology and does not offer a holistic framework
that clearly distinguishes between them. See, e.g., Randall P. Bezanson, The Libel Tort Today, 45
WASH. & LEE L. REV. 535, 537 (1988) (calling status “community reputation”).
87 See Yonathan A. Arbel & Murat Mungan, The Case Against Expanding Defamation Law,
71 ALA. L. REV. 453 (2019); Yonathan A. Arbel, The Protection of Reputation in Defamation
Law, work-in-progress (on file with author).
88 There are important and deep ties between status and reputation. See DAVID ROLPH,
REPUTATION, CELEBRITY AND DEFAMATION LAW, 3-6 (2008).
89 Economists have considered the role of status, dating back at least to Adam Smith. ADAM
SMITH, THE THEORY OF MORAL SENTIMENTS 112-13 (D.D. Raphael & A.L. Macfie eds.,
Clarendon Press 1976) (1759). Yet, the role of status is often overshadowed by more tractable and
simple models of human behavior. See Richard H. McAdams, Relative Preferences, 102 YALE L.J.
1, 10-14 (1992).
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 17/51
individuals because of the abstract positions they occupy rather than because of
immediately observable behavior.”90 Status arises organically in social groups—
from the small fraternity to the modern complex society—91 and reflects a social
hierarchy within the group; a pecking order.92 Possessing status is a matter of
great importance to individuals, and satiates what psychologists and sociologists
believe to be a basic human desire.93 Various studies show that “status difference
determines the observable power and prestige within the group.”94 95With status
comes “deference behavior”-—that is, “compl[iance] with that individual’s
wishes, desires, and suggestions . . . unaccompanied by threat or coercion.”96
Status matters in some unexpected ways. Within the aircrew positions of a B-
26 bomber, there is a clear military hierarchy: pilots rank over navigators who
rank over gunners. Interestingly, this military hierarchy, based on operational
considerations, carries over also to purely social settings, where researchers find
that pilots’ opinions are given a dominant role at the expense of gunners.97 In
science, a distinct domain, one finds that high-status scientists will attract many
more citations for similar ideas than their low-status peers.98 While status is often
sought for its own ends, high status also opens doors in market settings by giving
high-status individuals greater access to opportunities and capital. 99 It is not just
90 Roger V. Gould, The Origins of Status Hierarchies: A Formal Theory and Empirical Test,
107 AM. J. SOCIOL. 1143, 1147 (2002). This definition helps distinguish between status and
reputation.
91 Id. at 1143 (2002) (Social differentials have a “near-universality…across a wide range of
scales and contexts, actors are sorted into social positions that carry unequal rewards, obligations,
and expectations.”); see also Bernardo A. Huberman, Christoph H. Loch & Ayse Önçüler, Status
as a Valued Resource, 67 SOC. PSYCH. Q. 103–114 (2004) (finding a strong preference for status
in an experiment across five cultures); Jessica Kiski, Hongling Xie, & Ingrid R. Olson,
Understanding Social Hierarchies: The Neural and Psychological Foundations of Status
Perception, 10 SOC. NEUROSCIENCE 527 (2015) (“A wealth of evidence indicates social hierarchies
are endemic, innate, and most likely, evolved to support survival within a group-living context.”).
92 The commonly used term ‘pecking order’ reflects a real phenomenon which showcases the
ubiquity of status and status games, as chicken direct most of their pecks at lower status fowls. A.
M. Guhl, The Social Order of Chickens, 194 SCI. AM. 42 (1956).
93 See SMITH, supra note 89, at 336 (“The desire of being believed, the desire of persuading,
of leading, and directing other people, seems to be one of the strongest of all our natural desires.”).
94 Joseph Berger, Bernard P. Cohen & Morris Zelditch, Jr., Status Characteristics and Social
Interaction, AM. SOCIO. REV. 241, 243 (1972).
95 Joan C. Williams, Marina Multhaup, & Sky Mihyalo, Why Companies Should Add Class
to Their Diversity Discussions, Harv. Bus. Rev. (Sept 5, 2018) (“[I]n the United States, talking
about class is taboo”).
96 Cameron Anderson, John A. D. Hildreth & Laura Howland, Is the Desire for Status a
Fundamental Human Motive? A Review of the Empirical Literature, 141 PSYCH. BULL. 574, 575
(2015). See also JOEL M. PODOLNY, STATUS SIGNALS 14 (2005).
97 Id. at 241-42.
98 Robert K. Merton, The Matthew Effect in Science, 159 SCI. 56 (1968); see also Michael
Sauder, Freda Lynn & Joel M. Podolny, Status: Insights from Organizational Sociology, 38 ANN.
REV. SOCIO. 267 (2012). On biased citation practices in law, see Richard Delgado, The Imperial
Scholar: Reflections on a Review of Civil Rights Literature, 132 PENN. L. REV. 561 (1984).
99 Sauder, Lynn & Podolny, supra note 98, at 272-73; PODOLNY, supra note 96, at 27-29
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 18/51
that high-status signals merit;100 it is also that having high status makes
everyone’s evaluations of you more favorable.101 For instance, prestigious law
firms can charge higher prices, and one wonders how much of that is attributable
to higher quality.102
Lawyers, and future lawyers in particular, may find special interest and
concern in learning that interview invitations to elite law firms are highly
influenced by status markers. “[E]mployers discriminate on the basis of status
characteristics,” write Rivera and Tilcsik, who find that adding a high-status
marker to a student’s resume—being on the sailing team or listening to classical
music—results in a significantly higher rate of interview invitations than listing
low status markers—being on the track and field team or enjoying country
music.103
As to why individuals, firms, and countries compete for status, sociologists
propose three possibilities.104 First, individuals may pursue status instrumentally
to achieve those material advantages just noted.105 Second, individuals may seek
status as a terminal value simply because having status is pleasurable and losing
it is painful.106 Consistently, psychologists find that “[p]eople’s emotional state,
their short-term moods and long-term happiness, often depend on their ranking
in comparison with others.”107 Lastly, the pursuit of status may also be explained
by evolutionary adaptations to collaboration in group settings, a view supported
by the ubiquity of competition for status within the animal kingdom.108
Let us pause to briefly consider the interrelated concept of reputation. Much
like status, reputation is also a communal concept. It is aggregated information
regarding the quality of a person, service, or product based on past experience.
(“[S]tatus lowers the transaction costs associated with the exchange between buyer and seller.”);
Michael Jensen, Bo Kyung Kim & Heeyon Kim, The Importance of Status in Markets: A Market
Identity Perspective, STATUS IN MGMT. AND ORG. 87, 87 (2010).
100 Huberman, Loch & Önçüler, supra note 91, at 105 (reporting a "strong theoretical basis
as well as empirical support for the fact that status signals competence [and] provides access to
power and resources").
101 Gould, supra note 90, at 1158.
102 Brian Uzzi & Ryon Lancaster, Embeddedness and Price Formation in the Corporate Law
Market, 69 AM. SOCIO. R. 319, 341 (2004) (finding in the market for corporate legal services that
“status has an effect on prices that is independent of the quality of the firm”).
103 Lauren A. Rivera & András Tilcsik, Class Advantage, Commitment Penalty, 81 AM. SOC.
REV. 1115, at 1122 (2016).
104 See RIDGEWAY, supra note 2, at 20-47 (arguing that status serves to coordinate cooperation
within groups).
105 See, e.g., PODOLNY, supra note 96, at 30 (“[H]igher-status actors will be able to offer
goods of a given quality at a lower cost.”).
106 Anderson, Hildreth, & Howland, supra note 96, at 591-93. (reviewing diverse literature
and finding that status pursuits appear to be a fundamental human desire with important effects
on wellbeing); see also Huberman, Loch & Önçüler, supra note 91, at 104.
107 Richard H. McAdams, Relative Preferences, 102 YALE L. J. 1, 31 (1992).
108 Joey T. Cheng & Jessica L. Tracy, Toward a Unified Science of Hierarchy: Dominance
and Prestige are Two Fundamental Pathways to Human Social Rank, THE PSYCHOLOGY OF
SOCIAL STATUS 3 (2014).
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 19/51
109 Barring reputation failures, 110 one could expect higher quality from a product
that has good reviews. 111 The common observation that a firm or brand “has” a
good reputation means that many share a favorable expectation of its quality.112
I think it is easiest to think of reputation as a prediction, although I am also
partial to political scientist Robert Axelrod’s calling reputation “a shadow of the
future.”113 It follows quite naturally why reputation is valuable: it allows its
subject to capitalize on it.114
In reality, there is often overlap between the distinct concepts of status and
reputation, so it is understandable why the literature conflated them.115 It is
especially easy to mistake them in a society that has an ethos of meritocratic
allocation of status, where supposedly those admired are the most competent.116
Still, even twins are different people. While reputation measures quality, status
measures relative social standing; while reputation is mostly an instrumental
value, status is ta least partly a terminal one; and while one builds reputation by
accumulating positive reviews of past experiences, status is earned through the
accumulation of ‘deference behavior.’117 Quality is key to reputation, but
secondary to status. The late sociologist Roger Gould went as far as showing that
109 See RICHARD A. POSNER, THE ECONOMICS OF JUSTICE 272 (1981) (“A person’s
reputation is other people’s valuation of him as a trading, social, marital, or other kind of partner.
It is an asset of potentially great value which can be damaged both by false and by true
defamation.”); Yonathan Arbel, Reputation Failure: The Limits of Market Discipline in Consumer
Markets, 54 WAKE FOREST L. REV. 1239, 1255 (2019) (“reputational information is like a poll” as
it “helps consumers predict their own experiences based on the distribution and valence of
experiences of past consumers”). See also Roy Shapira, Reputation Through Litigation: How the
Legal System Shapes Behavior by Producing Information, 91 WASH. L. REV. 1193, 1203–04
(2016).
110 See generally Arbel, supra note 109 (exploring factors leading to reputation failures in
markets).
111 See Simon Board & Moritz Meyer-Ter-Vehn, Reputation for Quality, 81
ECONOMETRICA 2381, 2381 (2013) (defining reputation “as the market’s belief about …
quality”); Benjamin Klein & Keith B. Leffler, The Role of Market Forces in Assuring Contractual
Performance, 89 J. POL. ECON. 615, 616 (1981).
112 The statement that a brand enjoys a good reputation is intelligible, but it would be highly
confusing from a perspective of honor, property, or dignity.
113 ROBERT AXELROD, THE EVOLUTION OF COOPERATION 126 (1984).
114 See Benjamin Klein & Keith B. Leffler, The Role of Market Forces in Assuring
Contractual Performance, 89 J. POL. ECON. 615, 616 (1981).
115 See Olav Sorenson, Status and Reputation: Synonyms or Separate Concepts?, 12 STRATEG.
ORG. 62, 63 (2014). Economists have frequently conflated the two meanings by redefining status
as a measure of quality. See, e.g., Michael Jensen, Bo Kyung Kim & Heeyon Kim, The Importance
of Status in Markets: A Market Identity Perspective, STATUS MANAGEMENT ORG. 87–117 (2010).
The distinction developed here maps into a distinction in trademark law, which considers brands
as either informational signals of quality or markers of prestige. See Shahar J. Dillbary, Famous
Trademarks and the Rational Basis for Protecting Irrational Beliefs, 605 GEO. MASON L. REV.
605, 610-15 (2011).
116 See generally RIDGEWAY, supra note 2, at 6-7 (offering a merit based view of status
allocation). Podolny posits that status is also a predictor of quality in market relations used to
complement gaps in reputational information. PODOLNY, supra note 96, at 18.
117 See, e.g., Sauder, Lynn & Podolny, supra note 98, at 268.
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 20/51
“it is possible for a stable system of ranked social positions to emerge
endogenously in the absence of underlying variation in individual attributes.”118
Another difference is that it would be relatively easy for an outsider to evaluate
the reputation of various agents based on the demand for their services. However,
an outsider will find it difficult to track and quantify the allocation of status,
which manifests in non-market behavior.119 And even the internal experience is
different: reputation is about what we expect to get and status is what we should
get.120
A final point about status is that it is a relative property. We see that in the
way we talk about high status, as opposed to good reputation. There is no high
status without low status: leaders imply followers; cool kids, nerds; upper-class,
lower-class; patricians, plebeians; Brahmins, Sudras; gold medals, bronze medals;
and Ivy League schools, exposed-brick schools. We always measure status relative
to others.121 But reputation is different. Because reputation is a prediction of
future quality, it is not impossible for all the restaurants in a given area to enjoy
a high reputation for the quality of their food.122 This is not to say that reputation
is not competitive—every firm would like to be the only one with a high
reputation—but only to illustrate the different dynamics that sociologists and
economists attach to status and reputation.123
3. A Status Interpretation of Defamation Law
This brief introduction to status theory cannot do justice to such a mature
theory, but if we are willing to accept some sacrifice of nuance, we can test the
pudding by its eating. What follows is an attempt to read defamation law
through the lens of status theory. The analysis has no ambition of explaining all
of defamation law. Nor need it. Status theory could be valuable if it can
persuasively explain a fair portion of the law in a coherent manner—or, at least,
if it can do so better than our existing accounts of dignity, honor, or property.
The nexus between status and defamation law is first observed in the rhetoric
surrounding the doctrine. The Supreme Court endorsed a description of
defamation law as protecting individuals from loss of “standing in the
118 Gould, supra note 90, at 1149.
119 The difficulty of tracking status may explain defamation law’s liberal allowance of recovery
of presumed damages. See Dun & Bradstreet, Inc. v. Greenmoss Builders, Inc., 472 U.S. 749, 760
(1985) (quoting PROSSER, supra note 10, at 765) (“[P]roof of actual damage will be impossible in
a great many cases where, from the character of the defamatory words and the circumstances of
publication, it is all but certain that serious harm has resulted in fact.”); see also SACK ON
DEFAMATION, supra note 8, at § 2:4.2. Insiders develop a quick and intuitive sense of internal
social hierarchies, which they share among themselves with regularity.
120 The ‘should’ here is sociologically, not morally, normative. On the norms guiding the
attainment of status. See infra 0.1.
121 See, e.g., Jensen, Kim, & Kim, supra note 115, at 91 (“[S]tatus is best defined as a position
in a social system.”).
122 Cf. ROY SHAPIRA, LAW AND REPUTATION 120 (2020).
123 See infra Appendix.
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 21/51
community,”124 a telling reference to the social aspect of the tort. Consistently,
the common definition of defamatory expressions refers to statements that
expose individuals to “hatred, contempt, ridicule, or obloquy, or which cause[]...
any person to be shunned or avoided.”125 This is status-laden language, clearly
geared towards the social effects of statements.126 Some commentators have
likewise explained the need for monetary compensation in defamation cases in
the need to rehabilitate a “relational interest” that defamation jeopardizes.127
Dignity theory was criticized for its lax doctrinal fit. Let us measure how status
theory fares on this measure. Consider, first, the publication requirement. For
dignity theories, this requirement raises difficulty, because even private
statements can be extremely demeaning. In contrast, such a requirement is
inescapable if the value protected is social status. Unlike dignity, social status
cannot be lost in private. One can only lose status by being viewed negatively in
the eyes of others.128 The mechanisms of status loss render publication a non-
negotiable prerequisite.
We can also derive from first principles the community judgment
requirement. The measure of whether any given statement is defamatory is its
reception in the community. As Justice Holmes put it, statements are only
defamatory if they would “hurt the plaintiff in the estimation of an important
and respectable body of the community.”129 This requirement is puzzling if one
views dignitary harm as the crux of defamation. Why limit recovery to harm in
the eyes of the community and not, say, in the eyes of a loved one?130 And why
should it matter if those people are respectable or not? But from a status
perspective, these requirements are natural. Status only emerges within social
communities, and the existence of harm requires a change in the views of their
124 Gertz v. Robert Welch, Inc., 418 U.S. 323, 350 (1974).
125 ROBERT H. PHELPS & E. DOUGLAS HAMILTON, LIBEL: RIGHTS, RISKS, RESPONSIBILITIES
6 (1966); see also Phelan v. May Dep't Stores Co., 819 N.E.2d 550, 553 (Mass.
2004) (quoting Stone v. Essex Cty Newspapers, Inc., 330 N.E.2d 161, 165 (Mass. 1975))
(defining defamation as a statement that “would tend to hold the plaintiff up to scorn, hatred,
ridicule or contempt”).
126 W. PAGE KEETON ET AL., PROSSER AND KEETON ON THE LAW OF TORTS § 111, at 771
(5th ed. 1984) (citing “personal disgrace”) [hereinafter PROSSER AND KEETON]; Kimmerle v. New
York Evening J., 186 N.E. 217 (N.Y. 1933) (“[I]nduce an evil opinion of one in the minds of
right-thinking persons, and to deprive one of their confidence and friendly intercourse in
society.”).
127 RODNEY A. SMOLLA, 1 THE LAW OF DEFAMATION 18 (2d ed. 2020) (citing LEON GREEN,
CASES ON INJURIES TO RELATIONS 193-276 (1940)).
128 RIDGEWAY, supra note 2, at 65. Likewise, emotional pain and suffering are considered
‘parasitic’ on other harms and cannot exist by themselves.
129 Peck v. Tribune, 214 U.S. 185 (1909); see also RESTATEMENT OF TORTS, supra note 9, at
§559 cmt. e. (“[T]he communication would tend to prejudice [the victim] in the eyes of a
substantial and respectable minority.”); Mycroft v. Sleight, 90 L. J. K. B. 883 (1921) (“[I]n the
minds of ordinary, just and reasonable citizens.”).
130 Lidsky, supra note 33, at 19 (“[I]f the single individual who finds the statement defamatory
is the plaintiff's spouse or boss, the plaintiff will receive no recovery despite the very real and
substantial nature of his injury.”).
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 22/51
members. And the more respect and status audience members possess, the more
impactful the statement can be on one’s status.131
We can also revisit the libel-proof doctrine. Being libel proof means that one’s
standing is so low that no harm is visited by a defamatory allegation.132 From a
dignity perspective, such a doctrine is inexplicable, as all individuals have equal
dignity.133 But from a status perspective, it would make sense that those on the
lowest social rung are not losing much status from defamation.134 While one may
not be “dignity proof,” being status proof is entirely plausible (albeit tragic).
Finally, the most important and nuanced aspect of defamation doctrine is the
inveterate requirement that the statement is false, a requirement that long
predates modern concepts of free speech.135 From a dignity- or honor-based
perspective, this limitation is clearly puzzling.136 If anything, the humiliation a
person suffers from derogatory remarks is greater when those remarks prove true.
From a status perspective, however, the falsity requirement is a natural
corollary.137 The meaning of this requirement will be clarified once status games
are introduced; it is enough, for now, to note that status games’ integrity requires
some arbitration of which claims are false and which are not.138 After all, if a
person claims a status privilege on the basis of her honesty or piety, then it is
essential that others could truthfully expose her dishonesty or impiety.139
Majority of commentators today converge on a dignity of defamation law.
Beyond the Supreme Court itself, 140 a prominent exemplar is legal philosopher
Jeremey Waldron who espouses a full-throated dignity-based view of defamation
131 See, e.g., PODOLNY, supra note 96, at 15 (“[R]eceiving deference from a high-status actor
generally has a greater impact on one’s own status than receiving deference from a low-status
individual.”). The torts of intentional infliction of harm and breach of privacy are designed to
address cases that do not fit within this category. See David A. Logan, Tort Law and the Central
Meaning of the First Amendment, 51 U. PITT. L. REV. 493, 524 (1990).
132 Cardillo v. Doubleday & Co., 518 F.2d 638, 639 (2d Cir. 1975) (holding that a
“habitual criminal” was libel proof). SACK ON DEFAMATION, supra note 8, at § 2:4.18.
133 See supra notes 71-81 and accompanying text.
134 This is consistent with the mitigation of damages for individuals with low-status, as they
presumably suffer a lesser status harm. The reverse is true for individuals with high standing.
SMOLLA, supra note 127, at § 13.17(“Evidence that the plaintiff already has a bad reputation is
admissible in mitigation of damages.”); Mike Steenson, Presumed Damages in Defamation Law,
40 WM. MITCHELL L. REV. 1492, 1504 (2014).
135 Originally, as the law of libel was concerned with preserving the peace, it did not concern
itself with the veracity of allegations. See generally Van Vechten Veeder, The History and Theory
of the Law of Defamation, 8 COLU. L. REV. 546 (1903).
136 See Post, supra note 41, at 705-06.
137 For an early statement of the truth defense, see 3 WILLIAM BLACKSTONE, COMMENTARIES
*433-34.
138 In the context of bigoted defamation, our goal is to disrupt the underlying status game,
which is why the law does not regulate the veracity or mendacity of statements. See infra Part II.
139 See infra Part III.1.
140 Gertz v. Robert Welch, Inc., 418 U.S. 323, 342(1974) (quoting Rosenblatt v. Baer, 383
U.S. 75, 92 (1966) (Stewart, J., concurring)).
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 23/51
law.141 Another is defamation law expert Professor Lyrissa Lidsky, who argued
that “[b]y protecting reputation, defamation law safeguards the dignity of
citizens.”142 Beyond these leading voices, examples abound;143 in fact, despite the
voluminous literature on the topic, it is hard to locate different views who
challenge the dominance of the dignity or property view of defamation.
Despite its dominance, the dignity view faces difficult challenges. The
discussion above was sufficient, I hope, to illustrate how many difficult doctrinal
requirements that confound dignity, property, and honor interpretations appear
logical, even necessary, when viewed from the lens of status theory. The
publication requirement, the use of monetary compensation, the categories of
per-se libel, the indifference to naked dignitary harms, are but some of the
features that are easily interpretable from within the status perspective.
I would conjecture that the goodness of fit between theory and doctrine
suggests is not coincidental. Courts and commentators may have employed a
proto-theory of status for a long time, even if they lacked the theoretical
vocabulary to articulate it clearly. In fact, many of Post’s original ideas also fit
comfortably within this framework, suggesting theoretical consilience. Thus,
while status theory claims greater explanatory power, it does not claim much in
the way of novelty. At best, it is a refinement of older ideas and concepts that are
already present in the law. If it earns its interpretative keep, it is by expounding
these ideas clearly and parsimoniously.
DEFAMATION LAW AND STATUS GAMES
Interpreting defamation law from a status perspective offers a way to
coherently understand doctrine. Interpretative projects are important in
themselves, especially in areas that are cloudy and mystifying. Status theory,
however, is a more capacious framework; it offers prescriptive guidance to courts,
legislators, and reformers. The goal of this Part is to explicate the normative side
of the theory—employing a socio-economic-legal analysis.
In the traditional legal account, defamation law only protects good name
interests. The argument developed here is that the traditional account sells
defamation law short. By intervening in status competitions, that is, by
protecting some claims to status and denying others, defamation law inexorably
affects the choices individuals make in the first instance about how to acquire
141 See Waldron, supra note 37. See also infra notes 276-282 and accompanying text.
142 Lyrissa Barnett Lidsky, Silencing John Doe: Defamation & Discourse in Cyberspace, 49
DUKE L.J. 855, 885 (2000)
143 See generally Leslie Meltzer Henry, The Jurisprudence of Dignity, 160 U. PA. L. REV. 169,
217-18 (2011) (discussing courts defamation dignity jurisprudence); Neomi Rao, Three Concepts
of Dignity in Constitutional Law, 86 Notre Dame L. Rev. 183, 253-58 (2011) (considering, from
a comparative perspective, dignity in the context of defamation law).
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 24/51
status.144 This Part introduces the concept of status games –the status
competitions through which individuals gain and lose status. On this basis, it
explains how determinations in particular cases affect the choice of status games,
and then develops the normative argument that courts should openly recognize
their role and consider status games explicitly when they adjudicate cases. In
other words, courts play a regulatory role in status games and, while there are
important institutional constraints, a judicial reckoning of this role is vital.
1. Status Games
Defamatory remarks jeopardize social status. Obviously, not all statements
about me jeopardize my status. To falsely say that I cannot grow yams will have
nil effect on my social status, but it would be devastating if I lived in a different
society.145 There are specific social rules at play that govern status, rules that
define how status is lost, but also how it is gained. It is those rules that give
structure to the status competitions that are played within status games.
Status games are emergent social systems of recognized rules of status
acquisition, maintenance, and deprivation, alongside the entailments of status.146
Status games are a near-universal property of any social group, forming
spontaneously and often informally and subconsciously.147 The variety of status
games is dazzling, and they span the immediately recognizable (the consumption
of expensive items, titles, and physical appearance) to the nuanced and seemingly
“natural” (accent, manners, and even body language). A 19th-century lady is
instructed to “never read in company,” and when encountering a “gentleman at
the foot of flight of stairs,” she shall “Stop, bow, and motion him to precede.”148
These overtly sexist rules were codified in etiquette books, which themselves fell
out of fashion, only to be replaced by modern, less legible, status games. Probably
the most radical account of status games is the one proposed by sociologist Ervin
Goffman, who famously claimed that the entire presentation of the self is a form
of soliciting social impressions.149
Because many native status games are invisible from within, a useful tell is the
feeling one experiences when imagining transgressions. A faux-pas evokes such a
visceral response that it even redounds to third parties—in what is known as
144 The Technical Appendix develops a critique of the protection account on its own terms.
It explains why, from a social perspective, defamatory remarks are not necessarily harmful, thus
undermining the protection imperative. See infra Technical Appendix.
145 William R. Bascom, Ponapean Prestige Economy, 4 SW. J. ANTH. 211, 213 (1948).
146 See e.g., Anderson et al., supra note 96, at 116.
147 See Jessica Koski et al., Understanding Social Hierarchies: The Neural and Psychological
Foundations of Status Perception, 10 SOC. NEUROSCIENCE 527 (2015) (“[S]ocial hierarchies are
highly pervasive across human cultures … and they appear to emerge naturally in social groups …
this group organization is not strictly a product of human cognition, as almost every group-living
species demonstrates a natural tendency to organize into a social hierarchy.”).
148 FLORENCE HARTLEY, THE LADIES' BOOK OF ETIQUETTE, AND MANUAL OF POLITENESS,
284-85 (1860).
149 ERVING GOFFMAN, THE PRESENTATION OF SELF IN EVERYDAY LIFE (1959).
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 25/51
‘cringe’—just imagine accidentally calling the President ‘dude’, burping in a
company dinner, or the ubiquitous nightmare of coming to school with no
pants.150
Reports of social ‘fouls’ diffuse in the community through informal
mechanisms such as gossip.151 This is obviously relevant to defamation law,
which is focused on sanctioning false reports. It is worth noting that individuals
seem to take pleasure in diffusing and consuming such reports: the promise of
‘juicy’ gossip sustains much of the tabloid industry.
A key observation is that status games do more than allocate status, they also
affect society itself. Status games consume resources, time, and psychic energy,
while the form of the status game impacts individuals, for better or worse.152
These effects are not always visible because, as economist Roger Congleton
recognized, they accrue to non-participants.153 For instance, politeness and
civility is a form of a status game that directly improves the well-being of others
by rewarding only those who act amicably, patiently, and pro-socially.
Philanthropy is another example of a status game that creates a real difference in
the lives of others—sometimes literally saving them.154 It has become fairly
common to complain today of virtue-signaling: the conspicuous display of
prosocial attitudes motivated by selfish concerns of status and public image. But
perhaps the effects of such signaling is not all bad. In an intriguing account,
economist Robert Frank argued that status games sometimes help economic
inequality. He argued that workers do not only care about their own absolute
wages but also about their relative earnings. When employers use differential
wages, they sow the seeds of discontent. To preserve morale, managers must
maintain a certain degree of pay equality or else risk attrition.155
Alongside their more salutary implications, one must admit, there is also a
dark side to status games.156 In The Darwin Economy, Frank argues that status
150 Modern television shows engage in intensive exploration of this feeling. See generally Julia
Havas and Maria Sulimma, Through the Gaps of My Fingers: Genre, Femininity, and Cringe
Aesthetics in Dramedy Television, 21 TELEVISION & NEW MEDIA 75–94 (2021) (theorizing the
role of cringe in contemporary television shows).
151 See Terence D. Dores Cruz et al., The Bright and Dark Side of Gossip for Cooperation
in Groups, 10 FRONTIERS IN PSYCH. 1374 (2019) (noting the function of gossip in enforcing
group norms and its universality in human societies).
152 See SHAPIRA, supra note 122, at 137-38 (discussing the signal “broadcast efficiency” based
on its social effects).
153 Congleton, supra note 3, at 176.
154 See Amihai Glazer & Kai A. Konrad, A Signaling Explanation for Charity, 86 AM. ECON.
REV. 1019 (1996).
155 See Robert Frank, Are Workers Paid Their Marginal Products?, 74 AM. ECON. REV. 549
(1984).
156 See, e.g., Douadia Bougherara et al., Do Positional Preferences Cause Welfare Gains?, 39
ECON. BULL. 1228, 1229 (2019) (“In an economy with private consumption goods, positional
preferences lead to a welfare loss.”); Congleton, supra note 3, at 176 (“A substantial portion of the
investment in positional goods may be regarded as a dead-weight loss.”). The term “positional
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 26/51
games often result in socially destructive outcomes. Frank’s ideas draw on
Darwinian competitions for female attention in nature, where elks grow
unwieldy antlers, elephant seals grow to unsustainable sizes, and peacocks boast
heavy and flashy tails—all features that make survival harder.157 Similarly, Frank
argues, human competition for status can result in races that consume resources
but produce no improvement. We might compare this to an overcrowded
concert. If one person stands on her tiptoes, she can see the show better. But this
can lead to a cycle where everybody else also stands on their toes, resulting in
everyone standing uncomfortably and no one seeing any better for it. 158
A powerful example of pernicious status races comes from Nobel Laureate
economist Thomas Schelling. He noted the oddity that, given a choice, hockey
players would choose to skate without a helmet, but if asked to vote on a league
rule, the overwhelming majority of players would require helmets. The reason
for this dissonance stems from the underlying race. Any player not wearing a
helmet sees the field slightly better and is thus more likely to reap status and
financial rewards. But if all players skate without a helmet, this advantage
vanishes, and the original competitive ranking is maintained while leaving all
players more vulnerable to serious injuries.159
Keeping up with the Joneses is a familiar status race among neighbors, where
entire neighborhoods are drawn into a one-upmanship game of maintaining
large lawns, driving lavish cars, and donning expensive brands. The game is not
played because of the inherent utility of these actions, it is played in order to save
face.160 As Veblen noted, individuals often engage in such “conspicuous
consumption” to impress others and win their envy, although they would rarely
admit to such motives.161 While status races may not always be conscious, their
existence in our daily life is illuminated by the common and seemingly
innocuous pursuit of “decent” clothes, a “good” car, or a “nice” house, which,
“upon analysis, turn out to be (at least partly) relative to what others have.”162
arms race” is due to Robert H. Frank, Should Public Policy Respond to Positional Externalities?,
92 J. PUB. ECON. 1777, 1778 (2008).
157 ROBERT H. FRANK, THE DARWIN ECONOMY: LIBERTY, COMPETITION, AND THE
COMMON GOOD 8-9 (2011).
158 FRED HIRSCH, SOCIAL LIMITS TO GROWTH 5 (1995); see also Erzo Luttmer, Neighbors
as Negatives: Relative Earnings and Well-Being, 120 Q.J. ECON. 963 (2005) (arguing that
individuals feel worse off when their neighbors do better).
159 FRANK, supra note 157.
160 See, e.g., Frank, supra note 156, at 1778 (suggesting large houses are a source of positional
utility); FRANK, supra note 157, at 68-69.
161 THORSTEIN VEBLEN, THE THEORY OF THE LEISURE CLASS, 33-48 (1925) (noting that
status pursuits may not be entirely conscious “so much as it is a desire to live up to a conventional
standard of decency in the amount and grade of goods consumed."). See also ROGER S. MASON,
CONSPICUOUS CONSUMPTION: A STUDY OF EXCEPTIONAL CONSUMER BEHAVIOUR 42 (1981)
(stating that a conspicuous consumer, "anxious to display wealth and gain in prestige, will rarely if
ever explicitly admit to any such intentions").
162 Richard H. McAdams, Relative Preferences, 102 YALE L.J. 1, 43 (1992).
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 27/51
In such situations, one can easily identify a “positional treadmill,”163 where
players end up roughly where they started, only poorer. If resources are spent,
used up, or misused along the way—say, land that could be used for habitation
is converted into a fancy lawn—society itself suffers. Thus, we can see how status,
the product of status games, is a matter of great social interest. Some games
played can enrich society while other can greatly impoverish it.164
2. Defamation Law as a Regulator of Status Games
People invest their best efforts, most precious resources, and a great degree of
mental energy in status competitions. “To exalt my station,” wrote Alexander
Hamilton to one of his friends, “I would willingly risk my life, though not my
character.”165 What if some of this energy could be directed toward positive
games, or at least redirected away from antisocial ones?
This basic insight occurred to media magnate Ted Turner.166 He looked at
the Forbes 400 list of the top wealthiest Americans and realized that it might
actually be keeping people from giving. If one is located in the 10th place, giving
away money could cost them their spot in the top-10—and the same goes for
those in the 50th place, 100th, up to the person in the 400th place, who would be
especially anxious to give away any money. He also realized that status games are
somewhat malleable. With his influence, the journal Slate created an exclusive
list of top-sixty donors. The effects, at least according to Slate, were large and
noticeable. “Whether by coincidence or not, philanthropy has blossomed since
Slate’s list was created.”167
To understand the place of defamation law in this ecosystem, consider the
fundamental problem of status. The ROI of status games, the return on all of
this investment, is measured in the currency of status: the name we develop for
ourselves. The trouble is that unlike hard currency, which can be safeguarded in
banks and vaults, good name is easily deprived. The fruits of the investment are
a vicious rumor away from being destroyed. In other words, status is fragile.
Indeed, investment mogul Warren Buffett observed that “It takes 20 years to
163 Cass R. Sunstein & Robert H. Frank, Cost-Benefit Analysis and Relative Position, SSRN
237665 (Aug. 14, 2000), https://papers.ssrn.com/sol3/papers.cfm?abstract_id=237665.
164 HIRSCH, supra note 158, at 10-11 (“Positional goods … become an increasing brake on
the expansion and extension of economic welfare.”); see also id. at 37-38. Even philanthropic
activity may be excessive. See Glazer & Konrad, supra note 154.
165 Alexander Hamilton to Edward Stevens, November 11, 1769, quoted in Ron Chernow,
ALEXANDER HAMILTON 31 (2004).
166 Nicholas Kristof, How Giving Became Cool, N.Y. TIMES (Dec. 27, 2012),
https://perma.cc/C4LT-4EFL.
167 Sebastian Mallaby, The Slate 60 Turns 10, SLATE (Feb. 20, 2006, 8:19 AM),
https://perma.cc/D2RX-AEPJ.
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 28/51
build a reputation and five minutes to ruin it.”168 With this comes the idea of
“status anxiety,” the recurring fear that we will become ‘nobodies’ overnight.169
Most people are attracted to safer investments, ones that promise predictable
and durable returns. If some forms of status are more fragile than others, this
creates a problem for the underlying status game—but also an opportunity for
the law to play a role.
In the Hobbesian society, attaining the status of a great inventor is for
suckers.170 Take a budding researcher, who daydreams about making a name for
himself by making groundbreaking discoveries in the natural world. He does not
harbor this aspiration because it is good for his financial welfare. In fact, his
relatives keep telling him that he would do better plowing the fields than the
skies. But our scholar is determined. However, despite all of his determination
and willingness to forgo more material pursuits, he is shaken by the realization
that his accomplishments will always be tentative. He will be forever exposed to
gossipmongers who can sully his good name by spreading rumors on how he
stole his best ideas from others. The scholar’s only recourse to defending his hard-
earned status might be violence or duels—and as 20-year-old French genius
Galois discovered, being a prodigal genius is quite distinct from being a good
marksman.171 Foreseeing all of this difficulty, the scholar decides to abandon the
innovation race altogether. He sets his aims at attaining status that cannot be so
easily deprived, through conspicuous consumption, the hoarding of property,
and ostentatious displays of power.
This thought experiment reveals the potential of defamation law. Vicious
rumors and false allegations, left unchecked, deprive their subjects of the fruit of
their investment. Defamation law secures the ROI in status games, by both
sanctioning the gossipmonger and compensating the subject. We know that
when judges protect property, they guard the incentive to maintain it; when they
protect contractual claims, they invite reliance and investment; and when they
deny enforcement of illegal contracts, they discourage illicit transactions.172 By
the same token, when judges extend defamation protection, they secure status
168 Snyder, Benjamin. “7 Insights from Legendary Investor Warren Buffett.” CNBC, May 1,
2017. https://www.cnbc.com/2017/05/01/7-insights-from-legendary-investor-warren-
buffett.html.
169 See generally ALAIN DE BOTTON, STATUS ANXIETY (2008) (exploring the concept).
170 It is perhaps no coincidence that Hobbes was skeptical about ‘glory’. See generally Strong,
supra note 103.
171 See PROSSER AND KEETON, supra note 126, at § 111 p. 772 (arguing that defamation law
came to replace duels and blood feuds). See also JOHN LYDE WILSON, THE CODE OF HONOR OR
RULES FOR THE GOVERNMENT OF PRINCIPALS AND SECONDS IN DUELLING 6 (1858) (“[I]n cases
where the laws … give no redress … it is needless and a waste of time to denounce the [dueling.]”);
see also STARKIE, supra note 36, at 6 (recounting a case where the plaintiff said that if he could not
expect recovery in court “he would have cut the defendant’s throat”).
172 See generally Harold Demsetz, Toward a Theory of Property Rights, 57 AM. ECON. REV.
347 (1967).
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 29/51
and thus promote participation in the underlying status game.173 With
defamation law in place, the Scholar will be less likely to be defamed and, even
if he will be, could recover some compensation from the offender. Conversely, if
one day the law decides that allegations of plagiarism, real or false, are beyond its
purview, then this can have a significant impact on certain intellectual status
games. What defamation law protects, it encourages; what it shuns, it
discourages.
To be sure, defamation law affects status games only on the margin. The law
is hardly the only thing that binds them together. There were status games
between Cain and Abel long before any legal system evolved.174 Still, along the
relevant margins, defamation law affects the choice of status games that
individuals play—and the intensity with which they play them.175 Not a small
feat.
Yet the standard account is seemingly oblivious to these regulatory effects. It
talks in terms of protection, not regulation. But it should be clear that by
protecting, the law also has broader behavioral effects. And it is not that modern
defamation law does not consider social effects. After all, the primary thrust of
New York Times v. Sullivan is the fear that stringent defamation standards can
chill the incentive to participate in public debates.176 What courts fail to
recognize, however, is that stringent defamation standards also secure rewards in
the underlying status game.177 It is like thinking about liability for accidents only
in terms of compensating victims and the good drivers it may deter from driving,
not considering at all how safer roads benefit drivers and pedestrians alike.
The truth is that when courts decide cases, they cannot help but play a key
role in the regulation of status games. Any time the court announces, for
example, that calling a person a “slacker” for avoiding the draft is defamatory, it
promotes status based on service to the nation.178 When the courts decide to treat
allegations of female ‘promiscuity’ in a special manner, as per-se defamatory, they
promote a status game built around chastity.179 And when courts deny the aegis
of defamation law in the case of bigoted defamation, they weaken the bonds that
hold bigoted status games together.180
173 Cf. Waldron, supra note 37, at 1605 (“A democratic republic might equally be concerned
with upholding and vindicating important aspects of legal and social status… and with protecting
that status (as a matter of public order) from being undermined by various forms of obloquy.”).
174 Genesis 4.
175 Other laws also interact with status games. See infra Part III.3.
176 New York Times Co. v. Sullivan, 376 U.S. 254, 300 (1964).
177 For a critique of the theory of the chilling of speech, see Daniel Jacob Hemel & Ariel
Porat, Free Speech and Cheap Talk, 11 J. LEGAL ANALYSIS 46 (2019); see also Arbel & Mungan,
supra note 87.
178 Choctaw Coal & Mining Co. v. Lillich, 86 So. 383, 384-85 (1920) (holding that ‘slacker’
is per se libelous as it is “unquestionably a term of the severest reproach”).
179 See infra Part III.3.
180 See infra Part III.1.
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 30/51
In sum, status games involve intense investment, but their rewards can be
fragile. Defamation law secures these rewards when it sanctions false allegations,
thus contributing to the integrity of status games. Any judicial decision that
either expands or contracts the scope of defamation law inevitably affects the
status games that individuals play, in terms of both the choice of which games to
play and the intensity with which individuals play them. This understanding
elucidates what judges are really doing when they decide cases; it also opens the
way to a more normative view on what judges should be doing.
3. The Judicial Role in Defamation cases
The diagnosis invites a prescription. For reasons of judicial integrity and
accountability, I believe courts must openly acknowledge that they regulate
status games when they decide defamation cases. It is what judges—and to a
lesser extent legislators—do, and recognition of this fact is necessary. Professor
Lidsky once noted that a troubling aspect of defamation doctrine today is “not
that value choices are made but rather that they are cloaked in the deceptively
neutral language of determining defamatoriness.”181 It is now time to decloak the
courts’ role in regulating status games by turning away from the obfuscatory
protection-from-harm view. Perhaps judges are not always conscious of the full
effects of their decisions, but this only means that they cannot hope to perform
their role well. Courts should openly admit the role they are playing, and have
been playing for generations, in regulating status games.182
Acknowledgment is the first step; the more ambitious one is the license to
shape defamation law’s scope on the basis of the underlying status game. In short,
courts should extend status protection when society approves of the status game
in which the status was attained, and deny protection when society views the
status game with opprobrium. If society wants to protect innovation in the
Hobbesian society, it may need to create some form of defamation law; and if it
wants to disrupt status games based on bigoted hierarchies, it should let anyone
spread rumors freely. Protection-from-harm, the center of the standard account,
is an instrument in status theory. Harm to status should only be remedied when
the claim to status is legitimate. Harm is the tail, not the dog.
This means a reconstruction of the judicial role, from protecting status to
regulating status. Robert Post had the foresight to recognize this point. He
explained that “the meaning and significance of reputation will depend upon the
kinds of social relationships that defamation law is designed to uphold.”183 But
181 Lidsky, supra note 33, at 9.
182 See, e.g., Bovard v. Am. Horse Enters., 247 Cal. Rptr. 340 (Cal. Ct. App. 1988).
183 Post, supra note 41, at 693.
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 31/51
his focus on fostering social relationships was left mostly unheeded.184 Perhaps
status theory can help amplify this message and guide decision-making.
4. Which Status Games are Worth Protecting?
The difficult question is not whether courts and legislators should regulate
status games, but how.
They say that the first step is admitting that you have a problem. It will be an
important first step to explicitly recognize that the hidden underbelly of
defamation law determinations is the craft of regulating status games. Such an
open admission could lead, over time, to the evolution of decision and
accountability norms. To be sure, regulation of status games is an involved
decision, but this point calls for moderation and caution as much as it does for
transparency and accountability. Thus, on its own, admitting that defamation
law is in the business of regulating status games would be the first step towards
developing a robust caselaw.
There are also some good reasons to think that a principled approach to the
regulation of status games is within reach. When a claim involves loss of status,
the next question should ascertain the origin of that status—what is the
underlying status game that gave rise to the plaintiff’s status in the first place?
Such an investigation will reveal some status games that are virtuous, many that
are of ambiguous value, and some that are clearly noxious.185 Then depending
on the nature of the status game, judges can craft the scope of protection that
best fulfills society’s goals. 186
Virtuous Status Games. Status games that are valuable can be discerned by
their positive spillovers. Scholastic status races fit well in this category, as they
lead scholars to exert themselves to become the first to discover a vaccine, observe
an important physical phenomenon, or develop a new theory. 187 A different
example comes from the Bill and Melinda Gates Foundation, a non-profit
184 Defamation law scholar David Anderson, in a road to Damascus moment, recognized the
regulatory role of defamation law. “Compensation is not the only legitimate purpose of defamation
law. Robert Post is right . . . . [the Law also aims] to enforce society’s civility norms.” David A.
Anderson, Rethinking Defamation Law, 48 ARIZ. L. REV. 1047, 1049 (1991).
185 See Congleton, supra note 3, at 182-183 (arguing that status games involving positive
externalities may need to be subsidized, while negative externalities should be met with a Pigouvian
tax); see also Huberman, Loch & Önçüler, supra note 91, at 103 (“Intrinsic status seeking by
individuals has important implications for social and economic systems because it can provide a
powerful motivation to perform; it also can lead to unproductive competitions … such as in the
overconsumption of positional goods.”).
186 See Congleton, supra note 3, at 181 (“If status-seeking activities affect only the welfare of
others in the status game, it is relatively straightforward to demonstrate that too many resources
will be invested in the quest for status.”).
187 Some ancient texts recognize the motivating force of envy on scholarship, holding that
“jealousy among teachers increases wisdom.” Talmud Bava Batra, 21a
https://www.sefaria.org/Bava_Batra.21a?lang=bi. For a skeptical account, see Brian L. Frye,
Plagiarize This Paper, 60 IDEA 294 (2020) (“[A]cademic plagiarism norms are primarily an
inefficient and illegitimate form of extra-legal academic rent-seeking that should be ignored.”).
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 32/51
devoted to fighting poverty, improving healthcare, and expanding access to
information technology.188 Here, the pursuit of legacy—sometimes maligned as
selfish and narcissistic—led the Gates’ to donate $36 billion dollars to help
improve the world. Status games around magnanimity and generosity of spirit
are key drivers of philanthropy everywhere. Similarly, the quest for fame
harnesses the creative energies of many individuals, directing them to use those
energies to create art that will make everyone “remember my name.”189
Ambivalent status games. Other status games are not so clearly virtuous, yet
they feature some positive elements. Through surveys, sociologists have mapped
the way individuals perceive the distribution of status among occupations.190
The distribution sometimes appears justified, other times arbitrary and even
unjust. A typical survey found that biologists (ranked at 6.9) outperform bankers
(6.1) and that barbers (4.0) outperform bartenders (3.6). The extremes are
particularly telling. On the lowest end, one finds corner street drug dealers (1.9)
and panhandlers (2.1), as well as table clearers (2.3) and the loaded category of
agricultural migrant workers (2.7). On the opposite extreme, one finds surgeons
(7.7),191 astronauts (7.4), and mayors of large cities (7.2). While some of these
allocations are sensible, others appear harsh and unfair. In fact, there are signs of
racism, ageism, and sexism in the allocation of status among occupations.192 So
there is nothing particularly compelling about the current occupational status
distribution. But as long as the idea of occupational status is not contested, courts
may want to extend defamation law’s protection in this area.
Noxious status games. The last set of status games are those that prove
pernicious due to their negative social externalities. At one point in history,
alleging that a person was a “bastard” was a matter of great offense, involving
deeply held social mores of wedlock and matrimony. 193 This view reflected what
188 BILL & MELINDA GATES FOUNDATION, https://www.gatesfoundation.org (last visited Jan.
29, 2022).
189 IRENE CARA, FAME (RSO Records 1980).
190 Tom W. Smith & Jaesok Son, Measuring Occupational Prestige on the 2012 General
Social Survey, 122 GSS METHODOLOGICAL REP. (2014), https://perma.cc/58ZW-AKV7. The
relative ranking of occupations appears fairly robust to the manner in which the question is asked.
Margaret M. Marini, Occupational and Career Mobility, ENCYCLOPEDIA SOCIO. 1989 (2d ed.
2000).
191 Worryingly, lawyers (6.4) rate below medical doctors and narrowly overtake social
scientists (6.2).
192 See, e.g., Wun Xu & Ann Leffler, Gender and Race Effects on Occupational Prestige,
Segregation, and Earnings, GENDER & SOC. 377, 383-84 (finding race and gender effects);
Michael Hout et al., Prestige and Socioeconomic Scores for the 2010 Census Codes, 124 GSS
METHODOLOGICAL REP. 13 (2016) (reporting some evidence of race and gender effects); Anthony
Lemelle, The Effects of the Intersection of Race, Gender and Educational Class on Occupational
Prestige, 26 WESTERN J. BLACK STUD. 89 (2002) (finding that “race, gender and educational class
are important in the distribution of occupational prestige”).
193 BLACKSTONE, supra note 139 at *433; Harris v. Nashville Tr. Co., 162 S.W. 584, 585
(Tenn. 1914) (holding that it is “libelous per se to charge one in print or writing with being
illegitimate”); Jerald v. Huston, 242 P. 472, 474 (Kan. 1926) (“[C]ast[ing] aspersions on a man's
pedigree … [is] slanderous per se.”).
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 33/51
sociologists call a “closed stratification system” where status is “ascribed” based
on one’s pedigree.194 Slowly, society moved to a more open stratification system
where status is “achieved,” meaning that status mobility was possible based on
one’s accomplishments. 195 As society opened,196 judicial attitudes towards
bastardy started changing in the 1960s, culminating in a 1997 decision, where a
court simply shrugged off such allegations as patently unimportant.197
Disputes in this area point to an underlying status game of “legitimacy,” and
it is one that modern society rejects. It joins a larger class of noxious status games
involving immutable characteristics, such as race, ethnicity, and sex.198 The task
of identifying the status game involved does not require expertise that courts lack.
In many of these cases, one does not even feel the need for an overarching status
theory to know that the underlying status games are socially venomous.
5. Regulating Status Games: Legitimacy and Institutional Capacity
Calling for courts to openly embrace their role in regulating status games
raises several objections. One objection comes from James Whitman’s analysis of
the laws of civility.199 His account suggests that American law lacks the cultural
foundations to deal with the regulation of norms of civility. Another related issue
is the law’s legitimacy in regulating status games, a deeply social phenomenon.
Finally, there is also a narrower but no less important institutional concern about
the capability of courts to intervene in status games effectively.
This section grapples with these issues. The response to the first and second
concerns is a demonstration of the depth of American law’s interest in the
regulation of status. The response to the third concern is, for the most part, an
open acknowledgment of the problem.
Let us first consider the concern that American society either lacks status
games or is disinterested in regulating them. This objection is found most
forcefully in Whitman’s influential critique of norms of civility.200 This account
holds that American law either lacks the interest or the foundation to regulate
what he calls “civility rules,” a concept that roughly maps onto status games.
194 HUGHES & KROEHLER, supra note 207, at 176-177.
195 Id.
196 Congleton suggests that the move to status on the basis of merit rather than heritage is one
of the sources of strength for capitalist societies. Congleton, supra note 3, at 188. Under this view,
the move from status to contract may be seen as a change not so much in legal technology but in
the type of status games played.
197 Levinsky's, Inc. v. Wal-Mart Stores, Inc., 127 F.3d 122, 128 (1st Cir. 1997) (“For better
or worse, our society has long since passed the stage at which the use of the word “bastard” would
occasion an investigation into the target's lineage or the cry “you pig” would prompt a probe for a
porcine pedigree.”); Bolton v. Strawbridge, 156 N.Y.S.2d 722, 723 (Sup. Ct. 1956) (“Despite
their vulgarity and profanity, the words ‘bastard’ and ‘no good’ have been held not slanderous per
se and not actionable without proof of special damage.”).
198 Society seems to tolerate allocation of status on the basis of some immutable traits, such
as beauty, intelligence, height, and physique.
199 Whitman, supra note 54.
200 See id.
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 34/51
According to Whitman, modern era American, German, and French cultures
have all leveled the distribution of status, motivated by an egalitarian ideal of
social equality. The difference is that German and French societies, drawing on
their aristocratic traditions, decided to level-up status—treating everyone equally
with the respect once reserved for aristocrats.201 But the U.S., which lacks these
traditions, has “leveled-down” civility and thus endorses treating everyone as a
commoner, with equal (dis)respect.202 The upshot is that “American incivility is
woven into the cloth of the American egalitarian tradition,”203 which means that
“in general, America has no law of civility.”204
If this theory only claims that there are differences in the manifestations of
civility norms or the status games played between these societies, it is obviously
true. But if the contention is that the essence of American society is “incivility,”
i.e., lacking strict rules and regulation of social behavior due to the elision of
social status differences, this is a profound error—and a harmful one at that. To
be sure, the presence of status and class in American society is not always explicit
and so it may be easy to miss. People always take their own culture to be a natural
reflection of the way things should be.205 This is especially true in the context of
American culture, which is said to deem taboo the acknowledgment of class and
status differences.206 But as sociologists universally recognize, civility norms are
dyed in the wool of American society,207 and the law is hardly disinterested in
the regulation of status games.
When I first immigrated to the United States, I found myself hurtling at an
invisible wall of civility rules and status games. Each person was carrying around
them an invisible perimeter of space that belonged to them—their ‘personal
space’—and while everybody knew the metes and bounds of each person’s
perimeter, I did not. 208 Worse, the dimension of the invisible perimeter changes
201 Whitman contrasts American law in particular with the German doctrine of insult, which
gives rise to an action for simply showing disrespect. See generally id. at 1297.
202 See id. at 1387-90.
203 See id. at 1398.
204 See id. at 1384.
205 See, e.g., DAVID GRAEBER, DEBT: THE FIRST 5000 YEARS, 122 (2012) (“Consider the
custom, in American society, of constantly saying "please" and "thank you." To do so is often
treated as basic morality … but [based on comparative cultural analysis] it is not.”).
206 Joan C. Williams, Marina Multhaup, & Sky Mihyalo, Why Companies Should Add Class
to Their Diversity Discussions, Harv. Bus. Rev. (Sept 5, 2018) (“[I]n the United States, talking
about class is taboo”).
207 See MICHAEL HUGHES & CAROLYN J. KROEHLER, SOCIOLOGY: THE CORE 177 (2011)
(“The United States is founded neither on the idea that all people should enjoy equal status nor
on the notion of a classless society.”); see also GRAEBER, supra note 205, at 122-24 (arguing that it
is American middle-class behavior that treats everyone with “feudal deference”).
208 See Agnieszka Sorokowka et al., Preferred Interpersonal Distances: A Global Comparison,
48 J. CROSS-CULTURAL PSYCH. 577 (2017) (reporting the preferred interpersonal distances in a
survey of 42 countries).
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 35/51
based on one’s gender, age, and social status.209 You wouldn’t stand as close to
your boss or teacher as you would with a classmate, but I did. And how about
the rules that dictate how and when you can interject in your interlocutor’s
speech? The rules are not written anywhere, but as someone who still unwittingly
breaks them, I assure you that they exist.210 Time norms are another sore topic,
for someone from a culture who has a different set of rules on timeliness. The
idea that American society is somehow ‘incivil’ at its core thus appears far-
fetched.
Status norms are not only everywhere in the United States, they are also
hierarchical and not leveled by any measure. A working-class person goes to the
beach; a high-class person summers at “the Vineyard.” The CEO can approach
the frontline employee, tap them on the shoulder, and say “good work”—but
the worker may not return the favor after a positive earnings call. Between the
person who repairs your car and the person who repairs your body, you can only
“Hi, man” one of them.
Paul Fussell catalogs class differences in the 1980s that still feel mordant
today.211 Working-class Americans are fans of Football, middle-upper class
Americans follow tennis and golf; one class wears clothes with conspicuous brand
names plastered on them, the higher class finds such behavior ‘tacky’; one class
finds the possessive apostrophe redundant in communication, the other finds
mixing ‘its’ and ‘it’s’ to be an affront against all that is sacred in this world.212
The empirically minded could gauge a town’s class, Fussell argues, by measuring
its bowling-alleys-per-capita.213 What defines status games in the United States
is not their absence, but the pretense of their absence— the ethos of having
abolished class and status in favor of merit and mobility.
Any fair observer of social life in America is bound to conclude that status
permeates society even large parts of the legal system.214 It takes willpower to
resist the call of cultural Marxism to construe the sustained insistence that
“America has no law of civility”215 as being itself a mark of class and class
ideology. The legal system (or state apparatus, to those lacking willpower) takes
civility with great solemnity. There is, after all, an entire branch of government
that dresses its members in special regalia, insists on referring to them as your
209 Ann Leffler et al., The Effects of Status Differentiation on Nonverbal Behavior, 45 SOCIAL
PSYCHOLOGY Q. 153, 154 (1982) (Summarizing research showing that “persons of higher status
have more and better [personal] space for their use than do persons of lower status”).
210 Id. at 159 (finding that status differences predict speech interruptions).
211 See generally FUSELL, supra note 5. For a review and discussion, see Scott Siskind, Book
Review: Fussell on Class, ASTRAL CODEX TEN, https://perma.cc/W9M2-U3F4.
212 FUSELL, supra note 5, 114-16.
213 Id.
214 Nestor M. Davidson, Property and Relative Status, 107 MICH. L. REV. 757, 812 (2009)
(noting that “law both reinforces and undermines property's hierarchical signaling” and the
“intimate involvement of the state in what might at a remove seem a private dynamic”).
215 See id. at 1384.
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 36/51
honor, and makes the expression of contempt towards them a criminal offense.
In 2020, a Rhode Island man discovered these civility laws; when the judge read
his judgment, he impolitely said, “that’s bullshit”—to which the judge
responded by condemning him to three years in prison.216
Beyond the judicial branch, buried in history are numerous attempts to
regulate away unwanted status games such as honor duels,217 blood feuds,
potlatch traditions, and, more contemporarily, street racing.218 Sometimes the
law is invoked not to outlaw status games, but to moderate them. The use of tax
law often exemplifies this.219 Between 1990 and 2002 a luxury tax was applied
to yachts, jewelry, expensive furs, and private jet planes.220 The idea, owing to
John Stuart Mill, was that if these goods are purchased because they are
expensive, no harm will befall society from taxing them.221 Tellingly, the
regulation of status games is selective; higher education was not taxed, despite
the surge in costs.222 While Congress has since abandoned the luxury tax, some
216 In re Lamontagne, 228 A.3d 631 (R.I. 2020) (remanding for resentencing and finding a
sentence of more than six months excessive); see also People v. Sweat, 23 N.E.3d 955 (N.Y. 2014)
(“[A] court may hold a person in criminal contempt for … contemptuous, or insolent behavior”
that may “impair the respect due to [the court’s] authority.”).
217 Weber notes that in Germany, army officers were legally required to participate in duels
even though the criminal code prohibited this practice. MAX WEBER, ECONOMY AND SOCIETY: AN
OUTLINE OF INTERPRETATIVE SOCIOLOGY, 318 (1968). See also Hassani Mahmooei et al.,
Dueling for Honor and Identity Economics, MUNICH PERSONAL REPEC ARCHIVE (Jul. 30, 2012),
https://perma.cc/27MC-58ZH (arguing that duels served an organizing social function and
emerged within the aristocracy, but became a middle-class institution in France and Germany).
218 Congleton, supra note 3, at 183 (discussing potlatch). The potlatch serves as a cautionary
tale because many think it insidious to impose European values on indigenous people.
219 On the idea of Pigouvian taxes and subsidies on status games, see Congleton, supra note
3, at 182-183; David Jinkins, Conspicuous Consumption in the United States and China, 127 J.
ECON. BEHAV. & ORG. 115 (2016) (“Luxury taxes on … conspicuous goods skew consumption
back toward the no-signaling optimum.”). One study finds that status-driven concerns lead to
excessive consumption and under-saving. Nick Feltovich & Ourega-Zoe Ejebu, Do Positional
Goods Inhibit Saving? Evidence from a Life-Cycle Experiment, 107 J. ECON. BEHAV. & ORG. 440
(2014).
220 Omnibus Budget Reconciliation Act of 1990 Pub. L. No. 101–508, 104 Stat. 1388
(1990), 42 U.S.C. § 1396 (1990). The scope of the tax changed throughout this period. On the
reaction to the tax, see, e.g., Kevin E. Cullinane, The Bush Budget: Luxury Tax is a Luxury Nation
Cannot Afford, Industries Say, L.A. TIMES (Jan. 31, 1992 12:00 AM),
https://www.latimes.com/archives/la-xpm-1992-01-31-fi-1159-
story.html#:~:text=In%20his%201993%20annual%20budget,was%20proposed%20for%20thos
e%20goods.
221 5 JOHN S. MILL, PRINCIPLES OF POLITICAL ECONOMY WITH SOME OF THEIR
APPLICATIONS TO SOCIAL PHILOSOPHY, ch. 6, pt. 7 (1848). For an alternative view of the luxury
tax, see Joseph Bankman & David A. Weisbach, The Superiority of an Ideal Consumption Tax
over an Ideal Income Tax, 58 STAN. L. REV. 1413, 1428 (2006).
222 Higher education is not only a status good, but it is hard to ignore the status qualities of
“being educated.” On costs, see https://research.collegeboard.org/trends/college-pricing/figures-
tables/growth-in-published-charges.
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 37/51
states still employ a “mansion” tax on luxury homes,223 and similar taxation is
common around the world.224
There are other status games that the law wants to encourage. and one way of
doing it is gatekeeping who can claim status. A particularly clear demonstration
comes from the Stolen Valor Act—an attempt to regulate status by rationing its
allocation only to war heroes.225 The Supreme Court struck down its original
version on First Amendment grounds, but Congress has shown incredible vigor
and rare unanimity in passing a revised (albeit weaker) version.226 Or consider
the 1978 Act designating the exclusive right to use the word “Olympic” to the
United States Olympic Committee (USOC).227 When a California non-profit
sought to promote awareness to gay rights by organizing an event called the “Gay
Olympic Games,” the USOC objected. The case reached the Supreme Court,
which approved limiting freedom of speech to foster status exclusivity.228
And, of course, trademark law is a central locus of status regulation. While
the law has various functions—prevention of confusion chief among them229—
it is hard to understand other parts without invoking notions of status games.
Consider the post-sale confusion doctrine, which has little to do with confusion
and a lot to do with status dilution.230 This doctrine applies to a situation where
a competitor sells counterfeit items to buyers who know (and are thus not
confused) that they are purchasing a counterfeit at a presumably lower price.231
This doctrine fights such sales because of concern with status:232 If they can
afford a nice Rolex, why should we buy one?233
223 Michael Leachman & Samantha Waxman, State “Mansion Taxes” on Very Expensive
Homes, CENTER ON BUDGET AND POLICY PRIORITIES (Oct. 1, 2019)
https://www.cbpp.org/research/state-budget-and-tax/state-mansion-taxes-on-very-expensive-
homes.
224 Nadine Schmidt & Sheena McKenzie, Tampons Will No Longer be Taxed as Luxury
Items After Landmark German Vote, CNN: WORLD (Nov. 8, 2019, 4:24 AM),
https://www.cnn.com/2019/11/08/europe/tampon-tax-germany-luxury-item-grm-
intl/index.html.
225 18 U.S.C § 704 (2006).
226 U.S. v. Alvarez, 567 U.S. 709, 737 (2012). Stolen Valor Act of 2013 (Pub.L. 113–12)
(passed unanimously in the senate and 390-3 in the House), https://perma.cc/9SET-L2V7.
227 Amateur Sports Act of 1978, Pub. L. No. 95-606, 92 Stat. 3045 (1978).
228 San Francisco Arts & Athletics v. U.S. Olympic Comm., 483 U.S. 522, 539 (1987)
(“[M]uch of the word's value comes from its limited use.”).
229 15 U.S.C. § 1066 (2018).
230 See United States. v. Gillette Co., 828 F. Supp. 78, 80-82 (D.D.C. 1993); Mark P.
McKenna, A Consumer Decision-Making Theory of Trademark Law, 98 VA. L. REV. 67, 104
(2012). On the history of the doctrine, see Connie D. Powell, We All Know It’s a Knock Off -
Re-Evaluating the Need for the Post-Sale Confusion Doctrine in Trademark Law, 14 N.C. J.L.
TECH. 1, 17-24 (2012).
231 See Irina D. Manta, Hedonic Harms, 11 OHIO ST. L.J. 241, 268-69 (2013).
232 See, e.g., Rolex Watch U.S.A. v. Canner, 645 F. Supp. 484, 493, 495 (S.D. Fla. 1986)
(offering a mixed reputation and status based rationale for the doctrine).
233 Id. at 495; see also Hermes Int’l v. Lederer de Paris Fifth Ave., Inc., 219 F.3d 104 (2d Cir.
2000) (describing the harm as individuals “achieving the status of owning the genuine article at a
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 38/51
As these examples illustrate, American law is brimming with status concerns
which reflect a deep cultural interest in status games.234 Indeed, status is so deeply
embedded in the American system that some believe the law should intervene to
shelter individuals from status games. As Martha Nusbaum argues: 235
Social groups will continue to inflict shame on others with or
without the cooperation of the law, so the law needs to do more
than simply refuse to join in this behavior. It should actively
protect the individual who may want a place of retreat from the
shame that inevitably will continue to attach to unusual people
and behavior.
Seeing the deep involvement of the law in status games helps assuage concerns
of legitimacy. But it does leave open the question of capacity—how capable are
sitting judges and legislators of making good determinations on the regulation
of status games? This is a larger question, and it involves not just technical
expertise but also questions of ideology. My personal view is that courts and
legislators should make such determinations cautiously and rarely—but I think
this question requires much deeper analysis than the present scope allows. What
matters most, I think, is that judges are already making such determinations
when they determine defamation law cases and that there is no divorcing case
outcomes from the regulation of status games.
***
The social drama associated with loss of status, insult, and humiliation has
captivated audiences throughout human history. This drama, inherent to any
defamation lawsuit, can easily distract us and make us lose sight of broader
considerations, mainly, the status games that produced the lost status and whose
preservation is now at stake. When we turn our attention to these status games,
we see that American law has a keen interest in them, although it tends to do so
in a particularly American fashion—focusing on commercial trademarks and
military valor. The investigation revealed that not only are status claims
implicated in all defamation lawsuits, but also that the law proactively maintains
some and dismantles others.
The recognition that defamation law is enlisted to stabilize and destabilize
status games opens the way to new defamation law jurisprudence. In the new
jurisprudence, judges openly confront the underlying status game and ask
whether it is one worth preserving. If so, judges perhaps question whether more
knockoff price”). See also Jeremy N. Sheff, Veblen Brands, 96 MINN. L. REV. 769, 790-804 (2012)
(discussing Kal Raustiala & Christopher J. Sprigman, Rethinking Post-Sale Confusion, 108
TRADEMARK REP. (2018) (noting the framing of the doctrine in “consumers’ generalized desire for
exclusivity and specialness”).
234 For a comprehensive analysis, see Richard H. McAdams, Relative Preferences, 102 YALE
L.J. 1 (1992).
235 MARTHA C. NUSSBAUM, HIDING FROM HUMANITY: DISGUST, SHAME, AND THE LAW 297
(2004).
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 39/51
zealous protection is actually needed. If it is a noxious status game, judges would
explicitly refuse to lend it any protection. Large legal pivots in light of changing
mores are especially apt to legislative determination, and legislators may come to
play a larger role in the new jurisprudence. Status theory and the evolving
jurisprudence around it is not just an abstract way of thinking about defamation
law, it directly contributes to the decisions courts and legislators make in specific
cases. Demonstrating this claim is the task of the following Part.
CASE STUDIES
1. Racist and Bigoted Speech
One of the deepest quagmires of defamation doctrine is that of bigoted
defamation. This category contains such allegations as asserting that a Christian
person is a Jew, that a white person is of a different race, that a straight person is
gay, or that a cisgender person is transgender.236 These cases feature a plaintiff
who claims that these allegations are false and that they have impugned her social
standing within her bigoted community; thus, she seeks recompense for the harm
she suffered.
Such claims are deeply disturbing: a person sues so they can continue to keep
a privilege that results from a social hierarchy that humiliates others. Yet lawsuits
in this fashion are frequently made, presenting a vexing problem to the standard
view of defamation law that portrays defamation law as simply reacting to and
redressing harm. If harm to social standing is all there is, then it is difficult to
deny that bigoted statements cause a person to lose their (illegitimate) standing
within a bigoted community—indeed, the bigoted accusation was probably
made for that reason.
This reasoning emboldened courts to find defamation in a host of
problematic cases, such as when a straight person was called gay,237 or when a
white person was alleged to be black.238 Worse, if judges are seen as simply
reacting to harm—rather than enforcing social norms of behavior—they evade
accountability for the implications of their decisions. As noted earlier, Judges
were actively portraying themselves as disinterested social scientists who make
236 See generally John Watson, Defamation by a Racial Misidentification: A Study of the
Social Tort, 4 RUTGERS RACE & L. REV. 77 (2002).
237 See generally Anthony Michael Kreis, Lawrence Meets Libel: Squaring Constitutional
Norms with Sexual-Orientation Defamation, 122 Yale L.J. Forum 125, 128 (2012). See also
Lidsky, supra note 33 (“Courts have been slow to embrace a progressive view by declaring that an
allegation of homosexuality cannot be libelous.”).
238 See, e.g., Eden v. Legare 1 S.C.L. 68, 1 Bay 171, at 71 (1791) (finding that allegation that
a white person is black is “calculated to inflict injury”); Bowen v. Independent Pub. Co., 96 S.E.2d
564 (S.C. 1957); Stultz v. Cousins, 242 F. 794 (6th Cir. 1917) (holding that it was libelous to
allege that a white man was black). See generally John C. Watson, Defamation by a Racial
Misidentification: A Study of the Social Tort, 4 RUTGERS RACE & L. REV. 77, 104 (2002).
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 40/51
“objective” determinations of harm, factual determinations that led them to the
‘inevitable’ conclusion that belonging to an ethnic group is defamatory.239
Fortunately, courts are increasingly pivoting away from their old positions.240
This development is tempered by a pernicious compromise compelled by the
standard model of reaction to harm: court can avoid a finding of defamation
only by engaging in the subterfuge that society abolished bigotry.241
The logical chain leading to such artifice starts and ends with the issue of
harm. The problem in bigoted defamation cases is that bigoted statements to
bigoted communities do cause harm—it is the very definition of a bigoted
community. When courts decided they wanted to circumvent a finding of
defamation, the only route they thought was open to them was to deny the
existence of this harm. Courts did so by limiting the scope of the audience,
focusing not on the actual, relevant audience, but on “right-thinking” parts of
society or a “substantial and respectable minority” to the exclusion of bigoted
groups.242 Courts realized, however, that even this measure may lead to
unwanted conclusions: to this day, there are non-negligible parts of society that
are bigoted yet enjoy social esteem. Thus, courts had to stretch the ‘finding’ of
harm further, and hold that evidence for what these “right-thinking” people
actually think is not a matter for factual determination, but rather, one of judicial
intuition.243 Armed with this double artifice, courts were able to ‘find’ that a
statement that a woman “would do anything for five dollars” did not impute
unchastity in 1956,244 or that a derogatory term hurtled at an alleged gay man
ceased being derogatory in 1977.245 If concluding that there is nothing
defamatory in belonging to a specific race or ethnicity requires erasing bigotry
and invalidating the experiences of those marginalized by bigoted
communities—well, this is a serious issue for the legal system.
Scholars wrestled with these tensions. In an insightful early article,
defamation law scholar Lyrissa Lidsky suggested to resolve the tensions by
dividing the investigation into an objective, empirical determination of harm
239 See supra note 27
240 See, e.g., Mitchell v. Tribune, 99 N.E. 2d 397 (Ill. App. Ct. 1951) (holding that it was
not libelous to refer to a white man as black); Thomason v. Time-Journal, Inc., 379 S.E.2d 551
(Ga. App. 1989); Jay Barth, Is False Imputation of Being Gay, Lesbian, or Bisexual Still
Defamatory? The Arkansas Case, 34 U. ARK. LITTLE ROCK L. REV. 527, 528 (2012) (“In recent
years, however, courts have become conflicted on whether a false imputation of a person as LGB
is defamatory.”).
241 Lidsky, supra note 33, at 10 (“The resulting subterfuge is a natural outgrowth of an inquiry
that has little to do with actual harm and even less to do with the actual community segment whose
opinion the plaintiff values.”).
242 Id. at 7.
243 Id. at 8 (“[C]ourts rely on their own intuitive judgments about who constitutes the relevant
community, what values that community shares, and whether those values are respectable.”).
244 Bolton v. Strawbridge, 156 N.Y.S.2d 722, 724 (Sup. Ct. Westchester Co. 1956).
245 Moricoli v. Schwartz, 361 N.E.2d 74, 76 (Ill. App. Ct. 1977). Illinois was indeed the first
state to decriminalize sodomy, in 1961. See generally Kreis, supra note 237, at 125.
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 41/51
within a community, and an explicitly normative element: choosing a specific
community whose opinions matter. 246 Making the normative step explicit, she
argued, would “reinforce[] defamation’s symbolic role in the definition,
affirmation, and enforcement of community values in America.”247 In a later
article, Lidsky explains that this requires an evaluation of the reaction of a
“rational” audience, rather than the actual audience.248 In contrast, David Han
argued that courts should focus on the empirical element and predict how a
“targeted audience will likely process the speech, rather than on a strong
normative view of how an idealized ‘rational audience’ should process the
speech.”249
Their differences aside, these scholars both compellingly argue that there is
something deeply artificial about courts using a harm-based standard while
ignoring evidence of actual harm.250 Both of these accounts, however, still rely
on a reactionary, harm-centered approach. While they offer plausible solutions
to the problems of harm from bigoted speech within the existing framework,
status theory offers a direct approach that avoids the problem altogether. Under
the status approach, the question posed in bigoted speech cases is not whether,
as a matter of fact, a given community would judge a person negatively based on
their race. It is also not about whether the communities that view gay individuals
in a negative light are ‘substantial’, ‘respectable,’ or ‘rational’—intolerance to
LGBTQs is still a live issue, only recently starting to retreat from the mainstream.
What matters, instead, is the nexus between defamation lawsuits and status
games. When the plaintiff prevails in a defamation lawsuit, she receives money
damages which allow her to recoup her investment in status attainment. The
shadow of such payments deters would-be defamers from making false
allegations. The lawsuit also vindicates the plaintiff’s good name, alerting other
players in the status game as to whether her claim to status is rightful or not. On
the flip side, when a court denies the ability to bring defamation lawsuits, it
disrupts the status game. The denial makes it harder to know who claims status
honestly, and easier to make unfounded claims.
The solution to bigoted defamation is plain: disrupt the underlying racial
status game by denying defamation protection. If an individual suffers harm to
a status privilege in a racial status game, feigning that harm does not ‘really’ exist
is counterproductive. It is exactly because a harm exists that the status game is
worrisome. Courts of law should openly acknowledge that the claim is
246 Id. at 48.
247 Id. at 49.
248 See Lyrissa B. Lidsky, Nobody's Fools: The Rational Audience as First Amendment Ideal,
2010 U. Ill. L. Rev. 799 (2010).
249 David Han, The Mechanics of First Amendment Audience Analysis, WM. & MARY L.
REV. 1647, 1653 (2014).
250 Lidsky, supra note 248, at 838-49 (explaining why focusing on actual audiences rather
than “rational” audiences can result in various democratic harms).
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 42/51
illegitimate because it arises out of a status game that the court will not reinforce,
legitimize, or even ignore.251 Courts should pointedly say: we reject bigoted
status games.252
Such a principled holding would communicate to members of bigoted
communities that not only are ethno-racial status games illicit, but their gains
are also fleeting. When participating in such status competitions, members of the
majority often confront the threat that some rumor would undermine the
‘purity’ of their pedigree. With defamation law, they can deter such rumors and,
once made, they can enlist the legal system to publicly disprove them, thus
vindicating their claim to privilege. Without defamation law, however, majority
members will find their gains tentative and unprotected, because they are always
a rumor away from being challenged. Shrinking defamation law would work,
then, to destabilize the status game for majority members. At the same time,
shrinking defamation law’s reach would a blessing for minority members. After
all, defamation law was never available to them. This is because in their case, a
public statement about their identity would never be granted the protection of
defamation law—no matter how harmful within his bigoted community, a
public allegation that a gay man is gay cannot be defamatory if it is true.253
2. Collaborators & Snitches
Law enforcement requires the assistance of collaborators, but, in some parts
of society, cooperating with the government carries a social stigma.254 This gives
rise to an interesting dilemma in defamation law jurisprudence.
Take, for example, the Saunders case.255 There, a local TV station reported
that the plaintiff, an inmate, was an FBI informant. Saunders sued for
defamation, alleging harm to his social standing among his community of
inmates which caused him “physical and mental damage.”256 If the lawsuit was a
matter of negligent exposure to bodily harm or intentional infliction of
emotional harm, the matter would not be so problematic. But as the lawsuit was
brought in defamation, the court faced a palpable conundrum. Holding that
cooperation with law enforcement is a source of humiliation sends the wrong
message. But it is also hard to deny the fact that the plaintiff suffered real harm
within his community—a fact that the Saunders court reluctantly recognized.
251 It should be obvious that tort law, and defamation law in particular, cannot completely
eradicate status games. Their modest goal is only to increase the fragility of them.
252 See Richard Delgado, Words that Wound: A Tort Action for Racial Insults, Epithets, and
Name-Calling, 17 Harv. CIV. RTS.-CIV. LIBERTIES L. REV. 133, 140-41 (1982) (noting the
harmful social effects of racial stigmatization). I am unable to address here status games that involve
falsely passing as a member of a minority group, but the criteria developed here offers a clue.
253 Other laws may be available in such scenarios—intentional infliction of emotional harm
is a prominent example. But such laws remain unaffected by the scope of defamation law.
254 See generally ALEXANDRA NATAPOFF, SNITCHING: CRIMINAL INFORMANTS AND THE
EROSION OF AMERICAN JUSTICE (2009).
255 Saunders v. Bd. of Dirs., WHYY-TV (Channel 12), 382 A.2d 257 (Del. Super. Ct. 1978).
256 Id. at 258.
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 43/51
The court’s holding in the Saunders case involved a maneuver that should be
familiar by now. The court said that the “opprobrium” suffered by the informant
was insufficient because its effect was confined only to a “limited community in
which attitudes and social values may depart substantially from those generally
prevailing which an action for defamation is designed to protect.”257 Thus, the
court rejected the plaintiff’s claim. The problem is that the reactionary
protection-from-harm view has no problem protecting other minority views only
held by small communities.258 In Air Wisconsin Airlines v. Hoeper, for example,
the Supreme Court found that the statement that the plaintiff “was an FFDO
who may be armed,” was defamatory in the eyes of the “reasonable TSA
officer”—hardly a large segment of the population.259 There is also nothing that
hangs on the divergence of values within a small community. Courts were willing
to find defamation even though the defamatory statement was only offensive
within the sub-ethnic community of Vietnamese immigrants.260
What the court should have done in the Saunders case is radically
straightforward. Rather than employing the condescending criteria that inmates
are not “right-thinking individuals,”261 the court should have said that it
recognizes that some communities play status games around contempt for law
enforcement and fidelity to violent organizations. Consistent with the analysis
above,262 the court does not endorse such status games, and so it refuses to lend
defamation’s law protection to status claims resulting from these games. To an
extent, such a decision can destabilize status pursuits in these illicit status games,
and it is therefore justified.
257 Id. at 259. See also RESTATEMENT OF TORTS, supra note 7, at § 559 (defamation, even in
the eyes of “a substantial group is not enough if the group is one whose standards are … anti-
social”).
258 PROSSER AND KEETON, §111, at 777 ("[A] plaintiff may suffer real damage if he is lowered
in the esteem of any substantial and respectable group, even though it may be quite a small
minority."). Courts have also held that a statement which only hints at the identity of the plaintiff
is still defamatory, as long as there are “some who reasonably” identify the plaintiff. SMOLLA, supra
note 127, at § 4:44.
259 See Air Wisconsin Airlines Corp. v. Hoeper. See also David S. Ardia, Reputation in a
Networked World: Revisiting the Social Foundations of Defamation Law, 45 HARV. CIV. RTS.-
CIV .LIBERTIES L. REV. 261, 283 (2010) (arguing that defamation law is primarily concerned with
“the impact of the statement on those who make up the plaintiff’s community.”); PROSSER AND
KEETON., supra note 126.
260 Clay Calvert, Difficulties and Dilemmas Regarding Defamatory Meaning in Ethnic
Micro-Communities: Accusations of Communism, Then and Now, 54 U. LOUISVILLE L. Rev. 1
(2016).
261 In Connelly v. McKay, 176 Misc. 685, 28 N.Y.S.2d 327 (N.Y. Misc. 1941) the court
ignored the views of interstate truck drivers, who shunned a service station managed by the plaintiff
who was alleged to have been an informant for the Interstate Commerce Commission.
262 See supra Part II.4.
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 44/51
3. Female Sexual Autonomy
The last case study involves allegations that a woman is unchaste. According
to modern doctrine, allegations that a woman is “promiscuous” are so harmful
that they belong in a special category of “per se” defamatory statements.263 Per
se defamation means that harm to reputation is assumed, relieving the plaintiff
of a significant hurdle. A modern reader would likely be puzzled why this rule’s
existence and its special application to women. Because jurists are used to
thinking about defamation law through the prism of harm prevention, rather
than the regulation of social norms, they might be tempted to use this logic here
as well. Indeed, Prosser himself argued that this rule simply reflects harm
differentials: “Such a rule never has been applied to a man, since the damage to
his reputation is assumed not to be as great.”264 On close examination, however,
the reactionary model conceals a far more troubling reality.
The first sign of trouble was noted by Post, who showed that Prosser’s logic
is incongruous with the fact that the presumption of harm to women is
irrebuttable.265 It is one thing to say that harm to social standing is hard to prove,
it is entirely another to say that it cannot be proven not to exist.
A second sign comes from the rule’s history. While female chastity was a
central theme of 18th century England,266 it did not emerge at its moment of
zenith. Rather, it only emerged a century later267—and by a special act of
Parliament, no less268—when chastity concerns (and with them, good name
harms) have started declining.
263 RESTATEMENT (SECOND) OF TORTS § 574 (1977). For some modern examples, see Bryson
v. News Am. Publications, Inc., 174 Ill. 2d 77, 94, (Ill. 1996) (holding that an article referring to
the female plaintiff a “slut” was per-se defamatory); Doe v. Simone, No. CIV.A. 12-5825, 2013
WL 3772532, at *5 (D.N.J. July 17, 2013) (accusations that the female plaintiff was a “slut,” the
“queen of sluts,” and a “whore.”); Walia v. Vivek Purmasir & Assocs., Inc., 160 F. Supp. 2d 380,
394–95 (E.D.N.Y. 2000) (holding as slander per-se the defendant’s statement that the female
plaintiff was a “whore” and a “slut”).
264 PROSSER, supra note 10, at 760; see also Sexton v. Todd, Ohio (Wright) 316, 320-21
(1833) (“[An allegation of sexual impropriety] is vastly more injurious to a female than to our
sex.”).
265See Post, supra note 41 (“The fact that the presumption of general damages is
irrebuttable is inexplicable from the standpoint of the concept of reputation as property.”).
266 See Soile Ylivuori, Rethinking Female Chastity and Gentlewomen’s Honour in
Eighteenth-Century England, 59 HISTORICAL J. 71 (2016).
267 The common law before then did not consider allegations of unchastity to be slanderous
per se. PROSSER AND KEETON, supra note 126, at § 92. However, “[b]y the late 1800s the vast
majority of states had responded to the proliferation of sexual slander suits by designating
statements that impugned a woman’s chastity to be slander per se.” LISA PRUITT, HER OWN GOOD
NAME 4 (2004). See, e.g., Ala. Code § 7359 (1923) (cited in Marion v. Davis, 114 So. 357, 358
(Ala. 1927) and Note, Bases of Slander Per Se in Ohio: Comments, 15 OHIO STATE L.J. 312, 322-
323 (1954)).
268 In Roberts v. Roberts, 122 Eng. Rep. 874 (1864), a man told the plaintiff’s husband that
she was “as great a whore as any in the town of Liverpool.” Lord Cockburn C.J. lamented that he
could provide no remedy absent a showing of special damages, decrying the law as “cruel.” This
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 45/51
The harm perspective cannot explain the rule’s scope, its emergence at a
specific point in time, and the need to use special legislation. But status analysis
provides a considerably more viable explanation. As would be clear to most
modern readers, the underlying status game played in chastity defamation cases
is a sexist status game of ‘purity’, whereby a woman’s status is gained or lost
through exercise of her sexual autonomy.269 This mainstream 18th century status
game faced new challenges in the 19th century, as women started entering the
labor market and were starting to reevaluate their social fetters.270 The hypothesis
would be, then, that Parliament realized that it could breath life into this status
game by reinforcing it through defamation law: making it expensive to impugn
a chaste woman’s name, and cheap to besmirch a ‘promiscuous’ woman. It was
now easier than ever to distinguish oneself based on one’s sexual history, and
‘innocent’ women could easily claim the law’s protection. If true, this explains
why the per se rule only emerged after the status game’s zenith, band then only
by special legislative intervention. This is not to say that any of this was
conscious, but as Professor Pruitt notes, “Nineteenth-century legal rules around
sexual slander thus had unfortunate consequences for women, reinforcing the
social significance of their sexual virtue.”271
This perspective also allows us to rethink a seemingly progressive reform in
the Second Restatement of Torts. As an attempt to make the rule more modern,
the drafters restated it in a gender-neutral fashion,272 winning the praise of
was resolved with the enactment of the Slander of Women Act 54 & 55 Vict c. 51. See generally
LAURENCE H. ELDREDGE, THE LAW OF DEFAMATION 118-19 (1978).
269 Rejent v. Liberation Publ’n, Inc., 197 A.D.2d 240, 245 (N.Y. App. Div. 1994) (“[T]he
notion that while the imputation of sexual immorality to a woman is defamatory per se, but is not
so with respect to a man, has no place in modern jurisprudence.”); SMOLLA, supra note 127, at §
7.05[5] at 7–11, 7–12 (noting that this rule is “quite blatantly sexist and discriminatory, and is
based on outmoded assumptions about sexual behavior”). See generally Wendy N. Hess, Slut-
Shaming in the Workplace: Sexual Rumors & Hostile Environment Claims, 40 N.Y.U. REV. L. &
SOC. CHANGE 581 (2016) (exploring the social double-standard regarding male and female
sexuality). The concept of chastity is far more nuanced than engagement in sexual activity. See
generally Ylivuori, supra note 266.
270 For a critique of these laws, see Lisa R. Pruitt, Her Own Good Name: Two Centuries of
Talk about Chastity, 63 MD. L. REV. 401, 404 (2004). (“[T]he law’s adjudication of [per-se
defamation lawsuits for lack of chastity] has been negative in its reinforcement of society’s
expectations regarding women’s sexual behavior.”).
271 Lisa R. Pruitt, “On the Chastity of Women All Property in the World Depends”: Injury
from Sexual Slander in the Nineteenth Century, 78 IND. L.J. 965, 1015 (2003). Anthony Kreis
criticizes courts that treated allegations of homosexuality as per-se defamatory as being stigmatizing
and inconsistent with substantive due process. Kreis, supra note 237, at 128.
272 RESTATEMENT (SECOND) OF TORTS § 574 (AM. L. INST. 1977). Cf. RESTATEMENT
(FIRST) OF TORTS §574 (AM. L. INST. 1938) (“One who falsely and without a privilege to do so,
publishes a slander which imputes to a woman unchastity is liable to her.”).
Courts have applied this rule to men as well. See, e.g., Sullivan v. Malta Park, 156 So. 3d
1200, 1213 (La. Ct. App. 2014) (holding that the allegation of an extra-marital affair, directed at
a man, was per se defamatory). All cases found that cite to § 574 which are applied to men do not
concern sexual promiscuity in general, only adultery.
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 46/51
commentators for the progressive stance.273 Never mind the fact that this rule is
almost never applied to men,274 the critical point here is that the rule itself should
be abolished. The problem is not with its inequal application, but the chastity
status game itself. The progressive stance is not that both men and women
should be equally shamed for exercising their sexual autonomy. Rather, it is that
both sexes should be free to make sexual choices without being subject to
ridicule, judgment, or humiliation—in other words, that society should refrain
from sexual chastity status games for all sexes. Reforming the Restatement to
‘protect’ both women and men is about as sensible as reforming it to protect both
gays and straights against false allegations concerning one’s sexuality. Neither the
allegation that one is gay nor that one is straight should be considered
defamatory, as the very status game is repugnant. The reactionary harm model
obscures this issue.275 Thus, while policymakers correctly identified defamation
as a vehicle for social change, their application of this insight was misguided due
to a fundamental confusion about the nature of status games.
***
These three case studies illustrate how a clear-eyed view of status games can
guide more principled decision-making in this confused area of law. Still, one
might leave the present discussion with the impression that status theory only
works to restrict the scope of the law, such as by denying the claims of majority
members of bigoted groups.
This is only partly true. Status theory may also be used to ground a much
more capacious role for defamation law, as the regulation of hate speech
demonstrates. In Jeremy Waldron’s Holmes Lecture, he made the case that hate
speech should be regulated by allowing defamation suits based on group libel. 276
Waldron posed what he called “a dignitarian rationale” to the regulation of hate
speech.277
273 See SMOLLA, supra note 127, at § 7.05[5] 7–11, 7–12 (“The Restatement (Second) takes
a laudable lead in this area, modifying the traditional rule to a sex neutral standard that renders
any imputation of ‘sexual misconduct’ by a man or woman slanderous per se.”).
274 Based on an analysis of all cases citing to § 574, only one exception to this rule was found.
See, e.g., Hickerson v. Masters, 226 S.W. 1072, 1073 (Ky. 1921). Modern examples include
Dellefave v. Access Temporaries, No. 99 Civ. 6098(RWS), 2001 WL 286771 (S.D.N.Y. Jan. 11,
2001) (holding that an allegation of a sexual relationship in the workplace was not per-se
defamatory, in particular because the relationship was heterosexual); Ricciardi v. Weber, 795 A.2d
914, 927 (N.J. Super. Ct. App. Div. 2010) (expressing doubt that per-se slander applies to a
“statement made about men as well”). A relatively recent affirmation of the rule is found in Regehr
v. Sonopress, Inc., No. 2:99CV690K, 2000 WL 33710902, at 4 (D. Utah 2000), but cf. Rejent
v. Liberation Publications, Inc., 197 A.D.2d 240, 243 (N.Y. App. Div. 1994) (holding that the
imputation that a male model was lustful was capable of being held as libel per-se).
275 See Kreis, supra note 228, at 128.
276 See generally Waldron, supra note 37.
277 Id. at 1612.
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 47/51
The problem, for Waldron, is that dignity can only do so much.278 Dignity
only offers a limited basis to expanding group libel. Racist speech is surely an
affront to their target’s sense of dignity, safety, and autonomy. But the dignitary
effect is ultimately an empirical fact contingent on the target’s subjective
reaction.279 Many may experience a deep offense, but others may ignore the
malarkey of racists (or at least not feel denigrated by them).280 However, status
theory is not so limited by subjective offense; the problem it identifies is the
objective social status games that racist speech engenders. If hate speech
contributes to the evolution of racial social hierarchies, then this effect is
cognizable social harm, independent of any individual’s experience of dignitary
harm.281 Unfettered by any individual’s reactions, status-based regulation can
thus achieve more than the dignity view alone.282
CONCLUSION
There is no escaping status games—we all play them. Almost every choice we
make is infused with status considerations: the clothes we wear, our choice of
vocabulary, the unconscious decision whether to state a request with “can you”
or “would you.” There is no opting out.
We should care greatly about which status games are played. Some status
games are virtuous, if not in their intentions, then in their effects on the world.
Being a pro-bono lawyer carries with it some status, as it should. Being a social
activist fighting to feed the hungry and vindicate the downtrodden should be a
matter of pride. If we cannot compensate schoolteachers, we might as well respect
them. But with at least the same fervor we should reject those status games that
act to create social hierarchies based on race, ethnicity, sexual identity, and the
like. In the words of Isabel Wilkerson:283
The tyranny of caste is that we are judged on the very things we cannot change: a
chemical in the epidermis, the shape of one’s facial features, the signposts on our bodies
278 This is especially the case with respect to views of dignity as a negative right. See generally
Neomi Rao, Three Concepts of Dignity in constitutional Law, 86 NOTRE DAME L. REV. 183
(2011).
279 Robert Mark Simpson, Dignity, Harm, and Hate Speech, 32 LAW PHILOS. 701, 723
(2013) (critiquing Waldron’s account of harm to dignity as an “exercise in consequentialist
speculation”); see also Eric Barendt, What Is the Harm of Hate Speech?, 22 ETHICAL THEORY
MORAL PRACT. 539–553 (2019) (critiquing Waldron’s account of the harm caused by hate
speech).
280 See Barendt, supra note 279, at 542 (critiquing Waldron’s injury to dignity view).
281 See Robert Mark Simpson, Dignity, Harm, and Hate Speech, 32 LAW PHILOS. 701, 727
(2013) (questioning whether, as a matter of fact, hate speech “contributes to identity-based social
hierarchies”). Notably, Simpson’s account takes a status-based view of dignity.
282 Status theory would thus support the decision in cases such as Taylor v. Metzger, 706 A.2d
685 (N.J. 1998), where racist epithets were held to be capable of amounting to intentional
inflection of emotional harm. Limiting the ability of social agents to enforce racial hierarchies
threatens the underlying status game.
283 WILKERSON, supra note 3.
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 48/51
of gender and ancestry—superficial differences that have nothing to do with who we
are inside.
This Article builds on the sociological concept of status to argue that
defamation law is best seen as regulating status games. This is not all that
defamation law does, of course, but status concerns help explain and justify a
large part of a tort often described as mystifying. Courts are heavily implicated
in the regulation of status games, but their role is often cloaked and
misunderstood. With an explicit understanding of the link between defamation
and status games, society can decide which ones to nurture and which ones to
abandon in the 21st century.
TECHNICAL APPENDIX
The purpose of this Appendix is to offer a technical argument for why the
central justification today of defamation law is unpersuasive. As developed in the
Article, many commentators justify defamation law as a tool of redressing harm
to social standing. This is a decidedly ex-post perspective that ignores any
behavioral effects on the attainment of status. The primary argument developed
is that defamatory remarks do not harm status but redistribute it. As such,
defamatory remarks can be socially neutral or even beneficial in their effect, and
although they would occasionally be detrimental, the law of defamation does not
even attempt to track these instances.
The analysis is motivated by a common idea in the sociological and economic
literature—according to which social status is a zero-sum game. 284 While
individuals belong to a large number of social communities, and status is
evaluated on the basis of multiple dimensions—the social ranking is ultimately
considered to be a closed system. This is why economists consider status to be
the ultimate “positional good,” one “whose value is only defined in reference to
their position on an imaginary scale or ladder.”285 If society can be likened to this
imaginary ladder, it will follow that “[e]ach step up the status ladder for one
person logically requires a step down for another.”286 Sociologist Joel Podlonsky
summarizes this view:
284 Frederic C. Godart & Matthew S. Bothner, What is Social Status, Comparisons and
Contrasts with Cognate Concepts, SEMANTICS SCHOLAR (2009) (defining status as a “zero-sum
relational asset.”). See also Cecilia L. Ridgeway and Henri A. Walker, Status Structures, SOCIO.
PERSPECTIVES ON SOC. PSYCH. 281 (1995) (defining status structures as “rank-ordered
relationships,” which implicitly denotes the zero-sum character of the system); Richard H.
McAdams, Relative Preferences, 102 YALE L. J. 1, 5 (1992).
285 See, e.g., Bougherara et al., supra note 156, at 1229 (“[S]tatus being the ultimate positional
good.”); Congleton, supra note 3, at 178 (“The common element of all status games is that relative
performance rather than absolute performance ultimately determines individual utility levels,
where ‘performance’ is measured by the status-assigning rules of the game of interest.”).
286 Bougherara et al., supra note 156, at 1228.
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 49/51
“Within any social system, status is … zero-sum in character.
One actor cannot increase his status without another losing status.
As a consequence, to the extent that status is the indicator of
interest, it is necessarily the case that high status will not be
available to all actors within a social system.”287
This is not meant as an argument that participation in the status game is zero-
sum; the major takeaway from the Article is that participation is all but zero-
sum. Rather, this Appendix seeks to explore the logical consequences of the
common justification. Taking its ex-post perspective at face-value, and likewise
ignoring any ex-ante effects, the question is whether defamatory statements
indeed cause a social harm. Using two models of status, the analysis shows that,
from a social perspective, the protection justification is unpersuasive.
Ordinal Status. Suppose there is a community of individuals.
The individuals are assumed to be identical in all but their strict rank ordering.
𝑛𝑛 ∈ ℕ ,𝑛𝑛 ≥ 2
The relative social position of individual i is given by , with .
The lowest ranking individual is thus noted by and the highest by . Utility
𝑟𝑟𝑖𝑖 𝑟𝑟𝑖𝑖 < 𝑟𝑟𝑗𝑗,∀𝑖𝑖 < 𝑗𝑗
from status is given by the monotonically increasing utility function , with
𝑟𝑟1 𝑟𝑟𝑛𝑛
the utility of the individual with a status rank and .
𝑈𝑈(.)
Assume that the effect of defamation is to reduce one’s social standing
𝑟𝑟𝑖𝑖 = 𝑈𝑈(𝑟𝑟𝑖𝑖) 𝑈𝑈(0)= 0
by spots. It then follows that:
Corollary 1: With ordinal status, defamation is welfare neutral.
𝑥𝑥 > 0
Proof:
Defamation moves the tth individual from position to position .
But because rankings are relative, this implies that the individual previously in
𝑟𝑟𝑡𝑡 𝑟𝑟𝑡𝑡−𝑥𝑥 < 𝑟𝑟𝑡𝑡
the is moved up to the position, and similarly for any individual
above them with . Thus, total utility, , remains unchanged.
𝑟𝑟𝑡𝑡−𝑥𝑥 𝑟𝑟𝑡𝑡−𝑥𝑥+1
Note that individuals are assumed here to h𝑛𝑛ave similar utility functions, but
𝑟𝑟𝑖𝑖 < 𝑟𝑟𝑡𝑡 ∑1𝑈𝑈(𝑟𝑟𝑖𝑖) ∎
one might plausibly argue that individuals differ in the utility they draw from
status. In such a case, defamation may lead to a net increase or decrease in utility,
but there is no a-priori reason to assume any specific allocation of utility
functions.
Cardinal Status. Suppose that individuals in the community are identical in
all but their initial endowment of an intangible status good, with the total
endowment being .288 The endowment of status goods of the ith
individual is denoted by . This time, however, individuals do not care about
𝑆𝑆 ∈ (0,∞]
their rank directly, but about their distance from others. Individual i'’s total
𝑠𝑠𝑖𝑖
distance from others is given by .
𝑑𝑑𝑖𝑖
𝑛𝑛
𝑑𝑑𝑖𝑖 = ∑𝑗𝑗=1�𝑠𝑠𝑖𝑖 −𝑠𝑠𝑗𝑗� = 𝑛𝑛𝑠𝑠𝑖𝑖 −𝑆𝑆
287 PODOLNY, supra note 93, at 25.
288 Status goods can be thought of as the “accumulation of deference behavior”, as in Michael
Sauder, Freda Lynn & Joel M. Podolny, Status: Insights from Organizational Sociology, 38 ANNU.
REV. SOCIOL. 267, 268 (2012).
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 50/51
Because status is defined as the accumulation of deference behavior or “pellets
of peer recognition,”289 we can think of defamation as destroying some of the
target’s ‘pellets.’ That is, defamation reduces the target’s status goods by x units.
Let be the target’s status ranking post defamation. We can now state the
private′ harm to the target from defamation as:
𝑑𝑑𝑡𝑡
That is, the target’s ′original utility from her status goods less her utility from
|𝑈𝑈(𝑑𝑑𝑡𝑡)|−|𝑈𝑈(𝑑𝑑𝑡𝑡)|=|𝑈𝑈(𝑛𝑛𝑠𝑠𝑡𝑡 −S)|−|𝑈𝑈�𝑛𝑛(𝑠𝑠𝑡𝑡 −𝑥𝑥)−(𝑆𝑆−𝑥𝑥)�|
having x fewer units of the status good.290 Note that, while defamation destroys
some of the target’s status goods, this loss is partially offset by the fact that there
are fewer status goods to go around, which lowers the community average.
Finally, as we are considering the idea of risk to one’s status, it will be fairly
natural to assume that U'<0, U''>0.291 We can now state the following
proposition.
Proposition: Punching up: Defamation, on the margin, increases welfare if
it is directed at a high-status individual. Punching down: Conversely,
defamation reduces welfare on the margin if it is directed at a low-status
individual.
Proof. Defamation destroys x units of status goods, and so the effect of
defamation on total welfare is:
𝑊𝑊
𝑛𝑛−1
The deriva𝑊𝑊tive= w�ith r𝑈𝑈es(p𝑛𝑛e𝑠𝑠c 𝑖𝑖 t −to𝑆𝑆 x +is:𝑥𝑥 )+𝑈𝑈(𝑛𝑛𝑠𝑠𝑡𝑡 −𝑛𝑛𝑥𝑥−𝑆𝑆+𝑥𝑥)
1
𝑛𝑛−1
𝑑𝑑𝑊𝑊 ′
To under=sta�nd t𝑈𝑈he (e𝑛𝑛ff𝑠𝑠e 𝑖𝑖 ct− o𝑆𝑆f s+ma𝑥𝑥l)l c+ha(n1g−es,𝑛𝑛 e)v𝑈𝑈al′u(a𝑛𝑛t(e𝑠𝑠 a 𝑡𝑡 t −𝑥𝑥)− 𝑆𝑆+𝑥𝑥)
𝑑𝑑𝑥𝑥 1
𝑥𝑥 = 0
𝑛𝑛−1
𝑑𝑑𝑊𝑊 ′
Rearranging � = �𝑈𝑈 (𝑛𝑛𝑠𝑠𝑖𝑖 −S)+(1−n)𝑈𝑈′(𝑛𝑛𝑠𝑠𝑡𝑡 −S)
𝑑𝑑𝑥𝑥 𝑥𝑥=0 1
𝑛𝑛
′
By concavity, the to�tal m𝑈𝑈ar(g𝑛𝑛i𝑠𝑠n 𝑖𝑖 al− chSa)n−ge𝑛𝑛 i𝑈𝑈n ′u(t𝑛𝑛il𝑠𝑠i 𝑡𝑡 ty− (tSh)e first expression on the
1
left) will be larger than the change in the target’s utility multiplied by n if and
only if .
𝑆𝑆
𝑠𝑠𝑡𝑡 > 𝑛𝑛 ∎
289 Robert K. Merton, The Matthew Effect in Science, II: Cumulative Advantage and the
Symbolism of Intellectual Property, 79 ISIS 606, 620 (1988).
290 We consider the absolute value of the difference, as and may be negative.
291 See GEOGGREY BRENNAN & PHILIP PETTIT, THE ECONOMY OF ES′TEEM, 83-105 (2005).
𝑈𝑈(𝑟𝑟𝑡𝑡) 𝑈𝑈(𝑟𝑟𝑡𝑡)
Another way to rationalize this assumption is to see that it is plausibly more consequential to a
person’s well-being to move close to those around him rather than to move farther apart. To the
extent that wealth is also a status symbol, the marginal benefit from having a car rather than taking
the bus is arguably greater than being able to afford a newer model.
Electronic copy available at: https://ssrn.com/abstract=4021605

<<PAGE_BREAK>>

ARBEL, A STATUS THEORY OF DEFAMATION LAW 51/51
Corollary 2: The optimal amount of defamation of an individual t is given
by .
P∗roof 𝑛𝑛 . B 𝑠𝑠𝑡𝑡 y − 𝑆𝑆 the Proposition, defamation of any high-status individual is utility-
𝑥𝑥 = 𝑛𝑛−1
maximizing on the margin, hence the optimal degree of defamation x is defined
by . Below that point, the individual becomes low-
status, and punching-down is welfare-minimizing. Rearranging, this implies that
𝑑𝑑𝑡𝑡 = 𝑛𝑛(𝑠𝑠𝑡𝑡 −𝑥𝑥)−𝑆𝑆+𝑥𝑥 = 0
.
∗
𝑛𝑛𝑠𝑠𝑡𝑡−𝑆𝑆
𝑥𝑥 = 𝑛𝑛−1 ∎
Electronic copy available at: https://ssrn.com/abstract=4021605`

func AsMap() map[string]any {
  return map[string]any{
    "paper_id": PaperID,
    "title": Title,
    "ssrn_url": SSRNURL,
    "year": Year,
    "authors": Authors,
    "keywords": Keywords,
    "summary_md": SummaryMD,
    "summary_zh_md": SummaryZHMD,
    "one_pager_md": OnePagerMD,
    "study_pack_md": StudyPackMD,
    "article_text": ArticleText,
  }
}

func AsJSON() string {
  b, err := json.MarshalIndent(AsMap(), "", "  ")
  if err != nil { return "{}" }
  return string(b)
}

func main() {
  if len(os.Args) > 1 && os.Args[1] == "--json" {
    fmt.Print(AsJSON())
    return
  }
  fmt.Print(ArticleText)
}
