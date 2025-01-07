package main

import (
	"bytes"
	"embed"
	"io/fs"
	"net/http"

	"github.com/troygilman0/gong"
	"github.com/troygilman0/gong-wiki/ui"
	"github.com/yuin/goldmark"
)

//go:embed public
var publicFS embed.FS

//go:embed content
var contentFS embed.FS

func main() {
	content := compileMarkdownFS(contentFS, map[string]string{
		"content/getting-started/introduction.md": "introduction",
		"content/getting-started/installation.md": "installation",
	})

	mux := http.NewServeMux()

	mux.Handle("/public/", http.StripPrefix("/", http.FileServer(http.FS(publicFS))))

	g := gong.New(mux)

	g.Route("/", ui.RootView{}, func(r gong.Route) {
		r.Route("getting-started/{item}", ui.DocsView{
			Content: content,
		}, nil)
	})

	http.ListenAndServe(":8080", g)
}

func compileMarkdownFS(fileSystem fs.FS, filePaths map[string]string) map[string]string {
	content := make(map[string]string)
	for filePath, item := range filePaths {
		html, err := compileMarkdown(fileSystem, filePath)
		if err != nil {
			panic(err)
		}
		content[item] = html
	}
	return content
}

func compileMarkdown(fileSystem fs.FS, filePath string) (string, error) {
	source, err := fs.ReadFile(fileSystem, filePath)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := goldmark.Convert(source, &buf); err != nil {
		return "", err
	}

	return buf.String(), nil

}
