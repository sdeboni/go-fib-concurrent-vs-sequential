package main

import "testing"


func BenchmarkConcurrent(b *testing.B) {
  for n := 0; n < b.N; n++ {
    fibConcurrent(30)
  }
}

func BenchmarkSequential(b *testing.B) {
  for n := 0; n < b.N; n++ {
    fibSequential(30)
  }
}
