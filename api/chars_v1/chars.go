package charsv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetAllUserChars is to get all characters for the user
func GetAllUserChars(c *gin.Context) {
	c.JSON(http.StatusOK, "users comments")
}

//PostUserNewChar is to post new char
func PostUserNewChar(c *gin.Context) {
	c.JSON(http.StatusOK, "users comments")
}

//PutUserEditedChar is to edit the char
func PutUserEditedChar(c *gin.Context) {
	c.JSON(http.StatusOK, "users comments")
}
