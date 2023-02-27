package pm6

import (
	"strconv"
	"time"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	"github.com/gocolly/colly"
)

const URL = "https://www.6pm.com"

// Пропасить стрицы товара и записать данные в канал
func ParsePages(ProductChan chan<- bases.Product2) {
	ProductChan <- bases.Product2{Name: "123"}
	time.Sleep(time.Second) // Задержка
}

// Пропасить стрицу с товарами
func ParsePage(link string, page int) (prod []bases.Product2) {
	c := colly.NewCollector()
	c.OnHTML("article[class='qda-z']", func(e *colly.HTMLElement) {
		link, ErrorAttrLink := e.DOM.Find("a[class='OP-z']").Attr("href") // Получить ссылку на товар
		if ErrorAttrLink {
			color := e.DOM.Find("dd[class='UP-z']").Text() // Получить цвет товара
			color = formingColorEng(color)                 // Преобразовать в ссылку
			prod = append(prod, bases.Product2{
				Link: link,
				Name: e.DOM.Find("dd[class='SP-z']").Text(),
				Item: make(map[string]bases.ProdParam),
			})
			prod[len(prod)-1].Item[color] = bases.ProdParam{}
		}

	})
	c.Visit(URL + link + "&p=" + strconv.Itoa(page))
	return prod
}

// Получить все страницы с товарами на сайте
func AllPages(link string) (pages int) {
	c := colly.NewCollector()
	c.OnHTML("span[class='vm-z']", func(e *colly.HTMLElement) {
		pagesStr := e.DOM.Find("a:last-of-type").Text()
		pages, _ = strconv.Atoi(pagesStr)

	})
	c.Visit(URL + link)
	return pages
}
