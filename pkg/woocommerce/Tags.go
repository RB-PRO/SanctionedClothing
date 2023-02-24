// Файл, который поможет Вам создать товар.
package woocommerce

import (
	"encoding/json"
	"errors"
	"strings"
)

// Структура тэгов
type Tag struct {
	Name string
	Slug string
	Id   int
}

// Структура списка Тэгов
type TagsWC_Response []struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Count       int    `json:"count"`
	Links       struct {
		Self []struct {
			Href string `json:"href"`
		} `json:"self"`
		Collection []struct {
			Href string `json:"href"`
		} `json:"collection"`
	} `json:"_links"`
}

// Метод [products] поможет Вам добавить товар
//
// # Использую для добавления товара
// Передаю внутрь структуру загрузки для WC ProductWC
//
// [products]: http://woocommerce.github.io/woocommerce-rest-api-docs/?shell#create-a-product
func (user *User) AllTags_WC() ([]Tag, error) {
	var tags []Tag

	// Выполнить запрос
	bodyBytes, errData := user.quering("GET", "/products/tags?per_page=100", nil)
	if errData != nil {
		return nil, errData
	}

	// Получить ответ
	var TagsWC_Resp TagsWC_Response
	errUnmarshal := json.Unmarshal(bodyBytes, &TagsWC_Resp)
	if errUnmarshal != nil { // Если ошибка распарсивания в структуру данных
		return nil, errors.New("AllTags_WC: Не удалось распарсить ответ сервера: " + string(bodyBytes))
	}

	// Преобразуем в свою структуру
	for _, valueTags := range TagsWC_Resp {
		tags = append(tags, Tag{
			Name: valueTags.Name,
			Slug: valueTags.Slug,
			Id:   valueTags.ID,
		})
	}

	return tags, nil
}

// Поиск в массиве меток для параметра slug
// Пробелы заменяются на _
// Всё в нижнем регистре
func FindIdTagSlug(tags []Tag, slug string) int {
	slug = strings.ReplaceAll(slug, "  ", " ")
	slug = strings.ReplaceAll(slug, " ", "_")
	slug = strings.ToLower(slug)
	for _, val := range tags {
		if slug == val.Slug {
			return val.Id
		}
	}
	return 0
}
