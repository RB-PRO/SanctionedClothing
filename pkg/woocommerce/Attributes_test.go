package woocommerce_test

import (
	"io"
	"os"
	"testing"

	"github.com/RB-PRO/SanctionedClothing/pkg/woocommerce"
)

func TestProductsAttributes(t *testing.T) {

	consumer_key, _ := DataFile("consumer_key") //  Пользовательский ключ
	secret_key, _ := DataFile("secret_key")     // Секретный код пользователя
	//yandexToken, _ := DataFile("yandexToken")   // Секретный код пользователя

	// Авторизация
	userWC, _ := woocommerce.New(consumer_key, secret_key)

	attr, attrError := userWC.ProductsAttributes()
	if attrError != nil {
		t.Error(attrError)
	} else {
		for ind, val := range attr.Attribute {
			t.Log(ind, val.Name)
		}
	}

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
