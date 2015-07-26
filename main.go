package main

import "./goldprice"

const debug = true

func main() {

	today := goldprice.Date{2015, 7, 24}
	dateArray, yearPrice := goldprice.GetTaiwanBankGoldPriceYear()
	timeArray, dayPrice := goldprice.GetTaiwanBankGoldPriceDay(today)

	goldprice.UpdateYear(dateArray, yearPrice)
	goldprice.UpdateToday(today, timeArray, dayPrice)
}
