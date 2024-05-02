package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CurrencyResponse struct {
	Time      map[string]interface{} `json:"time"`
	ChartName string                 `json:"chartName"`
	USD       map[string]interface{} `json:"bpi"`
}

func main() {
	var currencyResponse CurrencyResponse
	url := fmt.Sprintf("https://api.coindesk.com/v2/bpi/currentprice.json")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&currencyResponse)
	if err != nil {
		fmt.Println("Ошибка при декодировании JSON:", err)
		return
	}
	a := currencyResponse.Time
	timeData, ok := a["updated"].(string)
	if !ok {
		fmt.Println("Значение не является map[string]int")
		return
	}
	parsedTime, err := time.Parse("Jan 02, 2006 15:04:05 UTC", timeData)
	if err != nil {
		fmt.Println("Ошибка при парсинге времени:", err)
		return
	}
	parsedTime = parsedTime.Add(5 * time.Hour)
	fmt.Println(parsedTime)

	for _, val := range currencyResponse.USD {
		mapData, ok := val.(map[string]interface{})
		if !ok {
			fmt.Println("Значение не является map[string]int")
			return
		}
		btcToUzs, ok := mapData["rate_float"].(float64)
		if !ok {
			fmt.Println("Значение не является string")
			return
		}

		uzs := btcToUzs * 12500.0
		strUzs := fmt.Sprintf("%.2f", uzs)
		fmt.Println(mapData["code"], "---->", mapData["rate_float"], "=", strUzs, "UZS")
	}
}
