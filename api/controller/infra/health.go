package infra

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaidev0711/go-project/api/controller"
)

func Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, controller.NewResponseMessage("ok"))
}
