package cancerigo

import (
	"sort"
	"strings"
	"unicode"
)

type precomputed struct {
	message    string
	length     int
	words      []string
	wordsCount int
}

// precompute adds metadata to a message make diagnosis quicker
func precompute(message string) (p precomputed) {
	p.message = message
	p.length = len(message)
	p.words = strings.Fields(message)
	p.wordsCount = len(p.words)
	return
}

// minimumWordCount awards points for one word messages
func minimumWordCount(p precomputed) (points int) {
	// one point per missing word below 2 words
	missing := 2 - p.wordsCount
	if missing > 0 {
		points = missing
	}
	return
}

// minimumMessageLength awards points to short messages
func minimumMessageLength(p precomputed) (points int) {
	// one point for the first character missing
	// then one point every 3 characters missing
	missing := 2 - p.length
	if missing > 0 {
		points = 1 + int(missing/3)
	}
	return
}

// maximumMessageLength awards points to long messages
func maximumMessageLength(p precomputed) (points int) {
	// one point for the first character over the limit
	// then one point every 5 characters
	excess := p.length - 80
	if excess > 0 {
		points = 1 + int(excess/5)
	}
	return
}

// capsRatio awards points to messages with lots of caps
func capsRatio(p precomputed) (points int) {
	if p.length == 0 {
		return
	}

	caps := 0
	for _, char := range p.message {
		if unicode.IsUpper(char) {
			caps++
		}
	}

	ratio := float32(caps)/float32(p.length) - 0.2
	if ratio > 0 {
		points = 1 + int(ratio*2)
	}
	return
}

// countEmotes counts emotes in p (based on emotes.txt)
func countEmotes(p precomputed) (count int) {
	for _, word := range p.words {
		if sort.SearchStrings(Emotes[:], word) < EmotesCount {
			count++
		}
	}
	return
}

// emoteCountAndRatio awards points to messages with lots of emotes
func emoteCountAndRatio(p precomputed) (points int) {
	emotesCount := countEmotes(p)
	count := emotesCount - 1
	if count > 0 {
		points += 1 + int(count/2)
	}

	ratio := float32(emotesCount)/float32(p.wordsCount) - 0.49
	if ratio > 0 {
		points += 1 + int(ratio*2)
	}
	return
}

// countUnique returns the number of unique words
func countUnique(words []string) (count int) {
	seen := make(map[string]bool)
	for _, word := range words {
		if _, found := seen[word]; !found {
			seen[word] = true
			count++
		}
	}
	return
}

// echoRatio awards points to messages with the words repeated
func echoRatio(p precomputed) (points int) {
	if p.wordsCount <= 1 {
		return
	}

	unique := countUnique(p.words)
	ratio := 0.7 - float32(unique)/float32(p.wordsCount)
	if ratio > 0 {
		points = 1 + int(ratio/0.3)
	}
	return
}
