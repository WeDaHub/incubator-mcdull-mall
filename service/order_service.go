package service

import (
	"App-CloudBase-mcdull-mall/env"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"time"

	"App-CloudBase-mcdull-mall/dao"
	"App-CloudBase-mcdull-mall/defs"
	"App-CloudBase-mcdull-mall/model"
	"App-CloudBase-mcdull-mall/utils"
)

var NotFoundOrderError = errors.New("not found order error")
var CancelOrderError = errors.New("cancel order error")
var OrderStatusError = errors.New("order status error")
var NotFoundRefundError = errors.New("not found refund error")

// QueryOrderList 查询订单列表
func QueryOrderList(ctx context.Context, userId, status, page, size int) ([]*model.OrderDO, int, error) {
	cList, err := dao.ListOrderByParams(ctx, userId, status)
	if err != nil {
		log.Printf("call ListOrderByParams failed, err:%v", err)
		return nil, 0, err
	}
	total := len(cList)
	offset, endpos := pagePos(total, page, size)
	return cList[offset:endpos], total, nil
}

// ExtraceOrderGoods 提取订单商品
func ExtraceOrderGoods(ctx context.Context, orderNo string) ([]defs.PortalOrderGoodsVO, int, error) {
	goodsList, err := dao.QueryOrderGoods(ctx, orderNo)
	if err != nil {
		log.Printf("call QueryOrderGoods failed, err:%v", err)
		return nil, 0, err
	}
	goodsNum := 0
	goodsVOList := make([]defs.PortalOrderGoodsVO, 0)
	for _, v := range goodsList {
		specList := make([]*defs.SkuSpecs, 0)
		// base64解码
		specBytes, err := base64.StdEncoding.DecodeString(v.Specs)
		if err != nil {
			log.Printf("call base64 DecodeString failed, err:%v", err)
			return nil, 0, err
		}
		err = json.Unmarshal(specBytes, &specList)
		if err != nil {
			log.Printf("call json Unmarshal failed, err:%v", err)
			return nil, 0, err
		}
		specs := ""
		for _, v := range specList {
			specs += v.Value + "; "
		}
		if len(specs) > 2 {
			specs = specs[0 : len(specs)-2]
		}
		domain := env.LoadConf().Domain
		goodsVO := defs.PortalOrderGoodsVO{}
		goodsVO.GoodsId = v.GoodsId
		goodsVO.Title = v.Title
		goodsVO.Price, _ = strconv.ParseFloat(v.Price, 2)
		goodsVO.Picture = domain + v.Picture
		goodsVO.SkuId = v.SkuId
		goodsVO.Specs = specs
		goodsVO.Num = v.Num
		goodsVOList = append(goodsVOList, goodsVO)
		goodsNum += v.Num
	}
	return goodsVOList, goodsNum, nil
}

// QueryOrderDetail 查询订单详情
func QueryOrderDetail(ctx context.Context, userId int, orderNo string) (*defs.PortalOrderDetailVO, error) {
	// 1.查询订单信息
	orderDO, err := dao.QueryOrderByOrderNo(ctx, orderNo)
	if err != nil {
		if err == dao.NotFoundRecord {
			return nil, NotFoundOrderError
		}
		log.Printf("call QueryOrderByOrderNo failed, err:%v", err)
		return nil, err
	}
	if orderDO.Del == defs.DELETE || orderDO.UserId != userId {
		return nil, NotFoundOrderError
	}
	// 2.收货地址
	address := defs.AddressSnapshot{}
	addressSnapshot, err := base64.StdEncoding.DecodeString(orderDO.AddressSnapshot)
	if err != nil {
		log.Printf("call base64 DecodeString failed, err:%v", err)
		return nil, err
	}
	err = json.Unmarshal(addressSnapshot, &address)
	if err != nil {
		log.Printf("call json Unmarshal failed, err:%v", err)
		return nil, err
	}
	// 3.订单商品
	goodsList, goodsNum, err := ExtraceOrderGoods(ctx, orderDO.OrderNo)
	if err != nil {
		log.Printf("call ExtraceOrderGoods failed, err:%v", err)
		return nil, err
	}
	// 4.退款信息
	refundDO, err := dao.QueryOrderRefundRecord(ctx, orderNo)
	if err != nil {
		if err != dao.NotFoundRecord {
			log.Printf("call QueryOrderRefundRecord failed, err:%v", err)
			return nil, err
		}
	}
	var refundNo string
	if refundDO != nil {
		refundNo = refundDO.RefundNo
	}
	// 5.渲染数据
	orderVO := &defs.PortalOrderDetailVO{}
	orderVO.Id = orderDO.Id
	orderVO.OrderNo = orderDO.OrderNo
	orderVO.GoodsAmount, _ = strconv.ParseFloat(orderDO.GoodsAmount, 2)
	orderVO.DiscountAmount, _ = strconv.ParseFloat(orderDO.DiscountAmount, 2)
	orderVO.DispatchAmount, _ = strconv.ParseFloat(orderDO.DispatchAmount, 2)
	orderVO.PayAmount, _ = strconv.ParseFloat(orderDO.PayAmount, 2)
	orderVO.Status = orderDO.Status
	orderVO.PlaceTime = orderDO.CreateTime.Format("2006-01-02 15:04:05")
	orderVO.PayTime = orderDO.PayTime.Format("2006-01-02 15:04:05")
	orderVO.DeliverTime = orderDO.DeliverTime.Format("2006-01-02 15:04:05")
	orderVO.FinishTime = orderDO.FinishTime.Format("2006-01-02 15:04:05")
	orderVO.GoodsList = goodsList
	orderVO.GoodsNum = goodsNum
	orderVO.Address = address
	orderVO.RefundApply = defs.OrderRefundApplyVO{
		RefundNo: refundNo,
	}
	return orderVO, nil
}

