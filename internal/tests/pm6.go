package tests

import (
	"fmt"

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
func Run_pm6_adventing_Sortered() {
	linkPages := "/null/.zso?s=brandNameFacetLC/asc/productName/asc/" // Ссылка на страницу товаров
	pagesInt := pm6.AllPages(linkPages)                               // Получить сколько всего страниц товаров есть
	pagesInt = 5

	var varient bases.Variety2                                 // Массив базы данных товаров
	varient = pm6.ParsePageWithVarienty(varient, linkPages, 0) // Парсим первую страницу товаров
	for i := 1; i <= pagesInt; i++ {                           // Цикл по всем страницам товаров
		fmt.Println(i, "/", pagesInt)
		// Сортируем товары и записываем их в готовую базу данных varient
		varient = pm6.ParsePageWithVarienty(varient, linkPages, i) // Парсим первую страницу товаров

		for j := 0; j < len(varient.Product); j++ {
			fmt.Println(">>", j, "/", len(varient.Product))
			if varient.Product[j].Manufacturer == "" {
				for key := range varient.Product[j].Item {
					//fmt.Println("parse", varient.Product[j].Item[key].Link)
					pm6.ParseProduct(&varient.Product[j], varient.Product[j].Item[key].Link)
				}

			}
		}
	}

	varient.SaveXlsxCsvs("TEST")

	fmt.Println("len", len(varient.Product))
	/*
		for index, value := range varient.Product {
			strs := ""
			for key := range value.Item {
				strs += key + ", "
			}
			fmt.Println(index, ":", ">"+value.Name+"<", "color:", strs)
		}
	*/
}
