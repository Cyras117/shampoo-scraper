package model

/*
	TODO Add last read,total chapters
*/

type Manga struct {
	Title         string
	Path          string
	SiteURL       string
	AlternateLink string
	ImgURL        string
	CurrentCh     float64
}
