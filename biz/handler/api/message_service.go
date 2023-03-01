// Code generated by hertz generator.

package api

import (
	"context"
	"douyin/pkg/global"

	"douyin/biz/model/api"
	"douyin/biz/service"
	"douyin/pkg/constant"
	"douyin/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// SendMessage .
// @router /douyin/message/action/ [POST]
func SendMessage(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinMessageActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, &api.DouyinResponse{
			StatusCode: errno.UserRequestParameterError.ErrCode,
			StatusMsg:  err.Error(),
		})
		return
	}

	hlog.Info("handler.message_service.SendMessage Request:", req)
	fromUserID := c.GetUint64(global.Config.JWTConfig.IdentityKey)
	hlog.Info("handler.message_service.SendMessage GetFromUserID:", fromUserID)
	var resp *api.DouyinMessageActionResponse
	if req.ActionType == constant.SendMessageAction {
		resp, err = service.SendMessage(fromUserID, uint64(req.ToUserID), req.Content)
	} else {
		err = errno.UserRequestParameterError
		hlog.Info("handler.message_service.SendMessage err:", err.Error())
	}

	if err != nil {
		errNo := errno.ConvertErr(err)
		c.JSON(consts.StatusOK, &api.DouyinMessageActionResponse{
			StatusCode: errNo.ErrCode,
			StatusMsg:  &errNo.ErrMsg,
		})
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// GetMessageChat .
// @router /douyin/message/chat/ [GET]
func GetMessageChat(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinMessageChatRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, &api.DouyinResponse{
			StatusCode: errno.UserRequestParameterError.ErrCode,
			StatusMsg:  err.Error(),
		})
		return
	}

	hlog.Info("handler.message_service.GetMessageChat Request:", req)
	userID := c.GetUint64(global.Config.JWTConfig.IdentityKey)
	hlog.Info("handler.message_service.GetMessageChat GetUserID:", userID)
	if req.PreMsgTime == nil {
		preMsgTime := int64(0)
		req.PreMsgTime = &preMsgTime
	}
	resp, err := service.GetMessageChat(userID, uint64(req.ToUserID), *req.PreMsgTime)
	if err != nil {
		errNo := errno.ConvertErr(err)
		c.JSON(consts.StatusOK, &api.DouyinMessageChatResponse{
			StatusCode: errNo.ErrCode,
			StatusMsg:  &errNo.ErrMsg,
		})
		return
	}

	c.JSON(consts.StatusOK, resp)
}
