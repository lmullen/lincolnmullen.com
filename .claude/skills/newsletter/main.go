// Command newsletter is the helper for the newsletter skill. It lists posts
// published since the last issue of Working on It, converts a post to
// email-ready Markdown, and renders an issue to HTML for pasting into Kit.
//
// Run it from anywhere in the repository:
//
//	go -C .claude/skills/newsletter run . <command> [args]
//
// It changes to the repository root first, so file arguments are relative to
// the root. Commands:
//
//	last            Print the last issue's number and date (JSON).
//	candidates      List posts published since the last issue (JSON).
//	items <dir>...  Print the posts as one block for an issue: titled posts,
//	                then untitled notes, each in date order, separated by
//	                horizontal rules. Each post is shown as by item.
//	item <dir>      Print a post as email-ready Markdown, as the website shows
//	                it: the text before <!--more--> plus a link to the full
//	                post if there is a break, otherwise the whole post.
//	render <issue>  Write <issue>.html next to the issue's Markdown file.
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
	site   = "https://lincolnmullen.com"
	blog   = "content/blog"
	emails = "emails"
	more   = "<!--more-->"
)

// The last issue published as a blog post, before issues moved to emails/.
// Used when emails/ has no issues yet.
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
	files, err := filepath.Glob(filepath.Join(emails, "*.md"))
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

// expandShortcodes turns the site's shortcodes into plain Markdown, with
// absolute URLs to the original files so email clients can show images.
func expandShortcodes(body, base string, warn func(string)) string {
	abs := func(src string) string {
		switch {
		case strings.HasPrefix(src, "http://"), strings.HasPrefix(src, "https://"):
			return src
		case strings.HasPrefix(src, "/"):
			return site + src
		}
		return base + src
	}
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

	var b strings.Builder
	if fm.Title != "" {
		fmt.Fprintf(&b, "## [%s](%s)\n\n", fm.Title, permalink)
	} else {
		fmt.Fprintf(&b, "**[%s](%s)**\n\n", date.Format("January 2, 2006"), permalink)
	}
	if cp := fm.Crosspost; cp != nil && cp.URL != "" {
		src := cp.Source
		if src == "" {
			src = cp.URL
		}
		fmt.Fprintf(&b, "*Originally published at [%s](%s).*\n\n", src, cp.URL)
	}
	if fm.Linkpost != "" {
		fmt.Fprintf(&b, "Link: <%s>\n\n", fm.Linkpost)
	}
	b.WriteString(text + "\n\n")
	if truncated {
		fmt.Fprintf(&b, "[Read the full post →](%s)\n", permalink)
	} else if fm.Via != "" {
		fmt.Fprintf(&b, "*Via: <%s>*\n", fm.Via)
	}
	return strings.TrimRight(b.String(), "\n") + "\n", nil
}

// separator goes between posts. It is written as "* * *" rather than "---",
// which Markdown would read as a heading underline after a line of text.
const separator = "\n* * *\n\n"

// items converts the given posts and orders them for an issue: titled posts
// first, then untitled notes, each group in chronological order, with a
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
		return posts[i].date.Before(posts[j].date)
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

// Inline styles, since email clients ignore most external and <head> CSS.
var styles = []struct{ tag, style string }{
	{"h2", "font-size:20px;line-height:1.3;margin:32px 0 8px;"},
	{"img", "max-width:100%;height:auto;display:block;margin:16px 0;"},
	{"blockquote", "margin:16px 0;padding-left:16px;border-left:3px solid #ccc;color:#555;"},
}

func render(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	var fm struct {
		Title       string `yaml:"title"`
		PreviewText string `yaml:"preview_text"`
	}
	body, err := splitFrontMatter(string(b), &fm)
	if err != nil {
		return "", err
	}
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM, extension.Typographer),
		goldmark.WithRendererOptions(gmhtml.WithUnsafe()),
	)
	var buf bytes.Buffer
	if err := md.Convert([]byte(body), &buf); err != nil {
		return "", err
	}
	content := buf.String()
	for _, s := range styles {
		re := regexp.MustCompile(`<` + s.tag + `([\s>/])`)
		content = re.ReplaceAllString(content, `<`+s.tag+` style="`+s.style+`"$1`)
	}
	doc := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>%[1]s</title>
</head>
<body style="font-family:Georgia,serif;font-size:17px;line-height:1.55;max-width:640px;margin:2em auto;padding:0 1em;color:#222;">
<!-- Subject: %[1]s -->
<!-- Preview text: %[2]s -->
%[3]s</body>
</html>
`, html.EscapeString(fm.Title), html.EscapeString(fm.PreviewText), content)
	out := strings.TrimSuffix(path, filepath.Ext(path)) + ".html"
	return out, os.WriteFile(out, []byte(doc), 0o644)
}

func printJSON(v any) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}

func run(args []string) error {
	usage := errors.New("usage: newsletter last | candidates | items <dir>... | item <dir> | render <issue.md>")
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
		out, err := render(args[1])
		if err != nil {
			return err
		}
		fmt.Println(out)
		return nil
	}
	return usage
}

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "newsletter:", err)
		os.Exit(1)
	}
}
