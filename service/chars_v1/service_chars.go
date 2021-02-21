package servicecharsv1

import (
	"github.com/achmadsy/game-char-go/db"
	"github.com/achmadsy/game-char-go/models"
)

//GetUserAllCharsService to query to get all user characters
func GetUserAllCharsService(userID *uint64, user *models.User) error {
	if err := db.DB.Model(&user).Preload("Chars").Where("id = ?", userID).Find(&user).Error; err != nil {
		return err
	}
	return nil
}

//UpdateDetailChar to update user character
func UpdateDetailChar(userID *uint64, userCharID *int, userCharData *models.Character) error {
	if err := db.DB.Model(&userCharData).Where("id = ? and user_id = ?", userCharID, userID).Updates(models.Character{Name: userCharData.Name, Power: userCharData.Power}).Error; err != nil {
		return err
	}
	return nil
}

//PostNewChar to post new user character
func PostNewChar(userCharData *models.Character) error {
	if result := db.DB.Create(&userCharData); result.Error != nil {
		return result.Error
	}
	return nil
}

//GetUserByID to get user by ID
func GetUserByID(userID *uint64) (models.User, error) {
	user := models.User{}
	if err := db.DB.Model(&user).Where("id = ?", userID).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
