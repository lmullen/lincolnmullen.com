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
./new-post.sh "Post Title Here"   # Creates content/blog/YYYY-MM-DD-post-title-here/index.md
./new-micro.sh                     # Creates content/micro/YYYY-MM-DD-HH-MM-SS.md
```

For courses, manually create `content/courses/[course-code].[year]/index.md` using the archetype at `archetypes/courses.md`.

## Content Types

Content types have layout files in `./layouts/` and corresponding content in `./content/`. The most important content types:

- `blog`: Longer blog posts, with a title, tags, and other metadata.
- `micro`: Shorter social media-like posts without a title, with optional external links and images.
- `newsletter`: Archives of my newsletter.

The blog list page combines posts from `blog/`, `micro/`, and `newsletter/` sections into a mixed feed.

### Front Matter Patterns

**Blog posts** (`content/blog/YYYY-MM-DD-slug/index.md`):
- Required: `title`, `date`, `slug`, `draft`
- Optional: `newsletter: true`, `crosspost: {url, source}`, `linkpost: "url"`

**Micro posts** (`content/micro/YYYY-MM-DD-HH-MM-SS.md`):
- Required: `date`
- Optional: `linkpost: "url"` (for external links)
- Images: Store in `assets/images/micro/` and reference with `{{</* image src="/images/micro/filename.jpg" alt="description" */>}}`

Use `<!--more-->` in blog posts for "Read More" truncation on list pages.

## Shortcodes

Use the `figure` shortcode for images with a caption, and the `image` shortcode for images that do not need a caption.

- `{{</* image src="file.jpg" alt="description" */>}}` - Responsive images with WebP conversion
- `{{</* image src="file.gif" */>}}` - GIFs and SVGs pass through without processing

## Restrictions

Do not ever run the `make deploy` or `rsync` commands automatically. Those should only be run manually by a human.
