package main

import "fmt"

const debug = true

func main() {

	yearPrice := GetTaiwanBankGoldPriceYear()

	for date, price := range yearPrice {
		if debug {
			fmt.Printf("%d/%02d/%02d buy %04d sell %04d\n", date.year, date.month, date.day, price.buy, price.sell)
		}
	}
}
