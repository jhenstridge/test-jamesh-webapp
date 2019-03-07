package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/coreos/go-systemd/activation"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain;encoding=UTF-8")
	w.WriteHeader(200)
	fmt.Fprintf(w, "Hello, %q\n", r.URL.Path)
}

func main() {
	// Set up routes
	http.HandleFunc("/", rootHandler)

	// Receive listening socket from systemd
	var listener net.Listener
	if listeners, err := activation.Listeners(); err != nil {
		log.Fatal(err)
	} else if len(listeners) != 1 {
		log.Fatal("Expected one socket activated file descriptor")
	} else {
		listener = listeners[0]
	}

	// Serve requests forever
	log.Fatal(http.Serve(listener, nil))
}
