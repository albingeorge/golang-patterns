package main

import (
	"fmt"
	"strings"

	"github.com/tjarratt/babble"
)

func main() {
	babbler := babble.NewBabbler()
	babbler.Separator = " "
	babbler.Count = 100

	words := strings.Split(babbler.Babble(), " ")

	fmt.Printf("Input :\n%v\n", words)
	// res := BadReverseMultipleGoroutines(words)
	// res := BadReverseSingleGoroutine(words)
	// res := WorkingReverseSingleGoroutine(words)
	res := ReverseMultipleGoroutines(words)

	fmt.Printf("Results :\n%+v\n", res)
	if len(words) != len(res) {
		fmt.Printf("Missing data in result")
	}
}
