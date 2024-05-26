package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

// Сортирует слайс ключей на основе кол-во вхождения в строку.
// slcWord - слайс ключей мапы mpWord.
// mpWord - мапа где ключ слово, а значений кол-во вхождений.
func sortSliceByCount(slcWord []string, mpWord map[string]int) []string {
	sort.SliceStable(slcWord, func(i, j int) bool {
		return mpWord[slcWord[i]] > mpWord[slcWord[j]]
	})
	return slcWord
}

// Лексикографическая сортировка слайса с учетом кол-во вхождений.
func finalSortSliceByLexicAndCount(slcWord []string, mpWord map[string]int) []string {
	sort.Slice(slcWord, func(i, j int) bool {
		if mpWord[slcWord[i]] == mpWord[slcWord[j]] {
			return slcWord[i] < slcWord[j]
		}
		return mpWord[slcWord[i]] > mpWord[slcWord[j]]
	})
	return slcWord
}

func Top10(in string) []string {
	if len(in) == 0 {
		return []string{}
	}

	splitedIn := strings.Fields(in)
	countMap := make(map[string]int)
	for _, word := range splitedIn {
		_, ok := countMap[word]
		if !ok {
			countMap[word] = 1
		} else {
			countMap[word]++
		}
	}

	freqWords := make([]string, 0, len(countMap))

	for key := range countMap {
		freqWords = append(freqWords, key)
	}

	freqWords = sortSliceByCount(freqWords, countMap)
	if len(freqWords) > 10 {
		return finalSortSliceByLexicAndCount(freqWords[:10], countMap)
	}
	return finalSortSliceByLexicAndCount(freqWords, countMap)
}
