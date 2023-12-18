package functions

import (
	pkg "Udemy/pkg/database"
	"fmt"
)

func GetUsersExpense() {
	var (
		name    string
		has     bool
		balance int
	)
	fmt.Println("Введите имя студента")
	fmt.Scan(&name)

	for _, course := range pkg.Subscription {
		if course.User.Name == name {
			balance += course.Price
			has = true
		}
	}

	if has {
		fmt.Println("Сумма которую потратил студент", balance)
	} else {
		fmt.Println("Ошибка!, данного пользователя нет в нашей бд")
	}
}
