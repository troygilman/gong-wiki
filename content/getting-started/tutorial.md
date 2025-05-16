---
{ "label": "Tutorial", "position": 3 }
---

# Tutorial

All Gong web apps are built using a combination of Components and Routes. A Component defines an HTML view and a set of back-end actions, while a Route defines the path-based hierarchy of your components.

## A Simple Component

Let's start with a simple component which renders static HTML in the browser.

Create a file named `simple.templ` and add the following templ code:

```go
type SimpleComponent struct {}

templ (c SimpleComponent) View() {
	<div>
		Hello World
	</div>
}
```

Run the templ command line compiler to generate native Go code from your templ file:

```bash
templ generate
```

This generates a file named `simple_templ.go` in the same directory as your templ file.

Now define a Route that uses your `SimpleComponent`. A Route requires a path and a View. At a minimum, components must implement the View interface, though additional Gong interfaces can be implemented for more dynamic behavior.

```go
func main() {
	simpleComponent := gong.NewComponent(SimpleComponent{})

	svr := gong.NewServer()
	svr.Route(gong.NewRoute("/", simpleComponent))

	if err := svr.Run(":8080"); err != nil {
		panic(err)
	}
}
```

This main function sets up an Gong HTTP server on port `8080` that routes all requests to your `SimpleComponent`. Run the server:

```bash
go run .
```

You should now see "Hello World" when you navigate to `localhost:8080`.

## Adding an Action

Let's add interactivity by allowing users to submit a request that prints "Hello Universe" to the server's console output.

Add a `Button` to your component:

```go
templ (c SimpleComponent) View() {
	Hello World
	@gong.Button() {
		Submit
	}
}
```

Next, implement the `Action` interface on your SimpleComponent:

```go
templ (c SimpleComponent) Action() {
	{{
		fmt.Println("Hello Universe")
	}}
}
```

After compiling and running your app, clicking the button will print "Hello Universe" to the server's console.

### How Actions Work

When a user clicks the button, an AJAX request is sent to the server, and Gong routes this request to the SimpleComponent's `Action` method. A request issued by a component's `Button` is always received by that same component's `Action`.

## Dynamic Content Swapping

Now, let's make the text change from "Hello World" to "Hello Universe" when the user clicks the button.

Wrap your text with a `Target`:

```go
templ (c SimpleComponent) View() {
	@gong.Target() {
		Hello World
	}
	@gong.Button() {
		Submit
	}
}
```

Update the `Action` to render the new text:

```go
templ (c SimpleComponent) Action() {
	{{
		fmt.Println("Hello Universe")
	}}
	Hello Universe
}
```

Now when the user clicks the button, the server responds with "Hello Universe", and the client swaps this content into the `Target` element.

This completes the basics of building reactive web apps with Gong. Using this pattern of Components, Routes, Buttons, Forms, and Targets, you can create dynamic, interactive web applications.
