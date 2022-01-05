package model

import (
	"App-CloudBase-mcdull-mall/pkg/database"
	"context"
)

// GridCategoryDO 宫格模型
type GridCategoryDO struct {
	Id         int    `gorm:"column:id" json:"id"`
	Title      string `gorm:"column:title" json:"title"`             // 标题
	Name       string `gorm:"column:name" json:"name"`               // 名称
	CategoryId int    `gorm:"column:category_id" json:"category_id"` // 分类ID
	Picture    string `gorm:"column:picture" json:"picture"`         // 图标
	Del        int    `gorm:"column:is_del" json:"del"`
}

func (g *GridCategoryDO) TableName() string {
	return "wechat_mall_grid_category"
}

// GetGridCategoryList 查询宫格记录
func GetGridCategoryList(ctx context.Context) ([]*GridCategoryDO, error) {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return nil, err
	}
	emtpy := new(GridCategoryDO)
	cList := make([]*GridCategoryDO, 0)
	err = db.Table(emtpy.TableName()).Where("is_del = 0").Find(&cList).Error
	if err != nil {
		return nil, err
	}
	return cList, nil
}

// GetAllGridCategoryList 查询宫格记录
func GetAllGridCategoryList(ctx context.Context) ([]*GridCategoryDO, error) {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return nil, err
	}
	emtpy := new(GridCategoryDO)
	cList := make([]*GridCategoryDO, 0)
	err = db.Table(emtpy.TableName()).Find(&cList).Error
	if err != nil {
		return nil, err
	}
	return cList, nil
}
