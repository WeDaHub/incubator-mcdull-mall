package model

import (
	"App-CloudBase-mcdull-mall/pkg/database"
	"context"
	"time"
)

// UserDO 小程序用户
type UserDO struct {
	Id         int       `gorm:"column:id" json:"id"`
	Openid     string    `gorm:"column:openid" json:"openid"`     // 微信openid
	Nickname   string    `gorm:"column:nickname" json:"nickname"` // 昵称
	Avatar     string    `gorm:"column:avatar" json:"avatar"`     // 头像
	Mobile     string    `gorm:"column:mobile" json:"mobile"`     // 手机号
	City       string    `gorm:"column:city" json:"city"`         // 城市编码
	Province   string    `gorm:"column:province" json:"province"` // 省份
	Country    string    `gorm:"column:country" json:"country"`   // 国家
	Gender     int       `gorm:"column:gender" json:"gender"`     // 性别 0：未知、1：男、2：女
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

func (u *UserDO) TableName() string {
	return "wechat_mall_user"
}

// FindUser 查找用户记录
func (u *UserDO) FindUser(ctx context.Context) (*UserDO, error) {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return nil, err
	}
	out := new(UserDO)
	err = db.Take(out, u).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreateUser 新增用户
func CreateUser(ctx context.Context, openid string) (int, error) {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return 0, err
	}
	u := &UserDO{
		Openid:     openid,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	err = db.Table(u.TableName()).Create(u).Error
	return u.Id, err
}

// UpdateUser 更新用户
func (u *UserDO) UpdateUser(ctx context.Context) error {
	db, err := database.GetGormDB(ctx)
	if err != nil {
		return err
	}
	err = db.Table(u.TableName()).Updates(u).Error
	return err
}
