package usecase

import (
	"mahasiswa_api/model"
	"mahasiswa_api/repository"
)

type StudentUsecase interface {
	FindStudents() any
	FindStudentById(id int) any
	Register(newstudent *model.Students) string
	Edit(student *model.Students) string
	Unreg(id int) string
}

type studentUsecase struct {
	studentRepo repository.StudentRepo
}

func (u *studentUsecase) FindStudents() any {
	return u.studentRepo.GetAll()
}

func (u *studentUsecase) FindStudentById(id int) any {
	return u.studentRepo.GetById(id)
}

func (u *studentUsecase) Register(newStudent *model.Students) string {
	return u.studentRepo.Create(newStudent)
}

func (u *studentUsecase) Edit(student *model.Students) string {
	return u.studentRepo.Update(student)
}

func (u *studentUsecase) Unreg(id int) string {
	return u.studentRepo.Delete(id)
}

func NewStudentUsecase(studentRepo repository.StudentRepo) StudentUsecase {
	return &studentUsecase{
		studentRepo: studentRepo,
	}
}
