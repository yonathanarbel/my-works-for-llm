"""
Thinking the Unthinkable: AI in the Service of Justice (thinking-unthinkable-ai-justice) — corpus code wrapper

This file intentionally embeds the paper text and study assets in code form.
It helps code-centric ingestion pipelines and makes the corpus easy to load programmatically.
"""

from __future__ import annotations

PAPER_ID = "thinking-unthinkable-ai-justice"
TITLE = 'Thinking the Unthinkable: AI in the Service of Justice'
SSRN_URL = None
YEAR = 2024
AUTHORS = ['Yonathan A. Arbel']
KEYWORDS = ['generative interpretation', 'large language models', 'ordinary meaning', 'legal interpretation', 'judicial decision-making', 'prompt documentation', 'human oversight', 'access to justice']

SUMMARY_MD = """
# Thinking the Unthinkable: AI in the Service of Justice — structured summary

**Author:** Yonathan A. Arbel

**Source:** *Capstone Lawyer* (2024), University of Alabama

1.  ## TL;DR
    Professor Yonathan Arbel argues that large language models can provide useful, context-sensitive evidence about ordinary meaning and thereby supplement dictionaries, precedent, interpretive canons, and judicial intuition. Judge Kevin Newsom's concurrence in the Eleventh Circuit's *Snell* trampoline-insurance dispute illustrates the approach's movement into judicial practice. Arbel recommends multiple-model cross-checking, preservation of prompts and answers, advisory rather than determinative use, and continuing human oversight.

2.  ## Section Summaries

    *   **The interpretive problem**
        Conventional sources promote stability but may be static, context-poor, or detached from current and vernacular usage. Interpretive error can affect coverage, contractual obligation, statutory meaning, litigation cost, and consistency.

    *   **Generative interpretation**
        Because LLMs learn statistical patterns from unusually broad text corpora, they may offer evidence about how language is actually used across domains and communities. The method draws on Arbel's joint work with David Hoffman and treats model output as linguistic evidence rather than legal authority.

    *   **The Snell concurrence**
        Judge Newsom asked several models whether installing an in-ground trampoline counts as landscaping. Their answers were affirmative and useful to his ordinary-meaning inquiry, but other factors determined the legal result. The episode is significant as an early judicial engagement with generative interpretation, not as binding acceptance of it.

    *   **Applications and safeguards**
        The essay identifies insurance, prenuptial agreements, and commercial jargon as possible use cases. It recommends querying multiple models, recording prompts and outputs for replication, treating results as advisory, and maintaining human review for bias, hallucination, and error.

    *   **Institutional impact**
        Lawyers may need new prompting and evaluation skills; judges may gain broader language evidence; and clients may benefit from more predictable interpretation. These possibilities coexist with risks of overreliance, bias, and erosion of discretion or empathy.

    *   **Broader agenda and limits**
        Generative interpretation sits within a wider program of AI legal education, consumer tools, systemic-risk governance, and simulation of legal standards such as the reasonable person. Its success should be measured by careful monitoring, methodological refinement, and contribution to a fair and just legal system—not by adoption alone.
"""

SUMMARY_ZH_MD = """
# 《思考不可想象之事：让人工智能服务于正义》结构化摘要

**作者：** Yonathan A. Arbel

**来源：** 阿拉巴马大学 *Capstone Lawyer*（2024）

## 核心观点

Arbel 教授认为，大语言模型可以提供关于日常语义的情境化证据，从而补充词典、先例、解释准则与法官直觉。美国第十一巡回上诉法院 Kevin Newsom 法官在 *Snell* 蹦床保险争议中的协同意见，展示了“生成式解释”进入司法实践的早期形态。模型输出应当被视为语言证据，而非法律权威或自动答案。

## 解释难题与生成式方法

传统解释工具能够维持稳定性，却可能静态、缺乏语境或落后于现代口语。大语言模型从广泛文本中学习统计模式，因此可能呈现不同领域与社群中的实际用法。该方法建立在 Arbel 与 David Hoffman 的共同研究之上。

## Snell 案与司法承认

Newsom 法官向多个模型询问安装地下蹦床是否属于“景观美化”，模型均作肯定回答。答案对日常语义分析有启发，但并未决定案件结论。其重要性在于司法开始认真考虑人工智能辅助解释，而不在于形成有约束力的规则。

## 应用与保障

论文以飓风保险、婚前协议和商业术语为例，说明模型可能帮助理解有争议的词语。作者建议使用多个模型交叉核验，完整保存提示与回答以便复核，把输出限定为咨询性材料，并保留人类监督以发现偏见、幻觉与错误。

## 制度影响与边界

律师可能需要新的提示与评估能力，法官可能获得更丰富的语言材料，客户也可能从更可预测的解释中受益。但过度依赖、偏见以及对司法裁量与同理心的影响仍是开放问题。生成式解释的成败应以持续监测、方法改进以及是否促进公平正义来衡量，而不能仅看采用数量。
"""

