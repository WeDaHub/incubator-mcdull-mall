package dao

import (
	"App-CloudBase-mcdull-mall/model"
	"context"
)

// AddVisitorRecord 增加访客记录
func AddVisitorRecord(ctx context.Context, do *model.VisitorRecordDO) error {
	return model.CreateVisitorRecord(ctx, do)
}
