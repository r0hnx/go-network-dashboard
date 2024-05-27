package handler

import (
	"encoding/json"
	"net/http"
	"nmap-scanner/utils"
)

func HandleNmapScan(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("target")
	if target == "" {
		http.Error(w, "Missing target parameter", http.StatusBadRequest)
		return
	}

	results, err := utils.RunNmap(target)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
