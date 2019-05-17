package main

import (
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ola mundo!"))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		panic("$PORT not set")
	}
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
