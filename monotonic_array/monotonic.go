package main

import "fmt"

func monotonic(arr []int){
  if (non_decreasing(arr) == true || non_increasing(arr) == true) {
    fmt.Println("Monotonic")
  } else {
    fmt.Println("Not Monotonic")
  }
}

func non_decreasing(arr []int) bool {
  for i := 1; i<len(arr); i++ {
    if arr[i] < arr[i-1] {
      return false
    }
  }
  return true
}

func non_increasing(arr []int) bool {
  for i := 1; i<len(arr); i++ {
    if arr[i] > arr[i-1] {
      return false
    }
  }
  return true
}

func main(){
  var decreasing_array_not = []int {1000, 2, 3, 17, 50}
  var increasing_array_not = []int {2, 4, 5, 7, 3, 6}
  var  array = []int {23, 34, 34, 56}
  monotonic(decreasing_array_not)
  monotonic(increasing_array_not)
  monotonic(array)
}
