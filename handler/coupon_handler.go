package handler

import (
	"App-CloudBase-mcdull-mall/dao"
	"App-CloudBase-mcdull-mall/defs"
	"App-CloudBase-mcdull-mall/errs"
	"App-CloudBase-mcdull-mall/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetCouponList 店铺-优惠券列表
func (h *Handler) GetCouponList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])

	// 1.查询优惠券
	cList, total, err := service.GetOnlineCouponList(h.ctx, page, size)
	if err != nil {
		log.Printf("call GetOnlineCouponList failed, err:%v", err)
		panic(errs.ErrorInternalFaults)
	}
	// 2.渲染
	voList := make([]*defs.PortalCouponVO, 0)
	for _, v := range cList {
		couponVO := &defs.PortalCouponVO{}
		couponVO.Id = v.Id
		couponVO.Title = v.Title
		couponVO.FullMoney = v.FullMoney
		couponVO.Minus = v.Minus
		couponVO.Rate = v.Rate
		couponVO.Type = v.Type
		couponVO.StartTime = v.StartTime.Format("2006-01-02 15:04:05")
		couponVO.EndTime = v.EndTime.Format("2006-01-02 15:04:05")
		couponVO.Description = v.Description
		voList = append(voList, couponVO)
	}
	resp := make(map[string]interface{})
	resp["list"] = voList
	resp["total"] = total
	defs.SendNormalResponse(w, resp)
}

// TakeCoupon 领取优惠券
func (h *Handler) TakeCoupon(w http.ResponseWriter, r *http.Request) {
	req := defs.PortalTakeCouponReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	couponId := req.CouponId
	userId := r.Context().Value(defs.ContextKey).(int)
	// 1.检查优惠券
	coupon, err := service.GetOnlineCouponById(h.ctx, couponId)
	if err != nil {
		log.Printf("call GetOnlineCouponById failed, err:%v", err)
		if err == dao.NotFoundRecord {
			panic(errs.NewErrorCoupon("优惠券不存在或下架了"))
		}
		panic(errs.ErrorInternalFaults)
	}
	// 2.检查领取的数量
	totalTakeNum, err := service.GetTakeCouponNum(h.ctx, couponId)
	if err != nil {
		log.Printf("call GetOnlineCouponById failed, err:%v", err)
		panic(errs.ErrorParameterValidate)
	}
	if totalTakeNum >= coupon.GrantNum {
		panic(errs.NewErrorCoupon("来晚了，优惠券领光了！"))
	}
	// 3.单人限领
	userTakeNum, err := service.GetUserTakeCouponNum(h.ctx, couponId, userId)
	if err != nil {
		log.Printf("call GetUserTakeCouponNum failed, err:%v", err)
		panic(errs.ErrorInternalFaults)
	}
	if userTakeNum >= coupon.LimitNum {
		panic(errs.NewErrorCoupon("单人限领！"))
	}
	// 4.记录领取
	err = service.RecordCouponLog(h.ctx, userId, coupon)
	if err != nil {
		panic(errs.ErrorInternalFaults)
	}
	defs.SendNormalResponse(w, "ok")
}

// GetUserCouponList 用户领取的优惠券
// status: 0-未使用 1-已使用 2-已过期
func (h *Handler) GetUserCouponList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	status, _ := strconv.Atoi(vars["status"])
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])
	userId := r.Context().Value(defs.ContextKey).(int)

	// 1.查询记录
	cList, total, err := service.GetUserCouponLog(h.ctx, userId, status, page, size)
	if err != nil {
		log.Printf("call GetUserCouponLog failed, err:%v", err)
		panic(errs.ErrorInternalFaults)
	}
	// 2.渲染
	voList := make([]defs.PortalUserCouponVO, 0)
	for _, v := range cList {
		couponDO, err := service.GetCouponById(h.ctx, v.CouponId)
		if err != nil {
			if err == dao.NotFoundRecord {
				panic(errs.ErrorCoupon)
			}
			log.Printf("call GetCouponById failed, err:%v", err)
			panic(errs.ErrorInternalFaults)
		}
		couponVO := defs.PortalUserCouponVO{}
		couponVO.CLogId = v.Id
		couponVO.CouponId = v.CouponId
		couponVO.Title = couponDO.Title
		couponVO.FullMoney = couponDO.FullMoney
		couponVO.Minus = couponDO.Minus
		couponVO.Rate = couponDO.Rate
		couponVO.Type = couponDO.Type
		couponVO.StartTime = couponDO.StartTime.Format("2006-01-02 15:04:05")
		couponVO.EndTime = v.ExpireTime.Format("2006-01-02 15:04:05")
		couponVO.Description = couponDO.Description
		voList = append(voList, couponVO)
	}
	resp := make(map[string]interface{})
	resp["list"] = voList
	resp["total"] = total
	defs.SendNormalResponse(w, resp)
}

// DoDeleteCouponLog 删除领取的优惠券
func (h *Handler) DoDeleteCouponLog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	couponLogId, _ := strconv.Atoi(vars["id"])
	userId := r.Context().Value(defs.ContextKey).(int)

	err := service.DeleteCouponLog(h.ctx, userId, couponLogId, defs.DELETE)
	if err != nil {
		log.Printf("call DeleteCouponLog failed, err:%v", err)
		panic(errs.ErrorInternalFaults)
	}
	defs.SendNormalResponse(w, "ok")
}
