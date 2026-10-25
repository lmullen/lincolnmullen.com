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

## Additional Directories (Docker Sandbox Only)

When running in a Docker sandbox, these directories may be mounted as additional workspaces. Outside a sandbox they will not be available, so do not assume they exist, and do not try to find or recreate them.

- `/Users/lmullen/github/lmullen/projects.lincolnmullen.com`: Standalone projects served off a subdomain of this site, separate from the Hugo site.
- `/Users/lmullen/github/lmullen/domesday`: My custom web analytics, served from Cloudflare. It has its own `CLAUDE.md`; read it before working in that repository.
- `/Users/lmullen/notes/writing/blog-drafts`: Drafts that I write in Obsidian on the go. These are source material for new blog posts, not part of the site. Do not edit or delete them unless I ask.

To check whether you are in the sandbox, test whether these directories exist before relying on them.

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

## Git and GitHub

### Branches and worktrees

Whenever the user asks you to work on a separate branch, use a git worktree instead of switching branches in the current checkout.

- Never run `git switch` or `git checkout <branch>` in the main working directory. Leave it on whatever branch it's on, with any uncommitted changes untouched.
- Create worktrees inside the project at `.worktrees/<branch-name>` (replace `/` in branch names with `-`).
- Before creating the first worktree, make sure `.worktrees/` is listed in `.gitignore`. If it isn't, add it and tell the user.
- For a new branch, base it on `main` unless the user says otherwise:
  `git worktree add -b <branch> .worktrees/<branch> main`
- For an existing branch: `git worktree add .worktrees/<branch> <branch>`
- If a worktree for that branch already exists (check `git worktree list`), reuse it rather than creating another.
- After creating a worktree, do all reading, editing, building, and testing from inside that directory. Install dependencies there if needed (e.g., `go mod download`, `npm install`).
- Tell the user the worktree path when you create it.
- Commit work in the worktree's branch. Don't merge into `main` unless the user asks. Push the branch (never `main`) whenever you make changes so Cloudflare builds a preview; see "Preview servers" below.
- Don't remove a worktree or delete its branch unless the user asks. When they do, use `git worktree remove <path>` followed by `git branch -d <branch>`, and run `git worktree prune` if needed.
- When the user needs to run a command that must run against the branch's code (a migration, a build, a database command), give the command to switch to its worktree (`cd .worktrees/<branch-name>` from the project root) before the command itself, so it runs against the branch rather than whatever the main checkout has checked out.

### Preview servers

Every branch gets two previews: a local Hugo server and a Cloudflare branch preview. Whenever you make changes on a branch, give the user both URLs.

**Local preview.** Ports 1313, 1314, 1315, and 1316 are published from the sandbox to the host. When you create a branch, start a Hugo server for it on one of those ports:

- Pick the first free port: `(exec 3<>/dev/tcp/127.0.0.1/<port>) 2>/dev/null` succeeds if the port is taken. If all four are taken, ask the user which server to stop.
- Run it from inside the branch's worktree, in the background, bound to all interfaces so the published port can reach it:
  `hugo serve --port <port> --bind 0.0.0.0 --baseURL http://localhost:<port>/ --buildDrafts --buildFuture --disableFastRender`
- Run `npm ci` in the worktree first; the SCSS imports Bootstrap from `node_modules`.
- If `hugo` is not on the `PATH`, install the version pinned in `wrangler.jsonc` into `~/.local/bin`, choosing `linux-arm64` or `linux-amd64` to match `uname -m`.
- The local preview URL is `http://localhost:<port>/`. Hugo rebuilds on file changes, so the server keeps working as you edit; restart it only if it stops.
- Keep one server per branch. Stop a branch's server before removing its worktree.

**Remote preview.** Cloudflare Workers Builds deploys every pushed branch, so push the branch after each change (`git push -u origin <branch>` the first time). The branch preview URL is:

`https://<branch>-lincolnmullen-dot-com.lincoln-07e.workers.dev/`

`<branch>` is the branch name, lowercased, with characters other than letters, digits, and `-` replaced by `-`. Once a PR is open, Cloudflare's bot comment on the PR shows the exact Branch Preview URL and build status; prefer that if it differs. The URL works only after the build finishes (`gh pr checks <n>`). The sandbox firewall may block `*.workers.dev`, so don't treat a failed fetch from the sandbox as a failed deploy.

### GitHub

Use the `gh` CLI for GitHub operations — creating pull requests, viewing issues, checking CI status, and commenting.

- Open a PR: `gh pr create --base main --head <branch> --title "..." --body "..."`
- Reference the issue a PR closes in the commit or PR body with `Closes #<n>` so it auto-closes on merge.
- View an issue: `gh issue view <n>`
- Check CI on a PR: `gh pr checks <n>`

Do not hand-build GitHub API calls with `curl`; prefer `gh`.
