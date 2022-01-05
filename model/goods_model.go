package model

import (
	"App-CloudBase-mcdull-mall/pkg/database"
	"context"
	"time"
)

// GoodsDO 商城-商品
type GoodsDO struct {
	Id            int       `gorm:"column:id" json:"id"`
	BrandName     string    `gorm:"column:brand_name" json:"brand_name"`         // 品牌
	Title         string    `gorm:"column:title" json:"title"`                   // 标题
	Price         string    `gorm:"column:price" json:"price"`                   // 价格
	DiscountPrice string    `gorm:"column:discount_price" json:"discount_price"` // 折扣
	CategoryId    int       `gorm:"column:category_id" json:"category_id"`       // 分类ID
	Online        int       `gorm:"column:online" json:"online"`                 // 是否上线
	Picture       string    `gorm:"column:picture" json:"picture"`               // 主图
	BannerPicture string    `gorm:"column:banner_picture" json:"banner_picture"` // 详情图
	DetailPicture string    `gorm:"column:detail_picture" json:"detail_picture"` // 轮播图
	Tags          string    `gorm:"column:tags" json:"tags"`                     // 标签
	Description   string    `gorm:"column:description" json:"description"`       // 详情
	SaleNum       int       `gorm:"column:sale_num" json:"sale_num"`             // 销量
	Del           int       `gorm:"column:is_del" json:"del"`
	UpdateTime    time.Time `gorm:"column:update_time" json:"update_time"`
}

func (g *GoodsDO) TableName() string {
	return "wechat_mall_goods"
}

// GetGoodsList 查询全量的商品
func GetGoodsList(ctx context.Context) ([]*GoodsDO, error) {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return nil, err
	}
	empty := new(GoodsDO)
	gList := make([]*GoodsDO, 0)
	err = db.Table(empty.TableName()).Find(&gList).Error
	if err != nil {
		return nil, err
	}
	return gList, nil
}

// UpdateGoodsSaleNum 更新商品销量
func UpdateGoodsSaleNum(ctx context.Context, gid int, num int) error {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return err
	}
	do := &GoodsDO{
		Id:         gid,
		SaleNum:    num,
		UpdateTime: time.Now(),
	}
	err = db.Table(do.TableName()).Updates(do).Error
	return err
}
