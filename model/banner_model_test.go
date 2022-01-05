package model

import (
	"App-CloudBase-mcdull-mall/pkg/database"
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
)

func init() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// 切换working目录
	os.Chdir(path[:strings.LastIndex(path, "/")])
	path, err = os.Getwd()
	if err != nil {
		panic(err)
	}
}

func Test_GetBanner(t *testing.T) {
	ctx := context.Background()
	db, err := database.GetGormDB(ctx)
	if err != nil {
		t.Fatal(err)
	}
	banner := new(BannerDO)
	err = db.Table("wechat_mall_banner").Where("id = ?", 1).Find(banner).Error
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v", banner)
}

func Test_GetBannerList(t *testing.T) {
	ctx := context.Background()
	bannerList, err := GetBannerList(ctx, 1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v", bannerList)
}
