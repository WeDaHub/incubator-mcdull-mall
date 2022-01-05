package service

import (
	"App-CloudBase-mcdull-mall/dao"
	"App-CloudBase-mcdull-mall/model"
	"App-CloudBase-mcdull-mall/utils"
	"context"
	"log"
	"net/http"
	"time"
)

// RecordVisitorRecord 记录访客
func RecordVisitorRecord(ctx context.Context, userId int, r *http.Request) {
	ip := utils.ReadUserIP(r)
	record := &model.VisitorRecordDO{}
	record.UserId = userId
	record.Ip = ip
	record.CreateTime = time.Now()
	record.UpdateTime = time.Now()
	err := dao.AddVisitorRecord(ctx, record)
	if err != nil {
		log.Printf("call AddVisitorRecord failed, err:%v", err)
	}
}
