package utils

import (
	"net/http"
	"path/filepath"
	"runtime"
	"strings"
)

func PageNotFound(w http.ResponseWriter, r *http.Request, status int) {
	LogWarn("Page not found", r)
	w.WriteHeader(status)
	RenderErrorTemplate(w, r, "error", nil)
}

func FunctionName() string {
	pc, _, _, _ := runtime.Caller(1)
	nameFull := runtime.FuncForPC(pc).Name() // main.foo
	nameEnd := filepath.Ext(nameFull)        // .foo
	name := strings.TrimPrefix(nameEnd, ".") // foo
	return name
}
