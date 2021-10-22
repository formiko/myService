package moonController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"github.com/formiko/myService/service/moonService"
)

func Admire(ctx *gin.Context) {
	moon, err := moonService.Admire(ctx)
	if err != nil {
		ctx.JSON(200, fmt.Sprintf("moonService.Admire err:[%+v]", err))
		log.Printf("moonService.Admire err:[%+v]", err)
		return
	} else {
		ctx.JSON(200, moon)
		log.Printf("moon:[%+v]", moon)
		return
	}
}

func ChangeMoon(ctx *gin.Context) {
	log.Printf("moon:[%+v]", ctx.Query("moon"))
	moon := ctx.Query("moon")
	fmt.Println(moon)
	resp, err := moonService.ChangeMoon(ctx, moon)
	if err != nil {
		ctx.JSON(200, fmt.Sprintf("moonService.ChangeMoon err:[%+v]", err))
		log.Printf("moonService.ChangeMoon err:[%+v]", err)
		return
	} else {
		ctx.JSON(200, resp)
		log.Printf("moon:[%+v]", resp)
		return
	}
}