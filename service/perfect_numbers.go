package service

import (
	"runtime"
	"sort"
	"sync"
)

func IsPerfect(n int) bool {
	if n <= 1 {
		return false
	}

	sum := 1
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			sum += i
			if i != n/i {
				sum += n / i
			}
		}
	}

	return sum == n
}

func FindPerfectNumbers(start, end int) []int {
	var results []int
	for i := start; i <= end; i++ {
		if IsPerfect(i) {
			results = append(results, i)
		}
	}
	return results
}

func FindPerfectNumbersParallel(start, end int) []int {
	numWorkers := runtime.GOMAXPROCS(runtime.NumCPU())
	chunkSize := (end - start + 1) / numWorkers
	if chunkSize == 0 {
		chunkSize = 1
	}

	results := make(chan int, end-start+1)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		s := start + i*chunkSize
		e := s + chunkSize - 1
		if i == numWorkers-1 || e > end {
			e = end
		}

		wg.Add(1)
		go func(startChunk, endChunk int) {
			defer wg.Done()
			for n := startChunk; n <= endChunk; n++ {
				if IsPerfect(n) {
					results <- n
				}
			}
		}(s, e)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var perfects []int
	for n := range results {
		perfects = append(perfects, n)
	}
	sort.Ints(perfects)
	return perfects
}
