package repository

import (
	"database/sql"
	"fmt"
	"log"

	"mahasiswa_api/model"
)

type StudentRepo interface {
	GetAll() any
	GetById(id int) any
	Create(newStudentr *model.Students) string
	Update(student *model.Students) string
	Delete(id int) string
}

type studentRepo struct {
	db *sql.DB
}

func (r *studentRepo) GetAll() any {
	var students []model.Students

	query := `SELECT s.name, s.age, s.major,  c.username 
				FROM students s
				JOIN credentials c ON s.c_username = c.username`
	rows, err := r.db.Query(query)

	if err != nil {
		log.Println(err)
	}
	if rows == nil {
		return "no data"
	}
	defer rows.Close()

	for rows.Next() {
		var student model.Students

		if err := rows.Scan(&student.Name, &student.Age, &student.Major, &student.Username); err != nil {
			log.Println(err)
		}

		students = append(students, student)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	if len(students) == 0 {
		return "no data"
	}

	return students
}

func (r *studentRepo) GetById(id int) any {
	var studentInDb model.Students

	query := "SELECT id, name, age, major FROM students WHERE id = $1"
	row := r.db.QueryRow(query, id)

	err := row.Scan(&studentInDb.ID, &studentInDb.Name, &studentInDb.Age, &studentInDb.Major)

	if err != nil {
		log.Println(err)
	}

	if studentInDb.ID == 0 {
		return "student not found"
	}

	return studentInDb
}

func (r *studentRepo) Create(newstudent *model.Students) string {
	tx, err := r.db.Begin()
	if err != nil {
		log.Println(err)
		return "failed to create student"
	}

	// insert into students table
	query1 := "INSERT INTO students (name, age, major,c_username) VALUES ($1, $2, $3,$4)"
	_, err = tx.Exec(query1, newstudent.Name, newstudent.Age, newstudent.Major, newstudent.C_Username)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return "failed to create student"
	}

	// insert into credentials table
	query2 := "INSERT INTO credentials (username, password) VALUES ($1, $2)"
	_, err = tx.Exec(query2, newstudent.Username, newstudent.Password)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return "failed to create student"
	}

	if err = tx.Commit(); err != nil {
		log.Println(err)
		return "failed to create student"
	}

	return "student created successfully"
}

func (r *studentRepo) Update(student *model.Students) string {
	res := r.GetById(student.ID)

	// jika tidak ada, return pesan
	if res == "student not found" {
		return res.(string)
	}

	// jika ada, maka update student
	query := "UPDATE students SET name = $1, age = $2, major = $3 WHERE id = $4"
	_, err := r.db.Exec(query, student.Name, student.Age, student.Major, student.ID)

	if err != nil {
		log.Println(err)
	}

	// jika update berhasil, return pesan sukses
	return fmt.Sprintf("student with id %d updated successfully", student.ID)
}

func (r *studentRepo) Delete(id int) string {
	res := r.GetById(id)

	// jika tidak ada, return pesan
	if res == "student not found" {
		return res.(string)
	}

	// jika ada, delete student
	query := "DELETE FROM students WHERE id = $1"
	_, err := r.db.Exec(query, id)

	if err != nil {
		log.Println(err)
		return "failed to delete student"
	}

	return fmt.Sprintf("student with id %d deleted successfully", id)
}

func NewStudentRepo(db *sql.DB) StudentRepo {
	repo := new(studentRepo)
	repo.db = db

	return repo
}
