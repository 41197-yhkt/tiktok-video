package service

import (
	"context"
	"fmt"
	"strconv"
	dal "tiktok-video/gen/dal"
	"tiktok-video/gen/dal/model"
	vedio "tiktok-video/kitex_gen/video"
)

// service/publishlist.go
type PublishListService struct {
	ctx context.Context
}

func NewPublishListService(ctx context.Context) *PublishListService {
	return &PublishListService{ctx: ctx}
}

func (s *PublishListService) PublishList(req *vedio.DouyinPublishListRequest) ([]*vedio.Video, error) {
	//vedioDatabase := q.Vedio.WithContext(s.ctx)
	userFavoriteDatabase := q.UserFavorite.WithContext(s.ctx)

	// 先根据 user_id 选出 vedios
	var vedios []*model.Vedio
	var isFavorite bool
	var err error
	dal.DB.WithContext(s.ctx).Where("author_id = ?", req.UserId).Find(&vedios)
	fmt.Println(vedios)

	// 根据 vedio_id 查 Vedio
	res := []*vedio.Video{}
	for _, vd := range vedios {
		// vedio, err := vedioDatabase.FindByID(int64(vd.Id))
		// if err != nil {
		//     panic(err)
		// }

		// 查询点赞数目
		var favoriteCount int64
		dal.DB.WithContext(s.ctx).Where("author_id = ?", vd.Id).Count(&favoriteCount)

		// TODO: 查询评论数
		var commentCount int64
		commentCount, err = strconv.ParseInt(vd.CommentCount, 10, 64)

		// 查询自己是不是也点了赞

		err = userFavoriteDatabase.WithContext(s.ctx).FindByUseridAndVedioid(req.UserId, int64(vd.Id))
		if err != nil {
			isFavorite = false
		} else {
			isFavorite = true
		}

		// 封装
		res = append(res, &vedio.Video{
			Id: int64(vd.Id),
			Author: &vedio.User{
				Id: vd.AuthorId,
			},
			PlayUrl:       vd.PlayUrl,
			CoverUrl:      vd.CoverUrl,
			FavoriteCount: favoriteCount,
			CommentCount:  commentCount,
			IsFavorite:    isFavorite,
			Title:         vd.Title,
		})
	}

	return res, nil
}
