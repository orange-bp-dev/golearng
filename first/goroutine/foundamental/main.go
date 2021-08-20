package main

import (
	"fmt"
	"sync"
	"time"
)

func goroutine (s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(s)
	}
}

func normal (s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup
	go goroutine("world", &wg)
	normal("Hello")
  wg.Wait()
}