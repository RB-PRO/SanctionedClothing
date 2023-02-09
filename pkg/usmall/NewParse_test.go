package usmall

import (
	"testing"
)

// Тест, который парсит страницы каталога
func TestParsePage(t *testing.T) {
	var variety Variety

	link := `/products/all/home/vyazanye-pledy` // Ссылка на категорию
	variety.ParsePage(link)

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
