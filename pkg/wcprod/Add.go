package wcprod

import (
	"fmt"
	"io"
	"os"

	wc "github.com/hiscaler/woocommerce-go"
)

// Добавить аттрибут в Woocommerce
func AddAttr(wooClient *wc.WooCommerce, idAttrColor int, newName, NewSlug string) (tecalAttrId int, tecalAttrName string, tecalAttrSlug string) {
	items, total, _, _, _ := wooClient.Services.ProductAttributeTerm.All(idAttrColor, wc.ProductAttributeTermsQueryParaTerms{Slug: NewSlug})
	// Если такого цвета не существует, то создаём его
	if total == 0 {
		AttributeTermCreate, errorCreate := wooClient.Services.ProductAttributeTerm.Create(idAttrColor, wc.CreateProductAttributeTermRequest{
			Name:        newName,
			Slug:        NewSlug,
			Description: "Создано автоматически при загрузке товара",
		})
		if errorCreate != nil {
			fmt.Println(errorCreate)
		}
		tecalAttrId = AttributeTermCreate.ID
		tecalAttrName = AttributeTermCreate.Name
		tecalAttrSlug = AttributeTermCreate.Slug
	} else {
		//fmt.Println("total", total)
		//fmt.Println("totalPages", totalPages)
		//fmt.Println("isLastPage", isLastPage)
		//fmt.Println("ProductAttributeTermAll", ProductAttributeTermAll)
		tecalAttrId = items[0].ID
		tecalAttrName = items[0].Name
		tecalAttrSlug = items[0].Slug
	}
	return tecalAttrId, tecalAttrName, tecalAttrSlug
}

// Получение значение из файла
func DataFile(filename string) (string, error) {
	// Открыть файл
	fileToken, errorToken := os.Open(filename)
	if errorToken != nil {
		return "", errorToken
	}

	// Прочитать значение файла
	data := make([]byte, 64)
	n, err := fileToken.Read(data)
	if err == io.EOF { // если конец файла
		return "", errorToken
	}
	fileToken.Close() // Закрытие файла

	return string(data[:n]), nil
}
