package main

import (
	"os"
	"strings"
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