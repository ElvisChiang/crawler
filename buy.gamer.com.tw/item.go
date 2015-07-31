package buygamer

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var cookie = http.Cookie{
	Name:   "ckBUY_item18UP",
	Value:  "18UP",
	Path:   "/",
	Domain: "buy.gamer.com.tw"}

const urlPrefix = "http://buy.gamer.com.tw/atmItem.php?sn="

// ItemData baha item data
type ItemData struct {
	url      string
	platform string
	vendor   string
	version  string
	date     string
	duedate  string
	price    string
}

// GetItemData get item from baha
func GetItemData(id int) (data *ItemData, ok bool) {
	url := urlPrefix + strconv.Itoa(id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, false
	}

	req.AddCookie(&cookie)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, false
	}

	body, _ := ioutil.ReadAll(resp.Body)

	html := strings.Split(string(body), "\n")

	// since incorrect serial number also has its page, have to parse it anyway
	data, ok = parseContent(html)
	return
}

func parseContent(html []string) (data *ItemData, ok bool) {
	var itemString string
	for _, line := range html {
		has := strings.Contains(line, `<p class="ES-lbox4A"`)
		if !has {
			continue
		}
		itemString = line
		break
	}
	r := regexp.MustCompile(`<img src=\"(.+)\" border="0">` +
		`.+<span>品　　名</span><strong>(.+)</strong>` +
		`.+<span>遊戲平台</span><i>(.+?)</i>` +
		`.+<span>製作發行</span>(.+)</li>` +
		`.+<span>版本資訊</span>(.+)</li>` +
		`.+<span>發售日期</span>(.+)</li>` +
		`.+<span>訂單修改期限</span><i>(.+)</i>` +
		`.+<span>本站售價</span><i>(.+)</i>`)
	res := r.FindStringSubmatch(itemString)
	if res == nil {
		fmt.Println("unmatch")
		return nil, false
	}
	data = &ItemData{res[1], res[2], res[3], res[4], res[5], res[6], res[7]}
	return data, true
}
