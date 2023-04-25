# slack_emoji_downloader

![example workflow](https://github.com/labynocle/slack_emoji_downloader/actions/workflows/go.yml/badge.svg)

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

* use https://github.com/spf13/viper to create real CLI
