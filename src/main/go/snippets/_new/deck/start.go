package deckk

import (
	"fmt"

	"github.com/pwera/Playground/src/main/go/snippets/_new/deck/deckk"
)

func ExampleDeck() {
	dm := deckk.DeckManager{Data: deckk.Deck{"RealMessage", "MessageExtractor"}}
	dm2 := append(dm.Data, "RealMessage")

	fmt.Println(len(dm.Data), len(dm2))

	for _, message := range dm2 {
		fmt.Println(message)
	}
	for i, message := range dm2 {
		fmt.Println(message, i)
	}

	dm.Print()

	fmt.Println(dm.Deal(1))

	dm.SaveToFile("x.data")

	deck, err := deckk.NewDeckFromFile("x.data")
	if err != nil {
		fmt.Printf("Error while newDeckFromFile : %v", err)
	}
	fmt.Printf("Deck : %v", deck)
}
