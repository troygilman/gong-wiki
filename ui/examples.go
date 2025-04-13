package ui

import (
	"github.com/troygilman/gong"
	"github.com/troygilman/gong/example/click_to_edit"
	"github.com/troygilman/gong/example/nested_components"
	"github.com/troygilman/gong/example/tabs"
)

type Example struct {
	path  string
	route gong.RouteBuilder
}

var examples = []Example{
	{
		path:  "/nested-components",
		route: nested_components.Route(),
	},
	{
		path:  "/click-to-edit",
		route: click_to_edit.Route(),
	},
	{
		path:  "/tabs",
		route: tabs.Route(),
	},
}

func ExampleRoute() gong.RouteBuilder {
	routes := []gong.RouteBuilder{}

	for _, example := range examples {
		routes = append(routes, gong.NewRoute(example.path, gong.NewComponent(OutletComponent{})).WithRoutes(example.route))
	}

	return gong.NewRoute("/example", gong.NewComponent(OutletComponent{})).WithRoutes(routes...)
}
