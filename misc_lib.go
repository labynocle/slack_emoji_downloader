package main

import (
	"regexp"
)

func isAlias(emoji_url string) (bool, error) {

	return regexp.MatchString("^alias", emoji_url)
}
