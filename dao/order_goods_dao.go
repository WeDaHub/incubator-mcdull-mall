package dao

import (
	"App-CloudBase-mcdull-mall/model"
	"context"
)

// CountBuyUserNum 统计商品购买人数
func CountBuyUserNum(ctx context.Context, goodsId int) (int, error) {
	return model.CountBuyUserNum(ctx, goodsId)
}

// AddOrderGoods 订单商品快照
func AddOrderGoods(ctx context.Context, do *model.OrderGoodsDO) error {
	return model.AddOrderGoods(ctx, do)
}
