package main

import (
	"sync"
	"time"

	"github.com/albingeorge/golang-patterns/libs"
)

func BadReverseMultipleGoroutines(words []string) []libs.ReverseResult {
	result := []libs.ReverseResult{}
	for _, word := range words {
		// Create separate goroutines for each word to be processed
		// Pass each word as an input to the lambda function, so that
		// when the loop continues, word variable is not replaced by the next
		// loop's instance
		go func(word string) {
			// Multiple goroutines appending to the same result variable
			// This is a potential race condition
			result = append(result, libs.Reverse(word))
		}(word)
	}

	// Wait for a roughly estimated amount of time
	// Else, the goroutines won't be finished on time and the result would be empty or partial
	time.Sleep(50 * time.Millisecond)

	return result
}

func BadReverseSingleGoroutine(words []string) []libs.ReverseResult {
	result := []libs.ReverseResult{}

	// Single goroutine handles all words
	go func() {
		for _, word := range words {
			result = append(result, libs.Reverse(word))
		}
	}()

	// Same issue as above
	time.Sleep(50 * time.Millisecond)

	return result
}

func WorkingReverseSingleGoroutine(words []string) []libs.ReverseResult {
	result := []libs.ReverseResult{}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	// Single goroutine handles all words
	go func() {
		for _, word := range words {
			result = append(result, libs.Reverse(word))
		}
		wg.Done()
	}()

	wg.Wait()

	return result
}
