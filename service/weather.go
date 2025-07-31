package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type WeatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func BuscarTemperatura(cidade string) (float64, error) {
	key := os.Getenv("WEATHER_API_KEY")
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&lang=pt", key, cidade)

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return 0, fmt.Errorf("erro na weather api")
	}
	defer resp.Body.Close()

	var data WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, err
	}
	return data.Current.TempC, nil
}
