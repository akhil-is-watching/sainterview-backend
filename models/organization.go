package models

type Organization struct {
	ID         string `gorm:"primaryKey" json:"id"`
	Name       string `json:"name"`
	Email      string `gorm:"unique" json:"email"`
	Password   string
	Credits    int         `json:"credits"`
	Interviews []Interview `json:"interviews"`
}
