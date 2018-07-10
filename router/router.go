package router

import (
	"os"
	"net/http"
	"github.com/samoslab/samos-storage-caculator/action"
	"path"
)

type Router struct {
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "" {
		action.IndexAction(w, r)
	} else if r.URL.Path == "/cal" {
		action.CalAction(w, r)
	} else if r.URL.Path == "/static" {
		dir := path.Dir(os.Args[0])
		staticHandler := http.FileServer(http.Dir(dir))
		staticHandler.ServeHTTP(w, r)
	}
}
