package crawler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/google/go-querystring/query"
)

const rootURL = "https://rent.591.com.tw/?"

type Document struct {
	doc *goquery.Document
}

type Crawler struct {
	URL string
	//items     int
	Houselist []*HouseInfo
	Options   *Options
	wg        sync.WaitGroup
}

type Options struct {
	Region      int    `url:"region"`                // 地區 - 預設：`1`
	Section     string `url:"section,omitempty"`     // 鄉鎮 - 可選擇多個區域，例如：`section=7,4`
	Kind        int    `url:"kind"`                  // 租屋類型 - `0`：不限、`1`：整層住家、`2`：獨立套房、`3`：分租套房、`4`：雅房、`8`：車位，`24`：其他
	RentPrice   string `url:"rentprice,omitempty"`   // 租金 - `0`：no limit、`1`：0k - 5k、`2`：5k - 10k、`3`：10k - 20k、`4`: 20k - 30k；自定義範圍如`5000,15000`：5k - 15k
	Area        string `url:"area,omitempty"`        // 坪數格式 - `10,20`（10 到 20 坪）
	Order       string `url:"order"`                 // 貼文時間 - 預設使用刊登時間：`posttime`，或是使用價格排序：`money`
	OrderType   string `url:"orderType"`             // 排序方式 - `desc` 或 `asc`
	Sex         int    `url:"sex,omitempty"`         // 性別 - `0`：不限、`1`：男性、`2`：女性
	HasImg      string `url:"hasimg,omitempty"`      // 過濾是否有「房屋照片」 - ``：空值（不限）、`1`：是
	NotCover    string `url:"not_cover,omitempty"`   // 過濾是否為「頂樓加蓋」 - ``：空值（不限）、`1`：是
	Role        string `url:"role,omitempty"`        // 過濾是否為「屋主刊登」 - ``：空值（不限）、`1`：是
	Shape       string `url:"shape,omitempty"`       // 房屋類型 - `1`：公寓、`2`：電梯大樓、`3`：透天厝、`4`：別墅
	Pattern     string `url:"pattern,omitempty"`     // 格局單選 - `0`：不限、`1`：一房、`2``：兩房、`3`：三房、`4`：四房、`5`：五房以上
	PatternMore string `url:"patternMore,omitempty"` // 格局多選 - 參考「格局單選」，可以選多種格局，例如：`1,2,3,4,5`
	Floor       string `url:"floor,omitempty"`       // 樓層 - `0,0`：不限、`0,1`：一樓、`2,6`：二樓到六樓、`6,12`：六樓到十二樓、`12,`：十二樓以上
	Option      string `url:"option,omitempty"`      // 提供設備 - `tv`：電視、`cold`：冷氣、`icebox`：冰箱、`hotwater`：熱水器、`naturalgas`：天然瓦斯、`four`：第四台、`broadband`：網路、`washer`：洗衣機、`bed`：床、`wardrobe`：衣櫃、`sofa`：沙發。可選擇多個設備，例如：option=tv,cold
	Other       string `url:"other,omitempty"`       // 其他條件 - `cartplace`：有車位、`lift`：有電梯、`balcony_1`：有陽台、`cook`：可開伙、`pet`：可養寵物、`tragoods`：近捷運、`lease`：可短期租賃。可選擇多個條件，例如：other=cartplace,cook
	FirstRow    int    `url:"firstRow"`
}

type HouseInfo struct {
	ID          int    `json:"id"`
	ImgSrc      string `json:"imgsrc"`
	Title       string `json:"title"`
	URL         string `json:"url"`
	Kind        string `json:"kind"`
	Layout      string `json:"layout"`
	Ping        string `json:"ping"`
	Floor       string `json:"floor"`
	Address     string `json:"address"`
	Price       string `json:"price"`
	NewItem     bool   `json:"newItem"`
	CrawlerTime string `json:"crawlertime"`
	UpdateTime  string `json:"updatetime"`
	Author      string `json:"author"`
}

func Newcrawler() *Crawler {
	return &Crawler{
		URL: rootURL,
		//items:     50,
		Houselist: make([]*HouseInfo, 0),
		Options: &Options{
			Kind:      0,
			Region:    1,
			Order:     "posttime",
			OrderType: "desc",
		},
	}
}

func NewHouseInfo() *HouseInfo {
	return &HouseInfo{}
}
func (c *Crawler) NewURL() string {
	v, err := query.Values(c.Options)
	if err != nil {
		log.Fatal(err)
	}
	return rootURL + v.Encode()
}

func NewDoc(url string) *goquery.Document {
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
	return doc
}

