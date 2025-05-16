---
{ "label": "Data Binding", "position": 7 }
---

# Data Binding

Gong provides many functions to bind and query data from the current request.

## PathParam()

Use `gong.PathParam(context.Context, string)` to get a dynamic path paramteter from the current request.

```go
// Route definition
gong.NewRoute("/user/{name}", UserComponent)

// Access parameter in your component
name := gong.PathParam(ctx, "name")
```

## QueryParam()

Use `gong.QueryParam(context.Context, string)` to get a query parameter from the current request.

```go
// URL: https://my-app.com?name=Joe
name := gong.QueryParam(ctx, "name")
```

## FormValue()

Use `gong.FormValue(context.Context, string)` to get a form value from the current request.

```go
// Form Data: name=Joe
name := gong.FormValue(ctx, "name")
```
