# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## About This Site

This website (<https://lincolnmullen.com>) is my professional, scholarly website. I am a professor of history.

## Technical Details

This site uses the static site generator, [Hugo](https://gohugo.io).

### Build Commands

- Build: `hugo --cleanDestinationDir --minify`
- Local preview: `hugo serve --port 2323 --buildDrafts --buildFuture --disableFastRender`

### Creating New Content

```bash
./new-post.sh   # Interactive: prompts for title and optional image
```

- **Titled post**: Enter a title when prompted. Creates `content/blog/YYYY-MM-DD-slug/index.md`.
- **Untitled post** (short-form/micro): Leave the title blank. Creates `content/blog/YYYY-MM-DD-HH-MM-SS/index.md` with the timestamp as the slug.
- Both types optionally include an image (drag-and-drop friendly).

For courses, manually create `content/courses/[course-code].[year]/index.md` using the archetype at `archetypes/courses.md`.

## Content Types

Content types have layout files in `./layouts/` and corresponding content in `./content/`. The most important content types:

- `blog`: All blog posts, both titled long-form articles and untitled short-form notes. Every post is a page bundle (`content/blog/<dir>/index.md`). Templates use title presence to distinguish: titled posts appear in the archive and get headings; untitled posts do not.
- `newsletter`: Archives of my newsletter.

The blog list page combines posts from `blog/` and `newsletter/` sections into a mixed feed.

### Front Matter Patterns

**Blog posts** (`content/blog/YYYY-MM-DD-slug/index.md`):
- Required: `date`, `slug`, `draft`
- Title: Optional. When present, the post appears in the archive and gets a heading. When absent, the post is a short-form note.
- Optional: `description`, `newsletter: true`, `crosspost: {url, source}`, `linkpost: "url"`, `via: "url"`, `image: "filename.jpg"`, `social: {bluesky, mastodon, linkedin}`, `tags`
- Images: Store in the page bundle directory and reference with `{{</* image src="filename.jpg" alt="description" */>}}`; the `image:` front matter field (relative filename) enables automatic OpenGraph image previews.

Use `<!--more-->` in blog posts for "Read More" truncation on list pages.

## Shortcodes

Use the `figure` shortcode for images with a caption, and the `image` shortcode for images that do not need a caption.

- `{{</* image src="file.jpg" alt="description" */>}}` - Responsive images with WebP conversion
- `{{</* image src="file.gif" */>}}` - GIFs and SVGs pass through without processing

## Style Guide

- Use sentence case for blog post titles: capitalize only the first word and proper nouns.
- Proper nouns to capitalize: people's names, place names, organization and institution names, acronyms, software and tool names, religious denominations (Catholic, Protestant, Baptist, Jewish, Presbyterian, etc.), and named works (book titles, journal names, project names, course names).
- Field and discipline names are lowercase: "digital humanities", "digital history", "religious history".
- The word "evangelical(s)" is lowercase when used as a general descriptor.
- Historical period names are lowercase: "early republic", "colonial".
- After a colon or semicolon within a title, do not capitalize the next word (unless it is a proper noun).
- After sentence-ending punctuation (? or !) within a title, capitalize the next word.
- Software package names keep their conventional capitalization (e.g., knitr, ggplot2, htmlwidgets stay lowercase; Pandoc, Jekyll, Vim stay capitalized).

### Permalink URLs

- **Blog posts** (titled): `https://lincolnmullen.com/blog/<slug>/`, where `<slug>` is the `slug` field in the front matter.
- **Blog posts** (untitled): `https://lincolnmullen.com/blog/<timestamp>/`, where `<timestamp>` is in `YYYY-MM-DD-HH-MM-SS` format.
- **Newsletter posts**: `https://lincolnmullen.com/newsletter/<slug>/`, where `<slug>` is the `slug` field in the front matter.

Include the permalink URL at the end of each social media text.
