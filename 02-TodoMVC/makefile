build:
	@GOARCH=wasm GOOS=js gop build -o public/web/app.wasm && \
	cp -R static public/web && \
	gop build -o public/app

gen: build
	@cd public && ./app

run: gen
	@cd public && go run ../../serve $(port)

clear:
	@rm -rf gop_autogen.go public