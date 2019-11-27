package webpageanalyser

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

//problemWords := 

type Website struct {
	Url string `json:"url"`
}

type Paragraph struct {
	text string
}

// This will get called for each Paragraph
func processElement(index int, element *goquery.Selection) {
        fmt.Println(element.Text())
}

func grabWebpage(url string) (numberOfParagraphs int){
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

	defer response.Body.Close()

    // Create a goquery document from the HTTP response
    document, err := goquery.NewDocumentFromReader(response.Body)
    if err != nil {
        log.Fatal("Error loading HTTP response body. ", err)
	}
	
	paragraphs := make([]string,0)

    // Find all paragraphs, process with function
    document.Find("p").Each(func(index int, element *goquery.Selection) {
		paragraphs = append(paragraphs, element.Text())
	  })
	
	numberOfParagraphs = len(paragraphs)
	return
}

//This will handle the webpage sent
func AnalyseWebpage(w http.ResponseWriter, r *http.Request) {
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var requestMessage Website
	err = json.Unmarshal(b, &requestMessage)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//Print Locally
	fmt.Printf("Website: %v \n", requestMessage)

	//Send to Webpage Handling Function
	numberOfParagraphs := grabWebpage(requestMessage.Url)
	
	//Create Response Map
	res := make(map[string]interface{})
	res["numberOfParagraphs"] = numberOfParagraphs
	res["url"] = requestMessage.Url
	
	//Package Up Response
	output, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//Send
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
