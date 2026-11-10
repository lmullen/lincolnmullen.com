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
