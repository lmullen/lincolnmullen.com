# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## About This Site

This website (<https://lincolnmullen.com>) is my professional, scholarly website. I am a professor of history.

## Technical Details

This site uses the static site generator, [Hugo](https://gohugo.io).

### Build Commands

- Build: `hugo --cleanDestinationDir --minify`
- Local preview: `hugo serve --port 2323 --buildDrafts --buildFuture --disableFastRender`
- Update the processed-image cache: `./gen-cache.sh` (see below)

Tool versions are pinned in `.mise.toml`, and `mise.lock` records each tool's download URL and checksum for every platform. `mise install` sets up a checkout; the GitHub Actions check installs with mise too; and Cloudflare's build command (`wrangler.jsonc`) runs `./cloudflare-build.sh`, which downloads the Hugo and Dart Sass pinned in `.mise.toml` from the URLs in `mise.lock` and verifies their checksums. Cloudflare doesn't use mise for Node, so `.node-version` repeats the Node pin for its build image. To upgrade Hugo, Dart Sass, or Node, change `.mise.toml` (and `.node-version` for Node), run `mise lock` (`mise install` also updates it), then run `./gen-cache.sh`, and commit all of it, including whatever the script changes in the cache, since cache keys can change between Hugo versions. In a cloud session without mise, install it with `npm install -g @jdxcode/mise` and run `mise install`.

### Sass

`layouts/_default/baseof.html` compiles `assets/scss/main.scss`, which imports Bootstrap from `node_modules`, with Dart Sass. Hugo has no built-in Sass compiler since 0.170.0: it runs the `sass` binary it finds on the `PATH`, so Dart Sass must be installed wherever the site is built (`mise install` locally and in CI; `cloudflare-build.sh` on Cloudflare). Dart Sass's deprecation warnings for `@import` and for Bootstrap's own files are silenced in the `css.Sass` options; remove the silencing when Bootstrap moves to `@use`.

### Processed-image cache

Hugo writes the WebP variants it makes from images to `resources/_gen/images`, and that directory is committed. Cloudflare Workers Builds starts every deploy from a fresh clone, so without the committed cache it re-encodes every image, which once took most of a four-minute build (issue #176). The cache must therefore hold exactly what a production build needs: every variant the live site uses, and nothing else.

`./gen-cache.sh` (also `make gen-cache`) keeps it that way. It builds the site with the production settings (no drafts, no future posts) against a copy of the cache, runs Hugo's garbage collection on the copy, and then adds whatever is missing from `resources/_gen/images` and deletes whatever is stale, printing both lists. `./gen-cache.sh --check` (`make check-gen-cache`) only reports and exits non-zero on drift, without touching the working tree; the "Processed-image cache" GitHub Actions workflow runs it on every pull request, so a pull request with an out-of-date cache fails its check.

Run `./gen-cache.sh` and commit the result (`git add -A resources/_gen/images`) in the same commit as any change that alters which images get processed or how:

- adding, replacing, or removing an image in a post or page, or changing a post's `image:` front matter (OpenGraph previews are processed images too);
- changing `imaging` in `config.yaml`, `layouts/partials/responsive-image.html`, the `image`, `figure`, and `sideimage` shortcodes, or any other template that calls `.Resize`, `.Fill`, `.Fit`, or `.Process`;
- upgrading Hugo.

The script refuses to run unless `hugo` is the version pinned in `.mise.toml`, `mise.lock` records the same version, and `.node-version` matches the Node pin, because a cache built with a different Hugo may use different keys and go unused on Cloudflare. It needs `node_modules` (`npm ci`) and Dart Sass (`sass`, installed by `mise install`), since the build compiles the SCSS. Don't commit cache files written by `hugo serve --buildDrafts --buildFuture`: variants for drafts are not needed in production, and the script removes them. A post dated in the future is likewise left out until its date passes; the first deploy after that encodes its images, and the next `./gen-cache.sh` run adds them.

### Creating New Content

```bash
./new-post.sh   # Interactive: prompts for title and optional image
```

- **Titled post**: Enter a title when prompted. Creates `content/blog/YYYY-MM-DD-slug/index.md`.
- **Untitled post** (short-form/micro): Leave the title blank. Creates `content/blog/YYYY-MM-DD-HH-MM-SS/index.md` with the timestamp as the slug.
- Both types optionally include an image (drag-and-drop friendly).

The script is for my own use. For any work on a post (creating it, turning a draft or an Obsidian note into one, reviewing, tagging, or adding social media text), follow the `new-post` skill (`.claude/skills/new-post/SKILL.md`), which does not run the script. A post always reaches the site through a pull request; never commit or push a post to `main`.

For courses, manually create `content/courses/[course-code].[year]/index.md` using the archetype at `archetypes/courses.md`.

## Additional Directories (Docker Sandbox Only)

When running in a Docker sandbox, these directories may be mounted as additional workspaces. Outside a sandbox they will not be available, so do not assume they exist, and do not try to find or recreate them.

- `/Users/lmullen/github/lmullen/projects.lincolnmullen.com`: Standalone projects served off a subdomain of this site, separate from the Hugo site.
- `/Users/lmullen/github/lmullen/domesday`: My custom web analytics, served from Cloudflare. It has its own `CLAUDE.md`; read it before working in that repository.
- `/Users/lmullen/notes/writing/blog`: Drafts that I write in Obsidian on the go. These are source material for new blog posts, not part of the site. Do not edit or delete them unless I ask, with one exception: when a note becomes a post, the `new-post` skill sets the note's `writing-status` to `completed`.

To check whether you are in the sandbox, test whether these directories exist before relying on them.

## Content Types

Content types have layout files in `./layouts/` and corresponding content in `./content/`. The most important content types:

- `blog`: All blog posts, both titled long-form articles and untitled short-form notes. Every post is a page bundle (`content/blog/<dir>/index.md`). Templates use title presence to distinguish: titled posts appear in the archive and get headings; untitled posts do not.

The blog's title is *Working On It*, after the former newsletter. Spell it with a capital "On" everywhere: in prose, post titles, citations, feeds, and social text. The title comes from `blogFeedTitle` in `config.yaml`, which supplies the label heading the home page's stream, the RSS and JSON feed titles, the `rel="alternate"` links, the suggested citations, and the Zotero "Blog" field; change it there rather than hard-coding it. The nav item and the featured and archive page titles still say "blog," and the subscribe bars say "Subscribe."

The site has no separate newsletter. Issues of the former newsletter, *Working On It*, are ordinary blog posts titled "Working On It #<n>: …" and tagged `newsletter`; their `feedid` keeps the old `/newsletter/<n>/` feed GUID. Email signup offers new posts bundled into an occasional digest: every "Email" link opens a modal (`layouts/partials/subscribe-modal.html`) whose form posts to Kit (`emailSignupAction` in `config.yaml`).

Since #19, *Working On It* is an occasional digest of new blog posts, sent by email: it has no content of its own. Issues live in `digests/`, which Hugo does not publish. Call it a digest, not a newsletter, and don't promise a cadence (weekly, monthly). To make an issue, follow the `digest` skill (`.claude/skills/digest/SKILL.md`). It asks which recent posts to include (each appears as it does on the website), writes the issue, and renders HTML for the user to paste into an HTML block in a Kit broadcast.

### Front Matter Patterns

**Blog posts** (`content/blog/YYYY-MM-DD-slug/index.md`):
- Required: `date`, `slug`, `draft`
- Title: Optional. When present, the post appears in the archive and gets a heading. When absent, the post is a short-form note.
- Optional: `description`, `crosspost: {url, source}`, `linkpost: "url"`, `via: "url"`, `image: "filename.jpg"`, `social: {bluesky, mastodon, linkedin}`, `tags`
- Images: Store in the page bundle directory and reference with `{{</* image src="filename.jpg" alt="description" */>}}`; the `image:` front matter field (relative filename) enables automatic OpenGraph image previews.

Use `<!--more-->` in blog posts for "Read More" truncation on list pages.

When working on a post, do not paste the full Markdown file into the chat unless the user asks for it. Summarize the changes and quote only the relevant lines.

## Shortcodes

Use the `figure` shortcode for images with a caption, the `image` shortcode for images that do not need a caption, and the `sideimage` shortcode for an illustrative image floated beside the text.

- `{{</* image src="file.jpg" alt="description" */>}}` - Responsive images with WebP conversion
- `{{</* image src="file.gif" */>}}` - GIFs and SVGs pass through without processing
- `{{</* figure src="file.jpg" caption="Markdown caption" */>}}` - Numbered figure, breaking out to the wide column; `link`, `alt`, `class`, and `width` (1x pixels) are optional
- `{{</* sideimage src="file.jpg" alt="description" caption="Optional caption" */>}}` - Small image floated right, half in the text column and half in the gutter, with an unnumbered caption; `link` and `side="left"` are optional. Works in both blog posts and pages.

Image `src` values are page-bundle filenames, or global assets with a leading slash (`/img/photo.jpeg`).

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

Include the permalink URL at the end of each social media text, or the `linkpost` URL when the post has one.

## Git and GitHub

The rules for working with Git and GitHub are the same in all my repositories and are not repeated here: the `lmullen-agent` identity, an issue for every change, the project board, branches and worktrees, pull requests, merging, and cleanup after a merge. In a Docker sandbox they come from the lmullen-agent kit, whose context file you must read before any Git or GitHub operation. In a Claude Code cloud session they are in `~/.claude/CLAUDE.md`. Their source is [`lmullen-agent-context.md`](https://github.com/lmullen/dotfiles/blob/main/sandboxes/kits/lmullen-agent/lmullen-agent-context.md) in `lmullen/dotfiles`; change a rule there, not here.

This section holds only the facts about this repository.

### Project board

- The board is <https://github.com/users/lmullen/projects/4> (owner `lmullen`, project number `4`). It spans all of my repositories.
- Its `Status` columns are `Triage`, `TODO`, `In progress`, and `Done`. There is no review column: an issue stays in `In progress` while its pull request is open.
- Issues opened in this repository are added to the board automatically, in `Triage`, so there is no need to run `gh project item-add`. You still need the item id for status moves.

### Branches

- Push the branch after every change, so Cloudflare builds a preview; see "Preview servers" below.
- Branch names must start with a letter, with any issue number at the end (`speed-up-builds-176`, not `176-speed-up-builds`). Cloudflare cannot make a branch preview URL from a name that starts with a digit.

### Preview servers

Every branch gets two previews: a local Hugo server and a Cloudflare branch preview. Whenever you make changes on a branch, give the user both URLs.

**Local preview.** Ports 1313, 1314, 1315, and 1316 are published from the sandbox to the host. When you create a branch, start a Hugo server for it on one of those ports:

- Pick the first free port: `(exec 3<>/dev/tcp/127.0.0.1/<port>) 2>/dev/null` succeeds if the port is taken. If all four are taken, ask the user which server to stop.
- Run it from inside the branch's worktree, in the background, bound to all interfaces so the published port can reach it:
  `hugo serve --port <port> --bind 0.0.0.0 --baseURL http://localhost:<port>/ --buildDrafts --buildFuture --disableFastRender --poll 500ms`
- Keep `--poll 500ms`. The checkout is mounted from the host over virtiofs, which does not deliver file-change events for edits made on the host, so without polling Hugo misses the user's edits and live reload never fires.
- Run `npm ci` in the worktree first; the SCSS imports Bootstrap from `node_modules`.
- If `hugo` or `sass` is not on the `PATH`, run `mise install` in the worktree (or checkout) to get the Hugo and Dart Sass versions pinned in `.mise.toml`, and run Hugo through mise (`mise exec -- hugo ...`) if it is still not found. Keep `mise.lock` (and `.node-version`) in step with `.mise.toml`.
- The local preview URL is `http://localhost:<port>/`. Hugo rebuilds on file changes, so the server keeps working as you edit; restart it only if it stops.
- Keep one server per branch. Stop a branch's server before removing its worktree.

**Remote preview.** Cloudflare Workers Builds deploys every pushed branch, so push the branch after each change (`git push -u origin <branch>` the first time). Each build gets two preview URLs:

- **Commit preview**, for that build only: `https://<version-id>-lincolnmullen-dot-com.lincoln-07e.workers.dev/`. Every successful build has one.
- **Branch preview**, which follows the branch's latest build: `https://<branch>-lincolnmullen-dot-com.lincoln-07e.workers.dev/`, where `<branch>` is the branch name, lowercased, with characters other than letters, digits, and `-` replaced by `-`. Cloudflare creates it only when it can turn the branch name into an alias. When it can't, the build log warns `Preview alias generation requested, but could not be autogenerated` and only the commit preview exists. Branch names that start with a digit (e.g. `176-speed-up-builds`) fail this way, so start branch names with a letter and put any issue number at the end (`speed-up-builds-176`).

Don't build these URLs from the pattern and report them. Read them from the build: the log prints `Version Preview URL` (the commit preview) and, when the alias exists, `Version Preview Alias URL` (the branch preview). Once a PR is open, Cloudflare's bot comment on the PR shows the same URLs and the build status. The URLs work only after the build finishes (`gh pr checks <n>`). The sandbox firewall may block `*.workers.dev`, so don't treat a failed fetch from the sandbox as a failed deploy.

In the sandbox, `CLOUDFLARE_API_TOKEN` can read builds and their logs through the Cloudflare API (account `07e6662687bbf114d240a98d738e4cc8`, worker `lincolnmullen-dot-com`, script tag `3d0c2a242d7840cda4e81fd7a219d999`):

- List builds: `GET https://api.cloudflare.com/client/v4/accounts/<account>/builds/workers/<script-tag>/builds?per_page=10`. Each result has `build_uuid`, `build_trigger_metadata.branch` and `.commit_hash`, `status`, `build_outcome`, and `created_on` / `running_on` / `stopped_on`.
- Read a build's log: `GET https://api.cloudflare.com/client/v4/accounts/<account>/builds/builds/<build_uuid>/logs`. `result.lines` is a list of `[timestamp_ms, text]` pairs. The endpoint can briefly return a non-JSON response right after a build stops; retry.
- Measure total build time as `stopped_on` minus `created_on`. The gap between `created_on` and `running_on` combines two things: waiting in the queue behind another build of this worker (builds run one at a time) and initializing the build environment. The log separates them: when a build was queued, `Initializing build environment...` appears twice, and the second one marks the end of the wait.

**Status line.** `.claude/statusline.sh` (registered in `.claude/settings.json`) prints the sandbox kit's status line and then `preview: local remote`, linking the Hugo server running from the current checkout and its Cloudflare branch preview. It finds the server by matching a `hugo` process's working directory to the checkout, so start servers from inside the worktree as above.
