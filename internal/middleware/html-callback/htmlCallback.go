package html_callback

import (
	"github.com/gocolly/colly"
	dbs "web-scrapper/internal/db-handler/db-structure"
)

var PokemonArr []dbs.ScrappedData

func HTMLCallback(elm *colly.HTMLElement) {
	toScrap := dbs.ScrappedData{}

	toScrap.Url = elm.ChildAttr("a", "href")
	toScrap.Name = elm.ChildText("h2")
	toScrap.Price = elm.ChildText(".price")
	toScrap.Img = elm.ChildAttr("img", "src")

	PokemonArr = append(PokemonArr, toScrap)
}
