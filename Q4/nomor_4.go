package main

import (
	"sync"
)

const N = 10

func main() {
	m := make(map[int]int)

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	wg.Add(N)
	// fmt.Println(wg)
	for i := 0; i < N; i++ {
		go func(wg *sync.WaitGroup, mu *sync.Mutex, val int) {
			defer wg.Done()
			mu.Lock()
			m[val] = val
			mu.Unlock()
		}(&wg, &mu, i)

	}
	wg.Wait()
	println(len(m))
}
