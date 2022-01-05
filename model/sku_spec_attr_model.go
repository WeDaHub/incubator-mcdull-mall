package model

import (
	"App-CloudBase-mcdull-mall/pkg/database"
	"context"
)

// SkuSpecAttrDO 商城-SKU规格属性
type SkuSpecAttrDO struct {
	Id     int `gorm:"column:id" json:"id"`
	SkuId  int `gorm:"column:sku_id" json:"sku_id"`
	SpecId int `gorm:"column:spec_id" json:"spec_id"`
	AttrId int `gorm:"column:attr_id" json:"attr_id"`
	Del    int `gorm:"column:id" json:"is_del"`
}

func (s *SkuSpecAttrDO) TableName() string {
	return "wechat_mall_sku_spec_attr"
}

// GetSkuSpecAttrList 查询SKU规格属性
func GetSkuSpecAttrList(ctx context.Context) ([]*SkuSpecAttrDO, error) {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return nil, err
	}
	empty := new(SkuSpecAttrDO)
	sList := make([]*SkuSpecAttrDO, 0)
	err = db.Table(empty.TableName()).Where("is_del = 0").Find(&sList).Error
	if err != nil {
		return nil, err
	}
	return sList, nil
}
