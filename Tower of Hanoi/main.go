//Tower of Hanoi Function

package main

import "fmt"


func towerofHanoi(num int, src byte, dst byte, temp byte) {
    if num < 1 {
        return
    }
    
    towerofHanoi(num-1, src, temp, dst)
    fmt.Printf("Move %d dis, from peg %c to peg %c \n", num, src, dst)
    towerofHanoi(num-1, temp, dst, src)
}

//Main function to run and test
func main() {
    num := 3
    fmt.Println("The sequence of moves used for the Tower of Hanoi are listed  below:")
    towerofHanoi(num, 'A', 'C', 'B')
}
