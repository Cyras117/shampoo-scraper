package model

import "go.mongodb.org/mongo-driver/bson/primitive"

/*
	TODO Add last read,total chapters
*/

type Manga struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Title         string             `bson:"title"`
	Path          string             `bson:"path"`
	SiteURL       string             `bson:"siteUrl"`
	AlternateLink string             `bson:"alternatreLink"`
	ImgURL        string             `bson:"imgUrl"`
	CurrentCh     float64            `bson:"currentCh"`
}
