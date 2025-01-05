package main

import (
	"embed"
	"net/http"

	"github.com/troygilman0/gong"
	"github.com/troygilman0/gong-wiki/ui"
)

//go:embed public
var publicFS embed.FS

func main() {
	mux := http.NewServeMux()

	mux.Handle("/public/", http.StripPrefix("/", http.FileServer(http.FS(publicFS))))

	g := gong.New(mux)

	g.Route("/", ui.RootView{}, nil)

	http.ListenAndServe(":8080", g)
}
