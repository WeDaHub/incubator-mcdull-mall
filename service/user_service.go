package service

import (
	"App-CloudBase-mcdull-mall/dao"
	"App-CloudBase-mcdull-mall/defs"
	"App-CloudBase-mcdull-mall/env"
	"App-CloudBase-mcdull-mall/model"
	"App-CloudBase-mcdull-mall/utils"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

var WechatServiceError = errors.New("wechat service error")
var NotFoundUserError = errors.New("not found user error")
var InvalidTokenError = errors.New("invalid token error")

// LoginCodeAuth 微信Code换用户信息
func LoginCodeAuth(ctx context.Context, code string) (string, int, error) {
	conf := env.LoadConf()
	appId := conf.Mini.AppId
	appSecret := conf.Mini.Secret
	baseUrl := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	url := fmt.Sprintf(baseUrl, appId, appSecret, code)
	rspBody, err := utils.DefaultClient.DoReq("GET", url, nil, nil)
	if err != nil {
		log.Printf("http get failed, err:%v", err)
		return "", 0, err
	}
	ret := make(map[string]interface{})
	err = json.Unmarshal(rspBody, &ret)
	if err != nil {
		return "", 0, WechatServiceError
	}
	// {"session_key":"TppZM2zEd6\/dGzkqbbrriQ==","expires_in":7200,"openid":"oQOru0EUuLdidBZH0r_F8fDURPjI"}
	if ret["errcode"] != nil {
		return "", 0, WechatServiceError
	}
	// 注册用户
	userId, err := registerUser(ctx, ret["openid"].(string))
	if err != nil {
		log.Printf("call registerUser failed, err:%v", err)
		return "", 0, err
	}
	token, err := utils.CreateToken(userId, defs.AccessTokenExpire)
	if err != nil {
		log.Printf("call CreateToken failed, err:%v", err)
		return "", 0, err
	}
	// 缓存Token
	cacheKey := defs.MiniappTokenPrefix + token
	dao.SetLoadCache(cacheKey, ret, defs.AccessTokenExpire)
	return token, userId, nil
}

func registerUser(ctx context.Context, openid string) (int, error) {
	user, err := dao.GetUserByOpenid(ctx, openid)
	if err != nil {
		if err != dao.NotFoundRecord {
			log.Printf("call GetUserByOpenid failed, err:%v", err)
			return 0, err
		}
	}
	uid := 0
	if user != nil {
		uid = user.Id
	}
	if user == nil {
		// 新增记录
		uid, err = dao.AddUser(ctx, openid)
		if err != nil {
			log.Printf("call AddUser failed, err:%v", err)
			return 0, err
		}
	}
	return uid, nil
}

// GetUserInfoById 查询用户信息
func GetUserInfoById(ctx context.Context, userId int) (*model.UserDO, error) {
	user, err := dao.GetUserById(ctx, userId)
	if err != nil {
		if err == dao.NotFoundRecord {
			return nil, NotFoundUserError
		}
		log.Printf("call GetUserById failed, err:%v", err)
		return nil, err
	}
	return user, nil
}

// GetSeesionKeyFromCache 从缓存中提取sessionKey
func GetSeesionKeyFromCache(token string) (string, error) {
	cacheKey := defs.MiniappTokenPrefix + token
	if val, ok := dao.GetLocalCache(cacheKey); ok {
		ret := val.(map[string]interface{})
		return ret["session_key"].(string), nil
	}
	return "", InvalidTokenError
}

// DoWxUserPhoneSignature 解密手机号
func DoWxUserPhoneSignature(ctx context.Context, userId int, sessionKey, encryptedData, iv string) error {
	// 1.解密数据
	conf := env.LoadConf()
	appId := conf.Mini.AppId
	wxSensitiveData := utils.WxSensitiveData{AppId: appId, SessionKey: sessionKey, Iv: iv, EncryptedData: encryptedData}
	decrypt, err := wxSensitiveData.Decrypt()
	if err != nil {
		panic(err)
	}
	// 2.查询用户信息
	userDO, err := dao.GetUserById(ctx, userId)
	if err != nil {
		if err == dao.NotFoundRecord {
			return NotFoundUserError
		}
		return err
	}
	// 3.更新用户信息
	userDO.Mobile = decrypt["phoneNumber"].(string)
	err = dao.UpdateUserInfo(ctx, userDO)
	if err != nil {
		log.Printf("call UpdateUserInfo failed, err:%v", err)
		return err
	}
	return nil
}

// DoUserAuthInfo 授权用户信息
func DoUserAuthInfo(ctx context.Context, userId int, req defs.WxappAuthUserInfoReq) error {
	userDO, err := dao.GetUserById(ctx, userId)
	if err != nil {
		if err != dao.NotFoundRecord {
			return NotFoundUserError
		}
		log.Printf("call GetUserById failed, err:%v", err)
		return err
	}
	userDO.Nickname = req.NickName
	userDO.Avatar = req.AvatarUrl
	userDO.Gender = req.Gender
	userDO.Country = req.Country
	userDO.Province = req.Province
	userDO.City = req.City
	err = dao.UpdateUserInfo(ctx, userDO)
	if err != nil {
		log.Printf("call UpdateUserInfo failed, err:%v", err)
		return err
	}
	return nil
}
