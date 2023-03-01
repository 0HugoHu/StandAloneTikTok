// Code generated by hertz generator.

package api

import (
	"context"
	"douyin/pkg/global"

	"douyin/biz/model/api"
	"douyin/biz/service"
	"douyin/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Register .
// @router /douyin/user/register/ [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinUserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, &api.DouyinResponse{
			StatusCode: errno.UserRequestParameterError.ErrCode,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 用户信息就不用打印到日志里了
	resp, err := service.Register(req.Username, req.Password)
	if err != nil {
		errNo := errno.ConvertErr(err)
		c.JSON(consts.StatusOK, &api.DouyinUserRegisterResponse{
			StatusCode: errNo.ErrCode,
			StatusMsg:  &errNo.ErrMsg,
		})
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// Login .
// @router /douyin/user/login/ [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinUserLoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, &api.DouyinResponse{
			StatusCode: errno.UserRequestParameterError.ErrCode,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 用户信息就不用打印到日志里了
	resp, err := service.Login(req.Username, req.Password)
	if err != nil {
		errNo := errno.ConvertErr(err)
		c.JSON(consts.StatusOK, &api.DouyinUserRegisterResponse{
			StatusCode: errNo.ErrCode,
			StatusMsg:  &errNo.ErrMsg,
		})
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// GetUserInfo .
// @router /douyin/user/ [GET]
func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, &api.DouyinResponse{
			StatusCode: errno.UserRequestParameterError.ErrCode,
			StatusMsg:  err.Error(),
		})
		return
	}

	hlog.Info("handler.user_service.GetUserInfo Request:", req)
	userID := c.GetUint64(global.Config.JWTConfig.IdentityKey)
	hlog.Info("handler.user_service.GetUserInfo GetUserID:", userID)
	resp, err := service.GetUserInfo(userID, uint64(req.UserID))
	if err != nil {
		errNo := errno.ConvertErr(err)
		c.JSON(consts.StatusOK, &api.DouyinUserResponse{
			StatusCode: errNo.ErrCode,
			StatusMsg:  &errNo.ErrMsg,
		})
		return
	}

	c.JSON(consts.StatusOK, resp)
}
