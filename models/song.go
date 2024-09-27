package models

type Song struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Group       string `json:"group"`
	Song        string `json:"song"`
	Text        string `json:"text"`
	Link        string `json:"link"`
	ReleaseDate string `json:"releaseDate"`
}
