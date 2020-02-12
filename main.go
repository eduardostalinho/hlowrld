package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/endocode/hlowrld/logging"
	"github.com/endocode/hlowrld/server"

	"github.com/gorilla/mux"
)

func run(args []string) int {

	bindAddress := flag.String("ip", "0.0.0.0", "IP address to bind")
	listenPort := flag.Int("port", 25478, "port number to listen on")
	debugFlag := flag.Bool("debug", false, "Turn debug on/off")
	flag.Parse()

	logger := logging.Start(*debugFlag)
	logger.Printf("[DEBUG] Debug mode is on")

	server := server.New(logger)

	r := mux.NewRouter()

	r.HandleFunc("/healthz", server.HealthCheckHandler)
	r.HandleFunc("/", server.GetHandler).Methods("Get")

	r.Use(server.LoggingMiddleware)

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *bindAddress, *listenPort))
	if err != nil {
		logger.Printf("[FATAL] Failed to initialize listener: %v", err)
	}

	logger.Printf("[INFO] Listening on %v", fmt.Sprintf("%s:%d", *bindAddress, *listenPort))
	log.Fatal(http.Serve(listener, r))
	return 0
}

func main() {
	result := run(os.Args)
	os.Exit(result)
}
