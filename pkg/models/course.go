package models

type Course struct {
	Name     string
	Price    float64
	Category Category
	Duration int64
}
