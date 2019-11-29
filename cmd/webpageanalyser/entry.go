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


type Website struct {
	Url string `json:"url"`
}

type Paragraph struct {
	text string
}

type ProblemWords struct {
	Word string `json:"word"`
	Locations []int `json:"locations"`
	Message string `json:"message"`
} 

type Response struct {
	NumberOfParagraphs int `json:"numberOfParagraphs"`
	Url string `json:"url"`
	Words []ProblemWords `json:"problemWords"`
}

// This will get called for each Paragraph
func processElement(index int, element *goquery.Selection) {
        fmt.Println(element.Text())
}

func grabWebpage(url string) (numberOfParagraphs int, problemWords []ProblemWords){
	grabWebpageClient := http.Client{
		Timeout: time.Second * 5,
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

	numberOfParagraphs = 0

	//Dictionary
	dictionary := []struct{
		word string
		message string
	}{
		{"think", "'Think' implies that this is an opinion, can you be sure this is a factual article?"},
		{"communist", "This term is often used in clickbait articles, is this article trying to rile you up or scare you?"},
		{"cancer", "This term is often used in clickbait articles, is this article trying to rile you up or scare you?"},
		{"immigrant", "This term is often used in clickbait articles, is this article trying to rile you up or scare you?"},
		{"sex", "This term is often used in clickbait articles, is this article trying to rile you up or scare you?"},
		{"studies show", "What studies? Does this article give a source?"},
		{"fact", "Facts are sometimes claimed when there are none. Is there a source?"},
		{"addicted", "This is a very charged term. Is the article fair and unbaised?"},
		{"lied", "This is a very charged term. Is the article fair and unbaised?"},
		{"liar", "This is a very charged term. Is the article fair and unbaised?"},
		{"millenial", "This term is often used in clickbait articles, is this article trying to rile you up or scare you?"},
	}

	//Response Prep
	problemWords = make([]ProblemWords,0)
	
	//Loop Through Dictionary

	for _, dictionaryItem := range dictionary {
		indexes := make([]int,0)
		
		// Find all paragraphs, process with function
		document.Find("p").Each(func(index int, element *goquery.Selection) {
			numberOfParagraphs++
			text := element.Text()
			if strings.Contains(text, dictionaryItem.word){
				indexes = append(indexes, index)
			}
		})
		
		if len(indexes) > 0 {
			object := ProblemWords{dictionaryItem.word,indexes,dictionaryItem.message}
		
			problemWords = append(problemWords,object) 
		}

	}


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
	numberOfParagraphs, problemWords := grabWebpage(requestMessage.Url)
	

	fmt.Printf("Words: \n %v \n", problemWords)
	//Create Response Map
	res := Response{
		numberOfParagraphs,
		requestMessage.Url,
		problemWords,
	}

	fmt.Printf("Res: \n %v \n", res)
	// res["numberOfParagraphs"] = 
	// res["url"] = requestMessage.Url

	// res["problemWords"] = problemWords	
	
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
		w.Header().Set("Access-Control-Allow-Origin", "https://us-central1-graphite-bliss-260202.cloudfunctions.net/AnalyseWebpage")
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
