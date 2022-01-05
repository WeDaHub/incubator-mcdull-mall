package handler

import (
	"App-CloudBase-mcdull-mall/defs"
	"App-CloudBase-mcdull-mall/errs"
	"App-CloudBase-mcdull-mall/model"
	"App-CloudBase-mcdull-mall/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// GetAddressList 查询-收货地址列表
func (h *Handler) GetAddressList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])
	userId := r.Context().Value(defs.ContextKey).(int)
	// 1.查询收货地址
	addressList, total, err := service.GetUserAddressList(h.ctx, userId, page, size)
	if err != nil {
		log.Printf("call GetUserAddressList failed, err:%v", err)
		panic(errs.ErrorInternalFaults)
	}
	// 2.渲染数据
	addressVOList := make([]defs.PortalAddressVO, 0)
	for _, v := range addressList {
		addressVO := defs.PortalAddressVO{}
		addressVO.Id = v.Id
		addressVO.Contacts = v.Contacts
		addressVO.Mobile = v.Mobile
		addressVO.ProvinceId = v.ProvinceId
		addressVO.CityId = v.CityId
		addressVO.AreaId = v.AreaId
		addressVO.ProvinceStr = v.ProvinceStr
		addressVO.CityStr = v.CityStr
		addressVO.AreaStr = v.AreaStr
		addressVO.Address = v.Address
		addressVO.IsDefault = v.IsDefault
		addressVOList = append(addressVOList, addressVO)
	}
	resp := make(map[string]interface{})
	resp["list"] = addressVOList
	resp["total"] = total
	defs.SendNormalResponse(w, resp)
}

// GetAddress 查询-单个地址
func (h *Handler) GetAddress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	addressId, _ := strconv.Atoi(vars["id"])
	userId := r.Context().Value(defs.ContextKey).(int)
	// 1.查询地址
	addressDO, err := service.GetAddressById(h.ctx, addressId, userId)
	if err == service.NotFoundAddressError {
		panic(errs.ErrorAddress)
	}
	if err != nil {
		log.Printf("call GetAddressById failed, err:%v", err)
		panic(errs.ErrorInternalFaults)
	}
	// 2.渲染
	addressVO := defs.PortalAddressVO{}
	addressVO.Id = addressDO.Id
	addressVO.Contacts = addressDO.Contacts
	addressVO.Mobile = addressDO.Mobile
	addressVO.ProvinceId = addressDO.ProvinceId
	addressVO.CityId = addressDO.CityId
	addressVO.AreaId = addressDO.AreaId
	addressVO.ProvinceStr = addressDO.ProvinceStr
	addressVO.CityStr = addressDO.CityStr
	addressVO.AreaStr = addressDO.AreaStr
	addressVO.Address = addressDO.Address
	addressVO.IsDefault = addressDO.IsDefault
	defs.SendNormalResponse(w, addressVO)
}

// EditAddress 新增/更新-收货地址
func (h *Handler) EditAddress(w http.ResponseWriter, r *http.Request) {
	req := defs.PortalAddressReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	userId := r.Context().Value(defs.ContextKey).(int)
	if req.Id == defs.ZERO {
		defaultAddr, err := service.GetDefaultAddress(h.ctx, userId)
		if err != nil {
			if err != service.NotFoundAddressError {
				log.Printf("call GetDefaultAddress failed, err:%v", err)
				panic(errs.ErrorInternalFaults)
			}
		}
		isDefault := 0
		if defaultAddr == nil {
			isDefault = 1
		}
		address := &model.UserAddressDO{}
		address.UserId = userId
		address.Contacts = req.Contacts
		address.Mobile = req.Mobile
		address.ProvinceId = req.ProvinceId
		address.CityId = req.CityId
		address.AreaId = req.AreaId
		address.ProvinceStr = req.ProvinceStr
		address.CityStr = req.CityStr
		address.AreaStr = req.AreaStr
		address.Address = req.Address
		address.IsDefault = isDefault
		address.CreateTime = time.Now()
		address.UpdateTime, _ = time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")
		err = service.AddUserAddress(h.ctx, address)
		if err != nil {
			log.Printf("call AddUserAddress failed, err:%v", err)
			panic(errs.ErrorInternalFaults)
		}
	} else {
		address, err := service.GetAddressById(h.ctx, req.Id, userId)
		if err != nil {
			if err == service.NotFoundAddressError {
				panic(errs.ErrorAddress)
			}
			log.Printf("call GetAddressById failed, err:%v", err)
			panic(errs.ErrorInternalFaults)
		}
		address.UserId = userId
		address.Contacts = req.Contacts
		address.Mobile = req.Mobile
		address.ProvinceId = req.ProvinceId
		address.CityId = req.CityId
		address.AreaId = req.AreaId
		address.ProvinceStr = req.ProvinceStr
		address.CityStr = req.CityStr
		address.AreaStr = req.AreaStr
		address.Address = req.Address
		address.IsDefault = req.IsDefault
		err = service.UpdateUserAddress(h.ctx, address)
		if err != nil {
			log.Printf("call UpdateUserAddress failed, err:%v", err)
			panic(errs.ErrorInternalFaults)
		}
	}
	defs.SendNormalResponse(w, "ok")
}

// DoDeleteAddress 删除-收货地址
func (h *Handler) DoDeleteAddress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	addressId, _ := strconv.Atoi(vars["id"])
	userId := r.Context().Value(defs.ContextKey).(int)

	address, err := service.GetAddressById(h.ctx, addressId, userId)
	if err != nil {
		if err == service.NotFoundAddressError {
			panic(errs.ErrorAddress)
		}
		log.Printf("call GetAddressById failed, err:%v", err)
		panic(errs.ErrorInternalFaults)
	}
	address.Del = defs.DELETE
	err = service.UpdateUserAddress(h.ctx, address)
	if err != nil {
		log.Printf("call UpdateUserAddress failed, err:%v", err)
		panic(errs.ErrorInternalFaults)
	}
	defs.SendNormalResponse(w, "ok")
}

// GetDefaultAddress 查询-默认收货地址
func (h *Handler) GetDefaultAddress(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(defs.ContextKey).(int)
	// 1.查询默认收货地址
	addressDO, err := service.GetDefaultAddress(h.ctx, userId)
	if err != nil {
		if err == service.NotFoundAddressError {
			panic(errs.ErrorAddress)
		}
		log.Printf("call GetDefaultAddress failed, err:%v", err)
		panic(errs.ErrorInternalFaults)
	}
	// 2.渲染数据
	addressVO := defs.PortalAddressVO{}
	addressVO.Id = addressDO.Id
	addressVO.Contacts = addressDO.Contacts
	addressVO.Mobile = addressDO.Mobile
	addressVO.ProvinceId = addressDO.ProvinceId
	addressVO.CityId = addressDO.CityId
	addressVO.AreaId = addressDO.AreaId
	addressVO.ProvinceStr = addressDO.ProvinceStr
	addressVO.CityStr = addressDO.CityStr
	addressVO.AreaStr = addressDO.AreaStr
	addressVO.Address = addressDO.Address
	addressVO.IsDefault = addressDO.IsDefault
	defs.SendNormalResponse(w, addressVO)
}
