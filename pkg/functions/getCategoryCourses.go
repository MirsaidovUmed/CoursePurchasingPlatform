package functions

import (
	pkg "Udemy/pkg/database"
	"fmt"
)

func GetCategoryCourses() {
	var (
		category string
		has      bool
	)
	fmt.Println("Введите категорию")
	fmt.Scan(&category)

	for _, val := range pkg.Category {
		if category == val.Name {
			has = true
		}
	}

	if has {
		fmt.Println("Курсы по категории:")
		for _, val := range pkg.Courses {
			if val.Category.Name == category {
				fmt.Println(val.Name)
			}
		}
	} else {
		fmt.Println("Ошибка!, данного пользователя нет в нашей бд")
	}
}
