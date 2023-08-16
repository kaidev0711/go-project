package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaidev0711/go-project/api/controller"
	student_usecase "github.com/kaidev0711/go-project/usecase/student"
)

func List(ctx *gin.Context) {
	students, err := student_usecase.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, controller.NewResponseMessageError(err.Error()))
	}
	ctx.JSON(http.StatusOK, students)
}
