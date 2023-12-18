package functions

import (
	pkg "Udemy/pkg/database"
	"fmt"
)

func GetUsers() {
	var (
		region string
		has    bool
	)
	fmt.Println("Введите регион")
	fmt.Scan(&region)

	for _, val := range pkg.Region {
		if region == val.Name {
			has = true
		}
	}

	if has {
		fmt.Println("Пользователи по региону:")
		for _, val := range pkg.User {
			if val.City.Region.Name == region {
				fmt.Println(val.Name)
			}
		}
	} else {
		fmt.Println("Ошибка!, данного региона нет в нашей бд")
	}
}
