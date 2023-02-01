package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"

	tcg "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"
	"github.com/gin-gonic/gin"
	controllers "github.com/huashenliang/pkm-go/controllers"
	database "github.com/huashenliang/pkm-go/database"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	pokemonClient := tcg.NewClient(os.Getenv("POKEMON_API_KEY"))
	db := database.New()

	routers := setupRouter(pokemonClient, db)

	_ = routers.Run(":8080")
	log.Println("running")
}

func setupRouter(pokemonClient tcg.Client, db *database.Repo) *gin.Engine {
	routers := gin.Default()
	routers.Use(cors.Default())

	routers.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	routers.GET("/getPokemonTypes", func(c *gin.Context) {
		controllers.GetPokemonTypes(c, pokemonClient)
	})

	routers.GET("/getDeckList", db.GetDeckList)

	routers.POST("/generateDeck", func(c *gin.Context) {
		controllers.GenerateDeck(c, pokemonClient)
	})

	routers.POST("/saveDeck", db.SaveDeck)

	return routers
}
