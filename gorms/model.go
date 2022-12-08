package gorms

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        int64          `gorm:"primarykey" gen:"detail;edit;" json:"id"`
	CreatedAt time.Time      `gorm:"not null" gen:"detail;list;order;where;"`
	UpdatedAt time.Time      `gorm:"not null" gen:"detail;list;order;where;"`
	DeletedAt gorm.DeletedAt `gorm:"index" gen:"ignore"`
}

// 权限Model
type AuthModel struct {
	// 权限-部门ID
	AuthDepartmentId *int64
}
