package service

import (
	"App-CloudBase-mcdull-mall/dao"
	"App-CloudBase-mcdull-mall/defs"
	"App-CloudBase-mcdull-mall/model"
	"context"
	"log"
)

// GetBannerList 查询首页Banner列表
func GetBannerList(ctx context.Context, page, size int) ([]*model.BannerDO, error) {
	bList, err := dao.GetBannerListCache(ctx, defs.ONLINE)
	if err != nil {
		log.Printf("call GetBannerList failed, err:%v", err)
		return nil, err
	}
	offset, endpos := pagePos(len(bList), page, size)
	return bList[offset:endpos], nil
}

// GetGridCategoryList 查询首页宫格
func GetGridCategoryList(ctx context.Context, page, size int) ([]*model.GridCategoryDO, error) {
	cList, err := dao.GetGridCategoryListCache(ctx)
	if err != nil {
		log.Printf("call GetGridCategoryListCache failed, err:%v", err)
		return nil, err
	}
	offset, endpos := pagePos(len(cList), page, size)
	return cList[offset:endpos], nil
}

// GetSubCategoryList 查询二级分类
func GetSubCategoryList(ctx context.Context) ([]*model.CategoryDO, error) {
	// 1.查询缓存
	cList, err := dao.GetCategoryListCache(ctx)
	if err != nil {
		log.Printf("GetCategoryListCache failed, err:%v", err)
		return nil, err
	}
	// 2.过滤一级分类
	newList := make([]*model.CategoryDO, 0)
	for _, v := range cList {
		if v.ParentId != 0 {
			newList = append(newList, v)
		}
	}
	return newList, nil
}
