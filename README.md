_This was created during my time as a student at [Code Chrysalis](https://www.codechrysalis.io/)._

_This was the Polyglottal Project, where I had to learn a new language and build a project within 5 days_

_Through this project I learnt Go, Browser Extensions & Google Cloud Functions_

_Therefore, please excuse the messy code, but I hope to continue working on this as I think its a useful product._


![alt text](./misc/logo-200x200.png 'Go Examine Logo')

# go-examine
An API for highlighting problematic phrases in news articles to train users to spot fake news.

Upon receiving a website URL, the API will download the content from that URL.
It will then identify all the paragraphs within the webpage.
It will check the content of each webpage against a list of predefined "problem phrases".

The Browser Extension part of this project can be found at [go_examine_browser](https://github.com/FraserTooth/go_examine_browser).


### Next Steps in Development
The next things I would like to do to develop this browser extension are:

1. Sort out the local testing/running issues.
2. Use MongoDB Atlas as an external datastore to get the dictionary out of the app.
3. Write Tests
4. Use Concurrency to speed up the search process
5. Send back a more specific location of the highlighted words to the front end.
6. Gamify the app.

## Use
This is designed to be deployed on [Google Cloud Functions](https://cloud.google.com/functions/docs/concepts/go-runtime) hence it does not use any external API libraries.

The [goquery](https://github.com/PuerkitoBio/goquery) library is used to parse the webpages a little easier.

The API is publicly available at https://us-central1-graphite-bliss-260202.cloudfunctions.net/AnalyseWebpage

## Example

![alt text](./misc/wikipediaExample.png 'Wikipedia Example of Go Examine')

_Example from the [Fake News](https://en.wikipedia.org/wiki/Fake_news) entry on Wikipedia_

## API Interaction

Send the following in a POST request to the API:
```json
{
     "url": "yourwebsiteurl"
}
```

You will receive a response with the following structure:
```json
{
  "numberOfParagraphs": 2530,
  "url": "yourwebsiteurl",
  "problemWords": [
    {
      "word": " example",
      "locations": [
        16,
        24
      ],
      "message": "Example Message which will be shown in tooltip"
    },
   ...
   ...
   ]
}
```

## Development

Ensure that you have the latest version of Go installed.

You 'can' run this locally by deleting the `go.mod` and `go.sum` files in `/cmd/webpageanalyser` and then running this from the project root:
```bash
go run cmd/server/server.go
```
But this is a bit messy, I'll figure out a way of sorting the `go.mod` issues out someday.
