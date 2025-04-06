# Basic goroutine

Here we have multiple implementations.

1. BadReverseMultipleGoroutines

    Here, we use multiple goroutines, each of which processes a single word to reverse.

    There are multiple problems in this approach.

    1. When we return the result value from the function, it won't have processed all the goroutines. Hence, we'd need to wait for a roughly estimated time before returning the result.
    2. Since each goroutine appends the ReverseResult to the same result variable, it's a possibility of race condition.

2. BadReverseSingleGoroutine

    This implementation uses a single goroutine and the looping of words to generate ReverseResult happens within this goroutine. Hence, there's no possibility of a race condition here.

    However, since this has other problems.

    1. Since we loop within the goroutine, this goroutine would take as much time as the program would have taken without a goroutine. Hence, there's no much point in having this in a separate goroutine. I consider this a classic case of over-engineering.
    2. Again, if we don't wait before returning the result, the goroutine would not have finished and would result in return value not having results of all the words in input.

3. WorkingReverseSingleGoroutine

    This implementation solves the problem of having to wait using `time.Sleep`. This means that the result would always contain the result of all the processed words. However, it still loops through the words within the single goroutine and hence is not much useful in most cases.

    This can however come in handy in certain scenarios.

    For example, say `WorkingReverseSingleGoroutine` takes 50ms to process the words. And in the meantime if the main thread want to make an http call which takes longer to process. In this case, if the host has multiple cores, the main thread and the goroutine would run in parallel, thus saving time.
