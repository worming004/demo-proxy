package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/proxy", repeatHandler)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	mux.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func repeatHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/proxy")
	client := http.DefaultClient
	res, err := client.Get("http://demoendpoint:9090/demo-endpoint/text")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(200)
	res.Write(w)
}
