package main

import (
	"fmt"

	"github.com/Pwera/Playground/src/main/go/snippets/deck/deckk"
)

func main() {
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
