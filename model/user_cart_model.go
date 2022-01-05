package model

import (
	"App-CloudBase-mcdull-mall/pkg/database"
	"context"
	"time"
)

// UserCartDO 商城-购物车
type UserCartDO struct {
	Id         int       `gorm:"column:id" json:"id"`
	UserId     int       `gorm:"column:user_id" json:"user_id"`   // 用户ID
	GoodsId    int       `gorm:"column:goods_id" json:"goods_id"` // 商品ID
	SkuId      int       `gorm:"column:sku_id" json:"sku_id"`     // sku ID
	Num        int       `gorm:"column:num" json:"num"`           // 数量
	Del        int       `gorm:"column:is_del" json:"del"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

func (u *UserCartDO) TableName() string {
	return "wechat_mall_user_cart"
}

// GetUserCartList 查询用户购物车
func GetUserCartList(ctx context.Context, userId int) ([]*UserCartDO, error) {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return nil, err
	}
	empty := new(UserCartDO)
	cList := make([]*UserCartDO, 0)
	err = db.Table(empty.TableName()).Where("is_del = 0 AND user_id = ?", userId).Find(&cList).Error
	if err != nil {
		return nil, err
	}
	return cList, nil
}

// CreateUserCart 新增购物车商品
func CreateUserCart(ctx context.Context, cartDO *UserCartDO) error {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return err
	}
	err = db.Table(cartDO.TableName()).Create(cartDO).Error
	return err
}

// UpdateUserCart 更新购物车商品
func UpdateUserCart(ctx context.Context, cartDO *UserCartDO) error {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return err
	}
	err = db.Table(cartDO.TableName()).Updates(cartDO).Error
	return err
}
