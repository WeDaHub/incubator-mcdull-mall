package service

import (
	"App-CloudBase-mcdull-mall/dao"
	"App-CloudBase-mcdull-mall/defs"
	"App-CloudBase-mcdull-mall/errs"
	"App-CloudBase-mcdull-mall/model"
	"App-CloudBase-mcdull-mall/utils"
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
	"time"

	"github.com/shopspring/decimal"
)

// GenerateOrder 下订单
func GenerateOrder(ctx context.Context, userId, addressId, couponLogId int, dispatchAmount, expectAmount decimal.Decimal,
	cartList []defs.PortalCartGoods) (*defs.PortalPlaceOrderVO, error) {
	goodsAmount := checkCartGoodsAndStock(ctx, userId, cartList)
	discountAmount := calcGoodsDiscountAmount(ctx, goodsAmount, userId, couponLogId)
	if !goodsAmount.Sub(discountAmount).Add(dispatchAmount).Equal(expectAmount) {
		panic(errs.NewErrorOrder("订单金额不符！"))
	}
	addressSnap := getAddressSnapshot(ctx, userId, addressId)
	orderNo := time.Now().Format("20060102150405") + utils.RandomNumberStr(6)
	prepayId := generateWxpayPrepayId(orderNo, expectAmount.String())

	// 聚合订单
	orderDO := &model.OrderDO{}
	orderDO.OrderNo = orderNo
	orderDO.UserId = userId
	orderDO.PayAmount = goodsAmount.Sub(discountAmount).Add(dispatchAmount).String()
	orderDO.GoodsAmount = goodsAmount.String()
	orderDO.DiscountAmount = discountAmount.String()
	orderDO.DispatchAmount = dispatchAmount.String()
	orderDO.PayTime, _ = time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")
	orderDO.DeliverTime, _ = time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")
	orderDO.FinishTime, _ = time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")
	orderDO.Status = 0
	orderDO.AddressId = addressId
	orderDO.AddressSnapshot = addressSnap
	orderDO.WxappPrepayId = prepayId
	orderDO.CreateTime = time.Now()
	orderDO.UpdateTime = time.Now()
	err := dao.AddOrder(ctx, orderDO)
	if err != nil {
		log.Printf("call AddOrder failed, err:%v", err)
		return nil, err
	}
	// WARN 异步任务，可能失败
	go func() {
		// 订单商品快照
		err = orderGoodsSnapshot(ctx, userId, orderNo, cartList)
		if err != nil {
			log.Printf("call orderGoodsSnapshot failed, err:%v", err)
			return
		}
		// 清理购物车
		err = clearUserCart(ctx, userId, cartList)
		if err != nil {
			log.Printf("call clearUserCart failed, err:%v", err)
			return
		}
		// 优惠券核销
		err = couponCannel(ctx, userId, couponLogId)
		if err != nil {
			log.Printf("call couponCannel failed, err:%v", err)
			return
		}
	}()
	orderInfo := &defs.PortalPlaceOrderVO{OrderNo: orderNo, PrepayId: prepayId}
	return orderInfo, nil
}

// 检查-购物车以及商品的库存
func checkCartGoodsAndStock(ctx context.Context, userId int, cartList []defs.PortalCartGoods) decimal.Decimal {
	goodsAmount := decimal.NewFromInt(0)
	for _, v := range cartList {
		if v.CartId != 0 {
			// 检查购物车
			cartDO, err := dao.GetUserCartById(ctx, userId, v.CartId)
			if err != nil {
				if err == dao.NotFoundRecord {
					panic(errs.ErrorGoodsCart)
				}
				log.Printf("call GetUserCartById failed, err:%v", err)
				panic(errs.ErrorInternalFaults)
			}
			if cartDO.Del == defs.DELETE {
				panic(errs.ErrorGoodsCart)
			}
		}
		// 检查商品
		goodsDO, err := dao.GetGoodsById(ctx, v.GoodsId)
		if err != nil {
			if err == dao.NotFoundRecord {
				panic(errs.ErrorGoods)
			}
			log.Printf("call GetGoodsById failed, err:%v", err)
			panic(errs.ErrorGoodsCart)
		}
		if goodsDO.Del == defs.DELETE || goodsDO.Online == defs.OFFLINE {
			panic(errs.NewErrorOrder("商品下架，无法售出"))
		}
		skuDO, err := dao.GetSkuListBySkuId(ctx, v.SkuId)
		if err != nil {
			if err == dao.NotFoundRecord {
				panic(errs.ErrorSKU)
			}
			log.Printf("call GetSkuListBySkuId failed, err:%v", err)
			panic(errs.ErrorInternalFaults)
		}
		if skuDO.Del == defs.DELETE || skuDO.Online == defs.OFFLINE {
			panic(errs.NewErrorOrder("商品下架，无法售出"))
		}
		if skuDO.Stock < v.Num {
			panic(errs.NewErrorOrder("商品库存不足！"))
		}
		price, err := decimal.NewFromString(skuDO.Price)
		if err != nil {
			panic(errs.ErrorInternalFaults)
		}
		num := decimal.NewFromInt(int64(v.Num))
		goodsAmount = goodsAmount.Add(price.Mul(num))
	}
	return goodsAmount
}

