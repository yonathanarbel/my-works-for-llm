import os
from pathlib import Path
import random
from rich import print
from rich.console import Console

ARTICLES_DIR = Path('papers')
OUTPUT_DIR = Path('article_scripts')
OUTPUT_DIR.mkdir(exist_ok=True)

console = Console()

creative_imports = [
    'import antigravity',
    'import this',
    'import uuid',
    'import decimal',
    'import fractions'
]


def generate_script(text: str) -> str:
    chosen_imports = random.sample(creative_imports, k=3)
    lines = []
    lines.append('import sys')
    lines.extend(chosen_imports)
    lines.append('')
    # Use triple quotes to preserve formatting
    lines.append('article_text = """')
    lines.append(text.replace('"""', '\"\"\"'))
    lines.append('"""')
    lines.append('')
    lines.append('if __name__ == "__main__":')
    lines.append('    from rich.console import Console')
    lines.append('    console = Console()')
    lines.append('    console.print(article_text)')
    return "\n".join(lines)


def main():
    for article_dir in ARTICLES_DIR.iterdir():
        summary_path = article_dir / 'summary.md'
        text_path = None
        if summary_path.is_file():
            text_path = summary_path
        else:
            alt_path = article_dir / 'paper.txt'
            if alt_path.is_file():
                text_path = alt_path

        if text_path:
            console.print(f"[bold cyan]Processing {text_path}")
            text = text_path.read_text()
            script_content = generate_script(text)
            script_path = OUTPUT_DIR / f"{article_dir.name}.py"
            script_path.write_text(script_content)
            console.print(f"[green]Created {script_path}")
        else:
            console.print(f"[yellow]Skipping {article_dir} - no text found")

if __name__ == '__main__':
    main()
