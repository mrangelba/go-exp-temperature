package viacep

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mrangelba/go-exp-temperature/internal/domain"
	"github.com/mrangelba/go-exp-temperature/internal/domain/entity"
	"github.com/mrangelba/go-exp-temperature/pkg/error_handle"
)

var (
	BASE_URL = "https://viacep.com.br"
)

type gateway struct {
	client *http.Client
}

type cepResponse struct {
	Cep        string `json:"cep"`
	Localidade string `json:"localidade"`
	Erro       bool   `json:"erro"`
}

func NewCepGateway(client *http.Client) domain.CEPGateway {
	return &gateway{
		client: client,
	}
}

func (g gateway) Get(ctx context.Context, cep string) (*entity.CEP, error) {
	var cepOutput cepResponse
	url := fmt.Sprintf("%s/ws/%s/json/", BASE_URL, cep)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, bytes.NewReader(nil))

	if err != nil {
		return nil, err
	}

	defer request.Body.Close()

	response, err := g.client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode == 200 {
		err = json.NewDecoder(response.Body).Decode(&cepOutput)

		if err != nil {
			return nil, err
		}

		if cepOutput.Erro {
			return nil, error_handle.ErrNotFound
		}

		return &entity.CEP{
			Cep:      cepOutput.Cep,
			CityName: cepOutput.Localidade,
		}, nil
	}

	return nil, error_handle.ErrUnprocessableEntity
}
