package main

import (
    "net/http"
    "time"
	"errors"
	"os"
	"encoding/json"
	"io"
)

var SLACK_TOKEN = os.Getenv("SLACK_TOKEN")
var SLACK_URL = os.Getenv("SLACK_URL")

type slack_emoji_resp struct {
	Ok bool `json:"ok"`
	Emoji map[string]string `json:"emoji"`
	Cache_ts string `json:"cache_ts"`
}


func httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

func httpDownloadEmoji(client *http.Client, filepath string, url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return errors.New("httpDownloadEmoji - Unable to forge GET request")
	}

    resp, err := client.Do(req)
	if err != nil {
		return errors.New("httpDownloadEmoji - Unable to download")
	}

	defer resp.Body.Close()

	dirpath := dirDefinePath(filepath)
	if _, err := os.Stat(dirpath); os.IsNotExist(err) {
		if err := os.MkdirAll(dirpath, os.ModePerm); err != nil {
			return errors.New("httpDownloadEmoji - Unable to create directory")
		}
	}

	out, err := os.Create(filepath)
	if err != nil {
		return errors.New("httpDownloadEmoji - Unable to write file locally")
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return errors.New("httpDownloadEmoji - Unable to finish writing file locally")
	}

	return nil
}

func httpGetListEmojis(client *http.Client) (map[string]string, error) {
	req, err := http.NewRequest("GET", SLACK_URL + "/api/emoji.list", nil)
	if err != nil {
		return nil, errors.New("httpGetListEmojis - Unable to forge GET request")
	}

    req.Header.Add("Authorization", "Bearer " + SLACK_TOKEN)

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("httpGetListEmojis - Unable to process GET")
	}

    defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("httpGetListEmojis - Unable to read response body")
	}

	var target slack_emoji_resp
	json.Unmarshal(body, &target)

	return target.Emoji, nil
}
