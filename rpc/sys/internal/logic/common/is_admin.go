package common

import (
	"context"
	"gorm.io/gorm"
)

// IsAdmin 判断是不是超级管理员
func IsAdmin(ctx context.Context, userId int64, db *gorm.DB) bool {
	sql := `select count(1) from sys_user_role t1 where t1.role_id = 1 and t1.user_id = ?`
	var count int64

	db.WithContext(ctx).Raw(sql, userId).Scan(&count)
	return count > 0
}
