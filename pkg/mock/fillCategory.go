package mock

import (
	pkg "Udemy/pkg/database"
	"Udemy/pkg/models"
)

func FillCategory() {
	pkg.Category = append(pkg.Category, models.Category{
		Name: "Language",
	})

	pkg.Category = append(pkg.Category, models.Category{
		Name: "IT",
	})
}
