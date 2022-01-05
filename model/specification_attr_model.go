package model

import (
	"App-CloudBase-mcdull-mall/pkg/database"
	"context"
)

// SpecificationAttrDO 商城-规格属性表
type SpecificationAttrDO struct {
	Id     int    `gorm:"column:id" json:"id"`
	SpecId int    `gorm:"column:id" json:"spec_id"` // 规格主键
	Value  string `gorm:"column:id" json:"value"`   // 规格值
	Extend string `gorm:"column:id" json:"extend"`  // 扩展
	Del    int    `gorm:"column:id" json:"del"`
}

func (s *SpecificationAttrDO) TableName() string {
	return "wechat_mall_specification_attr"
}

// GetSpecAttrList 查询商城规格属性
func GetSpecAttrList(ctx context.Context) ([]*SpecificationAttrDO, error) {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return nil, err
	}
	empty := new(SpecificationAttrDO)
	slist := make([]*SpecificationAttrDO, 0)
	err = db.Table(empty.TableName()).Where("is_del = 0").Find(&slist).Error
	if err != nil {
		return nil, err
	}
	return slist, nil
}
