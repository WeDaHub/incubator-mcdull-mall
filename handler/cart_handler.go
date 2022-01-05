package handler

import (
	"App-CloudBase-mcdull-mall/defs"
	"App-CloudBase-mcdull-mall/errs"
	"App-CloudBase-mcdull-mall/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetCartGoodsList 查询-购物车商品
func (h *Handler) GetCartGoodsList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])
	userId := r.Context().Value(defs.ContextKey).(int)
	cartGoodsList, total, err := service.GetCartGoods(h.ctx, userId, page, size)
	if err != nil {
		log.Printf("call GetCartGoods failed, err:%v", err)
		panic(errs.ErrorInternalFaults)
	}
	resp := make(map[string]interface{})
	resp["list"] = cartGoodsList
	resp["total"] = total
	defs.SendNormalResponse(w, resp)
}

// AddCartGoods 添加-购物车商品
func (h *Handler) AddCartGoods(w http.ResponseWriter, r *http.Request) {
	req := defs.PortalCartGoodsReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	userId := r.Context().Value(defs.ContextKey).(int)
	err = service.DoEditCart(h.ctx, userId, req.GoodsId, req.SkuId, req.Num)
	if err != nil {
		log.Printf("call DoEditCart failed, err:%v", err)
		panic(err)
	}
	defs.SendNormalResponse(w, "ok")
}

// EditCartGoods 编辑-购物车
func (h *Handler) EditCartGoods(w http.ResponseWriter, r *http.Request) {
	req := defs.PortalEditCartReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	userId := r.Context().Value(defs.ContextKey).(int)
	err = service.EditCartGoods(h.ctx, userId, req.Id, req.Num)
	if err != nil {
		log.Printf("call EditCartGoods failed, err:%v", err)
		panic(err)
	}
	defs.SendNormalResponse(w, "ok")
}

// GetCartGoodsNum 查询-购物车商品数量
func (h *Handler) GetCartGoodsNum(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(defs.ContextKey).(int)
	num, err := service.CountCartGoodsNum(h.ctx, userId)
	if err != nil {
		log.Printf("call CountCartGoodsNum failed, err:%v", err)
		panic(err)
	}
	resp := map[string]interface{}{}
	resp["num"] = num
	defs.SendNormalResponse(w, resp)
}
