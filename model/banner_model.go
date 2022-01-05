package model

import (
	"App-CloudBase-mcdull-mall/pkg/database"
	"context"
)

// BannerDO Banner模型
type BannerDO struct {
	Id           int    `gorm:"column:id" json:"id"`
	Picture      string `gorm:"column:picture" json:"picture"`
	Name         string `gorm:"column:name" json:"name"`
	BusinessType int    `gorm:"column:business_type" json:"business_type"`
	BusinessId   int    `gorm:"column:business_id" json:"business_id"`
	Status       int    `gorm:"column:status" json:"status"`
	Del          int    `gorm:"column:is_del" json:"del"`
}

func (b *BannerDO) TableName() string {
	return "wechat_mall_banner"
}

// GetBannerList 查询Banner列表
func GetBannerList(ctx context.Context, status int) ([]*BannerDO, error) {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return nil, err
	}
	bannerList := make([]*BannerDO, 0)
	err = db.Where("is_del = 0 AND status = ?", status).Find(&bannerList).Error
	if err != nil {
		return nil, err
	}
	return bannerList, nil
}
