package gorms

import (
	"context"
	"gorm.io/gorm"
)

type HealthCheckProbe struct {
	DB *gorm.DB

	ProbeName string
}

func (h HealthCheckProbe) IsHealth(ctx context.Context) error {
	return h.DB.Exec("select 1").Error
}

func (h HealthCheckProbe) Name() string {
	if h.ProbeName == "" {
		return "redis"
	}
	return h.ProbeName
}
