package mock

import (
	"Udemy/pkg"
	"Udemy/pkg/models"
)

func FillSubscriptions() {
	pkg.Subscriptions = append(pkg.Subscriptions, models.Subscription{
		User:   pkg.Users[1],
		Course: pkg.Courses[0],
		Price:  pkg.Courses[0].Price,
	})

	pkg.Subscriptions = append(pkg.Subscriptions, models.Subscription{
		User:   pkg.Users[0],
		Course: pkg.Courses[1],
		Price:  pkg.Courses[1].Price,
	})

	pkg.Subscriptions = append(pkg.Subscriptions, models.Subscription{
		User:   pkg.Users[0],
		Course: pkg.Courses[3],
		Price:  pkg.Courses[3].Price,
	})

	pkg.Subscriptions = append(pkg.Subscriptions, models.Subscription{
		User:   pkg.Users[1],
		Course: pkg.Courses[2],
		Price:  pkg.Courses[2].Price,
	})
}
