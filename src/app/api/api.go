package api

import (
	"app/models"
	"encoding/json"
	"log"
	"net/http"
)

var (
	counter int
)

func Handlers() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/favicon.ico", nil)
	mux.HandleFunc("/", count)
	mux.HandleFunc("/stats", stats)

	return mux
}

func count(w http.ResponseWriter, r *http.Request) {
	log.Println("Calling count handler")
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Invalid HTTP Method")
		return
	}
	counter++

	myMessage := models.MessageAPI{
		Message: "Hello T3chFest 2018!!",
	}
	respondWithJSON(w, http.StatusOK, myMessage)
}

func stats(w http.ResponseWriter, r *http.Request) {
	log.Println("Calling stats handler")
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Invalid HTTP Method")
		return
	}
	myVisits := models.VisitAPI{
		Visits: counter,
	}

	respondWithJSON(w, http.StatusOK, myVisits)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
