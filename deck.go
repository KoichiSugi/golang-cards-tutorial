package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// a slice of strings

type deck []string //custom type

func newDeck() deck {
	cards := deck{} //creates a slice of deck type
	cardSuites := []string{"Spades", "Diamond", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

//here is a custom method of type
// any variables with type deck can aceess to this method
//d is called reciever
func (d deck) print() {
	//iterate deck usign for loop
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//Multiple return values in one function
//(d deck, handSize int) is argument
func deal(d deck, handSize int) (deck, deck) {
	// 0 to handSize-1, handSize to the end
	return d[:handSize], d[handSize:]
}

//converts deck type to string[]
func (d deck) toString() string {
	//use join method in strings package
	return strings.Join([]string(d), ",")

}

//func WriteFile(filename string, data []byte, perm fs.FileMode) error
//0666 default anyone can read + write
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

//func ReadFile(filename string) ([]byte, error)
func newDeckFromFile(filename string) deck {
	//bs for byte slice
	//err will be populated if there is something wrong
	bs, err := ioutil.ReadFile(filename) //search through directory using given filename
	if err != nil {
		//error handling
		fmt.Println("Error: ", err)
		os.Exit(1) //exit the program using Exit method of os package
	}
	//Split slices s into all substrings separeted by sep parameter and returns
	// a slice of the substrings between those separators
	s := strings.Split(string(bs), ",") //casting bs (byte slice )to string
	return deck(s)                      //cast s to deck
}

func (d deck) shuffleCards() {
	//time.Now().UnixNano() to generate new seed value
	source := rand.NewSource(time.Now().UnixNano()) //rand.NewSource() takes int64 as seed values
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1) //get a random value btw length of the slice -1
		//assign to swap
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
