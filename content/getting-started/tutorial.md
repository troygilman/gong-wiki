---
{ "label": "Tutorial", "position": 3 }
---

# Tutorial

All Gong web apps are built using a combination of Components and Routes. A Component defines an HTML view and a set of back-end actions. A Route defines the path based heirarchy of your components.

## A Simple Component

Lets start with a simple component which renders a bit of static HTML in the browser.

Create a file named `simple.templ` and add the following templ code.

```go
type SimpleComponent struct {}

templ (component SimpleComponent) View() {
	<div>
		Hello World
	</div>
}
```

Run the templ command line compiler to generate native golang code from our templ file.

```bash
templ generate
```

This will generate a file named `simple_templ.go` in the same directory as our templ file.

Now we can define a Route that uses our `SimpleComponent`. You will notice that a Route requires a path and a View. This means that, at a minimum, our components must implement the View interface. We can implement additional Gong interfaces to give our components more dynamic behavior, but we will come back to that later.

```go
import (
	"net/http"
	"github.com/troygilman/gong"
)

func main() {
	g := gong.New(http.NewServeMux()).Routes(
		gong.NewRoute("/", SimpleComponent{}),
	)

	if err := http.ListenAndServe(":8080", g); err != nil {
		panic(err)
	}
}
```

This main function will set up an http server on port `8080` which serves all requests using our `SimpleComponent`. Run the server using the following command.

```bash
go run .
```

You should now see the words "Hello World" rendered when you navigate to `localhost:8080`.

Great job making it this far. Next, we will be adding some interactivity to our app. Get excited!

## Adding an Action

Our app is pretty boring at the moment. Lets allow the user to submit a request to the server which will print "Hello Universe" to the console output.

To get this behavior, we can add a `Form` with a button inside. Lets use the `NewForm` function to acheive this.

```go
templ (component SimpleComponent) View() {
	Hello World
	@gong.NewForm() {
		<button>
			Submit
		</button>
	}
}
```

Next, we need to implement the `Action` interface on our SimpleComponent. Our action will simply print "Hello Universe" to the servers console. Note that we are using templ's `{{ }}` syntax to define raw golang code that will be executed when rendering the component.

```go
templ (component SimpleComponent) Action() {
	{{
		fmt.Println("Hello Universe")
	}}
}
```

Try this out by compiling the templ files and running your app. You should now see the text "Hello Universe" printed to the server's console output when you click the button.

But how does this work?

When a user clicks the button, the `Form` will send a special AJAX request to the server and Gong will route the request to our SimpleComponent's `Action` method. A request issued by a component's `Form` will always be received by that same components `Action`.

## Swapping Content

It's time to spice things up. Lets have the text "Hello World" change to "Hello Universe" when the user clicks on the button.

To get this behavior, we can wrap our text with a `Target`. Lets use the `NewTarget` function to achieve this. We will also need to add the `WithSwap(gong.SwapInnerHTML)` configuration to our `Form` so that HTMX knows how to handle the servers response.

```go
templ (component SimpleComponent) View() {
	@gong.NewTarget() {
		Hello World
	}
	@gong.NewForm().WithSwap(gong.SwapInnerHTML) {
		<button>
			Submit
		</button>
	}
}
```

Next, we need to update the `Action` to render the text "Hello Universe". We can put it right below the print statement.

```go
templ (component SimpleComponent) Action() {
	{{
		fmt.Println("Hello Universe")
	}}
	Hello Universe
}
```

Now, When the user clicks the button and the request is sent to the server, the `Action` will respond to the request with the HTML text "Hello Universe". Finally, the client will take the HTML from the response and swap it in for the `Target`'s inner HTML content.

There you go! Now you know the basics of building reactive web apps with Gong. Using this simple pattern, you can build some pretty amazing things.
