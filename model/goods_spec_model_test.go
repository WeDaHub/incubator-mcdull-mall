package model

import (
	"context"
	"fmt"
	"testing"
)

func Test_GetGoodsSpecList(t *testing.T) {
	ctx := context.Background()
	list, err := GetGoodsSpecList(ctx)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(list))
}
