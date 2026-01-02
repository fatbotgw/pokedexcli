package main

import (
	"strings"
)

// The purpose of this function will be to split the user's input into "words"
// based on whitespace. It should also lowercase the input and trim any leading
// or trailing whitespace.
func cleanInput(text string) []string {

	retString := strings.Fields(strings.ToLower(text))

    return retString
}
