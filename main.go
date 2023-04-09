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

	how_many_emojis := 0
	how_many_aliases := 0

	for emojiName, emojiUrl := range listEmojis {
		emojiFullPath := fileDefinePath(emojiName, emojiUrl)

		is_alias, _ := isAlias(emojiUrl)
    	if is_alias {
			log.Printf("INFO: Do not download %v", emojiName)
			how_many_aliases++
			continue
		}

		log.Printf("INFO: Download %v as %v", emojiUrl, emojiFullPath)

		if downloadErr := httpDownloadEmoji(http_client, emojiFullPath, emojiUrl); downloadErr != nil {
			log.Fatal("FATAL: ", downloadErr)
		}

		// Every 100 emojis let's sleep 2 seconds
		how_many_emojis++
		if how_many_emojis%100 == 0 {
			time.Sleep(2 * time.Second)
		}
	}

	log.Printf("------------------------------")
	log.Printf("INFO: Success to download %v emojis", how_many_emojis)
	log.Printf("INFO: Don't download %v aliases", how_many_aliases)
}
