package handler

import (
	"App-CloudBase-mcdull-mall/defs"
	"App-CloudBase-mcdull-mall/env"
	"App-CloudBase-mcdull-mall/errs"
	"App-CloudBase-mcdull-mall/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// UserBrowseHistory 用户-历史浏览
func (h *Handler) UserBrowseHistory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])
	userId := r.Context().Value(defs.ContextKey).(int)

	// 1.查询浏览记录
	recordList, total, err := service.ListBrowseRecord(h.ctx, userId, page, size)
	if err != nil {
		log.Printf("call ListBrowseRecord failed, err:%v", err)
		panic(errs.ErrorInternalFaults)
	}
	// 2.渲染
	domain := env.LoadConf().Domain
	recordVOs := make([]*defs.PortalBrowseRecordVO, 0)
	for _, recordDO := range recordList {
		recordVO := &defs.PortalBrowseRecordVO{}
		recordVO.Id = recordDO.Id
		recordVO.GoodsId = recordDO.GoodsId
		recordVO.Picture = domain + recordDO.Picture
		recordVO.Title = recordDO.Title
		recordVO.Price = recordDO.Price
		recordVO.BrowseTime = recordDO.UpdateTime.Format("2006-01-02 15:04:05")
		recordVOs = append(recordVOs, recordVO)
	}
	resp := map[string]interface{}{}
	resp["list"] = recordVOs
	resp["total"] = total
	defs.SendNormalResponse(w, resp)
}

// ClearBrowseHistory 清理-浏览历史
func (h *Handler) ClearBrowseHistory(w http.ResponseWriter, r *http.Request) {
	var ids []int
	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	err = service.ClearBrowseHistory(h.ctx, ids)
	if err != nil {
		log.Printf("call ClearBrowseHistory failed, err:%v", err)
		panic(errs.ErrorInternalFaults)
	}
	defs.SendNormalResponse(w, "ok")
}
