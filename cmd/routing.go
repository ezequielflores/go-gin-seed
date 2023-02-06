package cmd

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoutes(dependencies *Dependencies) *gin.Engine {
	router := gin.Default()

	router.GET("/pokemon/:name", dependencies.PokemonController.GetPokemon)
	/*router.POST("/dumpPokemons", dependencies.PokemonController.DumpPokemons)
	router.POST("/dumpPokemonsGoRoutine", dependencies.PokemonController.DumpPokemonsWithGoRoutines)*/
	http.Handle("/", router)

	return router
}
