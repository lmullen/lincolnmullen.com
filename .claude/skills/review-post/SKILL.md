---
name: review-post
description: Review a blog post against the site's style guide and front matter requirements. Use when the user asks to review, check, or validate a post before publishing.
---

# Review post

Review the specified blog post file. If no file is given, use the file currently open in VS Code (provided via IDE context). If no file is open or specified, ask which post to review.

Read the post file, then check every item below. Report all issues found, organized by category. After reporting, ask the user if they want you to fix the issues.

## Title

- The `title` field must use sentence case: only the first word and proper nouns are capitalized.
- Proper nouns to keep capitalized: people's names, place names, organization and institution names, acronyms, software and tool names, religious denominations (Catholic, Protestant, Baptist, Jewish, Presbyterian, etc.), and named works (book titles, journal names, project names, course names).
- Field and discipline names are lowercase: "digital humanities", "digital history", "religious history".
- The word "evangelical(s)" is lowercase when used as a general descriptor.
- Historical period names are lowercase: "early republic", "colonial".
- After a colon or semicolon within a title, do not capitalize the next word (unless it is a proper noun).
- After sentence-ending punctuation (? or !) within a title, capitalize the next word.
- Software package names keep their conventional capitalization (e.g., knitr, ggplot2, htmlwidgets stay lowercase; Pandoc, Jekyll, Vim stay capitalized).

## Required front matter

Check that these fields are present and non-empty:

- `title`
- `date`: Must be present in ISO 8601 format with timezone offset (e.g., `2026-02-10T23:24:10-05:00`). If the date is more than one hour before the current date/time, ask the user whether it should be updated to now.
- `slug`: Must be URL-safe: lowercase letters, numbers, and hyphens only. No punctuation, no spaces. Should match what `new-post.sh` generates from the title.
- `draft`: Must be `false`. Flag as an issue if it is `true` or missing.

## Optional front matter

Flag if these are missing, but as suggestions rather than errors:

- `description`: Should be 1-2 sentences summarizing the post for meta tags and search results. Platform-neutral and informative. This field is used for OpenGraph meta tags (`og:description`) and directly affects how the post appears when shared on social media. Flag prominently when missing.
- `social` block with `bluesky`, `mastodon`, and `linkedin` fields. Posts with empty `social.bluesky` or `social.mastodon` fields will not appear in the platform-specific RSS feeds, so these fields are important for distribution.

If present, validate:

- `bluesky`: Must be 300 characters or fewer (including URL). Should end with the permalink URL or an external URL (e.g., for linkposts or crossposts where the external URL is more appropriate).
- `mastodon`: Must be 500 characters or fewer (including URL). Should end with the permalink URL or an external URL.
- `linkedin`: No strict character limit. Often uses multi-line `|` YAML syntax for longer content. Should end with the permalink URL or an external URL.

It would be rare that the `bluesky` and `mastodon` fields would be empty. If `linkedin` is empty that will be more common. Note it but do not offer to correct it.

## Other optional front matter

These fields are used by templates. If present, validate them; if absent, do not flag unless context warrants it.

- `crosspost`: If present, must have `url` (a valid URL string) and `source` (name of the external site) sub-fields. Displayed as a notice that the post was originally published elsewhere.
- `linkpost`: If present, must be a valid URL string. Indicates the post links to an external resource; the URL is used as the canonical link.
- `newsletter`: If `true`, the post is styled as a newsletter entry (title displayed as "Working on It #<slug>: <title>").
- `tags`: If present, must be a YAML list. Tags appear in OpenGraph meta tags and the JSON feed. No tag pages are generated.

## Spelling and grammar

- Check the post for spelling errors and obvious grammatical mistakes.
- Only flag errors that are clearly wrong. Do not offer general writing suggestions, stylistic preferences, or rewording ideas.

## Content

- For longer posts (roughly 500+ words), suggest adding `<!--more-->` if absent. This provides "Read More" truncation on list pages. This is a suggestion, not a required check.
- Check that images use the `image` or `figure` shortcodes rather than raw Markdown image syntax.
- Verify that image files referenced in `image` or `figure` shortcodes actually exist in the post's page bundle directory or in `assets/`.

## Building up the style guide

As you review posts, you may encounter style questions not covered by the rules above (e.g., preferred spellings, hyphenation conventions, formatting patterns). When this happens, ask the user what their preference is. After they answer, update this skill file with the new rule so it is enforced consistently in future reviews.

## Output format

Organize findings as:

1. **Issues** (things that should be fixed)
2. **Suggestions** (things that could be improved but are not strictly wrong)
3. **OK** (brief confirmation of what passed)

Then ask: "Want me to fix the issues found?"
