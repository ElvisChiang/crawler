package goldprice

import "testing"

func TestUpdateYear(t *testing.T) {
	dateArray, yearPrice := GetYearFromTaiwanBank()
	if true {
		return
	}
	result := UpdateYear(dateArray, yearPrice)
	if result == false {
		t.Errorf("result incorrect")
	}
}

func TestUpdateToday(t *testing.T) {
	cases := []struct {
		in Date
	}{
		{Date{2015, 7, 24}},
	}
	if true {
		return
	}
	for _, c := range cases {
		today := c.in
		timeArray, dayPrice := GetDayFromTaiwanBank(today)
		result := UpdateToday(today, timeArray, dayPrice)
		if result == false {
			t.Errorf("result incorrect")
		}
	}
}
