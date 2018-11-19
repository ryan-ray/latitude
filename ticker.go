package main

import "errors"

type Ticker []int

var ErrInsufficientData = errors.New("Need at least 2 stock prices to calculate profit")
var ErrNoProfitMade = errors.New("No profit can be made")

func (t Ticker) MaxProfit() (int, error) {
	if len(t) <= 1 {
		return 0, ErrInsufficientData
	}

	profit := t[1] - t[0]
	min := t[0]
	for _, p := range t[1:] {
		if p-min > profit {
			profit = p - min
		}

		if p < min {
			min = p
		}
	}

	if profit <= 0 {
		return 0, ErrNoProfitMade
	}

	return profit, nil
}
