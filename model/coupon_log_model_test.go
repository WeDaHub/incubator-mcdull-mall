package model

import (
	"App-CloudBase-mcdull-mall/utils"
	"context"
	"fmt"
	"testing"
	"time"
)

func Test_GetTakeCouponNum(t *testing.T) {
	ctx := context.Background()
	num, err := GetTakeCouponNum(ctx, 1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(num)
}

func Test_GetUserCouponLog(t *testing.T) {
	ctx := context.Background()
	cList, err := GetUserCouponLog(ctx, 1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(cList))
}

func Test_UpdateCouponLog(t *testing.T) {
	ctx := context.Background()
	cLog := &CouponLogDO{
		Id:         1,
		Status:     3,
		UpdateTime: time.Now(),
	}
	err := UpdateCouponLog(ctx, cLog)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_GetUserTakeCouponNum(t *testing.T) {
	ctx := context.Background()
	num, err := GetUserTakeCouponNum(ctx, 2, 1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(num)
}

func Test_AddCouponLog(t *testing.T) {
	ctx := context.Background()
	couponLog := &CouponLogDO{}
	couponLog.CouponId = 1
	couponLog.UserId = 2
	couponLog.UseTime = time.Now()
	couponLog.ExpireTime = time.Now()
	couponLog.Status = 0
	couponLog.Code = utils.RandomNumberStr(12)
	couponLog.OrderNo = "1234"
	couponLog.CreateTime = couponLog.UseTime
	couponLog.UpdateTime = couponLog.UseTime
	err := AddCouponLog(ctx, couponLog)
	if err != nil {
		t.Fatal(err)
	}
}
