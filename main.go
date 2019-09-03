package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"runtime/debug"
)

var db *sql.DB

type Student struct {
	StudentNo int
	FullName string
	Age int
}

func main() {
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
}

func insert(student Student) error {
	_, err := db.Exec("INSERT student VALUES (?, ?, ?)", student.StudentNo, student.FullName, student.Age)

	return err;
}

func query(studentNo int) (Student, error) {
	student := Student{}

	err := db.QueryRow("SELECT * FROM student WHERE studentNo = ?", studentNo).Scan(&student.StudentNo, &student.FullName, &student.Age)

	return student, err
}


