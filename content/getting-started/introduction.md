---
{ "position": 1 }
---

# Introduction

## Building Modern Web Apps with Gong

Gong is a lightweight framework for building responsive web applications using Go, Templ, and HTMX. Drawing inspiration from the design of [Remix](https://remix.run/), Gong provides an opinionated yet straightforward approach to web development.

- **Component-Based Architecture**: Build reusable UI elements for better maintainability and testing
- **Client-Side Routing**: Enable smooth transitions while preserving browser history
- **Server-Side Rendering**: Deliver optimized performance and SEO with SSR

```go
type usersView struct {
	db       *userDatabase
}

templ (view usersView) Action() {
	switch gong.GetRequest(ctx).Method {
		case http.MethodGet:
			{{
			users := view.db.ReadAll()
			}}
			for _, name := range users {
				<div>{ name }</div>
			}
		case http.MethodPost:
			{{
			name := gong.GetParam(ctx, "name")
			view.db.Create(name)
			}}
			<div>{ name }</div>
	}
}

templ (view usersView) View() {
	<div>
		@gong.Form(gong.FormWithSwap(gong.SwapBeforeEnd)) {
			<input name="name" type="text"/>
			<button type="submit">Add</button>
		}
		@gong.Target(gong.TargetWithTrigger(gong.TriggerLoad))
	</div>
}
```
