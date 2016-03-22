package main

import (
	"io"
	"net/http"
)

func healthcheck(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "{\"status\": \"ok\"}")
}

func main() {
	http.HandleFunc("/healtchcheck", healthcheck)
	http.ListenAndServe(":8888", nil)
}
