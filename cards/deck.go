package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of deck
// slice of strings
type deck []string

func newDeck() deck {
	d := deck{} // empty deck

	suits := []string{"Spades", "Diamond", "Hearts", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range suits {
		for _, value := range values {
			card := value + " of " + suit
			d = append(d, card)
		}
	}

	return d
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, hs int) (deck, deck) {
	return d[:hs], d[hs:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		// handle err
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	content := strings.Split(string(bs), ",")

	return deck(content)
}

func (d deck) shuffle() {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	for i := range d {
		np := r.Intn(len(d) - 1)

		d[i], d[np] = d[np], d[i]
	}
}
