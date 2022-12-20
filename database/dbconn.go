package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name   string `json:"name"`
	Status string `json:"status"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Sid@2002"
	dbname   = "todo-list-api"
)

var dsn string = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
	host, port, user, password, dbname)

func InitDB() error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&Task{})

	return nil
}

func CreateTask(name string, status string) (Task, error) {
	var newTask = Task{Name: name, Status: status}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return newTask, err
	}
	db.Create(&Task{Name: name, Status: status})

	return newTask, nil
}

func GetallTasks() ([]Task, error) {
	var tasks []Task

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return tasks, err
	}

	db.Find(&tasks)

	return tasks, nil
}

func Gettask(id int) (Task, error) {
	var task Task

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return task, err
	}

	db.Where("ID = &id").Find(&task)
	return task, nil
}
