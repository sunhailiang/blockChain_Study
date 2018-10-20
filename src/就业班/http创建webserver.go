package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/first", test)
	http.ListenAndServe("127.0.0.1:8001", nil)
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("AAAAAAAAABBBBBBBBBBBBB"))
}
