package main

import (
    "net/http"
    "time"
	"os"
	"encoding/json"
	"io"
	"fmt"
)

var slackToken = os.Getenv("SLACK_TOKEN")
var slackUrl = os.Getenv("SLACK_URL")

type slackEmojiResp struct {
	Ok       bool              `json:"ok"`
	Emoji    map[string]string `json:"emoji"`
	Cache_ts string            `json:"cache_ts"`
}

func httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

func httpDownloadEmoji(client *http.Client, filepath string, url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("httpDownloadEmoji - Unable to forge GET request: %w", err)
	}

    resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("httpDownloadEmoji - Unable to downlowd: %w", err)
	}

	defer resp.Body.Close()

	dirpath := dirDefinePath(filepath)
	if _, err := os.Stat(dirpath); os.IsNotExist(err) {
		if err := os.MkdirAll(dirpath, os.ModePerm); err != nil {
			return fmt.Errorf("httpDownloadEmoji - Unable to create directory: %w", err)
		}
	}

	out, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("httpDownloadEmoji - Unable to write file locally: %w", err)
	}

	defer out.Close()

	if _, err = io.Copy(out, resp.Body); err != nil {
		return fmt.Errorf("httpDownloadEmoji - Unable to finish writing file locally: %w", err)
	}

	return nil
}

func httpGetListEmojis(client *http.Client) (map[string]string, error) {
	req, err := http.NewRequest("GET", slackUrl + "/api/emoji.list", nil)
	if err != nil {
		return nil, fmt.Errorf("httpGetListEmojis - Unable to forge GET request: %w", err)
	}

    req.Header.Add("Authorization", "Bearer " + slackToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("httpGetListEmojis - Unable to process GET: %w", err)
	}

    defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("httpGetListEmojis - Unable to read response body: %w", err)
	}

	var target slackEmojiResp
	json.Unmarshal(body, &target)

	return target.Emoji, nil
}
