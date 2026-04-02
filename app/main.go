package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	version := os.Getenv("APP_VERSION")
	if version == "" {
		version = "v1.0.0"
	}
	color := os.Getenv("APP_COLOR")
	if color == "" {
		color = "#16a34a"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head><title>Latitude Demo App</title></head>
<body style="background:%s;color:white;font-family:sans-serif;display:flex;align-items:center;justify-content:center;height:100vh;margin:0">
<div style="text-align:center">
<h1>🚀 Latitude K8s Demo</h1>
<h2>Version: %s</h2>
<p>Hostname: %s</p>
<p>Time: %s</p>
<p>Deployed via Flux CD</p>
</div>
</body>
</html>`, color, version, os.Getenv("HOSTNAME"), time.Now().Format("2006-01-02 15:04:05"))
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprintf(w, `{"status":"ok","version":"%s"}`, version)
	})

	log.Printf("Starting server version %s on :8080", version)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
