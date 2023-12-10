package server

import (
	"encoding/json"
	"net/http"
)

func ProcessSingleHandler(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload
	if err := json.NewDecoder(r.Body).Decode(&requestPayload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	sortedArrays, timeTaken := sortArraysSequential(requestPayload.ToSort)

	responsePayload := ResponsePayload{
		SortedArrays: sortedArrays,
		TimeNs:       timeTaken,
	}

	writeJSONResponse(w, responsePayload)
}

func ProcessConcurrentHandler(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload
	if err := json.NewDecoder(r.Body).Decode(&requestPayload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	sortedArrays, timeTaken := sortArraysConcurrent(requestPayload.ToSort)

	responsePayload := ResponsePayload{
		SortedArrays: sortedArrays,
		TimeNs:       timeTaken,
	}

	writeJSONResponse(w, responsePayload)
}
