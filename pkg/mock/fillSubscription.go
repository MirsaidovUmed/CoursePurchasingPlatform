package mock

import (
	pkg "Udemy/pkg/database"
	"Udemy/pkg/models"
)

func FillSubscriptions() {
	for i := range pkg.User {
		if i < len(pkg.Courses) {
			pkg.User[i].Balance -= pkg.Courses[i].Price

			subscription := models.Subscription{
				User:   pkg.User[i],
				Course: pkg.Courses[i],
				Price:  pkg.Courses[i].Price,
			}
			pkg.Subscription = append(pkg.Subscription, subscription)
		}
	}
}
