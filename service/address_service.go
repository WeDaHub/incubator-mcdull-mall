package service

import (
	"App-CloudBase-mcdull-mall/dao"
	"App-CloudBase-mcdull-mall/model"
	"context"
	"errors"
	"log"
)

var NotFoundAddressError = errors.New("not found address error")

// GetUserAddressList 查询用户收货地址
func GetUserAddressList(ctx context.Context, userId, page, size int) ([]*model.UserAddressDO, int, error) {
	aList, err := dao.GetUserAddressList(ctx, userId)
	if err != nil {
		log.Printf("call GetUserAddressList failed, err:%v", err)
		return nil, 0, err
	}
	// 分页
	total := len(aList)
	offset, endpos := pagePos(total, page, size)
	return aList[offset:endpos], total, nil
}

// GetAddressById 查询单个地址
func GetAddressById(ctx context.Context, id, userId int) (*model.UserAddressDO, error) {
	addrDO, err := dao.GetAddressById(ctx, id, userId)
	if err != nil {
		if err == dao.NotFoundRecord {
			return nil, NotFoundAddressError
		}
		log.Printf("call GetAddressById failed, err:%v", err)
		return nil, err
	}
	return addrDO, nil
}

// GetDefaultAddress 查询默认地址
func GetDefaultAddress(ctx context.Context, userId int) (*model.UserAddressDO, error) {
	addrDO, err := dao.GetDefaultAddress(ctx, userId)
	if err == dao.NotFoundRecord {
		return nil, NotFoundAddressError
	}
	if err != nil {
		log.Printf("call GetDefaultAddress failed, err:%v", err)
		return nil, err
	}
	return addrDO, nil
}

// AddUserAddress 添加收货地址
func AddUserAddress(ctx context.Context, do *model.UserAddressDO) error {
	return dao.AddUserAddress(ctx, do)
}

// UpdateUserAddress 更新收货地址
func UpdateUserAddress(ctx context.Context, do *model.UserAddressDO) error {
	if do.IsDefault == 1 {
		err := clearDefaultAddress(ctx, do.UserId)
		if err != nil {
			log.Printf("call clearDefaultAddress failed, err:%v", err)
			return err
		}
	}
	return dao.UpdateUserAddress(ctx, do)
}

// clearDefaultAddress 清理默认的收货地址
func clearDefaultAddress(ctx context.Context, userId int) error {
	return dao.ClearDefaultUserAddress(ctx, userId)
}
