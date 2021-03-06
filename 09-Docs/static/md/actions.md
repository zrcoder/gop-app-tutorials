## What is an Action?

**An `action` is a custom event propagated across the app**. It can be asynchronously handled in a separate goroutine or by any component.

![Actions diagram](/web/static/images/actions.svg)

## Create

An action is created from a `Context` by calling `NewAction(name string, tags ...Tagger)`:

```go
func (h *hello) onInputChange(ctx app.Context, e app.Event) {
	ctx.newAction "greet"
}
```

A payload can also be attached with `NewActionWithValue(name string, v interface{}, tags ...Tagger):`

```go
func (h *hello) onInputChange(ctx app.Context, e app.Event) {
	name := ctx.jsSrc.get("value").string

	ctx.newActionWithValue "greet", name
}
```

A bit like an HTTP header, additional info can be attached to actions by setting tags:

```go
func (h *hello) onInputChange(ctx app.Context, e app.Event) {
	name := ctx.jsSrc.get("value").string

	ctx.newActionWithValue("greet", name,
		app.T("source", "input"),
		app.T("event", "change"),
	)
}
```

## Handling

Once an `action` is created, it is propagated across the app. It can then be handled at global and/or component levels with an `ActionHandler`:

```go
type ActionHandler func(Context, Action)
```

### Global Level

Dealing with actions at a global level is done by registering an `ActionHandler`  with the `Handle` function:

```go
app.handle "greet", handleGreet // Registering action handler.

app.route "/", &hello{}
app.runWhenOnBrowser
// ...

// Action handler that is called on a separate goroutine when a "greet" action is created.
func handleGreet(ctx app.Context, a app.Action) {
	name, ok := a.Value.(string) // Checks if a name was given.
	if !ok {
		println "Hello, World"
		return
	}
	println "Hello,", name
}
```

**Executed asynchronously on a separate goroutine**, handling an action globally is **used to centralize and separate the business logic from the UI**.

### Component Level

Actions can also be handled at component level by registering an `ActionHandler` from a `Context`:

```go
func (h *hello) OnMount(ctx app.Context) {
	ctx.Handle("greet", handleGreet) // Registering action handler.
}

func (h *hello) onChange(ctx app.Context, e app.EventHandler) {
	name := ctx.JSSrc().Get("value").String()
	ctx.NewActionWithValue("greet", name) // Creating "greet" action.
}

// Action handler that is called on the UI goroutine when a "greet" action is
// created.
func (h *hello) handleGreet(ctx app.Context, a app.Action) {
	name, ok := a.Value.(string) // Checks if a name was given.
	if !ok {
		return
	}
	h.name = name
}
```

**Executed on the UI goroutine**, handling actions from components can help **to send data from a component to another**.
