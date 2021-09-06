package main

import (
	"fmt"
	"net/http"
)

func main() {
	println("Open http://localhost")
	http.HandleFunc("/photo", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "Interstelar.jpg")
		fmt.Println("GET:", "Interstelar photo")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.Path

		fmt.Fprintf(w, `
		<div style="margin: 0 auto;width: 500px;" >
		 	<h1>Hello, your path is: <span style="color: blue;">%s</span></h1>
		 	<p>Please see this photo -> <a href="/photo">Interstelar</a></p>
		 </div>`, path)
		fmt.Println("GET: ", path)
	})
	http.ListenAndServe(":80", nil)
}

// Open http://localhost
