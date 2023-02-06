package main

import (
	"context"
	video "tiktok-video/kitex_gen/video"
	"tiktok-video/pack"
	service "tiktok-video/service"

	errno "github.com/41197-yhkt/pkg/errno"
)

// DouyinServiceImpl implements the last service interface defined in the IDL.
type DouyinServiceImpl struct{}

// DouyinPublishActionMethod implements the DouyinServiceImpl interface.
func (s *DouyinServiceImpl) DouyinPublishActionMethod(ctx context.Context, req *video.DouyinPublishActionRequest) (resp *video.DouyinPublishActionResponse, err error) {
	// TODO: Your code here...
	resp = new(video.DouyinPublishActionResponse)

	//检验参数规范性
	if req.Author < 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.InvalidParams)
		return resp, nil
	}

	//调用服务
	err = service.NewPublishActionService(ctx).PublishAction(req)

	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// DouyinPublishListMethod implements the DouyinServiceImpl interface.
func (s *DouyinServiceImpl) DouyinPublishListMethod(ctx context.Context, req *video.DouyinPublishListRequest) (resp *video.DouyinPublishListResponse, err error) {
	// TODO: Your code here...
	resp = new(video.DouyinPublishListResponse)

	//检验参数规范性
	if req.UserId < 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.InvalidParams)
		return resp, nil
	}

	//调用服务
	resp.VideoList, err = service.NewPublishListService(ctx).PublishList(req)

	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
