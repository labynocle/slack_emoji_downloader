package main

import (
	"path/filepath"
	"github.com/gosimple/slug"
)

var DOWNLOAD_DIRECTORY = "/tmp/emojis"

func fileDefinePath(emoji_name string, emoji_url string) string {
	emoji_slug_name := slug.Make(emoji_name)
	return DOWNLOAD_DIRECTORY + "/" + emoji_slug_name[0:1] + "/" + emoji_slug_name + filepath.Ext(emoji_url)
}


func dirDefinePath(emoji_full_path string) string {
	return filepath.Dir(emoji_full_path)
}
