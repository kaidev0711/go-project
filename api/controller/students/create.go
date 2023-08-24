package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaidev0711/go-project/api/controller"
	// student_usecase "github.com/kaidev0711/go-project/usecase/student"
)

func (sc *StudentController) Create(ctx *gin.Context) {
	var input Input

	if err := ctx.Bind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	student, err := sc.StudentUsecase.Create(input.FullName, input.Age)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
		return
	}

	ctx.JSON(http.StatusAccepted, student)
}
