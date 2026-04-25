package views

import "strings"

func decodePath(encoded string) string {
	if encoded == "" {
		return ""
	}
	s := strings.TrimPrefix(encoded, "-")
	s = strings.ReplaceAll(s, "--", "\x00")
	s = strings.ReplaceAll(s, "-", "/")
	s = strings.ReplaceAll(s, "\x00", "-")
	return "/" + s
}
