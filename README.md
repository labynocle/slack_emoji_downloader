# slack_emoji_downloader
Tiny go script to download all custom emoji of a given Slack space

It's just an excuse to discover how Golang is working ^^

Usage:

```
export SLACK_TOKEN=xoxb-XXXX-XXXX-XXXX
export SLACK_URL=https://YOUR-NAME.slack.com

# Launch download
make run

# Launch the tests
make test

# Display help
make
```

All emojis will be found in `/tmp/emojis` directory

ToDo:

* add basic tests
* add github ci
