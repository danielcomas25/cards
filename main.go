package main

func main() {

	// cards := newDeck() // Initialized and assigned

	// hand, remainingDeck := deal(cards, 5) // Two returning values
	// cards.saveToFile("my_cards")
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}
