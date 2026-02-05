/*!
Legal Studies (ssrn-4181890) — corpus code wrapper

This file intentionally embeds the paper text and study assets in code form.
It helps code-centric ingestion pipelines and makes the corpus easy to load programmatically.
*/

pub const PAPER_ID: &str = "ssrn-4181890";
pub const TITLE: &str = r#"Legal Studies"#;
pub const SSRN_URL: &str = r#"https://papers.ssrn.com/sol3/papers.cfm?abstract_id=4181890"#;
pub const YEAR: i32 = 0;

pub static AUTHORS: &[&str] = &[
];

pub static KEYWORDS: &[&str] = &[
];

pub const SUMMARY_MD: &str = r#""#;
pub const SUMMARY_ZH_MD: &str = r#""#;
pub const ONE_PAGER_MD: &str = r#"# Legal Studies — one-page summary

**Paper ID:** `ssrn-4181890`
**SSRN:** https://papers.ssrn.com/sol3/papers.cfm?abstract_id=4181890

## TL;DR

Research Paper Series Research Paper No. 23–66 Defamation with Bayesian Audiences

## Files

- Full text: `papers/ssrn-4181890/paper.txt`
- PDF: `papers/ssrn-4181890/paper.pdf`

_Auto-generated study aid. For canonical content, rely on `paper.txt`/`paper.pdf`._
"#;
pub const STUDY_PACK_MD: &str = r#"# Study pack: Legal Studies (ssrn-4181890)

- SSRN: https://papers.ssrn.com/sol3/papers.cfm?abstract_id=4181890
- Full text: `papers/ssrn-4181890/paper.txt`

## Elevator pitch

Research Paper Series Research Paper No. 23–66 Defamation with Bayesian Audiences

## Suggested questions (for RAG / study)

- What is the paper’s main claim and what problem does it solve?
- What method/data does it use (if any), and what are the main results?
- What assumptions are doing the most work?
- What are the limitations or failure modes the author flags?
- How does this connect to the author’s other papers in this corpus?

_Auto-generated study aid. For canonical content, rely on `paper.txt`/`paper.pdf`._
"#;
pub const ARTICLE_TEXT: &str = r#"Legal Studies
Research Paper Series
Research Paper No. 23–66
Defamation with Bayesian Audiences
Yonathan A. Arbel
Murat C. Mungan
This paper can be downloaded without charge from the
Social Science Research Network Electronic Paper Collection
Electronic copy available at: https://ssrn.com/abstract=4181890

<<PAGE_BREAK>>

J,V0N0 1
Defamation with Bayesian Audiences
YonathanA.Arbel,* MuratC.Mungan†
Howstrictlyshouldthelawregulatefalsedefamatorystatements? Wefirstshowthat
thepresenceofjudicialerrorsoftenputsdefamationlawonaLaffercurve: regulation
thatistoolaxortoostrictisinferiortomoderateregulation.Whilemoderateregulation
is ideal, it is not always attainable, due to practical and legal constraints. With these
constraints,thepresenceofBayesianaudiencescancausetheoptimalregulationtobe
laxerthanisprescribedbystandardmodelswithna¨ıveaudiences.Thishighlightsthe
importanceofaccountingfortheimpactofdefamationlawsonbeliefformation.
Keywords:Defamation,Bayesianaudience,informationregulation,disclosure.
JELclassification:C72;D82;D83;K10;K13;K39.
1. Introduction
Whenstatementsaremadeinpublic,audiencesassesstheircredibilitybased
onavarietyofcues.Oneofthesecuesishowstrictlythelawsanctionsfalse
statements, i.e., whether talk is cheap. Such audience effects complicate the
standardanalysisofdefamationlaw,whichtraditionallyonlyfocusesonhow
the law affects speakers and the subjects of their speech. We investigate here
theoptimalstrictnessofdefamationlawwhenaccountingforaudienceeffects.
Defamation law imposes tort liability on speakers who publish false state-
ments that harm their target’s good name. A typical example comes from a
recentcasewhereahospitalstafferfalselyallegedthatadoctorwasworking
undertheinfluenceofalcohol.Thedoctorwonalawsuitagainstthestafferand
recoveredmillionsofdollarsindamages(Denmanv.St.Vincent,2020).While
the basic structure of defamation law is well established, there is an ongoing
socialdebateondefamationlaw’sproperscope.WithcallsfromtheSupreme
Court, legal scholars, politicians, and various pundits, there is growing pres-
suretodaytoreformdefamationlaw(Arbel&Mungan,2019).Inthemidstof
thesecalls,anewRestatementprojectwasrecentlyannounced.
Theliteratureondefamationlawisvast,buttheeconomicanalysisofdefama-
tion law is quite limited (Hemel, 2020). In deciding the level of strictness of
defamationlaw,standardlegalanalysesaredominatedbyatwo-sidedbalanc-
ing act. On the one hand, society considers the interests of the target of the
speech—her need for compensation and the need to protect her by deterring
defamatory speech against her. On the other hand, society also considers the
speaker,hisrighttofreespeech,andthesocialconcernwithchillingvaluable
speech(NYTimesv.Sullivan,1964).
*UniversityofAlabamaSchoolofLaw.E-mail:yarbel@law.ua.edu
†GeorgeMasonUniversity,AntoninScaliaLawSchool.E-mail:mmungan@gmu.edu.
Draft,Vol.0,No.0,
doi:/ewmxxx
©.
Allrightsreserved.ForPermissions,pleaseemail:
EElleeccttrroonniicc ccooppyy aavvaaiillaabbllee aatt:: hhttttppss::////ssssrrnn..ccoomm//aabbssttrraacctt==44118811889900

<<PAGE_BREAK>>

2 .V0N0
Inreality,wenoted,defamationlawalsoaffectspartiesbeyondthespeaker
andhistarget.Inparticular, defamationlawalsoaffectsaudiencesofspeech.
Thisisnotjustduetothefamiliarideathatstrictdefamationlawwouldlimit
thesupplyoffalsespeechthroughthedeterrenceofspeakers.Ifthatwerethe
case,theprotectionofaudienceswouldbeasimplematterofsettingsanctions
asstrictaslegallypossible.Defamationlawalsoaffectsaudiencesbychanging
theirassessmentofthecredibilityofspeechand,thus,itaffectstheaudience’s
propensitytoactuponstatements(Pennycooketal.2020, Arbel2022).Such
effectsaddcomplexitythatthestandardanalysisneglects.Theneglectofau-
dienceeffectsmaybebecauselawyersnaturallyfocusonthepartiesthattake
anactivepartinthelegalprocess–thevictimastheplaintiffandthespeakeras
thedefendant(Heymann,2012).Whateverthereason,tworecentinformalac-
counts(Hemel&Porat,2019&Arbel&Mungan,2019)suggestthatomission
ofaudienceeffectsisconsequentialtotheoptimalstrictnessofdefamationlaw.
Our object here is to bridge this gap by offering a general framework that
analyzesbehaviorandevaluateswelfarebasedonthestrictnessofdefamation
law. To do so, we construct a model that includes three key features: (i) a
Bayesian (rather than a na¨ıve) audience, (ii) errors in the court’s judgment
(wrongfulliabilityaswellaswrongfulfailuretofindliability),and(iii)acap
onrecoverabledamages.Weexplaintheroleeachofthesefeaturesplays,after
brieflyreviewingthestructureofourmodelanditsimplications.
Weconsiderthebehaviorofthreeparties.Aspeaker,whohasprivateinfor-
mation about a certain target – a business or an individual. The speaker may
make claims about the target to an audience member. The audience member
then decides whether to interact—trade, collaborate, socialize—with the tar-
get. Targets can be a high- or low-type, and audiences would rather interact
only with the former kind. If the target loses an interaction, he may bring a
lawsuitagainstthespeakerallegingdefamation.Westudybehaviorunderdif-
ferentlevelsofdamagesfordefamationandtheirsocialwelfareimplications.
Our analysis reveals three central findings. First, we find that there is an
optimallevelofdamagesthatsupportsseparatingequilibriainwhichwould-be
defamersaredeterredfromsharingfalseinformationandspeakersonlyshare
informationhonestlywiththeiraudience.Audiencemembersbelievespeakers
and act upon this information. Naturally, social welfare is highest under this
regime.
Second, we find that defamation laws often follow a ‘Laffer Curve.’ Lax
regulationresultsinafloodofcheaptalk,whichleadstoaudiencesdiscount-
ing all statements—true or false—and simply acting on their priors. This re-
sultsinequilibriawhereimportantinformationisleftuncommunicated.Sanc-
tioningdefamationtoostrictlyisalsounwanted,becausehighsanctionsinvite
frivolous litigation, which in turn chills true negative statements. In between
thesetwoextremes,theoptimallevelofdamagesfollowsaninverseUshape,
with a range of optimal damages. Thus, our unified framework shows that
bothcheaptalkand‘overpriced’talkareundesirableastheydepriveaudiences
ofrelevantinformationthatcouldbemadeavailabletothemundermoderate
damages.
EElleeccttrroonniicc ccooppyy aavvaaiillaabbllee aatt:: hhttttppss::////ssssrrnn..ccoomm//aabbssttrraacctt==44118811889900

<<PAGE_BREAK>>

-- 3
Our third result emerges when the cap on damages is lower than the level
necessary to support equilibria in which the target’s type is fully revealed. In
reality, it is difficult to calculate the exact level of damages, and even when
possible,damagesareboundedbybothconstitutionalconsiderationsandlim-
its on defendants’ wealth. In such cases, we find that a lax approach can be
superiortoamorestringentone.Thereasonisthatstringentregulationinvites
audience trust, but because some statements are false, this trust can be mis-
placed, leading to the deterrence of valuable interactions. Compounding the
issueisthatstringentregulationincreaseslitigation.Laxregulation,however,
invitesaudiencestorelymoreontheirpriorsandreduceslitigation.Thus,per-
hapscounter-intuitively,laxregulationbecomespreferabletostrictregulation
whenreputationalharmsarelarge,andtheoppositeconclusionmayholdwhen
reputationalharmsaresmall.
All three features of our model (Bayesian audience, judgment errors, and
damagecaps)playimportantrolesintheproductionoftheseresults.
First, when courts make no errors in judgment, people have no incentive
tobringfrivolousclaimsagainstspeakers.Inthiscaseverylargedamages(if
feasible) are always preferable to smaller damages, because they only deter
falsespeechwithouthavinganyimpactontruenegativespeech.Thisisady-
namicthatemergesinmanyothercontextsaswell,andhighlightstheroleof
judgment errors in explaining the inefficiency of very large damages in the
defamation context, and the emergence of the Laffer Curve to which we al-
luded.
Second, the prior economics literature on defamation law assumes that a
publisher(e.g.atabloid, journal,anindividualetc.)canalwaysharmanother
personbymakingnegativestatementsaboutthem,andtheextentofthisharm
isindependentofthelawsinplace(e.g., Garoupa1999a,b, Bar-Gill&Ham-
dani 2003).1 This is equivalent to the audience –whose beliefs and behavior
isnotconsideredinthepriorliterature–naivelyformingitsbeliefsandacting
uponthem.Thus,inthepriorliterature,themainfunctionofreformingthelaw
is to alter the expected costs and benefits of making disparaging statements,
but not the harmful impact of defamatory statements. With a na¨ıve audience,
increasingdamagesleadstoareductionintheexpectedharmtothetarget,be-
causeitdetersnegativespeech.ThisisnotsowhentheaudienceisBayesian.
Becauseverylowdamagesresultinfrequentfalseallegations,theydilutethe
informational content of speech, and therefore causes the audience to act ac-
cordingtoitspriors.Thus,inadditiontoprovidingstraightforwardrationales
for some behavioral responses in the defamation context (e.g., disregarding
certainfalsespeech),theincorporationofBayesianaudiencesalsohasimpor-
tantnormativeimplications, e.g., loweringdamagescanreducetheharmthat
resultsfromdefamatorystatements.
1. Themirrorimageofthisassumptionisalsoinvokedinthisliterature:thespeaker’sbenefit
frommakinganegativestatementisindependentoftheaudience’sbeliefs,becausetheaudienceis
notconsideredinthisliterature.Thisassumptionismade,forinstance,inDalviandRefalo(2008),
whichfocusesexclusivelyonthespeakers’incentivesandignoresnotonlytheaudience’sbeliefs
andbehaviorbutalsothetarget’s.
EElleeccttrroonniicc ccooppyy aavvaaiillaabbllee aatt:: hhttttppss::////ssssrrnn..ccoomm//aabbssttrraacctt==44118811889900

