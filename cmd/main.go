package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
	_ "github.com/mattn/go-sqlite3"
)

type ScrappedData struct {
	Url   string
	Name  string
	Price string
	Img   string
}

var PokemonArr []ScrappedData

const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

func main() {
	c := colly.NewCollector()
	c.UserAgent = userAgent

	c.OnHTML("li.product", HTMLElement)

	c.OnScraped(func(response *colly.Response) {
		WriteDataToDatabase(PokemonArr)
	})

	err := c.Visit(fmt.Sprintf("https://scrapeme.live/shop/page/7/"))
	if err != nil {
		log.Fatal(err)
	}

	SaveImgToFile()
}

func HTMLElement(elm *colly.HTMLElement) {
	toScrap := ScrappedData{}

	toScrap.Url = elm.ChildAttr("a", "href")
	toScrap.Name = elm.ChildText("h2")
	toScrap.Price = elm.ChildText(".price")
	toScrap.Img = elm.ChildAttr("img", "src")

	PokemonArr = append(PokemonArr, toScrap)
}
