package main

import (
	"path/filepath"
	"github.com/gosimple/slug"
)

var downloadDirectory = "/tmp/emojis"

func fileDefinePath(emojiName string, emojiURL string) string {
	emojiSlugName := slug.Make(emojiName)
    emojiExtension := filepath.Ext(emojiURL)
    return filepath.Join(downloadDirectory, emojiSlugName[0:1], emojiSlugName+emojiExtension)
}


func dirDefinePath(emojiFullPath string) string {
	return filepath.Dir(emojiFullPath)
}
