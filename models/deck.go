package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Deck struct {
	gorm.Model
	ID   int
	Deck datatypes.JSON
}

// create a card deck
func CreateDeck(db *gorm.DB, Deck *Deck) (err error) {
	err = db.Create(Deck).Error
	if err != nil {
		return err
	}
	return nil
}

// get decks
func GetDecks(db *gorm.DB, Deck *[]Deck) (err error) {
	err = db.Find(Deck).Error
	if err != nil {
		return err
	}
	return nil
}
