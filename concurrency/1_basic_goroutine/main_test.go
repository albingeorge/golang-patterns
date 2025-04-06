package main

import (
	"strings"
	"testing"

	"github.com/tjarratt/babble"
)

var input = generateWords(1000)

func BenchmarkWorkingReverseSingleGoroutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		b.StartTimer()
		res := WorkingReverseSingleGoroutine(input)
		if len(res) != len(input) {
			b.Errorf("not all words processed")
		}
	}
}

func BenchmarkReverseMultipleGoroutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		b.StartTimer()
		res := ReverseMultipleGoroutines(input)
		if len(res) != len(input) {
			b.Errorf("not all words processed")
		}
	}
}

func generateWords(length int) []string {
	babbler := babble.NewBabbler()
	babbler.Separator = " "
	babbler.Count = length

	return strings.Split(babbler.Babble(), " ")
}
