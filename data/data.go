package data

import (
	"database/sql"
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

func (task *Task) CreateTask() (lastinsertid int64, err error) {
	stmt, err := Db.Prepare("INSERT INTO Tasks(Subject,Priority,CreatedAt)VALUES(?,?,NOW())")

	if err != nil {
		return 0, nil
	}

	defer stmt.Close()

	ret, err := stmt.Exec(task.Subject, task.Priority)

	if err != nil {
		return 0, err
	}

	id, _ := ret.LastInsertId()

	return id, nil
}

func (task *Task) UpdateTask() (err error) {
	stmt, err := Db.Prepare("UPDATE Tasks SET Subject = ?,Priority = ? WHERE Id = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(task.Subject, task.Priority, task.Id)

	if err != nil {
		return err
	}

	return
}

func (task *Task) DeleteTask() (err error) {
	stmt, err := Db.Prepare("DELETE FROM Tasks WHERE Id = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(task.Id)

	if err != nil {
		return err
	}

	return
}

func ReadTasks() (Tasks []Task, err error) {
	rows, err := Db.Query("SELECT Id,Subject,Priority,CreatedAt FROM Tasks")

	if err != nil {
		return nil, err
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

func ReadTask(id int) (task Task, err error) {
	task = Task{}

	err = Db.QueryRow("SELECT Id,Subject,Priority,CreatedAt FROM Tasks WHERE Id = ?", id).
		Scan(&task.Id, &task.Subject, &task.Priority, &task.CreatedAt)

	if err != nil {
		return task, err
	}

	return
}
