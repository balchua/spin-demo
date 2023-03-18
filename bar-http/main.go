package main

import (
	"fmt"
	"html"
	"net/http"

	spinconfig "github.com/fermyon/spin/sdk/go/config"
	spinhttp "github.com/fermyon/spin/sdk/go/http"
)

func barHandler(w http.ResponseWriter, r *http.Request) {
	appText, err := spinconfig.Get("app_text")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "missing config app_text")
		return
	}
	fmt.Fprintf(w, "Hello, %q, appText: %s", html.EscapeString(r.URL.Path), html.EscapeString(appText))
}

func init() {
	mux := http.NewServeMux()
	mux.HandleFunc("/bar", barHandler)
	spinhttp.Handle(mux.ServeHTTP)
}

func main() {}
