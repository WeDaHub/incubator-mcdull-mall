package model

import (
	"App-CloudBase-mcdull-mall/pkg/database"
	"context"
)

// CategoryDO 商城-分类
type CategoryDO struct {
	Id          int    `gorm:"column:id" json:"id"`
	ParentId    int    `gorm:"column:parent_id" json:"parent_id"`     // 父级分类
	Name        string `gorm:"column:name" json:"name"`               // 分类名称
	Sort        int    `gorm:"column:sort" json:"sort"`               // 排序
	Online      int    `gorm:"column:online" json:"online"`           // 是否上线
	Picture     string `gorm:"column:picture" json:"picture"`         // 图标
	Description string `gorm:"column:description" json:"description"` // 描述
	Del         int    `gorm:"column:is_del" json:"del"`
}

func (c *CategoryDO) TableName() string {
	return "wechat_mall_category"
}

// GetCategoryList 查询分类列表
func GetCategoryList(ctx context.Context, status int) ([]*CategoryDO, error) {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return nil, err
	}
	empty := new(CategoryDO)
	cList := make([]*CategoryDO, 0)
	err = db.Table(empty.TableName()).Where("is_del = 0 AND online = ?", status).Find(&cList).Error
	if err != nil {
		return nil, err
	}
	return cList, nil
}

// GetAllCategoryList 查询分类列表
func GetAllCategoryList(ctx context.Context) ([]*CategoryDO, error) {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return nil, err
	}
	empty := new(CategoryDO)
	cList := make([]*CategoryDO, 0)
	err = db.Table(empty.TableName()).Find(&cList).Error
	if err != nil {
		return nil, err
	}
	return cList, nil
}
