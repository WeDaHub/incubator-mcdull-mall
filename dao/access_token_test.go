package dao

import (
	"context"
	"fmt"
	"testing"
)

func Test_GetAccessToken(t *testing.T) {
	for i := 0; i < 5; i++ {
		accessToken, err := GetAccessTokenCache(context.Background(), false)
		if err != nil {
			fmt.Printf("getAccessToken failed, err:%v", err)
			return
		}
		fmt.Printf("%s\n", accessToken)
	}
}
