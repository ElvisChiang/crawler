package main

import "fmt"

const debug = true

func main() {

	yearPrice := GetTaiwanBankGoldPriceYear()
	dayPrice := GetTaiwanBankGoldPriceDay(Date{2015, 7, 24})

	if debug {
		for date, price := range yearPrice {
			fmt.Printf("%d/%02d/%02d buy %04d sell %04d\n", date.year, date.month, date.day, price.buy, price.sell)
		}

		for time, price := range dayPrice {
			fmt.Printf("%02d:%02d buy %04d sell %04d\n", time.hour, time.minute, price.buy, price.sell)
		}
	}

	// TOOD: dump to json
}
