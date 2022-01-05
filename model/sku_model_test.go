package model

import (
	"context"
	"fmt"
	"testing"
)

func Test_GetSkuList(t *testing.T) {
	ctx := context.Background()
	list, err := GetSkuList(ctx)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(list))
}

func Test_UpdateSkuStockById(t *testing.T) {
	ctx := context.Background()
	err := UpdateSkuStockById(ctx, 1, 8)
	if err != nil {
		t.Fatal(err)
	}
}
