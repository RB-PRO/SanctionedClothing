package pm6

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	"github.com/gocolly/colly"
)

// Парсинг страницы товара
// Парсинг будет выглядеть в виде редактирования структуры bases.Product2 со своевременным добавлением цвета
func ParseProduct(prod *bases.Product2, ProductColorLink string) {
	c := colly.NewCollector()
	c.UserAgent = "Golang"
	var tecalColor string // Цвет текущей страницы

	// Создаём структуру цвета
	c.OnHTML("span[class='NU-z']", func(e *colly.HTMLElement) {
		tecalColor = e.Text
		tecalColor = bases.FormingColorEng(tecalColor)
		fmt.Println("tecalColor", ">"+tecalColor+"<")
		for key := range prod.Item {
			fmt.Println(">" + key + "<")
		}
		prod.Item[tecalColor] = bases.ProdParam{ColorEng: tecalColor}
		prod.Specifications = make(map[string]string)
	})

	// Артикул, описание товара
	c.OnHTML("div[role='region'] ul li", func(e *colly.HTMLElement) {
		// Если это Артикул. В классер артикула всегда стоит "OR-z"
		if e.Attr("class") == "OR-z" {
			prod.Article = e.DOM.Find("span").Text()
			return
		}

		if strings.Contains(e.DOM.Text(), "Measurements:") {
			// Обработка дополнительных полей
			// Вынес в отдельный обработчик.
			return
		}

		// Обработка обычного описания товара
		if prod.Description.Eng == "" {
			prod.Description.Eng = e.Text
		} else {
			prod.Description.Eng += "\n" + e.Text
		}
	})

	// Описание товара по ключам
	c.OnHTML("div[role='region'] ul li ul li", func(e *colly.HTMLElement) {
		KeyValStr := strings.Split(e.DOM.Text(), ":")
		if len(KeyValStr) == 2 {
			KeyValStr[0] = strings.TrimSpace(KeyValStr[0])
			KeyValStr[1] = strings.TrimSpace(KeyValStr[1])
			prod.Specifications[KeyValStr[0]] = KeyValStr[1]
		}
	})

	// Размеры товара dpa-z epa-z
	c.OnHTML("div[class='dpa-z epa-z'] div[class^='Jqa-z'] input", func(e *colly.HTMLElement) {
		fmt.Println(e.DOM.Text())
		if attr, ok := e.DOM.Attr("data-label"); ok { // Если такой атрибут существует
			if entry, ok := prod.Item[tecalColor]; ok {
				entry.Size = append(entry.Size, attr)
				prod.Size = append(prod.Size, attr)
				prod.Item[tecalColor] = entry
			}
		}
	})

	// Картинки - Не работает.
	c.OnHTML("img[class^='IU-z']", func(e *colly.HTMLElement) {
		if attr, ok := e.DOM.Attr("src"); ok { // Если такой атрибут существует
			if entry, oks := prod.Item[tecalColor]; oks { // То добавляем его
				entry.Image = append(entry.Image, attr)
				prod.Item[tecalColor] = entry
			}
		}
	})

	// Картинки 2 div[class^='yq-z'] https://m.media-amazon.com/images/I/91GJ2hRcTeL.AC_SS144.jpg
	// https://m.media-amazon.com/images/I/${S}.AC_SS144.jpg
	c.OnHTML("img[itemprop=image]", func(e *colly.HTMLElement) {
		if attr, ok := e.DOM.Attr("srcSet"); ok { // Если такой атрибут существует
			if entry, oks := prod.Item[tecalColor]; oks { // То добавляем его
				entry.Image = append(entry.Image, attr)
				prod.Item[tecalColor] = entry
			}
		}
	})

	// Категории + производитель
	c.OnHTML("div[id=breadcrumbs] div", func(e *colly.HTMLElement) {
		prod.Cat[1].Name = e.DOM.Find("a:nth-of-type(2)").Text()
		prod.Cat[1].Slug, _ = e.DOM.Find("a:nth-of-type(2)").Attr("href")
		prod.Cat[1].Slug = formSlump(prod.Cat[1].Slug, 1) // Редактирование ссылки

		prod.Cat[2].Name = e.DOM.Find("a:nth-of-type(3)").Text()
		prod.Cat[2].Slug, _ = e.DOM.Find("a:nth-of-type(3)").Attr("href")
		prod.Cat[2].Slug = formSlump(prod.Cat[2].Slug, 1) // Редактирование ссылки

		prod.Cat[3].Name = e.DOM.Find("a:nth-of-type(4)").Text()
		prod.Cat[3].Slug, _ = e.DOM.Find("a:nth-of-type(4)").Attr("href")
		prod.Cat[3].Slug = formSlump(prod.Cat[3].Slug, 2) // Редактирование ссылки
		prod.Manufacturer = prod.Cat[3].Name              // Производитель
	})

	// Название Товара
	c.OnHTML("span[class=yn-z]", func(e *colly.HTMLElement) {
		prod.Name = e.DOM.Text()
	})
	// Ссылка на товар
	c.OnHTML("meta[itemprop=url]", func(e *colly.HTMLElement) {
		prod.Link, _ = e.DOM.Attr("content")
		if entry, oks := prod.Item[tecalColor]; oks { // То добавляем его
			entry.Link = prod.Link
			prod.Item[tecalColor] = entry
		}
	})
	// Полное название товара, оно же краткое описание товара
	c.OnHTML("li[class=LR-z]", func(e *colly.HTMLElement) {
		prod.FullName = e.DOM.Text()
	})
	// Гендер товара
	c.OnHTML("span[class=hpa-z]", func(e *colly.HTMLElement) {
		prod.GenderLabel = e.DOM.Text()                                          // Получить текст
		prod.GenderLabel = strings.ReplaceAll(prod.GenderLabel, "'s Sizes:", "") // Удалить лишнее из гендера
		prod.Cat[0].Name = prod.GenderLabel                                      // Название главной категории товара
		prod.GenderLabel = strings.ToLower(prod.GenderLabel)                     // Понизить регистр
		prod.Cat[0].Slug = prod.GenderLabel                                      // Название главной ссылки категории товара
	})

	// Цена
	c.OnHTML("span[class=eq-z]", func(e *colly.HTMLElement) {
		fmt.Println(e.DOM.Text())
		coast := e.DOM.Text()
		coast = strings.ReplaceAll(coast, "$", "")
		floaCoast, errCoast := strconv.ParseFloat(coast, 64) // Преобразование типов
		if errCoast == nil {
			if entry, oks := prod.Item[tecalColor]; oks { // То добавляем его
				entry.Price = floaCoast
				prod.Item[tecalColor] = entry
			}
		}
	})

	// Обработка ошибки
	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Ошибка. Поэтому немного ждём. Ошибка", e)
		time.Sleep(5 * time.Second)
	})

	// Обработка ошибки после ответа сервера
	c.OnResponse(func(r *colly.Response) {
		if r.StatusCode != http.StatusOK { // Если нет ответа
			fmt.Println("Ошибка. Поэтому немного ждём. Статус", r.StatusCode)
			time.Sleep(5 * time.Second)
		}
	})

	fmt.Println("Product.go", URL+ProductColorLink)
	c.Visit(URL + ProductColorLink)

}

