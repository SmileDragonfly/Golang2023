package business

import "gormtest/GormTest/storage"

type mainBiz struct {
	createBiz
	deleteBiz
	readBiz
	updateBiz
	store storage.Storage
}

func NewMainBiz(storage storage.Storage) *mainBiz {
	return &mainBiz{NewCreateBiz(storage),
		NewDeleteBiz(storage),
		NewReadBiz(storage),
		NewUpdateBiz(storage),
		storage}
}
