package services

import (
	"Udemy/pkg"
	"Udemy/pkg/models"
)

func GetUsersByRegionIndex(index int) []models.User {
	region := pkg.Regions[index]

	users := make([]models.User, 0)

	for _, v := range pkg.Users {
		if v.City.Region == region {
			users = append(users, v)
		}
	}

	return users
}
