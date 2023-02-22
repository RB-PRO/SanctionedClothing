package woocommerce

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// Набор констант базового набора запроса
const (
	URL = "https://clikshop.ru" // Сам сайт
	REQ = "/wp-json/wc/v3"      // API-шлюз
)

// Структура пользователя API
type User struct {
	consumer_key string // Пользовательский ключ
	secret_key   string // Секретный код пользователя
}

type Orders struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Status int `json:"status"`
	} `json:"data"`
}

// Авторизация пользователя "https://clikshop.ru/wp-json/wc/v3/orders"
func New(consumer_key, secret_key string) (*User, error) {
	return &User{consumer_key, secret_key}, nil
}

// Метод [Orders] помогает вам просматривать все заказы
//
// # Использую для проверки авторизации
//
// [Orders]: http://woocommerce.github.io/woocommerce-rest-api-docs/?shell#orders
func (user *User) IsOrder() error {
	bodyBytes, errData := user.quering("GET", "/orders", nil)
	if errData != nil {
		return errData
	}

	var order Orders
	errUnmarshal := json.Unmarshal(bodyBytes, &order)
	if errUnmarshal == nil { // Если ошибка
		return errors.New(order.Code + " - " + order.Message)
	}

	// Если всё верно сработало
	return nil
}

// Ядро запроса с входными параметрами:
// - methodURL - Метод GET, POST, PUT, ...
// - methodSite - Метод API
// - data - Массив byte с передаваемыми данным
func (user *User) quering(methodURL, methodApi string, data []byte) ([]byte, error) {
	client := &http.Client{}

	req, errReq := http.NewRequest(methodURL, URL+REQ+methodApi, bytes.NewBuffer(data))
	if errReq != nil {
		return nil, errReq
	}

	req.Header.Add("Authorization", "Basic "+basicAuth(user.consumer_key, user.secret_key))
	if data != nil {
		req.Header.Add("Content-Type", "application/json")
	}
	res, errRes := client.Do(req)
	if errRes != nil {
		return nil, errRes
	}
	defer res.Body.Close()

	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		return nil, errBody
	}

	return body, nil
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
