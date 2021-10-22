package user

import (
	"go-web/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"users": model.Get(),
	})
}
