package bob

import "strings"

const TestVersion = 1

func Hey(drivel string) string {
	switch {
	case silent(drivel):
		return "Fine. Be that way!"
	case yelling(drivel):
		return "Whoa, chill out!"
	case asking(drivel):
		return "Sure."
	default:
		return "Whatever."
	}
}

func yelling(drivel string) bool {
	return strings.ToUpper(drivel) == drivel &&
		strings.ToLower(drivel) != strings.ToUpper(drivel)
}

func asking(drivel string) bool {
	return strings.HasSuffix(drivel, "?")
}

func silent(drivel string) bool {
	return strings.TrimSpace(drivel) == ""
}
