---
name: newsletter
description: Put together the next issue of the Working on It email from posts published since the last issue, asking which posts to include, then write it to emails/ and render HTML to paste into Kit. Use when the user asks to make, draft, or prepare a newsletter or email issue.
---

# Newsletter

Make the next issue of *Working on It*, the occasional email of new blog posts. The user chooses the posts, the skill writes the issue to `emails/`, and the user pastes it into Kit and sends it there. Nothing is uploaded or sent automatically.

## The helper program

The mechanical work is done by a Go program in this folder (`main.go`). Run it from anywhere in the repository:

```bash
go -C .claude/skills/newsletter run . <command>
```

It changes to the repository root itself, so file arguments are relative to the root. If `go` isn't installed in a Docker sandbox, install it with `sudo apt-get install -y golang-1.26`. Don't rewrite the helper in Python or another language.

Commands:

- `last`: the last issue's number and date. If `emails/` has no issues yet, this is #18, the last issue published as a blog post.
- `candidates`: JSON list of published posts dated after the last issue, oldest first, with word counts and whether each has a `<!--more-->` break. Posts tagged `newsletter` are excluded.
- `item <dir>`: one post as email-ready Markdown, the way the website shows it. If the post has a `<!--more-->` break, this is the text before the break plus "Read the full post →". Otherwise it is the whole post. It also adds the heading (or a date line for untitled notes) and the permalink, and turns `image`/`figure`/`youtube`/`video` shortcodes into plain Markdown with absolute URLs to the original image files, which email clients can show. Warnings about unhandled shortcodes go to stderr. Fix those by hand.
- `render <issue.md>`: writes `<issue>.html` next to the issue, with simple inline styles.

## Steps

1. **Find the posts.** Run `last` and `candidates`. Show the user a short numbered list with each post's date, title (or "note:" and the first line for untitled posts), word count, and whether it will be cut at a break. If there are no posts, say so and stop.

2. **Choose posts.** Ask about each post with AskUserQuestion, up to four questions per call. There are two options:
   - **As on the website** (Recommended): included the way `item` produces it, cut at `<!--more-->` if there is one, in full if not.
   - **Skip**: left out.

3. **Subject and preview text.**
   - The subject is `Working on It #<n>: <topic>`. Base the topic on the most substantive post included in the issue: usually the longest titled essay, not a short note or a linkpost. If it isn't obvious which post that is, say which one you picked.
   - Write the topic in sentence case, following the style guide in `CLAUDE.md`. Make it a short phrase in the spirit of past topics like "Behind, ahead", not just the post's title repeated.
   - Draft `preview_text`, one plain sentence of about 90–140 characters.
   - Show the user both and let them change them.

4. **Intro.** Ask whether the user wants to write an intro paragraph or a "Currently" list, like the one that ended #18 (**Teaching:** …, **Reading:** …). Use the user's wording verbatim. Don't make up personal news or a "Currently" list. If the user gives none, leave those sections out.

5. **Write the issue.** Copy `emails/_template.md` to `emails/YYYY-MM-DD-working-on-it-<n>.md`, using today's date in America/New_York. Then fill it in:
   - `{{number}}`, `{{topic}}`, and `{{preview_text}}`.
   - `{{date}}`: the current time, ISO 8601 with offset (`TZ=America/New_York date -Iseconds`). The next issue starts after this date, so use the time of writing.
   - `{{posts}}`: one YAML list item per included post directory, `  - <dir>`.
   - `{{items}}`: the output of `item` for each included post, oldest first, separated by blank lines.
   - `{{intro}}` and `{{currently}}`: the user's text, or delete the placeholder.
   - Remove any line left with only a placeholder. Keep everything else in the template, which the user maintains.

6. **Render** with `render emails/<issue>.md`. Check the stderr warnings, and spot-check the HTML: links absolute, images present, nothing left in `{{< >}}` or `{{ }}`.

7. **Hand off.** Don't paste the whole issue into the chat. Report:
   - The subject line and preview text, ready to copy.
   - The included posts.
   - The full path of the `.html` file. The repo is mounted from the user's Mac, so they can open it there.
   - How to send it: open the `.html` in a browser, select all, copy, and paste into a new Kit broadcast that uses a Classic template. Set the subject and preview text. Leave the broadcast **not public**: issues are not posted to Kit's web archive. Send a test email before sending to everyone.

8. **Commit.** Once the user says the issue is final, commit the `.md` file. The `.html` is ignored by git. Issues aren't published on the site, so commit on whatever branch the user is on unless they ask otherwise, and don't push `main` without asking. If the user edited the text in Kit after pasting, ask whether to bring those edits back into the `.md` so the archive matches what was sent.
