package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"time"
)

const (
	Port       = "8795"
	TimeFormat = time.RFC3339
)

func main() {
	r := http.NewServeMux()

	r.HandleFunc("GET /time", getTimeHandler)

	s := &http.Server{
		Addr:    net.JoinHostPort("", Port),
		Handler: r,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

type timeResponse struct {
	Time string `json:"time"`
}

func getTimeHandler(w http.ResponseWriter, r *http.Request) {
	resp := timeResponse{
		Time: time.Now().Format(TimeFormat),
	}

	data, err := json.Marshal(&resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
