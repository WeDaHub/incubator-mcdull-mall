package service

import (
	"App-CloudBase-mcdull-mall/dao"
	"App-CloudBase-mcdull-mall/defs"
	"App-CloudBase-mcdull-mall/model"
	"App-CloudBase-mcdull-mall/utils"
	"context"
	"log"
	"time"
)

// GetOnlineCouponList 查询优惠券列表（上架）
func GetOnlineCouponList(ctx context.Context, page, size int) ([]*model.CouponDO, int, error) {
	// 1.查询优惠券
	cList, err := dao.GetCouponListCache(ctx)
	if err != nil {
		log.Printf("call GetCouponListCache failed, err:%v", err)
		return nil, 0, err
	}
	// 2.过滤
	newList := make([]*model.CouponDO, 0)
	for _, v := range cList {
		if v.Del == defs.DELETE || v.Online == defs.OFFLINE {
			continue
		}
		newList = append(newList, v)
	}
	// 3.分页
	total := len(newList)
	offset, endpos := pagePos(total, page, size)
	return newList[offset:endpos], total, nil
}

// GetOnlineCouponById 查询优惠券（上架）
func GetOnlineCouponById(ctx context.Context, cid int) (*model.CouponDO, error) {
	couponDO, err := dao.GetCouponById(ctx, cid)
	if err != nil {
		log.Printf("call GetCouponById failed, err:%v", err)
		return nil, err
	}
	// 剔除下架
	if couponDO.Del == defs.DELETE || couponDO.Online == defs.OFFLINE {
		return nil, dao.NotFoundRecord
	}
	return couponDO, nil
}

// GetCouponById 查询优惠券（含删除态）
func GetCouponById(ctx context.Context, cid int) (*model.CouponDO, error) {
	couponDO, err := dao.GetCouponById(ctx, cid)
	if err != nil {
		log.Printf("call GetCouponById failed, err:%v", err)
		return nil, err
	}
	return couponDO, nil
}

// GetTakeCouponNum 查询优惠券的领取数量
func GetTakeCouponNum(ctx context.Context, cid int) (int, error) {
	num, err := dao.GetTakeCouponNum(ctx, cid)
	return num, err
}

// GetUserTakeCouponNum 查询优惠券个人领取的数量
func GetUserTakeCouponNum(ctx context.Context, cid, userId int) (int, error) {
	num, err := dao.GetUserTakeCouponNum(ctx, cid, userId)
	return num, err
}

// RecordCouponLog 记录优惠券领取
func RecordCouponLog(ctx context.Context, userId int, coupon *model.CouponDO) error {
	couponLog := &model.CouponLogDO{}
	couponLog.CouponId = coupon.Id
	couponLog.UserId = userId
	couponLog.UseTime = time.Now()
	couponLog.ExpireTime = coupon.EndTime
	couponLog.Status = 0
	couponLog.Code = utils.RandomNumberStr(12)
	couponLog.OrderNo = ""
	couponLog.CreateTime = couponLog.UseTime
	couponLog.UpdateTime = couponLog.UseTime
	err := dao.AddCouponLog(ctx, couponLog)
	if err != nil {
		return err
	}
	return nil
}

// GetUserCouponLog 查询用户领取的优惠券
func GetUserCouponLog(ctx context.Context, userId, status, page, size int) ([]*model.CouponLogDO, int, error) {
	// 1.查询用户优惠券
	cList, err := dao.GetUserCouponLog(ctx, userId)
	if err != nil {
		log.Printf("call GetUserCouponLog failed, err:%v", err)
		return nil, 0, err
	}
	// 2.过滤
	idList := make([]int, 0)
	newList := make([]*model.CouponLogDO, 0)
	for _, v := range cList {
		if v.Status != status {
			continue
		}
		// 未使用的优惠券，检查时间
		if v.Status == 0 {
			// 过期
			if v.ExpireTime.Before(time.Now()) {
				idList = append(idList, v.Id)
				continue
			}
		}
		newList = append(newList, v)
	}
	// 3.刷优惠券状态（异步）
	go refreshCouponLogStatus(ctx, userId, idList)
	// 4.分页
	total := len(newList)
	offset, endpos := pagePos(total, page, size)
	return newList[offset:endpos], total, nil
}

// refreshCouponLogStatus 刷新过期的优惠券状态
func refreshCouponLogStatus(ctx context.Context, userId int, idList []int) {
	for _, id := range idList {
		// 1.查询记录
		record, err := dao.GetUserCouponLogById(ctx, userId, id)
		if err != nil {
			log.Printf("call GetUserCouponLogById failed, err:%v", err)
			continue
		}
		// 2.更新状态
		record.Status = 2
		record.UpdateTime = time.Now()
		err = dao.UpdateCouponLog(ctx, record)
		if err != nil {
			log.Printf("call UpdateCouponLog failed, err:%v", err)
		}
	}
}

// DeleteCouponLog 删除优惠券
func DeleteCouponLog(ctx context.Context, userId, logId, status int) error {
	// 1.查询记录
	record, err := dao.GetUserCouponLogById(ctx, userId, logId)
	if err != nil {
		log.Printf("call GetUserCouponLogById failed, err:%v", err)
		return err
	}
	// 2.更新记录状态
	record.Del = defs.DELETE
	record.UpdateTime = time.Now()
	err = dao.UpdateCouponLog(ctx, record)
	if err != nil {
		log.Printf("call UpdateCouponLog failed, err:%v", err)
		return err
	}
	return nil
}
