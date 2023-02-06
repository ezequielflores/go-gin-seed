package pokemon

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/redbeestudios/go-seed/pkg"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/redbeestudios/go-seed/internal/application/port/in"
)

type PathParam struct {
	Name string `uri:"name"  binding:"required,alphanum"`
}

type PokemonController struct {
	getPokemonByName in.GetByName
	savePokemon      in.SavePokemon
}

func NewPokemonController(
	getPokemonByName in.GetByName,
	savePokemon in.SavePokemon,
) *PokemonController {
	return &PokemonController{
		getPokemonByName: getPokemonByName,
		savePokemon:      savePokemon,
	}
}

func getErrorDetail(error error) (int, string) {
	switch errorType := error.(type) {
	case pkg.BadRequestException:
		return http.StatusBadRequest, errorType.Msj
	case pkg.NotFoundException:
		return http.StatusNotFound, errorType.Msj
	case pkg.BadGatewayException:
		return http.StatusBadGateway, errorType.Msj
	default:
		return http.StatusInternalServerError, errorType.Error()
	}
}

func (c *PokemonController) GetPokemon(context *gin.Context) {

	name := &PathParam{}
	err := context.ShouldBindUri(name)

	log.Printf("Request parameter Name: %s", name.Name)

	if err != nil {
		log.Printf("Invalid param request %s", err)
		context.JSON(http.StatusBadRequest, "Invalid Pokemon Name")
		return
	}

	pokemon, err := c.getPokemonByName.Get(context, name.Name)

	if err != nil {
		log.Printf("ERROR controller - GetPokemon: %s", err.Error())
		status, msj := getErrorDetail(err)
		context.JSON(status, msj)
		return
	}

	context.Header("custom-header", "hello word")
	context.JSON(http.StatusOK, fromDomain(pokemon))
}

func (c *PokemonController) DumpPokemons(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	go func() {
		start := time.Now()
		for i := 1; i <= 906; i++ {
			c.retrieveAndSavePokemon(response, ctx, i, true)
		}

		log.Printf("Execution Finalized, elapsed time: %s", time.Since(start))
	}()

}

func (c *PokemonController) DumpPokemonsWithGoRoutines(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	var times [906]int64

	var totalTime int64 = 0

	var wg sync.WaitGroup

	start := time.Now()
	for i := 1; i < 906; i++ {
		index := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.retrieveAndSavePokemon(response, ctx, index, false)
			times[index] = time.Since(start).Milliseconds()
		}()
	}

	wg.Wait()
	for _, eachTime := range times {
		totalTime = totalTime + eachTime
	}
	log.Printf("Execution Finalized, avg time: %d milliseconds", totalTime/905)

}

func (c *PokemonController) retrieveAndSavePokemon(response http.ResponseWriter, ctx context.Context, i int, logProcessingPokemon bool) {
	pokemon, err := c.getPokemonByName.Get(ctx, strconv.Itoa(i))
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	if logProcessingPokemon {
		log.Println("Processing: " + strconv.Itoa(pokemon.Id()) + " - " + pokemon.Name())
	}

	if pokemon == nil {
		http.Error(response, err.Error(), http.StatusNotFound)
		return
	}

	_ = c.savePokemon.Save(ctx, pokemon)
}
