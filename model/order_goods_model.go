package model

import (
	"App-CloudBase-mcdull-mall/pkg/database"
	"context"
	"time"
)

// OrderGoodsDO 订单商品
type OrderGoodsDO struct {
	Id         int       `gorm:"column:id" json:"id"`
	OrderNo    string    `gorm:"column:order_no" json:"order_no"`       // 订单号
	UserId     int       `gorm:"column:user_id" json:"user_id"`         // 用户ID
	GoodsId    int       `gorm:"column:goods_id" json:"goods_id"`       // 商品ID
	SkuId      int       `gorm:"column:sku_id" json:"sku_id"`           // sku ID
	Picture    string    `gorm:"column:picture" json:"picture"`         // 图片
	Title      string    `gorm:"column:title" json:"title"`             // 标题
	Price      string    `gorm:"column:price" json:"price"`             // 价格
	Specs      string    `gorm:"column:specs" json:"specs"`             // sku属性
	Num        int       `gorm:"column:num" json:"num"`                 // 数量
	LockStatus int       `gorm:"column:lock_status" json:"lock_status"` // 锁定状态：0-锁定 1-解锁
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

func (o *OrderGoodsDO) TableName() string {
	return "wechat_mall_order_goods"
}

// CountBuyUserNum 统计商品购买人数
func CountBuyUserNum(ctx context.Context, goodsId int) (int, error) {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return 0, err
	}
	empty := new(OrderGoodsDO)
	num := 0
	err = db.Table(empty.TableName()).Select("COUNT(DISTINCT user_id)").Where("lock_status = 1 AND goods_id = ?", goodsId).Find(&num).Error
	if err != nil {
		return 0, err
	}
	return num, nil
}

// AddOrderGoods 订单商品快照
func AddOrderGoods(ctx context.Context, do *OrderGoodsDO) error {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return err
	}
	err = db.Table(do.TableName()).Create(do).Error
	return err
}

// QueryOrderGoods 查询订单商品
func QueryOrderGoods(ctx context.Context, orderNo string) ([]*OrderGoodsDO, error) {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return nil, err
	}
	empty := new(OrderGoodsDO)
	gList := make([]*OrderGoodsDO, 0)
	err = db.Table(empty.TableName()).Where("order_no = ?", orderNo).Find(&gList).Error
	if err != nil {
		return nil, err
	}
	return gList, nil
}
