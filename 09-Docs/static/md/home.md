## Write front-end with Go+

`app` is a package for **building [progressive web apps (PWA)](https://developer.mozilla.org/en-US/docs/Web/Progressive_web_apps) with the [Go+](https://goplus.org) programming language and [WebAssembly (Wasm)](https://webassembly.org)**.

Shaping a UI is done by using a **declarative syntax that creates and composes HTML elements only by using the Go programing language**.

**Served with Go standard HTTP model**, apps created with go-app are **SEO friendly, installable, and support offline mode**.

## Declarative Syntax

We uses a declarative syntax so you can **write reusable component-based UI elements just by using the Go programming language**.

```go
// A component that displays a Hello world by composing with HTML elements,
// conditions, and binding.
type Hello struct {
	app.Compo
}

func (h *Hello) Render() app.UI {
	return app.h1.text("Hello, world!")
}
```

## Standard HTTP Server

Serving an app built with go-app is done by using the [Go standard HTTP model](https://golang.org/pkg/net/http).

```go
app.route "/", &Hello{}
app.runWhenOnBrowser

http.handle "/", &app.Handler{}
println "serving on [http://localhost:9990]"
println http.listenAndServe(":9990", nil)
```

## Other Features

- Works as a [Single-page application](https://en.wikipedia.org/wiki/Single-page_application)
- SEO friendly
- Installable
- Offline mode support
- State management
