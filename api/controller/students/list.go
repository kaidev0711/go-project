package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaidev0711/go-project/api/controller"
	// student_usecase "github.com/kaidev0711/go-project/usecase/student"
)

func (sc *StudentController) List(ctx *gin.Context) {
	students, err := sc.StudentUsecase.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
	}
	ctx.JSON(http.StatusOK, students)
}
