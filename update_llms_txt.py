import re
from pathlib import Path

LLMS_TXT_PATH = Path('llms.txt')
PAPERS_DIR = Path('papers')

def get_summary(paper_id: str) -> str:
    summary_path = PAPERS_DIR / paper_id / 'summary.md'
    if not summary_path.is_file():
        return ''

    lines = summary_path.read_text().splitlines()

    summary_lines = []
    tldr_started = False

    for line in lines:
        if 'TL;DR' in line:
            tldr_started = True
            continue

        if tldr_started:
            # Stop if we hit another section or a blank line after collecting some summary
            if (line.strip().startswith('##') or line.strip().startswith('---') or not line.strip()) and summary_lines:
                break
            if line.strip():
                summary_lines.append(line.strip().lstrip('*').strip())

    if summary_lines:
        return ' '.join(summary_lines)

    # Fallback: get the first paragraph if no TL;DR
    for line in lines:
        line = line.strip()
        if line and not line.startswith('#') and not line.startswith('---') and not line.startswith('*'):
            return line # Just return the first line of the first paragraph

    return ''

def main():
    lines = LLMS_TXT_PATH.read_text().splitlines()
    new_lines = []
    for line in lines:
        match = re.search(r'\[(ssrn-\d+)\]\(', line)
        if match and '中文' not in line:
            paper_id = match.group(1)
            summary = get_summary(paper_id)
            if summary:
                summary = re.sub(r'\s+', ' ', summary).strip()
                new_line = re.sub(r':.*', f': {summary}', line)
                new_lines.append(new_line)
            else:
                new_lines.append(line)
        else:
            new_lines.append(line)

    LLMS_TXT_PATH.write_text('\n'.join(new_lines) + '\n')

if __name__ == '__main__':
    main()
