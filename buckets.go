package main

import (
	"math/rand/v2"
)

var collection = map[string][]string{
	"shock": {
		"shocked reaction",
		"jaw drop",
		"mind blown",
		"surprised face",
	},
	"laugh": {
		"laughing reaction",
		"dying laughing",
		"spit take",
		"can't stop laughing",
	},
	"hype": {
		"lets go reaction",
		"celebration dance",
		"winning reaction",
		"slow clap",
	},
	"panic": {
		"panic reaction",
		"oh no reaction",
		"screaming internally",
		"disaster",
	},
	"cringe": {
		"awkward reaction",
		"secondhand embarrassment",
		"side eye",
		"disgusted face",
	},
	"sad": {
		"crying reaction",
		"devastated reaction",
		"disappointed face",
		"sad reaction",
	},
	"confused": {
		"confused reaction",
		"what is happening",
		"blinking confused",
		"question marks",
	},
}

func fetchfromBucket(arg string) string {
	bucketPhrases := collection["shock"]
	randomPhraseIndex := rand.N(len(bucketPhrases))
	selectedPhrase := bucketPhrases[randomPhraseIndex]
	return selectedPhrase
}