ONE_PAGER_MD = """
# Thinking the Unthinkable: AI in the Service of Justice — one-page summary

**Paper ID:** `thinking-unthinkable-ai-justice`
**Year:** 2024
**Author(s):** Yonathan A. Arbel

## TL;DR

Professor Yonathan Arbel argues that large language models can provide useful, context-sensitive evidence about ordinary meaning and thereby supplement dictionaries, precedent, interpretive canons, and judicial intuition. Judge Kevin Newsom's concurrence in the Eleventh Circuit's *Snell* trampoline-insurance dispute illustrates the approach's movement into judicial practice. Arbel recommends multiple-model cross-checking, preservation of prompts and answers, advisory rather than determinative use, and continuing human oversight.

## Key Sections (from `summary.md`)

- **The interpretive problem:** Conventional sources promote stability but may be static, context-poor, or detached from current and vernacular usage. Interpretive error can affect coverage, contractual obligation, statutory meaning, litigation cost, and consistency.
- **Generative interpretation:** Because LLMs learn statistical patterns from unusually broad text corpora, they may offer evidence about how language is actually used across domains and communities. The method draws on Arbel's joint work with David Hoffman and treats model output as linguistic evidence rather than legal authority.
- **The Snell concurrence:** Judge Newsom asked several models whether installing an in-ground trampoline counts as landscaping. Their answers were affirmative and useful to his ordinary-meaning inquiry, but other factors determined the legal result. The episode is significant as an early judicial engagement with generative interpretation, not as binding acceptance of it.
- **Applications and safeguards:** The essay identifies insurance, prenuptial agreements, and commercial jargon as possible use cases. It recommends querying multiple models, recording prompts and outputs for replication, treating results as advisory, and maintaining human review for bias, hallucination, and error.
- **Institutional impact:** Lawyers may need new prompting and evaluation skills; judges may gain broader language evidence; and clients may benefit from more predictable interpretation. These possibilities coexist with risks of overreliance, bias, and erosion of discretion or empathy.

## Keywords

generative interpretation; large language models; ordinary meaning; legal interpretation; judicial decision-making; prompt documentation; human oversight; access to justice

## Files

- Full text: `papers/thinking-unthinkable-ai-justice/paper.txt`
- Summary (EN): `papers/thinking-unthinkable-ai-justice/summary.md`
- Summary (ZH): `papers/thinking-unthinkable-ai-justice/summary.zh.md`

_Auto-generated study aid. For canonical content, rely on `paper.txt` and the cited authoritative online source._
"""

