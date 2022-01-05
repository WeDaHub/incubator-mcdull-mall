package web

import (
	"App-CloudBase-mcdull-mall/handler"
	"context"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func NewRouter(ctx context.Context) *mux.Router {
	router := mux.NewRouter()
	h := handler.NewHandler(ctx)
	registerHandler(router, h)
	return router
}

// 注册路由
func registerHandler(router *mux.Router, h *handler.Handler) {
	mw := &Middleware{}
	chain := alice.New(mw.LoggingHandler, mw.RecoverPanic, mw.CORSHandler, mw.ValidateAuthToken)
	router.Handle("/api/wxapp/login", chain.ThenFunc(h.Login)).Methods("GET").Queries("code", "{code}")
	router.Handle("/api/wxapp/user-info", chain.ThenFunc(h.UserInfo)).Methods("GET")
	router.Handle("/api/wxapp/auth-phone", chain.ThenFunc(h.AuthPhone)).Methods("POST")
	router.Handle("/api/wxapp/auth-info", chain.ThenFunc(h.AuthUserInfo)).Methods("POST")
	router.Handle("/api/home/banner", chain.ThenFunc(h.HomeBanner)).Methods("GET").Queries("page", "{page}").Queries("size", "{size}")
	router.Handle("/api/home/grid", chain.ThenFunc(h.GetGridCategoryList)).Methods("GET").Queries("page", "{page}").Queries("size", "{size}")
	router.Handle("/api/category/list", chain.ThenFunc(h.GetSubCategoryList)).Methods("GET")
	router.Handle("/api/goods/list", chain.ThenFunc(h.GetGoodsList)).Methods("GET").Queries("k", "{k}").Queries("s", "{s}").Queries("c", "{c}").Queries("page", "{page}").Queries("size", "{size}")
	router.Handle("/api/goods/detail", chain.ThenFunc(h.GetGoodsDetail)).Methods("GET").Queries("id", "{id}")
	router.Handle("/api/cart/list", chain.ThenFunc(h.GetCartGoodsList)).Methods("GET").Queries("page", "{page}").Queries("size", "{size}")
	router.Handle("/api/cart/add", chain.ThenFunc(h.AddCartGoods)).Methods("POST")
	router.Handle("/api/cart/edit", chain.ThenFunc(h.EditCartGoods)).Methods("POST")
	router.Handle("/api/cart/goods_num", chain.ThenFunc(h.GetCartGoodsNum)).Methods("GET")
	router.Handle("/api/coupon/list", chain.ThenFunc(h.GetCouponList)).Methods("GET").Queries("page", "{page}").Queries("size", "{size}")
	router.Handle("/api/coupon/take", chain.ThenFunc(h.TakeCoupon)).Methods("POST")
	router.Handle("/api/user/coupon/list", chain.ThenFunc(h.GetUserCouponList)).Methods("GET").Queries("status", "{status}").Queries("page", "{page}").Queries("size", "{size}")
	router.Handle("/api/user/coupon", chain.ThenFunc(h.DoDeleteCouponLog)).Methods("DELETE").Queries("id", "{id}")
	router.Handle("/api/user/address/list", chain.ThenFunc(h.GetAddressList)).Methods("GET").Queries("page", "{page}").Queries("size", "{size}")
	router.Handle("/api/user/address/edit", chain.ThenFunc(h.EditAddress)).Methods("POST")
	router.Handle("/api/user/address", chain.ThenFunc(h.GetAddress)).Methods("GET").Queries("id", "{id}")
	router.Handle("/api/user/address", chain.ThenFunc(h.DoDeleteAddress)).Methods("DELETE").Queries("id", "{id}")
	router.Handle("/api/user/default_address", chain.ThenFunc(h.GetDefaultAddress)).Methods("GET")
	router.Handle("/api/placeorder", chain.ThenFunc(h.PlaceOrder)).Methods("POST")
	router.Handle("/api/order/list", chain.ThenFunc(h.GetOrderList)).Methods("GET").Queries("status", "{status}").Queries("page", "{page}").Queries("size", "{size}")
	router.Handle("/api/order/detail", chain.ThenFunc(h.GetOrderDetail)).Methods("GET").Queries("orderNo", "{orderNo}")
	router.Handle("/api/order/cancel", chain.ThenFunc(h.CancelOrder)).Methods("PUT").Queries("id", "{id}")
	router.Handle("/api/order", chain.ThenFunc(h.DeleteOrder)).Methods("DELETE").Queries("id", "{id}")
	router.Handle("/api/order/confirm_goods", chain.ThenFunc(h.ConfirmTakeGoods)).Methods("PUT").Queries("id", "{id}")
	router.Handle("/api/order/refund_apply", chain.ThenFunc(h.RefundApply)).Methods("PUT")
	router.Handle("/api/order/refund_detail", chain.ThenFunc(h.RefundDetail)).Methods("GET").Queries("refundNo", "{refundNo}")
	router.Handle("/api/order/refund_undo", chain.ThenFunc(h.UndoRefundApply)).Methods("PUT").Queries("refundNo", "{refundNo}")
	router.Handle("/api/order/remind", chain.ThenFunc(h.GetOrderRemind)).Methods("GET")
	router.Handle("/api/browse/list", chain.ThenFunc(h.UserBrowseHistory)).Methods("GET").Queries("page", "{page}").Queries("size", "{size}")
	router.Handle("/api/browse/clear", chain.ThenFunc(h.ClearBrowseHistory)).Methods("POST")
	// 静态资源
	path, _ := os.Getwd()
	router.PathPrefix("/assets/images/").Handler(http.StripPrefix("/assets/images/", http.FileServer(http.Dir(path+"/assets/images"))))
}