<<PAGE_BREAK>>

4 .V0N0
Third, this normative distinction becomes quite significant with binding
damagecaps,2 inwhichcasefalsedisparagingremarkscannotbefullyelimi-
nated.Thus,withBayesianaudiences,thechoiceisbetweenahighdegreeof
interactionsbetweentheaudienceandtargets(goodandbad)achievedthrough
lowdamages,andthemaximumlevelofdamagesthatcausesbadinteractions
tobedeterredalongwithsomegoodinteractions.Theformeroptionisprefer-
ablewhenthevalueofgoodinteractionsarelarge.Ontheotherhand, witha
na¨ıveaudience,maximumdamagesarealwayspreferable,becauseevenwith
low damages the audience believes false disparaging remarks, which are in
highsupplyduetothelackofdeterrentdamages.
In short, the main impact of judicial errors in our analysis is to rule out
theoptimalityofverylargedamages.Thisbecomesanimportantissuewhen
the damage cap is very large (or non-existent), in which case the presence
of judicial errors supplies an independent rationale for not having very large
damages.Ontheotherhand,whenthedamagecapisbinding,ana¨ıveaudience
impliesthatthemaximumdamageisoptimal,andthisresultisoverturnedwith
Bayesianaudiences.
While our analysis focuses on defamation law, the basic question we pose
here is relevant for a broad range of legal contexts. The law regulates false
speechindomainsasdiverseascorporatedisclosures,falseadvertising,whistle-
blowers, and law enforcement. Common to these domains is a basic tension
betweenthestrictnessofsanctionsformisreportingandtheinformativenessof
speech,andwecommentonpotentialimplications.
The next section offers some brief background and reviews the related lit-
erature.Section3presentsthemodelanditsanalysis.Section4evaluatesthe
welfare implications of different damages regimes, and highlights the impor-
tanceofaccountingforaudienceeffects.Section5containsseveralextensions
and discussions of the basic model, such as the public enforcement case, the
generalizationofthemodeltocaseswherespeakersmaybemotivatedtospeak
truthfullyortoexcessivelypraisethetarget,anddiscussionsofcontextsother
thandefamationlaw.Section6providesconcludingremarks.
2. LiteratureReview
Defamationlawregulatesthedisseminationoffalsestatementsthatare‘defam-
atory.’ To be defamatory, a statement must not only be false but also made
publicandbecapableofharmingone’sreputationandstandinginthecommu-
nity.Defamationlawisconsideredtobeabranchoftorts,anditencompasses
severaldistincttorts,mostnotablylibelandslander.Today,however,thedis-
tinctionhaslesspracticalsignificancethaninthepast,andinwhatfollows,we
abstractfromit.
Many defamation lawsuits are brought by individuals, but businesses and
firmscanalsobringsuit.Arecenthigh-profileexampleinvolvesalawsuitby
‘Dominion,’ a firm that sells voting hardware and software, against various
2. Asimilardynamicalsoemergeswhencourtsfrequentlymakejudgmenterrors,aswebriefly
explaininsection5.4.,below,andingreaterdetailinArbeland&Mungan2020.
EElleeccttrroonniicc ccooppyy aavvaaiillaabbllee aatt:: hhttttppss::////ssssrrnn..ccoomm//aabbssttrraacctt==44118811889900

<<PAGE_BREAK>>

-- 5
publicfiguresandmediaoutlets,whoallegeditwasinvolvedinthemanipula-
tionofelectionvotes.3
Defamation law evolved in the common and ecclesiastical courts of Eng-
land. In the United States, the states took the doctrine and used it to develop
theirownvariants.Amajordevelopmenttookplacein1964,whentheSupreme
Court decided the seminal case of NYTimes v. Sullivan. There, the Court re-
viewedtheexistingbodyofdoctrineinlightoftheFirstAmendmentprotec-
tionoffreespeechandpress.TheCourtmadeitconsiderablyharderforpublic
figures to bring lawsuits on matters of public interest. In the years that fol-
lowed,thedoctrinewasrefinedand,whilestillcarryingsignsofitsconvoluted
history,reachedacertaindegreeofbalance.Inrecentyears,however,therehas
beengrowingpressuretoreformthelaw.CommentsfromtheSupremeCourt
(McKee v. Cosby, 2019; Berrisha v. Lawson, 2021), the political sphere, le-
galcommentators, andpundits—allrevealdissatisfactionwiththelaw.Many
ofthesecommentssuggestthatdefamationlawshouldbemadestricter; e.g.,
CassSunsteincalledtheNYTimesv.Sullivandecision‘anachronistic’andar-
guedthatpublicfiguresshouldbeallowedtobringsuitmoreeasily(Sunstein,
2021). Interestingly, the reason why the law should protect good name inter-
estsisnotwellunderstood.Somegroundthelaw’sinterventioninaproperty
likeinterestingoodname,orgoodname’sbasisindignity,property,andhonor
(Post, 1984) while others relate it to concepts of social status and reputation
(Arbel,2021).
The legal literature on defamation law is rich and vast, and it explores a
varietyoftopics,involvingdeepquestionsofpoliticalphilosophyandconsti-
tutionalcommitments.Itisthereforequitesurprisingthattheliteratureonthe
economicsofdefamationlaw‘haslagged’andissparse(Hemel,2020).Some
notablecontributionsinthisspaceincludesRichardPosner’spioneeringanal-
ysis (Posner, 1973, 1986), which highlighted the applicability of cost-benefit
analysis to defamation law. More recent work focuses on the law’s effect on
media’sincentivestoinvestigateandreporttopicsofpublicinterest(Bar-Gill
&Hamdani,2003,Dalvi&Refalo,2008,Acheson&Wohlschlegel,2018)and
on political dishonesty (Garoupa, 1999a,b). As noted, this paper differs from
theseanalysesbyconsideringaBayesianaudience,alongsidedamagecapsand
judicialerrors.
Despitethesecontributions,courtsandlegalcommentatorsarelimitedtoa
fairlyrudimentaryunderstandingoftheincentivesfosteredbydifferentdefama-
tion law regimes. Here we amplify on two informal contributions that recog-
nizetherelevanceandimportanceofaudiencestotheanalysisofdefamation
law (Hemel & Porat, 2019, Arbel & Mungan, 2019). Methodologically, our
article borrows tools from the rich literature on signaling (Spence 1973) and
cheaptalk(Crawford&Sobel,1982).Ouranalysiscanalsobeinterpretedas
part of emerging literature that looks at how laws can be used to create in-
formalsanctionsthroughthebehaviorofthirdparties(e.g.,Deffains&Fluet,
3. USDominion,Inc.v.FoxNewsNetwork,LLC,C.A.N21C-03-257EMD(Del.Super.Ct.
Dec.16,2021).
EElleeccttrroonniicc ccooppyy aavvaaiillaabbllee aatt:: hhttttppss::////ssssrrnn..ccoomm//aabbssttrraacctt==44118811889900

<<PAGE_BREAK>>

6 .V0N0
2019,Mungan2016,Be´nabou&Tirole,2006,2011,Rasmusen1996).
3. Model
We model the interactions between three parties: the speaker (S, she), the
target of the speech (T, he), and the audience, captured by a representative
member (A, it). A faces an informational problem: T is either a good or a
bad type, and A’s value of interacting with T depends on T’s type, which
is unknown to A. Before A decides whether to interact with the target, S,
who knows T’s type, communicates with A and may either disparage T or
makeanon-disparagingcomment.Becausewestudydefamation,weconsider
thepossibilitythatS mayfalselydisparageT inordertodeteraninteraction
betweenAandT.Wedeferthediscussionofspeakersbeing(atleastpartly)
motivatedbyadesiretotruthfullyshareinformation,asthishaslimitedimpact
onouranalysis.4 WemodeltheinteractionsasaBayesiangame,anduseitto
identifyPerfectBayesianEquilibria.5
3.1 Preliminaries
Thetarget,T,obtainsabenefitofr fromtheinteraction,wheret∈{B,G}
t
denotes his type and where the letters abbreviate bad and good, respectively.
T’s type is privately known to himself and S, but not to A, who only knows
thattheproportionofgoodtypesisγ ∈(0,1).6Apreferstointeractwithgood
types,butnotbadtypes,becausethisresultsinautilityofg >0>−bwhere
bisthedisutilityAbearsfrominteractingwithabadtype.Thus,absentfurther
information,AwouldprefertointeractwithT ifγg−(1−γ)b > 0,andwe
assumethisinequalityholds,sinceotherwisenointeractionswouldtakeplace
between A and T even without (negative) input from S.7 Thus, the audience
preferstointeractwiththetargetifitsupdatedbelief(basedonthestatementit
receivesfromS)ofT’slikelihoodofbeingagoodtypeexceedsthethreshold
b
x≡ <γ, (1)
(cid:98) g+b
wheretheinequalityfollowsfromtheassumptionthatAwouldprefertointer-
actwithT absentinputfromS.
ThespeakerhasaninterestinwhetherAandT interact:Sobtainsagainofv
whenAavoidsinteraction(alternatively,vcanbeinterpretedasalossincurred
whenAchoosestointeractwithT).v isarandomvariablewiththecumula-
tive distribution function F(v) with support (0,1] where the upper-bound of
the support is normalized to simplify notation. The specific v-draw is private
4. Consistentlywiththelaw,truthfulnegativestatementsarenotconsidereddefamatory.How-
ever,thecourtmaymakeerrorsinascertainingwhetheranegativestatementistruthful,andthis
possibilityisincorporatedinourmodel,asweexplainbelow.
5. Figure5intheAppendixdepictstheinteractionsbetweenthethreepartiesandishelpfulin
followingthedetaileddescriptionsoftheinteractionsthatweprovide,next.
6. Insection5wediscusstheconsequencesofendogenizingγ.
7. Ananalysisofthiscasecanbefoundinanearlierversionofthisarticle, andyieldsno
furtherinsights(seeArbel&Mungan(2020)).
EElleeccttrroonniicc ccooppyy aavvaaiillaabbllee aatt:: hhttttppss::////ssssrrnn..ccoomm//aabbssttrraacctt==44118811889900

<<PAGE_BREAK>>

