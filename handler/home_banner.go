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

// HomeBanner 查询首页Banner
// 示例：http://127.0.0.1:80/api/home/banner?page=1&size=10
func (h *Handler) HomeBanner(w http.ResponseWriter, r *http.Request) {
	// 1.请求参数
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])
	log.Printf("reqData{page:%d,size:%d}", page, size)
	// 2.查询banner
	bList, err := service.GetBannerList(h.ctx, page, size)
	if err != nil {
		log.Printf("call service.GetBannerList failed, err:%v", err)
		panic(errs.ErrorInternalFaults)
	}
	// 3.渲染数据
	domain := env.LoadConf().Domain
	voList := make([]defs.PortalBannerVO, 0)
	for _, v := range bList {
		voList = append(voList, defs.PortalBannerVO{
			Id:           v.Id,
			Picture:      domain + v.Picture,
			BusinessType: v.BusinessType,
			BusinessId:   v.BusinessId,
		})
	}
	// 4.响应请求
	defs.SendNormalResponse(w, voList)
}

// GetGridCategoryList 查询宫格列表
func (h *Handler) GetGridCategoryList(w http.ResponseWriter, r *http.Request) {
	// 1.请求参数
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])
	log.Printf("reqData{page:%d,size:%d}", page, size)
	// 2.查询宫格列表
	gridList, err := service.GetGridCategoryList(h.ctx, page, size)
	if err != nil {
		log.Printf("call service.GetGridCategoryList failed, err:%v", err)
		panic(errs.ErrorInternalFaults)
	}
	// 3.渲染数据
	domain := env.LoadConf().Domain
	gridVOList := make([]defs.PortalGridCategoryVO, 0)
	for _, v := range gridList {
		gridVO := defs.PortalGridCategoryVO{}
		gridVO.Id = v.Id
		gridVO.Name = v.Name
		gridVO.Picture = domain + v.Picture
		gridVO.CategoryId = v.CategoryId
		gridVOList = append(gridVOList, gridVO)
	}
	// 4.响应请求
	defs.SendNormalResponse(w, gridVOList)
}

// GetSubCategoryList 查询-全部二级分类
func (h *Handler) GetSubCategoryList(w http.ResponseWriter, r *http.Request) {
	// 1.查询分类
	cList, err := service.GetSubCategoryList(h.ctx)
	if err != nil {
		log.Printf("call service.GetSubCategoryList failed, err:%v", err)
		panic(errs.ErrorInternalFaults)
	}
	// 2.渲染数据
	voList := make([]*defs.PortalCategoryVO, 0)
	for _, v := range cList {
		categoryVO := &defs.PortalCategoryVO{}
		categoryVO.Id = v.Id
		categoryVO.Name = v.Name
		voList = append(voList, categoryVO)
	}
	defs.SendNormalResponse(w, voList)
}
