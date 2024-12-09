package main

import "fmt"

func main() {
  array := make([]int, 101)
  for i := range array {
    array[i] = i
  }
  //Brute force example with for loop for FizzBuzz;
  for i := 0; i < len(array); i++ {
    if array[i]%3 == 0 && array[i]%5 == 0 {
      fmt.Println("FizzBuzz!")
    }
    if array[i]%5 == 0 {
      fmt.Println("Buzz!")
    }
    if array[i]%3 == 0 {
      fmt.Println("Fizz!")
    }
    fmt.Println(array[i])
  }

  //Recursive example for FizzBuzz;

  recurse(array, 0)
}

func recurse(arr []int, n int) {
  if n == len(arr) {
    fmt.Println("Finished")
    return
  }

  if n%3 == 0 && n%5 == 0 {
    fmt.Println("FizzBuzz!")
  } else if n%5 == 0 {
    fmt.Println("Buzz!")
  } else if n%3 == 0 {
    fmt.Println("Fizz!")
  } else {
    fmt.Println(arr[n])
  }

  recurse(arr, n+1)
}
