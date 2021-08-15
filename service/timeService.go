package service

import (
	"github.com/formiko/mygoutil"
	"github.com/gin-gonic/gin"
)

func StampToTime(ctx *gin.Context, stamp int64) (time string, err error) {
	return mygoutil.StampToStr(stamp), nil
}
