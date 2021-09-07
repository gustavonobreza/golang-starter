package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", YourHandler).Methods("GET")
	router.HandleFunc("/prod/{id}", ProdHandler).Methods("GET")

	println("App is running on http://localhost")
	http.ListenAndServe(":80", router)
}

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	http.ServeFile(w, r, "./public/index.html")
}

func ProdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, ok := vars["id"]

	if !ok {
		panic("error: id is not provided")
	}

	res := fmt.Sprintf(`<h1>Product where id is #%v</h1>`, id)

	w.Write([]byte(res))
}
