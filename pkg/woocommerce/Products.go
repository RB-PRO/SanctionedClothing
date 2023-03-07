// Файл, который поможет Вам создать товар.
package woocommerce

import (
	"encoding/json"
	"errors"

	"github.com/RB-PRO/SanctionedClothing/pkg/bases"
)

// Структура добавления товара
type ProductWC struct {
	Name             string `json:"name"`
	Type             string `json:"type"`
	RegularPrice     string `json:"regular_price"`
	Description      string `json:"description"`
	ShortDescription string `json:"short_description"`
	Categories       []struct {
		ID int `json:"id"`
	} `json:"categories"`
	Tags []struct { // List of tags. See Product - Tags properties - http://woocommerce.github.io/woocommerce-rest-api-docs/?shell#product-tags-properties
		ID int `json:"id"`
	} `json:"tag"`

	Images []struct {
		Src string `json:"src"`
	} `json:"images"`

	// Если ошибка
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Status int `json:"status"`
		Params struct {
			Type string `json:"type"`
		} `json:"params"`
		Details struct {
			Type struct {
				Code    string      `json:"code"`
				Message string      `json:"message"`
				Data    interface{} `json:"data"`
			} `json:"type"`
		} `json:"details"`
	} `json:"data"`
}

// Метод [products] поможет Вам добавить товар
//
// # Использую для добавления товара
// Передаю внутрь структуру загрузки для WC ProductWC
//
// [products]: http://woocommerce.github.io/woocommerce-rest-api-docs/?shell#create-a-product
func (user *User) AddProduct_WC(ProdWC ProductWC) error {

	// Сделать json добавления категории
	bytesData, errMarshal := json.Marshal(ProdWC)
	if errMarshal != nil {
		return errMarshal
	}

	// Выполнить запрос
	bodyBytes, errData := user.quering("POST", "/products", bytesData)
	if errData != nil {
		return errData
	}

	// Получить ответ
	var ProductWC_Resp ProductWC
	errUnmarshal := json.Unmarshal(bodyBytes, &ProductWC_Resp)
	if errUnmarshal != nil { // Если ошибка распарсивания в структуру данных
		return errors.New("AddProduct_WC: Не удалось распарсить ответ сервера: " + string(bodyBytes))
	}

	//fmt.Println(string(bodyBytes))

	// Если всё верно сработало и произошло добавление
	return nil
}

// Перевести базовую структуру товара в структу запроса для woocommerce
func Product2ProductWC(prod bases.Product2, CatIDcreate, tagId int) (prodWC ProductWC) {
	prodWC.Name = prod.Name                   // Назвние товара
	prodWC.ShortDescription = prod.FullName   // краткое описание товара
	prodWC.Description = prod.Description.Rus // Описакние товара на Русском
	prodWC.Type = "variable"                  // simple, grouped, external, variable и woosb.

	// Категория
	prodWC.Categories = []struct {
		ID int "json:\"id\""
	}{{CatIDcreate}}

	// Метка
	prodWC.Tags = []struct {
		ID int "json:\"id\""
	}{{tagId}}

	return prodWC
}
