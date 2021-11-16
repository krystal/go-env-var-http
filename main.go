package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	msg := os.Getenv("MESSAGE")
	httpServ := http.Server{
		Addr: ":9000",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.WriteHeader(200)
			_, _ = w.Write([]byte(msg))
		}),
	}

	err := httpServ.ListenAndServe()
	if err != nil {
		log.Fatalf("listening failed: %s", err)
	}
}
