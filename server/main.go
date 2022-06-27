package main

import (
	"euchre/deck"
	"log"
)

func main() {
	d := deck.NewDeck()
	log.Print(d.Shuffle())

	for i, v := range d.Shuffle() {
		log.Print(i)
		log.Print(v)
	}
}
