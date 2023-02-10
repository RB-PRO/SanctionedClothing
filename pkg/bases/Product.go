package bases

// Структура массива товаров
type Variety2 struct {
	Product []Product2 // Массив продуктов
}

// Структура товара
type Product2 struct {
	Catalog    string // Каталог: Женщины, Мужчины, Здоровье, Девочки, Мальчики
	PodCatalog string // ПодКаталог: Одежда, Обувь, Сумки
	Section    string // Верхняя одежда, Платья, Юбки
	PodSection string // Шубы, Пуховики,Пальто

	Name         string // Название товара
	FullName     string // Полное название товара
	Link         string // ССсылка на товар базового цвета
	Article      string // Артикул
	Manufacturer string // Производитель

	Description struct { // Описание товара
		Eng string
		Rus string
	}

	// Описание товара по значению "цвет"
	// "Цвет" будет определять, как вариацию товара
	// "Цвет на русском"
	Item map[string]ProdParam
}

// Структура параметров товара
type ProdParam struct {
	Link           string            // Ссылка на товар нужного цвета
	ColorEng       string            // Цвет на английском
	Price          float64           // Цена
	Size           []string          // Размеры
	Image          []string          // Картинки
	Specifications map[string]string // Остальные характеристики
}
