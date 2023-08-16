package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaidev0711/go-project/api/controller"
	"github.com/kaidev0711/go-project/entities/shared"
	student_usecase "github.com/kaidev0711/go-project/usecase/student"
)

func Update(ctx *gin.Context) {
	var input Input
	var err error
	err = ctx.Bind(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	input.ID = ctx.Params.ByName("id")

	input.UUID, err = shared.GetUuidByString(input.ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	student, err := student_usecase.Update(input.UUID, input.FullName, input.Age)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
		return
	}

	ctx.JSON(http.StatusAccepted, student)
}
