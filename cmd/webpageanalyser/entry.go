package webpageanalyser

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"fmt"
)

type Website struct {
	Url string `json:"url"`
}

func grabWebpage(url string){
	grabWebpageClient := http.Client{
		Timeout: time.Second * 3,
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	response, err2 := grabWebpageClient.Do(request)
	if err2 != nil {
		log.Fatal(err2)
		return
	}

	dataInBytes, err3 := ioutil.ReadAll(response.Body)
	if err3 != nil {
		log.Fatal(err3)
	}

	// Get the response body as a string
	pageContent := string(dataInBytes)
	
	//Print Body
	fmt.Println(pageContent)
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

	grabWebpage(msg.Url)
	

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
