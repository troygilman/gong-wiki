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

	mux := http.NewServeMux()

	mux.Handle("/public/", http.StripPrefix("/", http.FileServer(http.FS(publicFS))))

	g := gong.New(mux)

	g.Route("/", ui.RootView{}, func(r gong.Route) {
		r.Route("home", ui.HomeView{}, nil)
		r.Route("docs/", ui.DocumentView{
			DocManager: docManager,
		}, nil)
	})

	http.ListenAndServe(":8080", g)
}
