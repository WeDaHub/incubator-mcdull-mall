package model

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func Test_FindUser(t *testing.T) {
	ctx := context.Background()
	c := new(UserDO)
	c.Id = 1

	out, err := c.FindUser(ctx)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v", out)
}

func Test_CreateUser(t *testing.T) {
	ctx := context.Background()
	autoid, err := CreateUser(ctx, "12345")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v", autoid)
}

func Test_UpdateUser(t *testing.T) {
	ctx := context.Background()
	u := &UserDO{
		Id:         2,
		Nickname:   "mars",
		UpdateTime: time.Now(),
	}
	err := u.UpdateUser(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
