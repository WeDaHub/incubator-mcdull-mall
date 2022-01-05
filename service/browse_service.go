package service

import (
	"App-CloudBase-mcdull-mall/dao"
	"App-CloudBase-mcdull-mall/model"
	"context"
	"log"
	"time"
)

// ListBrowseRecord 查询浏览记录
func ListBrowseRecord(ctx context.Context, userId, page, size int) ([]*model.GoodsBrowseRecordDO, int, error) {
	records, err := dao.GetGoodsBrowseRecord(ctx, userId)
	if err != nil {
		log.Printf("call GetGoodsBrowseRecord failed, err:%v", err)
		return nil, 0, err
	}
	total := len(records)
	offset, endpos := pagePos(total, page, size)
	return records[offset:endpos], total, nil
}

// ClearBrowseHistory 清理浏览记录
func ClearBrowseHistory(ctx context.Context, ids []int) error {
	for _, id := range ids {
		err := dao.DeleteGoodsBrowseRecord(ctx, id)
		if err != nil {
			return err
		}
	}
	return nil
}

// RecordGoodsBrowse 记录商品浏览
func RecordGoodsBrowse(ctx context.Context, userId, goodsId int) {
	// 1.清理旧的记录
	err := clearOldBrowseRecord(ctx, userId, goodsId)
	if err != nil {
		log.Printf("call clearOldBrowseRecord failed, err:%v", err)
		return
	}
	// 2.查询商品
	goods, err := dao.GetGoodsById(ctx, goodsId)
	if err != nil {
		log.Printf("call GetGoodsById failed, err:%v", err)
		return
	}
	// 3.组织数据
	browseDO := &model.GoodsBrowseRecordDO{}
	browseDO.UserId = userId
	browseDO.GoodsId = goods.Id
	browseDO.Picture = goods.Picture
	browseDO.Title = goods.Title
	browseDO.Price = goods.Price
	browseDO.CreateTime = time.Now()
	browseDO.UpdateTime = time.Now()
	err = dao.AddGoodsBrowseRecord(ctx, browseDO)
	if err != nil {
		log.Printf("call AddGoodsBrowseRecord failed, err:%v", err)
	}
}

// clearOldBrowseRecord 清理历史浏览记录（只保留最新的一条）
func clearOldBrowseRecord(ctx context.Context, userId, goodId int) error {
	record, err := dao.GetBrowseRecordByGoodsId(ctx, userId, goodId)
	if err == dao.NotFoundRecord {
		return nil
	}
	if err != nil {
		log.Printf("call GetBrowseRecordByGoodsId failed, err:%v", err)
		return err
	}
	return dao.DeleteGoodsBrowseRecord(ctx, record.Id)
}
