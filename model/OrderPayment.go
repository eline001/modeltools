package models

import (
	"database/sql"
	"time"
)

// 订单支付记录
type OrderPayment struct {
	Id           int          `json:"id"`             // 自增ID
	OrderId      int          `json:"order_id"`       // 订单ID
	PaymentType  int          `json:"payment_type"`   // 支付方式 1:微信
	TotalPrice   float32      `json:"total_price"`    // 支付金额
	PaymentNo    string       `json:"payment_no"`     // 内部支付订单号
	ThirdTradeNo string       `json:"third_trade_no"` // 外部支付订单号
	PaymentState int          `json:"payment_state"`  // 支付状态 0: 未支付1:已支付
	PaiedTime    time.Time    `json:"paied_time"`     // 支付时间
	IsDeleted    int          `json:"is_deleted"`     // 是否删除;0:未删除,1:已删除
	DeletedAt    sql.NullTime `json:"deleted_at"`     // 删除日期
	CreatedAt    time.Time    `json:"created_at"`     // 创建时间
	UpdatedAt    time.Time    `json:"updated_at"`     // 最后更新时间
}
