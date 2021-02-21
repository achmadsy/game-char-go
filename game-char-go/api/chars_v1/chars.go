package apicharsv1

import (
	"net/http"
	"strconv"

	"github.com/achmadsy/game-char-go/models"
	charsservicev1 "github.com/achmadsy/game-char-go/service/chars_v1"
	"github.com/gin-gonic/gin"
)

//GetAllUserChars is to get all characters for the user
// @Summary Show all user's characters
// @Description get all characters for the user
// @Produce json
// @Param userID path int true "User ID"
// @Success 200 {object} models.Character
// @Failure 400 {object} models.OutStruct
// @Failure 404 {object} models.OutStruct
// @Failure 500 {object} models.OutStruct
// @Failure default {object} models.OutStruct
// @Router /v1/user/{userID}/chars/all [get]
func GetAllUserChars(c *gin.Context) {
	user := models.User{}
	userID, _ := strconv.ParseUint(c.Params.ByName("userid"), 10, 32)
	if err := checkUserID(userID); err {
		out := models.SetResponse(false, "UserID not found", nil)
		c.JSON(http.StatusBadRequest, out)
		return
	}
	if err := charsservicev1.GetUserAllCharsService(&userID, &user); err != nil {
		out := models.SetResponse(false, "Failed to fetch characters", err)
		c.JSON(http.StatusInternalServerError, out)
		return
	}
	c.JSON(http.StatusOK, user.Chars)
}

//PostUserNewChar is to post new char
// @Summary Post new character
// @Description Post new character for the user
// @Produce json
// @Accept  json
// @Param account body models.Character true "New Character"
// @Param userID path int true "User ID"
// @Success 200 {object} models.Character
// @Failure 400 {object} models.OutStruct
// @Failure 404 {object} models.OutStruct
// @Failure 500 {object} models.OutStruct
// @Failure default {object} models.OutStruct
// @Router /v1/user/{userID}/chars/ [post]
func PostUserNewChar(c *gin.Context) {
	userChar := models.Character{}
	userID, _ := strconv.ParseUint(c.Params.ByName("userid"), 10, 32)
	c.Bind(&userChar)
	calcValue(&userChar)
	userChar.ID = 0
	userChar.UserID = uint(userID)
	if err := checkCharCode(userChar.CharacterCode); err {
		out := models.SetResponse(false, "Wrong character code", nil)
		c.JSON(http.StatusBadRequest, out)
		return
	}
	if err := checkUserID(userID); err {
		out := models.SetResponse(false, "UserID not found", nil)
		c.JSON(http.StatusBadRequest, out)
		return
	}
	if err := charsservicev1.PostNewChar(&userChar); err != nil {
		out := models.SetResponse(false, "Failed to create character detail", err)
		c.JSON(http.StatusInternalServerError, out)
		return
	}
	out := models.SetResponse(true, "Post user's new character success", nil)
	c.JSON(http.StatusOK, out)
}

//PutUserEditedChar is to edit the char
// @Summary Edit character
// @Description Edit character for the user
// @Produce json
// @Accept  json
// @Param account body models.Character true "Edit Character"
// @Param userID path int true "User ID"
// @Param charID path int true "Character ID"
// @Success 200 {object} models.Character
// @Failure 400 {object} models.OutStruct
// @Failure 404 {object} models.OutStruct
// @Failure 500 {object} models.OutStruct
// @Failure default {object} models.OutStruct
// @Router /v1/user/{userID}/chars/{charID} [put]
func PutUserEditedChar(c *gin.Context) {
	userChar := models.Character{}
	userID, _ := strconv.ParseUint(c.Params.ByName("userid"), 10, 32)
	charID, _ := strconv.Atoi(c.Params.ByName("id"))
	c.Bind(&userChar)
	calcValue(&userChar)
	if err := checkCharCode(userChar.CharacterCode); err {
		out := models.SetResponse(false, "Wrong character code", nil)
		c.JSON(http.StatusBadRequest, out)
		return
	}
	if err := checkUserID(userID); err {
		out := models.SetResponse(false, "UserID not found", nil)
		c.JSON(http.StatusBadRequest, out)
		return
	}
	if err := charsservicev1.UpdateDetailChar(&userID, &charID, &userChar); err != nil {
		out := models.SetResponse(false, "Failed to update character detail", err)
		c.JSON(http.StatusInternalServerError, out)
		return
	}
	out := models.SetResponse(true, "Update user's character success", nil)
	c.JSON(http.StatusOK, out)
}

//calcValue to calculate char's value
func calcValue(char *models.Character) {
	switch char.CharacterCode {
	case 1:
		char.Value = (150.00 / 100.00) * (float32(char.Power))
	case 2:
		char.Value = 2.00 + (110.00/100.00)*(float32(char.Power))
	case 3:
		if char.Power < 20 {
			char.Value = (200.00 / 100.00) * (float32(char.Power))
		} else {
			char.Value = (300.00 / 100.00) * (float32(char.Power))
		}
	}
}

func checkCharCode(code int) bool {
	if code > 3 || code < 1 {
		return true
	}
	return false
}

func checkUserID(userID uint64) bool {
	if user, _ := charsservicev1.GetUserByID(&userID); user.ID != 0 {
		return false
	}
	return true
}
