package database

import (
	"errors"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	URL      string
	Title    string
	MarkDown string
}

type Operator interface {
	Write(r Record) error
	Read() error
}

type operatorImpl struct {
	db *gorm.DB
}

func New() (Operator, error) {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database; %w", err)
	}
	if err := db.AutoMigrate(&Record{}); err != nil {
		return nil, fmt.Errorf("failed to migrate; %w", err)
	}
	return &operatorImpl{db: db}, nil
}

func (o *operatorImpl) Write(r Record) error {
	var rc Record
	err := o.db.First(&rc, "URL = ?", r.URL).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		if err := o.db.Create(r).Error; err != nil {
			return fmt.Errorf("failed to create; %w", err)
		}
	} else {
		if err := o.db.Model(&rc).Updates(r).Error; err != nil {
			return fmt.Errorf("failed to update; %w", err)
		}
	}
	return nil
}

func (o *operatorImpl) Read() error {
	return nil
}
