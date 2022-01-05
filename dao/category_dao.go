package dao

import (
	"App-CloudBase-mcdull-mall/defs"
	"App-CloudBase-mcdull-mall/model"
	"context"
	"log"
)

// getCategoryList 查询分类分类
func getCategoryList(ctx context.Context) ([]*model.CategoryDO, error) {
	return model.GetCategoryList(ctx, defs.ONLINE)
}

const (
	categoryCacheKey    = "category_key"
	categoryCacheExpire = 300 // 缓存时间，单位：秒
)

// GetCategoryListCache 查询分类集合缓存
func GetCategoryListCache(ctx context.Context) ([]*model.CategoryDO, error) {
	cList := make([]*model.CategoryDO, 0)
	if val, ok := GetLocalCache(categoryCacheKey); ok {
		return val.([]*model.CategoryDO), nil
	}
	cList, err := getCategoryList(ctx)
	if err != nil {
		log.Printf("getCategoryList failed, err:%v", err)
		return nil, err
	}
	SetLoadCache(categoryCacheKey, cList, categoryCacheExpire)
	return cList, nil
}
