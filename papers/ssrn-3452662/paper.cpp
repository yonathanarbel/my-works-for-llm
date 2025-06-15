#include <iostream>
#include <string>
#include <vector>
#include <map>
#include <fstream>
#include <random>
#include <algorithm>
#include <ctime>
#include <thread>

const std::string ARTICLE_TEXT = R"EOF(J,V0N0 1
Regulating Information With Bayesian Audiences
YonathanA.Arbel1andMuratMungan2
1SchoolofLaw,UniversityofAlabama
2ScaliaSchoolofLaw,GeorgeMasonUniversity
We analyze the regulation of false statements in the presence of Bayesian
audiences. We find that: (a) Often, moderate sanctions are optimal even
though strict sanctions can fully deter all false statements; (b) the existence
of separating equilibria—where only truthful statements are made—critically
depends on judicial accuracy; (c) the magnitude of sanctions trades-off
false information, chilling of truthful statements, and litigation costs; and (d)
private enforcement often dominates public enforcement despite the lack of
commitment. We emphasize the case of defamation law, and discuss other
contexts including securities regulation, whistle-blower incentives, jury trials,
andreportsofcriminalactivity.
WearethankfulforthecommentsofScottBaker,AlbertChoi,EzraFried-
man,NunoGaroupa,AlexLee,BenMcMichael,AlanMiller,SepehrShahsha-
hani, Kathy Spier, Bruno Srulovici, Abe Wickelgren, and the participants of
the2019LawandEconomicTheoryConference.
1. Introduction
In many contexts, we use the law to regulate the exchange of information
betweenprivateparties.Acommonconcernisthataninterestedspeakerwould
spreadfalseinformationtoadvanceitsownprivategoals.Topreventthis,the
lawwillsometimespunishfalsestatementsorrewardtruthfulones.
A common neglect in the literature is the interaction between the severity
of the law and the audience’s beliefs and actions. In reality, audiences pro-
cessinformationdifferentlywhenitsveracityisstrictlyregulated.Thisneglect
maybeduetothenaturaltendencytofocusonthepartiesthattakeanactive
partinthelegalprocess(thevictim-defendantandthespeaker-plaintiff)andto
abstractfromnon-participatingparties,namelythepublic(Heymann, 2012).
Whateverthereason,regulationoftheinformationenvironment—theflowand
qualityofinformationtothepublic—affectsaudiencesandtheirbeliefsquite
directly.
Ourobjecthereistobridgetheaudiencegapbyformalizingtheinteraction
betweenspeakers,thetargetsoftheirspeech,andmembersoftheaudience.We
employatoolthatisnaturallyaptatanalyzingthisissue,namely,aBayesian
game,andweinvestigatetheimpactofthestrictnessofthelawontheemerg-
ingPerfectBayesianEquilibria(PBE).Underthisframework,aspeaker,who
has private information about a business or individual (“target”), may make
Draft,Vol.0,No.0,
doi:/ewmxxx
(cid:13)c .
Allrightsreserved.ForPermissions,pleaseemail:
Electronic copy available at: https://ssrn.com/abstract=3452662
2 .V0N0
claims about the target to an audience member. The audience member then
decideswhethertointeract—trade,trust,socialize—withthetarget.Ifthetar-
get loses an interaction, he may bring a lawsuit against the speaker. Within
thisframework,itissociallyoptimalforaudiencestoonlyinteractwithhigh-
quality targets and avoid low-quality ones. The key variable of interest is the
strictness of the law, which we operationalize through the level of damages
awarded to the target if the lawsuit is successful—this reflects the relatively
broaddiscretioncourtshaveinthedeterminationofdamages(Steenson2014).
Our model contains four key features: (i) The information is provided by
aparty(thespeaker)whoisinterestedininfluencingtheaudience’sbehavior,
(ii)theaudiencemakesdecisionsinlightofthecontentofthesuppliedinfor-
mation,(iii)thespeaker’sobjectiveconflictswiththatofanotherparty(thetar-
get),and(iv)thelawpenalizesthesupplyoffalsenegativeinformationbythe
speaker.Thesekeyfeaturesarepresentinmanycontexts,including: defama-
tion law, whistle-blower rewards, complaint-driven law enforcement, and se-
curities regulation. In some of these contexts, legal proceedings are initiated
by the target (private enforcement) and in others by a governmental agency
(publicenforcement).Giventhegrowingpressuretoincreasetheregulationof
defamatoryspeechcomingfromtheSupremeCourt,politicalleaders,lawyers,
andscholars(Arbel&Mungan, 2019),wefocusondefamationlawasourrun-
ningexamplewithprivateenforcementinourbaselinemodel(Sections3and
4).Wesubsequentlyextendtheanalysistocomparepublicandprivateenforce-
ment,anddiscussspecificfieldsbesidesdefamationlaw(inSection5).
Our analysis reveals five central findings. First, the harmful effect of dis-
paraging statements is deeply related to the strictness of the law itself. A
speaker’sstatementsmayinformtheaudience’sbeliefsandactions.Inchoos-
ing whether to make disparaging statements, speakers will consider the ex-
pectedcostofapotentiallawsuitagainstthem.Stricterlawsincreasethiscost.
Thus, inequilibrium, thestrictnessofthelawaffectsspeakersand, anticipat-
ingthis,alsotargetsandaudiences.Theseeffectssometimesresultincounter-
intuitive implications, such that targets of speech who are ‘good’ types may
preferlaxerlaws,eventhoughitwouldlimittheirrecoveryinasuccessfullaw-
suit.Suchaconclusionispossiblebecausestrictlawsmakestatementsamore
costlysignal,andthus, amorereliableoneintheeyesofaudiencemembers.
Adeterminedspeakercouldabusethistrustandspreadfalsitieseffectively.
The second conclusion is closely related to the dynamics which we just
highlighted. We find that both very strict and lax laws have similar negative
informationalconsequences.Whenthelawislax,i.e.damagesarelow,speak-
ersfrequentlymisstatethetruthandaudiencesrelymoreontheirpriorsrather
than on statements (akin to babbling equilibria under cheap talk). However,
if defamation laws are very strict, i.e., expected damages are high, then this
may deter speakers from making even truthful assertions (“overpriced talk”).
Whereas truth is a defense to a lawsuit, the risk of judicial mistake may be
too great, and so speakers would refrain from sharing negative private infor-
mation.Therefore,overlystrictlawsdeprivetheaudienceofmeaningfulinfor-
Electronic copy available at: https://ssrn.com/abstract=3452662
-- 3
mation.1Thus,ouranalysisrevealsabasicinsightwithrespecttoregulationof
the information environment: Both cheap and overpriced talk can undermine
informationdissemination.
Third,ouranalysisilluminatestheimportanceofinstitutionalconsiderations
indesigninginformationregulatinglaws.Onekeyconsiderationisthecourt’s
subject-matter expertise and likelihood of delivering accurate judgments. If,
in a given area, judges can fairly accurately detect false statements, impos-
ingrelativelylargedamagesthatdeterfalsestatementscanleadtoseparating
equilibriawhereonlytruthfulstatementsaremade.Usingthelawtoregulate
informationcontinuestobeoptimalincaseswherecourtsdonothavetheac-
curacynecessarytoimplementseparatingequilibria,but,candetermostfalse
statementswithoutchillingtruthfulstatements.Whencourtsarelesscapableof
accuratelyadjudicatingstatements,thesocialcostofusingthecourtsystem—
operationalizedbylitigationcosts—iskeyindeterminingwhetherinformation
shouldberegulated.Eveninthesecases,whenthegainsfromfacilitatingbene-
ficialinteractionsanddeterringharmfulonesdwarfslitigationcosts,moderate
damagesemergeastheoptimalchoice.2
Anotherimplicationpertainstothepotentialdynamicimpactofinformation
regulating laws. Specifically, moderate laws that cause the audience to ratio-
nally rely on speakers’ statements broadens the gap between the frequency
withwhichtheaudienceinteractswithgoodtypesversusbadtypes.Thisnat-
urally increases the returns from being a good versus a bad type, thereby in-
centivizing individuals and firms to increase the quality of their products or
services.
Lastly, our comparison of public and private enforcement reveals the rela-
tivemeritsofprivateenforcement.Apublicagencymaybeabletocommitin
advancetoacertainlevelofenforcement.Whereasprivatepartiesarelessca-
pableofcommitment,theyenjoyanaturalinformationaladvantageregarding
the merit of the lawsuit, as they know their own type. Consequently, private
enforcement leads to more accurate litigation decisions, and an intuitive ad-
vantage of private enforcement emerges in our model: separating equilibria
canonlybeachievedthroughprivateenforcement.
Overall,ourframeworkandresultsaddtotheliteratureoninformationreg-
ulation by spotlighting the importance of audience effects, offering a formal
framework that accounts for audiences, and emphasizing the risks of overly-
stringentandlaxregulatoryregimes.
Thenextsectionofferssomebackgroundandreviewstherelatedliterature.
Section3presentsthemodelanditsanalysiswithafocusoncaseswherethe
courts are relatively accurate. Section 4 explains, in detail, the more compli-
cated trade-offs that emerge when courts are not accurate enough to achieve
separatingequilibria.Section5includesseveralpotentialextensionsoftheba-
1. Whenweconsiderhonestandothertypesofspeakers,wealsoshowthatstrictlawscanbe
worsethanlaxones,forsimilarinformationalreasons.
2. Incidentally,thisconclusioncanofferarationaletothelongstandingdistinctionindefama-
tionlawbetweenfactsandopinions,whicharegenerallyunregulated.
Electronic copy available at: https://ssrn.com/abstract=3452662
4 .V0N0
sicmodel,suchasthepublicenforcementcase,thegeneralizationofthemodel
tocaseswherespeakersmaybemotivatedtospeaktruthfullyortoexcessively
praisethetarget, anddiscussionsofcontextsotherthandefamationlaw.Sec-
tion6providesconcludingremarks.
2. Backgroundandrelatedliterature
Variouslawsregulateinformationbysanctioningfalsedisclosureorreward-
ing truthful sharing of information. Defamation is a classic example of the
formerandwhistle-blowersofthelatter.Theliteratureonthesetopicsisdis-
parate, but contain the same question: How to design sanctions and rewards
that would incentivize the optimal sharing of information. A common recur-
ring omission is the possibility that the audience may update its beliefs, in a
Bayesianmanner,basedonthesizeofthesanctionsorrewards.Becauseofthe
fragmentednatureoftheliterature,wewillconsiderfourexamples.
Defamationlawisperhapsthequintessentialexampleoftheproblemofin-
formationregulationandthusservesasourrunningexample.Underdefama-
tionlaw, atargetofa(1)publicstatementthatis(2)falseand(3)harmfulto
one’sreputation,cansueforallresultingdamages.Judgmentsinthisareacan
resultinhighpayments, withsomecasesreportingjuryjudgmentsoftensof
millions of dollars (Lesher v. Does, 2013). While courts and legislators un-
derstandthebehavioraleffectsofdefamationlaw,theyaremostlypreoccupied
withtheeffectofdefamationlawonspeakers’incentives(’chillingeffect’)and
victim’srights(Bar-Gill&Hamdani, 2003,Acheson&Wohlschlegel, 2018).
Consequently, they share a virtually axiomatic belief that stricter defamation
lawswouldbetterprotectvictims(McNamara, 2007).
Until very recently, scant attention has been given to the audience effects
ofdefamationlaw.Thisomissionissignificant, asdefamatoryspeechisonly
harmfulifitisbothbelievedandactedupon.Thefocusofeconomicworkin
thisareawasmediaoutlets,responsibleinvestigativejournalism,andpolitical
corruption(Garoupa, 1999a,b,Bar-Gill&Hamdani, 2003,andDalvi&Re-
falo, 2007).Weamplifyhereontwoinformalcontributionsthatrecognizethe
potentialimplicationsofaudienceeffects(Arbel&Mungan, 2019,Hemel&
Porat, 2019)byofferingaformalandbroaderaccount.
Another example of information regulation comes from the literature on
whistle-blowers,whichstudiestheoptimalrewardspaidtothewhistleblower.
There,aprimaryconcernisfalsereportsbythewhistle-blowerstoanenforce-
mentagency(Givati, 2016,Buccirossietal.2017,Deoorter&DeMot, 2005).
Onefindingisthatwhentheriskoffalsereportingishigh,itmightbeneces-
sarytoavoidrewardingwhistleblowersaltogether,eventhoughthismeansloss
of information. What is not accounted for is how the agency, the ”audience”
of the report, reacts to information, given the size of the reward. With large
rewards,theagencywouldbemorelikelytoexpectfalsereports.
Lawenforcementprovidesanotherillustrativeexample.Althoughthepolice
oftenhastoweighthecredibilityofacriminalactivityreport,thisrealityisnot
capturedinthestandardlawandeconomicsliterature(forareview,seePolin-
Electronic copy available at: https://ssrn.com/abstract=3452662
-- 5
skyandShavell2017),whichtypicallyreliesonmodelswheretheprobability
ofdetectionisonlyafunctionofenforcementexpenditures.Inreality,thepo-
lice seeks to economize resources by investigating more thoroughly reports
that appear credible—and its estimation is likely influenced by the sanctions
leviedagainstthosewhofilefalsereports.
A final example comes from securities regulation. There, a company self-
reports its performance, under an enforcement threat by the Securities and
ExchangeCommission(SEC).Theliteraturerecognizesthattheagency’sen-
forcementcanbeanimportantcredibilitymechanism(Stulz2009),butitpays
littleattentiontohowstrictenforcementinteractswithinvestorsandthetrust
theyplaceincompanydisclosure.
Methodologically, our article borrows tools from the rich literature on sig-
naling(Spence1973)andcheaptalk(Crawford&Sobel, 1982).Ouranalysis
canalsobeinterpretedaspartofanemergingliteraturethatlooksathowlaws
canbeusedtocreateinformalsanctionsthroughthebehaviorofthirdparties
(e.g.,Deffains&Fluet, 2019,Mungan2016,Be´nabou&Tirole, 2011,2006,
Rasmusen1996.)
3. Model
Tostudythebehavioraleffectsofinformationregulationwefocusontheex-
ampleofdefamationlaw,forthereasonsnotedintheintroduction.Wemodel
theinteractionsbetweenthreetypesofparties: thespeaker(S,she),thetarget
ofthespeech(T,he),andtheaudience,capturedbyarepresentativemember
(A,it).Afacesaninformationalproblem:T iseitheragoodorabadtype,and
A’svalueofinteractingwithT dependsonT’stype,whichisunknowntoA.
Before deciding, S, who knows T’s type, communicates with A and may ei-
therdisparageT ormakeanon-disparagingcomment.Asweareinterestedin
defamation,weassumethatS mightbenefitfromblockinganinteractionbe-
tweenAandT,andsoSmaychoosetodefameT–i.e.,liethatT isabadtype3.
Ofcourse,manyspeakersmaybemotivatedbyadesiretospeaktruthfullyor
tofacilitateinteractionsbetweenT andA,andweconsiderthispossibilityin
section5.2.
WemodeltheinteractionsasaBayesiangame,anduseittoidentifyPerfect
Bayesian Equilibria. Figure 1, below, depicts the interactions between these
threepartiesandishelpfulinfollowingthedetaileddescriptionsoftheinter-
actionsthatweprovide,next.4
3.1 PreliminaryNotation
We consider a game where T may be one of two types t ∈ {B,G} where
thelettersabbreviatebad andgood, respectively.T’stypeisprivatelyknown
to himself and S, but not to A, who only knows that the proportion of good
3. Consistentlywiththelaw,truthfulnegativestatementsarenotconsidereddefamatory.How-
ever,thecourtmaymakeerrorsinascertainingwhetheranegativestatementistruthful,andthis
possibilityisincorporatedinourmodel,asweexplainbelow.
4. ThefiguredoesnotdepictNature’sdrawofS’stype,duetoreasonsweexplain,below.
Electronic copy available at: https://ssrn.com/abstract=3452662
6 .V0N0
Figure1Extendedgametreeofthemodel.
Electronic copy available at: https://ssrn.com/abstract=3452662
-- 7
typesisγ ∈ (0,1).5 Apreferstointeractwithgoodtypes, butnotbadtypes,
because this results in a payoff of g > 0 > −b where b is the cost A bears
from interacting with a bad type. On the other hand, T always prefers to in-
teractwithAandobtainsabenefitofrfromtheinteraction.Finally,S hasan
interestinwhetherAandT interactandobtainsagainofv whentheydonot
interact(alternatively,v canbeinterpretedasalossincurredwhenAchooses
to interact with T); v is a random variable drawn from the continuum (0,v]
withthecumulativedistributionfunctionF(v).Thespecificv-drawisprivate
informationavailableonlytoS,andwecallv thespeaker’stype.Weassume
thatinteractionsbetweenAandT aresociallyvaluableif,andonlyif,T isa
goodtype,i.e.r+g >v >0>r−b.
AfterNaturedeterminesthetypesofT andS, T’stypebecomescommon
knowledge among T and S (but not A). At this point, S chooses what type
of statement to send A regarding T’s type. The types of possible statements
follow defamation law’s distinction between disparaging statements, which
are potentially actionable, and non-disparaging statements, which are non-
actionable(e.g.,positiveremarks,silence,opinion,etc.).6
Subsequently, A decides on whether to interact with T or to avoid him,
and, finally, T, decides whether to bring a lawsuit against S if a disparaging
remarkwasfollowedbyA’schoicetoavoidinteractingwithT.Wenotethat
this setting includes the possibility of T suing S, even if T is in fact a bad
type,i.e.,afrivolouslawsuitmaybebrought.Thisisanimportantpossibility
becausecourtsmayerrintheirjudgment.7 Tocapturetheparties’payoffs,we
definethefollowing:
d: damagespaidbyS toT whenthecourtfindsforT.
l: totallitigationcosts.Weassumethatlitigationcostsarenotprohibitive(l<v)
and,withoutlossofgenerality,thatthecostsareequallysharedbytheparties.
q : probabilityofplaintiffvictorywhenT isoftypet∈{B,G}.
t
Figure1summarizestheparties’payoffsattheterminalnodesonthebottom
in the order S,A,T. We note two graphical limitations of Figure 1. First, it
doesnotshowinformationsetsdescribingA’sknowledgeregardingS’stype,
due to the depiction difficulty caused by S drawing her type from a contin-
uum.Second,foreaseofexposition,Figure1doesnotdepictNature’svdraw
determiningS’sinclinationtodisparage.
Weusethevariabledasthekeypolicylevertooperationalizedifferentkinds
ofdefamationlawregimes, since, aswenoteintheintroduction, courtshave
verywidediscretioninsettingdamages.dcanbeinterpretedmostdirectlyas
thelevelofdamagesawardedtoavictimofdefamation.Thecasewhered=0
representsasituationwheretherearenodamagesfordefamation,whichis,in
oursetting,functionallyequivalenttonodefamationlaw.
5. Insection5wediscusstheconsequencesofendogenizingγ.
6. Weexplainhowouranalysisinformsthediscussionofwhattypesofstatementsoughtto
giverisetodefamationclaimsinSection5.4
7. TherequirementofharmmakesalawsuitbyT whenthepartiesdointeractunlikely
Electronic copy available at: https://ssrn.com/abstract=3452662
8 .V0N0
Itisalsoworthnotingthatwetaketheoddsofwinningattrial,q andq ,as
B G
exogenouslygiven.Thisimpliesthatcourtsarecommittedtoreviewingcases
onlyontheirmerit,andwithoutbringingintheirinformedestimatesaboutthe
proportion of frivolous cases. This is a standard commitment assumption in
the enforcement literature, and with it in place, the ratio between q and q
B G
correspondstoameasureofjudicialaccuracy.Inouranalysis,below,wevary
thisparametertostudyhowitaffectsequilibria.
3.2 Players’Actions,BeliefsandStrategies
Next,wedescribetheplayersstrategies,beliefs,andactions.Forsimplicity,
eachplayer’sactionislabelledaseither0or1,asfollows:
Table1Players’PotentialActions
Player Action
0 1
S Don’tDisparage Disparage
A Interact Avoid
T Don’t Litigate
Thus,asindicatedinFigure1,asuitisfiledonlyincaseswhereallplayers’
actionsare1.Usingthisnotation,wecandescribethestrategiesofeachplayer
asfollows:
Table2Players’Strategies
Player Strategy
S s(t,v):{B,G}×(0,v]→{0,1}
A a(z):{0,1}→{0,1}
T p(t):{B,G}→{0,1}
Here,inspecifyingA’sstrategy,zdenotesthestatementreceivedbyA.
InordertoidentifyPerfectBayesianEquilibria(henceforth:PBE),wespec-
ifyA’sbeliefsregardingT’stype,as:8
x =BeliefthatT =ggivenz =0
0
(1)
x =BeliefthatT =ggivenz =1
1
3.3 PerfectBayesianEquilibriumRequirementsandDefinitions
As our solution concept is PBE, there are standard requirements that must
be satisfied. Specifying these necessitates references to the conditional prob-
abilitiesofT beingaparticulartype, givenastrategyplayedbyS.Todoso,
we first specify the unconditional (or ex ante) probability with which S will
8. BecauseA’svaluationofhisinteractionwithT dependsonlyonT’stype, weneednot
specifyA’sbeliefsregardingS’stypeforpurposesofidentifyingthePBE.
Electronic copy available at: https://ssrn.com/abstract=3452662
-- 9
disparageT givenanystrategy,s,asfollows:
(cid:90) v
µ(s)≡ [γs(G,v)+(1−γ)s(B,v)]dF(v) (2)
0
When µ(s) ∈ (0,1), we can use Bayes’ rule to calculate the probability
ofT’stype, goodorbad, conditionalonthestatementmadeaboutT.Onthe
otherhand,whenµ(s) ∈ {0,1},itfollowsthatS isplayingastrategywhere
he(almost)always avoidsdisparaging(0)ordisparages(1) T, inwhichcase
Bayes’rulecannotbeusedtocalculatetheprobabilityofT beingaparticular
type, conditional on the strategy which is (almost) never played by S. Thus,
wedenotebothpossibilities,asfollows:
(cid:82)vs(G,v)dF(v)
γ 0 if µ(s)(cid:54)=0
Γ(t=G|1,s)≡ µ(s) (3)
Υ otherwise
(cid:82)v(1−s(G,v))dF(v)
γ 0 if µ(s)(cid:54)=1
Γ(t=G|0,s)≡ 1−µ(s) (4)
Υ otherwise
Here, the symbol Υ indicates that the strategy in question is (almost) never
chosenbythespeaker.
GiventhisnotationwemaycharacterizePBEasanassessmentconsistingof
thestrategyprofilea∗,s∗ andp∗ alongwithasetofbeliefsx∗ andx∗,which
0 1
satisfiesthefollowingfourrequirements.
Requirement 1 (R1): A has no profitable deviation given its beliefs. Let
x≡ b ,then:
(cid:98) g+b
a∗(z)=0 if x >x forz ∈{0,1}
z (cid:98) (5)
a∗(z)=1 if x <x forz ∈{0,1}
z (cid:98)
Here, x represents A’s risk threshold for engaging with T. Requirement 1
(cid:98)
states that A interacts with T only if A believes, given S’s statement, that
the probability that T is a good type exceeds the threshold probability of x.
(cid:98)
Similarly, if A believes that T is a good type with a probability that is lower
thanx, AdoesnotinteractwithT.Intheexceptionalcasewherex = x, A
(cid:98) z (cid:98)
isindifferentbetweeninteractingwithT andnot,and,thusitmayplayeither
strategy. When γ = x, this possibility is realized in all equilibria where the
(cid:98)
audience disregards newly acquired information in forming its beliefs. This
unnecessarilycomplicatesformalderivationsandmakesexpositionsmoredif-
ficult.Therefore,intheremainderofouranalysisweignoretheseexceptional
casesbyassumingthatγ (cid:54)= x,buttheanalysiscaneasilybeextendedtothis
(cid:98)
case.
Requirement2(R2):T hasnoprofitabledeviationsinsub-games:
0 if q d<l/2 fort∈{B,G}
p∗(t)= t (6)
1 if q d>l/2 fort∈{B,G}
t
Requirement2statesthatthePBEstrategyofT mustbesuchthatinsubgames
whereS disparageshim, T litigateswheneverthecostsofdoingso(l/2)are
Electronic copy available at: https://ssrn.com/abstract=3452662
10 .V0N0
lower than the expected damage rewards that he can obtain from litigation.
Conversely,T choosesnottolitigatewhenthecostsarehigherthanexpected
damages. In the exceptional case where q d = l/2, T is indifferent between
t
litigatingandnot.
Requirement 3 (R3): S has no profitable deviations: For all t,v pairs,
s∗(t,v)maximizesplayerS’spayoff,whichcanbeexpressedas
l
U ≡a∗(s(t,v))(v−p∗(t)s(t,v){q d+ }) (7)
S t 2
The requirement with respect to S appears more complex than the re-
quirementsthatpertaintoT andA’sstrategies,becauseS choosesheractions
in anticipation of the other players’ actions. Still, the requirement is simply
that,givenherowntype,T’stype,andtheanticipatedbehaviorofAandT,S
mustchoosethecourseofactionthatwouldmaximizeherpayoff.
Requirement4(R4):A’sbeliefsareconsistent:
x∗ =Γ(t=G|z,s∗)wheneverΓ(t=G|z,s∗)(cid:54)=Υforbothz ∈{0,1} (8)
z
Requirement 4 simply states that A’s beliefs must be consistent with the im-
plied conditional probability of T being a particular type based on the equi-
libriumstrategyofS.Thisrequirementisapplicableonlytostrategieswhich
haveapositiveprobabilityofbeingplayedbyS.
Ouranalysisrevealsthattherearetwotypesofassessmentswhichsatisfyre-
quirements1-4,i.e.twotypesofequilibrium.One,inwhichthespeaker’sstate-
mentshavenobearingontheaudience’sbehavior,inthesensethattheydonot
causetheaudiencetochangetheirbehaviorrelativetowhattheywouldhave
doneiftheyreliedonlyontheirpriors.Becausethespeaker’sstatementhasno
effectonaudience’sbehavior,wetermthesePBEIneffectiveCommunication
Equilibria. By contrast, when statements may affect behavior, the resulting
PBE are dubbed Effective Communication Equilibria. To avoid any ambigu-
ities in our usage of these terms, we define these two types of equilibria, as
follows.
Definition 1: A PBE is an effective communication equilibrium if, and only
if,thereexistsz ∈{0,1}suchthata∗(z)= xˆ−min{γ,xˆ} andµ∗(s∗)(cid:54)=1−z.
xˆ−γ
In classifying equilibria, we use these new definitions, instead of concepts
like babbling equilibria and informative equilibria, because, although these
concepts are related to our defined categories, they differ from each other in
meaningfulways.Specifically,althoughallbabblingequilibriaareineffective
communication equilibria, the converse is not true. This can be seen by not-
ing that, in some equilibria, S can play type-dependent strategies which do
not impact the behavior of A. These equilibria would not fit the definition of
babblingequilibria,butwouldnotcauseachangeinA’sbehaviorcomparedto
babblingequilibria.Sinceweareinterestedinclassifyingequilibriabasedon
Electronic copy available at: https://ssrn.com/abstract=3452662
-- 11
behavior,werelyonourbehavior-baseddefinitionofeffectivecommunication
equilibria.
3.4 ImpactofDefamationLawsonEquilibriumBehavior
ByusingRequirements1-4weidentifyandinterpretthePBEobtainedwith
different damages, through the help of four propositions, below. Our obser-
vations can be briefly summarized as follows. Proposition 1 shows that, re-
gardlessofthelevelofdamages, therearealwaysineffectivecommunication
equilibria where A acts according to its priors, i.e., where A essentially ig-
noresthecontentofSsstatement.Intheseequilibria,partiescannoteffectively
communicateprivateinformation.Infact.whendefamationlawsareextreme,
i.e. either too lax or too strict, ineffective communication equilibria are the
onlyPBEofthegame, aswenoteviaProposition2.Onlymoderatedefama-
tionlawscanengendereffectivecommunicationequilibria.Then,wequestion
whethereffectivecommunicationequilibriaaresociallypreferabletoineffec-
tiveones.Theanswertothisquestionissurprisinglyambiguousanddepends
in part on the accuracy of the courts. Proposition 3 shows that when courts
are sufficiently accurate, it is possible to set damages moderately such that
defamatorystatementsarefullydeterred,withoutinvitingfrivolouslitigation.
Thus, separating equilibria are obtainable, and they are socially preferable to
anyotherequilibria.Proposition3alsonotesthatevencourtswhicharefairly
accurate, but not accurate enough to facilitate separating equilibria, can en-
hancewelfarethroughmoderatedamagesthroughsemi-separatingequilibria.
Finally, Proposition 4 reveals that when the value of A’s returns from inter-
actionsdwarfsotherconsiderations,PBEassociatedwitheffectivedefamation
lawsarealwayssociallypreferable.
Proposition 1. (i) Under all defamation regimes, there exists ineffective
communication equilibria. (ii) In these equilibria, A either always interacts
(γ > x) or never interacts (γ < x) with the target, and litigation never takes
(cid:98) (cid:98)
place.
Proof. (i)Theassessmentconsistingofx∗ =x∗ =γ,
1 0
0forallz if γ >x
a∗(z)= (cid:98) ;
1forallz if γ <x
(cid:98)
s∗(t,v)=0forallvandt;
0 if q d(cid:54)l/2 forallt∈{B,G}
andp∗(t)= t satisfiesRequirements
1 if q d>l/2 forallt∈{B,G}
t
1-4,andthusconstitutesaPBEwhereAactsbasedonitspriors.
(ii)Bydefinition,inineffectivecommunicationequilibriaAactsaccordingto
itspriors, and, thusitalwaysinteractsifγ > xandneverinteractsifγ < x.
(cid:98) (cid:98)
In the former case, litigation never takes place as there is always interaction.
In the latter case, if a∗(0) = 1, S could profitably deviate from her strategy
byneverdefamingsincethiswouldsaveherlitigationcosts.Thus,itmustbe
thecasethata∗(0) = 0,whichispossibleonlyifµ(s∗) = 1since,bydefini-
tion,interactionnevertakesplace.ButthenS canprofitablydeviatefromher
Electronic copy available at: https://ssrn.com/abstract=3452662
12 .V0N0
strategy s∗ by choosing not to defame whenever t = G and v < q d+ l.
G 2
Thus,litigationcannotbetakingplaceinanineffectivecommunicationequi-
librium.
Proposition 1 reveals that it is always possible in equilibrium for the audi-
encetoactaccordingtoitspriors.GiventhisresponsebyA,S hasnothingto
gainbydisparagingthepotentialplaintiff,becauseherstatementshavenoef-
fectonA’sbehavior,yetitmaycauseT toinitiatealawsuit.Thus,nolitigation
canbeobservedinsuchequilibria.
Next, we turn to the question of whether defamation laws can cause A to
change its behavior relative to its behavior based on its priors. Because the
answer to this question depends on d, the magnitude of damages, it is worth
identifyingfourcriticaldamagelevelswhichplayakeyroleintheinterpreta-
tionofresults.Figure2belowdepictstheselevels.
Figure2Criticallevelsofdamages
Theupperlinedepictsthefirsttwolevels( l and l )whichrelatetothe
2qG 2qB
potentialplaintiff’sincentives,whereasthesecondlineincludestheothertwo
levels (2v−l and 2v−l) which relate to the speaker’s incentives. These levels
2qG 2qB
aredepictedontwoseparatelinesbecause,absentfurtherassumptions,twoof
thesevalues(namely l and 2v−l)cannotbeunambiguouslyranked.Wecan,
2qB 2qG
however,notethatthecriticalvaluesthatrelatetothespeaker’sincentivesare
greaterthanthecorrespondingcriticalvaluesthatrelatetothetarget’sincen-
tives(i.e. l < 2v−l fori∈B,G),givenourassumptionthatlitigationcosts
2qi 2qi
arenotprohibitivelyhigh,i.e.l<v.
We observe that when damages are low, i.e., d < l , as Figure 2 notes
2qG
T lackstheincentivestobringsuitevenwhenheisfalselydisparaged.When
damagesareveryhigh,i.e.,d > 2v−l,itfollowsthatalleffectivedisparaging
2qB
Electronic copy available at: https://ssrn.com/abstract=3452662
-- 13
statementsaredeterred.9Thus,inneithercasedostatementshaveanimpacton
theaudience’sbehavior.Wedistinguishbetweentheseextremedamages(i.e.,
(cid:104) (cid:105) (cid:104) (cid:105)
d (cid:54)∈ l ,2v−l )andmoderatedamages(i.e., d ∈ l ,2v−l .)Theabove
2qG 2qB 2qG 2qB
observationshighlightthatextremedamagescanonlyleadtoineffectivecom-
munication PBE. A question that remains is whether moderate damages can
lead to effective communication equilibria. Proposition 2 answers this ques-
tionaffirmativelyandformalizesrelatedobservations.
Proposition2. (i)Extremedefamationlawsonlygenerateineffectivecom-
municationequilibria.(ii)Effectivecommunicationequilibriacanbeobtained
only when the audience acts consistently with the speaker’s statement, i.e.
a∗(z) = z. (iii) There are moderate defamation laws which generate effec-
tivecommunicationequilibria.
Proof. SeeAppendix.
Proposition2holdsthatextremedefamationlawsonlyallowforineffective
communicationequilibria,and,asnotedinproposition1,theseequilibriaalso
exist under moderate defamation laws. However, moderate defamation laws
alsogenerateeffectivecommunicationequilibria.Thisimpliesthatswitching
fromanextremedefamationlawregimetoamoderateregimecanexpandthe
typesofequilibriathatmaybeobtained.Thus, itbecomesimportanttocom-
parethepropertiesofthetwotypesofequilibriatoascertaintheirwelfareim-
pacts,amongotherthings.Thiscomparisonhingesonhowaccuratethecourt
isinreturningcorrectverdicts.Byaccuracy,wemeanthefollowing:
Definition2(i) qG ∈(1,∞)measuresthecourts’accuracy.(ii)π ≡ 2v−1is
qB l
acriticallevelofcourtaccuracyusedtoevaluatethepotentialwelfareimpacts
ofdefamationlaws.
Wereporttherelationshipbetweenthecourt’saccuracy,asdefinedabove,and
thePBEobtainable,asfollows.
Proposition3. (i)SeparatingEquilibrium:Whenthecourtissufficientlyac-
curate(i.e. qG (cid:62) π)therearemoderatedefamationlawsassociatedwithPBE
qB
where: S disparagesT if,andonlyif,heisabadtype;theaudienceactscon-
sistentlywiththisinformation(i.e.a∗(z) = z); andthereisnolitigation.(ii)
Separating equilibria lead to greater expected welfare than all other equilib-
ria.(iii)Whenthecourtisinsufficientlyaccurate(i.e. qG < π), allequilibria
qB
involve a positive likelihood with which the audience does not interact with
a good type, interacts with a bad type, or both. (iv) When the court is only
9. Weintentionallyrefertothedeterrenceofeffectivedisparagingstatements,becausethere
couldbeequilibriawheretheaudiencedisregardsdisparagingcommentsandinteractswithT,
and,insuchinstances,disparagingcommentswouldnotbedeterredbecausetheywouldnotgive
risetolitigation.
Electronic copy available at: https://ssrn.com/abstract=3452662
14 .V0N0
slightly inaccurate, i.e. π − qG > 0 is sufficiently small, there exist moder-
qB
atedefamationlawswhichgenerateequilibriathatleadtogreaterwelfarethan
thosegeneratedbyineffectivecommunicationequilibria.
Proof. SeeAppendix
Intuitively, when courts are sufficiently accurate it ought to be possible to
set damages large enough to deter defamatory statements without generating
frivolous lawsuits. When qG (cid:62) π, this is in fact the case, because the am-
qB
biguousrankingbetweenthecriticaldamagelevelsdepictedinFigure2, l
2qB
and 2v−l, vanishes, and it follows that 2v−l < l , as depicted in Figure 3
2qG 2qG 2qB
below. Therefore, by choosing damages in between these two threshold val-
(cid:16) (cid:17)
ues,i.e.d∈ 2v−l, l ,onecanachievetwoimportantgoalsatonce: deter
2qG 2qB
defamationaswellasfrivolouslawsuits.
Figure3Criticallevelsofdamages
Separating equilibria that achieve these two goals at once naturally maxi-
mizewelfare,because(1)theyleadtointeractionsonlywhentheseinteractions
enhance welfare and (2) there are no litigation costs. This reasoning extends
to the case where the court is only somewhat accurate through a simple con-
tinuity argument. In this case, moderate defamation laws are associated with
semi-separatingequilibria,whereinanon-disparagingcommentrevealsthatT
is a good type, but where good types face a very small likelihood of being
disparaged.Theseequilibrialeadtoonlyslightlylowerexpectedwelfarethan
separating equilibria and, thus, are associated with greater expected welfare
thanineffectivecommunicationequilibria.
Propositions 1-3, together, reveal that when courts perform well in distin-
guishinggoodandbadtypes,moderatedefamationlawscanbeusedwithrel-
ativeeasetoenhancewelfareandtoincreasetheinformationalvalueofstate-
ments made by speakers. In these cases, (semi-)separating equilibria lead to
obviousandunambiguousimprovementscomparedtoequilibriawheretheau-
dienceislefttouseitspriorstomakedecisions.Inpractice,however,thereare
manycaseswherethereisexpressedconcernamongjudgesandlawyersthat
discoveringthetruthisdifficultandthatlitigationisfraughtwithinaccuracies.
Theanalysisinthenextsectionthusfocusesonthesesituations.
Electronic copy available at: https://ssrn.com/abstract=3452662
-- 15
4. DynamicswhenCourtsareInaccurate
The previous section explained why it is impossible to obtain separating
equilibriaifcourtsareinaccurate.AsnotedinProposition3,thisimpliesthat
withsomepositiveprobabilityeitherinteractionswithgoodtypesaredeterred
(i.e.type-1errors),interactionswithbadtypesareundeterred(type-2errors),
orboth.Inthesecases,usingstricterdefamationlaws(i.e.higherd)cangener-
ateatrade-offbetweencostsassociatedwiththesetwotypesoferrorsandmay
alsoimpactexpectedlitigationcosts.Inthissection, wedescribethesetrade-
offs. We focus exclusively on moderate defamation laws and the impact of
changingdoneffectivecommunicationequilibriabecause,asnotedinPropo-
sition 3, in all other cases the audience acts according to its priors and no
litigation takes place. Subsequently, we identify a sufficient condition under
which achieving effective communication equilibria through moderate dam-
agescontinuestobesociallypreferabletohavingextremedefamationlaws.
Toexplainthedynamicsthatemerge,wefirststartbycalculatingtheequi-
libriumbeliefs,i.e.x∗ andx∗ thatwouldemergeinaPBEwherea∗(z) = z,
0 1
assuming that such an equilibrium exists. We plot these beliefs in Figure 4,
below,throughaspecificbutrepresentativeexample.Thehorizontalaxisrep-
resentsdamages, onwhichwemarkthefourcriticaldamageslevelslistedin
Figure2.Thistime,however,thecourt’saccuracyislowerthanπ,sotherank-
ingoftheintermediatecriticaldamages(i.e. l and 2v−l)istheoppositeof
thatdepictedinFigure3.Inadditiontoplottin 2 g qB beliefs, 2 i q . G e.,x∗andx∗,inFig-
0 1
ure4wealsoplottheex-anteprobabilityofT beingdisparagedinthesePBE.
Thesearelabeledδ andδ forgoodtypesandbadtypes,respectively.Next,
G B
we explain how these expressions are derived in the three relevant ranges of
damages.
(1) In the range ( l , l ), damages are too low to incentivize bad types
2qG 2qB
to sue. Thus, S faces no consequences from disparaging bad types. Whereas
goodtypeswillbringalawsuit,Smightstilldisparagethemifitsbenefitfrom
blockinganinteractionissufficientlyhigh, i.e., v > v ≡ q d+ l.Thus, a
G G 2
badtypeisdisparagedwithcertainty,i.e.δ =1,andagoodtypemayormay
B
not be disparaged with positive probability. From A’s perspective this means
thatapersonwhoisnotdisparagedisdefinitelyagoodtype,i.e.x∗ =1,while
0
a target who is disparaged may or may not be a good type, but is no more
likelytobeagoodtypethanarandomdrawfromthepopulation,i.e.x∗ < γ.
1
The ex-ante probability with which S draws a benefit that is higher than v
G
isδ = 1−F(v ),and,thus,thisistheprobabilitywithwhichagoodtype
G G
is disparaged. Using this expression, x∗ can be more precisely expressed as
1
x∗ = γδG <γ.
1 γδG+1−γ
(2)Intherange( l ,2v−l),damagesaresufficienttotriggerfrivoloussuits
2qB 2qG
by bad types who are disparaged. The threat of a suit causes the speaker to
refrainfromdisparagingevenabadtype,unlessherbenefitfromblockingan
interactionissufficientlyhigh.Still,theminimumbenefitthatleadsaspeaker
to disparage a bad type, v ≡ q d+ l, is lower than the minimum benefit
B B 2
thatwouldmakeherdisparageagoodtype,v =q d+ l,asfrivolousclaims
G G 2
Electronic copy available at: https://ssrn.com/abstract=3452662
16 .V0N0
Figure4. IllustrationofBeliefsandtheLikelihoodofaDisparagingStatement.
Damages=d,x∗
0
,x∗
1
arebeliefs.qG=0.8,qB =0.2,l=0.3,andF(v)=vwithsupport
(0,1].
are less likely to be successful. Thus, the ex-ante probability with which S
disparages a bad type, δ = 1−F(v ), is greater than the likelihood with
B B
which she disparages a good type, δ = 1 − F(v ). Consequently, in this
G G
range,x∗ = γ(1−δG) >γ >x∗ = γδG asδ <δ .
0 γ(1−δG)+(1−γ)(1−δB) 1 γδG+(1−γ)δB G B
(3)Intherange(2v−l,2v−l),damagesaresufficientlyhightodeterS from
2qG 2qB
disparaginggoodtypes,evenifherbenefitfromblockinginteractionsismaxi-
mal,i.e.v.Shewillonlydisparagebadtypesifherbenefitfromblockinganin-
teractionissufficientlyhigh.Thus,inthisrange:x∗ =0<γ < γ =
x∗andδ =0<δ =1−F(v ).
1 γ+(1−γ)(1−δB)
0 G B B
This brief analysis, and its depiction in Figure 4 can be used to identify
someofthewelfareimplicationsofalteringthelevelofdamages.Inthelower-
moderate range (i.e. ( l , l )) damages are insufficient to completely pre-
2qG 2qB
ventdisparagingremarksagainstgoodtypes,buttheyarealsolowenoughto
deter frivolous litigation by bad types, leading to their disparagement. Thus,
in this range, the only impact of increasing damages is to reduce the number
ofgoodtypesbeingdisparaged.Thisreductionconsequentlyreducesexpected
litigation costs, and increases the likelihood of interactions with good types.
Therefore,increasingdamagesinthisrangemonotonicallyenhanceswelfare,
becauseinteractionswithgoodtypesaresociallydesirable,andlitigationcosts
reducewelfare.
Intheintermediate-moderaterange(i.e.( l ,2v−l))damagesarelargeenough
2qB 2qG
toinducefrivolouslitigationbybadtypes,butnotlargeenoughtocompletely
deterdisparagingremarksagainstgoodtypes.Therefore,increasingdamages
Electronic copy available at: https://ssrn.com/abstract=3452662
-- 17
in this range generates more meaningful trade-offs by increasing the likeli-
hoodofbeneficialinteractionsaswellasharmfulinteractions,whilereducing
the likelihood of litigation. Thus, it is desirable to increase damages in this
rangeonlyifthesavingsfromlowerlitigationcostsandtheincreasedvalueof
beneficialinteractionsexceedthecostinvolvedwithharmfulinteractions.Ab-
sent more restrictive assumptions, one cannot unambiguously compare these
benefitsandcosts, becausetheirmagnitudesdepend, inpart, onthemarginal
changesinδ andδ ,whichcantakemanyformsdependingontheshapeof
B G
thedistributionofspeakerbenefits(i.e.F(v)).
Finally,inthehigher-moderaterange(i.e.(2v−l,2v−l)),damagesarehigh
2qG 2qB
enoughtodeterdisparagingcommentsagainstallgoodtypes,butarenotsuf-
ficiently high to deter disparaging remarks against bad types. An increase in
damages in this range causes an increase in the expected costs from harmful
interactions, but reduces litigation costs. Thus, as long as litigation costs are
lowerthanthegainsfromblockingharmfulinteractions,socialwelfareisim-
provedbyreducingdamagesinthisrange.
This analysis reveals the complex nature of trade-offs involved when the
courtisinaccurateinmakingdecisions.Thereisnogeneralreasonwhyhigher
damageswouldbebetterthanlowerdamages.Courtsandpolicymakersmust
account for domain-specific considerations which can tilt the balance in any
givendirection.
Asomewhatcounterintuitiveconclusionisthat,withinaccuratecourts,itis
noteventrueingeneralthatonecanimproveuponineffectivecommunication
equilibria where the audience acts upon its priors. Moving to an equilibrium
wheretheaudienceactsconsistentlywiththeinformationitreceivesfromthe
speakercanbehelpfulinpromotingbeneficialinteractionsordissuadingharm-
fulones.However,itcomesatthecostofincreasedlitigation,andmayreduce
thetarget’sbenefitfromincreasedmissedinteractionsorthespeaker’sbenefit
from blocking interactions. An aspect of this analysis is that higher damages
inthemoderaterangesometimessacrificethewell-beingofsomegoodtypes,
thus calling into question a widely-held belief among lawyers that stronger
defamationlawsprotectgoodtypes(Arbel&Mungan, 2019,Hemel&Porat,
2019).
Adding to these complexities is the fact that, given any damage level, d,
effective communication PBE are possible only if A’s risk tolerance, xˆ, lies
inbetweenx∗ andx∗,asdepictedinFigure4.Despitetheseambiguities,one
1 0
canalwaysusemoderatedamagesthatleadtoeffectivecommunicationPBE.
Thus,onecanimprovetheoddsofbeneficialinteractionstakingplaceand/or
harmfulinteractionsnottakingplace.Thus,iftheaudience’swell-beingisthe
predominant consideration in the welfare analysis, it follows that moderate
damages can always improve upon extreme damages. The next proposition
formalizesthisresult.
Proposition4. Thereexistmoderatedamagesleadingtoeffectivecommu-
nicationequilibria,whichgenerategreaterwelfarethanineffectivecommuni-
cationequilibria,aslongasgandbarelargerelativetoothercostsandbenefits.
Electronic copy available at: https://ssrn.com/abstract=3452662
18 .V0N0
Proof. Theexpectedpay-offoftheaudienceinanequilibriumwherea∗(z)=
zforallzis
U =γ(1−δ )g−(1−γ)(1−δ )b (9)
A G B
Ontheotherhand,0andγg−(1−γ)baretheexpectedpay-offsthattheau-
diencewouldhavereceivedbyactingaccordingtoitspriors,whenγ <xand
(cid:98)
γ >x,respectively.InthesePBE,itfollowsthatU =γ(1−δ )g >0when
(cid:98) A G
d∈( l , l ),and,similarly,U =γg−(1−γ)(1−δ )b>γg−(1−γ)b
2qG 2qB A B
when d ∈ (2v−l,2v−l). Thus, for any given x, the increase in the expected
2qG 2qB (cid:98)
pay-offoftheaudiencestemmingfromamovefromaPBEwhereitactsac-
cordingtoitspriorstoonewhereitactsaccordingtotheinformationitreceives
fromthespeakerislinearlyincreasinging andb,respectively.Moreover,the
magnitudes of g and b only affect A’s payoff, and, hence, there exist large
enoughg andbwhichcausethesePBEtogenerategreaterwelfarethanPBE
wheretheaudienceactsaccordingtoitspriors.
Proposition 4 reveals that when the value of interactions are large in com-
parison to other considerations, like litigation costs and the benefits that the
speakergetsfromblockinginteractions,moderatedefamationlawscanbeused
toenhancewelfare.Thisisbecause,undertheseconditions,thedominantcon-
siderationbecomesthemaximizationoftheaudience’spay-off,whichbenefits
fromhavingeffectivecommunicationequilibria.
5. Discussion
In Sections 3 and 4, we provided a model that allowed us to clearly focus
ondefamationlaws’impactontheaudience’sequilibriumbeliefsandactions.
Indoingso,weabstractedfrommanyissuesthatbearontheregulationofin-
formationinmoregeneralsettings,particularly,thepossibilityofacommitted
publicenforcer,qualitybeingendogenouslychosenbythetarget,andtheexis-
tenceofhonestandothertypesofspeakers.Hereweturnourattentiontothese
issues.
5.1 EndogenousTypesandDynamicEfficiencies
Inouranalysisthusfar,weassumedthatthetarget’stypetwasexogenously
determinedbynaturetobeeitherGorBwithprobabilitiesγand1−γ,respec-
tively.Onemightquestiontherealityofthisassumption,aspeoplecanmake
investmentsthatwouldmakethembetterorworsetradingpartners, e.g., cre-
atehigherqualityproducts,maintainsafetystandards,orkeephigherhygiene
standards.
One option of incorporating quality investments into our analysis is to re-
place Nature’s choice of types with a preliminary stage where the target, T,
makes a costly investment (c) that can increase her likelihood of becoming
a good type. Formally, we may assume that γ = γ(c) with γ(cid:48) > 0 > γ(cid:48)(cid:48),
limγ(cid:48)(c) = ∞, γ(0) = γ and limγ(c) = γ where1 > γ > γ > 0.More-
c→0 c→∞
over,tokeepthedescriptionofthisextensionbrief,wefocusonthecasewhere
Electronic copy available at: https://ssrn.com/abstract=3452662
-- 19
γ >x.
(cid:98)
Thequalityinvestmentdecisionisnowpartofalargergame.Givenanysub-
gameequilibrium,thebestresponseofT istomakeaninvestmenttomaximize
hisexpectedpay-off,whichcanbedenotedasγ(c)m +(1−γ(c))m −c
G B
wherem andm refertothepay-offsheobtainsinthesub-gameequilibria.
G B
This observation reveals a very clear result: When the laws are extreme,
(cid:104) (cid:105)
i.e.d(cid:54)∈ l ,2v−l ,thetargethasnoreasontoinvestinquality.Thisfollows
2qG 2qB
fromPropositions1&2,whichshowthatwithextremelaws,theaudienceacts
basedonitspriorsandinteractswiththetargetifγ issufficientlyhigh.Thus,
investmentshavenoprivatereturnsforthetarget.
Itisonlywhenthelawsaremoderatethattargetsmayhaveanincentiveto
invest in quality. This can be demonstrated by focusing on the lower bound
of intermediate damages, i.e. l . In this case, in PBE with a∗(z) = z, it
2qG
followsthatm =0(asallbadtypesaredisparaged)whilem =(1−δ )r
B G G
(because good types are disparaged with probability δ , in which case there
G
isalawsuitwhichpaysthetargetexpecteddamagesequaltolitigationcosts).
Thus,thetarget’spay-offisγ(c)(1−δ )r−c,and,therefore,thetargetprofits
G
(inexpectation)frominvesting.Whetherthisissociallygoodorbad,depends,
ofcourse,onwhethertherearenetsocialgainsfromsuchinvestments.Inour
context, this is socially valuable as long as the expected benefits from good
interactions ((1−δ )g)—which are not internalized by T—are greater than
G
the expected litigation costs l and the loss of benefit to S from blocking an
interaction, i.e. (1 − δ )E[v|v > l]. In fact, if investments in quality are
G 2
sociallyvaluable,thenincreasingdamageswithintheintermediaterangeupto
l will be desirable. This is because these higher damages lead to a lower
2qB
probability of disparaging remarks made against good types (as illustrated in
Figure 4) and, thus, increase m , while still keeping expected payoffs from
G
beingabadtypeatm =0.
B
Thediscussionherehighlightstheimportanceofinformationregulationfor
broadermarketdynamics.Theintuitionunderlyingourresultsarestraightfor-
ward.Extremelawsleadtoineffectivecommunicationequilibria.Incontrast,
moderate laws create an environment with more reliable information regard-
ing types, thus generating a greater gap between the payoffs obtainable by
goodtypesversusbadtypes.This,inturn,increasesthereturnsfrombeinga
goodtype,andleadstomoreinvestments.Inrealisticsettings,providingsuch
additionalincentivesissociallydesirablewhenthepotentialinvestorisunder-
incentivized due to problems like information asymmetries. The gains from
suchinvestmentsinqualityshouldbeaddedtotheotherbenefitsofmoderate
lawsthatwehaveidentified.
5.2 TruthSpeakersandEulogists
So far, we only considered speakers who had something to gain from sev-
ering the relationship between the audience and the target. This abstraction
followstheideaofspeaker’s’bias’inthecheaptalkliterature.Inreality,how-
ever,somespeakersmaynothavesuchmotivations.Quiteimportantly,many
Electronic copy available at: https://ssrn.com/abstract=3452662
20 .V0N0
people, when asked their opinion, provide an honest assessment of others.
Moreover,therearealsopeoplewhoaremotivatedbydoingtheexactopposite
of what the speakers in our model are motivated by; namely, promoting the
relationship between the target and the audience. In what follows we distin-
guishbetweenthefirsttype, “truthspeakers,”thelattertype, “eulogists,”and
theonesweformerlydiscussedinSection3as“disparagers.”Webriefly,and
informally, explain now what occurs when these kinds of speakers are incor-
poratedintoouranalysis.
Inourdiscussion,weconceiveofthesetypesasfollows.Disparagers,aswe
noted,receiveapositivevaluefromblockinganinteraction;truth-speakersare
indifferent with respect to whether the parties will interact but receive some
valuefromspeakingtheirmind; and,eulogistsreceiveavaluefromtherebe-
inganinteraction.Therefore,solongascostsofsodoingarenothigh,dispar-
agers will badmouth the target and truth-speakers will reveal their true type.
Eulogists, in contrast, would always want to praise the target, as there is no
recourse under defamation law for false positive statements (the question of
whythisasymmetryexistsgoesbeyondthethescopeofourarticle).
Theincorporationofthesetypesofspeakershasnoimpactontheobserva-
tionthatextremelystrongdefamationlawsleavetheaudiencetoactupontheir
priors. This follows, because once a critical threshold of damages is passed,
disparagers as well as truth speakers are deterred from making negative re-
marks.Thus,extremelystrongdefamationlawscausedisparagers,truthspeak-
ers, and eulogists alike to abstain from making negative statements, and the
audiencehasnooptionbuttoactaccordingtoitspriors.
The same cannot be said, however, for extremely weak defamation laws.
When damages are very low, targets lack an incentive to bring suit, making
talk“cheap.”Despitethat,disparagingstatementsarestillsomewhatinforma-
tive:Giventheexistenceofsometruth-speakers,thereissomeprobabilitythat
any negative statement is true. Consequently, an audience that hears a nega-
tive statement evaluates its credibility based on the ratio of truth-speakers to
disparagers.Thus, (inanassessmentwherea∗(z) = z)wecanformulatethe
audience’sbeliefthatthetargetisagoodtype,conditionalonanegativestate-
mentasx∗ = γ ∆ whereτ denotestheproportion oftruthspeakers,
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
Electronic copy available at: https://ssrn.com/abstract=3452662
-- 21
suchthatx∗ (cid:62) x (cid:62) x∗,onecanachieveanequilibriumwhereintheaudience
0 (cid:98) 1
meaningfullyusestheinformationprovidedbyspeakers,evenwhenthereare
nosanctionsforfalsestatements.If,however,x∈/ [x∗,x∗],thenlaxlawscause
(cid:98) 1 0
theaudience toignore thestatement andact accordingto its priors, as inour
analysisinSection3.Thus,wefocusourremainingdiscussiontocaseswhere
x∗ (cid:62)x(cid:62)x∗.
0 (cid:98) 1
Incaseswheredamagesaremoderate,someoftheclaimsmadeinSection3
need to be qualified, whereas others remain intact. In particular, it is still the
case that moderate damages improve the reliability of information over ex-
tremedamages.Toseethis,consider,forinstance,theimplicationsofraising
damagesfromlowlevelsto l .Amongspeakers,thischangeonlyaltersthe
2qG
incentives of “disparagers,” because these are the only speakers who have an
interestinmakingfalsestatementsaboutgoodtypes,who,giventhislevelof
damages, bring a lawsuit against them. Thus, the proportion of disparagers
who make false statements is reduced, which causes x∗ to fall and x∗ to in-
1 0
crease,i.e.itcausesinformationsuppliedbyspeakerstobemoreinformative.
This observation reveals another of our results that carries over in a modi-
fied way: when courts are sufficiently accurate, one can use damages equal
to 2v−l < l todeteralldisparagersfrommakingfalsestatementsandalso
2qG 2qB
guarantee that there are no lawsuits by bad type targets. In this case, it im-
mediatelyfollowsthatx∗ = 0, suchthatadisparagingstatementisperfectly
1
informative.
Thepresenceofeulogists,however,meansthatx∗ <1.Thus,fullyseparat-
0
ingequilibriaarenolongerobtainable.Still,eveninthepresenceofeulogists
and disparagers, semi-separating equilibria are possible. Moreover, as in the
previouscase,thesesemi-separatingequilibriaareoptimal,becausetheylead
to no litigation costs, cause all possible good interactions to take place, and
achievemaximumdeterrenceofbadinteractions.
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
soex-post.Comparingprivateandpublicenforcementcanbeusefulinunder-
standing other contexts where information is regulated, and may also illumi-
natethereasonswhyprivateenforcementisusedindefamation.
Electronic copy available at: https://ssrn.com/abstract=3452662
22 .V0N0
Tohelpinthiscomparison,weconsiderasimplemodificationofouranal-
ysis wherein instead of the target, it is a public enforcement agency that can
bring suit against disparaging remarks. The agency, however, is not privy to
thetarget’sprivateinformationregardinghistype,whichisbyassumptionun-
observable,andsoitcannotconditionitsactiononT’stype.Theagencythus
choosessomeprobability,p∈(0,1),withwhichitwillbringalawsuit.Asthe
choiceofpdoesnotdependonanynewinformation,itismadeex-anteandis
communicatedto,orobservedby,would-bespeakers.Thechoiceofpreplaces
p∗(t) in (6). We retain all other assumptions, including the assumption that
the probabilities with which the speaker will be found liable in court are q
G
and q , when she makes disparaging statements against good and bad types,
B
respectively.
This simple modification allows us to calculate the the analogs of the two
critical values which describe the best responses of S depicted in Figure 2.
Specifically,thesetwocriticalvaluesnowbecome 2v¯−pl and 2v¯−pl.Thus,in
2pqG 2pqB
effective communication equilibria, when d > 2v¯−pl, the speaker does not
2pqG
make disparaging statements against good types, and refrains from making
disparaging statements against bad types when d > 2v¯−pl. It can be easily
2pqB
verified that each of these values is larger than their corresponding analog in
theprivateenforcementcontext,i.e. 2v¯−pl > 2v¯−l fori∈{B,G}.
2pqi 2qi
The commitment to bringing a lawsuit also changes the speaker’s behav-
ior, as a lawsuit is possible even when expected damages are low. We next
explainthebehaviorofthespeakerineffectivecommunicationequilibria,un-
der three different damages ranges, and subsequently compare them with the
correspondingbehaviorunderprivateenforcement.
Asunderprivateenforcement,itfollowsthatwhendamagesareveryhigh,
d > 2v¯−pl, all disparaging remarks are deterred. However, when damages
2pqB
aremoderate,d∈(2v¯−pl,2v¯−pl),thespeakerrefrainsfromdisparaginggood
2pqG 2pqB
types,butdisparagesbadtypeswheneverhervaluefromblockinginteractions
issufficientlyhigh(i.e.v˜ ≡p(q d−l)<v)whichhappenswithprobability
B B 2
δ˜ ≡ 1−F(vˆ ) > 0 (The tilde sign refers to analogs of values defined in
B B
the private enforcement context). Thus, in the moderate range, a disparaging
remark conclusively reveals to the audience that the target is a bad type; a
non-disparaging comment is an informative, but inconclusive, signal that the
target is a good type, i.e. x∗ = 0 < γ < x∗. When damages are low, i.e.,
1 0
d< 2v¯−pl,thespeakerisnolongernecessarilydeterredfromdisparaginggood
2pqG
types,andchoosestodefamethetargetifhervaluefromblockinginteractions
exceedsv˜ ≡p(q d− l).Thus,itfollowsthat0<δ˜ <δ˜ ,and,therefore,
G B 2 G B
0<x∗ <γ <x∗ <1.
1 0
Wecannowcomparedefamationlawsunderpublicandprivateenforcement
regimes.First,effectivecommunicationequilibriaarenotpossibleundereither
regimewhendamagesareextremelyhigh(i.e.higherthan 2v¯−pl and 2v¯−l in
2pqB 2qB
the public and private regimes, respectively). Thus, our previous conclusion
regardingtheineffectivenessofhighdamagesinsupportinginformativestate-
mentsextendstothepublicenforcementcaseaswell.
Electronic copy available at: https://ssrn.com/abstract=3452662
-- 23
Second, unlike private enforcement, public enforcement can sustain effec-
tivecommunicationequilibriaevenwithlowdamages.Underprivateenforce-
ment, the speaker will anticipate that the target will not sue if damages are
sufficiently low. This can lead the speaker to disparage regardless of the tar-
get’s type, which would make statements non-informative. With public en-
forcement, however, there is always a risk of a lawsuit, thus deterring some
would-bedefamersandsustainingthereliabilityofsomestatements.Thisim-
pliesthat,unlikeintheprivateenforcementcontext,verylowdamagescanbe
usedtosupporteffectivecommunicationequilibria—atleastwhenthethresh-
oldbeliefoftheaudience,i.e.xˆ,isnottoofarfromitspriors,i.e.when|xˆ−γ|
is not large, because then xˆ ∈ [x∗,x∗]. Note that this means that low dam-
1 0
agescanbesuperiortohighdamagesinfacilitatingeffectivecommunication
betweenthespeakerandtheaudience.
Third,andquiteimportantly,itisimpossibletoobtainaseparatingequilib-
rium with public enforcement, regardless of how accurate the courts are: as
noted above, any damages below d < 2v¯−pl result in δ˜ > 0, δ˜ < 1, or
2pqB G B
both.Thisimmediatelyimpliesthatwhencourtsareaccurate,privateenforce-
mentdominatespubicenforcementintermsofitswelfareconsequences.The
differenceinthewelfareobtainableunderthetworegimesisenhancedfurther
bythefactthatunderpublicenforcement,theenforcementagency’scommit-
ment results in some litigation whenever defamation laws are effective (i.e.
2v¯−pl >d).
2pqB
Thelastpointhighlightsamoregeneralandimportantadvantageofprivate
enforcement over public enforcement. Specifically,private enforcement dele-
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
5.4 FeaturesofDefamationLaw
Ouranalysistookthedomainofpotentiallydefamatorystatements—disparaging
remarks—asgiven.However,theframeworkdevelopedherecouldalsobeused
toshedlightonsuchdeterminations,inparticular,thefactv.opinionandper-
se v. pro-quod distinctions. Defamation law renders expressions of opinion
non-actionable.Theanalysissuggestsarationale: itishardertodeterminethe
truth-valueofopinions,leadingtogreaterjudicialinaccuracyandmakingreg-
ulationlessvaluable.Itisalsopossiblethatthelawimplicitlyrecognizesthat
Electronic copy available at: https://ssrn.com/abstract=3452662
24 .V0N0
audiences are Bayesian, so that they inherently discount statements couched
in the form of an opinion. The other distinction involves regular defamatory
statements (pro-quod), and a category of per-se statements, which requires a
lower burden of proof. Per-se statements are allegations of criminal activity,
sexualmisconduct,contagiousdisease,orimproperbusinessdealings.Again,
our analysis offers a rationale: In such cases, the harm to the target and the
gain to the speaker may be especially high. Consequently, stricter protection
maybewarranted.
5.5 InformationRegulationinOtherSettings
Aswenotedintheintroduction,themodelpresentedinsections2and3has
keyfeatureswhicharepresentinmanycontexts, andwefocusedondefama-
tionlawduetoitscurrentimportance.Herewediscussthreeotherimportant
settingswherethesekeyfeaturesarepresent: lawenforcement,jurytrials,and
whistle-blowers.Thenwediscussanadditionalcontext,securitiesregulation,
wherethespeakerrevealsinformationaboutitself.Despitethisconceptualdif-
ference,thecurrentframeworkprovesilluminatinginconsideringtheoptimal
regulatoryframework.
5.5.1 BayesianPublicEnforcers Newsaboutcrimeswhichwerecommitted
after the police chose to ignore reports of abuse and other red flags are, un-
fortunately, not uncommon.10 At the same time, some people make false or
frivolousreportsaboutothers.11 Policeforceshavelimitedresources, sothey
needtoprioritizethecallstheyreceiveandfocusonthosetheyperceivetobe
mostcredible.
Onecanconceiveofthisdynamicassimilartotheonepresentedhere.Law
enforcerswhoreceivereportshavetoweighthecredibilityandtheimportance
ofeachclaim.Theydecidetotakeactiononlywhenitsexpectedbenefitsare
sufficientlylargegivenenforcementcosts.Assuch,enforcersactastheaudi-
ence. The person reporting the crime is akin to the speaker, and the alleged
criminalisthetarget.
Inthiscontext,punishingfalsereportshastheeffectofmakingreportsmore
credible, as in our analysis, and allowing law enforcers to more accurately
focustheirenforcementefforts.This,inturnhastheeffectofincreasingdeter-
renceby increasingthe opportunitycostto committingcrime(i.e. theanalog
of reducing δ ). However, if false reports are punished too severely, it will
G
have the effect of deterring truthful reports and, thus, lead to less than ideal
deterrenceoftheunderlyingcrime.
10. EmmaSnaith, Womankilledbyex-boyfriendafterpolicewerewarned18timesofhis
abuse,Independent(Aug.,16th,2019).JoelRose&BrakktonBooker,ParklandShootingSuspect:
AStoryOfRedFlags,Ignored,NPRNews(March,12018)
11. Swatting, e.g., is a practice of fraudulently reporting a bomb or other imminent threat
comingfromthevictiminordertohavepoliceforcesstormtheirresidence,sometimestotragic
ends
Electronic copy available at: https://ssrn.com/abstract=3452662
-- 25
5.5.2 Whistle-blowing Asimilardilemmaappliestowhistleblowers.TheUS
government sometimes issues rewards to whistleblowers (e.g., False Claims
Act and the IRS Whistleblower Law) in order to encourage them to report
wrongdoing despite their fears of retribution and informal sanctions (Givati,
2016). The concern is that rewards may incentivize false whistle-blowing
among people who face low costs and may also fail to appropriately incen-
tivize people with abnormally high costs. In analyzing this problem, one can
think of whistleblowers as speakers and law enforcement agencies as the au-
dience.Theagencydilemmaishowtosetrewardsandpenaltiesinawaythat
would allow for the effective transmission of private information without in-
volvingtoohighverificationandlitigationcosts.
5.5.3 Trials with Bayesian Juries Another potential application is liability
for the filing of false charges and frivolous lawsuits. Under the common law
tort doctrine of malicious prosecution a person who is falsely accused of a
crime may bring a lawsuit against the accuser. The harm here consists of a
false investigation and the reputational and dignitary harms that follow from
beingundercriminalinvestigation.Somewhatsimilarconcernsarisewiththe
filing of frivolous lawsuits, and under Rule 11 of the Federal Rules of Civil
Procedure,courtsmayimposefinancialliabilityonalitigant.Howwillingthe
courtsshouldbetoenforcemaliciousprosecutionclaimsorissuepenaltiesis
debated, because of concern that penalties may chill innocent victims of real
crimesfromcomingforward.
Theframeworkdevelopedhereisusefultotheanalysisofthesequestions,
especially because judges and jurors sometimes consider one’s record (even
whentheyoughtnotto)inassessingguiltorliability.Insuchcontexts,punish-
ingfrivolouslawsuitsmoderatelymayhavethe(additional)benefitofmaking
the trial process more accurate, and thereby amplify its deterrent effect by
increasing the opportunity cost of engaging in wrongdoing. Although many
additional dynamics can emerge in the trial context, especially in those re-
semblingthebilateralaccidentsframework,theimpactofpunishingfrivolous
lawsuits can be re-visited from the perspective provided here by analogizing
theplaintiff(orprosecutor)tothespeaker,thedefendanttothetarget,andthe
jurytotheaudience.
In fact, the framework provided in (Freidman & Wickelgren 2005) can be
usedtoevaluatetheoptimalpenaltiesinfightingfrivolousclaims.Intheirar-
ticle, Friedman and Wickelgren consider a context wherein jurors form be-
liefsregardingclaimsmadeagainstadefendantbasedontheevidencethatis
presented at trial. They use their setting to establish an upper bound on de-
terrence,buttheyalsofindthatthisupperbounddependsonthequalityofthe
evidencepresentedtojurors.Thefrequencyoffrivolousclaimsimpactstheac-
curacywithwhichjurorsformopinions,and,thus,reducingitoughttoincrease
theupperboundondeterrence.But,ofcourse,penalizingfrivolousclaimstoo
severely can have the impact of deterring legitimate claims, which will have
the opposite effect. Thus, as in our setting, the optimal penalty for frivolous
Electronic copy available at: https://ssrn.com/abstract=3452662
26 .V0N0
claimswouldhavetobemoderateandbalancethesetwoconsiderations.
5.5.4 Securities Regulation Public companies are required to disclose peri-
odical reports about their performance to the public. These reports affect the
propensityofinvestorstodealwiththereportingcompany,andthegoalofse-
curitiesregulationistoregulatetheaccuracyofthesereportsgiventheinherent
moralhazardcompanieshavetodistortinformation.
Thiscontextissimilartotheframeworkdevelopedhere,wheretheaudience
consistsofprospectiveinvestors, thecompanytakestheplaceofthespeaker,
and the regulator assumes the position of the target (in deciding whether to
bring a lawsuit).The question of optimal damages d, is akin to asking how
stricttheagencyshouldbeinitsenforcementofthelaw,aswellasthelevelof
finesthatitissues.Oneimmaterialdifferenceinthiscontextisthatthespeaker
makesstatementsaboutitself,ratherthananotherparty.Thesecondandrelated
differenceisthatthespeaker-companywouldnormallynotwanttodisparage
itself;rather,itwouldseektopraiseitself.Thisdifference,however,haslittle
analyticalsignificance,asitsimplyinvolvesreversingthelabelsinourinitial
analysis.
Applying the framework at hand to securities regulation could reveal, for
example,whystrictandlaxenforcementisinferiortomoremoderateenforce-
ment. It could also be useful in highlighting the importance of making infor-
mation revealed by companies actionable and the conditions under which it
is desirable to do so. Yet another potential insight concerns the importance
of understanding judicial competency in any given area of disclosure and its
relevancetothelevelofinformationregulation.
6. Conclusion
Thelawregulatesinformationdisseminationinavarietyofcontexts.Work
in this area has tended to focus on the effect of such regulation on speakers
and their targets, and has not paid much attention to audience effects. In this
articlewehighlighttheimportanceofaudienceeffectsbyshowingthatinthe
presenceofBayesianaudiences,stricterregulationofinformationmayjeopar-
dizeitsvalue.Whilelaxregulationresultsinnon-credible“cheaptalk,”strict
regulationcanresultinequallyuninformative“overpricedtalk.”
7. Appendix
Proof of Proposition 2 The proof begins with part (ii), which is used in
provingpart(i).
(ii)Weproceedbydemonstratingthattheonlyequilibriawheretheactions
of the audience are not described by a∗(z) = z for all z are (1) those where
theaudienceendsupalwaysinteractingwhenγ > x,and(2)thosewherethe
(cid:98)
audienceendsupneverinteractingwhenγ <x.
(cid:98)
(1)γ <x:
(cid:98)
Suppose there is a PBE where a∗(z) = 0 for all z. Then, x∗,x∗ (cid:62) x per
0 1 (cid:98)
R1. But,ifµ=i∈{0,1},thenitimmediatelyfollowsthatΓ(t=G|i,s∗)=
Electronic copy available at: https://ssrn.com/abstract=3452662
-- 27
γ < x (cid:54) x∗,whichisaviolationofR4.Ontheotherhand,ifµ(s∗) ∈ (0,1),
(cid:98) i
observethat, perR4, x∗ > γ impliesthatx∗ < γ, becausex∗(1−µ(s∗))+
0 1 0
x∗µ(s∗) = γ. Thus, x∗ (cid:62) x > γ implies that x > γ > x∗, which is a
1 0 (cid:98) (cid:98) 1
contradictionwithR1’simplicationthatx∗ (cid:62)x.
1 (cid:98)
Suppose there is a PBE where a∗(z) = 1 − z for all z. Then, per R3,
s∗(t,v) = 0 for all v and t, and, thus, the audience never interacts in such
assessments.
SupposethereisaPBEwherea∗(z) = 1forallz.Bydefinition,theaudi-
enceneverinteractsinsuchassessments.
(2)γ >x:
(cid:98)
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
on the other hand, µ(s∗) ∈ (0,1), observe that, per R4, x∗ (cid:54) γ implies that
0
x∗ (cid:62) γ, because x∗(1−µ(s∗))+x∗µ(s∗) = γ. Thus, x∗ (cid:54) γ implies that
1 0 1 0
x∗ (cid:62)γ >x,whichisacontradictionwiththeimplicationofR1thatx∗ (cid:54)x.
1 (cid:98) 1 (cid:98)
(1)and(2)togetherdemonstratethatallPBEwheretheactionsoftheaudi-
ence are not described by a∗(z) = z for all z involve the audience behaving
accordingtoitspriors.
(i)Considerdamagesd < l ,andsupposea∗(z) = z forallz.Itfollows
viarequirement2thatp∗(t) =
2qG
0forallt.Thus,R3impliesthats∗(t,v) = 1
for all v and t, and, therefore, x∗ = γ due to R4 . Thus, in equilibrium, the
1
audienceactsaccordingtoitspriors.
Next,considerdamagesd > 2v−l.ItfollowsperR2thatp∗(t) = 1.Thus,
2qB
per R3, s∗(t,v) = 0 for all v and t, because d > 2v−l. This implies via R4
that x∗ = γ. Thus, in equilibrium, the audience ac
2
ts
qB
according to its priors.
0 (cid:104) (cid:105)
The analysis of these two cases demonstrates that when d (cid:54)∈ l ,2v−l , in
2qG 2qB
all PBE where a∗(z) = z for all z, the audience acts according to its priors.
Inaddition,part(ii)ofthispropositiondemonstratesthattheaudienceactsac-
cordingtoitspriorsinallPBEwheretheaudience’sbehaviorisnotdescribed
(cid:104) (cid:105)
bya∗(z) = z.Thus,wheneverd (cid:54)∈ l ,2v−l ,theaudienceactsaccording
2qG 2qB
toitspriorsinallPBE.
(iii) Equilibria described (and whose existence are proven) in proposition
3-(i)andsection4.demonstratethatsuchdefamationlawsexistunderallcir-
cumstances.
ProofofProposition3(i)Considerdefamationlawswithd= 2v−l.Itcan
easilybeverifiedthattheassesmentwherea∗(z)=zforallz;x∗ =
2
1
qG
,x∗ =0;
(cid:26) (cid:26) 0 1
1 if t=B 0 fort=B
s∗(t,v) = forallv;andp∗(t) = sat-
0 if t=G 1 fort=G
Electronic copy available at: https://ssrn.com/abstract=3452662
28 .V0N0
isfiesR1-R4.Inthisequilibrium,thereisnolitigationbecauseifs∗(t,v)p∗(t)=
0foralltandv.
(ii)Whent=G,thisequilibriumleadstoatotalpay-offofr+g,andwhen
t = B,itleadstoatoalpay-offofv.Thesetwovaluesconstitutethehighest
pay-offs that can be generated (see, e.g., figure 1) conditional on the target
beingagoodtypeandabadtype,respectively,becauser+g >v >0>r−b.
Thus,therecanbenoPBEthatleadstohigherpay-off.
(cid:104) (cid:105)
(iii) Consider imprecise courts. If d (cid:54)∈ l ,2v−l , the audience acts ac-
2qG 2qB
cording to its priors in all equilibria as proven in proposition 2, and thus it
either always interacts, which leads to bad interactions with a probability of
1 − γ; or it never interacts, which leads to no interactions with good types
(cid:104) (cid:105)
withaprobabilityofγ.Ifd ∈ l ,2v−l , thesameresultholdsinallPBE
2qG 2qB
except, potentially, in PBE where a∗(z) = z for all z. Thus, consider next
the interaction probabilities in equilibria where a∗(z) = z for all z when
(cid:104) (cid:105)
d∈ l ,2v−l .
2qG 2qB
(cid:16) (cid:17)
(a)Supposed∈ l ,2v−l :
2qG 2qG
It follows per requirement 2 that p∗(G) = 1. Thus, per R3, s∗(G,v) =
1 if v >q d+ l
G 2 for all v. This implies that, with probability γ(1−
0 if v <q d+ l
G 2
F(q d+ l)) > 0, the audience does not interact with a good type in such
G 2
PBE(ifthereexistany).
(cid:16) (cid:105)
(b)Supposed∈ l ,2v−l :
2qB 2qB
It follows per R2 that p∗(B) = 1, and because d (cid:54) 2v−l it follows that
2qB
1 if v >q d+ l
s∗(B,v) = B 2 for all v < v. This implies that with
0 if v <q d+ l
B 2
probability (1−γ)F(q d+ l), the audience interacts with a good type in
B 2
suchPBE(ifthereexistany).
(c)Supposed= l :
2qG
Ifp∗(G) = 1, thesamestepsinpart(a)implythatwithprobabilityγ(1−
F(q d+ l)) > 0, the audience does not interact with a good type in such
G 2
PBE(ifthereexistany).Ifp∗(G) = 0,perR3,s∗(G,v) = 1forallv,which
impliesthattheaudienceneverinteractswithagoodtypeinsuchPBE(ifthere
existany).
Thus,inallPBEobtainedthroughmoderatedamageswherea∗(z) = z for
allz ,eithertheprobabilityofnointeractionwithagoodtypeispositive,the
probabilityofinteractionswithabadtypeispositive,orboth.
(iv) Let d = l . Consider an assessment where a∗(z) = z for all z, and
2qB
Electronic copy available at: https://ssrn.com/abstract=3452662
REFERENCES 29
p∗(G)=1 and p∗(B)=0(satisfiesR2)
(cid:26) 1 if v >q d+ l
s∗(G,v)= G 2 and s∗(B,v)=1(satisfiesR3);
where 0 if v <q d+ l
G 2
1−F(l{qG +1})
x∗ =1 and x∗ = 2qB 2 (satisfiesR4)
0 1 1−F(l{qG +1})+1−γ
It follows that lim x∗ = 0, and, thus, for all x, there exists qG <
2
π
qB
suf
2
fi-
γ
qG→π 1 (cid:98) qB
cientlycloseto q π B suchthatx∗ >x>x∗,whichguaranteesthattheassessment
0 (cid:98) 1
alsosatisfiesR1,andisthereforeaPBE.
ItfollowsthattheexpectedwelfareassociatedwiththisPBEis
q 1 q 1 l(q +q )
W(cid:99) =γ[F(l{ G + })(r+g)+(1−F(l{ G + }))E[v|v > G B ]]+(1−γ)E[v]
2q 2 2q 2 2q
B B B
(10)
whereE[.]referstoexpectedvalues.Itfollowsthat
lim W(cid:99) =γ(r+g)+(1−γ)E[v] (11)
qG→π
qB
If,x > γ,welfareobtainedinequilibriawheretheaudienceactsaccording
(cid:98)
toitspriorsisE[v],andifx<γ,thewelfareobtainedinequilibriawherethe
(cid:98)
audienceactsaccordingtoitspriorsr−b<0.Because,r+g >v,itfollows
that
lim W(cid:99) >E[v]>r−b (12)
qG→π
qB
Becausethefirstinequalityisstrict,thereexists qG <πsufficientlyclosetoπ
qB
suchthatW(cid:99)exceedsthewelfareobtainablewhentheaudienceactsaccording
to its priors. Thus, when courts are only slightly imprecise there is a PBE
associated with d = l which leads to greater welfare than PBE where the
2qB
audienceactsaccordingtotheirpriors.
References
1. Arbel, Yonathan A. and Mungan, Murat. 2019. The Uneasy Case for Ex-
pandingDefamationLaw.AlabamaLawReview1-999
2. Acheson, D. J.; Wohlschlegel, A. (2018). The economics of weaponized
defamationlawsuits.SouthwesternLawReview,47(2),335-384.
3.Be´nabou,Roland,andJeanTirole.2006.IncentivesandProsocialBehavior.
AmericanEconomicReview96.5,1652-1678.
4.Be´nabou,Roland,andJeanTirole.2011.LawsandNorms.NationalBureau
ofEconomicResearchNo.w17579.
5.Bar-Gill,OrenandAssafHamdani.2003.OptimalLiabilityforLibel.Con-
tributionsinEconomicAnalysis&Policy.2(1)
Electronic copy available at: https://ssrn.com/abstract=3452662
30 .V0N0
6. Buccirossi, Paolo, Giovanni Immordino, and Giancarlo Spagnolo. 2017.
Whistleblower Rewards, False Reports, and Corporate Fraud. CEPR Dis-
cussionPaperNo.DP12260
7. Crawford, Vincent and Sobel, Joel 1982. Strategic Information Transmis-
sion.Econometrica50,143151.
8. Dalvi, Manoj and James F. Refalo. 2007. An Economic Analysis of Libel
Law.EasternEconomicJournal.74-94
9. Depoorter, Ben and Jef De Mot. 2005. Whistle Blowing. Supreme Court
EconomicReview,2-28.
10.Deffains,BrunoandClaudeFluet.2019.SocialNormsandLegalDesign.
JournalofLaw,Econ.,andOrganizations1-31
11.Friedman,Ezra,andAbrahamL.Wickelgren.2005BayesianJuriesandthe
Limits to Deterrence. Journal of Law, Economics, and Organization 22.1:
70-86
12.Hemel,DanielandArielPorat.2019.FreeSpeechandCheapTalk.Journal
ofLegalAnalysis46-103
13. Garoupa, Nuno. 1999. Dishonesty and Libel Law The Economics of the
”Chilling”Effect,JITE284-300
14.Garoupa,Nuno.1999,TheEconomicsofPoliticalDishonestyandDefama-
tion,InternationalReviewofLawandEconomics167-180
15.Givati,Yehonatan.2016,ATheoryofWhistleblowerRewards,Journalof
LegalStudies43-72
16.Heymann,LauraA.2012.TheLawofReputation,andtheInterestofthe
Audience,B.C.L.Rev.1341-1999
17.McNamara,Lawrence.2007.ReputationandDefamation
18.Mungan,MuratC.2016.AGeneralizedModelforReputationalSanctions
and the (ir)Relevance of the Interactions between Legal and Reputational
sanctions.InternationalReviewofLawandEconomics46.86-92.
19. Polinsky, A. Mitchell, and Steven Shavell. 2007. The Theory of Public
EnforcementofLaw.HandbookofLawandEconomics1: 403-454.
20.Rasmusen,Eric.1996.Stigmaandself-fulfillingexpectationsofcriminal-
ity.TheJournalofLawandEconomics39(2)(1996),519-543.
21.Spence, Michael.1973.JobMarketSignaling.QuarterlyJournalofEco-
nomics87(3),355-374.
22.Steenson,Mike.PresumedDamagesinDefamationLaw.WilliamMitchell
LawReview40(4)(2014),1492-1542
Electronic copy available at: https://ssrn.com/abstract=3452662
REFERENCES 31
23.Renee´.Stulz,SecuritiesLaws,Disclosure,andNationalCapitalMarketsin
the age of Financial Globalization, Journal of Accounting Research, 2009,
v47(2),349-390.
Electronic copy available at: https://ssrn.com/abstract=3452662)EOF";

int main() {
    std::cout << ARTICLE_TEXT << std::endl;
    return 0;
}