package webpageanalyser

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
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

func grabWebpage(url string) (numberOfParagraphs int, indexesOfProblemParagraphs []int){
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
	
	indexesOfProblemParagraphs = make([]int,0)
	numberOfParagraphs = 0

    // Find all paragraphs, process with function
    document.Find("p").Each(func(index int, element *goquery.Selection) {
		numberOfParagraphs++
		text := element.Text()
		if strings.Contains(text, "think"){
			indexesOfProblemParagraphs = append(indexesOfProblemParagraphs, index)
		}
	  })

	fmt.Println(indexesOfProblemParagraphs)

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
	numberOfParagraphs, indexesOfProblemParagraphs := grabWebpage(requestMessage.Url)
	
	//Create Response Map
	res := make(map[string]interface{})
	res["numberOfParagraphs"] = numberOfParagraphs
	res["url"] = requestMessage.Url
	res["indexesOfProblemParagraphs"] = indexesOfProblemParagraphs	
	
	//Package Up Response
	output, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")


	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
