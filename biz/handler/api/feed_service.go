// Code generated by hertz generator.

package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	api "douyin/biz/model/api"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// GetFeed .
// @router /douyin/feed/ [GET]
func GetFeed(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinFeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	hlog.Info(*req.LatestTime)

	var followCount int64 = 3
	var followerCount int64 = 3
	u := &api.User{
		ID:            1,
		Name:          "name",
		FollowCount:   &followCount,
		FollowerCount: &followerCount,
		IsFollow:      false,
		Avatar:        "https://picture-bucket-01.oss-cn-beijing.aliyuncs.com/DouYin/cover/cover01.png",
	}

	videos := make([]*api.Video, 0)
	videos = append(videos, &api.Video{
		ID:            2,
		Author:        u,
		PlayURL:       "https://picture-bucket-01.oss-cn-beijing.aliyuncs.com/DouYin/video/video01.mp4",
		CoverURL:      "https://picture-bucket-01.oss-cn-beijing.aliyuncs.com/DouYin/cover/cover01.png",
		FavoriteCount: 2,
		CommentCount:  2,
		IsFavorite:    false,
		Title:         "video01",
	})
	var nextTime int64 = 100000
	resp := &api.DouyinFeedResponse{
		StatusCode: 0,
		StatusMsg:  nil,
		VideoList:  videos,
		NextTime:   &nextTime,
	}

	c.JSON(consts.StatusOK, resp)
}
