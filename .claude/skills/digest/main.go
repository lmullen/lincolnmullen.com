// Command digest is the helper for the digest skill. It lists posts
// published since the last issue of Working On It, converts a post to
// email-ready Markdown, and renders an issue to HTML for a Kit HTML block.
//
// Run it from anywhere in the repository:
//
//	go -C .claude/skills/digest run . <command> [args]
//
// It changes to the repository root first, so file arguments are relative to
// the root. Commands:
//
//	last            Print the last issue's number and date (JSON).
//	candidates      List posts published since the last issue (JSON).
//	items <dir>...  Print the posts as one block for an issue: titled posts,
//	                then untitled notes, each newest first, separated by
//	                horizontal rules. Each post is shown as by item.
//	item <dir>      Print a post as email-ready Markdown, as the website shows
//	                it: the text before <!--more--> plus a link to the full
//	                post if there is a break, otherwise the whole post.
//	render <issue>  Write <issue>.kit.html (the issue's HTML, to paste into a
//	                Kit HTML block) and <issue>.preview.html (that HTML in the
//	                Kit template, to view in a browser) next to the issue.
package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	gmhtml "github.com/yuin/goldmark/renderer/html"
	"gopkg.in/yaml.v3"
)

const (
	site    = "https://lincolnmullen.com"
	blog    = "content/blog"
	digests = "digests"
	more    = "<!--more-->"
	// The blog's "long" date form (site.Params.dateForms.long in hugo.yaml).
	longDate = "2 January 2006"
)

// The last issue published as a blog post, before issues moved to digests/.
// Used when digests/ has no issues yet.
var legacyLast = issue{Number: 18, Date: "2026-03-08T22:32:13-04:00"}

type frontMatter struct {
	Title       string `yaml:"title"`
	Date        string `yaml:"date"`
	Slug        string `yaml:"slug"`
	Draft       bool   `yaml:"draft"`
	Description string `yaml:"description"`
	Linkpost    string `yaml:"linkpost"`
	Via         string `yaml:"via"`
	Crosspost   *struct {
		URL    string `yaml:"url"`
		Source string `yaml:"source"`
	} `yaml:"crosspost"`
	Tags any `yaml:"tags"`
}

type issue struct {
	Number int    `yaml:"number" json:"number"`
	Date   string `yaml:"date" json:"date"`
	File   string `yaml:"-" json:"file,omitempty"`
}

type candidate struct {
	Dir         string `json:"dir"`
	Title       string `json:"title,omitempty"`
	Date        string `json:"date"`
	Description string `json:"description,omitempty"`
	FirstLine   string `json:"first_line"`
	HasMore     bool   `json:"has_more"`
	Words       int    `json:"words"`
	Linkpost    string `json:"linkpost,omitempty"`
}

var fmRE = regexp.MustCompile(`(?s)\A---\s*\n(.*?)\n---\s*\n?(.*)\z`)

func splitFrontMatter(text string, v any) (string, error) {
	m := fmRE.FindStringSubmatch(text)
	if m == nil {
		return text, nil
	}
	if err := yaml.Unmarshal([]byte(m[1]), v); err != nil {
		return "", err
	}
	return m[2], nil
}

func parseDate(s string) (time.Time, error) {
	for _, layout := range []string{time.RFC3339, "2006-01-02T15:04:05", "2006-01-02 15:04:05 -0700", "2006-01-02"} {
		if t, err := time.Parse(layout, s); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("cannot parse date %q", s)
}

func hasTag(tags any, want string) bool {
	switch t := tags.(type) {
	case string:
		return t == want
	case []any:
		for _, x := range t {
			if s, ok := x.(string); ok && s == want {
				return true
			}
		}
	}
	return false
}

// findRoot walks up from the working directory to the repository root.
func findRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, blog)); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", errors.New("cannot find the repository root (no content/blog above here)")
		}
		dir = parent
	}
}

