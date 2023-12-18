package services

import "Udemy/pkg"

func GetUserExpensesByUserIndex(index int) float64 {
	user := pkg.Users[index]

	var expenses float64

	for _, v := range pkg.Subscriptions {
		if user == v.User {
			expenses += v.Price
		}
	}

	return expenses
}
