package main

import (
	"fmt"
	"testing"
)

func TestGetTaiwanBankGoldPriceYear(t *testing.T) {
	yearPrice := GetTaiwanBankGoldPriceYear()
	if len(yearPrice) == 0 {
		t.Errorf("cannot fetch data by year")
	}
	for date, price := range yearPrice {
		fmt.Printf("%d/%02d/%02d buy %04d sell %04d\n", date.year, date.month, date.day, price.buy, price.sell)
	}
}

func TestGetTaiwanBankGoldPriceDay(t *testing.T) {
	cases := []struct {
		in Date
	}{
		{Date{2015, 7, 24}},
	}
	for _, c := range cases {
		dayPrice := GetTaiwanBankGoldPriceDay(c.in)
		if len(dayPrice) == 0 {
			t.Errorf("cannot fetch data by day")
		}
		for time, price := range dayPrice {
			fmt.Printf("%02d:%02d buy %04d sell %04d\n", time.hour, time.minute, price.buy, price.sell)
		}
	}
}
