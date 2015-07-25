package main

const debug = true

func main() {

	today := Date{2015, 7, 24}
	dateArray, yearPrice := GetTaiwanBankGoldPriceYear()
	timeArray, dayPrice := GetTaiwanBankGoldPriceDay(today)

	UpdateYear(dateArray, yearPrice)
	UpdateToday(today, timeArray, dayPrice)
}
