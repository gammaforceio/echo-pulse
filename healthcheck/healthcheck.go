// Package healthcheck provides HTTP health check functionality
package healthcheck

import (
	"net/http"
)

// Start initializes the health check server on the given port
func Start(port string) {
	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "OK"}`))
	})

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