// 计算优惠金额
func calcGoodsDiscountAmount(ctx context.Context, goodsAmount decimal.Decimal, userId, couponLogId int) decimal.Decimal {
	// 1.优惠券记录
	if couponLogId == 0 {
		return decimal.NewFromInt(0)
	}
	couponLog, err := dao.GetUserCouponLogById(ctx, userId, couponLogId)
	if err != nil {
		if err == dao.NotFoundRecord {
			panic(errs.ErrorCoupon)
		}
		log.Printf("call GetUserCouponLogById failed, err:%v", err)
		panic(errs.ErrorInternalFaults)
	}
	if couponLog.Del == defs.DELETE || couponLog.Status != 0 || couponLog.UserId != userId {
		panic(errs.NewErrorCoupon("无效的优惠券！"))
	}
	// 2.优惠券信息
	coupon, err := dao.GetCouponById(ctx, couponLog.CouponId)
	if err != nil {
		if err == dao.NotFoundRecord {
			panic(errs.ErrorCoupon)
		}
		log.Printf("call GetCouponById failed, err:%v", err)
		panic(errs.ErrorInternalFaults)
	}
	var discountAmount decimal.Decimal
	switch coupon.Type {
	case 1:
		fullMoney, err := decimal.NewFromString(coupon.FullMoney)
		if err != nil {
			panic(err)
		}
		if goodsAmount.LessThan(fullMoney) {
			panic(errs.NewErrorCoupon("未达到满减要求！"))
		}
		minus, err := decimal.NewFromString(coupon.Minus)
		if err != nil {
			panic(err)
		}
		discountAmount = minus
	case 2:
		rate, err := decimal.NewFromString(coupon.Rate)
		if err != nil {
			panic(err)
		}
		discountAmount = goodsAmount.Sub(goodsAmount.Mul(rate).Round(2))
	case 3:
		minus, err := decimal.NewFromString(coupon.Minus)
		if err != nil {
			panic(err)
		}
		discountAmount = minus
	case 4:
		fullMoney, err := decimal.NewFromString(coupon.FullMoney)
		if err != nil {
			panic(err)
		}
		if goodsAmount.LessThan(fullMoney) {
			panic(errs.NewErrorCoupon("未达到满减要求！"))
		}
		rate, err := decimal.NewFromString(coupon.Rate)
		if err != nil {
			panic(err)
		}
		discountAmount = goodsAmount.Sub(goodsAmount.Mul(rate).Round(2))
	default:
		discountAmount = decimal.NewFromInt(0)
	}
	if discountAmount.GreaterThan(goodsAmount) {
		discountAmount = goodsAmount
	}
	return discountAmount
}

// 收货地址快照（base64编码，支持云数据库存储）
func getAddressSnapshot(ctx context.Context, userId, addressId int) string {
	addressDO, err := dao.GetAddressById(ctx, addressId, userId)
	if err != nil {
		if err == dao.NotFoundRecord {
			panic(errs.ErrorAddress)
		}
		panic(errs.ErrorInternalFaults)
	}
	if addressDO.Del == defs.DELETE {
		panic(errs.ErrorAddress)
	}
	snapshot := defs.AddressSnapshot{}
	snapshot.Contacts = addressDO.Contacts
	snapshot.Mobile = addressDO.Mobile
	snapshot.ProvinceId = addressDO.ProvinceId
	snapshot.ProvinceStr = addressDO.ProvinceStr
	snapshot.CityId = addressDO.CityId
	snapshot.CityStr = addressDO.CityStr
	snapshot.AreaStr = addressDO.AreaStr
	snapshot.Address = addressDO.Address
	bytes, err := json.Marshal(snapshot)
	if err != nil {
		panic(errs.ErrorInternalFaults)
	}
	return base64.StdEncoding.EncodeToString(bytes)
}

