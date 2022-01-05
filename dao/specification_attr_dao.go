package dao

import (
	"App-CloudBase-mcdull-mall/model"
	"context"
	"log"
)

// getSpecAttrList 查询商城规格属性
func getSpecAttrList(ctx context.Context) ([]*model.SpecificationAttrDO, error) {
	return model.GetSpecAttrList(ctx)
}

const (
	specificationAttrCacheKey    = "specification_attr_key"
	specificationAttrCacheExpire = 300 // 缓存时间，单位：秒
)

// GetSpecAttrListCache 查询商城规格属性（缓存）
func GetSpecAttrListCache(ctx context.Context) ([]*model.SpecificationAttrDO, error) {
	cList := make([]*model.SpecificationAttrDO, 0)
	if val, ok := GetLocalCache(specificationAttrCacheKey); ok {
		return val.([]*model.SpecificationAttrDO), nil
	}
	cList, err := getSpecAttrList(ctx)
	if err != nil {
		log.Printf("getSpecAttrList failed, err:%v", err)
		return nil, err
	}
	SetLoadCache(specificationAttrCacheKey, cList, specificationAttrCacheExpire)
	return cList, nil
}
