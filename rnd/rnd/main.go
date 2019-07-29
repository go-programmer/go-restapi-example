package main

import (
	"log"
	"net/http"
	"time"
)

// func main() {
// 	mux := http.NewServeMux()

// 	rh := http.RedirectHandler("http://example.org", 307)
// 	mux.Handle("/foo", rh)

// 	log.Println("Listening...")
// 	http.ListenAndServe(":3000", mux)
// }

type timeHandler struct {
	format string
}

func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(th.format)
	w.Write([]byte("The time is: " + tm))
}

func main() {
	mux := http.NewServeMux()

	th := &timeHandler{format: time.RFC1123}
	mux.Handle("/time", th)

	th1123 := &timeHandler{format: time.RFC1123}
	mux.Handle("/time/rfc1123", th1123)

	th3339 := &timeHandler{format: time.RFC3339}
	mux.Handle("/time/rfc3339", th3339)

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}
