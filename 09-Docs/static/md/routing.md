## Intro

Routing is about **associating a component with an URL path**.

## Define a route

Defining a route is done by **associating a URL path with a given component type**.

When a page is requested, its URL path is compared with the defined routes. Then **a new instance of the component type associated with the route is created and displayed**.

Routes are defined by using a simple pattern or by a regular expression.

### Simple route

Simple routes are when a component type matches an exact URL path. They are defined with the `Route()` function:

```go
app.route "/", &hello{}  // hello component type is associated with default path "/".
app.route "/foo", &foo{} // foo component type is associated with "/foo".
app.runWhenOnBrowser    // Launches the app when in a web browser.
```

### Route with regular expression

Routes with regular expressions are when a component type matches an URL path with a given pattern. They are defined with the `RouteWithRegexp()`function:

```go
app.routeWithRegexp "^/bar.*", &bar // bar component is associated with all paths that start with /bar.
app.runWhenOnBrowser               // Launches the app when in a web browser.
```

Regular expressions follow [Go standard syntax](https://github.com/google/re2/wiki/Syntax).

## How it works?

Progressive web apps created with the **go-app** package are working as a [single page application](https://en.wikipedia.org/wiki/Single-page_application). At first navigation, the app is loaded in the browser. Once loaded, each time a page is requested, the navigation event is intercepted and **go-app**'s routing mechanism reads the URL path, then loads a new instance of the associated [component](/components).

![routing.png](/web/static/images/routing.svg)

## Detect navigation

Some scenarios may require additional actions to be done when a page is navigated on. Components can detect when a page is navigated on by implementing the `Navigator` interface:

```go
type foo struct {
    app.Compo
}

func (f *foo) OnNav(ctx app.Context) {
    println "component navigated:", u
}
```

See [component lifecycle](/components#nav).