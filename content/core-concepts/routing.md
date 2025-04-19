---
{ "label": "Routing", "position": 6 }
---

# Routing

Gong provides a powerful and flexible routing system that supports both simple and nested routing patterns. This documentation covers the core routing concepts and how to implement them in your Gong applications.

## Basic Routing

Create a basic route using the `NewRoute` function:

```go
gong.NewRoute("/", HomeComponent)
```

This renders the `HomeComponent` when the root path ("/") is accessed.

### Data Binding

Gong provides many functions to bind and query data from the current request.

#### PathValue()

Define dynamic route parameters using the `{name}` syntax.

```go
// Route definition
gong.NewRoute("user/{name}", userComponent)

// Access parameter in your component
name := gong.PathValue(ctx, "name")
```

#### QueryParam()

Use `QueryParam(context.Context, string)` to get a query parameter from the current request.

```go
// URL: https://my-app.com?name=Joe
name := gong.QueryParam(ctx, "name")
```

#### FormValue()

Use `FormValue(context.Context, string)` to get a form value from the current request.

```go
// Form Data: name=Joe
name := gong.FormValue(ctx, "name")
```

## Nested Routing

Create hierarchical routes with the `WithRoutes(...RouteBuilder)` method:

```go
gong.NewRoute("/home", HomeComponent).WithRoutes(
    gong.NewRoute("/users", UserListComponent),
)
```

In this example, the UserListComponent will be rendered inside the Outlet of the HomeComponent.

### Outlets

Outlets define where child route components are rendered.

- Automatically render child routes based on the current path
- Can be targeted by links for partial page updates
- Will render the first child route as a default

```go
templ (c HomeComponent) View() {
    HOME
    @gong.NewOutlet()
}
```

## Navigation

Use the `Link` component for client-side navigation with partial page updates.

- Maintains browser history
- Updates only the necessary parts of the page

```go
@gong.NewLink("/users") {
    Users
}
```
