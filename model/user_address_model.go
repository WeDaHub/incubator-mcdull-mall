package model

import (
	"App-CloudBase-mcdull-mall/pkg/database"
	"context"
	"time"
)

// UserAddressDO 商城-用户收货地址
type UserAddressDO struct {
	Id          int       `gorm:"column:id" json:"id"`
	UserId      int       `gorm:"column:user_id" json:"user_id"`           // 用户ID
	Contacts    string    `gorm:"column:contacts" json:"contacts"`         // 联系人
	Mobile      string    `gorm:"column:mobile" json:"mobile"`             // 手机号
	ProvinceId  string    `gorm:"column:province_id" json:"province_id"`   // 省份编码
	CityId      string    `gorm:"column:city_id" json:"city_id"`           // 城市编码
	AreaId      string    `gorm:"column:area_id" json:"area_id"`           // 地区编码
	ProvinceStr string    `gorm:"column:province_str" json:"province_str"` // 省份
	CityStr     string    `gorm:"column:city_str" json:"city_str"`         // 城市
	AreaStr     string    `gorm:"column:area_str" json:"area_str"`         // 地区
	Address     string    `gorm:"column:address" json:"address"`           // 详细地址
	IsDefault   int       `gorm:"column:is_default" json:"is_default"`     // 默认收货地址：0-否 1-是
	Del         int       `gorm:"column:is_del" json:"del"`
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime  time.Time `gorm:"column:update_time" json:"update_time"`
}

func (u *UserAddressDO) TableName() string {
	return "wechat_mall_user_address"
}

// GetUserAddressList 查询用户收货地址
func GetUserAddressList(ctx context.Context, userId int) ([]*UserAddressDO, error) {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return nil, err
	}
	empty := new(UserAddressDO)
	aList := make([]*UserAddressDO, 0)
	err = db.Table(empty.TableName()).Where("is_del = 0 AND user_id = ?", userId).Find(&aList).Error
	if err != nil {
		return nil, err
	}
	return aList, nil
}

// GetAddressById 查询单个地址
func GetAddressById(ctx context.Context, id, userId int) (*UserAddressDO, error) {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return nil, err
	}
	out := new(UserAddressDO)
	err = db.Table(out.TableName()).Where("id = ? AND user_id = ? AND is_del = 0", id, userId).Take(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetDefaultAddress 查询默认收货地址
func GetDefaultAddress(ctx context.Context, userId int) (*UserAddressDO, error) {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return nil, err
	}
	out := new(UserAddressDO)
	err = db.Table(out.TableName()).Where("user_id = ? AND is_del = 0 AND is_default = 1", userId).Take(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddUserAddress 新增收货地址
func AddUserAddress(ctx context.Context, do *UserAddressDO) error {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return err
	}
	err = db.Table(do.TableName()).Create(do).Error
	return err
}

// UpdateUserAddress 更新收货地址
func UpdateUserAddress(ctx context.Context, do *UserAddressDO) error {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return err
	}
	err = db.Table(do.TableName()).Updates(do).Error
	return err
}

// ClearDefaultUserAddress 清理用户默认收货地址
func ClearDefaultUserAddress(ctx context.Context, userId int) error {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return err
	}
	empty := new(UserAddressDO)
	return db.Table(empty.TableName()).Where("user_id = ?", userId).Update("is_default", 0).Error
}
