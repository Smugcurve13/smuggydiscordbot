package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"google.golang.org/genai"
)

func getWhitelistedIDS() ([]string) {
	ids := strings.Split(os.Getenv("WHITELISTED_USER_IDS"), ",")
	WHITELISTED_IDS := []string{}
	for i := range ids {
		ids[i] = strings.TrimSpace(ids[i])
		WHITELISTED_IDS = append(WHITELISTED_IDS, ids[i])
	}
	return WHITELISTED_IDS
}

func geminiClient(ctx context.Context, apiKey string) (*genai.Client) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		fmt.Printf("The Error is : %s",err)
	}
	return client
}

func cleanGeminiResponse(resp *genai.GenerateContentResponse) string {
	var result string

	for _, candidate := range resp.Candidates {
		for _, part := range candidate.Content.Parts {
			if part != nil && part.Text != "" {
				result += part.Text
			}
		}
	}

	return result
}