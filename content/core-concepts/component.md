---
{ "label": "Component", "position": 4 }
---

# Component

Gong's server-side component based design allows for modular UI architectures that previously would only be possible with a front-end solution. The component pattern can be used accross your application to structurally package UI and reactivity in one elegant solution.

## Interfaces

In practice, a Component is a golang type that implements a set of Gong defined interfaces. Gong provides the following interfaces for customizing your components.

### View
The View defines the initial `templ.Component` that is rendered when a user views your application. It is the only interface that is required for all components. This is were the bulk of your UI will be defined.

```go
type View interface {
	View() templ.Component
}
```

### Action

The Action defines how a component should handle user requests. Actions are typically used to update server-side state and replace the existing Target's content. The Action can also be used for lazy-loading data.

```go
type Action interface {
	Action() templ.Component
}
```

### Loader

The Loader defines a data loading operation. This is typically used for expensive operations that return crucial data for your component. You might want to consider using a Loader if your component is a nested component.

```go
type Loader interface {
	Loader(ctx context.Context) any
}
```

The Loader can be called elsewhere in the component by calling `gong.LoaderData[Type](ctx)`. Gong will try to cast the data returned by the Loader to the `Type` provided. Gong will panic if the types are not compatible. Do not call the Loader function directly from your component.

### Head

The Head can define your own HTML `<head>` tag that will replace Gong's default `<head>` tag.

```go
type Head interface {
	Head() templ.Component
}
```

If you do this, make sure to add this script tag which loads the HTMX library onto the client.

```html
<script src="https://unpkg.com/htmx.org@2.0.4" integrity="sha384-HGfztofotfshcF7+8n44JQL2oJmowVChPTg48S+jvZoztPfvwD79OC/LTtG6dMp+" crossorigin="anonymous"></script>
```

## Nested Components
