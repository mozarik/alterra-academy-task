package main

import (
	"fmt"
	"strings"
	"sync"
)

type SafeMap struct {
	mu     sync.Mutex
	result map[string]int
}

func (s *SafeMap) letterFrequency(char string, size chan int, group *sync.WaitGroup) {
	s.mu.Lock()
	s.result[char]++
	size <- s.result[char]
	fmt.Println(char, ": ", <-size)
	group.Done()
	s.mu.Unlock()
}

func (s *SafeMap) letterValue() map[string]int {
	return s.result
}

func main() {
	var wg sync.WaitGroup
	s := SafeMap{result: make(map[string]int)}
	str := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"
	strReplace := strings.ReplaceAll(str, " ", "")
	size := make(chan int, len(strReplace))
	wg.Add(len(strReplace))
	for _, v := range strReplace {
		go s.letterFrequency(string(v), size, &wg)
	}
	wg.Wait()
	fmt.Println("===== RESULT =====")
	fmt.Println(s.letterValue())
}
