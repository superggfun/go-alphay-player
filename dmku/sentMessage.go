package dmku

import (
	"fmt"
	"net/http"
	"player/config"
	"player/dmku/database"
	"time"

	"github.com/gin-gonic/gin"
)

type list struct {
	Author    string  `json:"author"`
	Color     string  `json:"color"`
	Player    string  `json:"player"`
	Size      string  `json:"size"`
	Text      string  `json:"text"`
	Videotime float32 `json:"time"`
	Typee     string  `json:"type"`
}

func SentMessage(c *gin.Context) {
	ip := c.ClientIP()
	var l list
	c.BindJSON(&l)
	frequency, err := m.Qy_frequency(ip)
	if err != nil {
		fmt.Println(err)
	}
	if frequency.C == 0 {
		err := m.Ins_danmaku(database.List{l.Player, 0, l.Typee, l.Text, l.Color, l.Size, l.Videotime, ip, 0})
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    -3,
				"danmuku": err,
			})
		}

		err = m.Ins_frequency(database.Ip{ip, 1, 0})
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    -3,
				"danmuku": err,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    23,
			"danmuku": "true",
		})
	} else {
		conf, err := config.Read()
		if err != nil {
			panic(err)
		}
		if (frequency.Time + conf.Limit_time) > time.Now().Unix() {
			if frequency.C < conf.Limit_requence {
				err := m.Ins_danmaku(database.List{l.Player, 0, l.Typee, l.Text, l.Color, l.Size, l.Videotime, ip, 0})
				if err != nil {
					c.JSON(http.StatusOK, gin.H{
						"code":    -3,
						"danmuku": err,
					})
				}
				m.Upd_frequency(ip, false)
				c.JSON(http.StatusOK, gin.H{
					"code":    23,
					"danmuku": "true",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code":    -2,
					"danmuku": "你tm发送的太频繁了,请问你单身几年了？",
				})
			}
		} else {
			err := m.Ins_danmaku(database.List{l.Player, 0, l.Typee, l.Text, l.Color, l.Size, l.Videotime, ip, 0})
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code":    -3,
					"danmuku": err,
				})
			}
			m.Upd_frequency(ip, true)
			c.JSON(http.StatusOK, gin.H{
				"code":    23,
				"danmuku": "true",
			})
		}
	}
}
