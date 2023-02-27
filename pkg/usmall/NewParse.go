// Файл для сбора данных по запросу по конкретному продукту
package usmall

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	"github.com/gocolly/colly"
)

// Ссылка на сайт
const URL string = "https://usmall.ru"

// https://usmall.ru/api/product/477964

type WareUsmall struct {
	ID         int        `json:"id"`          // ID товара
	Name       string     `json:"name"`        // Название товара Рус
	OriginName string     `json:"origin_name"` // Название товара оригинальное Eng
	Categories []struct { // Категория товара
		ID              int    `json:"id"` // ID категории
		Name            string `json:"name"`
		NameForSeo      string `json:"name_for_seo"`
		NameMaleRu      string `json:"name_male_ru"`
		NameMaleExtRu   string `json:"name_male_ext_ru"`
		NameFemaleRu    string `json:"name_female_ru"`
		NameFemaleExtRu string `json:"name_female_ext_ru"`
		NameGirlRu      string `json:"name_girl_ru"`
		NameGirlExtRu   string `json:"name_girl_ext_ru"`
		NameBoyRu       string `json:"name_boy_ru"`
		NameBoyExtRu    string `json:"name_boy_ext_ru"`
		Tree            struct {
			Num2  string `json:"2"`
			Num21 string `json:"21"`
			Num34 string `json:"34"`
		} `json:"tree"`
		Alias []interface{} `json:"alias"`
	} `json:"categories"`
	GenderLabel       string   `json:"genderLabel"`        // Пол товара
	Description       string   `json:"description"`        // Описание Рус
	OriginDescription string   `json:"origin_description"` // Описание Eng
	Brand             struct { // Бренд товара
		ID   int    `json:"id"`
		Name string `json:"name"` // Название бренда
	} `json:"brand"`

	Variants []struct {
		Images []struct { // Картинки
			URL    string `json:"url"` // Ссылка на картинку
			Width  int    `json:"width"`
			Height int    `json:"height"`
			Size   int    `json:"size"`
		} `json:"images"`
		Price       int    `json:"price"`         // Цена
		OriginColor string `json:"origin_color"`  // Цвет Eng иной
		ColorName   string `json:"color_name"`    // Цвет Eng
		ColorNameRu string `json:"color_name_ru"` // Цвет Рус
		OriginSize  string `json:"origin_size"`   // Размер одежды Eng
		RussianSize struct {
			Name string `json:"name"` // Размер одежды Рус
		} `json:"russianSize"`
	} `json:"variants"`
}

// Получить по коду данные с API по товару
//
// Перед этим необходимо взять конкретный код товара
func Ware(code string) (WareUsmall, error) {
	// Ответ от сервера
	var WareUsmallRes WareUsmall

	// Выполнить запрос
	resp, responseError := http.Get(URL + "/api/product/" + code)
	if responseError != nil {
		return WareUsmall{}, responseError
	}

	// Закрыть канал в коце работы
	defer resp.Body.Close()

	// Получить тело ответа
	body, errIoReadAll := io.ReadAll(resp.Body)
	if errIoReadAll != nil {
		return WareUsmall{}, errIoReadAll
	}

	//fmt.Println(string(body))

	// Распарсить данные
	responseErrorUnmarshal := json.Unmarshal(body, &WareUsmallRes)
	if responseErrorUnmarshal != nil {
		return WareUsmall{}, responseErrorUnmarshal
	}

	return WareUsmallRes, nil
}

/*
// Преобразовать результат работы API во внутреннюю структуру данных
//
//	WareUsmall > Product
//
// NE MENYAT
func (product *Product) WareInProduct(ware WareUsmall) {
	product.Name = ware.Name                // Название
	product.FullName = ware.OriginName      // Полное название
	product.Article = strconv.Itoa(ware.ID) // Артикул
	product.Manufacturer = ware.Brand.Name  // Производитель

	product.Description.eng = ware.OriginDescription // Описание на английском
	product.Description.rus = ware.Description       // Описание на русском

	product.Image = make(map[string][]string) // Выделить память в мапу
	for _, valueWare := range ware.Variants { // Цикл по всем возможным вариантам товара
		product.Size = append(product.Size, valueWare.OriginSize)      // Размеры
		product.Colors = append(product.Colors, valueWare.ColorNameRu) // Цвета
		product.Price = float64(valueWare.Price)                       // Цена

		// Цвет
		product.Image[valueWare.ColorNameRu] = make([]string, 0) // Выделить память в часть мапы
		for _, valueImage := range valueWare.Images {
			product.Image[valueWare.ColorNameRu] = append(product.Image[valueWare.ColorNameRu], valueImage.URL)
		}
	}

	// Удалить дубликаты в ссылках
	product.Size = RemoveDuplicateStr(product.Size)
	product.Colors = RemoveDuplicateStr(product.Colors)
	for keyImage := range product.Image {
		product.Image[keyImage] = RemoveDuplicateStr(product.Image[keyImage])
	}
}
*/

