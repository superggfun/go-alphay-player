package admin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"player/config"

	"github.com/gin-gonic/gin"
)

func Admin(c *gin.Context) {
	f, err := ioutil.ReadFile("admin.json")
	if err != nil {
		fmt.Print(err)
	}

	var d data
	err = json.Unmarshal(f, &d)
	if err != nil {
		fmt.Println(err)
	}

	c.HTML(http.StatusOK, "admin.html", &d)

}

func Referer(c *gin.Context) {
	con, err := config.Read()
	if err != nil {
		panic(err)
	}
	if len(con.Allow_url) == 0 {
		c.Next()
	} else {
		for k, v := range c.Request.Header {
			if k == "Referer" {
				for _, vv := range con.Allow_url {
					if v[0] == vv {
						c.Next()
						break
					}
				}
			}
		}
		c.JSON(http.StatusForbidden, gin.H{
			"code": 403,
			"msg":  "Permission Denied!",
		})
		c.Abort()
	}

}