// 支付-获取微信预支付ID
func generateWxpayPrepayId(orderNo string, payAmount string) string {
	// todo: 请求微信支付订单
	return "prepay_id:" + orderNo
}

// 订单详情-快照
func orderGoodsSnapshot(ctx context.Context, userId int, orderNo string, goodsList []defs.PortalCartGoods) error {
	for _, v := range goodsList {
		goodsDO, err := dao.GetGoodsById(ctx, v.GoodsId)
		if err != nil {
			log.Printf("call GetGoodsById failed, err:%v", err)
			return err
		}
		skuDO, err := dao.GetSkuListBySkuId(ctx, v.SkuId)
		if err != nil {
			log.Printf("call GetSkuListBySkuId failed, err:%v", err)
			return err
		}
		// base64编码，兼容云数据库
		encodeSpecs := base64.StdEncoding.EncodeToString([]byte(skuDO.Specs))
		orderGoodsDO := &model.OrderGoodsDO{}
		orderGoodsDO.OrderNo = orderNo
		orderGoodsDO.UserId = userId
		orderGoodsDO.GoodsId = v.GoodsId
		orderGoodsDO.SkuId = v.SkuId
		orderGoodsDO.Picture = skuDO.Picture
		orderGoodsDO.Title = goodsDO.Title
		orderGoodsDO.Price = skuDO.Price
		orderGoodsDO.Specs = encodeSpecs
		orderGoodsDO.Num = v.Num
		orderGoodsDO.LockStatus = 0
		orderGoodsDO.CreateTime = time.Now()
		orderGoodsDO.UpdateTime, _ = time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")
		err = dao.AddOrderGoods(ctx, orderGoodsDO)
		if err != nil {
			log.Printf("call AddOrderGoods failed, err:%v", err)
			return err
		}
		// SKU减库存
		sku, err := dao.GetSkuListBySkuId(ctx, v.SkuId)
		if err != nil {
			log.Printf("call GetSkuListBySkuId failed, err:%v", err)
			return err
		}
		err = dao.UpdateSkuStockById(ctx, v.SkuId, sku.Stock-v.Num)
		if err != nil {
			log.Printf("call UpdateSkuStockById failed, err:%v", err)
			return err
		}
		// 更新商品销量
		goods, err := dao.GetGoodsById(ctx, v.GoodsId)
		if err != nil {
			log.Printf("call GetGoodsById failed, err:%v", err)
			return err
		}
		err = dao.UpdateGoodsSaleNum(ctx, v.GoodsId, goods.SaleNum+v.Num)
		if err != nil {
			log.Printf("call UpdateGoodsSaleNum failed, err:%v", err)
			return err
		}
	}
	return nil
}

// 下单成功-清理购物车
func clearUserCart(ctx context.Context, userId int, goodsList []defs.PortalCartGoods) error {
	for _, v := range goodsList {
		if v.CartId != 0 {
			cartDO, err := dao.GetUserCartById(ctx, userId, v.CartId)
			if err != nil {
				log.Printf("call GetUserCartById failed, err:%v", err)
				continue
			}
			if cartDO.Del == defs.DELETE {
				continue
			}
			cartDO.Del = defs.DELETE
			err = dao.UpdateUserCart(ctx, cartDO)
			if err != nil {
				log.Printf("call UpdateUserCart failed, err:%v", err)
				return err
			}
		}
	}
	return nil
}

// couponCannel 优惠券-核销
func couponCannel(ctx context.Context, userId, couponLogId int) error {
	if couponLogId == 0 {
		return nil
	}
	couponLogDO, err := dao.GetUserCouponLogById(ctx, userId, couponLogId)
	if err != nil {
		log.Printf("call GetUserCouponLogById failed, err:%v", err)
		return err
	}
	couponLogDO.Status = 1
	couponLogDO.UseTime = time.Now()
	err = dao.UpdateCouponLog(ctx, couponLogDO)
	if err != nil {
		log.Printf("call UpdateCouponLog failed, err:%v", err)
		return err
	}
	return nil
}
