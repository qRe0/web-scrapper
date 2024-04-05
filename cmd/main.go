package main

import (
	"fmt"
	"github.com/gocolly/colly"
	_ "github.com/mattn/go-sqlite3"
	"log"

	dbp "web-scrapper/internal/db_processing"
	dbs "web-scrapper/internal/db_structure"
	imgp "web-scrapper/internal/img_processing"
)

const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

var PokemonArr []dbs.ScrappedData

func main() {
	c := colly.NewCollector()
	c.UserAgent = userAgent

	c.OnHTML("li.product", HTMLElement)

	c.OnScraped(func(response *colly.Response) {
		dbp.WriteDataToDatabase(PokemonArr)
	})

	RangeScrapping(c, 21, 21)

	imgp.SaveImgToFile()
}

func HTMLElement(elm *colly.HTMLElement) {
	toScrap := dbs.ScrappedData{}

	toScrap.Url = elm.ChildAttr("a", "href")
	toScrap.Name = elm.ChildText("h2")
	toScrap.Price = elm.ChildText(".price")
	toScrap.Img = elm.ChildAttr("img", "src")

	PokemonArr = append(PokemonArr, toScrap)
}

func RangeScrapping(c *colly.Collector, startPage, endPage int) {
	for i := startPage; i <= endPage; i++ {
		err := c.Visit(fmt.Sprintf("https://scrapeme.live/shop/page/%d/", i))
		if err != nil {
			log.Fatal(err)
		}
	}
}
