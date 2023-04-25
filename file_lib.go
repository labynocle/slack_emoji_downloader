package main

import (
	"path/filepath"
	"github.com/gosimple/slug"
	"fmt"
)

var downloadDirectory = "/tmp/emojis"

func fileDefinePath(emojiName string, emojiURL string) (string,error) {
	emojiSlugName := slug.Make(emojiName)
	emojiExtension := filepath.Ext(emojiURL)

	if len(emojiSlugName) == 0 {
		return "", fmt.Errorf("fileDefinePath - emojiSlugName is empty")
	}

    return filepath.Join(downloadDirectory, emojiSlugName[0:1], emojiSlugName+emojiExtension), nil
}


func dirDefinePath(emojiFullPath string) string {
	return filepath.Dir(emojiFullPath)
}
