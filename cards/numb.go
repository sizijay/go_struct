package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {
	for i, x := range d {
		fmt.Println(i, x)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Hearts", "Clubs", "Diamonds"}

	cardValues := []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

	for _, value := range cardValues {
		for _, suit := range cardSuites {
			//fmt.Println(cards, value+" of "+suit)
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, amount int) (deck, deck) {
	return d[:amount], d[amount:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0644)
}

func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(byteSlice), ","))

}

func (d deck) shuffle() deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	fmt.Println(r)
	for i := range d {
		newIdx := r.Intn(len(d))
		d[i], d[newIdx] = d[newIdx], d[i]
	}
	return d
}
