package dao

import (
	"App-CloudBase-mcdull-mall/model"
	"context"

	"gorm.io/gorm"
)

// GetGoodsBrowseRecord 查询商品浏览记录（降序）
func GetGoodsBrowseRecord(ctx context.Context, userId int) ([]*model.GoodsBrowseRecordDO, error) {
	return model.GetGoodsBrowseRecord(ctx, userId)
}

// GetBrowseRecordByGoodsId 查询商品浏览记录（指定商品）
func GetBrowseRecordByGoodsId(ctx context.Context, userId, goodsId int) (*model.GoodsBrowseRecordDO, error) {
	record, err := model.GetBrowseRecordByGoodsId(ctx, userId, goodsId)
	if err == gorm.ErrRecordNotFound {
		return nil, NotFoundRecord
	}
	if err != nil {
		return nil, err
	}
	return record, nil
}

// AddGoodsBrowseRecord 新增商品浏览记录
func AddGoodsBrowseRecord(ctx context.Context, do *model.GoodsBrowseRecordDO) error {
	return model.AddGoodsBrowseRecord(ctx, do)
}

// DeleteGoodsBrowseRecord 删除浏览记录
func DeleteGoodsBrowseRecord(ctx context.Context, id int) error {
	return model.DeleteGoodsBrowseRecord(ctx, id)
}
