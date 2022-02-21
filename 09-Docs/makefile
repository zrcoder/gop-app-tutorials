build:
	@GOARCH=wasm GOOS=js gop build -o public/web/app.wasm && \
	cp -R static public/web | true && \
	gop build -o public/app

gen: build
	@cd public && ./app

gen4repo: build
	@cd public && ./app $(repo)

run: gen
	@cd public && \
	go run ../../serve $(port)

clear:
	@rm -rf gop_autogen.go public