package api

import (
	"encoding/json"
	"github.com/gojektech/heimdall/httpclient"
	"io/ioutil"
	"time"
)

type Quote struct {
	Content string `json:"content"`
}

func GetQuote() string {
	timeout := 10000 * time.Millisecond
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))
	res, err := client.Get("http://api.quotable.io/random", nil)
	if err != nil{
		panic(err)
	}

	body, _ := ioutil.ReadAll(res.Body)
	var quote Quote
	jsonString := string(body)
	json.Unmarshal([]byte(jsonString), &quote)
	return quote.Content
}

func GetQuoteAsync(result chan string) {
	timeout := 1000 * time.Millisecond
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))
	res, err := client.Get("http://api.quotable.io/random", nil)
	if err != nil{
		panic(err)
	}

	body, _ := ioutil.ReadAll(res.Body)
	var quote Quote
	jsonString := string(body)
	json.Unmarshal([]byte(jsonString), &quote)
	result <- quote.Content
}