package routes

import (
	"github.com/achmadsy/game-char-go/api/charsv1"
	"github.com/gin-gonic/gin"
)

func addUserCharsRoutes(rg *gin.RouterGroup) {
	chars := rg.Group("/user/chars")

	chars.GET("/all", charsv1.GetAllUserChars)
	chars.POST("/", charsv1.PostUserNewChar)
	chars.PUT("/:id", charsv1.PutUserEditedChar)
}
