package models

import (
	"database/sql"
	"time"
)

// 商品SKU信息表
type GoodsSku struct {
	Id              int          `json:"id"`               // SKU_Id
	GoodsSkuNo      string       `json:"goods_sku_no"`     // sku编号
	GoodsId         int          `json:"goods_id"`         // 商品ID
	SkuName         string       `json:"sku_name"`         // 组合名称
	CombinationRule string       `json:"combination_rule"` // 组合规则
	CostPrice       float32      `json:"cost_price"`       // 成本价
	MarketPrice     float32      `json:"market_price"`     // 市场价
	SalePrice       float32      `json:"sale_price"`       // 售价
	SaleState       int          `json:"sale_state"`       // 销售状态 0:在售, 1: 禁售, 2: 下架
	SkuNum          int          `json:"sku_num"`          // 当前数量
	IsDeleted       int          `json:"is_deleted"`       // 是否删除;0:未删除,1:已删除
	DeletedAt       sql.NullTime `json:"deleted_at"`       // 删除日期
	CreatedAt       time.Time    `json:"created_at"`       // 创建时间
	UpdatedAt       time.Time    `json:"updated_at"`       // 最后更新时间
}
