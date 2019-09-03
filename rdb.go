package data_migration

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"runtime/debug"
)

var db *sql.DB

type Student struct {
	StudentNo int
	FullName string
	Age int
	OldStudentNo int
}

func init() {
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		debug.PrintStack()
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		debug.PrintStack()
		log.Fatal(err)
	}

	fmt.Println("Successfully pinged database.")
}

func Insert(student Student) error {
	_, err := db.Exec("INSERT student (fullName, age, oldStudentNo) VALUES (?, ?, ?)", student.StudentNo, student.FullName, student.Age)

	return err
}

func Find(studentNo int) (Student, error) {
	student := Student{}

	err := db.QueryRow("SELECT * FROM student WHERE oldStudentNo = ?", studentNo).Scan(&student.StudentNo, &student.FullName, &student.Age, &student.OldStudentNo)

	return student, err
}


