package functions

import (
	pkg "Udemy/pkg/database"
	"fmt"
)

func GetUserCourses() {
	var (
		name string
		has  bool
	)
	fmt.Println("Введите имя студента")
	fmt.Scan(&name)

	for _, val := range pkg.User {
		if name == val.Name {
			has = true
		}
	}

	if has {
		fmt.Println("Курсы на которые подписан студент:")
		for _, val := range pkg.Subscription {
			if val.User.Name == name {
				fmt.Println(val.Course.Name)
			}
		}
	} else {
		fmt.Println("Ошибка!, данного пользователя нет в нашей бд")
	}
}
