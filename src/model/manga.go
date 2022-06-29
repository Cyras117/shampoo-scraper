package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Manga struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Title         string             `bson:"title"`
	Path          string             `bson:"path"`
	SiteURL       string             `bson:"siteUrl"`
	AlternateLink string             `bson:"alternatreLink"`
	ImgURL        string             `bson:"imgUrl"`
	CurrentCh     float64            `bson:"currentCh"`
	LastReadCh    float64            `bson:"lastReadCh"`
}
