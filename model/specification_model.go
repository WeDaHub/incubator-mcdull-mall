package model

import (
	"App-CloudBase-mcdull-mall/pkg/database"
	"context"
)

// SpecificationDO 商城-规格
type SpecificationDO struct {
	Id          int    `gorm:"column:id" json:"id"`
	Name        string `gorm:"column:name" json:"name"`               // 名称
	Description string `gorm:"column:description" json:"description"` // 描述
	Unit        string `gorm:"column:unit" json:"unit"`               // 单位
	Standard    int    `gorm:"column:standard" json:"standard"`       // 是否标准
	Del         int    `gorm:"column:is_del" json:"del"`
}

func (s *SpecificationDO) TableName() string {
	return "wechat_mall_specification"
}

// GetSpecList 查询商城规格
func GetSpecList(ctx context.Context) ([]*SpecificationDO, error) {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return nil, err
	}
	empty := new(SpecificationDO)
	doList := make([]*SpecificationDO, 0)
	err = db.Table(empty.TableName()).Where("is_del = 0").Find(&doList).Error
	if err != nil {
		return nil, err
	}
	return doList, nil
}
