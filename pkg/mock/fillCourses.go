package mock

import (
	pkg "Udemy/pkg/database"
	"Udemy/pkg/models"
)

func FillCourses() {
	pkg.Courses = append(pkg.Courses, models.Course{
		Name:     "BackEnd",
		Price:    5000,
		Category: pkg.Category[1],
		Duration: "6_month",
	})

	pkg.Courses = append(pkg.Courses, models.Course{
		Name:     "FrontEnd",
		Price:    1500,
		Category: pkg.Category[1],
		Duration: "6_month",
	})

	pkg.Courses = append(pkg.Courses, models.Course{
		Name:     "English",
		Price:    3000,
		Category: pkg.Category[0],
		Duration: "6_month",
	})
}
