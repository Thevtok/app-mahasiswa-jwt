package controller

import (
	"net/http"
	"strconv"

	"mahasiswa_api/model"
	"mahasiswa_api/usecase"

	"github.com/gin-gonic/gin"
)

type StudentController struct {
	usecase usecase.StudentUsecase
}

func (c *StudentController) FindStudents(ctx *gin.Context) {
	res := c.usecase.FindStudents()

	ctx.JSON(http.StatusOK, res)
}

func (c *StudentController) FindStudentById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid user ID")
		return
	}

	res := c.usecase.FindStudentById(id)
	ctx.JSON(http.StatusOK, res)
}

func (c *StudentController) Register(ctx *gin.Context) {
	var newStudent model.Students

	if err := ctx.BindJSON(&newStudent); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		return
	}

	res := c.usecase.Register(&newStudent)
	ctx.JSON(http.StatusCreated, res)
}

func (c *StudentController) Edit(ctx *gin.Context) {
	var student model.Students

	if err := ctx.BindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		return
	}

	res := c.usecase.Edit(&student)
	ctx.JSON(http.StatusOK, res)
}

func (c *StudentController) Unreg(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid user ID")
		return
	}

	res := c.usecase.Unreg(id)
	ctx.JSON(http.StatusOK, res)
}

func NewStudentController(u usecase.StudentUsecase) *StudentController {
	controller := StudentController{
		usecase: u,
	}

	return &controller
}
