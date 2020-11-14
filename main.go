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

var newsession591 string
var xcsrf string

func fetch (url string) []byte {
    fmt.Println("Fetch Url", url)

    // 建立請求
	req, _ := http.NewRequest("GET", url, nil)
	// must set Header 591_new_session and X-CSRF-TOKEN or it will return status 419
	req.Header.Set("Cookie","591_new_session=eyJpdiI6ImpyRENDYzh0Y3hYdDlYNzlmMEFORHc9PSIsInZhbHVlIjoiN1wvck1jVGdUblZvR1VwQVUyTGlJdW1HektwcHNwc1BqYkhHTHJDSVk4QnZyeUpxOHVLNHJwN3UwbUZDTmh5VUg0dCtyM3VBVW9YbDh4M3JrblRwSEhnPT0iLCJtYWMiOiI0NDMwMjdmYWJhODlmNmEwNmQ3ZjlhMGYyYzcyZDBjMzY2NTkyNTgyY2YyZGY0NGVjMzY3ZmJlNTU1ZTlhZTZlIn0%3D")
	/*cookie := &http.Cookie{
		Name:  "591_new_session",
		Value: newsession591,
	}
	req.AddCookie(cookie)*/
	//req.Header.Set("X-CSRF-TOKEN", xcsrf)
	req.Header.Set("X-CSRF-TOKEN", "kAI3Ye9RMKcJkkbiI3RpsmaJ5W9bUN1i5Kn3AeaB")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.183 Safari/537.36")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://rent.591.com.tw/?kind=0&region=1")

	// 建立HTTP客戶端
	client := &http.Client{}

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
	fmt.Println("********")

	//fmt.Println(string(body))
	fmt.Println("********")
    return body
}

// Export with struct HouseInfo
func houselistJSON(b []byte) {
	var tmpHouseList HouseList
	err := json.Unmarshal(b, &tmpHouseList)
	if err != nil {
		fmt.Println("error:", err)
	}

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

func getHeaders(url string) {
	// ** get X-CSRF-TOKEN
	  // Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	/*if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}*/
	
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	// Find the review items
	  doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		if name, _ := s.Attr("name"); name == "csrf-token" {
			xcsrf, _ = s.Attr("content")
			fmt.Printf("XCSRF field: %s\n", xcsrf)
		}
	})
	// **

	// ** get 591_new_session
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	
	fmt.Println("cookies")
	for _, v := range resp.Cookies() {
		if v.Name == "591_new_session" {
			newsession591 = v.Value
			fmt.Printf("%+v\n", v)
		}
	}
	fmt.Println("set - cookies")
	//fmt.Println(resp.Header.readSetCookies("Set-Cookie"))
}

func Scrape(url string) {
	// Request the HTML page.
	res, err := http.Get(url)
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

	fmt.Println("find")
	// Find the review items
	doc.Find("#content").Each(func(i int, s *goquery.Selection) {
	  // For each item found, get the band and title
	  s.Find(".listInfo.clearfix").Each(func(li int, listInfo *goquery.Selection) {
		// For each item found, get the band and title
		  title := listInfo.Find(".pull-left.infoContent > h3 > a").Text()
		  fmt.Println(title)
	  })
	})
	//fmt.Println(doc.Find(".content"))
}

func init() {
	//var url = "https://rent.591.com.tw/?kind=0&region=1"
	//getHeaders(url)
}

func main() {
	var url = "https://rent.591.com.tw/?kind=0&region=1"
	Scrape(url)
	// this will return status 419
	//fmt.Println(fetch("https://rent.591.com.tw/home/search/rsList?is_new_list=1&type=1&kind=0&searchtype=1&region=1&rentprice=0"))
	// this one work for raw html page
	//fmt.Println(fetch("https://rent.591.com.tw/?kind=0&region=1&rentprice=3"))

	//bs := fetch(url)
	//ExportJSON(bs)
}
	
func check(e error) {
    if e != nil {
        panic(e)
    }
}