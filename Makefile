dev/templ:
	templ generate --watch --proxy="http://localhost:3000"

dev/tailwind:
	npx tailwindcss -o ./public/output.css --watch

dev/server:
	air \
	--build.cmd "go build -o tmp/bin/main" --build.bin "tmp/bin/main" --build.delay "100" \
	--build.exclude_dir "node_modules" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true

live/sync_assets:
	air \
	--build.cmd "templ generate --notify-proxy" \
	--build.bin "true" \
	--build.delay "100" \
	--build.exclude_dir "" \
	--build.include_dir "public" \
	--build.include_ext "js,css"

# start all 5 watch processes in parallel.
dev: 
	make -j5 dev/templ dev/server live/tailwind
