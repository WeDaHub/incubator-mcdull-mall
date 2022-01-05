package service

import (
	"App-CloudBase-mcdull-mall/dao"
	"App-CloudBase-mcdull-mall/defs"
	"App-CloudBase-mcdull-mall/env"
	"App-CloudBase-mcdull-mall/errs"
	"App-CloudBase-mcdull-mall/model"
	"context"
	"log"
	"math"
	"strconv"
	"time"
)

// GetCartGoods 查询用户购物车
func GetCartGoods(ctx context.Context, userId, page, size int) ([]*defs.PortalCartGoodsVO, int, error) {
	// 1.查询购物车记录
	rList, err := dao.GetUserCartList(ctx, userId)
	if err != nil {
		log.Printf("call GetUserCartList failed err:%v", err)
		return nil, 0, err
	}
	// 2.剔除无效的商品和SKU
	cartGoodsVOList := make([]*defs.PortalCartGoodsVO, 0)
	for _, v := range rList {
		// 查询商品
		goodsDO, err := dao.GetGoodsById(ctx, v.GoodsId)
		if err != nil {
			log.Printf("call GetGoodsById failed, err:%v", err)
			return nil, 0, err
		}
		// 查询SKU
		skuDO, err := dao.GetSkuListBySkuId(ctx, v.SkuId)
		if err != nil {
			log.Printf("call GetSkuListBySkuId failed, err:%v", err)
			return nil, 0, err
		}
		status := 0
		if goodsDO.Id == defs.ZERO || goodsDO.Del == defs.DELETE || goodsDO.Online == defs.OFFLINE ||
			skuDO.Id == defs.ZERO || skuDO.Del == defs.DELETE || skuDO.Online == defs.OFFLINE {
			status = 2
		} else {
			if skuDO.Stock < v.Num {
				status = 1
			}
		}
		domain := env.LoadConf().Domain
		cartGoodsVO := &defs.PortalCartGoodsVO{}
		cartGoodsVO.Id = v.Id
		cartGoodsVO.GoodsId = v.GoodsId
		cartGoodsVO.SkuId = v.SkuId
		cartGoodsVO.Title = goodsDO.Title
		cartGoodsVO.Price, _ = strconv.ParseFloat(skuDO.Price, 2)
		cartGoodsVO.Picture = domain + skuDO.Picture
		cartGoodsVO.Specs = skuDO.Specs
		cartGoodsVO.Num = v.Num
		cartGoodsVO.Status = status
		cartGoodsVOList = append(cartGoodsVOList, cartGoodsVO)
	}
	// 3.分页
	total := len(cartGoodsVOList)
	offset, endpos := pagePos(total, page, size)
	return cartGoodsVOList[offset:endpos], total, nil
}

// DoEditCart 编辑购物车
func DoEditCart(ctx context.Context, userId, goodsId, skuId, num int) error {
	if math.Abs(float64(num)) > defs.CartMax {
		return errs.ErrorParameterValidate
	}
	// 1.查询商品
	goodsDO, err := dao.GetGoodsById(ctx, goodsId)
	if err != nil {
		if err == dao.NotFoundRecord {
			return errs.ErrorGoods
		}
		return errs.ErrorInternalFaults
	}
	if goodsDO.Del == defs.DELETE || goodsDO.Online == defs.OFFLINE {
		return errs.ErrorGoods
	}
	// 2.查询SKU
	skuDO, err := dao.GetSkuListBySkuId(ctx, skuId)
	if err != nil {
		if err == dao.NotFoundRecord {
			return errs.ErrorSKU
		}
		return errs.ErrorInternalFaults
	}
	if skuDO.Id == defs.ZERO || skuDO.Del == defs.DELETE || skuDO.Online == defs.OFFLINE {
		panic(errs.ErrorSKU)
	}
	if skuDO.Stock <= 0 {
		return errs.NewErrorGoodsCart("库存不足！")
	}
	// 3.查询购物车记录
	cartDO, err := dao.GetUserCartListByParams(ctx, userId, goodsId, skuId)
	if err != nil && err != dao.NotFoundRecord {
		return errs.ErrorInternalFaults
	}
	if num > 0 {
		if cartDO == nil {
			userCartDO := &model.UserCartDO{}
			userCartDO.UserId = userId
			userCartDO.GoodsId = goodsId
			userCartDO.SkuId = skuId
			userCartDO.Num = num
			userCartDO.CreateTime = time.Now()
			userCartDO.UpdateTime = time.Now()
			err = dao.CreateUserCart(ctx, userCartDO)
			if err != nil {
				log.Printf("call UpdateUserCart failed, err:%v", err)
				return errs.ErrorInternalFaults
			}
		} else {
			if skuDO.Stock < cartDO.Num+num {
				return errs.NewErrorGoodsCart("库存不足！")
			}
			if cartDO.Num+num > defs.CartMax {
				cartDO.Num = defs.CartMax
			} else {
				cartDO.Num += num
			}
			err = dao.UpdateUserCart(ctx, cartDO)
			if err != nil {
				log.Printf("call UpdateUserCart failed, err:%v", err)
				return errs.ErrorInternalFaults
			}
		}
	} else {
		if cartDO == nil {
			panic(errs.ErrorGoodsCart)
		}
		if cartDO.Num+num >= 1 {
			cartDO.Num += num
			err = dao.UpdateUserCart(ctx, cartDO)
			if err != nil {
				log.Printf("call UpdateUserCart failed, err:%v", err)
				return errs.ErrorInternalFaults
			}
		}
	}
	return nil
}

// EditCartGoods 编辑购物车
func EditCartGoods(ctx context.Context, userId, cartId, num int) error {
	// 1.查询购物车记录
	record, err := dao.GetUserCartById(ctx, userId, cartId)
	if err != nil {
		if err == dao.NotFoundRecord {
			return errs.ErrorGoodsCart
		}
		log.Printf("call GetUserCartById failed, err:%v", err)
		return errs.ErrorInternalFaults
	}
	// 2.更新数量
	if num == 0 {
		record.Del = defs.DELETE
		err = dao.UpdateUserCart(ctx, record)
		if err != nil {
			log.Printf("call UpdateUserCart failed, err:%v", err)
			return errs.ErrorInternalFaults
		}
	} else {
		return DoEditCart(ctx, userId, record.GoodsId, record.SkuId, num)
	}
	return nil
}

// CountCartGoodsNum 统计购物车商品数量
func CountCartGoodsNum(ctx context.Context, userId int) (int, error) {
	cList, err := dao.GetUserCartList(ctx, userId)
	if err != nil {
		log.Printf("call GetUserCartList failed, err:%v", err)
		return 0, errs.ErrorInternalFaults
	}
	var num int
	for _, v := range cList {
		num += v.Num
	}
	return num, nil
}
