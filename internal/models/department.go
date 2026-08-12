package models

import "time"

type DepartmentStatus string

const (
	DepartmentActive  DepartmentStatus = "active"
	DepartmentDeleted DepartmentStatus = "deleted"
)

type Department struct {
	DepartmentID uint `gorm:"primaryKey;column:department_id"`

	DepartmentName string `gorm:"size:100;not null;uniqueIndex"`

	Description string `gorm:"type:text"`

	Status DepartmentStatus `gorm:"type:varchar(20);not null;default:'active'"`

	CreatedAt time.Time `gorm:"autoCreateTime"`

	//Doctors []Doctor `gorm:"foreignKey:DepartmentID"`
}
