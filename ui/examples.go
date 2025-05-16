package ui

import (
	"github.com/troygilman/gong"
	"github.com/troygilman/gong/example/bulk_update"
	"github.com/troygilman/gong/example/click_to_edit"
	"github.com/troygilman/gong/example/tabs"
)

type Example struct {
	path  string
	route gong.Route
}

var examples = []Example{
	{
		path:  "/click-to-edit",
		route: click_to_edit.Route(),
	},
	{
		path:  "/tabs",
		route: tabs.Route(),
	},
	{
		path:  "/bulk-update",
		route: bulk_update.Route(),
	},
}

func ExampleRoute() gong.Route {
	routes := []gong.Route{}

	for _, example := range examples {
		routes = append(routes,
			gong.NewRoute(example.path, gong.NewComponent(OutletComponent{}), gong.WithChildren(example.route)),
		)
	}

	return gong.NewRoute("/example", gong.NewComponent(OutletComponent{}), gong.WithChildren(routes...))
}
