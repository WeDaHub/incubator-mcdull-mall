package database

import (
	"App-CloudBase-mcdull-mall/env"
	"context"
	"fmt"
	"log"
	"sync"

	"gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	MaxIdleConns    = 20
	MaxOpenConns    = 100
	ConnMaxLifetime = 1800
)

var gormDb *gorm.DB

type MySQL struct {
	db   *gorm.DB
	once sync.Once
}

// GetGormDB 获取Gorm连接
func GetGormDB(ctx context.Context) (*gorm.DB, error) {
	if gormDb != nil {
		return gormDb, nil
	}
	m := new(MySQL)
	dsn := genDsn(ctx)
	db, err := m.connect(dsn)
	if err != nil {
		log.Printf("call connect failed, err:%v", err)
		return nil, err
	}
	gormDb = db
	return gormDb, nil
}

func (m *MySQL) connect(dsn string) (db *gorm.DB, err error) {
	m.once.Do(func() {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			log.Fatal(err)
		}
		pool, err := db.DB()
		if err != nil {
			log.Fatal(err)
		}
		pool.SetMaxIdleConns(MaxIdleConns)
		pool.SetMaxOpenConns(MaxOpenConns)
		pool.SetConnMaxLifetime(ConnMaxLifetime)
		m.db = db
	})
	return
}

func genDsn(ctx context.Context) string {
	conf := env.LoadConf()
	user := conf.Db.User
	pass := conf.Db.Pass
	addr := conf.Db.Addr
	dbname := conf.Db.DbName
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, addr, dbname)
}
