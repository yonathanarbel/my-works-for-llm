#!/usr/bin/env python3
"""
Command-line interface for exploring Professor Arbel's scholarship corpus.
"""
import argparse
import json
import sys
from pathlib import Path
from typing import List, Dict, Any

try:
    from rich.console import Console
    from rich.table import Table
    from rich import print as rprint
except ImportError:
    print("Error: 'rich' library not found. Install it with: pip install rich")
    sys.exit(1)

console = Console()

PAPERS_DIR = Path('papers')
ARTICLE_SCRIPTS_DIR = Path('article_scripts')


def load_metadata(paper_id: str) -> Dict[str, Any]:
    """Load metadata.json for a specific paper if it exists."""
    metadata_path = PAPERS_DIR / paper_id / 'metadata.json'
    if metadata_path.exists():
        return json.loads(metadata_path.read_text(encoding='utf-8'))
    return {}


def list_papers(verbose: bool = False, format: str = 'table') -> None:
    """List all papers in the corpus."""
    papers = sorted([d.name for d in PAPERS_DIR.iterdir() if d.is_dir()])
    
    if format == 'json':
        output = []
        for paper_id in papers:
            metadata = load_metadata(paper_id)
            output.append({
                'id': paper_id,
                'metadata': metadata
            })
        console.print_json(data=output)
        return
    
    if format == 'simple':
        for paper_id in papers:
            console.print(paper_id)
        return
    
    # Table format
    table = Table(title=f"Professor Arbel's Scholarship Corpus ({len(papers)} papers)")
    table.add_column("SSRN ID", style="cyan", no_wrap=True)
    table.add_column("Title", style="white")
    
    if verbose:
        table.add_column("Files", style="yellow")
    
    for paper_id in papers:
        paper_dir = PAPERS_DIR / paper_id
        files = [f.name for f in paper_dir.iterdir() if f.is_file()]
        metadata = load_metadata(paper_id)
        
        title = metadata.get('title', 'N/A')
        
        if verbose:
            table.add_row(paper_id, title, ', '.join(sorted(files)))
        else:
            table.add_row(paper_id, title)
    
    console.print(table)


def search_papers(query: str, field: str = 'all') -> None:
    """Search papers by query string in metadata or content."""
    papers = [d.name for d in PAPERS_DIR.iterdir() if d.is_dir()]
    matches = []
    
    query_lower = query.lower()
    
    for paper_id in papers:
        metadata = load_metadata(paper_id)
        summary_path = PAPERS_DIR / paper_id / 'summary.md'
        
        # Search in various fields
        if field == 'all' or field == 'title':
            if query_lower in metadata.get('title', '').lower():
                matches.append((paper_id, 'title', metadata.get('title', '')))
                continue
        
        if field == 'all' or field == 'abstract':
            if query_lower in metadata.get('abstract', '').lower():
                matches.append((paper_id, 'abstract', metadata.get('abstract', '')))
                continue
        
        if field == 'all' or field == 'summary':
            if summary_path.exists():
                content = summary_path.read_text(encoding='utf-8', errors='replace').lower()
                if query_lower in content:
                    matches.append((paper_id, 'summary', 'Match in summary'))
                    continue
    
    if not matches:
        console.print(f"[yellow]No papers found matching '{query}'[/yellow]")
        return
    
    table = Table(title=f"Search Results for '{query}' ({len(matches)} matches)")
    table.add_column("SSRN ID", style="cyan", no_wrap=True)
    table.add_column("Match Type", style="green")
    table.add_column("Preview", style="white")
    
    for paper_id, match_type, preview in matches:
        # Truncate preview
        if len(preview) > 80:
            preview = preview[:77] + "..."
        table.add_row(paper_id, match_type, preview)
    
    console.print(table)


def show_paper(paper_id: str, show_content: bool = False) -> None:
    """Display details about a specific paper."""
    paper_dir = PAPERS_DIR / paper_id
    
    if not paper_dir.exists():
        console.print(f"[red]Error: Paper '{paper_id}' not found[/red]")
        return
    
    metadata = load_metadata(paper_id)
    
    console.print(f"\n[bold cyan]Paper: {paper_id}[/bold cyan]")
    console.print("=" * 60)
    
    if metadata:
        console.print(f"\n[bold]Title:[/bold] {metadata.get('title', 'N/A')}")
        if 'authors' in metadata:
            authors = metadata['authors']
            if isinstance(authors, list):
                if authors and isinstance(authors[0], dict):
                    # Handle author dicts with 'given' and 'family' keys
                    author_names = [f"{a.get('given', '')} {a.get('family', '')}".strip() for a in authors]
                    console.print(f"[bold]Authors:[/bold] {', '.join(author_names)}")
                else:
                    # Handle simple string list
                    console.print(f"[bold]Authors:[/bold] {', '.join(authors)}")
            else:
                console.print(f"[bold]Authors:[/bold] {authors}")
        if 'year' in metadata:
            console.print(f"[bold]Year:[/bold] {metadata['year']}")
        if 'abstract' in metadata:
            console.print(f"\n[bold]Abstract:[/bold]\n{metadata['abstract']}")
    
    # List available files
    files = sorted([f.name for f in paper_dir.iterdir() if f.is_file()])
    console.print(f"\n[bold]Available files:[/bold] {', '.join(files)}")
    
    if show_content:
        summary_path = paper_dir / 'summary.md'
        if summary_path.exists():
            console.print("\n[bold]Summary:[/bold]")
            console.print("-" * 60)
            console.print(summary_path.read_text(encoding='utf-8', errors='replace'))