func (c *Crawler) Scrape(page int) {
	defer c.wg.Done()
	c.URL = c.NewURL()
	// Request the HTML page.
	u, err := url.Parse(c.URL)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Set("firstRow", strconv.Itoa(page*30))
	fmt.Println("Crawlering ------> " + rootURL + q.Encode())
	res, err := http.Get(rootURL + q.Encode())
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

	var thisid = page * 30
	// Find the content section
	doc.Find("#content").Each(func(i int, s *goquery.Selection) {
		// For each listInfo section, and loop to get the detail
		s.Find(".listInfo.clearfix").Each(func(li int, listInfo *goquery.Selection) {
			// For each item found, get the band and title
			newHouseInfo := NewHouseInfo()
			newHouseInfo.CrawlerTime = time.Now().Format("2006-01-02 3:4:5 pm")
			if img, ok := listInfo.Find(".pull-left.imageBox > img").Attr("data-original"); ok {
				newHouseInfo.ImgSrc = strings.Replace(img, "210x158.crop.jpg", "765x517.water3.jpg", 1)
				//fmt.Println(img)
			}

			newHouseInfo.ID = thisid
			newHouseInfo.Title = listInfo.Find(".pull-left.infoContent > h3 > a").Text()
			//fmt.Println(title)

			if link, ok := listInfo.Find(".pull-left.infoContent > h3 > a").Attr("href"); ok {
				newHouseInfo.URL = "https:" + link
			}

			housetype := listInfo.Find(".pull-left.infoContent > .lightBox").First().Text()
			typearray := strings.Split(stringReplacer(housetype), "|")
			// some time it will contain Room layout in index 1
			if len(typearray) == 4 {
				newHouseInfo.Kind = typearray[0]
				newHouseInfo.Layout = typearray[1]
				newHouseInfo.Ping = typearray[2]
				newHouseInfo.Floor = typearray[3]
			} else {
				newHouseInfo.Kind = typearray[0]
				newHouseInfo.Layout = ""
				newHouseInfo.Ping = typearray[1]
				newHouseInfo.Floor = typearray[2]
			}

			newHouseInfo.Address = listInfo.Find(".pull-left.infoContent > .lightBox > em").Text()
			//fmt.Println(address)

			publishInfo := listInfo.Find(".pull-left.infoContent > .lightBox").Eq(1).Next().Text()
			pInfoarray := strings.Split(stringReplacer(publishInfo), "/")
			newHouseInfo.Author = pInfoarray[0]
			newHouseInfo.UpdateTime = pInfoarray[1]

			newHouseInfo.Price = stringReplacer(listInfo.Find(".price > i").Text())
			newHouseInfo.Price = strings.Replace(newHouseInfo.Price, ",", "", -1)

			newHouseInfo.NewItem = false
			listInfo.Find(".newArticle").Each(func(_ int, n *goquery.Selection) {
				newHouseInfo.NewItem = true
			})
			c.Houselist[thisid] = newHouseInfo
			thisid++
		})
	})
}

func (c *Crawler) Start(page int) {
	totalitems, totalpages := c.ItemandPageNum()
	fmt.Println(totalitems)
	if page == -1 {
		page = totalpages
	}
	if totalpages < page {
		fmt.Println("There are only " + strconv.Itoa(totalpages) + " pages!")
		page = totalpages
	}
	c.Houselist = make([]*HouseInfo, page*30)
	for i := 0; i < page; i++ {
		c.wg.Add(1)
		go c.Scrape(i)
	}
	c.wg.Wait()
	if page == totalpages {
		fmt.Println("Total items: " + strconv.Itoa(totalitems))
	} else {
		fmt.Println("Total items: " + strconv.Itoa(page*30))
	}
	fmt.Println("Total pages: " + strconv.Itoa(page))
}

func stringReplacer(s string) string {
	// notice that there are two different spaces
	r := strings.NewReplacer("\n", "", " ", "", " ", "")
	return r.Replace(s)
}

func (c *Crawler) ItemandPageNum() (int, int) {
	v, err := query.Values(c.Options)
	if err != nil {
		log.Fatal(err)
	}
	url := rootURL + v.Encode()
	doc := NewDoc(url)
	r := strings.NewReplacer("\n", "", " ", "", ",", "")
	stritems := r.Replace(doc.Find(".pull-left.hasData > i").Text())
	items, _ := strconv.Atoi(stritems)
	pages := items/30 + 1

	return items, pages
}

func (c *Crawler) Jsonformat() []byte {
	b, err := json.MarshalIndent(c.Houselist, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func (c *Crawler) ExportJSON() {
	b, err := json.MarshalIndent(c.Houselist, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	_ = ioutil.WriteFile("houselist.json", b, 0644)
}
