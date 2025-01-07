package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/troygilman0/gong"
	"github.com/troygilman0/gong-wiki/ui"
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

	g.Route("/docs", ui.RootView{}, func(r gong.Route) {
		r.Route("/", ui.DocumentView{
			DocManager: docManager,
		}, nil)
	})

	http.ListenAndServe(":8080", g)
}
