"""Helpers for the newsletter skill: list posts since the last issue, convert a
post to email-ready Markdown, and render an issue to HTML for pasting into Kit.

Run from the repository root with:

    uv run --with pyyaml --with markdown .claude/skills/newsletter/newsletter.py <command> ...

Commands:
    last                       Print the last issue's number and date.
    candidates                 List posts published since the last issue (JSON).
    item <dir> <full|summary>  Print a post as email-ready Markdown.
    render <issue.md>          Write <issue>.html next to the issue file.
"""

import datetime as dt
import html
import json
import re
import sys
from pathlib import Path

import markdown
import yaml

SITE = "https://lincolnmullen.com"
BLOG = Path("content/blog")
EMAILS = Path("emails")
TZ = dt.timezone(dt.timedelta(hours=-5))  # fallback when a date has no offset

# The last issue of Working on It published as a blog post, before issues moved
# to emails/. Used when emails/ has no issues yet.
LEGACY_LAST = {"number": 18, "date": "2026-03-08T22:32:13-04:00"}

MORE = "<!--more-->"


def split_front_matter(text):
    m = re.match(r"^---\s*\n(.*?)\n---\s*\n?(.*)$", text, re.S)
    if not m:
        return {}, text
    return yaml.safe_load(m.group(1)) or {}, m.group(2)


def parse_date(value):
    if isinstance(value, dt.datetime):
        d = value
    elif isinstance(value, dt.date):
        d = dt.datetime(value.year, value.month, value.day)
    else:
        d = dt.datetime.fromisoformat(str(value))
    return d if d.tzinfo else d.replace(tzinfo=TZ)


def last_issue():
    issues = []
    for p in EMAILS.glob("*.md"):
        if p.name.startswith("_") or p.name == "README.md":
            continue
        fm, _ = split_front_matter(p.read_text())
        if "number" in fm and "date" in fm:
            issues.append((int(fm["number"]), str(fm["date"]), p.name))
    if issues:
        n, d, name = max(issues)
        return {"number": n, "date": d, "file": name}
    return dict(LEGACY_LAST, file=None)


def load_post(d):
    path = BLOG / d / "index.md"
    fm, body = split_front_matter(path.read_text())
    return path, fm, body


def candidates():
    since = parse_date(last_issue()["date"])
    out = []
    for path in sorted(BLOG.glob("*/index.md")):
        fm, body = split_front_matter(path.read_text())
        if fm.get("draft") or "date" not in fm:
            continue
        if "newsletter" in (fm.get("tags") or []):
            continue
        date = parse_date(fm["date"])
        if date <= since:
            continue
        first_line = next((l for l in body.strip().splitlines() if l.strip()), "")
        out.append(
            {
                "dir": path.parent.name,
                "title": fm.get("title") or None,
                "date": date.isoformat(),
                "description": fm.get("description") or "",
                "first_line": first_line[:160],
                "has_more": MORE in body,
                "linkpost": fm.get("linkpost") or None,
                "via": fm.get("via") or None,
                "crosspost": fm.get("crosspost") or None,
                "words": len(body.split()),
            }
        )
    out.sort(key=lambda p: p["date"])
    return out


SHORTCODE = re.compile(r"\{\{<\s*(\w+)\s*(.*?)\s*>\}\}", re.S)
PARAM = re.compile(r'(\w+)\s*=\s*"((?:[^"\\]|\\.)*)"')


