package main

import "fmt"

func main() {

	t := Ticker([]int{10, 9, 8, 7, 5, 6, 4, 3, 2, 1})

	profit, err := t.MaxProfit()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Profit: $%d\n", profit)
}
