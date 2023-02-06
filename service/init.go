package service

import (
	dal "tiktok-video/gen/dal"
	query "tiktok-video/gen/dal/query"
)

var q *query.Query

func Init() {
	q = query.Use(dal.DB.Debug())
}