def expand_shortcodes(body, base, warnings):
    """Turn the site's shortcodes into plain Markdown with absolute URLs."""

    def url(src):
        if src.startswith(("http://", "https://")):
            return src
        if src.startswith("/"):
            return SITE + src
        return base + src

    def repl(m):
        name, args = m.group(1), m.group(2)
        params = dict(PARAM.findall(args))
        if name in ("image", "figure"):
            src = params.get("src", "")
            alt = params.get("alt") or params.get("caption", "")
            alt = re.sub(r"[\[\]]", "", re.sub(r"<[^>]+>", "", alt))
            img = f"![{alt}]({url(src)})"
            link = params.get("link")
            if link:
                img = f"[{img}]({url(link)})"
            caption = params.get("caption")
            if name == "figure" and caption:
                img += f"\n\n*{caption}*"
            return img
        if name == "youtube":
            vid = params.get("id") or args.strip().strip('"')
            title = params.get("title", "Watch the video on YouTube")
            return f"[{title}](https://www.youtube.com/watch?v={vid})"
        if name == "video":
            caption = params.get("caption") or "Watch the video"
            return f"[{caption}]({url(params.get('src', ''))})"
        warnings.append(f"unhandled shortcode: {m.group(0)}")
        return m.group(0)

    return SHORTCODE.sub(repl, body)


def item(d, mode):
    path, fm, body = load_post(d)
    slug = fm.get("slug") or d
    permalink = f"{SITE}/blog/{slug}/"
    date = parse_date(fm["date"])
    warnings = []

    if mode == "summary":
        if MORE in body:
            text = body.split(MORE, 1)[0].strip()
        elif fm.get("description"):
            text = fm["description"]
        else:
            text = body.strip().split("\n\n", 1)[0]
    else:
        text = body.replace(MORE, "").strip()
    text = expand_shortcodes(text, permalink, warnings)

    lines = []
    title = fm.get("title")
    if title:
        lines.append(f"## [{title}]({permalink})")
    else:
        lines.append(f"**[{date.strftime('%B %-d, %Y')}]({permalink})**")
    lines.append("")
    if fm.get("crosspost"):
        cp = fm["crosspost"]
        lines.append(f"*Originally published at [{cp.get('source', cp.get('url'))}]({cp.get('url')}).*")
        lines.append("")
    if fm.get("linkpost"):
        lines.append(f"Link: <{fm['linkpost']}>")
        lines.append("")
    lines.append(text)
    lines.append("")
    if mode == "summary":
        lines.append(f"[Read the full post →]({permalink})")
        lines.append("")
    elif fm.get("via"):
        lines.append(f"*Via: <{fm['via']}>*")
        lines.append("")

    for w in warnings:
        print(f"WARNING ({d}): {w}", file=sys.stderr)
    return "\n".join(lines).rstrip() + "\n"


STYLE = {
    "body": "margin:0;padding:0;",
    "h2": "font-size:20px;line-height:1.3;margin:32px 0 8px;",
    "img": "max-width:100%;height:auto;display:block;margin:16px 0;",
    "blockquote": "margin:16px 0;padding-left:16px;border-left:3px solid #ccc;color:#555;",
}


def render(issue):
    issue = Path(issue)
    fm, body = split_front_matter(issue.read_text())
    content = markdown.markdown(body, extensions=["extra", "sane_lists", "smarty"])
    for tag, style in STYLE.items():
        if tag == "body":
            continue
        content = re.sub(rf"<{tag}(?=[\s>])", f'<{tag} style="{style}"', content)
    title = html.escape(str(fm.get("title", "")))
    preview = html.escape(str(fm.get("preview_text", "")))
    doc = f"""<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{title}</title>
</head>
<body style="font-family:Georgia,serif;font-size:17px;line-height:1.55;max-width:640px;margin:2em auto;padding:0 1em;color:#222;">
<!-- Subject: {title} -->
<!-- Preview text: {preview} -->
{content}
</body>
</html>
"""
    out = issue.with_suffix(".html")
    out.write_text(doc)
    print(out)


def main(argv):
    if len(argv) < 2:
        print(__doc__)
        return 2
    cmd = argv[1]
    if cmd == "last":
        print(json.dumps(last_issue(), indent=2))
    elif cmd == "candidates":
        print(json.dumps(candidates(), indent=2, ensure_ascii=False))
    elif cmd == "item" and len(argv) == 4 and argv[3] in ("full", "summary"):
        sys.stdout.write(item(argv[2], argv[3]))
    elif cmd == "render" and len(argv) == 3:
        render(argv[2])
    else:
        print(__doc__)
        return 2
    return 0


if __name__ == "__main__":
    sys.exit(main(sys.argv))
