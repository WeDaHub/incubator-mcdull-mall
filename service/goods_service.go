package service

import (
	"App-CloudBase-mcdull-mall/dao"
	"App-CloudBase-mcdull-mall/defs"
	"App-CloudBase-mcdull-mall/env"
	"App-CloudBase-mcdull-mall/model"
	"context"
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"sync"
)

// QueryGoodsList 查询商品列表（匹配和排序）
func QueryGoodsList(ctx context.Context, keyword string, sort, categoryId, page, size int) ([]*model.GoodsDO, int, error) {
	// 1.查询商品缓存
	goodsList, err := dao.GetGoodsListCache(ctx)
	if err != nil {
		log.Printf("call GetGoodsListCache failed, err:%v", err)
		return nil, 0, err
	}
	// 2.过滤
	newList := make([]*model.GoodsDO, 0)
	for _, v := range goodsList {
		// 过滤删除态和下架
		if v.Del == defs.DELETE || v.Online == defs.OFFLINE {
			continue
		}
		if categoryId != 0 {
			if categoryId != v.CategoryId {
				continue
			}
		}
		if keyword != "" {
			if strings.Index(strings.ToLower(v.Title), strings.ToLower(keyword)) == -1 {
				continue
			}
		}
		newList = append(newList, v)
	}
	// 3.排序：0-综合 1-新品 2-销量 3-价格 todo:
	//var order string
	//switch sort {
	//case 1:
	//	order = "create_time"
	//case 2:
	//	order = "sale_num"
	//case 3:
	//	order = "price"
	//default:
	//	order = ""
	//}
	// 4.分页
	total := len(newList)
	offset, endpos := pagePos(total, page, size)
	return newList[offset:endpos], total, nil
}

// BatchCountBuyUserNum 批量查询用户数量
func BatchCountBuyUserNum(ctx context.Context, goodsIds []int) (*sync.Map, error) {
	var numMap sync.Map
	var wg sync.WaitGroup
	wg.Add(len(goodsIds))
	for _, v := range goodsIds {
		go func(goodsId int) {
			defer wg.Done()
			num, err := dao.CountBuyUserNum(ctx, goodsId)
			if err != nil {
				log.Printf("call CountBuyUserNumCache failed, err:%v", err)
				return
			}
			numMap.Store(goodsId, num)
		}(v)
	}
	wg.Wait()
	return &numMap, nil
}

// QueryPortalGoodsDetail 查询商品详情
func QueryPortalGoodsDetail(ctx context.Context, goodsId int) (*defs.PortalGoodsInfo, error) {
	// 1.查询商品
	goodsDO, err := dao.GetGoodsById(ctx, goodsId)
	if err != nil {
		log.Printf("call GetGoodsById failed, err:%v", err)
		return nil, err
	}
	if goodsDO.Del == defs.DELETE || goodsDO.Online == defs.OFFLINE {
		return nil, dao.NotFoundRecord
	}
	// 2.查询商品SKU
	skuDOList, err := dao.GetOnlineSkuListByGoodsId(ctx, goodsId)
	if err != nil {
		log.Printf("call GetSkuListByGoodsId failed, err:%v", err)
		return nil, err
	}
	domain := env.LoadConf().Domain
	skuList := extractSkuVOList(domain, skuDOList)
	specList, err := extraceSpecVOList(ctx, goodsId, skuDOList)
	if err != nil {
		log.Printf("call extraceSpecVOList failed, err:%v", err)
		return nil, err
	}
	// 3.渲染数据
	goodsVO := &defs.PortalGoodsInfo{}
	goodsVO.Id = goodsDO.Id
	goodsVO.Title = goodsDO.Title
	goodsVO.Price, _ = strconv.ParseFloat(goodsDO.Price, 2)
	goodsVO.Picture = domain + goodsDO.Picture
	goodsVO.BannerPicture = combinationPictureUrl(domain, goodsDO.BannerPicture)
	goodsVO.DetailPicture = combinationPictureUrl(domain, goodsDO.DetailPicture)
	goodsVO.Tags = goodsDO.Tags
	goodsVO.Description = goodsDO.Description
	goodsVO.SkuList = skuList
	goodsVO.SpecList = specList
	return goodsVO, nil
}

