---
{ "label": "Data Binding", "position": 7 }
---

# Data Binding

Gong provides many functions to bind and query data from the current request.

## PathParam()

Use `hooks.PathParam(context.Context, string)` to get a dynamic path paramteter from the current request.

```go
// Route definition
route.New("/user/{name}", UserComponent)

// Access parameter in your component
name := hooks.PathParam(ctx, "name")
```

## QueryParam()

Use `hooks.QueryParam(context.Context, string)` to get a query parameter from the current request.

```go
// URL: https://my-app.com?name=Joe
name := hooks.QueryParam(ctx, "name")
```

## FormValue()

Use `hooks.FormValue(context.Context, string)` to get a form value from the current request.

```go
// Form Data: name=Joe
name := hooks.FormValue(ctx, "name")
```
