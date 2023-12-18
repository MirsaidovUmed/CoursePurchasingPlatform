package mock

import (
	"Udemy/pkg"
	"Udemy/pkg/models"
)

func FillCities() {
	pkg.Cities = append(pkg.Cities, models.City{
		Name:   "Душанбе",
		Region: pkg.Regions[2],
	})

	pkg.Cities = append(pkg.Cities, models.City{
		Name:   "Худжанд",
		Region: pkg.Regions[1],
	})

	pkg.Cities = append(pkg.Cities, models.City{
		Name:   "Куляб",
		Region: pkg.Regions[0],
	})

	pkg.Cities = append(pkg.Cities, models.City{
		Name:   "Хорог",
		Region: pkg.Regions[3],
	})
}
