package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type (
	response struct {
		responseData `json:"responseData"`
	}

	responseData struct {
		Results []responseItem `json:"results"`
	}

	responseItem struct {
		GsearchResultClass string `json:"GsearchResultClass"`
		UnescapedURL       string `json:"unescapedUrl"`
		URL                string `json:"url"`
		VisibleURL         string `json:"visibleUrl"`
		CacheURL           string `json:"cacheUrl"`
		Title              string `json:"title"`
		TitleNoFormatting  string `json:"titleNoFormatting"`
		Content            string `json:"content"`
	}
)

func TestToObject(t *testing.T) {
	data := "{ \"responseData\": { \"results\": [ { \"GsearchResultClass\": \"GwebSearch\", \"unescapedUrl\": \"https://www.reddit.com/r/golang\", \"url\": \"https://www.reddit.com/r/golang\", \"visibleUrl\": \"www.reddit.com\", \"cacheUrl\": \"http://www.google.com/search?q=cache:W...\", \"title\": \"r/\u003cb\u003eGolang\u003c/b\u003e - Reddit\", \"titleNoFormatting\": \"r/Golang - Reddit\", \"content\": \"First Open Source\" }, { \"GsearchResultClass\": \"GwebSearch\", \"unescapedUrl\": \"http://tour.golang.org/\", \"url\": \"http://tour.golang.org/\", \"visibleUrl\": \"tour.golang.org\", \"cacheUrl\": \"http://www.google.com/search?q=cache:O...\", \"title\": \"A Tour of Go\", \"titleNoFormatting\": \"A Tour of Go\", \"content\": \"Welcome to a tour of the Go programming ...\" } ] }}"

	var res response

	json.Unmarshal([]byte(data), &res)

	fmt.Println(res.responseData.Results[0].Title)
	fmt.Println(res.responseData.Results[0].Content)
}

func TestJsonMap(t *testing.T) {
	data := "{ \"responseData\": { \"results\": [ { \"GsearchResultClass\": \"GwebSearch\", \"unescapedUrl\": \"https://www.reddit.com/r/golang\", \"url\": \"https://www.reddit.com/r/golang\", \"visibleUrl\": \"www.reddit.com\", \"cacheUrl\": \"http://www.google.com/search?q=cache:W...\", \"title\": \"r/\u003cb\u003eGolang\u003c/b\u003e - Reddit\", \"titleNoFormatting\": \"r/Golang - Reddit\", \"content\": \"First Open Source\" }, { \"GsearchResultClass\": \"GwebSearch\", \"unescapedUrl\": \"http://tour.golang.org/\", \"url\": \"http://tour.golang.org/\", \"visibleUrl\": \"tour.golang.org\", \"cacheUrl\": \"http://www.google.com/search?q=cache:O...\", \"title\": \"A Tour of Go\", \"titleNoFormatting\": \"A Tour of Go\", \"content\": \"Welcome to a tour of the Go programming ...\" } ] }}"

	var res map[string]interface{}

	json.Unmarshal([]byte(data), &res)

	responseData, _ := (res["responseData"]).(map[string]interface{})
	results, _ := (responseData["results"]).([]interface{})
	for _, value := range results {
		item := value.(map[string]interface{})
		for k, v := range item {
			fmt.Printf("%v = %v ; ", k, v)
		}
		fmt.Print("\n\n")
	}
}

func Test2Json(t *testing.T) {
	data := "{ \"responseData\": { \"results\": [ { \"GsearchResultClass\": \"GwebSearch\", \"unescapedUrl\": \"https://www.reddit.com/r/golang\", \"url\": \"https://www.reddit.com/r/golang\", \"visibleUrl\": \"www.reddit.com\", \"cacheUrl\": \"http://www.google.com/search?q=cache:W...\", \"title\": \"r/\u003cb\u003eGolang\u003c/b\u003e - Reddit\", \"titleNoFormatting\": \"r/Golang - Reddit\", \"content\": \"First Open Source\" }, { \"GsearchResultClass\": \"GwebSearch\", \"unescapedUrl\": \"http://tour.golang.org/\", \"url\": \"http://tour.golang.org/\", \"visibleUrl\": \"tour.golang.org\", \"cacheUrl\": \"http://www.google.com/search?q=cache:O...\", \"title\": \"A Tour of Go\", \"titleNoFormatting\": \"A Tour of Go\", \"content\": \"Welcome to a tour of the Go programming ...\" } ] }}"

	var res response

	json.Unmarshal([]byte(data), &res)

	noIndent, _ := json.Marshal(res)
	fmt.Println(string(noIndent))

	indent, _ := json.MarshalIndent(res, "", " ")
	fmt.Println(string(indent))
}
