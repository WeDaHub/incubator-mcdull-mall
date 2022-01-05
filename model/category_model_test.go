package model

import (
	"App-CloudBase-mcdull-mall/defs"
	"context"
	"fmt"
	"testing"
)

func Test_GetCategoryList(t *testing.T) {
	ctx := context.Background()
	cList, err := GetCategoryList(ctx, defs.ONLINE)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(cList))
}
