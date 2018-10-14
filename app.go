package main

import (
	"log"
	"net/http"
	"fmt"
)

var remoteAddr string = ""

func curlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	remoteAddr = r.RemoteAddr
	w.Write([]byte("The remoteAddr is: " + remoteAddr))
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("The remoteAddr is: " + remoteAddr))
}

func main() {
	ch := http.HandlerFunc(curlHandler)
	get := http.HandlerFunc(getHandler)

	http.Handle("/curl", ch)
	http.Handle("/get", get)

	log.Println("Listening 3000... ")
	http.ListenAndServe(":3000", nil)
}
