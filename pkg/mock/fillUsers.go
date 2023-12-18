package mock

import (
	pkg "Udemy/pkg/database"
	"Udemy/pkg/models"
)

func FillUsers() {
	pkg.User = append(pkg.User, models.User{
		Name:      "Oskar",
		BirthYear: 2004,
		Balance:   100000,
		City:      pkg.City[0],
		Address:   "someAddress",
		Mail:      "someMail.com",
		Phone:     "+992565656",
	})

	pkg.User = append(pkg.User, models.User{
		Name:      "someName",
		BirthYear: 1999,
		Balance:   100000,
		City:      pkg.City[2],
		Address:   "someOtherAddress",
		Mail:      "someMail.com",
		Phone:     "+992565656",
	})

	pkg.User = append(pkg.User, models.User{
		Name:      "someOtherName",
		BirthYear: 2004,
		Balance:   100000,
		City:      pkg.City[3],
		Address:   "someOtherOtherAddress",
		Mail:      "someMail.com",
		Phone:     "+992565656",
	})
}
