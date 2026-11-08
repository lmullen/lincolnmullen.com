# Digests

Issues of *Working On It*, the occasional digest of new blog posts, sent by email through Kit.

Each issue is a Markdown file, `YYYY-MM-DD-working-on-it-<n>.md`, made by the `digest` skill (`.claude/skills/digest/SKILL.md`). Its front matter records which posts it included, so the next issue starts after them. `_template.md` holds the boilerplate.

Hugo only reads `content/`, so nothing here is published on the site. The rendered `.html` files are for pasting into Kit and are not committed.

`kit-template.html` is the Kit email template for the digest, styled like the website. It is the one `.html` file here that is committed. To install or update it, paste it into Kit under Email Templates → New Email Template → Create HTML Template (or edit the existing "Working On It" template). Kit fills in Liquid tags even inside HTML comments, so don't write a tag with its braces in the template's comments.

Issues #1–18 are from when *Working On It* was a newsletter. They are blog posts tagged `newsletter`.
