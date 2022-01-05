package model

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func Test_AddOrder(t *testing.T) {
	ctx := context.Background()
	orderDO := &OrderDO{}
	orderDO.OrderNo = "12345"
	orderDO.UserId = 1
	orderDO.PayAmount = "100"
	orderDO.GoodsAmount = "100"
	orderDO.DiscountAmount = "100"
	orderDO.DispatchAmount = "100"
	orderDO.PayTime = time.Now()
	orderDO.DeliverTime = time.Now()
	orderDO.FinishTime = time.Now()
	orderDO.Status = 0
	orderDO.AddressId = 1
	orderDO.AddressSnapshot = ""
	orderDO.WxappPrepayId = ""
	orderDO.CreateTime = time.Now()
	orderDO.UpdateTime = time.Now()
	err := AddOrder(ctx, orderDO)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_ListOrderByParams(t *testing.T) {
	ctx := context.Background()
	cList, err := ListOrderByParams(ctx, 1, 0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(cList))
}

func Test_QueryOrderByOrderNo(t *testing.T) {
	ctx := context.Background()
	do, err := QueryOrderByOrderNo(ctx, "20200503144717754429")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(do)
}

func Test_QueryOrderById(t *testing.T) {
	ctx := context.Background()
	do, err := QueryOrderById(ctx, 1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(do)
}

func Test_UpdateOrder(t *testing.T) {
	ctx := context.Background()
	do := &OrderDO{
		Id:         1,
		Status:     3,
		UpdateTime: time.Now(),
	}
	err := UpdateOrder(ctx, do)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_CountOrderNum(t *testing.T) {
	ctx := context.Background()
	num, err := CountOrderNum(ctx, 1, 3)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(num)
}
