package main

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	content, _ := goquery.NewDocument("http://sr48.net/follower?e=6")
	for i := 1; i < 420; i++ {
		name := content.Find("#content > table > tbody > tr:nth-child(" + strconv.Itoa(i) + ") > td:nth-child(3) > a:nth-child(3) > span").Text()
		follow :=strings.Split(content.Find("#content > table > tbody > tr:nth-child(" + strconv.Itoa(i) + ") > td:nth-child(3) > p.lay_r > span").Text(),"+")
		if len(follow)==1{follow=append(follow,"0")}
		if name != ""  {
			fmt.Println(name,  follow[0],follow[1])
		}
	}
}
