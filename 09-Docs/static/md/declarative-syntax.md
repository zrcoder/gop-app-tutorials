## Intro

The go-app declarative syntax is to customize [components](/components)' look.

It uses a chaining mechanism made from the [Go programming language](https://golang.org) syntax that allows composing HTML elements and components in order to craft beautiful and usable UIs.

Here is an example where HTML elements are used to display a title and a paragraph:

```go
func (c *myCompo) Render() app.UI {
	return app.Div().Body(
		app.h1.class("title").text("Build a GUI with Go"),
		app.p.class("text").text("Just because Go and this package are really awesome!"),
	)
}
```

## HTML Elements

Go-app provides interfaces for each standard HTML element. Those interfaces describe setters for attributes and event handlers.

Here is a simplified version of the interface for a `<div>`:

```go
type HTMLDiv interface {
    // Attributes:
    Body(children ...UI) HTMLDiv
    Class(v string) HTMLDiv
    ID(v string) HTMLDiv
    Style(k, v string) HTMLDiv

    // Event handlers:
    OnClick(h EventHandler, scope ...interface{}) HTMLDiv
    OnKeyPress(h EventHandler, scope ...interface{}) HTMLDiv
    OnMouseOver(h EventHandler, scope ...interface{}) HTMLDiv
}
```

### Create

An HTML element is created by calling a function named after its name. The example below shows how to create a `\<div>`:

```go
func (c *myCompo) Render() app.UI {
	return app.div
}
```

### Standard Elements

A standard HTML element is an element that can contain other UI elements. Other HTML elements, texts, and [components](/components) are nested by using the `Body()` method:

```go
func (c *myCompo) Render() app.UI {
	return app.Div().Body(       // Div Container
		app.h1.text("Title"),  // First child
		app.p.text("Content"), // Second child
	)
}
```

### Self Closing Elements

A self-closing element is an HTML element that cannot contain other UI elements.

```go
func (c *myCompo) Render() app.UI {
	return app.img.src("/myImage.png")
}
```

### Attributes

HTML element interfaces provide methods to set element attributes. Here is an example that set a `<div>`'s class:

```go
func (c *myCompo) Render() app.UI {
	return app.div.class("my-class")
}
```

Multiple attributes are set by using the chaining mechanism:

```go
func (c *myCompo) Render() app.UI {
	return app.div.ID("id-name").class("class-1").class("class-2")
}
```

### Style

Style is an attribute that sets the element style with CSS.

```go
func (c *myCompo) Render() app.UI {
	return app.div.style("width", "400px")
}
```

Like the `Class()` attribute, multiple styles are set by using the chaining mechanism:

```go
func (c *myCompo) Render() app.UI {
	return app.div.
  	style("width", "400px").
  	style("height", "200px").
  	style("background-color", "deepskyblue")
}
```

### Event handlers

`Event handlers` are functions that are called when an HTML event occurs. They must have the following signature:

```go
func(ctx app.Context, e app.Event)
```

Like attributes, HTML element interfaces provide methods to associate an event to a given handler:

```go
func (c *myCompo) Render() app.UI {
	return app.div.onClick(c.onClick)
}

func (c *myCompo) onClick(ctx app.Context, e app.Event) {
	println "onClick is called"
}
```

The `Context` argument embeds several go-app tools that help in creating responsive UIs. Usable with any function accepting a [Go standard context](https://golang.org/pkg/context/#Context), it is canceled when the source of the event is dismounted. The source element value can be retrieved with the JSSrc field:

```go
func (c *myCompo) Render() app.UI {
	return app.div.onChange(c.onChange)
}

func (c *myCompo) onChange(ctx app.Context, e app.Event) {
	v := ctx.jsSrc.get("value")
}
```

`ctx.jSSrc` and `Event` are `JavaScript objects wrapped in Go interfaces`.

## Text

`Text()` represents simple HTML text. Here is an example that display a `Hello World` text:

```go
func (c *myCompo) Render() app.UI {
	return app.div.body( // Container
		app.text("Hello"), // First text
		app.text("World"), // Second text
	)
}
```

When an HTML element embeds a single text element, HTML element's `Text()` method can be used instead:

```go
func (c *myCompo) Render() app.UI {
	return app.div.text("Hello World")
}
```

## Raw elements

`Raw elements` are elements representing plain HTML code. Be aware that using them is **unsafe since there is no check on HTML format**.

Here is an example that creates a `<svg>` element.

```go
func (c *myCompo) Render() app.UI {
	return app.raw(`
	<svg width="100" height="100">
		<circle cx="50" cy="50" r="40" stroke="green" stroke-width="4" fill="yellow" />
	</svg>
	`)
}
```

## Nested Components

[Components](/components) are structs that let you split the UI into independent and reusable pieces. They can be used within other components to achieve more complex UIs.

Here is an example where a component named `foo` embeds a [text](#text) and another component named `bar`.

```go
// foo component:
type foo struct {
	app.Compo
}

func (f *foo) Render() app.UI {
	return app.P().Body(
		app.text("Foo, "), // Simple HTML text
		&bar{},            // Nested bar component
	)
}

// bar component:
type bar struct {
	app.Compo
}

func (b *bar) Render() app.UI {
	return app.text("Bar!")
}
```

## Condition

A ·Condition· is a construct that selects the UI elements that satisfy a condition. They are created by calling the `If()` function.

### If

Here is an If example that shows a title only when the `showTitle` value is `true`:

```go
type myCompo struct {
	app.Compo

	showTitle bool
}

func (c *myCompo) Render() app.UI {
	return app.div.body(
		app.If(c.showTitle,
			app.h1.text("hello"),
		),
	)
}
```

### ElseIf

Here is an ElseIf example that shows a title in different colors depending on an `int` value:

```go
type myCompo struct {
	app.Compo

	color int
}

func (c *myCompo) Render() app.UI {
	return app.Div().Body(
		app.If(c.color > 7,
			app.h1.
				style("color", "green").
				text("Good!"),
		).ElseIf(c.color < 4,
			app.h1.
				style("color", "red").
				text("Bad!"),
		).Else(
			app.h1.
				style("color", "orange").
				text("So so!"),
		),
	)
}
```

### Else

Here is an Else example that shows a simple text when the `showTitle` value is `false`:

```go
type myCompo struct {
	app.Compo

	showTitle bool
}

func (c *myCompo) Render() app.UI {
	return app.Div().Body(
		app.If(c.showTitle,
			app.h1.text("hello"),
		).Else(
			app.text("world"), // Shown when showTitle == false
		),
	)
}
```

## Range

Range represents a `range loop` that shows UI elements generated from a `slice` or [map](#map). They are created by calling the `Range()` function.

### Slice

Here is a slice example that shows an unordered list from a `[]string`:

```go
func (c *myCompo) Render() app.UI {
	data := []string{
		"hello",
		"go-app",
		"is",
		"sexy",
	}

	return app.Ul().Body(
		app.Range(data).slice(func(i int) app.UI {
			return app.li.text(data[i])
		}),
	)
}
```

### Map

Here is a map example that shows an unordered list from a `map[string]int`:

```go
func (c *myCompo) Render() app.UI {
	data := map[string]int{
		"Go":         10,
		"JavaScript": 4,
		"Python":     6,
		"C":          8,
	}

	return app.Ul().Body(
		app.Range(data).Map(func(k string) app.UI {
			s := fmt.Sprintf("%s: %v/10", k, data[k])

			return app.li.text(s)
		}),
	)
}
```

## Form helpers

Form helpers are [component](/components) methods that help to map HTML form element values to [component fields](/components#fields).

### ValueTo

`ValueTo` returns an `event handler` which maps an HTML element value property to a given variable.

Here is a Hello component version that uses the `ValueTo()` method to get the username from its input rather than defining an [event handler](/declarative-syntax#event-handlers):

```go
type hello struct {
	app.Compo

	name string
}

func (h *hello) Render() app.UI {
	return app.Div().Body(
		app.h1.text("Hello " + h.name),
		app.p.body(
			app.input.
				type("text").
				value(h.name).
				placeholder("What is your name?").
				autoFocus(true).
				// Here the username is directly mapped from the input's change
				// event.
				onChange(h.valueTo(&h.name)),
		),
	)
}
```
