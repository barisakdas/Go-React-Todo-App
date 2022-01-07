package entity

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Todo struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedDate string `json:"created-date"`
	EndDate     string `json:"end-date"`
	IsCompleted bool   `json:"is-completed"`
	IsDeleted   bool   `json:"is-deleted"`
}

func (todo Todo) Migrate() {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
	}
	db.AutoMigrate(&todo)
}

func (todo Todo) Add() (Todo, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		return todo, err
	}

	db.Create(&todo)
	return todo, nil
}

func (todo Todo) Get(where ...interface{}) (Todo, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		return todo, err
	}

	db.First(&todo, where...)
	return todo, nil
}

func (todo Todo) GetAll(where ...interface{}) ([]Todo, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	var todos []Todo
	db.Find(&todos, where...)
	return todos, nil
}

func (todo Todo) Update(data Todo) (Todo, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		return todo, err
	}

	db.Model(&todo).Updates(data)
	return todo, nil
}

func (todo Todo) Delete() error {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	db.Delete(&todo)
	return nil
}
