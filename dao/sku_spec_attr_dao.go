package dao

import (
	"App-CloudBase-mcdull-mall/model"
	"context"
	"log"
)

// getSkuSpecAttrList 查询SKU规格属性
func getSkuSpecAttrList(ctx context.Context) ([]*model.SkuSpecAttrDO, error) {
	return model.GetSkuSpecAttrList(ctx)
}

const (
	skuSpecAttrCacheKey    = "sku_spec_attr_key"
	skuSpecAttrCacheExpire = 300 // 缓存时间，单位：秒
)

// GetSkuSpecAttrListCache 查询SKU规格属性（缓存）
func GetSkuSpecAttrListCache(ctx context.Context) ([]*model.SkuSpecAttrDO, error) {
	cList := make([]*model.SkuSpecAttrDO, 0)
	if val, ok := GetLocalCache(skuSpecAttrCacheKey); ok {
		return val.([]*model.SkuSpecAttrDO), nil
	}
	cList, err := getSkuSpecAttrList(ctx)
	if err != nil {
		log.Printf("getSkuSpecAttrList failed, err:%v", err)
		return nil, err
	}
	SetLoadCache(skuSpecAttrCacheKey, cList, skuSpecAttrCacheExpire)
	return cList, nil
}