func lastIssue() (issue, error) {
	files, err := filepath.Glob(filepath.Join(digests, "*.md"))
	if err != nil {
		return issue{}, err
	}
	var issues []issue
	for _, f := range files {
		base := filepath.Base(f)
		if strings.HasPrefix(base, "_") || base == "README.md" {
			continue
		}
		b, err := os.ReadFile(f)
		if err != nil {
			return issue{}, err
		}
		var is issue
		if _, err := splitFrontMatter(string(b), &is); err != nil {
			return issue{}, fmt.Errorf("%s: %w", f, err)
		}
		if is.Number > 0 && is.Date != "" {
			is.File = base
			issues = append(issues, is)
		}
	}
	if len(issues) == 0 {
		return legacyLast, nil
	}
	sort.Slice(issues, func(i, j int) bool { return issues[i].Number < issues[j].Number })
	return issues[len(issues)-1], nil
}

func loadPost(dir string) (frontMatter, string, error) {
	var fm frontMatter
	b, err := os.ReadFile(filepath.Join(blog, dir, "index.md"))
	if err != nil {
		return fm, "", err
	}
	body, err := splitFrontMatter(string(b), &fm)
	if err != nil {
		return fm, "", fmt.Errorf("%s: %w", dir, err)
	}
	return fm, body, nil
}

func candidates() ([]candidate, error) {
	last, err := lastIssue()
	if err != nil {
		return nil, err
	}
	since, err := parseDate(last.Date)
	if err != nil {
		return nil, err
	}
	paths, err := filepath.Glob(filepath.Join(blog, "*", "index.md"))
	if err != nil {
		return nil, err
	}
	out := []candidate{}
	for _, p := range paths {
		dir := filepath.Base(filepath.Dir(p))
		fm, body, err := loadPost(dir)
		if err != nil {
			return nil, err
		}
		if fm.Draft || fm.Date == "" || hasTag(fm.Tags, "newsletter") {
			continue
		}
		date, err := parseDate(fm.Date)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", dir, err)
		}
		if !date.After(since) {
			continue
		}
		first := ""
		for _, l := range strings.Split(strings.TrimSpace(body), "\n") {
			if strings.TrimSpace(l) != "" {
				first = l
				break
			}
		}
		if r := []rune(first); len(r) > 160 {
			first = string(r[:160])
		}
		out = append(out, candidate{
			Dir: dir, Title: fm.Title, Date: date.Format(time.RFC3339),
			Description: fm.Description, FirstLine: first,
			HasMore: strings.Contains(body, more), Words: len(strings.Fields(body)),
			Linkpost: fm.Linkpost,
		})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Date < out[j].Date })
	return out, nil
}

var (
	shortcodeRE = regexp.MustCompile(`(?s)\{\{<\s*(\w+)\s*(.*?)\s*>\}\}`)
	paramRE     = regexp.MustCompile(`(\w+)\s*=\s*"((?:[^"\\]|\\.)*)"`)
	tagRE       = regexp.MustCompile(`<[^>]+>`)
)

var schemeRE = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9+.-]*:`)

// absURL makes a URL from a post absolute. Site-relative URLs ("/blog/...")
// resolve against the site, and page-relative ones ("file.pdf") against the
// post's permalink, base. URLs with a scheme and fragments are unchanged.
func absURL(u, base string) string {
	switch {
	case u == "", strings.HasPrefix(u, "#"), schemeRE.MatchString(u):
		return u
	case strings.HasPrefix(u, "//"):
		return "https:" + u
	case strings.HasPrefix(u, "/"):
		return site + u
	}
	return base + strings.TrimPrefix(u, "./")
}

var (
	fenceRE      = regexp.MustCompile("(?m)^[ \t]*(```|~~~)")
	inlineCodeRE = regexp.MustCompile("``[^`]*``|`[^`\n]*`")
)

// outsideCode applies f to the parts of a Markdown text that are not fenced
// code blocks or inline code spans, so rewriting links and footnotes cannot
// change code samples.
func outsideCode(text string, f func(string) string) string {
	var b strings.Builder
	inFence := false
	for _, line := range strings.SplitAfter(text, "\n") {
		if fenceRE.MatchString(line) {
			inFence = !inFence
			b.WriteString(line)
			continue
		}
		if inFence {
			b.WriteString(line)
			continue
		}
		last := 0
		for _, loc := range inlineCodeRE.FindAllStringIndex(line, -1) {
			b.WriteString(f(line[last:loc[0]]))
			b.WriteString(line[loc[0]:loc[1]])
			last = loc[1]
		}
		b.WriteString(f(line[last:]))
	}
	return b.String()
}

