package dao

import (
	"App-CloudBase-mcdull-mall/model"
	"context"
	"fmt"
	"log"
)

// getBannerList 查询Banner集合
func getBannerList(ctx context.Context, status int) ([]*model.BannerDO, error) {
	return model.GetBannerList(ctx, status)
}

const (
	bannerCacheKeyPrefix = "banner_status_"
	bannerCacheExpire    = 300 // 缓存时间，单位：秒
)

// GetBannerListCache 查询Banner集合缓存
func GetBannerListCache(ctx context.Context, status int) ([]*model.BannerDO, error) {
	cacheKey := getBannerCacheKey(status)
	bList := make([]*model.BannerDO, 0)
	if val, ok := GetLocalCache(cacheKey); ok {
		return val.([]*model.BannerDO), nil
	}
	bList, err := getBannerList(ctx, status)
	if err != nil {
		log.Printf("GetBannerList failed, err:%v", err)
		return nil, err
	}
	SetLoadCache(cacheKey, bList, bannerCacheExpire)
	return bList, nil
}

func getBannerCacheKey(status int) string {
	return fmt.Sprintf("%s%d", bannerCacheKeyPrefix, status)
}
