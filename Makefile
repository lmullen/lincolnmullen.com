preview :
	hugo serve --buildDrafts --buildFuture --disableFastRender

build :
	rm -rf public/*
	hugo --cleanDestinationDir

deploy : build
	rsync --delete --exclude-from=rsync-excludes \
		--checksum -avz \
		public/ reclaim:~/www/lincolnmullen.com/ \
		| tee deploy.log

