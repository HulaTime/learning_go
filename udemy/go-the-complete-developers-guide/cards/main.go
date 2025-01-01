// The cards package is a basic simulator for dealing with a pack of cards
package main

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")
	newdeck := newDeckFromFile("my_cards")
	newdeck.shuffle().print()
}
