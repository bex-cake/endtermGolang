package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"time"
)

type Word struct {
	bytes   []byte
	counter int
}

func MySolution() {
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

	var words []*Word
	isFound := false
	var tempWord []byte

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		b:= scanner.Bytes()[0]
		if (b >= 65 && b <= 90) || (b >= 97 && b <= 122) {
			if b >= 65 && b <= 90 {
				b += 32
			}
			tempWord = append(tempWord, b)
		} else {
			if len(tempWord) > 0 {
				for _, w := range words {
					if bytes.Equal(w.bytes, tempWord) {
						w.counter++
						isFound = true
						break
					}
				}
				if !isFound {
					words = append(words, &Word{
						bytes: tempWord,
						counter:  1,
					})
				}
				isFound = false
				tempWord = nil
			}
		}
	}
	sort.Slice(words, func(i, j int) bool {
		return words[i].counter > words[j].counter
	})
	for _, w := range words[:20] {
		fmt.Println(w.counter, string(w.bytes))
	}
	fmt.Printf("Process took %s\n", time.Since(start))
}
