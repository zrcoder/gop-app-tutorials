## Intro

**app** is a package to build [progressive web apps (PWA)](https://developers.google.com/web/progressive-web-apps/) with the [Go programming language](https://golang.org) and [WebAssembly](https://webassembly.org).

You are about to learn how to get started with this package by **building and running an app that displays `Hello World`**.

## Prerequisite

Look at the `go.mod`/`gop.mod` files, version of Go and Go+ required are:

```
go 1.17
go+ 1.0
```

## Install

Visite the official site to install the required version of Go and Go+ :

[Go](https://go.dev)

[Go+](https://goplus.org)

Create your own project now:

```
mkdir hello && cd hello
gop mod init hello
```

## Code

main.gop

```go

import (
	"net/http"

	"github.com/zrcoder/app"
)

type Hello struct {
	app.Compo
}

func (h *Hello) Render() app.UI {
	return app.h1.text("Hello, world!")
}

app.route "/", &Hello{}
app.runWhenOnBrowser

http.handle "/", &app.Handler{}
println "serving on [http://localhost:9990]"
println http.listenAndServe(":9990", nil)
```

## Build and run

Running a progressive app with **app** requires 2 Go+ programs:

- A client that runs in a web browser
- A server that serves the client and its resources

### Build the client

```bash
GOARCH=wasm GOOS=js gop build -o web/app.wasm
```

Note that the build output is explicitly set to `web/app.wasm`. The reason why is that the `Handler` expects the client to be a **static resource** located at the `/web/app.wasm` path.

Note the build output, you should get gop `builtin` and `github.com/zrcoder/app` dependencies by:

```shell
go get github.com/goplus/gop/builtin
```

And

```shell
go get github.com/zrcoder/app
```

At this point, the package has the following content:

```bash
.
├── go.mod
├── go.sum
├── gop.mod
└── main.gop

```

### Build the server

```bash
gop build
```

### Run the app

Now the client and server built, the package has the following content:

```bash
.
├── go.mod
├── go.sum
├── gop.mod
├── hello
├── main.gop
└── web
    └── app.wasm
```

The server is ran with the following command:

```bash
./hello
```

The app is now accessible from a web browser at http://localhost:9990.

### Use a Makefile

The build process can be simplified by writing a makefile:

```makefile
build:
	GOARCH=wasm GOOS=js gop build -o web/app.wasm
	gop build

run: build
	./hello
```

It can now be built and run with this single command:

```bash
make run
```

