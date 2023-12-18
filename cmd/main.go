package main

import (
	"Udemy/pkg/functions"
	pkg "Udemy/pkg/mock"
	"fmt"
)

func main() {
	pkg.FillRegion()
	pkg.FillCity()
	pkg.FillUsers()
	pkg.FillCategory()
	pkg.FillCourses()
	pkg.FillSubscriptions()
	for {
		var choice int

		fmt.Println("1. Получить курсы определенного пользователя")
		fmt.Println("2. Получить общий расход на курсы определенного пользователя")
		fmt.Println("3. Получить курсы по определенной категории")
		fmt.Println("4. Получить пользователей по региону")
		fmt.Println("5. Средний возраст покупателей определенного курса")
		fmt.Println("6. Выйти")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			functions.GetUserCourses()
		case 2:
			functions.GetUsersExpense()
		case 3:
			functions.GetCategoryCourses()
		case 4:
			functions.GetUsers()
		case 5:
			functions.AverageAge()
		case 6:
			return
		}
	}
}
