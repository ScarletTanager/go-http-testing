package server

import (
	"net/http"
)

const (
	HEADER_KEY_X_ACCOUNT = "X-Account"
)

func HandleGET(w http.ResponseWriter, r *http.Request) {
	// Check our required header
	account := r.Header.Get(HEADER_KEY_X_ACCOUNT)
	if account == "" {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		// Ordinarily you should not need to do this, assuming you're using a
		// packaged mux/routing layer, but if you're writing your own,
		// it might be a good idea.
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			// Process request
			w.Write([]byte("Content!"))
		}
	}
}

func HandlePOST(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}
