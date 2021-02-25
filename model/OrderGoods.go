package models

import (
	"database/sql"
	"time"
)

// 订单-商品-信息
type OrderGoods struct {
	Id              int          `json:"id"`               // 自增ID
	OrderId         int          `json:"order_id"`         // 订单ID
	SkuId           int          `json:"sku_id"`           // sku_id
	GoodsId         int          `json:"goods_id"`         // 商品ID
	SkuName         string       `json:"sku_name"`         // 组合名称
	CombinationRule string       `json:"combination_rule"` // 组合规则
	CostPrice       float32      `json:"cost_price"`       // 成本价
	MarketPrice     float32      `json:"market_price"`     // 市场价
	SalePrice       float32      `json:"sale_price"`       // 售价
	BuyNum          int          `json:"buy_num"`          // 购买数量
	IsDeleted       int          `json:"is_deleted"`       // 是否删除;0:未删除,1:已删除
	DeletedAt       sql.NullTime `json:"deleted_at"`       // 删除日期
	CreatedAt       time.Time    `json:"created_at"`       // 创建时间
	UpdatedAt       time.Time    `json:"updated_at"`       // 最后更新时间
}
