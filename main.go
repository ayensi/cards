package main

func main() {
	card := newCard()
	cards := deck{}
	cards = append(cards, "Two Of Hearts")
	cards = append(cards, card)
	cards.print()
}

func newCard() string {
	return "Five Of Hearts"
}
