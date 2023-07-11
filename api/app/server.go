package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AleksanderWWW/fin-app/core"
	"github.com/gorilla/mux"
)


func HandleRequest(w http.ResponseWriter, r *http.Request) {
	provider := mux.Vars(r)["provider"]

	reader := core.GetReaderFromProviderString(provider)
	

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ") // Indentation of 4 spaces
	
	if reader == nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := fmt.Sprintf("Invalid data provider - '%s'", provider)
		err := encoder.Encode(msg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	records := reader.FetchData()
	err := encoder.Encode(records)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
