package model

import (
	"App-CloudBase-mcdull-mall/pkg/database"
	"context"
)

// GoodsSpecDO 商城-商品规格
type GoodsSpecDO struct {
	Id      int `gorm:"column:id" json:"id"`
	GoodsId int `gorm:"column:goods_id" json:"goods_id"` // 商品ID
	SpecId  int `gorm:"column:spec_id" json:"spec_id"`   // 规格ID
	Del     int `gorm:"column:is_del" json:"del"`
}

func (g *GoodsSpecDO) TableName() string {
	return "wechat_mall_goods_spec"
}

// GetGoodsSpecList 查询商品规格表
func GetGoodsSpecList(ctx context.Context) ([]*GoodsSpecDO, error) {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return nil, err
	}
	empty := new(GoodsSpecDO)
	specList := make([]*GoodsSpecDO, 0)
	err = db.Table(empty.TableName()).Where("is_del = 0").Find(&specList).Error
	if err != nil {
		return nil, err
	}
	return specList, nil
}
