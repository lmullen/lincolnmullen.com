---
name: newsletter
description: Put together the next issue of the Working on It email from posts published since the last issue, asking which posts to include in full or as a summary, then write it to emails/ and render HTML to paste into Kit. Use when the user asks to make, draft, or prepare a newsletter or email issue.
---

# Newsletter

Make the next issue of *Working on It*, the occasional email of new blog posts. The user chooses the posts, the skill writes the issue to `emails/`, and the user pastes it into Kit and sends it there. Nothing is uploaded or sent automatically.

## Running the helper script

`.claude/skills/newsletter/newsletter.py` does the mechanical work. Run it from the repository root. It needs `pyyaml` and `markdown`:

- Normally: `uv run --with pyyaml --with markdown .claude/skills/newsletter/newsletter.py <command>`
- If `uv` crashes (it does in the Docker sandbox on ARM, with `<jemalloc>: Unsupported system page size`), make a venv in the scratchpad: `python3 -m venv <scratch>/venv && <scratch>/venv/bin/pip install -q pyyaml markdown`, then run the script with `<scratch>/venv/bin/python`.

Commands:

- `last`: the last issue's number and date. If `emails/` has no issues yet, this is #18, the last issue published as a blog post.
- `candidates`: JSON list of published posts dated after the last issue, oldest first. Posts tagged `newsletter` are excluded.
- `item <dir> <full|summary>`: one post as email-ready Markdown. It adds the heading or date line and the permalink, and turns `image`/`figure`/`youtube`/`video` shortcodes into plain Markdown with absolute URLs to the original image files, which email clients can show. Warnings about unhandled shortcodes go to stderr. Fix those by hand.
- `render <issue.md>`: writes `<issue>.html` next to the issue, with simple inline styles.

## Steps

1. **Find the posts.** Run `last` and `candidates`. Show the user a short numbered list: date, title (or "note:" and the first line for untitled posts), and word count. If there are none, say so and stop.

2. **Choose posts.** Ask about each post with AskUserQuestion, up to four questions per call. The options are **Full**, **Summary**, and **Skip**. Put the suggested option first and mark it "(Recommended)":
   - Untitled notes: Full. They are short and have no summary.
   - Titled posts with `<!--more-->`, or over about 600 words: Summary.
   - Other titled posts: Full.
   - Summary means the text before `<!--more-->`, or the `description` if there is no break, followed by "Read the full post →".

3. **Subject, preview text, and intro.**
   - Draft a topic phrase for the subject, `Working on It #<n>: <topic>`, in sentence case, following the style guide in `CLAUDE.md`. Past topics: "Behind, ahead".
   - Draft `preview_text`, one plain sentence of about 90–140 characters.
   - Ask whether the user wants to write an intro paragraph or a "Currently" list, like the one that ended #18 (**Teaching:** …, **Reading:** …). Use the user's wording verbatim. Don't make up personal news or a "Currently" list. If the user gives none, leave those sections out.
   - Show the subject and preview text and let the user change them.

4. **Write the issue.** Copy `emails/_template.md` to `emails/YYYY-MM-DD-working-on-it-<n>.md`, using today's date in America/New_York. Then fill it in:
   - `{{number}}`, `{{topic}}`, and `{{preview_text}}`.
   - `{{date}}`: the current time, ISO 8601 with offset (`TZ=America/New_York date -Iseconds`). The next issue starts after this date, so use the time of writing, not a past date.
   - `{{posts}}`: one YAML item per included post, `  - {dir: <dir>, mode: <full|summary>}`.
   - `{{items}}`: the output of `item` for each chosen post, oldest first, separated by blank lines.
   - `{{intro}}` and `{{currently}}`: the user's text, or delete the placeholder line.
   - Delete any line left with only a placeholder. Leave the footer in place.

5. **Render** with `render emails/<issue>.md`. Check the stderr warnings, and spot-check the HTML: links absolute, images present, nothing left in `{{< >}}` or `{{ }}`.

6. **Hand off.** Don't paste the whole issue into the chat. Report:
   - The subject line and preview text, ready to copy.
   - The included posts and their modes.
   - The full path of the `.html` file. The repo is mounted from the user's Mac, so they can open it there.
   - How to send it: open the `.html` in a browser, select all, copy, and paste into a new Kit broadcast that uses a Classic template. Set the subject and preview text. Leave the broadcast **not public**: issues are not posted to Kit's web archive. Send a test email before sending to everyone.

7. **Commit.** Once the user says the issue is final, commit the `.md` file. The `.html` is ignored by git. Issues aren't published on the site, so commit on whatever branch the user is on unless they ask otherwise, and don't push `main` without asking. If the user edited the text in Kit after pasting, ask whether to bring those edits back into the `.md` so the archive matches what was sent.
