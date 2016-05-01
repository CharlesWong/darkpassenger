package pac

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/CharlesWong/darkpassenger/config"
	"log"
	"net/http"
)

var pacHandler = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(bytes)))
	if _, err := w.Write(bytes); err != nil {
		log.Println("unable to write image.")
	}
}

func StartRemotePacService() {
	mux := http.NewServeMux()
	mux.HandleFunc("/dp.pac", pacHandler)
	http.ListenAndServe(config.PacServiceAddr, mux)
}

func StartLocalPacService() {
	mux := http.NewServeMux()
	mux.HandleFunc("/dp.pac", pacHandler)
	http.ListenAndServe(config.PacServiceAddr, mux)
}
