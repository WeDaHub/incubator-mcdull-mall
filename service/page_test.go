package service

import (
	"fmt"
	"testing"
)

func Test_PagePos(t *testing.T) {
	total := 10
	page := 2
	size := 3
	offset, endpos := pagePos(total, page, size)
	fmt.Printf("offset:%d, endpos:%d\n", offset, endpos)
}
