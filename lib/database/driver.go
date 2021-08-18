package database

import (
	"etmarket/project/config"
	"etmarket/project/middlewares"
	"etmarket/project/models"
)

func CreateDriver(driver models.Driver) (interface{}, error) {
	if err := config.DB.Save(&driver).Error; err != nil {
		return nil, err
	}
	return driver, nil
}

//login driver with matching data from database
func LoginDriver(email string) (interface{}, error) {
	var driver models.Driver
	var err error
	if err = config.DB.Where("email = ?", email).First(&driver).Error; err != nil {
		return nil, err
	}
	driver.Token, err = middlewares.CreateToken(int(driver.ID))
	if err != nil {
		return nil, err
	}
	if err := config.DB.Save(driver).Error; err != nil {
		return nil, err
	}
	return driver, err
}

func CheckEmailOnDriver(email string) bool {
	var check bool
	var driver models.Driver

	config.DB.Model(&driver).Select("email").Where("email=?", email).First(&driver.Email)

	if driver.Email == "" {
		check = false
	} else {
		check = true
	}

	return check
}
