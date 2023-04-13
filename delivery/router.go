package delivery

import (
	"log"
	"mahasiswa_api/config"
	"mahasiswa_api/controller"
	"mahasiswa_api/repository"
	"mahasiswa_api/usecase"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func StartServer() {

	serverPort := ":8080"

	db := config.LoadDatabase()

	router := gin.Default()

	studentsRouter := router.Group("/students")

	studentRepo := repository.NewStudentRepo(db)
	studentUsecase := usecase.NewStudentUsecase(studentRepo)
	studentCtrl := controller.NewStudentController(studentUsecase)

	studentsRouter.GET("", studentCtrl.FindStudents)
	studentsRouter.GET("/:id", studentCtrl.FindStudentById)
	studentsRouter.POST("", studentCtrl.Register)
	studentsRouter.PUT("", studentCtrl.Edit)
	studentsRouter.DELETE("/:id", studentCtrl.Unreg)

	// run server

	if err := router.Run(serverPort); err != nil {
		log.Fatal(err)
	}
}
