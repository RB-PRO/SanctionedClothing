package woocommerce

import (
	"encoding/json"
	"errors"
	"fmt"
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
	var attrib Attributes // Структура аттрибутов

	for i, TotalPages := 1, 2; i <= TotalPages; i++ {
		var bodyBytes []byte
		var errData error

		bodyBytes, TotalPages, errData = user.queringProductsCategories("GET", "/products/attributes", i)
		if errData != nil {
			return Attributes{}, errData
		}
		fmt.Println("TotalPages", TotalPages)

		// Распарсим входную инормацию
		var Attribute []ProductListAttributes
		errUnmarshal := json.Unmarshal(bodyBytes, &Attribute)
		if errUnmarshal != nil { // Если ошибка распарсивания в структуру данных
			return Attributes{}, errors.New("ProductsAttributes: Не удалось распарсить ответ сервера: " + string(bodyBytes))
		}
		attrib.Attribute = append(attrib.Attribute, Attribute...)

	}

	// Если всё верно сработало
	return attrib, nil
}

// Поиск ID аттрибута по имени
//
//	Attributes.name -> Attributes.id
func (attr Attributes) Find_id_of_name(name string) (int, error) {
	for _, value := range attr.Attribute {
		if value.Name == name {
			return value.ID, nil
		}
	}
	return 0, errors.New("Find_id_of_name: Не найден аттрибут с таким именем (Name)")
}

// Поиск ID аттрибута по ссылке
//
//	Attributes.slug -> Attributes.id
func (attr Attributes) Find_id_of_slug(slug string) (int, error) {
	for _, value := range attr.Attribute {
		if value.Slug == "pa_"+slug {
			return value.ID, nil
		}
	}
	return 0, errors.New("Find_id_of_name: Не найден аттрибут с такой ссылкой (Slug) " + slug)
}
