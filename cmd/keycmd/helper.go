package keycmd

import "strings"

func sanitizeName(name string) string {
	sanitized := strings.ReplaceAll(name, " ", "-")
	if sanitized == "" {
		return "unnamed"
	}
	return sanitized
}
