package main

import (
	"bytes"
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var pings = -1
var txt = ""

func ifNewPin() {
	res, err := http.Get(UrlZhihu)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatal("status code error: %d %s\n",res.StatusCode,res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	//to get the ping num
	doc.Find("ul[class=\"Tabs ProfileMain-tabs\"]").Eq(0).Find("li[aria-controls=\"Profile-pins\"]").Each(func(i int, selection *goquery.Selection) {
		a := selection.Find("a")
		number,_ := a.Attr("meta")
		number = strings.Replace(number,",","",-1)
		num ,err:= strconv.Atoi(number)
		if err != nil {
			log.Fatal(err)
		}
		if num != pings {
			log.Println("************there is new pin****************")
			doc.Find("div[class=\"ContentItem PinItem\"]").Eq(0).
				Find("div[class=\"RichContent-inner\"]").Find("span").Each(func(i int, selection *goquery.Selection) {
				if i == 0 {
					txt = selection.Text()
					log.Println(txt)
					values := map[string]string{"txt":txt}
					jsonValue,_ := json.Marshal(values)
					_, _ = http.Post("http://localhost:8080/newping", "application/json", bytes.NewReader(jsonValue))
				}

			})
			_ = SendMail()
			pings = num
		}else {

		}
	})

}

func main() {
	rand.Seed(time.Now().UnixNano())
	for{
		ifNewPin()
		offset := rand.Intn(20) - 10
		time.Sleep(time.Duration(30+offset)*time.Second)
	}

}