---
name: new-post
description: Take a blog post (titled or untitled short-form) from wherever it starts to a pull request, covering the page bundle, images, review, tags, description, and social media text. Use when the user asks to create, write, add, publish, review, tag, or finish a blog post or micro post, or to turn an Obsidian note or a draft into a post.
---

# New post

Take a post from its source to a pull request that the user can check on a preview and merge. This one skill covers the whole job: the page bundle, images, review, tags, description, social media text, the build check, and the pull request.

Three rules hold throughout.

- **Every post ends as a pull request.** Never commit or push a post to `main`, wherever it came from and however small it is.
- **The words are the user's.** Fix obvious errors, but do not reword the post, and do not write the description or social media text yourself.
- **Work straight through.** Do the steps in order without stopping to ask permission for each one. Pause only where a step below says to: for an actual error you cannot resolve, or for text that needs the user's judgment. When you pause, ask everything you have at once, then carry on.

When the user asks for only part of this on a post that already has a branch (for example, "tag this post" or "review this post"), do just that step, on that branch.

## 1. Find the source

A post starts in one of three places. If the user doesn't say which post, and more than one candidate exists, ask.

- **A bundle in `content/blog/`.** Usually an untracked directory in the main checkout with `draft: true`, which the user wrote in their editor. `git status --short content/blog` in the main checkout lists them.
- **A note in Obsidian**, in `/Users/lmullen/notes/writing/blog`. That folder is mounted only in the Docker sandbox; if it is absent, say so rather than looking elsewhere. The filename, without `.md`, is the working title. The note's own front matter (`writing-status` and the other `writing-*` fields) belongs to Obsidian and does not go into the post. Notes whose status is `completed` are already posts.
- **The prompt.** Text the user types or a file they provide. When the user gives exact wording, use it verbatim.

Read the whole source before going on. If the source is plainly unfinished (an outline, a stub of a sentence or two under a long-form title, notes to self, `TODO` markers), pause and ask whether to go ahead.

## 2. Issue, worktree, and branch

Follow the Git and GitHub rules (the lmullen-agent kit, and the "Git and GitHub" section of `CLAUDE.md`). What is specific to posts:

- Find or create an issue for the post, and make sure it carries the `post` label (`gh issue create --label post ...`, or `gh issue edit <n> --add-label post` for an existing issue). Move it to `In progress` on the board.
- Name the branch `draft-post-<short-slug>-<issue number>`, where `<short-slug>` is a short, lowercase, hyphenated description of the post (e.g. `draft-post-zelda-on-paper-212`). The branch name becomes part of the preview URL, so keep it short.
- Work in a worktree based on the current `origin/main`, and start the local Hugo server there as `CLAUDE.md` describes.

## 3. Build the page bundle

Do not run `./new-post.sh`; it is the user's own tool and ends by opening VS Code. Create the bundle in the worktree by hand, using `archetypes/blog/index.md` as the model for the front matter.

- **Untitled post**: the directory and `slug` are the timestamp, `YYYY-MM-DD-HH-MM-SS`. Omit the `title` field entirely; do not leave it as `""`.
- **Titled post**: the directory is `YYYY-MM-DD-<slug>`, where the slug is the lowercased, hyphenated title with punctuation removed.
- `date` is the current time: ISO 8601 with the UTC offset, e.g. `'2026-10-25T09:01:22-04:00'`, in America/New_York (`TZ=America/New_York date +%Y-%m-%dT%H:%M:%S%:z`). It is set again before every commit; see step 8.
- Set `draft: false`. The branch preview and the pull request are the draft stage. Mention this in the handoff.
- Remove the archetype's commented-out fields that the post doesn't use.

By source:

- **Existing bundle**: move it out of the main checkout and into the worktree (`mv`, from the main checkout's `content/blog/<dir>` to the worktree's), so that nothing is left behind to collide with `main` after the merge. Then bring it up to date, since the bundle may have sat for a while: set `date` to now, and rename the directory to match (the timestamp, with the same change to `slug`, for an untitled post; the `YYYY-MM-DD` prefix for a titled one).
- **Obsidian note**: copy the body into a new bundle dated now. Do not change the note yet (see step 9). Convert Obsidian syntax to what the site uses: `![[image.png]]` embeds become shortcodes (find the image in the vault and copy it into the bundle), and `[[wikilinks]]` become plain text or, if they point at something public, ordinary links. If an embed or link can't be resolved, pause and ask.
- **Prompt**: write the text into a new bundle dated now.

### Images

- Store each image in the page bundle, with a short, hyphenated filename that conveys the content (e.g. `zelda-paper-map.jpg`), not just a timestamp.
- Set `image:` in the front matter to the filename of the main image, so link previews show it. A post whose body has an image must have this field.
- Use the `image` shortcode, `figure` if the image needs a caption, or `sideimage` for a small illustrative image; never raw Markdown image syntax. Look at the image and write descriptive alt text if the user hasn't.
- Put the body text above the image unless the user says otherwise.

## 4. Review

Check the post, fix what is obviously wrong, and keep a list of what you changed for the handoff.

**Fix without asking:**

- Spelling errors, typos, doubled words, and plainly broken grammar or punctuation.
- A title not in sentence case. The title rules are in the "Style Guide" section of `CLAUDE.md`.
- Titles of works (books, journals, games, websites, etc.) that are not italicized in the body.
- Front matter: `date` in ISO 8601 with offset; `slug` lowercase letters, digits, and hyphens only, matching the directory; `draft: false`; `tags` a YAML list; `crosspost` with both `url` and `source`; `linkpost` and `via` valid URLs.
- Raw Markdown images converted to shortcodes; `image:` set when the body has an image.

