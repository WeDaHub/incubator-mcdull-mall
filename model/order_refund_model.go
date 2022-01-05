package model

import (
	"App-CloudBase-mcdull-mall/pkg/database"
	"context"
	"time"
)

// OrderRefund 订单退款
type OrderRefund struct {
	Id           int       `gorm:"column:id" json:"id"`
	RefundNo     string    `gorm:"column:refund_no" json:"refund_no"`         // 退款编号
	UserId       int       `gorm:"column:user_id" json:"user_id"`             // 用户ID
	OrderNo      string    `gorm:"column:order_no" json:"order_no"`           // 订单编号
	Reason       string    `gorm:"column:reason" json:"reason"`               // 退款原因
	RefundAmount string    `gorm:"column:refund_amount" json:"refund_amount"` // 退款金额
	Status       int       `gorm:"column:status" json:"status"`               // 状态：0-退款申请 1-商家处理申请 2-退款完成
	Del          int       `gorm:"column:is_del" json:"del"`                  // 是否删除：0-否 1-是
	RefundTime   time.Time `gorm:"column:refund_time" json:"refund_time"`     // 退款时间
	CreateTime   time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`
}

func (o *OrderRefund) TableName() string {
	return "wechat_mall_order_refund"
}

// QueryOrderRefundRecord 查询订单退款记录
func QueryOrderRefundRecord(ctx context.Context, orderNo string) (*OrderRefund, error) {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return nil, err
	}
	out := new(OrderRefund)
	err = db.Table(out.TableName()).Where("order_no = ?", orderNo).Take(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryRefundRecord 查询退款单
func QueryRefundRecord(ctx context.Context, refundNo string) (*OrderRefund, error) {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return nil, err
	}
	out := new(OrderRefund)
	err = db.Table(out.TableName()).Where("refund_no = ?", refundNo).Take(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateRefundApply 更新退款记录
func UpdateRefundApply(ctx context.Context, do *OrderRefund) error {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return err
	}
	return db.Table(do.TableName()).Updates(do).Error
}

// AddRefundRecord 新增退款记录
func AddRefundRecord(ctx context.Context, do *OrderRefund) (int, error) {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return 0, err
	}
	err = db.Table(do.TableName()).Create(do).Error
	if err != nil {
		return 0, err
	}
	return do.Id, nil
}
