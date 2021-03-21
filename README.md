# Crawler591
a web crawler to scrape rent information from 591 website


# Usage:


```go
import (

	"log"

	crawler "github.com/kevin51034/Crawler591"
)

var c *crawler.Crawler


func main() {

  c = crawler.Newcrawler()
  
  // set your options
  c.Options.RentPrice = "10000,15000"
  c.Options.Kind = 2
  c.Options.HasImg = "1"
  c.Options.NotCover = "1"
  c.Options.Role = "1"
  
  // get the items number within your condition.
  // the page number is items/30 because it's 30 items per page on 591 website.
  items, pages := c.ItemandPageNum()
  
  // Start crawling
  // choose the pages you want to crawler, and the result will store in c.Houselist with HouseInfo data struct.
  c.start(pages)
  
  // Export data
  c.ExportJSON()
}
```
# Crawler struct
```go
type Crawler struct {
	URL string
	Houselist []*HouseInfo
	Options   *Options
	wg        sync.WaitGroup
}
```
# Options struct
you can use this struct to define your crawler options.
```go
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
```

# HouseInfo struct in crawler.Houselist
```go
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
```
# Export json format:
```json
  {
    "id": 0,
    "imgsrc": "https://hp1.591.com.tw/house/active/2014/10/14/141326459274107405_765x517.water3.jpg",
    "title": "近古亭站優質套房出租一房一廳一衛",
    "url": "https://rent.591.com.tw/rent-detail-10119232.html",
    "kind": "獨立套房",
    "layout": "",
    "ping": "6.5坪",
    "floor": "樓層：頂樓加蓋/7",
    "address": "中正區-廈門街xxx巷",
    "price": "8500",
    "newItem": true,
    "crawlertime": "2020-11-17 4:23:49 pm",
    "updatetime": "2分鐘內更新",
    "author": "屋主許媽媽"
  },
```