STUDY_PACK_MD = """
# Study pack: Thinking the Unthinkable: AI in the Service of Justice (thinking-unthinkable-ai-justice)

- Full text: `papers/thinking-unthinkable-ai-justice/paper.txt`
- Summary (EN): `papers/thinking-unthinkable-ai-justice/summary.md`
- Summary (ZH): `papers/thinking-unthinkable-ai-justice/summary.zh.md`

## Elevator pitch

Professor Yonathan Arbel argues that large language models can provide useful, context-sensitive evidence about ordinary meaning and thereby supplement dictionaries, precedent, interpretive canons, and judicial intuition. Judge Kevin Newsom's concurrence in the Eleventh Circuit's *Snell* trampoline-insurance dispute illustrates the approach's movement into judicial practice. Arbel recommends multiple-model cross-checking, preservation of prompts and answers, advisory rather than determinative use, and continuing human oversight.

## Structured outline (high-signal)

### The interpretive problem

Conventional sources promote stability but may be static, context-poor, or detached from current and vernacular usage. Interpretive error can affect coverage, contractual obligation, statutory meaning, litigation cost, and consistency.

### Generative interpretation

Because LLMs learn statistical patterns from unusually broad text corpora, they may offer evidence about how language is actually used across domains and communities. The method draws on Arbel's joint work with David Hoffman and treats model output as linguistic evidence rather than legal authority.

### The Snell concurrence

Judge Newsom asked several models whether installing an in-ground trampoline counts as landscaping. Their answers were affirmative and useful to his ordinary-meaning inquiry, but other factors determined the legal result. The episode is significant as an early judicial engagement with generative interpretation, not as binding acceptance of it.

### Applications and safeguards

The essay identifies insurance, prenuptial agreements, and commercial jargon as possible use cases. It recommends querying multiple models, recording prompts and outputs for replication, treating results as advisory, and maintaining human review for bias, hallucination, and error.

### Institutional impact

Lawyers may need new prompting and evaluation skills; judges may gain broader language evidence; and clients may benefit from more predictable interpretation. These possibilities coexist with risks of overreliance, bias, and erosion of discretion or empathy.

### Broader agenda and limits

Generative interpretation sits within a wider program of AI legal education, consumer tools, systemic-risk governance, and simulation of legal standards such as the reasonable person. Its success should be measured by careful monitoring, methodological refinement, and contribution to a fair and just legal system—not by adoption alone.

## Keywords / concepts

generative interpretation; large language models; ordinary meaning; legal interpretation; judicial decision-making; prompt documentation; human oversight; access to justice

## Suggested questions (for RAG / study)

- What is the paper’s main claim and what problem does it solve?
- What method/data does it use (if any), and what are the main results?
- What assumptions are doing the most work?
- What are the limitations or failure modes the author flags?
- How does this connect to the author’s other papers in this corpus?

_Auto-generated study aid. For canonical content, rely on `paper.txt` and the cited authoritative online source._
"""

