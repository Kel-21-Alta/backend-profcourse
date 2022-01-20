package helpers

import "strings"

func CheckItemInSlice(arr []string, item string) bool {
	for _, e := range arr {
		if strings.ToLower(e) == strings.ToLower(item) {
			return true
		}
	}
	return false
}
