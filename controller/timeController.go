package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myService/service"
	"regexp"
	"strconv"
)

func StampToTime(ctx *gin.Context) {
	stamp := ctx.Query("stamp")
	fmt.Println(stamp)
	if "" == stamp {

	}
	isNum, err := regexp.MatchString(`\d+`, stamp)
	if err != nil {
		ctx.JSON(200, "正则表达式错误")
		return
	}
	if !isNum {
		ctx.JSON(200, "请传入非负整数")
		return
	}
	int64Stamp, err := strconv.ParseInt(stamp, 10, 64)
	if err != nil {
		ctx.JSON(200, fmt.Sprintf("strconv.ParseInt err:[%+v]", err))
		return
	}
	time, err := service.StampToTime(ctx, int64Stamp)
	if err != nil {
		ctx.JSON(200, fmt.Sprintf("service.StampToTime err:[%+v]", err))
		return
	} else {
		ctx.JSON(200, time)
		return
	}
}
