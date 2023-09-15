package model

import "time"

type BaseModel struct {
	ID         int        `gorm:"primaryKey"`
	CreateTime *time.Time `gorm:"autoCreateTime"`
	UpdateTime *time.Time `gorm:"autoCreateTime"`
}

type Teacher struct {
	BaseModel
	Name   string `gorm:"type:varchar(32);unique;not null"`
	Tno    int
	Pwd    string `gorm:"type:varchar(100);not null"`
	Tel    string `gorm:"type:char(11);"`
	Birth  *time.Time
	Remark string `gorm:"type:varchar(255);"`
}

type Class struct {
	BaseModel
	Name string `gorm:"type:varchar(32);unique;not null"`
	Num  int
	// 导员/班主任
	TeacherID int
	Teacher   Teacher `gorm:"constraint:OnDelete:CASCADE;"`

	Students []Student
}

type Course struct {
	BaseModel
	Name   string `gorm:"type:varchar(32);unique;not null"`
	Credit int
	Period int
	// 多对一
	TeacherID int
	Teacher   Teacher `gorm:"constraint:OnDelete:CASCADE;"`
}

type Student struct {
	BaseModel
	Name   string `gorm:"type:varchar(32);unique;not null"`
	Sno    int
	Pwd    string `gorm:"type:varchar(100);not null"`
	Tel    string `gorm:"type:char(11);"`
	Gender byte   `gorm:"default:1"`
	Birth  *time.Time
	Remark string `gorm:"type:varchar(255);"`

	// 多对一
	ClassID int
	Class   Class `gorm:"foreignKey:ClassID;constraint:OnDelete:CASCADE;"`

	// 多对多
	Courses []Course `gorm:"many2many:student2course;constraint:OnDelete:CASCADE;"`
}