// CancelOrder 订单-取消订单
func CancelOrder(ctx context.Context, userId, orderId int) error {
	// 1.查询订单信息
	orderDO, err := dao.QueryOrderById(ctx, orderId)
	if err != nil {
		if err == dao.NotFoundRecord {
			return NotFoundOrderError
		}
		log.Printf("call QueryOrderById failed, err:%v", err)
		return err
	}
	if orderDO.Del == defs.DELETE || orderDO.UserId != userId {
		return NotFoundOrderError
	}
	// 2.检查订单状态
	if orderDO.Status != 0 {
		return CancelOrderError
	}
	orderDO.Status = -1
	orderDO.FinishTime = time.Now()
	// 3.更新订单状态
	err = dao.UpdateOrderFinishStauts(ctx, orderDO)
	if err != nil {
		log.Printf("call UpdateOrderById failed, err:%v", err)
		return err
	}
	// 库存回退（异步）
	go orderStockRollback(ctx, orderDO.OrderNo)
	return nil
}

// DeleteOrderRecord 订单-删除记录
func DeleteOrderRecord(ctx context.Context, userId, orderId int) error {
	// 1.查询订单信息
	orderDO, err := dao.QueryOrderById(ctx, orderId)
	if err != nil {
		if err == dao.NotFoundRecord {
			return NotFoundOrderError
		}
		return err
	}
	if orderDO.Del == defs.DELETE || orderDO.UserId != userId {
		return NotFoundOrderError
	}
	// 2.检查状态
	if orderDO.Status == -1 || orderDO.Status == 3 || orderDO.Status == 5 {
		orderDO.Del = defs.DELETE
		err := dao.DeleteOrderRecord(ctx, orderId)
		if err != nil {
			return err
		}
	} else {
		return CancelOrderError
	}
	return nil
}

// ConfirmTakeGoods 订单-确认收货
func ConfirmTakeGoods(ctx context.Context, userId, orderId int) error {
	// 1.查询订单信息
	orderDO, err := dao.QueryOrderById(ctx, orderId)
	if err != nil {
		if err == dao.NotFoundRecord {
			return NotFoundOrderError
		}
		return err
	}
	if orderDO.Del == defs.DELETE || orderDO.UserId != userId {
		return NotFoundOrderError
	}
	// 2.更新状态
	if orderDO.Status == 2 {
		orderDO.Status = 3
		orderDO.FinishTime = time.Now()
		err = dao.UpdateOrderFinishStauts(ctx, orderDO)
		if err != nil {
			log.Printf("call UpdateOrderFinishStauts failed, err:%v", err)
			return err
		}
	}
	return nil
}

// OrderPaySuccessNotify 订单支付成功通知
func OrderPaySuccessNotify(ctx context.Context, orderNo string) error {
	// 1.查询订单信息
	orderDO, err := dao.QueryOrderByOrderNo(ctx, orderNo)
	if err != nil {
		if err == dao.NotFoundRecord {
			return NotFoundOrderError
		}
		return err
	}
	// 2.幂等性
	if orderDO.Status != 0 {
		log.Printf("orderNo = %v 重复回调", orderDO)
		return errors.New("break repeated request")
	}
	// 3.更新订单
	orderDO.Status = 1
	orderDO.PayTime = time.Now()
	err = dao.UpdateOrderPayStauts(ctx, orderDO)
	if err != nil {
		log.Printf("call UpdateOrderFinishStauts failed, err:%v", err)
		return err
	}
	return nil
}