var (
	mdLinkRE   = regexp.MustCompile(`(\]\(\s*<?)([^)\s>]+)`)
	htmlLinkRE = regexp.MustCompile(`((?:href|src)\s*=\s*")([^"]*)(")`)
	refDefRE   = regexp.MustCompile(`^(\s{0,3}\[[^\]^][^\]]*\]:\s*<?)([^\s>]+)`)
)

// absoluteLinks makes the URLs of Markdown links and images, reference
// definitions, and HTML href and src attributes absolute, since links
// relative to the website are broken in an email.
func absoluteLinks(text, base string) string {
	return outsideCode(text, func(s string) string {
		repl := func(re *regexp.Regexp) {
			s = re.ReplaceAllStringFunc(s, func(m string) string {
				g := re.FindStringSubmatch(m)
				return g[1] + absURL(g[2], base) + strings.Join(g[3:], "")
			})
		}
		repl(mdLinkRE)
		repl(htmlLinkRE)
		repl(refDefRE)
		return s
	})
}

var (
	footnoteDefRE = regexp.MustCompile(`^\[\^([^\]]+)\]:[ \t]*(.*?)\n?$`)
	footnoteRefRE = regexp.MustCompile(`\[\^([^\]]+)\]`)
)

// footnotes renders a post's Markdown footnotes without relying on the
// renderer: each referenced footnote becomes a superscript number, numbered
// within the post so that posts in one issue cannot collide, and its note is
// listed at the end of the post. A reference whose definition is not in the
// text, as when the post is cut at <!--more--> above its notes, is dropped.
func footnotes(text string) string {
	// Definitions are whole lines, so find them line by line, skipping
	// fenced code blocks. (The site's footnotes are all one line long.)
	defs := map[string]string{}
	var kept strings.Builder
	inFence := false
	for _, line := range strings.SplitAfter(text, "\n") {
		if fenceRE.MatchString(line) {
			inFence = !inFence
		} else if m := footnoteDefRE.FindStringSubmatch(line); !inFence && m != nil {
			defs[m[1]] = m[2]
			continue
		}
		kept.WriteString(line)
	}
	text = kept.String()
	var order []string
	num := map[string]int{}
	text = outsideCode(text, func(s string) string {
		return footnoteRefRE.ReplaceAllStringFunc(s, func(ref string) string {
			id := footnoteRefRE.FindStringSubmatch(ref)[1]
			if _, ok := defs[id]; !ok {
				return ""
			}
			if num[id] == 0 {
				order = append(order, id)
				num[id] = len(order)
			}
			return fmt.Sprintf("<sup>%d</sup>", num[id])
		})
	})
	text = strings.TrimSpace(text)
	for _, id := range order {
		text += fmt.Sprintf("\n\n<sup>%d</sup> %s", num[id], defs[id])
	}
	return text
}

// expandShortcodes turns the site's shortcodes into plain Markdown, with
// absolute URLs to the original files so email clients can show images.
func expandShortcodes(body, base string, warn func(string)) string {
	abs := func(src string) string { return absURL(src, base) }
	return shortcodeRE.ReplaceAllStringFunc(body, func(sc string) string {
		m := shortcodeRE.FindStringSubmatch(sc)
		name, args := m[1], m[2]
		params := map[string]string{}
		for _, p := range paramRE.FindAllStringSubmatch(args, -1) {
			params[p[1]] = p[2]
		}
		switch name {
		case "image", "figure":
			alt := params["alt"]
			if alt == "" {
				alt = params["caption"]
			}
			alt = strings.NewReplacer("[", "", "]", "").Replace(tagRE.ReplaceAllString(alt, ""))
			img := fmt.Sprintf("![%s](%s)", alt, abs(params["src"]))
			if link := params["link"]; link != "" {
				img = fmt.Sprintf("[%s](%s)", img, abs(link))
			}
			if c := params["caption"]; name == "figure" && c != "" {
				img += "\n\n*" + c + "*"
			}
			return img
		case "youtube":
			id := params["id"]
			if id == "" {
				id = strings.Trim(strings.TrimSpace(args), `"`)
			}
			title := params["title"]
			if title == "" {
				title = "Watch the video on YouTube"
			}
			return fmt.Sprintf("[%s](https://www.youtube.com/watch?v=%s)", title, id)
		case "video":
			c := params["caption"]
			if c == "" {
				c = "Watch the video"
			}
			return fmt.Sprintf("[%s](%s)", c, abs(params["src"]))
		}
		warn("unhandled shortcode: " + sc)
		return sc
	})
}

