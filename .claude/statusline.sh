#!/bin/bash
# Project status line: the sandbox kit's status line (~/.claude/statusline.sh),
# plus a second line linking the previews for the checkout being edited:
#
#   preview: local remote
#
# "local" links to the `hugo serve` process started from this checkout (see
# "Preview servers" in .claude/CLAUDE.md); "remote" links to the Cloudflare
# branch preview, derived from the branch name, so it is omitted on main.
# The line is omitted when neither exists. Links use OSC 8 terminal hyperlinks.

input=$(cat)

[ -x ~/.claude/statusline.sh ] && printf '%s' "$input" | ~/.claude/statusline.sh

cur_dir=$(printf '%s' "$input" | jq -r '.workspace.current_dir // .cwd // empty')
top=$(git -C "${cur_dir:-.}" rev-parse --show-toplevel 2>/dev/null) || exit 0
branch=$(git -C "$top" branch --show-current 2>/dev/null)

link() { printf '\033]8;;%s\033\\%s\033]8;;\033\\' "$1" "$2"; }

# Find the Hugo server whose working directory is this checkout.
port=""
for pid in $(pgrep -x hugo 2>/dev/null); do
  [ "$(readlink "/proc/$pid/cwd" 2>/dev/null)" = "$top" ] || continue
  port=$(tr '\0' '\n' < "/proc/$pid/cmdline" | sed -n '/^--port$/{n;p;q}; s/^--port=//p' | head -n 1)
  [ -n "$port" ] && break
done

links=()
[ -n "$port" ] && links+=("$(link "http://localhost:$port/" local)")
if [ -n "$branch" ] && [ "$branch" != "main" ]; then
  slug=$(printf '%s' "$branch" | tr '[:upper:]' '[:lower:]' | sed 's/[^a-z0-9-]/-/g')
  links+=("$(link "https://$slug-lincolnmullen-dot-com.lincoln-07e.workers.dev/" remote)")
fi

[ "${#links[@]}" -gt 0 ] && printf '\npreview: %s' "${links[*]}"
exit 0
