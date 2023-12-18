package main

import (
	"Udemy/pkg/mock"
	"Udemy/pkg/services"
	"fmt"
)

func main() {
	mock.FillAll()

	avg := services.GetAverageUserAgeByCourseIndex(0)

	fmt.Println(avg)

}
