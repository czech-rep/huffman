package main

import (
	"math"
	"strings"
)

func stringToSlice(word string) []string {
	result := []string{}
	for _, letter := range word {
		result = append(result, string(letter))
	}
	return result
}
func sliceToString(slice []string) string {
	return strings.Join(slice, "")
}
func MinimumKeyByValue(mapped map[string]int) string {
	var minKey string
	minValue := math.MaxInt

	for key, value := range mapped {
		if value < minValue {
			minValue = value
			minKey = key
		}
	}
	return minKey
}
func Counted(items []string) map[string]int {
	counter := make(map[string]int)
	for _, item := range items {
		counter[item]++
	}
	return counter
}
