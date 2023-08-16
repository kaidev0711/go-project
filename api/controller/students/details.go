package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaidev0711/go-project/api/controller"
	"github.com/kaidev0711/go-project/entities"
	"github.com/kaidev0711/go-project/entities/shared"
	student_usecase "github.com/kaidev0711/go-project/usecase/student"
)

func Details(ctx *gin.Context) {
	var input Input
	var err error

	var studentFound entities.Student

	input.ID = ctx.Params.ByName("id")
	input.UUID, err = shared.GetUuidByString(input.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, controller.NewResponseMessageError("Don't find uuid"))
		return
	}

	studentFound, err = student_usecase.SearchByID(input.UUID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, controller.NewResponseMessageError(err.Error()))
	}

	ctx.JSON(http.StatusOK, studentFound)
}
