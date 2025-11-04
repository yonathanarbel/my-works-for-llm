"""Setup script for my-works-for-llm corpus tools."""
from setuptools import setup, find_packages
from pathlib import Path

# Read the README file
readme_file = Path(__file__).parent / 'README.md'
long_description = readme_file.read_text() if readme_file.exists() else ''

setup(
    name='arbel-corpus-tools',
    version='1.0.0',
    description='Tools for working with Professor Yonathan Arbel\'s scholarship corpus',
    long_description=long_description,
    long_description_content_type='text/markdown',
    author='Yonathan A. Arbel',
    author_email='',
    url='https://github.com/yonathanarbel/my-works-for-llm',
    license='CC BY-NC 4.0',
    py_modules=['generate_article_scripts', 'corpus_cli', 'corpus_api'],
    install_requires=[
        'rich>=10.0.0',
    ],
    entry_points={
        'console_scripts': [
            'arbel-corpus=corpus_cli:main',
            'generate-article-scripts=generate_article_scripts:main',
        ],
    },
    classifiers=[
        'Development Status :: 4 - Beta',
        'Intended Audience :: Science/Research',
        'Intended Audience :: Developers',
        'License :: Other/Proprietary License',
        'Programming Language :: Python :: 3',
        'Programming Language :: Python :: 3.8',
        'Programming Language :: Python :: 3.9',
        'Programming Language :: Python :: 3.10',
        'Programming Language :: Python :: 3.11',
        'Programming Language :: Python :: 3.12',
        'Topic :: Scientific/Engineering :: Artificial Intelligence',
        'Topic :: Text Processing',
    ],
    python_requires='>=3.8',
    keywords='legal scholarship nlp llm corpus machine-learning',
)
