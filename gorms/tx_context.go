package gorms

import (
	"context"
	"gorm.io/gorm"
)

var key = "gorm_tx"

// 创建事务
func CreateTx(ctx context.Context, db *gorm.DB) (context.Context, *gorm.DB) {
	value := ctx.Value(key)
	if value == nil {
		tx := db.WithContext(ctx).Begin()
		newCtx := context.WithValue(ctx, key, tx)
		return newCtx, tx
	} else {
		tx := value.(*gorm.DB)
		return ctx, tx
	}
}

// 获取db，如果上下文中有事务db,则返回，否则返回默认的db
func GetDb(ctx context.Context, db *gorm.DB) *gorm.DB {
	value := ctx.Value(key)
	if value == nil {
		return db.WithContext(ctx)
	}
	tx := value.(*gorm.DB)
	return tx
}

// 提交事务，如果有事务db，则提交
func CommitTx(ctx context.Context) {
	value := ctx.Value(key)
	if value != nil {
		db := value.(*gorm.DB)
		db.Commit()
	}
}

// 回滚事务，如果有事务db，则回滚
func RollbackTx(ctx context.Context) {
	value := ctx.Value(key)
	if value != nil {
		db := value.(*gorm.DB)
		db.Commit()
	}
}

func NewWithContext(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, key, db)
}
