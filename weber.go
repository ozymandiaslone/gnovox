package main

import (
  "io/ioutil"
  "net/http"
  "fmt"
  "net/url"
  "encoding/xml"
)

type Entry struct {
  XMLName xml.Name `xml:"entry"`
  Id      string   `xml:"id"`
  Title   string   `xml:"title"`
}

type Feed struct {
  XMLName xml.Name `xml:"feed"`
  Entry   []Entry  `xml:"entry"`
}

func ScrapeArXiv() {
  searchTerms := ReadMultiLineInput()
  encodedSearchTerms := url.QueryEscape(searchTerms)
  fmt.Println("Encoded Search Terms: ", encodedSearchTerms)
  url := fmt.Sprintf(`http://export.arxiv.org/api/query?search_query=all:%s&start=0&max_results=10`, encodedSearchTerms)
  fmt.Println("URL:", url)
  resp, err := http.Get(url)
  if err != nil {
    fmt.Printf("Error getting response %s", err)
    return
  }
  defer resp.Body.Close()
  bodyBytes, _ := ioutil.ReadAll(resp.Body)
  var feed Feed
  xml.Unmarshal(bodyBytes, &feed)
  for _, entry := range feed.Entry {
    id := entry.Id
    title := entry.Title
    fmt.Printf("Title: %s \n", title)
    fmt.Printf("Id: %s \n\n", id)
  }
}

func ScrapeSciHub(string url) {
  sciUrl := fmt.Sprintf(`https://sci-hub.wf/%s`, url)
}
