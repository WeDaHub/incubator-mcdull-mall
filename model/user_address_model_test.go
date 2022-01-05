package model

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func Test_GetUserAddressList(t *testing.T) {
	ctx := context.Background()
	aList, err := GetUserAddressList(ctx, 1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(aList))
}

func Test_GetAddressById(t *testing.T) {
	ctx := context.Background()
	addr, err := GetAddressById(ctx, 1, 1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(addr)
}

func Test_GetDefaultAddress(t *testing.T) {
	ctx := context.Background()
	addr, err := GetDefaultAddress(ctx, 1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(addr)
}

func Test_AddUserAddress(t *testing.T) {
	ctx := context.Background()
	addr := &UserAddressDO{
		UserId:     1,
		Contacts:   "mars",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	err := AddUserAddress(ctx, addr)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_UpdateUserAddress(t *testing.T) {
	ctx := context.Background()
	addr := &UserAddressDO{
		Id:         2,
		UserId:     2,
		Contacts:   "mars",
		UpdateTime: time.Now(),
	}
	err := UpdateUserAddress(ctx, addr)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_ClearDefaultUserAddress(t *testing.T) {
	ctx := context.Background()
	err := ClearDefaultUserAddress(ctx, 1)
	if err != nil {
		t.Fatal(err)
	}
}
