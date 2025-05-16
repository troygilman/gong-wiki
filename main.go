package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/troygilman/gong"
	"github.com/troygilman/gong-wiki/document"
	"github.com/troygilman/gong-wiki/ui"
)

//go:embed public
var publicFS embed.FS

//go:embed content
var contentFS embed.FS

func main() {
	contentFS, err := fs.Sub(contentFS, "content")
	if err != nil {
		panic(err)
	}

	docManager, err := document.NewManager(contentFS)
	if err != nil {
		panic(err)
	}

	// COMPONENTS
	navbarComponent := gong.NewComponent(ui.NavbarComponent{
		SearchComponent: gong.NewComponent(ui.SearchComponent{
			DocumentRepository: docManager.Repository,
		}),
	})

	menuComponent := gong.NewComponent(ui.MenuComponent{
		Props: []ui.MenuGroupProps{
			{
				Label: "Getting Started",
				Items: []ui.MenuItemProps{
					{
						Label: "Introduction",
						Link:  "/docs/getting-started/introduction",
					},
					{
						Label: "Installation",
						Link:  "/docs/getting-started/installation",
					},
					{
						Label: "Tutorial",
						Link:  "/docs/getting-started/tutorial",
					},
					{
						Label: "Examples",
						Link:  "/docs/getting-started/examples",
					},
				},
			},
			{
				Label: "Core Concepts",
				Items: []ui.MenuItemProps{
					{
						Label: "Components",
						Link:  "/docs/core-concepts/components",
					},
					{
						Label: "Routing",
						Link:  "/docs/core-concepts/routing",
					},
					{
						Label: "Data Binding",
						Link:  "/docs/core-concepts/data-binding",
					},
				},
			},
		},
	})

	rootComponent := gong.NewComponent(ui.RootView{
		Menu:   menuComponent,
		Navbar: navbarComponent,
	})

	landingPageComponent := gong.NewComponent(ui.LandingPageComponent{Navbar: navbarComponent})

	// ROUTES
	docRoutes := []gong.Route{}
	for _, path := range docManager.AllPaths() {
		docRoutes = append(docRoutes, newDocumentRoute(docManager, path))
	}

	svr := gong.NewServer()

	svr.Route(gong.NewRoute("/", landingPageComponent))
	svr.Route(gong.NewRoute("/docs", rootComponent, gong.WithChildren(docRoutes...)))
	svr.Route(ui.ExampleRoute())

	svr.Handle("/public/", http.StripPrefix("/", http.FileServer(http.FS(publicFS))))

	if err := svr.Run(":8080"); err != nil {
		panic(err)
	}
}

func newDocumentRoute(manager document.Manager, path string) gong.Route {
	doc := manager.GetByPath(path)
	prev := manager.GetByPosition(doc.Metadata().Position - 1)
	next := manager.GetByPosition(doc.Metadata().Position + 1)
	return gong.NewRoute(path, ui.NewDocumentComponent(doc, prev, next))
}
