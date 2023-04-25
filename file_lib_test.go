package main

import (
    "testing"
)

type isFileDefinePathTest struct {
    input_name string
    input_url  string
    expected   string
}

var isFileDefinePathTests = []isFileDefinePathTest{
    {"plop", "https://slack.com/plop.png", "/tmp/emojis/p/plop.png"},
    {"plop%toto", "https://slack.com/toto.jpeg", "/tmp/emojis/p/plop-toto.jpeg"},
}

var isFileDefinePathErrorTests = []isFileDefinePathTest{
    {"++", "https://slack.com/plop.png", ""},
}

func TestFileDefinePath(t *testing.T) {
    for _, tt := range isFileDefinePathTests {
        actual, err := fileDefinePath(tt.input_name, tt.input_url)
        if actual != tt.expected || err != nil {
            t.Errorf("fileDefinePath(%s,%s): expected %s, actual %s, err %v", tt.input_name, tt.input_url, tt.expected, actual, err)
        }
    }
}

func TestFileDefinePathError(t *testing.T) {
	for _, tt := range isFileDefinePathErrorTests {
        actual, err := fileDefinePath(tt.input_name, tt.input_url)
        if actual != "" || err == nil {
            t.Errorf("fileDefinePath(%s,%s): expected %s, actual %s, err %v", tt.input_name, tt.input_url, tt.expected, actual, err)
        }
    }
}
