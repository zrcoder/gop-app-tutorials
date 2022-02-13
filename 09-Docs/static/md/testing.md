## Intro

Testing is an essential step to achieve app reliability. Since go-app is working on 2 different environments (web browser and server), it provides 2 testing `dispatchers` to emulate [components lifecycle](/components#lifecycle-events) behaviors.

## Component server prerendering

[Prerendering](/components#prerender) is a component lifecycle step where a component can be initialized on the server-side before being converted into HTML. The server-side environment can be emulated with a dispatcher created with the `NewServerTester()` function.

Here is an example that tests if a component has the expected values after the PreRenderer interface call:

```go
type aTitle struct {
	app.Compo

	title string
}

func (t *aTitle) OnPreRender(ctx app.Context) {
	t.title = "Testing Prerendering"
}

func (t *aTitle) Render() app.UI {
	return app.h1.
		class("title").
		text(t.title)
}

func TestComponentPreRendering(t *testing.T) {
	compo := &aTitle{}

	// Creating the server emulator:
	disp := app.NewServerTester(compo)
	defer disp.Close() // Releases alocated resources.

	if compo.title == "Testing Prerendering" {
		t.Fatal("bad component title:", compo.title)
	}

	// Call OnPreRender() from PreRenderer interface:
	disp.preRender

	// Executes all the queued UI instructions.
	disp.consume

	if compo.title != "Testing Prerendering" {
		t.Fatal("bad component title:", compo.title)
	}
}
```

## Component client lifecycle

Like on the [server-side](#testing-component-server-prerendering), testing a component on the client-side is done by emulating the corresponding environment. On the client-side, it is done with the `NewClientTester()` function.

Here is an example that tests if a component has the expected values after mounting and navigation:

```go
type aTitle struct {
	app.Compo

	title string
}

func (t *aTitle) OnMount(ctx app.Context) {
	t.title = "Testing Mounting"
}

func (t *aTitle) OnNav(ctx app.Context) {
	t.title = "Testing Nav"
}

func (t *aTitle) Render() app.UI {
	return app.h1.
		class("title").
		text(t.title)
}

func TestComponentLifcycle(t *testing.T) {
	compo := &aTitle{}

	disp := app.newClientTester(compo)
	defer disp.Close()

	disp.nav(&url.URL{})
	disp.consume
	if compo.title != "Testing Nav" {
		t.Fatal("bad component title:", compo.title)
	}

}
```

See `ClientDispatcher` for other lifecycle and component extension events.

## Asynchronous operations

Asynchronous operations are started with the context's [Async()](/concurrency#async) method. Once started, they can be awaited during testing with the dispatcher `Wait()` method.

Here is an example that launches a goroutine and modifies a component field:

```go
type aTitle struct {
	app.Compo

	title string
}

func (t *aTitle) Render() app.UI {
	return app.h1.
		class("title").
		text(t.title)
}

func (t *aTitle) setAsyncTitle(ctx app.Context) {
	ctx.async(func() {
		time.Sleep(time.Millisecond * 100)
		t.Defer(func(ctx app.Context) {
			t.title = "Testing Async"
		})
	})
}

func TestComponentAsync(t *testing.T) {
	compo := &aTitle{}

	disp := app.newClientTester(compo)
	defer disp.Close()

	compo.setAsyncTitle(disp.Context()) // Async operation queued.
	disp.Consume()                      // Async operation launched but not completed.
	if compo.title == "Testing Async" {
		t.Fatal("bad component title:", compo.title)
	}

	disp.wait    // Wait for the async operations do complete.
	disp.consume // Apply changes.
	if compo.title != "Testing Async" {
		t.Fatal("bad component title:", compo.title)
	}
}
```

## UI elements

UI elements can be tested with the help of the `TestMatch()` function and the `TestUIDescriptor` struct, by allowing a comparison between matching UI elements.

```go
func TestMatch(tree UI, d TestUIDescriptor) error
```

```go
type TestUIDescriptor struct {
    // The location of the node. It is used by the TestMatch to find the
    // element to test.
    //
    // If empty, the expected UI element is compared with the root of the tree.
    //
    // Otherwise, each integer represents the index of the element to traverse,
    // from the root's children to the element to compare
    Path []int

    // The element to compare with the element targeted by Path. Compare
    // behavior varies depending on the element kind.
    //
    // Simple text elements only have their text value compared.
    //
    // HTML elements have their attribute compared and check if their event
    // handlers are set.
    //
    // Components have their exported field values compared.
    Expected UI
}
```

Here is an example that tests the `h1` content of the Hello component:

```go
type aTitle struct {
	app.Compo

	title string
}

func (t *aTitle) OnMount(ctx app.Context) {
	t.title = "Testing Mounting"
}

func (t *aTitle) Render() app.UI {
	return app.h1.
		class("title").
		text(t.title)
}

func TestUIElement(t *testing.T) {
	compo := &aTitle{}
	disp := app.NewClientTester(compo)
	defer disp.Close()

	app.TestMatch(compo, app.TestUIDescriptor{
		Path:     app.TestPath(0), // Component root.
		Expected: app.H2().Text("Testing Mounting"),
	})
}

```
