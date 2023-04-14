package controller

import (
	"net/http"
	"strconv"

	"mahasiswa_api/model"
	"mahasiswa_api/model/response"
	"mahasiswa_api/usecase"

	"github.com/gin-gonic/gin"
)

type StudentController struct {
	usecase usecase.StudentUsecase
}

func (c *StudentController) FindStudents(ctx *gin.Context) {
	res := c.usecase.FindStudents()
	if res == nil {
		response.JSONErrorResponse(ctx.Writer, http.StatusInternalServerError, "Failed to get students")
		return
	}
	response.JSONSuccess(ctx.Writer, http.StatusOK, res)
}

func (c *StudentController) FindStudentById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.JSONErrorResponse(ctx.Writer, http.StatusBadRequest, "Invalid user ID")
		return
	}

	res := c.usecase.FindStudentById(id)
	response.JSONSuccess(ctx.Writer, http.StatusOK, res)
}
func (c *StudentController) Register(ctx *gin.Context) {
	var newStudent model.Students

	if err := ctx.BindJSON(&newStudent); err != nil {
		response.JSONErrorResponse(ctx.Writer, http.StatusBadRequest, "Invalid Input")
		return
	}

	res := c.usecase.Register(&newStudent)
	response.JSONSuccess(ctx.Writer, http.StatusCreated, res)
}

func (c *StudentController) Edit(ctx *gin.Context) {
	var student model.Students

	if err := ctx.BindJSON(&student); err != nil {
		response.JSONErrorResponse(ctx.Writer, http.StatusBadRequest, "Invalid Input")
		return
	}

	res := c.usecase.Edit(&student)
	response.JSONSuccess(ctx.Writer, http.StatusOK, res)
}

func (c *StudentController) Unreg(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.JSONErrorResponse(ctx.Writer, http.StatusBadRequest, "Invalid user ID")
		return
	}

	res := c.usecase.Unreg(id)
	response.JSONSuccess(ctx.Writer, http.StatusOK, res)
}

func NewStudentController(u usecase.StudentUsecase) *StudentController {
	controller := StudentController{
		usecase: u,
	}

	return &controller
}
