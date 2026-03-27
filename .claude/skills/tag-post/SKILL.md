---
name: tag-post
description: Add or update tags on a blog or newsletter post using the site's canonical tag taxonomy. Use when the user asks to tag a post, review tags, or add tags.
---

# Tag post

Add or update tags on the specified post file. If no file is given, use the file currently open in VS Code (provided via IDE context). If no file is open or specified, ask which post to tag.

Read the post's front matter and content, then select 1-5 tags from the taxonomy below. Prefer using tags from this taxonomy, but if the post is about a different subject, ask whether to use a new tag. If you create a new tag, add it to this taxonomy. It is OK to use both a general and specific tag (e.g., `digital history` and `computational history`). It is OK to leave a post untagged if nothing fits well.

## YAML format

Use non-indented list format:

```yaml
tags:
- digital history
- American religious history
```

## Tag taxonomy

### Subject tags

| Tag | Use for |
|-----|---------|
| `American religious history` | American religion as a historical subject: denominations, conversion, pluralism, demographics, evangelicalism, Catholicism, Mormonism, secularization |
| `history of Christianity` | Christianity broadly, including non-American or transatlantic contexts; distinct from American religious history |
| `history of Judaism` | Jewish history, distinct from antisemitism |
| `antisemitism` | History of antisemitism specifically |
| `legal history` | Nineteenth-century legal codes, civil procedure, Field Code, law broadly |
| `slavery` | History of slavery, abolition, mapping slavery data |

### Method/approach tags

| Tag | Use for |
|-----|---------|
| `digital history` | Digital history methods, argumentation, the field's development; umbrella for computational approaches to the past |
| `computational history` | Subset of digital history emphasizing quantitative/computational methods specifically |
| `data visualization` | Charts, graphs, interactive visualizations of historical data |
| `spatial history` | Maps, GIS, geographic analysis, spatial approaches to history |
| `text analysis` | Computational text analysis: text reuse, NLP, corpus analysis |
| `data` | Working with historical datasets, data sources, data cleaning |
| `R` | R programming language, R packages, ggplot2, the R ecosystem |
| `Go` | Go programming language |
| `Ruby` | Ruby programming language |
| `JavaScript` | JavaScript, D3.js, web development with JS |
| `programming` | General programming topics, software development practices |

### Activity/context tags

| Tag | Use for |
|-----|---------|
| `teaching` | Syllabi, course design, pedagogy, classroom practices |
| `scholarly publishing` | Posts *about* the publishing process, peer review, academic publishing |
| `publication` | Announcement of something the author published (article, book, preprint, dataset) |
| `book reviews` | Reviews of scholarly books |
| `RRCHNM` | Roy Rosenzweig Center for History and New Media -- its people, projects, news |
| `tools` | Practical technology for scholars: plain text workflows, Markdown, Pandoc, Vim, Git, Make, etc. |

### Project tags

| Tag | Use for |
|-----|---------|
| `America's Public Bible` | The America's Public Bible project |
| `American Religious Ecologies` | The American Religious Ecologies project and Census of Religious Bodies work |
| `Chance of Salvation` | The book *The Chance of Salvation* and the American Converts Database |
| `Legal Modernism` | The Legal Modernism project |
| `Collecting These Times` | The Collecting These Times project (pandemic Jewish experiences) |
| `Pandemic Religion` | The Pandemic Religion project |
| `DataScribe` | The DataScribe project |

## Judgment calls

- **`digital history` vs `tools`**: Post about using tech *for historical research* -> `digital history`. How-to about the tool itself -> `tools`. Can overlap.
- **`digital history` vs `computational history`**: Use both when a post is about specifically computational/quantitative methods within digital history. Use only `digital history` for broader DH topics.
- **Language tags** (`R`/`Go`/`Ruby`/`JavaScript`) can coexist with `programming`.
- **`data visualization` vs `spatial history`**: Maps and GIS -> `spatial history`. Charts and graphs -> `data visualization`. Can overlap.
- **`legal history` vs `Legal Modernism`**: `Legal Modernism` for posts about the specific project. `legal history` for general legal history. Can overlap.
- **`scholarly publishing` vs `publication`**: `scholarly publishing` = posts *about* the process. `publication` = announcing something the author published.
- **`American religious history` vs `history of Christianity`**: Use `American religious history` for US-focused. Use `history of Christianity` for broader/non-American contexts. Can overlap.
- **Micro/untitled posts** (timestamp slugs): Generally leave untagged unless content clearly fits a tag.

## Output

Show the proposed tags with a brief rationale for each, then ask: "Want me to apply these tags?"
