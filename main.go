package main

func main() {
	cards := newDeck()

	cards.saveToFile("saved_cards")
}
