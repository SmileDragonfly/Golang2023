package transport

import (
	"github.com/gin-gonic/gin"
	"gormtest/GormTest/business"
	"gormtest/GormTest/model"
	"gormtest/GormTest/storage"
	"net/http"
)

type gin_transport struct {
	db  storage.Storage
	biz business.Bussiness
}

func NewGinTransport(store storage.Storage, biz business.Bussiness) *gin_transport {
	return &gin_transport{store, biz}
}

func (g gin_transport) HandlerCreateWallet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var wallet model.Wallet
		if err := ctx.ShouldBindJSON(&wallet); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}
		// TODO-preprocess

	}
}