-- 7
informationavailableonlytoS,andwecallv thespeaker’stype.Weassume
thatinteractionsbetweenAandT aresociallyvaluableif,andonlyif,T isa
goodtype,i.e.r +g >1>0>r −b.
g b
After Nature determines the types of T and S, the target’s type becomes
commonknowledgeamongT andS(butnotA).Atthispoint,Schooseswhat
type of statement to send A regarding T’s type. The types of possible state-
ments follow defamation law’s distinction between disparaging statements,
which are potentially actionable, and non-disparaging statements, which are
non-actionable(e.g.,positiveremarks,silence,opinion,etc.).
Subsequently, A decides on whether to interact with T or to avoid him,
and, finally, T, decides whether to bring a lawsuit against S if a disparaging
remarkwasfollowedbyA’schoicetoavoidinteractingwithT.8 Wenotethat
this setting includes the possibility of T suing S, even if T is in fact a bad
type,i.e.,afrivolouslawsuitmaybebrought.Thisisanimportantpossibility
becausecourtsmayerrintheirjudgment.Tocapturetheparties’payoffsfrom
litigation,wedefinethefollowing:
d: damagespaidbyS toT whenthecourtfindsforT
l: totallitigationcosts.Weassumethatlitigationcostsarenotprohibitive(l<1)
and,withoutlossofgenerality,thatthecostsareequallysharedbytheparties.
q : probabilityofplaintiffvictorywhenT isoftypet∈{B,G}
t
(cid:16) (cid:17)
Weassumetheprobabilityofwrongfulliabilityissmall: q <q l/2
B G 1−l/2
Becausecourtswieldbroadlatitudeinsettingthelevelofdamages,9wefol-
low the existing literature (see, e.g., Garoupa 1999a,b) and use d as a policy
lever to operationalize different defamation regimes. The value of d can also
be interpreted as expected damages, i.e. d = E[d˜] where d˜is a random vari-
ablerepresentingactualdamagesawardedwhosedistributionisaffectedbythe
court’sinterpretationofthelaw.Wealsonotethatdcanbechosenseparately
for each type of statement (represented by all variables described above, r ,
g
r ,g,b,etc.)beinganalyzed,whichwetakeasgiveninouranalysis.
b
It is worth noting that we take the odds of winning at trial, q and q ,
B G
as exogenously given. This is a standard commitment assumption in the en-
forcementliterature,andimpliesthatcourtsarecommittedtoreviewingcases
onlyontheirmerit,i.e.,withoutbringingintheirinformedestimatesaboutthe
proportion of frivolous cases. Our assumption relating q to q , above, cor-
G B
respondstothecasewherejudicialerrorsarepossible,butoccurinfrequently.
Wenotethattheanalysisofthealternativecasewithlargejudicialerrorsleads
8. Inpractice,courtsoftenrequiresomeproofofharmtoallowmonetaryrecovery,hencethe
requirementherethatthedefamatoryremarkpreventedaninteraction.
9. Courtscan,withinlimitations,awardnominal,economic,non-economic,andpunitivedam-
ages, andhavedemonstratedconsiderablediscretioninpractice.SomeexamplesincludeLoth-
schuetzv.Carpenter,898F.2d1200,1205(6thCir.1990)($1innominaldamages);WasteMgmt.
ofTexas,Inc.v.TexasDisposalSys.Landfill,Inc.,434S.W.3d142,162(Tex.2014)($450,000
ineconomicharms); Cantuv.Flanigan,705F.Supp.2d220(E.D.N.Y.2010)($150millionin
non-economicdamages);Armstrongv.Shirvell,596F.App’x433,448(6thCir.2015)($500,000
inpunitivedamages).
EElleeccttrroonniicc ccooppyy aavvaaiillaabbllee aatt:: hhttttppss::////ssssrrnn..ccoomm//aabbssttrraacctt==44118811889900

<<PAGE_BREAK>>

8 .V0N0
tosimilarresults,andacompleteanalysisofthiscasecanbefoundinArbel&
Mungan(2020).
3.2 Players’Actions,Beliefs,Strategies,andPayoffs
Next,wedescribetheplayersstrategies,beliefs,andactions.Forsimplicity,
eachplayer’sactionislabelledaseither0or1,asfollows:
Table1: Players’PotentialActions
Player Action
0 1
S Don’tDisparage Disparage
A Interact Avoid
T Don’tLitigate Litigate
We note that labeling A’s action of interacting as 0 may appear counter-
intuitive.However,thebenefitofthisnotationisthatasuitisfiledonlyincases
where all players’ actions are 1. This makes it simpler to express the pay-off
ofthespeaker(asinTable3,below),sinceshefacesexpectedlitigationcosts
onlywhenallactionsequal1.
Usingthisnotation,wecandescribethestrategiesofeachplayerasfollows:
Table2: Players’Strategies
Player Strategy
S s(t,v):{B,G}×(0,1]→{0,1}
A a(z):{0,1}→{0,1}
T p(t):{B,G}→{0,1}
Here,inspecifyingA’sstrategy,zdenotesthestatementreceivedbyA.
BecauseoursolutionconceptisaPerfectBayesianEquilibrium(henceforth
PBE),wealsospecifyA’sbeliefsregardingT’stype,as:10
x : BeliefthatT isagoodtypegivenz =i
i
Withthisnotation,weexpresstheexpectedpay-offsofeachplayer,giventheir
beliefsandinformation,asfollows:
Table3: Players’Payoffs
Player Payoff
S a(s(t,v))(v−p(t)s(t,v)(q d+ l))
t 2
A a(z)(x g−(1−x )b)
z z
T (1−a(s(t,v)))p(t)(q d− l)+a(s(t,v))r
t 2 t
10. BecauseA’svaluationofhisinteractionwithT dependsonlyonT’stype,weneednot
specifyA’sbeliefsregardingS’stypeforpurposesofidentifyingthePBE.
EElleeccttrroonniicc ccooppyy aavvaaiillaabbllee aatt:: hhttttppss::////ssssrrnn..ccoomm//aabbssttrraacctt==44118811889900

<<PAGE_BREAK>>

-- 9
3.3 EffectiveandIneffectiveCommunicationEquilibria
Perfect Bayesian Equilibria consist of assessments (i.e. a profile of beliefs
and strategies) that satisfy sequential rationality and consistency of beliefs.
SincetherequirementsforPBEarewellknown,werelegatetheirformaldef-
initions to Appendix A, below. As in many other contexts, communications
canbedisregardedbytheaudienceinsomeequilibria.Wedistinguishbetween
theseandothertypesofequilibriabyusingthefollowingdefinition.
Definition1APBEisaneffectivecommunicationequilibriumif,andonly
if,theaudiencechoosesnottointeractwiththetargetwithsomepositiveprob-
abilitybasedontheinformationitreceivesfromthespeaker.
We start by noting that defamation law cannot eliminate ineffective com-
munication equilibria. This is because when the audience’s beliefs regarding
the target’s types are unconditional on the speaker’s statement and equal to
its prior (i.e. x = x = γ), it chooses to interact with the target regardless
0 1
of what it hears from the target (i.e. a∗(z) = 0). This results in payoffs of 0
andrtothespeakerandtarget,respectively.Thesepayoffsareindependentof
the actions of the speaker and target, which makes them indifferent between
playing any of the strategies available to them. Thus, any assessment where
the speaker plays a strategy that supports the audience’s beliefs constitutes a
PBE.Thesimplestexampleisonewherethespeakerneverchoosestodispar-
agethetarget(i.e.s(t,v)=0foralltandv).Weformalizethisobservationas
follows.
Proposition 1. Under all defamation regimes, there exist ineffective com-
municationequilibria.
Proof. The assessment consisting of x∗ = x∗ = γ, a∗(z) = 0, s∗(t,v) =
(cid:26) 0 if q d⩽l/2 1 fort 0 ∈{B,G}
0, and p∗(t) = t satisfies sequential
1 if q d⩾l/2 fort∈{B,G}
t
rationality and consistency of beliefs (i.e. requirements 1-4 in Appendix A),
andthusisaPBE.
Proposition 1 notes that ineffective communication equilibria are always
present, regardless of the defamation regime in place. If these were the only
equilbria, defamation law would be irrelevant. Thus, we proceed by showing
thatsomelevelsofdamagesinfactgenerateeffectivecommunicationequilib-
ria.
Proposition 2. (i) Extremely low damages (i.e. d < l ) and extremely
2qG
highdamages(i.e.d> 2−l)onlygenerateineffectivecommunicationequilib-
2qB
(cid:16) (cid:17)
ria. (ii) There exist a range of moderate damages, D ⊂ l , 2−l , which
2qG 2qB
generate effective communication equilibria. (iii) The audience acts consis-
tentlywiththespeaker’sstatement,i.e.a∗(z)=z,inalleffectivecommunica-
tionequilibria.
Proof. SeeAppendix.
EElleeccttrroonniicc ccooppyy aavvaaiillaabbllee aatt:: hhttttppss::////ssssrrnn..ccoomm//aabbssttrraacctt==44118811889900

<<PAGE_BREAK>>

10 .V0N0
The intuition behind the first part of proposition 2 is relatively straightfor-
ward.Whendamagesareextremelylow,thetargetisdeterredfromsuingthe
speaker, even when he has a meritorious case, since expected damages (i.e.
q d)arelowerthanlitigationcosts.Thiscausesthespeaker’sstatementstobe
G
perceived as cheap-talk by the audience, since the speaker faces no negative
consequence from making disparaging statements. Thus, the audience disre-
gards the speaker’s statements and acts according to its priors. On the other
hand,whendamagesareextremelyhigh,alltargetsareincentivizedtolitigate,
andexpecteddamagesarehighenoughtodeterallspeakertypesfrommaking
disparaging statements. Thus, the audience is once again left without any in-
formativestatements,thistimeduetotheover-pricingofspeechasopposedto
thepresenceofcheap-talk.
Itisonlymoderatedamagesthatsupporteffectivecommunicationsbetween
speakers and audience members, and this is formalized in proposition 2-(ii).
Part (iii) of proposition 2 simply rules out the possibility of counter-intuitive
equilibria,forinstance,inwhichtheaudienceinfersfromadisparagingremark
thatthetargetmustbeagoodtypeandvice-versa(i.e.wherea∗(z)=1−z).
These preliminary findings indicate that if defamation laws are to have any
impact, theymustdosothrougheffectivecommunicationequilibriaobtained
under moderate damages. Thus, we analyze these equilibria in further detail,
next.
3.4 ModerateDamagesandEffectiveCommunicationEquilibria
The damages in place affect the target’s incentive to sue when disparaged,
aswellasthespeaker’sincentivestodisparagethetargetinthefirstplace.We
notetwopairsofcriticaldamagesthatpertaintoeachparty’sincentives.First,
l l
d ≡ andd ≡ (2)
1 2q 3 2q
G B
are the smallest damages that causes a type G and B target, respectively, to
bring suit whenever the speaker disparages him.11 On the other hand, when
damagesaregreaterthan
2−l 2−l
d ≡ andd ≡ (3)
2 2q 4 2q
G B
typeGandBtargets,respectively,areexpectedtobringsuit,andthisdetersthe
speakerfromdefamingthetarget.Ourassumptionofnon-prohibitivelitigation
costsandsmalljudicialerrorsimpliesthatthesefourcriticaldamagelevelsare
orderedasfollows:
d <d <d <d (4)
1 2 3 4
Thus,therearethreecategoriesofmoderatedamages,whichwecalllow,inter-
mediate,andhighdamages.Weexplaintheincentivesthateachplayerfacesin
11. Weassume,onlytosimplifyexposition,thatanindifferenttargetchoosesnottolitigate.
EElleeccttrroonniicc ccooppyy aavvaaiillaabbllee aatt:: hhttttppss::////ssssrrnn..ccoomm//aabbssttrraacctt==44118811889900

<<PAGE_BREAK>>

-- 11
eachofthesedamagecategoriesunderaneffectivecommunicationequilibrium
(whenoneexists).
LowDamages(d∈(d ,d ))
1 2
In this range, the target has the incentive to litigate only if he is type G,
sinceq d > l > q d.Thus,inaneffectivecommunicationequilibrium,the
G 2 B
speaker faces no threat of litigation from disparaging a bad type, and thus a
type B target is disparaged with certainty. On the other hand, if the speaker
encountersatypeGtarget,sheexpectsthatdisparaginghimwillleadtoacost
of
l
v (d)≡q d+ . (5)
G G 2
Thus,thespeakerchoosestodisparageatypeGtargetifhertypeexceedsthis
value.Therefore,atypeGtargetisdisparagedwithaprobabilityof
1−F(v (d)) (6)
G
Ouranalysisthusfaridentifiesthebehaviorofthespeakerandtargetinan
effectivecommunicationequilibrium,assumingthatitexists.But,forthistype
ofequilibriumtobesupportable,theaudience’sbeliefsmustbeconsistentwith
theequilibriumbehavioroftheotherparties.Thus,theaudiencemustholdthe
belief that a target who is not disparaged must be a good type, since all bad
typesaredisparaged,i.e.
x∗ =P(t=G|z =0)=1
0
Ontheotherhand,whentheaudiencereceivesadisparagingstatement,itmust
believethatthetargetisneverthelessagoodtypewithaprobabilityof
γ[1−F(v (d))]
x∗(d)=P(t=G|z =1)= G <γ (7)
1 γ[1−F(v (d))]+(1−γ)
G
This is because a type G target is disparaged with probability 1−F(v (d))
G
whereas a type B individual is disparaged with certainty, and the likelihood
withwhichthetargetisagoodtypeisγ.
Aswenotedvia(1)theaudiencefindsitinitsbestinteresttointeractwith
thetargetifitbelievesheisagoodtypewithaprobabilityexceedingxˆ.Thus,
aneffectivecommunicationequilibriumissupportableinthisrangeif
x∗(d)<xˆ<x∗ =1 (8)
1 0
In the upper portion of this range, this condition certainly holds, since 1 −
F(v (d ))=0,i.e.thespeakerisalwaysdeterredfromdisparagingatypeG
G 2
target.Moreover,x∗(d)isclearlydecreasinginthisdamagerange,whichcan
1
easily be verified by inspecting (7). Therefore, within low damages, there is
some critical damage dˆ, above which effective communication equilibria are
obtained.Whetherdˆ=d ,i.e.whethereffectivecommunicationequilibriaare
1
supportable with all lower moderate damages, depends on parametric values
andhasnoimportantimplications.
EElleeccttrroonniicc ccooppyy aavvaaiillaabbllee aatt:: hhttttppss::////ssssrrnn..ccoomm//aabbssttrraacctt==44118811889900

