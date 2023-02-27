package tests

import (
	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	"github.com/RB-PRO/SanctionedClothing/pkg/pm6"
)

func Run_pm6() {
	linkPages := "/null/.zso?s=brandNameFacetLC/asc/productName/asc/" // Ссылка на страницу товаров
	pagesInt := pm6.AllPages(linkPages)                               // Получить сколько всего страниц товаров есть

	var AddProducts []bases.Product2          // дополнительный массив товаров
	var varient bases.Variety2                // Массив базы данных товаров
	AddProducts = pm6.ParsePage(linkPages, 1) // Парсим первую страницу товаров
	for i := 2; i <= pagesInt; i++ {          // Цикл по всем страницам товаров
		// Сортируем товары и записываем их в готовую базу данных varient
		varient, AddProducts = pm6.SortingProducts(varient, AddProducts)

		//AddProducts = append(AddProducts, pm6.ParsePage(linkPages, i)...) // Собираем массив товаров

		//varient.Product = append(varient.Product, AddProducts...) // Добавляем в исходный массив товаров
	}

}
