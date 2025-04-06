---
{ "label": "Component", "position": 5 }
---

# Component

Gong's server-side component-based design enables modular UI architectures that traditionally required front-end frameworks. The component pattern allows you to package UI elements and their reactivity together in one elegant solution across your application.

## Interfaces

A Component in Gong is a Go type that implements one or more Gong-defined interfaces. Gong provides the following interfaces for customizing your components:

### View

The View interface defines the initial `templ.Component` that renders when a user accesses your application. It is the only required interface for all components and is where most of your UI will be defined.

```go
type View interface {
	View() templ.Component
}
```

### Action

The Action interface defines how a component handles user interactions. Actions typically update server-side state and replace existing Target content. Actions can also be used for lazy-loading data.

```go
type Action interface {
	Action() templ.Component
}
```

### Loader

The Loader interface defines data loading operations. This is useful for expensive operations that fetch crucial data for your component, particularly important for nested components.

```go
type Loader interface {
	Loader(ctx context.Context) any
}
```

You can access loader data elsewhere in your component by calling `gong.LoaderData[Type](ctx)`. Gong will attempt to cast the data returned by the Loader to the specified `Type` and will panic if the types are incompatible. Never call the Loader function directly from your component.

### Head

The Head interface allows you to define a custom HTML `<head>` tag that replaces Gong's default `<head>` tag.

```go
type Head interface {
	Head() templ.Component
}
```

When implementing a custom Head, ensure you include this script tag to load the HTMX library:

```html
<script src="https://unpkg.com/htmx.org@2.0.4" integrity="sha384-HGfztofotfshcF7+8n44JQL2oJmowVChPTg48S+jvZoztPfvwD79OC/LTtG6dMp+" crossorigin="anonymous"></script>
```

## Nested Components

Components can be nested within other components to create complex UI hierarchies. This approach promotes code reuse and maintainable architecture.
