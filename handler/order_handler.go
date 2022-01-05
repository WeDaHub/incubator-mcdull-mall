package handler

import (
	"App-CloudBase-mcdull-mall/errs"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/shopspring/decimal"

	"App-CloudBase-mcdull-mall/defs"
	"App-CloudBase-mcdull-mall/service"

	"github.com/gorilla/mux"
)

// PlaceOrder 订单-C端下订单
func (h *Handler) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	req := &defs.PortalCartPlaceOrderReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	userId := r.Context().Value(defs.ContextKey).(int)
	// 1.参数类型状态
	dispatchAmount, err := decimal.NewFromString(req.DispatchAmount)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	expectAmount, err := decimal.NewFromString(req.ExpectAmount)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	// 2.下订单
	orderInfo, err := service.GenerateOrder(h.ctx, userId, req.AddressId, req.CouponLogId, dispatchAmount, expectAmount, req.GoodsList)
	if err != nil {
		log.Printf("call GenerateOrder failed, err:%v", err)
		panic(errs.ErrorInternalFaults)
	}
	defs.SendNormalResponse(w, orderInfo)
}

// GetOrderList 订单-C端列表
func (h *Handler) GetOrderList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	status, _ := strconv.Atoi(vars["status"])
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])
	userId := r.Context().Value(defs.ContextKey).(int)
	// 1.查询订单列表
	orderList, total, err := service.QueryOrderList(h.ctx, userId, status, page, size)
	if err != nil {
		log.Printf("call QueryOrderList failed, err:%v", err)
		panic(errs.ErrorInternalFaults)
	}
	// 2.渲染
	voList := make([]*defs.PortalOrderListVO, 0)
	for _, v := range orderList {
		// 3.查询订单商品
		goodsList, num, err := service.ExtraceOrderGoods(h.ctx, v.OrderNo)
		if err != nil {
			log.Printf("call ExtraceOrderGoods failed, err:%v", err)
			panic(errs.ErrorInternalFaults)
		}
		orderVO := &defs.PortalOrderListVO{}
		orderVO.Id = v.Id
		orderVO.OrderNo = v.OrderNo
		orderVO.PlaceTime = v.CreateTime.Format("2006-01-02 15:04:05")
		orderVO.PayAmount, _ = strconv.ParseFloat(v.PayAmount, 2)
		orderVO.Status = v.Status
		orderVO.GoodsList = goodsList
		orderVO.GoodsNum = num
		voList = append(voList, orderVO)
	}
	resp := make(map[string]interface{})
	resp["list"] = voList
	resp["total"] = total
	defs.SendNormalResponse(w, resp)
}

// GetOrderDetail 查询订单详情
func (h *Handler) GetOrderDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderNo := vars["orderNo"]
	log.Printf("reqData{orderNo:%s}", orderNo)
	userId := r.Context().Value(defs.ContextKey).(int)
	// 查询订单
	orderVO, err := service.QueryOrderDetail(h.ctx, userId, orderNo)
	if err != nil {
		if err == service.NotFoundOrderError {
			panic(errs.ErrorOrder)
		}
		log.Printf("call QueryOrderDetail failed, err:%v", err)
		panic(errs.ErrorInternalFaults)
	}
	defs.SendNormalResponse(w, orderVO)
}

// CancelOrder 订单-C端取消订单
func (h *Handler) CancelOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := r.Context().Value(defs.ContextKey).(int)

	orderId, _ := strconv.Atoi(vars["id"])
	err := service.CancelOrder(h.ctx, userId, orderId)
	if err != nil {
		if err == service.NotFoundOrderError {
			panic(errs.ErrorOrder)
		}
		if err == service.CancelOrderError {
			panic(errs.NewErrorOrder("进行中的订单，无法取消"))
		}
		log.Printf("call CancelOrder failed, err:%v", err)
		panic(errs.ErrorInternalFaults)
	}
	defs.SendNormalResponse(w, "ok")
}

