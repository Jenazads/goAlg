package sortfunctions

import (
  "math/rand";
  "github.com/jenazads/goutils";
)

// return min between a, b
func EvaluateMin(a, b interface{}, comp goutils.TypeComparator) interface{} {
  if comp(a ,b) == 1 {
    return b
  } else {
    return a
  }
}

// Return min index
func MinIndex(arr []interface{}, comp goutils.TypeComparator, i, j int) int {
  if i == j {
    return i
  }
  k := MinIndex(arr, comp, i + 1, j)
  if comp(arr[i], arr[k]) == -1 {
    return i
  } else {
    return k
  }
}

// To generate permuatation of the array
func Shuffle(arr []interface{}, low, high int){
  for i:=low; i<high; i++ {
    rand_index := rand.Intn(high-low)+low
    arr[rand_index], arr[i] = arr[i], arr[rand_index]
  }
}

// verify is Sorted
func IsSorted(arr []interface{}, comp goutils.TypeComparator, low, high int) (bool){
  for i:=low; i<high-1; i++ {
    if comp(arr[i], arr[i+1]) == 1 {
      return false;
    }
  }
  return true;
}

// Reverses arr[low..high]
func Reverse(arr []interface{}, low, high int){
  for i:=low; i < high ; i++ {
    arr[high], arr[i] = arr[i], arr[high]
    high--;
  }
}

// Returns index of the maximum element in arr[0..size-1] 
func FindMaxMinElementIndex(arr []interface{}, comp goutils.TypeComparator, low, high int) (int, int){
  index_min, index_max := low, low
  for i := low+1; i < high; i++ {
    if comp(arr[i], arr[index_max]) == 1 {
      index_max = i
    }
    if comp(arr[i], arr[index_min]) == -1{
      index_min = i
    }
  }
  return index_max, index_min;
}

// return True if is power of 2
func IsPowerOfTwo(x int) (bool, uint){
  var expo uint = 0
  for ((x % 2) == 0) && (x > 1) {
    x /= 2
    expo++
  }
  return ((x != 0) && ((x & (^x + 1)) == x)), expo
}

// Power Two
func SetPowerOfTwo(x uint) int {
  return 2<<(x-1)
}