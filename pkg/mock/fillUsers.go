package mock

import (
	"Udemy/pkg"
	"Udemy/pkg/models"
)

func FillUsers() {
	pkg.Users = append(pkg.Users, models.User{
		Name:      "Jacob Akhmedov",
		BirthYear: 2004,
		Balance:   1000,
		City:      pkg.Cities[1],
		Address:   "улица хз, дом хз, квартира хз",
		Mail:      "jacob@akhmedov.tech",
		Phone:     "928560139",
	})

	pkg.Users = append(pkg.Users, models.User{
		Name:      "Maksim Gorkiy",
		BirthYear: 1868,
		Balance:   2500,
		City:      pkg.Cities[3],
		Address:   "улица Максима Горького, дом 3, квартира 27",
		Mail:      "maksim@gorkiy.tech",
		Phone:     "889300475",
	})
}
