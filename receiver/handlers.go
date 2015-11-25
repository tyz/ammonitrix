package receiver

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/eBayClassifiedsGroup/ammonitrix/config"
)

func handleRegister(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "registered!\n")
}

func handleDeregister(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "deregistered!\n")
}

func handleData(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Unsupported method", 405)
		return
	}
	var valid bool
	var datagram config.Datagram
	if valid, datagram = validateDataRequest(r.Body); valid != true {
		http.Error(w, "Invalid request received", 500)
		return
	}
	be_elastic.StoreDatagram(datagram)
	err := json.NewEncoder(w).Encode(datagram)
	if err != nil {
		http.Error(w, "Failed to encode datagram", 500)
	}
}
