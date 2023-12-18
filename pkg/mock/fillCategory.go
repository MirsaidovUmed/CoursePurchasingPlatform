package mock

import (
	"Udemy/pkg"
	"Udemy/pkg/models"
)

func FillCategories() {
	pkg.Categories = append(pkg.Categories, models.Category{
		Name: "IT",
	})

	pkg.Categories = append(pkg.Categories, models.Category{
		Name: "Cooking",
	})

	pkg.Categories = append(pkg.Categories, models.Category{
		Name: "Design",
	})

	pkg.Categories = append(pkg.Categories, models.Category{
		Name: "Marketing",
	})
}
