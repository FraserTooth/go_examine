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
	//Remember to add a space to the start of each string
	dictionary := []struct{
		word string
		message string
	}{
		{" think", "'Think' implies that this is an opinion, can you be sure this is a factual article?"},
		{" belive", "'Belive' implies that this is an opinion, can you be sure this is a factual article?"},
		
		{" communist", "This term is often used in clickbait articles, is this article trying to rile you up or scare you?"},
		{" immigrant", "This term is often used in clickbait articles, is this article trying to rile you up or scare you?"},
		{" sex", "This term is often used in clickbait articles, is this article trying to rile you up or scare you?"},
		{" millenial", "This term is often used in clickbait articles, is this article trying to rile you up or scare you?"},
		{" cancer", "This term is often used in clickbait articles, is this article trying to rile you up or scare you?"},
		{" vaccination", "This term is often used in clickbait articles, is this article trying to rile you up or scare you?"},
		
		{" studies", "What studies? Does this article give a source?"},
		{" study", "What study? Does this article give a source?"},
		{" scientist", "Which scientists? Does this article give a source?"},
		
		{" fact", "False articles often make appeals to 'facts', even when they are untrue. Is there a source?"},
		{" truth", "False articles often make appeals to 'truth', even when they are untrue. Is there a source?"},
		{" known", "False articles often make appeals to 'knowledge', even when they are untrue. Is there a source?"},
		
		{" addicted", "This is a very charged term. Is the article fair and unbaised?"},
		{" lied", "This is a very charged term. Is the article fair and unbaised?"},
		{" liar", "This is a very charged term. Is the article fair and unbaised?"},
		
		
		{" ridiculed", "This is an objective term. Is this article trying to tell you how to feel?"},
		{" destroyed", "This is an objective term. Is this article trying to tell you how to feel?"},
		
		{" psychotic", "This is a very negative term. Is the article fair and unbaised?"},
		{" coward", "This is a very negative arged term. Is the article fair and unbaised?"},
		
		{" was cancelled", "Does what Twitter thinks always matter?"},
		{" tweeted", "Twitter is a record of what people have said on twitter, not what is objectively true."},
		{" twitter", "Twitter is a record of what people have said on twitter, not what is objectively true."},
		
		{" Corbyn", "This person is often featured in untruthful articles. Is the article representing them fairly?"},
		{" Boris Johnson", "This person is often featured in untruthful articles. Is the article representing them fairly?"},
		{" Trump", "This person is often featured in untruthful articles. Is the article representing them fairly?"},
		{" Clinton", "This person is often featured in untruthful articles. Is the article representing them fairly?"},
		{" Biden", "This person is often featured in untruthful articles. Is the article representing them fairly?"},
		{" Lindsay Lohan", "This person is often featured in untruthful articles. Is the article representing them fairly?"},
		{" Einstein", "This person is often featured in untruthful articles. Is the article representing them fairly?"},
		{" Obama", "This person is often featured in untruthful articles. Is the article representing them fairly?"},
		
		{" stormed", "This is very dramatic language. Is this article trying to tell you how to feel?"},
		{" invaded", "This is very dramatic language. Is this article trying to tell you how to feel?"},
		{" invasion", "This is very dramatic language. Is this article trying to tell you how to feel?"},
		{" stripped", "This is very dramatic language. Is this article trying to tell you how to feel?"},

		{" liberal", "Grouping a diverse group of people under one heading is oversimplification. Is this article trying to make you feel a certain way about this group?"},
		{" the left", "Grouping a diverse group of people under one heading is oversimplification. Is this article trying to make you feel a certain way about this group?"},
		{" the right", "Grouping a diverse group of people under one heading is oversimplification. Is this article trying to make you feel a certain way about this group?"},
		{" rightwing", "Grouping a diverse group of people under one heading is oversimplification. Is this article trying to make you feel a certain way about this group?"},
		{" under-25", "Grouping a diverse group of people under one heading is oversimplification. Is this article trying to make you feel a certain way about this group?"},
		{" working-class", "Grouping a diverse group of people under one heading is oversimplification. Is this article trying to make you feel a certain way about this group?"},
		{" middle-class", "Grouping a diverse group of people under one heading is oversimplification. Is this article trying to make you feel a certain way about this group?"},
		{" upper-class", "Grouping a diverse group of people under one heading is oversimplification. Is this article trying to make you feel a certain way about this group?"},
		{" leftwing", "Grouping a diverse group of people under one heading is oversimplification. Is this article trying to make you feel a certain way about this group?"},
		{" mob", "Grouping a diverse group of people under one heading is oversimplification. Is this article trying to make you feel a certain way about this group?"},

		{" total victory", "This is an objective term. Is this article trying to tell you how to feel?"},
		{" extremely", "This is an objective term. Is this article trying to tell you how to feel?"},
		{" very", "This is an objective term. Is this article trying to tell you how to feel?"},
		
		{" support the troops", "This statement is emotionally charged. Is this article emotional or factual?"},
		{" take your guns", "This statement is emotionally charged. Is this article emotional or factual?"},
		{" are under attack", "This statement is emotionally charged. Is this article emotional or factual?"},
		
		{" accused", "Accusations are not always evidence-based. Has there been an offical investigation or legal case?"},
		
		{" offended", "Who was offended? Articles sometimes claim people are more offended than they really are."},
		{" triggered", "Who was offended? Articles sometimes claim people are more offended than they really are."},
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
			textUpper, wordUpper := strings.ToUpper(text), strings.ToUpper(dictionaryItem.word)
			if strings.Contains(textUpper, wordUpper){
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

	//Send to Webpage Handling Function
	numberOfParagraphs, problemWords := grabWebpage(requestMessage.Url)
	
	//Create Response Map
	res := Response{
		numberOfParagraphs,
		requestMessage.Url,
		problemWords,
	}
	
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
