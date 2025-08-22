package main

import "regexp"

func RemoveANSI(data []byte) []byte {
	text := string(data)
	ansiRegex := regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)
	return []byte(ansiRegex.ReplaceAllString(text, ""))
}
