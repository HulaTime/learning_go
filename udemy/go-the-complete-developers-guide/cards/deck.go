package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type deck []string

// Create a new standard deck of playing cards
func newDeck() deck {
	suits := [4]string{"C", "D", "H", "S"}
	newdeck := make(deck, 52)

	for di := 0; di < len(newdeck)/4; di++ {
		for si, suit := range suits {
			position := di*4 + si
			if di == 0 {
				newdeck[position] = "A" + suit
			} else if di == 10 {
				newdeck[position] = "J" + suit
			} else if di == 11 {
				newdeck[position] = "Q" + suit
			} else if di == 12 {
				newdeck[position] = "K" + suit
			} else {
				newdeck[position] = strconv.Itoa(di+1) + suit
			}
		}
	}

	return newdeck
}

func newDeckFromFile(filename string) deck {
	filebytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to read deck from file", err)
		os.Exit(1)
	}
	return strings.Split(string(filebytes), ",")
}

func (deck deck) print() {
	for i, card := range deck {
		fmt.Println(i+1, card)
	}
}

func (deck deck) deal(handsize int8) (deck, deck) {
	return deck[:handsize], deck[handsize:]
}

func (d deck) toString() string {
	s := ""
	// could just be done with strings.Join(deck, ",")
	for i, card := range d {
		if i == 0 {
			s += card
		} else {
			s += "," + card
		}
	}
	return s
}

func (d deck) saveToFile(filename string) {
	os.WriteFile(filename, []byte(d.toString()), 0666)
}

func (d deck) shuffle() deck {
	for i := range d {
		new_position := rand.Intn(len(d) - 1)
		d[i], d[new_position] = d[new_position], d[i]
	}
	return d
}
