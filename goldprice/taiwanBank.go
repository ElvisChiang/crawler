package goldprice

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const debug = true

const urlBankYear = "http://rate.bot.com.tw/Pages/UIP005/UIP005INQ3.aspx?view=1&amp;lang=zh-TW"

const viewStatKey = "__VIEWSTATE"
const viewStat = "/wEPDwUKLTc3OTg4NTEyM2QYAgUeX19Db250cm9sc1JlcXVpcmVQb3N0QmFja0tleV9fFgoFBlJhZGlvNQUGUmFkaW8xBQZSYWRpbzIFBlJhZGlvMwUGUmFkaW80BQVjdGwwMgUFY3RsMDMFBWN0bDA0BQVjdGwwNQUFY3RsMDYFCW11bHRpVGFicw8PZAIBZPZJVCxGUh1sKGSIy7aqPTTqBGsH"
const validationKey = "__EVENTVALIDATION"
const validation = "/wEWEQKouumRBAKOkoCCCwKMhc76CQLWssLjAgKf5L6bDQLL3LrjCgLW3JbjCgLS3JbjCgLR3JbjCgLM3JbjCgLl18nSCwLm1/nSCwL9/oqMCQKLts3CAQKUts3CAQKM54rGBgLWlM+bAvA8FJroJ9FZZI52UscKHadSwtMt"

// form input parameter for date
const (
	dateParam   = "term"
	recentDay   = 99
	threeMonth  = 6
	halfYear    = 2
	year        = 3
	specifyDate = 0
	monthParam  = "month"
	yearParam   = "year"
)

// form input parameter for current type
const (
	currentParam = "curcd"
	currentTWN   = "TWD"
	currentUSD   = "USD"
	currentCNY   = "CNY"
)

// form input parameter for when in a day
const (
	whenParam = "afterOrNot"
	before    = 0
	after     = 1
)

// Date 1234/5/6
type Date struct {
	Year, Month, Day int
}

// Time 23:59
type Time struct {
	Hour, Minute int
}

// Price of buy and sell
type Price struct {
	Buy, Sell int
}

// GetYearFromTaiwanBank get whole year gold price from taiwan bank
func GetYearFromTaiwanBank() (dateArray []Date, ret map[Date]Price) {
	var term int //
	var y int
	var m int
	var curcd string // should be TWD
	var when int
	ret = make(map[Date]Price)
	dateArray = make([]Date, 0)

	now := time.Now()
	term = year
	y = now.Year()
	m = int(now.Month())
	curcd = currentTWN
	when = before

	resp, err := http.PostForm(urlBankYear,
		url.Values{
			"Button1":     {"查詢"}, // I don't know what the fuck is this
			viewStatKey:   {viewStat},
			validationKey: {validation},
			dateParam:     {strconv.Itoa(term)},
			yearParam:     {strconv.Itoa(y)},
			monthParam:    {strconv.Itoa(m)},
			currentParam:  {curcd},
			whenParam:     {strconv.Itoa(when)}})
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	html := strings.Split(string(body), "\n")

	for _, line := range html {
		// if strings.Contains(line, "class=\"color0\"") || strings.Contains(line, "class=\"color1\"") {
		r := regexp.MustCompile(`date=(\d{8}).+class="decimal">(\d+)</td><td class="decimal">(\d+)`)
		res := r.FindStringSubmatch(line)
		if res == nil {
			continue
		}
		tmpDate, _ := strconv.Atoi(res[1])
		date := Date{tmpDate / 10000, (tmpDate % 10000) / 100, tmpDate % 100}
		buy, _ := strconv.Atoi(res[2])
		sell, _ := strconv.Atoi(res[3])
		price := Price{buy, sell}
		ret[date] = price
		dateArray = append(dateArray, date)
	}

	return
	// TODO return map of a year
}

const urlBankDay = "http://rate.bot.com.tw/Pages/UIP005/UIP00511.aspx"

// GetDayFromTaiwanBank get specifiy date gold price from taiwan bank
func GetDayFromTaiwanBank(date Date) (timeArray []Time, ret map[Time]Price) {
	ret = make(map[Time]Price)
	timeArray = make([]Time, 0)
	dateString := fmt.Sprintf("%d%02d%02d", date.Year, date.Month, date.Day)

	url := urlBankDay +
		"?" +
		"&lang=zh-TW" +
		"&whom=GB0030001000" + // don't know what
		"&date=" + dateString +
		"&afterOrNot=" + strconv.Itoa(before) +
		"&curcd=" + currentTWN

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	html := strings.Split(string(body), "\n")
	for _, line := range html {
		r := regexp.MustCompile(`">(\d{2}):(\d{2})<.+class="decimal">(\d+)</td><td class="decimal">(\d+)`)
		res := r.FindStringSubmatch(line)
		if res == nil {
			continue
		}
		hour, _ := strconv.Atoi(res[1])
		min, _ := strconv.Atoi(res[2])
		buy, _ := strconv.Atoi(res[3])
		sell, _ := strconv.Atoi(res[4])
		eventTime := Time{hour, min}
		price := Price{buy, sell}
		ret[eventTime] = price
		timeArray = append(timeArray, eventTime)
	}

	return
}
