package routes

import (
	"github.com/achmadsy/game-char-go/api/chars/api_chars_v1"
	"github.com/gin-gonic/gin"
)

func addUserCharsRoutes(rg *gin.RouterGroup) {
	chars := rg.Group("/user/chars")

	chars.GET("/all", api_chars_v1.GetAllUserChars)
	chars.POST("/", api_chars_v1.PostUserNewChar)
	chars.PUT("/:id", api_chars_v1.PutUserEditedChar)
}
