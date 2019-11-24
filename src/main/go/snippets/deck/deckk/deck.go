package deckk

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type Deck []string

type DeckManager struct {
	Data Deck
}

func newDeck() Deck {
	d := Deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearth", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			d = append(d, value+" of "+suit)
		}
	}
	return d
}

func (d *DeckManager) Print() {
	for i, m := range d.Data {
		fmt.Println(i, m)
	}
}

func (d *DeckManager) Deal(size int) (Deck, Deck) {
	return d.Data[:size], d.Data[size:]
}

func (d *DeckManager) toString() string {
	ar := []string(d.Data)
	return strings.Join(ar, ",")
}

func NewDeckFromFile(name string) (DeckManager, error) {
	bs, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Printf("Error while ReadFile : %v", err)
		return DeckManager{}, err
	}
	s := strings.Split(string(bs), ",")
	fmt.Printf("%v", s)
	return DeckManager{Data: s}, nil

}

func (dec *DeckManager) Shuffle() {
	d := dec.Data
	source := rand.NewSource(time.Now().UnixNano())
	for i := range d {
		newPosition := rand.New(source).Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]

	}
}

func (d *DeckManager) SaveToFile(fName string) {
	bytes := []byte(d.toString())
	err := ioutil.WriteFile(fName, bytes, 0666)
	if err != nil {
		fmt.Printf("Error while WriteFile : %v", err)
	}
}
