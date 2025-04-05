package libs

import "time"

type ReverseResult struct {
	Input   string
	Reverse string
}

func Reverse(in string) ReverseResult {
	res := make([]rune, len(in))

	for i, val := range in {
		res[len(in)-i-1] = val
	}

	// Introduce a delay to check if long running processes are impacted
	time.Sleep(50 * time.Millisecond)

	return ReverseResult{
		Input:   in,
		Reverse: string(res),
	}
}
