package solver

import (
	"bufio"
	"errors"
	"io/ioutil"
	"log"
	"sort"
	"strings"
	"sync"
)

type byLength []string

var wg sync.WaitGroup

// FindWordsForLetters finds possible words for the letters provided
func FindWordsForLetters(letters string) ([]string, error) {
	var resultsChannel = make(chan string)
	letters = strings.ToLower(letters)
	if isNotValid(letters) {
		return []string{}, errors.New("Not valid letters")
	}
	for _, letter := range uniqueLetters(letters) {
		wg.Add(1)
		go findWords(letters, letter, resultsChannel)
	}

	go func() {
		wg.Wait()
		close(resultsChannel)
	}()

	results := make([]string, 0)
	for result := range resultsChannel {
		results = append(results, result)
	}
	sort.Sort(byLength(results))
	return results, nil
}

func findWords(word, firstLetter string, resultsChannel chan string) {
	bs, err := ioutil.ReadFile("./word-lists/" + firstLetter + "-list")
	if err != nil {
		log.Fatal(err)
	}
	lengthOfWord := len([]rune(word))
	wordFrequency := buildHistogram(word)

	scanner := bufio.NewScanner(strings.NewReader(string(bs)))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		// scanner.Text() for current word
		currentWord := scanner.Text()

		// check length
		if len([]rune(currentWord)) > lengthOfWord {
			continue
		}

		possibleWord := true
		currentWordFrequency := buildHistogram(currentWord)
		// iterate over word and check letters
		for letter, frequency := range currentWordFrequency {
			possibleWord = strings.Contains(word, letter) && possibleWord
			possibleWord = (frequency <= wordFrequency[letter]) && possibleWord
		}

		if !possibleWord {
			continue
		}
		resultsChannel <- currentWord
	}
	wg.Done()
}

func buildHistogram(word string) map[string]int {
	histogram := make(map[string]int)
	for _, letter := range []rune(word) {
		histogram[string(letter)]++
	}
	return histogram
}

func uniqueLetters(word string) []string {
	result := make([]string, 0)
	for _, letter := range []rune(word) {
		if !sliceContains(result, string(letter)) {
			result = append(result, string(letter))
		}
	}
	return result
}

func sliceContains(slice []string, letter string) bool {
	for _, currentLetter := range slice {
		if currentLetter == letter {
			return true
		}
	}
	return false
}

func isNotValid(letters string) bool {
	valid := true
	for _, letter := range []rune(letters) {
		valid = valid && 96 < letter && letter < 123
	}
	return !valid
}

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}