ARTICLE_TEXT = """
Canonical citation:
Yonathan A. Arbel, Thinking the Unthinkable: AI in the Service of Justice, Capstone Lawyer (2024).

Authoritative source:
https://stories.ua.edu/thinking-the-unthinkable-ai-in-the-service-of-justice/index.html

The following text was extracted from the authoritative University of Alabama page on 2026-09-04.

Thinking the Unthinkable: AI in the Service of Justice

By Yonathan Arbel
Professor of Law
Director, Artificial Intelligence Initiative

A case involving a simple trampoline installation in the quiet town of Fairhope, Alabama, may unexpectedly become the catalyst for a profound shift in legal interpretation. In 2020, James Snell, owner of Outdoor Expressions, installed an in-ground trampoline for a local family. When a young girl later injured herself on the unprotected trampoline, Snell found himself embroiled in a lawsuit. His insurance company refused coverage, claiming that trampoline installation wasn’t part of the “landscaping” his policy covered.

Though seemingly straightforward, this dispute has brought to light a fundamental challenge in legal practice — the interpretation of language — potentially marking the dawn of a new era where artificial intelligence (AI) assists in the pursuit of justice and revolutionizes how we understand and apply the law.

The Challenge of Legal Interpretation

For centuries, courts have relied on a limited toolkit to interpret the meaning of words in legal contexts. This toolkit primarily consists of dictionaries, precedents, and Latin maxims like “expressio unius” (the expression of one thing implies the exclusion of others) and “contra proferentem” (ambiguous terms should be interpreted against the drafter). Judges, as native speakers, often rely on their intuition about language. However, these traditional methods have clear limitations.

Dictionaries, while useful, can be outdated and miss nuanced meanings. They cannot account for context, which is crucial in understanding language use. Precedents, while valuable for consistency, often lag behind societal changes, potentially leading to interpretations that are out of touch with current language use. Judges, despite their expertise, are frequently criticized for being disconnected from modern or vernacular language.

These limitations can have significant consequences. Misinterpretation of contract terms, statutes, or other legal documents can lead to unjust outcomes, protracted litigation, and inconsistent application of the law. The need for a more robust, context-aware method of interpretation has long been apparent, but, until recently, no viable solution has emerged.

Enter Generative Interpretation

The trampoline case was the first time that a judge used the novel approach of “generative interpretation” as a form of legal interpretation.

The essence of this approach is elegantly simple yet profound. Large language models (LLMs), trained on vast amounts of internet text data, can provide insights into how language is actually used and understood in real-world contexts. These models are exposed to more textual data during their training than any single human could read in a hundred lifetimes. Through this process, they develop a statistical understanding of language so comprehensive that some models can even pass the bar exam.

This vast reservoir of language knowledge can be tapped to generate a more comprehensive and nuanced understanding of ordinary meaning, potentially revolutionizing how we interpret legal texts. Unlike traditional methods, LLMs can provide context-aware interpretations that reflect current language use across various domains and communities.

This approach is based on my work with Professor David Hoffman of the University of Pennsylvania Law School. Soon to be published in the New York University Law Review, it develops the principles for harnessing the power of LLMs like ChatGPT to inform legal interpretation.

Judicial Recognition and Future Prospects

Back to the trampoline. In a groundbreaking concurring opinion, Judge Kevin Newsom of the U.S. Court of Appeals for the Eleventh Circuit boldly proposed what he calls a “heresy” — that AI-powered language models might inform legal interpretation of ordinary meaning. He relied extensively on our work in informing his analysis of the case.

Judge Newsom posed the pivotal question to various models like ChatGPT: “Is installing an in-ground trampoline ‘landscaping’?” The answer, by all models, was “yes.” They interpreted that in common usage with context provided by the judge, installing an in-ground trampoline would indeed be considered part of a landscaping job.

The affirmative response, while not determinative in this case due to other factors, led Judge Newsom to a profound realization: “At the very least, it no longer strikes me as ridiculous to think that an LLM like ChatGPT might have something useful to say about the common, everyday meaning of the words and phrases used in legal texts.”

This judicial recognition marks a significant milestone in the acceptance of AI-assisted legal interpretation. It suggests that the legal community is beginning to seriously grapple with the potential of these technologies to enhance, and perhaps, transform traditional legal reasoning.

Practical Applications and Best Practices

In our work, Hoffman and I sought to go beyond theory and offer practical guidance for implementing generative interpretation in legal practice. We developed best practices for using LLMs, addressing potential pitfalls like prompt bias and artificial intelligence hallucinations. We also demonstrated the method’s potential through case studies, showing how LLMs could have aided courts in resolving complex interpretative puzzles.

For instance, in cases involving the interpretation of insurance policies after Hurricane Katrina, LLMs could have provided insights into the common understanding of terms like “flood” and “storm surge.” In family court disputes over prenuptial agreements, they could offer nuanced interpretations of potentially ambiguous terms. In commercial agreements, LLMs could help clarify industry-specific jargon that might be misunderstood by generalist judges.

Importantly, generative interpretation doesn’t seek to replace traditional methods but to complement them. Where dictionaries offer static definitions and precedents may lag behind societal changes, LLMs can provide real-time insights into language usage across diverse contexts. For instance, in a case involving technical jargon, an LLM could offer nuanced context-aware interpretations based on usage in specific industries, outperforming general-purpose dictionaries.

The implementation of this approach, however, requires careful consideration. We suggest four best practices:

1. Use multiple LLMs to cross-verify results, mitigating the risk of relying on a single model’s potential biases or errors.
2. Clearly document prompts and responses for transparency, allowing for scrutiny and replication of the interpretation process.
3. Always view LLM outputs as advisory rather than determinative, preserving the crucial role of human judgment in legal decision-making.
4. Emphasize the importance of human oversight to catch potential biases or errors in AI responses.

Impact on the Legal System

This paradigm shift could significantly impact various stakeholders in the legal system. Lawyers might need to develop new skills in prompt engineering and AI interpretation, potentially changing how they prepare and argue cases. Judges could have access to more comprehensive language data, potentially leading to more informed and consistent decisions. Clients might also benefit from more predictable and consistent legal interpretations, potentially reducing litigation costs and uncertainty.

However, the adoption of generative interpretation also raises important questions about the evolving role of human judgment in law and the potential for over-reliance on AI tools. These questions will likely occupy legal philosophers and ethicists for years to come. How do we balance the benefits of AI-assisted interpretation with the need for human discretion and empathy in the legal process? How can we ensure that the use of LLMs doesn’t inadvertently perpetuate or amplify existing biases in the legal system?

Broader Implications and Future Research

This work on generative interpretation is part of a broader initiative, which I am leading as the director of AI studies at The University of Alabama School of Law. This new initiative seeks to position the Law School at the frontier of legal pedagogy in the nation. As part of this initiative, the school offers an innovative “AI for Lawyers” course, where law students build practical AI tools to assist in legal tasks.

My broader research agenda spans various aspects of AI in law, from managing systemic risks from AI, to integrating AI into the judiciary. I explore offering consumer AI tools to deal with complex legal documents and building new AI tools to simulate how juries and ordinary people determine who is the “reasonable person.” These efforts represent a comprehensive approach to understanding and shaping the role of AI in the legal system and mark the Law School's commitment to innovation in pedagogy and research.

As we stand on the brink of this new era, the potential of AI in legal interpretation is both exciting and challenging. It promises more accurate, consistent, and accessible legal interpretations, potentially reducing litigation and improving access to justice. However, it also raises important questions about the role of human judgment in law and the ethical implications of AI in legal decision-making.

Conclusion

In the end, the story of a trampoline in Fairhope, Alabama, may be remembered not for the legal dispute it sparked, but for the transformative ideas it helped to propel into the mainstream of legal thought. As we move forward, frameworks developed by scholars, and forward-thinking jurists like Judge Newsom, will undoubtedly play a crucial role in shaping the future of legal interpretation in the age of artificial intelligence.

The integration of AI into legal interpretation represents a significant shift in how we approach the law. It offers the potential to make legal interpretation more accurate, consistent, and responsive to real-world language use. At the same time, it challenges us to rethink fundamental aspects of legal reasoning and decision-making.

As this new approach gains traction, it will be crucial to monitor its impacts carefully, refine the methodologies, and ensure that the use of AI in legal interpretation serves the ultimate goal of a fair and just legal system. The symbiotic relationship, with scholarly and judicial work developing in tandem, gives us reason to be optimistic about the future of AI in law.

In the coming years, we can expect to see further developments in this field, with more courts potentially adopting AI-assisted interpretation methods, law schools incorporating AI into their curricula, and legal technology companies developing specialized tools for generative interpretation. The legal profession, known for its traditionalism, stands at the cusp of a technological revolution that could fundamentally change how we understand and apply the law.

As we navigate this new frontier, it will be essential to maintain a balance between embracing the potential of AI and preserving the human elements that are central to the practice of law. The story of generative interpretation is not just about technology; it’s about reimagining how we can make the legal system more accurate, efficient, and ultimately, more just.
"""


def as_dict() -> dict[str, object]:
    return {
        "paper_id": PAPER_ID,
        "title": TITLE,
        "ssrn_url": SSRN_URL,
        "year": YEAR,
        "authors": AUTHORS,
        "keywords": KEYWORDS,
        "summary_md": SUMMARY_MD,
        "summary_zh_md": SUMMARY_ZH_MD,
        "one_pager_md": ONE_PAGER_MD,
        "study_pack_md": STUDY_PACK_MD,
        "article_text": ARTICLE_TEXT,
    }


if __name__ == "__main__":
    print(ARTICLE_TEXT)
