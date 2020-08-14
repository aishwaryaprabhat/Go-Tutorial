package main

func main() {

	cards := newDeckFromFile("my_cards")

	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
	cards.shuffle()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"

}
