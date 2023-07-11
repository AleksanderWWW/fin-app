package app

import (
	"fmt"
	"net/http"

	"github.com/AleksanderWWW/fin-app/core"
	"github.com/gorilla/mux"
)


func HandleRequest(w http.ResponseWriter, r *http.Request) {
	provider := mux.Vars(r)["provider"]

	reader, ok := core.ReaderMap[provider]
	
	if !ok {
		fmt.Println("invalid provider")
	}
	fmt.Println(reader)
}
