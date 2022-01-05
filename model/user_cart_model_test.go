package model

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func Test_GetUserCartList(t *testing.T) {
	ctx := context.Background()
	list, err := GetUserCartList(ctx, 1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(list))
}

func Test_CreateUserCart(t *testing.T) {
	ctx := context.Background()
	userCartDO := &UserCartDO{}
	userCartDO.UserId = 2
	userCartDO.GoodsId = 1
	userCartDO.SkuId = 1
	userCartDO.Num = 1
	userCartDO.CreateTime = time.Now()
	userCartDO.UpdateTime = time.Now()
	err := CreateUserCart(ctx, userCartDO)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_UpdateUserCart(t *testing.T) {
	ctx := context.Background()
	userCartDO := &UserCartDO{
		Id:         1,
		Num:        4,
		UpdateTime: time.Now(),
	}
	err := UpdateUserCart(ctx, userCartDO)
	if err != nil {
		t.Fatal(err)
	}
}
