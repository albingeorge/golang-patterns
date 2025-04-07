# Basic goroutine

Here we have multiple implementations.

1. `BadReverseMultipleGoroutines`

    Here, we use multiple goroutines, each of which processes a single word to reverse.

    There are multiple problems in this approach.

    1. When we return the result value from the function, it won't have processed all the goroutines. Hence, we'd need to wait for a roughly estimated time before returning the result.
    2. Since each goroutine appends the ReverseResult to the same result variable, it's a possibility of race condition.

2. `BadReverseSingleGoroutine`

    This implementation uses a single goroutine and the looping of words to generate ReverseResult happens within this goroutine. Hence, there's no possibility of a race condition here.

    However, since this has other problems.

    1. Since we loop within the goroutine, this goroutine would take as much time as the program would have taken without a goroutine. Hence, there's no much point in having this in a separate goroutine. I consider this a classic case of over-engineering.
    2. Again, if we don't wait before returning the result, the goroutine would not have finished and would result in return value not having results of all the words in input.

3. `WorkingReverseSingleGoroutine`

    This implementation solves the problem of having to wait using `time.Sleep`. This means that the result would always contain the result of all the processed words. However, it still loops through the words within the single goroutine and hence is not much useful in most cases.

    This can however come in handy in certain scenarios.

    For example, say `WorkingReverseSingleGoroutine` takes 50ms to process the words. And in the meantime if the main thread want to make an http call which takes longer to process. In this case, if the host has multiple cores, the main thread and the goroutine would run in parallel, thus saving time.

    Diagram:

    ```mermaid
    flowchart LR
        A[start] --> |main thread| C[end]
        A --> B(WorkingReverseSingleGoroutine)
        B --> C

    ```

4. `ReverseMultipleGoroutines`

    `ReverseMultipleGoroutines` solves all of the problems mentioned above.

    1. Has multiple goroutines, each which can run the core logic which takes time concurrently
    2. Does not have a race condition, since we're using a Mutex when writing to the result array
    3. Makes use of WaitGroup to ensure that all the goroutines are processed before returning result

    Diagram:

    ```mermaid
    flowchart LR
        A[start] --> C1(word1)
        A --> C2(word2)
        A --> C3(word3)
        subgraph ReverseMultipleGoroutines
            C1
            C2
            C3
            end

        C1 --> B[end]
        C2 --> B
        C3 --> B
    ```

## Benchmarks

These are the benchmark results.

```sh
$ go test -bench=.
goos: darwin
goarch: arm64
pkg: github.com/albingeorge/golang-patterns/concurrency/1_basic_goroutine
cpu: Apple M1
BenchmarkWorkingReverseSingleGoroutine-8   	       1	50968679417 ns/op
BenchmarkReverseMultipleGoroutines-8       	      22	  51789165 ns/op
PASS
ok  	github.com/albingeorge/golang-patterns/concurrency/1_basic_goroutine	53.178s
```

Here, the `ReverseMultipleGoroutines` has much higher lower ns/op, since it ran each call of the `libs.Reverse()` function concurrently.

> Note: I've not written benchmarks for the functions `BadReverseMultipleGoroutines` and `BadReverseSingleGoroutine`, since they run in constant time and would not yield the right results depending on the number of words to process.
