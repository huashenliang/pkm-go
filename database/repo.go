package database

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/huashenliang/pkm-go/models"
	"gorm.io/gorm"
)

type Repo struct {
	Db *gorm.DB
}

// &
func New() *Repo {
	db := InitDb()
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Deck{})
	return &Repo{Db: db}
}

// function to save deck
func (repository *Repo) SaveDeck(c *gin.Context) {
	var deck models.Deck
	c.BindJSON(&deck)
	err := models.CreateDeck(repository.Db, &deck)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, deck)
}

// function to get deck list
func (repository *Repo) GetDeckList(c *gin.Context) {
	var deck []models.Deck
	err := models.GetDecks(repository.Db, &deck)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, deck)
}

// get users
func (repository *Repo) GetUsers(c *gin.Context) {
	var user []models.User
	err := models.GetUsers(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

// get user by id
func (repository *Repo) GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	err := models.GetUser(repository.Db, &user, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}
