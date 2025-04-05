package main

import (
	"log"
	"sync"

	"github.com/albingeorge/golang-patterns/libs"
)

func reverse(words []string) []libs.ReverseResult {
	log.Printf("%v\n", words)
	inputCh, outChan := make(chan string), make(chan libs.ReverseResult, 10)

	// Run worker
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		worker(inputCh, outChan)
	}()

	// Close outChan after all workers finish
	go func() {
		wg.Wait()      // Wait for all workers to finish
		close(outChan) // Close the output channel
	}()

	go func() {
		for _, word := range words {
			inputCh <- word
		}
		close(inputCh)
	}()

	count := 0
	out := []libs.ReverseResult{}
	for res := range outChan {
		log.Printf("Input: %s; Reverse: %s\n", res.Input, res.Reverse)
		out = append(out, res)
		count++
	}

	if count != len(words) {
		log.Fatalf("Some words not processed! Word count: %d; Processed count: %d\n", len(words), count)
	}
	log.Println("Count:", count)

	return out
}

func worker(ch <-chan string, out chan<- libs.ReverseResult) {
	for in := range ch {
		out <- libs.Reverse(in)
	}
}
