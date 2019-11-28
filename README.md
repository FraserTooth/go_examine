# go-examine
An API for highlighting problematic phrases in news articles to train users to spot fake news.

Upon receiving a website URL, the API will download the content from that URL.
It will then identify all the paragraphs within the webpage.
It will check the content of each webpage against a list of predefined "problem phrases".

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
  "indexesOfProblemParagraphs": [
    16,
    24,
    43
  ],
  "numberOfParagraphs": 57,
  "url": "yourwebsiteurl"
}
```

## Development

Ensure that you have the latest version of Go installed.