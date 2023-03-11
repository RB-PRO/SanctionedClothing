package pm6wp

import (
	"fmt"
	"log"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	"github.com/RB-PRO/SanctionedClothing/pkg/pm6"
	"github.com/RB-PRO/SanctionedClothing/pkg/wcprod"
)

// Функция, которая занимается вызовом функций парсинга и загрузки товаров
// Входные параметры:
// PageStart int  // Стартовая страница
// walrus float64 // Моржа
// delivery int   // Стоймость доставки
func Work(PageStart int, walrus float64, delivery int) {
	Adding, errorInitWcAdd := wcprod.NewWcAdd()

	if errorInitWcAdd != nil {
		log.Fatalln(errorInitWcAdd)
	}

	linkPages := "/null/.zso?s=brandNameFacetLC/asc/productName/asc/" // Ссылка на страницу товаров
	PageEnd := pm6.AllPages(linkPages)                                // Получить сколько всего страниц товаров есть
	PageEnd = 5                                                       // До этого мы парсим

	var varient bases.Variety2                                         // Массив базы данных товаров
	varient = pm6.ParsePageWithVarienty(varient, linkPages, PageStart) // Парсим первую страницу товаров
	for i := PageStart + 1; i <= PageEnd; i++ {                        // Цикл по всем страницам товаров
		fmt.Println("[pmwp]: Парсинг страниц", i, "/", PageEnd)

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
}
