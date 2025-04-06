# Basic concurrency with channels

Use a single worker to reverse a slice of words.

Sicne there's only a single worker, the output array would be in the same order as the input words and hence writes to output array are non-concurrent.