<<PAGE_BREAK>>

12 .V0N0
Based on these observations, we can summarize the impacts of increasing
damagesintherange(dˆ,d )onthebehaviorofallplayersineffectivecommu-
2
nication equilibria. As damages are increased, the speaker disparages type G
individualslessfrequently,sincethethresholdspeechbenefitthatsherequires
isincreasingindamagesper(5).Thisleadstolessfrequentlitigationaswell
as less frequent blocking of beneficial interactions between the audience and
thetypeGtarget.Thelevelofdamagesleadstonofurthereffects, becausea
typeBtargetisdisparagedwithcertaintyinthisrange.
IntermediateDamages(d∈[d ,d ])
2 3
When damages are increased into the intermediate range the speaker is al-
waysdeterredfromdisparagingatypeGtarget.Thisisbecausetheexpected
damages and litigation costs associated with doing so exceed the benefit that
she obtains from blocking the target’s interaction with the audience. More-
over,becaused<d ,atypeBtargetlackstheincentivestolitigate.Thus,the
3
speakeralwaysdisparagesatypeB target.Basedonthisbehavior, itfollows
thatinaneffectivecommunicationequilibriumtheaudience’sbeliefsaregiven
by
x∗=P(t=G|z =0)=1and
0
x∗=P(t=G|z =1)=0
1
This trivially implies that effective communication equilibria are supportable
byalldamagesinthisrange,since0=x∗ <xˆ<x∗ =1.
1 0
Quite importantly, in this range all effective communication equilibria are
separatingequilibriawhereinthespeakerchoosesto[not]disparagewhenever
sheknowsthatthetargetistypeB [G].Welaternotethesuperiorityofthese
equilibriawhenweconductawelfareanalysis.
HighDamages(d∈(d ,d ))
3 4
With high damages, a type B target is given the incentive to frivolously
litigate. This, in turn, causes the speaker to be deterred from disparaging a
typeB targetwhenherbenefitfromblockinginteractionsislow.Specifically,
shechoosestodisparageatypeBindividualonlyifherbenefitexceeds
l
v (d)≡q d+ (9)
B B 2
Thus,abadtypeisdisparagedwithaprobabilityof
1−F(v (d))<1 (10)
B
On the other hand, when the speaker encounters a type G target, she never
disparages him given the high damages. Based on this behavioral profile, in
equilibriumtheaudience’sbeliefsaregivenby
γ
x∗(d)=P(t=G|z =0)= >γ and
0 γ+(1−γ)F(v (d))
B
x∗=P(t=G|z =1)=0
1
EElleeccttrroonniicc ccooppyy aavvaaiillaabbllee aatt:: hhttttppss::////ssssrrnn..ccoomm//aabbssttrraacctt==44118811889900

<<PAGE_BREAK>>

-- 13
Figure1
Sincexˆ<γ,itfollowsthatx∗ >xˆ>x∗,andthereforeeffectivecommunica-
0 1
tionequilibriaaresupportablebyallhighdamages.
As the above discussion indicates, increasing damages in this range only
reduces the frequency with which speakers disparage a type B target. Thus,
increasing damages in this range has countervailing effects: it increases the
frequencyofinteractionswithbadtypesbutreducesthefrequencyoflitigation.
Wesummarizeourfindingsinthissubsectionthroughfigure1,below,which
depictsthequalitativerelationshipbetweendamagesandthebeliefsofA;the
likelihood with which a type t ∈ {B,G} target is disparaged; and the likeli-
hood of litigation in effective communication equilibria. Next, we conduct a
welfareanalysiswhichbuildsonthesekeyfindings.
EElleeccttrroonniicc ccooppyy aavvaaiillaabbllee aatt:: hhttttppss::////ssssrrnn..ccoomm//aabbssttrraacctt==44118811889900

<<PAGE_BREAK>>

