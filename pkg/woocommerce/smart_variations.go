package woocommerce

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// https://www.smart-variations.com/knowledge-base/api/post/
type SmartImagePost struct {
	Data []struct {
		Slugs      []string `json:"slugs"`
		Imgs       []string `json:"imgs"`
		LoopHidden bool     `json:"loop_hidden"`
	} `json:"data"`
}

// https://www.smart-variations.com/knowledge-base/api/post/
type SmartImagePostResponse struct {
	Status bool `json:"response"`
}

func (user *User) PostSmartImage(product_id int) error {
	//url := "/wp-json/wc/svi/" + strconv.Itoa(product_id)

	// Выполнить запрос
	bodyBytes, errData := user.queringSmartImage("POST", strconv.Itoa(product_id), nil)
	if errData != nil {
		return errData
	}

	fmt.Println("bodyBytes", string(bodyBytes))

	// Получить ответ
	var SmartImagePostResp SmartImagePostResponse
	errUnmarshal := json.Unmarshal(bodyBytes, &SmartImagePostResp)
	if errUnmarshal != nil { // Если ошибка распарсивания в структуру данных
		return errors.New("PostSmartImage: Не удалось распарсить ответ сервера: " + string(bodyBytes))
	}

	return nil
}

// Ядро запроса с входными параметрами:
// - methodURL - Метод GET, POST, PUT, ...
// - methodSite - Метод API
// - data - Массив byte с передаваемыми данным
func (user *User) queringSmartImage(methodURL, product_id string, data []byte) ([]byte, error) {
	client := &http.Client{}

	req, errReq := http.NewRequest(methodURL, URL+"/wp-json/wc/svi/"+product_id, bytes.NewBuffer(data))
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
