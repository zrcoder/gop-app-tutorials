## Intro

Since WebAssembly is browser-based technology, some scenarios may require DOM access and JavaScript calls.

This is usually done with the help of [syscall/js](https://golang.org/pkg/syscall/js/) but for compatibility and tooling reasons, **go-app wraps the JS standard package**. Interacting with JavaScript is done by using the `Value` interface.

This article provides examples that show common interactions with JavaScript.

## Include JS files

Building UIs can sometimes require the need of third-party JavaScript libraries. Those libraries can either be included at the [page](/architecture#html-pages) level or inlined in a [component](/components).

### Page's scope

JS files can be included on a page by using the `Handler` `Scripts` field:

```go
handler := &app.Handler{
	Name: "My App",
	Scripts: []string{
		"/web/myscript.js",                // Local script
		"https://foo.com/remoteScript.js", // Remote script
	},
}
```

Or by directly putting JS markup in the `RawHeaders` field:

```go
handler := &app.Handler{
	Name: "My App",
	RawHeaders: []string{
		`<!-- Global site tag (gtag.js) - Google Analytics -->
		<script async src="https://www.googletagmanager.com/gtag/js?id=UA-xxxxxxx-x"></script>
		<script>
		  window.dataLayer = window.dataLayer || [];
		  function gtag(){dataLayer.push(arguments);}
		  gtag('js', new Date());

		  gtag('config', 'UA-xxxxxx-x');
		</script>
		`,
	},
}
```

### Inlined in Components

JS files can also be included directly inlined into [components](/components) in the `Render()` method by using the `\<script\>` HTML element.

The following example asynchronously loads a YouTube video into an `<iframe>`, using a YouTube JavaScript file:

```go
type youtubePlayer struct {
	app.Compo
}

func (p *youtubePlayer) Render() app.UI {
	return app.div.body(
		app.script.
			src("//www.youtube.com/iframe_api").
			async(true),
		app.iFrame.
			ID("yt-container").
			allow("autoplay").
			allow("accelerometer").
			allow("encrypted-media").
			allow("picture-in-picture").
			sandbox("allow-presentation allow-same-origin allow-scripts allow-popups").
			src("https://www.youtube.com/embed/LqeRF_0DDCg"),
	)
}
```

## Using window global object

The `window` JS global object is usable from the `Window` function.

```go
app.window
```

### Get element by ID

`GetElementByID()` is to get a DOM element from an ID.

```js
// JS version:
let elem = document.getElementById("YOUR_ID");
```

```go
// Go equivalent:
elem := app.window.getElementByID("YOUR_ID")
```

It is a helper function equivalent to:

```go
elem := app.window.
    get("document").
    call("getElementById","YOUR_ID")
```

### Create JS object

Creating an object from a library is done by getting its name from the `Window` and call the `New()` function.

Here is an example about how to create a Youtube player:

```js
// JS version:
let player = new YT.Player("yt-container", {
  height: "390",
  width: "640",
  videoId: "M7lc1UVf-VE",
});
```

```go
// Go equivalent:
player := app.window.
	get("YT").
	get("Player").
	New("yt-container", map[string]interface{}{
		"height":  390,
		"width":   640,
		"videoId": "M7lc1UVf-VE",
    })
```

## Cancel an event

When implementing an `event handler`, the event can be canceled by calling `PreventDefault()`.

```go
type foo struct {
	app.Compo
}

func (f *foo) Render() app.UI {
	return app.div.
		onChange(f.onContextMenu).
		text("Don't copy me!")
}

func (f *foo) onContextMenu(ctx app.Context, e app.Event) {
	e.preventDefault
}
```

## Get input value

Input are usually used to get a user inputed value. Here is how to get that value when implementing an `event handler`:

```go
type foo struct {
    app.Compo
}

func (f *foo) Render() app.UI {
    return app.input.onChange(f.onInputChange)
}

func (f *foo) onInputChange(ctx app.Context, e app.Event) {
    v := ctx.jsSrc.get("value").string
}
```