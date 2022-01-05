package model

import (
	"context"
	"fmt"
	"testing"
)

func Test_GetSpecAttrList(t *testing.T) {
	ctx := context.Background()
	list, err := GetSpecAttrList(ctx)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(list))
}
