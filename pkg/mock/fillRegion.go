package mock

import (
	pkg "Udemy/pkg/database"
	"Udemy/pkg/models"
)

func FillRegion() {
	pkg.Region = append(pkg.Region, models.Region{
		Name: "RRP",
	})

	pkg.Region = append(pkg.Region, models.Region{
		Name: "Согдийская область",
	})

	pkg.Region = append(pkg.Region, models.Region{
		Name: "Хатлонская область",
	})
}
