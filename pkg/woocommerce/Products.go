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
	Images []struct {
		Src string `json:"src"`
	} `json:"images"`
}

// Структура добавления товара
type ProductWC_Response struct {
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
	var ProductWC_Resp ProductWC_Response
	errUnmarshal := json.Unmarshal(bodyBytes, &ProductWC_Resp)
	if errUnmarshal != nil { // Если ошибка распарсивания в структуру данных
		return errors.New("AddProduct_WC: Не удалось распарсить ответ сервера: " + string(bodyBytes))
	}

	// Если всё верно сработало и произошло добавление
	return nil
}
func Product2ProductWC(prod bases.Product2) (prodWC ProductWC) {

	return prodWC
}
