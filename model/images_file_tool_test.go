package model

import (
	"App-CloudBase-mcdull-mall/pkg/database"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"testing"
)

// =================== 警告 请勿执行下面的单元测试 ====================
// 目的：为了将图片资源打入容器，因此需要调整Db表中的图片路径，执行步骤如下：
//
// 1.图片路径-使用本地文件路径
// 2.mux路由注册静态资源访问
// 3.修改Dockerfile 将图片资源打到容器中
// 4.配置文件-访问域名
// 5.接口处理图片路径（前缀域名）

// Test_BannerImages 1.处理Banner图片路径
func Test_BannerImages(t *testing.T) {
	ctx := context.Background()
	bannerList, err := GetBannerList(ctx, 1)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range bannerList {
		// 1.组织数据
		url := v.Picture
		filename := url[strings.LastIndex(url, "/")+1:]
		subdir := "banner"
		table := "wechat_mall_banner"

		// 2.更新路径
		db, _ := database.GetGormDB(ctx)
		err = db.Table(table).Where("id = ?", v.Id).Update("picture", "/assets/images/"+subdir+"/"+filename).Error
		if err != nil {
			t.Fatal(err)
		}
		// 3.下载图片
		download(url, filename, subdir)
	}
}

// Test_CategoryImages 2.处理category图片路径
func Test_CategoryImages(t *testing.T) {
	ctx := context.Background()
	cList, err := GetAllCategoryList(ctx)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range cList {
		// 1.组织数据
		url := v.Picture
		filename := url[strings.LastIndex(url, "/")+1:]
		subdir := "category"
		table := "wechat_mall_category"

		// 2.更新路径
		db, _ := database.GetGormDB(ctx)
		err = db.Table(table).Where("id = ?", v.Id).Update("picture", "/assets/images/"+subdir+"/"+filename).Error
		if err != nil {
			t.Fatal(err)
		}
		// 3.下载图片
		download(url, filename, subdir)
	}
}

// Test_GridCategoryImages 3.处理grid_category图片路径
func Test_GridCategoryImages(t *testing.T) {
	ctx := context.Background()
	cList, err := GetAllGridCategoryList(ctx)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range cList {
		// 1.组织数据
		url := v.Picture
		filename := url[strings.LastIndex(url, "/")+1:]
		subdir := "grid_category"
		table := "wechat_mall_grid_category"

		// 2.更新路径
		db, _ := database.GetGormDB(ctx)
		err = db.Table(table).Where("id = ?", v.Id).Update("picture", "/assets/images/"+subdir+"/"+filename).Error
		if err != nil {
			t.Fatal(err)
		}
		// 3.下载图片
		download(url, filename, subdir)
	}
}

// Test_SkuImages 4.处理sku图片路径
func Test_SkuImages(t *testing.T) {
	ctx := context.Background()
	cList, err := GetSkuList(ctx)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range cList {
		// 1.组织数据
		url := v.Picture
		filename := url[strings.LastIndex(url, "/")+1:]
		subdir := "sku"
		table := "wechat_mall_sku"

		// 2.更新路径
		db, _ := database.GetGormDB(ctx)
		err = db.Table(table).Where("id = ?", v.Id).Update("picture", "/assets/images/"+subdir+"/"+filename).Error
		if err != nil {
			t.Fatal(err)
		}
		// 3.下载图片
		download(url, filename, subdir)
	}
}

// Test_GoodsImages 5.处理商品图片路径
func Test_GoodsImages(t *testing.T) {
	ctx := context.Background()
	cList, err := GetGoodsList(ctx)
	if err != nil {
		t.Fatal(err)
	}
	subdir := "goods"
	table := "wechat_mall_goods"

	// 处理picture字段
	for _, v := range cList {
		// 1.组织数据
		url := v.Picture
		filename := url[strings.LastIndex(url, "/")+1:]

		// 2.下载图片
		download(url, filename, subdir)

		// 3.更新路径
		db, _ := database.GetGormDB(ctx)
		err = db.Table(table).Where("id = ?", v.Id).Update("picture", "/assets/images/"+subdir+"/"+filename).Error
		if err != nil {
			t.Fatal(err)
		}
	}
	// 处理banner_picture字段
	for _, v := range cList {
		bList := make([]string, 0)
		_ = json.Unmarshal([]byte(v.BannerPicture), &bList)
		nbList := make([]string, 0)
		for _, v := range bList {
			// 1.处理url
			url := v
			filename := url[strings.LastIndex(url, "/")+1:]
			// 2.下载图片
			download(url, filename, subdir)
			// 3.新路径
			nbList = append(nbList, "/assets/images/"+subdir+"/"+filename)
		}
		// 4.更新路径（数组）
		nbListBytes, _ := json.Marshal(nbList)
		db, _ := database.GetGormDB(ctx)
		err = db.Table(table).Where("id = ?", v.Id).Update("banner_picture", string(nbListBytes)).Error
		if err != nil {
			t.Fatal(err)
		}
	}

	// 处理detail_picture字段
	for _, v := range cList {
		bList := make([]string, 0)
		_ = json.Unmarshal([]byte(v.DetailPicture), &bList)
		nbList := make([]string, 0)
		for _, v := range bList {
			// 1.处理url
			url := v
			filename := url[strings.LastIndex(url, "/")+1:]
			// 2.下载图片
			download(url, filename, subdir)
			// 3.新路径
			nbList = append(nbList, "/assets/images/"+subdir+"/"+filename)
		}
		// 4.更新路径（数组）
		nbListBytes, _ := json.Marshal(nbList)
		db, _ := database.GetGormDB(ctx)
		err = db.Table(table).Where("id = ?", v.Id).Update("detail_picture", string(nbListBytes)).Error
		if err != nil {
			t.Fatal(err)
		}
	}
}

func download(url, filename, subdir string) {
	res, err := http.Get(url)
	if err != nil {
		log.Printf("http.Get -> %v", err)
		return
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("ioutil.ReadAll -> %s", err.Error())
		return
	}
	defer res.Body.Close()
	if err = ioutil.WriteFile("assets/images"+string(filepath.Separator)+subdir+string(filepath.Separator)+filename, data, 0777); err != nil {
		log.Println("Error Saving:", filename, err)
	} else {
		log.Println("Saved:", filename)
	}
}
