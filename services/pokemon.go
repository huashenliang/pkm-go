package services

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	tcg "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"
	"github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg/request"
)

type Deck struct {
	TYPE    string
	Name    string
	Pokemon []*tcg.PokemonCard
	Energy  []*tcg.PokemonCard
	Trainer []*tcg.PokemonCard
}

// Main entry to generate a deck
// Total 60 cards for a deck
func GenerateDeck(pokemonClient tcg.Client, pokemonType string, deckName string) Deck {
	var newDeck Deck

	// Set the deck type and name
	newDeck.TYPE = pokemonType
	newDeck.Name = deckName

	// Generating cards for pokemon with random number between 12 - 16
	numOfPkmCards := RandomNumInRange(12, 16)
	newDeck.Pokemon = GetRandomPokemonByType(pokemonClient, pokemonType, numOfPkmCards)

	// Generating 10 energy cards
	newDeck.Energy = GetEnegryCards(pokemonClient, pokemonType)

	// Generating the rest of cards as trainer cards
	numOfTrainerCards := 60 - numOfPkmCards
	newDeck.Trainer = GetTrainerCards(pokemonClient, numOfTrainerCards)

	return newDeck
}

// Get a list of pokemon typs
func GetPokemonTypes(pokemonClient tcg.Client) []string {
	pokemonTypes, err := pokemonClient.GetTypes()
	if err != nil {
		log.Fatal(err)
	}

	return pokemonTypes
}

// Get 12-16 pokemon of a given type
// Max page size is 250
func GetRandomPokemonByType(pokemonClient tcg.Client, pokemonType string, numOfPkmCards int) []*tcg.PokemonCard {
	totalCount := GetTotalCount(fmt.Sprintf("page=1&pageSize=1&q=types:%s", pokemonType))
	maxPage := totalCount / numOfPkmCards
	randPageNum := RandomNumInRange(1, maxPage)

	//optimize: query smaller page size but with multiple times
	allCards, err := pokemonClient.GetCards(
		request.Query(fmt.Sprintf("types:%s", pokemonType)),
		request.PageSize(numOfPkmCards),
		request.Page(randPageNum),
	)
	if err != nil {
		log.Fatal(err)
	}

	return allCards
}

// Get 10 enegry cards by type
func GetEnegryCards(pokemonClient tcg.Client, pokemonType string) []*tcg.PokemonCard {
	searchEnegry := fmt.Sprintf("%s energy", pokemonType)
	totalCount := GetTotalCount(fmt.Sprintf("page=1&pageSize=1&q=supertype:energy name:%s", strconv.Quote(searchEnegry)))
	maxPage := totalCount / 10
	randPageNum := RandomNumInRange(1, maxPage)

	energyCards, err := pokemonClient.GetCards(
		request.Query("supertype:energy", fmt.Sprintf("name:%s energy", pokemonType)),
		request.PageSize(10),
		request.Page(randPageNum),
	)
	if err != nil {
		log.Fatal(err)
	}

	return energyCards
}

// Get random trainer cards
func GetTrainerCards(pokemonClient tcg.Client, numOfTrainerCards int) []*tcg.PokemonCard {
	totalCount := GetTotalCount("page=1&pageSize=1&q=supertype:Trainer")
	maxPage := totalCount / numOfTrainerCards
	randPageNum := RandomNumInRange(1, maxPage)

	trainerCards, err := pokemonClient.GetCards(
		request.Query("supertype:Trainer"),
		request.PageSize(numOfTrainerCards),
		request.Page(randPageNum),
	)
	if err != nil {
		log.Fatal(err)
	}

	return trainerCards
}

// Quick-dirty hack for getting the totalCount from pokemon api
// They do not provide this attribute in their go sdk
// searchString -
//
//	e.g supertype:Trainer&page=1&pageSize=1
func GetTotalCount(searchString string) int {
	resp, err := http.Get(fmt.Sprintf("https://api.pokemontcg.io/v2/cards?%s", searchString))
	if err != nil {
		log.Fatal(err)
	}
	var generic map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&generic)
	if err != nil {
		log.Fatal(err)
	}

	if generic["totalCount"] == nil {
		return 0
	}

	return int(generic["totalCount"].(float64))
}

// Helper function to randomly pick from cards
func RandomPick(originData []*tcg.PokemonCard, totalPicks int) []*tcg.PokemonCard {
	var result []*tcg.PokemonCard
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= totalPicks; i++ {
		result = append(result, originData[rand.Intn(len(originData)-1)])
	}

	return result
}

func RandomNumInRange(min int, max int) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max-min+1) + min
}
