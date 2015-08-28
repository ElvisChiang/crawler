package buygamer

import (
	"fmt"
	"strconv"
	"testing"
)

func TestGetItemData(t *testing.T) {
	cases := []struct {
		in int
	}{
		{18719}, // item with 18+
		{18772},
	}

	fmt.Println("----------------")
	for _, c := range cases {
		data, ok := GetItemData(c.in)
		if ok == false {
			t.Errorf("cannot fetch data " + strconv.Itoa(c.in))
			continue
		}
		fmt.Printf("img url: %s\nitem name: %s\nplatform: %s\nvendor: %s\nversion: %s\ndate: %s\nduedate: %s\nprice: %s\n",
			data.url, data.itemName, data.platform, data.vendor, data.version, data.date, data.duedate, data.price)
		fmt.Println("----------------")
	}
}
