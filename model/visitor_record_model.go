package model

import (
	"App-CloudBase-mcdull-mall/pkg/database"
	"context"
	"time"
)

// VisitorRecordDO 访问记录
type VisitorRecordDO struct {
	Id         int       `gorm:"column:id" json:"id"`
	UserId     int       `gorm:"column:user_id" json:"user_id"`
	Ip         string    `gorm:"column:ip" json:"ip"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

func (v *VisitorRecordDO) TableName() string {
	return "wechat_mall_visitor_record"
}

// CreateVisitorRecord 增加访客记录
func CreateVisitorRecord(ctx context.Context, do *VisitorRecordDO) error {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return err
	}
	return db.Table(do.TableName()).Create(do).Error
}
