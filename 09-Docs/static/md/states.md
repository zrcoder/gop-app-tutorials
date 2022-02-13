## What is a state?

A state is a value identified by a key, that is available across the app, observable, and concurrency safe.

## Set

A state is set from a `Context` with the `SetState(state string, v interface{}, opts ...StateOption)` method:

```go
// Handling the "greet" action:
func handleGreet(ctx app.Context, a app.Action) {
	name, ok := a.Value.(string)
	if !ok {
		return
	}

	// Setting a state named "greet-name" with the name value.
	ctx.setState("greet-name", name)
}
```

By default a state lives within app memory, It gets deleted when the app is closed. The way a state is set can be modified by using options.

### Options

| Name                              | Description                                                                      | Note                                                                                 |
| --------------------------------- | -------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------ |
| `Persist`     | The state is persisted on local storage, making it available for later sessions. | The value must be compatible with [encoding/json](https://pkg.go.dev/encoding/json). |
| `Encrypt`     | The state is encrypted when persisted on local storage.                          | Requires the use of the `Persist` option.                        |
| `ExpiresIn` | The state is deleted after the given duration.                                   |                                                                                      |
| `ExpiresAt` | The state is deleted at the given time.                                          |                                                                                      |
| `Broadcast` | The state is propagated to other browser tabs and windows.                       | The value must be compatible with [encoding/json](https://pkg.go.dev/encoding/json). |

Options are set by appending the options at the end of the `SetState` method. Here is an example where a state is persisted in local storage and propagated across browsers tabs and windows:

```go
func handleGreet(ctx app.Context, a app.Action) {
	name, ok := a.Value.(string)
	if !ok {
		return
	}

	ctx.SetState("greet-name", name,
		app.Persist,
		app.Broadcast,
	)
}
```

## Observe

Observing a state is to get its value and get notified whenever it is modified with `SetState`. It is done from a `Context` with the `ObserveState` method.

```go
type hello struct {
	app.Compo
	name string
}

func (h *hello) OnMount(ctx app.Context) {
	ctx.observeState("greet-name").value(&h.name)
}
```

`ObserveState` creates an `Observable`. The `Observable.Value` method stores the `"greet-name"` state value into the `name` field, then associates the observable with the state, which will trigger the `name` field update each time the state is modified.

### Conditional Observation

`Observable.While` set a condition to the observation. Here is an example where the `"greet-name"` state will be observed only until a name reaches a length of 5 characters:

```go
func (h *hello) OnMount(ctx app.Context) {
	ctx.ObserveState("greet-name").
		While(func() bool {
			return len(h.name) < 5
		}).
		value(&h.name)
}
```

### Additional Instructions

When a state is modified, `Observable.OnChange` sets additional instructions to be executed after a change occurs:

```go
func (h *hello) OnMount(ctx app.Context) {
	ctx.ObserveState("greet-name").
		onChange(func() {
			fmt.Println("greet-name was changed at", time.Now())
		}).
		value(&h.name)
}
```

## Get

For scenarios where a state value is just to be retrieved without being observed, there is the `Context` `GetState` method:

```go
func handleGreet(ctx app.Context, a app.Action) {
	var name string
	ctx.getState("greet-name", &name)

	// ...
}
```
