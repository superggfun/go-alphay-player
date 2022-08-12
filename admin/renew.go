package admin

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Renew(c *gin.Context) {
	if act, ok := c.GetQuery("act"); ok {
		if act == "reset" {
			var out bytes.Buffer
			out.WriteString(cc)
			x, err := os.Create("admin.json")
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"err": err,
				})
			}
			out.WriteTo(x)
			x.Close()
		} else if act == "setting" {
			if edit := c.PostForm("edit"); edit == "1" {
				var form dataForm
				if c.ShouldBind(&form) == nil {
					var data = make(map[string]interface{})
					var ads = make(map[string]interface{})
					var set = make(map[string]interface{})
					var pic = make(map[string]interface{})
					var vod = make(map[string]interface{})
					var pause = make(map[string]interface{})
					if form.Danmuon != "" {
						data["danmuon"] = form.Danmuon
					}
					data["color"] = form.Color
					data["logo"] = form.Logo
					data["trytime"] = form.Trytime
					data["waittime"] = form.Waittime
					data["sendtime"] = form.Sendtime
					data["dmrule"] = form.Dmrule
					data["pbgjz"] = form.Pbgjz
					data["ads"] = ads
					if form.AdsState != "" {
						ads["state"] = form.AdsState
					}
					ads["set"] = set
					set["state"] = form.AdsSetState
					set["group"] = form.AdsSetGroup
					set["pic"] = pic
					pic["time"] = form.AdsSetPicTime
					pic["img"] = form.AdsSetPicImg
					pic["link"] = form.AdsSetPicLink
					set["vod"] = vod
					vod["url"] = form.AdsSetVodUrl
					vod["link"] = form.AdsSetVodLink
					ads["pause"] = pause
					if form.AdsPauseState != "" {
						pause["state"] = form.AdsPauseState
					}
					pause["pic"] = form.AdsPausePic
					pause["link"] = form.AdsPauseLink
					s, err := json.MarshalIndent(&data, "", "\t")
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{
							"err": err,
						})
					}
					var out bytes.Buffer
					out.Write(s)
					x, err := os.Create("config.json")

					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{
							"err": err,
						})
					}
					out.WriteTo(x)
					x.Close()
				}
			}
		}
	}

}
