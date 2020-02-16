//Binary Search Recursion
package main

import "fmt"

func BinarySearchRecursive(data []int, value int) bool {
    size := len(data)
    return BinarySearchRecursiveUtil(data, 0, size - 1, value)
}

func BinarySearchRecursiveUtil(data []int, low int, high int, value int) bool {
    if low > high {
        return false
    }
    mid := (low + high)/2
    if data[mid] == value {
        return true
    } else if data [mid] < value {
        return BinarySearchRecursiveUtil(data, mid+1, high, value)
    } else {
        return BinarySearchRecursiveUtil(data, low, mid-1, value)
    }
}

func main() {
    //sample array for search
    arr := []int{1, 2, 3, 4, 5, 6,7, 9,}
    fmt. Println("BinarySearchRecursive:", BinarySearchRecursive(arr, 7)) //true
    fmt. Println("BinarySearchRecursive:", BinarySearchRecursive(arr, 8)) //false
}
