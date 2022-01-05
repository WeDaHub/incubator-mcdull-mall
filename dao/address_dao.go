package dao

import (
	"App-CloudBase-mcdull-mall/model"
	"context"

	"gorm.io/gorm"
)

// GetUserAddressList 查询用户收货地址
func GetUserAddressList(ctx context.Context, userId int) ([]*model.UserAddressDO, error) {
	return model.GetUserAddressList(ctx, userId)
}

// GetAddressById 查询单个地址
func GetAddressById(ctx context.Context, id, userId int) (*model.UserAddressDO, error) {
	addr, err := model.GetAddressById(ctx, id, userId)
	if err == gorm.ErrRecordNotFound {
		return nil, NotFoundRecord
	}
	if err != nil {
		return nil, err
	}
	return addr, nil
}

// GetDefaultAddress 查询默认收货地址
func GetDefaultAddress(ctx context.Context, userId int) (*model.UserAddressDO, error) {
	addr, err := model.GetDefaultAddress(ctx, userId)
	if err == gorm.ErrRecordNotFound {
		return nil, NotFoundRecord
	}
	if err != nil {
		return nil, err
	}
	return addr, nil
}

// AddUserAddress 新增收货地址
func AddUserAddress(ctx context.Context, do *model.UserAddressDO) error {
	return model.AddUserAddress(ctx, do)
}

// UpdateUserAddress 更新收货地址
func UpdateUserAddress(ctx context.Context, do *model.UserAddressDO) error {
	return model.UpdateUserAddress(ctx, do)
}

// ClearDefaultUserAddress 清理用户默认收货地址
func ClearDefaultUserAddress(ctx context.Context, userId int) error {
	return model.ClearDefaultUserAddress(ctx, userId)
}
