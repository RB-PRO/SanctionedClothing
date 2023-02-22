package woocommerce

import (
	"encoding/json"
	"errors"
)

// Структура ответа API на создание категории
type AddCatResponse struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Slug        string      `json:"slug"`
	Parent      int         `json:"parent"`
	Description string      `json:"description"`
	Display     string      `json:"display"`
	Image       interface{} `json:"image"`
	MenuOrder   int         `json:"menu_order"`
	Count       int         `json:"count"`
	Links       struct {
		Self []struct {
			Href string `json:"href"`
		} `json:"self"`
		Collection []struct {
			Href string `json:"href"`
		} `json:"collection"`
	} `json:"_links"`

	// Если ошибка, то заполняются эти данные:
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Status     int `json:"status"`
		ResourceID int `json:"resource_id"`
	} `json:"data"`
}

// Метод [create-a-product-category] поможет Вам добавить категорию
//
// # Использую для добавления категории товаров
// Возвращает ID категории
//
// [create-a-product-category]: http://woocommerce.github.io/woocommerce-rest-api-docs/?shell#create-a-product-category
func (user *User) AddCat_WC(valetCat MeCat) (int, error) {

	// Сделать json добавления категории
	bytesData, errMarshal := json.Marshal(valetCat)
	if errMarshal != nil {
		return 0, errMarshal
	}

	// Выполнить запрос
	bodyBytes, errData := user.quering("POST", "/products/categories", bytesData)
	if errData != nil {
		return 0, errData
	}

	// Получить ответ
	var AddCatRes AddCatResponse
	errUnmarshal := json.Unmarshal(bodyBytes, &AddCatRes)
	if errUnmarshal != nil { // Если ошибка распарсивания в структуру данных
		return 0, errors.New("AddCat_WC: Не удалось распарсить ответ сервера: " + string(bodyBytes))
	}
	if AddCatRes.Code == "term_exists" { // Обработка случая, когда существует такая категория
		return AddCatRes.Data.ResourceID, nil
	}

	// Если всё верно сработало и произошло добавление
	return AddCatRes.ID, nil
}
