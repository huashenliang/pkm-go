package controllers

import (
	"net/http"

	tcg "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"
	"github.com/gin-gonic/gin"
	pkmService "github.com/huashenliang/pkm-go/services"
)

type GenerateDeckReq struct {
	PokemonType string
	DeckName    string
}

// function to get a list of pokemon types
func GetPokemonTypes(c *gin.Context, pokemonClient tcg.Client) {
	typeList := pkmService.GetPokemonTypes(pokemonClient)

	c.JSON(http.StatusOK, typeList)
}

// function to generate a pokemon deck by type
func GenerateDeck(c *gin.Context, pokemonClient tcg.Client) {
	var pokemonType GenerateDeckReq
	// Call BindJSON to bind the received JSON to pokemonType
	if err := c.BindJSON(&pokemonType); err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(http.StatusOK, pkmService.GenerateDeck(pokemonClient, pokemonType.PokemonType, pokemonType.DeckName))
}
