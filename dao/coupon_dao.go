package dao

import (
	"App-CloudBase-mcdull-mall/model"
	"context"
	"log"
)

// getCouponList 查询优惠券（含删除态）
func getCouponList(ctx context.Context) ([]*model.CouponDO, error) {
	return model.GetCouponList(ctx)
}

const (
	couponCacheKey    = "coupon_key"
	couponCacheExpire = 300 // 缓存时间，单位：秒
)

// GetCouponListCache 查询优惠券（缓存）
func GetCouponListCache(ctx context.Context) ([]*model.CouponDO, error) {
	cList := make([]*model.CouponDO, 0)
	if val, ok := GetLocalCache(couponCacheKey); ok {
		return val.([]*model.CouponDO), nil
	}
	cList, err := getCouponList(ctx)
	if err != nil {
		log.Printf("getCouponList failed, err:%v", err)
		return nil, err
	}
	SetLoadCache(couponCacheKey, cList, couponCacheExpire)
	return cList, nil
}

// GetCouponById 通过ID查询优惠券
func GetCouponById(ctx context.Context, cid int) (*model.CouponDO, error) {
	cList, err := GetCouponListCache(ctx)
	if err != nil {
		log.Printf("call GetCouponListCache failed, err:%v", err)
		return nil, err
	}
	var coupon *model.CouponDO
	for _, v := range cList {
		if v.Id == cid {
			coupon = v
		}
	}
	if coupon == nil {
		return nil, NotFoundRecord
	}
	return coupon, nil
}
