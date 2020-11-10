package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"os"
	"bufio"


	"github.com/PuerkitoBio/goquery"
)

func fetch (url string) string {
    fmt.Println("Fetch Url", url)

    // 建立請求
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36`,)
	//User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36
	req.Header.Set("Cookie","webp=1; PHPSESSID=93cpqj399lqjithki5eoekvmd7; urlJumpIp=1; urlJumpIpByTxt=%E5%8F%B0%E5%8C%97%E5%B8%82; new_rent_list_kind_test=1; T591_TOKEN=93cpqj399lqjithki5eoekvmd7; c10f3143a018a0513ebe1e8d27b5391c=1; _ga=GA1.3.180757881.1604990332; _gid=GA1.3.1170198405.1604990332; _ga=GA1.4.180757881.1604990332; _gid=GA1.4.1170198405.1604990332; tw591__privacy_agree=0; XSRF-TOKEN=eyJpdiI6IlhVUG5aS3hrbmRXYmFWSFVpckpaVFE9PSIsInZhbHVlIjoiaTV4bWVvbWtTVnhYVVI0TGFGZ1ozeFcxaEUzblczYktMc21sVDQxbndBdlhNa1c1Z1RIRXE0c3M3YTVDaDdhOEZFdVwveXVZRFlhMFgwVlVyVFwvazMrdz09IiwibWFjIjoiNTcyODEwOWMzNGQwNWMxOGQ2MWRhZmQxZjBmNTYyYTlhZGI1YjIzNzgxYjA0NjZkMDI1ZjFmMGUyYTg5MWU2OCJ9; _fbp=fb.2.1604990334076.260835423; 591_new_session=eyJpdiI6Im9lRWI4UTFkbWVoNUFVbFBVZmRzb1E9PSIsInZhbHVlIjoiYlM2dERxbWs1K1VFemxnVDh1dlgwUHJrYm9zV2pZbXVkZ1pidTdFSzZBTkdia0hYT2hTVG9NSVBKcnNzbUJzYU10NTlmQm9PeFQwUWI5M201WWpKZXc9PSIsIm1hYyI6IjhiOWI3MDcxYmY3YWYyYzRmYjVlMDFhYWMxMmQ5M2I4OTAwMjZjYjI5MzliY2M3MWMwOTI0MjM1Yzg4ZGYzYzgifQ%3D%3D; _gat=1; _gat_UA-97423186-1=1")
	req.Header.Set("X-CSRF-TOKEN", "0tArIGrO6cAtCSyhtWsh1z9XyzWAw4znm3hBk5Ze")
	req.Header.Set("host", "rent.591.com.tw")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
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
        return ""
    }
    if resp.StatusCode != 200 {
        fmt.Println("Http status code:", resp.StatusCode)
        return ""
    }
    // 讀取HTTP響應正文
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Read error", err)
        return ""
    }
    return string(body)
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

	d1 := []byte(fetch("https://rent.591.com.tw/home/search/rsList?is_new_list=1&type=1&kind=0&searchtype=1&region=1&rentprice=0"))
	//d1 := []byte("hello\ngo\n")
    err := ioutil.WriteFile("./data1", d1, 0644)
    check(err)

    f, err := os.Create("./data2")
    check(err)

    defer f.Close()

    d2 := []byte{115, 111, 109, 101, 10}
    n2, err := f.Write(d2)
    check(err)
    fmt.Printf("wrote %d bytes\n", n2)

    n3, err := f.WriteString("writes\n")
    check(err)
    fmt.Printf("wrote %d bytes\n", n3)

    f.Sync()

    w := bufio.NewWriter(f)
    n4, err := w.WriteString("buffered\n")
    check(err)
    fmt.Printf("wrote %d bytes\n", n4)

    w.Flush()
}

	
func check(e error) {
    if e != nil {
        panic(e)
    }
}