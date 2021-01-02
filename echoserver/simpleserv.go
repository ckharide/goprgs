package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ckharide/metainfo"
	"io/ioutil"
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
	/*postBody, _ := json.Marshal(map[string]string{
		"name":  "test",
		"email": "Toby@example.com",
	})*/

	data := metainfo.Message{Name: "Chandra", Body: "Living in Hyd", Time: time.Now().UnixNano() / int64(time.Millisecond)}
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error:", err)
	}
	responseBody := bytes.NewBuffer(b)
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post("https://postman-echo.com/post", "application/json", responseBody)
	//Handle Error
	if err != nil {
		fmt.Println("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	fmt.Fprintf(w, "Hello Chandra Mouli New")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	sb := string(body)
	fmt.Fprintf(w, sb)
	//data := metainfo.Message{Name: "Chandra", Body: "Living in Hyd", Time: time.Now().UnixNano() / int64(time.Millisecond)}
	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusCreated)
	//json.NewEncoder(w).Encode(data)
}
