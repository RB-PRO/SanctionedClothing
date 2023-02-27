package pm6

import (
	"fmt"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
)

// Сортировка массива товаров.
// Это значит, что данная функция должна обрабатывать
// На вход подаём значения массива всех товаров обработанных и массив товаров на добавление
func SortingProducts(varient bases.Variety2, addProds []bases.Product2) (bases.Variety2, []bases.Product2) {

	// При входных:
	// varient = AAABBBC
	// addProds = CCDDGG
	// Ищем такое    ^ ExitLastIndex = 3
	ExitLastIndex := LastIndex(addProds)

	var TecalName string
	TecalName = addProds[0].Name // Исходное имя товара
	for i := 0; i <= ExitLastIndex; i++ {
		if TecalName != addProds[i].Name {
			TecalName = addProds[i].Name

		}
	}
	fmt.Println("last:", ExitLastIndex)

	return varient, addProds
}

// Ищем самую последнюю позиция вхождения товаров одинакового цвета.
// Если товары AAABBCCCCDD, то будет индекс в конце товара C.
// т.е 8
func LastIndex(prod []bases.Product2) (ExitLastIndex int) {
	var TecalName string
	for index := range prod {
		// Если название новое
		if TecalName != prod[index].Name {
			TecalName = prod[index].Name
			ExitLastIndex = index - 1
		}
	}
	return ExitLastIndex
}

// Скрестить структуры продукта
// bases.Product2 in bases.Product2
func ConcatenateProduct2() {

}
