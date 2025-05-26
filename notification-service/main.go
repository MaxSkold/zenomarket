package main

import (
	"fmt"
	"math/rand"
	"time"
)

var phrasesList = []string{
	"Hello, I'm your friend",
	"How are you?",
	"What do you do?",
	"Bye bye!!!",
}

func randomPhrase() string {
	return phrasesList[rand.Intn(len(phrasesList))]
}

func main() {
	for {
		fmt.Println(randomPhrase())
		time.Sleep(time.Second)
	}
}
