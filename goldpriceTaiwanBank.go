package main

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

const urlBank = "http://rate.bot.com.tw/Pages/UIP005/UIP005INQ3.aspx?view=1&amp;lang=zh-TW"
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
	year, month, day int
}

// Price of buy and sell
type Price struct {
	buy, sell int
}

// GetTaiwanBankGoldPriceYear get whole year gold price from taiwan bank
func GetTaiwanBankGoldPriceYear() (ret map[Date]Price) {
	var term int //
	var y int
	var m int
	var curcd string // should be TWD
	var when int
	ret = make(map[Date]Price)

	now := time.Now()
	term = year
	y = now.Year()
	m = int(now.Month())
	curcd = currentTWN
	when = before

	resp, err := http.PostForm(urlBank,
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
	}

	return ret
	// TODO return map of a year
}
