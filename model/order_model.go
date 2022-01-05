package model

import (
	"App-CloudBase-mcdull-mall/defs"
	"App-CloudBase-mcdull-mall/pkg/database"
	"context"
	"time"
)

// OrderDO 商城-订单表
type OrderDO struct {
	Id              int       `gorm:"column:id" json:"id"`
	OrderNo         string    `gorm:"column:order_no" json:"order_no"`                 // 订单号
	UserId          int       `gorm:"column:user_id" json:"user_id"`                   // 用户ID
	PayAmount       string    `gorm:"column:pay_amount" json:"pay_amount"`             // 付款金额
	GoodsAmount     string    `gorm:"column:goods_amount" json:"goods_amount"`         // 商品小计
	DiscountAmount  string    `gorm:"column:discount_amount" json:"discount_amount"`   // 优惠金额
	DispatchAmount  string    `gorm:"column:dispatch_amount" json:"dispatch_amount"`   // 运费
	PayTime         time.Time `gorm:"column:pay_time" json:"pay_time"`                 // 支付时间
	DeliverTime     time.Time `gorm:"column:deliver_time" json:"deliver_time"`         // 发货时间
	FinishTime      time.Time `gorm:"column:finish_time" json:"finish_time"`           // 完成时间
	Status          int       `gorm:"column:status" json:"status"`                     // 状态 -1 已取消 0-待付款 1-待发货 2-待收货 3-已完成 4-（待发货）退款申请 5-已退款
	AddressId       int       `gorm:"column:address_id" json:"address_id"`             // 地址ID
	AddressSnapshot string    `gorm:"column:address_snapshot" json:"address_snapshot"` // 地址快照
	WxappPrepayId   string    `gorm:"column:wxapp_prepay_id" json:"wxapp_prepay_id"`   // 微信预支付ID
	TransactionId   string    `gorm:"column:transaction_id" json:"transaction_id"`     // 微信支付单号
	Remark          string    `gorm:"column:remark" json:"remark"`                     // 订单备注
	Del             int       `gorm:"column:is_del" json:"del"`
	CreateTime      time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime      time.Time `gorm:"column:update_time" json:"update_time"`
}

func (o *OrderDO) TableName() string {
	return "wechat_mall_order"
}

// AddOrder 新增订单
func AddOrder(ctx context.Context, do *OrderDO) error {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return err
	}
	return db.Table(do.TableName()).Create(do).Error
}

// ListOrderByParams 查询用户订单
func ListOrderByParams(ctx context.Context, userId, status int) ([]*OrderDO, error) {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return nil, err
	}
	empty := new(OrderDO)
	oList := make([]*OrderDO, 0)
	if status == defs.ALL {
		err = db.Table(empty.TableName()).Where("is_del = 0 AND user_id = ?", userId).Order("id DESC").Find(&oList).Error
	} else {
		err = db.Table(empty.TableName()).Where("is_del = 0 AND user_id = ? AND status = ?", userId, status).Order("id DESC").Find(&oList).Error
	}
	if err != nil {
		return nil, err
	}
	return oList, nil
}

// QueryOrderByOrderNo 查询订单详情
func QueryOrderByOrderNo(ctx context.Context, orderNo string) (*OrderDO, error) {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return nil, err
	}
	out := new(OrderDO)
	err = db.Table(out.TableName()).Where("order_no = ?", orderNo).Take(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryOrderById 查询订单信息
func QueryOrderById(ctx context.Context, id int) (*OrderDO, error) {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return nil, err
	}
	out := new(OrderDO)
	err = db.Table(out.TableName()).Where("id = ?", id).Take(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateOrder 更新订单
func UpdateOrder(ctx context.Context, do *OrderDO) error {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return err
	}
	return db.Table(do.TableName()).Updates(do).Error
}

// CountOrderNum 统计订单数量
func CountOrderNum(ctx context.Context, userId, status int) (int, error) {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return 0, err
	}
	empty := new(OrderDO)
	num := 0
	err = db.Table(empty.TableName()).Select("COUNT(id)").Where("is_del = 0 AND user_id = ? AND status = ?", userId, status).Find(&num).Error
	if err != nil {
		return 0, err
	}
	return num, nil
}
