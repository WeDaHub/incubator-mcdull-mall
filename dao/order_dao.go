package dao

import (
	"context"
	"time"

	"gorm.io/gorm"

	"App-CloudBase-mcdull-mall/model"
)

// AddOrder 新增订单
func AddOrder(ctx context.Context, do *model.OrderDO) error {
	return model.AddOrder(ctx, do)
}

// ListOrderByParams 查询用户订单
func ListOrderByParams(ctx context.Context, userId, status int) ([]*model.OrderDO, error) {
	return model.ListOrderByParams(ctx, userId, status)
}

// QueryOrderGoods 查询订单商品
func QueryOrderGoods(ctx context.Context, orderNo string) ([]*model.OrderGoodsDO, error) {
	return model.QueryOrderGoods(ctx, orderNo)
}

// QueryOrderByOrderNo 查询订单详情
func QueryOrderByOrderNo(ctx context.Context, orderNo string) (*model.OrderDO, error) {
	do, err := model.QueryOrderByOrderNo(ctx, orderNo)
	if err == gorm.ErrRecordNotFound {
		return nil, NotFoundRecord
	}
	if err != nil {
		return nil, err
	}
	return do, nil
}

// QueryOrderRefundRecord 查询订单退款记录
func QueryOrderRefundRecord(ctx context.Context, orderNo string) (*model.OrderRefund, error) {
	record, err := model.QueryOrderRefundRecord(ctx, orderNo)
	if err == gorm.ErrRecordNotFound {
		return nil, NotFoundRecord
	}
	if err != nil {
		return nil, err
	}
	return record, nil
}

// QueryOrderById 查询订单信息
func QueryOrderById(ctx context.Context, id int) (*model.OrderDO, error) {
	do, err := model.QueryOrderById(ctx, id)
	if err == gorm.ErrRecordNotFound {
		return nil, NotFoundRecord
	}
	if err != nil {
		return nil, err
	}
	return do, nil
}

// UpdateOrderFinishStauts 更新订单完成状态
func UpdateOrderFinishStauts(ctx context.Context, order *model.OrderDO) error {
	return model.UpdateOrder(ctx, order)
}

// UpdateOrderPayStauts 更新订单支付状态
func UpdateOrderPayStauts(ctx context.Context, order *model.OrderDO) error {
	return model.UpdateOrder(ctx, order)
}

// UpdateOrderStatus 更新订单状态
func UpdateOrderStatus(ctx context.Context, orderId, status int) error {
	do := &model.OrderDO{
		Id:         orderId,
		Status:     status,
		UpdateTime: time.Now(),
	}
	return model.UpdateOrder(ctx, do)
}

// DeleteOrderRecord 删除订单
func DeleteOrderRecord(ctx context.Context, orderId int) error {
	do := &model.OrderDO{
		Id:         orderId,
		Del:        1,
		UpdateTime: time.Now(),
	}
	return model.UpdateOrder(ctx, do)
}

// QueryRefundRecord 查询退款单
func QueryRefundRecord(ctx context.Context, refundNo string) (*model.OrderRefund, error) {
	record, err := model.QueryRefundRecord(ctx, refundNo)
	if err == gorm.ErrRecordNotFound {
		return nil, NotFoundRecord
	}
	if err != nil {
		return nil, err
	}
	return record, nil
}

func CountOrderNum(ctx context.Context, userId, status int) (int, error) {
	return model.CountOrderNum(ctx, userId, status)
}

// UpdateRefundApply 更新退款记录
func UpdateRefundApply(ctx context.Context, id int, status int) error {
	do := &model.OrderRefund{
		Id:         id,
		Status:     status,
		UpdateTime: time.Now(),
	}
	return model.UpdateRefundApply(ctx, do)
}

// AddRefundRecord 新增退款记录
func AddRefundRecord(ctx context.Context, do *model.OrderRefund) (int, error) {
	return model.AddRefundRecord(ctx, do)
}
