package main

import (
    "testing"
)

type isAliasTest struct {
    input    string
    expected bool
}

var isAliasTests = []isAliasTest{
    {"https://emoji.slack-edge.com/T03LS0K23/drapeau_basque/577c2e6efe9dabc6.png", false},
    {"alias-plop", true},
    {"", false},
}

func TestIsAlias(t *testing.T) {
    for _, tt := range isAliasTests {
        actual, err := isAlias(tt.input)
        if actual != tt.expected || err != nil {
            t.Errorf("isAlias(%s): expected %t, actual %t, err %v", tt.input, tt.expected, actual, err)
        }
    }
}
