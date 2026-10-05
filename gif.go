package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/tidwall/gjson"
)

func searchgifFunc(arg string) string {
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

func klipySearchResponseParserFunc(body string) string {
	url := gjson.Get(body, "data.data.0.file.hd.gif.url")
	return url.Str
}