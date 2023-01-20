package usmall

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

// Ссылка на сайт
const URL string = "https://usmall.ru"

// Структура массива товаров
type Variety struct {
	Product []Product // Массив продуктов
}

// Структура товара
type Product struct {
	Catalog    string // Каталог: Женщины, Мужчины, Здоровье, Девочки, Мальчики
	PodCatalog string // ПодКаталог: Одежда, Обувь, Сумки
	Section    string // Верхняя одежда, Платья, Юбки
	PodSection string // Шубы, Пуховики,Пальто

	Name         string   // Название товара
	FullName     string   // Полное название товара
	Link         string   // Ссылка на товар
	Article      string   // Артикул
	Manufacturer string   // Производитель
	Price        float64  // Цена
	Description  struct { // Описание товара
		eng string
		rus string
	}
	ImageLink      []string          // Ссылки на картинки
	Colors         []string          // Цвета
	Size           []string          // Размеры
	Specifications map[string]string // Остальные характеристики
}

// Метод, который парсит [страницу товара]
//
// [страницу товара]: https://usmall.ru/product/477964-cropped-faux-fur-jacket-avec-les-filles
func (product *Product) ParseProduct() {

	c := colly.NewCollector()
	mapas := make(map[string]string)

	// Названия пути к товару и полное название товара
	c.OnHTML("nav[class='c-crumbs wrapper']", func(e *colly.HTMLElement) {
		product.Catalog = e.DOM.Find("span:nth-child(2) a").Text()
		product.PodCatalog = e.DOM.Find("span:nth-child(3) a").Text()
		product.Section = e.DOM.Find("span:nth-child(4) a").Text()
		product.PodSection = e.DOM.Find("span:nth-child(5) a").Text()
		product.FullName = e.DOM.Find("span:nth-child(6) span").Text()
	})

	// Артикул
	c.OnHTML("div[class='p-product__vendor-code'] span", func(e *colly.HTMLElement) {
		product.Article = e.DOM.Text()
	})

	// Производитель
	c.OnHTML("h1[class='p-product__h1'] span", func(e *colly.HTMLElement) {
		product.Manufacturer = e.DOM.Text()
	})

	// Название товара
	c.OnHTML("span[class='__product-name']", func(e *colly.HTMLElement) {
		product.Name = e.DOM.Text()
	})
	// Цена
	c.OnHTML("span[class='p-product__price-value']", func(e *colly.HTMLElement) {
		cost := e.DOM.Text()
		//fmt.Println("cost", cost)
		reg := regexp.MustCompile("[^0-9.]+")
		replaceStr := reg.ReplaceAllString(cost, "")
		//fmt.Println("replaceStr", replaceStr)
		if n, err := strconv.ParseFloat(replaceStr, 64); err == nil {
			product.Price = n
		}
	})

	// Цвета
	c.OnHTML("div[class='p-product__color-list'] span[class^='__pseudo']", func(e *colly.HTMLElement) {
		colorSet, colorError := e.DOM.Attr("title")
		if colorError {
			product.Colors = append(product.Colors, colorSet)
		}
	})

	// Размеры
	c.OnHTML("div[class='p-product__size-list __col-6'] label[class^='__pseudo']", func(e *colly.HTMLElement) {
		product.Size = append(product.Size, e.DOM.Text())
	})

	// Картинки
	// background-image:url(https://usmall.ru/image/047/79/64/bac5359c2f1a5671ee8fdfc565fc55fb.jpeg);
	c.OnHTML("div[class^='thumb-scroller__thumb']", func(e *colly.HTMLElement) {
		imgHref, isImage := e.DOM.Attr("style")
		if isImage {
			imgHref = strings.ReplaceAll(imgHref, "background-image:", "")
			imgHref = strings.ReplaceAll(imgHref, "url(", "")
			imgHref = strings.ReplaceAll(imgHref, ");", "")
			imgHref = strings.ReplaceAll(imgHref, "\"", "")
			imgHref = strings.ReplaceAll(imgHref, "https://usmall.ru", "")
			imgHref = strings.TrimSpace(imgHref)

			product.ImageLink = append(product.ImageLink, imgHref)
		}
	})

	// Дополнительные параметры в Мапе
	c.OnHTML("div[class='__text __facets'] tr", func(e *colly.HTMLElement) {
		mapas[e.DOM.Find("td[class='__label']").Text()] = e.DOM.Find("td[class='__value']").Text()
	})

	// Описание
	c.OnHTML("div[class='__content'] div[class='__text'] ul", func(e *colly.HTMLElement) {
		product.Description.rus = strings.ReplaceAll(e.DOM.Text(), "  ", "")
	})

	c.Visit(URL + product.Link)

	product.Specifications = mapas // Заполнение дополнительных параметров
}

// Пропарсить PodSection на ссылки на товары
//
// Спарсить один подраздел и создать карточку с товарами
// [PodSection]: /products/boy/clothes/kids-robes
func (variety *Variety) ParsePage(link string) {
	var exit bool // Переменная, отвечающая за выходиз цикла, который просматривает все товары
	c := colly.NewCollector()

	// Поиск ссылки на товар
	c.OnHTML("a[class='__img __fg']", func(e *colly.HTMLElement) {
		hrefLink, isHref := e.DOM.Attr("href")
		if isHref {
			variety.Product = append(variety.Product, Product{Link: hrefLink})
		}
	})

	// Тег a или span, который должен сожержать или нет ссылку на следующую страницу. Если нет - выходим
	c.OnHTML("[class^=__next]", func(e *colly.HTMLElement) {
		hrefLink, isHref := e.DOM.Attr("href")
		if isHref {
			link = hrefLink
		} else {
			exit = true
		}
	})

	// Бесконечный цикл по pages
	for !exit {
		c.Visit(URL + link)
	}
}
