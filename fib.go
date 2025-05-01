package fib

import (
  "fmt"
  "sync"
  "sort"
)


func fib(n int) int {
  // calcate the nth number in fiberna.. sequence

  if n == 1 {
    return 1
  }

  prevPrev := 0
  prev := 1
  val := 0

  for i := 1; i < n; i++ {
    val = prev + prevPrev
    prevPrev = prev
    prev = val
  }
  
  return val
}

type result struct {
   idx int
   val int
}

func fibSequential(n int) []int {
  results := make([]int, 0, n)
  for i := 1; i <= n; i++ {
    results = append(results, fib(i))
  }
  return results
}

func fibConcurrent(n int) []int {
  results := make(chan result)
  var wg sync.WaitGroup

  for i := 1; i <= n; i++ {
    wg.Add(1)
    go func(idx int, results chan<- result) {
      defer wg.Done()
      results <- result { idx: idx, val: fib(idx) }
    }(i, results)
  }

  go func(results chan result, wg *sync.WaitGroup) {
    wg.Wait()
    close(results)
  }(results, &wg)

  
  sortedResults := make([]int, 0, 31)
  for r := range results {
    sortedResults = append(sortedResults, r.val)
  }

  sort.Slice(sortedResults, func(i, j int) bool { 
    return sortedResults[i] < sortedResults[j]
  })
  return sortedResults
}
