package services

import (
	"Udemy/pkg"
	"Udemy/pkg/models"
)

func GetCoursesByUserIndex(index int) []models.Course {
	courses := make([]models.Course, 0)

	user := pkg.Users[index]

	for _, v := range pkg.Subscriptions {
		if user == v.User {
			courses = append(courses, v.Course)
		}
	}

	return courses
}
