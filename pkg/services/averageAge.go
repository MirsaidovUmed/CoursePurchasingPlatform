package services

import "Udemy/pkg"

func GetAverageUserAge(index int) float64 {

	var summ float64

	for _, v := range pkg.Users {

		summ += float64(2023 - v.BirthYear)
	}

	return summ / float64(len(pkg.Users))
}
