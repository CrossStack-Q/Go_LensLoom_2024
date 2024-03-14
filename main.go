package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	fmt.Fprintf(w, "<h1>Namaste  Parashar Ji!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	fmt.Fprintf(w, `<h1>
	<a href="mailto:crossrehk@gmail.com">
	crossrehk@gmail.com
	</a>
	</h1>`)
}

func ErrorWalaPage(w http.ResponseWriter, r *http.Request) {
	// http.Error(w, "404 Page Not Found", http.StatusNotFound)
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, http.StatusText(http.StatusNotFound))
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<html><h1>FAQ questions</h1> <ul><li>Pixel vs Iphone</li><ul> </html>`)
}

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		handlerFunc(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		ErrorWalaPage(w, r)
	}
}

func main() {
	var router Router
	fmt.Println("Starting server on :3000")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		panic(err)
	}
}
