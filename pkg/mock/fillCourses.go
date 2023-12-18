package mock

import (
	"Udemy/pkg"
	"Udemy/pkg/models"
)

func FillCourses() {
	pkg.Courses = append(pkg.Courses, models.Course{
		Name:     "Основы GO",
		Category: pkg.Categories[0],
		Duration: 10000,
		Price:    0,
	})

	pkg.Courses = append(pkg.Courses, models.Course{
		Name:     "Как приготовить оливье к Новому году",
		Category: pkg.Categories[1],
		Duration: 120,
		Price:    100,
	})

	pkg.Courses = append(pkg.Courses, models.Course{
		Name:     "Как увеличить продажи",
		Category: pkg.Categories[3],
		Duration: 360,
		Price:    10000,
	})

	pkg.Courses = append(pkg.Courses, models.Course{
		Name:     "Figma",
		Category: pkg.Categories[2],
		Duration: 7200,
		Price:    1500,
	})
}
