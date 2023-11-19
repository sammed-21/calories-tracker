package models

import (
	"go.mongodb.org/mongo-driver/bson/prmitive"
)

type Entry struct {
	ID      prmitive.ObjectId 	`bson:"id"`
	Dish	*string				`json:"dish"`
	Fat      *float64 				`json:"fat"`
	Ingredients	*string 			`json:"ingredients"`
	Calories 	*string				`json:"calories"`

}