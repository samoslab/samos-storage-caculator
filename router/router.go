package router

import (
	"os"
	"net/http"
	"path/filepath"
	"github.com/samoslab/samos-storage-caculator/action"
)

type Router struct {
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		action.IndexAction(w, r)
	} else if r.URL.Path == "/cal" {
		action.CalAction(w, r)
	} else {
		dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		dir = dir + "/public"
		staticHandler := http.FileServer(http.Dir(dir))
		staticHandler.ServeHTTP(w, r)
	}
}
