package handler

import (
	"App-CloudBase-mcdull-mall/defs"
	"App-CloudBase-mcdull-mall/env"
	"App-CloudBase-mcdull-mall/errs"
	"App-CloudBase-mcdull-mall/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetGoodsList 查询商品列表
func (h *Handler) GetGoodsList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keyword := vars["k"]
	sort, _ := strconv.Atoi(vars["s"])
	categoryId, _ := strconv.Atoi(vars["c"])
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])
	log.Printf("reqData{k:%s,s:%d,c:%d,page:%d,size:%d}\n", keyword, sort, categoryId, page, size)
	// 1.查询商品
	goodsList, total, err := service.QueryGoodsList(h.ctx, keyword, sort, categoryId, page, size)
	if err != nil {
		log.Printf("call QueryGoodsList failed, err:%v", err)
		panic(errs.ErrorInternalFaults)
	}
	goodsIds := make([]int, 0)
	for _, v := range goodsList {
		goodsIds = append(goodsIds, v.Id)
	}
	// 2.统计购买人数
	numMap, err := service.BatchCountBuyUserNum(h.ctx, goodsIds)
	if err != nil {
		log.Printf("call BatchCountBuyUserNum failed, err:%v", err)
		panic(errs.ErrorInternalFaults)
	}
	// 3.渲染数据
	domain := env.LoadConf().Domain
	goodsVOList := make([]*defs.PortalGoodsListVO, 0)
	for _, v := range goodsList {
		goodsVO := &defs.PortalGoodsListVO{}
		goodsVO.Id = v.Id
		goodsVO.Title = v.Title
		goodsVO.Price, _ = strconv.ParseFloat(v.Price, 2)
		goodsVO.Picture = domain + v.Picture
		goodsVO.HumanNum = 0
		if num, ok := numMap.Load(v.Id); ok {
			goodsVO.HumanNum = num.(int)
		}
		goodsVOList = append(goodsVOList, goodsVO)
	}
	resp := make(map[string]interface{})
	resp["list"] = goodsVOList
	resp["total"] = total
	defs.SendNormalResponse(w, resp)
}

// GetGoodsDetail 查询商品详情
func (h *Handler) GetGoodsDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	goodsId, _ := strconv.Atoi(vars["id"])
	// 1.聚合商品详情
	goodsInfo, err := service.QueryPortalGoodsDetail(h.ctx, goodsId)
	if err != nil {
		log.Printf("call QueryPortalGoodsDetail failed, err:%v", err)
		panic(errs.ErrorInternalFaults)
	}
	// 2.记录浏览记录（异步）
	userId := r.Context().Value(defs.ContextKey).(int)
	go service.RecordGoodsBrowse(h.ctx, userId, goodsId)
	defs.SendNormalResponse(w, goodsInfo)
}
