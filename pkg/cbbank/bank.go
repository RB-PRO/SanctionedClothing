package cbbank

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Bank struct {
	Date         time.Time `json:"Date"`
	PreviousDate time.Time `json:"PreviousDate"`
	PreviousURL  string    `json:"PreviousURL"`
	Timestamp    time.Time `json:"Timestamp"`
	Valute       struct {
		Usd struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"USD"`
		Eur struct {
			ID       string  `json:"ID"`
			NumCode  string  `json:"NumCode"`
			CharCode string  `json:"CharCode"`
			Nominal  int     `json:"Nominal"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
			Previous float64 `json:"Previous"`
		} `json:"EUR"`
	} `json:"Valute"`
}

// Получить коээфициент USD/RUB
func USD() float64 {
	resp, err := http.Get("https://www.cbr-xml-daily.ru/daily_json.js")
	if err != nil {
		return 0.0
	}

	body, errBody := io.ReadAll(resp.Body)
	if errBody != nil {
		return 0.0
	}

	// Получить ответ
	var bankResp Bank
	errUnmarshal := json.Unmarshal(body, &bankResp)
	if errUnmarshal != nil { // Если ошибка распарсивания в структуру данных
		return 0.0
	}

	return bankResp.Valute.Usd.Value
}
