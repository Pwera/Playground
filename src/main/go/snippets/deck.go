package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	d := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearth", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			d = append(d, value+" of "+suit)
		}
	}
	return d
}

func (d deck) print() {
	for i, m := range d {
		fmt.Println(i, m)
	}
}

func (d deck) deal(size int) (deck, deck) {
	return d[:size], d[size:]
}

func (d deck) toString() string {
	ar := []string(d)
	return strings.Join(ar, ",")
}

func newDeckFromFile(name string) (deck, error) {
	bs, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Printf("Error while ReadFile : %v", err)
		return nil, err
	}
	s := strings.Split(string(bs), ",")
	fmt.Printf("%v", s)
	return deck(s), nil

}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	for i := range d {
		newPosition := rand.New(source).Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]

	}
}

func (d deck) saveToFile(fName string) {
	bytes := []byte(d.toString())
	err := ioutil.WriteFile(fName, bytes, 0666)
	if err != nil {
		fmt.Printf("Error while WriteFile : %v", err)
	}
}
