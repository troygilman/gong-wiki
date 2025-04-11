---
{ "label": "Routing", "position": 5 }
---

# Routing

Gong provides a powerful and flexible routing system that supports both simple and nested routing patterns. This documentation covers the core routing concepts and how to implement them in your Gong applications.

## Basic Routing

Create a basic route using the `NewRoute` function:

```go
gong.NewRoute("/", homeComponent)
```

This renders the `homeComponent` when the root path ("/") is accessed.

## Nested Routing

Create hierarchical routes with the `WithRoutes` method:

```go
gong.NewRoute("/", homeComponent).WithRoutes(
    gong.NewRoute("users", listComponent),
    gong.NewRoute("user/{name}", userComponent),
)
```

In this example:

- The root route ("/") renders the `homeComponent`
- "/users" renders the `listComponent`
- "/user/{name}" renders the `userComponent` with a dynamic parameter

### Outlets

Outlets define where child route components are rendered:

```go
templ (view homeView) View() {
    <div>
        HOME
        <div>
            @gong.NewLink("/users") {
                Users
            }
        </div>
        @gong.NewOutlet()
    </div>
}
```

Key features:

- Automatically render child routes based on the current path
- Can be targeted by links for partial page updates

## Navigation

Use the `Link` component for client-side navigation:

```go
// Basic link - targets an outlet in the current component
gong.NewLink("/users") {
    Users
}

// Link targeting the closest parent outlet
gong.NewLink("/users").WithClosestOutlet() {
    View User
}
```

The `Link` component:

- Uses HTMX for client-side navigation
- By default, looks for an outlet in the current component
- With `WithClosestOutlet()`, targets the closest parent outlet
- Maintains browser history
- Updates only the necessary parts of the page

## Route Parameters

Define dynamic route parameters using the `{param}` syntax:

```go
// Route definition
gong.NewRoute("user/{name}", userComponent)

// Access parameter in your component
name := gong.FormValue(ctx, "name")
```
