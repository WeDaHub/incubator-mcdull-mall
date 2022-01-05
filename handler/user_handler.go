package handler

import (
	"App-CloudBase-mcdull-mall/defs"
	"App-CloudBase-mcdull-mall/errs"
	"App-CloudBase-mcdull-mall/service"
	"App-CloudBase-mcdull-mall/utils"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Login 小程序-用户登录
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]
	if code == "" {
		panic(errs.NewParameterError("缺少code"))
	}
	// 注册用户，生成Token
	token, userId, err := service.LoginCodeAuth(h.ctx, code)
	if err != nil {
		log.Printf("call LoginCodeAuth failed, err:%v", err)
		if err == service.WechatServiceError {
			panic(errs.ErrorWechatError)
		}
		panic(errs.ErrorInternalFaults)
	}
	// 访客记录
	go service.RecordVisitorRecord(h.ctx, userId, r)
	defs.SendNormalResponse(w, defs.WxappLoginVO{Token: token})
}

// UserInfo 小程序-查询用户信息
func (h *Handler) UserInfo(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(defs.ContextKey).(int)
	// 1.查询用户
	userDO, err := service.GetUserInfoById(h.ctx, userId)
	if err != nil {
		if err == service.NotFoundUserError {
			panic(errs.ErrorMiniappUser)
		}
		log.Printf("call GetUserInfoById failed, err:%v", err)
		panic(errs.ErrorInternalFaults)
	}
	// 2.渲染
	userVO := defs.WxappUserInfoVO{}
	userVO.Uid = userDO.Id
	userVO.Nickname = userDO.Nickname
	userVO.Avatar = userDO.Avatar
	if userDO.Mobile != "" {
		userVO.Mobile = utils.PhoneMark(userDO.Mobile)
	}
	defs.SendNormalResponse(w, userVO)
}

// AuthPhone 小程序-授权手机号
func (h *Handler) AuthPhone(w http.ResponseWriter, r *http.Request) {
	req := defs.WxappAuthPhoneReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	// 1.提取Token
	authorization := r.Header.Get("Authorization")
	accessToken := strings.Split(authorization, " ")[1]
	// 2.提取缓存sessionKey
	sessionKey, err := service.GetSeesionKeyFromCache(accessToken)
	if err != nil {
		log.Printf("call GetSeesionKeyFromCache failed, err:%v", err)
		panic(errs.ErrorTokenInvalid)
	}
	// 3.解密数据
	userId := r.Context().Value(defs.ContextKey).(int)
	err = service.DoWxUserPhoneSignature(h.ctx, userId, sessionKey, req.EncryptedData, req.Iv)
	if err != nil {
		log.Printf("call DoWxUserPhoneSignature failed, err:%v", err)
		panic(errs.ErrorInternalFaults)
	}
	defs.SendNormalResponse(w, "ok")
}

// AuthUserInfo 小程序-授权用户信息
func (h *Handler) AuthUserInfo(w http.ResponseWriter, r *http.Request) {
	req := defs.WxappAuthUserInfoReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	userId := r.Context().Value(defs.ContextKey).(int)
	err = service.DoUserAuthInfo(h.ctx, userId, req)
	if err != nil {
		if err == service.NotFoundUserError {
			panic(errs.ErrorMiniappUser)
		}
		log.Printf("call DoUserAuthInfo failed, err:%v", err)
		panic(errs.ErrorInternalFaults)
	}
	defs.SendNormalResponse(w, "ok")
}
