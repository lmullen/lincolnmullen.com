---
name: draft-newsletter
description: Draft a new issue of the newsletter by gathering recent blog and micro posts. Use when the user asks to draft, create, or start a new newsletter issue.
---

# Draft newsletter

Draft a new issue of the _Working on It_ newsletter. Follow every step below in order.

## Step 1: Find the latest newsletter issue

Glob `content/newsletter/*/index.md` and determine the highest-numbered directory. Read that file's front matter to get:
- The issue number (from the directory name / slug)
- The `date` field (this is the "last issue date")

## Step 2: Create the new newsletter file

- Compute the new issue number: last issue number + 1
- Create the directory `content/newsletter/{new issue number}/`
- Create `content/newsletter/{new issue number}/index.md` with this front matter:

```yaml
---
title: ""
date: "{TODAY}T21:00:00-04:00"
slug: {new issue number}
draft: true
newsletter: true
social:
  bluesky: ""
  mastodon: ""
  linkedin: ""
---
```

Replace `{TODAY}` with today's date in `YYYY-MM-DD` format and `{new issue number}` with the computed number.

## Step 3: Opening paragraph

Calculate the number of days between today's date and the last issue's `date` field. Write this as the first content after the front matter:

```
Welcome back to the (very) occasional newsletter. It has been {NUMDAYS} days since the last issue.

<!--more-->
```

Replace `{NUMDAYS}` with the calculated number of days.

## Step 4: Blog post sections

Find all blog posts published since the last newsletter issue:
- Glob `content/blog/*/index.md`
- Read each file's front matter
- Include only posts where `date` is **after** the last newsletter's date AND `draft` is not `true`
- Sort by date, oldest first

For each qualifying blog post, add a section:

1. A level-three heading with the post's `title` exactly as it appears in front matter: `### {title}`
2. The body text from after the front matter up to and NOT including the `<!--more-->` marker, copied verbatim with one exception: **remove all footnote markers** (e.g., `[^1]`, `[^2]`) and **remove all footnote definitions** (lines starting with `[^1]:`, etc.). Footnotes don't work well in the newsletter context.
3. **Images**: If any `image` or `figure` shortcodes appear in the text before `<!--more-->`, copy the referenced image file from the blog post's page bundle directory into the new newsletter's directory. Update the `src` attribute in the shortcode to reference just the filename (since it will now be in the same page bundle). For images with paths starting with `/` (absolute paths), copy from `assets/` instead.
4. After the body excerpt, add a blank line and then: `[Read the rest](https://lincolnmullen.com/blog/{slug}/)`

   where `{slug}` is the `slug` field from the blog post's front matter.

## Step 5: Links and observations section

Add a section heading: `### Links and observations`

Find all micro posts published since the last newsletter issue:
- Glob `content/micro/*.md`
- Read each file's front matter
- Include only posts where `date` is **after** the last newsletter's date AND `draft` is not `true`
- Sort by date, oldest first

For each qualifying micro post, add a paragraph:

1. Draft a **one-sentence summary** that captures the essence of the micro post. Do NOT reproduce the body text verbatim — borrow from it and from any social media content in the front matter, but write an original summary sentence.
2. **Images**: If the micro post contains `image` or `figure` shortcodes, copy the referenced image file into the newsletter directory. For micro post images (paths like `/images/micro/filename.jpg`), the source file is at `assets/images/micro/filename.jpg`. Update the shortcode `src` to just the filename. Add the shortcode on its own line before the summary paragraph.
3. In the same paragraph as the summary, add a link:
   - If the post has a `linkpost` field: `[Read the post](https://lincolnmullen.com/micro/{date-slug}/) or [follow the link]({linkpost-url}).`
   - If no `linkpost`: `[Read the post](https://lincolnmullen.com/micro/{date-slug}/).`

**Micro post permalink format**: The `{date-slug}` is derived from the post's `date` field formatted as `YYYY-MM-DD-HH-MM-SS`. For example, a post dated `2026-02-06T15:09:00-05:00` has permalink `https://lincolnmullen.com/micro/2026-02-06-15-09-00/`.

## Step 6: Updates section

Add the following section. The three standard categories (`Reading`, `Listening`, `Working`) come first, followed by all other unique gerund-style labels found in `### Updates` sections across all previous newsletter issues.

To find the gerund labels: read the Updates sections of all previous newsletters and collect every line that starts with a capitalized gerund followed by a colon (e.g., `Playing:`, `Watching:`, `Hiking:`). Deduplicate them and exclude `Reading`, `Listening`, and `Working` (which are already listed). Also exclude `Also reading` and `Also listening` style variants.

```markdown
### Updates

Reading:

Listening:

Working:

{other gerund}:

{other gerund}:

...
```

List each gerund label with a colon and nothing after it. The user will fill in content and remove unwanted labels.

## Step 7: Review

After writing the complete newsletter draft, briefly summarize:
- How many blog posts were included
- How many micro posts were included
- The calculated days since last issue
- Remind the user to fill in the title, Updates section, and social media fields
