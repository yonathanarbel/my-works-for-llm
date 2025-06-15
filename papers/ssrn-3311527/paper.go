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
4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
THE CASE AGAINST EXPANDING
DEFAMATION LAW
Yonathan A. Arbel & Murat Mungan
INTRODUCTION ..................................................................................................... 454
I. DEFAMATION LAW & REPUTATION ......................................................... 461
A. Goals & Functions of Defamation Law .................................................. 461
B. Defamation Law: Doctrine, Nature of Reputation, and the Audience ....... 465
C. Where Defamation Law Ends ................................................................ 469
II. DEFAMATION & AUDIENCES .................................................................... 471
A. Basic Example ....................................................................................... 472
B. Analysis ................................................................................................. 475
C. Audience Effects: Evaluation & Generalization ...................................... 481
D. Qualifications & Richer Considerations .................................................. 483
1. Ignorant Audiences .......................................................................... 484
2. Audience Rationality & Sticky Falsehoods ...................................... 485
3. Imperfect and Costly Enforcement, Litigation, and Execution ........... 486
4. Social and Moral Norms ................................................................. 487
E. Would Expanding Defamation Law Advance Its Goals? ........................ 488
1. Protecting Reputation Through Deterrence ........................................ 488
2. Protecting Reputation Through Redress ............................................. 488
3. Protecting Reputation Through Vindication ...................................... 490
III. DEFAMATION LAW & AUDIENCE EFFECTS ............................................ 491
A. The Desirable Scope of Defamation Law ................................................. 491
1. Defamation Law in Employment ..................................................... 492
2. Defamation Law in Consumer Markets ........................................... 493
3. Political Speech ................................................................................ 494
B. Some Constitutional Ramifications .......................................................... 495
CONCLUSION ......................................................................................................... 496
453
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
THE CASE AGAINST EXPANDING
DEFAMATION LAW
Yonathan A. Arbel & Murat Mungan*
It is considered axiomatic that defamation law protects reputation. This proposition—commonsensical,
pervasive, and influential—is faulty. Underlying this fallacy is the failure to appreciate audience effects:
the interaction between defamation law and members of the audience.
Defamation law seeks to affect the behavior of speakers by making them bear a cost for spreading un-
truthful information. Invariably, however, the law will also affect members of the audience, as statements
made in a highly regulated environment tend to appear more reliable than statements made without
accountability. Strict defamation law would tend to increase the perceived reliability of statements, which
in some cases can have harmful effects on the reputation of the targets of the speech.
This unrecognized complexity of defamation law has the potential to tip the scales in First Amendment
jurisprudence towards greater protection of free speech and free press. Audience effects should also be
considered within the newly announced Restatement project on defamation law. Most urgently, the conse-
quences of audience effects should give pause to the recent calls to expand libel laws to fight fake news by
showing that such laws may well backfire and exaggerate the consequences of falsehoods.
INTRODUCTION
In New York Times Co. v. Sullivan, the Supreme Court famously circum-
scribed the tort of defamation to protect freedom of speech and press.1 Now,
winds blowing from Washington augur that the Times they are a-changin’.2 Re-
cently, Justice Thomas called on the Supreme Court to reconsider its holding
in New York Times, seeing the scope of defamation law as a state issue, rather
than one having a “constitutional status.”3 In the name of fighting “fake news,”
many others are now calling to erode the safeguards set out by this case and its
progeny. President Donald Trump has made defamation law a repeated theme
* Assistant Professor of Law, the University of Alabama School of Law, and Professor of Law, George
Mason University, Antonin Scalia Law School. We are grateful for the insightful comments of Jack Balkin,
Oren Bar-Gill, Omri Ben-Shahar, Mark Brandon, William Brewbaker, Zachary Clopton, Shahar Dillbary,
Heather Elliott, Janet Freilich, John Goldberg, Christine Jolls, Ronald Krotoszynski, Yair Listokin, Irina
Manta, Benjamin McMichael, Michael Pardo, Kish Parella, Steve Shavell, and Rebecca Tushnet. We are also
grateful for comments received in the Yale/Harvard/Stanford Junior Faculty Forum and Midwestern Law &
Economics Conference. The editors of the Alabama Law Review gave many helpful and thoughtful sugges-
tions. Katherine Johnson, Hamilton Millwee, Victoria Moffa, and Brenton Smith provided invaluable re-
search assistance.
1. 376 U.S. 254, 277–83 (1964).
2. BOB DYLAN, THE TIMES THEY ARE A-CHANGIN’ (Columbia Records 1964).
3. McKee v. Cosby, 139 S. Ct. 675, 682 (2019) (Thomas, J., concurring in the denial of certiorari).
454
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
2019] The Case Against Expanding Defamation Law 455
of his public communications, and he promised decisive action: “We are going
to take a strong look at our country’s libel laws, so that when somebody says
something that is false and defamatory about someone, that person will have
meaningful recourse in our courts . . . .”4 At the same time, in January of 2019,
the American Law Institute (ALI) announced the start of a new Restatement
project for defamation law.5 The confluence of political will, support on the
Supreme Court, and the ALI project suggests that, indeed, writers and critics
who prophesize with their pens should keep their eyes wide open, for the Times
they are a-changin’.6
The calls to expand defamation law—by removing safeguards and increas-
ing money damages7—are grounded in the widely held theory that defamation law
protects reputation.8 And because defamation law is believed to protect reputation,
4. Michael M. Grynbaum, Trump Renews Pledge to ‘Take a Strong Look’ at Libel Laws, N.Y. TIMES (Jan.
10, 2018), https://www.nytimes.com/2018/01/10/business/media/trump-libel-laws.html; see also David
Jackson, Donald Trump Maintains Attacks on Bob Woodward, Calls for Changes in Libel Laws, USA TODAY (Sept.
5, 2018, 3:55 PM), https://www.usatoday.com/story/news/politics/2018/09/05/donald-trump-maintains-
attacks-bob-woodward-calls-changes-libel-laws/1199794002/; Donald J. Trump (@realDonaldTrump),
TWITTER (Sept. 5, 2018, 6:33 AM), https://twitter.com/realdonaldtrump/status/1037302649199177728?
lang=en (“Isn’t it a shame that someone can write an article or book, totally make up stories and form a
picture of a person that is literally the exact opposite of the fact, and get away with it without retribution or
cost. Don’t know why Washington politicians don’t change libel laws?”). Other figures have insinuated that
the “fake news” media is so untrustworthy that it must be circumvented. See Sophia Tesfaye, Trump’s Daughter-
in-Law Pushes His Propaganda: Lara Trump Launches “Real News” Show to Praise the President, SALON (Aug. 2, 2017,
5:03 PM), https://www.salon.com/2017/08/02/trumps-daughter-in-law-pushes-his-propaganda-lara-
trump-launches-real-news-show-to-praise-the-president/. Despite Trump’s calls for stronger defamation
laws, he himself has been a defendant in defamation suits. Jessica Levinson, Stormy Daniels and Summer Zervos
Are Using Defamation Laws to Try and Reveal the Truth About Trump, NBC NEWS (Mar. 28, 2018, 2:18 PM),
https://www.nbcnews.com/think/opinion/stormy-daniels-summer-zervos-are-using-defamation-laws-try-
reveal-ncna860851.
5. Restatement of the Law Third, Torts: Defamation and Privacy, ALI, https://www.ali.org/projects/show/
torts-defamation-and-privacy/ (last visited Oct. 8, 2019).
6. DYLAN, supra note 2.
7. See, e.g., Hadas Gold, Donald Trump: We’re Going to ‘Open Up’ Libel Laws, POLITCO (Feb. 26, 2016,
2:31 PM), https://www.politico.com/blogs/on-media/2016/02/donald-trump-libel-laws-219866 (quoting
Trump’s assertion that he is “going to open up our libel laws so when [newspapers] write purposely negative
and horrible and false articles, we can sue them and win lots of money”).
8. Milkovich v. Lorain Journal Co., 497 U.S. 1, 12 (1990) (“Defamation law developed not only as a
means of allowing an individual to vindicate his good name, but also for the purpose of obtaining redress for
harm caused by such statements.”); see also, e.g., PETER N. AMPONSAH, LIBEL LAW, POLITICAL CRITICISM,
AND DEFAMATION OF PUBLIC FIGURES 2 (2004) (“To protect people from injury to their reputation, socie-
ties create laws of defamation to settle issues of truth or falsehood and reputational harm that result from
defamatory statements.”); DAN B. DOBBS ET AL., HORNBOOK ON TORTS § 37.1, at 936 (2d ed. 2015) (“Def-
amation law . . . aims at protecting reputation and good name against false and derogatory communica-
tions.”); RICHARD A. EPSTEIN & CATHERINE M. SHARKEY, CASES AND MATERIALS ON TORTS 1017 (10th
ed. 2012) (describing the common law “premise that an individual’s reputation should be protected against
false words”); WILLIAM K. JONES, INSULT TO INJURY: LIBEL, SLANDER, AND INVASIONS OF PRIVACY 1, 9
(2003) (“Briefly stated, the law of defamation protects a person against falsehoods that expose him to hatred,
contempt, or ridicule or cause him to be shunned by his fellows or that tend to injure him in his trade or
occupation. . . . The objective of the law of defamation is to protect reputations against derogatory false-
hoods.”); W. PAGE KEETON ET AL., PROSSER AND KEETON ON THE LAW OF TORTS § 111, at 771 (5th ed.
1984) [hereinafter PROSSER AND KEETON] (“[D]efamation is an invasion of the interest in reputation and
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
456 ALABAMA LAW REVIEW [Vol. 71:2:453
many think that stricter defamation law means stricter protection of reputa-
tional interests. It is remarkable just how common the reputation-protection
theory is. Whereas lawyers often decry the incoherency and vagaries of the com-
mon law tort of defamation,9 they share what has been described as a “virtually
axiomatic” belief that defamation law shields reputation from harm.10 Courts,
commentators, policy makers, and lay people all seem to share this common
theory.11 Despite decades of jurisprudence and debate, few have contested it,
perhaps because of how enticing it is in its simplicity—this theory that offen-
sive, pejorative, and vituperative comments cause harm to the victim’s standing
in society. To prevent wrongdoers from causing this harm and to compensate
victims, tort law must impose a fine on those who make statements that are
found to be false. On the basis of this theory, courts will let a tort overcome a
constitutional right that is right at the heart of the American ethos.12
This Article exposes the shortcomings of the reputation-protection theory
by demonstrating how it fails to consider audience effects.13 Unlike the harms from
good name.”); LAWRENCE MCNAMARA, REPUTATION AND DEFAMATION 1 (2007) (noting that “[t]he clarity
of the rationale and the extent of agreement surrounding it is quite remarkable”). For further discussion, see
infra Part I.A.
9. EPSTEIN & SHARKEY, supra note 8, at 1017 (“Of all the areas of tort law, defamation is perhaps the
most difficult to organize and to understand.”); PROSSER AND KEETON, supra note 8, § 111, at 771 (“[T]here
is a great deal of the law of defamation which makes no sense. It contains anomalies and absurdities for which
no legal writer ever has had a kind word . . . .”); Randall P. Bezanson, The Libel Tort Today, 45 WASH. & LEE
L. REV. 535, 543 (1988) (“[R]ecovery by any plaintiff is more likely to be the product of chance than of any
systematic pattern reflecting reputational interests.”); Sheldon W. Halpern, Values and Value: An Essay on Libel
Reform, 47 WASH. & LEE L. REV. 227, 230 (1990). For an early discussion of defamation, see Van Vechten
Veeder, The History and Theory of the Law of Defamation, 3 COLUM. L. REV. 546, 546 (1903) (“English law of
defamation . . . is a mass which has grown by aggregation . . . . [P]erhaps no other branch of the law is as
open to criticism for its doubts and difficulties, its meaningless and grotesque anomalies. It is, as a whole,
absurd in theory, and very often mischievous in its practical operation.”).
10. MCNAMARA, supra note 8, at 1. The concept of reputation in legal scholarship is quite thin. See
generally Yonathan A. Arbel, Reputation Failure: The Limits of Market Discipline in Consumer Markets, WAKE
FOREST L. REV. (forthcoming 2019) [hereinafter Arbel, Reputation Failure] (exploring the concept of reputa-
tion and arguing that systems of reputation are highly susceptible to significant bias, given the divergence
between private incentives people have to share reputational information and its status as a public good).
11. See infra Part I.A.
12. See, e.g., Ventura v. Kyle, 63 F. Supp. 3d 1001, 1011 (D. Minn. 2014), rev’d in part, vacated in part, 825
F.3d 876, 885–86 (8th Cir. 2016) (vacating the jury verdict due to a prejudicial error during trial); Blake v.
Ann-Marie Giustibelli, P.A., 182 So. 3d 881, 883–84 (Fla. Dist. Ct. App. 2016) (affirming $350,000 in damages
for online defamatory reviews of attorney services); Anagnost v. Mortg. Specialists, Inc., No. 216-2016-cv-
277, 2017 WL 7693149 (N.H. Super. Ct. Sept. 1, 2017), aff’d, No. 2017-0311, 2018 WL 4940850 (N.H. Sept.
25, 2018) (upholding a $275 million judgment, see Jury Verdict Form, Anagnost, No. 216-2016-cv-277, 2017
WL 7690898, for false accusations of running a drug ring and conspiracy to murder). However, other torts
often yield to constitutional protections. See, e.g., Snyder v. Phelps, 562 U.S. 443, 460–61 (2011) (holding that
the First Amendment bars the father of a deceased soldier from recovering for intentional infliction of emo-
tional distress due to picketing by members of the Westboro Baptist Church); United States v. Eichman, 496
U.S. 310, 319 (1990) (holding that a federal law prohibiting flag burning was unconstitutional and that “[p]un-
ishing desecration of the flag dilutes the very freedom that makes this emblem so revered, and worth rever-
ing”).
13. This statement is less bold when one considers Professor Laura Heymann’s argument that, until
recently, the existence of the audience and its interests were all but neglected in both scholarship and case
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
2019] The Case Against Expanding Defamation Law 457
traditional torts, such as assault and battery, reputational harms are not imme-
diate. Rather, they are mediated by third parties, namely, the audience.14 Repu-
tational harm is the result of the audience believing, at least to some extent, in
a negative statement. As a result, any legal analysis of defamation is incomplete
without considering audience effect. As communication theorists agree, audi-
ence effects can be complex: “The modern view, informed by decades of em-
pirical research, supports an understanding antithetical to the assumption of
direct and uniform effects [of defamatory statements].”15 Yet the standard
model of defamation law only accounts for speakers and victims, relegating the
audience to a secondary and often invisible role.16
This Article offers a framework for defamation law that accounts for audi-
ence effects. Audiences form judgments on the basis of a multitude of factors.
We focus here on one important, but neglected, factor—the legal environment
itself. Borrowing from the rich and well-established signaling theory,17 we con-
sider how defamation law affects audience beliefs. A key finding in this litera-
ture is that expensive signals appear more reliable than cheap signals, “cheap
talk.”18 This makes intuitive sense: puffery and gossip appear less reliable than
a statement by a person testifying under the threat of perjury. Because stricter
defamation law makes inaccurate statements more expensive (by increasing the
likelihood of an adverse judgment), they make the signals appear to be more
reliable. Indeed, this is why the regulation of false advertising is believed to
increase consumer confidence in the marketplace.19 The increased perception
of the reliability of statements is salutary when the statements are indeed truth-
law. See Laura A. Heymann, The Law of Reputation and the Interest of the Audience, 52 B.C. L. REV. 1341, 1341
(2011).
14. Other dignitary torts, such as intentional infliction of emotional harm or infringement of privacy,
are contingent on the reaction of the victim, rather than the audience, and thus can be tortious even when
they involve true statements.
15. Jeremy Cohen & Albert C. Gunther, Libel as Communication Phenomena, 9 COMM. & L. 9, 21 (1987).
16. See, e.g., Heymann, supra note 13, at 1341; Lyrissa Barnett Lidsky, Nobody’s Fools: The Rational Audience
as First Amendment Ideal, 2010 U. ILL. L. REV. 799, 801 (2010) [hereinafter Lidsky, Rational Audience] (“The
Court rarely articulates its assumptions about the presumed audience of core speech, but its assumptions
shape the outcomes of First Amendment cases.”).
17. See generally Rebecca Bliege Bird & Eric Alden Smith, Signaling Theory, Strategic Interaction, and Symbolic
Capital, 46 CURRENT ANTHROPOLOGY 221 (2005); Brian L. Connelly et al., Signaling Theory: A Review and
Assessment, 37 J. MGMT. 39 (2011); Michael Spence, Job Market Signaling, 87 Q.J. ECON. 355 (1973).
18. See, e.g., William Boulding & Amna Kirmani, A Consumer-Side Experimental Examination of Signaling
Theory: Do Consumers Perceive Warranties as Signals of Quality?, 20 J. CONSUMER RES. 111 (1993).
19. Everette MacIntyre, Member, Fed. Trade Comm’n, Address Before the Better Business Division
Miami-Dade County Chamber of Commerce: FTC Promotes Confidence in Advertising (July 18, 1963),
https://www.ftc.gov/system/files/documents/public_statements/683591/19630718_macintyre_ftc_pro-
motes_confidence_in_adertising.pdf. We also find a similar idea in securities law: close regulation of infor-
mation streams would increase investor confidence. See, e.g., Mary Schapiro, Chairman, U.S. Sec. & Exch.
Comm’n, Speech at the SIFMA Annual Conference: The Road to Investor Confidence (Oct. 27, 2009),
https://www.sec.gov/news/speech/2009/spch102709mls.htm (noting that the way to restore investor con-
fidence, post crisis, is through high standards and stricter enforcement of securities laws).
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
458 ALABAMA LAW REVIEW [Vol. 71:2:453
ful. But it can have negative and deleterious effects when the occasional state-
ment proves to be false, as audiences are more likely to believe it to be true than
they would absent strict defamation law.20 Hence, we conclude that strict defama-
tion law may damage reputational interests.
The framework presented here offers another dimension to standard anal-
yses of defamation law.21 In the standard bilateral-tort model, courts and com-
mentators see expansions to defamation law as involving a simple balance be-
tween better protection of the victim’s reputation and the chilling effect of such
laws on speakers.22 The framework here adds another player—the audience.
The chilling effect in the bilateral model assumes that would-be speakers will
change their behavior in response to stricter defamation laws. In the trilateral
model, stricter defamation laws also affect the audience. In particular, we argue
that considering the impact of defamatory statements on audiences produces a
“seesaw dynamic,” where strict defamation laws chill false statements but in-
crease the reliability and thus the harm of surviving false statements.23 Using
the trilateral model, courts should engage in richer balancing of the effects of
defamation law, which can sometimes lead to very different conclusions than
the standard analysis.
20. A poetic illustration of some of the audience effects is encapsulated in writer Bertolt Brecht’s
satirical poem, The Burning of the Books:
When the Regime commanded that books with harmful knowledge
Should be publicly burned and on all sides
Oxen were forced to drag cartloads of books
To the bonfires, a banished
Writer, one of the best, scanning the list of the
Burned, was shocked to find that his
Books had been passed over. He rushed to his desk
On wings of wrath, and wrote a letter to those in power
Burn me! he wrote with flying pen, burn me! Haven’t my books
Always reported the truth? And here you are
Treating me like a liar! I command you:
Burn me!
BERTOLT BRECHT, POEMS 1913-1956, at 294 (1998).
21. Our paper was developed contemporaneously and independently of a similar project pursued by
Professors Ariel Porat and Daniel Hemel. Both projects conclude that audience effects undermine the stand-
ard narrative of defamation law, although we differ in focus and in some of our conclusions. See Daniel Jacob
Hemel & Ariel Porat, Free Speech and Cheap Talk, 11 J. LEGAL ANALYSIS 46 (2019).
22. See, e.g., Petro-Lubricant Testing Labs., Inc. v. Adelman, 184 A.3d 457, 461 (N.J. 2018) (“Defama-
tion law balances two competing interests—an individual’s right to protect his reputation from unjustified
and false aspersions and our citizens’ right to free expression and robust debate in our democratic society.”).
23. See infra Part II.C.
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
2019] The Case Against Expanding Defamation Law 459
Audience effects, of the kind we identify here, do not depend on audiences
being rational, sophisticated, or well versed in defamation law. Just like the well-
accepted chilling effect, the key dynamics we identify here accommodate dif-
ferent views regarding these issues. In fact, under some conditions, greater au-
dience ignorance and gullibility will bolster the harmful effect of expanding def-
amation law.24 What we find implausible is the idea inherent in the standard
model that defamation law affects speakers and their subjects but that somehow
audiences are completely insulated from these effects.25 As noted, it is well rec-
ognized in other domains that close regulation of the information environment
would lead to increased trust by members of the audience.26 Our analysis here
frames the key dynamics and invites judges, policy makers, and scholars to ex-
plicitly consider these effects.
Our policy prescriptions call for a more nuanced and careful approach to
the evolution of defamation law. There will be contexts wherein expansions of
defamation law can be shown to be helpful, but the law’s efficacy is neither self-
evident nor easily supported. Lawmakers, courts, and commentators contem-
plating a reform or a new application of defamation law—such as President
Trump’s proposal—should consider the trade-off between defamation law’s
potential benefits and its countervailing effects. When considering the expan-
sion of defamation law, policy makers should ask: what is the harm prevented
by having fewer falsehoods, and how does it compare to the harm caused by
falsehoods becoming more believable? These regulators should only expand
liability if they judge that the harm prevented by the expansion is greater than
the harm it creates. Indeed, this question is not always easy to answer, and there
will often be some factual uncertainty about this trade-off. In such cases, we
proffer a default position with a venerable history in the courts: err on the side
of free speech. If one cannot prove that a new application of defamation law
would result in a net benefit to reputation, then defamation’s ambit should not
be expanded.
24. Audiences may easily underestimate or overestimate the strictness of defamation law. The effects
of expanding defamation law depend on the audience members’ starting positions and how likely and quick
they are to learn about changes to the law, among other considerations. See generally Christine Jolls, Debiasing
Through Law and the First Amendment, 67 STAN. L. REV 1411 (2015) (listing empirical evidence of how the law
might correct consumer misperceptions). But see Karen Russo France & Paula Fitzgerald Bone, Policy Makers’
Paradigms and Evidence from Consumer Interpretations of Dietary Supplement Labels, 39 J. CONSUMER AFF. 27 (2005)
(finding that it is difficult to correct the misperception that the safety of certain supplements is regulated).
Psychological biases and heuristics add another dimension of complexity, as audiences may react strongly to
salience and repetition (even of suspected lies) or may be overly cynical as a defense mechanism. These
considerations, alongside many other practical, legal, and political considerations and the findings of this
paper, are relevant for the determination of the scope of defamation law. This is why it is important to
abandon the simplistic bilateral model of reputation protection.
25. Perhaps in some contexts, speakers are better informed about the law than audiences. While rele-
vant, what matters most for the analysis here is that audiences have some very general (even if mistaken) idea
about the severity of the law, which is plausible given the common media coverage of libel lawsuits.
26. See supra note 19 and accompanying text.
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
460 ALABAMA LAW REVIEW [Vol. 71:2:453
We show in three Parts that the case for expanding defamation law is an
uneasy one. In Part I, we provide general background on defamation law, First
Amendment jurisprudence, and the protection of reputation. A central obser-
vation here is that the making of defamatory statements is already permissible
in many important areas of social life, and commentators generally accept such
carve outs without much alarm, presumably because audiences take assertions
in these spheres more skeptically. This observation grounds our discussion and
demonstrates the practicality and sensibility of limiting defamation law. In Part
II, we develop our main argument using concrete examples. While we base our
general analysis on a formal model,27 here we use a simple, stylistic, and easy-
to-follow example. We then consider a variety of limitations, qualifications, and
considerations. We pay particular attention to issues such as audience sophisti-
cation, behavioral biases, reliability and accessibility of courts, and the stakes of
the allegations. With a few exceptions, the dynamics we identify in Part II turn
out to be general, although their strength is context-specific.
In Part III, we start to explore some of the legal implications of our analysis.
We examine the proper ambit of defamation law in diverse contexts, such as
employment, consumer goods, and political speech. We conclude that, first,
courts should reexamine traditional balances in First Amendment jurisprudence
and tilt the balance further in favor of free speech. Historically, when courts
decided defamation cases, they labored under the assumption that defamation
law definitively protects reputation. The upshot of our analysis is that defama-
tion law is less protective than traditionally believed, and so the weight given by
defamation law to the goal of protecting reputation should be considerably
lighter. The recent announcement of a new Restatement project presents a rare
opportunity to engage in such reflection.28
The second and perhaps the most urgent consequence of our analysis is
that the war on fake news should not result in an expansion of libel laws. This
consequence is significant, because our analysis takes the concerns of fake news
seriously. Based on this premise, the analysis shows that expanding defamation
law can easily backfire, resulting in the unintended consequence of aggravating
the harm of fake publications by lulling the populace into a sense of assurance
that if it’s in the news, it must be true.
27. See Yonathan A. Arbel & Murat Mungan, Regulating Speech with Bayesian Audiences (U. of Ala. Sch. of
Law Legal Studies, Research Paper No. 3452662, 2019), https://papers.ssrn.com/sol3/abstract=3452662.
28. See Restatement of Law Third, Torts: Defamation and Privacy, supra note 5.
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
2019] The Case Against Expanding Defamation Law 461
I. DEFAMATION LAW & REPUTATION
This Part explores modern defamation law, as it formed after the turbu-
lence of the 1960s with New York Times Co. v. Sullivan.29 This Part demonstrates
how the belief in defamation law’s reputation-protecting powers has shaped its
doctrine, highlights how courts have relegated the role of audiences to second-
ary consideration, and explores some of the exceptions to defamation law.
These exceptions, it is argued, are domains where audiences tend to be more
critical of information.
A. Goals & Functions of Defamation Law
Under defamation law, a plaintiff can recover pecuniary damages from any
person who makes a public, false defamatory statement against her, unless spe-
cial conditions or privileges exist.30 We defer discussion of doctrine until after
we have examined a pressing preliminary question of why the law is needed to
regulate such statements in the first place. In the United States, free speech is a
core value, ethos, and right; the commitment to protecting and promoting free
speech runs deep. Indeed, most forms of speech go unregulated, even if the
content is inflammatory, inane, provocative, flippant, or frivolous.31 Why, then,
regulate defamatory speech?
Defamation law exists to protect reputation. This, at least, is the common
view, one that is so broadly held and so commonsensical that it is in the small
province of statements to which lawyers do not attach a footnote. Leading com-
mentators call this idea axiomatic and universal.32 Courts frequently recite the
idea that “[i]n American law, defamation is . . . about protecting a good reputa-
tion honestly earned.”33 Several casebooks matter-of-factly state that “[d]efa-
mation protects an individual’s interest in reputation,” omitting—for reasons
of blatant obviousness—any support or authority.34 Note that this justification
is distinct from the related, but often confused, goal of protecting emotions or
29. 376 U.S. 254 (1964).
30. See PROSSER AND KEETON, supra note 8, § 111, at 771.
31. See, e.g., Snyder v. Phelps, 562 U.S. 443, 458 (2011) (allowing Westboro Baptist Church’s picketing
of a soldier’s funeral); Reno v. ACLU, 521 U.S. 844 (1997) (discussing “indecent” content on the Internet);
United States v. Eichman, 496 U.S. 310, 319 (1990) (regarding flag burning).
32. See supra note 10 and accompanying text; see also Spencer v. Kemna, 523 U.S. 1, 24 n.5 (1998)
(Stevens, J., dissenting) (“[V]indicating one’s reputation is the main interest at stake in a defamation
case . . . .”); Albright v. Oliver, 510 U.S. 266, 283 (1994) (Kennedy, J., concurring) (“[T]he interests granted
historical protection by the common law of torts [include] freedom from defamation . . . .”); David J. Ache-
son & Ansgar Wohlschlegel, The Economics of Weaponized Defamation Lawsuits, 47 SW. L. REV. 335, 335 (2018)
(“The law of defamation is the principal legal mechanism in both the United States and England for protect-
ing the interest in reputation.”).
33. Bustos v. A&E Television Networks, 646 F.3d 762, 764 (10th Cir. 2011).
34. RUSSEL L. WEAVER ET AL., TORTS: CASES, PROBLEMS, AND EXERCISES 803 (2d ed. 2005).
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
462 ALABAMA LAW REVIEW [Vol. 71:2:453
privacy. These goals are served by other tort doctrines, such as intentional in-
fliction of harm and privacy infringement, that are not conditioned on the men-
dacity of the allegation.35 The Supreme Court has emphasized that emotional
harm alone cannot support a defamation lawsuit but that a defamation lawsuit
must rest on harm to reputation, which is, after all, its raison d’etre.36
As to why reputation should be protected, there is an unsurprising differ-
ence of opinion. Rights-based accounts tend to view reputation as being part
of a person’s property, dignity, or honor.37 In such views, reputation is some-
thing one has, a view which is quite dominant in the scholarship.38 For such
accounts, the wrongfulness of defamation consists in the taking and the viola-
tion of a right.39 A more social view sees defamation as a replacement of the
old customs of honor duels and blood feuds.40 By channeling certain behaviors
to the court system, defamation law is believed to serve a civilizing function—
so much so that some think that “a civilized society cannot refuse to protect
reputation.”41
Economic theories of reputation shift the point of view from the subject
to the public. Reputation consists of an ununified mass of opinions others have
of us.42 Such opinions have predictive values. They help members of the com-
munity—or market participants—judge the “affinity” of a particular partner, be
35. PROSSER AND KEETON, supra note 8, § 111, at 771 (“Defamation is not concerned with the plain-
tiff’s own humiliation, wrath or sorrow . . . .”). In Gertz v. Robert Welch, Inc., the Court took a more inclusive
approach. 418 U.S. 323, 349–50 (1974) (“We need not define ‘actual injury,’ . . . . Suffice it to say that actual
injury is not limited to out-of-pocket loss. Indeed, the more customary types of actual harm inflicted by
defamatory falsehood include impairment of reputation and standing in the community, personal humiliation,
and mental anguish and suffering.”).
36. See Hustler Magazine, Inc. v. Falwell, 485 U.S. 46, 52–57 (1988).
37. See Robert C. Post, The Social Foundations of Defamation Law: Reputation and the Constitution, 74 CALIF.
L. REV. 691, 693–719 (1986); see also Richard A. Epstein, Was New York Times v. Sullivan Wrong?, 53 U. CHI.
L. REV. 782, 800–01 (1986).
38. See, e.g., David S. Ardia, Reputation in a Networked World: Revisiting the Social Foundations of Defamation
Law, 45 HARV. CIV. RTS.-CIV. LIBERTIES L. REV. 261, 290 (2010) (“The . . . most dominant[] conception of
reputation embodied in American defamation law is that of reputation as property.”); Post, supra note 37, at
730 (discussing defamation “within the framework of reputation as property”).
39. See generally Dun & Bradstreet, Inc. v. Greenmoss Builders, Inc., 472 U.S. 749 (1985) (Powell, J.,
plurality opinion). In common law, punitive damages are conditional on malice but in the sense of a wrongful
attitude rather than in the Sullivan sense of a disregard for the truth. See Sheldon W. Halpern, Of Libel, Language,
and Law: New York Times v. Sullivan at Twenty-Five, 68 N.C. L. REV. 273, 278 (1990).
40. See PROSSER AND KEETON, supra note 8, § 111, at 772. It is puzzling what social function duels
had; after all, the ability to win a duel is not correlated with the veracity of the offending statement. See generally
Ben Merriman, Duels in the European Novel: Honor, Reputation, and the Limits of a Bourgeois Form, 9 CULTURAL
SOC. 203 (2015) (rooting duels in a system of honor, dispute resolution, and diffusion of conflict).
41. David A. Anderson, Is Libel Law Worth Reforming?, 140 U. PA. L. REV. 487, 490 (1991).
42. RICHARD A. POSNER, THE ECONOMICS OF JUSTICE 272 (1981) (“A person’s reputation is other
people’s valuation of him as a trading, social, marital, or other kind of partner. It is an asset of potentially
great value which can be damaged both by false and by true defamation.”); Christian Hahn et al., Social Repu-
tation: A Mechanism for Flexible Self-Regulation of Multiagent Systems, 10 J. ARTIFICIAL SOCIETIES & SOC.
SIMULATION ¶ 3.1 (Jan. 31, 2007), http://jasss.soc.surrey.ac.uk/10/1/2/2.pdf (treating reputation as a form
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
2019] The Case Against Expanding Defamation Law 463
it in a social, romantic, or transactional setting.43 How trustworthy is a potential
business partner? How good is a specific product? How loyal is a potential date?
In economic parlance, reputation is a signal about past behavior that is predic-
tive of future behavior;44 it is a “shadow of the future”45 that helps economize
search costs.46
Defamation law is thought to promote the protection of reputation in three
ways. One idea, perhaps the most prominent, is deterrence. By imposing a sanc-
tion on false allegations, defamation law disincentivizes such wrongful behavior
ex ante.47 Prospective liars would worry that lying exposes them to civil liability
and so would likely refrain from falsely defaming others.48 A second way is
through redress; by offering compensation to victims of defamation, defama-
tion law protects their reputation against wrongful attacks.49 The Supreme
Court has declared both of these goals to be important.50 A third way, although
less commonly articulated, is vindication.51 The idea of vindication focuses on
the legal process. By being able to bring a lawsuit, a victim gains access to a
procedure that allows for a public determination of truth—either through the
outcome of the case or through information revealed in the litigation process.52
Thus, victims can clear their names against nefarious allegations and have court
judgments as records of their innocence.
of “symbolic capital”); Roy Shapira, Reputation Through Litigation: How the Legal System Shapes Behavior by Producing
Information, 91 WASH. L. REV. 1193, 1203–04 (2016).
43. POSNER, supra note 42, at 272.
44. Ardia, supra note 38, at 264 (describing reputation as “a heuristic for predicting the behavior of
others”).
45. Id. (quoting ROBERT AXELROD, THE EVOLUTION OF COOPERATION 126 (1984)).
46. POSNER, supra note 42, at 287.
47. AMPONSAH, supra note 8, at 2; Stanley Ingber, Defamation: A Conflict Between Reason and Decency, 65
VA. L. REV. 785, 792 (1979) (“[A] successful defamation suit can act as a public rebuke and as an economic
penalty for the publishers of defamatory statements. Anticipation of such suits presumably will deter individ-
uals from making false statements that cause ridicule and loss of reputation.”). See LAURENCE H. ELDREDGE,
THE LAW OF DEFAMATION § 3, at 6 (1978); see also RODNEY A. SMOLLA, LAW OF DEFAMATION § 1:26 (2d
ed. 2019) (“The law of defamation serves an important social interest as a deterrent on the publication of
false and injurious speech.”).
48. Gertz v. Robert Welch, Inc., 418 U.S. 323, 350 (1974) (describing punitive damages in defamation
suits as “private fines levied by civil juries to punish reprehensible conduct and to deter its future occur-
rence”).
49. AMPONSAH, supra note 8, at 2.
50. Milkovich v. Lorain Journal Co., 497 U.S. 1, 12 (1990).
51. Id. at 23 (“[A]n action for damages is the only hope for vindication or redress the law gives to a
man whose reputation has been falsely dishonored.” (quoting Rosenblatt v. Baer, 383 U.S. 75, 93 (1966)
(Stewart, J., concurring))).
52. Shapira, supra note 42, at 1211–23.
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
464 ALABAMA LAW REVIEW [Vol. 71:2:453
By design, defamation law falls short of protecting these goals—it does not
deter all falsehoods, compensate all victims, or vindicate all claims. There are
both institutional and substantive reasons for these limitations. On the institu-
tional side, it is widely understood that courts cannot always identify falsehoods,
the judicial process is expensive and affords limited access, offenders are some-
times judgment-proof, and courts have limited jurisdiction.53 On the substan-
tive side, First Amendment considerations,54 concerns about chilling speech,55
and aversion to the idea that state courts should declare what is true and what
is false56 have led courts to approach defamation lawsuits with a willingness to
err on the side of unfettered speech.57 This is summed up by the idea that def-
amation law must give “breathing space” to free speech58 and that some degree
of abuse is to be tolerated in a free society.59 As such, courts have shown reluc-
tance to award injunctions.60 They have also required heightened standards of
proof,61 and in some cases, they have even struck down prohibitions on false
speech motivated by ill-intent.62
In closing, it should be noted that the idea that defamation protects repu-
tation is not some impractical abstraction; it has significant constitutional and
legal ramifications. Starting in the 1960s—surprisingly late—courts started to
53. See generally Lyrissa Barnett Lidsky, Silencing John Doe: Defamation & Discourse in Cyberspace, 49 DUKE
L.J. 855 (2000).
54. St. Amant v. Thompson, 390 U.S. 727, 732 (1968) (“[T]o insure the ascertainment and publication
of the truth about public affairs, it is essential that the First Amendment protect some erroneous publications
as well as true ones.”).
55. See generally Gary L. Lee, Comment, Strict Liability Versus Negligence: An Economic Analysis of the Law
of Libel, 1981 BYU L. REV. 398, 400 (“[F]alse defamatory statements are properly viewed as an unavoidable
cost of publishing true defamatory statements.”).
56. Gertz v. Robert Welch, Inc., 418 U.S. 323, 349 (1974) (warning that “juries [might] punish unpop-
ular opinion[s]”).
57. Id. at 341 (“The First Amendment requires that we protect some falsehood in order to protect
speech that matters.”).
58. Hustler Magazine, Inc. v. Falwell, 485 U.S. 46, 56 (1988).
59. See Gertz, 418 U.S. at 341; see also Snyder v. Phelps, 562 U.S. 443, 458 (2011) (“If there is a bedrock
principle underlying the First Amendment, it is that the government may not prohibit the expression of an
idea simply because society finds the idea itself offensive or disagreeable.” (quoting Texas v. Johnson, 491
U.S. 397, 414 (1989))); James Madison, Report on the Virginia Resolutions, reprinted in 4 DEBATES IN THE
SEVERAL STATE CONVENTIONS ON THE ADOPTION OF THE FEDERAL CONSTITUTION 546, 571 (Jonathan
Elliot ed., Philadelphia, J.B. Lippincott Co. 2d ed. 1891) (1800) (“Some degree of abuse is inseparable from
the proper use of every thing; and in no instance is this more true than in that of the press.”).
60. See, e.g., Kramer v. Thompson, 947 F.2d 666, 677 (3d Cir. 1991) (“[T]he maxim that equity will not
enjoin a libel has enjoyed nearly two centuries of widespread acceptance at common law.”); Konigsberg v.
Time, Inc., 288 F. Supp. 989, 989 (S.D.N.Y. 1968) (“A court of equity will not, except in special circum-
stances, issue an injunctive order restraining libel or slander or otherwise restricting free speech. To enjoin
any publication, no matter how libelous, would be repugnant to the First Amendment to the Constitution,
and to historic principles of equity.” (citations omitted)).
61. Milkovich v. Lorain Journal Co., 497 U.S. 1, 15 (1990) (“The Court has also determined that both
for public officials and public figures, a showing of New York Times malice is subject to a clear and convincing
standard of proof.”).
62. Illinois ex rel. Madigan v. Telemarketing Assocs., Inc., 538 U.S. 600, 612–16 (2003).
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
2019] The Case Against Expanding Defamation Law 465
recognize the tension between defamation law and freedom of expression.63 In
the seminal case New York Times Co. v. Sullivan, the Supreme Court held that
protection of reputation should sometimes cede to First Amendment consid-
erations and that a defamation lawsuit by public officials cannot proceed with-
out a showing of actual malice.64 This holding cemented the idea that protecting
reputation and freedom of expression requires balancing antagonistic consider-
ations.65 First Amendment considerations could be overcome, even if narrowly,
when limitations on speech are required to protect reputation. Epstein summa-
rized this idea: “Everyone agrees that the central task of the modern law of
defamation is to reconcile the interest in reputation with that in freedom of
speech.”66
B. Defamation Law: Doctrine, Nature of Reputation, and the Audience
Defamation has been on the books for a little longer than there have been
books. First came the moral exhortations. Ancient Sumerian cuneiform tablets
admonish him who would associate with a slanderer.67 The Bible reproves
“speaking guile”68 and “spread[ing] a false report.”69 At some unknown point
in history, defamation transitioned from the improper to the illegal. One early
example comes from the Frankish Lex Salica (dated to around 500 C.E.) which
imposes a three-shilling penalty on him who calls another a “wolf” or a “hare,”
63. See generally Lillian R. BeVier, The First Amendment and Political Speech: An Inquiry into the Substance and
Limits of Principle, 30 STAN. L. REV. 299, 348 (1978); Cynthia L. Estlund, Speech on Matters of Public Concern: The
Perils of an Emerging First Amendment Category, 59 GEO. WASH. L. REV. 1, 10 (1990); id. at 12 (“The First Amend-
ment was a late entrant into the fields of public employee speech and defamation law and has never held full
sway within the two areas.”). On this tension, see AMPONSAH, supra note 8, at 2; Epstein, supra note 37, at
783.
64. New York Times Co. v. Sullivan, 376 U.S. 254, 283 (1964). Confusingly, malice refers not to
attitude towards the subject but rather to the making of a statement while knowing its falsity or with reckless
disregard for its veracity. See Halpern, supra note 39, at 278.
65. See ROBERT D. SACK, SACK ON DEFAMATION: LIBEL, SLANDER, AND RELATED PROBLEMS § 1:1
(3d ed. 2009) (“Even though reputation may be seriously injured by defamation . . . courts concluded that on
balance the damage [to free speech due to chilling effects was] too great to permit the defamed person to
recover.”); J. M. Balkin, Some Realism About Pluralism: Legal Realist Approaches to the First Amendment, 1990 DUKE
L.J. 375, 403–04; Ingber, supra note 47, at 789 (“The tort interest of protecting the individual—the interest of
decency—may circumscribe the ambit of free debate. How should this conflict be resolved?” (footnote omit-
ted)); Gerald R. Smith, Of Malice and Men: The Law of Defamation, 27 VAL. U. L. REV. 39, 40 (1992) (“The Court
has struggled to find a reasonable balance between protecting reputation and protecting free speech by fash-
ioning rules of general applicability in order to provide certainty and predictability and to avoid chilling free
speech.”).
66. EPSTEIN & SHARKEY, supra note 8, at 1022.
67. The Instructions of Shuruppag: Translation, ELEC. TEXT CORPUS OF SUMERIAN LITERATURE §§ 65–
66, http://etcsl.orinst.ox.ac.uk/section5/tr561.htm (last visited Oct. 10, 2019) (“The eyes of the slanderer
always move around as shiftily as a spindle. You should never remain in his presence . . . .”).
68. Psalms 34:13.
69. Exodus 23:1.
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
466 ALABAMA LAW REVIEW [Vol. 71:2:453
or a forty-five shilling penalty if one calls a woman a “harlot.”70 In common
law, defamation made its debut in the latter half of the sixteenth century.71 In
England, two different doctrines coevolved: slander (spoken defamation) and
libel (written defamation).72 Each of these doctrines evolved within its own ven-
ues—ecclesiastical or common courts—and consequently developed different
rules, remedies, and procedures.73 For a period of time, defamation coupled
with sedition to form the much-maligned offense of “seditious libel” (speech
meant to criticize the government).74 Only in modernity did the common law
fuse the disparate torts of libel and slander into what is now known as defama-
tion, although remnants of this distinction still persist.75
The modern doctrine of defamation has a misleadingly simple structure. To
prevail, the plaintiff must show that the defendant has made (i) a public state-
ment76 that is (ii) false and (iii) may diminish the plaintiff’s reputation.77 In prac-
tice, applying these tests proves Herculean, with little in the way of an organiz-
ing theory.78 Here, we explicate some aspects of these doctrinal requirements,
as their application is based on some fundamental assumptions about the nature
of defamation, reputation, and audiences.
For a statement to be defamatory, it must be “published” in the very ex-
pansive sense of being communicated to any other person besides the subject.79
This distinction separates defamation law from emotional or dignitary harms
because those can arise even in private communications, as is the case with
racist or vituperative comments. By imposing a publicity requirement, courts
have taken the position that defamation is not about protecting victims from
emotional harm resulting from offensive speech per se, but rather defamation
70. PACTUS LEGIS SALICAE, tit. XXX, reprinted in KATHERINE FISCHER DREW, THE LAWS OF THE
SALIAN FRANKS 59, 94 (1991). For the history of the Lex Salica, see generally DREW, supra.
71. ELDREDGE, supra note 47, § 3, at 5.
72. EPSTEIN & SHARKEY, supra note 8, at 1051–53.
73. On the origins of defamation in English law, see generally PAUL MITCHELL, THE MAKING OF
THE MODERN LAW OF DEFAMATION (2005); Veeder, supra note 9.
74. WEAVER ET AL., supra note 34, at 803.
75. The law once distinguished between oral defamation (slander) and written defamation (libel) and
afforded stronger protection from the latter. This led to absurd outcomes, with a public address to 3,000
people considered to be less harmful than a private written letter. See ELDREDGE, supra note 47, § 12, at 77.
This distinction still reverberates in modern law, and California, for example, still has different rules for libel
and slander. Compare CAL. CIV. CODE § 45 (West 2018) (defining libel), with CAL. CIV. CODE § 46 (defining
slander).
76. JONES, supra note 8, at 19 (“[I]f the denunciation is not communicated to third persons, it is not
an actionable defamation.”).
77. See RESTATEMENT (SECOND) OF TORTS §§ 558–59 (AM. LAW INST. 1977).
78. See Nat Stern, The Certainty Principle ss Justification for the Group Defamation Rule, 40 ARIZ. ST. L.J. 951,
970 n.114 (2008) (collecting commentary condemning defamation law as “confusing and even incoherent”).
79. See RESTATEMENT (SECOND) OF TORTS §§ 558–59; see also Lambert v. Whiting Turner Contractor,
No. 1:15-cv-958-GBL-MSN, 2016 WL 2946176, at *6 (E.D. Va. May 19, 2016) (“[A] defamatory statement
must be ‘communicated to a third party “so as to be heard and understood by such person.”’” (quoting Katz
v. Odin, Feldman & Pittleman, P.C., 332 F. Supp. 2d 909, 915 (E.D. Va. 2004))).
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
2019] The Case Against Expanding Defamation Law 467
is something that is inextricably tethered to the existence of an audience, i.e.,
harm to reputation. But this observation should not be overstated; courts do
not place any quantitative requirements on what counts as “public.”80 A state-
ment made to even one other person could be considered defamatory, although
the damages will be conditioned in part on the scope of audience exposure to
the defamatory statement.81
The statement, or at least its “gist,” must be false.82 Historically, the veracity
of statements was a defense in common law.83 It was generally presumed that
statements were false unless the defendant could prove otherwise.84
In 1964, the Court decided the landmark case of New York Times & Co. v.
Sullivan.85 This case famously involved an advertisement of allegations made
against the police department of Montgomery, Alabama, for abusing their pow-
ers to suppress the Civil Rights movement.86 The Court held that in defamation
lawsuits by public officials, the plaintiff must show actual malice in order to
prevail.87 This led to the “constitutionalization” of defamation law and to the
broader protection of speech, even when the subject is not strictly a public of-
ficial.88 For example, the assertion that a restaurant fails to meet sanitation
standards may involve a broader interest in public health, and as such, the
speech may be protected.89 Another large effect is that today, in contrast to the
past, truth is treated by most courts as an absolute bar to liability.90
An important qualification on the meaning of truth and falsity is in order.
Recall Justice Jackson’s famous insight that “[courts] are not final because [they]
are infallible, but [they] are infallible only because [they] are final.”91 In a like
80. Huegerich v. IBP, Inc., 547 N.W.2d 216, 221 (Iowa 1996) (“Publication . . . simply means a com-
munication of statements to one or more third persons.”); RESTATEMENT (SECOND) OF TORTS § 577 cmt.
b.
81. See RESTATEMENT (SECOND) OF TORTS § 621 (limiting damages to compensation for “the proved,
actual harm”).
82. Id. §§ 558–59.
83. Truth is mostly an “absolute” defense, but a sizable minority of states still allow recovery even for
truthful defamatory statements if they were made with bad intentions. See, e.g., FLA. CONST. art. 1, § 4. (“If
the matter charged as defamatory is true and was published with good motives, the party shall be acquit-
ted . . . .”); DEL. CODE ANN. tit. 10, § 3919 (2017) (noting that truth could be a defense in libel only if the
information was “published properly for public information, and with no malicious . . . motives”).
84. PROSSER AND KEETON, supra note 8, § 116, at 84 (noting that it is “[o]ut of a tender regard for
reputations [that] the law presumes in the first instance that all defamation is false”).
85. 376 U.S. 254 (1964).
86. Id. at 257–58.
87. Id. at 283.
88. See, e.g., Gertz v. Robert Welch, Inc., 418 U.S. 323, 345–50 (1974).
89. See, e.g., Journal-Gazette Co. v. Bandido’s, Inc., 712 N.E.2d 446 (Ind. 1999).
90. See SACK, supra note 65, § 3.3.2, at 3–7. But see Noonan v. Staples, Inc., 707 F. Supp. 2d 85, 90–92
(D. Mass. 2010) (holding that a Massachusetts firm that sent its employees a truthful but embarrassing recount
of the reasons why the plaintiff was fired may be held liable under Massachusetts law and thus allowing
recovery for malicious truthful defamatory speech).
91. Brown v. Allen, 344 U.S. 443, 540 (1953) (Jackson, J., concurring).
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
468 ALABAMA LAW REVIEW [Vol. 71:2:453
manner, most courts treat their finding of falsity with humility, recognizing that
their findings constitute a “legal” truth rather than an ontological one.92 The
determinations in defamation lawsuits are not the result of exhaustive investi-
gative work by the judge but rather the determination of whether the evidence
presented supports the allegation.
The third and last requirement is for the speech to be defamatory—i.e.,
capable of harming one’s reputation, esteem, good name, or standing in soci-
ety.93 Here, courts do not normally articulate a clear understanding of reputa-
tion. Still, it seems like one reducible part of reputation is that it consists of
beliefs: A person held in high esteem, good standing, or well repute is believed to
be of high prestige, high trust, or high value.94 Understood this way, it is easier
to see how defamatory speech harms reputation—the allegation leads people
to negatively update their opinions so that they will be less willing to trade,
socialize, or partner with its subject.
This understanding of reputation seems to suggest a concern with a per-
son’s actual standing in the community. Courts, however, have treated reputa-
tion as a normative concept, focusing on an imagined, normative audience.95 Ac-
cording to various rulings in this area, harm to reputation requires showing only
that the statement “would tend to hold the plaintiff up to scorn, hatred, ridicule
or contempt.”96 That is, there is no need to show that any member of the audi-
ence actually changed her mind about the plaintiff in response to the allegations.
Moreover, courts sometimes limit attention only to “the minds of any consid-
erable and respectable segment in the community.”97 Hence, not even all segments
of the imagined public are relevant, only those that reason as a “respectable”
92. See also SACK, supra note 65, § 3:12 (noting the tension and suggesting truth in this context would
be understood narrowly, as the product of a legal process rather than an ontological truth).
93. See RESTATEMENT (SECOND) OF TORTS § 559 (AM. LAW INST. 1977) (“A communication is de-
famatory if it tends so to harm the reputation of another as to lower him in the estimation of the community
or to deter third persons from associating or dealing with him.”).
94. See sources cited supra note 42.
95. Because the effect on the actual audience figures into the calculation of damages, the determination
of defamatory nature seems to be inconsequential. If no one found the statement to be defamatory, the scope
of damages would be zero. Speculating on why this filter is even used, two possible reasons emerge. First,
courts may use it to economize on litigation costs (if damages are likely to be zero, dismissing the case for
failing to meet the normative standard avoids the need for a trial). Alternatively, courts recognize that damages
are speculative and juries may compensate individuals even when they did not suffer any harm. Such a pre-
liminary screen thus protects against such outcomes, as the trier of fact will not be permitted to consider
cases that fail to meet the normative threshold. See Note, The Community Segment in Defamation Actions: A Dis-
senting Essay, 58 YALE L.J. 1387, 1388 (1949) (arguing that “[t]he emphasis is usually on normalcy: the eccentric
or ‘wrong-thinking’ segments [of the audience] albeit of considerable size, are disregarded”).
96. See Phelan v. May Dep’t Stores Co., 819 N.E.2d 550, 553 (Mass. 2004) (emphasis added) (quot-
ing Stone v. Essex Cty. Newspapers, Inc., 330 N.E.2d 161, 165 (Mass. 1975)).
97. Id. (emphasis added) (quoting Stone, 330 N.E.2d at 165); SIR P.H. WINFIELD, A TEXT-BOOK OF
THE LAW OF TORT § 72, at 242 (5th ed. 1950) (stating that “[d]efamation [consists of] statement[s] which
tend[] to lower a person in the estimation of right-thinking members of society generally; or which tends to
make them shun or avoid that person”).
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
2019] The Case Against Expanding Defamation Law 469
person, which supposedly means that they are rational. Note, also, that the
“considerable . . . segment” requirement here is different from the one under
the publication requirement, which only requires a single third party.98 For an
example of the operation of this rule, a demeaning fake interview in Hustler
magazine with Andrea Dworkin was held not libelous.99 The ruling was not the
result of testimony from any reader of Hustler but rather the result of the court’s
assessment that the interview’s ludicrous nature made it unbelievable.100 The
Ninth Circuit expressed confidence “that the outrageous and the outlandish will
be recognized for what they are,” without consulting any actual reader of the
magazine.101 If the court believes the statement is believable, it will not help the
defendant to bring contrary evidence; such questions, courts hold, might miti-
gate damages but will not undermine the existence of defamation.102
C. Where Defamation Law Ends
Our analysis here demonstrates the complexity involved in trying to protect
reputation through stricter defamation law. This complexity belies some estab-
lished positions on defamation law, which tend to support an absolute relation-
ship between defamation law and reputation protection. For example, Richard
Epstein warns that failing to protect against defamation would result in “a world
with too much defamation, too much misinformation—in a word, too much
public fraud.”103 Similarly, Justice White warned in Gertz v. Robert Welch, Inc. that
limiting defamation law would “frustrate th[e] search [for truth]” and contribute
to “assaults on individuality and personal dignity.”104
Fortunately, such concerns are assuaged by the fact that defamatory speech
is already unregulated in broad domains, such as statements of opinions; de-
famatory statements regarding public officials; and group-based slander, or the
online hosting of defamatory statements made by third parties.105 In addition
to the legal exceptions to defamation law, there are important practical ones.
Litigation costs, judicial inaccuracy, jurisdictional limits on foreign speech, and
98. See Phelan, 819 N.E.2d at 553 (quoting Stone, 330 N.E.2d at 165).
99. Dworkin v. Hustler Magazine, Inc., 867 F.2d 1188, 1195–96 (9th Cir. 1989).
100. Id. at 1194.
101. Id.
102. Luster v. Retail Credit Co., 575 F.2d 609, 615 (8th Cir. 1978) (deciding under Arkansas law that
“[e]ven if the statement is disbelieved . . . damages may be mitigated, but nevertheless awarded”); Roeben v.
BG Excelsior Ltd. P’ship, 344 S.W.3d 93, 98 (Ark. Ct. App. 2009).
103. Epstein, supra note 37, at 799. In fact, Epstein shrewdly recognizes that audiences in this world
will be less credulous, but he sees that distrust as a negative: “The influence of the press will diminish as there
will be no obvious way to distinguish the good reports from the bad . . . .” Id. at 800.
104. Gertz v. Robert Welch, Inc., 418 U.S. 323, 392, 400 (1974) (White, J., dissenting).
105. See Communications Decency Act of 1996, 47 U.S.C. § 230 (2012); SACK, supra note 65, §§ 1.4–
1.5.
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
470 ALABAMA LAW REVIEW [Vol. 71:2:453
anonymous speech all make it very difficult to pursue even meritorious ac-
tions.106 These considerations have led some to think that even in the domains
where defamation law applies, it is so ineffectual that its “[a]bolition would leave
victims of defamation little worse off than they are today.”107
In all of these diverse and broad contexts, false defamatory statements do
not result in legal liability. If defamation law was necessary to protect reputation,
we would expect to see broad discussions of how audiences are led astray by
expressions of opinions and political discussions—which clearly tend to include
hyperbole and falsehoods. Yet we find broad consensus that protection of opin-
ions should continue.108 The reasons for this protection are varied. Some courts
ground the distinction in liberty: “[T]he freedom to speak one’s mind is . . . an
aspect of individual liberty—and thus a good unto itself . . . .”109 Others look
more broadly at the social and democratic good brought about by the unham-
pered discussion of opinions.110 Still others suggest that allowing lawsuits for
expressions of opinion would bring about an unworkable deluge of litigation.111
Consequently, the Restatement provides that opinions are completely immune
unless the opinion implies some facts that are false and defamatory.112 These
views take a very pragmatic approach, seeing audiences as able to filter and
discount some falsities, and we observe the lack of serious pressure by courts,
lawmakers, and the public to reform this rule.
* * *
Defamation law, we saw, is understood as a bulwark against harm to repu-
tation. The doctrine itself is haphazard and confused, but the consensus on the
law’s aspirations is broad and deep. It is not surprising, then, that some harbor
anxieties about deregulation of speech, worrying that it would lead to the wan-
ton disregard of the truth and the trampling of individual rights. At the same
time, the existence of large pockets of unregulated speech in diverse areas is,
106. See generally Lidsky, supra note 53.
107. Anderson, supra note 41, at 490.
108. See, e.g., PROSSER AND KEETON, supra note 8, § 113A, at 813 (“The distinction [between opinions
and facts] is a necessary and important one.”); SACK, supra note 65, § 4:1, at 4–3.
109. Bose Corp. v. Consumers Union of U.S., Inc., 466 U.S. 485, 503 (1984).
110. Curtis Pub. Co. v. Butts, 388 U.S. 130, 149 (1967).
111. Wun-ee Chelsea Chen, Note, Pinning Opinion to the First Amendment Mat, 11 LOY. ENT. L.J. 567,
601–03 (1991) (responding to the Milkovich decision that loosened the protections of opinion in defamation
suits by predicting the courts “will be highly susceptible to a flood of litigation”).
112. RESTATEMENT (SECOND) OF TORTS § 566 (AM. LAW INST. 1977). As one court explains, “A
statement couched as an opinion that presents or implies the existence of facts which are capable of being
proven true or false can be actionable.” Levinsky’s, Inc. v. Wal-Mart Stores, Inc., 127 F.3d 122, 127 (1st Cir.
1997). The line between opinions and facts is fine. For example, some courts have held that a statement
indicating that a person did not pay for goods he owned is not defamatory but that a person refuses to pay a
debt is defamatory. Compare Sim v. Stretch [1936] 2 All E.R. 1237, with Neaton v. Lewis Apparel Stores, Inc.,
48 N.Y.S.2d 492, 495 (N.Y. App. Div. 1944).
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
2019] The Case Against Expanding Defamation Law 471
hopefully, an effective anxiolytic. We now turn to develop the main argument
of this Article, that defamation law can actually undermine reputational inter-
ests.
II. DEFAMATION & AUDIENCES
The typical defamation case involves a charge that a speaker cast a false
aspersion on another person and, because of this statement, that person lost his
esteem, which led to a loss of business, social, and romantic opportunities.113
This is a simple narrative and one that is common to many actions in tort: a
tortfeasor-wrong-harm model. The application of the model is deceptively sim-
ple: a wrongdoer committed a wrong (false aspersion) that was the proximate
cause of an identifiable injury (loss of social standing).
The problem with this account is that it obscures the individual capacity of
audience members to judge the truth. In other torts, the harm is either direct or
mediated by objects: an aggressive assailant hits the victim with an iron pipe; a
reckless woodcutter fells a tree that falls on the victim’s house; a distracted
driver runs his car over the victim; an obsessed ex-lover hacks his partner’s
email account and exposes private information. In defamation, the harm does
not follow immediately from the wrongful act itself but is rather mediated by
the audience of the offending speech. For speech to be harmful, it must be
believed.
The bilateral model of speaker–victim treats audiences deterministically, as
if they immediately believe whatever they hear. This dim view of audiences is
unpersuasive. Members of the public are not perfect reasoners, but neither are
they opiated masses.114 Regarding political speech, the Supreme Court has re-
fused to judge audiences according to their lowest possible denominator.115
Professor Lyrissa Lidsky powerfully notes that “[d]emocratic theory demands
faith in the rationality of citizens, and several members of the founding gener-
ation, steeped in the ideals of the Enlightenment, publicly professed this
faith.”116 Even in the commercial-speech context, the Supreme Court noted
that the states could not ban advertising just because they mistrust their citizens’
abilities to reach the “right” conclusions.117 In advertising, courts permit “mere
113. See supra Part I.B. We focus on “retail”-level defamation—i.e., defamation by individuals against
individuals or businesses. While much of what we have to say applies to libel by media outlets as well, ad-
dressing that rich context will needlessly muddle the discussion.
114. Lidsky, Rational Audience, supra note 16, at 815 (“First Amendment doctrines rely on a model of
the audience as rational, skeptical, and capable of sorting through masses of information . . . .”).
115. See, e.g., FEC v. Wis. Right to Life, Inc., 551 U.S. 449, 469 (2007).
116. Lidsky, Rational Audience, supra note 16, at 811.
117. Va. State Bd. of Pharm. v. Va. Citizens Consumer Council, Inc., 425 U.S. 748, 769–70 (1976).
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
472 ALABAMA LAW REVIEW [Vol. 71:2:453
puffery” because they believe that the public can sort between factual state-
ments and vapid boasting.118 Interestingly, members of the public show healthy
skepticism of speech in contexts where it is shielded from defamation law. For
example, a study of the credibility of news stories found that stories based on
anonymous sources were seen as far less credible than those with identifiable
sources.119 When individuals browse the web for information, they tend to
screen information given by unknown sources.120 To emphasize, both the
Founding Fathers and courts have recognized that people can be deceived and
misled.121 Overall, much of our understanding of the public involves the idea
that people, with differing degrees of ability, attempt to discern truth from false-
hoods.122
In what follows, we seek to tease out the implications of the basic insight
that audiences are also involved in reputational harms. We add to the bilateral
model of speaker–victim a third agent, the audience, and investigate the impli-
cations of various defamation law rules on all the relevant stakeholders. To de-
velop an intuitive explanation of our ideas, we use a stylistic example, but our
conclusions are general in nature, subject to a few caveats that we explore
later.123
A. Basic Example
Consider a case where 100 individuals each own a restaurant that advertises
its exclusive use of organic ingredients (or that it observes kosher, halal, or hy-
genic norms, or any other facts that would make patrons more likely to patron-
ize the restaurant). That is, we assume a situation where a person would like
others to believe certain favorable facts about her for owning such a restaurant,
as these would make others more likely to engage with her.
118. See generally David A. Hoffman, The Best Puffery Article Ever, 91 IOWA L. REV. 1395 (2006); Adi
Osovsky, Puffery on the Market: A Behavioral Economic Analysis of the Puffery Defense in the Securities Arena, 6 HARV.
BUS. L. REV. 333 (2016).
119. Miglena Mantcheva Sternadori & Esther Thorson, Anonymous Sources Harm Credibility of All Stories,
NEWSPAPER RES. J., Sept. 1, 2009, at 54; see also Ivanka Pjesivac & Rachel Rui, Anonymous Sources Hurt Credi-
bility of News Stories Across Cultures: A Comparative Experiment in America and China, 76 INT’L COMM. GAZETTE
641, 654 (2014) (finding that both American and Chinese audiences found stories citing anonymous sources
less credible than those using identified sources).
120. Miriam J. Metzger et al., Social and Heuristic Approaches to Credibility Evaluation Online, 60 J. COMM.
413, 416 (2010).
121. Lidsky, Rational Audience, supra note 16, at 811–17.
122. See, e.g., Abrams v. United States, 250 U.S. 616, 629–30 (1919) (Holmes, J., dissenting) (“[T]he
ultimate good desired is better reached by free trade in ideas . . . .”).
123. For a formal general analysis of this model, see Arbel & Mungan, supra note 27.
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
2019] The Case Against Expanding Defamation Law 473
For reputation to have any meaning, differences between restaurants must
exist that are not immediately apparent to a casual observer. If casual inspection
were sufficient, reputation would not play such an important role in our lives
in choosing which restaurant to attend, which product to buy, whom to trade
with, and even whom to date.124 Thus, suppose that 80% of the restaurants are
high quality—i.e., they really do, as they claim, serve organic food, maintain
sanitary standards, and refrain from mixing transfats in their ingredients—and
that 20% of the restaurants that claim to do so actually do not, making them
low quality.
Next, we consider the existence of an audience. Suppose that 100 different
patrons each contemplate going to a different restaurant. The patrons do not
know whether the restaurant they are about to attend is high or low quality. At
best, they can estimate that twenty of these restaurants are low quality.125 Be-
cause we are interested in the effect of defamatory speech, we assume that,
despite the risk, all patrons have already formed an intention to visit a restau-
rant. In particular, we assume that patrons are willing to attend a restaurant as
long as the risk of it being low quality does not exceed 40%. That patrons, and
more generally individuals, are willing to bear some risk in real life is self-evi-
dent.126 In this case, since only 20% of restaurants are low quality, patrons judge
this risk acceptable and, thus, would choose to attend all 100 restaurants.
The last element of this hypothetical is the speaker. Suppose that, while on
the way to the restaurant, each patron meets an ex-employee of the restaurant.
After exchanging pleasantries, the patron tells the employee that she is headed
to the restaurant and asks for his opinion. The ex-employee is privy to the in-
ternal workings of the restaurant, so he knows whether the owner adheres to
high- or low-quality standards. The employee can respond by either giving a
positive or negative review of the quality standards of the restaurant. As the
quality standards are a factual matter, such an assertion can be actionable. For
simplicity, we will denote a favorable review by “thumbs-up” and a negative
review by “thumbs-down.”127
124. See generally Luís M B Cabral, The Economics of Trust and Reputation: A Primer (June 2005)
(unpublished manuscript) (on file with the New York University Library Systems), http://pages.stern.nyu.
edu/~lcabral/reputation/Reputation_June05.pdf.
125. It is not important for our general argument that patrons accurately gauge the distribution since
we are interested in how their perceptions change with different defamation regimes.
126. Consider risks associated with choosing where to dine (food poisoning), whom to date (infidelity),
and whom to lend money (default). In all of these cases, there tends to be some threshold level beyond which
the risk is unacceptable.
127. Reviews, such as a thumbs-up, tend to be protected. See, e.g., Browne v. Avvo, Inc., 525 F. Supp.
2d 1249, 1251–53 (W.D. Wash. 2007) (holding that a rating is constitutionally protected if it results from a
subjective interpretation of facts, even if it claims to be based on objectively verifiable criteria); see also Avia-
tion Charter, Inc. v. Aviation Research Grp./US, 416 F.3d 864, 868–71 (8th Cir. 2005) (applying Minnesota
law to a ratings system), abrogated on other grounds by Lexmark Int’l, Inc. v. Static Control Components, Inc.,
572 U.S. 118 (2014), as recognized in Syngenta Seeds, Inc. v. Bunge N. Am., Inc., 773 F.3d 58 (8th Cir. 2014).
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
474 ALABAMA LAW REVIEW [Vol. 71:2:453
Defamation law exists because sometimes people choose to lie.128 If all em-
ployees were honest, each patron would trust the employee’s statement and
know for sure whether the restaurant she hopes to attend is high or low quality.
Unfortunately, this is not the case. Instead, to allow for the possibility of false
defamatory statements, we assume that some employees are prone to lying to
varying degrees.129 We capture the propensity to defame by thinking of ex-em-
ployees as deriving some value from denouncing their former employers, which
we denote by “v.” Specifically, ex-employees belong to one of three subgroups.
One subgroup of, say, sixty employees derives no value (i.e., “zero v”) from
besmirching its ex-employers; we will call such employees “honest.” Another
twenty ex-employees benefit slightly from besmirching their former employers
(“low v”), and the remaining twenty are disgruntled and thus derive intense
pleasure or value from badmouthing their former employers (“high v”). We will
call all positive v employees “dishonest” because they may want to lie by pre-
senting a high-quality restaurant as low-quality. Note also that employees’ val-
ues for v could very well be on a continuum rather than partitioned into the
discrete groups provided above, but this distinction does not matter for the
analysis.130 Figure 1 summarizes these basic assumptions. The stage is now set
to analyze the possible consequences of introducing defamation law into this
hypothetical.
Here we consider factual assertions such as “they never use organic materials” or “the kitchen is infested with
rats.”
128. Why people lie is a complicated question. See generally Arbel, Reputation Failure, supra note 10, at
18–21 & 28–29.
129. We take into account only false disparaging statements and table the discussion of “positive def-
amation.”
130. Indeed, each subgroup could be understood as comprised of heterogeneous employees who have
a value of “v” below or above the threshold (above zero or above “low”).
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
2019] The Case Against Expanding Defamation Law 475
Figure 1. Basic Setup
B. Analysis
Imagine, first, a world with very lax defamation law, which we denote as
D . One can think of this world as one where it is too costly to file a lawsuit,
0
where damages are too low, or where standards of proof are too high. In D
0,
victims of defamation have no recourse.131 How would restaurant owners, em-
ployees, and, ultimately, patrons behave in this world?
In answering this question, it is useful to start first with the restaurant
owner and work backward towards the patron’s choice of whether to dine at
the restaurant. The owner’s choice without defamation law is nonexistent: even
if she is the unfortunate subject of a false aspersion, she cannot file a lawsuit.
131. As argued supra Part I.C, there are various areas today that are outside of defamation law’s ambit.
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
476 ALABAMA LAW REVIEW [Vol. 71:2:453
She can only insist that her restaurant really uses fresh ingredients but so can
the owner of a low-quality restaurant. Thus, her protestations would not be very
credible.132
Next is the employee. He meets with the prospective patron and decides
whether to tell the truth. If he is honest, this decision is automatic: if the patron
is headed to a high-quality restaurant, the employee will give a thumbs-up; oth-
erwise, he will give a thumbs-down. In both cases, the review will be factually
correct. If, on the other hand, the employee is a positive v person, i.e., he derives
some benefit from besmirching his former employer, his decision-making will
be modified. Whether the restaurant is high or low quality, he will want to give
a thumbs-down to punish his former employer. The absence of defamation law
means that the employee will not be subject to liability; therefore, he can do as
he wishes—in this case, give a negative review regardless of the actual quality
of the restaurant. In other words, the employee will lie if the restaurant is high
quality or tell the truth if it is low quality.
Because there are 100 restaurants (eighty high-quality and twenty low-qual-
ity) and 100 employees (sixty honest and forty dishonest), we can assume the
following pairings of employees to restaurants.133 The sixty honest employees
will be matched with sixty restaurants, forty-eight high-quality and twelve low-
quality. The dishonest employees will be matched with forty restaurants, thirty-
two high-quality and eight low-quality. Because honest employees tell the truth,
they will give forty-eight thumbs-ups and twelve thumbs-downs. The dishonest
employees always give a thumbs-down, meaning they will give a thumbs-down
in forty cases (for a total of fifty-two thumbs-downs). Table 1 below summa-
rizes the ex-employees’ ratings.
132. For defamation to be relevant, accusations must “stick.” Of course, if the restaurant owner can
open her kitchen for all to see and the veracity of her claims can be checked, then the harm from defamation
may be negated, but reputation also becomes less important in the first place.
133. For simplicity of exposition, the matching is uniform with the same proportion of high- and low-
quality restaurants for honest and dishonest employees.
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
2019] The Case Against Expanding Defamation Law 477
Table 1.
Finally, we consider the patron. Recall that we assume the patron is willing
to attend a restaurant so long as the risk of it being low quality falls under 40%.
Before meeting with the employee, she assesses the risk at 20%, as she estimates
that there are only twenty low-quality restaurants out of 100 restaurants. She
finds the risk acceptable and plans to attend. Stopped on the way, she meets the
employee, who in forty-eight cases gives a thumbs-up and in fifty-two cases
gives a thumbs-down. The patron then tries to assess the probability that a res-
taurant is low quality given the review she hears.
If the patron receives a positive review, it is likely to come from an honest
employee because disgruntled employees will want to paint their former em-
ployers negatively.134 Hence, she knows she can trust a positive review to be
true and can dine at the restaurant with confidence. This will occur in forty-
eight cases, as described in Table 2 below. When she hears a negative review,
however, she finds it more difficult to assess its validity. She can tell that, in the
absence of defamation law, falsehoods are prevalent. Of the fifty-two negative
statements she expects to receive, only twenty can be true, because there are
only twenty low-quality restaurants. In other words, if she hears a negative re-
view, there is a twenty in fifty-two, or a 38%, chance that it is accurate and that
the restaurant is low quality.135 Now, because the patron’s risk tolerance is 40%,
the patron should still be willing to attend a negatively reviewed restaurant—
she will simply reason that falsehoods are so prevalent in this harsh D world
0
134. In the example, the only positive reviews come from honest employees. This is because no dis-
honest employees would want to spread favorable lies. As we noted supra note 129, we table issues of “positive
defamation,” as they fit more naturally under discussion of false advertising.
135. As noted supra note 125, the example makes the admittedly strong assumption that patrons have
a good sense of the distribution of low- and high-quality restaurants, but the analysis and theses developed
here apply even when patrons have inaccurate estimates.
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
478 ALABAMA LAW REVIEW [Vol. 71:2:453
that it is difficult to trust negative statements; so, she will discount the negative
review accordingly. This is not to say that negative reviews are not informative
at all; before hearing the negative review, she only considered 20% of
restaurants to be low quality, and now she realizes that she cannot identify those
20% from negative reviews since a negatively reviewed restaurant has a 38%
chance of being low quality. Still, such negative reviews are insufficiently trust-
worthy to make her abandon her plan to attend the restaurant. Hence, she might
choose to undertake the risk of attending a negatively reviewed restaurant in a
D world.
0
Table 2.
We move now to examine an alternative world, D , where defamation law
1
is very strict. In reality, the difference between D and D can be quite gradual;
0 1
it can simply be the difference between low and high damages; low and high
burdens of proof; lax and strict enforcement; or broad and narrow privileges.
In every case, we are contemplating an expansion of defamation law. For now,
however, let us maintain the clear distinction between these two “worlds.”
Starting again with the restaurant owner, she now has the option of bring-
ing a lawsuit against the ex-employee. Here, we must make some assumptions
about the relative costs and benefits of a defamation lawsuit; we assume for
now that a lawsuit is brought whenever an employee gives a false negative re-
view. This assumption is fairly innocuous; if the statement were positive, there
would be no incentive to bring a lawsuit, even if the statement were false. When
a statement is negative, there may be an incentive to bring a lawsuit, and our
assumption means that courts are sufficiently competent such that a plaintiff
will only consider it worthwhile to sue if the negative statement was false.136
How would employees behave in D ? An honest employee is not expected
1
to change her behavior. As we just analyzed, the restaurant owner is unlikely to
136. To be clear, it is not assumed that a meritorious lawsuit will always succeed but that the expected
recovery in such a case—given the probability of losing—is sufficient to motivate the restaurant owner to
file a lawsuit.
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
2019] The Case Against Expanding Defamation Law 479
bring a lawsuit because true communications are always protected, meaning that
the employer would tend to lose in court. Of course, there is some chance that
the employer would win—some courts make mistakes, and some lawsuits are
brought strategically—but omitting these considerations favors defamation law,
so we can table them for the time being.137 Hence, we again expect a total of
forty-eight thumbs-ups and twelve thumbs-downs produced by honest employ-
ees. If the employee is dishonest and the restaurant is low quality, the analysis
will again remain unchanged. After all, a truthful negative review generally is
not actionable—that fact does not change even if the motivation is malicious.138
The main change from D will be with respect to dishonest employees
0
matched with high-quality restaurants. Here, we will assume—in line with the
Court’s reasoning and reality—that it is impossible to deter all falsehoods.139
Because some people will still make false allegations, defamation law as a
deterrent will be imperfect. Therefore, we assume that the potential cost of a
lawsuit (i.e., the combination of the risk of losing with the eventual money judg-
ment) is enough to dissuade some employees from lying but not others. Specif-
ically, we will assume that the costs of a lawsuit are enough to dissuade the low
v group but not the high v group, which may include disgruntled employees or
those who moved to work for a competitor. Here, there would be sixteen low
v employees matched with high-quality restaurants. Although they would want
to say something negative, they will worry about the prospects of a lawsuit and
thus truthfully provide a positive review instead.140 The high v employees will
remain unperturbed and continue to defame, despite the risk of a lawsuit. Table
3 below summarizes the employees’ behavior.
137. In other words, these effects imply a chilling effect on truthful speech, which is a strong reason
to limit defamation law.
138. See SMOLLA, supra note 47, § 5:1. Contra MASS. GEN. LAWS ch. 231, § 92 (2018); Noonan v. Sta-
ples, Inc., 556 F.3d 20, 28 (1st Cir. 2009) (“[U]nder Massachusetts law, even a true statement can form the
basis of a libel action if the plaintiff proves that the defendant acted with ‘actual malice.’”).
139. See supra notes 53–59 and accompanying text.
140. Disgruntled employees might choose not to endorse the business but rather remain silent; how-
ever, it will become apparent that their silence means that they would like to lie but are afraid to do so. In
other words, silence will suggest that the quality is actually high.
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
480 ALABAMA LAW REVIEW [Vol. 71:2:453
Table 3.
Finally, let us consider the patron. In this case, there are sixty-four thumbs-
ups, and these are reliable measures of quality. Hence, the patron can safely
attend these sixty-four restaurants. In contrast, there are thirty-six thumbs-
downs. The patron will realize that some of these reviews are likely false despite
the presence of defamation law. Because there are only twenty low-quality res-
taurants, the probability that any negatively reviewed restaurant is indeed low
quality is twenty out of thirty-six, or 55%. In other words, negative reviews
become stronger indicators that the restaurant is indeed low quality. Because
this risk greatly exceeds 40%, the patron will be reluctant to attend any nega-
tively reviewed restaurant.141 Table 4 below summarizes the patron’s analysis:
Table 4.
141. The threshold of 40% is an arbitrary choice; the actual threshold in any specific case may be
different. But this will not change the point we make here. If, in this example, the risk tolerance was 50%,
patrons would still avoid the restaurant, as the perceived risk (55%) exceeds their threshold. Indeed, if we
were to imagine an even higher risk tolerance of, say, 80%, the result would change—patrons would assume
the risk in both D0 and D1. But this only demonstrates that the value and credibility of speech are lower when
people care less about its content.
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
2019] The Case Against Expanding Defamation Law 481
C. Audience Effects: Evaluation & Generalization
Based on the stylized example above, we can now analyze the effects of
expanding defamation law by comparing the (extreme) move from a world
where defamation law is effectively nonexistent (D ) to a world with it (D ).142
0 1
The primary lesson from this analysis is that defamation law can harm owners of
high-quality goods. In the baseline case, where the patron does not encounter any
employee, the patron judges the risk that a specific restaurant is low quality to
be small (20%). Therefore, the patron assumes the risk and attends all restau-
rants—including all eighty high-quality restaurants. As we saw, the encounter
with the employee in D does not change this conclusion.143 Here, patrons meet
0
employees, and some of the employees indeed say negative things. However,
the flood of false negative reviews reduces the reliability of each negative re-
view. Hence, patrons—while still finding negative reviews useful—will not find
them to be sufficiently indicative of low quality. As a result, the risk of a nega-
tively reviewed restaurant being low quality is still low, even if it is somewhat
higher than the baseline case. And so, in D patrons will continue to attend all
0,
eighty high-quality restaurants.
The shift to D —the introduction of strict defamation law—changes these
1
results quite drastically. There are now fewer negative reviews overall, because
some employees are deterred by the prospect of being sued for falsely defaming
their previous employer. The decrease in quantity means that each negative re-
view is now more credible than it was when there was a greater quantity of false
negative reviews. As such, patrons feel much more confident in relying on these
negative reviews and are reluctant to attend any negatively reviewed restaurant.
This outcome is benign when the restaurant is indeed low quality, but the in-
creased credence is problematic for those honest owners of high-quality restau-
rants who are besmirched. Patrons who previously used their discretion and
discounted negative reviews are now more trusting and are misled into avoiding
some high-quality restaurants.
Let us focus on the interpretation of this key result before offering qualifi-
cations. The expansion of defamation law makes “false signals” more costly.144
The speaker would be more likely to face liability for false statements. This leads
to a decline in the frequency of such statements, as speakers worry about the
specter of civil liability. With the decline of false signals, audiences will learn
either from experience or from observing the legal norms that statements are
142. See infra Part II.D.
143. See supra Part II.A.
144. More accurately, “false signals” are signals the speaker knows to be untrue. But the veracity of
statements is often a matter of belief; a statement may communicate a strong conviction or a weak one.
Stronger defamation law could also cull statements that are likely to be true but for which the speaker lacks
conclusive evidence.
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
482 ALABAMA LAW REVIEW [Vol. 71:2:453
generally reliable and will come to rely on them more often. The greater trust
the public places in statements, based on common-sense reasoning, the more
this makes victims susceptible to attacks, as the higher trust people place in
statements amplifies the pernicious effects of false aspersions and unfounded
statements.145 In this sense, stricter defamation law may undermine reputation
interests.
To illustrate, in the stylized example, the reputational harm to some owners
of high-quality restaurants is caused by the introduction of strict defamation
law (D ). Stated differently, without strict defamation law, lies are drowned out
1
by their own noise. When defamation law is expanded, falsehoods become rarer
and thus more harmful, as their rarity makes audiences more trusting.
A more general way to describe this result is by looking at signaling theory
again. Strict defamation law increases the cost of false statements; hence, it
strengthens the credibility of the signal a speaker sends. The greater credibility
of statements makes audience members more likely to act upon them. And
when the statement proves to be false, the audience is more easily deceived than
they otherwise would be. In sum, then, consideration of audience effects reveals
a basic trade-off, a “seesaw dynamic,” according to which reducing the number of
falsehoods increases their credibility.
Figure 2.
145. In the model, some speakers may have intentionally lied; however, statements of fact of which
the speakers are uncertain can be equally damaging, and no bad faith is required.
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
2019] The Case Against Expanding Defamation Law 483
D. Qualifications & Richer Considerations
Our analysis of audience effects has shown that stricter defamation law
could, in some cases, result in weaker protection of reputational interests. We
now move to consider a few important qualifications regarding the scope of
audience effects.
The most immediate qualification is that stricter defamation law does not
always harm reputational interests. In some cases, defamation law will harm
reputational interests, and in others, it will protect them, although to a lesser
extent than scholars have previously assumed. The key point is not that defa-
mation law invariably undermines reputation but that the “virtually axiomatic”
understanding of defamation law, which holds that defamation law unambigu-
ously protects reputation, is erroneous.146 The trilateral model calls for a more
nuanced examination than the standard bilateral model admits.
Another important qualification is that our discussion so far is only focused
on estimating behavioral effects, not social welfare. Like the courts, we restrict
attention only to whether defamation law protects reputation, not whether def-
amation law is, on net, desirable.147 But of course, the net social effect of ex-
panding or restricting defamation law is of paramount importance. Investigat-
ing these issues is complex, as it requires consideration of all the relevant
stakeholders—victims, audience members, and speakers. In terms of the res-
taurant example, the investigation would involve an assessment of the harm
suffered by patrons from attending a low-quality restaurant measured against
the gains from attending a high-quality one. Then, the analysis would need to
evaluate the potential costs imposed on speakers, taking into account their risk
aversion and the value they place on speaking, truthfully or otherwise. From
there, the analysis should account for the potential victims, the restaurant own-
ers, by measuring their increased losses in terms of business volume. Finally,
and outside our basic model, one might worry about the incentive to tell the
truth and the supply of ideas to the market. Even that, however, will not con-
clude the investigation, because the expansion of defamation law will also have
ex ante effects. These effects may be wide-ranging, from the decision of whether
to invest in restaurant quality to the pricing of entrees or how to treat ex-em-
ployees. In conducting such an analysis, different commentators will place dif-
ferent weights on deterrence, redress, and vindication. Engagement with these
questions is also context dependent, so we do not make any categorical state-
ments besides rejecting the court’s supposition that defamation law is good be-
cause it protects reputation.148 The effect of defamation law is nuanced, com-
plex, and, at times, self-contradictory. We shall address the meta question of
146. See supra notes 9–10.
147. For further reflections on social effects and social welfare, see infra Part III.
148. See, e.g., Hustler Magazine, Inc. v. Falwell, 485 U.S. 46, 52 (1988) (citing Gertz v. Robert Welch,
Inc., 418 U.S. 323, 340, 344 n.9 (1974)).
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
484 ALABAMA LAW REVIEW [Vol. 71:2:453
how we think courts should deal with this uncertainty, but for now, let us just
highlight our conclusion that even from the internal perspective of defamation
law, having stricter defamation law does not necessarily lead to better protection
of reputation.
A third, related qualification is that there are other benefits associated with
strict defamation law—ones rarely recognized by the courts. In particular, there
can be some value in increasing the perceived reliability of statements. If, in D ,
0
the patrons were hesitant to attend any restaurant, D would make speech more
1
credible and may persuade some patrons to attend some of the high-quality
restaurants. More generally, if one takes the view that audiences are overly cyn-
ical and skeptical, defamation law will have more value than the examples imply.
This qualification is important where the most pressing imperative is fighting
public skepticism. Inasmuch as the national context today is one of concern
that people are too trusting and reliant on fake news, we believe that there is a
good reason to be concerned with the credibility of falsehoods.
With these qualifications in mind, we now move to examine more closely
some of the assumptions used in the example. As we hope to show, the trilateral
model is applicable in broad settings. While the conclusions will differ based on
one’s assumptions, the framework itself continues to be useful under broad
settings.
1. Ignorant Audiences
People are sometimes ignorant, confused, or indifferent regarding the con-
tent of any law—and defamation law in particular. Therefore, it is possible that
members of the audience will erroneously think that a statement in a given do-
main is protected when it is not (“in my opinion, she is a deadbeat”) or unpro-
tected when it actually is (“she is good for nothing”).149 The possibility of igno-
rance of the law seems especially germane with respect to changes in
defamation law: would audiences know and react to the expansion (or contrac-
tion) of defamation law?
On reflection, ignorance of the law is less consequential than it appears.
Audience effects will continue to be important even when people only have a
very vague idea of the law. The only necessary trigger for the seesaw dynamics
is a vague (and possibly even erroneous) belief among audience members that
defamation law has changed. We believe that such a lax assumption is well
within reason; some libel cases receive considerable media coverage, and libel
policy is at the center of some political campaigns. More generally, defamation
law reformers essentially agree on this point: expanders believe that by making
149. Opinions that impute facts are actionable. RESTATEMENT (SECOND) OF TORTS § 566 (AM. LAW
INST. 1977); see also DOBBS, supra note 8, at 1113–14.
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
2019] The Case Against Expanding Defamation Law 485
defamation law more strict, members of the public will be restrained from mak-
ing negative statements,150 whereas free-speech advocates worry that strict laws
will chill prospective speakers.151 Both sides agree, however, that legal reform
can translate to different behavioral responses by the public. We only extend
these behavioral effects to members of the public as listeners rather than speak-
ers. We do not deny the possibility that, in some cases, audiences will be com-
pletely unaware of changes in the law—in which case audience effects would
be negligible. But whether that is the case requires analysis and empirical sup-
port. Consideration of audience effects is warranted, then, even in these do-
mains.
2. Audience Rationality & Sticky Falsehoods
A related issue is audience rationality. Insofar as individuals diverge from
the model of a perfectly rational agent, they may treat suspect assertions with
either too much or too little trust. Of the two, there may be a particular reason
to worry that individuals will be overly trusting, i.e., that falsehoods will be
“sticky” even when they are clearly dubious. Such a concern is motivated in part
by the psychological phenomenon of “anchoring,” where individuals are said
to react to irrelevant information in a way contrary to what rational decision
theory would predict.152 Experiments suggest negotiations, for example, may be
influenced by introducing numbers and figures that have little bearing on the
issue.153 Moreover, one study found that telling subjects that a speaker is likely
to lie because of a financial incentive made them more likely to believe the
speaker.154
If falsehoods are completely sticky, then audience effects are unimportant.
It is more urgent to control the quantity of false statements, and the perceived
reliability of false statements can be neglected, as they are, by assumption, very
reliable.
However, it is highly unlikely that falsehoods are completely sticky and that
they hopelessly captivate individuals.155 Political campaigns frequently involve
a stream of lies, misstatements, and spurious allegations.156 Many marketing
150. See supra notes 4–11 and accompanying text.
151. See supra notes 54–59 and accompanying text.
152. Russell B. Korobkin & Thomas S. Ulen, Law and Behavioral Science: Removing the Rationality Assump-
tion from Law and Economics, 88 CALIF. L. REV. 1051, 1100–02 (2000).
153. See, e.g., Adam D. Galinsky & Thomas Mussweiler, First Offers as Anchors: The Role of Perspective-
Taking and Negotiator Focus, 81 J. PERSONALITY & SOC. PSYCHOL. 657, 658 (2001).
154. George Loewenstein et al., The Limits of Transparency: Pitfalls and Potential of Disclosing Conflicts of
Interest, 101 AM. ECON. REV. 423, 425–26 (2011).
155. Higher stakes might make lies more sticky, a point we develop infra Part III.A.1.
156. See generally Jacob Rowbottom, Lies, Manipulation and Elections—Controlling False Campaign Statements,
32 OXFORD J. LEGAL STUD. 507 (2012).
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
486 ALABAMA LAW REVIEW [Vol. 71:2:453
campaigns involve deception, puffery, and images of people cheerfully laughing
while taking a payday loan.157 Even socializing and dating involve dissembling,
white lies, and putting on appearances.158 It is impossible to survive the modern
world without an ability to dismiss hyperbole and discount patent falsehoods.159
Importantly, even those who are concerned with the stickiness of lies would
take those concerns to their logical conclusion and support the imposition of
prior restraints, criminalization of lies, and agency review of media articles.160
The stickiness of falsehoods, then, is a matter of degree. And when false-
hoods are not completely sticky, audience effects continue to be important. One
corollary of the seesaw dynamic is that if it can be shown that a lie is especially
sticky, or that a certain type of falsehood is especially hard to refute, then there
is a stronger case for preferring to decrease the number of lies rather than con-
trol their credibility. And the obverse is also true: if some lies are simply “hot
air,” “bullshit,” or “puffery” and easily dismissed, the priority should be reduc-
ing their credibility.
3. Imperfect and Costly Enforcement, Litigation, and Execution
Enforcement is one of the biggest problems in law161: it is costly, emotion-
ally draining, and lengthy to pursue a lawsuit; winning the case may turn on
luck; damages may undercompensate the victim; and collecting the judgment is
always risky.162 All of these considerations make it less likely that a victim will
bring a lawsuit against the defaming party. Now, if it were possible to correct
these issues—i.e., to punish defamers with full expropriation of their wealth or
even to exact corporal punishment at all costs—then it is theoretically possible
to completely eliminate the incentive to lie. In such a world, the seesaw dynamic
would disappear: people could safely rely on statements, as all statements would
be truthful and none would dare to mislead. But because such a policy is neither
157. See, e.g., Payday Lender Withdraws Ad that Encourages ‘Worry-Free’ Borrowing, GUARDIAN (July 23,
2014, 5:13 AM), https://www.theguardian.com/money/2014/jul/23/pounds-to-pocket-payday-lender-
withdraws-advert-borrowing.
158. See generally Irina D. Manta, Tinder Lies, 54 WAKE FOREST L. REV. 207 (2019) (exploring various
lies used in the dating context).
159. Indeed, part of the tragedy of individuals suffering from neurodegeneration is that they start
falling prey to manipulations that they would otherwise avoid. See Tal Shany-Ur et al., Comprehension of Insincere
Communication in Neurodegenerative Disease: Lies, Sarcasm, and Theory of Mind, 48 CORTEX 1329, 1335 (2012) (doc-
umenting the loss of the ability to detect deception by patients with neurodegenerative diseases).
160. See Brandenburg v. Ohio, 395 U.S. 444, 447–48 (1969) (holding that even advocacy of violence,
unless it presents an imminent threat, is protected by the First Amendment).
161. See generally Yonathan A. Arbel, Adminization: Gatekeeping Consumer Contracts, 71 VAND. L. REV.
121 (2018) (analyzing the difficulty of enforcement in the context of private credit contracts).
162. See generally Yonathan A. Arbel, Shielding of Assets and Lending Contracts, 48 INT’L REV. L. & ECON.
26 (2016) (analyzing the ways debtors can circumvent creditors and the circumstances under which asset
shielding is most likely).
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
2019] The Case Against Expanding Defamation Law 487
plausible nor advisable, under any feasible policy, there will always be some lies.
Thus, the seesaw principle is relevant.163
Still, audience effects bear on changes to the enforceability of defamation
lawsuits. Any removal of practical limitations to enforcement is equivalent to
increasing the scope of defamation law. Fewer practical hurdles mean more
lawsuits, increased litigation means an increase in potential defamation law en-
forcement, and increased enforcement creates further disincentives to keep
speakers from lying. Per the seesaw dynamics, more efficient enforcement
would lead to fewer, but more harmful, falsehoods and misstatements. Hence,
we can draw a broad conclusion regarding a host of interventions—cheaper
enforcement methods; larger or punitive judgments; fee-shifting arrangements;
subsidies to support plaintiffs; lower proof thresholds; or speedier legal pro-
cesses. In all of these cases, despite important nuances, audience effects would
lead to similar seesaw dynamics.
4. Social and Moral Norms
Most societies have social norms against lying, and individuals may have an
internal moral sense that spreading falsehoods is wrong. How does the presence
of such norms change the audience effects on defamation law?
Generally speaking, internalization of norms, social or moral, increases the
proportion of honest people. And if all men were angels, defamation law would
become obsolete.164
Of course, any discussion of defamation law rests on the recognition that
some people lie. And with this assumption, norms naturally fit into the trilateral
model. Recall that in the example, some individuals were assumed to be honest
(i.e., zero v). This honest attitude could be a matter of outside preference—
perhaps the example employees bear no ill will towards their ex-employers—or
it could be the result of more innate preferences—such as from internalizing
the norm against lying. Whatever the reason, some individuals could be ex-
pected to tell the truth independent of the law. As the proportion of such em-
ployees rises, consideration of audience effects reveals that all statements will
be viewed, on average, as more credible. In other words, the proportion of
honest speakers affects the baseline level of credibility of statements, independ-
ent of any specific defamation law regime. Changing defamation law would only
change the reliability of statements relative to this baseline.
This observation highlights the importance of considering audience effects
when making decisions regarding the scope of defamation law. The common
approach—assuming a simple relationship between stricter defamation law and
163. In Arbel & Mungan, supra note 27, we explain that overly strict defamation laws would also deter
truth speaking, rendering speech uninformative to audiences.
164. In terms of our analysis, this is equal to the possibility that enforcement is perfect.
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
488 ALABAMA LAW REVIEW [Vol. 71:2:453
greater protection of reputation—fails to account for the prevalence of such
background norms.
E. Would Expanding Defamation Law Advance Its Goals?
Having surveyed audience effects and some of their qualifications, we can
now evaluate the outcomes of expanding defamation law on the basis of three
main goals of the law—deterrence, redress, and vindication.
1. Protecting Reputation Through Deterrence
Defamation law is often justified on the belief that it protects reputation by
deterring individuals from falsely besmirching others. The consideration of au-
dience effects adds an important qualification. While defamation law deters
some false statements—just as advertised—it also increases the perceived reli-
ability of undeterred statements.
The discussion of the example focused on this point exactly. On the posi-
tive side, defamation law reduced the thirty-two lies in D to only twenty-four
0
in D , or a 25% reduction. But, on the other hand, the false negative reviews in
1
D are less believable precisely because there are more lies. In this construct,
0
negative reviews are split between twenty truthful to thirty-two false, a truth-
to-lie ratio of 0.625. In D , there are, again, twenty truthful negative reviews but
1
only twenty-four false reviews, a higher truth-to-lie ratio of 0.833. As a result,
every negative statement in D is more credible (i.e., the diagnosticity of
1
negative reviews is higher). If lies are more credible, people are more likely to
trust and act on them, as the patrons do in the example, hence enhancing their
social harm. In conclusion, the deterrent effect of defamation law can under-
mine reputation, contrary to the literature on the topic.165
2. Protecting Reputation Through Redress
The second goal of defamation law is to provide redress to victims, a func-
tion that the Court emphasized in Milkovich.166 Justice Rehnquist’s narrative
highlights that defamation law is not only about protecting reputation but also
compensating injured victims of the false allegations.167 It is widely understood
that current defamation law often leads to under-compensation of victims.168
With that in mind, would the expansion of defamation law not better protect
victims?
165. See supra Part I.A.
166. Milkovich v. Lorain Journal Co., 497 U.S. 1, 12 (1990).
167. See id.; see also ELDREDGE, supra note 47, § 3, at 5–6.
168. On the intentional imperfection of defamation law, see supra notes 54–59 and accompanying text.
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
2019] The Case Against Expanding Defamation Law 489
We believe that it is illogical to think about redress without considering
audience effects. The underlying assumption of the redress argument is that a
harm already exists and the law needs to allocate compensation. This argument,
however, ignores the effect that offering compensation will have on the crea-
tion of the harm. Consider again the restaurant example. There, shifting from
D to D led to reputational harm to all owners of high-quality restaurants be-
0 1
cause falsehoods became more reliable. In other words, one reason why owners
of some high-quality restaurants suffered a commercial loss—for which they
would seek redress—was defamation law itself.
A vivid illustration of some problems with the redress argument is a recent
spat on Twitter. Elon Musk, the outspoken billionaire, engaged in a heated de-
bate on Twitter with a British citizen, Vernon Unsworth.169 Angrily, Musk
tweeted to his 22 million followers an accusation that Unsworth was a pedo-
phile: “Sorry pedo guy, you really did ask for it.”170 Critically, no evidence was
presented to support this devastating allegation. Six weeks later, with still no
evidence being adduced, Musk tweeted a response to an inquiry made by a
Twitter user going by the handle “@yoda”171:
Musk used the absence of a lawsuit as positive evidence of Unsworth’s guilt.
It is because a lawsuit was not brought that any member of the audience could
surmise that Unsworth had something to hide. Rather than the standard narra-
tive of defamation law protecting reputation, defamation law was effectively
used here as a sword to undo it. This Tweet seems to have worked, and it “set
off a chain reaction on Twitter and in the media,” leading many members of
the public to accept Musk’s statement as true and greatly tarnish Unsworth’s
169. Elon Musk in New Rant at Thai Cave Rescuer, BBC (Sept. 5, 2018), https://www.bbc.com/news/
world-us-canada-45418245.
170. Elon Musk (@elonmusk), TWITTER (July 15, 2018, 10:10 AM), https://i.imgur.com/kNzmBOJ.
png. The original tweet has since been deleted.
171. Elon Musk (@elonmusk), TWITTER (Aug. 28, 2018, 12:41 PM), https://twitter.com/elonmusk/
status/1034481160783585280, https://i.imgur.com/xmXZMlD.png.
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
490 ALABAMA LAW REVIEW [Vol. 71:2:453
reputation.172 A lawsuit was eventually filed,173 but one worries that it was too
late to reverse the harm. In other words, the harm that redress wishes to solve
can, in some cases, be exacerbated by the existence of financial compensation.
At the same time, we recognize some countervailing considerations in favor
of the redress argument. While defamation law may exacerbate the harm, it is
not the only reason why harm exists. It is possible that even in the absence of
any defamation law, falsehoods would have some negative effects. Moreover,
redress offers a venue for compensation for at least some victims, even if not
all sue. These are two important points, but they should not detract from the
existence of other considerations. When considering such effects, it is still crit-
ical to evaluate whether they are sufficiently strong to overcome the negative
effect that some victims will suffer from more expansive defamation law. How-
ever one chooses to answer this question, it must involve a consideration of
audience effects.
3. Protecting Reputation Through Vindication
The last justification for defamation law is that victims of defamation can
vindicate their good names by proving the falsity of the allegation. The im-
portance of vindication goes beyond judicial remedies. Defamation, after all,
involves the audience and one’s esteem in the community. Through the process
of adjudication, individuals are able to redeem themselves in the eyes of others
and regain their former esteem.
Much like redress, we start by highlighting that even proponents of defa-
mation law would agree that vindication is limited.174 Courts are not perfect
truth finders, and the costs and difficulties of bringing a lawsuit mean that some
innocent victims will not be able to vindicate their good names. So, again, we
are contemplating the imperfect fulfillment of this goal.
Beyond this preliminary concern, vindication suffers from three previously
unrecognized flaws: “circularity,” “entrenchment,” and “antivindication.” Like
its circularity in providing redress, defamation law may be the reason why vin-
dication is needed in the first place—i.e., defamation law self-perpetuates the
harm it attempts to remedy. The more audiences believe falsehoods, the more
important it becomes to offer victims a way to vindicate their good names.
Conversely, if falsehoods are not believed (as is possible under D ), then there
0
is much less to vindicate. The need for vindication is predicated on audiences
believing the statement, and while audiences may believe statements for reasons
172. Complaint at 14, Unsworth v. Musk, No. 2:18-cv-8048, 2018 WL 4403350 (C.D. Cal. Sept. 17,
2018); see also id. at 18–20 (arguing that the tweet was harmful to the plaintiff’s reputation).
173. Id. at 1.
174. See generally Anderson, supra note 41, at 509.
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
2019] The Case Against Expanding Defamation Law 491
that are independent of the existence of defamation law, the availability of def-
amation law may exacerbate the problem.
The second issue is “entrenchment.” Judges and juries are limited in their
ability to determine factual truth.175 Given that, a world where the courts are
used as the final arbiters of truth—as the vindication argument assumes—is a
world where the stakes of judicial mistakes are dramatic. If an innocent victim
brings a lawsuit and—erroneously—loses the case, then this will entrench the
allegations against her in the public eye. After all, courts play a key role in cre-
ating reputational information.176
Third, consider the problem of “antivindication.” Recall the Elon Musk
example.177 There, the fact that a lawsuit was not brought was evidence that the
allegation was true—and in the court of public opinion, it was the only piece of
evidence. Antivindication is the (rational) inference that failure to litigate is ev-
idence of the veracity of the allegations, which is an acute problem if the victim
chooses not to bring a lawsuit for other reasons, such as funding, time, or avail-
ability of evidence. This is also troubling on the macro level, as stronger defa-
mation law may engender an environment, a social equilibrium, where failure
to bring suit is a signal that the allegation is true, thus inviting extensive litiga-
tion.
Taken together, circularity, entrenchment, and antivindication undermine
the claim that defamation law can be unambiguously justified on the basis of pro-
tecting reputation through vindication. While there might be cases where vin-
dication could justify an expansion of defamation law, this justification is not
general in nature and requires domain-specific analysis.
III. DEFAMATION LAW & AUDIENCE EFFECTS
A. The Desirable Scope of Defamation Law
Beyond critique, the trilateral model of defamation offers a new way of
thinking about the proper scope of defamation law. In what follows, we offer
a preliminary sketch of how such an analysis will proceed, although we empha-
size that we are not engaging in a full social-welfare analysis here. We follow
this analysis by showing how it sheds new light on defamation law doctrine by
evaluating political defamation.
175. See Frederick F. Schauer, Language, Truth, and the First Amendment: An Essay in Memory of Harry
Canter, 64 VA. L. REV. 263, 279 (1978).
176. The legal process also reveals important reputational information not through case outcomes but
through the process itself. See Kishanthi Parella, Reputational Regulation, 67 DUKE L.J. 907, 910 (2018).
177. See supra Part II.E.2.
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
492 ALABAMA LAW REVIEW [Vol. 71:2:453
1. Defamation Law in Employment
Defamation doctrine has special rules that apply more strictly in the context
of employment than in any other context. In employment cases, a libelous state-
ment is actionable per se (i.e, without the need to prove special damages) as
long as it is likely to “injure a man in his profession and calling.”178 This added
protection is exceptional and is only afforded outside the employment context
against false criminal allegations or assertions that the victim suffers from a
“loathsome disease.”179 But the reason why these special protections apply is
poorly understood. Why is it more exigent to protect employment than, say,
social standing in the community? Should such a distinction be maintained?
Our framework resolves this puzzle.
In the context of employment, a perennial concern is that a vindictive em-
ployer may defame an ex-employee.180 While an employer may have valid com-
plaints about the ex-employee, the employer is also incentivized to besmirch
the employee. This may be done to prevent talent from defecting to the com-
petition; to send a warning message to current employees that they will face
difficulties in finding another place of employment; or to protect brand image
or save face.181
From the employee’s perspective, false allegations by a former employer
can be especially disastrous. In our society, employment is important not only
as a source of income but also as a matter of fulfillment and social status. But
prospective employers are likely to deny employment based on even a small risk
of employee malfeasance, as a single bad employee can cost the organization
dearly.182 Putting aside embezzlement and other criminal activities, undesirable
employee traits such as tardiness, insubordination, and poor attitude can cost a
business significant resources and production efficiencies. Perhaps the clearest
example is that of a nanny for whom any allegations of abuse can ruin any
chance of ever finding employment again.183
In such circumstances, even a low probability that an allegation against an
employee is true can lead to denial of future employment. Thus, the reliability
of the allegation is a secondary concern to the existence of an allegation. The
178. See, e.g., Lega Siciliana Soc. Club, Inc. v. St. Germaine, 825 A.2d 827, 832 (Conn. App. Ct. 2003)
(quoting Proto v. Bridgeport Herald Corp., 72 A.2d 820, 825 (Conn. 1950)).
179. See, e.g., Lent v. Huntoon, 470 A.2d 1162, 1168 (Vt. 1983).
180. See, e.g., Frank B. Hall & Co. v. Buck, 678 S.W.2d 612, 630 (Tex. App. 1984) (affirming an award
of $1.9 million for a wrongful reference). Note, however, that most states provide immunity to employers
who write letters of reference. See Markita D. Cooper, Job Reference Immunity Statutes: Prevalent but Irrelevant, 11
CORNELL J.L. & PUB. POL’Y 1, 11 (2001).
181. This context is also marked by a relatively unique form of liability for favorable false reviews, but
we leave this topic for future investigation.
182. See Cooper, supra note 180, at 2.
183. Insurance policies are unlikely to cover loss of employment due to loss of reputation.
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
2019] The Case Against Expanding Defamation Law 493
focus should be on limiting the supply of false allegations. As the seesaw dy-
namics demonstrate, this goal can be achieved by making defamation law
stricter. This is why per se protection in such cases may be justified.
Interestingly, this logic may also explain the other special exceptions of per
se libel, allegations of crime or a “loathsome disease.”184 These are both contexts
where the small probability of the truth of the allegation can lead to ex-com-
munication or loss of critical social opportunities and the credibility of allega-
tions is a less pressing concern. Hence, limiting the supply of false allegations
can be more important than controlling their credibility—given that even a
small probability of truth is sufficiently damaging.
2. Defamation Law in Consumer Markets
Reputation for quality plays a key role in consumer markets. In this context,
reputational information is usually manifested through product and business
reviews and ratings. This information is then used by consumers to compara-
tively evaluate different products. In consumer markets, having a robust repu-
tation usually supplants more direct regulatory measures.185 In this context, def-
amation lawsuits usually contend that a spiteful consumer lied or that a
competing business faked a review.186 Importantly, such lawsuits are rarely suc-
cessful.187
The analysis suggests that the case for defamation law is weaker in the con-
text of consumer transactions. First, the distribution of reviews tends to follow
a “j-shaped” distribution with a small mass of negative reviews, very few mid-
dling reviews, and an overwhelming mass of positive reviews.188 Positive re-
views are obviously not sponsored by competitors; they may be the result of
either organic consumer reaction or deliberate attempts by the firm to bolster
its own reputation. If most of these positive reviews are from honest reviewers,
then the stakes for defamation law are low, as false reviews by disgruntled con-
sumers can be overwhelmed by the volume of positive reviews. There is another
reason to question the relevancy of defamation law in this context. There is
good reason to worry about the accuracy of judicial determinations in this area
because most reviews are based on intimate and not easily verifiable interactions
184. See Lent, 470 A.2d at 1168.
185. See Yonathan A. Arbel & Roy Shapira, Consumer Activism: From the Informed Minority to the Crusading
Majority, 69 DEPAUL L. REV. (forthcoming 2019) (discussing reputational pressures as a mode of consumer
governance).
186. See generally Yonathan A. Arbel & Roy Shapira, Theory of the Nudnik: The Future of Consumer Activism
and What We Can Do to Stop It, 73 VAND. L. REV. (forthcoming 2020) (evaluating how idiosyncratically moti-
vated consumers can enforce market norms).
187. See id.; see also Lidsky, supra note 53, at 883 n.144.
188. See Arbel, Reputation Failure, supra note 10.
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
494 ALABAMA LAW REVIEW [Vol. 71:2:453
between a consumer and a product or between a consumer and a service pro-
vider. It is not an accident that most cases resolve in favor of the consumer,
albeit after protracted and costly litigation. Lastly, and outside our model, it is
unclear what benefits consumers receive from sharing opinions, and the impo-
sition of large costs—in the form of litigation risks—may deter many from pro-
ducing this public good.189
These considerations, although preliminary in nature, suggest a presump-
tion against liability or, at the very least, against the expansion of defamation
liability for consumer reviews. Elsewhere, one of us advocated the use of a con-
sumer-review privilege for different reasons, and the analysis presented here is
consistent with this recommendation.190
3. Political Speech
For political speech, the general rule is no liability under defamation law
absent a showing of some hard-to-prove conditions.191 Is this a wise rule? This
question is especially timely now due to the commonly heard allegation of “fake
news” and the President’s attempts to expand the reach of defamation law.192
Ironically, it is the very salience of the fake news debate that abates the risk
of fake news. Under the standard bilateral-tort model, a deluge of fake news
would provide a compelling reason to expand libel laws.193 However, under the
trilateral model presented here, which accounts for how audiences process in-
formation, it is clear to see that one way to combat the harmful effect of fake
news is to educate the public about the phenomenon. Now that the message is
out and it is widely recognized that some stories are outright fake or otherwise
politically biased, some—although not all—of the harm of fake news is miti-
gated. That this conclusion is not surprising suggests the practical relevance of
our main argument—one cannot think about the behavioral effects of laws
without accounting for how they affect beliefs. In any case, given the competing
considerations at play in this context—the importance of public speech, the
danger of letting the government decide what is politically true, and the im-
portance of free press—the default position should be for defamation law to
genuflect to free speech.
189. See id.
190. See id.
191. See supra Part I.B.
192. See supra note 4 and accompanying text.
193. See supra Part II.D.2.
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
2019] The Case Against Expanding Defamation Law 495
B. Some Constitutional Ramifications
The application of constitutional protections to defamation law dates back
to the seminal 1964 decision in New York Times Co. v. Sullivan.194 The body of
constitutional law evolving from this decision has been built on one central
dilemma: How does one balance freedom of expression against the protection
of reputation? In the case law, one finds a careful balancing of these two com-
peting interests, with reputation prevailing in some cases and speech in the
rest.195
The analysis presented here has the potential to upend some parts of this
standard constitutional analysis. We currently view defamation law as a trade-
off between free speech and reputation, but in fact, the protection of reputation
is much weaker than traditionally recognized, and stronger defamation law may
undermine some reputational interests. As the seesaw dynamic highlights that
from the internal viewpoint of protection of reputation, the inevitable trade-off
is between fewer but more credible lies versus a greater number of lies of lesser
credibility. Whether one effect is deemed more important than the other is a
matter of analysis and cannot be ignored when considering the constitutionality
of expanding defamation law.
In many of the Supreme Court decisions, a media outlet was involved.196
Liability for defamation in this context raises more difficulties, as reporters
would often have less than full certainty about the accuracy of their own stories.
This difficulty is compounded by the problem of exposing sources. While our
focus here was on defamation between private parties, audience effects are also
relevant to the examination of regulation of libel by the media.
Without engaging in a full analysis, audience effects suggest that tighter reg-
ulation of the media would have the salutary effect of increasing the credibility
of reports, the undesirable effect of chilling the dissemination of some infor-
mation, and mixed effects on the reputation of the subjects of the reports. As
we have already noted, in the context of reporting about political figures, it is
the case today that the media has little exposure to defamation law. And indeed,
it would seem that the public treats political reporting with a greater degree of
suspicion.197
Doctrinally, courts should consider this trade-off in every case involving a
potential expansion—or contraction—of defamation law: would having fewer
falsehoods outweigh the cost of having more harmful falsehoods? Courts
should only adjust defamation law doctrine if the answer is a net positive. Only
then can a court appropriately balance protecting reputation against other val-
ues.
194. 376 U.S. 254 (1964).
195. See, e.g., Petro-Lubricant Testing Labs., Inc. v. Adelman, 184 A.3d 457, 461 (N.J. 2018).
196. E.g., New York Times Co., 376 U.S. at 256.
197. Kathleen M. McGraw et al., The Pandering Politician of Suspicious Minds, 64 J. POL. 362, 365 (2002).
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
496 ALABAMA LAW REVIEW [Vol. 71:2:453
Even if one thinks that the net effect of defamation law is positive—i.e.,
that defamation law tends to protect reputation more than it harms it—it is still
critical to understand that the net effect is smaller than the gross effect of def-
amation law. Today, judges think of defamation law in terms of its gross effect
(protection of reputation), without recognizing that there is an internal trade-
off due to the seesaw principle, which makes the net effect of defamation law
on reputation smaller. Realizing the difference, even in cases where the net ef-
fect is positive, is important when judges purport to balance free speech against
reputational interests. That the net effect is smaller has especially acute impli-
cations for all of these cases where judges declared that the cases involved a
close call.198 If the net effect is smaller, it means that the close call was only
reached by overweighing the perceived importance of defamation law. Realiz-
ing this, the case no longer becomes a close call and may, indeed, come out the
other way.
Assessing the net effects of defamation law is not always straightforward,
and in some cases, there will be considerable uncertainty. Such difficulties are
common when policy changes are contemplated. However, in the context of
defamation law, the significance of these difficulties is quite limited, because
freedom of speech and the press are constitutionally protected. Consequently,
if there is uncertainty over whether the restrictions of these freedoms will result
in a benefit, such restrictions are unwarranted. In other words, if there is real
doubt whether expanding defamation law would provide any benefit, courts
should err on the side of free speech.
CONCLUSION
Defamation law is a feature of most modern systems of law. In the United
States, it challenges values enshrined in the First Amendment. Courts and
scholars have battled over the proper balance between truth finding, freedom
of speech, and protection of reputation. But despite the perennial struggle over
these issues, a fundamental idea was uncritically accepted—that defamation law
is necessary for the protection of reputation. But the idea that defamation law
protects reputation is misguided. Due to the failure to properly evaluate audi-
ence effects, courts failed to notice that stricter defamation law can actually
undermine reputational interests.
To properly account for audience effects, one has to recognize that defam-
atory speech necessarily involves a thinking, autonomous audience whose will-
ingness to believe any given proposition is not divorced from the legal rules. In
a world with no defamation law, the cost of sending false signals—misstate-
ments of fact—is low. Audiences would, therefore, tend to be more skeptical
and cautious. This skepticism, however, guards against the harmful effects of
198. See, e.g., Petro-Lubricant Testing Labs, Inc., 184 A.3d 457.
Electronic copy available at: https://ssrn.com/abstract=3311527

<<PAGE_BREAK>>

4 ARBELMUNGAN 453-497 (DO NOT DELETE) 12/4/2019 7:18 PM
2019] The Case Against Expanding Defamation Law 497
falsehoods. With defamation law, individuals would be much more willing to
believe statements of unknown veracity because they would believe that the
threat of legal liability would deter many from misstating the truth. The greater
credulity would make any residual falsehoods much more pernicious.
We imagine future work that seeks to establish and test the limits of these
ideas empirically, alongside work that seeks to deepen the investigation in the
context of tort doctrines and related First Amendment scholarship—expanding
it to adjacent fields, such as privacy law and advertisement law. We think there
is good reason to pause and reflect on the necessity of this costly social institu-
tion in the age of anonymous and untraceable online speech. For the time being,
however, our greatest hope is that our analysis will encourage the explicit con-
sideration of audience effects in the jurisprudence of defamation law.
Electronic copy available at: https://ssrn.com/abstract=3311527`

func main() {
    fmt.Println(articleText)
}
