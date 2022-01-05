package model

import (
	"context"
	"testing"
	"time"
)

func Test_CreateVisitorRecord(t *testing.T) {
	ctx := context.Background()
	do := &VisitorRecordDO{
		UserId:     11111,
		Ip:         "127.0.0.1",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	err := CreateVisitorRecord(ctx, do)
	if err != nil {
		t.Fatal(err)
	}
}
