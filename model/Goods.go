package models

import (
	"database/sql"
	"time"
)

// 商品基础信息表
type Goods struct {
	Id          int          `json:"id"`           // 商品Id
	GoodsNo     string       `json:"goods_no"`     // 商品编号
	GoodsName   string       `json:"goods_name"`   // 商品名称
	CostPrice   float32      `json:"cost_price"`   // 成本价
	MarketPrice float32      `json:"market_price"` // 市场价
	SalePrice   float32      `json:"sale_price"`   // 售价
	SaleState   int          `json:"sale_state"`   // 销售状态 0:在售, 1: 禁售, 2: 下架
	IsDeleted   int          `json:"is_deleted"`   // 是否删除;0:未删除,1:已删除
	DeletedAt   sql.NullTime `json:"deleted_at"`   // 删除日期
	CreatedAt   time.Time    `json:"created_at"`   // 创建时间
	UpdatedAt   time.Time    `json:"updated_at"`   // 最后更新时间
}
