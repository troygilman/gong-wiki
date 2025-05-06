---
{ "label": "Components", "position": 5 }
---

# Components

## Introduction

Components are the building blocks of Gong applications. They encapsulate the UI and its behavior, making it easier to manage and reuse code. By using components, developers can create modular, maintainable, and scalable applications.

### Why Use Components?

- **Modularity:** Break down complex UIs into smaller, manageable pieces.
- **Reusability:** Share components across different parts of the application.
- **Maintainability:** Simplify debugging and updates by isolating functionality.

## Interfaces

A Component in Gong is a Go type that implements one or more Gong-defined interfaces. Gong provides the following interfaces for customizing your components:

### View

The View interface defines the initial `templ.Component` that renders when a user accesses your application. It is the only required interface for all components and is where most of your UI will be defined.

```go
type View interface {
	View() templ.Component
}
```

**Example:**

```go
type SimpleViewComponent struct {}

templ (c SimpleViewComponent) View() {
	<div>
		Hello, Gong!
	</div>
}
```

### Action

The Action interface defines how a component handles user interactions. Actions typically update server-side state and replace existing Target content. Actions can also be used for lazy-loading data.

```go
type Action interface {
	Action() templ.Component
}
```

**Example:**

```go
import (
	"github.com/troygilman/gong/button"
)

type ButtonComponent struct {}

templ (c ButtonComponent) Action() {
	{{
		fmt.Println("Button clicked!")
	}}
}

templ (c ButtonComponent) View() {
	@button.New() {
		Click Me
	}
}
```

### Loader

The Loader interface defines data loading operations. This is useful for expensive operations that fetch crucial data for your component, particularly important for nested components.

```go
type Loader interface {
	Loader(ctx context.Context) any
}
```

**Example:**

```go
import (
	"github.com/troygilman/gong/hooks"
)

type DataLoaderComponent struct {}

func (c DataLoaderComponent) Loader(ctx context.Context) any {
	return fetchNameFromDB(ctx)
}

templ (c DataLoaderComponent) View() {
	<div>
		{ hooks.LoaderData[string](ctx) }
	</div>
}
```

You can access loader data elsewhere in your component by calling `hooks.LoaderData[Type](ctx)`. Gong will attempt to cast the data returned by the Loader to the specified `Type` and will panic if the types are incompatible. Never call the Loader function directly from your component.

### Head

The Head interface allows you to define a custom `<head>` element that replaces Gong's default head element.

```go
type Head interface {
	Head() templ.Component
}
```

**Example:**

```go
type CustomIndexComponent struct {}

templ (c CustomIndexComponent) View() {}

templ (c CustomIndexComponent) Head() {
	<head>
		<title>Custom Page</title>
	</head>
}
```

When implementing a custom Head, ensure you include this script tag to load the HTMX library:

```html
<script
    src="https://unpkg.com/htmx.org@2.0.4"
    integrity="sha384-HGfztofotfshcF7+8n44JQL2oJmowVChPTg48S+jvZoztPfvwD79OC/LTtG6dMp+"
    crossorigin="anonymous"
></script>
```

The Head will only be used if it is implemented by the Component in the root level Route.

## Nested Components

Components can be nested within other components to create complex UI hierarchies. This approach promotes code reuse and maintainable architecture.

### Configuring

In order to properly configure a nested component, the child component must be set as a publicly accessable field within the parent component.

```go
import (
	"github.com/troygilman/gong"
	"github.com/troygilman/gong/component"
)

type ParentComponent struct {
	Child gong.Component
}

func NewParentComponent(child gong.Component) gong.Component {
	return component.New(ParentComponent{
		Child: child,
	})
}
```

### Rendering

To render the child component, simply render it within your templ function like any other `templ.Component`.

```go
templ (c ParentComponent) View() {
	<div>
		This is the parent!
		@c.Child
	</div>
}
```

#### WithLoaderData()

To render a child component with data from the parent, use the `WithLoaderData(any)` function.
If the child component uses `hooks.LoaderData[Type](ctx)` then the parent defined data will be used.

```go
templ (c ParentComponent) View() {
	<div>
		This is the parent!
		@c.Child.WithLoaderData("Hello Child")
	</div>
}
```

#### WithLoaderFunc()

To render a child component with a loader function from the parent, use the `WithLoaderFunc(LoaderFunc)` function.
If the child component uses `hooks.LoaderData[Type](ctx)` then the parent defined loader will be used.

```go
templ (c ParentComponent) View() {
	<div>
		This is the parent!
		@c.Child.WithLoaderFunc(func(ctx context.Context) any {
			return "Hello Child"
		})
	</div>
}
```
