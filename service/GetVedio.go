package service

import (
	"context"
	"fmt"
	"log"
	"strconv"
	dal "tiktok-video/gen/dal"
	"tiktok-video/gen/dal/model"
	vedio "tiktok-video/kitex_gen/video"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// service/GetVedio.go
type GetVedioService struct {
	ctx context.Context
}

func NewGetVedioService(ctx context.Context) *GetVedioService {
	return &GetVedioService{ctx: ctx}
}

func (s *GetVedioService) GetVedio(req *vedio.GetVedioRequest) (*vedio.Video, error) {
	//vedioDatabase := q.Vedio.WithContext(s.ctx)
	//client为阿里云oss对象
	client, err := oss.New("oss-cn-beijing.aliyuncs.com", "LTAI5tGdrFczu9cP7RX8LgrC", "I0P6eEUAk740O5jM1VLbvfePs5yGAf")
	if err != nil {
		log.Panic(err)
	}
	//选择视频bucket
	bucket, err := client.Bucket("video-bucket0")
	if err != nil {
		log.Panic(err)
	}

	userFavoriteDatabase := q.UserFavorite.WithContext(s.ctx)
	//userDatabase := q.User.WithContext(s.ctx)

	// 先根据 user_id 选出 vedios
	var vedios *model.Vedio
	var isFavorite bool
	var res *vedio.Video

	dal.DB.WithContext(s.ctx).Where("Id = ?", req.TargetVedioId).Find(&vedios)
	fmt.Println(vedios)

	// vedio, err := vedioDatabase.FindByID(int64(vd.Id))
	// if err != nil {
	//     panic(err)
	// }
	var user *vedio.User
	dal.DB.WithContext(s.ctx).Select("name").Where("id = ?", vedios.AuthorId).Find(&user.Name)

	// 查询点赞数目
	var favoriteCount int64
	dal.DB.WithContext(s.ctx).Where("author_id = ?", vedios.Id).Count(&favoriteCount)

	// TODO: 查询评论数
	var commentCount int64
	commentCount, err = strconv.ParseInt(vedios.CommentCount, 10, 64)

	// 查询自己是不是也点了赞

	err = userFavoriteDatabase.WithContext(s.ctx).FindByUseridAndVedioid(req.UserId, req.TargetVedioId)
	if err != nil {
		isFavorite = false
	} else {
		isFavorite = true
	}

	playurl, err := bucket.SignURL(vedios.Title, oss.HTTPGet, 30)
	if err != nil {
		panic(err)
	}

	coverurl, err := bucket.SignURL("cover.png", oss.HTTPGet, 30)
	if err != nil {
		panic(err)
	}

	// 封装
	res = &vedio.Video{
		Id:            req.TargetVedioId,
		Author:        user,
		PlayUrl:       playurl,
		CoverUrl:      coverurl,
		FavoriteCount: favoriteCount,
		CommentCount:  commentCount,
		IsFavorite:    isFavorite,
		Title:         vedios.Title,
	}

	return res, nil
}
