package database

import (
	"etmarket/project/config"
	"etmarket/project/middlewares"
	"etmarket/project/models"
)

func CreateDriver(driver models.Driver) (models.Driver, error) {
	if err := config.DB.Save(&driver).Error; err != nil {
		return driver, err
	}
	return driver, nil
}

//login driver with matching data from database
func LoginDriver(email string) (models.Driver, error) {
	var driver models.Driver
	var err error
	if err = config.DB.Where("email = ?", email).First(&driver).Error; err != nil {
		return driver, err
	}
	driver.Token, err = middlewares.CreateToken(int(driver.ID))
	if err != nil {
		return driver, err
	}
	if err := config.DB.Save(driver).Error; err != nil {
		return driver, err
	}
	return driver, err
}

/*
Author: Riska
This function for search password user by email
*/
func GetPwdDriver(email string) string {
	var driver models.Driver
	config.DB.Where("email = ?", email).First(&driver)
	return driver.Password
}

func CheckEmailOnDriver(email string) (interface{}, error) {
	var driver models.Driver

	if err := config.DB.Model(&driver).Select("email").Where("email=?", email).First(&driver.Email).Error; err != nil {
		return nil, err
	}

	return driver, nil
}

func GetDriverById(driver_id int) (models.Driver, error) {
	var driver models.Driver
	if err := config.DB.Where("id=?", driver_id).First(&driver).Error; err != nil {
		return driver, err
	}

	return driver, nil
}

//get 1 specified driver with Driver struct output
func GetDriver(id int) (models.Driver, error) {
	var driver models.Driver
	if err := config.DB.Find(&driver, "id=?", id).Error; err != nil {
		return driver, err
	}
	return driver, nil
}

//get email driver
func GetEmailDriverById(driver_id int) (string, error) {
	var driver models.Driver

	if err := config.DB.Model(&driver).Select("email").Where("id=?", driver_id).First(&driver.Email).Error; err != nil {
		return "nil", err
	}

	return driver.Email, nil
}

//update driver info from database
func UpdateDriver(driver models.Driver) (models.Driver, error) {
	if err := config.DB.Save(&driver).Error; err != nil {
		return driver, err
	}

	return driver, nil
}
