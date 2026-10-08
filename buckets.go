package main

import (
	"fmt"
	"maps"
	"math/rand/v2"
	"slices"
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

var bucketDescriptions = map[string]string{
	"shock":    "A message expresses surprise, disbelief, or something mind-blowing.",
	"laugh":    "A message is funny, ridiculous, or makes people laugh.",
	"hype":     "A message celebrates a win, success, excitement, or an impressive achievement.",
	"panic":    "A message describes a problem, disaster, danger, or an alarming situation.",
	"cringe":   "A message is awkward, embarrassing, uncomfortable, or painfully bad.",
	"sad":      "A message shares disappointment, loss, bad news, or emotional sadness.",
	"confused": "A message is strange, unclear, nonsensical, or difficult to understand.",
}

func fetchfromBucket(arg string) string {
	closestBucket := findClosestBucket(arg)
	bucketPhrases := collection[closestBucket]
	randomPhraseIndex := rand.N(len(bucketPhrases))
	selectedPhrase := bucketPhrases[randomPhraseIndex]
	return selectedPhrase
}

func findClosestBucket(arg string) string {
	bucketDescriptionsList := slices.Collect(maps.Values(bucketDescriptions))
	vectorList := cloudflareEmbedFunc(arg, bucketDescriptionsList)
	fmt.Println("/n", vectorList)
	return "shock"
}