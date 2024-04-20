package fetch

import (
	"fmt"
	"os"

	"github.com/anaskhan96/soup"
)

// Author:	DragonChen https://github.com/dragonchen-tw/
// Title:	Package - fetch
// Date:	2024/04/20

func Fetch() {
	url := "https://dragonchen.tw/links/"
	res, err := soup.Get(url)
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(res)
	links := doc.FindAll("a")
	for _, link := range links {
		fmt.Println(link.Text(), "| Link :", link.Attrs()["href"])
	}
}
