package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"os"
	//"bufio"
	"encoding/json"


	"github.com/PuerkitoBio/goquery"
)

func fetch (url string) []byte {
    fmt.Println("Fetch Url", url)

    // 建立請求
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36`,)
	//User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36
	// must set Header 591_new_session and X-CSRF-TOKEN or it will return status 419
	req.Header.Set("Cookie","591_new_session=eyJpdiI6ImJTRThla3FRUXNcL2ljRUdIRXY1YVpnPT0iLCJ2YWx1ZSI6IjkzVUFHbnJVMmM1Y2RpY2VEQmZhQ1RUQ1g0bjljSXh2Q2ZuenYraWczVGJWOGFWNk16SHFtcGdjQVBPVGl0K3ZcL29OT1lQNUQ4SlFqRlNqTU9cL05Pb1E9PSIsIm1hYyI6IjQxMGQwYzRkMTc1N2Q5OWI3N2NlNmQ5NzViMzkwMTNmMDc1MmUwNDY3NTA5MjhiOGM4ODQxY2NkODI3NWM0NjkifQ%3D%3D")
	req.Header.Set("X-CSRF-TOKEN", "tNgLhP2SCl3BNMmXduikBVg1svMo5KhqfgO3HbaP")
	//req.Header.Set("host", "rent.591.com.tw")
	//req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//fmt.Println(req.Header)
	// 建立HTTP客戶端
	client := &http.Client{}
	/*cookie := &http.Cookie{
		Name:  "urlJumpIp",
		Value: "1",
	}
	req.AddCookie(cookie)*/
    // 發出請求
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Http get err:", err)
        return []byte{}
    }
    if resp.StatusCode != 200 {
        fmt.Println("Http status code:", resp.StatusCode)
        return []byte{}
    }
    // 讀取HTTP響應正文
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Read error", err)
        return []byte{}
	}
	houselistJSON(body)
    return body
}

// Export with struct HouseInfo
func houselistJSON(b []byte) {
	var tmpHouseList HouseList
	err := json.Unmarshal(b, &tmpHouseList)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", tmpHouseList)
	fmt.Println(tmpHouseList)
	file, _ := json.MarshalIndent(tmpHouseList, "", " ")

	_ = ioutil.WriteFile("houselist.json", file, 0644)
}

// ExportJSON export json file.
func ExportJSON(b []byte) {
	f, err := os.Create("./rent.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	f.Write(b)
	log.Println("✔️  Export Path: \x1b[42m/rent.json\x1b[0m")
}

func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("https://rent.591.com.tw/home/search/rsList?is_new_list=1&type=1&kind=0&searchtype=1&region=1&rentprice=0")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	/*
	// Find the review items
	doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("a").Text()
		title := s.Find("i").Text()
		fmt.Printf("Review %d: %s - %s\n", i, band, title)
	})*/
	fmt.Print(doc)
}

func main() {
	//ExampleScrape()

	// this will return status 419
	fmt.Println(fetch("https://rent.591.com.tw/home/search/rsList?is_new_list=1&type=1&kind=0&searchtype=1&region=1&rentprice=0"))
	// this one work for raw html page
	//fmt.Println(fetch("https://rent.591.com.tw/?kind=0&region=1&rentprice=3"))

	d1 := fetch("https://rent.591.com.tw/home/search/rsList?is_new_list=1&type=1&kind=0&searchtype=1&region=1&rentprice=0")
	ExportJSON(d1)
}

	
func check(e error) {
    if e != nil {
        panic(e)
    }
}