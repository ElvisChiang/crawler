package main

import "github.com/ElvisChiang/crawler/goldprice"

const debug = true

func main() {

	today := goldprice.Date{2015, 7, 24}
	dateArray, yearPrice := goldprice.GetYearFromTaiwanBank()
	timeArray, dayPrice := goldprice.GetDayFromTaiwanBank(today)

	goldprice.UpdateYear(dateArray, yearPrice)
	goldprice.UpdateToday(today, timeArray, dayPrice)
}
