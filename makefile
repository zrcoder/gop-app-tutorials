build:
	@cd $(dir) && \
	GOARCH=wasm GOOS=js gop build -o public/web/app.wasm && \
	gop build -o public/app && \
	cp -R static public/web | true

run: build
	@cd $(dir)/public && ./app
genSite: build
	@cd $(dir)/public && ./app genSite && rm app

clear:
	@cd $(dir) && rm -rf web gop_autogen.go public

