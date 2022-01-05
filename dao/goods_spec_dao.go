package dao

import (
	"App-CloudBase-mcdull-mall/model"
	"context"
	"log"
)

// getGoodsSpecList 查询商品规格
func getGoodsSpecList(ctx context.Context) ([]*model.GoodsSpecDO, error) {
	return model.GetGoodsSpecList(ctx)
}

const (
	goodsSpecCacheKey    = "goods_spec_key"
	goodsSpecCacheExpire = 300 // 缓存时间，单位：秒
)

// GetGoodsSpecListCache 查询商品规格集合缓存
func GetGoodsSpecListCache(ctx context.Context) ([]*model.GoodsSpecDO, error) {
	cList := make([]*model.GoodsSpecDO, 0)
	if val, ok := GetLocalCache(goodsSpecCacheKey); ok {
		return val.([]*model.GoodsSpecDO), nil
	}
	cList, err := getGoodsSpecList(ctx)
	if err != nil {
		log.Printf("getGoodsSpecList failed, err:%v", err)
		return nil, err
	}
	SetLoadCache(goodsSpecCacheKey, cList, goodsSpecCacheExpire)
	return cList, nil
}

// GetGoodsSpecByGoodsId 查询商品规格
func GetGoodsSpecByGoodsId(ctx context.Context, goodsId int) ([]*model.GoodsSpecDO, error) {
	specList, err := GetGoodsSpecListCache(ctx)
	if err != nil {
		log.Printf("call GetGoodsSpecListCache failed, err:%v", err)
		return nil, err
	}
	newList := make([]*model.GoodsSpecDO, 0)
	for _, v := range specList {
		if v.GoodsId == goodsId {
			newList = append(newList, v)
		}
	}
	return newList, nil
}
