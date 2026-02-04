# lincolnmullen.com

This website (<https://lincolnmullen.com>) is my professional, scholarly website. I am a professor of history. 

## Technical details

This site uses the static site generator, [Hugo](https://gohugo.io).

Here are key commands for testing the site:

- Build: `hugo --cleanDestinationDir --minify`
- Local preview: `hugo serve --port 2323 --buildDrafts --buildFuture --disableFastRender`

The site has several content types. They have layout files in `./layouts/` and corresponding content in `./content/`. Here are the most important content types:

- `blog`: Longer blog posts, with a title, tags, and other metadata.
- `links`: Shorter social media—like posts without a title.
- `newsletter`: Archives of my newsletter.

Use the `figure` shortcodes for images with a caption, and the `image` shortcode for images that do not need a caption.

## Restrictions

Do not every run the `make deploy` or `rsync` commands automatically. Those should only be run manually by a human. 
