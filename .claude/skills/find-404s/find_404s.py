#!/usr/bin/env python3
"""Helper for the find-404s skill. Standard library only.

  classify <rows.json> [--public DIR] [--redirects FILE] [--all]
      Drop the 404 paths that a built site (or a rule in _redirects) now
      handles, filter out scanner and crawler noise, and print what is left,
      most-viewed first. Writes the remaining rows to <rows>.open.json.

  probe <rows.json>
      For 404s that look like files, look for the file on
      files.lincolnmullen.com (R2) under the folders it was moved to, and print
      a ready-made redirect rule for each one found.

  verify <rows.json> [--base URL]
      Request every path from a local `wrangler dev` server (default
      http://127.0.0.1:8799) with and without a trailing slash, and report
      which now redirect and which still return 404.

<rows.json> is the result of the D1 query in SKILL.md, as the Cloudflare
connector returns it (a list holding {"results": [...]}) or a bare list of rows.
"""
import argparse, concurrent.futures, json, os, re, sys, urllib.error, urllib.parse, urllib.request

ROOT = os.path.abspath(os.path.join(os.path.dirname(__file__), "..", "..", ".."))
R2 = "https://files.lincolnmullen.com/"
R2_FOLDERS = ("pdf/", "syllabi/", "cv/", "history-datasets/", "figures/", "downloads/", "files/")

# Requests that no redirect should answer: vulnerability scanners, old
# WordPress archive pages, assets from retired themes, and Hugo image variants
# whose hashed names change with every Hugo release.
NOISE = re.compile(
    r"wp-|\.php|xmlrpc|/cgi-bin|\.env|\.git/|/admin|^/contact|^/galleries|"
    r"^/(archives?|tag|(19|20)\d\d)(/|$)|/feed$|comment-page|"
    r"^/(assets|css|js|fonts|images|img)/|\.(css|js|woff2?|ttf|eot|swf)$|"
    r"_hu[0-9a-f_]",
    re.I,
)


def load_rows(path):
    data = json.load(open(path))
    if isinstance(data, list) and data and isinstance(data[0], dict) and "results" in data[0]:
        data = data[0]["results"]
    return data


def compile_rule(src):
    pat = ""
    for part in re.split(r"(\*|:[A-Za-z]\w*)", src):
        if part == "*":
            pat += "(.*)"
        elif part.startswith(":"):
            pat += "([^/]+)"
        else:
            pat += re.escape(part)
    return re.compile("^" + pat + "$", re.I)


def classify(args):
    rows = load_rows(args.rows)
    files = set()
    for root, _, names in os.walk(args.public):
        for n in names:
            files.add(("/" + os.path.relpath(os.path.join(root, n), args.public)).lower())
    rules = []
    for line in open(args.redirects):
        line = line.strip()
        if line and not line.startswith("#"):
            rules.append(compile_rule(line.split()[0]))

    def handled(p):
        bare = urllib.parse.unquote(p).lower().rstrip("/")
        if bare == "" or any(c in files for c in (bare, bare + "/index.html", bare + ".html")):
            return True
        return any(rx.match(c) for c in (p, p.rstrip("/") + "/") for rx in rules)

    open_rows = [r for r in rows if not handled(r["path"])]
    shown = open_rows if args.all else [r for r in open_rows if not NOISE.search(r["path"])]
    out = os.path.splitext(args.rows)[0] + ".open.json"
    json.dump(open_rows, open(out, "w"))
    print(f"{len(rows)} paths ({sum(r['views'] for r in rows)} views) in the data; "
          f"{len(open_rows)} ({sum(r['views'] for r in open_rows)} views) still 404; "
          f"{len(shown)} shown after filtering noise. Remaining rows written to {out}\n")
    for r in sorted(shown, key=lambda r: (-r["views"], r["path"])):
        print(f"{r['views']:4} {r.get('latest', '')[:10]}  {r['path'][:120]}  {r.get('refs') or ''}")


def status(url):
    req = urllib.request.Request(url, method="HEAD")
    try:
        return urllib.request.urlopen(req, timeout=15).status
    except urllib.error.HTTPError as e:
        return e.code
    except Exception:
        return 0


def probe(args):
    rows = [r for r in load_rows(args.rows)
            if re.search(r"\.(pdf|zip|png|jpe?g|gif|csv|txt|docx?)$", r["path"], re.I)
            and "_hu" not in r["path"] and r["path"] != "/robots.txt"]  # R2 has its own robots.txt
    # Domesday lowercases paths, so try the original URL's path when it has one.
    jobs = {}
    with concurrent.futures.ThreadPoolExecutor(16) as ex:
        for r in rows:
            src = urllib.parse.urlparse(r["url"]).path if r.get("url") else r["path"]
            rel = urllib.parse.unquote(src).lstrip("/")
            base = rel.rsplit("/", 1)[-1]
            cands = [rel] + [d + base for d in R2_FOLDERS]
            for pre in ("files/", "downloads/", "files/downloads/", "figures/", "files/figures/"):
                if rel.startswith(pre):
                    cands += [rel[len(pre):]] + [d + rel[len(pre):] for d in R2_FOLDERS]
            for c in dict.fromkeys(cands):
                jobs[(src, c)] = ex.submit(status, R2 + urllib.parse.quote(c))
    found = {}
    for (src, c), fut in jobs.items():
        if fut.result() == 200 and src not in found:
            found[src] = c
    for src, c in sorted(found.items()):
        print(f"{src} {R2}{c} 301")
    print(f"\n{len(found)} of {len(rows)} file paths found on R2.", file=sys.stderr)


class _NoRedirect(urllib.request.HTTPRedirectHandler):
    def redirect_request(self, *a, **k):
        return None


def verify(args):
    opener = urllib.request.build_opener(_NoRedirect)
    base = args.base.rstrip("/")

    def get(path):
        try:
            return opener.open(base + path, timeout=20).status, None
        except urllib.error.HTTPError as e:
            return e.code, e.headers.get("Location")

    fixed, still = [], []
    for r in load_rows(args.rows):
        p = r["path"]
        hit = None
        for cand in dict.fromkeys((p, p.rstrip("/") + "/")):
            code, loc = get(cand)
            if code in (301, 302, 307, 308):
                hit = (cand, code, loc)
                break
        (fixed if hit else still).append((r, hit))
    print(f"redirected: {len(fixed)} paths ({sum(r['views'] for r, _ in fixed)} views); "
          f"still 404: {len(still)} paths ({sum(r['views'] for r, _ in still)} views)\n")
    for r, (cand, code, loc) in sorted(fixed, key=lambda x: -x[0]["views"]):
        print(f"{r['views']:4} {code} {cand} -> {loc}")


def main():
    ap = argparse.ArgumentParser(description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter)
    sub = ap.add_subparsers(dest="cmd", required=True)
    c = sub.add_parser("classify")
    c.add_argument("rows")
    c.add_argument("--public", default=os.path.join(ROOT, "public"))
    c.add_argument("--redirects", default=os.path.join(ROOT, "static", "_redirects"))
    c.add_argument("--all", action="store_true", help="don't filter out noise")
    p = sub.add_parser("probe")
    p.add_argument("rows")
    v = sub.add_parser("verify")
    v.add_argument("rows")
    v.add_argument("--base", default="http://127.0.0.1:8799")
    args = ap.parse_args()
    {"classify": classify, "probe": probe, "verify": verify}[args.cmd](args)


if __name__ == "__main__":
    main()
