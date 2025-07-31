package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ViaCepResponse struct {
	Localidade string `json:"localidade"`
}

func BuscarCidadePorCEP(cep string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
	if err != nil || resp.StatusCode != 200 {
		return "", errors.New("erro ao consultar CEP")
	}
	defer resp.Body.Close()

	var data ViaCepResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil || data.Localidade == "" {
		return "", errors.New("cidade não encontrada")
	}

	return data.Localidade, nil
}
