package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
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
			_, _ = w.Write([]byte(fmt.Sprintf("Path: %s\n\nEnvironment:\n", r.URL.String())))
			_, _ = w.Write([]byte(strings.Join(os.Environ(), "\n")))
		}),
	}

	slowStart := os.Getenv("SLOW_START")
	if slowStart != "" {
		log.Printf("slow start configured: %s", slowStart)

		duration, err := time.ParseDuration(slowStart)
		if err != nil {
			log.Fatalf("failed to parse SLOW_START: %s", err)
		}

		time.Sleep(duration)
	}

	log.Printf("listening on %s", lisAddr)
	err := httpServ.ListenAndServe()
	if err != nil {
		log.Fatalf("listening failed: %s", err)
	}
}
