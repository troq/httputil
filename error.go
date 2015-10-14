package httputil

import "net/http"

func Error400(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusBadRequest)
}

func Error400S(w http.ResponseWriter, err string) {
	http.Error(w, err, http.StatusBadRequest)
}

func Error500(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
func Error500S(w http.ResponseWriter, err string) {
	http.Error(w, err, http.StatusInternalServerError)
}
