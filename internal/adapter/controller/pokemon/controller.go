package pokemon

import (
	"encoding/json"
	"net/http"

	"github.com/redbeestudios/go-seed/internal/application/port/in"
	"github.com/redbeestudios/go-seed/pkg"
)

type PokemonController struct {
	getPokemonByName in.GetPokemonByName
}

func NewPokemonController(getPokemonByName in.GetPokemonByName) *PokemonController {
	return &PokemonController{
		getPokemonByName: getPokemonByName,
	}
}

type pokemonResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func (c *PokemonController) GetPokemon(
	response http.ResponseWriter,
	request *http.Request,
) {
	name, err := pkg.GetStringFromPath("name", request)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	pokemon, err := c.getPokemonByName.Get(name)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	if pokemon == nil {
		http.Error(response, err.Error(), http.StatusNotFound)
		return
	}

	pokemonResponse := pokemonResponse{
		Id:   pokemon.Id(),
		Name: pokemon.Name(),
		Type: pokemon.PokemonType(),
	}

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(pokemonResponse)

}
