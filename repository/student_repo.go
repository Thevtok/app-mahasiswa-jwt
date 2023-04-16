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

	// insert into credentials table
	query1 := "INSERT INTO credentials (username, password) VALUES ($1, $2) RETURNING username"
	_, err = tx.Exec(query1, newstudent.Username, newstudent.Password)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return "failed to create student"
	}

	// insert into students table
	query2 := "INSERT INTO students (name, age, major, c_username) VALUES ($1, $2, $3, $4)"
	_, err = tx.Exec(query2, newstudent.Name, newstudent.Age, newstudent.Major, newstudent.Username)
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

	// start transaction
	tx, err := r.db.Begin()
	if err != nil {
		log.Println(err)
		return "failed to update student"
	}

	// update credentials table
	query1 := "UPDATE credentials SET password = $1 WHERE username = $2"
	_, err = tx.Exec(query1, student.Password, student.Username)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return "failed to update student"
	}

	// update students table
	query2 := "UPDATE students SET name = $1, age = $2, major = $3 WHERE id = $4"
	_, err = tx.Exec(query2, student.Name, student.Age, student.Major, student.ID)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return "failed to update student"
	}

	if err = tx.Commit(); err != nil {
		log.Println(err)
		return "failed to update student"
	}

	// return success message
	return fmt.Sprintf("student with id %d updated successfully", student.ID)
}

func (r *studentRepo) Delete(id int) string {
	// retrieve the c_username of the student with the given id
	var cUsername string
	query1 := "SELECT c_username FROM students WHERE id = $1"
	err := r.db.QueryRow(query1, id).Scan(&cUsername)
	if err == sql.ErrNoRows {
		return "student not found"
	} else if err != nil {
		log.Println(err)
		return "failed to delete student"
	}

	// delete the student with the given id
	query2 := "DELETE FROM students WHERE id = $1"
	_, err = r.db.Exec(query2, id)
	if err != nil {
		log.Println(err)
		return "failed to delete student"
	}

	// delete the record in the credentials table with the retrieved c_username value
	query3 := "DELETE FROM credentials WHERE username = $1"
	_, err = r.db.Exec(query3, cUsername)
	if err != nil {
		log.Println(err)
		return "failed to delete student's credentials"
	}

	return fmt.Sprintf("student with id %d and credentials  deleted successfully", id)
}

func NewStudentRepo(db *sql.DB) StudentRepo {
	repo := new(studentRepo)
	repo.db = db

	return repo
}
