package main

import (
	"fmt"
	"io"
	"log"
	"math/rand/v2"
	"net/http"
	"net/url"
	"os"
	"sync"

	"github.com/bwmarrin/discordgo"
	"resty.dev/v3"

	"github.com/tidwall/gjson"
)

var lastGifURL string
var lastGifURLMu sync.Mutex

func cloudflareEmbedFunc(arg string, bucketDescriptionList []string) []string {
	AccountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	ApiKey := os.Getenv("CLOUDFLARE_API_TOKEN")
	url := "https://api.cloudflare.com/client/v4/accounts/%s/ai/run/@cf/baai/bge-small-en-v1.5"
	formattedUrl := fmt.Sprintf(url, AccountID)
	payloadList := append([]string{arg}, bucketDescriptionList...)
	payload := map[string]any{"text": payloadList}
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application-json").
		SetAuthToken(ApiKey).
		SetBody(payload).
		Post(formattedUrl)

	if err != nil {
		log.Printf("Error in CloudFlare Embed Func : %v , status code : %v", err, resp.StatusCode())
		return []string{err.Error()}
	}
	vectorList := gjson.Get(resp.String(), "result.data")
	// fmt.Println("Response:", vectorList.String())
	vectorResult := vectorList.Array()
	var vectorSlice []string
	for _, vector := range vectorResult {
		vectorSlice = append(vectorSlice, vector.String())
	}
	return vectorSlice
}

func klipySearchResponseParserFunc(body string) string {
	gifSlice := gjson.Get(body, "data.data").Array()
	if len(gifSlice) == 0 {
		return ""
	}

	lastGifURLMu.Lock()
	defer lastGifURLMu.Unlock()

	// ponytail: remembers only the last URL in-process; add per-query history only if repeats remain a problem.
	gifURLs := make([]string, 0, len(gifSlice))
	for _, gif := range gifSlice {
		gifURL := gif.Get("file.hd.gif.url").String()
		if gifURL != "" && gifURL != lastGifURL {
			gifURLs = append(gifURLs, gifURL)
		}
	}

	if len(gifURLs) == 0 {
		for _, gif := range gifSlice {
			if gifURL := gif.Get("file.hd.gif.url").String(); gifURL != "" {
				gifURLs = append(gifURLs, gifURL)
			}
		}
	}
	if len(gifURLs) == 0 {
		return ""
	}

	gifURL := gifURLs[rand.IntN(len(gifURLs))]
	lastGifURL = gifURL
	return gifURL
}

func searchgifFunc(arg string) string {
	arg = url.QueryEscape(arg)
	klipyappKey := os.Getenv("KLIPY_APP_KEY")
	url := "https://api.klipy.com/api/v1/%s/gifs/search?page=1&per_page=10&q=%s"
	formattedURL := fmt.Sprintf(url, klipyappKey, arg)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, formattedURL, nil)

	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	body_str := (string(body))
	gif_url := klipySearchResponseParserFunc(body_str)
	return gif_url
}

func reactgifFunc(arg string, reply *discordgo.MessageReference) string {
	if reply != nil {
		phrase := fetchfromBucket(arg)
		url := searchgifFunc(phrase)
		return url
	} else {
		return "Beep Boop!!\nINCORRECT USAGE DETECTED!!\nReply to someone using \n```!gif react```"
	}

}
