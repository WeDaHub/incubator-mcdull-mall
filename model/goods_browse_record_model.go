package model

import (
	"App-CloudBase-mcdull-mall/pkg/database"
	"context"
	"time"
)

// GoodsBrowseRecordDO 商城-商品浏览记录
type GoodsBrowseRecordDO struct {
	Id         int       `gorm:"column:id" json:"id"`
	UserId     int       `gorm:"column:user_id" json:"user_id"`   // 用户ID
	GoodsId    int       `gorm:"column:goods_id" json:"goods_id"` // 商品ID
	Picture    string    `gorm:"column:picture" json:"picture"`   // 商品图片
	Title      string    `gorm:"column:title" json:"title"`       // 商品名称
	Price      string    `gorm:"column:price" json:"price"`       // 商品价格
	Del        int       `gorm:"column:is_del" json:"del"`        // 是否删除：0-否 1-是
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

func (g *GoodsBrowseRecordDO) TableName() string {
	return "wechat_mall_goods_browse_record"
}

// GetGoodsBrowseRecord 查询商品浏览记录（降序）
func GetGoodsBrowseRecord(ctx context.Context, userId int) ([]*GoodsBrowseRecordDO, error) {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return nil, err
	}
	empty := new(GoodsBrowseRecordDO)
	rList := make([]*GoodsBrowseRecordDO, 0)
	err = db.Table(empty.TableName()).Where("is_del = 0 AND user_id = ?", userId).Order("id DESC").Find(&rList).Error
	if err != nil {
		return nil, err
	}
	return rList, nil
}

// GetBrowseRecordByGoodsId 查询商品浏览记录（指定商品）
func GetBrowseRecordByGoodsId(ctx context.Context, userId, goodsId int) (*GoodsBrowseRecordDO, error) {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return nil, err
	}
	out := new(GoodsBrowseRecordDO)
	err = db.Table(out.TableName()).Where("is_del = 0 AND user_id = ? AND goods_id = ?", userId, goodsId).Take(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddGoodsBrowseRecord 新增商品浏览记录
func AddGoodsBrowseRecord(ctx context.Context, do *GoodsBrowseRecordDO) error {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return err
	}
	return db.Table(do.TableName()).Create(do).Error
}

// DeleteGoodsBrowseRecord 删除浏览记录
func DeleteGoodsBrowseRecord(ctx context.Context, id int) error {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return err
	}
	do := &GoodsBrowseRecordDO{
		Id:         id,
		Del:        1,
		UpdateTime: time.Now(),
	}
	return db.Table(do.TableName()).Updates(do).Error
}
