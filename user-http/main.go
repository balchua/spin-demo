package main

import (
	"fmt"
	"html"
	"net/http"

	spinconfig "github.com/fermyon/spin/sdk/go/config"
	spinhttp "github.com/fermyon/spin/sdk/go/http"
)

func userHandler(w http.ResponseWriter, r *http.Request) {
	_, err := spinconfig.Get("user_token")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "missing config user_token")
		return
	}
	user := r.URL.Query().Get("userName")
	if user == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "user cannot be empty")
		return
	}
	fmt.Fprintf(w, "Welcome user, %q", html.EscapeString(user))
}

func init() {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", userHandler)
	spinhttp.Handle(mux.ServeHTTP)
}

func main() {}
