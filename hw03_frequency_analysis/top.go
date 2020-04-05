package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"sort"
	"strings"
)

const MaxWords = 10

type WordCounter struct {
	Word  string
	Count int
}

func TrimNewlineAndTabInString(text string) string {
	text = strings.ReplaceAll(strings.ReplaceAll(text, "	", ""), "\n", " ")
	return text
}

func Top10(text string) []string {
	words := strings.Split(TrimNewlineAndTabInString(text), " ")
	//
	counters := make(map[string]*WordCounter)

	for _, word := range words {
		// skip empty word
		if word == "" {
			continue
		}
		if _, ok := counters[word]; ok {
			counter := counters[word]
			counter.Count++
		} else {
			counters[word] = &WordCounter{word, 1}
		}
	}

	// convert map to slice of values
	values := []WordCounter{}
	for _, value := range counters {
		values = append(values, *value)
	}
	// sort values by count
	sort.Slice(values, func(i, j int) bool {
		return values[i].Count > values[j].Count
	})

	result := []string{}

	for _, value := range values {
		result = append(result, value.Word)
	}

	if len(result) > MaxWords {
		result = result[:MaxWords]
	}

	return result
}
