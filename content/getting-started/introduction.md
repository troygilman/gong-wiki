---
{ "label": "Introduction", "position": 1 }
---

# Introduction

## Building Modern Web Apps with Gong

Gong is a lightweight framework for building responsive web applications using Go, Templ, and HTMX. Drawing inspiration from the design of [Remix](https://remix.run/), Gong provides an opinionated yet straightforward approach to web development.

- **Component-Based Architecture**: Build reusable UI elements for better maintainability and testing
- **Client-Side Routing**: Enable smooth transitions while preserving browser history
- **Server-Side Rendering**: Deliver optimized performance and SEO with SSR

```go
type UsersComponent struct {
	db       UserDatabase
}

templ (component UsersComponent) View() {
	{{
		users := component.db.ReadAll()
	}}
	<div>
		@gong.NewForm().WithSwap(gong.SwapBeforeEnd) {
			<input name="name" type="text"/>
			<button type="submit">Add</button>
		}
		@gong.NewTarget() {
			for _, name := range users {
				<div>{ name }</div>
			}
		}
	</div>
}

templ (component UsersComponent) Action() {
	{{
	name := gong.FormValue(ctx, "name")
	component.db.Create(name)
	}}
	<div>{ name }</div>
}
```
