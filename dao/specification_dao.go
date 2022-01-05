package dao

import (
	"App-CloudBase-mcdull-mall/model"
	"context"
	"log"
)

// getSpecList 查询商城规格
func getSpecList(ctx context.Context) ([]*model.SpecificationDO, error) {
	return model.GetSpecList(ctx)
}

const (
	specificationCacheKey    = "specification_key"
	specificationCacheExpire = 300 // 缓存时间，单位：秒
)

// GetSpecListCache 查询商城规格（缓存）
func GetSpecListCache(ctx context.Context) ([]*model.SpecificationDO, error) {
	cList := make([]*model.SpecificationDO, 0)
	if val, ok := GetLocalCache(specificationCacheKey); ok {
		return val.([]*model.SpecificationDO), nil
	}
	cList, err := getSpecList(ctx)
	if err != nil {
		log.Printf("getSpecList failed, err:%v", err)
		return nil, err
	}
	SetLoadCache(specificationCacheKey, cList, specificationCacheExpire)
	return cList, nil
}
