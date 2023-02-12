// Code generated by hertz generator.

package api

import (
	"context"
	"douyin/biz/service"
	"douyin/constant"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	api "douyin/biz/model/api"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// CommentAction .
// @router /douyin/comment/action/ [POST]
func CommentAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinCommentActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	hlog.Infof("%#v comment action", req)

	userID := c.GetInt64(constant.IdentityKey)
	hlog.Info(userID)

	resp := service.CommentAction(&req)

	c.JSON(consts.StatusOK, resp)
}

// GetCommentList .
// @router /douyin/comment/list/ [GET]
func GetCommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinCommentListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	hlog.Infof("%#v", req)

	resp := service.CommentList(&req)

	c.JSON(consts.StatusOK, resp)
}
