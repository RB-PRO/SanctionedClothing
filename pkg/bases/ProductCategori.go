package bases

import "strings"

// Структура массива товаров
type Variety2 struct {
	Product []Product2 // Массив продуктов
}

// Категория Name Slug
type Cat [4]struct { // Категория товаров
	Name string // Название подкатегории
	Slug string // транслитом категория
}

// Структура товара
type Product2 struct {
	Cat // Подкатегория

	Name         string // Название товара
	FullName     string // Полное название товара
	Link         string // Сссылка на товар базового цвета
	Article      string // Артикул
	Manufacturer string // Производитель

	// Используется для tag
	GenderLabel string

	Size []string // Все возможные размеры

	Description struct { // Описание товара
		Eng string
		Rus string
	}

	// Описание товара по значению "цвет"
	// "Цвет" будет определять, как вариацию товара
	// "Цвет на русском"
	Item           map[string]ProdParam
	Specifications map[string]string // Остальные характеристики
}

// Структура параметров товара
type ProdParam struct {
	Link     string   // Ссылка на товар нужного цвета
	ColorEng string   // Цвет на английском
	Price    float64  // Цена
	Size     []string // Размеры
	Image    []string // Картинки
}

// Перевести /sweaters/CKvXARDQ1wHiAgIBAg.zso в sweaters
func FormingColorEng(input string) (output string) {
	input = strings.ReplaceAll(input, " ", "-")
	input = strings.ReplaceAll(input, "'", "")
	input = strings.ReplaceAll(input, "/", "_")
	output = strings.ToLower(input)
	return output
}

// Словарь, который используется для Name в GenderLabel
// и
// роидетльской категории. Например Женщины/woman
//
//	Функция принимает Woman[или]woman, а отдаёт Женщины
func GenderBook(key string) (string, bool) {
	keyLower := strings.ToLower(key) // Сделать нижний шрифт
	switch keyLower {
	case "women":
		return "Женщины", true
	case "man":
		return "Мужчины", true
	case "kid":
		return "Дети", true
	default:
		return key, false
	}
}
