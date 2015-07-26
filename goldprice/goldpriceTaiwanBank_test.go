package goldprice

import (
	"fmt"
	"testing"
)

func TestGetTaiwanBankGoldPriceYear(t *testing.T) {
	dateArray, yearPrice := GetTaiwanBankGoldPriceYear()
	if len(yearPrice) == 0 {
		t.Errorf("cannot fetch data by year")
	}
	for _, date := range dateArray {
		price := yearPrice[date]
		fmt.Printf("%d/%02d/%02d buy %04d sell %04d\n", date.Year, date.Month, date.Day, price.buy, price.sell)
	}
}

func TestGetTaiwanBankGoldPriceDay(t *testing.T) {
	cases := []struct {
		in Date
	}{
		{Date{2015, 7, 24}},
	}
	for _, c := range cases {
		timeArray, dayPrice := GetTaiwanBankGoldPriceDay(c.in)
		if len(dayPrice) == 0 {
			t.Errorf("cannot fetch data by day")
		}
		for _, time := range timeArray {
			price := dayPrice[time]
			fmt.Printf("%02d:%02d buy %04d sell %04d\n", time.hour, time.minute, price.buy, price.sell)
		}
	}
}
