package main

import (
	"fmt"
	"sync"
)

func main() {
	var balance int
	var wg sync.WaitGroup
	var mu sync.Mutex

	deposit := func(amount int) {
		mu.Lock()
		defer mu.Unlock()
		balance += amount
	}

	withdrawal := func(amount int) {
		mu.Lock()
		defer mu.Unlock()
		balance -= amount
	}

	wg.Add(100)
	for i := 1; i < 101; i++ {
		go func() {
			defer wg.Done()
			deposit(1)
		}()

	}

	wg.Add(100)
	for i := 1; i < 101; i++ {
		go func() {
			defer wg.Done()
			withdrawal(1)
		}()
	}

	wg.Wait()
	fmt.Println(balance)

}
