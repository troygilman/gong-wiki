package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/troygilman/gong"
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

	docManager, err := ui.NewDocumentManager(contentFS)
	if err != nil {
		panic(err)
	}

	// COMPONENTS
	rootComponent := gong.NewComponent(ui.RootView{})
	docComponent := gong.NewComponent(ui.DocumentView{
		DocManager: docManager,
	})

	// ROUTES
	mux := http.NewServeMux()

	mux.Handle("/public/", http.StripPrefix("/", http.FileServer(http.FS(publicFS))))

	g := gong.New(mux).Routes(
		gong.NewRoute("/", rootComponent).WithRoutes(
			gong.NewRoute("docs/", docComponent),
		),
	)

	http.ListenAndServe(":8080", g)
}
