package main

import "github.com/ElvisChiang/crawler/goldprice"

const debug = true

func main() {

	today := goldprice.Date{Year: 2016, Month: 4, Day: 18}
	dateArray, yearPrice := goldprice.GetYearFromTaiwanBank()
	timeArray, dayPrice := goldprice.GetDayFromTaiwanBank(today)

	goldprice.UpdateYear(dateArray, yearPrice)
	goldprice.UpdateToday(today, timeArray, dayPrice)
}
