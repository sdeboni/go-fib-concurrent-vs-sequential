package fib

import (
  "testing"
  "reflect"
)


func TestFib1(t *testing.T) {
  actual := fib(1)
  expected := 1
  if expected != actual {
    t.Fatalf("expected: %d, got: %d", expected, actual)
  } 
}

func TestFib2(t *testing.T) {
  actual := fib(2)
  expected := 1
  if expected != actual {
    t.Fatalf("expected: %d, got: %d", expected, actual)
  } 
}
func TestFib3(t *testing.T) {
  actual := fib(3)
  expected := 2
  if expected != actual {
    t.Fatalf("expected: %d, got: %d", expected, actual)
  } 
}
func TestFib4(t *testing.T) {
  actual := fib(4)
  expected := 3
  if expected != actual {
    t.Fatalf("expected: %d, got: %d", expected, actual)
  } 
}
func TestFib6(t *testing.T) {
  actual := fib(6)
  expected := 8
  if expected != actual {
    t.Fatalf("expected: %d, got: %d", expected, actual)
  } 
}
func TestSequential(t *testing.T) {
  actual := fibSequential(30)
  expected := []int{1,1,2,3,5,8,13,21,34,55,89,144,233,377,610,987,1597,2584,4181,6765,10946,17711,28657,46368,75025,121393,196418,317811,514229,832040}
  if !reflect.DeepEqual(expected, actual) {
    t.Fatalf("got     :\n%v\nexpected:\n%v\n", actual, expected)
  }
} 
func TestConcurrent(t *testing.T) {
  actual := fibConcurrent(30)
  expected := []int{1,1,2,3,5,8,13,21,34,55,89,144,233,377,610,987,1597,2584,4181,6765,10946,17711,28657,46368,75025,121393,196418,317811,514229,832040}
  if !reflect.DeepEqual(expected, actual) {
    t.Fatalf("got     :\n%v\nexpected:\n%v\n", actual, expected)
  }
} 

