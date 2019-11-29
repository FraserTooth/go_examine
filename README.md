_This was created during my time as a student at [Code Chrysalis](https://www.codechrysalis.io/)._

_This was the Polyglottal Project, where I had to learn a new language and build a project within 5 days_

_Through this project I learnt Go, Browser Extensions & Google Cloud Functions_

_Therefore, please excuse the messy code, but I hope to continue working on this as I think its a useful product._

# go-examine
An API for highlighting problematic phrases in news articles to train users to spot fake news.

Upon receiving a website URL, the API will download the content from that URL.
It will then identify all the paragraphs within the webpage.
It will check the content of each webpage against a list of predefined "problem phrases".

The Browser Extension part of this project can be found at [go_examine_browser](https://github.com/FraserTooth/go_examine_browser).

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
