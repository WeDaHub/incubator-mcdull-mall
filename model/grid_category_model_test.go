package model

import (
	"context"
	"fmt"
	"testing"
)

func Test_GetGridCategoryList(t *testing.T) {
	ctx := context.Background()
	cList, err := GetGridCategoryList(ctx)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(cList))
}
