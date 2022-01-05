package dao

import (
	"App-CloudBase-mcdull-mall/model"
	"context"
	"log"
)

// getGridCategoryList 查询首页宫格
func getGridCategoryList(ctx context.Context) ([]*model.GridCategoryDO, error) {
	return model.GetGridCategoryList(ctx)
}

const (
	gridCategoryKey         = "grid_category_key"
	gridCategoryCacheExpire = 300 // 缓存时间，单位：秒
)

// GetGridCategoryListCache 查询宫格列表缓存
func GetGridCategoryListCache(ctx context.Context) ([]*model.GridCategoryDO, error) {
	cList := make([]*model.GridCategoryDO, 0)
	if val, ok := GetLocalCache(gridCategoryKey); ok {
		return val.([]*model.GridCategoryDO), nil
	}
	cList, err := getGridCategoryList(ctx)
	if err != nil {
		log.Printf("getGridCategoryList failed, err:%v", err)
		return nil, err
	}
	SetLoadCache(gridCategoryKey, cList, gridCategoryCacheExpire)
	return cList, nil
}
