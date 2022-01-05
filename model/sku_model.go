package model

import (
	"App-CloudBase-mcdull-mall/pkg/database"
	"context"
	"time"
)

// SkuDO 商城-SKU
type SkuDO struct {
	Id         int       `gorm:"column:id" json:"id"`
	Title      string    `gorm:"column:title" json:"title"`       // 标题
	Price      string    `gorm:"column:price" json:"price"`       // 价格
	Code       string    `gorm:"column:code" json:"code"`         // 编码
	Stock      int       `gorm:"column:stock" json:"stock"`       // 库存量
	GoodsId    int       `gorm:"column:goods_id" json:"goods_id"` // 商品ID
	Online     int       `gorm:"column:online" json:"online"`     // 是否上线
	Picture    string    `gorm:"column:picture" json:"picture"`   // 图片
	Specs      string    `gorm:"column:specs" json:"specs"`       // 多规格属性
	Del        int       `gorm:"column:is_del" json:"del"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

func (s *SkuDO) TableName() string {
	return "wechat_mall_sku"
}

// GetSkuList 查询商品SKU（含删除态）
func GetSkuList(ctx context.Context) ([]*SkuDO, error) {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return nil, err
	}
	empty := new(SkuDO)
	skuList := make([]*SkuDO, 0)
	err = db.Table(empty.TableName()).Find(&skuList).Error
	if err != nil {
		return nil, err
	}
	return skuList, nil
}

// UpdateSkuStockById 更新SKU库存
func UpdateSkuStockById(ctx context.Context, id, stock int) error {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return err
	}
	sku := &SkuDO{
		Id:         id,
		Stock:      stock,
		UpdateTime: time.Now(),
	}
	err = db.Table(sku.TableName()).Updates(sku).Error
	return err
}
