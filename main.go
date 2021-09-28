package main

func main() {
	//var card string = "Ace" full
	// card := "Ace of Spades"   //shorthanded := is only for declaring new variables
	// card = "Five of Diamonds" //reassign
	//cards = append(cards, "Six of Spades") //append a create a new Slice and assign to cards
	// cards := newDeck()
	// fmt.Println(cards.toString()) //get comma separated string
	// cards.saveToFile("myCards")

	//cards.print() //iterate a Slice
	// hand.print()
	// remaingDeck.print()
	//cards := newDeckFromFile("mycards")
	cards := newDeck()
	//hand, remaingDeck := deal(cards, 5)
	// hand.print()
	// remaingDeck.print()
	cards.shuffleCards()
	cards.print()
}
