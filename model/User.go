package models

import (
	"database/sql"
	"time"
)

// 用户表
type User struct {
	Id        int          `json:"id"`         // 用户ID
	UserPhone string       `json:"user_phone"` // 用手机号
	UserName  string       `json:"user_name"`  // 用户名称
	Gender    int          `json:"gender"`     // 性别: 0:保密,1:女,2:男
	IsDeleted int          `json:"is_deleted"` // 是否删除;0:未删除,1:已删除
	DeletedAt sql.NullTime `json:"deleted_at"` // 删除日期
	CreatedAt time.Time    `json:"created_at"` // 创建时间
	UpdatedAt time.Time    `json:"updated_at"` // 最后更新时间
}
