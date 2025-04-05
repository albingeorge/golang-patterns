package main

import (
	"testing"

	"github.com/albingeorge/golang-patterns/libs"
	"github.com/go-test/deep"
)

func TestWorker(t *testing.T) {
	type test struct {
		input  []string
		output []libs.ReverseResult
	}

	tests := []test{
		{
			input: []string{
				"test",
				"CapsTest",
			},
			output: []libs.ReverseResult{
				{
					Input:   "test",
					Reverse: "tset",
				},
				{
					Input:   "CapsTest",
					Reverse: "tseTspaC",
				},
			},
		},
	}

	for _, tt := range tests {
		res := reverse(tt.input)
		if diff := deep.Equal(res, tt.output); diff != nil {
			t.Error(diff)
		}
	}
}
