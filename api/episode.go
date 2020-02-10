package api

import (
	"encoding/json"
	"github.com/gojektech/heimdall/httpclient"
	"io/ioutil"
	"time"
)

type Character struct{
	Episodes []string `json:"episode"`
}

type Episode struct{
	Name string `json:"name"`
}

func GetEpisodes() []string {
	timeout := 1000 * time.Millisecond
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))
	res, err := client.Get("https://rickandmortyapi.com/api/character/1", nil)
	if err != nil{
		panic(err)
	}

	body, _ := ioutil.ReadAll(res.Body)
	var character Character
	jsonString := string(body)
	json.Unmarshal([]byte(jsonString), &character)
	return character.Episodes
}

func GetEpisodeName(episodeUrl string, name chan string) {
	timeout := 1000 * time.Millisecond
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))
	res, err := client.Get(episodeUrl, nil)
	if err != nil{
		panic(err)
	}

	body, _ := ioutil.ReadAll(res.Body)
	var episode Episode
	jsonString := string(body)
	json.Unmarshal([]byte(jsonString), &episode)
	name <- episode.Name
}