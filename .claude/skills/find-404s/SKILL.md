---
name: find-404s
description: Find recent 404s on lincolnmullen.com in the Domesday analytics database and add redirects for the ones that have a sensible destination, as a pull request. Use when the user asks to check for 404s, broken links, missing pages, or redirects that need adding.
---

# Find 404s

Domesday, the site's own analytics (`/Users/lmullen/github/lmullen/domesday` in the sandbox), records every view of the 404 page. This skill pulls those views, drops the ones the site already handles, and adds rules to `static/_redirects` for the rest that have a real destination. The work goes to the user as a pull request, like any other change.

## 1. Get the 404s from Domesday

Domesday's data is in the Cloudflare D1 database `domesday` (id `86193ef2-c373-47eb-8a93-d4b52ecf62b8`). Query it with the Cloudflare connector's `d1_database_query` tool. The sandbox's `CLOUDFLARE_API_TOKEN` can't read D1. Run only `SELECT` statements: the connector can also write.

```sql
SELECT path, MIN(url) AS url, COUNT(*) AS views,
       COUNT(DISTINCT visitor_hash) AS visitors,
       MIN(timestamp) AS first, MAX(timestamp) AS latest,
       GROUP_CONCAT(DISTINCT referrer_hostname) AS refs
FROM page_views
WHERE hostname = 'lincolnmullen.com'
  AND (status_code = 404 OR title = '404 Page not found | Lincoln Mullen')
  AND timestamp >= '<since, e.g. the date of the last run>'
GROUP BY path
ORDER BY views DESC
```

- The `title` condition catches 404s recorded between July and November 2026, when a template change dropped the page's `data-status="404"` and they were stored with no status code. If the 404 page's title changes, update it here.
- Domesday stores `path` lowercased, without the query string or trailing slash. `url` keeps the original, so use it for case and trailing slash. Redirect rules match exactly, so `/files/Foo.pdf` is not `/files/foo.pdf`.
- The result is usually too large to show inline. The tool saves it to a file; copy that file to `$CLAUDE_JOB_DIR/tmp/404s.json` (or another temporary directory) and work from there.

## 2. Drop what the site already handles

Build the site with the production settings and run the helper:

```bash
npm ci
mise exec -- hugo --minify --cleanDestinationDir
python3 -I .claude/skills/find-404s/find_404s.py classify $CLAUDE_JOB_DIR/tmp/404s.json
```

`classify` removes paths that the built `public/` now serves or that a rule in `static/_redirects` matches. It then hides noise: vulnerability scanners (`wp-`, `.php`), old WordPress archive, tag, and date pages, retired theme assets, and Hugo image variants (`_hu…`), whose names change with every Hugo release. It prints the rest, most-viewed first, and writes them to `404s.open.json`. Add `--all` to see the noise too.

## 3. Decide where each one should go

Work down the list. Redirect a path only when a page or file exists that the visitor was clearly looking for. Leave the rest as 404s, and don't redirect them to the home page just to make them go away.

- **Blog posts at old slugs.** Old URLs are often shortened or older slugs of current posts. Compare them with the `slug:` values in `content/blog/*/index.md`; `difflib.get_close_matches` finds most. Confirm each match by its title. Untitled posts were once addressed by a 12-hour timestamp: `/blog/2022-10-20-05-43-16` is now `/blog/2022-10-20-17-43-16`.
- **Files.** PDFs, syllabi, the CV, and datasets moved to R2 at `files.lincolnmullen.com`, but under different folders (`pdf/`, `syllabi/`, `cv/`, `history-datasets/`), and many old files were never moved. `probe` looks for each file there and prints a rule for each one it finds:

  ```bash
  python3 -I .claude/skills/find-404s/find_404s.py probe $CLAUDE_JOB_DIR/tmp/404s.open.json
  ```

  Don't add a blanket `/files/* https://files.lincolnmullen.com/:splat` rule. Most of those paths don't exist on R2, so it would trade a 404 here for one there.
- **Projects and software.** Static project sites live on `projects.lincolnmullen.com`; R packages, on `docs.ropensci.org` or GitHub. Check that a destination answers before you point at it (`curl -sI`).
- **Typos and mangled links** (`/pojects/slavery`, `/blog/isnt-`, encoded junk): leave them, unless one path gets steady traffic from a single bad link, like `/https%3A/lincolnmullen.com/`.

When a choice is unclear (a post that might be the one meant, or a project that is offline), ask the user instead of guessing.

## 4. Write the rules

Add the rules to `static/_redirects`, following the comments at the top of the file:

- **Exact paths in the top half, wildcards (`*`, `:name`) at the bottom.** Cloudflare counts every rule after the first wildcard as dynamic, even an exact one, and silently ignores everything after the 100th dynamic rule. `wrangler dev` reports `Maximum number of dynamic rules supported is 100` when that happens.
- **Write each exact rule both with and without the trailing slash.** A rule matches only the path it names, and `/x/*` doesn't match `/x`.
- An exact rule must come before any wildcard that would also match its path.
- Use 301 for permanent moves and 302 for anything that may change.

## 5. Check the rules against Cloudflare's own matcher

Rebuild, then serve the site the way Cloudflare does:

```bash
mise exec -- hugo --minify --cleanDestinationDir
npx wrangler dev --port 8799 --ip 127.0.0.1   # in the background
```

Wait for `Ready on`. Then confirm that the log says `Parsed N valid redirect rules` with no warning, and run:

```bash
python3 -I .claude/skills/find-404s/find_404s.py verify $CLAUDE_JOB_DIR/tmp/404s.open.json
```

Every rule you added should show up as a redirect. Also spot-check a few existing rules (`/micro/x`, `/projects/slavery/x`, `/cv/`) to make sure the new rules haven't shadowed them. Stop `wrangler dev` when you're done.

## 6. Open the pull request

Follow the usual Git and GitHub rules: an issue, a branch from `origin/main`, and a PR with `lmullen` as reviewer. In the PR body, report:

- the date range checked, and the number of 404 paths and views;
- each group of rules added and the views they cover;
- the paths left as 404s, grouped by kind (noise, files never moved to R2, posts that don't exist), so the user can decide whether any need more.
