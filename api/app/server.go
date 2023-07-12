package app

import (
	"encoding/json"
	"net/http"

	"github.com/AleksanderWWW/fin-app/core"
	"github.com/gorilla/mux"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// -------------------------------------------------------------------------
	// Parse request
	provider := mux.Vars(r)["provider"]

	var initArgs interface{}
	err := json.NewDecoder(r.Body).Decode(&initArgs)
	if err != nil {
		http.Error(w, "Failed to parse JSON request body", http.StatusBadRequest)
		return
	}

	// -------------------------------------------------------------------------
	// Get correct data reader

	reader, err := core.GetReaderFromProviderString(provider, initArgs)

	// -------------------------------------------------------------------------
	// Prepare response encoder
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ") // Indentation of 4 spaces

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// -------------------------------------------------------------------------
	// Fetch financial data
	records, err := reader.FetchData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// -------------------------------------------------------------------------
	// Respond with results
	err = encoder.Encode(records)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
