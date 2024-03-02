package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch, please send an email to<a href=\"mailto:support@appsavvytalk.com\">support@appsavvytalk.com</a>.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Sorry! We could not find what you were looking for")
	}

}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
