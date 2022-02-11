package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	lisAddr := ":9000"
	portEnv := os.Getenv("PORT")
	if portEnv != "" {
		lisAddr = ":" + portEnv
	}

	httpServ := http.Server{
		Addr: lisAddr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			_, _ = w.Write([]byte(strings.Join(os.Environ(), "\n")))
		}),
	}

	log.Printf("Listening on %s", lisAddr)
	err := httpServ.ListenAndServe()
	if err != nil {
		log.Fatalf("listening failed: %s", err)
	}
}
