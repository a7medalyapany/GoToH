package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

type Pair struct {
	Key   string
	Value int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run main.go <file-path>")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	wordsFrequency, err := freq(file)
	if err != nil {
		log.Fatal(err)
	}

	sortedWords := sortWordsFrequency(wordsFrequency)

	fmt.Println("Word frequencies:")
	for _, p := range sortedWords {
		fmt.Printf("%s: %d\n", p.Key, p.Value)
	}
}

func freq(r io.Reader) (map[string]int, error) {
	wordsFrequency := make(map[string]int)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)

	for s.Scan() {
		word := strings.ToLower(s.Text())
		word = strings.Trim(word, ".,!?;:\"'()[]{}")
		wordsFrequency[word]++
	}

	return wordsFrequency, s.Err()
}

func sortWordsFrequency(m map[string]int) []Pair {
	pairs := make([]Pair, 0, len(m))
	for k, v := range m {
		pairs = append(pairs, Pair{Key: k, Value: v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})

	return pairs
}
