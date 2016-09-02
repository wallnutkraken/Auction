package main

import (
	"math/rand"
	"time"
)

func AddRandomPimps() {
	for {
		time.Sleep(time.Second * time.Duration(rand.Intn(30)))
		auction.NewPimp(GeneratePimp())
	}
}
