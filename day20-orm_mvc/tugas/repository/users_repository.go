package repository

import (
	"tugas-mvc/config"
	"tugas-mvc/models"
)

func GetAllUsers() ([]models.User, error) {
	var usersGorm []models.UsersGorm
	var users []models.User
	err := config.DB.Find(&users).Error
	if err != nil {
		return users, err
	}
	// Format usersGorm to users
	for _, user := range usersGorm {
		data := models.User{
			Name:  user.Name,
			Email: user.Email,
			DOB:   user.DOB.String(),
		}
		users = append(users, data)
	}
	return users, nil
}

func AddUser(userReq models.AddUserRequest) (user models.User, err error) {
	err = config.DB.Create(&userReq).Error
	if err != nil {
		return user, err
	}
	// Format userReq to user
	user = models.User{
		Name:  userReq.Name,
		Email: userReq.Email,
		DOB:   userReq.DOB,
	}
	return user, nil
}

func GetUser(id int) (user models.User, err error) {
	err = config.DB.First(&user, id).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func UpdateUser(userReq models.User) (user models.User, err error) {
	err = config.DB.Save(&userReq).Error
	if err != nil {
		return user, err
	}

	return userReq, nil
}
