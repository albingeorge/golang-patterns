package main

import (
	"log"
	"strings"

	"github.com/tjarratt/babble"
)

func main() {
	babbler := babble.NewBabbler()
	babbler.Separator = " "
	babbler.Count = 100

	words := strings.Split(babbler.Babble(), " ")
	res := reverse(words)
	log.Printf("Results:\n%+v\n", res)
}
