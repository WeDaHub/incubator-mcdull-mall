package dao

import (
	"App-CloudBase-mcdull-mall/env"
	"App-CloudBase-mcdull-mall/utils"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/url"
)

const (
	wxToken      = "https://api.weixin.qq.com/cgi-bin/token"
	TokenInvalid = 40001 // 无效的token
	TokenExpired = 42001 // token过期错误码
)

type TokenRet struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

var InvalidAccessToken = errors.New("invalid access_token")

// Token本地缓存
var tokenCacheKey = "access_token_cache_key"

// RefreshAccessToken 刷新Token
func RefreshAccessToken(ctx context.Context) {
	_, err := GetAccessTokenCache(ctx, true)
	if err != nil {
		log.Printf("refreshAccessToken failed, err:%v", err)
	}
}

func GetAccessTokenCache(ctx context.Context, refresh bool) (string, error) {
	if !refresh {
		// 从缓存中获取access_token
		if val, ok := GetLocalCache(tokenCacheKey); ok {
			accessToken := val.(string)
			return accessToken, nil
		}
	}
	// get access_token by http
	ret, err := getAccessToken(ctx)
	if err != nil {
		log.Printf("getAccessToken failed, err:%v", err)
		return "", err
	}
	// write cache
	SetLoadCache(tokenCacheKey, ret.AccessToken, uint32(ret.ExpiresIn-200))
	return ret.AccessToken, nil
}

func getAccessToken(ctx context.Context) (*TokenRet, error) {
	conf := env.LoadConf()
	appId := conf.Mini.AppId
	appSecret := conf.Mini.Secret
	params := url.Values{}
	params.Set("grant_type", "client_credential")
	params.Set("appid", appId)
	params.Set("secret", appSecret)
	reqParams := params.Encode()
	tokenUrl := wxToken + "?" + reqParams
	rspBody, err := utils.DefaultClient.DoReq("GET", tokenUrl, nil, nil)
	if err != nil {
		log.Printf("http DoReq failed, err:%v", err)
		return nil, err
	}
	ret := new(TokenRet)
	err = json.Unmarshal(rspBody, ret)
	if err != nil {
		return nil, InvalidAccessToken
	}
	return ret, nil
}
