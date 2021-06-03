package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"sync"
	"time"
)

func contain(words []*Word, word []byte) (bool, int) {
	for i := 0; i < len(words); i++ {
		if string(words[i].bytes) == string(word) {
			return true, i
		}
	}
	return false, -1
}

type Word struct {
	bytes   []byte
	counter int32
}

func SecondSolution() {
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

	scanner := bufio.NewScanner(file)
	var words []*Word
	wg := &sync.WaitGroup{}
	for scanner.Scan() {
		if len(scanner.Bytes()) != 0 {
			in := make(chan []byte)
			out := make(chan []byte)
			wg.Add(1)

			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				var tempWord []byte

				byteLines := <-in
				for i := 0; i < len(byteLines)+1; i++ {
					var b byte
					if i != len(byteLines) {
						b = byteLines[i]
					}
					if !((b >= 65 && b <= 90) || (b >= 97 && b <= 122)) || (i == len(byteLines)) {
						if len(tempWord) > 0 {
							out <- tempWord
						}
						tempWord = nil
					} else {
						if b >= 65 && b <= 90 {
							b += 32
						}
						tempWord = append(tempWord, b)
					}
				}
				close(in)
				close(out)
			}(wg)

			in <- scanner.Bytes()

			for word := range out {
				isContain, index := contain(words, word)
				if !isContain {
					words = append(words, &Word{
						bytes:   word,
						counter: 1,
					})
				} else {
					words[index].counter += 1
				}
			}
		}
	}
	wg.Wait()
	sort.Slice(words, func(i, j int) bool {
		return words[i].counter > words[j].counter
	})
	for _, w := range words[:20] {
		fmt.Println(w.counter, string(w.bytes))
	}
	fmt.Printf("Process took %s\n", time.Since(start))
}
