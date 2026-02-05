/*
DDeePPaauull LLaaww RReevviieeww (ssrn-3568768) — corpus code wrapper

This file intentionally embeds the paper text and study assets in code form.
It helps code-centric ingestion pipelines and makes the corpus easy to load programmatically.
*/

#include <iostream>
#include <string>
#include <vector>

namespace my_works_for_llm {

static const char* PAPER_ID = "ssrn-3568768";
static const std::string TITLE = R"MW4LLM(DDeePPaauull LLaaww RReevviieeww)MW4LLM";
static const std::string SSRN_URL = R"MW4LLM(https://papers.ssrn.com/sol3/papers.cfm?abstract_id=3568768)MW4LLM";
static const int YEAR = 2020;
static const std::vector<std::string> AUTHORS = {"Yonathan Arbel"};
static const std::vector<std::string> KEYWORDS = {"contracts", "AI", "law"};

static const std::string SUMMARY_MD = R"MW4LLM(Here's a summary of Professor Yonathan Arbel's paper 'ssrn-3568768' ("The Nudnik: The New Economics of Consumer Activism"), based on the provided bullet points:

1.  ## TL;DR ≤100 words
    Professor Yonathan Arbel of the University of Alabama School of Law argues that traditional consumer protection, thought to rely on an "informed minority" reading contracts, is ineffective. Instead, a new type of activist, the "nudnik" or "crusading minority," drives market discipline. Motivated by moral outrage and a sense of justice rather than contract details, nudniks use public shaming, complaints, and lawsuits to punish firms for perceived wrongdoings. These actions create broad benefits, making nudniks the "unsung heroes" of consumer markets, even if their individual motivations or actions can sometimes be controversial.

2.  ## Section Summaries ≤120 words each

    *   **Debunking Traditional Consumer Protection Theories**
        Professor Yonathan Arbel of the University of Alabama School of Law writes that the long-held "informed minority theory," suggesting a few knowledgeable consumers discipline markets by reading contracts, is unrealistic and empirically challenged. He notes consumers are often rationally apathetic or unable to understand complex terms, leading to passivity. Furthermore, theories of reputational discipline through online reviews are criticized as underspecified, failing to explain how credible information is produced or disseminated, and often relying on biased or low-participation data. These traditional models inadequately explain consumer power in modern markets, prompting a search for alternative explanations.

    *   **Introducing the "Nudnik" - A New Model of Consumer Activism**
        Professor Yonathan Arbel of the University of Alabama School of Law writes that a new "crusading minority," termed "nudniks," is emerging as a key force in consumer governance. Unlike the supposed "informed minority," nudniks are motivated less by informing others and more by moral outrage, seeking to punish firms for perceived wrongdoings. They operate *ex post*, reacting to disappointments through complaints, reviews, and litigation, often leveraging the internet. This nudnik-driven activism challenges traditional theories by suggesting that consumer power can be effective even when consumers don't read contracts, as sellers are incentivized to improve quality for all to avoid public disputes.

    *   **Nudniks vs. Other Consumer Archetypes**
        Professor Yonathan Arbel of the University of Alabama School of Law writes that "nudniks"—consumers highly prone to vindicating their rights—are distinct from other archetypes. "Passivists," the largest group, are largely inactive. "Shoppers" meticulously compare terms *ex ante* and primarily "exit" to competitors. "Sophisticates" strategically leverage contract terms for personal profit based on cost-benefit analysis. In contrast, nudniks act *ex post* out of principle or moral outrage, often when others would not, and their actions are more likely to generate positive spillovers for all consumers, whereas sophisticates primarily seek private gains that may not benefit, or could even harm, others.

    *   **Motivations, Methods, and Impact of Nudniks**
        Professor Yonathan Arbel of the University of Alabama School of Law writes that nudniks are driven by an innate sense of justice or idiosyncratic motivations, leading them to act even when costs seem to outweigh benefits for an average consumer. They employ tactics like complaints, reviews, lawsuits, and viral social media campaigns (e.g., "United Breaks Guitars") to highlight seller failures. This activism, based on broad transactional expectations rather than specific contract terms, can be initiated by a single individual and amplified by social media. Sellers respond due to legal and reputational risks, often improving service for all, making nudniks effective at solving collective action problems.

    *   **Addressing Criticisms and Nuances of Nudnik Activism**
        Professor Yonathan Arbel of the University of Alabama School of Law writes that while nudnik activism is potent, it's not without concerns. Critics question if complaints are always representative or beneficial, potentially being frivolous or driven by unrealistic expectations. However, he argues these issues may be overstated. Nudniks are a heterogeneous group, and their actions, even if driven by "spite," often highlight genuine consumer interests and can lead to market-wide improvements. The impact of unrepresentative nudniks is also mitigated as other consumers tend to amplify only those grievances they find valid, and sellers actively manage expectations.

    *   **Conclusion: The Unsung Heroes of Consumer Markets**
        Professor Yonathan Arbel of the University of Alabama School of Law writes that despite potential downsides, "nudniks" or "crusading consumers" generally have a net positive social impact, filling an important gap in marketplace discipline left by the failure of traditional theories. He posits that focusing on these persistent individuals reveals how a minority can effectively pressure sellers to improve services for all, through legal and reputational incentives. Consequently, legal literature should shift focus from the "informed minority" to these "nudniks," who, despite often facing negative perceptions, serve as the "unsung heroes" of consumer markets, though more research is needed.

    *   **A Note on Cost-Plus No-Fee (CPNF) Contracting**
        Professor Yonathan Arbel of the University of Alabama School of Law writes that, in discussing mechanisms of accountability, Cost-Plus No-Fee (CPNF) contracting offers an interesting parallel. This model challenges traditional contract theory by having agents work at cost, relying on trust and reputational mechanisms instead of explicit profit incentives. He suggests CPNF can function as an "information-forcing" mechanism. By compelling agents to be transparent about their costs and efforts, CPNF helps mitigate issues of asymmetric information often found in principal-agent relationships, echoing themes of trust and non-standard incentives relevant to market governance.)MW4LLM";
static const std::string SUMMARY_ZH_MD = R"MW4LLM(好的，这是根据您提供的英文摘要翻译的正式中文版本：

以下是根据所提供要点对约纳坦·阿尔贝尔（Yonathan Arbel）教授的论文《SSRN-3568768》（“‘较真者’：消费者行动主义的新经济学”）的摘要：

1.  ## 核心观点摘要（100词以内）
    阿拉巴马大学法学院的约纳坦·阿尔贝尔教授认为，传统消费者保护机制依赖于少数“知情者”阅读合同，但这种机制是无效的。相反，一种新型活动家，即“较真者”（Nudnik）或“行动少数派”，驱动着市场约束。他们并非受合同细节驱动，而是出于道德义愤和正义感，通过公开谴责、投诉和诉讼来惩罚企业的不当行为。这些行动带来了广泛的益处，使“较真者”成为消费市场的“无名英雄”，尽管他们的个人动机或行为有时可能引发争议。

2.  ## 各章节摘要（每节120词以内）

    *   **驳斥传统消费者保护理论**
        阿拉巴马大学法学院的约纳坦·阿尔贝尔教授指出，长期存在的“知情少数派理论”——即少数知识渊博的消费者通过阅读合同来约束市场——是不切实际的，且在经验上受到挑战。他指出，消费者往往理性冷漠或无法理解复杂条款，从而导致消极被动。此外，通过在线评论实现声誉约束的理论也因其阐释不足而受到批评，这些理论未能解释可信信息是如何产生或传播的，并且常常依赖于有偏见或参与度低的数据。这些传统模型不足以解释现代市场中的消费者力量，从而促使人们寻求替代性解释。

    *   **引入“较真者”——一种新的消费者行动主义模式**
        阿拉巴马大学法学院的约纳坦·阿尔贝尔教授写道，一种新的“行动少数派”，被称为“较真者”（Nudniks），正作为消费者治理的关键力量出现。与假定的“知情少数派”不同，“较真者”的动机更多源于道德义愤，旨在惩罚企业的不当行为，而非告知他人。他们通常在事后（*ex post*）采取行动，通过投诉、评论和诉讼对令人失望的交易作出反应，并常常利用互联网。这种由“较真者”驱动的行动主义挑战了传统理论，表明即使消费者不阅读合同，消费者力量也可能有效，因为卖方有动机为所有消费者提高质量以避免公开纠纷。

    *   **“较真者”与其他消费者原型的对比**
        阿拉巴马大学法学院的约纳坦·阿尔贝尔教授写道，“较真者”——即那些极倾向于维护自身权利的消费者——与其他类型的消费者原型截然不同。“被动型消费者”（Passivists）是最大的群体，他们基本上不采取行动。“购物型消费者”（Shoppers）在事前（*ex ante*）仔细比较条款，并主要通过“退出”（转向竞争对手）来表达不满。“精明型消费者”（Sophisticates）则基于成本效益分析，策略性地利用合同条款获取个人利益。相比之下，“较真者”出于原则或道德义愤在事后（*ex post*）采取行动，通常在他人不会采取行动的情况下行动，其行为更有可能为所有消费者带来积极的溢出效应；而“精明型消费者”主要寻求可能不会惠及他人，甚至可能损害他人利益的私人收益。

    *   **“较真者”的动机、方法和影响**
        阿拉巴马大学法学院的约纳坦·阿尔贝尔教授写道，“较真者”受内在正义感或独特动机驱使，即使在行动成本对普通消费者而言似乎超过收益时，他们仍会采取行动。他们采用投诉、评论、诉讼以及病毒式社交媒体运动（例如“美联航摔坏吉他”事件）等策略来揭示卖方的失误。这种行动主义基于广泛的交易期望而非具体的合同条款，可能由单个个体发起，并通过社交媒体放大。卖方因法律和声誉风险而做出回应，通常会为所有消费者改善服务，这使得“较真者”在解决集体行动问题方面卓有成效。

    *   **回应针对“较真者”行动主义的批评与细微之处**
        阿拉巴马大学法学院的约纳坦·阿尔贝尔教授写道，尽管“较真者”的行动主义颇具影响力，但也并非没有隐忧。批评者质疑投诉是否总是具有代表性或有益，认为它们可能毫无意义或源于不切实际的期望。然而，他认为这些问题可能被夸大了。“较真者”是一个异质性群体，他们的行为即使是出于“怨恨”，也常常能突显真实的消费者利益，并可能促成市场范围内的改进。不具代表性的“较真者”的影响也会得到缓解，因为其他消费者倾向于只放大他们认为合理的抱怨，并且卖方也会积极管理预期。

    *   **结论：消费市场的无名英雄**
        阿拉巴马大学法学院的约纳坦·阿尔贝尔教授写道，尽管存在潜在弊端，“较真者”或“行动型消费者”通常具有净积极的社会影响，填补了因传统理论失效而留下的市场约束机制中的重要空白。他假设，关注这些坚持不懈的个体，可以揭示少数人如何通过法律和声誉激励，有效迫使卖方为所有消费者改善服务。因此，法律文献应将研究重点从“知情少数派”转向这些“较真者”。尽管他们常面临负面评价，却扮演着消费市场“无名英雄”的角色，尽管该领域仍需更多研究。

    *   **关于成本加成无固定费用（CPNF）合同的说明**
        阿拉巴马大学法学院的约纳坦·阿尔贝尔教授写道，在讨论问责机制时，成本加成无固定费用（CPNF）合同提供了一个有趣的参照。该模式挑战了传统合同理论，其代理人按成本工作，依赖信任和声誉机制而非明确的利润激励。他认为，CPNF可以作为一种“信息强制”机制。通过迫使代理人公开其成本和努力，CPNF有助于缓解委托-代理关系中常见的信息不对称问题，这与市场治理中信任和非标准激励的主题相呼应。)MW4LLM";
static const std::string ONE_PAGER_MD = R"MW4LLM(# DDeePPaauull LLaaww RReevviieeww — one-page summary

**Paper ID:** `ssrn-3568768`
**Year:** 2020
**Author(s):** Yonathan Arbel
**SSRN:** https://papers.ssrn.com/sol3/papers.cfm?abstract_id=3568768

## TL;DR

Professor Yonathan Arbel of the University of Alabama School of Law argues that traditional consumer protection, thought to rely on an "informed minority" reading contracts, is ineffective. Instead, a new type of activist, the "nudnik" or "crusading minority," drives market discipline. Motivated by moral outrage and a sense of justice rather than contract details, nudniks use public shaming, complaints, and lawsuits to punish firms for perceived wrongdoings. These actions create broad benefits, making nudniks the "unsung heroes" of consumer markets, even if their individual motivations or actions can sometimes be controversial.

## Key Sections (from `summary.md`)

- **Debunking Traditional Consumer Protection Theories:** Professor Yonathan Arbel of the University of Alabama School of Law writes that the long-held "informed minority theory," suggesting a few knowledgeable consumers discipline markets by reading contracts, is unrealistic and empirically challenged. He notes consumers are often rationally apathetic or unable to understand complex terms, leading to passivity. Furthermore, theories of reputational discipline through online reviews are criticized as underspecified, failing to explain how credible information is produced or disseminated, and often relying on biased or low-participation data. These traditional models inadequately explain consumer power in modern markets, prompting a search for alternative explanations.
- **Introducing the "Nudnik" - A New Model of Consumer Activism:** Professor Yonathan Arbel of the University of Alabama School of Law writes that a new "crusading minority," termed "nudniks," is emerging as a key force in consumer governance. Unlike the supposed "informed minority," nudniks are motivated less by informing others and more by moral outrage, seeking to punish firms for perceived wrongdoings. They operate *ex post*, reacting to disappointments through complaints, reviews, and litigation, often leveraging the internet. This nudnik-driven activism challenges traditional theories by suggesting that consumer power can be effective even when consumers don't read contracts, as sellers are incentivized to improve quality for all to avoid public disputes.
- **Nudniks vs. Other Consumer Archetypes:** Professor Yonathan Arbel of the University of Alabama School of Law writes that "nudniks"—consumers highly prone to vindicating their rights—are distinct from other archetypes. "Passivists," the largest group, are largely inactive. "Shoppers" meticulously compare terms *ex ante* and primarily "exit" to competitors. "Sophisticates" strategically leverage contract terms for personal profit based on cost-benefit analysis. In contrast, nudniks act *ex post* out of principle or moral outrage, often when others would not, and their actions are more likely to generate positive spillovers for all consumers, whereas sophisticates primarily seek private gains that may not benefit, or could even harm, others.
- **Motivations, Methods, and Impact of Nudniks:** Professor Yonathan Arbel of the University of Alabama School of Law writes that nudniks are driven by an innate sense of justice or idiosyncratic motivations, leading them to act even when costs seem to outweigh benefits for an average consumer. They employ tactics like complaints, reviews, lawsuits, and viral social media campaigns (e.g., "United Breaks Guitars") to highlight seller failures. This activism, based on broad transactional expectations rather than specific contract terms, can be initiated by a single individual and amplified by social media. Sellers respond due to legal and reputational risks, often improving service for all, making nudniks effective at solving collective action problems.
- **Addressing Criticisms and Nuances of Nudnik Activism:** Professor Yonathan Arbel of the University of Alabama School of Law writes that while nudnik activism is potent, it's not without concerns. Critics question if complaints are always representative or beneficial, potentially being frivolous or driven by unrealistic expectations. However, he argues these issues may be overstated. Nudniks are a heterogeneous group, and their actions, even if driven by "spite," often highlight genuine consumer interests and can lead to market-wide improvements. The impact of unrepresentative nudniks is also mitigated as other consumers tend to amplify only those grievances they find valid, and sellers actively manage expectations.

