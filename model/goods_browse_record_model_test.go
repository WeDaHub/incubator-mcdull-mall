package model

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func Test_GetGoodsBrowseRecord(t *testing.T) {
	ctx := context.Background()
	rList, err := GetGoodsBrowseRecord(ctx, 1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(rList))
}

func Test_GetBrowseRecordByGoodsId(t *testing.T) {
	ctx := context.Background()
	record, err := GetBrowseRecordByGoodsId(ctx, 1, 7)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(record)
}

func Test_AddGoodsBrowseRecord(t *testing.T) {
	ctx := context.Background()
	browseDO := &GoodsBrowseRecordDO{}
	browseDO.UserId = 1
	browseDO.GoodsId = 9
	browseDO.Price = "200"
	browseDO.CreateTime = time.Now()
	browseDO.UpdateTime = time.Now()
	err := AddGoodsBrowseRecord(ctx, browseDO)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_DeleteGoodsBrowseRecord(t *testing.T) {
	ctx := context.Background()
	err := DeleteGoodsBrowseRecord(ctx, 20)
	if err != nil {
		t.Fatal(err)
	}
}
