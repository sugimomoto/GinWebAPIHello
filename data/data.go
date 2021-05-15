package data

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "user:Password!@/sample?parseTime=true")

	if err != nil {
		panic(err)
	}
}

type Task struct {
	Id        int
	Subject   string
	Priority  string
	CreatedAt time.Time
}

func (task *Task) Createtask() (lastinsertid int64, err error) {
	stmt, err := Db.Prepare("INSERT INTO Tasks(Subject,Priority,CreatedAt)VALUES(?,?,NOW())")

	if err != nil {
		log.Fatal("Insert Prepare error : ", err)
		return
	}

	defer stmt.Close()

	ret, err := stmt.Exec(task.Subject, task.Priority)

	if err != nil {
		log.Fatal("Insert Exec error : ", err)
		return
	}

	id, _ := ret.LastInsertId()

	return id, nil
}

func (task *Task) Updatetask() (err error) {
	stmt, err := Db.Prepare("UPDATE Tasks SET Subject = ?,Priority = ? WHERE Id = ?")

	if err != nil {
		log.Fatal("Insert Prepare error : ", err)
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(task.Subject, task.Priority, task.Id)

	if err != nil {
		log.Fatal("Insert Exec error : ", err)
		return
	}

	return
}

func (task *Task) Deletetask() (err error) {
	stmt, err := Db.Prepare("DELETE FROM Tasks WHERE Id = ?")

	if err != nil {
		log.Fatal("Insert Prepare error : ", err)
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(task.Id)

	if err != nil {
		log.Fatal("Insert Exec error : ", err)
		return
	}

	return
}

func ReadTasks() (Tasks []Task, err error) {
	rows, err := Db.Query("SELECT Id,Subject,Priority,CreatedAt FROM Tasks")

	if err != nil {
		log.Fatal("Connection Error : ", err)
		return
	}

	for rows.Next() {
		task := Task{}

		if err = rows.Scan(&task.Id, &task.Subject, &task.Priority, &task.CreatedAt); err != nil {
			return
		}

		Tasks = append(Tasks, task)
	}

	rows.Close()
	return
}

func Readtask(id int) (task Task, err error) {
	task = Task{}

	err = Db.QueryRow("SELECT Id,Subject,Priority,CreatedAt FROM Tasks WHERE Id = ?", id).
		Scan(&task.Id, &task.Subject, &task.Priority, &task.CreatedAt)

	if err != nil {
		log.Fatal("Connection Error : ", err)
		return
	}

	return
}
