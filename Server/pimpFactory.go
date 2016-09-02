package main

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/wallnutkraken/Auction/Server/Auction"
)

var words chan string = make(chan string, 16)

func GeneratePimp() Auction.Pimp {
	return Auction.NewPimp(rand.Int(), Auction.NewItem(getWord()))
}

func getWord() string {
	return <-words
}

func fillWords() {
	for {
		resp, err := http.Get("http://www.setgetgo.com/randomword/get.php")
		if err != nil {
			logger.Println("Error getting word:", err.Error(), ". Retry after 1 second.")
			time.Sleep(time.Second)
			continue
		}
		text, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.Println("Error reading result from word get:", err.Error(), ". Retry after 1 second.")
			time.Sleep(time.Second)
			continue
		}
		words <- string(text)
	}
}

func init() {
	go fillWords()
}
