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

templ (c UsersComponent) View() {
	{{
		users := c.db.ReadAll()
	}}
	<div>
		@form.New().WithSwap(gong.SwapBeforeEnd) {
			<input name="name" type="text"/>
			<button type="submit">Add</button>
		}
		@target.New() {
			for _, name := range users {
				<div>{ name }</div>
			}
		}
	</div>
}

templ (c UsersComponent) Action() {
	{{
		name := hooks.FormValue(ctx, "name")
		c.db.Create(name)
	}}
	<div>{ name }</div>
}
```
