package player

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/http"
	"player/config"
	"plugin"
	"strings"

	"github.com/gin-gonic/gin"
)

var f plugin.Symbol

func init() {
	plug, err := plugin.Open("plugin.so")
	if err != nil {
		fmt.Println(err)
	}

	f, err = plug.Lookup("Paras")
	if err != nil {
		fmt.Println(err)
	}
}

func Index(c *gin.Context) {
	Ourl, ok := c.GetQuery("url")
	if !ok || Ourl == "" {
		c.HTML(http.StatusOK, "empty.html", nil)
		return
	}
	con, err := config.Read()
	var tt string
	if err != nil {
		goto there
	}

	if con.Paras == 1 {
		url, t, err := f.(func(string) (string, string, error))(Ourl)
		if err != nil {
			goto there
		} else {
			Ourl = url
		}
		tt = t
	}

there:
	url := base64.StdEncoding.EncodeToString([]byte(Ourl))
	h := md5.New()
	h.Write([]byte(Ourl))
	id := hex.EncodeToString(h.Sum(nil))[:8]
	av, _ := c.GetQuery("av")
	sid, _ := c.GetQuery("sid")
	pic, _ := c.GetQuery("pic")
	title, _ := c.GetQuery("title")
	next, _ := c.GetQuery("next")
	user, _ := c.GetQuery("user")
	group, _ := c.GetQuery("group")
	color, _ := c.GetQuery("color")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"av":    av,
		"url":   url,
		"id":    id,
		"sid":   sid,
		"pic":   pic,
		"title": title,
		"next":  next,
		"user":  user,
		"group": group,
		"color": color,
		"mode": func() int {
			if strings.Contains(Ourl, "m3u8") || tt == "hls" {
				return 0
			} else if strings.Contains(Ourl, "flv") || tt == "mp4" {
				return 1
			} else {
				return 1
			}
		}(),
		"tj": tj(c.ClientIP()),
	})
}