// Перевести /sweaters/CKvXARDQ1wHiAgIBAg.zso в sweaters
func formSlump(input string, selection int) (output string) {
	input = strings.ReplaceAll(input, ":", "")
	strs := strings.Split(input, "/")
	if len(strs) >= selection+1 {
		return strs[selection]
	}
	return ""
}

// Распечатать продукт
func PrintProduct2(prod bases.Product2) (output string) {
	output = "Название: " + prod.Name + ". Артикул: " + prod.Article + "\n" +
		"Производитель: " + prod.Manufacturer + ". Гендер: " + prod.GenderLabel + "\n" + " Название(Полн): " + prod.FullName + "\n" +
		"Ссылка: " + prod.Link + "\n" +
		"Размеры: " + strings.Join(prod.Size, ",") + "\n" +
		"Подкатегория: " + fmt.Sprintf("%+v", prod.Cat) + "\n" +
		"Описание(Рус): " + prod.Description.Rus + "\n" +
		"Описание(Eng): " + prod.Description.Eng + "\n" +
		"Дополнительные поля: " + fmt.Sprintf("%+v", prod.Specifications) + "\n" +
		"Подробнее по каждому цвету:\n" + PrintItems(prod.Item)

	return output
}
func PrintItems(items map[string]bases.ProdParam) (output string) {
	for key, val := range items {
		output += key + " - " + val.ColorEng + "\n" +
			"\tЦена: " + fmt.Sprintf("%.2f", val.Price) + ". Ссылка: " + val.Link + "\n" +
			"\tРазмеры(" + strconv.Itoa(len(val.Size)) + "): " + strings.Join(val.Size, ",") + "\n" +
			"\tКартинка: " + strings.Join(val.Image, ",") + "\n"
	}
	return output
}
