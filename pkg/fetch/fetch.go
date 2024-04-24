package fetch

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
)

// Author:	DragonChen https://github.com/dragonchen-tw/
// Title:	Package - fetch
// Date:	2024/04/20

func Fetch() {
	url := "https://dragonchen.tw/links/"
	// url := "https://min-api.cryptocompare.com/data/price?fsym=ETH&tsyms=ETH,USD"

	c := colly.NewCollector()

	c.OnHTML("li > a", func(e *colly.HTMLElement) {
		fmt.Println(e.Text, e.Attr("href"))
	})

	// c.OnHTML("#caller_2038-0", func(e *colly.HTMLElement) {
	// 	fmt.Println(e.Text)
	// })

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(url)
	c.Wait()

	// // JSON
	// var data map[string]float32
	// json.Unmarshal([]byte(res), &data)
	// fmt.Println("ETH price", data)
}

func DynamicFetch() {
	url := "https://www.mainpi.com/query?i=2038"

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var waitBefore, waitAfter string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(time.Second*1),
		chromedp.OuterHTML("#caller_2038-0", &waitBefore, chromedp.ByQuery),
		chromedp.OuterHTML("#caller_1435-0", &waitAfter, chromedp.ByQuery),
	); err != nil {
		log.Fatal(err)
	}

	fmt.Println(waitBefore, waitAfter)

	document1, err := NewDocuemntFromString(waitBefore)
	if err != nil {
		log.Fatal(err)
	}
	document2, err := NewDocuemntFromString(waitAfter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(document1.Text(), document2.Text())
}

func NewDocuemntFromString(HTMLString string) (*goquery.Document, error) {
	document, err := goquery.NewDocumentFromReader(strings.NewReader(HTMLString))
	if err != nil {
		return nil, err
	}
	return document, err
}
