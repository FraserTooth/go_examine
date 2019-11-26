package webpageanalyser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Website struct {
	Url string `json:"url"`
}

func AnalyseWebpage(w http.ResponseWriter, r *http.Request) {
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var msg Website
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//Print Locally
	fmt.Printf("Website: %v \n", msg)

	//Package Up Response
	res := msg
	output, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//Send
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
