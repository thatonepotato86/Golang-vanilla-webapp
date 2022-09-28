package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	var port = os.Getenv("port")

	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe("localhost:"+port, nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<h1 style=\"font-family: Muli\">Bom Bom Ding Ping &#9889<h1>"))
}
