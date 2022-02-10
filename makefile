build:
	@cd $(dir) && GOARCH=wasm GOOS=js gop build -o web/app.wasm && cp -R static web | true

run: build
	@cd $(dir) && gop run .

clear:
	@cd $(dir) && rm -rf web gop_autogen.go