## Keywords

contracts; AI; law

## Files

- Full text: `papers/ssrn-3568768/paper.txt`
- PDF: `papers/ssrn-3568768/paper.pdf`
- Summary (EN): `papers/ssrn-3568768/summary.md`
- Summary (ZH): `papers/ssrn-3568768/summary.zh.md`

_Auto-generated study aid. For canonical content, rely on `paper.txt`/`paper.pdf`._
)MW4LLM";
static const std::string STUDY_PACK_MD = R"MW4LLM(# Study pack: DDeePPaauull LLaaww RReevviieeww (ssrn-3568768)

- SSRN: https://papers.ssrn.com/sol3/papers.cfm?abstract_id=3568768
- Full text: `papers/ssrn-3568768/paper.txt`
- Summary (EN): `papers/ssrn-3568768/summary.md`
- Summary (ZH): `papers/ssrn-3568768/summary.zh.md`

## Elevator pitch

Professor Yonathan Arbel of the University of Alabama School of Law argues that traditional consumer protection, thought to rely on an "informed minority" reading contracts, is ineffective. Instead, a new type of activist, the "nudnik" or "crusading minority," drives market discipline. Motivated by moral outrage and a sense of justice rather than contract details, nudniks use public shaming, complaints, and lawsuits to punish firms for perceived wrongdoings. These actions create broad benefits, making nudniks the "unsung heroes" of consumer markets, even if their individual motivations or actions can sometimes be controversial.

## Structured outline (high-signal)

### Debunking Traditional Consumer Protection Theories

Professor Yonathan Arbel of the University of Alabama School of Law writes that the long-held "informed minority theory," suggesting a few knowledgeable consumers discipline markets by reading contracts, is unrealistic and empirically challenged. He notes consumers are often rationally apathetic or unable to understand complex terms, leading to passivity. Furthermore, theories of reputational discipline through online reviews are criticized as underspecified, failing to explain how credible information is produced or disseminated, and often relying on biased or low-participation data. These traditional models inadequately explain consumer power in modern markets, prompting a search for alternative explanations.

### Introducing the "Nudnik" - A New Model of Consumer Activism

Professor Yonathan Arbel of the University of Alabama School of Law writes that a new "crusading minority," termed "nudniks," is emerging as a key force in consumer governance. Unlike the supposed "informed minority," nudniks are motivated less by informing others and more by moral outrage, seeking to punish firms for perceived wrongdoings. They operate *ex post*, reacting to disappointments through complaints, reviews, and litigation, often leveraging the internet. This nudnik-driven activism challenges traditional theories by suggesting that consumer power can be effective even when consumers don't read contracts, as sellers are incentivized to improve quality for all to avoid public disputes.

### Nudniks vs. Other Consumer Archetypes

Professor Yonathan Arbel of the University of Alabama School of Law writes that "nudniks"—consumers highly prone to vindicating their rights—are distinct from other archetypes. "Passivists," the largest group, are largely inactive. "Shoppers" meticulously compare terms *ex ante* and primarily "exit" to competitors. "Sophisticates" strategically leverage contract terms for personal profit based on cost-benefit analysis. In contrast, nudniks act *ex post* out of principle or moral outrage, often when others would not, and their actions are more likely to generate positive spillovers for all consumers, whereas sophisticates primarily seek private gains that may not benefit, or could even harm, others.

### Motivations, Methods, and Impact of Nudniks

Professor Yonathan Arbel of the University of Alabama School of Law writes that nudniks are driven by an innate sense of justice or idiosyncratic motivations, leading them to act even when costs seem to outweigh benefits for an average consumer. They employ tactics like complaints, reviews, lawsuits, and viral social media campaigns (e.g., "United Breaks Guitars") to highlight seller failures. This activism, based on broad transactional expectations rather than specific contract terms, can be initiated by a single individual and amplified by social media. Sellers respond due to legal and reputational risks, often improving service for all, making nudniks effective at solving collective action problems.

### Addressing Criticisms and Nuances of Nudnik Activism

Professor Yonathan Arbel of the University of Alabama School of Law writes that while nudnik activism is potent, it's not without concerns. Critics question if complaints are always representative or beneficial, potentially being frivolous or driven by unrealistic expectations. However, he argues these issues may be overstated. Nudniks are a heterogeneous group, and their actions, even if driven by "spite," often highlight genuine consumer interests and can lead to market-wide improvements. The impact of unrepresentative nudniks is also mitigated as other consumers tend to amplify only those grievances they find valid, and sellers actively manage expectations.

### Conclusion: The Unsung Heroes of Consumer Markets

Professor Yonathan Arbel of the University of Alabama School of Law writes that despite potential downsides, "nudniks" or "crusading consumers" generally have a net positive social impact, filling an important gap in marketplace discipline left by the failure of traditional theories. He posits that focusing on these persistent individuals reveals how a minority can effectively pressure sellers to improve services for all, through legal and reputational incentives. Consequently, legal literature should shift focus from the "informed minority" to these "nudniks," who, despite often facing negative perceptions, serve as the "unsung heroes" of consumer markets, though more research is needed.

### A Note on Cost-Plus No-Fee (CPNF) Contracting