func item(dir string) (string, error) {
	fm, body, err := loadPost(dir)
	if err != nil {
		return "", err
	}
	slug := fm.Slug
	if slug == "" {
		slug = dir
	}
	permalink := fmt.Sprintf("%s/blog/%s/", site, slug)
	date, err := parseDate(fm.Date)
	if err != nil {
		return "", fmt.Errorf("%s: %w", dir, err)
	}

	// Do what the website does: cut at <!--more--> if there is one.
	text, _, truncated := strings.Cut(body, more)
	text = strings.TrimSpace(text)
	text = expandShortcodes(text, permalink, func(w string) {
		fmt.Fprintf(os.Stderr, "WARNING (%s): %s\n", dir, w)
	})
	text = absoluteLinks(footnotes(text), permalink)

	// The lines the helper adds are written as HTML with classes, which the
	// Kit template (digests/kit-template.html) styles: the issue is pasted
	// into a Kit HTML block, which keeps them. Each is its own HTML block in
	// the Markdown, so it needs blank lines around it.
	//
	// The date is written as the blog writes it (its "long" date form) and
	// links to the post. It goes below a titled post's heading and atop an
	// untitled note.
	var b strings.Builder
	if fm.Title != "" {
		fmt.Fprintf(&b, "## [%s](%s)\n\n", fm.Title, permalink)
	}
	fmt.Fprintf(&b, "<p class=\"woi-date\"><a href=\"%s\">%s</a></p>\n\n",
		html.EscapeString(permalink), date.Format(longDate))
	b.WriteString(text + "\n\n")
	// The read-more action, as the website sets it: the words on an orange
	// rule, then an orange arrow.
	if truncated {
		fmt.Fprintf(&b, "<p class=\"woi-more\"><a href=\"%s\"><span class=\"woi-more-text\">Read the post</span> <span class=\"woi-arrow\">→</span></a></p>\n\n",
			html.EscapeString(permalink))
	}
	// As on the blog, the post's outside links go in one line after the
	// post, each as a labeled link rather than a bare URL.
	var links []string
	aff := func(label, u string) string {
		return fmt.Sprintf("<a class=\"woi-aff\" href=\"%s\">%s <span class=\"woi-arrow\">→</span></a>",
			html.EscapeString(u), label)
	}
	if fm.Linkpost != "" {
		links = append(links, aff("Link", fm.Linkpost))
	}
	if fm.Via != "" {
		links = append(links, aff("Via", fm.Via))
	}
	if cp := fm.Crosspost; cp != nil && cp.URL != "" {
		links = append(links, aff("Crossposted", cp.URL))
	}
	if len(links) > 0 {
		b.WriteString("<p class=\"woi-links\">" + strings.Join(links, " ") + "</p>\n")
	}
	return strings.TrimRight(b.String(), "\n") + "\n", nil
}

// separator goes between posts. It is written as "* * *" rather than "---",
// which Markdown would read as a heading underline after a line of text.
const separator = "\n* * *\n\n"

// items converts the given posts and orders them for an issue: titled posts
// first, then untitled notes, each group newest first like the blog, with a
// separator between every post.
func items(dirs []string) (string, error) {
	type post struct {
		dir    string
		titled bool
		date   time.Time
	}
	var posts []post
	for _, dir := range dirs {
		fm, _, err := loadPost(dir)
		if err != nil {
			return "", err
		}
		date, err := parseDate(fm.Date)
		if err != nil {
			return "", fmt.Errorf("%s: %w", dir, err)
		}
		posts = append(posts, post{dir, fm.Title != "", date})
	}
	sort.SliceStable(posts, func(i, j int) bool {
		if posts[i].titled != posts[j].titled {
			return posts[i].titled
		}
		return posts[i].date.After(posts[j].date)
	})
	parts := make([]string, 0, len(posts))
	for _, p := range posts {
		s, err := item(p.dir)
		if err != nil {
			return "", err
		}
		parts = append(parts, s)
	}
	return strings.Join(parts, separator), nil
}

