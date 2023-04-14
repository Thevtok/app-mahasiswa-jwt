package delivery

import (
	"log"
	"mahasiswa_api/config"
	"mahasiswa_api/controller"
	"mahasiswa_api/repository"
	"mahasiswa_api/usecase"
	"mahasiswa_api/utils"

	"github.com/Thevtok/auth/cont"
	"github.com/Thevtok/auth/repo"
	"github.com/Thevtok/auth/service"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func StartServer() {
	s_key := []byte(utils.DotEnv("KEY"))

	db := config.LoadDatabase()
	defer db.Close()
	authMiddleware := cont.AuthMiddleware(s_key)

	r := gin.Default()

	studentsRouter := r.Group("/students")
	studentsRouter.Use(authMiddleware)

	sr := repository.NewStudentRepo(db)
	su := usecase.NewStudentUsecase(sr)
	sc := controller.NewStudentController(su)
	studentRepo := repo.NewStudentRepo(db)
	loginService := service.NewLoginService(studentRepo)
	loginJwt := cont.NewUserJwt(loginService)

	r.POST("/login", loginJwt.Login)

	studentsRouter.GET("", sc.FindStudents)
	studentsRouter.GET("/:id", sc.FindStudentById)

	studentsRouter.PUT("", sc.Edit)
	studentsRouter.DELETE("/:id", sc.Unreg)

	r.POST("/register", sc.Register)

	// run server

	if err := r.Run(utils.DotEnv("SERVER_PORT")); err != nil {
		log.Fatal(err)
	}
}
