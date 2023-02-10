package usmall

import (
	"testing"
	"time"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
)

// Тест, который парсит страницы каталога
func TestParsePage(t *testing.T) {
	var variety bases.Variety2

	link := `/products/all/home/vyazanye-pledy` // Ссылка на категорию
	ParsePage(&variety, link)

	if len(variety.Product) == 0 {
		t.Fatal(`Товары не найдены. Вообще.`)
	}

	if variety.Product[0].Link == "" {
		t.Fatalf(`Ссылка "%s" пустая, а должно хоть что-нибудь.\nСсылка на категорию "%s"`, variety.Product[0].Link, URL+link)
	}
}

// Количество страниц в подсекции
func TestLenPodSection(t *testing.T) {
	var link string
	var lens int

	// Тесторирование, когда 1 страница каталога
	link = `/products/all/home/vyazanye-pledy` // Ссылка на категорию
	lens = LenPodSection(link)
	if lens != 1 {
		t.Fatalf(`Страницы категории не найдены. Результат: %d, а должно быть 1.\nСсылка на каталог: %s`, lens, URL+link)
	}

	// Тесторирование, когда более 1-й страницы каталога
	link = `/products/women/clothes/shuby` // Ссылка на категорию
	lens = LenPodSection(link)
	if lens == 0 {
		t.Fatalf(`Страницы категории некорректны. Результат: %d\nСсылка на каталог: %s`, lens, URL+link)
	}

}

// Тестирование конвентора
func TestCodeOfLink(t *testing.T) {
	link := `https://usmall.ru/product/477964-cropped-faux-fur-jacket-avec-les-filles?color=red`
	code, errorCode := CodeOfLink(link)
	if code != "477964" && errorCode != nil {
		t.Fatalf(`Код страницы иной. Результат: %s, а должно быть %s.\n%v`, code, "477964", errorCode)
	}
}

// Полная страница парсинга на примере вязаных пледов
func TestMain(t *testing.T) {
	var podsec PodSection
	podsec.Link = make([]string, 1)
	podsec.Link[0] = "/products/all/home/vyazanye-pledy"

	var variety bases.Variety2
	ParsePage(&variety, podsec.Link[0])

	//variety.Product = variety.Product[:5] // debug

	for i := 0; i < len(variety.Product); i++ {
		MyCode := variety.Product[i].Link         // Код товара
		MyCode, _ = CodeOfLink(MyCode)            // Вычленить код товара
		ware, _ := Ware(MyCode)                   // Получить запрос с API
		WareInProduct2(&variety.Product[i], ware) // Преобразовать в домашнюю структуру
		time.Sleep(20 * time.Microsecond)
	}

	// *************************************************
	variety.SaveXlsxCsvs("strPodSection") // Cохранить в формате из ТЗ
}

func TestWareInProduct2(t *testing.T) {
	var variety bases.Variety2
	variety.Product = make([]bases.Product2, 1)
	MyCode := "https://usmall.ru/product/334932-mens-quilted-hooded-bomber-jacket-dkny?color=oxblood" // Код товара
	MyCode, _ = CodeOfLink(MyCode)                                                                    // Вычленить код товараw
	ware, _ := Ware(MyCode)                                                                           // Получить запрос с API
	WareInProduct2(&variety.Product[0], ware)                                                         // Преобразовать в домашнюю структуру

	t.Log("\n   All:", variety.Product[0].Item["Красный (Oxblood)"], "\n")
	t.Log("\n Sizes:", variety.Product[0].Item["Красный (Oxblood)"].Size, "\n")
	variety.SaveXlsxCsvs("test")
}
