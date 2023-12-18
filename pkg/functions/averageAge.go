package functions

import (
	pkg "Udemy/pkg/database"
	"fmt"
)

func AverageAge() {
	var (
		average int
		course  string
		has     bool
	)
	fmt.Println("Введите курс")
	fmt.Scan(&course)

	for _, val := range pkg.Courses {
		if course == val.Name {
			has = true
		}
	}

	if has {
		fmt.Println("Средний возраст студентов")
		for _, val := range pkg.User {
			average += 2023 - val.BirthYear
		}
		fmt.Println(average / len(pkg.User))
	} else {
		fmt.Println("Ошибка!, данного курса нет в нашей бд")
	}
}
