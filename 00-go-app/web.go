package main

import (
    "fmt"
	"log"
	"strings"
	"net/http"
	"os/exec"
	"time"
)

var hostname = getHostName()
var startedAt = time.Now()

func namedHandler(w http.ResponseWriter, r *http.Request) {
	var path = r.URL.Path[1:]
	if strings.Compare("ping", path) == 0 {
		fmt.Fprintf(w, "%s : pong\n", hostname)
	} else if strings.Compare("pong", path) == 0 {
		fmt.Fprintf(w, "%s : ping\n", hostname)
	} else if strings.Compare("health", path) == 0 {
		duration := time.Now().Sub(startedAt)
		if duration.Seconds() > 30 {
			w.WriteHeader(500)
			w.Write([]byte(fmt.Sprintf("Dead for %v seconds...\n", duration.Seconds())))
		} else {
			w.WriteHeader(200)
			w.Write([]byte("OK\n"))
		}
	} else {
		fmt.Fprintf(w, "%s : Welcome to Ping/Pong (v2)...\n", hostname)
	}
}

func getHostName() string {
	out, err := exec.Command("hostname").Output()
    if err != nil {
        log.Fatal(err)
	}
	return strings.TrimSpace(string(out))
}

func main() {
	fmt.Printf("Started Server on %s:8080\n",hostname)
    http.HandleFunc("/", namedHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}