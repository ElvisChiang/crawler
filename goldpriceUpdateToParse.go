package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Parse definition
const (
	appID      = "" // input your application id here before running
	restAPIKey = "" // put your REST API key here before running
	version    = 1
	classDay   = "dayPrice"
	classYear  = "yearPrice"
)

// UpdateYear data to parse
func UpdateYear(dateArray []Date, yearPrice map[Date]Price) (ok bool) {
	tr := &http.Transport{
		// TLSClientConfig:    &tls.Config{RootCAs: false},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	urlYearPrice := fmt.Sprintf("https://api.parse.com/%d/classes/%s", version, classYear)

	// TOOD: delete all first
	for _, date := range dateArray {
		price := yearPrice[date]

		json := fmt.Sprintf(`{"date":%d%02d%02d, "buy":%d, "sell":%d}`,
			date.year, date.month, date.day,
			price.buy, price.sell)
		if debug {
			fmt.Println(json)
		}
		var jsonStr = []byte(json)
		req, err := http.NewRequest("POST", urlYearPrice, bytes.NewBuffer(jsonStr))
		if err != nil {
			fmt.Printf("%v\n", err)
			return false
		}
		req.Header.Add("X-Parse-Application-Id", appID)
		// req.Header.Add("X-Parse-REST-API-Key", restAPIKey)
		req.Header.Add("X-Parse-Master-Key", restAPIKey)
		req.Header.Add("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("%v\n", err)
			return false
		}
		// 4XX, bump error
		if resp.StatusCode/100 == 4 {
			body, _ := ioutil.ReadAll(resp.Body)
			if debug {
				fmt.Printf("code %d : %s\n", resp.StatusCode, body)
			}
			return false
		}
		defer resp.Body.Close()
	}
	return true
}

// UpdateToday data to parse
func UpdateToday(today Date, timeArray []Time, dayPrice map[Time]Price) (ok bool) {
	tr := &http.Transport{
		// TLSClientConfig:    &tls.Config{RootCAs: false},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	urlTodayPrice := fmt.Sprintf("https://api.parse.com/%d/classes/%s", version, classDay)

	// TOOD: delete all first
	for _, time := range timeArray {
		price := dayPrice[time]

		json := fmt.Sprintf(`{"date":%d%02d%02d, "hour":%d, "minute":%d, "buy":%d, "sell":%d}`,
			today.year, today.month, today.day,
			time.hour, time.minute, price.buy, price.sell)

		if debug {
			fmt.Println(json)
		}
		var jsonStr = []byte(json)

		req, err := http.NewRequest("POST", urlTodayPrice, bytes.NewBuffer(jsonStr))
		if err != nil {
			fmt.Printf("%v\n", err)
			return false
		}
		req.Header.Add("X-Parse-Application-Id", appID)
		// req.Header.Add("X-Parse-REST-API-Key", restAPIKey)
		req.Header.Add("X-Parse-Master-Key", restAPIKey)
		req.Header.Add("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("%v\n", err)
			return false
		}
		// 4XX, bump error
		if resp.StatusCode/100 == 4 {
			body, _ := ioutil.ReadAll(resp.Body)
			if debug {
				fmt.Printf("code %d : %s\n", resp.StatusCode, body)
			}
			return false
		}
		defer resp.Body.Close()
	}
	return true
}
