package main

import (
	"sync"

	"github.com/albingeorge/golang-patterns/libs"
)

func ReverseMultipleGoroutines(words []string) []libs.ReverseResult {
	result := []libs.ReverseResult{}

	wg := &sync.WaitGroup{}
	wg.Add(len(words))
	m := &sync.Mutex{}
	for _, word := range words {
		go func(word string) {
			defer wg.Done()
			res := libs.Reverse(word)

			m.Lock()
			result = append(result, res)
			m.Unlock()
		}(word)
	}

	wg.Wait()

	return result
}