// kitTemplate is the Kit email template the issue is sent with. render uses
// it only for the preview; in Kit, the template wraps the pasted fragment.
const kitTemplate = digests + "/kit-template.html"

// render converts an issue to HTML and writes two files next to it:
//
//   - <issue>.kit.html: the issue's content alone, with no <html>, <head>,
//     or <body> and no inline styles, to paste into a Kit HTML block. The
//     Kit template supplies the styles, and Kit inlines them when it sends.
//   - <issue>.preview.html: that content wrapped in the Kit template, with
//     Kit's tags filled in with placeholders, to see the email in a browser.
//
// It returns the two paths.
func render(path string) ([]string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var fm struct {
		Title       string `yaml:"title"`
		PreviewText string `yaml:"preview_text"`
	}
	body, err := splitFrontMatter(string(b), &fm)
	if err != nil {
		return nil, err
	}
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM, extension.Typographer),
		goldmark.WithRendererOptions(gmhtml.WithUnsafe()),
	)
	var buf bytes.Buffer
	if err := md.Convert([]byte(body), &buf); err != nil {
		return nil, err
	}
	fragment := buf.String()

	tmpl, err := os.ReadFile(kitTemplate)
	if err != nil {
		return nil, err
	}
	page := strings.NewReplacer(
		"{{ message_content }}", fragment,
		"{{ unsubscribe_url }}", "#",
		"{{ subscriber_preferences_url }}", "#",
		"{{ address }}", "[Your mailing address, from Kit]",
	).Replace(string(tmpl))
	if strings.Contains(page, "{{") {
		fmt.Fprintln(os.Stderr, "WARNING: the preview has a Liquid tag the helper doesn't fill in")
	}
	preview := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta name="color-scheme" content="light dark">
<title>Preview: %[1]s</title>
</head>
<body style="margin:0;background:#f3f4f6;">
<!-- Subject: %[1]s -->
<!-- Preview text: %[2]s -->
%[3]s
</body>
</html>
`, html.EscapeString(fm.Title), html.EscapeString(fm.PreviewText), page)

	base := strings.TrimSuffix(path, filepath.Ext(path))
	outs := []string{base + ".kit.html", base + ".preview.html"}
	if err := os.WriteFile(outs[0], []byte(fragment), 0o644); err != nil {
		return nil, err
	}
	return outs, os.WriteFile(outs[1], []byte(preview), 0o644)
}

func printJSON(v any) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}

func run(args []string) error {
	usage := errors.New("usage: digest last | candidates | items <dir>... | item <dir> | render <issue.md>")
	if len(args) == 0 {
		return usage
	}
	root, err := findRoot()
	if err != nil {
		return err
	}
	if err := os.Chdir(root); err != nil {
		return err
	}
	switch {
	case args[0] == "last" && len(args) == 1:
		is, err := lastIssue()
		if err != nil {
			return err
		}
		return printJSON(is)
	case args[0] == "candidates" && len(args) == 1:
		c, err := candidates()
		if err != nil {
			return err
		}
		return printJSON(c)
	case args[0] == "item" && len(args) == 2:
		s, err := item(args[1])
		if err != nil {
			return err
		}
		_, err = os.Stdout.WriteString(s)
		return err
	case args[0] == "items" && len(args) >= 2:
		s, err := items(args[1:])
		if err != nil {
			return err
		}
		_, err = os.Stdout.WriteString(s)
		return err
	case args[0] == "render" && len(args) == 2:
		outs, err := render(args[1])
		if err != nil {
			return err
		}
		fmt.Println(strings.Join(outs, "\n"))
		return nil
	}
	return usage
}

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "digest:", err)
		os.Exit(1)
	}
}