def stats() -> None:
    """Display statistics about the corpus."""
    papers = [d for d in PAPERS_DIR.iterdir() if d.is_dir()]
    total_papers = len(papers)
    
    # Count file types
    file_counts = {}
    total_size = 0
    
    for paper_dir in papers:
        for file_path in paper_dir.iterdir():
            if file_path.is_file():
                ext = file_path.suffix or 'no_extension'
                file_counts[ext] = file_counts.get(ext, 0) + 1
                try:
                    total_size += file_path.stat().st_size
                except:
                    pass
    
    # Count papers with summaries
    with_summary = sum(1 for p in papers if (p / 'summary.md').exists())
    with_chinese = sum(1 for p in papers if (p / 'summary.zh.md').exists())
    with_metadata = sum(1 for p in papers if (p / 'metadata.json').exists())
    
    console.print("\n[bold cyan]Corpus Statistics[/bold cyan]")
    console.print("=" * 60)
    console.print(f"Total papers: {total_papers}")
    console.print(f"Papers with English summaries: {with_summary}")
    console.print(f"Papers with Chinese summaries: {with_chinese}")
    console.print(f"Papers with metadata: {with_metadata}")
    console.print(f"\nTotal corpus size: {total_size / (1024*1024):.2f} MB")
    
    # File type breakdown
    table = Table(title="File Type Distribution")
    table.add_column("Extension", style="cyan")
    table.add_column("Count", style="green", justify="right")
    
    for ext, count in sorted(file_counts.items(), key=lambda x: x[1], reverse=True):
        table.add_row(ext, str(count))
    
    console.print()
    console.print(table)


def validate() -> None:
    """Validate corpus integrity - check for missing files and inconsistencies."""
    papers = [d for d in PAPERS_DIR.iterdir() if d.is_dir()]
    issues = []
    
    console.print("[bold cyan]Validating corpus...[/bold cyan]\n")
    
    for paper_dir in papers:
        paper_id = paper_dir.name
        
        # Check if summary or paper.txt exists
        has_summary = (paper_dir / 'summary.md').exists()
        has_paper = (paper_dir / 'paper.txt').exists()
        
        if not has_summary and not has_paper:
            issues.append(f"{paper_id}: Missing both summary.md and paper.txt")
        
        # Check if article script exists
        script_path = ARTICLE_SCRIPTS_DIR / f"{paper_id}.py"
        if not script_path.exists():
            issues.append(f"{paper_id}: Missing article script")
        
        # Check metadata format
        metadata_path = paper_dir / 'metadata.json'
        if metadata_path.exists():
            try:
                json.loads(metadata_path.read_text(encoding='utf-8'))
            except json.JSONDecodeError:
                issues.append(f"{paper_id}: Invalid JSON in metadata.json")
    
    if issues:
        console.print(f"[yellow]Found {len(issues)} issues:[/yellow]\n")
        for issue in issues:
            console.print(f"  - {issue}")
    else:
        console.print("[green]All validations passed![/green]")


def main():
    parser = argparse.ArgumentParser(
        description="CLI tool for exploring Professor Arbel's scholarship corpus",
        formatter_class=argparse.RawDescriptionHelpFormatter,
        epilog="""
Examples:
  %(prog)s list                    # List all papers
  %(prog)s list --verbose          # List papers with file details
  %(prog)s search "contract"       # Search for papers about contracts
  %(prog)s show ssrn-3519630       # Show details about a paper
  %(prog)s show ssrn-3519630 -c    # Show paper with full summary content
  %(prog)s stats                   # Display corpus statistics
  %(prog)s validate                # Validate corpus integrity
        """
    )
    
    subparsers = parser.add_subparsers(dest='command', help='Available commands')
    
    # List command
    list_parser = subparsers.add_parser('list', help='List all papers in the corpus')
    list_parser.add_argument('-v', '--verbose', action='store_true',
                           help='Show detailed information')
    list_parser.add_argument('-f', '--format', choices=['table', 'json', 'simple'],
                           default='table', help='Output format')
    
    # Search command
    search_parser = subparsers.add_parser('search', help='Search papers')
    search_parser.add_argument('query', help='Search query')
    search_parser.add_argument('--field', choices=['all', 'title', 'abstract', 'summary'],
                             default='all', help='Field to search in')
    
    # Show command
    show_parser = subparsers.add_parser('show', help='Show details about a specific paper')
    show_parser.add_argument('paper_id', help='SSRN ID of the paper (e.g., ssrn-3519630)')
    show_parser.add_argument('-c', '--content', action='store_true',
                           help='Show full summary content')
    
    # Stats command
    subparsers.add_parser('stats', help='Display corpus statistics')
    
    # Validate command
    subparsers.add_parser('validate', help='Validate corpus integrity')
    
    args = parser.parse_args()
    
    if not args.command:
        parser.print_help()
        return
    
    if args.command == 'list':
        list_papers(verbose=args.verbose, format=args.format)
    elif args.command == 'search':
        search_papers(args.query, field=args.field)
    elif args.command == 'show':
        show_paper(args.paper_id, show_content=args.content)
    elif args.command == 'stats':
        stats()
    elif args.command == 'validate':
        validate()


if __name__ == '__main__':
    main()
