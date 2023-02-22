package woocommerce

import (
	"io"
	"os"
	"testing"
)

// Тестируем добавление категории
func TestAddCat_WC(t *testing.T) {

	consumer_key, _ := dataFile("consumer_key") //  Пользовательский ключ
	secret_key, _ := dataFile("secret_key")     // Секретный код пользователя

	// Авторизация
	userWC, _ := New(consumer_key, secret_key)
	cat := MeCat{
		Name: "Кофты",
		Slug: "kofta",
	}
	ParentID, ParentError := userWC.AddCat_WC(cat)
	t.Log(ParentID)
	if ParentError != nil {
		t.Error(ParentError)
	}

	cat = MeCat{
		Name:     "Кофты_подкатегория",
		Slug:     "kofta_child",
		ParentID: ParentID,
	}
	t.Log(userWC.AddCat_WC(cat))

	// Должно вывестись два числа.
	// Первое - ID категории, второе - ID подкатегории
}

// Получение значение из файла
func dataFile(filename string) (string, error) {
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
