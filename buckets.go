package main

import (
	"encoding/json"
	"errors"
	"log"
	"math"
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
	var bucketDescriptionsList []string

	bucketKeys := []string{"shock", "laugh", "hype", "panic", "cringe", "sad", "confused"}

	for _, bucket := range bucketKeys {
		description := bucketDescriptions[bucket]
		bucketDescriptionsList = append(bucketDescriptionsList, description)
	}

	vectorList := cloudflareEmbedFunc(arg, bucketDescriptionsList)
	replyVector, err := parseEmbeddingVector(vectorList[0])
	if err != nil {
		log.Printf("Error in parseEmbeddingVector: %s", err)
	}

	var bucketVectorCollection [][]float64

	for _, vector := range vectorList[1:] {
		bucketVector, err := parseEmbeddingVector(vector)
		if err != nil {
			log.Printf("Error in parseEmbeddingVector: %s", err)
		}
		bucketVectorCollection = append(bucketVectorCollection, bucketVector)
	}

	var similarityList []float64

	for _, vector := range bucketVectorCollection {
		simScore, err := CosineSimilarity(replyVector, vector)
		if err != nil {
			log.Printf("Error : %v", err)
		}
		similarityList = append(similarityList, simScore)
	}

	simIndex, err := findHighestIndex(similarityList)

	simKey := bucketKeys[simIndex]

	return simKey
}

func parseEmbeddingVector(vectorString string) ([]float64, error) {
	var vector []float64

	if err := json.Unmarshal([]byte(vectorString), &vector); err != nil {
		return nil, err
	}

	return vector, nil
}

func CosineSimilarity(vecA, vecB []float64) (float64, error) {
	if len(vecA) != len(vecB) {
		return 0, errors.New("vectors must have the same length")
	}
	if len(vecA) == 0 {
		return 0, errors.New("vectors cannot be empty")
	}

	var dotProduct, normA, normB float64

	for i := 0; i < len(vecA); i++ {
		dotProduct += vecA[i] * vecB[i]
		normA += vecA[i] * vecA[i]
		normB += vecB[i] * vecB[i]
	}

	if normA == 0 || normB == 0 {
		return 0, errors.New("vector magnitude cannot be zero")
	}

	return dotProduct / (math.Sqrt(normA) * math.Sqrt(normB)), nil
}

func findHighestIndex(values []float64) (int, error) {
	if len(values) == 0 {
		return 0, errors.New("the slice cannot be empty")
	}

	highest := values[0]
	highestIndex := 0

	for idx, val := range values {
		if val > highest {
			highestIndex = idx
			highest = val
		}
	}
	return highestIndex, nil
}
