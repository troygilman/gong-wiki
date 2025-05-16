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
type CounterComponent struct {}

templ (c CounterComponent) View() {
	@gong.Target() {
		@counter(0)
	}
}

templ (c CounterComponent) Action() {
	{{
		count, err := strconv.Atoi(gong.FormValue(ctx, "count"))
		if err != nil {
			return err
		}
	}}
	@counter(count+1)
}

templ counter(count int) {
	<p>Count: { strconv.Itoa(count) }</p>
	@gong.Button() {
		Increment
		<input type="hidden" name="count" value={ strconv.Itoa(count) }/>
	}
}
```
