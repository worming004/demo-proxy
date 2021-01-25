package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/proxy", repeatHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func repeatHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/proxy")
	client := http.DefaultClient
	res, err := client.Get("http://demo-endpoint:9090/demo-endpoint/text")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(200)
	res.Write(w)
}
