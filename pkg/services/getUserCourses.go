package services

import "Udemy/pkg"

func GetAverageUserAgeByCourseIndex(index int) float64 {
	course := pkg.Courses[index]

	var summ float64
	var count_of_users float64

	for _, v := range pkg.Subscriptions {
		if v.Course == course {
			summ += 2023 - float64(v.User.BirthYear)

			count_of_users++
		}
	}

	return summ / count_of_users
}
