package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("handler http request")
	fmt.Fprintf(w, "Hello %s", r.URL.Path[1:])
}

func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "server port")
	flag.Parse()

	http.HandleFunc("/", handler)
	log.Println("http server run on " + port + " port")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
