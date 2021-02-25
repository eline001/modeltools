package models

import (
	"database/sql"
	"time"
)

// 基本订单信息
type Order struct {
	Id                int          `json:"id"`                 // Id
	OrderNo           string       `json:"order_no"`           // sku编号
	OrderPrice        float32      `json:"order_price"`        // 订单总金额
	PreferentialPrice float32      `json:"preferential_price"` // 优惠金额
	ActuallyPrice     float32      `json:"actually_price"`     // 实际金额，用户实付金额
	OrderState        int          `json:"order_state"`        // 订单状态   0:未支付, 1: 已支付, 2: 已过期,
	PaiedTime         time.Time    `json:"paied_time"`         // 支付时间
	IsDeleted         int          `json:"is_deleted"`         // 是否删除;0:未删除,1:已删除
	DeletedAt         sql.NullTime `json:"deleted_at"`         // 删除日期
	CreatedAt         time.Time    `json:"created_at"`         // 创建时间
	UpdatedAt         time.Time    `json:"updated_at"`         // 最后更新时间
}
