package main

import (
	"regexp"
)

func isAlias(emojiUrl string) (bool, error) {
	return regexp.MatchString("^alias", emojiUrl)
}