// 组合本地资源图片路径
func combinationPictureUrl(domain, arrstr string) string {
	temArr := make([]string, 0)
	err := json.Unmarshal([]byte(arrstr), &temArr)
	if err != nil {
		return arrstr
	}
	newArr := make([]string, 0)
	for _, v := range temArr {
		newArr = append(newArr, domain+v)
	}
	newArrBytes, err := json.Marshal(newArr)
	if err != nil {
		return arrstr
	}
	return string(newArrBytes)
}

func extractSkuVOList(domain string, skuDOList []*model.SkuDO) []defs.PortalSkuVO {
	skuList := make([]defs.PortalSkuVO, 0)
	for _, v := range skuDOList {
		skuVO := defs.PortalSkuVO{}
		skuVO.Id = v.Id
		skuVO.Picture = domain + v.Picture
		skuVO.Title = v.Title
		skuVO.Price, _ = strconv.ParseFloat(v.Price, 2)
		skuVO.Code = v.Code
		skuVO.Stock = v.Stock
		skuVO.Specs = v.Specs
		skuList = append(skuList, skuVO)
	}
	return skuList
}

func extraceSpecVOList(ctx context.Context, goodsId int, skuDOList []*model.SkuDO) ([]defs.PortalSpecVO, error) {
	specVOMap, specAttrVOMap := extraceSpecAttrVOList(skuDOList)
	specList, err := dao.GetGoodsSpecByGoodsId(ctx, goodsId)
	if err != nil {
		log.Printf("call GetGoodsSpecByGoodsId failed, err:%v", err)
		return nil, err
	}
	specVOList := make([]defs.PortalSpecVO, 0)
	for _, v := range specList {
		specId := v.SpecId
		if specVOMap[specId] == "" {
			continue
		}
		specVO := defs.PortalSpecVO{}
		specVO.SpecId = specId
		specVO.Name = specVOMap[specId]
		specVO.AttrList = specAttrVOMap[specId]
		specVOList = append(specVOList, specVO)
	}
	return specVOList, nil
}

func extraceSpecAttrVOList(skuDOList []*model.SkuDO) (map[int]string, map[int][]defs.PortalSpecAttrVO) {
	specVOMap := map[int]string{}
	specAttrVOMap := map[int][]defs.PortalSpecAttrVO{}
	for _, v := range skuDOList {
		// [{"key": "颜色", "value": "青芒色", "keyId": 1, "valueId": 42}, {"key": "尺寸", "value": "7英寸", "keyId": 2, "valueId": 5}]
		specs := make([]defs.SkuSpecs, 0)
		err := json.Unmarshal([]byte(v.Specs), &specs)
		if err != nil {
			panic(err)
		}
		for _, item := range specs {
			specName := specVOMap[item.KeyId]
			if specName == "" {
				specVOMap[item.KeyId] = item.Key
			}
			attrVOList := specAttrVOMap[item.KeyId]
			if attrVOList == nil {
				attrVOList = []defs.PortalSpecAttrVO{}
				specAttrVOMap[item.KeyId] = attrVOList
			}
			flag := false
			for _, attrVO := range attrVOList {
				if attrVO.AttrId == item.ValueId {
					flag = true
					break
				}
			}
			if flag {
				continue
			}
			attrVO := defs.PortalSpecAttrVO{}
			attrVO.AttrId = item.ValueId
			attrVO.Value = item.Value
			attrVOList = append(attrVOList, attrVO)
			specAttrVOMap[item.KeyId] = attrVOList
		}
	}
	return specVOMap, specAttrVOMap
}
