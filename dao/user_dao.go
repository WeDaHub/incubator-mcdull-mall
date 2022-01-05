package dao

import (
	"App-CloudBase-mcdull-mall/model"
	"context"

	"gorm.io/gorm"
)

// AddUser 新增用户
func AddUser(ctx context.Context, openid string) (int, error) {
	return model.CreateUser(ctx, openid)
}

// GetUserByOpenid 通过openid查询
func GetUserByOpenid(ctx context.Context, openid string) (*model.UserDO, error) {
	c := model.UserDO{Openid: openid}
	out, err := c.FindUser(ctx)
	if err == gorm.ErrRecordNotFound {
		return nil, NotFoundRecord
	}
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetUserById 通过Id查询用户
func GetUserById(ctx context.Context, userId int) (*model.UserDO, error) {
	c := model.UserDO{Id: userId}
	out, err := c.FindUser(ctx)
	if err == gorm.ErrRecordNotFound {
		return nil, NotFoundRecord
	}
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateUserInfo 更新用户信息
func UpdateUserInfo(ctx context.Context, do *model.UserDO) error {
	return do.UpdateUser(ctx)
}
