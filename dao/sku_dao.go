package dao

import (
	"App-CloudBase-mcdull-mall/defs"
	"App-CloudBase-mcdull-mall/model"
	"context"
	"log"
)

// getSkuList 查询商品SKU（含删除态）
func getSkuList(ctx context.Context) ([]*model.SkuDO, error) {
	return model.GetSkuList(ctx)
}

const (
	skuCacheKey    = "sku_key"
	skuCacheExpire = 300 // 缓存时间，单位：秒
)

// GetSkuListCache 查询商品SKU（缓存）
func GetSkuListCache(ctx context.Context) ([]*model.SkuDO, error) {
	cList := make([]*model.SkuDO, 0)
	if val, ok := GetLocalCache(skuCacheKey); ok {
		return val.([]*model.SkuDO), nil
	}
	cList, err := getSkuList(ctx)
	if err != nil {
		log.Printf("getSkuList failed, err:%v", err)
		return nil, err
	}
	SetLoadCache(skuCacheKey, cList, skuCacheExpire)
	return cList, nil
}

// GetOnlineSkuListByGoodsId 查询商品SKU（仅上架）
func GetOnlineSkuListByGoodsId(ctx context.Context, goodsId int) ([]*model.SkuDO, error) {
	skuList, err := GetSkuListCache(ctx)
	if err != nil {
		return nil, err
	}
	newList := make([]*model.SkuDO, 0)
	for _, v := range skuList {
		if v.Del == defs.DELETE || v.Online == defs.OFFLINE {
			continue
		}
		if v.GoodsId == goodsId {
			newList = append(newList, v)
		}
	}
	return newList, nil
}

// GetSkuListBySkuId 查询SKU（含删除态）
func GetSkuListBySkuId(ctx context.Context, skuId int) (*model.SkuDO, error) {
	skuList, err := GetSkuListCache(ctx)
	if err != nil {
		log.Printf("call GetSkuListCache failed, err:%v", err)
		return nil, err
	}
	var skuDO *model.SkuDO
	for _, v := range skuList {
		if skuId == v.Id {
			skuDO = v
		}
	}
	if skuDO == nil {
		return nil, NotFoundRecord
	}
	return skuDO, nil
}

// UpdateSkuStockById 更新SKU库存
func UpdateSkuStockById(ctx context.Context, id, stock int) error {
	return model.UpdateSkuStockById(ctx, id, stock)
}
