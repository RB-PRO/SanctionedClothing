package woocommerce

import (
	"encoding/json"
	"errors"
)

type Attributes struct {
	Attribute []ProductListAttributes
}

type ProductListAttributes struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Type        string `json:"type"`
	OrderBy     string `json:"order_by"`
	HasArchives bool   `json:"has_archives"`
	Links       struct {
		Self []struct {
			Href string `json:"href"`
		} `json:"self"`
		Collection []struct {
			Href string `json:"href"`
		} `json:"collection"`
	} `json:"_links"`
}

// Метод [product/categories] позволяет Вам извлекать все аттрибуты продуктов.
//
// # Использую для создания структуры всех категорий
//
// [Orders]: http://woocommerce.github.io/woocommerce-rest-api-docs/?shell#list-all-product-categories
func (user *User) ProductsAttributes() (Attributes, error) {
	// Структура по категории
	var attrib Attributes
	//var TotalPages int = 2

	for i, TotalPages := 1, 2; i <= TotalPages; i++ {
		var categ Attributes
		var bodyBytes []byte
		var errData error

		bodyBytes, TotalPages, errData = user.queringProductsCategories("GET", "/products/attributes", i)
		if errData != nil {
			return Attributes{}, errData
		}

		// Получить ответ
		errUnmarshal := json.Unmarshal(bodyBytes, &categ)
		if errUnmarshal != nil { // Если ошибка распарсивания в структуру данных
			return Attributes{}, errors.New("ProductsAttributes: Не удалось распарсить ответ сервера: " + string(bodyBytes))
		}

		attrib.Attribute = append(attrib.Attribute, categ.Attribute...)

	}

	// Если всё верно сработало
	return attrib, nil
}
