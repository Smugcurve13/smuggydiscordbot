package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"resty.dev/v3"
	"github.com/bwmarrin/discordgo"
	"github.com/tidwall/gjson"
)

func cloudflareEmbedFunc(arg string, description string) string {
	AccountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	ApiKey := os.Getenv("CLOUDFLARE_API_TOKEN")
	url := "https://api.cloudflare.com/client/v4/accounts/%s/ai/run/@cf/baai/bge-small-en-v1.5"
	formattedUrl := fmt.Sprintf(url, AccountID)

	payload := map[string]any{"text": []string{arg,description}}
	client := resty.New()

	resp, err := client.R().
			SetHeader("Content-Type", "application-json").
			SetAuthToken(ApiKey).
			SetBody(payload).
			Post(formattedUrl)
	
	if err != nil {
		log.Printf("Error in CloudFlare Embed Func : %v , status code : %v", err, resp.StatusCode())
		return ""
	}
	vector := gjson.Get(resp.String(), "result.data.1")
	fmt.Println("Response:", vector.String())
	return ""
}

func klipySearchResponseParserFunc(body string) string {
	url := gjson.Get(body, "data.data.0.file.hd.gif.url")
	return url.Str
}

func searchgifFunc(arg string) string {
	arg = url.QueryEscape(arg)
	klipyappKey := os.Getenv("KLIPY_APP_KEY")
	url := "https://api.klipy.com/api/v1/%s/gifs/search?page=1&per_page=1&q=%s"
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
	// search the arg against 6 buckets
	// classify the arg with a bucket
	// randomise a phrase from that bucket
	// return the url

	// we need buckets to store 
	// names of buckets
	// phrases in the buckets
	// smash the arg at the bucket
	if reply != nil{
		phrase := fetchfromBucket(arg)
		url := searchgifFunc(phrase)
		return url
	} else {
		return "Beep Boop!!\nINCORRECT USAGE DETECTED!!\nReply to someone using \n```!gif react```"
	}
	
}