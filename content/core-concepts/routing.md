---
{ "label": "Routing", "position": 6 }
---

# Routing

Gong provides a powerful and flexible routing system that supports both simple and nested routing patterns. This documentation covers the core routing concepts and how to implement them in your Gong applications.

## Basic Routing

Create a basic route using the `route.New()` function:

```go
import "github.com/troygilman/gong/route"

route.New("/", HomeComponent)
```

This renders the `HomeComponent` when the root path ("/") is accessed.

### Dynamic Routing

Define dynamic path parameters using the `{name}` syntax.

```go
route.New("/user/{name}", UserComponent)
```

## Nested Routing

Create hierarchical routes with the `route.WithChildren(...Route)` functional option:

```go
route.New("/home", HomeComponent,
	route.WithChildren(
		route.New("/users", UserListComponent),
	),
)
```

In this example, the UserListComponent will be rendered inside the Outlet of the HomeComponent.

### Outlet

Use the `Outlet` component to define where child route components are rendered.

- Automatically render child routes based on the current path
- Can be targeted by links for partial page updates
- Will render the first child route as a default

```go
import "github.com/troygilman/gong/outlet"

templ (c HomeComponent) View() {
    HOME
    @outlet.New()
}
```

### Link

Use the `Link` component for client-side navigation with partial page updates.

- Maintains browser history
- Updates only the necessary parts of the page

```go
import "github.com/troygilman/gong/link"

templ (c HomeComponent) View() {
	@link.New("/users") {
    	Users
	}
}
```
