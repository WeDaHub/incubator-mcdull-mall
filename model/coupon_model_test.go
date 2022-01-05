package model

import (
	"context"
	"fmt"
	"testing"
)

func Test_GetCouponList(t *testing.T) {
	ctx := context.Background()
	cList, err := GetCouponList(ctx)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(cList))
	if len(cList) > 0 {
		fmt.Println(cList[0].StartTime.Format("2006-01-02 15:04:05"))
	}
}
