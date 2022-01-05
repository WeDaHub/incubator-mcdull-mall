package model

import (
	"App-CloudBase-mcdull-mall/pkg/database"
	"context"
	"time"
)

// CouponLogDO 商城-优惠券领取记录
type CouponLogDO struct {
	Id         int       `gorm:"column:id" json:"id"`
	CouponId   int       `gorm:"column:coupon_id" json:"coupon_id"`     // 优惠券ID
	UserId     int       `gorm:"column:user_id" json:"user_id"`         // 用户ID
	UseTime    time.Time `gorm:"column:use_time" json:"use_time"`       // 核销时间
	ExpireTime time.Time `gorm:"column:expire_time" json:"expire_time"` // 过期时间
	Status     int       `gorm:"column:status" json:"status"`           // 状态：0-未使用 1-已使用 2-已过期
	Code       string    `gorm:"column:code" json:"code"`               // 券码
	OrderNo    string    `gorm:"column:order_no" json:"order_no"`       // 核销的订单号
	Del        int       `gorm:"column:is_del" json:"del"`              // 是否删除
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

func (c *CouponLogDO) TableName() string {
	return "wechat_mall_coupon_log"
}

// GetTakeCouponNum 查询优惠券领取的数量
func GetTakeCouponNum(ctx context.Context, cid int) (int, error) {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return 0, err
	}
	empty := new(CouponLogDO)
	num := 0
	err = db.Table(empty.TableName()).Select("COUNT(id)").Where("coupon_id = ?", cid).Find(&num).Error
	if err != nil {
		return 0, err
	}
	return num, nil
}

// GetUserCouponLog 查询用户领取的优惠券
func GetUserCouponLog(ctx context.Context, userId int) ([]*CouponLogDO, error) {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return nil, err
	}
	empty := new(CouponLogDO)
	cList := make([]*CouponLogDO, 0)
	err = db.Table(empty.TableName()).Where("is_del = 0 AND user_id = ?", userId).Find(&cList).Error
	if err != nil {
		return nil, err
	}
	return cList, nil
}

// UpdateCouponLog 更新优惠券记录
func UpdateCouponLog(ctx context.Context, do *CouponLogDO) error {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return err
	}
	err = db.Table(do.TableName()).Updates(do).Error
	return err
}

// GetUserTakeCouponNum 查询优惠券个人领取的数量
func GetUserTakeCouponNum(ctx context.Context, cid, userid int) (int, error) {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return 0, err
	}
	empty := new(CouponLogDO)
	num := 0
	err = db.Table(empty.TableName()).Select("COUNT(id)").Where("coupon_id = ? AND user_id = ?", cid, userid).Find(&num).Error
	if err != nil {
		return 0, err
	}
	return num, nil
}

// AddCouponLog 添加优惠券领取记录
func AddCouponLog(ctx context.Context, do *CouponLogDO) error {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return err
	}
	err = db.Table(do.TableName()).Create(do).Error
	return err
}
