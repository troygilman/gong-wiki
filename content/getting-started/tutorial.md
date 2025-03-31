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

Now we can define a Route that uses our `SimpleComponent`. You will notice that a Route requires a path and a View. This means that at a minimum, our components must implement the View interface. We can implement additional Gong interfaces to give our components more dynamic behavior, but we will come back to that later.

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

You should now see the words `Hello World` rendered when you navigate to `localhost:8080`.

Great job making it this far. Next, we will be adding some interactivity to our app. Get excited!

## Adding Actions

Our app is pretty boring at the moment. Lets have the text `Hello World` change to `Hello Universe` when the user clicks on a button.

To get this behavior, we can wrap our text with a Target and add a Form with a button inside.

```go
templ (component SimpleComponent) View() {
	@gong.NewTarget() {
		Hello World
	}
	@gong.NewForm() {
		<button>
			Change
		</button>
	}
}
```

Next, we need to implement the Action interface for our `SimpleComponent`. Our action will simply render the text `Hello Universe`.

```go
templ (component SimpleComponent) Action() {
	Hello Universe
}
```

When the user clicks the button, an AJAX request is sent to the server and Gong will route the request to our `SimpleComponent`'s action method. The action method will respond to the request with the HTML text `Hello Universe`. Finally, the client will take the HTML from the response and swap it in for the Target's content.

There you go! Now you know the basics of building reactive web apps with Gong. Using this pattern, you can build some pretty amazing things.
