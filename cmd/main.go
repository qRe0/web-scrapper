package main

import (
	"github.com/gocolly/colly"
	_ "github.com/mattn/go-sqlite3"
	
	dbp "web-scrapper/internal/db-handler/db-processing"
	imp "web-scrapper/internal/img-processing"
	call "web-scrapper/internal/middleware/html-callback"
	rs "web-scrapper/internal/middleware/range-scrap"
)

const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

func main() {
	c := colly.NewCollector()
	c.UserAgent = userAgent

	c.OnHTML("li.product", call.HTMLCallback)

	c.OnScraped(func(response *colly.Response) {
		dbp.WriteDataToDatabase(call.PokemonArr)
	})

	rs.RangeScrapping(c, 46, 46)

	imp.SaveImgToFile()
}
