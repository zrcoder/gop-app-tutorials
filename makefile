build:
	@cd $(dir) && GOARCH=wasm GOOS=js gop build -o web/app.wasm && cp -R static web | true

run: build
	@gop run serve.gop $(dir) $(port)

clear:
	@rm -f gop_autogen.go
	@cd $(dir) && rm -rf web gop_autogen.go