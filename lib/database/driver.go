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

func CheckEmailOnDriver(email string) (interface{}, error) {
	var driver models.Driver

	if err := config.DB.Model(&driver).Select("email").Where("email=?", email).First(&driver.Email).Error; err != nil {
		return nil, err
	}

	return driver, nil
}

func GetDriverById(driver_id int) (interface{}, error) {
	var driver models.Driver
	if err := config.DB.Where("id=?", driver_id).First(&driver).Error; err != nil {
		return nil, err
	}
	return driver, nil
}
