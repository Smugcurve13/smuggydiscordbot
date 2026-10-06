package main

import (
	"math/rand/v2"
)

var shockedCollection = map[string][]string{
	"shocked": {
		"shocked reaction",
		"jaw drop",
		"mind blown",
		"surprised face",
	},
}

func fetchfromBucket(arg string) string {
	bucketPhrases := shockedCollection["shocked"]
	randomPhraseIndex := rand.N(len(bucketPhrases))
	selectedPhrase := bucketPhrases[randomPhraseIndex]
	return selectedPhrase
}