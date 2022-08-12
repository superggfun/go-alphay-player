package main

import (
	"fmt"
	"log"
	"net/http"
	"player/admin"
	"player/config"
	"player/dmku"
	"player/player"
	"strconv"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("---AlphaY 播放器，永久免费---")
	fmt.Println("---开源地址 <https://github.com/superggfun/go-alphay-player> ---")
}

func main() {

	r := gin.Default()

	r.LoadHTMLGlob("./static/*")
	r.StaticFS("/assets", http.Dir("./assets"))

	con, err := config.Read()
	if err != nil {
		panic(err)
	}

	p := r.Group(con.Player)
	{
		p.Use(admin.Referer)
		p.GET("", player.Index)

	}

	r.GET("/api", player.Api)

	r.POST(con.Admin+"/login", admin.Login)

	ad := r.Group(con.Admin)
	{
		ad.Use(admin.Wall)
		ad.GET("", admin.Admin)
		ad.POST("/renew", admin.Renew)
	}

	dm := r.Group("/dmku")
	{
		dm.GET("/", dmku.ShowMessage)
		dm.POST("/", dmku.SentMessage)
	}
	log.Println("---AlphaY 播放器，永久免费---")
	log.Println("---开源地址 <https://github.com/superggfun/go-alphay-player> ---")
	r.Run(":" + strconv.Itoa(con.Port))
}