// DeleteOrder 订单-C端删除订单
func (h *Handler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := r.Context().Value(defs.ContextKey).(int)

	orderId, _ := strconv.Atoi(vars["id"])
	err := service.DeleteOrderRecord(h.ctx, userId, orderId)
	if err != nil {
		if err == service.NotFoundOrderError {
			panic(errs.ErrorOrder)
		}
		if err == service.CancelOrderError {
			panic(errs.NewErrorOrder("进行中的订单，无法删除"))
		}
		log.Printf("call DeleteOrderRecord failed, err:%v", err)
		panic(errs.ErrorInternalFaults)
	}
	defs.SendNormalResponse(w, "ok")
}

// ConfirmTakeGoods 订单-C端确认收货
func (h *Handler) ConfirmTakeGoods(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := r.Context().Value(defs.ContextKey).(int)

	orderId, _ := strconv.Atoi(vars["id"])
	err := service.ConfirmTakeGoods(h.ctx, userId, orderId)
	if err != nil {
		if err == service.NotFoundOrderError {
			panic(errs.ErrorOrder)
		}
		log.Printf("call ConfirmTakeGoods failed, err:%v", err)
		panic(errs.ErrorInternalFaults)
	}
	defs.SendNormalResponse(w, "ok")
}

// GetOrderRemind 订单-红点提醒
// 订单状态：-1 已取消 0-待付款 1-待发货 2-待收货 3-已完成 4-（待发货）退款申请 5-已退款
func (h *Handler) GetOrderRemind(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(defs.ContextKey).(int)

	remindVO := defs.OrderRemindVO{}
	remindVO.WaitPay = service.CountOrderNum(h.ctx, userId, 0)
	remindVO.NotExpress = service.CountOrderNum(h.ctx, userId, 1)
	remindVO.WaitReceive = service.CountOrderNum(h.ctx, userId, 2)
	defs.SendNormalResponse(w, remindVO)
}

// WxPayNotify 订单-微信支付回调
func (h *Handler) WxPayNotify(w http.ResponseWriter, r *http.Request) {
	// todo: 解析数据，从 attach 字段获取订单号，响应微信服务器
	_ = service.OrderPaySuccessNotify(h.ctx, "")
}

// RefundApply 退款-C端退款申请
func (h *Handler) RefundApply(w http.ResponseWriter, r *http.Request) {
	req := defs.OrderRefundApplyReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(err)
	}
	userId := r.Context().Value(defs.ContextKey).(int)
	// 订单退款
	refundNo, err := service.RefundApply(h.ctx, userId, req.OrderNo, req.Reason)
	if err != nil {
		if err == service.NotFoundOrderError {
			panic(errs.ErrorOrder)
		}
		if err == service.OrderStatusError {
			panic(errs.NewErrorOrder("非法操作"))
		}
		panic(errs.ErrorInternalFaults)
	}
	resp := defs.OrderRefundApplyVO{RefundNo: refundNo}
	defs.SendNormalResponse(w, resp)
}

// RefundDetail 退款-C端退款详情
func (h *Handler) RefundDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	refundNo := vars["refundNo"]
	userId := r.Context().Value(defs.ContextKey).(int)
	// 查询退款详情
	refundDetail, err := service.QueryRefundDetail(h.ctx, userId, refundNo)
	if err != nil {
		if err == service.NotFoundRefundError {
			panic(errs.ErrorOrderRefund)
		}
		log.Printf("call QueryRefundDetail failed, err:%v", err)
		panic(errs.ErrorInternalFaults)
	}
	defs.SendNormalResponse(w, refundDetail)
}

// UndoRefundApply 退款-C端撤销申请
func (h *Handler) UndoRefundApply(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	refundNo := vars["refundNo"]
	userId := r.Context().Value(defs.ContextKey).(int)

	err := service.UndoRefundApply(h.ctx, userId, refundNo)
	if err != nil {
		if err == service.NotFoundRefundError {
			panic(errs.ErrorOrderRefund)
		}
		if err == service.OrderStatusError {
			panic(errs.NewErrorOrderRefund("状态异常"))
		}
		log.Printf("call UndoRefundApply failed, err:%v", err)
		panic(errs.ErrorInternalFaults)
	}
	defs.SendNormalResponse(w, "ok")
}
