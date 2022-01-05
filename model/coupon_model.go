package model

import (
	"App-CloudBase-mcdull-mall/pkg/database"
	"context"
	"time"
)

// CouponDO 商城-优惠券
type CouponDO struct {
	Id          int       `gorm:"column:id" json:"id"`
	Title       string    `gorm:"column:title" json:"title"`             // 标题
	FullMoney   string    `gorm:"column:full_money" json:"full_money"`   // 满减额
	Minus       string    `gorm:"column:minus" json:"minus"`             // 优惠金额
	Rate        string    `gorm:"column:rate" json:"rate"`               // 折扣
	Type        int       `gorm:"column:type" json:"type"`               // 类型
	GrantNum    int       `gorm:"column:grant_num" json:"grant_num"`     // 发券数量
	LimitNum    int       `gorm:"column:limit_num" json:"limit_num"`     // 领取数量
	StartTime   time.Time `gorm:"column:start_time" json:"start_time"`   // 开始时间
	EndTime     time.Time `gorm:"column:end_time" json:"end_time"`       // 截止时间
	Description string    `gorm:"column:description" json:"description"` // 规则描述
	Online      int       `gorm:"column:online" json:"online"`           // 是否上线
	Del         int       `gorm:"column:is_del" json:"del"`
}

func (c *CouponDO) TableName() string {
	return "wechat_mall_coupon"
}

// GetCouponList 查询优惠券（含删除态）
func GetCouponList(ctx context.Context) ([]*CouponDO, error) {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return nil, err
	}
	empty := new(CouponDO)
	cList := make([]*CouponDO, 0)
	err = db.Table(empty.TableName()).Where("start_time < now() AND end_time > now()").Find(&cList).Error
	if err != nil {
		return nil, err
	}
	return cList, nil
}
