package routes

import (
	apicharsv1 "github.com/achmadsy/game-char-go/api/chars_v1"
	"github.com/gin-gonic/gin"
)

func addUserCharsRoutes(rg *gin.RouterGroup) {
	chars := rg.Group("/user/:userid/chars")

	chars.GET("/all", apicharsv1.GetAllUserChars)
	chars.POST("/", apicharsv1.PostUserNewChar)
	chars.PUT("/:id", apicharsv1.PutUserEditedChar)
}
