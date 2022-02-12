## What is go-app?

Go-app is a package for **building [progressive web apps (PWA)](https://developer.mozilla.org/en-US/docs/Web/Progressive_web_apps) with the [Go programming language (Golang)](https://golang.org) and [WebAssembly (Wasm)](https://webassembly.org)**.

Shaping a UI is done by using a **declarative syntax that creates and composes HTML elements only by using the Go programing language**.

**Served with Go standard HTTP model**, apps created with go-app are **SEO friendly, installable, and support offline mode**.



## Declarative Syntax

Go-app uses a declarative syntax so you can **write reusable component-based UI elements just by using the Go programming language**.

```go
// A component that displays a Hello world by composing with HTML elements,
// conditions, and binding.
type hello struct {
	app.Compo

	name string
}

func (h *hello) Render() app.UI {
	return app.Div().Body(
		app.H1().Body(
			app.Text("Hello, "),
			app.If(h.name != "",
				app.Text(h.name),
			).Else(
				app.Text("World!"),
			),
		),
		app.P().Body(
			app.Input().
				Type("text").
				Value(h.name).
				Placeholder("What is your name?").
				AutoFocus(true).
				OnChange(h.ValueTo(&h.name)),
		),
	)
}
```

## Standard HTTP Server

Serving an app built with go-app is done by using the [Go standard HTTP model](https://golang.org/pkg/net/http).

```go
func main() {
    // Go-app component routing (client-side):
	app.Route("/", &hello{})
	app.Route("/hello", &hello{})
	app.RunWhenOnBrowser()

    // Standard HTTP routing (server-side):
	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
```

## Other Features

- Works as a [Single-page application](https://en.wikipedia.org/wiki/Single-page_application)
- SEO friendly
- Installable
- Offline mode support
- State management