Professor Yonathan Arbel of the University of Alabama School of Law writes that, in discussing mechanisms of accountability, Cost-Plus No-Fee (CPNF) contracting offers an interesting parallel. This model challenges traditional contract theory by having agents work at cost, relying on trust and reputational mechanisms instead of explicit profit incentives. He suggests CPNF can function as an "information-forcing" mechanism. By compelling agents to be transparent about their costs and efforts, CPNF helps mitigate issues of asymmetric information often found in principal-agent relationships, echoing themes of trust and non-standard incentives relevant to market governance.

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
static const std::string ARTICLE_TEXT = R"MW4LLM(DDeePPaauull LLaaww RReevviieeww
Volume 69
Article 3
Issue 2 Winter 2020
CCoonnssuummeerr AAccttiivviissmm:: FFrroomm TThhee IInnffoorrmmeedd MMiinnoorriittyy TToo TThhee
CCrruussaaddiinngg MMiinnoorriittyy
Yonathan A. Arbel
Roy Shapira
Follow this and additional works at: https://via.library.depaul.edu/law-review
Part of the Law Commons
RReeccoommmmeennddeedd CCiittaattiioonn
Yonathan A. Arbel & Roy Shapira, Consumer Activism: From The Informed Minority To The Crusading
Minority, 69 DePaul L. Rev. (2020)
Available at: https://via.library.depaul.edu/law-review/vol69/iss2/3
This Article is brought to you for free and open access by the College of Law at Via Sapientiae. It has been
accepted for inclusion in DePaul Law Review by an authorized editor of Via Sapientiae. For more information,
please contact digitalservices@depaul.edu.
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 1 21-APR-20 11:48
CONSUMER ACTIVISM: FROM THE INFORMED
MINORITY TO THE CRUSADING MINORITY
Yonathan A. Arbel and Roy Shapira*
CONTENTS
INTRODUCTION................................................. 234 R
I. CONSUMER GOVERNANCE VIA INFORMED MINORITY
AND REPUTATIONAL SANCTIONS........................ 237 R
A. The Limits of Consumer Governance ............... 237 R
B. The Informed Minority Theory and Its Limits ...... 240 R
C. Reputational Discipline Theory and Its Limits ...... 242 R
II. THE CRUSADING MINORITY THEORY................... 243 R
A. The Nudnik and Other Types of Consumers ........ 244 R
B. Nudniks in Action: Motivating Examples............ 250 R
C. Why do Nudniks Prevail where the Informed
Minority Fails?...................................... 255 R
D. Why and How Sellers Accommodate Nudniks’
Concerns............................................ 259 R
III. LIMITATIONS OF NUDNIK-BASED CONSUMER
GOVERNANCE........................................... 261 R
CONCLUSION ................................................... 266 R
Legal scholars have long recognized that market norms are respected
not only because of consumer protection laws, but also because of
internal market dynamics. Consumers, the argument goes, fend for
themselves and hold sellers accountable. But how exactly do consumers
discipline sellers? The most influential model has been the informed
minority theory, according to which a critical mass of informed
consumers reads and negotiates contracts in advance, thereby
pressuring sellers to offer better contracts to all consumers. Recent
empirical studies, however, cast doubt on the existence of such a mass,
leading many to view the informed minority theory as unrealistic. What,
* The University of Alabama, School of Law; Interdisciplinary Center (IDC). We thank par-
ticipants at the Clifford Symposium Rising Stars: A New Generation of Scholars Looks at Civil
Justice, as well as Lisa Bernstein, Eric Goldman, Stephen Laandsman, Ben McMichael, Tony
Sebok, Catherine Sharkey, Steve Shavell, and Rory Van Loo for helpful comments and discus-
sions. Cade McGavin Brown provided invaluable research assistance.
233
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 2 21-APR-20 11:48
234 DEPAUL LAW REVIEW [Vol.69:233
then, may explain bottom-up governance in a world where consumers
do not read contracts?
In this contribution to the Clifford Symposium, we aim at exposing a
different mechanism of market discipline: one that works not through
ex ante reading and negotiating, but rather through ex post pressures to
meet buyers’ expectations. We specifically emphasize the role of a small
subset of consumers that we dub “nudniks.” Nudniks are those
consumers who call in to complain, fill out satisfaction surveys, post
online reviews, and file lawsuits. Driven by an innate sense of justice
and atypical motivations, these nudniks act as crusading consumers
against underperforming sellers. Through their actions, nudniks direct
attention to seller failure, leading to a variety of formal and informal
sanctions, thus presenting a more realistic form of consumer activism in
today’s overwhelming information environment.
INTRODUCTION
Market discipline comes not only from legal protections, but also
from consumers themselves. Understanding the effectiveness of con-
sumer-driven market discipline mechanisms is key, as it dictates the
scope and design of legal interventions.
The leading theory of market-based discipline has traditionally been
the informed minority theory.1 The theory concedes that most con-
sumers lack sophistication or time to read their contracts and shop for
better terms.2 Yet, it suggests that consumer-based governance of
market discipline can be powerful.3 As long as a minority of consum-
ers are engaged with these aspects of the transaction, one could still
expect sellers to provide favorable terms. Sellers would compete over
who wins the segment of informed buyers, and in the process will have
to modify their standard form contracts in ways that benefit the entire
consumer body, or so the theory goes. While enjoying large influence,
over the years, the informed minority theory has encountered increas-
ing opposition. Perhaps most critically, recent empirical evidence sug-
gests that the number of consumers that actually read and understand
contracts is too low to justify a change in sellers’ behavior.4 Even the
1. Alan Schwartz & Louis L. Wilde, Intervening in Markets on the Basis of Imperfect Informa-
tion: A Legal and Economic Analysis, 127 U. PA. L. REV. 630, 655 (1979).
2. Id. at 642.
3. We use the term “consumer governance” in ways that bear similarities to the more oft-used
“corporate governance” term: denoting the set of formal and informal rules that govern the
interactions between sellers and buyers.
4. Yannis Bakos et al., Does Anyone Read the Fine Print? Consumer Attention to Standard
Form Contracts, 43 J. LEGAL STUD. 1, 2 (2014) (providing empirical data that few consumers
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 3 21-APR-20 11:48
2020] CONSUMER ACTIVISM 235
theory’s progenitors now seem to question its practicality.5 This has
left a gap in our understanding of market discipline through consumer
governance: If market discipline does not come from a critical mass of
informed readers, where does it come from?
This Essay suggests looking elsewhere: Instead of focusing on buy-
ers who read and negotiate before the purchase, focus on buyers who
feel compelled to respond strongly whenever sellers disappoint. In-
stead of focusing on avid readers, focus on avid “enforcers”—those
consumers who demand to speak with the manager, fill out satisfac-
tion surveys, post online reviews, and file lawsuits. We dub these con-
sumers “nudniks.”6 Nudniks do not operate like most of us. They
possess an innate sense of justice, atypical motivations, and an idio-
syncratic cost structure, which lead them to fight sellers who disap-
point—even in situations where most of us would not notice, or notice
and stay passive. Nudniks are often perceived as petty and vindictive.
Yet, through their actions, nudniks provide an important public ser-
vice: directing attention to failures in the market, thus leading to a
variety of formal and informal sanctions against misbehaving firms. In
other words, nudniks generate underappreciated spillover effects that
reverberate throughout the economy. This Essay explores the role of
nudniks in the enforcement of market norms and consumer govern-
ance, evaluates their social contribution, and suggests this “crusading
minority” of nudniks as a missing piece in theories of consumer mar-
ket governance.
This Essay argues that consumer activism predicated on a crusading
minority of nudniks, who notice seller misbehavior and respond to it
through legal-reputational channels, is a more realistic depiction of
how market discipline works than the informed minority theory.7
Nudniks complain and fight sellers publicly regardless of whether they
read the contract in advance. They often complain based on their
transactional expectations from the seller. And transactional expecta-
read End Users License Agreements); Shmuel I. Becher & Esther Unger-Aviram, The Law of
Standard Form Contracts: Misguided Intuitions and Suggestions for Reconstruction, 8 DEPAUL
BUS. & COM. L.J. 199, 206 (2010) (providing empirical data that most consumers are not likely to
read contracts ex ante); Clayton P. Gillette, Rolling Contracts as an Agency Problem, 2004 WIS.
L. REV. 679, 680 (2004) (“[C]ommentators agree that buyers, or the vast majority of them, do
not read the terms presented to them by sellers.”). We note that most of the evidence is focused
on online contracts and in specific domains; more work is needed in other areas.
5. See infra note 35.
6. The word derives from Yiddish and can be loosely translated to “a busybody.” See infra
Part II.A.
7. Consumer activism here denotes activism with respect to the properties of the good, ser-
vice, contract, or transaction. We do not deal here with consumer activism with respect to social
or political goals.
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 4 21-APR-20 11:48
236 DEPAUL LAW REVIEW [Vol.69:233
tions are a function not only of the explicit terms in the contract, but
also of sellers’ oral representations, advertisements, market norms,
fairness standards, and so on.8 Even if the seller is contractually pro-
tected by a disclaimer nestled in the fine print, she will anticipate the
potential risk that comes from entering a public battle with nudniks
and may find it best to deliver better service ex ante.
Such a nudnik-driven mode of consumer activism creates positive
spillovers, but also comes with social costs. Some nudniks pursue nar-
row interests that do not benefit the rest of the consumer body and
impose unnecessary costs on sellers. While we do not venture to offer
a conclusive quantification of the net effect of nudniks, we do offer
here a synthesis of findings from the consumer complaining behavior
literature, suggesting that many nudniks positively contribute to the
market. At the minimum, our analysis suggests that legal scholars and
policymakers should pay more attention to nudniks’ effects.
The nudniks are a response to the problems with the informed mi-
nority theory. This theory essentially rests on two assumptions, re-
garding the what and the how of seller behavior. First, what sellers do:
The theory assumes that sellers compete over a small segment of con-
sumers who read the contract and care about its terms. Second, how
they do it: To win the segment of readers, sellers have to offer better
terms to all consumers across the board. Sellers operate through stan-
dard-form contracts, and cannot tell which consumer is a reader and
which is not before the fact; therefore, they are forced to offer better
terms for all. In this Essay, we respond to the first premise. Many have
taken the recent empirical evidence of low readership rates to as un-
dermining the possibility of internal market discipline. This Essay sug-
gests that market discipline does not have to be predicated on
consumers reading the contract before purchasing; it can also come
from consumers noticing and complaining publicly about sellers who
fail to meet consumers’ transactional expectations, regardless of the
contract. In a separate paper, we confront the second assumption of
the informed minority theory: the premise that sellers cannot distin-
guish between active and passive consumers.9 In today’s world, we ar-
gue there, sellers can, and to a growing extent already do, employ big
8. Contract law, and in particular, the Uniform Commercial Code, is sensitive to background
expectations, which form the penumbra of the rights and obligations the parties owe each other.
See RESTATEMENT (SECOND) OF CONTRACTS §§203(B), 211(A), (C), 220–23; U.C.C. §§1-303,
2-208(2).
9. See Yonathan A. Arbel & Roy Shapira, Theory of the Nudnik: The Future of Consumer
Activism and What We Can Do to Stop It, VAND. L. REV. (forthcoming 2020) [hereinafter Arbel
& Shapira, Theory of the Nudnik]. Another important contribution of this paper is the classifica-
tion of the nudnik and the identification of its role within theories of consumer law.
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 5 21-APR-20 11:48
2020] CONSUMER ACTIVISM 237
data tools to tell which consumer is most likely to be nudnik, and then
cater to these consumers personally.
This Essay proceeds in three parts. Part I explores the leading theo-
ries of market-based, consumer governance mechanisms and their
shortcomings. Part II suggests a new direction for thinking of market-
based consumer governance. Instead of counting how many consum-
ers read contracts, we need to shift attention to consumer dissatisfac-
tion behavior: How many consumers complain after the purchase?
How do other potential consumers react to these complaints? What
impact do such complaints have on sellers? We emphasize the rise of
the internet and social media as factors that greatly empowered
nudniks and increased their potential reach. As long as sellers are un-
able to spot nudniks in advance,10 they are incentivized to provide
higher-quality service to all consumers ex ante, so as not to risk the
reputation and legal risk that comes with nudniks. Part III evaluates
the shortcomings of nudnik-based activism. We conclude that while
not all nudnik-activity is socially beneficial, overall there is reason to
believe that nudniks are the unsung heroes of market governance.
I. CONSUMER GOVERNANCE VIA INFORMED MINORITY
AND REPUTATIONAL SANCTIONS
When parties enter into a contract, they assumedly select the terms
that advance their mutual goals.11 Yet in the context of consumer con-
tracts, this standard assumption too often does not apply: For a variety
of reasons, consumers are limited in their negotiation, enforcement,
and monitoring of contract terms, thus creating an opportunity for
sellers to offer inferior, one-sided terms. Do sellers take advantage of
this opportunity? If not, why? This Part provides a quick overview of
the extant literature on consumer governance. We start by noting the
factors that limit consumers’ ability to monitor sellers’ behavior di-
rectly. We then detail the two most influential theories of how con-
sumers can nevertheless discipline sellers—the informed minority
theory and the reputational discipline theory.
A. The Limits of Consumer Governance
We use the term “consumer governance” here to denote the idea
that consumers can exert pressure on sellers who “misbehave,”
10. For an exploration of how firms respond to the challenge posed by the nudnik based on
big-data and predictive algorithms, see id.
11. See, e.g., CHARLES FRIED, CONTRACTAS PROMISE 8 (1981); STEVEN SHAVELL, FOUNDA-
TIONSOF ECONOMIC ANALYSISOF LAW 293 (2004).
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 6 21-APR-20 11:48
238 DEPAUL LAW REVIEW [Vol.69:233
thereby disciplining sellers’ behavior and participating in the govern-
ance of market norms. Thus defined, consumer governance runs into
several well-known issues that ostensibly limit its effectiveness.
First, consumers often lack the necessary sophistication to under-
stand the terms of their contracts.12 The contractual language is com-
plex, the grammar is convoluted, the vocabulary is full of legalese, and
even the format tends to be quite jarring.13 The substance itself can
also be difficult to grasp, given the complex nature of some common
transactions (think, e.g., about the number of the parties involved in a
typical home purchase agreement). These difficulties present a chal-
lenge to many consumers, especially in groups that suffer from low
rates of financial and legal literacy.14
Second, many consumers are apathetic about the terms of their con-
tracts.15 Apathy is said to be rational if the costs of being engaged
outweigh the benefits. For many consumers, this is indeed the case:
The costs of reading contracts are immediate and certain, namely
one’s time and effort. The benefits of reading, by contrast, are remote
and uncertain. Even if a consumer identifies an unfavorable choice-of-
law clause, the odds of this clause mattering is quite low for any indi-
vidual consumer.16 Moreover, the consumer will often lack the bar-
gaining power necessary to negotiate any of the terms in contracts that
are mostly based on standard forms. This makes reading the contract a
losing proposition from the consumer’s standpoint in many cases.17 In
12. Michael I. Meyerson, The Efficient Consumer Form Contract: Law and Economics Meets
the Real World, 24 GA. L. REV. 583, 598–99 (1990) (noting consumer inability to discern legal
meaning of contractual terms, even those in plain language, due to high costs).
13. See Yonathan A. Arbel & Andrew Toler, All-Caps (U. of Ala. Legal Studies, Research
Paper No. 3519630, Jan. 15, 2020) (finding that all-caps blocks of text in contracts fail to improve
the quality of consumer consent and make it worse for older readers), available at http://
papers.ssrn.com/sol3/papers.cfm?abstract_id=3519630; Uri Benoliel & Shmuel I. Becher, The
Duty to Read the Unreadable, 60 B.C. L. REV. 2255 (2019).
14. See generally Annamaria Lusardi & Olivia S. Mitchell, Financial Literacy: An Overview,
10 J. PENSION ECON. & FIN. 297 (2011). See also Annamaria Lusardi & Olivia S. Mitchell, Finan-
cial Literacy and Retirement Planning in the United States, 10 J. PENSION ECON. & FIN. 509, 512
(2011) (finding, in a survey of American adults, that only 30% could answer correctly three basic
questions of financial literacy); Judy T. Lin et al., The State of U.S. Financial Capability: The 2018
National Financial Capability Study, https://www.usfinancialcapability.org/downloads/NF-CS_
2018_Report_Natl_Findings.pdf (finding lower rates of financial literacy among minorities).
15. See generally William M. Landes & Richard A. Posner, The Private Enforcement of Law, 4
J. LEGAL STUD. 1, 33 (1975); Roger Van Den Bergh & Louis Visscher, The Preventive Function
of Collective Actions for Damages in Consumer Law, 1 ERASMUS L. REV. 5, 6 (2008).
16. See Melvin Aron Eisenberg, The Limits of Cognition and the Limits of Contract, 47 STAN.
L. REV. 211, 243 (1995). For the consumer, a 1-in-100 chance of litigation is a remote possibility.
For a firm serving 100 consumers, it is a high likelihood.
17. Robert A. Hillman & Jeffrey J. Rachlinski, Standard-Form Contracting in the Electronic
Age, 77 N.Y.U. L. REV. 429, 445 (2002).
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 7 21-APR-20 11:48
2020] CONSUMER ACTIVISM 239
other cases, consumers can suffer from the other side, namely, irra-
tional apathy. That is, even if it is worthwhile for them to read and
negotiate terms, they are bound by a variety of cognitive limitations
and biases that would prevent them from reading. For example, con-
sumers may act myopically, by failing to consider future possibilities
or overly discount future events.18
Finally, consumers often ignore and tend to remain passive when
sellers disappoint. That is, consumers are not just passive pre-
purchase, in reading and negotiating, but also passive post-purchase.
Part of the reason for this passivity is that consumers are often una-
ware that their contractual rights were violated.19 Even when consum-
ers are sophisticated and sufficiently aware, they may opt to do
nothing, simply because they consider taking action to be too costly.
This is especially true for bringing lawsuits and waging a legal battle.
The costs of bringing a lawsuit, the risk of losing the case, and the
difficulty of collecting judgments if you have won, all lead consumers
to frequently abandon the pursuit of rights that they know were not
met.20 In fact, recent work suggests that reading the contract after the
consumer experienced a breach may actually make the consumer less
likely to complain, as sellers include unenforceable and otherwise mis-
leading terms in their contracts, which cause buyers to give up.21
On paper, sellers would be aware of the confluence of these
problems and offer buyers inferior terms in their contracts ex ante,
and fail to deliver on obligations ex post. Such seller behavior may
easily lead to an eventual breakdown of consumer trust—akin to a
“lemons problem,” whereby deep mistrust prevents many desirable
transactions.22 What stops this supposed race to the bottom? How do
consumer markets function given all these inherent problems? The le-
gal literature has offered several theories in response. The next two
Sections elaborate.
18. See, e.g., Oren Bar-Gill, The Behavioral Economics of Consumer Contracts, 92 MINN. L.
REV. 749, 755 (2008); Christine Jolls, Cass R. Sunstein & Richard Thaler, A Behavioral Ap-
proach to Law and Economics, 50 STAN. L. REV. 1471, 1476–80 (1998).
19. From experience teaching this subject (Arbel), even most law students are unaware of the
implied warranty of fitness for a particular purpose before it is covered in class. See U.C.C §2-
315 (AM. LAW INST. & UNIF. LAW COMM’N, amended 2011).
20. For a concrete example from the most common type of consumer cases—debt collection
cases—see generally Yonathan A. Arbel, Adminization: Gatekeeping Consumer Contracts, 71
VAND. L. REV. 121, 130–42 (2018).
21. Meirav Furth-Matzkin, On the Unexpected Use of Unenforceable Contract Terms: Evi-
dence from the Residential Rental Market, 9 J. LEGAL ANALYSIS 1, 3 (2017).
22. See George A. Akerlof, The Market for “Lemons”: Quality Uncertainty and the Market
Mechanism, 84 Q. J. ECON. 488 (1970).
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 8 21-APR-20 11:48
240 DEPAUL LAW REVIEW [Vol.69:233
B. The Informed Minority Theory and Its Limits
The most influential response to the abovementioned concerns has
been Schwartz & Wilde’s informed minority theory.23 The informed
minority theory readily concedes that the majority of consumers are
unsophisticated and do not engage with the contracts before they
purchase. Still, the theory argues that a minority of consumers are so-
phisticated and do read their contracts carefully and negotiate the
terms before purchasing. This informed minority diligently compares
the product and its terms to those offered by competitors. The efforts
of the minority create demand-side pressure on firms to offer better
contractual terms, as doing so will allow firms to win this segment of
the market. Now, either because many consumer contracts are stan-
dard-form agreements or because firms cannot distinguish between
consumers on the basis of their sophistication and tendency to read,24
the way a firm can win the hearts of the informed minority is by offer-
ing better terms across the board.25 Consequently, competitive pres-
sures created by an informed minority push the entire market towards
a more consumer-friendly equilibrium where firms are competing not
only on price, but also on the quality of their contracts.
The informed minority theory thus explains why despite the lack of
sophistication on the part of many consumers, contractual terms are
not the worst possible terms conceivable. The diligence of the in-
formed minority is a bulwark against sellers’ tendency to grossly favor
themselves.
While highly influential, the informed minority theory runs into the-
oretical and empirical problems.26 For the theory to work, there must
be a sufficiently sizable minority, a critical mass of readers (after all,
23. Alan Schwartz & Louis L. Wilde, Intervening in Markets on the Basis of Imperfect Infor-
mation: A Legal and Economic Analysis, 127 U. PA. L. REV. 630 (1979). For a review of its
influence, see generally Eyal Zamir, Contract Law and Theory: Three Views of the Cathedral, 81
U. CHI. L. REV. 2077 (2014) (reviewing DOUGLAS G. BAIRD, RECONSTRUCTING CONTRACTS
(2013); BRIAN H. BIX, CONTRACT LAW: RULES, THEORY, AND CONTEXT (2012)); R. Ted Cruz &
Jeffrey J. Hinck, Not My Brother’s Keeper: The Inability of an Informed Minority to Correct for
Imperfect Information, 47 U.C. HASTINGS L. J. 635 (1996).
24. Elsewhere we focus on the ways the assumption of inability to screen consumers might
break, given big data and predictive analytics, and detail the alarming consequences. Arbel &
Shapira, Theory of the Nudnik, supra note 9.
25. See also George L. Priest, A Theory of the Consumer Product Warranty, 90 YALE L.J.
1297, 1347 (1981) (“If a small group of consumers reads warranties and selects among products
according to warranty content, manufacturers may be forced to draft warranties responsive to
the group’s preferences, even though the large majority of consumers generally neglect warranty
terms.”).
26. For examples of notable objections, see Shmuel I. Becher, Asymmetric Information in
Consumer Contracts: The Challenge That Is Yet to Be Met, 45 AM. BUS. L.J. 723, 735–54 (2008);
Cruz & Hinck, supra note 23; Zamir, supra note 23; Jeff Sovern, Toward a New Model of Con-
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 9 21-APR-20 11:48
2020] CONSUMER ACTIVISM 241
why should sellers alter their form contracts in order to win the busi-
ness of only a handful of reading consumers?). But consumers gener-
ally dislike reading contracts,27 and the private gains from reading can
be quite marginal—there is limited opportunity to negotiate terms,
competitors might not offer significantly better terms, and there is a
good chance that the covered provision will never materialize.28 While
the costs are real and private, the benefits of reading are more remote
and public—leading each individual consumer to attempt to free-ride
the efforts of others.29 Some recent empirical work now documents
the theoretical prediction: The number of people who actually read
contracts is too small to reach the critical mass needed for the in-
formed minority mechanism to work.30
Even if the informed minority theory was plausible when Schwartz
and Wilde first developed it in the 1970s, the increase in length and
complexity of consumer contracts occasioned by the rise of the digital
age has gradually rendered the theory less and less fit for today’s
world.31 Contracts nowadays are also encumbered by the rise of ever-
increasing disclosures, which compete over, and overload, limited con-
sumer attention.32 And while the non-law-and-economics scholars
have long been skeptical of the theory,33 nowadays even the law and
economics community loses faith.34 Exhibit A: Schwartz himself
seems to believe that nobody reads contracts these days.35
sumer Protection: The Problem of Inflated Transaction Costs, 47 WM. & MARY L. REV. 1635,
166–72 (2006).
27. See, e.g., Omri Ben-Shahar, The Myth of the ‘Opportunity to Read’ in Contract Law
(Coase-Sander Working Paper Series in L. & ECON. No. 415, 2008), https://chicagounbound.uchi
cago.e-du/law_and_economics/549/. For a review of the literature on the no-reading problem, see
Arbel & Toler, supra note 13.
28. See supra notes 15–18 and accompanying text.
29. Id.
30. Bakos et al., supra note 4, at 4 (“We find that the fraction of consumers who read such
contracts is so small that it is unlikely that an informed minority alone is shaping software license
terms.”).
31. See Todd D. Rakoff, Contracts of Adhesion, An Essay in Reconstruction, 96 HARV. L. REV
1173, 1179 n.22 (1983); Zamir, supra note 23, at 2102–03; Robert A. Hillman & Jeffrey J. Rach-
linski, Standard-Form Contracting in the Electronic Age, 77 N.Y.U. L. REV. 429, 448 (2002).
32. SeeOMRI BEN-SHAHAR & CARL E. SCHNEIDER, MORE THAN YOU WANTEDTO KNOW:
THE FAILUREOF MANDATED DISCLOSURE 94–101 (2014).
33. See Rakoff, supra note 31; Zamir, supra note 23; Hillman & Rachlinski, supra note 31.
34. Cruz & Hinck, supra note 23.
35. Ian Ayres & Alan Schwartz, The No-Reading Problem in Consumer Contract Law, 66
STAN. L. REV. 545, 552 (2014) (“[T]he state should jettison the disclosure project of making all
terms accessible to consumers with the expectation that consumers can read the entire docu-
ment.”). See also Ian Ayres & Eyal Zamir, Mandatory Rules 12 n.54 (Hebrew Univ. of Jerusalem
Legal Res. Paper No. 19-12), https://papers.ssrn.com/sol3/papers.cfm?abstract_id=3420179 (not-
ing the absence of the informed minority theory from Schwartz’ recent work).
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 10 21-APR-20 11:48
242 DEPAUL LAW REVIEW [Vol.69:233
C. Reputational Discipline Theory and Its Limits
Other theories of bottom-up market discipline usually invoke the
concept of reputational sanctions.36 Legal scholars and economists
have suggested that reputational forces play an increasingly important
role in today’s consumer markets, given the rise of the internet and
social media.37 The idea is that a seller who underperforms can expect
buyers to post negative online reviews, thereby causing other poten-
tial buyers to avoid purchasing from the seller in the future. The pros-
pect of negative reputational information deters such seller
misbehavior ex ante, or so the argument goes.
The appeal of the reputational discipline theory is quite clear. And
indeed, there exists evidence that both negative and positive reputa-
tional information affect business revenues.38 However, theories that
invoke reputational discipline are usually underspecified and rely too
much on strong assumptions. Sure, reputation matters. But how ex-
actly is reputation produced? Who disseminates details of private in-
teractions between a seller and a buyer? Is the information considered
credible? Do other buyers act on it? Existing accounts do not develop
satisfying answers to these questions. They implicitly assume that dis-
satisfied consumers put in motion meaningful reputational sanctions
by transparently and mechanically sharing details of their dissatisfac-
tion with others online. But given that accurate reputational informa-
tion is a public good, why would an individual consumer find it
worthwhile to share this information with others? And why would
other consumers read, believe, and decide to act based on such infor-
mation from another buyer (who, in all likelihood, they never met)?
Writing reviews comes with certain costs—the time it takes to write
the review, the legal risk involved in defamation lawsuits,39 the social
36. On reputational sanctions more generally see Roy Shapira, A Reputational Theory of Cor-
porate Law, 26 STAN. L. & POL’Y REV. 1 (2015); Roy Shapira, Reputation through Litigation:
How the Legal System Shapes Behavior by Producing Information, 91 WASH. L. REV. 1193
(2016); Yonathan A. Arbel, Reputation Failure: The Limits of Market Discipline in Consumer
Markets, WAKE FOREST L. REV., 2019 (Univ. of Ala. Legal Stud. Res. Paper No. 3239995),
https://papers.ssrn.com/sol3/papers.cfm?abstract_id=3239995 [hereinafter Arbel, Reputation
Failure].
37. See, e.g., Alex Tabarrok & Tyler Cowen, The End of Asymmetric Information, CATO UN-
BOUND (2015); Adam Thierer et al., How the Internet, the Sharing Economy, and Reputational
Feedback Mechanisms Solve the “Lemons Problem,” 70 UNIV. MIAMI L. REV. 830, 830 (2016).
38. Michael Luca, Reviews, Reputation, and Revenue: The Case of Yelp.com 10, 12 (Harv. Bus.
Sch. Working Paper No. 12-016, 2016), https://www.hbs.edu/faculty/Publication%20Files/12-
016_a7e4a5a2-03f9-490d-b093-8f951238dba2.pdf.
39. See Yonathan A. Arbel, Reputation Failure: The Limits of Market Discipline in Consumer
Markets, 54 WAKE FOREST L. REV. 1239 (2019) (noting how consumer reviews can lead to litiga-
tion), available at https://papers.ssrn.com/sol3/papers.cfm?abstract_id=3239995.
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 11 21-APR-20 11:48
2020] CONSUMER ACTIVISM 243
pressure to conform, and so on.40 Indeed, evidence shows that few
consumers actually write reviews, and among those who choose to do
so, there is a strong self-selection that might bias reputational infor-
mation in unpredictable ways.41 The fast-growing body of evidence on
online reviews suggests that consumer-sourced reputational informa-
tion can be an unreliable guide to future consumers, and therefore a
weaker restraint on firm behavior than proponents admit.42
To emphasize, we do not claim here that reputational forces are un-
important or insufficiently potent. Both of us have written extensively
on the important role that reputation plays in market discipline.43 Our
point here is that too often consumer governance theories that invoke
reputation are under-specified. This gap calls for an explanation that
current theories fail to provide.
*
For bottom-up market governance to emerge, consumers must
wield sufficient power. Yet most consumers remain uninformed and
unengaged. The public good nature of market discipline makes it sus-
ceptible to free riding and consumer collective action problems.
Where does effective consumer activism come from, then? One influ-
ential theory suggested that a critical mass of consumers who do read
contracts make sellers change their behavior toward all other consum-
ers. Yet the accumulated evidence suggests that in many markets such
a critical mass does not exist. Another influential theory suggests that
dissatisfied consumers will complain online, thereby creating a reputa-
tional risk for sellers. However, this reputational theory fails to ex-
plain who invests in creating and diffusing credible information that
leads other consumers to stop purchasing from a given seller, why
they do so, and how. Accordingly, there is a gap in our understanding
of how consumer governance works. The next Part suggests a way to
narrow this gap by examining the role of a small subset of consumers
who notice and fight back whenever sellers underperform.
II. THE CRUSADING MINORITY THEORY
Bottom-up market discipline relies on the work of a small subset of
active consumers. Among those activists, we focus here on a specific
type that we call nudniks. Part II.A offers a typology of different types
of consumers to explain what exactly we mean by nudniks. Part II.B
40. Arbel, Reputation Failure, supra note 36.
41. Id. at 6.
42. Id. at 14–15, 33.
43. Id.; ROY SHAPIRA, LAWAND REPUTATION (forthcoming 2020).
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 12 21-APR-20 11:48
244 DEPAUL LAW REVIEW [Vol.69:233
offers a few motivating examples of nudniks in action. Part II.C ex-
plains why consumer activism driven by nudniks is a more realistic
possibility in today’s world than consumer activism driven by an in-
formed minority. Finally, Part II.D elaborates on how nudniks bring
change in seller behavior.
A. The Nudnik and Other Types of Consumers
Originally deriving from Yiddish, “nudnik” can be translated as “a
bore, a nag, a jerk”44 or a “busybody.” While all these terms carry
negative connotations, we use the relatively unfamiliar nudnik term in
a neutral way.45 A nudnik, in our framework, denotes a consumer who
is likely to vindicate her transactional rights. When she feels that her
rights were breached, she will not “let it go” until she has addressed
the issue, even if most other consumers will not do so. The nudnik is
the type of consumer who will demand to speak with the manager,
write an angry letter to the editor, or bring a lawsuit over a torn pair
of pants that cost $40.
To understand the attributes and the role that nudniks play in con-
sumer markets, it is useful to consider nudniks alongside various pro-
totypes of consumers. We can roughly separate the different types of
consumers into four categories: “Passivists,” “Shoppers,” “Sophisti-
cates,” and “Nudniks.” A caveat is in order at the outset: Each of the
categories inevitably generalizes, and the lines are murky. Still, for our
modest purposes here—understanding what makes nudniks unique—
the rough categorization works. Figure 1 schematically illustrates this
classification.
44. Nudnik, https://en.wiktionary.org/wiki/nudnik (last updated Oct. 14, 2019).
45. Internet jokes, while ephemeral, capture public sentiment, and so it is telling of wider
public reaction that a common recent ‘meme’ involves nudniks, and derides them for being privi-
leged and entitled. See, e.g., “Speak to the Manager” Haircut, KNOWYOURMEME (2015), https://
know-yourmeme.com/memes/speak-to-the-manager-haircut (last visited Nov. 27, 2019). For the
maltreatment of nudniks in the marketing literature and more generally, see LEON G. SCHIFF-
MAN & JOSEPH L. WISENBLIT, CONSUMER BEHAVIOR 44 (11th ed. 2015).
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 13 21-APR-20 11:48
2020] CONSUMER ACTIVISM 245
FIGURE 1: Species in the Homo Consumerus Genus
Sophisticates Shoppers
Passivists
Nudniks
Passivists are the largest group of consumers—these are those con-
sumers who tend to engage with products only at a basic level, taking
minimal action both when shopping for products and when problems
arise. In selecting products, passivists engage in a rudimentary price
and term comparison. When problems with the product arise (e.g., a
small overcharge, late shipping), they will not always notice, or note
the problem but do little about it. The most action a passivist would
take in response to service failure is refraining from purchasing the
product again or complaining in a way that does not entail much ef-
fort, such as asking the service representative about the issue.
The marketing literature has long documented that most consumers
are passivists.46 Marketing scholar Professor Richard Oliver summa-
rizes the typical consumer behavior by stating that “[c]onsumers do
not do anything, in the main, in response to consumption.”47 The 2006
Retail Customer Dissatisfaction Study indeed found that only six per-
46. SCHIFFMAN & WISENBLIT, supra note 45, at 421 (“Research indicates that only a few un-
satisfied customers actually complain.”); Jean-Charles Chebat et al., Silent Voices: Why Some
Dissatisfied Consumers Fail to Complain, 7 J. SERV. RES. 328 (2005); Clay M. Voorhees et al., A
Voice From the Silent Masses: An Exploratory and Comparative Analysis of Noncomplainers, 34
J. ACAD. MARK. SCI. 514, 514 (2006) (“The majority of dissatisfied customers fail to complain
. . . .”); TECH. ASSISTANCE RESEARCH PROGRAM INST. & U.S. OFFICEOF CONSUMER AFFAIRS,
CONSUMER COMPLAINT HANDLINGIN AMERICA: AN UPDATE STUDY (1979) (finding that 96%
do not complain); Stephen S. Tax & Stephen W. Brown, Recovering and Learning From Service
Failure, 40 MGMT. REV. 75, 75–88 (1998) (finding that 90% do not complain).
47. RICHARD L. OLIVER, SATISFACTION: A BEHAVIORAL PERSPECTIVE ON THE CONSUMER
385 (2d ed. 2015).
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 14 21-APR-20 11:48
246 DEPAUL LAW REVIEW [Vol.69:233
cent of consumers who experienced a problem voiced it to the firm.48
These rates increased somewhat when the value of the good was
higher, but the base rates of consumers who take action rarely ex-
ceeded ten percent.49
The paucity of active consumers is not surprising.50 Standing up for
one’s rights comes at an immediate cost. It involves social discord and
may require a considerable investment of time and effort.51 Activism
may be followed by social opprobrium. The time involved in com-
plaining can also be quite substantial; we leave it as an exercise for the
reader to estimate how long it would take to resolve a technical issue
with her cable company. The benefits of taking such action, by con-
trast, are uncertain: The seller may not yield to the consumer’s de-
mands, and even if the seller relents, the value of such remedial action
may not be significant. On net, the value of an uncertain replacement
of a faulty product can be easily outweighed by the certain investment
of time and effort.52 In a study of 149 dissatisfied consumers who did
not complain, the consumers explained their inaction as follows:
shortage of time was the leading reason (~21%), followed by lack of
responsiveness on part of sellers (~20%), and consumer personality
factors (some simply do not like the confrontation involved in com-
plaining) (~17%).53
Nudniks, in contrast, are active. They become “crusading consum-
ers” whenever their transactional expectations are defied, even when
most others would have decided that the costs of fighting sellers are
not worth the expected benefits. It may be time to interject with a
note on terminology. While consumer passivism is often labelled “ra-
tional,” we wish to avoid labeling consumer complaining as “irra-
tional.” A nudnik who serially complains does not necessarily act
irrationally. A consumer suing for a small overcharge can be cast as
irrational if one reduces rationality to the pursuit of material cost-ben-
48. Beware of Dissatisfied Consumers: They Like to Blab, MARKETING (Mar. 8, 2006), https://
knowledge.wharton.upenn.edu/article/beware-of-dissatisfied-consumers-they-like-to-blab/ (cit-
ing WHARTON BUS. SCH. & VERDE GRP., RETAIL CUSTOMER DISSATISFACTION STUDY (2006)).
49. John Goodman, Basic Facts on Customer Complaint Behavior and the Impact of Service on
the Bottom Line, 8 COMPETITIVE ADVANTAGE 1, 1–5 (1999).
50. John W. Huppertz, Firms’ Complaint Handling Policies and Consumer Complaint Voicing,
24 J. CONSUMER MARKETING. 428, 428 (2007).
51. See generally Robin M. Kowalski, Complaints and Complaining: Functions, Antecedents,
and Consequences, 119 PSYCHOL. BULL. 179 (1996) (examining how different personality types
experience the lodging of complaints); Marsha L. Richins, A Multivariate Analysis of Responses
to Dissatisfaction, 15 J. ACAD. MARKETING SCI. 24 (1987).
52. For a review of the marketing literature on the costs and benefits of complaints, see Hup-
pertz, supra note 50, at 429–30.
53. Voorhees et al., supra note 46, at 519.
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 15 21-APR-20 11:48
2020] CONSUMER ACTIVISM 247
efit analysis. But if one sees the nudnik’s preferences as consisting of
broader concerns, from spite to altruism to all other human motiva-
tions that are in-between, then labeling his actions irrational no longer
fits. For our purposes, it is immaterial what label one gives to their
behavior, so long as it is clear that nudniks defy the standard account
of rational apathy.
Nudniks are not the only type of active consumers. “Shoppers” pre-
sent another category:54 those consumers that the informed minority
theory envisions, who shop around, read contracts, and compare
among products based on price, quality, and the terms of the con-
sumer agreement. A shopper, for example, will not fly with a certain
airline if she reads on the website that this airline is not willing to offer
compensation for delays; avoid a car dealership if the contract does
not provide warranties; or will not sign up for a credit card if the
agreement shows cash advance fees that are too high. In other words,
for the shopper, the terms of the contract are the product. Nudniks, in
contrast, are not necessarily committed to studiously comparing
among sellers. A nudnik can form her transactional expectations
based on the same sources that most other consumers use—negotia-
tions with the seller, representations, advertisements, market norms,
and so on.55
Shoppers exert pressure ex ante, before the purchase; nudniks exert
pressures ex post, after the purchase. If a seller includes an unfavora-
ble term in the fine print, such as denying refunds for defective prod-
ucts, this term may end up costing the seller consumers who are
shoppers, as they will switch to a more consumer-friendly competitor.
Nudniks (like passivists) may not be as sensitive to the inclusion of
such terms, if only because they may not read the fine print. On the
other hand, a seller who actually enforces the term—e.g., denying the
nudnik a refund for a broken printer—risks invoking the nudnik’s
wrath. Herein lies another distinction. The shoppers’ mode of action is
exit: they do not engage with sellers who offer inferior terms. By con-
trast, the nudnik’s mode of action is more elaborate, and consists of a
variety of voice strategies, both private and public.56
54. Economic theories of search behavior focus on shoppers. Sara Fisher Ellison, Price Search
and Obfuscation: An Overview of the Theory and Empirics, in HANDBOOKONTHE ECONOMICS
OF RETAILINGAND DISTRIBUTION287 (2016).
55. These broader transactional expectations may sometimes become part of the contract it-
self—through tools of interpretation that focus on oral representations, trade customs, past deal-
ings, etc. See supra note 8. But that is not always the case (think, for example, of the effect of the
parol evidence rule). See RESTATEMENT(SECOND) OF CONTRACTS §213; U.C.C. §1-303.
56. We elaborate on nudniks’ different modes of activism in Arbel & Shapira, Theory of the
Nudnik, supra note 9.
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 16 21-APR-20 11:48
248 DEPAUL LAW REVIEW [Vol.69:233
The last category of consumers is sophisticates. These are the con-
sumers who take advantage of the most favorable terms of the trans-
action, strategize their use of the product to derive the most value,
and file lawsuits when they expect a large judgment.57 The sophisti-
cates are typified by higher levels of literacy, a better understanding of
legal concepts, and easier access to legal representation. The sophisti-
cate is one to know which credit card to use in each transaction when
choosing among a dozen that are tightly packed in her wallet, when to
make a claim on her insurance policy, and how to maximize the value
of her miles.58 Sophisticates are not just better at consuming, they are
also better at identifying profitable lawsuits. The sophisticate can tell
when her contractual rights are violated, the value of filing a lawsuit
or a complaint, and the most effective route to use.
The sophisticates are like nudniks, in a sense that they are “private
enforcers” in the market. But, unlike nudniks, sophisticates enforce
only when a cold cost-benefit calculation says it pays to do so. In their
enforcement actions, sophisticates are in a sense like bounty hunters.
They seek personal profit from vindicating their legal rights. A classic
example is the serial class action plaintiff who actively seeks wrongs so
as to file profitable lawsuits. In China, consumer protection laws gave
birth to “counterfeit hunters”, who purchase counterfeits just so that
they can file a complaint to the regulator and collect a reward under
the rules there.59 Indeed, policy makers sometimes leverage bounty
hunters to help with private enforcement of issues of public import—
think, for example, on the use of private rewards in qui tam claims.60
Sophisticates and nudniks are therefore birds of a different feather.
Sophisticates will complain or sue only when they are within their le-
gal right to do so and when it pays to complain. Nudniks complain
because it is in their blood. They do not like to be treated unfairly.
They believe that sellers should keep promises and will not stop until
the issue is rectified, regardless of a cold cost-benefit analysis of what
it would take to keep fighting. Sophisticates focus on the product or
57. A famous example is John Leonard, a consumer of Pepsi who found an apparent loophole
in their promotion offering—a harrier jet estimated at $22 million for anyone who could collect
700,000 points. As points were purchasable at a rate of cent/point, that meant that an investment
of $700,000 would net a profit of $21.3 million. Leonard v. Pepsico, Inc., 88 F. Supp. 2d 116,
(S.D.N.Y. 1999), aff’d, 210 F.3d 88 (2d Cir. 2000).
58. For example, the authors have, on average, 20 credit cards. One of the authors only has
one card.
59. Sui-Lee Wee, Though Awash in Fakes, China Rethinks Counterfeit Hunters, N.Y. TIMES,
Nov. 30, 2016.
60. See David Freeman Engstrom, Public Regulation of Private Enforcement: Empirical Anal-
ysis of DOJ Oversight of Qui Tam Litigation Under the False Claims Act, 107 NW. L. REV. 1689,
1690 (2013) (noting the shift “toward private lawsuits as a regulatory tools”).
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 17 21-APR-20 11:48
2020] CONSUMER ACTIVISM 249
contract: They will use the product optimally, as in using balance
transfers to roll payments and avoid interest payments, and will file
class actions when they think a favorable settlement or judgment is
likely. The nudniks, on the other hand, are more likely to take action
vis-a`-vis the seller; when they feel the seller mistreated them, they will
complain, bring a lawsuit, report to the regulator, and so on.
Another way to put it: Sophisticates operate based on the fine print,
while nudniks do just fine without the print. The sophisticates act
upon contractual rights they know they have; nudniks upon broad
transactional expectations.61 To illustrate, think about a restaurant
that notes in its terms and conditions a disclaimer that the chef holds
full discretion over how to cook the meat and will not replace a dish.
If the restaurant serves an overcooked steak, the irate nudnik will de-
mand to have it replaced and it will comfort her little even if the res-
taurant’s general counsel reads her the contractual disclaimer.
Meanwhile, the sophisticate will sit and angrily chew, recalling the ex-
clusion in Article 7, subsection (3).
An important distinction for our purposes is that the nudnik’s ac-
tions are more likely to generate positive spillovers than the sophisti-
cate’s actions. When sophisticates use products optimally, they tend to
extract private benefits. In fact, sophistication often leads to negative
spillovers and cross-subsidies from less sophisticated consumers.62
Even when sophisticates file lawsuits, or even class-actions, the private
profit motive suggests that they internalize a larger share of the value
of their activities. An even more important distinction stems from the
sophisticates’ dependence on the exogenous cost-benefit of private
enforcement. To illustrate, consider the example of the class action
plaintiff. Recent developments in American law, and specifically the
wave of mandatory arbitration clauses, severely limit the scope of col-
lective action, making it extremely difficult to benefit from bringing
them.63 In this environment, the bounty hunters who face an insur-
61. The idea that consumers make choices based on background knowledge, not deriving di-
rectly from the contract, was recently advanced by Ayres and Schwartz: “Consumers also learn
about the deals they make from visiting firms, their experience with similar deals, discussions
with friends, their observation of other consumers’ purchasing choices, and reading consumer
reports.” Ayres & Schwartz, supra note 35, at 550–51.
62. See, e.g., Xavier Gabaix & David Laibson, Shrouded Attributes, Consumer Myopia, and
Information Suppression in Competitive Markets, 121 Q. J. ECON. 505, 519 (2006) (analyzing
cross-subsidies between sophisticated and unsophisticated consumers).
63. See AT&T Mobility v. Concepcion, 563 U.S. 333, 352 (2011) (holding that the Federal
Arbitration Act preempts state power to limit class-action waivers); Zachary D. Clopton, Class
Actions and Executive Power, 92 N.Y.U. L. REV. 878, 880 (2017) (“The private-enforcement
class action faces strong ‘headwinds’ in the form of class certification, subject-matter jurisdiction,
and arbitration.”); Roy Shapira, Mandatory Arbitration and the Market for Reputation, 99 B.U.
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 18 21-APR-20 11:48
250 DEPAUL LAW REVIEW [Vol.69:233
mountable hurdle in collecting their bounty will stop hunting. By con-
trast, nudniks may still be irritated enough and have a strong sense of
commitment to make things right that they will file an individual law-
suit, attempt to enlist the help of journalists or report to the regulator.
In such scenarios, the positive spillovers from nudnik actions are even
more impactful.
To be sure, the distinction between sophisticates and nudniks is not
always clear. Consider for example the case of “gripe sites.”64 When
the Partingtons ran into issues with their contractor, Stantons and
Sons Contracting, they opened a blog hosted under danstantonissue
.com, where they detailed their negative experiences.65 Similarly, a
disappointed customer who bought an improperly installed fence from
Lowes opened the blog Lowes-Sucks.com.66 In both examples, the
gripe sites share key features of nudnik-based activism: a small con-
sumer resisting seller misbehavior and instead of doing nothing about
it (as most consumers would), going to great lengths to fight the mis-
behavior seller publicly.67 Yet many other gripe sites exist,68 and some
of them are used to blackmail companies and do not represent au-
thentic consumer sentiments.69 The operator of the website may be
solely motivated by the desire to receive payment from the company
to take down the website. Operators of the latter type of gripe sites
can be perceived as “trolls”; or, in our framework, and given that the
profit-motive seemingly drives them, as cynical sophisticates.
B. Nudniks in Action: Motivating Examples
We illustrate nudniks’ modes of activism and the role they play in
consumer markets with a few motivating examples.70 The examples
L. REV. 873 (2019) (arguing that the rise of mandatory arbitration clauses dilutes the effective-
ness of market discipline).
64. See generally Rachael Braswell, Consumer Gripe Sites, Intellectual Property Law, and the
Use of Cease-and-Desist Letters to Chill Protected Speech on the Internet, 17 FORDHAM INTELL.
PROP. MEDIA & ENT. L.J. 1241 (2007).
65. Dan Stanton, Contractor Issue, BLOGSPOT, http://danstantonissue.blogspot.com/; see also
Bruce Mohl, Constructive Criticism, BOSTON GLOBE(Sept. 9, 2007), http://archive.boston.com/-
business/technology/articles/2007/09/09/constructive_criticism/.
66. Jacqui Cheng, Intellectual Property Laws Abused in Question to Shutdown Lowes-sucks
.com, ARS TECHNICA (Sept. 26, 2007). In many of the gripe site cases, the seller claims trade-
mark infringement, but such claims are rarely successful.
67. See Felix T. Wu, Collateral Censorship and the Limits of Intermediary Immunity, 87 NOTRE
DAME L. REV. 293, 306 (2011) (noting the public-service orientation of gripe sites).
68. See Peter Johnson, Can You Quote Donald Duck?: Intellectual Property in Cyberculture,
13 YALE J.L. & HUMAN. 451, 478 (2001).
69. Wu, supra note 67, at 304.
70. We discuss some of these examples in greater detail in Arbel & Shapira, Theory of the
Nudnik. We also include some graphics and other illustrations in the companion website.
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 19 21-APR-20 11:48
2020] CONSUMER ACTIVISM 251
we use here are admittedly representative of only a small segment of
nudnik activities in the real world: we use examples that become pub-
licly salient, while most of nudnik activities happens away from the
limelight. Think for example about the consumer who consistently
writes detailed negative Yelp reviews whenever the seller disappoints,
posting rants on products’ Facebook pages, submitting letters to the
Better Business Bureau or complaints to the regulator, and so on.71 In
most cases, we will not hear about this nudnik’s activities. Still, the
publicly-salient examples we use are helpful in a sense that they tease
out certain recurring themes in nudnik-based activism.
A classic example of a nudnik in action comes from Harvard Busi-
ness School Professor Ben Edelman. In 2014, Professor Edelman or-
dered takeout from a local restaurant, Sichuan Garden.72 When he
reviewed the check, he found that he was overcharged $1 for each of
the four items ordered.73 Edelman wrote a message on the restau-
rant’s website and followed up with a detailed email, noting the $4
overcharge and asking for a clarification.74 The owner responded by
noting that the website prices have “been out of date” for quite some
time relative to the restaurant menu.75 Edelman then sought a com-
pensation of $12 for the overcharge, citing the Massachusetts Unfair
and Deceptive Trade Practices Act, which permits treble damages in
certain cases.76 The owner refused, Edelman reported the incident to
the regulator and the parties continued corresponding until the story
leaked to the media. The public response was highly negative, but the
Yonathan Arbel, Theory of the Nudnik—Battle of the Forms, BLOG (Feb. 1, 2019), http://battleof
theforms.com/theory-of-the-nudnik/ [hereinafter Arbel, Battle of the Forms].
71. To clarify: the emphasis here is on “consistently.” The reader who wonders if she would be
considered a nudnik in our framework because she once posted a review on Yelp or TripAdvisor
can ask herself this simple question: how often do I fight sellers when I am dissatisfied with the
product or service. If I frequently take action when sellers disappoint, I am a nudnik (nothing
wrong with it!). If I once wrote a negative review because I was so upset and needed to air out
the frustration, chances are I am not really a serial consumer crusader.
72. See Hilary Sargent, Ben Edelman, Harvard Business Professor, Goes to War over $4 Worth
of Chinese Food, BOSTON.COM (Dec. 9, 2014), https://www.boston.com/culture/restaurants/2014/
12/09/ben-edelman-harvard-business-school-professor-goes-to-war-over-4-worth-of-chinese-
food.
73. Id.
74. Id.
75. Id. For some other examples on a companion website, see Arbel, Battle of the Forms,
supra note 70.
76. MASS. GEN. LAWS. ch. 93A, §11, declared unconstitutional by Rev-Lyn Contracting Co. v.
Patriot Marine, LLC, 760 F. Supp. 2d 162 (D. Mass. 2010).
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 20 21-APR-20 11:48
252 DEPAUL LAW REVIEW [Vol.69:233
negativity was directed at Edelman and not at the overcharging res-
taurant. His behavior was portrayed as petty and privileged.77
The overwhelmingly negative public response fails to appreciate the
essential public service Edelman provided—namely, deterring
overcharging. Anecdotally, when we called the restaurant, five years
after the story broke and replicated the original order, we found the
pricing to be accurate: The online menu matched the actual prices
down to the penny. Edelman’s insistence here may strike some as “ir-
rational,” as the opportunity cost of the time Edelman—a well-paid
Harvard Business Professor and a sought-after consultant—spent on
fighting the overcharge easily dwarfed the $12 he was ostensibly after.
Yet if it were not for people like Edelman who go through the trouble,
restaurants would have a much easier time systematically overcharg-
ing the rest of us.78
Another example came in 2011, when Molly Katchpole, a 22-year-
old customer of Bank of America noticed an upcoming change in the
bank’s fee structure.79 The bank was about to add a $5 charge to cer-
tain debit-card users, a right the bank had under its contract with its
customers.80 Molly not only noticed the upcoming change, she decided
to fight back against it.81 She started an online petition, where she
wrote: “[T]his change will hit low income customers the worst - in-
cluding people like me, a recent college graduate working two part-
time jobs . . . . At some point, we’ve got to say enough is enough.”82
To promote her petition, she used social media in combination with
traditional media—she reached out to an ABC reporter who filmed a
segment about her petition.83 Her call resonated with many others and
the petition quickly garnered over 300,000 signatures.84 Soon after,
the bank announced that it would abandon its plan to add this charge.
The Katchpole example illustrates how nudniks pay attention to as-
pects of the seller behavior that many do not see; how they become
active when they feel wronged even when most of us would remain
77. Elizabeth Barber, A Harvard Professor Launched an Epic Rant Over an Extra $4 on his
Chinese Takeout Bill, TIME (Dec. 10, 2014), http://time.com/3627282/harvard-professor-chinese-
takeout-ben-edelman/.
78. Telephone Conversation with Victoria Moffa, Research Assistant (Mar. 1, 2019).
79. Molly Katchpole, Tell Bank of America: No $5 Debit Card Fees, CHANGE.ORG (Nov. 1,
2011), https://www.change.org/p/tell-bank-of-america-no-5-debit-card-fees.
80. Id.
81. Id.
82. Id.
83. Matt Gutman & Susanna Kim, BofA Site Problems Persist; Customers Petition Company,
ABC NEWS (Oct. 5, 2011), https://abcnews.go.com/Business/bank-america-customers-launch-
petition-debit-card-fee/story?id=14665531.
84. See Katchpole, supra note 79.
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 21 21-APR-20 11:48
2020] CONSUMER ACTIVISM 253
passive; and how the nudnik’s activism can draw others, normally pas-
sive consumers, into action and bring change in market behavior.
Another example of a high impact strategy is that of the Canadian
country singer, Dave Carroll. On his flight with United Airlines, his
favorite 710 Taylor guitar was broken due to mishandling by the
ground crew.85 Upset with the company’s indifference, he posted a
song on YouTube called “United Breaks Guitars,” which quickly went
viral with over 19 million views.86 The story was soon reported by the
mass media as well.87 By one estimate, the incident led to a decline of
over ten percent in the company’s stock price.88 Consequently, United
backed down and offered monetary compensation.89
In 2017, British Airways lost Hasan Syed’s father’s luggage in an
international flight to Paris. Syed took his grievance to social media,
where he tweeted: “Don’t fly with @British_Airways. They can’t keep
track of your luggage.”90 This tweet had a twist. Syed paid $1,000 to
Twitter to have this tweet promoted to as many followers of British
Airways as possible.91 Over a short period of time, his tweet was seen
by over 70,000 users.92 His tweet received considerable media atten-
tion and was dubbed as the first instance of “complaintvertising.”93
Soon after his campaign, British Airways sent him a special apologetic
message, located the bag, hand-delivered it to his dad in Paris, and
issued a public apology.94
85. UNITED BREAKS GUITARS, https://www.davecarrollmusic.com/songwriting/united-breaks-
guitars/?v=7516fd43adaa (last visited Nov. 27, 2019).
86. Sonsofmaxwell, United Breaks Guitars, YOUTUBE (July 6, 2009), https://www.youtube
.com/-watch?v=5YGc4zOqozo.
87. DAVID DUNNE, ROTMAN, UNITED BREAKS GUITARS CASE, https://3gz8rn1ntxn33t9p221
v8-mlgtsq-wpengine.netdna-ssl.com/wp-content/uploads/United-Breaks-Guitars-Case-Jan-11-10-
21.pdf.
88. Eddie Wrenn, The Sweet Music of Revenge: Singer Pens YouTube Hit After United Air-
lines Breaks His Guitar . . . and Shares Plunge 10%, DAILY MAIL (July 24, 2009), https://www
.dail-ymail.co.uk/news/article-1201671/Singer-Dave-Carroll-pens-YouTube-hit-United-Airlines-
breaks-guitar—shares-plunge-10.html.
89. Did Dave Carroll Lose United Airlines 4180m?, THE ECONOMIST (July 24, 2009), https://
www.economist.com/gulliver/2009/07/24/did-dave-carroll-lose-united-airlines-180m.
90. @HVSVN, TWITTER (Sept. 3, 2013, 3:46 PM), https://twitter.com/HVSVN/status/375026
96-3347304449.
91. Id.; @Kforesti, TWITTER (Sept. 4, 2013, 9:08 AM), https://twitter.com-/Kforesti/status/3752
89284276006912.
92. @Kforesti, supra note 91.
93. ANGRY TRAVELER PAYS BIG BUCKS FOR TWEETS, CNN MONEY, https://money.cnn.com/
vid-eo/news/2013/09/04/n-british-air-twitter-war-mclaughlin.cnnmoney/index.html [hereinafter
CNN MONEY]; Jason King, Complainvertising: Word of Mouth’s Evil Twin, HUFFPOST (Oct. 23,
2013, 1:14 PM), https://www.huffingtonpost.com/jasonking/complainvertising-word-of_b_4143-
073.html (last updated Dec. 6, 2017).
94. CNN MONEY, supra note 93; King, supra note 93.
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 22 21-APR-20 11:48
254 DEPAUL LAW REVIEW [Vol.69:233
There are various indications that these acts of consumer activism
have not only improved things for the specific complaining customer,
but also led to a broader change in sellers’ policies. When Katchpole’s
petition gained traction, Bank of America quickly retracted its policy.
After the Carroll video broke out, United Airlines promised to
change its customer service policy accordingly and asked permission
to use the video in its internal training seminars (one trusts as an ex-
emplar of what not to do).95 The ripple effects, interestingly, were felt
much more broadly, and Carroll reports that several other companies
licensed his video as well for their internal training seminars.96
The nudnik’s activism is often part of repeat behavior. Recall Pro-
fessor Edelman’s example;97 it turns out that Professor Edelman had
previously complained about various other restaurants around Bos-
ton.98 When his discount coupon was not accepted at a sushi restau-
rant, he threatened that he would write to the Boston Licensing Board
to have their business license revoked.99 Katchpole’s campaign against
Bank of Americas was not her last; she also fought against a planned
“convenience fee” by Verizon Wireless.100 The company retracted its
policy in less than twelve hours.101
The public reaction to nudniks is often negative. When Edelman
complained about Sichuan Garden over $12, many mocked him for
being petty, privileged, and ruthless.102 Even the marketing literature
treats nudniks quite unfavorably.103 An introductory textbook in mar-
keting, for example, calls them “terrorists,” that “companies must
take measures to get rid of.”104
95. Broken Guitar Song Gets Airline’s Attention, CBC NEWS (July 8, 2009, 3:00 PM), https://
www.cbc.ca/news/entertainment/broken-guitar-song-gets-airline-s-attention-1.802741 (last up-
dated July 8, 2009).
96. DAVE CARROLL, UNITED BREAKS GUITARS: THE POWEROF ONE VOICEINTHE AGEOF
SOCIAL MEDIA 81 (2013).
97. See supra Introduction.
98. Hilary Sargent, There’s More: Edelman Did This Before, And Worse, BOSTON.COM (Dec.
10, 2014), https://www.boston.com/culture/restaurants/2014/12/10/theres-more-edelman-did-this-
before-and-worse.
99. Id.
100. Minda Zetlin, Meet Fee-Fighting Vigilante Molly Katchpole, DEBTHELPER.COM (Jan. 19,
2012), https://www.debthelper.com/blog/2012/01/meet-fee-fighting-vigilante-molly-katchpole/.
101. Id.
102. Nathan J. Robinson, Stop Eviscerating the Harvard Professor Who Threatened to Sue a
Chinese Restaurant Over $4. He Has a Point, The NEW REPUBLIC (Dec. 13, 2014), https://newer-
public.com/article/120558/ben-edelman-harvard-prof-angry-over-4-overcharge-has-point (“By
now even Ben Edelman thinks Ben Edelman is fairly despicable . . . . The consensus is that he’s a
cheap, entitled bully and that the immigrant restaurant owner is a hapless victim.”).
103. See, e.g., Jagdip Singh, A Typology of Consumer Dissatisfaction Response Styles, 66 J.
RETAILING 57 (1990).
104. SCHIFFMAN & WISENBLIT, supra note 45, at 44.
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 23 21-APR-20 11:48
2020] CONSUMER ACTIVISM 255
In summary, we see in these examples how nudniks have an idio-
syncratic cost structure, which leads them to take action—sometimes
quite radical—in instances where most consumers would remain
quiet. Potentially attributable to their unique cost structure, public re-
ception of nudniks is not always favorable. The public mocks them
partly because they are different. Yet it is precisely because they are
different that nudniks can solve the collective action problems that
plague all other consumers. Nudniks take action when most of us do
not. And their action can lead to substantial changes in consumer
markets—even when nudniks pursue the enforcement of rights not
formally grounded in the contract.
C. Why do Nudniks Prevail where the Informed Minority Fails?
The theory of the crusading minority sidesteps the issues that
plague the informed minority theory along three key dimensions: (1)
nudnik-based activism is not predicated on financial cost-benefit anal-
ysis, and so is less susceptible to changes in the market environment;
(2) nudnik-based activism is not predicated on reading and compre-
hending contracts; and, (3) nudnik-based activism is not predicated on
the existence of a critical mass of similarly minded activists.
The first factor concerns the cost of being an active consumer. Both
shoppers and sophisticates are active because it pays for them to do
so. They shop around, read contracts, and examine reviews because it
allows them to find the best deals on the market, use the product opti-
mally, or find profitable lawsuit opportunities. But this makes their
activism contingent and unreliable when the costs of becoming active
rise. When the costs of becoming engaged increase, sophisticates and
shoppers stop being involved. They will read, negotiate, and sue less.
Nudniks, by contrast, operate mostly based on motivations that are,
in a sense, internal. They possess certain idiosyncratic personality
traits and beliefs that compel them to sink their teeth in and not let go
when they feel wronged, regardless of the financial cost-benefit analy-
sis. Elinor Ostrom argued that group norms are often enforced by a
subset of individuals within the group who are “willing punishers”:
those for whom fighting wrongs comes naturally (even at a personal
cost—think about approaching an able-bodied person parking at a
disabled parking spot to scold him).105 While some may view this as-
pect of nudnik behavior as irrational or impetuous, for the nudnik it is
the right thing to do—you do not let people get away with violating
105. Elinor Ostrom, Collective Action and the Evolution of Social Norms, 14 J. ECON. PERSP.
137, 142 (2000).
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 24 21-APR-20 11:48
256 DEPAUL LAW REVIEW [Vol.69:233
their promises. To be sure, there are limits to nudnik activism too.
Presumably, when the costs of complaining, enlisting the help of the
media, or filing a lawsuit are especially high, they will deter even the
nudniks (or at least some of them) from voicing their concerns pub-
licly. Our point here is therefore not absolute but rather relative: An
increase in the costs of becoming involved is more likely to deter
other types of consumers than it is to deter nudniks.
This difference in responding to the costs of activism makes nudniks
a better fit as active players in today’s consumer markets, relative to
shoppers, sophisticates, and obviously passivists. In today’s markets,
disclosures and boilerplate terms are copied and reproduced at almost
zero marginal cost, leading to an explosion of contractual data.106 The
increase in the length and complexity of contracts will most directly
impact shoppers, but less so the nudnik. The limitations on class ac-
tions and the trend towards individual arbitration pose the greatest
risk for legally active sophisticates, but not necessarily nudniks.107
Even companies’ increasing efforts to limit complaints have a limited
impact on nudniks, who tend to find a way.108
The same idiosyncratic motivations lead nudniks to produce reputa-
tional information, as in posting detailed negative reviews online or
sharing stories with other consumers, even when the costs of such ac-
tivities dwarf any direct benefit. Note that most users of online plat-
forms consult reviews but do not post them. The operation of these
reputation markets is predicated on the few who do invest in posting
informative reviews, even when it does not pay to do so. In other
words, the fact that nudniks are driven by internal motivations and
not external cost-benefit analysis helps to solve not just the problems
that plague the informed minority theory, but also the problems that
plague the reputational discipline theory.
The second key distinction is that nudniks can affect change even
without thoroughly reading and comprehending contracts. Nudniks
frequently assert claims based on broader transactional expectations,
that is, what rights they believe they should have, regardless of what
the contract stipulates. As legal scholars have started to acknowledge,
consumers often form transactional expectations that are based not
106. To reiterate our point from Part I.B: Recent market trends, concerning the rise in con-
sumer disclosure and the explosion of online contracts, privacy policies, and End User License
Agreements all put increasing demands on the already-overwhelmed minority of readers. BEN-
SHAHAR & SCHNEIDER, supra note 32, at 94–101.
107. See supra note 63 and accompanying text. R
108. Amy J. Schmitz, Remedy Realities in Business-to-Consumer Contracting, 58 ARIZ. L.
REV. 213, 233–38 (2016) (detailing companies’ efforts to make it increasingly harder for consum-
ers to complain).
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 25 21-APR-20 11:48
2020] CONSUMER ACTIVISM 257
necessarily on the specific contract and its fine print, but rather on
previous transactions, a general sense of fairness, and market
norms.109 When there is a mismatch between consumer expectations
and the contract, most consumers will not fight the company. Nudniks
will. When a cable company fails to arrive on time, the nudnik may
post a negative review online, even if the fine print actually provides
the company the right to delay appointments with no prior notice.110
Ironically, the reading of the contract might lead sophisticates to be
less active, as they will find the cable company’s actions permissible
under the contract.111
Reacting to transactional expectations, rather than the language of
the contract, makes nudnik-based activism more relevant in today’s
consumer markets relative to informed minority-based activity. While
the volume and complexity of contractual information has increased,
making it less likely that consumers will read and comprehend their
contractual rights, consumers have been developing an increasingly
richer set of transactional expectations. Consumers nowadays experi-
ence more market interactions and are exposed to more merchants
and dealing styles (think online shopping comparisons). This develop-
ment broadens and sharpens their ability to compare products and
sense what a viable market norm looks like. Indeed, one notable cus-
tomer service report stated that “the digitally empowered customer
[has led to] customer expectations soaring” and “54% of respondents
say they have higher expectations for customer service today than
they had one year ago.”112 This is to say that in today’s consumer mar-
kets, consumers read contracts less but expect from the seller more. In
such an environment, activism based on breached expectations after
109. Ayres & Schwartz, supra note 35.
110. As others have noted, such claims can be quite effective, even if not grounded in the four
corners of the contract. See, e.g., Lucian A. Bebchuk & Richard A. Posner, One-Sided Contracts
in Competitive Consumer Markets, 104 MICH. L. REV. 827, 830 (2006) (arguing that “reputa-
tional considerations” may “induce the seller to treat the buyer fairly even when such treatment
is not contractually required”); Jason Scott Johnston, The Return of the Bargain: An Economic
Theory of How Standard Form Contracts Enable Cooperative Negotiation Between Businesses
and Consumers, 104 MICH. L. REV. 857, 858 (2006) (“In practice, acting through its agents, a firm
will often provide benefits to consumers who complain beyond those that its standard form obli-
gates it to provide . . . .”). See also Clayton P. Gillette, Rolling Contracts as an Agency Problem,
2004 WIS. L. REV. 679 (2004).
111. The problem is especially acute when contract terms are misleading and unenforceable,
yet consumers tend to view them as binding. See Furth-Matzkin, supra note 21 (finding the com-
mon inclusion of unenforceable terms in residential agreements).
112. 2017 STATE OF GLOBAL CUSTOMER SERVICE REPORT, MICROSOFT (2017), https://
info.micros-oft.com/rs/157-GQE-382/images/EN-CNTNT-Report-DynService-2017-global-state-
customer-service.pdf.
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 26 21-APR-20 11:48
258 DEPAUL LAW REVIEW [Vol.69:233
the fact is more relevant than activism based on reading the terms of
contract in advance.
Finally, a critical mass of nudniks is not necessarily needed to affect
change; a single nudnik may be enough. For the informed minority
theory, size matters: If the subset of readers is not sufficiently large,
sellers would not have sufficient incentives to adapt the firm’s offering
to win over these comparison shoppers. Firms would simply ignore
this tiny segment of the market.
Nudniks work differently. They put reputational and legal sanctions
in motion. For this reason, their power comes not necessarily from
numbers. One Ben Edelman was enough to stop the restaurant’s prac-
tice of overcharging. A single Dave Caroll can create a public rela-
tions crisis of large proportions. Just one lawsuit of an aggrieved Bank
of America customer resulted in a multi-million-dollar award in puni-
tive damages.113 To be sure, a single nudnik would not always be
enough—indeed, sometimes even many nudniks’ complaints will not
move the seller. The point is that under certain conditions, a single
nudnik’s fight can draw the attention of many others—a fact that is
appreciated by crisis management firms.
Recent changes in consumer markets have therefore made nudnik-
based activism relatively more impactful. The increase in contract
length and complexity and the rise of class action waivers have made
non-nudnik activists (and the informed minority theory) less relevant.
At the same time, the internet and in particular social media have
increased the ability of a single nudnik’s fight to bring a seller to its
knees.114 Posting a negative review online increases the dissemination
potential, the permanence, and the ease of accessibility of the damn-
ing information about the seller. Everyone searching for that seller in
the future can run into the nudnik’s detailed concern with the seller’s
underperformance. In other words, changes in the information envi-
ronment have boosted the nudnik’s signal and ability to shame firms
into meeting market norms.
113. See, e.g., Sundquist v. Bank of Am., N.A., 566 B.R. 563, 620 (Bankr. E.D. Cal. 2017), va-
cated in part sub nom. In re Sundquist, 580 B.R. 536, 556 (Bankr. E.D. Cal. 2018) (ordering $45
million in punitive damages against bank). Punitive damages can lure consumers into action—
especially sophisticates—but the point is that for nudniks, the financial payment is not the pri-
mary motive.
114. See generally Matthew S. O’Hern & Lynn R. Kahle, The Empowered Customer: User-
Generated Content and the Future of Marketing, 18 GLOBAL ECON. & MGMT. REV. 21 (2013)
(arguing that the user-generated content on social media “represents a profound shift of power
from firms to consumers”).
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 27 21-APR-20 11:48
2020] CONSUMER ACTIVISM 259
D. Why and How Sellers Accommodate Nudniks’ Concerns
To further underscore how nudniks bring about market-wide
changes in behavior, let us explore the question of how and why sell-
ers respond to nudniks. When nudniks have legally cognizable claims,
the mechanism is fairly straightforward. The pressure comes from
nudniks’ complaints to the regulator, the filing of lawsuits, or the re-
fusal to pay until the issue is settled. All of these actions impose direct
costs on firms, and the threat of these costs can deter seller misbehav-
ior ex ante.
Less intuitively, sellers respond to nudniks even when the latter as-
sert claims based on transactional expectations that are not grounded
in the four corners of the contract.115 While firms are not legally obli-
gated to appease the nudnik in such cases, they often have reputa-
tional incentives to do so. A nudnik that posts a negative review
online, or talks with a reporter or a consumer organization may cause
more significant damage to the firm than any single lawsuit.116 Recall
how Molly Katchpole fought Bank of America for something that was
within the bank’s contractual rights—changing its fee structure.117 An-
other example is the 2017 incident whereby United Airlines evicted a
paying passenger from the flight to accommodate another passen-
ger.118 Even though the airline company’s contract with the passenger
stipulated that it can de-board him, treating a consumer that way
seemed, in the eyes of United’s various stakeholders, as uncalled for,
unfair, and bad business.119 The incident led to a swift and significant
115. For a broad and attentive treatment of these issues, see also Shmuel I. Becher & Tal Z.
Zarsky, Minding the Gap, 51 CONN. L. REV. 69 (2019).
116. Even complaints to the regulator may be based on violation of consumer expectations;
whether or not the regulator would respond to such complaints is a different matter.
117. See supra notes 61–64 and accompanying text.
118. Christina Zdanowicz & Emanuella Grinberg, Passenger Dragged Off Overbooked United
Flight, CNN TRAVEL (Apr. 10, 2018, 8:13 AM), https://www.cnn.com/2017/04/10/travel/passen
ger-removed-united-flight-trnd/index.html; Erin McCann, United’s Apologies: A Timeline, N.Y.
TIMES (Apr. 14, 2017), https://www.nytimes.com/2017/04/14/business/united-airlines-passenger-
doctor.html (recounting the various apologies issues. At first, the CEO noted that the “re-ac-
commodat[ion]” of the passenger was according to “established procedures”—a claim that was
not repeated in future apologies).
119. CONTRACT OF CARRIAGE DOCUMENT r. 25, UNITED, https://www.united.com/ual/en/us/
fly/co-ntract-of-carriage.html (last updated Nov. 7, 2019). For an example of the public response,
see Alex Abad-Santos, Why United Airlines Can Get Away With Treating Its Customers Like
Garbage, VOX.COM (Apr. 11, 2017, 12:30 PM), https://www.vox.com/culture/2017/4/11/15246632/
united-airlines-drag-man-off-plane (“Even if United was well within its rights . . . in bumping the
Chicago passenger from his flight, most people can agree that its public explanation in both cases
wasn’t a good look.”). See also David A. Hoffman, Relational Contracts of Adhesion, 85 U. CHI.
L. REV. 1395, 1401–02 (2018) (noting the potential reputational fallout of enforcing certain terms
against consumers).
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 28 21-APR-20 11:48
260 DEPAUL LAW REVIEW [Vol.69:233
decline in passengers’ willingness to fly United.120 In general, the
reputational effects of asserting contractual rights narrowly can be
devastating from the firm’s perspective. Such potential reputational
effects may be enough to drive firms to conform to consumers’ supra-
contractual expectations.121
The mere background threat of nudnik activity—the “shadow of the
nudnik”—can be enough to affect change. A seller that deals with a
large number of consumers and cannot identify in advance who will be
a nudnik,122 faces a choice. The seller could maintain its level of con-
tractual performance and address nudniks once they reveal them-
selves as such, hoping to be able to buy them off before they draw
negative attention to the firm’s behavior. Alternatively, the seller
could raise its level of contractual performance across the board, so
that not even nudniks would have something to complain about.
At least until recently, there were good reasons for firms to choose
the latter option and improve their contractual performance across all
consumers. The alternative—waiting until you realize you are dealing
with a nudnik—is simply too risky.123 If the firm only reaches a
nudnik after she aired her grievances online, the negative review al-
ready has a life of its own and the firm’s potential future consumers
may read it and decide to go elsewhere.124 Similarly, if a firm only
approaches a nudnik after she filed a lawsuit, then even if the firm
paid her enough to get her to settle, the mere filing may leave enough
120. A poll of nearly 2,000 individuals reported a sharp decline in willingness to fly with
United Airlines following the incident. Kevin Quealy, How Much Would You Put Up With to
Avoid United Airlines?, N.Y. TIMES(Apr. 17, 2017), https://www.nytimes.com/2017/04/17/upshot/
how-much-would-people-put-up-with-to-avoid-united-airlines.html.
121. The incident led to a “marked decrease” in the rate of bumping passengers, from 0.62 per
10,000 to 0.44, the lowest rate in decades. U.S. DEP’T OF TRANSP., AIRLINE BUMPING RATE
LOWEST IN DECADES, TRANSPORTATION.GOV (Sept. 7, 2017), https://www.transportation.gov-/
briefing-room/dot6417.
122. We relax this assumption in Arbel & Shapira, showing that when sellers can identify who
is a nudnik before the consumer even makes a purchase (assisted by new predictive analytics
tools), the prospect of market discipline takes a hit. Theory of the Nudnik, supra note 9.
123. Our analysis here diverges from that of Professor Amy Schmitz in one important aspect.
Professor Schmitz highlights the inefficiencies and unfairness (in terms of cross-subsidies) that
follows when firms discriminate in favor of active consumers and against passivists. See generally
Amy J. Schmitz, Access to Consumer Remedies in the Squeaky Wheel System, 39 PEPP. L. REV.
279 (2012) [Schmitz, Access to Consumer Remedies]. We believe that sellers’ ability to discrimi-
nate after the fact is limited and may come too late once one accounts for the way reputational
channels operate.
124. The existence might also alert disgruntled consumers to the existence of a systemic issue,
encouraging them to complain as well. The legitimizing power of a trailblazing complaint is seen
most powerfully in the context of complaints of sexual misconduct.
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 29 21-APR-20 11:48
2020] CONSUMER ACTIVISM 261
breadcrumbs for others to pick up on.125 As one of us documents else-
where, journalists often use lawsuit filings as a source for follow-up
investigative projects that hold defendants to account.126 Moreover,
paying the nudnik off to silence her tends to be more difficult after the
nudnik had already filed a lawsuit.127
*
Focusing on the small subset of crusading consumers holds the
promise of understanding how market governance works in an age
when only few take the trouble to read the fine print. Sellers have
strong incentives to respond to nudniks for reasons that are both legal
and reputational. And if sellers are unable to identify nudniks in ad-
vance, nudniks’ activism would prompt sellers to improve their level
of service for all consumers. To be sure, nudniks are not omnipotent,
and we cannot expect them to fix all (or even most) ills in consumer
markets. Our claim here is more modest: Nudniks represent an impor-
tant force in the marketplace that has remained understudied and
should receive more attention from legal scholars and policymakers.
We now turn to consider the limits to nudnik-based activism, and how
the magnitude of their social contribution varies greatly across
markets.
III. LIMITATIONS OF NUDNIK-BASED CONSUMER GOVERNANCE
Our argument thus far has been straightforward: Instead of focusing
on a minority of consumers who read and negotiate terms before the
purchase, legal scholars and policymakers should pay greater atten-
tion to a minority of crusading consumers who take action when sell-
ers do not meet their transactional expectations. This Part introduces
nuance. Nudnik-based behavior is not always socially beneficial. Some
nudniks complain frivolously. Their transactional expectations may be
unrealistic and untethered. Some complaints only reflect the nudniks’
idiosyncratic preferences. Such behavior does not contribute to effec-
tive seller deterrence, but rather imposes costs and hurts the function-
ing of consumer markets. Any theory of nudnik-based activism should
therefore attempt to identify the cross-sectional variation—the condi-
125. See generally Roy Shapira, Law as Source: How the Legal System Facilitates Investigative
Journalism, 37 YALE L. POL’Y REV. 153 (2018) (detailing how legal breadcrumbs lead to investi-
gative reports).
126. Id.
127. This is the result of, among others, the larger leverage plaintiffs have after sinking some
of the costs of litigation. See generally Lucian A. Bebchuk & Alon Klement, Negative Expected-
Value Suits, inELGAR ENCYCLOPEDIAOF LAWAND ECONOMICS (2d ed. 2009).
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 30 21-APR-20 11:48
262 DEPAUL LAW REVIEW [Vol.69:233
tions under which nudniks are more or less likely to bring positive
change in seller behavior.
One concern with nudniks is distributional. As the old adage goes,
“[t]he squeaky wheel gets the grease.”128 Applied here, the concern is
that sellers will reward only the noisy consumers. More accurately, the
concern is twofold: (1) that sellers will reward noisy consumers at the
expense of other consumers and, importantly, (2) that the noisy con-
sumers tend to be the “haves” rather than the “have-nots.”129 In other
words, to assess the social value of nudnik-based activism we have to
ask who gets what and at whose expense. But such distributional con-
cerns about nudnik-based activism should not be overstated, for the
following three factors.
First, to the extent these concerns are valid, they are not unique to
nudnik-based activism but rather apply equally, or even more force-
fully, to other forms of consumer activism. Consider the sophisticates,
who pursue legal actions for personal gain or find ways to optimally
use the products, or the shoppers who shop around for the best deals.
Consumers in both of these categories are likely to be more privileged
or sophisticated (by definition) than the rest of the consumer body
and extract private benefits—sometimes at the expense of others.
Second and more specifically to nudniks, it is hardly evident that
nudniks are overly privileged. While future empirical research is
needed on this issue, we can already point out studies in the consumer
complaint behavior literature, showing that there actually exists great
heterogeneity among serial complainers: nudniks cut across cultural,
economic, and social dimensions.130
Finally, even if we assume that nudniks are the privileged ones,
does it matter? If nudniks squeak loud enough, they can affect mar-
ket-wide changes that benefit the entire consumer body, including the
relatively less well off. The squeaky wheel can alert us all to the possi-
bility that there is a problem with a given seller or a product. And
providing this service may sometimes actually require privilege. To il-
128. Schmitz, Access to Consumer Remedies, supra note 123, at 280.
129. Id. at 290. See also Lauren E. Willis, Performance-Based Consumer Law, 82 U. CHI. L
REV. 1309, 1326 (2015); Arthur Best & Alan R. Andreasen, Consumer Response to Unsatisfac-
tory Purchases: A Survey of Perceiving Defects, Voicing Complaints, and Obtaining Redress, 11
LAW & SOC’Y REV. 701, 723 (1977) (finding that complaints underrepresent poor consumers and
racial minorities).
130. Some of the empirical research does suggest that, on average, serial complainers tend to
be more educated and affluent. But much more empirical research is needed before we can
make the leap to argue against nudniks in the name of distributional concerns. See, e.g., Michelle
A. Morganosky & Hilda M. Buckley, Complaint Behavior: Analysis by Demographics, Lifestyle,
and Consumer Values, 14 NA – ADVANCESIN CONSUMER RES. 223–26 (1987).
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 31 21-APR-20 11:48
2020] CONSUMER ACTIVISM 263
lustrate, recall the Ben Edelman example of going on a crusade
against an overcharge of $4. Perhaps it can be said that only someone
with sufficient financial security can spend so much time correspond-
ing over a $4 overcharge as a matter of principle.131
Another oft-mentioned concern with nudniks is their motivations.
The public, and sometimes even judges, tend to focus on nudniks’ sup-
posed motivations and cast them negatively as vengeful and petty.
Granted, some consumer crusaders are motivated by spite. But we do
not view spite as a strong argument against nudnik-behavior. In fact,
spite may be a virtue in this context.
Spite drives nudniks to be “willing punishers”—those who contrib-
ute to solve consumer collective action problems.132 While all consum-
ers may suffer from late deliveries or missed appointments by a seller,
each consumer sees the costs of a public fight with the seller as dwarf-
ing whatever benefit she might get from a refund. Spite allows
nudniks to transcend such cold cost-benefit calculation and motivate
public action that would provide a valuable service to all other con-
sumers. In this sense, punishing a misbehaving seller out of spite could
be considered “other-regarding” by the nudnik.133 In fact, there is a
certain irony in treating spiteful nudniks negatively: Why not cast the
majority’s silence in the face of seller violations as a selfish, disinter-
ested attitude towards their fellow consumers? Consider for example
how most economic models of market discipline invoke the rhetoric of
“consumer sharing” (with its positive connotations), to denote in-
stances where one consumer learns about a product failure, and im-
mediately shares the information with her fellow consumers. In
reality, the overwhelming majority of consumers do not share. They
do nothing. If spite or pettiness or obsession makes nudniks de facto
share information with others, then so be it. The only concern should
be with whether the information nudniks generate is valuable to
others or not.
There is a broader point at play here. Legal scholars and policymak-
ers should focus less on what drives nudniks and more on the social
impact that nudniks generate. Focus more on the outputs and less on
the inputs. Even if nudniks are after revenge, material compensation,
validation from others, or satisfying their own sense of entitlement,
131. Unfortunately, it would seem that one has to have sufficient social capital to feel legiti-
mized to complain and to have their complaints taken seriously.
132. See Ostrom, supra note 105.
133. SeeLYNN STOUT,CULTIVATING CONSCIENCE: HOW GOOD LAWS MAKE GOOD PEOPLE
13–15 (2011) (describing “other-regarding behavior” as actions that express “concern for some-
one or something beyond one’s own material interests.”).
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 32 21-APR-20 11:48
264 DEPAUL LAW REVIEW [Vol.69:233
this should not matter too much, as long as nudniks generate positive
spillovers for other consumers.
This brings us to the third and most important issue with nudniks’
behavior: representativeness. The biggest question when evaluating
nudniks’ social impact is whether nudniks’ preferences and concerns
are representative of the preferences and concerns of other, non-
nudnik consumers. Some nudniks may complain about esoteric as-
pects of seller behavior that are inconsequential to other consumers.
A particular issue is that some serial complainers may be “trolls,” that
is, individuals who make spurious arguments for no other reason than
to evoke a reaction from their target. They seek attention rather than
improvement. For such troll-based activism, the social contribution is
limited, and the negative treatment these troll-nudniks will receive in
the courtroom and the court of public opinion is justified.
We start our response to the “unrepresentative” argument by not-
ing that such a critique is not unique to nudnik-based activism. The
critique rather applies to other types of active consumers as well. A
sophisticate filing a class action may be hunting a “bounty,” even if
the underlying cause is only technical and they were not really
harmed. Indeed, we noted the example of counterfeit hunters in
China who go out of their way to find fake products to buy, so that
they can later complain that they bought a fake product.134 Whenever
one introduces a bounty (a profit-motive), one raises the risk of nega-
tive-value behavior on part of the hunters. There is actually reason to
believe that such negative-value behavior will be more common
among non-nudniks; but more research is needed before reaching de-
finitive conclusions.
As for nudniks, the data there is limited, and future research is
much needed. Yet the existing literature contains a few indications
that lead us to believe that the “unrepresentative” concern is grossly
overstated. Elsewhere we synthesized findings from the consumer
complaining behavior literature, indicating that serial complainers
often operate in good faith and implicate broader consumer inter-
est.135 For example, we pointed out studies establishing a link between
seller behavior and consumer complaining—better service leads to
fewer complaints. Sellers who want to avoid the wrath of nudniks are
able to do so by offering a better product. Another finding is that
serial complainers are more likely to be loyal to a seller who rectifies
134. See Sui-Lee Wee, Though Awash in Fakes, China Rethinks Counterfeit Hunters, N.Y.
TIMES (Nov. 30, 2016), https://www.nytimes.com/2016/11/30/business/china-fakes-counterfeit-
hunters.html.
135. Arbel & Shapira, Theory of the Nudnik, supra note 9.
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 33 21-APR-20 11:48
2020] CONSUMER ACTIVISM 265
past issues. Again, this stylized fact of consumer complaining behavior
suggests a positive dynamic: It is not that serial complainers are only
after making a seller’s life miserable. If consumers perceive the seller
as making a genuine effort to satisfy their expectations, they will pay
the seller back by continuously purchasing from her. To reiterate, we
acknowledge that we cannot offer here conclusive proof on the repre-
sentativeness of nudniks’ expectations, or the ratio of valid-to-frivo-
lous complaints. Still, these findings are at least suggestive that
nudniks’ interests are correlated with those of other consumers.
Another reason to not overstate the concern with idiosyncratic ex-
pectations is that sellers do not remain passive. Sellers can, and often
do, employ various tools to defend themselves from “bad” nudniks
with unwarranted complaints. For example, sellers can avoid unrealis-
tic buyers’ expectations through salient and repeated communication.
To illustrate, consider how stores that do not accept certain modes of
payment frequently communicate this fact clearly to the consumer in
advance, with large and visible signs, despite the absence of any legal
duty to do so. Consider also the large signs one often encounters with
limitations, disclaimers, and special conditions on how meals are non-
returnable, sales are final, or seats on a flight are tentative. In other
words, sellers can reduce the risk of unrealistic consumer expectations
by modifying consumer expectations in advance. This is a feature of
nudnik-based consumer governance, rather than a bug in the system.
If there are hidden aspects of the transaction that consumers care
about, advertising them in a salient manner helps both parties deter-
mine in advance if the transaction is mutually advantageous. The
background threat of nudniks attacking sellers for violating their ex-
pectations incentivizes sellers to mitigate with the gap between trans-
actional expectations and the actual transaction.
Yet another reason to not overstate the costs of bad nudnik behav-
ior is the involvement of other, non-nudnik consumers. Nudniks’ com-
plaints create a reputational sanction only to the extent that other
consumers learn about, share, and act upon the information they re-
ceived from nudniks. These other consumers are not clueless; they can
decide for themselves whether the nudnik’s complaint raises a valid
issue or not. Dave Carroll’s “United Breaks Guitars” hurt United not
just because of Dave’s singing talent, but probably also because it res-
onated with other United consumers, hitting on a widely-shared frus-
tration with how airline companies treat luggage. If Dave would have
written a song about how Amazon does not deliver on time, we sus-
pect that his complaint would not have gone viral, because most con-
sumers have a positive experience with Amazon shipments.
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 34 21-APR-20 11:48
266 DEPAUL LAW REVIEW [Vol.69:233
To illustrate with another infamous example, consider the case of
Taylor Chapman. When Chapman did not receive a receipt at her last
visit to Dunkin’ Donuts, she sought to avail herself of the “free-donut-
if-we-don’t-give-you-a-receipt” policy.136 She recorded an exchange
where she demanded—using expletives and blatantly racist and sexist
remarks—that the store manager will provide her with a free meal.137
Chapman posted the exchange on social media, and it received exten-
sive exposure (over a million views), yet the result was not a reputa-
tional fallout for Dunkin’ Donuts, but rather a hit to Chapman’s own
reputation.138 Unlike with the United Airlines example, Dunkin’
Donuts did not experience a drop in stock price or in consumers’ will-
ingness to purchase.
To reiterate, we readily acknowledge that some nudnik complaints
are frivolous. But there is reason to believe that nudniks’ net social
impact is positive and that they fill in an important gap in the market-
place. It is therefore imperative to not dismiss the contribution of
nudniks because they may strike some as spiteful, selfish, or idiosyn-
cratic. Instead, we should strive to continue studying the conditions
under which nudniks are more or less valuable to market discipline.
CONCLUSION
Where does consumer governance come from? How do consumer
markets maintain norms of behavior? Legal institutions are not the
only ones deterring seller misbehavior; market mechanisms deter, too.
Understanding the effectiveness of market forces is key for legal
scholars and policymakers, as these forces set the outer limits on the
need for legal intervention. Yet, the legal literature has overly focused
on one type of market mechanism, namely, an informed minority,
which is seemingly less relevant in today’s world. This Essay suggests
switching focus to a different mechanism, namely, a minority of cru-
sading consumers—a small subset of consumers who go to great
lengths to complain publicly about seller misbehavior, and in the pro-
cess draw others’ attention and put in motion reputational sanctions.
Directing attention to the nudnik phenomenon is a first step toward
understanding consumer governance in today’s world. But it is just
136. Tsg600, Taylor Chapman Dunkin Donuts Rant, YOUTUBE (June 10, 2013), https://www
.youtube.com/watch?v=juLHmG76P4Q [hereinafter Dunkin Donuts Rant].
137. Id.
138. See, e.g., Joe Patrice, Aspiring Lawyer’s Insane Rant at Dunkin’ Donuts Staff, ABOVETHE
LAW (June 12, 2013, 10:10 AM), https://abovethelaw.com/2013/06/aspiring-lawyers-insane-rant-
at-dunkin-donuts-staff/; Dunkin Donuts Rant, supra note 136 (generating 1,084,248 views, 586
‘likes’ and 17,000 ‘dislikes’ as of Nov. 28, 2019).
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 35 21-APR-20 11:48
2020] CONSUMER ACTIVISM 267
that: a step. Other necessary steps include understanding how sellers
react to nudniks, and how the legal system affects the interactions be-
tween sellers and nudniks.139 While much future research is still
needed, the existing evidence from the consumer complaining behav-
ior literature makes us comfortable enough to risk ending on a gener-
alization: Nudniks are the unsung heroes of consumer markets.
139. We develop these issues in Arbel & Shapira, Theory of the Nudnik, supra note 9.
Electronic copy available at: https://ssrn.com/abstract=3568768

<<PAGE_BREAK>>

\\jciprod01\productn\D\DPL\69-2\DPL206.txt unknown Seq: 36 21-APR-20 11:48
268 DEPAUL LAW REVIEW [Vol.69:233
Electronic copy available at: https://ssrn.com/abstract=3568768)MW4LLM";

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