// Удалить дубликаты в слайсе
func RemoveDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

// Получить код товара из полной ссылки на товар
//
//	https://usmall.ru/product/477964-cropped-faux-fur-jacket-avec-les-filles?color=red
//
// преобразуется в
//
//	477964
func CodeOfLink(link string) (string, error) {
	link = strings.ReplaceAll(link, URL, "")
	link = strings.ReplaceAll(link, "/product/", "")
	linkStrs := strings.Split(link, "-")
	if len(linkStrs[0]) == 0 {
		return "", errors.New("null link")
	} else {
		return linkStrs[0], nil
	}
}

// Получить категорию товара
//
// Он же путь к товару
func CatalogsCat(link string) (cat bases.Cat) {

	c := colly.NewCollector()
	c.UserAgent = "Golang"

	c.OnHTML("nav[class='c-crumbs wrapper']", func(e *colly.HTMLElement) {
		cat[0].Name = e.DOM.Find("span:nth-child(2) a").Text()
		cat[0].Slug, _ = e.DOM.Find("span:nth-child(2) a").Attr("href")
		cat[1].Name = e.DOM.Find("span:nth-child(3) a").Text()
		cat[1].Slug, _ = e.DOM.Find("span:nth-child(3) a").Attr("href")
		cat[2].Name = e.DOM.Find("span:nth-child(4) a").Text()
		cat[2].Slug, _ = e.DOM.Find("span:nth-child(4) a").Attr("href")
		cat[3].Name = e.DOM.Find("span:nth-child(5) a").Text()
		cat[3].Slug, _ = e.DOM.Find("span:nth-child(5) a").Attr("href")
		cat[0].Slug = lastSlush(cat[0].Slug)
		cat[1].Slug = lastSlush(cat[1].Slug)
		cat[2].Slug = lastSlush(cat[2].Slug)
		cat[3].Slug = lastSlush(cat[3].Slug)
	})
	c.Visit(link)

	return cat
}

// Конфертировать
//
//	/products/men/clothes/casual-jackets/levis
//
// в
//
//	levis
func lastSlush(inputStr string) string {
	inputsStrs := strings.Split(inputStr, "/")
	return inputsStrs[len(inputsStrs)-1]
}

// Преобразовать результат работы API во внутреннюю структуру данных
//
//	WareUsmall > Product
func WareInProduct2(product *bases.Product2, ware WareUsmall) {
	// product.Cat

	product.Name = ware.Name                // Название
	product.FullName = ware.OriginName      // Полное название
	product.Article = strconv.Itoa(ware.ID) // Артикул
	product.Manufacturer = ware.Brand.Name  // Производитель
	product.GenderLabel = ware.GenderLabel  // tag

	product.Description.Eng = ware.OriginDescription // Описание на английском
	product.Description.Rus = ware.Description       // Описание на русском
	product.Specifications = make(map[string]string)

	// product.Image = make(map[string][]string) // Выделить память в мапу

	// Выделяем память
	product.Item = make(map[string]bases.ProdParam)
	for _, valueWare := range ware.Variants {
		colorKey := valueWare.ColorNameRu + " (" + valueWare.OriginColor + ")"

		// Если НЕ существует мапа такого цвета
		if _, ok := product.Item[colorKey]; !ok {
			// То выделяем память в данные соответственно
			product.Item[colorKey] = bases.ProdParam{
				Image: make([]string, 0),
				Size:  make([]string, 0),
			}
			//product.Item[colorKey].Image = make([]string, 0)
			//product.Item[colorKey].Size = make([]string, 0)
			//product.Item[colorKey].Specifications = make(map[string]string)
		}

		// Если существует мапа такого цвета
		if entry, ok := product.Item[colorKey]; ok {
			// Название цвета
			entry.ColorEng = valueWare.OriginColor

			// Цена
			entry.Price = float64(valueWare.Price)

			// Картинки
			for _, valueImage := range valueWare.Images {
				entry.Image = append(entry.Image, valueImage.URL)
			}

			// Размер
			//entry.Size = append(entry.Image, valueWare.RussianSize.Name)
			entry.Size = append(entry.Size, valueWare.OriginSize)
			product.Size = append(product.Size, valueWare.OriginSize) // Все всех размеров

			// Ссылка на страницу с цветом
			entry.Link = URL + product.Link + "/?color=" + valueWare.OriginColor

			// Добавляем эту промежуточную строкуту в нашу
			product.Item[colorKey] = entry
		}
	}
	product.Size = RemoveDuplicateStr(product.Size)
}
