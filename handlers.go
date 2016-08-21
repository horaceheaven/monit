package main

import (
	"encoding/json"
	"net/http"
)

func MonitorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.WriteHeader(http.StatusOK)

	m := Monitor{
		UpTime: UpTime(),
	}

	if err := json.NewEncoder(w).Encode(m); err != nil {
		panic(err)
	}
}
