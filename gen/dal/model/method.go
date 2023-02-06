package model

import (
	"gorm.io/gen"
)

type UserFavoriteMethod interface {
	// //sql(select vedio_id from @@table where user_id = @userId)
	// FindByUserid(userId int64) (gen.T, error)

	// //sql(select user_id from @@table where vedio_id = @vedioId)
	// FindByVedioid(vedioId int64) (gen.T, error)

	//sql(select * from @@table where vedio_id = @vedioId and user_id = @userId)
	FindByUseridAndVedioid(userId, vedioId int64) error
}

type UserMethod interface {
	//where(id=@id)
	FindByID(id int64) (gen.T, error)
}

type VedioMethod interface {
	//where(id=@id)
	FindByID(id int64) (gen.T, error)
	//sql(select * from @@table where AuthorId = @Authorid)
	FindByAuthorId(Authorid int) (gen.T, error)
}
