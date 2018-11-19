package main

import (
	"fmt"
	"testing"
)

func TestErrValues(t *testing.T) {

	tests := []struct {
		name        string
		stockPrices []int
		profit      int
		err         error
	}{
		{"EmptyData", []int{}, 0, ErrInsufficientData},
		{"InsufficientData", []int{1}, 0, ErrInsufficientData},
		{"NoProfit", []int{10, 1}, 0, ErrNoProfitMade},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s", tt.name), func(t *testing.T) {
			ticker := Ticker(tt.stockPrices)
			p, err := ticker.MaxProfit()
			if p != tt.profit {
				t.Errorf("got: %d, want: %d", p, tt.profit)
			}
			if err != tt.err {
				t.Errorf("got: %v, want: %v", err, tt.err)
			}
		})
	}

}

func TestMaxProfit(t *testing.T) {

	tests := []struct {
		name        string
		stockPrices []int
		profit      int
		err         error
	}{
		{"LinearRise", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9, nil},
		{"FallingWithSpike", []int{10, 9, 8, 7, 5, 6, 4, 3, 2, 1}, 1, nil},
		{"RisingWithSpike", []int{1, 2, 3, 4, 5, 11, 7, 8, 9, 10}, 10, nil},
		{"FlatWithSpike", []int{3, 3, 3, 1, 3, 3, 3, 11, 3, 3, 3}, 10, nil},
		{"MultipleSpikes", []int{1, 2, 1, 3, 1, 4, 1, 5, 1, 6, 1, 7, 1, 8, 1, 9, 1, 10}, 9, nil},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s", tt.name), func(t *testing.T) {
			ticker := Ticker(tt.stockPrices)
			p, err := ticker.MaxProfit()
			if err != tt.err {
				t.Errorf("got: %v, want: %v", err, tt.err)
			}
			if p != tt.profit {
				t.Errorf("got: %d, want: %d", p, tt.profit)
			}
		})
	}
}
