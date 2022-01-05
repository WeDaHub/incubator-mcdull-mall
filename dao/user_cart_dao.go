package dao

import (
	"App-CloudBase-mcdull-mall/model"
	"context"
	"log"
)

// GetUserCartList 查询用户购物车
func GetUserCartList(ctx context.Context, userId int) ([]*model.UserCartDO, error) {
	return model.GetUserCartList(ctx, userId)
}

// GetUserCartById 通过主键查询
func GetUserCartById(ctx context.Context, userId int, id int) (*model.UserCartDO, error) {
	cList, err := GetUserCartList(ctx, userId)
	if err != nil {
		log.Printf("call GetUserCartList failed, err:%v", err)
		return nil, err
	}
	var record *model.UserCartDO
	for _, v := range cList {
		if v.Id == id {
			record = v
		}
	}
	if record == nil {
		return nil, NotFoundRecord
	}
	return record, nil
}

// GetUserCartListByParams 查询购物车记录
func GetUserCartListByParams(ctx context.Context, userId, goodsId, skuId int) (*model.UserCartDO, error) {
	cList, err := GetUserCartList(ctx, userId)
	if err != nil {
		log.Printf("call GetUserCartList failed, err:%v", err)
		return nil, err
	}
	var record *model.UserCartDO
	for _, v := range cList {
		if v.UserId == userId && v.GoodsId == goodsId && v.SkuId == skuId {
			record = v
		}
	}
	if record == nil {
		return nil, NotFoundRecord
	}
	return record, nil
}

// CreateUserCart 新增购物车商品
func CreateUserCart(ctx context.Context, cartDO *model.UserCartDO) error {
	return model.CreateUserCart(ctx, cartDO)
}

// UpdateUserCart 更新购物车商品
func UpdateUserCart(ctx context.Context, cartDO *model.UserCartDO) error {
	return model.UpdateUserCart(ctx, cartDO)
}
