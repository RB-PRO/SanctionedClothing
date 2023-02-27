package pm6

import (
	"errors"
	"fmt"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
	"golang.org/x/exp/maps"
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
	var startIndex int
	TecalName = addProds[0].Name // Исходное имя товара
	fmt.Println("ExitLastIndex", ExitLastIndex)
	for i := 0; i <= ExitLastIndex; i++ {
		if TecalName != addProds[i].Name || (ExitLastIndex == i && ExitLastIndex < len(addProds)) {
			indexTecalName, indexTecalNameError := FindFirstNameProducts(varient.Product, TecalName)
			if indexTecalNameError != nil { // Если найдено
				fmt.Println("Нашёл ", indexTecalName)
			} else {
				varient.Product = append(varient.Product, ConcatenateProduct2(addProds[startIndex:i]))
			}
			TecalName = addProds[i].Name
			startIndex = i

		}
	}
	fmt.Println("last:", ExitLastIndex)

	return varient, addProds[startIndex+1:]
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
func ConcatenateProduct2(prodInPut []bases.Product2) (prodOutput bases.Product2) {
	prodOutput = prodInPut[0]

	// Объединить мапы
	for _, valueKey := range prodInPut {
		//maps.Copy(prodOutput.Item, valueKey.Item)
		prodOutput = ConcatenateOneProduct2(prodOutput, valueKey)
	}

	return prodOutput
}

// Скрестить структуры 2 продуктов
// bases.Product2 in bases.Product2
func ConcatenateOneProduct2(prodInPut1, prodInPut2 bases.Product2) bases.Product2 {
	maps.Copy(prodInPut1.Item, prodInPut2.Item) // Объединить мапы
	return prodInPut1
}

// Поиск структуры по значению. На выходе инта - порядковый номер
func FindFirstNameProducts(prod []bases.Product2, findStr string) (int, error) {
	for index, valueProduct := range prod {
		if valueProduct.Name == findStr {
			return index, nil
		}
	}
	return 0, errors.New("Не найдено " + findStr)
}
