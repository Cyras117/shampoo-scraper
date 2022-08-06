package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type List struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
	//TODO check if put the manga itself is better
	Collection []string `bson:"Collection"`
}
