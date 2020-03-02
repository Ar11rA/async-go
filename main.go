package main

import (
	"async/api"
	"async/mocks"
	"fmt"
	"time"
)

func callSimulatorAsync(done chan string) {
	go mocks.AsyncA(done)
	go mocks.AsyncB(done)
}

func callSimulatorSync() {
	fmt.Println(mocks.SyncA())
	fmt.Println(mocks.SyncB())
}

func callQuoteSync() {
	for i := 0 ;i < 15; i++ {
		quote := api.GetQuote()
		fmt.Println(quote)
	}
}

func callQuoteAsync(done chan string) {
	for i := 0 ;i < 15; i++ {
		go api.GetQuoteAsync(done)
	}
	for i := 0 ;i < 15; i++ {
		fmt.Println(<-done)
	}
}

func getEpisodes() {
	episodes := api.GetEpisodes()
	names := make([]chan string, len(episodes))
	for i := 0; i < len(episodes); i++ {
		names[i] = make(chan string)
		go api.GetEpisodeName(episodes[i], names[i])
	}
	for i := 0; i < len(episodes); i++ {
		fmt.Println(<-names[i])
	}
}

func main() {
	fmt.Println("Getting episodes for Rick&Morty")
	getEpisodes()
	fmt.Println("Start simulator sync", time.Now())
	callSimulatorSync()
	fmt.Println("End simulator sync", time.Now())
	fmt.Println("Start simulator async", time.Now())
	done := make(chan string)
	callSimulatorAsync(done)
	fmt.Println("Start simulator async", time.Now())
	fmt.Println(<-done)
	fmt.Println(<-done)
	fmt.Println("Start quote sync", time.Now())
	callQuoteSync()
	fmt.Println("End quote sync", time.Now())
	fmt.Println("Start quote async", time.Now())
	fmt.Println(time.Now())
	callQuoteAsync(done)
	fmt.Println(time.Now())
	fmt.Println("End quote async", time.Now())
}
