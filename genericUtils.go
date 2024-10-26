package main

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"strings"
)

func substringBeforeFirst(input string, delimiter string) string {
	index := strings.Index(input, delimiter)
	if index == -1 {
		return ""
	}
	return strings.TrimSpace(input[:index])
}

func md5Hash(text string) string {
	data := []byte(text)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func removeDuplicates(elements []string) []string {
	seen := make(map[string]struct{})
	result := []string{}

	for _, element := range elements {
		if _, found := seen[element]; !found {
			seen[element] = struct{}{}
			result = append(result, element)
		}
	}

	return result
}

func sliceContainsString(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func baseURL(inputURL string) (string, error) {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}

	// Construct base URL with scheme and host
	baseURL := fmt.Sprintf("%s://%s", parsedURL.Scheme, parsedURL.Host)
	return baseURL, nil
}
