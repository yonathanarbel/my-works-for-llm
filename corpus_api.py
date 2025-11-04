"""
Python API for programmatic access to Professor Arbel's scholarship corpus.

Example usage:
    from corpus_api import ArbelCorpus
    
    corpus = ArbelCorpus()
    papers = corpus.list_papers()
    
    paper = corpus.get_paper('ssrn-3519630')
    summary = paper.get_summary()
    full_text = paper.get_full_text()
"""
import json
from pathlib import Path
from typing import List, Dict, Any, Optional


class Paper:
    """Represents a single paper in the corpus."""
    
    def __init__(self, paper_id: str, base_dir: Path):
        self.paper_id = paper_id
        self.base_dir = base_dir
        self.paper_dir = base_dir / paper_id
        self._metadata = None
    
    @property
    def metadata(self) -> Dict[str, Any]:
        """Load and cache paper metadata."""
        if self._metadata is None:
            metadata_path = self.paper_dir / 'metadata.json'
            if metadata_path.exists():
                self._metadata = json.loads(metadata_path.read_text())
            else:
                self._metadata = {'id': self.paper_id}
        return self._metadata
    
    def get_summary(self, language: str = 'en') -> Optional[str]:
        """
        Get paper summary.
        
        Args:
            language: Language code ('en' for English, 'zh' for Chinese)
        
        Returns:
            Summary text or None if not available
        """
        if language == 'en':
            summary_path = self.paper_dir / 'summary.md'
        elif language == 'zh':
            summary_path = self.paper_dir / 'summary.zh.md'
        else:
            raise ValueError(f"Unsupported language: {language}")
        
        if summary_path.exists():
            return summary_path.read_text()
        return None
    
    def get_full_text(self) -> Optional[str]:
        """Get full paper text."""
        text_path = self.paper_dir / 'paper.txt'
        if text_path.exists():
            return text_path.read_text()
        return None
    
    def get_pdf_path(self) -> Optional[Path]:
        """Get path to PDF file if available."""
        pdf_path = self.paper_dir / 'paper.pdf'
        return pdf_path if pdf_path.exists() else None
    
    def get_article_script(self) -> Optional[str]:
        """Get the generated article script content."""
        script_path = Path('article_scripts') / f'{self.paper_id}.py'
        if script_path.exists():
            return script_path.read_text()
        return None
    
    def get_available_files(self) -> List[str]:
        """List all available files for this paper."""
        if not self.paper_dir.exists():
            return []
        return sorted([f.name for f in self.paper_dir.iterdir() if f.is_file()])
    
    def get_title(self) -> str:
        """Get paper title from metadata or paper ID."""
        return self.metadata.get('title', self.paper_id)
    
    def get_authors(self) -> List[str]:
        """Get list of authors from metadata."""
        return self.metadata.get('authors', [])
    
    def get_year(self) -> Optional[int]:
        """Get publication year from metadata."""
        return self.metadata.get('year')
    
    def get_abstract(self) -> Optional[str]:
        """Get paper abstract from metadata."""
        return self.metadata.get('abstract')
    
    def has_summary(self, language: str = 'en') -> bool:
        """Check if summary exists in specified language."""
        if language == 'en':
            return (self.paper_dir / 'summary.md').exists()
        elif language == 'zh':
            return (self.paper_dir / 'summary.zh.md').exists()
        return False
    
    def __repr__(self) -> str:
        return f"Paper(id='{self.paper_id}', title='{self.get_title()}')"


class ArbelCorpus:
    """Main interface for accessing Professor Arbel's scholarship corpus."""
    
    def __init__(self, base_dir: Optional[Path] = None):
        """
        Initialize corpus access.
        
        Args:
            base_dir: Base directory of the corpus (defaults to current directory)
        """
        self.base_dir = base_dir or Path('.')
        self.papers_dir = self.base_dir / 'papers'
        
        if not self.papers_dir.exists():
            raise ValueError(f"Papers directory not found at {self.papers_dir}")
    
    def list_papers(self) -> List[Paper]:
        """
        Get list of all papers in the corpus.
        
        Returns:
            List of Paper objects
        """
        paper_ids = sorted([d.name for d in self.papers_dir.iterdir() if d.is_dir()])
        return [Paper(paper_id, self.papers_dir) for paper_id in paper_ids]
    
    def get_paper(self, paper_id: str) -> Paper:
        """
        Get a specific paper by ID.
        
        Args:
            paper_id: SSRN ID (e.g., 'ssrn-3519630')
        
        Returns:
            Paper object
        
        Raises:
            ValueError: If paper not found
        """
        paper = Paper(paper_id, self.papers_dir)
        if not paper.paper_dir.exists():
            raise ValueError(f"Paper '{paper_id}' not found")
        return paper
    
    def search_papers(self, query: str, field: str = 'all') -> List[Paper]:
        """
        Search papers by query string.
        
        Args:
            query: Search query
            field: Field to search in ('all', 'title', 'abstract', 'summary')
        
        Returns:
            List of matching Paper objects
        """
        papers = self.list_papers()
        matches = []
        query_lower = query.lower()
        
        for paper in papers:
            if field in ['all', 'title']:
                if query_lower in paper.get_title().lower():
                    matches.append(paper)
                    continue
            
            if field in ['all', 'abstract']:
                abstract = paper.get_abstract()
                if abstract and query_lower in abstract.lower():
                    matches.append(paper)
                    continue
            
            if field in ['all', 'summary']:
                summary = paper.get_summary()
                if summary and query_lower in summary.lower():
                    matches.append(paper)
                    continue
        
        return matches
    
    def get_stats(self) -> Dict[str, Any]:
        """
        Get corpus statistics.
        
        Returns:
            Dictionary with corpus statistics
        """
        papers = self.list_papers()
        
        stats = {
            'total_papers': len(papers),
            'with_english_summary': sum(1 for p in papers if p.has_summary('en')),
            'with_chinese_summary': sum(1 for p in papers if p.has_summary('zh')),
            'with_metadata': sum(1 for p in papers if (p.paper_dir / 'metadata.json').exists()),
            'with_pdf': sum(1 for p in papers if p.get_pdf_path() is not None),
            'with_full_text': sum(1 for p in papers if p.get_full_text() is not None),
        }
        
        return stats
    
    def iterate_papers(self):
        """
        Generator to iterate through all papers.
        
        Yields:
            Paper objects
        """
        for paper in self.list_papers():
            yield paper
    
    def __len__(self) -> int:
        """Return number of papers in corpus."""
        return len(self.list_papers())
    
    def __repr__(self) -> str:
        return f"ArbelCorpus(papers={len(self)})"


# Convenience function for quick access
def load_corpus(base_dir: Optional[Path] = None) -> ArbelCorpus:
    """
    Load the corpus for easy access.
    
    Args:
        base_dir: Base directory of the corpus
    
    Returns:
        ArbelCorpus instance
    
    Example:
        >>> corpus = load_corpus()
        >>> papers = corpus.list_papers()
        >>> print(f"Found {len(papers)} papers")
    """
    return ArbelCorpus(base_dir)
