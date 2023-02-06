package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/redbeestudios/go-seed/internal/application/model/pokemon"
	"github.com/redbeestudios/go-seed/internal/application/port/out"
	"github.com/redbeestudios/go-seed/pkg"
)

var _ out.PokemonRepository = (*PokemonRestAdapter)(nil)

type PokemonRestAdapter struct {
	client *resty.Client
}

func (a *PokemonRestAdapter) SavePokemon(ctx context.Context, pokemon *pokemon.Pokemon) error {
	//TODO implement me
	panic("implement me")
}

func NewPokemonRestAdapter(
	config pkg.RestClientConfig,
) *PokemonRestAdapter {

	client := resty.New().
		SetBaseURL(config.BaseUrl).
		SetTimeout(time.Duration(config.TimeoutMilliseconds) * time.Millisecond).
		SetRetryCount(config.RetryCount).
		SetRetryWaitTime(time.Duration(config.RetryWaitMilliseconds) * time.Millisecond)

	return &PokemonRestAdapter{
		client: client,
	}
}

func checkError(response *resty.Response) error {
	log.Printf("Status Response: %d - %s", response.StatusCode(), response.Status())
	switch response.StatusCode() {
	case 200:
		return nil
	case 404:
		return pkg.NotFoundException{Msj: "Not Found", Detail: "Pokemon Not Found"}
	case 502:
		return pkg.BadRequestException{Msj: "Bad Gateway", Detail: "Bad Gateway to Pokemon Api"}
	default:
		return fmt.Errorf("status response: %d - %s", response.StatusCode(), response.Status())
	}
}

func (a *PokemonRestAdapter) GetByName(
	ctx context.Context,
	name string,
) (*pokemon.Pokemon, error) {

	log.Printf("Call Pokemon Api with parameter: %s\n", name)

	response, err := a.client.
		R().
		SetContext(ctx).
		SetPathParam("name", name).
		Get("/pokemon/{name}")

	if err != nil {
		return nil, err
	}

	if errorResponse := checkError(response); errorResponse != nil {
		return nil, errorResponse
	}

	var responseObject = &pokemonResponse{}
	if err := json.Unmarshal(response.Body(), responseObject); err != nil {
		return nil, err
	}

	return responseObject.ToDomain()
}
