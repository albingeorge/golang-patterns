package main

import (
	"fmt"
	"log"
	"time"
)

func reverse(words []string) []result {
	log.Printf("%v\n", words)
	inputCh, outChan := make(chan string), make(chan result, 10)

	// Run worker
	go worker(inputCh, outChan)
	go worker(inputCh, outChan)

	go func() {
		for _, word := range words {
			inputCh <- word
		}
		close(inputCh)
	}()

	count := 0
	out := []result{}
	for res := range outChan {
		fmt.Printf("Input: %s; Reverse: %s\n", res.input, res.reverse)
		out = append(out, res)
		count++
	}

	if count != len(words) {
		log.Fatalf("Some words not processed! Word count: %d; Processed count: %d\n", len(words), count)
	}
	log.Println("Count:", count)

	return out
}

func worker(ch <-chan string, out chan<- result) {
	for in := range ch {
		out <- reverseimpl(in)
	}
	close(out)
}

func reverseimpl(in string) result {
	res := make([]rune, len(in))

	for i, val := range in {
		res[len(in)-i-1] = val
	}

	// Introduce a delay to check if long running processes are impacted
	time.Sleep(1000 * time.Millisecond)

	return result{
		input:   in,
		reverse: string(res),
	}
}
