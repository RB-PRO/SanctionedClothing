package pm6

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
)

func TestAllPages(t *testing.T) {
	pagesInt := AllPages("/null/.zso?s=brandNameFacetLC/asc/productName/asc/")
	if pagesInt != 1131 {
		t.Error("Неправильное к-во товаров. По ссылке https://www.6pm.com/null/.zso?s=brandNameFacetLC/asc/productName/asc/ Должно быть \"1131\", а получено " + "\"" + strconv.Itoa(pagesInt) + "\"")
	}
}
func TestParsePageWithVarienty(t *testing.T) {
	linkPages := "/null/.zso?s=brandNameFacetLC/asc/productName/asc/" // Ссылка на страницу товаров
	pagesInt := 2                                                     // Получить сколько всего страниц товаров есть

	var varient bases.Variety2                             // Массив базы данных товаров
	varient = ParsePageWithVarienty(varient, linkPages, 0) // Парсим первую страницу товаров
	for i := 1; i <= pagesInt; i++ {                       // Цикл по всем страницам товаров
		// Сортируем товары и записываем их в готовую базу данных varient
		varient = ParsePageWithVarienty(varient, linkPages, i) // Парсим первую страницу товаров
	}
	PrintVarient(varient) // Печать
}

func PrintVarient(varient bases.Variety2) {
	fmt.Println("len", len(varient.Product))
	for index, value := range varient.Product {
		strs := ""
		for key := range value.Item {
			strs += key + ", "
		}
		fmt.Println(index, ":", ">"+value.Name+"<", "color:", strs)
	}
}
