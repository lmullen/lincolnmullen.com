---
name: new-post
description: Create a new blog post (titled or untitled short-form), with optional image and social media text, on a preview branch with a pull request. Use when the user asks to create, write, or add a new blog post or micro post.
---

# New post

Create a new blog post as a page bundle on its own branch, push it so Cloudflare builds a preview, and open a pull request so the user can check the post before publishing.

## Create the post

Do not run `./new-post.sh`: it ends by opening the file in VS Code with `code`, which does not work in a sandbox or cloud session. Create the page bundle by hand, using `archetypes/blog/index.md` as the model for the front matter.

- Use the current time in America/New_York (`TZ=America/New_York date`).
- **Untitled post**: the directory and `slug` are the timestamp, `YYYY-MM-DD-HH-MM-SS`. Omit the `title` field entirely; do not set it to `""`.
- **Titled post**: the directory is `YYYY-MM-DD-<slug>`, where the slug is the lowercased, hyphenated title with punctuation removed. Follow the title style rules in `CLAUDE.md`.
- `date` is ISO 8601 with the UTC offset, e.g. `'2026-10-25T09:01:22-04:00'`.
- Set `draft: false`, even though the archetype defaults to `true`. The branch preview and pull request serve as the draft stage. Mention this in the reply.

## Images

- Store the image in the page bundle. Give it a short, hyphenated filename that conveys the content (e.g. `zelda-paper-map.jpg`), not just a timestamp.
- Set `image:` in the front matter to the filename so link previews show it.
- Include it with the `image` shortcode (or `figure` if it needs a caption). Look at the image and write descriptive alt text.
- Put the body text above the image unless the user says otherwise.

## Text and style

- Italicize titles of works (games, books, journals, etc.) in the post body. Keep `description` and the social media text plain, without Markdown.
- Each social media field gets the text followed by the permalink: `https://lincolnmullen.com/blog/<slug>/`. See the `summarize-for-social-media` skill for the tone and length of each field.
- If the user says to skip a platform, set its field to `""`.
- When the user gives exact wording, use it verbatim.

## Branch, preview, and pull request

Follow the "Git and GitHub" section of `CLAUDE.md` for the worktree, the local Hugo server, pushing, and reading the Cloudflare preview URLs from the build. The only post-specific rules are:

- Name the branch `draft-post-<short-slug>`, where `<short-slug>` is a short, lowercase, hyphenated description of the post (e.g. `draft-post-zelda-on-paper`). The branch name becomes part of the preview URL, so keep it short.
- Build the site before pushing to catch errors: `hugo --cleanDestinationDir --minify -d <scratch dir>`.
- Open the pull request as a draft.
- Report the preview URL for the post itself (the branch preview URL followed by `/blog/<slug>/`), not just the site root.

## Handoff

Show the user the full Markdown file in the reply. After each revision, show the full file again.