// RefundApply 退款申请
func RefundApply(ctx context.Context, userId int, orderNo, reason string) (string, error) {
	// 1.查询订单信息
	orderDO, err := dao.QueryOrderByOrderNo(ctx, orderNo)
	if err != nil {
		if err == dao.NotFoundRecord {
			return "", NotFoundOrderError
		}
		return "", err
	}
	if orderDO.Del == defs.DELETE || orderDO.UserId != userId {
		return "", NotFoundOrderError
	}
	// 2.检查状态
	if orderDO.Status != 1 {
		return "", OrderStatusError
	}
	// 3.新增退款记录
	refundNo := time.Now().Format("20060102150405") + utils.RandomNumberStr(6)
	refund := &model.OrderRefund{}
	refund.RefundNo = refundNo
	refund.UserId = userId
	refund.OrderNo = orderNo
	refund.Reason = reason
	refund.RefundTime, _ = time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")
	refund.RefundAmount = orderDO.PayAmount
	refund.CreateTime = time.Now()
	refund.UpdateTime, _ = time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")
	autoId, err := dao.AddRefundRecord(ctx, refund)
	if err != nil {
		log.Printf("call AddRefundRecord failed, err:%v", err)
		return "", err
	}
	// 4.更新订单状态
	err = dao.UpdateOrderStatus(ctx, autoId, 4)
	if err != nil {
		log.Printf("call UpdateOrderStatus failed, err:%v", err)
		return "", err
	}
	return refundNo, nil
}

// QueryRefundDetail 查询退款详情
func QueryRefundDetail(ctx context.Context, userId int, refundNo string) (*defs.OrderRefundDetailVO, error) {
	// 1.查询退款记录
	refundDO, err := dao.QueryRefundRecord(ctx, refundNo)
	if err != nil {
		if err == dao.NotFoundRecord {
			return nil, NotFoundRefundError
		}
		return nil, err
	}
	if refundDO.Del == defs.DELETE || refundDO.UserId != userId {
		return nil, NotFoundRefundError
	}
	// 2.查询订单商品
	goodsList, _, err := ExtraceOrderGoods(ctx, refundDO.OrderNo)
	if err != nil {
		log.Printf("call ExtraceOrderGoods failed, err:%v", err)
		return nil, err
	}
	// 3.渲染数据
	refundVO := &defs.OrderRefundDetailVO{}
	refundVO.RefundNo = refundDO.RefundNo
	refundVO.Reason = refundDO.Reason
	refundVO.RefundAmount, _ = strconv.ParseFloat(refundDO.RefundAmount, 2)
	refundVO.Status = refundDO.Status
	refundVO.ApplyTime = refundDO.CreateTime.Format("2006-01-02 15:04:05")
	refundVO.RefundTime = refundDO.RefundTime.Format("2006-01-02 15:04:05")
	refundVO.GoodsList = goodsList
	return refundVO, nil
}

// CountOrderNum 统计-订单数量
func CountOrderNum(ctx context.Context, userId, status int) int {
	orderNum, err := dao.CountOrderNum(ctx, userId, status)
	if err != nil {
		log.Printf("call CountOrderNum failed, err:%v", err)
		return 0
	}
	return orderNum
}

// UndoRefundApply 取消退款申请
func UndoRefundApply(ctx context.Context, userId int, refundNo string) error {
	// 1.查询退款记录
	refundDO, err := dao.QueryRefundRecord(ctx, refundNo)
	if err != nil {
		if err == dao.NotFoundRecord {
			return NotFoundRefundError
		}
		return err
	}
	if refundDO.Del == defs.DELETE || refundDO.UserId != userId {
		return NotFoundRefundError
	}
	// 2.检查退款状态
	if refundDO.Status != 0 {
		return OrderStatusError
	}
	// 3.查询订单详情
	orderDO, err := dao.QueryOrderByOrderNo(ctx, refundDO.OrderNo)
	if err != nil {
		log.Printf("call QueryOrderByOrderNo failed, err:%v", err)
		return err
	}
	// 4.更新订单状态
	err = dao.UpdateOrderStatus(ctx, orderDO.Id, 1)
	if err != nil {
		log.Printf("call UpdateOrderStatus failed, err:%v", err)
		return err
	}
	// 5.更新退款状态
	err = dao.UpdateRefundApply(ctx, refundDO.Id, 2)
	if err != nil {
		log.Printf("call UpdateRefundApply failed, err:%v", err)
		return err
	}
	return nil
}

// 订单-库存回滚
// 场景：取消订单（手动取消、超时未支付）、订单退款
func orderStockRollback(ctx context.Context, orderNo string) {
	// 1.查询订单商品
	orderGoods, err := dao.QueryOrderGoods(ctx, orderNo)
	if err != nil {
		log.Printf("call QueryOrderGoods failed, err:%v", err)
		return
	}
	for _, v := range orderGoods {
		// 2.更新Sku库存
		sku, err := dao.GetSkuListBySkuId(ctx, v.SkuId)
		if err != nil {
			log.Printf("call GetSkuListBySkuId failed, err:%v", err)
			return
		}
		err = dao.UpdateSkuStockById(ctx, v.SkuId, sku.Stock+v.Num)
		if err != nil {
			log.Printf("call UpdateSkuStockById failed, err:%v", err)
			return
		}
		// 3.更新商品销量
		goods, err := dao.GetGoodsById(ctx, v.GoodsId)
		if err != nil {
			log.Printf("call GetGoodsById faild, err:%v", err)
			return
		}
		err = dao.UpdateGoodsSaleNum(ctx, v.GoodsId, goods.SaleNum-v.Num)
		if err != nil {
			log.Printf("call UpdateGoodsSaleNum failed, err:%v", err)
			return
		}
	}
}
