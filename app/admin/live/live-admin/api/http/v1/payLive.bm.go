// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: payLive.proto

/*
Package v1 is a generated blademaster stub package.
This code was generated with go-common/app/tool/bmgen/protoc-gen-bm v0.1.

It is generated from these files:
	payLive.proto
*/
package v1

import (
	"context"

	bm "go-common/library/net/http/blademaster"
	"go-common/library/net/http/blademaster/binding"
)

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathPayLiveAdd = "/live.liveadmin.v1.PayLive/add"
var PathPayLiveUpdate = "/live.liveadmin.v1.PayLive/update"
var PathPayLiveGetList = "/live.liveadmin.v1.PayLive/getList"
var PathPayLiveClose = "/live.liveadmin.v1.PayLive/close"
var PathPayLiveOpen = "/live.liveadmin.v1.PayLive/open"

// =================
// PayLive Interface
// =================

type PayLiveBMServer interface {
	// `method:"POST" internal:"true"` 生成付费直播信息
	Add(ctx context.Context, req *PayLiveAddReq) (resp *PayLiveAddResp, err error)

	// `method:"POST" internal:"true"` 更新付费直播信息
	Update(ctx context.Context, req *PayLiveUpdateReq) (resp *PayLiveUpdateResp, err error)

	// `internal:"true"` 获取付费直播列表
	GetList(ctx context.Context, req *PayLiveGetListReq) (resp *PayLiveGetListResp, err error)

	// `method:"POST" internal:"true"` 关闭鉴权
	Close(ctx context.Context, req *PayLiveCloseReq) (resp *PayLiveCloseResp, err error)

	// `method:"POST" internal:"true"` 开启鉴权
	Open(ctx context.Context, req *PayLiveOpenReq) (resp *PayLiveOpenResp, err error)
}

var v1PayLiveSvc PayLiveBMServer

func payLiveAdd(c *bm.Context) {
	p := new(PayLiveAddReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1PayLiveSvc.Add(c, p)
	c.JSON(resp, err)
}

func payLiveUpdate(c *bm.Context) {
	p := new(PayLiveUpdateReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1PayLiveSvc.Update(c, p)
	c.JSON(resp, err)
}

func payLiveGetList(c *bm.Context) {
	p := new(PayLiveGetListReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1PayLiveSvc.GetList(c, p)
	c.JSON(resp, err)
}

func payLiveClose(c *bm.Context) {
	p := new(PayLiveCloseReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1PayLiveSvc.Close(c, p)
	c.JSON(resp, err)
}

func payLiveOpen(c *bm.Context) {
	p := new(PayLiveOpenReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1PayLiveSvc.Open(c, p)
	c.JSON(resp, err)
}

// RegisterV1PayLiveService Register the blademaster route with middleware map
// midMap is the middleware map, the key is defined in proto
func RegisterV1PayLiveService(e *bm.Engine, svc PayLiveBMServer, midMap map[string]bm.HandlerFunc) {
	v1PayLiveSvc = svc
	e.POST("/xlive/internal/live-admin/v1/payLive/add", payLiveAdd)
	e.POST("/xlive/internal/live-admin/v1/payLive/update", payLiveUpdate)
	e.GET("/xlive/internal/live-admin/v1/payLive/getList", payLiveGetList)
	e.POST("/xlive/internal/live-admin/v1/payLive/close", payLiveClose)
	e.POST("/xlive/internal/live-admin/v1/payLive/open", payLiveOpen)
}

// RegisterPayLiveBMServer Register the blademaster route
func RegisterPayLiveBMServer(e *bm.Engine, server PayLiveBMServer) {
	v1PayLiveSvc = server
	e.POST("/live.liveadmin.v1.PayLive/add", payLiveAdd)
	e.POST("/live.liveadmin.v1.PayLive/update", payLiveUpdate)
	e.GET("/live.liveadmin.v1.PayLive/getList", payLiveGetList)
	e.POST("/live.liveadmin.v1.PayLive/close", payLiveClose)
	e.POST("/live.liveadmin.v1.PayLive/open", payLiveOpen)
}