**Do not reword.** Leave the user's sentences, word choices, and structure alone, even where you would write them differently. You may suggest rewordings in the handoff, quoting the sentence and your alternative, but do not apply them without being told to.

**Pause and ask** when:

- a passage is garbled enough that the fix is a guess at what the user meant;
- a fact, name, date, quotation, or link looks wrong (check links that you can);
- an image the post refers to is missing;
- a style question comes up that the rules don't cover (a preferred spelling, a hyphenation). After the user answers, add the rule to the "Style Guide" in `CLAUDE.md`.

For a post of roughly 500 words or more with no `<!--more-->`, suggest a place for one in the handoff rather than adding it.

## 5. Tags

- **Titled post**: choose one to five tags from `tags.md` in this directory, following its judgment calls, and apply them. A general and a specific tag together are fine, and so is leaving a post untagged when nothing fits.
- **Untitled post**: leave untagged, unless the content clearly fits a tag.
- If the post needs a tag that is not in the taxonomy, pause and ask. If the user approves it, add it to `tags.md`.

Use the non-indented list format:

```yaml
tags:
- digital history
- American religious history
```

## 6. Description and social media text

Do not draft this text. It comes from the user, in this order of preference:

1. **Text the user supplied** for a field, in the source's front matter or in the prompt. Use it as given. If they supplied text for one platform only, use it for the others too, as long as it fits.
2. **The description.** If the post has a `description`, the social text for each platform is the description.
3. **The body of a short untitled post**, when the whole body fits within the platform's limit. Use it verbatim (as plain text).
4. **The title.** If the post has a title but no description, the social text is the title.
5. **Nothing to go on** (an untitled post, too long to quote, with no description): ask the user to write the social text, and a description if they want one, then fill in what they give you. Don't hold up the rest of the work for this. Finish through the pull request with the fields empty, ask in the handoff, and push the text to the branch when it comes.

Never write a `description` yourself. If it is missing, leave it empty and say so in the handoff, since it is what link previews and search results show.

Each of `social.bluesky`, `social.mastodon`, and `social.linkedin` is the text, a space, and then a link:

- the post's permalink, `https://lincolnmullen.com/blog/<slug>/`;
- or, when the post has a `linkpost`, that URL instead.

Rules:

- Plain text only: strip Markdown such as italics from the text. No hashtags.
- `bluesky` is at most 300 characters and `mastodon` at most 500, including the link. Do not guess: count with `printf '%s' "<text>" | wc -m`. If the text is too long, do not trim or rephrase it; pause and ask the user for a shorter version.
- If the user says to skip a platform, set its field to `""`. A post with an empty `bluesky` or `mastodon` field is left out of that platform's feed.
- Use double-quoted strings on one line. For multi-paragraph `linkedin` text, use a literal block scalar (`|`).

```yaml
description: "Why narrative visual essays work better than exploratory visualizations."
social:
  bluesky: "Why narrative visual essays work better than exploratory visualizations. https://lincolnmullen.com/blog/example-post/"
  mastodon: "Why narrative visual essays work better than exploratory visualizations. https://lincolnmullen.com/blog/example-post/"
  linkedin: "Why narrative visual essays work better than exploratory visualizations. https://lincolnmullen.com/blog/example-post/"
```

## 7. Build check

Run `./gen-cache.sh` from the worktree. It builds the site with the production settings, so it catches template and front matter errors, and it adds the post's image variants to the processed-image cache (`resources/_gen/images`) and removes stale ones. It refuses to run if `hugo` is not the version pinned in `.mise.toml`; see "Processed-image cache" in `CLAUDE.md`.

Then look at the post on the local server to confirm it renders: the title or its absence, images, and links. A build failure is an actual error: fix it if the cause is in the post, and otherwise pause and report it.

## 8. Commit, push, and pull request

- **Update the date before every commit.** Each time you commit to the post's branch, the first commit and every later one (review fixes, social text, changes the user asks for), first set the post's `date` to the current time, in the format from step 3. The post should carry the time it was last worked on, so that it is current whenever it is merged. Change only `date`: leave the directory name and `slug` as they were when the bundle was built, so that the permalink, the social media links, and the preview URL stay the same from one commit to the next. If the user has asked for a particular date, keep theirs.
- **Keep the slug in step with the title.** The exception to the rule above: if the title of a post that has not yet been published changes, rename the slug and the directory to match the new title without asking, update the permalink in the social media text, and run `./gen-cache.sh`, since the cache paths include the slug. Once a post is published, never change its slug.
- Commit the bundle together with the cache changes (`git add -A resources/_gen/images`), plus any change to `tags.md` or `CLAUDE.md` from the steps above.
- Push the branch, so Cloudflare builds a preview.
- Open the pull request **ready for review, not as a draft**, with `lmullen` as the reviewer and `Closes #<n>` in the body. In the body, list the fixes you made in step 4, so the user can see every change to their text.
- Never merge it yourself unless the user asks.

## 9. Close out the source

After the pull request is open:

- **Obsidian note**: never delete, move, or rewrite the note. Change only its `writing-status` front matter field to `completed` (adding the field if the note has none), and leave the body and the other fields as they are.
- **Existing bundle**: confirm that the directory is gone from the main checkout.

## 10. Handoff

Do not paste the full Markdown file into the reply unless the user asks for it. Report:

- the pull request and issue URLs, and the reviewer;
- the preview URLs **for the post itself**: the local server and the Cloudflare branch preview, each followed by `/blog/<slug>/` (read the Cloudflare URL from the build, as `CLAUDE.md` says);
- that `draft` is `false`, so merging publishes the post;
- every fix you made to the text, quoting the before and after;
- suggestions you did not apply: rewordings, a place for `<!--more-->`, a missing description;
- which text each social field uses, and anything you still need from the user, such as social text to write.
