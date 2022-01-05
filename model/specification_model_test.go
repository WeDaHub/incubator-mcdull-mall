package model

import (
	"context"
	"fmt"
	"testing"
)

func Test_GetSpecList(t *testing.T) {
	ctx := context.Background()
	list, err := GetSpecList(ctx)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(list))
}
