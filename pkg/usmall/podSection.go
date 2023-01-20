package usmall

import "github.com/gocolly/colly"

// Структура всех подсекций
type PodSection struct {
	Link []string // Ссылки из PodSectin
}

// Пропарсить PodSection
//
// [PodSection]: https://usmall.ru
func ParsePodSection() PodSection {
	var podSection PodSection
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("div[class='links mcol'] a[class!=subhead]", func(e *colly.HTMLElement) {
		hrefLink, isHref := e.DOM.Attr("href")
		if isHref {
			podSection.Link = append(podSection.Link, hrefLink)
		}
	})

	c.Visit(URL)

	return podSection
}
