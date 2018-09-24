package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)


func main() {
	var httpPort uint64
	flag.Uint64Var(&httpPort, "port", 8080, "HTTP server port")
	flag.Uint64Var(&httpPort, "p", 8080, "HTTP server port (shorthand)")
	flag.Parse()

	var addr string = ":" + strconv.FormatUint(httpPort, 10)

	mux := http.NewServeMux()

	staticFilesHandler := http.FileServer(http.Dir("./"))

	// serve static files from ./
	mux.Handle("/", http.StripPrefix("/", staticFilesHandler))

	log.Println("Listening at: http://" + addr + "/")
	http.ListenAndServe(addr, mux)

	// should not reach here
	log.Println("Error: Failed to start server.")
}

