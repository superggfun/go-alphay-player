package admin

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"player/config"

	"github.com/gin-gonic/gin"
)

type loginForm struct {
	User     string `form:"username" binding:"required"`
	Password string `form:"pass" binding:"required"`
}

func Wall(c *gin.Context) {
	if _, err := c.Request.Cookie("ay_player"); err != nil {
		if _, ok := c.GetQuery("l"); ok {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"ti": "用户名或密码错误！"})
		} else {
			c.HTML(http.StatusOK, "login.html", nil)
		}
		c.Abort()
	} else {
		c.Next()
	}
}

func Login(c *gin.Context) {
	var form loginForm
	if c.ShouldBind(&form) == nil {
		con, err := config.Read()
		if err != nil {
			panic(err)
		}
		if form.User == con.Username && form.Password == con.Password {
			h := md5.New()
			h.Write([]byte(con.Username + con.Password))
			h.Write([]byte(hex.EncodeToString(h.Sum(nil)) + "#"))
			uid := hex.EncodeToString(h.Sum(nil))
			c.SetCookie("ay_player", uid, con.Duration, "/", con.Domain, false, true)
			c.Redirect(http.StatusFound, "/admin")
		} else {
			c.Redirect(http.StatusFound, "/admin?l")
		}
	}
}
