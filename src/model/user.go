package model

import "go.mongodb.org/mongo-driver/bson/primitive"

/*
TODO:Delete this
User{
	id:xxxxxx,
	name:xxxxx,
	colections:[
		list1[{manga1},{manga2}],
		list2[{manga1},{manga2}]
	]
}
*/

type User struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Lists []List             `bson:"lists"`
}
