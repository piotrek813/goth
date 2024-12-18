dev/templ:
	templ generate --watch --proxy="http://localhost:3000"

dev/tailwind:
	npm run tailwind

dev/server:
	air \
	--build.cmd "go build -o tmp/bin/main" --build.bin "tmp/bin/main" --build.delay "100" \
	--build.exclude_dir "node_modules" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true

dev/sync_assets:
	air \
	--build.cmd "templ generate --notify-proxy" \
	--build.bin "true" \
	--build.delay "100" \
	--build.exclude_dir "" \
	--build.include_dir "public" \
	--build.include_ext "js,css"

dev:
	# for whatever reason dev/tailwind has to be the first thing to run
	make -j4 dev/tailwind dev/templ dev/server dev/sync_assets
