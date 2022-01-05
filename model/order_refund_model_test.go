package model

import (
	"App-CloudBase-mcdull-mall/utils"
	"context"
	"fmt"
	"testing"
	"time"
)

func Test_QueryOrderRefundRecord(t *testing.T) {
	ctx := context.Background()
	record, err := QueryOrderRefundRecord(ctx, "12345")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(record)
}

func Test_QueryRefundRecord(t *testing.T) {
	ctx := context.Background()
	record, err := QueryRefundRecord(ctx, "20210612191332559814")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(record)
}

func Test_UpdateRefundApply(t *testing.T) {
	ctx := context.Background()
	do := &OrderRefund{
		Id:         1,
		Reason:     "test",
		UpdateTime: time.Now(),
	}
	err := UpdateRefundApply(ctx, do)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_AddRefundRecord(t *testing.T) {
	ctx := context.Background()
	refundNo := time.Now().Format("20060102150405") + utils.RandomNumberStr(6)
	refund := &OrderRefund{}
	refund.RefundNo = refundNo
	refund.UserId = 1
	refund.OrderNo = "12345"
	refund.Reason = "todo"
	refund.RefundTime = time.Now()
	refund.RefundAmount = "100"
	refund.CreateTime = time.Now()
	refund.UpdateTime = time.Now()
	autoId, err := AddRefundRecord(ctx, refund)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(autoId)
}
