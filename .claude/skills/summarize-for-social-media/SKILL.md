---
name: summarize-for-social-media
description: Generate first drafts of social media posts. Use when the user asks to create social media content for a post.
---

## Social Media Text and Description Generation

When asked to generate social media posts for a blog post, read the post file and write the text directly into the post's YAML front matter fields. Use the post that was provided in the context of the chat. Use the stylistic suggestions in other Claude skills and documentation.

If the `linkpost` field is present in the YAML header, always uses that as the "call to action" link at the end of a post.

Generate all of the following at once:

### Description

Write the `description` field: 1–2 sentences summarizing the post for meta tags and OpenGraph previews. This should be platform-neutral, informative, and suitable for search engine results. It is not a social media post.

### Social media fields

Write the `social:` block fields (`bluesky`, `mastodon`, `linkedin`) in the front matter:

- **bluesky**: ≤300 characters (including the URL). Conversational, informal tone. End with the permalink URL. No hashtags.
- **mastodon**: ≤500 characters (including the URL). Similar tone to BlueSky, slightly more room for detail. End with the permalink URL. No hashtags.
- **linkedin**: One or two short paragraphs. More professional framing, emphasizing scholarly or institutional significance. End with the permalink URL. No hashtags.

For all platforms: capture the main point or most interesting detail of the post, not just the title. For linkposts, summarize what makes the linked item worth reading.


### YAML formatting

- For `description`, `bluesky`, and `mastodon`, use double-quoted strings on a single line.
- For `linkedin`, if the text is multi-paragraph, use the YAML literal block scalar (`|`) to preserve line breaks. If it fits on one line, a double-quoted string is fine.

### Example

Here is an example of what a blog post's front matter should look like after generation:

```yaml
description: "A summary of what the example post is about, suitable for meta tags and search results."
social:
  bluesky: "Here's an interesting take on the example topic. https://lincolnmullen.com/blog/example-post-title/"
  mastodon: "Here's a slightly longer take on the example topic, with a bit more room for detail and nuance. https://lincolnmullen.com/blog/example-post-title/"
  linkedin: |
    I recently wrote about the example topic, which has significant implications for the field.

    The post explores how this matters for scholars and institutions working in this area. https://lincolnmullen.com/blog/example-post-title/
```
