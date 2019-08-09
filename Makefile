preview :
	hugo serve --buildDrafts --buildFuture --disableFastRender

build :
	rm -rf public/*
	hugo --cleanDestinationDir --minify

deploy : build
	rsync --delete --omit-dir-times --exclude-from=rsync-excludes \
		--checksum -avz \
		--itemize-changes \
		public/ reclaim:~/www/lincolnmullen.com/ | egrep -v '^\.'
