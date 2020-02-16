//Using a Given list where the nth element is the price of stock on nth day an algorithm is used to find 
//what data you will buy and what date you will sell to get max profit
 package main
import "fmt"

func maxProfit(stocks []int) int{
    size := len(stocks)
    buy := 0
    sell := 0
    curMin := 0
    currProfit := 0
    maxProfit := 0
    
    for i := 0; i < size; i++ {
        if stocks[i] < stocks[curMin] {
            curMin = i
        }
        
        currProfit = stocks[i] - stocks[curMin]
        if currProfit > maxProfit {
            buy = curMin
            sell = i
            maxProfit = currProfit
        }
    }
    fmt.Println("Purchase day is ", buy, " at price ", stocks[buy])
    fmt.Println("Sell day is ", sell, " at price ", stocks[sell])
    fmt.Println("Max Profit :: ", maxProfit)
    return maxProfit
}

func main(){
    first := []int {10, 150, 6, 67, 61, 16, 86, 6, 67, 78, 150, 3, 28, 143 }
    maxProfit(first)
}
