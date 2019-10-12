package main

import (
	"fmt"
)

type Message interface {
	fun(s string) string
}

type Package struct {
	m1 Message
	m2 Message
}
type RealMessage struct {
}

func (rm RealMessage) fun(s string) string {
	return s + " ..."
}

type MessageExtractor struct {
	RealMessage
}

func main() {
	fmt.Println(".......\n......")
	var card string = "?"
	card = "asd"
	c := "asd"
	fmt.Println(card)
	fmt.Println(c)
	myPackage := Package{
		m1: RealMessage{},
		m2: RealMessage{}}
	fmt.Println(myPackage)
	fmt.Println(myPackage.m1.fun("23"))
	me := MessageExtractor{}
	fmt.Println(me)
	fmt.Println(me.fun("??"))

	ar := deck{"RealMessage", "MessageExtractor"}
	ar2 := append(ar, "RealMessage")

	fmt.Println(len(ar), len(ar2))

	for _, message := range ar2 {
		fmt.Println(message)
	}
	for i, message := range ar2 {
		fmt.Println(message, i)
	}

	ar.print()
	ar2.print()

	fmt.Println(ar.deal(1))

	ar.saveToFile("x.data")

	deck, err := newDeckFromFile("x.data")
	if err != nil {
		fmt.Printf("Error while newDeckFromFile : %v", err)
	}
	fmt.Printf("Deck : %v", deck)
}
