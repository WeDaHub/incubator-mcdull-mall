package model

import (
	"context"
	"fmt"
	"testing"
)

func Test_GetSkuSpecAttrList(t *testing.T) {
	ctx := context.Background()
	sList, err := GetSkuSpecAttrList(ctx)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(sList))
}
