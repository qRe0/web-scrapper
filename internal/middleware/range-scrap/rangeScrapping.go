package range_scrap

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

func RangeScrapping(c *colly.Collector, startPage, endPage int) {
	for i := startPage; i <= endPage; i++ {
		err := c.Visit(fmt.Sprintf("https://scrapeme.live/shop/page/%d/", i))
		if err != nil {
			log.Fatal(err)
		}
	}
}
