package cancerigo

//go:generate go run emotes.go

// Diagnose returns the sum of cancer points found in the message
func Diagnose(message string) (points int) {
	p := precompute(message)
	points += minimumWordCount(p)
	points += minimumMessageLength(p)
	points += maximumMessageLength(p)
	points += capsRatio(p)
	points += emoteCountAndRatio(p)
	points += echoRatio(p)
	return
}
