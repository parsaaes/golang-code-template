package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

const (
	templateName = "sql_template"
)

type Template struct {
	ID   int64  `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

type TemplateRepo interface {
	Create(template *Template) error
}

type SQLTemplateRepo struct {
	DB *gorm.DB
}

func (s SQLTemplateRepo) Create(template *Template) (finalErr error) {
	startTime := time.Now()

	defer func() { metrics.report(templateName, "create", startTime, finalErr) }()

	return s.DB.Create(template).Error
}
