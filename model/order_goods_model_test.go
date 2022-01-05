package model

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func Test_CountBuyUserNum(t *testing.T) {
	ctx := context.Background()
	num, err := CountBuyUserNum(ctx, 7)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(num)
}

// AddOrderGoods 订单商品快照
func Test_AddOrderGoods(t *testing.T) {
	ctx := context.Background()
	do := &OrderGoodsDO{
		OrderNo:    "12345",
		Price:      "10",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	err := AddOrderGoods(ctx, do)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_QueryOrderGoods(t *testing.T) {
	ctx := context.Background()
	gList, err := QueryOrderGoods(ctx, "20200503144717754429")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(gList))
}
