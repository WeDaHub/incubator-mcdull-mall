package dao

import (
	"App-CloudBase-mcdull-mall/model"
	"context"
	"errors"
	"log"
)

// getGoodsList 查询商品（含删除态）
func getGoodsList(ctx context.Context) ([]*model.GoodsDO, error) {
	return model.GetGoodsList(ctx)
}

const (
	goodsCacheKey    = "goods_key"
	goodsCacheExpire = 300 // 缓存时间，单位：秒
)

// GetGoodsListCache 查询商品集合缓存
func GetGoodsListCache(ctx context.Context) ([]*model.GoodsDO, error) {
	cList := make([]*model.GoodsDO, 0)
	if val, ok := GetLocalCache(goodsCacheKey); ok {
		return val.([]*model.GoodsDO), nil
	}
	cList, err := getGoodsList(ctx)
	if err != nil {
		log.Printf("getGoodsList failed, err:%v", err)
		return nil, err
	}
	SetLoadCache(goodsCacheKey, cList, goodsCacheExpire)
	return cList, nil
}

var NotFoundRecord = errors.New("not found record")

// GetGoodsById 查询商品信息（含删除态）
func GetGoodsById(ctx context.Context, goodsId int) (*model.GoodsDO, error) {
	goodsList, err := GetGoodsListCache(ctx)
	if err != nil {
		log.Printf("GetGoodsListCache failed, err:%v", err)
		return nil, err
	}
	var goods *model.GoodsDO
	for _, v := range goodsList {
		if v.Id == goodsId {
			goods = v
		}
	}
	if goods == nil {
		return nil, NotFoundRecord
	}
	return goods, nil
}

// UpdateGoodsSaleNum 更新商品销量
func UpdateGoodsSaleNum(ctx context.Context, goodsId, saleNum int) error {
	return model.UpdateGoodsSaleNum(ctx, goodsId, saleNum)
}
