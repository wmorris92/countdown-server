package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/wmorris92/countdown-server/solver"
)

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func solveCountdown(w http.ResponseWriter, r *http.Request) {
	type countdownResult struct {
		Words []string `json:"words"`
	}
	type errorResult struct {
		Error string `json:"error"`
	}
	vars := mux.Vars(r)
	solutions, err := solver.FindWordsForLetters(vars["letters"])
	var response []byte

	if err != nil {
		response, _ = json.Marshal(errorResult{err.Error()})
	} else {
		response, _ = json.Marshal(countdownResult{solutions})
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(string(response)))
}

func main() {
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/countdown/{letters}", solveCountdown)

	log.Printf("Listening on %s...\n", addr)
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(addr, r))
}
