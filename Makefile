preview :
	hugo serve --port 1313 --buildDrafts --buildFuture --disableFastRender

build :
	rm -rf public/*
	hugo --cleanDestinationDir --minify --baseURL "https://lincolnmullen.com"

# Bring Hugo's committed processed-image cache (resources/_gen/images) in line
# with what a production build needs; --check only reports drift.
gen-cache :
	./gen-cache.sh

check-gen-cache :
	./gen-cache.sh --check

# Proofread prose with Vale (run `vale sync` first). The first pass covers the
# body, title, and description; the second the social media text, in posts
# that have any. Pass FILES to check only some files: make vale FILES=content/blog/x
FILES ?= content $(wildcard digests/2*.md)
vale :
	-vale $(FILES)
	grep -rl --include='*.md' '^social:' $(FILES) | xargs vale --config .vale-social.ini
