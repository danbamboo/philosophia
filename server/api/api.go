package api

import(
	"strconv"
	"net/http"
)


// Healthcheck responds with a generic healthcheck
func Healthcheck(w http.ResponseWriter, r *http.Request) {
	sendBytes(w, r, []byte(`{"status":"ok"}`))
}


func sendBytes(w http.ResponseWriter, r *http.Request, b []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(b)))
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}