package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaidev0711/go-project/api/controller"
	"github.com/kaidev0711/go-project/entities/shared"
	student_usecase "github.com/kaidev0711/go-project/usecase/student"
)

func Delete(ctx *gin.Context) {
	var input Input
	var err error

	input.ID = ctx.Params.ByName("id")
	input.UUID, err = shared.GetUuidByString(input.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, controller.NewResponseMessageError("Don't get uuid"))
		return
	}
	if err = student_usecase.Delete(input.UUID); err != nil {
		ctx.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, controller.NewResponseMessage("Delete successing"))
}
