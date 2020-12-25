package main

import (
	"encoding/json"
	"fmt"
	"github.com/ckharide/metainfo"
	"net/http"
	"time"
)

func helloWorldServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Chandra Mouli")
}

func main() {
	// http.HandleFunc("/echoServer", helloWorldServer)
	http.HandleFunc("/echoJSON", echoJSONPayload)
	http.ListenAndServe(":9100", nil)
}

func echoJSONPayload(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Chandra Mouli New")
	data := metainfo.Message{Name: "Chandra", Body: "Living in Hyd", Time: time.Now().UnixNano() / int64(time.Millisecond)}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
