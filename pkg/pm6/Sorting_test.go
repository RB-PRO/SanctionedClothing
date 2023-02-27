package pm6

import (
	"strconv"
	"testing"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
)

func TestSortingProducts(t *testing.T) {
	var varient bases.Variety2
	var addProds []bases.Product2

	// Создаём базовый массив
	varient.Product = append(varient.Product, bases.Product2{Name: "A"})
	varient.Product = append(varient.Product, bases.Product2{Name: "B"})
	varient.Product = append(varient.Product, bases.Product2{Name: "C"})
	varient.Product = append(varient.Product, bases.Product2{Name: "C"})

	// Создаём массив на добавление
	addProds = append(addProds, bases.Product2{Name: "C"})
	addProds = append(addProds, bases.Product2{Name: "D"})
	addProds = append(addProds, bases.Product2{Name: "D"})
	addProds = append(addProds, bases.Product2{Name: "D"})
	addProds = append(addProds, bases.Product2{Name: "G"})
	addProds = append(addProds, bases.Product2{Name: "G"})

	varient, addProds = SortingProducts(varient, addProds)
	t.Log("varient")
	for index, value := range varient.Product {
		t.Log(index, value.Name)
	}
	t.Log("addProds")
	for index, value := range addProds {
		t.Log(index, value.Name)
	}
}

func TestLastIndex(t *testing.T) {
	var addProds []bases.Product2

	// Создаём массив на добавление
	addProds = append(addProds, bases.Product2{Name: "C"})
	addProds = append(addProds, bases.Product2{Name: "D"})
	addProds = append(addProds, bases.Product2{Name: "D"})
	addProds = append(addProds, bases.Product2{Name: "D"})
	addProds = append(addProds, bases.Product2{Name: "G"})
	addProds = append(addProds, bases.Product2{Name: "G"})

	answer := LastIndex(addProds)
	if answer != 3 {
		t.Error("Неправильное значение последнего элемента. Должно было получиться 3, а получилось " + strconv.Itoa(answer))
	}
}
