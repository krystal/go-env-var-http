// Package main provides the http handlers for testing out network connectivity
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	envPort           = "PORT"
	envSlowStart      = "SLOW_START"
	readHeaderTimeout = 20 * time.Second
	readTimeout       = 1 * time.Minute
	writeTimeout      = 2 * time.Minute
)

func httpServer(listenAddr string) *http.Server {
	return &http.Server{
		ReadHeaderTimeout: readHeaderTimeout,
		ReadTimeout:       readTimeout,

		WriteTimeout: writeTimeout,
		Addr:         listenAddr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			if _, err := w.Write(
				[]byte(
					fmt.Sprintf(
						"Path: %s\n\nEnvironment:\n%s",
						r.URL.String(),
						strings.Join(os.Environ(), "\n"),
					),
				),
			); err != nil {
				log.Println("failed to write response")
			}
		}),
	}
}

func slowStart(startTime string) {
	if startTime != "" {
		log.Printf("slow start configured: %s", startTime)

		duration, err := time.ParseDuration(startTime)
		if err != nil {
			log.Printf("failed to parse SLOW_START: %s", err)

			return
		}

		time.Sleep(duration)
	}
}

func main() {
	listenAddr := ":9000"
	port := os.Getenv(envPort)

	if port != "" {
		listenAddr = ":" + port
	}

	slowStart(os.Getenv(envSlowStart))

	log.Printf("listening on %s", listenAddr)

	err := httpServer(listenAddr).ListenAndServe()
	if err != nil {
		log.Fatalf("listening failed: %s", err)
	}
}
