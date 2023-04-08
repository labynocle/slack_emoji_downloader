package main

import (
	"time"
	"log"
)




func main() {

	prog_name := "slack_emoji_downloader"

	log.SetPrefix(prog_name + ": ")
    log.SetFlags(0)

    http_client := httpClient()

	log.Println("INFO: Get emojis list")
	listEmojis, err := httpGetListEmojis(http_client)
	if err != nil {
		log.Fatal("FATAL: ", err)
	}

	how_many_emojis := 0
	how_many_aliases := 0

	for emoji_name, emoji_url := range listEmojis {
		emoji_full_path := fileDefinePath(emoji_name, emoji_url)

		is_alias, _ := isAlias(emoji_url)
    	if is_alias {
			log.Printf("INFO: Do not download %v", emoji_name)
			how_many_aliases++
			continue
		}

		log.Printf("INFO: Download %v as %v", emoji_url, emoji_full_path)

		download_err := httpDownloadEmoji(http_client, emoji_full_path, emoji_url)
		if download_err != nil {
			log.Fatal("FATAL: ", download_err)
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
