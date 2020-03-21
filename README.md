# Cancerigo

A simple library used to award "cancer points" to chat messages from Twitch.

Golang re-implementation of [TwichCancer/symptom](https://github.com/Benzhaomin/twitchcancer/tree/master/twitchcancer/symptom).

## Usage

Basically, this lib can be summarized by its main function `func Diagnose(message string) (points int)`.

```go
>>> Diagnose("some message with lots of interesting things and Kappa Kappa emotes")
5
```

## Dev

Edit [emotes.txt](emotes.txt) and run `go generate`.

Test with `go test`

Benchmark with `go test -test.bench=.`
