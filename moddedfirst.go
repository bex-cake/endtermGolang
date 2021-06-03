package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"time"
)

func ModdedFirst() {
	start := time.Now()
	file, err := os.Open("mobydick.txt")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	if err != nil {
		panic(err)
	}
	words := make([]*Word, 20)
	actualSize := 0
	isFound := false
	var tempWord []byte

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		b := scanner.Bytes()[0]
		if (b >= 65 && b <= 90) || (b >= 97 && b <= 122) {
			if b >= 65 && b <= 90 {
				b += 32
			}
			tempWord = append(tempWord, b)
		} else {
			if len(tempWord) > 0 {
				for i := 0; i < actualSize; i++ {
					if bytes.Equal(words[i].bytes, tempWord) {
						words[i].counter++
						isFound = true
						break
					}
				}
				if !isFound {
					words[actualSize] = &Word{
						bytes:   tempWord,
						counter: 1,
					}
					actualSize += 1
					if actualSize >= len(words)-1 {
						tempWords := words
						words = make([]*Word, actualSize*3)
						for i := 0; i < len(tempWords); i++ {
							words[i] = tempWords[i]
						}
					}
				}
				isFound = false
				tempWord = nil
			}
		}
	}
	tempWords := words
	words = make([]*Word, actualSize)
	for i := 0; i < actualSize; i++ {
		words[i] = tempWords[i]
	}
	sort.Slice(words, func(i, j int) bool {
		return words[i].counter > words[j].counter
	})
	for _, w := range words[:20] {
		fmt.Println(w.counter, string(w.bytes))
	}
	fmt.Printf("Process took %s\n", time.Since(start))
}
