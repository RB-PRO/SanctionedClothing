package usmall

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

	Name           string            // Название товара
	FullName       string            // Полное название товара
	Link           string            // Ссылка на товар
	Article        string            // Артикул
	Manufacturer   string            // Производитель
	ImageLink      []string          // Ссылка на картинки
	Price          float64           // Цена
	Specifications map[string]string // Остальные характеристики
	Colors         []string          // Цвета
	Size           []string          // Размеры
	Description    struct {          // Описание товара
		eng string
		rus string
	}
}

// Пропарсить PodSection
//
// [PodSection]: https://usmall.ru/products/women/clothes/faux-fur-shearling-coats
func (items *Variety) ParsePodSection(link string) {

}
