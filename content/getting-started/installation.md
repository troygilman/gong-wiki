---
{ "label": "Installation", "position": 2 }
---

# Installation

## Dependencies

### go

Gong requires go v1.22 or higher.

### templ

Install templ by running the following command or by following the detailed instructions found on the [templ website](https://templ.guide/quick-start/installation/).

```bash
go install github.com/a-h/templ/cmd/templ@latest
```

Templ is the HTML templating language at the core of Gong's components based design. This documentation assumes the reader has a proficient understanding of templ and basic HTML.

## Adding Gong to a project

Gong can be added to an existing go project by running the following command.

```bash
go get github.com/troygilman/gong@latest
```
