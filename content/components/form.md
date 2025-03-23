---
{ "label": "Form", "position": 3 }
---

# Form

Form is a wrapper for HTML form elements that integrates with HTMX for handling form submissions with minimal JavaScript. Form makes it easy to create forms that use HTMX attributes for AJAX submissions. It supports various HTTP methods, CSS classes, and HTMX-specific configurations.

## Templ Component

The component generates an HTML form with appropriate HTMX attributes:

- `hx-post`, `hx-patch`, or `hx-delete` based on the method
- `hx-swap` for DOM update strategy
- `hx-target` for the element to update (if swap is not `SwapNone`)
- `hx-headers` for additional request metadata
- `class` for CSS styling

## Basic Usage

```go
templ UserForm() {
	@gong.Form() {
		<label for="name">Name:</label>
		<input type="text" id="name" name="name"/>
		<button type="submit">Submit</button>
	}
}
```

## API

### Constructor

#### `Form() FormComponent`

Creates a new form component with default settings:

- Method: POST
- Swap: SwapNone (no HTMX swap behavior)

### Methods

#### `WithMethod(method string) FormComponent`

Sets the HTTP method for the form submission.

**Parameters:**

- `method`: HTTP method string (e.g., "GET", "POST", "PATCH", "DELETE", or use constants from `http` package)

**Returns:**

- Modified `FormComponent` instance

#### `WithCSSClass(cssClass templ.CSSClass) FormComponent`

Applies CSS classes to the form element.

**Parameters:**

- `cssClass`: Class or classes to apply to the form

**Returns:**

- Modified `FormComponent` instance

#### `WithTargetID(targetID string) FormComponent`

Sets the target element ID for HTMX to update after form submission.

**Parameters:**

- `targetID`: ID of the element to update with the response

**Returns:**

- Modified `FormComponent` instance

#### `WithSwap(swap string) FormComponent`

Sets the HTMX swap method for updating the DOM.

**Parameters:**

- `swap`: HTMX swap method (e.g., "innerHTML", "outerHTML", "beforebegin")

**Returns:**

- Modified `FormComponent` instance

## Notes

- The form automatically uses the current request URI as the submission endpoint
- HTMX headers are automatically built using the `buildHeaders` function with `GongRequestTypeAction`
- Target IDs are processed through `buildComponentID` to ensure proper scoping
