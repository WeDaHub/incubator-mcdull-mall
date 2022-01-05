package dao

import (
	"App-CloudBase-mcdull-mall/model"
	"context"
)

// GetTakeCouponNum 查询优惠券领取的数量
func GetTakeCouponNum(ctx context.Context, cid int) (int, error) {
	return model.GetTakeCouponNum(ctx, cid)
}

// GetUserCouponLog 查询用户领取的优惠券
func GetUserCouponLog(ctx context.Context, userId int) ([]*model.CouponLogDO, error) {
	return model.GetUserCouponLog(ctx, userId)
}

// GetUserCouponLogById 通过ID查询优惠券记录
func GetUserCouponLogById(ctx context.Context, userId, logId int) (*model.CouponLogDO, error) {
	cList, err := GetUserCouponLog(ctx, userId)
	if err != nil {
		return nil, err
	}
	var cLog *model.CouponLogDO
	for _, v := range cList {
		if v.Id == logId {
			cLog = v
		}
	}
	if cLog == nil {
		return nil, NotFoundRecord
	}
	return cLog, nil
}

// UpdateCouponLog 更新优惠券记录
func UpdateCouponLog(ctx context.Context, do *model.CouponLogDO) error {
	return model.UpdateCouponLog(ctx, do)
}

// GetUserTakeCouponNum 查询优惠券个人领取的数量
func GetUserTakeCouponNum(ctx context.Context, cid, userid int) (int, error) {
	return model.GetUserTakeCouponNum(ctx, cid, userid)
}

// AddCouponLog 添加优惠券领取记录
func AddCouponLog(ctx context.Context, do *model.CouponLogDO) error {
	return model.AddCouponLog(ctx, do)
}
