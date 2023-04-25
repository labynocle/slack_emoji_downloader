package main

import (
	"time"
	"log"
)

func main() {
	progName := "slack_emoji_downloader"

	log.SetPrefix(progName + ": ")
	log.SetFlags(0)

	http_client := httpClient()

	log.Println("INFO: Get emojis list")
	listEmojis, err := httpGetListEmojis(http_client)
	if err != nil {
		log.Fatal("FATAL: ", err)
	}

	howManyEmojis := 0
	howManyAliases := 0
	howManyNoslugify := 0

	for emojiName, emojiUrl := range listEmojis {
		emojiFullPath, err := fileDefinePath(emojiName, emojiUrl)

		if err != nil {
			log.Printf("INFO: Do not download %v because unable to slugify", emojiName)
			howManyNoslugify++
			continue
		}

		is_alias, _ := isAlias(emojiUrl)
		if is_alias {
			log.Printf("INFO: Do not download %v because it's an alias", emojiName)
			howManyAliases++
			continue
		}

		log.Printf("INFO: Download %v as %v", emojiUrl, emojiFullPath)

		if downloadErr := httpDownloadEmoji(http_client, emojiFullPath, emojiUrl); downloadErr != nil {
			log.Fatal("FATAL: ", downloadErr)
		}

		// Every 100 emojis let's sleep 2 seconds
		howManyEmojis++
		if howManyEmojis%100 == 0 {
			time.Sleep(2 * time.Second)
		}
	}

	log.Printf("------------------------------")
	log.Printf("INFO: Success to download %v emojis", howManyEmojis)
	log.Printf("INFO: Don't download %v aliases", howManyAliases)
	log.Printf("INFO: Don't download %v slugify", howManyNoslugify)
}
