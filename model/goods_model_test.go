package model

import (
	"context"
	"fmt"
	"testing"
)

func Test_GetGoodsList(t *testing.T) {
	ctx := context.Background()
	gList, err := GetGoodsList(ctx)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(gList))
}

func Test_UpdateGoodsSaleNum(t *testing.T) {
	ctx := context.Background()
	err := UpdateGoodsSaleNum(ctx, 1, 1)
	if err != nil {
		t.Fatal(err)
	}
}
