package config

import (
	"errors"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	return errors.New("Not implemented")
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
