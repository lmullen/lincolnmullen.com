# Digests

Issues of *Working On It*, the occasional digest of new blog posts, sent by email through Kit.

Each issue is a Markdown file, `YYYY-MM-DD-working-on-it-<n>.md`, made by the `digest` skill (`.claude/skills/digest/SKILL.md`). Its front matter records which posts it included, so the next issue starts after them. `_template.md` is the skeleton of an issue: front matter, an optional intro, and the posts.

Hugo only reads `content/`, so nothing here is published on the site. The rendered files are not committed: `<issue>.kit.html` is the issue's HTML, to paste into an HTML block in a Kit broadcast, and `<issue>.preview.html` shows it inside the Kit template.

`kit-template.html` is the Kit email template for the digest, styled like the website. It holds what is the same in every issue: the blog's name, the headnote, and the footer, and the styles for the issue's HTML. It is the one `.html` file here that is committed. To install or update it, paste it into Kit under Email Templates → New Email Template → Create HTML Template (or edit the existing "Working On It" template). Kit fills in Liquid tags even inside HTML comments, so don't write a tag with its braces in the template's comments.

Issues #1–18 are from when *Working On It* was a newsletter. They are blog posts tagged `newsletter`.