14 .V0N0
4. WelfareAnalysis
Inanalyzingthesocialdesirabilityofdifferentdefamationregimes,weuse
a simple social welfare function which consists of the sum of each player’s
expectedpay-off.Weconductthiswelfareanalysistohighlightthethreemain
pointsthatweemphasizedintheintroduction.First,welfareisnon-monotonic
in damages for defamation. More precisely, we show that the shape of wel-
fare obtained through effective communication equilibria vis- a`-vis damages
resemblesaLafferCurve:welfareisincreasinginthelowerrangeofmoderate
damages (i.e. for d ∈ (dˆ,d )), is maximized in the intermediate range (i.e.
2
when d ∈ [d ,d ]), and is decreasing in the upper range of moderate dam-
2 3
ages (i.e. when d ∈ (d ,d )). Second, when there is a cap on damages (e.g.
3 4
reflecting the wealth of the defendant or a legal bound on permissible dam-
ages), then all effective defamation remedies may reduce welfare. Third, we
contrasttheimplicationsofamodelwithaBayesianversusana¨ıveaudience.
Whentheaudienceisna¨ıveandeasilymisledbyfalsestatements,typeGtar-
gets always prefer stricter defamation laws. The same is not true when the
audienceisBayesian,becauseabsentsizeabledamagestheaudienceperceives
the speaker’s disparaging statements as cheap-talk and disregards them. This
insight leads to a divergence between the normative implications of the two
models:withbindingcapsondamages,itmaybeoptimaltohavenodefama-
tionlawsatallwithaBayesianaudiencebutoptimaltohavemaximaldamages
withana¨ıveaudience.Next,weconsiderandformalizeeachofthesepoints.
4.1 TheLafferCurveofDefamationLaw
Underineffectivecommunicationequilibria,theaudienceactsaccordingto
its priors. Thus, it chooses to interact with the target regardless of the state-
ments by the speaker. There is no litigation since interaction always takes
place.Thus,expectedwelfareisindependentofdamages,andisgivenby
W ≡(1−γ)[r −b]+γ[r +g] (1)
I B G
On the other hand, under effective communication equilibria, the specific
functionalformofwelfarediffersdependingonwhichofthethreerangesdam-
agesarein,asexplainedtheprevioussection.Next,weconsiderwelfareunder
eachrange.
Withlowdamagesthatsupporteffectivecommunicationequilibria(i.e.d∈
(dˆ,d )), when the target is type B, the speaker disparages him, the audience
2
refuses to interact, and the target chooses not to litigate. Thus, with a proba-
bilityof(1−γ)expectedwelfareequalsthespeaker’sexpectedbenefitE[v].
WhenthetargetistypeG,thespeakerdisparageshimonlywhenv > v (d).
G
In those cases, the audience avoids an interaction with T, and the target liti-
gates.Thus,withaprobabilityofγ,expectedwelfareisF(v (d))[r +g]+
G G
(cid:82)1
(v−l)f(v)dv.Therefore,expectedwelfareisgivenby:
vG(d)
(cid:32) (cid:33)
(cid:90) 1
W (d)≡(1−γ)E[v]+γ F(v (d))[r +g]+ (v−l)f(v)dv (2)
L G G
vG(d)
EElleeccttrroonniicc ccooppyy aavvaaiillaabbllee aatt:: hhttttppss::////ssssrrnn..ccoomm//aabbssttrraacctt==44118811889900

<<PAGE_BREAK>>

-- 15
Whendamagesareintermediate(i.e.d∈[d ,d ]),effectivecommunication
2 3
leadstoseparatingequilibriawhereininteractionstakeplaceif,andonlyif,the
targetisagoodtype.Moreover,thereisnolitigationsincetypeB targetslack
theincentivestolitigate.Thus,welfareinthisrangeisgivenby
W ≡(1−γ)E[v]+γ[r +g] (3)
S G
Finally,whendamagesareintheuppermoderaterange(i.e.d ∈ (d ,d )),
3 4
thespeakerchoosesnottodisparageatypeGtarget.Thus,withaprobability
of γ, welfare is r + g. When the target is type B, the speaker chooses to
G
disparagehimonlywhenv > v (d),andthisleadstolitigation.Thus,witha
B
(cid:82)1
probabilityof1−γ,expectedwelfareisF(v )[r −b]+ (v−l)f(v)dv.
B B vB(d)
Thus,expectedwelfareis
(cid:32) (cid:33)
(cid:90) 1
W (d)≡(1−γ) F(v (d))[r −b]+ (v−l)f(v)dv +γ[r +g](4)
H B B G
vB(d)
A very simple yet important observation is that W is increasing whereas
L
W isdecreasingindamages.Thisisbecause,whendamagesareinthelower
H
moderaterange, theimpactofincreasingdamagesistoreducethelikelihood
ofdefamatorystatementsagainstatypeGtarget.Thisisbeneficial,becauseit
reducesthelikelihoodofblockedbeneficialinteractionsbetweenAandT as
well as wasteful litigation between T and S. Similarly, when damages are in
theuppermoderaterange,loweringdamagesleadstoanincreaseinthelikeli-
hoodwithwhichatypeBtargetisdisparaged.Thisincreasesthelikelihoodof
harmfulinteractionsbeingblocked,butattheexpenseofincreasedlitigation.
The former (beneficial) effect dominates the latter (detrimental) effect, since
thespeakerdisparagesatypeB targetonlyifherbenefitsfromdoingsomore
than off-set total litigation costs. This last point can be formalized by noting
that
W′ =(1−γ)f(v )v′ {(r −b)−(v −l)}⩽0 (5)
H B B B B
sincer <bandv (d)⩾v (d )=l.
B B B 3
To summarize these observations, we depict welfare as a function of dam-
agesinfigure2,below,whichisaccompaniedbyclarifyingnotes.12
The most striking feature of figure 2 is that welfare obtained under effec-
tivecommunicationequilibriamimicsaninverseUshape,notwithstandingthe
factsthatitcontainsadiscretedrop13 aroundd = d andthatwelfareiscon-
3
stant in the [d ,d ] range. This is what we interpret as the Laffer curve of
2 3
defamationlaw.Welfareisnotmonotonicallyincreasingindamages,because
12. Wenotethateffectivecommunicationequilibriamayormaynotbesupportableforalllow
damages,asexplainedintheprevioussection.Figure2depicts,onlyforexpositionalpurposes,a
casewherealllowdamagessupporteffectivecommunicationequilibria.
13. Once damages pass over to the upper moderate range, damages are sufficient to invite
firivilouslitigationbytypeBtargets.Thisleadstoadiscretereductioninwelfare,sinceitcauses
thelikelihoodofdisparagingstatementsagainsttypeBindividualstodiscretelydropfrom1to
1−F(vB(d))anditalsocausesadiscretejumpinthelikelihoodoflitigation(seefigure1).
EElleeccttrroonniicc ccooppyy aavvaaiillaabbllee aatt:: hhttttppss::////ssssrrnn..ccoomm//aabbssttrraacctt==44118811889900

<<PAGE_BREAK>>

16 .V0N0
Figure2
large damages (i.e. d > d ) reduce welfare by deterring accurate negative
3
speech against a type B target. On the other hand, reducing damages to low
levels (i.e. d < d ) is also detrimental because it leads to defamatory state-
2
mentsagainstatypeGindividual,whicharetakenseriouslybytheaudience.
Anotherfeatureofthewelfarecurvedepictedinfigure3isthattheseparat-
ing equilibria obtained through intermediate damages lead to greater welfare
thanineffectivecommunicationequilibria.Wenotethatthisisnocoincidence,
andoccursduetothefactthatwhiletheaudienceinteractswitheithertypein
ineffectiveequilibria,itinteractswithatargetif,andonlyif,itisagoodtype
when damages are in the intermediate moderate range. Thus, when there are
nocapsondamagesitfollowsthatsettingdamagesintheintermediatemoder-
aterangeissociallydesirable.Weformalizethisresultthroughthefollowing
proposition,whoseprooffollowsfromourcomments,above.
Proposition3. SeparatingequilibriawhereSchoosestodisparageT if,and
onlyif,heistypeB leadtogreaterexpectedwelfarethananyotherequilibria
andareobtainableonlythroughintermediatedamages.
Animplicationofproposition3isthatmaximizingwelfarethroughdefama-
tion law requires the implementation of separating equilibria through the use
of intermediate damages. We note that this implication is obtained under the
assumption that these damages are feasible. However, when these damages
are too large for the defendant to pay (i.e. when the defendant is judgement-
proof)orwhentherearelegalrestrictions(e.g.constitutional)onthedamages
EElleeccttrroonniicc ccooppyy aavvaaiillaabbllee aatt:: hhttttppss::////ssssrrnn..ccoomm//aabbssttrraacctt==44118811889900

<<PAGE_BREAK>>

-- 17
thatcanbechosen,damageslargeenoughtosupportequilibriainwhichtarget
typesarefullyrevealedmaynotbeavailable.Wediscussthiscase,next.
4.2 BoundedDamages
Weused¯todenotetheupperboundondamages.Animmediateimplication
ofproposition3isthatwhend¯⩾d ,theupperboundisnon-binding.Thus,we
2
focusonthecasewhered¯< d .Withthisrestrictioninplace, thereareonly
2
tworelevantrangesofdamagesthatonecanselectfrom;(i)verylowdamages
which only support ineffective communication equilibria (d ⩽ d ), and (ii)
1
damagesinthelowerrange(d∈(d ,d¯)).
1
In the latter range, if d¯is close to d it is possible for effective communi-
1
cationequilibriatobeunsupportable,14 andtheanalysisofthiscaseistrivial:
thereisnofeasiblelevelofdamagesthatcanresultineffectivecommunication
equilibria, and hence the choice of damages is irrelevant. Thus, we focus on
the more interesting case where maximum damages are sufficient to support
someeffectivecommunicationequilibria,i.e.x∗(d¯)<xˆ.
1
Inthiscase,wecancomparethewelfarebenefitsandcostsassociatedwith
ineffective and effective communication equilibria. The former equilibria, as
reflectedby(1), alwaysleadtointeractionsbetweenT andA.Theseinterac-
tionsarewelfarereducingwhenthetargetistypeB.Hence, whenthetarget
istypeB,effectivecommunicationequilibriaperformbetter,sincetheydeter
interactionsbetweenT andA.However,thiscomesatthecostofdeterringso-
ciallybeneficialinteractionsbetweenatypeGtargetandAwhenthespeaker
has a sufficiently high type (i.e. v > v ). This implies that a switch from an
G
ineffective communication equilibrium to an effective communication equi-
librium trades-off deterrence of good interactions against deterrence of bad
interactions. Therefore, when the harm to the target from defamatory state-
mentsislargerelativetootherconsiderations,arathercounter-intuitiveresult
is obtained. Even when it is possible to implement effective communication
throughdefamationlaws,itissociallymoredesirablenottodoso.Thishap-
pensbecauseputtingapriceonspeechlendsmorecredibilitytothespeaker’s
statements,whichshecanthenusetoinefficientlyblockagoodinteraction.In
such cases, the superior option is to not make speech credible and cause the
audiencetorelyonitspriors,whichcausesittointeractwiththetarget.
Weformalizethisresultviaproposition4below,andweprovideanexample
anditsgraphicaldepictionviafigure4toillustrateit.
Proposition4. Supposetherearebindingmaximumdamages(i.e.d¯<d ).
2
Then, given all other parameters, there exists a threshold harm from effec-
tive defamation, r¯ , such that ineffective communication equilibria lead to
G
(weakly) higher welfare than all effective communication equilibria if, and
onlyif,r ⩾r¯ .
G G
14. Thisoccurswhenmaximumdamagesarenotenoughtodeterspeechagainstgoodtypes
frequentlyenoughinaneffectivecommunicationequilibrium,asnotedvia(7)-(8)andtheaccom-
panyingdiscussion.
EElleeccttrroonniicc ccooppyy aavvaaiillaabbllee aatt:: hhttttppss::////ssssrrnn..ccoomm//aabbssttrraacctt==44118811889900

<<PAGE_BREAK>>

18 .V0N0
Proof. Using(1)and(2),wecanexpressthedifferenceinbetweenineffective
andeffectivecommunicationequilibriaas:
(cid:90) 1
W −W =(1−γ)[r −b−E[v]]+γ [r +g−v+l]f(v)dv
I L B G
vG(d)
Thisexpressionisincreasingandunboundedinr .Thus,thereexistsr¯ such
G G
thatW ⩾W iffr ⩾r¯ .
I L G G
In figure 3, below, we depict multiple cases which illustrate the rationale
behindproposition4.Inthisexample,visdistributeduniformlyandr¯ =1.7
G
isusedtoillustrateallthreepossibilities.15
Asthefigureillustrates,thegapbetweenwelfareobtainedthroughlowdam-
ages under the two types of equilibria is decreasing in damages but increas-
ingtheharmthatthetargetsuffersfromeffectivedefamation.Thus,forsmall
defamation harms to the target, effective communication equilibria obtained
through maximum damages are superior, and the opposite conclusion holds
forlargedefamationharms.Theexceptionalcasewherethetwotypesofequi-
librialeadtothesameamountofwelfarewhenmaximumdamagesareusedis
alsodepictedasanintermediatecase(i.e.thecasewherer =r¯ ).
G G
Theseobservationsimplythatwhendamagesinducingcompleterevelation
of target types are not feasible, it is socially desirable to strive for effective
communication equilibria only when the harms from defamation are small.
Thisresultappearscounter-intuitive,becauseitsuggeststhattheoptimalityof
effective defamation remedies ought to be inversely related to the size of the
alleged harm to the plaintiff. The rationale behind this result is that making
speechcredibleinanenvironmentwheredefamatoryspeechcannotbelargely
eliminated has the function of making some false speech credible, and thus
harmful to type G targets. When the size of the harm to these individuals is
large,itnaturallybecomesmoredesirabletotakeawaythecredibilityofneg-
ativespeechaltogether.
4.3 WelfarewithBayesianversusNa¨ıveAudiences
Inouranalysisthusfar,wehaveconsideredaBayesianaudiencewhosebe-
liefsareconsistentwithequilibriumbehavior.Analternativeassumptionoften
invokedintheliteratureisthattheharmfromdefamatorystatementsisinde-
pendentofthefrequencyoffalsestatements.Wecalltheaudienceunderthis
alternative assumption na¨ıve, and we consider the differences in the implica-
tionsofamodelwithana¨ıveversusBayesianaudience.Whentheaudienceis
na¨ıve, itavoidsaninteractionwiththetargetif, andonlyif, itreceivesadis-
paragingstatementfromthespeaker,anditdoessoregardlessofthefrequency
offalsestatements.
Thus,withana¨ıveaudience,thespeakerisabletosuccessfullyblockinter-
actionswithbothtargettypeswhendamagesareverylow(i.e.d ⩽ d ).This
1
15. Weusethefollowingvaluestoproducecurvesthatdonotoverlapwitheachotherfor
expositionalpurposes:γ=rB =l= 3
4
;g=b=1;andqG=
1
7
0
.
EElleeccttrroonniicc ccooppyy aavvaaiillaabbllee aatt:: hhttttppss::////ssssrrnn..ccoomm//aabbssttrraacctt==44118811889900

<<PAGE_BREAK>>

-- 19
Figure3
EElleeccttrroonniicc ccooppyy aavvaaiillaabbllee aatt:: hhttttppss::////ssssrrnn..ccoomm//aabbssttrraacctt==44118811889900

<<PAGE_BREAK>>

20 .V0N0
isbecausethesedamagesaretoolowtogenerateanylitigationthreatfromthe
target, and thus the speaker disparages the target independently of his type.
Thena¨ıveaudience, unliketheBayesianaudience, reliesonthestatementby
thespeakerinsteadofitsprior,andthereforealwaysavoidsaninteractionwith
thetarget.
Whendamagespassontothemoderaterange,theequilibriumbehaviorand
welfareinthena¨ıveaudiencecaseisidenticaltothosethatareobservedunder
aneffectivecommunicationequilibriumoftheBayesianaudiencecase.Thisis
becausetheBayesianaudience,likethena¨ıveaudience,actsinamannercon-
sistentwiththespeaker’sstatementsineffectivecommunicationequilibria.Fi-
nally,whendamagesareveryhigh(i.e.d>d ),thespeakerisdeterredagainst
4
makingdisparagingstatementsagainstbothtypes,andthena¨ıveaudiencein-
teracts with both types. Thus, in this range welfare with a na¨ıve audience is
equaltowelfarewithaBayesianaudience.
Inshort,themoststrikingdifferencearisingfromaswitchfromaBayesian
audience to a na¨ıve audience occurs when damages are too low to cause the
target to litigate (i.e. d ⩽ d ). The most prevalent normative impact of this
1
difference is observed when there is an upper bound on maximum damages,
since otherwise optimal equilibria are trivially obtained in the intermediate
moderate range (i.e. d ∈ [d ,d ]) under both models. Thus, we focus on the
2 3
case where d¯ < d to highlight the greatest difference between the models
2
withaBayesianandana¨ıveaudience.
Aswepreviouslynoted,withaBayesianaudience,typeGtargetsaremade
worseoffwhendamagesinthelowermoderaterangelendcredibilitytospeaker’s
statements. Thus, as we noted via proposition 4, when defamatory harms are
large,welfareisactuallyreducedwhendefamationlawsareeffectivecompared
towhentheyarenot.Theoppositeconclusionholdswithana¨ıveaudience.In
this case, increasing damages always leads to less frequent defamatory state-
ments,andtypeGtargetsalwayspreferstricterdefamationlaws.Thus,when
defamatoryharmsarelarge,contrarytothecasewithaBayesianaudience,it
isoptimaltousemaximaldamages.
We illustrate this result through figure 4, which depicts welfare obtained
in the example used to generate the high r case in figure 3,16 but this time
G
italso includeswelfareobtainedwith ana¨ıveaudience. Thefigureillustrates
thatwhileitisoptimaltoincreasedamagestotheirmaximallevelwithana¨ıve
audience,itisoptimaltousedamageslowenoughtoguaranteetheemergence
ofineffectivecommunicationequilibriawithaBayesianaudience.Thus,when
harms from defamatory statements are large, assuming that the audience is
na¨ıveisnotmerelyasimplifyingassumption,itisonethatcangeneratemis-
leadingnormativeimplications.
16. Asareminder,thiscorrespondstothecasewherevisuniformlydistributed;rG = 9
4
;
γ=rB =l= 3
4
;g=b=1;andqG=
1
7
0
.
EElleeccttrroonniicc ccooppyy aavvaaiillaabbllee aatt:: hhttttppss::////ssssrrnn..ccoomm//aabbssttrraacctt==44118811889900

<<PAGE_BREAK>>

-- 21
Figure4
EElleeccttrroonniicc ccooppyy aavvaaiillaabbllee aatt:: hhttttppss::////ssssrrnn..ccoomm//aabbssttrraacctt==44118811889900

<<PAGE_BREAK>>

22 .V0N0
5. ExtensionsandDiscussion
InSections3and4, wepresentedamodelthatallowedustoclearlyfocus
ondefamationlaws’impactontheaudience’sequilibriumbeliefsandactions.
In doing so, we abstracted from many issues that bear on the regulation of
information in more general settings, particularly, the possibility of there be-
ing a committed public enforcer, quality being endogenously chosen by the
target, the existence of honest and other types of speakers, and courts being
lessaccuratethanwepreviouslyassumed.Hereweturnourattentiontothese
issues.
5.1 EndogenousTypesandDynamicEfficiencies
In our analysis thus far, we assumed that the target’s type is exogenously
determined by nature to be either G or B with probabilities γ and 1−γ, re-
spectively. One might question the reality of this assumption, as people can
makeinvestmentsthatwouldmakethembetterorworsetradingpartners,e.g.,
create higher quality products, maintain safety standards, or keep higher hy-
gienestandards.Garoupa1999a,b,forinstance,takesasimilarapproach,and
assumesthatthetarget’sbehaviorisimpactedbywhatlawsareinplace.Here,
we explain how the types in our setting can be endogenized, and how doing
so yields results similar to those in prior work where the target’s behavior is
endogenous.
One option of incorporating quality investments into our analysis is to re-
place Nature’s choice of types with a preliminary stage where the target, T,
makes a costly investment (c) that can increase her likelihood of becoming
a good type. Formally, we may assume that γ = γ(c) with γ′ > 0 > γ′′,
limγ′(c)=∞,γ(0)=γ and limγ(c)=γ where1>γ >γ >xˆ>0.
c→0 c→∞
Thequalityinvestmentdecisionisnowpartofalargergame.Givenanysub-
gameequilibrium,thebestresponseofT istomakeaninvestmenttomaximize
hisexpectedpay-off,whichcanbedenotedasγ(c)m +(1−γ(c))m −c
G B
wherem andm refertothepayoffsheobtainsinthesub-gameequilibria.
G B
Thisobservationrevealsaveryclearresult: Whenthelawsareextreme,i.e.
d ̸∈ [d ,d ], the target has no reason to invest in quality. This follows from
1 4
Proposition 2, which shows that with extreme laws, the audience acts based
on its priors and interacts with the target. Thus, investments have no private
returnsforthetarget.
Itisonlywhenthelawsaremoderatethattargetsmayhaveanincentiveto
investinquality.Thiscanbedemonstratedbyfocusingonthelowerboundof
intermediatedamages,i.e.d .Inthiscase,ineffectivecommunicationequilib-
1
ria,itfollowsthatm =0(becauseallbadtypesaredisparaged)whilem =
B G
F(v (d))r(becausegoodtypesaredisparagedwithprobability1−F(v (d)),
G G
inwhichcasethereisalawsuitwhichpaysthetargetexpecteddamagesequal
tolitigationcosts).Thus,thetarget’spay-offisγ(c)F(v (d))r−c,and,there-
G
fore,thetargetprofits(inexpectation)frominvesting.Whetherthisissocially
good or bad, depends, of course, on whether there are net social gains from
such investments. In our context, this is socially valuable as long as the ex-
EElleeccttrroonniicc ccooppyy aavvaaiillaabbllee aatt:: hhttttppss::////ssssrrnn..ccoomm//aabbssttrraacctt==44118811889900

<<PAGE_BREAK>>

-- 23
pectedbenefitsfromgoodinteractions(F(v (d))g)—whicharenotinternal-
G
izedbyT—aregreaterthantheexpectedlitigationcostslandthelossofben-
efit to S from blocking an interaction, i.e. F(v (d))E[v|v > l]. In fact, if
G 2
investments in quality are socially valuable, as is implicitly assumed in the
literature (e.g., Garoupa 1999a,b), then increasing damages within the inter-
mediate range up to d will be desirable. This is because these higher dam-
2
ages lead to a lower probability of disparaging remarks made against good
types (as illustrated in Figure 2) and, thus, increase m , while still keeping
G
expectedpayoffsfrombeingabadtypeatm = 0.Therefore,theextension
B
of our model with endogenous types resonates with Garoupa’s (1999a,b) in-
sightsthatmoderatedamagescanincentivizeinvestmentsinbecomingagood
type.Moreover,ithighlightsthepotentialsocialcostsandbenefitsassociated
withsuchinvestmentmorespecifically.
Thediscussionherehighlightstheimportanceofinformationregulationfor
broadermarketdynamics.Theintuitionunderlyingourresultsarestraightfor-
ward.Extremelawsleadtoineffectivecommunicationequilibria.Incontrast,
moderatelawscreateanenvironmentwithmorereliableinformationregarding
types, thus generating a greater gap between the payoffs obtainable by good
typesversusbadtypes.17Inrealisticsettings,providingsuchadditionalincen-
tivesissociallydesirablewhenthepotentialinvestorisunderincentivizeddue
to problems like information asymmetries. The gains from such investments
inqualityshouldbeaddedtotheotherbenefitsofmoderatelawsthatwehave
identified.
5.2 HonestSpeakersandEulogists
Existinganalysesofdefamationlawtypicallyassumethatthespeaker’sneg-
ative statement always harms the target, which is equivalent to the audience
being na¨ıve. Moreover, these analyses (e.g. Garoupa 1999a,b and Bar-Gill &
Hamdani2003)considerstrategicspeakerswhobenefitfromdefamingthetar-
get,andwhosebenefitsfromdoingsoareindependentoftheveracityoftheir
statements.Inreality,however,manyspeakersmaynothavesuchmotivations.
Quite importantly, many people, when asked their opinion, provide an hon-
est assessment of others. Moreover, there are also people who are motivated
by doing the exact opposite of what the speakers in our model are motivated
by; namely, promoting the relationship between the target and the audience.
Inwhatfollowswedistinguishbetweenthefirsttype,truthspeakers,thelatter
type,eulogists,andtheonesweformerlydiscussedinsection3asdisparagers.
Here,webrieflyexplainwhatoccurswhenthesekindsofspeakersareincor-
poratedintoouranalysis.
Inourdiscussion,weconceiveofthesetypesasfollows.Disparagers,aswe
noted,receiveapositivevaluefromblockinganinteraction;truth-speakersare
17. Thisresultisreminiscientofthedeterrencereducingimpactsofjudicialerrorsobtained
inthelawenforcementliterature(see,e.g.,Png(1986),RizzolliandGraoupa(2012),Mungan
(2017), and Lando and Mungan (2018)) wherein judicial errors dilute the deterrence effect of
punishmentbycreatingadisconnectbetweenpunishmentandbehavior.
EElleeccttrroonniicc ccooppyy aavvaaiillaabbllee aatt:: hhttttppss::////ssssrrnn..ccoomm//aabbssttrraacctt==44118811889900

<<PAGE_BREAK>>

24 .V0N0
indifferent with respect to whether the parties will interact but receive some
valuefromspeakingtheirmind; and,eulogistsreceiveavaluefromtherebe-
inganinteraction.Therefore,solongascostsofsodoingarenothigh,dispar-
agers will badmouth the target and truth-speakers will reveal their true type.
Eulogists, in contrast, would always want to praise the target, as there is no
recourse under defamation law for false positive statements (the question of
whythisasymmetryexistsgoesbeyondthescopeofourarticle).
Theincorporationofthesetypesofspeakershasnoimpactontheobserva-
tionthatextremelystrongdefamationlawsleavetheaudiencetoactupontheir
priors. This follows, because once a critical threshold of damages is passed,
disparagers as well as truth speakers are deterred from making negative re-
marks.Thus,extremelystrongdefamationlawscausedisparagers,truthspeak-
ers, and eulogists alike to abstain from making negative statements, and the
audiencehasnooptionbuttoactaccordingtoitspriors.
The same cannot be said, however, for extremely weak defamation laws.
When damages are very low, targets lack an incentive to bring suit, making
talk cheap. Despite that, disparaging statements are still somewhat informa-
tive:Giventheexistenceofsometruth-speakers,thereissomeprobabilitythat
anynegativestatementistrue.Consequently,anaudiencethathearsanegative
statementevaluatesitscredibilitybasedontheratiooftruth-speakerstodispar-
agers.Thus,inanassessmentwitha∗(z)=z,wecanformulatetheaudience’s
consistentbeliefthatthetargetisagoodtype,conditionalonanegativestate-
ment as x∗ = γ ∆ where τ denotes the proportion of truth speakers,
1 ∆+(1−γ)τ
and ∆ is the proportion of disparagers. On the other hand, non-disparaging
remarksdonotnecessarilymeanthatT isagoodtype.Bysimilarlogic,there
is some probability that any given praise is false given the existence of eulo-
gists.Anaudiencewhichhearsapositivestatementevaluatesitsveracityasa
function of the ratio of eulogists to truth-speakers. Thus, we can express the
audience’sbeliefasx∗ =γ τ+ε ,whereεistheproportionofeulogists.
0 γτ+ε
Using these observations it is easy to verify that, under lax laws, both dis-
paraging and non-disparaging statements are somewhat informative of types.
Inotherwords,non-disparagingstatementsaremoreindicativeofgoodtypes
than no information at all (x∗ > γ), and disparaging statements are more
0
indicativeofbadtypesthannoinformationatall,i.e.x∗ <γ.Thus,iftheau-
1
dience’snecessarylevelofconfidenceforinteraction,(x),iscloseenoughtoγ
(cid:98)
suchthatx∗ ⩾ x ⩾ x∗,onecanachieveanequilibriumwhereintheaudience
0 (cid:98) 1
meaningfullyusestheinformationprovidedbyspeakers,evenwhenthereare
nosanctionsforfalsestatements.If,however,x∈/ [x∗,x∗],thenlaxlawscause
(cid:98) 1 0
theaudience toignore thestatement andact accordingto its priors, as inour
analysisinsection3.Thus,wefocusourremainingdiscussiontocaseswhere
x∗ ⩾x⩾x∗.
0 (cid:98) 1
Incaseswheredamagesaremoderate, someoftheclaimsmadeinsection
3 need to be qualified, whereas others remain intact. In particular, it is still
the case that moderate damages improve the reliability of information over
extremedamages.Toseethis,consider,forinstance,theimplicationsofrais-
EElleeccttrroonniicc ccooppyy aavvaaiillaabbllee aatt:: hhttttppss::////ssssrrnn..ccoomm//aabbssttrraacctt==44118811889900

<<PAGE_BREAK>>

-- 25
ingdamagesfromlowlevelsto l .Amongspeakers,thischangeonlyalters
2qG
the incentives of disparagers, because these are the only speakers who have
aninterestinmakingfalsestatementsaboutgoodtypes,who,giventhislevel
ofdamages,bringalawsuitagainstthem.Thus,theproportionofdisparagers
who make false statements is reduced, which causes x∗ to fall and x∗ to in-
1 0
crease,i.e.itcausesinformationsuppliedbyspeakerstobemoreinformative.
Thisobservationrevealsanotherofourresultsthatcarriesoverinamodified
way:onecanusedamagesequaltod <d todeteralldisparagersfrommak-
2 3
ingfalsestatementsandalsoguaranteethattherearenolawsuitsbybadtype
targets.Inthiscase,itimmediatelyfollowsthatx∗ =0,suchthatadisparaging
1
statementisperfectlyinformative.
The presence of eulogists, however, means that x∗ < 1. Thus, equilibria
0
that are completely informative of the target’s type are no longer obtainable.
Still,eveninthepresenceofeulogistsanddisparagers,imperfectlyinformative
equilibria are possible. Moreover, these equilibria are optimal, because they
lead to no litigation costs, cause all possible good interactions to take place,
andachievemaximumdeterrenceofbadinteractions.
We conclude that the introduction of honest speakers as well as what we
calledeulogists—peoplewhowishtopromotethetarget—doesnotaffectthe
superiorityofmoderatedamagesoverextremeformsofdamages.Whatdoes
changeisperhapssomewhatcounterintuitive: strictlawsturnouttobeworse
thanlaxlaws.Strictlawsleadtocompletelyuninformativespeechinequilib-
riumwhereaslaxlawsstillallowspeechtobesomewhatinformative,permit-
tingeffectivecommunicationequilibria.
5.3 CommitmentandPublicEnforcement
Our analysis so far focused on private enforcement of defamation laws,
where the target is the one to sue. However, private parties will only bring
alawsuitifitpaystodosoex-post, andthiscalculusexposesthemtostrate-
gicbehaviorbywould-bedefamers.Incontrast,someparties,typicallypublic
agencies, maybeabletocommitex-antetosue, evenifitdoesnotpaytodo
soex-post.Infact,thisistheimplicitassumptioninvokedinmuchoftheeco-
nomics of public law enforcement literature (see, e.g., Polinsky and Shavell
2007 for a survey). Comparing private and public enforcement can be useful
inunderstandingothercontextswhereinformationisregulated,andmayalso
illuminatethereasonswhyprivateenforcementisusedindefamation.
Tohelpinthiscomparison,weconsiderasimplemodificationofouranal-
ysis wherein instead of the target, it is a public enforcement agency that can
bring suit against disparaging remarks. The agency, however, is not privy to
thetarget’sprivateinformationregardinghistype,whichisbyassumptionun-
observable,andsoitcannotconditionitsactiononT’stype.Theagencythus
choosessomeprobability,p∈(0,1),withwhichitwillbringalawsuit.Asthe
choiceofpdoesnotdependonanynewinformation,itischosenex-anteandis
communicatedto,orobservedby,would-bespeakers.Thechoiceofpreplaces
p∗(t) in in the prior sections. We retain all other assumptions, including the
EElleeccttrroonniicc ccooppyy aavvaaiillaabbllee aatt:: hhttttppss::////ssssrrnn..ccoomm//aabbssttrraacctt==44118811889900

<<PAGE_BREAK>>

26 .V0N0
assumption that the probabilities with which the speaker will be found liable
incourtareq andq ,whenshemakesdisparagingstatementsagainstgood
G B
andbadtypes,respectively.
Thissimplemodificationallowsustocalculatetheanalogsofthetwocrit-
ical damages pertaining to the best responses of S noted in (3). Specifically,
these two critical values now become d˜ ≡ 2v¯−pl and d˜ ≡ 2v¯−pl. Thus, in
effectivecommunicationequilibriawith 2 d>d˜ 2pq th G espeake 4 rdoe 2 s p n q o B tmakedis-
2
paragingstatementsagainstgoodtypes,andrefrainsfrommakingdisparaging
statementsagainstbadtypeswhend > d˜.Itcanbeeasilyverifiedthateach
4
ofthesevaluesislargerthantheircorrespondinganalogintheprivateenforce-
mentcontext,i.e.d˜ >d andd˜ >d .
2 2 4 4
The commitment to bringing a lawsuit also changes the speaker’s behav-
ior, as a lawsuit is possible even when expected damages are low. We next
explainthebehaviorofthespeakerineffectivecommunicationequilibria,un-
der three different damages ranges, and subsequently compare them with the
correspondingbehaviorunderprivateenforcement.
Asunderprivateenforcement,itfollowsthatwhendamagesareveryhigh,
i.e.,d>d˜,alldisparagingremarksaredeterred.However,whend∈(d˜,d˜),
4 2 4
the speaker refrains from disparaging good types, but disparages bad types
wheneverhervaluefromblockinginteractionsissufficientlyhigh(i.e.v˜ (d)≡
B
p(q d− l) < v)whichhappenswithprobability1−F(v˜ (d)) > 0.Thus,
B 2 B
in this range, a disparaging remark conclusively reveals to the audience that
thetargetisabadtype;andanon-disparagingcommentisaninformative,but
inconclusive,signalthatthetargetisagoodtype,i.e.x∗ =0<γ <x∗.When
1 0
damages are low, i.e., d < d˜, the speaker is no longer necessarily deterred
2
from disparaging good types, and chooses to defame the target if her value
from blocking interactions exceeds v˜ ≡ p(q d− l). Thus, it follows that
G B 2
0<1−F(v˜ (d))<1−F(v˜ (d)),and,therefore,0<x∗ <γ <x∗ <1.
G B 1 0
We can now compare defamation laws under public and private enforce-
mentregimes.First,effectivecommunicationequilibriaarenotpossibleunder
either regime when damages are extremely high (i.e. higher than d˜ and d
4 4
in the public and private regimes, respectively). Thus, our previous conclu-
sion regarding the ineffectiveness of high damages in supporting informative
statementsextendstothepublicenforcementcaseaswell.
Second, unlike private enforcement, public enforcement can sustain effec-
tivecommunicationequilibriaevenwithlowdamages.Underprivateenforce-
ment, the speaker will anticipate that the target will not sue if damages are
sufficiently low. This can lead the speaker to disparage regardless of the tar-
get’s type, which would make statements non-informative. With public en-
forcement, however, there is always a risk of a lawsuit, thus deterring some
would-bedefamersandsustainingthereliabilityofsomestatements.Thisim-
pliesthat,unlikeintheprivateenforcementcontext,verylowdamagescanbe
usedtosupporteffectivecommunicationequilibria—atleastwhenthethresh-
oldbeliefoftheaudience,i.e.xˆ,isnottoofarfromitspriors,i.e.whenγ−xˆ
is not large, because then xˆ ∈ [x∗,x∗]. Note that this means that low dam-
1 0
EElleeccttrroonniicc ccooppyy aavvaaiillaabbllee aatt:: hhttttppss::////ssssrrnn..ccoomm//aabbssttrraacctt==44118811889900

<<PAGE_BREAK>>

-- 27
agescanbesuperiortohighdamagesinfacilitatingeffectivecommunication
betweenthespeakerandtheaudience.
Third, and quite importantly, it is impossible to obtain an equilibrium that
alwaysrevealsthetarget’stypewithpublicenforcement: asnotedabove,any
damages below d < d˜ result in good types being disparaged with a proba-
4
bilityof1−F(v˜ (d)) > 0,badtypesbeingdisparagedwithaprobabilityof
G
1−F(v˜ (d))<1,orboth.Thisimmediatelyimpliesthatprivateenforcement
B
dominatespubicenforcementintermsofitswelfareconsequences.Thediffer-
ence in the welfare obtainable under the two regimes is enhanced further by
thefactthatunderpublicenforcement,theenforcementagency’scommitment
resultsinlitigation.
Thelastpointhighlightsamoregeneralandimportantadvantageofprivate
enforcementoverpublicenforcement.Specifically, privateenforcementdele-
gates the decision to litigate to the party with the best information about the
meritsofthecase.Moderatedamagescanbecraftedtoseparategoodandbad
types based on their willingness to sue, and this enables the speaker’s state-
mentstobemoreinformativeofthetarget’stype.
Insum,thiscomparisonilluminatetherelativevalueofpublicversuspublic
enforcement.However,asourfocushereisoncommitment,weabstractfrom
otherrelevantconsiderations, suchastherelativecostsoflearningaboutdis-
paragingremarksorproducingevidence.Inasmuchaspublicagenciesemploy
discretion, they are also susceptible to capture and other public choice prob-
lems.Theseconsiderationsshouldalsobetakenintoaccountincomparingthe
relative social desirability of pubic versus private enforcement in regulating
speech.
5.4 InaccurateCourts
To keep our analysis focused, we presented results obtained in the case
wherethecourtisrelatively‘accurate’inrenderingdecisions,inthesensethat
itcommitserrorswithlowfrequency(i.e.q <q ( l/2 )).Themainroleof
B G 1−l/2
thisassumptionistoguaranteetheexistenceofarangeofmoderatedamages
whichleadstoseparatingequilibria.
We relax the assumption that courts are accurate in a prior version of this
article,18 which reveals that most results we presented extend to this case. A
notableexceptionisthattheremaybenoeffectivecommunicationequilibrium
whichleadstogreatersocialwelfarethanwhentheaudienceislefttoactbased
onitspriors,andthisistrueevenwhentherearenodamagecaps.Thereason
forthisisthatthesmallestdamageswhichincentivizetypeB targetstobring
suit(i.e.d )arethensmallerthanthesmallestdamagesthatdeterthespeaker
3
from disparaging type G targets. This causes there to be no damage amount
whichislowenoughtodisincentivizefrivolouslawsuitswhilealsocompletely
deterring defamatory speech. Thus, effective communication equilibria can-
noteliminatefalsestatements.Thiscancauseverylowdamagestobeoptimal,
duetothesamerationalesthatemergewhendamagesarebounded(seesection
18. SeeArbel&Mungan2020.
EElleeccttrroonniicc ccooppyy aavvaaiillaabbllee aatt:: hhttttppss::////ssssrrnn..ccoomm//aabbssttrraacctt==44118811889900

<<PAGE_BREAK>>

28 .V0N0
4.2).ThisresultfurtherhighlightstheimportanceofBayesianaudiences.With
ana¨ıveaudience,standardeconomicmodelswouldpredictthatoptimaldam-
agesaremoderate, becauselowdamageswouldinvitetoomuchfalsespeech
towhichtheaudiencelendscredence.
6. Conclusion
Existing economic analyses of defamation law typically assume that there
are no obstacles in the way of a person who wishes to harm another person
throughdefamatorystatements.However,forsuchharmstoberealized,people
must find the derogatory statements made by the person credible. The credi-
bility of these statements, in turn, depend on what types of consequences a
speaker faces by making such statements. Defamation law is a tool that can
be used to alter these consequences, and therefore the credibility of negative
statements.Wehaveformalizedthisdynamicbystudyingthebeliefformation
processofpeoplewhoaretherecipientsofsuchinformation.
Our analysis has revealed several interesting dynamics. When courts are
accurateandthemaximumrecoverabledamagesfromthedefendantarelarge,
onecanlargelyeliminatefalsespeech.However,whentheseconditionsdonot
hold,itispossiblefortheregulationofspeechthroughdefamationlawtocause
more harm than benefits. This may occur because the credibility of speech
obtainedthroughweakorinaccurateenforcementofdefamationlawcanlead
toanincreaseinfalsenegativespeechwhichisbelievedbytheaudience(and
thismayalsoleadtosignificantlitigationcosts).Inothercases,itpossiblefor
the increased informativeness obtained through defamation law to outweigh
its social costs. Which of these two cases is obtained depends, among other
things,ontheharminflictedtothetargetofspeech.
Onereasonwhytheseconclusionshavenotbeendevelopedinthevastlegal
literatureonthetopicispresumablytheinsufficientattentionpaidtotherole
of audiences, which have typically assumed to be na¨ıve. We believe that our
basic insights are also applicable to many other areas of law where the goal
is to regulate the veracity of information. Although our analysis represents a
stepforwardinunderstandingimportantdynamicsinthesecontexts,wewere
unabletoincorporatemanyotherpossibilitieswhichmayyieldadditionalin-
terestinginsights.Wehighlightsomeissueswehaveabstractedfromwiththe
hopeofhighlightingsomeavenuesforfutureresearch.
We have focused, for instance, exclusively on plaintiffs who bring suit to
increasetheirmonetarywell-being.But,therearemanyplaintiffswhoaremo-
tivatednotbythedamagestheymayrecover,butbytheprospectofadvancing
thetruth.Inthesecircumstances,largerexpecteddamagesmaycrowd-outthe
intrinsicmotivationsofthepotentialplaintifftobringsuitbymakingitharder
forthirdpartiestoidentifythetruemotivationsoftheplaintiff.Similarly, we
consideredahomogenousaudienceandassumedthatthespeakerhasperfect
informationregardingthetarget’stype.Relaxingtheseassumptionsmaycause
a greater wedge between the results obtained with na¨ıve audience members
and with Bayesian audience members. We hope that the framework we have
EElleeccttrroonniicc ccooppyy aavvaaiillaabbllee aatt:: hhttttppss::////ssssrrnn..ccoomm//aabbssttrraacctt==44118811889900

<<PAGE_BREAK>>

-- 29
Figure5
developed here can be used as starting point to incorporate these additional
considerations.
7. Appendix
GameTree
Figure5depictstheinteractionsbetweentheplayers.Itsummarizesthepar-
ties’payoffsattheterminalnodesonthebottomintheorderS,A,T.Thereare
twographicallimitationsofFigure5.First,itdoesnotshowinformationsets
describing A’s knowledge regarding S’s type, due to the depiction difficulty
causedbyS drawinghertypefromacontinuum.Second,foreaseofexposi-
tion, Figure5doesnotdepictNature’sv drawdeterminingS’sinclinationto
disparage.
PerfectBayesianEquilibriumRequirements
InformalizingtherequirementsforaPBE,wefirstspecifytheunconditional
(orexante)probabilitywithwhichSwilldisparageT givenanystrategy,s,as
follows:
(cid:90) 1
µ(s)≡ [γs(G,v)+(1−γ)s(B,v)]dF(v) (1)
0
When µ(s) ∈ (0,1), we can use Bayes’ rule to calculate the probability
ofT’stype, goodorbad, conditionalonthestatementmadeaboutT.Onthe
otherhand,whenµ(s) ∈ {0,1},itfollowsthatS isplayingastrategywhere
he(almost)always avoidsdisparaging(0)ordisparages(1) T, inwhichcase
Bayes’rulecannotbeusedtocalculatetheprobabilityofT beingaparticular
EElleeccttrroonniicc ccooppyy aavvaaiillaabbllee aatt:: hhttttppss::////ssssrrnn..ccoomm//aabbssttrraacctt==44118811889900

<<PAGE_BREAK>>

30 .V0N0
type, conditional on the strategy which is (almost) never played by S. Thus,
wedenotebothpossibilities,asfollows:
(cid:82)1s(G,v)dF(v)
γ 0 if µ(s)̸=0
Γ(t=G|z =1,s)≡ µ(s) (2)
Υ otherwise
(cid:82)1(1−s(G,v))dF(v)
γ 0 if µ(s)̸=1
Γ(t=G|z =0,s)≡ 1−µ(s) (3)
Υ otherwise
Here, the symbol Υ indicates that the strategy in question is (almost) never
chosenbythespeaker.
GiventhisnotationwemaycharacterizePBEasanassessmentconsistingof
thestrategyprofilea∗,s∗ andp∗ alongwithasetofbeliefsx∗ andx∗,which
0 1
satisfiesthefollowingfourrequirements.
Requirement1(R1): Ahasnoprofitabledeviationgivenitsbeliefs:
a∗(z)=0 if x >x forz ∈{0,1}
z (cid:98) (4)
a∗(z)=1 if x <x forz ∈{0,1}
z (cid:98)
R1statesthatAinteractswithT onlyifAbelieves,givenS’sstatement,that
the probability that T is a good type exceeds the threshold probability of x.
(cid:98)
Similarly, if A believes that T is a good type with a probability that is lower
thanx, AdoesnotinteractwithT.Intheexceptionalcasewherex = x, A
(cid:98) z (cid:98)
isindifferentbetweeninteractingwithT andnot,and,thusitmayplayeither
strategy.
Requirement2(R2):T hasnoprofitabledeviationsinsub-games:
0 if q d<l/2 fort∈{B,G}
p∗(t)= t (5)
1 if q d>l/2 fort∈{B,G}
t
R2statesthatthePBEstrategyofT mustbesuchthatinsubgameswhereS
disparageshim,T litigateswheneverthecostsofdoingso(l/2)arelowerthan
theexpecteddamagerewardsthathecanobtainfromlitigation.Conversely,T
choosesnottolitigatewhenthecostsarehigherthanexpecteddamages.Inthe
exceptionalcasewhereq d=l/2,T isindifferentbetweenlitigatingandnot.
t
Requirement3(R3):Shasnoprofitabledeviations:Forallt,vpairs,s∗(t,v)
maximizesplayerS’spayoff,whichcanbeexpressedas
l
U ≡a∗(s(t,v))(v−p∗(t)s(t,v){q d+ }) (6)
S t 2
The requirement with respect to S appears more complex than the re-
quirementsthatpertaintoT andA’sstrategies,becauseS choosesheractions
in anticipation of the other players’ actions. Still, the requirement is simply
that,givenherowntype,T’stype,andtheanticipatedbehaviorofAandT,S
mustchoosethecourseofactionthatwouldmaximizeherpayoff.
Requirement4(R4):A’sbeliefsareconsistent:
x∗ =Γ(t=G|z,s∗)wheneverΓ(t=G|z,s∗)̸=Υforbothz ∈{0,1} (7)
z
R4 simply states that A’s beliefs must be consistent with the implied condi-
tionalprobabilityofT beingaparticulartypebasedontheequilibriumstrategy
EElleeccttrroonniicc ccooppyy aavvaaiillaabbllee aatt:: hhttttppss::////ssssrrnn..ccoomm//aabbssttrraacctt==44118811889900

<<PAGE_BREAK>>

REFERENCES 31
of S. This requirement is applicable only to strategies which have a positive
probabilityofbeingplayedbyS.
Proof of Proposition 2: The proof begins with part (iii), which is used in
provingpart(i).
(iii) We show that the audience ends up always interacting with T, in all
equilibria where the actions of the audience are not described by a∗(z) = z
forz ∈{0,1}.
SupposethereisaPBEwherea∗(z) = 0forallz.Bydefinition,theaudi-
encealwaysinteractsinsuchassessments.
Suppose there is a PBE where a∗(z) = 1 − z for all z, then per R3,
s∗(t,v) = 0 for all v and t, and, therefore, µ(s∗) = 0, which implies that
Γ(t = G|0,s∗) = γ.ThisimpliesviaR4thatx∗ = γ,which,inturnimplies
0
viaR1that a∗(0)=0,whichcontradictstheassumptionthata∗(0)=1.
Suppose there is a PBE where a∗(z) = 1for all z. If µ(s∗) = i ∈ {0,1},
then Γ(t = G|i,s∗) = γ, which implies via R4 that x∗ = γ. This implies
i
viaR1thata∗(i)=0,whichisacontradictionwiththeinitialsupposition.If,
on the other hand, µ(s∗) ∈ (0,1), observe that, per R4, x∗ ⩽ γ implies that
0
x∗ ⩾ γ, because x∗(1−µ(s∗))+x∗µ(s∗) = γ. Thus, x∗ ⩽ γ implies that
1 0 1 0
x∗ ⩾γ >x,whichisacontradictionwiththeimplicationofR1thatx∗ ⩽x.
1 (cid:98) 1 (cid:98)
(i) Consider damages d < d , and suppose a∗(z) = z for all z. It follows
1
viaR2thatp∗(t)=0forallt.Thus,R3impliesthats∗(t,v)=1forallvand
t, and, therefore, x∗ = γ due to R4 . Thus, in equilibrium, the audience acts
1
accordingtoitspriors.
Next, consider damages d > d . It follows per R2 that p∗(t) = 1. Thus,
4
perR3,s∗(t,v) = 0forallv andt,becaused > d .ThisimpliesviaR4that
4
x∗ =γ.Thus,inequilibrium,theaudienceactsaccordingtoitspriors.
0
Theanalysisofthesetwocasesdemonstratesthatwhend ̸∈ [d ,d ],inall
1 4
PBE where a∗(z) = z for all z, the audience acts according to its priors. In
addition, part (ii) of this proposition demonstrates that the audience acts ac-
cordingtoitspriorsinallPBEwheretheaudience’sbehaviorisnotdescribed
bya∗(z) = z.Thus,wheneverd ̸∈ [d ,d ],theaudienceactsaccordingtoits
1 4
priorsinallPBE.
(ii)Thediscussionofseparatingequilibriainsection3.4demonstratesthat
suchdamagesexist.
References
Acheson, D. J. and A. Wohlschlegel. 2018. The Economics of Weaponized
DefamationLawsuits.47SouthwesternLawReview335-384.
Arbel, Y. and M. Mungan. 2019. The Case Against Expanding Defamation
Law.71AlabamaLawReview453-497.
Arbel,Y.andM.Mungan.2020.RegulatingInformationwithBayesianAudi-
ences,GeorgeMasonLaw&EconomicsResearchPaper19-28.
EElleeccttrroonniicc ccooppyy aavvaaiillaabbllee aatt:: hhttttppss::////ssssrrnn..ccoomm//aabbssttrraacctt==44118811889900

<<PAGE_BREAK>>

32 .V0N0
Arbel,Y.2021.AStatusTheoryofDefamationLaw.AlabamaWorkingPaper
Series2021.
Arbel,Y.2022.TheCredibilityEffect,AlabamaWorkingPaper.
Be´nabou,R.,andJ.Tirole.2006.IncentivesandProsocialBehavior.96Amer-
icanEconomicReview1652-1678.
Be´nabou, R., and J. Tirole. 2011. Laws and Norms. National Bureau of Eco-
nomicResearchNo.w17579.
Bar-Gill,OrenandAssafHamdani.2003.OptimalLiabilityforLibel.2Con-
tributionsinEconomicAnalysis&Policy,1-26.
Crawford, V. and J. Sobel. 1982. Strategic Information Transmission, 50
Econometrica1431-1451.
Dalvi, M. and J. Refalo. 2008. An Economic Analysis of Libel, 34 Eastern
EconomicJournal74-94.
Deffains,B.andC.Fluet.2020.SocialNormsandLegalDesign,36TheJour-
nalofLaw,Economics,andOrganization136-169.
Garoupa,N.1999.TheEconomicsofPoliticalDishonestyandDefamation,19
InternationalReviewofLawandEconomics167-180.
Garoupa, N. 1999. Dishonesty and Libel Law: The Economics of the ”Chill-
ing”Effect,155JournalofInstitutionalandTheoreticalEconomics284-300.
Garoupa,N.andM.Rizzolli.2012WrongfulConvictionsdoLowerDeterrence
168JournalofInstitutionalandTheoreticalEconomics224-231.
Hemel, D. and A. Porat. 2019. Free Speech and Cheap Talk, 11 Journal of
LegalAnalysis46-103.
Hemel,D.2020.EconomicPerspectivesonFreeSpeech.OxfordHandbookof
FreedomofSpeech,118-136.
Heymann, L.2012. TheLawofReputation, andtheInterestoftheAudience,
52BostonCollegeLawReview1341-1439.
Lando,H.andM.Mungan.2018.TheEffectofType-1ErroronDeterrence53
InternationalReviewofLawandEconomics1-8.
Mungan, M. 2016. A Generalized Model for Reputational Sanctions and the
(Ir)relevance of the Interactions between Legal and Reputational Sanctions,
46InternationalReviewofLawandEconomics86-92.
Mungan,M.2017.WrongfulConvictions,Deterrence,andStigmaDilution25
SupremeCourtEconomicReview199-216.
EElleeccttrroonniicc ccooppyy aavvaaiillaabbllee aatt:: hhttttppss::////ssssrrnn..ccoomm//aabbssttrraacctt==44118811889900

<<PAGE_BREAK>>

REFERENCES 33
Pennycook, G., Bear, A., Collins, E., and D. Rand (2020) The Implied Truth
Effect: Attaching Warnings to a Subset of Fake News Headlines Increases
PerceivedAccuracyofHeadlinesWithoutWarnings,66ManagementScience
4944-4957.
Png,IvanPL.1986.OptimalSubsidiesandDamagesinthePresenceofJudi-
cialError6InternationalReviewofLawandEconomics101-105.
Polinsky,M.andS.Shavell.2007.TheTheoryofPublicEnforcementofLaw
in1HandbookofLawandEconomics403-454.
Posner, R., 1986. Free Speech in an Economic Perspective, 20 Suffolk Law
Review1-54.
Posner,R.1973.EconomicAnalysisofLaw,1sted.
Post,R.1986.TheSocialFoundationsofDefamationLaw:Reputationandthe
Constitution.74CaliforniaLawReview691-742.
Rasmusen,E.1996.StigmaandSelf-fulfillingExpectationsofCriminality,39
TheJournalofLawandEconomics519-543.
Sunstein,C.2021.Liars:FalsehoodsandFreeSpeechinanAgeofDeception.
EElleeccttrroonniicc ccooppyy aavvaaiillaabbllee aatt:: hhttttppss::////ssssrrnn..ccoomm//aabbssttrraacctt==44118811889900"#;

#[derive(Clone, Debug)]
pub struct Paper<'a> {
  pub paper_id: &'a str,
  pub title: &'a str,
  pub ssrn_url: &'a str,
  pub year: i32,
  pub authors: &'a [&'a str],
  pub keywords: &'a [&'a str],
  pub summary_md: &'a str,
  pub summary_zh_md: &'a str,
  pub one_pager_md: &'a str,
  pub study_pack_md: &'a str,
  pub article_text: &'a str,
}

pub fn as_paper() -> Paper<'static> {
  Paper {
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
  }
}

fn main() {
  print!("{}", ARTICLE_TEXT);
}
